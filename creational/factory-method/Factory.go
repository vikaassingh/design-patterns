package factorymethod

type Factory struct {
	Name    string
	Address string
}

func (f *Factory) SetFactoryName(name string) {
	f.Name = name
}

func (f *Factory) SetFactoryAddress(addr string) {
	f.Address = addr
}

func (f *Factory) GetFactoryName() string {
	return f.Name
}

func (f *Factory) GetFactoryAddress() string {
	return f.Address
}
