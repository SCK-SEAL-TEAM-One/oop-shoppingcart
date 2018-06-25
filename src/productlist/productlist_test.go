package productlist

import (
	"product"
	"testing"
)

func Test_AddToList_Input_Products_Should_be_ArrayProduct(t *testing.T) {
	product1:= product.Product{}.SetProduct("Shoe",3000.33,"Adidas",15,"www.pic.com")
	product2 := product.Product{}.SetProduct("Shoe",4000.33,"Nike",15,"www.pic.com")
	product3 :=product.Product{}.SetProduct("Shoe",5000.33,"Reebok",15,"www.pic.com")
	
	productArray:=[]product.Product{}
	append(productArray,product1)
	append(productArray,product2)
	append(productArray,product3)
	expected := productArray

	productsList := new(ProductsList)
	productsList.addToList(product1);
	productsList.addToList(product2);
	productsList.addToList(product3);
	actual := productsList.getList();

	if expected != actual {
		t.Errorf("expect %s but got %s", len(expected), len(actual))
	}



}
