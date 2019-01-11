package IR

type Instruction struct {
	Op  *Op
	Imm Imm
	//TODO:need Index?
	Index int
}

func (i *Instruction) Type() ValueType {
	return TypeIns
}

func (i *Instruction) Value() interface{} {
	return i
}
