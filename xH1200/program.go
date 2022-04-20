package main

import (
	"fmt"
	"github.com/tidwall/gjson")

func main () {
	fmt.Println ("PHASE 1:")
	//_|^^1^^|
	//_|^^2^^|
	fmt.Println ("-------- Component 01: Starting...")
	
	var chan10_Clap chan []string
	chan10_Clap = make (chan []string)
	var chan10_Flap chan []string
	chan10_Flap = make (chan []string)
	
	go vX_10ErrorReporter (chan10_Clap, chan10_Flap)

	_x1100 := <- chan10_Flap

	if gjson.Get (_x1100 [0], "id").String () == "0" {
		fmt.Println (fmt.Sprintf ("-------- Component 01 (Error Reporter) could not " +
			"start up. [%s]", gjson.Get (_x1100 [0], "description").String ()))
		return
	}
	//_|==2==|
	fmt.Println ("-------- Component 01: Started!")

	//_|^^2^^|
	fmt.Println ("-------- Component 02: Starting...")

	var chan20_Clap chan []string
	chan20_Clap = make (chan []string)
	var chan20_Flap chan []string
	chan20_Flap = make (chan []string)
	
	go vX_20TaskExecutor (chan20_Clap, chan20_Flap)

	_x1200 := <- chan20_Flap

	if gjson.Get (_x1200 [0], "id").String () == "0" {
		fmt.Println (fmt.Sprintf ("-------- Component 02 (Task Executor) could not " +
			"start up. [%s]", gjson.Get (_x1200 [0], "description").String ()))
		return
	}
	//_|==2==|
	fmt.Println ("-------- Component 02: Started!")

	//_|^^2^^|
	fmt.Println ("-------- Component 03: Starting...")
	
	var chan30_Clap chan []string
	chan30_Clap = make (chan []string)
	var chan30_Flap chan []string
	chan30_Flap = make (chan []string)
	
	go vX_30HttpIM (chan30_Clap, chan30_Flap)

	_x1300 := <- chan30_Flap

	if gjson.Get (_x1300 [0], "id").String () == "0" {
		fmt.Println (fmt.Sprintf ("-------- Component 03 (HTTP I M) could not " +
			"start up. [%s]", gjson.Get (_x1300 [0], "description").String ()))
		return
	}
	fmt.Println ("-------- Component 03: Started!")
	//_|==2==|

	fmt.Println ("-------- Components started successfully.\n")

	//_|==1==|




	fmt.Println ("PHASE 2:")
	//_|^^1^^|
	for {
		var _x2100 []string

		select {
			case _x2100 = <- chan10_Flap: {
				_x2100 [0] = "10"
			}
			case _x2100 = <- chan20_Flap: {
				_x2100 [0] = "20"
			}
			case _x2100 = <- chan30_Flap: {
				_x2100 [0] = "30"
			}
		}

		if _x2100 [1] == "10" {
			chan10_Clap <- _x2100
		} else if _x2100 [1] == "20" {
			chan20_Clap <- _x2100
		} else if _x2100 [1] == "30" {
			chan30_Clap <- _x2100
		}	
	}
	//_|==1==|
}
