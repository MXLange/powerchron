package powerchron

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

type Gap struct {
	Time int64
	Type time.Duration
}

// Start starts the scheduler
// This scheduler will run the function indefinitely waiting for the time defined in the gap to run again
func (g *Gap) Start(fn interface{}, params ...any) {
	newGap(g, fn, params...)
}

func newGap(gap *Gap, fn interface{}, params ...any) {

	messages, err := gap.validate(fn)
	if err != nil {
		log.Println("Error validating scheduler")
		for _, message := range messages {
			log.Println(message)
		}
		panic(err.Error())
	}

	timeToWait := time.Duration(gap.Time) * gap.Type

	// Prepare parameters
	in := make([]reflect.Value, len(params))
	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}

	fnValue := reflect.ValueOf(fn)

	// Call the function imeadiatly
	fnValue.Call(in)

	for {

		time.Sleep(timeToWait)

		log.Printf("Running task at %v\n", time.Now().Format("2006-01-02 15:04:05"))
		fnValue.Call(in)
	}
}

func (g *Gap) validate(fn interface{}) (messages []string, err error) {
	err = validateFunc(fn)
	if err != nil {
		messages = append(messages, err.Error())
	}

	if g.Time <= 0 {
		messages = append(messages, "time must be greater than 0")
	}

	if len(messages) > 0 {
		return messages, fmt.Errorf("error validating scheduler")
	}

	return nil, nil
}
