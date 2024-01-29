package order

func (o *Order[X]) Element() X {
	return o.x
}

func (o *Order[X]) Equal(x X) bool {
	return o.Element() == x
}

func (o *Order[X]) LessThan(x X) bool {
	return o.Element() < x
}
