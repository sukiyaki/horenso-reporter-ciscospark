package helper

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func resetEnvs() {
	os.Setenv("HRS_SPARK_TOKEN", "")
	os.Setenv("HRS_SPARK_ROOMID", "")
	os.Setenv("HRS_SPARK_TO_EMAIL", "")
	os.Setenv("HRS_SPARK_NOTIFY_EVERYTHING", "")
}

func TestGetenvs(t *testing.T) {
	func() {
		defer func() {
			err := recover()
			if err != nil {
				assert.Equal(t, "HRS_SPARK_TOKEN environment variable is required.", err)
			} else {
				t.Fail()
			}
		}()

		resetEnvs()
		token, _, _, _, _ := Getenvs()
		if token == "" {
			t.Fail()
		}
	}()

	func() {
		defer func() {
			err := recover()
			if err != nil {
				assert.Equal(t, "HRS_SPARK_ROOMID or HRS_SPARK_TO_EMAIL environment variable is required.", err)
			} else {
				t.Fail()
			}
		}()

		resetEnvs()
		os.Setenv("HRS_SPARK_TOKEN", "token")
		token, _, _, _, _ := Getenvs()
		if token == "" {
			t.Fail()
		}
	}()

	func() {
		defer func() {
			err := recover()
			if err != nil {
				assert.Equal(t, "Either HRS_SPARK_ROOMID or HRS_SPARK_TO_EMAIL can be provided, but not both.", err)
			} else {
				t.Fail()
			}
		}()

		resetEnvs()
		os.Setenv("HRS_SPARK_TOKEN", "token")
		os.Setenv("HRS_SPARK_ROOMID", "roomid")
		os.Setenv("HRS_SPARK_TO_EMAIL", "example@example.com")
		os.Setenv("HRS_SPARK_NOTIFY_EVERYTHING", "0")

		token, roomId, toPersonEmail, items, notifyEverything := Getenvs()

		assert.Equal(t, "token", token)
		assert.Equal(t, "roomid", roomId)
		assert.Equal(t, "example@example.com", toPersonEmail)
		assert.Equal(t, []string{"all"}, items)
		assert.Equal(t, false, notifyEverything)
	}()

	func() {
		resetEnvs()
		os.Setenv("HRS_SPARK_TOKEN", "token")
		os.Setenv("HRS_SPARK_ROOMID", "roomid")
		os.Setenv("HRS_SPARK_NOTIFY_EVERYTHING", "0")

		token, roomId, toPersonEmail, items, notifyEverything := Getenvs()

		assert.Equal(t, "token", token)
		assert.Equal(t, "roomid", roomId)
		assert.Equal(t, "", toPersonEmail)
		assert.Equal(t, []string{"all"}, items)
		assert.Equal(t, false, notifyEverything)
	}()
}

func TestGetReport(t *testing.T) {
	func() {
		f, _ := os.Open("../fixtures/report_exit_0.json")
		r := GetReport(f)
		assert.Equal(t, 0, *r.ExitCode)
		assert.Equal(t, "command exited with code: 0", r.Result)
	}()

	func() {
		f, _ := os.Open("../fixtures/report_exit_1.json")
		r := GetReport(f)
		assert.Equal(t, 1, *r.ExitCode)
		assert.Equal(t, "command exited with code: 1", r.Result)
	}()

	func() {
		f, _ := os.Open("../fixtures/report_not_found.json")
		r := GetReport(f)
		assert.Equal(t, -1, *r.ExitCode)
		assert.Equal(t, "failed to execute command: exec: \"foobarbaz\": executable file not found in $PATH", r.Result)
	}()
}
