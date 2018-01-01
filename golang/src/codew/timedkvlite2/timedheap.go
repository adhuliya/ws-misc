package timedkvlite

import (
  "container/heap"
  "time"
  //"log"
)

const (
    MAX_SLEEP   = time.Millisecond << 8 // 256 millisecond
)

/* runHeap() goroutine initializes the heap for operation,
it starts the helper goroutine runTimeGc(). It only adds the
entries to the heap. All deletions are handled by runTimeGc()
*/
func (kv *TimedKv) runHeap() {
  kv.initHeap()

  stop := make(chan int, 1)
  go kv.runTimeGc(stop)

  for item := range kv.heapAdd {
    kv.heapMutex.Lock()

      heap.Push(&kv.heap, item)

    kv.heapMutex.Unlock()
  }

  stop <- 1
}

// Initializes the heap to be ready to start operations.
func (kv *TimedKv) initHeap() {
  // add all initialization code here
  heap.Init(&kv.heap)
}

// runTimeGc() runs as a goroutine, which deletes the keys from the kv that have expired. It deletes the top entry if its time has expired, and it sleeps for sometime if it has nothing to do, then wakesup to check again.
func (kv *TimedKv) runTimeGc(stop <-chan int) {
  var now int64
  var heapLen uint64
  var isItem = false
  var item heapItem
  var tmpDelTime int64
  var sleepDuration time.Duration = 0

  for {
    select {
    case <-stop:
      return
    default:
    }

    now = time.Now().UnixNano()
    isItem = false

    kv.heapMutex.Lock()

      heapLen = uint64(kv.heap.Len())

      if heapLen > 0 {
        tmpDelTime = kv.heap[heapLen-1].delTime
        if tmpDelTime <= now {
          isItem = true
          item = heap.Pop(&kv.heap).(heapItem)
        } else {
          sleepDuration = time.Duration(tmpDelTime - now)
          if sleepDuration > MAX_SLEEP {
            sleepDuration = MAX_SLEEP
          }
        }
      } else {
        sleepDuration = MAX_SLEEP
      }

    kv.heapMutex.Unlock()

    if sleepDuration < 0 {
      sleepDuration = 0
    }

    if isItem {
      value, isPresent := kv.get(item.key)
      if isPresent {
        kv.Delete(item.key)
        kv.Done <- value // raise alarm that time expired
      }
    } else {
      // No item found!
      // consumes less amount of CPU this way
      if sleepDuration > 0 {
        time.Sleep(time.Duration(sleepDuration))
      }
    }
  } // for {}
}


