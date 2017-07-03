package reporter

import (
	"encoding/json"
	"io/ioutil"
	"os"
	//"reflect"
	"testing"

	"github.com/Songmu/horenso"
	"github.com/stretchr/testify/assert"
)

func TestGetMakrDown(t *testing.T) {
	f, _ := os.Open("../fixtures/report_exit_0.json")
	jsonBytes, _ := ioutil.ReadAll(f)

	var r horenso.Report
	json.Unmarshal(jsonBytes, &r)

	items := []string{"all"}
	a := GetMarkDown(r, items)

	assert.Equal(t, "## horenso Reporter\n**Result**  \n`command exited with code: 0\n`\n\n**Output**  \n`1\n95030\n\n`\n\n**Stdout**  \n`1\n\n`\n\n**Stderr**  \n`95030\n\n`\n\n**Command**  \n`perl -E 'say 1;warn \"$$\\n\";'\n`\n\n**CommandArgs**  \n`[perl -E say 1;warn \"$$\\n\";]\n`\n\n**Pid**  \n95030\n\n**ExitCode**  \n0\n\n**StartAt**  \n2015-12-28 00:37:10.494282399 +0900 JST\n\n**EndAt**  \n2015-12-28 00:37:10.546466379 +0900 JST\n\n**Hostname**  \nwebserver.example.com\n\n**SystemTime**  \n0.000123\n\n**UserTime**  \n0.000456\n\n", a)
}

func TestIsSelectedItem(t *testing.T) {
	assert.True(t, IsSelectedItem("ExitCode", []string{"all"}))
	assert.True(t, IsSelectedItem("ExitCode", []string{"ExitCode", "Output"}))
	assert.False(t, IsSelectedItem("Stdout", []string{"ExitCode", "Output"}))
	assert.False(t, IsSelectedItem("ExitCode", []string{}))
}
