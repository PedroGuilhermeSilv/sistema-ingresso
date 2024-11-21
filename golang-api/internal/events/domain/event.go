package domain

import (
	"time"
)

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "10"
	Rating12    Rating = "12"
	Rating14    Rating = "14"
	Rating16    Rating = "16"
	Rating18    Rating = "18"
)

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        int
	PartnerID    int
	Spots        []Spot
	Tickets      []Ticket
}

func (e *Event) AddSpot(name string) (*Spot, error) {
	spot, err := NewSpot(e, name)
	if err != nil {
		return nil, err
	}
	e.Spots = append(e.Spots, *spot)
	return spot, nil
}

func (e Event) Validate() error {
	if e.Name == "" {
		return ErrInvalidName
	}
	if e.Location == "" {
		return ErrInvalidLocation
	}
	if e.Organization == "" {
		return ErrInvalidOrganization
	}
	if e.Rating == "" {
		return ErrInvalidRating
	}
	if e.Date.IsZero() {
		return ErrInvalidDate
	}
	if e.ImageURL == "" {
		return ErrInvalidImageURL
	}
	if e.Capacity <= 0 {
		return ErrInvalidCapacity
	}
	if e.Price <= 0 {
		return ErrInvalidPrice
	}
	if e.PartnerID == 0 {
		return ErrInvalidPartnerID
	}
	if len(e.Spots) == 0 {
		return ErrInvalidSpots
	}
	if len(e.Tickets) == 0 {
		return ErrInvalidTickets
	}
	return nil
}
