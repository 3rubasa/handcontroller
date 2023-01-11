package interfaces

type HorizontalMover interface {
	Init() error
	MoveLeft() error
	MoveRight() error
}

type VerticalMover interface {
	Init() error
	MoveUp() error
	MoveDown() error
}

type DepthMover interface {
	Init() error
	MoveForward() error
	MoveBackward() error
}
