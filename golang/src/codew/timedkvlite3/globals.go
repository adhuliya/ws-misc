package timedkvlite

import (
  "sync"
//  "time"
)

const (
    // max key that can be used in the map
    MAX_KEY uint64 = 1<<64 - 1
)

type TimedKv struct {
  kv            map[uint64]int64 //[Key]Value
  heap          timedHeap
  kvAdd         chan heapItem
  heapAdd       chan heapItem

  kvMutex       *sync.RWMutex
  heapMutex     *sync.Mutex
  kvBuffMutex   *sync.Mutex

  kvAddBuff     map[uint64]int64
  Done          chan uint64
}

// No update() allowed in this heap
type timedHeap []heapItem

type heapItem struct {
  // deliberately without the 'index' member
  // no updates or removal allowed; let it get garbage collected
  // It is only 16 bytes and will eventually get removed in time.
  key     uint64   // Copy of Key.Key

  delTime int64    // absolute time at which to delete this
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


