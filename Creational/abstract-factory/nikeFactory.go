package main

type nikeFactory struct{
}

func (n *nikeFactory)makeShoe()iShoe{
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 43,
		},
	}
}

func (n *nikeFactory)makeShirt()iShirt{
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 33,
		},
	}
}
