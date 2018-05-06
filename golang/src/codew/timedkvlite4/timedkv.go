package timedkvlite

import (
  "sync"
  "time"
  "errors"
//  "fmt"
)

const (
    HEAP_CHAN_SIZE    = 1 << 10
    DONE_CHAN_SIZE    = 1 << 10
)

/* TimedKv creates and prepares the kv store for use.
 * It returns the pointer to the kv object.
 */
func Make() *TimedKv {
  kv := &TimedKv{
    kv          : make(map[uint64]int64),
    heap        : make(timedHeap, 0, 1024),
    heapAdd     : make(chan heapItem, HEAP_CHAN_SIZE),
    kvMutex     : &sync.Mutex{},
    heapMutex   : &sync.Mutex{},
    Done        : make(chan uint64, DONE_CHAN_SIZE),
  }

  go kv.runHeap()

  return kv
}

/* destroykv() stops the kv store. Any further use of the kv store will result in an error.
 */
func (kv *TimedKv) Destroy() {
  // closing the channel automatically ends the
  // runHeap() goroutine eventually.
  //fmt.Println("HERE?")
  close(kv.heapAdd)
}

/* Add() sets a new kv pair in the kv store.
It returns error if key is already present.
*/
func (kv *TimedKv) Add(key uint64, timeToLive time.Duration) (uint64, error) {

  var isPresentKv = false

  if !kv.isInLimits(timeToLive) {
      return 0, errors.New("Duration not in limits.")
  }

  now := time.Now().UnixNano()
  delTime := now + int64(timeToLive)

  kv.kvMutex.Lock()
    _, isPresentKv = kv.kv[key]
    if !isPresentKv {
        kv.kv[key] = delTime
    }
  kv.kvMutex.Unlock()

  if isPresentKv {
      return 0, errors.New("Key already present.")
  }

  kv.heapAdd <- heapItem {
      key       : key,
      delTime   : delTime,
  }

  return key, nil
}


// Check timeToLive limits
func (kv *TimedKv) isInLimits(timeToLive time.Duration) bool {
  if timeToLive < 0 {
    return false
  }

  return true
}

// Get() returns the value if present, and true/false if the value
//returned is valid/invalid respectively.
func (kv *TimedKv) Get(key uint64) (int64, bool) {
  kv.kvMutex.Lock()

    delTime, isPresent := kv.kv[key]

  kv.kvMutex.Unlock()

  if isPresent {
    if delTime < time.Now().UnixNano() {
      isPresent = false
      // the runHeap() goroutine will automatically call
      // the delete on this kv, so no need to delete here
    }
  }

  return delTime, isPresent
}

// get() returns the value if present, even if time has expired.
// for internal/testing use only.
// Expects kv to be locked already.
func (kv *TimedKv) get(key uint64) (int64, bool) {

  mapVal, isPresent := kv.kv[key]

  return mapVal, isPresent
}

/* Delete() deletes a kv pair. If the key is not in the kv store, it
returns false, otherwise it removes the kv pair and returns true.
*/
func (kv *TimedKv) Delete(key uint64) bool {
  kv.kvMutex.Lock()

    _, isPresent := kv.kv[key]
    if isPresent {
      delete(kv.kv, key)
    }
    // No deletion in heap as in time it will automatically
    // get deleted.
    // It becomes inefficient to delete in heap explicitly.

  kv.kvMutex.Unlock()

  return isPresent
}

/* del() deletes a kv pair.
 * If the key is not in the kv store, it returns false, otherwise it removes the kv pair and returns true.
 * It expects kv to be already locked.
*/
func (kv *TimedKv) del(key uint64) bool {
    _, isPresent := kv.kv[key]
    if isPresent {
      delete(kv.kv, key)
    }
    // No deletion in heap as in time it will automatically
    // get deleted.
    // It becomes inefficient to delete in heap explicitly.

  return isPresent
}


