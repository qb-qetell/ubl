//-- p --
package main

//-- r --
import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

//-- i --
func vX_30HttpIM (clap <-chan []string, flap chan<- []string) {
	//_|^^1^^|
	//_|^^2^^|
	_x1100, _x1200 := os.ReadFile (".plate/X_30HttpIM/config.json")
	if _x1200 != nil {
		_x1300 := fmt.Sprintf ("{\"id\": \"0\", \"description\": \"Could not load " +
			"configuration info from plate. [%s]\"}", _x1200.Error ())
		flap <- []string {_x1300}
		return
	}
	//_|==2==|

	//_|^^2^^|
	_y1100 := gjson.Get (string (_x1100), "maxQuantityOfHeldRequests").String ()
	_y1200, _ := strconv.Atoi (_y1100)
	if _y1200 < 1 {
		_y1300 := fmt.Sprintf ("{\"id\": \"0\", \"description\": \"Vale of " +
			"'MaxQuantityOfHeldRequests' is less than 1; value not allowed to be " +
			"less than 1 or invalid.\"}")
		flap <- []string {_y1300}
		return
	}
	//_|==2==|

	//_|^^2^^|
	_y1400 := gjson.Get (string (_x1100), "maxRequestProcessingTime").String ()
	_y1500, _ := strconv.Atoi (_y1400)
	if _y1500 < 1 {
		_y1600 := fmt.Sprintf ("{\"id\": \"0\", \"description\": \"Vale of " +
			"'MaxRequestProcessingTime' is less than 1; value not allowed to be " +
			"less than 1 or invalid.\"}")
		flap <- []string {_y1600}
		return
	}
	//_|==2==|

	//_|^^2^^|
	flap <- []string {"{\"id\": \"1\"}"}
	//_|==2==|
	//_|==1==|

	//_|^^1^^|
	vX_30HttpIM_MaxQuantityOfHeldRequests = _y1200
	vX_30HttpIM_PresntQuantityOfHeldRequests = 0
	vX_30HttpIM_PresntQuantityOfHeldRequestsChangingAuthorization = &sync.Mutex {}
	vX_30HttpIM_Clap = clap
	vX_30HttpIM_RequestReceivingAuthorization = &sync.Mutex {}
	vX_30HttpIM_Flap = flap
	vX_30HttpIM_RequestSendingAuthorization = &sync.Mutex {}
	vX_30HttpIM_Crate = make (map[string]string)
	vX_30HttpIM_MaxRequestProcessingTime = int64 (_y1500)
	//_|==1==|

	//_|^^1^^|
	_x1400 := http.ListenAndServeTLS (
		gjson.Get (string (_x1100), "network.id.part1").String () + ":" +
			gjson.Get (string (_x1100), "network.id.part2").String (),
		gjson.Get (string (_x1100), "network.tls.crtFileId").String (),
		gjson.Get (string (_x1100), "network.tls.keyFileId").String (),
		&vX_30HttpIM_Task {},
	)
	if _x1400 != nil {
		_x1500 := fmt.Sprintf ("HTTP interface shutdown due to an error. [%s]",
			_x1400.Error ())
		fmt.Println (_x1500)
	}
	//_|==1==|
}
type vX_30HttpIM_Task struct {}
func (v *vX_30HttpIM_Task) ServeHTTP (w http.ResponseWriter,
	req *http.Request) {
	//_|^^1^^|
	if vX_30HttpIM_PresntQuantityOfHeldRequests == vX_30HttpIM_MaxQuantityOfHeldRequests {
		io.WriteString (w, "{\"success_status\": \"2\"}")
		return
	}
	//_|==1==|

	//_|^^1^^|
	_x1100, _x1200 := io.ReadAll (req.Body)
	_x1150 := string (_x1100)
	if _x1200 != nil {
		//_|^^2^^|
		_x1300 := fmt.Sprintf ("Reporter: Component 30 (HTTP I M)\n\n" +
			"Time: %s\n\n" +
			"Request Id: %s\n\n" +
			"Request Description: %s\n\n" +
			"Error: Could not get info from the info manager to determine " +
			"whether user exist. [%s]",
			time.Now ().Format ("Jan 02, 2006; 03:04 PM (-07:00)"),
			"-",
			_x1150,
			_x1200.Error (),
		)
		vX_30HttpIM_Flap <- []string {"30", "10", _x1300}
		//_|==2==|

		//_|^^2^^|
		io.WriteString (w, "{\"success_status\": \"3\"}")
		//_|==2==|

		return
	}
	//_|^^1^^|

	//_|^^1^^|
	vX_30HttpIM_PresntQuantityOfHeldRequestsChangingAuthorization.Lock ()
	request1400 := make (map[string]interface {})
	_x1450 := vX_30HttpIM_10GenerateRequestId ()
	request1400 ["id"] = _x1450
	_x1470, _ := gjson.Parse (_x1150).Value ().(map[string]interface{})
	request1400 ["description"] = _x1470
	_x1500, _ := json.Marshal (request1400)
	vX_30HttpIM_PresntQuantityOfHeldRequests = vX_30HttpIM_PresntQuantityOfHeldRequests + 1
	vX_30HttpIM_PresntQuantityOfHeldRequestsChangingAuthorization.Unlock ()
	//_|==1==|

	//_|^^1^^|
	vX_30HttpIM_RequestSendingAuthorization.Lock ()
	vX_30HttpIM_Flap <- []string {"30", "20", string (_x1500)}
	requestExpirationTime := time.Now ().Unix () + vX_30HttpIM_MaxRequestProcessingTime
	vX_30HttpIM_RequestSendingAuthorization.Unlock ()
	//_|==1==|

	//_|^^1^^|
	for true {
		vX_30HttpIM_RequestReceivingAuthorization.Lock ()

		done1550 := false
		if true {
			_x1600, _x1700 := vX_30HttpIM_Crate [_x1450]
			if _x1700 == true {
				delete (vX_30HttpIM_Crate, _x1450)
				_x1800 := gjson.Get (_x1600, "description").String ()
				io.WriteString (w, _x1800)
				done1550 = true
				goto done1575
			}

			select {
				case _x2100 := <- vX_30HttpIM_Clap: {
					_x2200 := gjson.Get (_x2100 [2], "id").String ()
					_x2300 := gjson.Get (_x2100 [2],
						"description").String ()
					if _x2200 == _x1450 {
						fmt.Println ("gotcha", _x2300)			
						io.WriteString (w, _x2300)
						done1550 = true
						goto done1575
					} else {
						_x2400, _ := vX_30HttpIM_Crate [_x2200]
						if _x2400 == "ignore" {
							delete (vX_30HttpIM_Crate, _x2200)
						} else {
							vX_30HttpIM_Crate [_x2200] = _x2300
						}
					}
				}
				default: {
					_x3100 := time.Now ().Unix ()
					
					if _x3100 > requestExpirationTime {
						vX_30HttpIM_Crate [_x1450] = "ignore"
						io.WriteString (w,
							"{\"success_status\": \"3\"}")
						done1550 = true
						goto done1575
					}
				}
			}		
		}

		done1575:
		vX_30HttpIM_RequestReceivingAuthorization.Unlock ()

		if done1550 == true {
			vX_30HttpIM_PresntQuantityOfHeldRequestsChangingAuthorization.Lock ()
			vX_30HttpIM_PresntQuantityOfHeldRequests =
				vX_30HttpIM_PresntQuantityOfHeldRequests - 1
			vX_30HttpIM_PresntQuantityOfHeldRequestsChangingAuthorization.Unlock ()
			return
		}
	}
}
func vX_30HttpIM_10GenerateRequestId () string {
	_x1400 := time.Now ()
	_x1500 := _x1400.Nanosecond ()

	_x1600 := strconv.Itoa (_x1500)
	for len (_x1600) < 9 {
		_x1600 = "0" + _x1600
	}

	_x1700 := fmt.Sprintf ("%s%s",
		_x1400.Format ("2006January02150405"),
		_x1600)

	_x1700 = strings.Replace (_x1700, "January",   "01", -1)
	_x1700 = strings.Replace (_x1700, "February",  "02", -1)
	_x1700 = strings.Replace (_x1700, "March",     "03", -1)
	_x1700 = strings.Replace (_x1700, "April",     "04", -1)
	_x1700 = strings.Replace (_x1700, "May",       "05", -1)
	_x1700 = strings.Replace (_x1700, "June",      "06", -1)
	_x1700 = strings.Replace (_x1700, "July",      "07", -1)
	_x1700 = strings.Replace (_x1700, "August",    "08", -1)
	_x1700 = strings.Replace (_x1700, "September", "09", -1)
	_x1700 = strings.Replace (_x1700, "October",   "10", -1)
	_x1700 = strings.Replace (_x1700, "November",  "11", -1)
	_x1700 = strings.Replace (_x1700, "December",  "12", -1)

	return _x1700
}
func vX_30HttpIM_20TellTime () string {
	_x1400 := time.Now ()
	
	_x1700 := _x1400.Format ("2006January02150405")
	_x1700 = strings.Replace (_x1700, "January",   "01", -1)
	_x1700 = strings.Replace (_x1700, "February",  "02", -1)
	_x1700 = strings.Replace (_x1700, "March",     "03", -1)
	_x1700 = strings.Replace (_x1700, "April",     "04", -1)
	_x1700 = strings.Replace (_x1700, "May",       "05", -1)
	_x1700 = strings.Replace (_x1700, "June",      "06", -1)
	_x1700 = strings.Replace (_x1700, "July",      "07", -1)
	_x1700 = strings.Replace (_x1700, "August",    "08", -1)
	_x1700 = strings.Replace (_x1700, "September", "09", -1)
	_x1700 = strings.Replace (_x1700, "October",   "10", -1)
	_x1700 = strings.Replace (_x1700, "November",  "11", -1)
	_x1700 = strings.Replace (_x1700, "December",  "12", -1)

	return _x1700
}


var (
	vX_30HttpIM_MaxQuantityOfHeldRequests int
	vX_30HttpIM_PresntQuantityOfHeldRequests int
	vX_30HttpIM_PresntQuantityOfHeldRequestsChangingAuthorization *sync.Mutex
	vX_30HttpIM_Clap <-chan []string
	vX_30HttpIM_RequestReceivingAuthorization *sync.Mutex
	vX_30HttpIM_Flap chan<- []string
	vX_30HttpIM_RequestSendingAuthorization *sync.Mutex
	vX_30HttpIM_Crate map[string]string
	vX_30HttpIM_MaxRequestProcessingTime int64	
)
