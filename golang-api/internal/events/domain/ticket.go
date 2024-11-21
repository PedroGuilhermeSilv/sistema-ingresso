package domain

import "github.com/google/uuid"

type TicketType string

const (
	TicketTypeHalf TicketType = "half"
	TicketTypeFull TicketType = "full"
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketType TicketType
	Price      int
}

func NewTicket(event *Event, spot *Spot, ticketType TicketType) (*Ticket, error) {
	t := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketType: ticketType,
		Price:      event.Price,
	}
	if err := t.Validate(); err != nil {
		return nil, err
	}
	t.CalculatePrice()
	return t, nil
}

func IsValidTicketType(t TicketType) bool {
	return t == TicketTypeHalf || t == TicketTypeFull
}

func (t *Ticket) CalculatePrice() {
	if t.TicketType == TicketTypeHalf {
		t.Price = t.Price / 2
	}
}

func (t *Ticket) Validate() error {
	if t.EventID == "" {
		return ErrInvalidEventID
	}
	if t.Spot == nil {
		return ErrInvalidSpot
	}
	if !IsValidTicketType(t.TicketType) {
		return ErrInvalidTicketType
	}
	if t.Price <= 0 {
		return ErrInvalidTicketPrice
	}
	return nil
}
