package decorator

type Peanuts struct {
	Decorator IIcecreamDecorator
}

func (c *Peanuts) GetIngredientDecorate(decorator IIcecreamDecorator) IIcecreamDecorator {
	c.Decorator = decorator
	return c.Decorator
}

func (c *Peanuts) GetPrice() int {
	if c.Decorator == nil {
		return 10
	}

	return c.Decorator.GetPrice() + 5
}

func (c *Peanuts) GetDescription() string {
	if c.Decorator == nil {
		return "Peanuts"
	}

	return c.Decorator.GetDescription() + ", Peanuts"
}
