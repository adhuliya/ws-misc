package timedkvlite

import (
  "sync"
  "time"
)

const (
    // max key that can be used in the map
    MAX_KEY uint64 = 1<<64 - 1
)

type TimedKv struct {
  kv            map[uint64]MapValue
  heap          timedHeap
  heapAdd       chan heapItem
  kvMutex       *sync.RWMutex
  heapMutex     *sync.Mutex

  keyCounter    uint64

  Done          chan Value

  maxDuration   time.Duration
  minDuration   time.Duration

  // for stats
  maxDurationYet  time.Duration
}

// Only this struct is returned on time expiry,
// through the TimedKv.Done channel.
type Value struct {
  // Deliberately 16 bytes for efficiency (mem aligned).

  // `value` is meaningful only to the user of this library.
  // `value` is returned on time expiry.
  value     uint64

  // `action` is meaningful only to the user of this library.
  // `action` is returned on time expiry.
  action    uint64
}

type MapValue struct {
  value     Value

  // `addTime` the time when this value is added.
  addTime   int64 // (nanosec) absolute time

  // `delTime` is not returned on time expiry.
  delTime   int64 // (nanosec) absolute time when to delete
}

// No update() allowed in this heap
type timedHeap []heapItem

type heapItem struct {
  // deliberately without the 'index' member
  // no updates or removal allowed; let it get garbage collected
  // It is only 16 bytes and will eventually get removed in time.
  key     uint64   // Related to TimedKv.kv[key]

  delTime int64    // absolute time at which to delete this
}

func (self *MapValue) GetDelTime() int64 {
  return self.delTime
}

func (self *MapValue) GetAddTime() int64 {
  return self.addTime
}

func (self *MapValue) GetValue() Value {
  return self.value
}

func (self *Value) GetValue() uint64 {
  return self.value
}

func (self *Value) GetAction() uint64 {
  return self.action
}

// Returns the next key to be used for the given value.
// MUST BE CALLED ATOMICALLY UNDER A MUTEX LOCK
func (kv *TimedKv) nextKey() uint64 {
  if kv.keyCounter == MAX_KEY {
    // go back to zero (must happen in years)
    // TODO : log this event
    kv.keyCounter = 0
  }
  kv.keyCounter++
  return kv.keyCounter
}

func (th timedHeap) Len() int {
  return len(th)
}

func (th timedHeap) Less(i, j int) bool {
  // we want earliest expiry first
  return th[i].delTime < th[j].delTime
}

func (th timedHeap) Swap(i, j int) {
  th[i], th[j] = th[j], th[i]
}

func (th *timedHeap) Push(x interface{}) {
  item := x.(heapItem)
  *th = append(*th, item)
}

func (th *timedHeap) Pop() interface{} {
  old := *th
  n := len(old)
  item := old[n-1]
  *th = old[:n-1]
  return item
}

// Update utitlity of heap is not used here; to use,
// add the 'index' member too, and copy it from the
// documentation of 'heap' package


