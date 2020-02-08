package main

import (
	"fmt"
	"time"
)

//Time is represented in seconds
type config struct {
	work_time       uint32
	break_time      uint32
	long_break_time uint32
}

func start_break(break_time uint32) {
	fmt.Printf("**************************************************************************\n")
	fmt.Printf("Break has started.\n")
	done_timer := make(chan bool)
	go start_timer(break_time, done_timer, "Break")
	<-done_timer

}

func start_pomodoro(work_time uint32, pomodoros_done int) {
	fmt.Printf("**************************************************************************\n")
	fmt.Printf("Timer has started.\n")
	fmt.Printf("Number of Pomodoros done so far %d\n", pomodoros_done)
	done_timer := make(chan bool)
	go start_timer(work_time, done_timer, "Timer")
	<-done_timer
}
func start_timer(break_time uint32, done chan<- bool, mode string) {
	counter := break_time
	fmt.Printf("%s set for: %2d minutes.\n", mode, counter/60)
	for ; counter > 0; counter-- {
		fmt.Printf("\r%02d:%02d minutes left.", counter/60, counter%60)
		time.Sleep(time.Second)
	}
	go beep()
	fmt.Println()
	done <- true
}
func beep() {
	for i := 0; i < 5; i++ {
		fmt.Printf("\a")
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	conf := config{work_time: 1500, break_time:300 , long_break_time: 900}
	pomodoros_done := 0
	for {
		start_pomodoro(conf.work_time, pomodoros_done)
		pomodoros_done++
		if pomodoros_done%4 != 0 {
			start_break(conf.break_time)
		} else {
			start_break(conf.long_break_time)
		}
	}
}