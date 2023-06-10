package main

import (
	"errors"
	"fmt"
	"github.com/atadzan/playground/log/hello"
	"log"
	"runtime"
)

// log flags which we can set into standard log package to improve our log message
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	hello.Hello()
	log.Println("Log me here")

	// Runtime caller which shows where it was called. It shows file path, called line and ok
	_, file, line, _ := runtime.Caller(0)

	fmt.Printf("Error: %s:%d %v \n", file, line, errors.New("invalid input"))
}
