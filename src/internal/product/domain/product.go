package domain

type Product struct {
	Id          uint
	CommerceId  uint
	Name        string
	Price       float32
	Description string
	Img         string
}

func NewProduct(name, description, img string, price float32, commerceid uint) (*Product, error) {
	if name == "" {
		return nil, ErrInvalidName
	}
	if description == "" {
		return nil, ErrInvalidDescription
	}
	if img == "" {
		return nil, ErrInvalidImg
	}
	if commerceid == 0 {
		return nil, ErrInvalidCommerceID
	}

	return &Product{
		Name:        name,
		Price:       price,
		Description: description,
		Img:         img,
		CommerceId:  commerceid,
	}, nil
}

func (p *Product) Update(name, description string, price float32) error {
	if name != "" {
		p.Name = name
	}
	if description != "" {
		p.Description = description
	}
	if price != 0 {
		p.Price = price
	}

	return nil
}
