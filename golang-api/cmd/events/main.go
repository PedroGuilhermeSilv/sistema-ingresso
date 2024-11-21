package events

import (
	"database/sql"
	"net/http"

	httpHandler "github.com/PedroGuilhermeSilv/sistema-ingresso/golang/internal/events/infra/http"
	"github.com/PedroGuilhermeSilv/sistema-ingresso/golang/internal/events/infra/repository"
	"github.com/PedroGuilhermeSilv/sistema-ingresso/golang/internal/events/infra/service"
	"github.com/PedroGuilhermeSilv/sistema-ingresso/golang/internal/events/usecase"
)

func main() {
	db, err := sql.Open("mysql", "test_user:test_password@tcp(localhost:3306)/test_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventRepo, err := repository.NewMysqlEventRepository(db)
	if err != nil {
		panic(err)
	}

	partnerBaseURLs := map[int]string{
		1: "http://localhost:3000",
		2: "http://localhost:3001",
	}

	partnerFactory := service.NewPartnerFactory(partnerBaseURLs)

	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)
	buyTicketsUseCase := usecase.NewBuyTicketsUseCase(eventRepo, partnerFactory)

	EventsHandler := httpHandler.NewEventsHandler(
		listEventsUseCase,
		getEventUseCase,
		listSpotsUseCase,
		buyTicketsUseCase,
	)
	r := http.NewServeMux()
	r.HandleFunc("/events", EventsHandler.ListEvents)
	r.HandleFunc("/events/{eventId}", EventsHandler.GetEvent)
	r.HandleFunc("/events/{eventId}/spots", EventsHandler.ListSpots)
	r.HandleFunc("POST /checkout", EventsHandler.BuyTickets)

	http.ListenAndServe(":8080", r)
}
