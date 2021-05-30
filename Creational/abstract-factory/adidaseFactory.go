package main

type adidaseFactory struct{
}

func (n *adidaseFactory)makeShoe()iShoe{
	return &adidaseShoe{
		shoe: shoe{
			logo: "addidas",
			size: 43,
		},
	}
}

func (n *adidaseFactory)makeShirt()iShirt{
	return &addidaseShirt{
		shirt: shirt{
			logo: "addidas",
			size: 33,
		},
	}
}
