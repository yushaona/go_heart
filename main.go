package main

import (
	"fmt"
	"os"

	"os/signal"
	"strings"
	"time"
)

func getIndex(v int, length int) int {
	if v < 0 {
		return length + v
	}

	return v
}

func main() {

	c := make(chan os.Signal)
	signal.Notify(c)

	var word string = "I LOVE YOU BABY"

	res := strings.Split(word, " ")

	for _, item := range res {
		item1 := string(item + " ")
		itemLength := len(item1)
		letterlist := make([]string, 1)
		for y := 12; y > -12; y-- {

			letters := ""

			for x := -30; x < 30; x++ {
				f := ((float32(x)*0.05)*(float32(x)*0.05) + (float32(y)*0.1)*(float32(y)*0.1) - 1)
				expression := f*f*f - (float32(x)*0.05)*(float32(x)*0.05)*(float32(y)*0.1)*(float32(y)*0.1)*(float32(y)*0.1)

				if expression <= 0 {

					letters += string(item1[getIndex((x-y)%itemLength, itemLength)])
				} else {
					letters += " "
				}
			}

			letterlist = append(letterlist, letters)
		}
		for i := 0; i < len(letterlist); i++ {
			fmt.Println(letterlist[i])
		}

		time.Sleep(time.Second * 2)
	}

	s := <-c
	fmt.Println("退出信号", s)
}
