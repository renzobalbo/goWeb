package product

import (
	"errors"

	"github.com/renzobalbo/goWeb/internal/domain"
	"github.com/renzobalbo/goWeb/pkg/store"
)

type Repository interface {
	GetAll() []domain.Product
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) []domain.Product
	Create(p domain.Product) (domain.Product, error)
	Update(product domain.Product) (domain.Product, error)
	UpdatePrice(product domain.Product, price float64) (domain.Product, error)
	Delete(id int) error
}

type repository struct {
	storage store.Storage
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.Storage) Repository {
	return &repository{storage}
}

// GetAll devuelve todos los productos
func (r *repository) GetAll() []domain.Product {
	products, err := r.storage.GetAll()
	if err != nil {
		return []domain.Product{}
	}
	return products
}

// GetByID busca un producto por su id
func (r *repository) GetByID(id int) (domain.Product, error) {
	product, err := r.storage.GetById(id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

// SearchPriceGt busca productos por precio mayor o igual que el precio dado
func (r *repository) SearchPriceGt(price float64) []domain.Product {
	var productsGt []domain.Product
	products, err := r.storage.GetAll()
	if err != nil {
		return productsGt
	}
	for _, p := range products {
		if p.Price > price {
			productsGt = append(productsGt, p)
		}
	}
	return productsGt
}

// Create agrega un nuevo producto
func (r *repository) Create(p domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exist")
	}
	err := r.storage.Create(p)
	if err != nil {
		return domain.Product{}, errors.New("error creating product")
	}
	return p, nil
}

// validateCodeValue valida que el codigo no exista en la lista de productos
func (r *repository) validateCodeValue(codeValue string) bool {
	products, err := r.storage.GetAll()
	if err != nil {
		return false
	}
	for _, product := range products {
		if product.CodeValue == codeValue {
			return false
		}
	}
	return true
}

// update actualiza todos los campos de un producto
func (r *repository) Update(product domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(product.CodeValue) {
		return domain.Product{}, errors.New("code value already exist")
	}
	err := r.storage.Update(product)
	if err != nil {
		return domain.Product{}, errors.New("error updating the product")
	}
	return product, nil
}

// updatePrice actualiza el precio de un producto
func (r *repository) UpdatePrice(product domain.Product, price float64) (domain.Product, error) {
	p, err := r.storage.UpdatePrice(product, price)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
