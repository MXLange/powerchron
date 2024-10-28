package powerchron

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

type Scheduler struct {
	Months  Month
	Days    Day
	Hours   Hour
	Minutes Minute
}

type Month struct {
	All    bool
	Months []int
}

type Day struct {
	All  bool
	Days []int
}

type Hour struct {
	All   bool
	Hours []int
}

type Minute struct {
	All     bool
	Minutes []int
}

type schedulerMaps struct {
	Months  map[int]bool
	Days    map[int]bool
	Hours   map[int]bool
	Minutes map[int]bool
}

func (s *Scheduler) Start(fn interface{}, params ...any) {
	newScheduler(s, fn, params...)
}

// newScheduler creates a new Scheduler
// months, days, hours and minutes are slices of integers that represent the time when the function will be executedxxz
func newScheduler(scheduler *Scheduler, fn interface{}, params ...any) {
	messages, err := scheduler.validate(fn)
	if err != nil {
		log.Println("Error validating scheduler")
		for _, message := range messages {
			log.Println(message)
		}
		panic(err.Error())
	}

	sch := scheduler.toMaps()

	// Prepare parameters
	in := make([]reflect.Value, len(params))
	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}

	fnValue := reflect.ValueOf(fn)
	// Call the function imeadiatly
	fnValue.Call(in)

	for {
		now := time.Now()

		next := now.Truncate(time.Minute).Add(time.Minute)
		sleepDuration := time.Until(next)

		time.Sleep(sleepDuration)

		month := int(now.Month())
		day := now.Day()
		hour := now.Hour()
		minute := now.Minute()

		_, m := sch.Months[month]
		_, d := sch.Days[day]
		_, h := sch.Hours[hour]
		_, mi := sch.Minutes[minute]

		if m && d && h && mi {
			log.Printf("Running task at %v\n", time.Now().Format("2006-01-02 15:04:05"))
			fnValue.Call(in)
		}
	}

}

func validateFunc(fn interface{}) error {
	fnValue := reflect.ValueOf(fn)

	//Verify if fn is a function
	if fnValue.Kind() != reflect.Func {
		return fmt.Errorf("task is not a function")
	}

	return nil
}

func (s *Scheduler) validate(fn interface{}) (messages []string, err error) {

	err = validateFunc(fn)
	if err != nil {
		messages = append(messages, err.Error())
	}

	messages = append(messages, validate(s.Months.All, s.Months.Months, 1, 12, "Month")...)
	messages = append(messages, validate(s.Days.All, s.Days.Days, 1, 31, "Day")...)
	messages = append(messages, validate(s.Hours.All, s.Hours.Hours, 0, 23, "Hour")...)
	messages = append(messages, validate(s.Minutes.All, s.Minutes.Minutes, 0, 59, "Minute")...)

	if len(messages) > 0 {
		return messages, fmt.Errorf("invalid parameters")
	}

	return nil, nil
}

func validate(all bool, list []int, fisrtNum, lastNum int, name string) (messages []string) {
	if (all && len(list) > 0) || (!all && len(list) == 0) {
		messages = append(messages, fmt.Sprintf("just set %ss All param as true or specify %ss list", name, name))
	} else if all {
		for _, num := range list {
			if num < fisrtNum || num > lastNum {
				messages = append(messages, fmt.Sprintf("%s %d is invalid", name, num))
			}

		}
	} else {
		for _, num := range list {
			if num < fisrtNum || num > lastNum {
				messages = append(messages, fmt.Sprintf("%s %d is invalid", name, num))
			}
		}
	}
	return messages
}

func (s *Scheduler) toMaps() schedulerMaps {
	sch := schedulerMaps{
		Months:  getMap(s.Months.All, s.Months.Months, 1, 12),
		Days:    getMap(s.Days.All, s.Days.Days, 1, 31),
		Hours:   getMap(s.Hours.All, s.Hours.Hours, 0, 23),
		Minutes: getMap(s.Minutes.All, s.Minutes.Minutes, 0, 59),
	}
	return sch
}

func getMap(all bool, list []int, fisrtNum, lastNum int) map[int]bool {
	m := make(map[int]bool)
	if all {
		for i := fisrtNum; i <= lastNum; i++ {
			m[i] = true
		}
		return m
	} else {
		for _, num := range list {
			m[num] = true
		}
	}
	return m
}

func (h *Hour) toMap() map[int]bool {
	m := make(map[int]bool)
	if h.All {
		for i := 0; i < 24; i++ {
			m[i] = true
		}
	} else {
		for _, hour := range h.Hours {
			m[hour] = true
		}
	}
	return m
}
