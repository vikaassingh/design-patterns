package prototype

type Registry struct {
	List map[ProtoType]IProtoType
}

func NewRegistry() *Registry {
	return &Registry{List: make(map[ProtoType]IProtoType, 0)}
}

func (r *Registry) AddRegistry(protoType IProtoType) *Registry {
	r.List[protoType.GetId()] = protoType
	return r
}

func (r *Registry) RemoveRegistry(protoType ProtoType) {
	delete(r.List, protoType)
}
