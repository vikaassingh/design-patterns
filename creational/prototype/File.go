package prototype

import "fmt"

type File struct {
	Name string
}

func (f *File) GetId() ProtoType {
	return ProtoTypeFile
}

func (f *File) Print(str string) {
	fmt.Println(str + f.Name)
}

func (f *File) Clone() IProtoType {
	return &File{
		Name: f.Name + "_clone",
	}
}
