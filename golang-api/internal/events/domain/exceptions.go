package domain

import "errors"

var (
	ErrInvalidName         = errors.New("invalid name")
	ErrInvalidLocation     = errors.New("invalid location")
	ErrInvalidOrganization = errors.New("invalid organization")
	ErrInvalidRating       = errors.New("invalid rating")
	ErrInvalidDate         = errors.New("invalid date")
	ErrInvalidImageURL     = errors.New("invalid image URL")
	ErrInvalidCapacity     = errors.New("invalid capacity")
	ErrInvalidPrice        = errors.New("invalid price")
	ErrInvalidPartnerID    = errors.New("invalid partner ID")
	ErrInvalidSpots        = errors.New("invalid spots")
	ErrInvalidTickets      = errors.New("invalid tickets")
)

var (
	ErrInvalidSpotNameEmpty  = errors.New("spot name cannot be empty")
	ErrInvalidSpotNameLength = errors.New("spot name must be at least 2 characters long")
	ErrInvalidSpotName       = errors.New("invalid spot name")
	ErrInvalidSpotEventID    = errors.New("invalid spot event ID")
	ErrInvalidSpotStatus     = errors.New("invalid spot status")
	ErrInvalidSpotTicketID   = errors.New("invalid spot ticket ID")
	ErrSpotAlreadySold       = errors.New("spot already sold")
	ErrInvalidQuantitySpots  = errors.New("invalid quantity of spots")
	ErrSpotNotFound          = errors.New("spot not found")
)

var (
	ErrInvalidTicketPrice = errors.New("invalid ticket price")
	ErrInvalidTicketType  = errors.New("invalid ticket type")
	ErrInvalidSpot        = errors.New("invalid spot")
	ErrInvalidEventID     = errors.New("invalid event ID")
	ErrEventNotFound      = errors.New("event not found")
	ErrInvalidEventId    = errors.New("invalid event ID")
)
