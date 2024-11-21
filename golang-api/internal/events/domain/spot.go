package domain
import (
	"github.com/google/uuid"
)
type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	ID       string
	EventId  string
	Name     string
	Status   SpotStatus
	TicketID string
}

func (s *Spot) Reserve(ticketID string) error {
	if s.Status == SpotStatusSold {
		return ErrSpotAlreadySold
	}
	s.Status = SpotStatusSold
	s.TicketID = ticketID
	return nil
}

func NewSpot(event *Event, name string) (*Spot, error) {
	s := &Spot{
		ID:   uuid.New().String(),
		EventId: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}
	if err := s.Validate(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Spot) Validate() error {
	if s.Name == "" {
		return ErrInvalidSpotNameEmpty
	}
	if len(s.Name) < 2 {
		return ErrInvalidSpotNameLength
	}
	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrInvalidSpotName
	}
	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrInvalidSpotName
	}
	if s.EventId == "" {
		return ErrInvalidSpotEventID
	}
	if s.Status == "" {
		return ErrInvalidSpotStatus
	}
	if s.TicketID == "" {
		return ErrInvalidSpotTicketID
	}
	return nil

}
