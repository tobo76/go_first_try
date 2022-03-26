package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

/* function to print the current date time - int arg 1 defines the number of prints */
func main() {

	fmt.Println("Date time prints:")

	progargs_cnt := len(os.Args)

	loop_cnt := 1
	if progargs_cnt > 1 {
		arg_loop_cnt, err := strconv.Atoi(os.Args[1])
		if err == nil {
			loop_cnt = arg_loop_cnt
		} else {
			fmt.Println("error parsing argument:", err)
		}
	}

	for i := 0; i < loop_cnt; i++ {

		dt := time.Now()
		fmt.Println("Current date and time is: ", dt.Weekday(), ":", dt.Year(), "-", dt.Month(), "-", dt.Day(), "::", dt.Hour(), ":", dt.Minute(), ":", dt.Second(), ":", dt.Nanosecond())
	}
}
