package main

import (
	"fmt"
	"time"

	"github.com/3rubasa/go-servo-picobber"
)

func main() {
	servoMin := 150 // Min pulse length out of 4096
	servoMax := 600 // Max pulse length out of 4096

	fmt.Println("Creating new servo...")
	sv, err := servo.NewServo()
	if err != nil {
		fmt.Println("Failed to create ne Servo: ", err)
		return
	}
	fmt.Println("New servo created!")

	fmt.Println("Settin frequency...")
	err = sv.SetPwmFreq(50) // Set frequency to 60 Hz
	if err != nil {
		fmt.Println("Failed to set servo frequency: ", err)
		return
	}
	fmt.Println("New servo created!")

	for {
		// Change speed of continuous servo on channel O
		err = sv.SetPwm(0, 0, servoMin)
		if err != nil {
			fmt.Println("Failed to set PWM 1", err)
			return
		}
		time.Sleep(1 * time.Second)
		err = sv.SetPwm(0, 0, servoMax)
		if err != nil {
			fmt.Println("Failed to set PWM 2", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
