package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/tidwall/gjson"
	"os"
	"strconv"
	"strings"
	"time"
)

func vX_20TaskExecutor_10Arm (clap <-chan []string, flap chan<- []string) {
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
	_f1400, _f1500 := sql.Open ("postgres",
		fmt.Sprintf ("postgres://%s:%s@%s:%s/%s?sslmode=require&sslrootcert=%s&" +
			"connect_timeout=8",
			gjson.Get (string (_x1100), "im.user.name").String (),
			gjson.Get (string (_x1100), "im.user.password").String (),
			gjson.Get (string (_x1100), "im.networkInfo.id.part1").String (),
			gjson.Get (string (_x1100), "im.networkInfo.id.part2").String (),
			gjson.Get (string (_x1100), "im.database2").String (),
			gjson.Get (string (_x1100),
				"im.networkInfo.connEncryptionEnablerCert").String (),
		),
	)
	if _f1500 != nil {
		_f1600 := fmt.Sprintf ("{\"id\": \"0\", \"description\": \"Could not " +
			"connect to DB. [%s]\"}", _f1500.Error ())
		flap <- []string {_f1600}
		return
	}
	_f1700 := _f1400.Ping ()
	if _f1700 != nil {
		_f1800 := fmt.Sprintf ("{\"id\": \"0\", \"description\": \"Could not ping " +
			"DB. [%s]\"}", _f1700.Error ())
		flap <- []string {_f1800}
		return
	}
	//_|==2==|

	//_|^^2^^|
	fmt.Println ("-------- Component 02: Arm started successfully.")

	flap <- []string {"{\"id\": \"1\"}"}
	//_|==2==|
	//_|==1==|

	//_|^^1^^|
	for {
		//_|^^2^^|
		_ = <- clap
		_x1900 := <- clap
		//_|==2==|

		//_|^^2^^|
		_x2000 := _f1400.Ping ()
		if _x2000 != nil {
			_x2100 := fmt.Sprintf ("Reporter: Component 20 (Task Executor)\n" +
				"Time: %s\n" +
				"Request Id: %s\n" +
				"Request Description: %s\n" +
				"Error: Could not connect the info manager. [%s]",
				time.Now ().Format ("Jan 02, 2006; 03:04 PM (-07:00)"),
				gjson.Get (_x1900 [2], "id").String (),
				gjson.Get (_x1900 [2], "description").String (),
				_x2000.Error ())
			flap <- []string {"20", "10", _x2100}
			
			_x2200 := make (map[string]interface {})
			_x2200 ["id"] = gjson.Get (_x1900 [2], "id").String ()
			_x2300 := make (map[string]interface {})
			_x2300 ["success_status"] = "3"
			_x2200 ["description"] = _x2300			
			_x2400, _ := json.Marshal (_x2200)
			flap <- []string {"20", "30", string (_x2400)}
			continue
		}
		//_|==2==|

		//_|^^2^^|
		_, _h2600 := _f1400.Exec ("update x1200____session set " +
			"expiry_time = $1 where __user = $2 and id = $3 and key = $4", "0",
			gjson.Get (_x1900 [2], "description.user_id").String (),
			gjson.Get (_x1900 [2], "description.session.id").String (),
			gjson.Get (_x1900 [2], "description.session.key").String (),
		)
		if _h2600 != nil {
			_h2700 := fmt.Sprintf ("Reporter: Component 20 (Task " +
				"Executor)\n" +
				"Time: %s\n" +
				"Request Id: %s\n" +
				"Request Description: %s\n" +
				"Error: Could not get info from the info manager to " +
				"determine whether user exist. [%s]",
				time.Now ().Format ("Jan 02, 2006; 03:04 PM (-07:00)"),
				gjson.Get (_x1900 [2], "id").String (),
				gjson.Get (_x1900 [2], "description").String (),
				_h2600.Error ())
			flap <- []string {"20", "10", _h2700}
			
			_h2800 := make (map[string]interface {})
			_h2800 ["id"] = gjson.Get (_x1900 [2], "id").String ()
			_h2900 := make (map[string]interface {})
			_h2900 ["success_status"] = "3"
			_h2800 ["description"] = _h2900			
			_h3000, _ := json.Marshal (_h2800)
			flap <- []string {"20", "30", string (_h3000)}
			continue
		}
		//_|==2==|

		//_|^^2^^|
		_i2200 := make (map[string]interface {})
		_i2200 ["id"] = gjson.Get (_x1900 [2], "id").String ()
		_i2300 := make (map[string]interface {})
		_i2300 ["success_status"] = "4"
		_i2200 ["description"] = _i2300			
		_i2400, _ := json.Marshal (_i2200)
		flap <- []string {"20", "30", string (_i2400)}
		//_|==2==|
	}
}
func vX_20TaskExecutor_10Arm_10GenerateRequestId () string {
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
