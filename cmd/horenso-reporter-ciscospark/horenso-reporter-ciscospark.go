package main

import (
	"crypto/tls"
	"os"
	"net/http"

	"github.com/sukiyaki/horenso-reporter-ciscospark/helper"
	"github.com/sukiyaki/horenso-reporter-ciscospark/reporter"
	"github.com/jbogarin/go-cisco-spark/ciscospark"
)

func main() {
	token, roomId, toPersonEmail, items, notifyEverything := helper.Getenvs()

	report := helper.GetReport(os.Stdin)

	if *report.ExitCode == 0 && !notifyEverything {
		os.Exit(0)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}
	api := ciscospark.NewClient(client)
	api.Authorization = "Bearer " + token

	reporter.SendReportToCiscoSpark(api, report, roomId, toPersonEmail, items)
}
