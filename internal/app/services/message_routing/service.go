package message_routing

import (
	"context"
	"log"

	"github.com/IASamoylov/tg_calories_observer/internal/domain"
)

type Service struct {
	requestRouting  RequestRouting
	commandRouting  CommandRouting
	messengerClient MessengerClient
	userGetter      UserGetter
}

func NewService(
	requestRouting RequestRouting,
	commandRouting CommandRouting,
	messengerClient MessengerClient,
	userGetter UserGetter,
) Service {
	return Service{
		requestRouting:  requestRouting,
		commandRouting:  commandRouting,
		messengerClient: messengerClient,
		userGetter:      userGetter,
	}
}

// Handle ...
func (svc Service) Handle(ctx context.Context, sender domain.User, message string) {
	var err error
	if sender, err = svc.userGetter.UpsertAndGet(ctx, sender); err != nil {
		svc.send(sender, "Временный сбой, повторите попытку позже")
		log.Println(err)
		return
	}

	var (
		recipient domain.User
		response  string
	)
	switch true {
	case svc.requestRouting.IsQuery(message):
		recipient, response, err = svc.requestRouting.Execute(ctx, sender, message)
	default:
		svc.send(sender, "Не известная команда")
	}

	if err != nil {
		svc.send(sender, "Временный сбой, повторите попытку позже")
		log.Println("")
	}

	svc.send(recipient, response)
}

func (svc Service) send(recipient domain.User, message string) {
	if err := svc.messengerClient.Send(recipient, message); err != nil {
		log.Println(err)
	}
}

//return types.NewResponse(domain.User{}, "Временный сбой, повторите попытку позже")
