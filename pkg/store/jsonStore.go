package store

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/renzobalbo/goWeb/internal/domain"
)

type Storage interface {
	GetAll() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	Create(product domain.Product) error
	Update(product domain.Product) error
	UpdatePrice(product domain.Product, price float64) (domain.Product, error)
	Delete(id int) error
	saveProducts(products []domain.Product) error
	loadProducts() ([]domain.Product, error)
}

type storage struct {
	json string
}

func (s *storage) loadProducts() ([]domain.Product, error) {
	filename := os.Getenv("DB_FILE_NAME")

	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var productsList []domain.Product

	err = json.Unmarshal(file, &productsList)
	if err != nil {
		panic(err)
	}

	return productsList, nil
}

func (s *storage) saveProducts(products []domain.Product) error {
	bytes, err := json.Marshal(products)
	if err != nil {
		return err
	}
	return os.WriteFile(s.json, bytes, 0644)
}

func NewStorage(path string) Storage {
	return &storage{
		json: path,
	}
}

func (s *storage) GetAll() ([]domain.Product, error) {
	products, err := s.loadProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *storage) GetById(id int) (domain.Product, error) {
	products, err := s.loadProducts()
	if err != nil {
		return domain.Product{}, err
	}
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")
}

func (s *storage) Create(product domain.Product) error {
	products, err := s.loadProducts()
	if err != nil {
		return err
	}
	product.Id = len(products) + 1
	products = append(products, product)
	return s.saveProducts(products)
}

func (s *storage) Update(product domain.Product) error {
	products, err := s.loadProducts()
	if err != nil {
		return err
	}
	for i, p := range products {
		if p.Id == product.Id {
			products[i] = product
			return s.saveProducts(products)
		}
	}
	return errors.New("product not found")
}

func (s *storage) UpdatePrice(product domain.Product, price float64) (domain.Product, error) {
	products, err := s.loadProducts()
	if err != nil {
		return domain.Product{}, err
	}
	for i, p := range products {
		if p.Id == product.Id {
			products[i].Price = price
			return products[i], s.saveProducts(products)
		}
	}
	return domain.Product{}, errors.New("product not found")
}

func (s *storage) Delete(id int) error {
	products, err := s.loadProducts()
	if err != nil {
		return err
	}
	for i, p := range products {
		if p.Id == id {
			products = append(products[:i], products[i+1:]...)
			return s.saveProducts(products)
		}
	}
	return errors.New("product not found")
}
