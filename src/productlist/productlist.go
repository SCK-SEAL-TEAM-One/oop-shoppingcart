package productlist

import (
	"product"
)

type ProductsList struct {
	productsList []product.Product
}

func (p *ProductsList) addToList(product product.Product) {
	p.productsList = append(p.productsList, product)
}

func (p *ProductsList) getList(product product.Product) []product.Product {
	return p.productsList
}
