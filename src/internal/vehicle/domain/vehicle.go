package domain

import (
	"regexp"
)

type Vehicle struct {
	Id               uint
	DeliveryPersonId uint
	Color            string
	Type             string
	Plate            string
	CardID           string
}

var (
	cardIdExpr = regexp.MustCompile(`^[A-Z]{3}-\d{2}-\d{2}$`)
	plateExpr  = regexp.MustCompile(`[A-Z]{3}\d{3}[A-Z\d]{1}`)
)

func NewVehicle(color, vtype, plate, cardid string, deliverypersonid uint) (*Vehicle, error) {
	if color == "" {
		return nil, ErrInvalidColor
	}
	if vtype != "motorcycle" && vtype != "bicycle" {
		return nil, ErrInvalidType
	}
	if deliverypersonid == 0 {
		return nil, ErrInvalidDeliveryPersonID
	}
	if cardid != "" && !cardIdExpr.MatchString(cardid) {
		return nil, ErrInvalidCardID
	}
	if plate != "" && !plateExpr.MatchString(plate) {
		return nil, ErrInvalidPlate
	}

	return &Vehicle{
		Color:            color,
		Type:             vtype,
		Plate:            plate,
		CardID:           cardid,
		DeliveryPersonId: deliverypersonid,
	}, nil
}

func (v *Vehicle) Update(color, vtype, plate, cardid string, deliverypersonid uint) error {
	if color != "" {
		v.Color = color
	}
	if vtype != "" {
		v.Type = vtype
	}
	if plate != "" {
		if !plateExpr.MatchString(plate) {
			return ErrInvalidPlate
		}
		v.Plate = plate
	}
	if cardid != "" {
		if !cardIdExpr.MatchString(cardid) {
			return ErrInvalidCardID
		}
		v.CardID = cardid
	}
	if deliverypersonid != 0 {
		v.DeliveryPersonId = deliverypersonid
	}

	return nil
}
