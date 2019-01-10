package eventhus

// EventStore saves the events from an aggregate0000
type EventStore interface {
	Save(events []Event, version int) error
	SafeSave(events []Event, version int) error
	Load(aggregateID string) ([]Event, error)
}
