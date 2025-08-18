package domain

import "time"

type Location struct {
	Latitude  float64
	Longitude float64
	Timestamp time.Time
}

type DeliveryPerson struct {
	Id              uint
	UserId          uint
	MainVehicleId   uint
	Birthday        time.Time
	CurrentLocation Location
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

func (d *DeliveryPerson) Update(mainvehicleid uint, birthday time.Time, personalid string) error {
	if mainvehicleid != 0 {
		d.MainVehicleId = mainvehicleid
	}
	if !birthday.IsZero() {
		d.Birthday = birthday
	}
	if personalid != "" {
		d.PersonalID = personalid
	}

	return nil
}

func (d *DeliveryPerson) UpdateCurrentLocation(location Location) error {
	if !location.Timestamp.IsZero() {
		d.CurrentLocation = location
		return nil
	}

	return ErrInvalidLocation
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
