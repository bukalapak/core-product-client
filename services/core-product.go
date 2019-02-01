package services

import (
	"fmt"

	structs "github.com/bukalapak/core-product/handler/util"
)

type CoreProductService struct {
	client *Client
}

func (c *CoreProductService) ShowProduct(id uint64) (interface{}, error) {
	path := fmt.Sprintf("/core-product-service/show/%s", id)
	resp, err := c.client.process("GET", path, nil, structs.ShowProductResponse{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
