package domain

type CommerceCategory struct {
	Id   uint
	Name string
}

func NewCommerceCategory(name string) (*CommerceCategory, error) {
	if name == "" {
		return nil, ErrInvalidName
	}

	return &CommerceCategory{
		Name: name,
	}, nil
}

func (cc *CommerceCategory) Update(name string) error {
	if name != "" {
		cc.Name = name
	}

	return nil
}
