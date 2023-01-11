package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/3rubasa/handcontroller/movers/dmover"
	"github.com/3rubasa/handcontroller/movers/hmover"
	"github.com/3rubasa/handcontroller/movers/vmover"
)

func main() {
	hMover := hmover.NewHorizontalMover(0, 10)
	err := hMover.Init()
	if err != nil {
		fmt.Println("Failed to init Horizontal Mover: ", err)
		return
	}

	vMover := vmover.NewVerticalMover(2, 10)
	err = vMover.Init()
	if err != nil {
		fmt.Println("Failed to init Vertical Mover: ", err)
		return
	}

	dMover := dmover.NewDepthMover(1, 10)
	err = dMover.Init()
	if err != nil {
		fmt.Println("Failed to init Depth Mover: ", err)
		return
	}

	ch := make(chan string)
	go func(ch chan string) {
		// disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		var b = make([]byte, 1)
		for {
			os.Stdin.Read(b)
			ch <- string(b)
		}
	}(ch)

	for {
		stdin, _ := <-ch

		fmt.Println("Keys pressed:", stdin)

		switch stdin {
		case "e":
			hMover.MoveRight()
		case "a":
			hMover.MoveLeft()
		case ",":
			dMover.MoveForward()
		case "o":
			dMover.MoveBackward()
		case "c":
			vMover.MoveUp()
		case "t":
			vMover.MoveDown()
		case "'":
			os.Exit(0)
		}
	}

	// for {
	// 	consoleReader := bufio.NewReaderSize(os.Stdin, 1)
	// 	input, _ := consoleReader.ReadByte()
	// 	ascii := input

	// 	// ESC = 27 and Ctrl-C = 3
	// 	//if ascii == 27 || ascii == 3 {
	// 	//	fmt.Println("Exiting...")
	// 	//		os.Exit(0)
	// 	//	}

	// 	fmt.Println("ASCII : ", ascii)

	// 	switch ascii {
	// 	case 101:
	// 		hMover.MoveRight()
	// 	case 97:
	// 		hMover.MoveLeft()
	// 	case 44:
	// 		dMover.MoveForward()
	// 	case 111:
	// 		dMover.MoveBackward()
	// 	case 99:
	// 		vMover.MoveUp()
	// 	case 116:
	// 		vMover.MoveDown()
	// 	}
	// }

	// for {
	// 	// Change speed of continuous servo on channel O
	// 	err = sv.SetPwm(0, 0, servoMin)
	// 	if err != nil {
	// 		fmt.Println("Failed to set PWM 1", err)
	// 		return
	// 	}
	// 	time.Sleep(1 * time.Second)
	// 	err = sv.SetPwm(0, 0, servoMax)
	// 	if err != nil {
	// 		fmt.Println("Failed to set PWM 2", err)
	// 		return
	// 	}
	// 	time.Sleep(1 * time.Second)
	// }
}
