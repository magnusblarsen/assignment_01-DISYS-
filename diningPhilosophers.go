package main

import (
	"fmt"
	"time"
)

func philosoph(firstFork chan bool, secondFork chan bool, index int, satisfied chan bool) {
	timesEaten := 0

	for {
		canUseFirstFork := <-firstFork
		firstFork <- false
		canUseSecondFork := <-secondFork
		secondFork <- false

		if canUseFirstFork && canUseSecondFork {
			fmt.Printf("%d is eating\n", index)
			timesEaten++
			if timesEaten == 3 {
				satisfied <- true
			}
			<-firstFork
			firstFork <- true
			<-secondFork
			secondFork <- true
			time.Sleep(2 * time.Second)
			// time.Sleep(time.Duration(rand.Intn(5) * int(time.Second)))
		} else {
			if canUseFirstFork {
				<-firstFork
				firstFork <- true
			}
			if canUseSecondFork {
				<-secondFork
				secondFork <- true
			}
			fmt.Printf("%d is thinking\n", index)
			// time.Sleep(time.Duration(rand.Intn(5) * int(time.Second)))
			time.Sleep(2 * time.Second)
			fmt.Println("Done thinking")
		}

	}

}

// func fork() {
// 	for {
// 		if <-channel1 {

// 		}
// 	}
// }

func main() {
	fork0cha := make(chan bool, 1)
	fork1cha := make(chan bool, 1)
	fork2cha := make(chan bool, 1)
	fork3cha := make(chan bool, 1)
	fork4cha := make(chan bool, 1)
	fork0cha <- true
	fork1cha <- true
	fork2cha <- true
	fork3cha <- true
	fork4cha <- true

	// forkChannels := make([]chan bool, 5)
	// forkChannels[0] = fork0cha
	// forkChannels[1] = fork1cha
	// forkChannels[2] = fork2cha
	// forkChannels[3] = fork3cha
	// forkChannels[4] = fork4cha

	satisfied := make(chan bool, 5)
	go philosoph(fork0cha, fork1cha, 0, satisfied)
	go philosoph(fork1cha, fork2cha, 1, satisfied)
	go philosoph(fork2cha, fork3cha, 2, satisfied)
	go philosoph(fork3cha, fork4cha, 3, satisfied)
	go philosoph(fork4cha, fork0cha, 4, satisfied)

	for i := 0; i < 5; i++ {
		<-satisfied
	}
}
