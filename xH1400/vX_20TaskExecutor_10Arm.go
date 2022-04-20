package main

import (
	"crypto/sha256"
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
	_x1400, _x1500 := sql.Open ("postgres",
		fmt.Sprintf ("postgres://%s:%s@%s:%s/%s?sslmode=require&sslrootcert=%s&" +
			"connect_timeout=8",
			gjson.Get (string (_x1100), "im.user.name").String (),
			gjson.Get (string (_x1100), "im.user.password").String (),
			gjson.Get (string (_x1100), "im.networkInfo.id.part1").String (),
			gjson.Get (string (_x1100), "im.networkInfo.id.part2").String (),
			gjson.Get (string (_x1100), "im.database1").String (),
			gjson.Get (string (_x1100),
				"im.networkInfo.connEncryptionEnablerCert").String (),
		),
	)
	if _x1500 != nil {
		_x1600 := fmt.Sprintf ("{\"id\": \"0\", \"description\": \"Could not " +
			"connect to DB. [%s]\"}", _x1500.Error ())
		flap <- []string {_x1600}
		return
	}
	_x1700 := _x1400.Ping ()
	if _x1700 != nil {
		_x1800 := fmt.Sprintf ("{\"id\": \"0\", \"description\": \"Could not ping " +
			"DB. [%s]\"}", _x1700.Error ())
		flap <- []string {_x1800}
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
		_x2000 := _x1400.Ping ()
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

		_o2000 := _f1400.Ping ()
		if _o2000 != nil {
			_o2100 := fmt.Sprintf ("Reporter: Component 20 (Task Executor)\n" +
				"Time: %s\n" +
				"Request Id: %s\n" +
				"Request Description: %s\n" +
				"Error: Could not connect the info manager. [%s]",
				time.Now ().Format ("Jan 02, 2006; 03:04 PM (-07:00)"),
				gjson.Get (_x1900 [2], "id").String (),
				gjson.Get (_x1900 [2], "description").String (),
				_o2000.Error ())
			flap <- []string {"20", "10", _o2100}
			
			_o2200 := make (map[string]interface {})
			_o2200 ["id"] = gjson.Get (_x1900 [2], "id").String ()
			_o2300 := make (map[string]interface {})
			_o2300 ["success_status"] = "3"
			_o2200 ["description"] = _o2300			
			_o2400, _ := json.Marshal (_o2200)
			flap <- []string {"20", "30", string (_o2400)}
			continue
		}
		//_|==2==|

		//_|^^2^^|
		id := ""
		//_|==2==|

		//_|^^2^^|
		//_|^^3^^|
		if id == "" {
			_x2500, _x2600 := _x1400.Query ("select id from x1100_user where " +
				"id = $1",
				gjson.Get (_x1900 [2], "description.user_iupe").String (),
			)
			if _x2600 != nil {
				_x2700 := fmt.Sprintf ("Reporter: Component 20 (Task " +
					"Executor)\n" +
					"Time: %s\n" +
					"Request Id: %s\n" +
					"Request Description: %s\n" +
					"Error: Could not get info from the info manager " +
					"to determine whether user exist. [%s]",
					time.Now ().Format ("Jan 02, 2006; 03:04 PM (-07:00)"),
					gjson.Get (_x1900 [2], "id").String (),
					gjson.Get (_x1900 [2], "description").String (),
					_x2600.Error ())
				flap <- []string {"20", "10", _x2700}
			
				_x2800 := make (map[string]interface {})
				_x2800 ["id"] = gjson.Get (_x1900 [2], "id").String ()
				_x2900 := make (map[string]interface {})
				_x2900 ["success_status"] = "3"
				_x2800 ["description"] = _x2900			
				_x3000, _ := json.Marshal (_x2800)
				flap <- []string {"20", "30", string (_x3000)}
				continue
			}

			_x3200 := _x2500.Next ()

			if _x3200 == true {
				_x3300 := _x2500.Scan (&id)
				if _x3300 != nil {
					_x3400 := fmt.Sprintf ("Reporter: Component 20 " +
						"(Task Executor)\n" +
						"Time: %s\n" +
						"Request Id: %s\n" +
						"Request Description: %s\n" +
						"Error: Could not get info from the info " +
						"manager to determine whether user exist. [%s]",
						time.Now ().Format (
							"Jan 02, 2006; 03:04 PM (-07:00)"),
						gjson.Get (_x1900 [2], "id").String (),
						gjson.Get (_x1900 [2], "description").String (),
						_x3300.Error ())
					flap <- []string {"20", "10", _x3400}
				
					_x3500 := make (map[string]interface {})
					_x3500 ["id"] = gjson.Get (_x1900 [2], "id").String ()
					_x3600 := make (map[string]interface {})
					_x3600 ["success_status"] = "3"
					_x3500 ["description"] = _x3600
					_x3700, _ := json.Marshal (_x3500)
					flap <- []string {"20", "30", string (_x3700)}
					continue
				}
			}
		}
		//_|==3==|

		//_|^^3^^|
		if id == "" {
			_y2500, _y2600 := _x1400.Query ("select id from x1100_user where " +
				"username = $1",
				gjson.Get (_x1900 [2], "description.user_iupe").String (),
			)
			if _y2600 != nil {
				_y2700 := fmt.Sprintf ("Reporter: Component 20 (Task " +
					"Executor)\n" +
					"Time: %s\n" +
					"Request Id: %s\n" +
					"Request Description: %s\n" +
					"Error: Could not get info from the info manager to " +
					"determine whether user exist. [%s]",
					time.Now ().Format ("Jan 02, 2006; 03:04 PM (-07:00)"),
					gjson.Get (_x1900 [2], "id").String (),
					gjson.Get (_x1900 [2], "description").String (),
					_y2600.Error ())
				flap <- []string {"20", "10", _y2700}
			
				_y2800 := make (map[string]interface {})
				_y2800 ["id"] = gjson.Get (_x1900 [2], "id").String ()
				_y2900 := make (map[string]interface {})
				_y2900 ["success_status"] = "3"
				_y2800 ["description"] = _y2900			
				_y3000, _ := json.Marshal (_y2800)
				flap <- []string {"20", "30", string (_y3000)}
				continue
			}

			_y3200 := _y2500.Next ()

			if _y3200 == true {
				_y3300 := _y2500.Scan (&id)
				if _y3300 != nil {
					_y3400 := fmt.Sprintf ("Reporter: Component 20 " +
						"(Task Executor)\n" +
						"Time: %s\n" +
						"Request Id: %s\n" +
						"Request Description: %s\n" +
						"Error: Could not get info from the info " +
						"manager to determine whether user exist. [%s]",
						time.Now ().Format (
							"Jan 02, 2006; 03:04 PM (-07:00)"),
						gjson.Get (_x1900 [2], "id").String (),
						gjson.Get (_x1900 [2], "description").String (),
						_y3300.Error ())
					flap <- []string {"20", "10", _y3400}
				
					_y3500 := make (map[string]interface {})
					_y3500 ["id"] = gjson.Get (_x1900 [2], "id").String ()
					_y3600 := make (map[string]interface {})
					_y3600 ["success_status"] = "3"
					_y3500 ["description"] = _y3600
					_y3700, _ := json.Marshal (_y3500)
					flap <- []string {"20", "30", string (_y3700)}
					continue
				}
			}
		}
		//_|==3==|

		//_|^^3^^|
		if id == "" {
			_z2500, _z2600 := _x1400.Query ("select id from x1100_user where " +
				"phone_no = $1",
				gjson.Get (_x1900 [2], "description.user_iupe").String (),
			)
			if _z2600 != nil {
				_z2700 := fmt.Sprintf ("Reporter: Component 20 (Task " +
					"Executor)\n" +
					"Time: %s\n" +
					"Request Id: %s\n" +
					"Request Description: %s\n" +
					"Error: Could not get info from the info manager to " +
					"determine whether user exist. [%s]",
					time.Now ().Format ("Jan 02, 2006; 03:04 PM (-07:00)"),
					gjson.Get (_x1900 [2], "id").String (),
					gjson.Get (_x1900 [2], "description").String (),
					_z2600.Error ())
				flap <- []string {"20", "10", _z2700}
			
				_z2800 := make (map[string]interface {})
				_z2800 ["id"] = gjson.Get (_x1900 [2], "id").String ()
				_z2900 := make (map[string]interface {})
				_z2900 ["success_status"] = "3"
				_z2800 ["description"] = _z2900			
				_z3000, _ := json.Marshal (_z2800)
				flap <- []string {"20", "30", string (_z3000)}
				continue
			}

			_z3200 := _z2500.Next ()

			if _z3200 == true {
				_z3300 := _z2500.Scan (&id)
				if _z3300 != nil {
					_z3400 := fmt.Sprintf ("Reporter: Component 20 " +
						"(Task Executor)\n" +
						"Time: %s\n" +
						"Request Id: %s\n" +
						"Request Description: %s\n" +
						"Error: Could not get info from the info " +
						"manager to determine whether user exist. [%s]",
						time.Now ().Format (
							"Jan 02, 2006; 03:04 PM (-07:00)"),
						gjson.Get (_x1900 [2], "id").String (),
						gjson.Get (_x1900 [2], "description").String (),
						_z3300.Error ())
					flap <- []string {"20", "10", _z3400}
				
					_z3500 := make (map[string]interface {})
					_z3500 ["id"] = gjson.Get (_x1900 [2], "id").String ()
					_z3600 := make (map[string]interface {})
					_z3600 ["success_status"] = "3"
					_z3500 ["description"] = _z3600
					_z3700, _ := json.Marshal (_z3500)
					flap <- []string {"20", "30", string (_z3700)}
					continue
				}
			}
		}
		//_|==3==|

		//_|^^3^^|
		if id == "" {
			_a2500, _a2600 := _x1400.Query ("select id from x1100_user where " +
				"e_mail = $1",
				gjson.Get (_x1900 [2], "description.user_iupe").String (),
			)
			if _a2600 != nil {
				_a2700 := fmt.Sprintf ("Reporter: Component 20 (Task " +
					"Executor)\n" +
					"Time: %s\n" +
					"Request Id: %s\n" +
					"Request Description: %s\n" +
					"Error: Could not get info from the info manager to " +
					"determine whether user exist. [%s]",
					time.Now ().Format ("Jan 02, 2006; 03:04 PM (-07:00)"),
					gjson.Get (_x1900 [2], "id").String (),
					gjson.Get (_x1900 [2], "description").String (),
					_a2600.Error ())
				flap <- []string {"20", "10", _a2700}
			
				_a2800 := make (map[string]interface {})
				_a2800 ["id"] = gjson.Get (_x1900 [2], "id").String ()
				_a2900 := make (map[string]interface {})
				_a2900 ["success_status"] = "3"
				_a2800 ["description"] = _a2900			
				_a3000, _ := json.Marshal (_a2800)
				flap <- []string {"20", "30", string (_a3000)}
				continue
			}

			_a3200 := _a2500.Next ()

			if _a3200 == true {
				_a3300 := _a2500.Scan (&id)
				if _a3300 != nil {
					_a3400 := fmt.Sprintf ("Reporter: Component 20 " +
						"(Task Executor)\n" +
						"Time: %s\n" +
						"Request Id: %s\n" +
						"Request Description: %s\n" +
						"Error: Could not get info from the info " +
						"manager to determine whether user exist. [%s]",
						time.Now ().Format (
							"Jan 02, 2006; 03:04 PM (-07:00)"),
						gjson.Get (_x1900 [2], "id").String (),
						gjson.Get (_x1900 [2], "description").String (),
						_a3300.Error ())
					flap <- []string {"20", "10", _a3400}
				
					_a3500 := make (map[string]interface {})
					_a3500 ["id"] = gjson.Get (_x1900 [2], "id").String ()
					_a3600 := make (map[string]interface {})
					_a3600 ["success_status"] = "3"
					_a3500 ["description"] = _a3600
					_a3700, _ := json.Marshal (_a3500)
					flap <- []string {"20", "30", string (_a3700)}
					continue
				}
			}
		}
		//_|==3==|

		//_|^^3^^|
		if id == "" {
			_b2200 := make (map[string]interface {})
			_b2200 ["id"] = gjson.Get (_x1900 [2], "id").String ()
			_b2300 := make (map[string]interface {})
			_b2300 ["success_status"] = "4"
			_b2200 ["description"] = _b2300			
			_b2400, _ := json.Marshal (_b2200)
			flap <- []string {"20", "30", string (_b2400)}
			continue
		}
		//_|==3==|
		//_|==2==|

		//_|^^2^^|
		_e2500, _e2600 := _f1400.Query ("select __user from " +
			"x1100____authentication_info where __user = $1 and " +
			"password = $2",
			gjson.Get (_x1900 [2], "description.user_iupe").String (),
			fmt.Sprintf ("%x", sha256.Sum256 ([]byte (gjson.Get (_x1900 [2],
				"description.password").String ()))),
		)
		if _e2600 != nil {
			_e2700 := fmt.Sprintf ("Reporter: Component 20 (Task " +
				"Executor)\n" +
				"Time: %s\n" +
				"Request Id: %s\n" +
				"Request Description: %s\n" +
				"Error: Could not get info from the info manager to " +
				"determine whether user exist. [%s]",
				time.Now ().Format ("Jan 02, 2006; 03:04 PM (-07:00)"),
				gjson.Get (_x1900 [2], "id").String (),
				gjson.Get (_x1900 [2], "description").String (),
				_e2600.Error ())
			flap <- []string {"20", "10", _e2700}
			
			_e2800 := make (map[string]interface {})
			_e2800 ["id"] = gjson.Get (_x1900 [2], "id").String ()
			_e2900 := make (map[string]interface {})
			_e2900 ["success_status"] = "3"
			_e2800 ["description"] = _e2900			
			_e3000, _ := json.Marshal (_e2800)
			flap <- []string {"20", "30", string (_e3000)}
			continue
		}
		_e3200 := _e2500.Next ()
		if _e3200 == false {
			_e3500 := make (map[string]interface {})
			_e3500 ["id"] = gjson.Get (_x1900 [2], "id").String ()
			_e3600 := make (map[string]interface {})
			_e3600 ["success_status"] = "5"
			_e3500 ["description"] = _e3600
			_e3700, _ := json.Marshal (_e3500)
			flap <- []string {"20", "30", string (_e3700)}
			continue
		}
		//_|==2==|

		//_|^^2^^|
		sessionId := vX_20TaskExecutor_10Arm_10GenerateRequestId ()
		sessionKey := fmt.Sprintf("%x",
			sha256.Sum256 ([]byte (strconv.Itoa (time.Now ().Nanosecond ()))))
		sessionExpiryTime := time.Now ().Unix () + 240
		//_|==2==|
		
		//_|^^2^^|
		_j2500, _j2600 := _f1400.Query ("select __user from " +
			"x1200____session where __user = $1 and " +
			"id = $2",
			gjson.Get (_x1900 [2], "description.user").String (), id,
		)
		if _j2600 != nil {
			_j2700 := fmt.Sprintf ("Reporter: Component 20 (Task " +
				"Executor)\n" +
				"Time: %s\n" +
				"Request Id: %s\n" +
				"Request Description: %s\n" +
				"Error: Could not get info from the info manager to " +
				"determine whether user exist. [%s]",
				time.Now ().Format ("Jan 02, 2006; 03:04 PM (-07:00)"),
				gjson.Get (_x1900 [2], "id").String (),
				gjson.Get (_x1900 [2], "description").String (),
				_j2600.Error ())
			flap <- []string {"20", "10", _j2700}
			
			_j2800 := make (map[string]interface {})
			_j2800 ["id"] = gjson.Get (_x1900 [2], "id").String ()
			_j2900 := make (map[string]interface {})
			_j2900 ["success_status"] = "3"
			_j2800 ["description"] = _j2900			
			_j3000, _ := json.Marshal (_j2800)
			flap <- []string {"20", "30", string (_j3000)}
			continue
		}
		_j3200 := _j2500.Next ()
		if _j3200 == true {
			_j3500 := make (map[string]interface {})
			_j3500 ["id"] = gjson.Get (_x1900 [2], "id").String ()
			_j3600 := make (map[string]interface {})
			_j3600 ["success_status"] = "3"
			_j3500 ["description"] = _j3600
			_j3700, _ := json.Marshal (_j3500)
			flap <- []string {"20", "30", string (_j3700)}
			continue
		}
		//_|==2==|

		//_|^^2^^|
		_, _h2600 := _f1400.Exec ("insert into x1200____session " +
			"(__user, id, key, expiry_time) values ($1, $2, $3, $4)",
			id, sessionId, sessionKey, sessionExpiryTime,
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
		_i2200 := make (map[string]interface {})
		_i2200 ["id"] = gjson.Get (_x1900 [2], "id").String ()
		_i2300 := make (map[string]interface {})
		_i2300 ["success_status"] = "6"
		_i2300 ["user_id"] = id
		_i2350 := make (map[string]interface {})
		_i2350 ["id"] = sessionId
		_i2350 ["key"] = sessionKey
		_i2350 ["expiryTime"] = sessionExpiryTime
		_i2300 ["session"] = _i2350
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
