package domain

type ClientCard struct {
	Id                uint
	ClientId          uint
	Provider          string
	ExpYear           string
	ExpMonth          string
	Last4             string
	Brand             string
	ServiceCustomerId string
}

func NewClientCard(userid uint, provider, expyear, expmonth, last4, brand, servicecustomerid string) (*ClientCard, error) {
	if userid == 0 {
		return nil, ErrInvalidUserID
	}
	if provider == "" {
		return nil, ErrInvalidProvider
	}
	if expyear == "" {
		return nil, ErrInvalidExpYear
	}
	if expmonth == "" {
		return nil, ErrInvalidExpMonth
	}
	if last4 == "" {
		return nil, ErrInvalidLast4
	}
	if brand == "" {
		return nil, ErrInvalidBrand
	}
	if servicecustomerid == "" {
		return nil, ErrInvalidServiceCustomerID
	}

	return &ClientCard{
		ClientId:          userid,
		Provider:          provider,
		ExpYear:           expyear,
		ExpMonth:          expmonth,
		Last4:             last4,
		Brand:             brand,
		ServiceCustomerId: servicecustomerid,
	}, nil
}

func (c *ClientCard) Update(provider, expyear, expmonth, last4, brand, servicecustomerid string) error {
	if provider != "" {
		c.Provider = provider
	}
	if expyear != "" {
		c.ExpYear = expmonth
	}
	if expmonth != "" {
		c.ExpMonth = expmonth
	}
	if last4 != "" {
		c.Last4 = last4
	}
	if brand != "" {
		c.Brand = brand
	}
	if servicecustomerid != "" {
		c.ServiceCustomerId = servicecustomerid
	}

	return nil
}
