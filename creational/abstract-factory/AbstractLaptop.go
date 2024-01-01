package abstractfactory

type AbstractLaptop struct {
	Name  string
	Brand string
}

func (p *AbstractLaptop) SetName(name string) {
	p.Name = name
}

func (p *AbstractLaptop) SetBrand(name string) {
	p.Brand = name
}

func (p *AbstractLaptop) GetName() string {
	return p.Name
}

func (p *AbstractLaptop) GetBrand() string {
	return p.Brand
}
