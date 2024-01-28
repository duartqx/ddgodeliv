package delivery

import "reflect"

type statusChoices struct {
	Pending   uint8 // Has no Driver
	Assigned  uint8 // Assigned a Driver
	InTransit uint8
	Late      uint8
	Completed uint8
}

// Reflects upon the fields of statusChoices and builds an slice containing all
// status labels
func (sc statusChoices) getStatusChoicesLabels() *[]string {

	rfl := reflect.TypeOf(statusChoices{})

	statusLabels := make([]string, rfl.NumField())
	for i := 0; i < rfl.NumField(); i++ {
		statusLabels[i] = rfl.Field(i).Name
	}

	return &statusLabels
}

func (sc statusChoices) GetDisplay(s uint8) string {
	return (*sc.getStatusChoicesLabels())[s]
}

// Order/Values Must not be altered
const (
	pending   uint8 = 0
	assigned        = 1
	inTransit       = 2
	late            = 3
	completed       = 4
)

var StatusChoices *statusChoices = &statusChoices{
	Pending:   pending,
	Assigned:  assigned,
	InTransit: inTransit,
	Late:      late,
	Completed: completed,
}
