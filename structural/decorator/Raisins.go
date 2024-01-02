package decorator

type Raisins struct {
	Decorator IIcecreamDecorator
}

func (c *Raisins) GetIngredientDecorate(decorator IIcecreamDecorator) IIcecreamDecorator {
	c.Decorator = decorator
	return c.Decorator
}

func (c *Raisins) GetPrice() int {
	if c.Decorator == nil {
		return 10
	}

	return c.Decorator.GetPrice() + 5
}

func (c *Raisins) GetDescription() string {
	if c.Decorator == nil {
		return "Raisins"
	}

	return c.Decorator.GetDescription() + ", Raisins"
}
