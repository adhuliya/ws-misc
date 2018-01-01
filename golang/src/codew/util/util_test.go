package util

import (
	"bytes"
	logg "cs733/lib/log"
	"os"
	"strings"
	"testing"
)

//func autotest(f(int64, int64),

func TestMain(m *testing.M) {

	logg.InitLogger("app.log")

	x := m.Run()

	os.Exit(x)
}

func Test_MaxInt64(t *testing.T) {
	var testcase = [][]int64{{1, 2, 2}, {1234567, 2345678, 2345678}}

	for _, val := range testcase {
		res := MaxInt64(val[0], val[1])
		if res != val[2] {
			t.Error("Wrong resutl for", val, "result :", res)
		}
	}
}

func Test_MinInt64(t *testing.T) {
	var testcase = [][]int64{{1, 2, 1}, {1234567, 2345678, 1234567}}

	for _, val := range testcase {
		res := MinInt64(val[0], val[1])
		if res != val[2] {
			t.Error("Wrong resutl for", val, "result :", res)
		}
	}
}

func Test_GetInputClosure(t *testing.T) {
	// create input
	strs := []string{"hello\n", "what is your name\n", "hi"}

	SIZE := 20000
	bb := make([]byte, 0, SIZE)
	for i := 0; i < SIZE; i++ {
		bb = append(bb, 'a')
	}

	strs = append(strs, string(bb))

	// join all strings
	input := []byte(strings.Join(strs, ""))

	buffinput := bytes.NewBuffer(input)

	getLine, getBytes := GetInputClosure(buffinput)

	str, _ := getLine()
	logg.Trace.Println(string(str))
	if string(str) != strs[0] {
		t.Error("Invalid reading 1")
	}

	buff := make([]byte, len(strs[1]))
	_ = getBytes(buff)
	logg.Trace.Println(string(buff))
	if string(buff) != strs[1] {
		t.Error("Invalid reading 2")
	}

	_ = getBytes(buff[:2])

	buff = make([]byte, SIZE)
	_ = getBytes(buff)
	logg.Trace.Println(string(buff))
	if string(buff) != strs[3] {
		t.Error("Invalid reading 3")
	}
}
