package main

type adidas struct {
}

type adidasShoe struct {
	shoe
}

type adidasShort struct {
	short
}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *adidas) makeShort() iShort {
	return &adidasShort{
		short: short{
			logo: "adidas",
			size: 14,
		},
	}
}
