package prototype

type ProtoType int

const (
	ProtoTypeFile ProtoType = iota
	ProtoTypeFolder
	ProtoTypeCircle
	ProtoTypeSquare
	ProtoTypeTriangle
	ProtoTypeRectangle
)

func LoadRegistry(r *Registry) {
	file := &File{Name: "File.txt"}
	fold := &Folder{Name: "FoldText", Children: []IProtoType{file}}
	folder := &Folder{Name: "FolderText", Children: []IProtoType{fold, file, &File{Name: "Demo.txt"}}}
	circle := &Circle{Radius: 5.5}
	square := &Square{Side: 15.5}
	r.AddRegistry(file).AddRegistry(folder).AddRegistry(circle).AddRegistry(square)
}

func TestPrototype() {
	r := NewRegistry()
	LoadRegistry(r)
	// for _, protoType := range r.List {
	// 	protoType.Clone().Print("-")
	// }
	r.List[ProtoTypeFolder].Clone().Print("-")
}
