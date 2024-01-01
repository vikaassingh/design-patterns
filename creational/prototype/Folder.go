package prototype

import "fmt"

type Folder struct {
	Children []IProtoType
	Name     string
}

func (f *Folder) GetId() ProtoType {
	return ProtoTypeFolder
}

func (f *Folder) Print(str string) {
	fmt.Println(str + f.Name)
	for _, child := range f.Children {
		child.Print(str + str)
	}
}

func (f *Folder) Clone() IProtoType {
	var children []IProtoType

	for _, child := range f.Children {
		children = append(children, child.Clone())
	}

	return &Folder{
		Name:     f.Name + "_clone",
		Children: children,
	}
}
