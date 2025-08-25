package domain

type Client struct {
	Id            uint
	UserId        uint
	MainAddressId uint
}

func NewClient(userid, mainaddressid uint) (*Client, error) {
	if userid == 0 {
		return nil, ErrInvalidUserID
	}
	if mainaddressid == 0 {
		return nil, ErrInvalidMainAddressID
	}

	return &Client{
		UserId:        userid,
		MainAddressId: mainaddressid,
	}, nil
}

func (c *Client) Update(mainaddressid uint) error {
	if mainaddressid == 0 {
		return ErrInvalidMainAddressID
	}

	c.MainAddressId = mainaddressid
	return nil
}
