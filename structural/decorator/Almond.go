package decorator

type Almond struct {
	Decorator IIcecreamDecorator
}

func (c *Almond) GetIngredientDecorate(decorator IIcecreamDecorator) IIcecreamDecorator {
	c.Decorator = decorator
	return c.Decorator
}

func (c *Almond) GetPrice() int {
	if c.Decorator == nil {
		return 10
	}

	return c.Decorator.GetPrice() + 5
}

func (c *Almond) GetDescription() string {
	if c.Decorator == nil {
		return "Almond"
	}

	return c.Decorator.GetDescription() + ", Almond"
}
