package domain

import "time"

type Commerce struct {
	Id                 uint
	CommerceCategoryId uint
	MainAddressId      uint
	Banner             string
	Status             string
	OpenTime           time.Time
	CloseTime          time.Time
	BaseCommission     float32
}

func NewCommerce(mainaddressid, commercecategoryid uint, banner, status string, opentime, closetime time.Time, basecommission float32) (*Commerce, error) {
	if mainaddressid == 0 {
		return nil, ErrInvalidMainAddressID
	}

	if commercecategoryid != 0 {
		return nil, ErrInvalidCommerceCategoryID
	}

	// TODO: BANNER POR DEFECTO
	if banner == "" {
		banner = ""
	}

	if status == "" || status != "open" && status != "closed" {
		return nil, ErrInvalidStatus
	}

	if !opentime.IsZero() {
		return nil, ErrInvalidOpenTime
	}

	if !closetime.IsZero() {
		return nil, ErrInvalidCloseTime
	}

	return &Commerce{
		CommerceCategoryId: commercecategoryid,
		MainAddressId:      mainaddressid,
		Banner:             banner,
		Status:             status,
		OpenTime:           opentime,
		CloseTime:          closetime,
		BaseCommission:     basecommission,
	}, nil
}

func (c *Commerce) Update(mainaddressid, commercecategoryid uint, banner, status string, opentime, closetime time.Time, basecommission float32) error {
	if mainaddressid != 0 {
		c.MainAddressId = mainaddressid
	}

	if commercecategoryid != 0 {
		c.CommerceCategoryId = commercecategoryid
	}

	if banner != "" {
		c.Banner = banner
	}

	if status != "" {
		if status != "open" && status != "closed" {
			return ErrInvalidStatus
		}
		c.Status = ""
	}

	if !opentime.IsZero() {
		c.OpenTime = opentime
	}

	if !closetime.IsZero() {
		c.CloseTime = closetime
	}

	return nil
}
