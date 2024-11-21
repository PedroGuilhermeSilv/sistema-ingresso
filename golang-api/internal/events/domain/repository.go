package domain

type EventRepository interface {
	ListEvents() ([]Event, error)
	FindEventByID(id string) (*Event, error)
	FindSpotsByEventID(eventID string) ([]*Spot, error)
	FindSpotByName(eventID, spotName string) (*Spot, error)
	ReserveSpot(spotID string, ticketID string) error
	CreateTicket(ticket *Ticket) error
}
