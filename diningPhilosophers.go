package main

import (
	"fmt"
	"time"
)

//Made by otja, brml and kbej

// Everytime a philosopher tries to pick two forks up and one of them can't be used,
// he puts the other fork down and starts thinking.
// Due to the somewhat random runtime of each philosopher, they will inevitably find two free forks and eat.

func philosoph(firstForkSend chan bool, firstForkRecv chan bool, secondForkSend chan bool, secondForkRecv chan bool, index int, satisfied chan bool) {
	timesEaten := 0

	for {
		canUseFirstFork := <-firstForkSend
		firstForkRecv <- false
		canUseSecondFork := <-secondForkSend
		secondForkRecv <- false

		if canUseFirstFork && canUseSecondFork {
			fmt.Printf("Philosopher %d is eating\n", index)
			timesEaten++
			if timesEaten == 3 {
				satisfied <- true
			}
			<-firstForkSend
			firstForkRecv <- true
			<-secondForkSend
			secondForkRecv <- true
			time.Sleep(2 * time.Second)
		} else {
			if canUseFirstFork {
				<-firstForkSend
				firstForkRecv <- true
			}
			if canUseSecondFork {
				<-secondForkSend
				secondForkRecv <- true
			}
			fmt.Printf("Philosopher %d is thinking\n", index)
			time.Sleep(2 * time.Second)
		}
	}
}

func fork(sendChannel chan<- bool, reciveChannel <-chan bool) {
	sendChannel <- true
	for {
		available := <-reciveChannel
		sendChannel <- available
	}
}

func main() {
	fork0chRecv := make(chan bool, 1)
	fork1chRecv := make(chan bool, 1)
	fork2chRecv := make(chan bool, 1)
	fork3chRecv := make(chan bool, 1)
	fork4chRecv := make(chan bool, 1)

	fork0chSend := make(chan bool, 1)
	fork1chSend := make(chan bool, 1)
	fork2chSend := make(chan bool, 1)
	fork3chSend := make(chan bool, 1)
	fork4chSend := make(chan bool, 1)

	go fork(fork0chSend, fork0chRecv)
	go fork(fork1chSend, fork1chRecv)
	go fork(fork2chSend, fork2chRecv)
	go fork(fork3chSend, fork3chRecv)
	go fork(fork4chSend, fork4chRecv)

	satisfied := make(chan bool, 5)
	go philosoph(fork0chSend, fork0chRecv, fork1chSend, fork1chRecv, 0, satisfied)
	go philosoph(fork1chSend, fork1chRecv, fork2chSend, fork2chRecv, 1, satisfied)
	go philosoph(fork2chSend, fork2chRecv, fork3chSend, fork3chRecv, 2, satisfied)
	go philosoph(fork3chSend, fork3chRecv, fork4chSend, fork4chRecv, 3, satisfied)
	go philosoph(fork4chSend, fork4chRecv, fork0chSend, fork0chRecv, 4, satisfied)

	for i := 0; i < 5; i++ {
		<-satisfied
	}
}
