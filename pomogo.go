package main

import (
	"fmt"
	"time"
)

//Time is represented in seconds
type config struct {
	workTime      int
	breakTime     int
	longBreakTime int
}

func startBreak(breakTime int) {
	fmt.Printf("**************************************************************************\n")
	fmt.Printf("Break has started.\n")
	doneTimer := make(chan bool)
	go startTimer(breakTime, doneTimer, "Break")
	<-doneTimer

}

func startPomodoro(workTime int, pomodorosDone int) {
	fmt.Printf("**************************************************************************\n")
	fmt.Printf("Timer has started.\n")
	fmt.Printf("Number of Pomodoros done so far %d\n", pomodorosDone)
	doneTimer := make(chan bool)
	go startTimer(workTime, doneTimer, "Timer")
	<-doneTimer
}
func startTimer(breakTime int, done chan<- bool, mode string) {
	counter := breakTime
	fmt.Printf("%s set for: %2d minutes.\n", mode, counter/60)
	for ; counter >= 0; counter-- {
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
	conf := config{workTime: 1500, breakTime: 300, longBreakTime: 900}
	pomodorosDone := 0

	for {
		startPomodoro(conf.workTime, pomodorosDone)
		pomodorosDone++
		if pomodorosDone%4 != 0 {
			startBreak(conf.breakTime)
		} else {
			startBreak(conf.longBreakTime)
		}
	}
}
