package powerchron

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

type Week struct {
	All       WeekGeneral
	Monday    WeekGeneral
	Tuesday   WeekGeneral
	Wednesday WeekGeneral
	Thursday  WeekGeneral
	Friday    WeekGeneral
	Saturday  WeekGeneral
	Sunday    WeekGeneral
	weekMap   map[time.Weekday]map[int]bool
}

type WeekGeneral struct {
	All   bool
	Hours Hour
}

func (w *Week) Start(fn interface{}, params ...any) {
	newWeek(w, fn, params...)
}

func newWeek(w *Week, fn interface{}, params ...any) {
	messages, err := w.validate(fn)
	if err != nil {
		log.Println("Error validating week")
		for _, message := range messages {
			log.Println(message)
		}
		panic(err.Error())
	}

	w.toMaps()
	err = w.validateMap()
	if err != nil {
		log.Println("Error validating week")
		panic(err.Error())
	}

	in := make([]reflect.Value, len(params))
	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}

	fnValue := reflect.ValueOf(fn)
	// Call the function imeadiatly
	fnValue.Call(in)

	for {

		now := time.Now()

		next := now.Truncate(time.Hour).Add(time.Hour)
		sleepDuration := time.Until(next)

		time.Sleep(sleepDuration)

		today := time.Now().Weekday()

		if hours, ok := w.weekMap[today]; ok {
			hour := time.Now().Hour()
			if hours[hour] {
				fnValue.Call(in)
			}
		}
	}

}

func (w *Week) validateMap() error {

	isValid := false

	for _, value := range w.weekMap {
		for _, v := range value {
			if v {
				isValid = true
			}
		}
	}

	if !isValid {
		return fmt.Errorf("no valid hours set")
	}

	return nil
}

func (w *Week) toMaps() {

	w.weekMap = make(map[time.Weekday]map[int]bool)

	if w.All.All {
		w.weekMap[time.Monday] = w.All.Hours.toMap()
		w.weekMap[time.Tuesday] = w.All.Hours.toMap()
		w.weekMap[time.Wednesday] = w.All.Hours.toMap()
		w.weekMap[time.Thursday] = w.All.Hours.toMap()
		w.weekMap[time.Friday] = w.All.Hours.toMap()
		w.weekMap[time.Saturday] = w.All.Hours.toMap()
		w.weekMap[time.Sunday] = w.All.Hours.toMap()
	}

	if w.Monday.All {
		w.weekMap[time.Monday] = w.Monday.Hours.toMap()
	}

	if w.Tuesday.All {
		w.weekMap[time.Tuesday] = w.Tuesday.Hours.toMap()
	}

	if w.Wednesday.All {
		w.weekMap[time.Wednesday] = w.Wednesday.Hours.toMap()
	}

	if w.Thursday.All {
		w.weekMap[time.Thursday] = w.Thursday.Hours.toMap()
	}

	if w.Friday.All {
		w.weekMap[time.Friday] = w.Friday.Hours.toMap()
	}

	if w.Saturday.All {
		w.weekMap[time.Saturday] = w.Saturday.Hours.toMap()
	}

	if w.Sunday.All {
		w.weekMap[time.Sunday] = w.Sunday.Hours.toMap()
	}
}

func (w *Week) validate(fn interface{}) (messages []string, err error) {
	err = validateFunc(fn)
	if err != nil {
		messages = append(messages, err.Error())
	}

	err = w.validateAllWeek()
	if err != nil {
		messages = append(messages, err.Error())
	}

	if w.All.All {
		messages = append(messages, validate(w.All.All, w.All.Hours.Hours, 0, 23, "All")...)
	}

	if w.Monday.All {
		messages = append(messages, validate(w.Monday.All, w.Monday.Hours.Hours, 0, 23, "Monday")...)
	}

	if w.Tuesday.All {
		messages = append(messages, validate(w.Tuesday.All, w.Tuesday.Hours.Hours, 0, 23, "Tuesday")...)
	}

	if w.Wednesday.All {
		messages = append(messages, validate(w.Wednesday.All, w.Wednesday.Hours.Hours, 0, 23, "Wednesday")...)
	}

	if w.Thursday.All {
		messages = append(messages, validate(w.Thursday.All, w.Thursday.Hours.Hours, 0, 23, "Thursday")...)
	}

	if w.Friday.All {
		messages = append(messages, validate(w.Friday.All, w.Friday.Hours.Hours, 0, 23, "Friday")...)
	}

	if w.Saturday.All {
		messages = append(messages, validate(w.Saturday.All, w.Saturday.Hours.Hours, 0, 23, "Saturday")...)
	}

	if w.Sunday.All {
		messages = append(messages, validate(w.Sunday.All, w.Sunday.Hours.Hours, 0, 23, "Sunday")...)
	}

	if len(messages) > 0 {
		return messages, fmt.Errorf("invalid parameters")
	}

	return nil, nil
}

func (w *Week) validateAllWeek() error {

	if w.All.All {
		if w.Monday.All || w.Tuesday.All || w.Wednesday.All || w.Thursday.All || w.Friday.All || w.Saturday.All || w.Sunday.All {
			return fmt.Errorf("just set All param as true or specify days list")
		}
	} else {
		if !w.Monday.All && !w.Tuesday.All && !w.Wednesday.All && !w.Thursday.All && !w.Friday.All && !w.Saturday.All && !w.Sunday.All {
			return fmt.Errorf("if All.All param is false, specify at least one day as true")
		}
	}
	return nil
}
