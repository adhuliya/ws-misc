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
  var limit uint64 = 1000000
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
      //if i % 1000 == 0 {
      //    fmt.Println("testing add", i)
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

// Is key deleted in time? (Many Many Keys)
func Test_timelyDeletion3_5000K(t *testing.T) {
  tkv := Make()
  var limit uint64 = 5000000
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
  }

  time.Sleep(time.Duration(time.Second))

  // All keys should be deleted
  for i = 0; i < limit; i++ {
    _, isPresent := tkv.Get(i)
    if isPresent {
      t.Error("Keys not deleted in time.")
    }
  }

  tkv.Destroy()
}

// Is key deleted in time? (Many Keys)
// Are all values received in Done channel
func Test_alarmConfirmation(t *testing.T) {
  tkv := Make()
  limit := 5000000
  count := 0

  for i := 0; i < limit; i++ {
    _, err := tkv.Add(uint64(i), time.Second)
    if err != nil {
        t.Error("Cannot add key.", i)
    }
  }

  for _ = range tkv.Done {
      count++
      if count == limit {
          break
      }
  }

  select {
  case <-tkv.Done:
    t.Error("tkv.Done should be empty")
  default:
  }

  tkv.Destroy()
}

// Is key retained within time? (Many Keys)
func Test_keysRetained1(t *testing.T) {
  tkv := Make()
  limit := 5000000

  for i := 0; i < limit; i++ {
    _, err := tkv.Add(uint64(i), time.Second * 100)
    if err != nil {
        t.Error("Cannot add key.", i)
    }
  }

  // time.Sleep(time.Duration(tkv.minDuration))

  // All keys should be present
  for i := 0; i < limit; i++ {
    _, isPresent := tkv.Get(uint64(i))
    if !isPresent {
      t.Error("Keys deleted within time.")
    }
  }

  tkv.Destroy()
}

// Is key retained within time? (Many Keys)
// to check how much time is required to insert.
func Test_add5000000keys(t *testing.T) {
  tkv := Make()
  limit := 5000000

  for i := 0; i < limit; i++ {
    _, err := tkv.Add(uint64(i), time.Second * 100)
    if err != nil {
        t.Error("Cannot add key.", i)
    }
  }

  tkv.Destroy()
}


