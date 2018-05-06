package timedkvlite

import (
  "container/heap"
  "time"
//  "fmt"
)

const (
    MAX_SLEEP   = time.Millisecond << 6 // 64 millisecond
    MIN_SLEEP   = time.Microsecond << 0 //  1 microsecond
    HEAP_BATCH_DEL_SIZE = 1 << 12
    HEAP_BATCH_ADD_SIZE = 1 << 12
    HEAP_ADD_PAITENCE = time.Millisecond * 4
)

/* runHeap() goroutine initializes the heap for operation,
it starts the helper goroutine runTimeGc(). It only adds the
entries to the heap. All deletions are handled by runTimeGc()
*/
func (kv *TimedKv) runHeap() {
  kv.initHeap()

  var count = 0
  var open = true
  var item heapItem

  var items []heapItem
  items = make([]heapItem, HEAP_BATCH_ADD_SIZE)

  stop := make(chan int, 1)
  go kv.runTimeGc(stop)

  ticker := time.NewTicker(HEAP_ADD_PAITENCE)
  defer ticker.Stop()

  mainLoop:
  for {
    select {
      case item, open = <-kv.heapAdd:
        if !open {
          break mainLoop
        }

        items[count] = item
        count += 1
        if count >= HEAP_BATCH_ADD_SIZE {
          kv.heapMutex.Lock()

            for i := 0; i < count; i++ {
              heap.Push(&kv.heap, items[i])
            }

          kv.heapMutex.Unlock()
          count = 0
        }
        break

      case <-ticker.C:
        if count > 0 {

          kv.heapMutex.Lock()

            for i := 0; i < count; i++ {
              heap.Push(&kv.heap, items[i])
            }

          kv.heapMutex.Unlock()
          count = 0
        }
        break
    }
  }

  stop <- 1
}

// Initializes the heap to be ready to start operations.
func (kv *TimedKv) initHeap() {
  // add all initialization code here
  heap.Init(&kv.heap)
}

// runTimeGc() runs as a goroutine, which deletes the keys from the kv that have expired.
// It deletes the top entry if its time has expired.
// It sleeps for sometime if it has nothing to do.
func (kv *TimedKv) runTimeGc(stop <-chan int) {
  var now int64
  var heapLen uint64
  var topDelTime int64
  var sleepDuration time.Duration = 0
  var items []heapItem
  var itemCount uint
  var i uint

  items = make([]heapItem, HEAP_BATCH_DEL_SIZE)


  for {
    select {
    case <-stop:
      return
    default:
    }

    itemCount = 0

    kv.heapMutex.Lock()

      collectItems:
      for {
        heapLen = uint64(kv.heap.Len())
        if heapLen == 0 {
            break collectItems
        }

        now = time.Now().UnixNano()

        if heapLen > 0 {
          topDelTime = kv.heap[heapLen-1].delTime
          if topDelTime <= now && itemCount < HEAP_BATCH_DEL_SIZE {
            items[itemCount] = heap.Pop(&kv.heap).(heapItem)
            itemCount += 1
            //fmt.Println("Heap Deleted", items[itemCount-1].key)
          } else {
            sleepDuration = time.Duration(topDelTime - now)
            break collectItems // jump out of loop
          }
        }
      }

    kv.heapMutex.Unlock()

    // set lower limit to sleepDuration
    if sleepDuration < 0 {
      sleepDuration = 0
    }

    // set upper limit to sleepDuration
    if sleepDuration > MAX_SLEEP {
      sleepDuration = MAX_SLEEP
    }

    if itemCount > 0 {
      kv.kvMutex.Lock()

        for i = 0; i < itemCount; i++ {
          delTime, isPresent := kv.get(items[i].key)
          if isPresent {
            if delTime == items[i].delTime {
              kv.del(items[i].key)
              kv.Done <- items[i].key // ring the alarm 'time expired'
            }
          }
        }

      kv.kvMutex.Unlock()
    } else {
      // No items found!
      // Consume less CPU by sleeping for a while.
      if sleepDuration > MIN_SLEEP {
        time.Sleep(sleepDuration)
      }
    }
  } // for {}
}


