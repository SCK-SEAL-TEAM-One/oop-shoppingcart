package product

type Product struct {
	productName     string
	productPrice    float64
	productBrand    string
	productQuantity int
	productPicture  string
}

func (p *Product) SetProduct(productName string, productPrice float64, productBrand string, productQuantity int, productPicture string) {
	p.productName = productName
	p.productPrice = productPrice
	p.productBrand = productBrand
	p.productQuantity = productQuantity
	p.productPicture = productPicture
}

func (p *Product) getProduct() *Product {
	return p
}
