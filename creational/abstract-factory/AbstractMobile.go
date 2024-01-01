package abstractfactory

type AbstractMobile struct {
	Name  string
	Brand string
}

func (p *AbstractMobile) SetName(name string) {
	p.Name = name
}

func (p *AbstractMobile) SetBrand(name string) {
	p.Brand = name
}

func (p *AbstractMobile) GetName() string {
	return p.Name
}

func (p *AbstractMobile) GetBrand() string {
	return p.Brand
}
