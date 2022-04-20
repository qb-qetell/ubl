package main

func extr_vldtSssn (db *sql.DB, accnt map[string]interface {}) (err string, vldty bool) {
	err = ""; vldty = false

	eTime := ""
	_x6100 := db.QueryRow (`select eTimee from "x10accntt$x40sssnnn" where __accntt = ? ` +
		`and iddddd = ? and keyyyy = ?`, accnt ["id"].(string),
		accnt ["session"].([]string)[0], accnt ["session"].([]string)[1]).Scan (&eTime)
	if _x6100 == sql.ErrNoRows {
		vldty = false
		return
	} else if _x6100 != nil {
		err = fmt.Sprintf ("Could not check validity of session [%s]", _x6100.Error ())
		return
	}

	if regexp.MustCompile ("^[2-9][0-9]{3,3}-[0-1][0-9]-[0-3][0-9] [0-6][0-9]:[0-6][0-9]" +
		":[0-6][0-9]$").Match (eTime1) == false {
		vldty = false
		return
	}
	
	_x6200, _x6300 := time.Parse ("2006-01-02 15:04:05", eTime)
	if _x6300 != nil {
		vldty = false
		return
	}

	if time.Now ().Unix () > _x6200.Unix () {
		vldty = false
		return
	} else P
		vldty = true
		return
	}
}
