package hmover

import (
	"fmt"

	"github.com/3rubasa/go-servo-picobber"
	"github.com/3rubasa/handcontroller/interfaces"
)

const defaultPosition int = 326 // 90 deg
const frequency uint8 = 50      // 50 Hz
const rightMostPos int = 110    // 0 deg
const leftMostPos int = 542     // 180 deg

type horizontalMover struct {
	channel byte
	step    uint
	curPos  int
	sv      *servo.Servo
}

func NewHorizontalMover(channel byte, step uint) interfaces.HorizontalMover {
	return &horizontalMover{
		channel: channel,
		step:    step,
	}
}

func (p *horizontalMover) Init() error {
	var err error
	fmt.Println("Creating new Servo for Horizontal Mover ...")

	p.sv, err = servo.NewServo()
	if err != nil {
		fmt.Println("Failed to create new Servo: ", err)
		return err
	}

	fmt.Println("Setting frequency for Horizontal Mover...")
	err = p.sv.SetPwmFreq(frequency) // Set frequency to 50 Hz

	if err != nil {
		fmt.Println("Failed to set servo frequency: ", err)
		return err
	}

	err = p.sv.SetPwm(p.channel, 0, defaultPosition)

	if err != nil {
		fmt.Println("Failed to set default position of Horizontal Mover: ", err)
		return err
	}

	p.curPos = defaultPosition

	return nil
}

func (p *horizontalMover) MoveRight() error {
	if p.curPos == rightMostPos {
		return nil
	}

	newPos := p.curPos - int(p.step)
	if newPos < rightMostPos {
		newPos = rightMostPos
	}

	err := p.sv.SetPwm(p.channel, 0, newPos)

	if err != nil {
		fmt.Println("Failed to move right: ", err)
		return err
	}

	p.curPos = newPos

	return nil
}

func (p *horizontalMover) MoveLeft() error {
	if p.curPos == leftMostPos {
		return nil
	}

	newPos := p.curPos + int(p.step)
	if newPos > leftMostPos {
		newPos = leftMostPos
	}

	err := p.sv.SetPwm(p.channel, 0, newPos)

	if err != nil {
		fmt.Println("Failed to move left: ", err)
		return err
	}

	p.curPos = newPos

	return nil
}
