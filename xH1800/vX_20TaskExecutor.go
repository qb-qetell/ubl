package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/tidwall/gjson"
	"os"
)

func vX_20TaskExecutor (clap <-chan []string, flap chan<- []string) {
	//_|^^1^^|
	//_|^^2^^|
	_x1100, _x1200 := os.ReadFile (".plate/X_20TaskExecutor/config.json")
	if _x1200 != nil {
		_x1300 := fmt.Sprintf ("{\"id\": \"0\", \"description\": \"Could not load " +
			"configuration info from plate. [%s]\"}", _x1200.Error ())
		flap <- []string {_x1300}
		return
	}
	//_|==2==|

	//_|^^2^^|
	_x1400 := int (gjson.Get (string (_x1100), "maxServicingRate").Int ())
	if _x1400 < 1 {
		_x1500 := fmt.Sprintf ("{\"id\": \"0\", \"description\": \"Maximum servicing " +
			"rate must be greater than 0; rate specified is '%d'.\"}", _x1400)
		flap <- []string {_x1500}
		return
	}
	//_|==2==|

	//_|^^2^^|
	claps := make ([]chan []string, 0)
	flaps := make ([]chan []string, 0)
	//_|==2==|

	//_|^^2^^|
	for i1600 := 1; i1600 <= _x1400; i1600 ++ {
		armClap := make (chan []string)
		armFlap := make (chan []string)

		go vX_20TaskExecutor_10Arm (armClap, armFlap)

		_x1700 := <- armFlap

		if gjson.Get (_x1700 [0], "id").String () == "0" {
			_x1800 := fmt.Sprintf ("{\"id\": \"0\", \"description\": \"Arm %d " +
				"could not start up. [%s]\"}", i1600,
				gjson.Get (_x1700 [0], "description").String ())
			flap <- []string {_x1800}
			return
		}

		claps = append (claps, armClap)
		flaps = append (flaps, armFlap)
	}

	fmt.Println (fmt.Sprintf ("-------- Component 02: All %d arm(s) started successfully.",
		_x1400))
	//_|==2==|
	
	//_|^^2^^|
	flap <- []string {"{\"id\": \"1\"}"}
	//_|==2==|
	//_|==1==|

	//_|^^1^^|
	request := queue.New ()

	for {
		//_|^^2^^|
		for {
			select {
			case _x1900 := <- clap: {
				request.Enqueue (_x1900)
				continue
			}
			default: {
				//fmt.Println ("g1170")
			}
			}
			break
		}		
		//_|==2==|

		//_|^^2^^|
		for request.Len () > 0 {
			for _, _x2000 := range claps {
				select {
				case _x2000 <- []string {"", "", "get ready"}: {
					_x2100 := request.Dequeue ()
					_x2000 <- _x2100.([]string)
				}
				default: {}
				}
				break
			}
		}
		//_|==2==|

		//_|^^2^^|
		for _, _x2000 := range flaps {
			select {
			case _x2100 := <- _x2000: {
				flap <- _x2100
			}
			default: {}
			}
		}
		//_|==2==|
	}
	//_|==1==|
}
