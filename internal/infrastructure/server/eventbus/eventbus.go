package eventbus

import "github.com/asaskevich/EventBus"

func NewEventBus() EventBus.Bus {
	return EventBus.New()
}
