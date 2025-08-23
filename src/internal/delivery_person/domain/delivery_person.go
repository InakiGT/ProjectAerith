package domain

import "time"

type Location struct {
	Latitude  float64
	Longitude float64
}

type DeliveryPerson struct {
	Id              uint
	UserId          uint
	Birthday        time.Time
	CurrentLocation Location `gorm:"type:point"`
	Status          string
	PersonalID      string
}

func NewDeliveryPerson(userid uint, birthday time.Time, personalid string) (*DeliveryPerson, error) {
	if userid == 0 {
		return nil, ErrInvalidUserID
	}
	if birthday.IsZero() {
		return nil, ErrInvalidBirthday
	}
	if personalid == "" {
		return nil, ErrInvalidPersonalID
	}

	return &DeliveryPerson{
		UserId:     userid,
		Birthday:   birthday,
		Status:     "inactive",
		PersonalID: personalid,
	}, nil
}

func (d *DeliveryPerson) Update(birthday time.Time, personalid string) error {
	if !birthday.IsZero() {
		d.Birthday = birthday
	}
	if personalid != "" {
		d.PersonalID = personalid
	}

	return nil
}

func (d *DeliveryPerson) UpdateCurrentLocation(location Location) error {
	d.CurrentLocation = location

	return nil
}

func (d *DeliveryPerson) UpdateStatus(status string) error {
	if status != "" {
		if status == "active" || status == "inactive" {
			d.Status = status
			return nil
		}
	}

	return ErrInvalidStatus
}
