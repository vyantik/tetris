package block

type LBlock struct {
	Block block
}

func NewLBlock() *LBlock {
	block := newBlock()
	block.id = 1
	block.cells = map[int][]position{
		0: {
			*newPosition(0, 2),
			*newPosition(1, 0),
			*newPosition(1, 1),
			*newPosition(1, 2),
		},
		1: {
			*newPosition(0, 1),
			*newPosition(1, 1),
			*newPosition(2, 1),
			*newPosition(2, 2),
		},
		2: {
			*newPosition(1, 0),
			*newPosition(1, 1),
			*newPosition(1, 2),
			*newPosition(2, 0),
		},
		3: {
			*newPosition(0, 0),
			*newPosition(0, 1),
			*newPosition(1, 1),
			*newPosition(2, 1),
		},
	}
	block.Move(0, 3)

	return &LBlock{
		Block: *block,
	}
}

type JBlock struct {
	Block block
}

func NewJBlock() *JBlock {
	block := newBlock()
	block.id = 2
	block.cells = map[int][]position{
		0: {
			*newPosition(0, 0),
			*newPosition(1, 0),
			*newPosition(1, 1),
			*newPosition(1, 2),
		},
		1: {
			*newPosition(0, 1),
			*newPosition(0, 2),
			*newPosition(1, 1),
			*newPosition(2, 1),
		},
		2: {
			*newPosition(1, 0),
			*newPosition(1, 1),
			*newPosition(1, 2),
			*newPosition(2, 2),
		},
		3: {
			*newPosition(0, 1),
			*newPosition(1, 1),
			*newPosition(2, 0),
			*newPosition(2, 1),
		},
	}
	block.Move(0, 3)

	return &JBlock{
		Block: *block,
	}
}

type IBlock struct {
	Block block
}

func NewIBlock() *IBlock {
	block := newBlock()
	block.id = 3
	block.cells = map[int][]position{
		0: {
			*newPosition(1, 0),
			*newPosition(1, 1),
			*newPosition(1, 2),
			*newPosition(1, 3),
		},
		1: {
			*newPosition(0, 2),
			*newPosition(1, 2),
			*newPosition(2, 2),
			*newPosition(3, 2),
		},
		2: {
			*newPosition(2, 0),
			*newPosition(2, 1),
			*newPosition(2, 2),
			*newPosition(2, 3),
		},
		3: {
			*newPosition(0, 1),
			*newPosition(1, 1),
			*newPosition(2, 1),
			*newPosition(3, 1),
		},
	}
	block.Move(-1, 3)

	return &IBlock{
		Block: *block,
	}
}

type OBlock struct {
	Block block
}

func NewOBlock() *OBlock {
	block := newBlock()
	block.id = 4
	block.cells = map[int][]position{
		0: {
			*newPosition(0, 0),
			*newPosition(0, 1),
			*newPosition(1, 0),
			*newPosition(1, 1),
		},
		1: {
			*newPosition(0, 0),
			*newPosition(0, 1),
			*newPosition(1, 0),
			*newPosition(1, 1),
		},
		2: {
			*newPosition(0, 0),
			*newPosition(0, 1),
			*newPosition(1, 0),
			*newPosition(1, 1),
		},
		3: {
			*newPosition(0, 0),
			*newPosition(0, 1),
			*newPosition(1, 0),
			*newPosition(1, 1),
		},
	}
	block.Move(0, 4)

	return &OBlock{
		Block: *block,
	}
}

type SBlock struct {
	Block block
}

func NewSBlock() *SBlock {
	block := newBlock()
	block.id = 5
	block.cells = map[int][]position{
		0: {
			*newPosition(0, 1),
			*newPosition(0, 2),
			*newPosition(1, 0),
			*newPosition(1, 1),
		},
		1: {
			*newPosition(0, 1),
			*newPosition(1, 1),
			*newPosition(1, 2),
			*newPosition(2, 2),
		},
		2: {
			*newPosition(1, 1),
			*newPosition(1, 2),
			*newPosition(2, 0),
			*newPosition(2, 1),
		},
		3: {
			*newPosition(0, 0),
			*newPosition(1, 0),
			*newPosition(1, 1),
			*newPosition(2, 1),
		},
	}
	block.Move(0, 3)

	return &SBlock{
		Block: *block,
	}
}

type TBlock struct {
	Block block
}

func NewTBlock() *TBlock {
	block := newBlock()
	block.id = 6
	block.cells = map[int][]position{
		0: {
			*newPosition(0, 1),
			*newPosition(1, 0),
			*newPosition(1, 1),
			*newPosition(1, 2),
		},
		1: {
			*newPosition(0, 1),
			*newPosition(1, 1),
			*newPosition(1, 2),
			*newPosition(2, 1),
		},
		2: {
			*newPosition(1, 0),
			*newPosition(1, 1),
			*newPosition(1, 2),
			*newPosition(2, 1),
		},
		3: {
			*newPosition(0, 1),
			*newPosition(1, 0),
			*newPosition(1, 1),
			*newPosition(2, 1),
		},
	}
	block.Move(0, 3)

	return &TBlock{
		Block: *block,
	}
}

type ZBlock struct {
	Block block
}

func NewZBlock() *ZBlock {
	block := newBlock()
	block.id = 7
	block.cells = map[int][]position{
		0: {
			*newPosition(0, 0),
			*newPosition(0, 1),
			*newPosition(1, 1),
			*newPosition(1, 2),
		},
		1: {
			*newPosition(0, 2),
			*newPosition(1, 1),
			*newPosition(1, 2),
			*newPosition(2, 1),
		},
		2: {
			*newPosition(1, 0),
			*newPosition(1, 1),
			*newPosition(2, 1),
			*newPosition(2, 2),
		},
		3: {
			*newPosition(0, 1),
			*newPosition(1, 0),
			*newPosition(1, 1),
			*newPosition(2, 0),
		},
	}
	block.Move(0, 3)

	return &ZBlock{
		Block: *block,
	}
}
