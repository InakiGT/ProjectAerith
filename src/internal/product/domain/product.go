package domain

type Product struct {
	Id          uint
	Name        string
	Price       float32
	Description string
}

func NewProduct(name, description string, price float32) (*Product, error) {
	if name == "" {
		return nil, ErrInvalidName
	}
	if description == "" {
		return nil, ErrInvalidDescription
	}

	return &Product{
		Name:        name,
		Price:       price,
		Description: description,
	}, nil
}
