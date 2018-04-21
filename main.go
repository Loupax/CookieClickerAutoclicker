package main

import (
	"github.com/go-vgo/robotgo"
	"regexp"
	"time"
)

func main() {
	for {
		match, _ := regexp.MatchString("Cookie Clicker$", robotgo.GetTitle())
		if match == true {
			x, y := robotgo.GetMousePos()
			if x >= 122 && x <= 300 && y >= 311 && y <= 472 {
				robotgo.MouseClick()
				time.Sleep(time.Millisecond)
			}
		} else {
			time.Sleep(time.Second * 5)
		}
	}
}
