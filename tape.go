package bfuk

const (
	arraySize = 30_000
)

type Tape struct {
	array       []uint8
	cellPointer uint
}

func (t *Tape) MoveRight() {
	if t.cellPointer != arraySize-1 {
		t.cellPointer++
	}
}

func (t *Tape) MoveLeft() {
	if t.cellPointer != 0 {
		t.cellPointer--
	}
}

func (t *Tape) IncrementCurrentCell() {
	t.array[t.cellPointer]++
}

func (t *Tape) DecrementCurrentCell() {
	t.array[t.cellPointer]--
}

func (t *Tape) GetCurrentCell() uint8 {
	return t.array[t.cellPointer]
}

func (t *Tape) SetCurrentCell(value uint8) {
	t.array[t.cellPointer] = value
}

func NewTape() *Tape {
	return &Tape{
		array:       make([]uint8, arraySize),
		cellPointer: 0,
	}
}
