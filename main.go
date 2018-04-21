package main

import (
	"time"
	"regexp"
	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		select {
		case <-time.After(time.Millisecond):
			match, _ := regexp.MatchString("Cookie Clicker$", robotgo.GetTitle())
			if match == true {
				x, y := robotgo.GetMousePos()
				if x >= 122 && x <= 300 && y >= 311 && y <= 472 {
					robotgo.MouseClick()
				}
			}
		}
	}
}
