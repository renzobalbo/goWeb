package product

import (
	"errors"
	"fmt"

	"github.com/renzobalbo/goWeb/internal/domain"
)

type Repository interface {
	GetAll() []domain.Product
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) []domain.Product
	Create(p domain.Product) (domain.Product, error)
	Update(id int, name string, quantity int, codeValue string, isPublished bool, expiration string, price float64) (domain.Product, error)
	UpdatePrice(id int, price float64) (domain.Product, error)
	Delete(id int) error
}

type repository struct {
	list []domain.Product
}

// NewRepository crea un nuevo repositorio
func NewRepository(list []domain.Product) Repository {
	return &repository{list}
}

// GetAll devuelve todos los productos
func (r *repository) GetAll() []domain.Product {
	return r.list
}

// GetByID busca un producto por su id
func (r *repository) GetByID(id int) (domain.Product, error) {
	for _, product := range r.list {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")

}

// SearchPriceGt busca productos por precio mayor o igual que el precio dado
func (r *repository) SearchPriceGt(price float64) []domain.Product {
	var products []domain.Product
	for _, product := range r.list {
		if product.Price > price {
			products = append(products, product)
		}
	}
	return products
}

// Create agrega un nuevo producto
func (r *repository) Create(p domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	p.Id = len(r.list) + 1
	r.list = append(r.list, p)
	return p, nil
}

// validateCodeValue valida que el codigo no exista en la lista de productos
func (r *repository) validateCodeValue(codeValue string) bool {
	for _, product := range r.list {
		if product.CodeValue == codeValue {
			return false
		}
	}
	return true
}

// update actualiza todos los campos de un producto
func (r *repository) Update(id int, name string, quantity int, codeValue string, isPublished bool, expiration string, price float64) (domain.Product, error) {
	p := domain.Product{Name: name, Quantity: quantity, CodeValue: codeValue, IsPublished: isPublished, Expiration: expiration, Price: price}
	updated := false
	for i := range r.list {
		if r.list[i].Id == id {
			p.Id = id
			r.list[i] = p
			updated = true
		}
	}
	if !updated {
		return domain.Product{}, fmt.Errorf("couldn't find a product with the id: %d", id)
	}
	return p, nil
}

// updatePrice actualiza el precio de un producto
func (r *repository) UpdatePrice(id int, price float64) (domain.Product, error) {
	var p domain.Product
	updated := false
	for i := range r.list {
		if r.list[i].Id == id {
			r.list[i].Price = price
			updated = true
			p = r.list[i]
		}
	}
	if !updated {
		return domain.Product{}, fmt.Errorf("couldn't find a product with the id: %d", id)
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range r.list {
		if r.list[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("couldn't find a product with the id: %d", id)
	}
	r.list = append(r.list[:index], r.list[index+1:]...)
	return nil
}
