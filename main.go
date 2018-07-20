package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: %s unixdate number\n", os.Args[0])
		os.Exit(1)
	}

	i, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	num, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	s := time.Unix(i, 0)
	for i := 0; i < num; i++ {
		fmt.Println(s.Add(time.Hour * -time.Duration(i)).Format("Mon 02/01/06 15:04"))
	}
}
