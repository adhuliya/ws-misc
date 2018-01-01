package timedkvlite

import (
  "sync"
  "time"
  "errors"
  //"log"
)

const (
    HEAP_CHAN_SIZE  = 1024
    DONE_CHAN_SIZE  = 1024
    MAX_DURATION_DEFAULT = time.Hour << 5   // 32 hours
    MIN_DURATION_DEFAULT = time.Second << 3 // 8 seconds
    MODIFY_VALUE          = 1
    MODIFY_DURATION       = 2
    MODIFY_VALUE_DURATION = 3
)

/* TimedKv creates the kv store for use. It returns the kv object.
 */
func Make() *TimedKv {
  kv := &TimedKv{
    kv        : make(map[uint64]MapValue),
    heap      : make(timedHeap, 0, 1024),
    heapAdd   : make(chan heapItem, HEAP_CHAN_SIZE),
    kvMutex   : &sync.RWMutex{},
    heapMutex : &sync.Mutex{},
    Done      : make(chan Value, DONE_CHAN_SIZE),
    maxDuration : MAX_DURATION_DEFAULT,
    minDuration : MIN_DURATION_DEFAULT,
    maxDurationYet  : time.Duration(0),
  }

  go kv.runHeap()

  return kv
}

/* destroykv() stops the kv store. Any further use of the kv store will result in an error.
 */
func (kv *TimedKv) Destroy() {
  // closing the channel automatically ends the
  // runHeap() goroutine eventually.
  close(kv.heapAdd)
}

/* Add() sets a new kv pair in the kv store. It overwrites the value
if the key already exists.
*/
func (kv *TimedKv) Add(value Value, timeToLive time.Duration) (uint64, error) {

  if !kv.isInLimits(timeToLive) {
      return 0, errors.New("Duration not in limits.")
  }

  // for stats
  if timeToLive > kv.maxDurationYet {
    kv.maxDurationYet = timeToLive
  }

  now := time.Now().UnixNano()
  mapVal := MapValue {
    value   : value,
    addTime : now,
    delTime : now + int64(timeToLive),
  }


  kv.kvMutex.Lock()

    key        := kv.nextKey() // allocate unique key

    kv.kv[key]  = mapVal

  kv.kvMutex.Unlock()

  //log.Println("added", value)

  kv.heapAdd <- heapItem {
      key       : key,
      delTime   : mapVal.delTime,
  }

  return key, nil
}

// Check timeToLive limits
func (kv *TimedKv) isInLimits(timeToLive time.Duration) bool {
  if timeToLive < kv.minDuration {
    return false
  }
  if timeToLive > kv.maxDuration {
    return false
  }

  return true
}

// Get() returns the value if present, and true/false if the value
//returned is valid/invalid respectively.
func (kv *TimedKv) Get(key uint64) (Value, bool) {
  kv.kvMutex.RLock()

      mapVal, isPresent := kv.kv[key]

  kv.kvMutex.RUnlock()

  if isPresent {
    if mapVal.delTime < time.Now().UnixNano() {
      isPresent = false
      // the runHeap() goroutine will automatically call
      // the delete on this kv, so no need to delete here
    }
  }

  return mapVal.value, isPresent
}

// get() returns the value if present, even if time has expired.
// for internal use only.
func (kv *TimedKv) get(key uint64) (Value, bool) {
  kv.kvMutex.RLock()

      mapVal, isPresent := kv.kv[key]

  kv.kvMutex.RUnlock()

  return mapVal.value, isPresent
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
    // get deleted. (It becomes complex to del in heap)

  kv.kvMutex.Unlock()

  return isPresent
}

// Modify sets a new value to an already present kv pair, if the pair is present, it is modified. It returns false if the key doesnot exist.
func (kv *TimedKv) Modify(key uint64, value Value, timeToLive time.Duration, modify int) (bool, error) {
  success := true

  if modify & MODIFY_DURATION != 0 && !kv.isInLimits(timeToLive) {
      return false, errors.New("Duration not in limits.")
  }

  now     := time.Now().UnixNano()

  kv.kvMutex.Lock()

    mapVal, isPresent := kv.kv[key] // READ
    if isPresent {
      if modify & MODIFY_DURATION != 0 {
        mapVal.delTime  = now + int64(timeToLive)
      }
      if modify & MODIFY_VALUE != 0 {
        mapVal.value    = value
      }

      kv.kv[key] = mapVal // SET (reset/update)

      if modify & MODIFY_DURATION != 0 {
        kv.heapAdd <- heapItem {
          key     : key,
          delTime : mapVal.delTime,
        }
      }

    } else {
        success = false
    }

  kv.kvMutex.Unlock()

  return success, nil
}


