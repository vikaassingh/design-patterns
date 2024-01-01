package prototype

type IProtoType interface {
	GetId() ProtoType
	Print(string)
	Clone() IProtoType
}
