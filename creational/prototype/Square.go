package prototype

import "fmt"

type Square struct {
	Side float32
}

func (f *Square) GetId() ProtoType {
	return ProtoTypeSquare
}

func (f *Square) Print(str string) {
	fmt.Println(str+"Square Side:", f.Side)
}

func (f *Square) Clone() IProtoType {
	return &Square{
		Side: f.Side,
	}
}
