package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func fail(msg string) {
	fmt.Printf("ERR: %v\n", msg)
	os.Exit(2)
}

func readNumbers() chan int {
	bio := bufio.NewReader(os.Stdin)
	ch := make(chan int, 100)
	go func() {
		for {
			line, isPrefix, err := bio.ReadLine()
			if isPrefix {
				fail("too long")
			}
			if err != nil {
				if err == io.EOF {
					close(ch)
					break
				}
				fail(err.Error())
			}
			num, err := strconv.Atoi(string(line))
			if err != nil {
				fail("not a number")
			}
			ch <- num
		}
	}()
	return ch
}

func main() {
	lis := NewLIS()
	for num := range readNumbers() {
		lis.Add(num)
	}
	fmt.Printf("%d\n", lis.Len())
}
