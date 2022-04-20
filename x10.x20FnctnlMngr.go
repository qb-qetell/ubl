//-- p --
package main

//-- r --
import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/octamile/pckrAuth"
	"github.com/thanhpk/randstr"
	"github.com/tidwall/gjson"
	"os"
	"regexp"
	"strings"
	"time"
)
var (
	dbDtls1 *sql.DB
	userDtls string
)

//-- i --
func x10_x20FnctnlMngr (clap <-chan []string, flap chan<- []string) {
	_x5100 := <- clap
	
	_f1100, _f1200 := os.ReadFile (".plate/cnfg.json")
	if _f1200 != nil {
		_f1251 := fmt.Sprintf ("Could not fetch configuration data. [%s]",
			_f1200.Error ())
		_f1252 := make (map[string]string)
		_f1252 ["startupSccsssStatusId"] = "0"
		_f1252 ["startupSccsssStatusDscrpt"] = _f1251
		_f1253, _ := json.Marshal (_f1252)
		
		flap <- []string {"", "", string (_f1253)}
		return
	}
	userDtls = string (_f1100)
	
	_f1400, _f1500 := sql.Open ("postgres",
		fmt.Sprintf ("postgres://%s:%s@%s:%s/%s?sslmode=require&sslrootcert=%s&" +
			"connect_timeout=8",
			gjson.Get (userDtls, "figureMngr.accessInfo.userName").String (),
			gjson.Get (userDtls, "figureMngr.accessInfo.passWord").String (),
			gjson.Get (userDtls, "figureMngr.ntwrkId.part1").String (),
			gjson.Get (userDtls, "figureMngr.ntwrkId.part2").String (),
			gjson.Get (userDtls, "figureMngr.db1").String (),
			gjson.Get (userDtls, "figureMngr.ntwrkId.part2").String (),
		),
	)
	if _f1500 != nil {
		_f1551 := fmt.Sprintf ("Could not contact figure manager. [%s]",
			_f1500.Error ())
		_f1552 := make (map[string]string)
		_f1552 ["startupSccsssStatusId"] = "0"
		_f1552 ["startupSccsssStatusDscrpt"] = _f1551
		_f1553, _ := json.Marshal (_f1552)
		
		flap <- []string {"", "", string (_f1553)}
		return
	}
	_f2500 := _f1400.Ping ()
	if _f2500 != nil {
		_f2551 := fmt.Sprintf ("Could not contact figure manager. [%s]",
			_f2500.Error ())
		_f2552 := make (map[string]string)
		_f2552 ["startupSccsssStatusId"] = "0"
		_f2552 ["startupSccsssStatusDscrpt"] = _f2551
		_f2553, _ := json.Marshal (_f2552)
		
		flap <- []string {"", "", string (_f2553)}
		return
	}

	dbDtls1 = _f1400
	
	_f5200 := make (map[string]string)
	_f5200 ["startupSccsssStatusId"] = "1"
	_f5300, _ := json.Marshal (_f5200)
	flap <- []string {"", "", string (_f5300)}
	
	for {
		_f5400 := <- clap