package main

import (
	"fmt"
	"log"

	cpc "github.com/bukalapak/core-product-client/services"
)

func main() {
	c := cpc.NewClient(nil)
	prod, err := c.CoreProduct.ShowProduct(1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(prod)
}
