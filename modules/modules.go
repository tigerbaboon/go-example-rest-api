package modules

import "app/modules/product"

type Modules struct {
	Product *product.ProductModule
}

func Get() *Modules {
	return &Modules{
		Product: product.New(),
	}
}
