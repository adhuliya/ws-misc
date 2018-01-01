package timedkvlite

import (
  "sync"
  "time"
  //"log"
)

/* TimedKv creates the kv store for use. It returns the kv object.
 */
func Make() *TimedKv {
  kv := &TimedKv{
    kv        : make(map[uint64]Value),
    heap      : make(timedHeap, 0, 1000),
    heapAdd   : make(chan heapItem, 100),
    kvMutex   : &sync.RWMutex{},
    heapMutex : &sync.Mutex{},
    key       : 0,
    maxKey    : 1<<64 - 1,
    maxDuration : int64(time.Hour * 24 * 365 * 2), // ~ 2 years
    minDuration : int64(time.Second * 5),
    maxDurationYet  : 0,
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

// /* Makevalue() retuns an initialized Value, ready to be fed into
// the kv store.
// */
// func MakeValue(value interface{}, exptime time.Duration) Value {
//   if exptime <= 0 {
//     // if exptime is zero or below, the key-value pair is
//     // stored independent of time (or could be changed to few hours)
//     return Value{
//       value:   value,
//       exptime: 0,
//     }
//   }
// 
//   return Value{
//     value:   value,
//     exptime: time.Now().Add(exptime).UnixNano(),
//   }
// }

/* Add() sets a new kv pair in the kv store. It overwrites the value
if the key already exists.
*/
func (kv *TimedKv) Add(value uint64, timeToLive int64) uint64 {

  timeToLive = kv.limitDuration(timeToLive)
  if timeToLive > kv.maxDurationYet {
    kv.maxDurationYet = timeToLive
  }

  now := time.Now().UnixNano()
  val := Value {
    key     : 0,
    value   : value,
    addTime : now,
    delTime : now + int64(timeToLive),
  }


  kv.kvMutex.Lock()

    val.key         = kv.nextKey() // allocate unique key

    kv.kv[val.key]  = val

  kv.kvMutex.Unlock()

  //log.Println("added", value)

  if val.delTime > 0 {
    kv.heapAdd <- heapItem {
      key : val.key,
      delTime : val.delTime,
    }
  }

  return val.key
}

// Set timeToLive between max and min
func (kv *TimedKv) limitDuration(timeToLive int64) int64 {
  if timeToLive < kv.minDuration {
    timeToLive = kv.minDuration
  }
  if timeToLive > kv.maxDuration {
    timeToLive = kv.maxDuration
  }

  return timeToLive
}

// Get() returns the value if present, and true/false if the value
//returned is valid/invalid respectively.
func (kv *TimedKv) Get(key uint64) (Value, bool) {
  kv.kvMutex.RLock()

  value, isPresent := kv.kv[key]

  kv.kvMutex.RUnlock()

  if isPresent {
    if value.delTime < time.Now().UnixNano() {
      isPresent = false
      // the runHeap() goroutine will automatically call
      // the delete on this kv, so no need to delete here
    }
  }

  return value, isPresent
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

/* Modify sets a new value to an already present kv pair, if the pair is
present and its version is same as given during the function call.
*/
func (kv *TimedKv) Modify(key uint64, value uint64, timeToLive int64,
    modifyVal bool, modifyDuration bool) byte {
  var status byte = FAILURE

  timeToLive = kv.limitDuration(timeToLive)
  now     := time.Now().UnixNano()

  kv.kvMutex.Lock()

    val, isPresent := kv.kv[key] // READ
    if isPresent {
      if modifyDuration {
        val.delTime  = now + int64(timeToLive)
      }
      if modifyVal {
        val.value    = value
      }

      kv.kv[key] = val // SET (reset/update)

      if modifyDuration {
        kv.heapAdd <- heapItem {
          key     : val.key,
          delTime : val.delTime,
        }
      }

      status = SUCCESS
    }

  kv.kvMutex.Unlock()

  return status
}

// Cad() (Compare and delete) deletes a kv pair, if the pair is
// present and its delTime is same as given delTime.
func (kv *TimedKv) Cad(key uint64, delTime int64) byte {
  var status byte

  status = FAILURE

  kv.kvMutex.Lock()

    val, isPresent := kv.kv[key] // READ
    if isPresent {
      // if delTime is not same, then the value must have been
      // updated; and so another heap entry must exist with the
      // updated delTime
      if val.delTime == delTime {
        delete(kv.kv, key)
        status = SUCCESS
      }
    }
    // No deletion in heap as in time it will automatically
    // get deleted. (It becomes complex to del in heap)

  kv.kvMutex.Unlock()

  return status
}


