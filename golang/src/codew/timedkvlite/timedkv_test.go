package timedkvlite

import (
//  "fmt"
  "os"
//  "sync"
  "testing"
  "time"
)

func TestMain(m *testing.M) {
  x := m.Run()

  os.Exit(x)
}

// Single Element Add/Modify/Delete
func Test_singleElementAMD(t *testing.T) {
  tkv := Make()

  // Put a value
  key := tkv.Add(12, int64(time.Second*10))

  // Check its presence
  val, ispresent := tkv.Get(key)
  if ispresent {
    if val.value != 12 {
      t.Error("Got val", val)
    }
  } else {
    t.Error("key not present:", key)
  }

  // Modify the value
  tkv.Modify(key, 23, 0, true, false)
  val, ispresent = tkv.Get(key)
  if ispresent {
    if val.value != 23 {
      t.Error("Got val", val)
    }
  } else {
    t.Error("key not present:", key)
  }

  // Delete the value
  tkv.Delete(key)
  val, ispresent = tkv.Get(key)

  if ispresent {
    t.Error("key", "12", "is present!")
  }

  tkv.Destroy()
}

// Is key deleted in time? (Single Key)
func Test_timelyDeletion1(t *testing.T) {
  tkv := Make()
  value := uint64(444)

  key := tkv.Add(value, tkv.minDuration)

  time.Sleep(time.Duration(tkv.minDuration))

  _, isPresent := tkv.Get(key)
  if isPresent {
    t.Error("Key not deleted in time")
  }

  tkv.Destroy()
}

// Is key deleted in time? (Many Keys)
func Test_timelyDeletion2(t *testing.T) {
  tkv := Make()
  value := uint64(444)
  limit := 1000000
  arr   := make([]uint64, limit, limit)

  for i := 0; i < limit; i++ {
    arr[i] = tkv.Add(value, tkv.minDuration)
  }

  time.Sleep(time.Duration(tkv.minDuration))

  // All keys should be deleted
  for i := 0; i < limit; i++ {
    _, isPresent := tkv.Get(arr[i])
    if isPresent {
      t.Error("Keys not deleted in time.")
    }
  }

  tkv.Destroy()
}

// Is key retained within time? (Many Keys)
func Test_keysRetained1(t *testing.T) {
  tkv := Make()
  value := uint64(444)
  limit := 1000000
  arr   := make([]uint64, limit, limit)

  for i := 0; i < limit; i++ {
    arr[i] = tkv.Add(value, tkv.maxDuration)
  }

  // time.Sleep(time.Duration(tkv.minDuration))

  // All keys should be present
  for i := 0; i < limit; i++ {
    _, isPresent := tkv.Get(arr[i])
    if !isPresent {
      t.Error("Keys deleted within time.")
    }
  }

  tkv.Destroy()
}

// Is key retained within time? (Many Keys, updating duration)
// Add/Extend Duration/Check Retained
func Test_keysRetained2(t *testing.T) {
  tkv := Make()
  value := uint64(444)
  limit := 1000000
  var duration1 int64 = int64(time.Second * 3)
  var duration2 int64 = int64(time.Second * 5)
  arr   := make([]uint64, limit, limit)

  // Add values
  for i := 0; i < limit; i++ {
    arr[i] = tkv.Add(value, duration1)
  }

  // Modify values
  for i := 0; i < limit; i++ {
    status := tkv.Modify(arr[i], value, duration2, false, true)
    if status == FAILURE {
      t.Error("Error modifying key.")
    }
  }

  time.Sleep(time.Duration(duration1))

  // All keys should be present
  for i := 0; i < limit; i++ {
    _, isPresent := tkv.Get(arr[i])
    if !isPresent {
      t.Error("Keys deleted within time.")
    }
  }

  tkv.Destroy()
}

// func Test_timedeletion1(t *testing.T) {
//   tkv := Make()
// 
//   step := 100
//   goroutines := 100
//   wg := &sync.WaitGroup{}
//   wg.Add(goroutines)
// 
//   for i := 0; i < goroutines; i++ {
//     go addkeys1(tkv, step*i, step*(i+1), wg)
//   }
// 
//   wg.Wait()
// 
//   time.Sleep(time.Millisecond)
// 
//   for i := 0; i < step*goroutines; i++ {
//     _, ispresent := tkv.Get(fmt.Sprintf("%d", i))
//     if ispresent {
//       t.Error("KV pair not deleted")
//     }
//   }
// 
//   //time.Sleep(time.Second)
// 
//   tkv.Destroy()
// }
// 
// func addkeys1(tkv *Timedkv, from int, to int, wg *sync.WaitGroup) {
//   defer wg.Done()
// 
//   for i := from; i < to; i++ {
//     tkv.Set(fmt.Sprintf("%d", i), Makevalue([]byte("h"), time.Millisecond))
//   }
// }
// 
// func Test_versiontest(t *testing.T) {
//   tkv := Make()
// 
//   key := "hello"
//   tkv.Set(key, Makevalue([]byte("world0"), time.Second*10))
// 
//   val, ispresent := tkv.Get(key)
//   if !ispresent {
//     t.Error("key not present")
//   }
// 
//   version := val.version
// 
//   limit := int64(100000)
//   for i := int64(0); i < limit; i++ {
//     status := tkv.Cas(key, version+i, Makevalue(
//       []byte(fmt.Sprintf("world%d", i+1)),
//       time.Second*10))
//     if status != SUCCESS {
//       t.Error("status is", status)
//     }
//   }
// 
//   val, ispresent = tkv.Get(key)
//   if !ispresent {
//     t.Error("key not present")
//   }
// 
//   if val.version != version+limit {
//     t.Error("different version", val.version)
//   }
// 
//   if string(val.Value().([]byte)) != fmt.Sprintf("world%d", limit) {
//     t.Error("value is different", string(val.Value().([]byte)))
//   }
// 
//   status := tkv.Cad(key, version+limit)
//   if status != SUCCESS {
//     t.Error("key not deleted, diffrent version")
//   }
// 
//   tkv.Destroy()
// }
