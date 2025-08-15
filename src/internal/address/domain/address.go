package domain

type Address struct {
	Id         uint
	City       string
	Country    string
	Number     string
	Street     string
	PostalCode string
	Cologne    string
}

func NewAddress(city, country, number, street, postalCode, cologne string) (*Address, error) {

	if city == "" {
		return nil, ErrInvalidCity
	}
	if country == "" {
		return nil, ErrInvalidCountry
	}
	if number == "" {
		return nil, ErrInvalidNumber
	}
	if street == "" {
		return nil, ErrInvalidStreet
	}
	if postalCode == "" {
		return nil, ErrInvalidPostalCode
	}
	if cologne == "" {
		return nil, ErrInvalidCologne
	}

	return &Address{
		City:       city,
		Country:    country,
		Number:     number,
		Street:     street,
		PostalCode: postalCode,
		Cologne:    cologne,
	}, nil
}

func (a *Address) Update(city, country, number, street, postalCode, cologne string) error {
	if city != "" {
		a.City = city
	}
	if country != "" {
		a.Country = country
	}
	if number != "" {
		a.Number = number
	}
	if street != "" {
		a.Street = street
	}
	if postalCode != "" {
		a.PostalCode = postalCode
	}
	if cologne != "" {
		a.Cologne = cologne
	}

	return nil
}
