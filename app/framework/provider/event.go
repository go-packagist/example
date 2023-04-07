package provider

import (
	event "github.com/go-packagist/event/v2"
	"github.com/go-packagist/example/app/framework/container"
)

var listeners = map[string][]event.Listener{
	// "example": {
	// 	&Example1Listener{},
	// 	&Example2Listener{},
	// },
}

type DispatcherProvider struct {
	*container.Container
}

func NewDispatcherProvider(container *container.Container) *DispatcherProvider {
	return &DispatcherProvider{
		Container: container,
	}
}

func (p *DispatcherProvider) Register() {
	p.Dispatcher = event.NewDispatcher()

	for name, listener := range listeners {
		p.Dispatcher.Listen(name, listener...)
	}
}

func (p *DispatcherProvider) Boot() {

}
