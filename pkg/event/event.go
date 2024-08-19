package event

type Event struct {
	Name string
	Data interface{}
}

type EventListener interface {
	OnEvent(e Event)
}

type EventManager struct {
	listeners map[string][]EventListener
}

func NewEventManager() *EventManager {
	return &EventManager{
		listeners: make(map[string][]EventListener),
	}
}

func (em *EventManager) RegisterListener(eventName string, listener EventListener) {
	em.listeners[eventName] = append(em.listeners[eventName], listener)
}

func (em *EventManager) TriggerEvent(e Event) {
	if listeners, found := em.listeners[e.Name]; found {
		for _, listener := range listeners {
			listener.OnEvent(e)
		}
	}
}
