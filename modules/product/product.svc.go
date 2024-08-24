package product

import (
	productdto "app/modules/product/dto"
	"errors"
	"strconv"
)

type ProductService struct {
}

func newService() *ProductService {
	return nil
}

type ProductEntity struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	IsActived   bool    `json:"isActived"`
}

var products = []ProductEntity{}

func (s *ProductService) Create(req productdto.CreateProductRequest) (*ProductEntity, error) {

	if len(products) == 0 {
		products = append(products, ProductEntity{
			ID:          1,
			Name:        req.Name,
			Price:       req.Price,
			Description: req.Description,
			IsActived:   req.IsActived,
		})
	} else {
		products = append(products, ProductEntity{
			ID:          products[len(products)-1].ID + 1,
			Name:        req.Name,
			Price:       req.Price,
			Description: req.Description,
			IsActived:   req.IsActived,
		})
	}

	lastProduct := products[len(products)-1]

	return &lastProduct, nil
}

func (s *ProductService) Update(id productdto.GetProductByIDRequest, req productdto.UpdateProductRequest) (*ProductEntity, error) {

	productId, _ := strconv.Atoi(id.ID)

	var index *int
	for i, r := range products {
		if r.ID == productId {
			index = &i
			break
		}
	}

	if index == nil {
		return nil, errors.New("Product not found")
	}

	products[*index].Name = req.Name
	products[*index].Price = req.Price
	products[*index].Description = req.Description
	products[*index].IsActived = req.IsActived

	return &products[*index], nil
}

func (s *ProductService) Delete(id productdto.GetProductByIDRequest) (*ProductEntity, error) {

	productId, _ := strconv.Atoi(id.ID)

	var index *int
	for i, r := range products {
		if r.ID == productId {
			index = &i
			break
		}
	}

	if index == nil {
		return nil, errors.New("Product not found")
	}

	products = append(products[:*index], products[*index+1:]...)

	return nil, nil
}

func (s *ProductService) Get(id productdto.GetProductByIDRequest) (*ProductEntity, error) {

	productId, _ := strconv.Atoi(id.ID)

	var index *int
	for i, r := range products {
		if r.ID == productId {
			index = &i
			break
		}
	}

	if index == nil {
		return nil, errors.New("Product not found")
	}

	return &products[*index], nil
}

func (s *ProductService) List() ([]ProductEntity, error) {

	return products, nil
}
