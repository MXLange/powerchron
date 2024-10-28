package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/MXLange/powerchron"
)

func init() {
	// var err error = nil
	// configs.LoadConfig()

	// err = oracle.InitConnection()
	// helper.FatalError(err)

	// err = sqlserver.InitConnection()
	// helper.FatalError(err)

	// handlers.InitFetch()

}

func PrintSum() {
	time.Sleep(2 * time.Minute)
	fmt.Println("PrintSum")
}

func main() {

	s := powerchron.Scheduler{
		Months: powerchron.Month{
			All: true,
		},
		Days: powerchron.Day{
			All: true,
		},
		Hours: powerchron.Hour{
			All: true,
		},
		Minutes: powerchron.Minute{
			All: true,
		},
	}

	//Gera erro, com panic
	// s := powerchron.Scheduler{
	// 	Months: powerchron.Month{
	// 		All: true,
	// 	},
	// 	Days: powerchron.Day{
	// 		All: true,
	// 	},
	// 	Hours: powerchron.Hour{
	// 		All: true,
	// 	},
	// 	Minutes: powerchron.Minute{
	// 		Minutes: []int{50, 58, 59, 88},
	// 	},
	// }

	// w := powerchron.Week{
	// 	All: powerchron.WeekGeneral{
	// 		All: true,
	// 		Hours: powerchron.Hour{
	// 			All: true,
	// 		},
	// 	},
	// }

	//Gera erro, com panic
	// w := powerchron.Week{
	// 	All: powerchron.WeekGeneral{
	// 		All: true,
	// 	},
	// }

	//Gera erro, com panic
	// g := powerchron.Gap{
	// Time: -1,
	// Type: time.Second,
	// }

	// g := powerchron.Gap{
	// 	Time: 2,
	// 	Type: time.Second,
	// }

	//You can use the powerchron to keep the application running
	// powerchron.Start(PrintSum, 1, 2)

	//If you use with goroutines the program ends before executing the function
	//ensure that the program does not end before executing the function
	wg := sync.WaitGroup{}
	wg.Add(1)
	go s.Start(PrintSum)
	// go w.Start(PrintSum)
	// go g.Start(PrintSum)
	wg.Wait()
}
