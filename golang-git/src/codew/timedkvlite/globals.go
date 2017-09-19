package timedkvlite

import (
  "sync"
)

type TimedKv struct {
  kv        map[uint64]Value
  heap      timedHeap
  heapAdd   chan heapItem
  kvMutex   *sync.RWMutex
  heapMutex *sync.Mutex

  key       uint64
  maxKey    uint64

  maxDuration int64
  minDuration int64

  maxDurationYet  int64
}


type Value struct {
  // Deliberately 32 bytes for efficiency (mem aligned)
  key       uint64
  value     uint64

  addTime   int64 // (nanosec) time when added
  delTime   int64 // (nanosec) time when to delete
}

// No update() allowed in this heap
type timedHeap []heapItem

type heapItem struct {
  // deliberately without the 'index' member
  // no updates or removal allowed; let it get garbage collected
  // It is only 16 bytes and will eventually get removed in time.
  key     uint64   // Related to TimedKv.kv[key]

  delTime int64   // time at which to delete this
}

const (
  SUCCESS     byte = 2
  FAILURE     byte = 3 // key not present in CAS and CAD
)

func (self *Value) GetDelTime() int64 {
  return self.delTime
}

func (self *Value) GetValue() uint64 {
  return self.value
}

// Returns the next key to be used for the given value.
// MUST BE CALLED ATOMICALLY UNDER MUTEX
func (kv *TimedKv) nextKey() uint64 {
  if kv.key == kv.maxKey {
    // go back to zero (must happen in years)
    // TODO : log this event
    kv.key = 0
  }
  kv.key++
  return kv.key
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


