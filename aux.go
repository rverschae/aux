package aux

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var Verbose = false //true

func init() {
	flag.BoolVar(&Verbose, "v", false, "verbose")
}

func Took(t time.Duration, s string, args ...interface{}) string {
	tt := fmt.Sprintf("[%v] ", t)
	text := fmt.Sprintf(tt+s+"\n", args...)
	if Verbose {
		fmt.Printf(text)
	}
	return text
}

func ErrorM(s string, args ...interface{}) error {
	e := fmt.Sprintf("cv.Err: "+s+"\n", args...)
	if Verbose {
		fmt.Fprintf(os.Stderr, e)
	}
	return fmt.Errorf(e)
}
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//type BOOL int16
const (
	False = int16(iota)
	True
)
