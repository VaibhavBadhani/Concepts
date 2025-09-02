package main

import (
	"fmt"
	"math"
	"time"

	"golang.org/x/exp/rand"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}
func channelSum(c, co chan int) {
	for {
		x := <-c
		y := <-c
		co <- x + y
	}
}

func properChannels() {
	c := make(chan int, 6)
	co := make(chan int, 3)
	go channelSum(c, co)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	c <- 6
	r1 := <-co
	r2 := <-co
	r3 := <-co
	fmt.Printf("the first value for 1,2 is %d\n", r1)
	fmt.Printf("the first value for 3,4 is %d\n", r2)
	fmt.Printf("the first value for 5,6 is %d\n", r3)
}

type primemsg struct {
	num     int
	isPrime bool
}

const TEST_LEN = 100
const NUM_WOKERS = 3

func isPrime(cin, cout chan primemsg) {
	id := rand.Intn(1000)
	i := 0
	for {
		msg := <-cin
		num := msg.num
		fmt.Println(id, "is testing", num)
		sq_root := int(math.Sqrt(float64(num)))
		for i = 2; i <= sq_root; i++ {
			if num%i == 0 {
				msg.isPrime = false
				cout <- msg
				break
			}
		}
		if i > sq_root {
			msg.isPrime = true
			cout <- msg
		}

	}
}
func main() {
	// go say("hey")
	// say("vaibhav")

	// c := make(chan int) //default length 1  c := make(chan int,2) if i make two i get answer
	// 3 because of race condition that it was this much fast that it read either 3 or 4
	// go channelSum(c)
	// c <- 3
	// c <- 4
	// answer := <-c
	// fmt.Println(answer)

	// properChannels()

	msg := primemsg{42, false}

	//create channel
	cin := make(chan primemsg, TEST_LEN)
	cout := make(chan primemsg, TEST_LEN)

	//create workers
	for i := 0; i < NUM_WOKERS; i++ {
		go isPrime(cin, cout)
	}

	//fill the input queue
	for i := 0; i < TEST_LEN; i++ {
		msg.num = rand.Intn(1000) + 1000
		cin <- msg
	}

	//read the answer
	for i := 0; i < TEST_LEN; i++ {
		msg = <-cout
		fmt.Println(msg.num, msg.isPrime)
	}

}
