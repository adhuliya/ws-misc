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
  key,_ := tkv.Add(12, time.Second*2)

  time.Sleep(time.Millisecond * 10)
  // Check its presence
  _, ispresent := tkv.Get(key)
  if !ispresent {
    t.Error("key not present:", key)
  }

  // No modification allowed. Delete and then Add.

  // Delete the value
  tkv.Delete(key)
  _, ispresent = tkv.Get(key)

  if ispresent {
    t.Error("key", "12", "is present!")
  }

  tkv.Destroy()
}

// Is key deleted in time? (Single Key)
func Test_timelyDeletion1(t *testing.T) {
  tkv := Make()

  key,_ := tkv.Add(12, time.Second)

  time.Sleep(time.Duration(time.Second))

  _, isPresent := tkv.Get(key)
  if isPresent {
    t.Error("Key not deleted in time")
  }

  tkv.Destroy()
}

// Is key deleted in time? (Many Keys)
func Test_timelyDeletion2_1000K(t *testing.T) {
  tkv := Make()
  var limit uint64 = 20000000
  var i uint64
  //pause := time.Millisecond * 1

  go func() {
      for {
          // keep consuming the Done channel
          <-tkv.Done
      }
  }()

  for i = 0; i < limit; i++ {
      _, err := tkv.Add(i, time.Second)
      if err != nil {
          t.Error("Cannot add key.", i)
      }
      //if i % 100000 == 0 {
      //    //time.Sleep(pause)
      //}
  }

  time.Sleep(time.Duration(time.Second))

  //fmt.Println("hereN-10")
  // All keys should be deleted
  for i = 0; i < limit; i++ {
    _, isPresent := tkv.Get(i)
    if isPresent {
      t.Error("Keys not deleted in time.")
    }
    //if i % 100000 == 0 {
    //    time.Sleep(pause)
    //}
  }

  //fmt.Println("hereN")
  tkv.Destroy()
}

//// Is key deleted in time? (Many Many Keys)
//func Test_timelyDeletion3_50000K(t *testing.T) {
//  tkv := Make()
//  var limit uint64 = 50000000
//  var i uint64
//  //pause := time.Millisecond * 1
//
//  go func() {
//      for {
//          // keep consuming the Done channel
//          <-tkv.Done
//      }
//  }()
//
//  for i = 0; i < limit; i++ {
//      _, err := tkv.Add(i, time.Second)
//      if err != nil {
//          t.Error("Cannot add key.", i)
//      }
//  }
//
//  time.Sleep(time.Duration(time.Second))
//
//  // All keys should be deleted
//  for i = 0; i < limit; i++ {
//    _, isPresent := tkv.Get(i)
//    if isPresent {
//      t.Error("Keys not deleted in time.")
//    }
//  }
//
//  tkv.Destroy()
//}
//
//// Is key deleted in time? (Many Keys)
//// Are all values received in Done channel
//func Test_timelyAlarmedDeletion3(t *testing.T) {
//  tkv := Make()
//  limit := 20000000
//  arr   := make([]uint64, limit, limit)
//  count := 0
//
//  for i := 0; i < limit; i++ {
//    arr[i],_ = tkv.Add(Value{23,23}, tkv.minDuration)
//  }
//
//  for _ = range tkv.Done {
//      count++
//      if count == limit {
//          break
//      }
//  }
//
//
//  tkv.Destroy()
//}
//
//// Is key retained within time? (Many Keys)
//func Test_keysRetained1(t *testing.T) {
//  tkv := Make()
//  limit := 1000000
//  arr   := make([]uint64, limit, limit)
//
//  for i := 0; i < limit; i++ {
//    arr[i],_ = tkv.Add(Value{23,23}, tkv.maxDuration)
//  }
//
//  // time.Sleep(time.Duration(tkv.minDuration))
//
//  // All keys should be present
//  for i := 0; i < limit; i++ {
//    _, isPresent := tkv.Get(arr[i])
//    if !isPresent {
//      t.Error("Keys deleted within time.")
//    }
//  }
//
//  tkv.Destroy()
//}
//
//// Is key retained within time? (Many Keys, updating duration)
//// Add/Extend Duration/Check Retained
//func Test_keysRetained2(t *testing.T) {
//  tkv := Make()
//  limit := 1000000
//  var duration1 = time.Second * 8
//  var duration2 = time.Second * 15
//  arr   := make([]uint64, limit, limit)
//
//  // Add values
//  for i := 0; i < limit; i++ {
//    arr[i],_ = tkv.Add(Value{23,23}, duration1)
//  }
//
//  // Modify values
//  for i := 0; i < limit; i++ {
//    status,err := tkv.Modify(arr[i], Value{1,2}, duration2, MODIFY_DURATION)
//    if status == false || err != nil {
//      t.Error("Error modifying key.")
//    }
//  }
//
//  time.Sleep(time.Duration(duration1))
//
//  // All keys should be present
//  for i := 0; i < limit; i++ {
//    _, isPresent := tkv.Get(arr[i])
//    if !isPresent {
//      t.Error("Keys deleted within time.")
//    }
//  }
//
//  tkv.Destroy()
//}
//
//
