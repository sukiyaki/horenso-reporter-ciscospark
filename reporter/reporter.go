package reporter

import (
	"bytes"
	"fmt"

	"github.com/Songmu/horenso"
	"github.com/jbogarin/go-cisco-spark/ciscospark"
)

// GetMarkDown get a markdown message
func GetMarkDown(r horenso.Report, items []string) string {
	var buffer bytes.Buffer

	buffer.WriteString("## horenso Reporter\n")

	if IsSelectedItem("Result", items) {
		buffer.WriteString("**Result**  \n")
		buffer.WriteString("`")
		buffer.WriteString(fmt.Sprintf("%v", r.Result) + "\n")
		buffer.WriteString("`\n\n")
	}

	if IsSelectedItem("Output", items) {
		buffer.WriteString("**Output**  \n")
		buffer.WriteString("`")
                buffer.WriteString(fmt.Sprintf("%v", r.Output) + "\n")
		buffer.WriteString("`\n\n")
	}

	if IsSelectedItem("Stdout", items) {
		buffer.WriteString("**Stdout**  \n")
		buffer.WriteString("`")
		buffer.WriteString(fmt.Sprintf("%v", r.Stdout) + "\n")
		buffer.WriteString("`\n\n")
	}

	if IsSelectedItem("Stderr", items) {
                buffer.WriteString("**Stderr**  \n")
		buffer.WriteString("`")
		buffer.WriteString(fmt.Sprintf("%v", r.Stderr) + "\n")
		buffer.WriteString("`\n\n")
	}

	if IsSelectedItem("Command", items) {
                buffer.WriteString("**Command**  \n")
		buffer.WriteString("`")
		buffer.WriteString(fmt.Sprintf("%v", r.Command) + "\n")
		buffer.WriteString("`\n\n")
	}

	if IsSelectedItem("CommandArgs", items) {
                buffer.WriteString("**CommandArgs**  \n")
		buffer.WriteString("`")
		buffer.WriteString(fmt.Sprintf("%v", r.CommandArgs) + "\n")
		buffer.WriteString("`\n\n")
	}

	if IsSelectedItem("Pid", items) {
                buffer.WriteString("**Pid**  \n" + fmt.Sprintf("%d", *r.Pid) + "\n\n")
	}

	if IsSelectedItem("ExitCode", items) {
                buffer.WriteString("**ExitCode**  \n" + fmt.Sprintf("%d", *r.ExitCode) + "\n\n")
	}

	if IsSelectedItem("StartAt", items) {
                buffer.WriteString("**StartAt**  \n" + fmt.Sprintf("%v", r.StartAt) + "\n\n")
	}
	if IsSelectedItem("EndAt", items) {
                buffer.WriteString("**EndAt**  \n" + fmt.Sprintf("%v", r.EndAt) + "\n\n")
	}
	if IsSelectedItem("Hostname", items) {
                buffer.WriteString("**Hostname**  \n" + fmt.Sprintf("%v", r.Hostname) + "\n\n")
	}
	if IsSelectedItem("SystemTime", items) {
                buffer.WriteString("**SystemTime**  \n" + fmt.Sprintf("%f", *r.SystemTime) + "\n\n")
	}
	if IsSelectedItem("UserTime", items) {
                buffer.WriteString("**UserTime**  \n" + fmt.Sprintf("%f", *r.UserTime) + "\n\n")
	}

	return buffer.String()
}

// SendReportToCiscoSpark send Report to Cisco Spark
func SendReportToCiscoSpark(api *ciscospark.Client, r horenso.Report, roomId string, toPersonEmail string, items []string) {
	msg := GetMarkDown(r, items)

	markDownMessage := &ciscospark.MessageRequest{
		MarkDown: msg,
		RoomID:   roomId,
		ToPersonEmail: toPersonEmail,
	}

	_, _, err := api.Messages.Post(markDownMessage)
	if err != nil {
		panic(err)
	}
}

// IsSelectedItem returns key exists in slice
func IsSelectedItem(a string, list []string) bool {
	if len(list) == 0 {
		return false
	}

	if list[0] == "all" {
		return true
	}

	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
