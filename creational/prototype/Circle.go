package prototype

import "fmt"

type Circle struct {
	Radius float32
}

func (f *Circle) GetId() ProtoType {
	return ProtoTypeCircle
}

func (f *Circle) Print(str string) {
	fmt.Println(str+"Circle Radius:", f.Radius)
}

func (f *Circle) Clone() IProtoType {
	return &Circle{
		Radius: f.Radius,
	}
}
