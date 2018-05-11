package util
// utlity functions that makes programming convenient.

import (
	"bufio"
	logg "cs733/lib/log"
	"io"
	"sync"
	"time"
)

// maxInt64 returns the max of two int64 values given.
func MaxInt64(a, b int64) int64 {
	switch {
	case a > b:
		return a
	default:
		return b
	}
}

func MinInt64(a, b int64) int64 {
	switch {
	case a < b:
		return a
	default:
		return b
	}
}

func MinInt(a, b int) int {
	switch {
	case a < b:
		return a
	default:
		return b
	}
}

func ResetTimer(timer *time.Timer, dur time.Duration) {
	if !timer.Reset(dur) {
		select {
		// emtpy the timer channel off the previous timeout
		case <-timer.C:
		default:
		}
	}
}

// getInputClosure() returns two closures, first resturns a line of bytes
// ending with '\n' and the other returning bytes. Both the closures are
// thread safe.
func GetInputClosure(reader io.Reader) (func() ([]byte, error),
	func([]byte) error) {

	buffread := bufio.NewReaderSize(reader, 1024)
	mutex := &sync.Mutex{}
	SIZE := 256
	localbuff := make([]byte, SIZE)

	// returns a line of bytes ending with and including '\n'
	getLine := func() ([]byte, error) {
		mutex.Lock()
		defer mutex.Unlock()

		line, err := buffread.ReadBytes('\n')
		if err != nil {
			logg.Fatal.Println(err)
			// panic(err) // TODO commentit
			return nil, err
		}

		return line, err
	}

	// Reads bytes equal to the length of the buff,
	// otherwise returns an error
	getBytes := func(buff []byte) error {
		mutex.Lock()
		defer mutex.Unlock()

		length := len(buff)
		buff = buff[:0] // make zero length for append()

		for length > 0 {
			localbuff = localbuff[:MinInt(length, SIZE)]

			n, err := buffread.Read(localbuff)
			//logg.Trace.Println("Read", string(localbuff[:n]))
			if err != nil {
				logg.Fatal.Println(err)
				return err
			}

			buff = append(buff, localbuff[:n]...)

			length -= n
		}

		return nil
	}

	return getLine, getBytes
}


