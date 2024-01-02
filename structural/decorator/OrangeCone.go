package decorator

type OrangeCone struct {
	Decorator IIcecreamDecorator
}

func (c *OrangeCone) GetIngredientDecorate(decorator IIcecreamDecorator) IIcecreamDecorator {
	c.Decorator = decorator
	return c.Decorator
}

func (c *OrangeCone) GetPrice() int {
	if c.Decorator == nil {
		return 10
	}

	return c.Decorator.GetPrice() + 10
}

func (c *OrangeCone) GetDescription() string {
	if c.Decorator == nil {
		return "Orange Cone"
	}

	return c.Decorator.GetDescription() + ", Orange Cone"
}
