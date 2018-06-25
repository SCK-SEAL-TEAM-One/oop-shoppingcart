package product

import "testing"

func Test_setProduct_Input_ProductDetail_Should_Be_Set(t *testing.T) {
	inputProduct := "Soap"
	ProductPrice := 45.8
	ProductBrand := "Lux"
	ProductQuantity := 10
	ProductPicture := "www.picture.com"
	expected := Product{
		productName:     inputProduct,
		productPrice:    ProductPrice,
		productBrand:    ProductBrand,
		productQuantity: ProductQuantity,
		productPicture:  ProductPicture,
	}

	actualProduct := new(Product)
	actualProduct.setProduct(inputProduct, ProductPrice, ProductBrand, ProductQuantity, ProductPicture)
	actual := actualProduct.getProduct()

	if expected.productName != actual.productName {
		t.Errorf("expect %s but got %s", expected.productName, actual.productName)
	}

	if expected.productPrice != actual.productPrice {
		t.Errorf("expect %v but got %v", expected.productPrice, actual.productPrice)
	}
}
