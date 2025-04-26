package block

type LTypeBlock struct {
	*block
}

func NewLTypeBlock() *LTypeBlock {
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

	return &LTypeBlock{
		block: block,
	}
}

type JTypeBlock struct {
	*block
}

func NewJTypeBlock() *JTypeBlock {
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

	return &JTypeBlock{
		block: block,
	}
}

type ITypeBlock struct {
	*block
}

func NewITypeBlock() *ITypeBlock {
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

	return &ITypeBlock{
		block: block,
	}
}

type OTypeBlock struct {
	*block
}

func NewOTypeBlock() *OTypeBlock {
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

	return &OTypeBlock{
		block: block,
	}
}

type STypeBlock struct {
	*block
}

func NewSTypeBlock() *STypeBlock {
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

	return &STypeBlock{
		block: block,
	}
}

type TTypeBlock struct {
	*block
}

func NewTTypeBlock() *TTypeBlock {
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

	return &TTypeBlock{
		block: block,
	}
}

type ZTypeBlock struct {
	*block
}

func NewZTypeBlock() *ZTypeBlock {
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

	return &ZTypeBlock{
		block: block,
	}
}
