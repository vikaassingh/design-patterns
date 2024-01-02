package decorator

type Cashew struct {
	Decorator IIcecreamDecorator
}

func (c *Cashew) GetIngredientDecorate(decorator IIcecreamDecorator) IIcecreamDecorator {
	c.Decorator = decorator
	return c.Decorator
}

func (c *Cashew) GetPrice() int {
	if c.Decorator == nil {
		return 10
	}

	return c.Decorator.GetPrice() + 5
}

func (c *Cashew) GetDescription() string {
	if c.Decorator == nil {
		return "Cashew"
	}

	return c.Decorator.GetDescription() + ", Cashew"
}
