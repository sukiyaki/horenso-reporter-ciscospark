package helper

import (
	"os"
	"strings"
	"time"

	"github.com/Songmu/horenso"
	"github.com/antonholmquist/jason"
)

// Getenvs get environment varibles
func Getenvs() (string, string, string, []string, bool) {
	token := os.Getenv("HRS_SPARK_TOKEN")
	roomId := os.Getenv("HRS_SPARK_ROOMID")
	toPersonEmail := os.Getenv("HRS_SPARK_TO_EMAIL")

	if len(token) == 0 {
		panic("HRS_SPARK_TOKEN environment variable is required.")
	} else if len(roomId) == 0 && len(toPersonEmail) == 0 {
		panic("HRS_SPARK_ROOMID or HRS_SPARK_TO_EMAIL environment variable is required.")
	} else if len(roomId) > 0 && len(toPersonEmail) > 0 {
		panic("Either HRS_SPARK_ROOMID or HRS_SPARK_TO_EMAIL can be provided, but not both.")
	}

	itemsStr := os.Getenv("HRS_SPARK_ITEMS")
	var items []string
	if len(itemsStr) > 0 {
		items = strings.Split(itemsStr, ",")
	} else {
		items = []string{"all"}
	}

	notifyEverythingEnv := os.Getenv("HRS_SPARK_NOTIFY_EVERYTHING")
	notifyEverything := true
	if len(notifyEverythingEnv) != 0 && notifyEverythingEnv == "0" {
		notifyEverything = false
	}

	return token, roomId, toPersonEmail, items, notifyEverything
}

// GetReport get horenso report via STDIN
func GetReport(f *os.File) horenso.Report {
	v, err := jason.NewObjectFromReader(f)
	if err != nil {
		panic(err)
	}

	var r horenso.Report
	r.Command, _ = v.GetString("command")
	r.CommandArgs, _ = v.GetStringArray("commandArgs")
	r.Tag, _ = v.GetString("tag")
	r.Output, _ = v.GetString("output")
	r.Stdout, _ = v.GetString("stdout")
	r.Stderr, _ = v.GetString("stderr")
	r.Command, _ = v.GetString("command")
	r.Result, _ = v.GetString("result")
	r.Hostname, _ = v.GetString("hostname")
	exitCode, _ := getInt(v, "exitCode")
	r.ExitCode = &exitCode
	pid, _ := getInt(v, "pid")
	r.Pid = &pid
	r.StartAt, _ = getTime(v, "startAt")
	r.EndAt, _ = getTime(v, "endAt")
	systemTime, _ := v.GetFloat64("systemTime")
	r.SystemTime = &systemTime
	userTime, _ := v.GetFloat64("userTime")
	r.UserTime = &userTime

	return r
}

func getInt(v *jason.Object, key string) (int, error) {
	n, err := v.GetInt64(key)
	if err != nil {
		return 0, err
	}

	return int(n), nil
}

func getTime(v *jason.Object, key string) (*time.Time, error) {
	s, err := v.GetString(key)
	if err != nil {
		return nil, err
	}
	t, err := time.Parse("2006-01-02T15:04:05.999999-07:00", s)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
