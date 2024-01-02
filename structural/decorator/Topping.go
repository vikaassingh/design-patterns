package decorator

type Topping struct {
	Decorator IIcecreamDecorator
}

func (c *Topping) GetIngredientDecorate(decorator IIcecreamDecorator) IIcecreamDecorator {
	c.Decorator = decorator
	return c.Decorator
}

func (c *Topping) GetPrice() int {
	if c.Decorator == nil {
		return 10
	}

	return c.Decorator.GetPrice() + 5
}

func (c *Topping) GetDescription() string {
	if c.Decorator == nil {
		return "Topping"
	}

	return c.Decorator.GetDescription() + ", Topping"
}
