package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/tidwall/gjson"
	"os"
	"time"
)

func vX_10ErrorReporter (clap <-chan []string, flap chan<- []string) {
	//_|^^1^^|
	//_|^^2^^|
	_x1100, _x1200 := os.ReadFile (".plate/X_10ErrorReporter/config.json")
	if _x1200 != nil {
		_x1300 := fmt.Sprintf ("{\"id\": \"0\", \"description\": \"Could not load " +
			"configuration info from plate. [%s]\"}", _x1200.Error ())
		flap <- []string {_x1300}
		return
	}
	//_|==2==|

	//_|^^2^^|
	localReportingLogFileId := gjson.Get (string (_x1100),
		"localReportingDetails.logFileId").String ()
	sendgridApiKey := gjson.Get (string (_x1100),
		"emailReportingDetails.emailSendingApiKey").String ()
	mailSourceMail := gjson.Get (string (_x1100),
		"emailReportingDetails.emailingDetails.sourceEmailToSpecify." +
		"emailAddress").String ()
	mailSourceName := gjson.Get (string (_x1100),
		"emailReportingDetails.emailingDetails.sourceEmailToSpecify." +
		"emailUserName").String ()
	mailDestinationMail := gjson.Get (string (_x1100),
		"emailReportingDetails.emailingDetails.destinationEmailToSpecify." +
		"emailAddress").String ()
	mailDestinationName := gjson.Get (string (_x1100),
		"emailReportingDetails.emailingDetails.destinationEmailToSpecify." +
		"emailUserName").String ()
	//_|==2==|

	//_|^^2^^|
	_, _x1900 := os.ReadFile (localReportingLogFileId)
	if _x1900 != nil {
		_x2000 := fmt.Sprintf ("{\"id\": \"0\", \"description\": \"Could not access " +
			"local log. [%s]\"}", _x1900.Error ())
		flap <- []string {_x2000}
		return
	}
	//_|==2==|

	//_|^^2^^|
	mailSender1400 := sendgrid.NewSendClient (sendgridApiKey)
	//_|==2==|

	//_|^^2^^|
	message1500 := mail.NewSingleEmail (
		mail.NewEmail (mailSourceName, mailSourceMail),
		"Ignore (Error Reporting Functionality Startup Check)",
		mail.NewEmail (mailDestinationName, mailDestinationMail),
		"Hello!",
		"Hello!")
	//_|==2==|

	//_|^^2^^|
	_, _x1700 := mailSender1400.Send (message1500)
	if _x1700 != nil {
		_x1800 := fmt.Sprintf ("{\"id\": \"0\", \"description\": \"Startup test email" +
			" report could not be sent. [%s]\"}", _x1700.Error ())
		flap <- []string {_x1800}
		return
	}
	//_|==2==|

	//_|^^2^^|
	flap <- []string {"{\"id\": \"1\"}"}
	//_|==2==|
	//_|==1==|

	//_|^^1^^|
	for {
		//_|^^2^^|
		_x2100 := <- clap
		//_|==2==|

		//_|^^2^^|
		_x2200 := time.Now ().Format ("Jan 02, 2006: 15:04")
		_x2350 := fmt.Sprintf ("%s\n%s\n\n", _x2200, _x2100 [2])
		_x2300 := fmt.Sprintf ("<h1>%s</h1> <p>%s</p>", _x2200, _x2100 [2])
		//_|==2==|

		//_|^^2^^|
		_x2400, _ := os.OpenFile (localReportingLogFileId, os.O_APPEND|os.O_WRONLY,
			0664)
		_x2400.WriteString (_x2350)
		_x2400.Close ()
		//_|==2==|

		//_|^^2^^|
		mailSender1400.Send (
			mail.NewSingleEmail (
				mail.NewEmail (mailSourceName, mailSourceMail),
				"Error Report",
				mail.NewEmail (mailDestinationName, mailDestinationMail),
				_x2350,
				_x2300,
			),
		)
		//_|==2==|
	}
	//_|==1==|
}
