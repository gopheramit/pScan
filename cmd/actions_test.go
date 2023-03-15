package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/gopheramit/pScan/scan"
)

func setup(t *testing.T, hosts []string, initList bool) (string, func()) {
	tf, err := ioutil.TempFile("", "pScan")
	if err != nil {
		t.Fatal(err)
	}
	tf.Close()
	if initList {
		hl := &scan.HostList{}
		for _, h := range hosts {
			hl.Add(h)
		}
		if err := hl.Save(tf.Name()); err != nil {
			t.Fatal(err)
		}
	}
	return tf.Name(), func() {
		os.Remove(tf.Name())
	}
}

func TestHostAction(t *testing.T) {
	hosts := []string{
		"host1",
		"host2",
		"host3",
	}
	testCases := []struct {
		name           string
		args           []string
		expectdOut     string
		ininList       bool
		actionFunction func(io.Writer, string, []string) error
	}{
		{name: "AddAction",
			args:           hosts,
			expectdOut:     "Added host: host1\nAdded host: host2\nAdded host: host3\n",
			ininList:       false,
			actionFunction: addAction},
		{
			name:           "ListAction",
			expectdOut:     "host1\nhost2\nhost3\n",
			ininList:       true,
			actionFunction: listAction,
		},
		{
			name:           "DeleteAction",
			args:           []string{"host1", "host2"},
			expectdOut:     "Deleted host: host1\nDeleted host: host2\n",
			ininList:       true,
			actionFunction: deleteAction,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tf, cleanup := setup(t, hosts, tc.ininList)
			defer cleanup()

			var out bytes.Buffer

			if err := tc.actionFunction(&out, tf, tc.args); err != nil {
				t.Fatalf("Expected no error,got%q\n", err)
			}
			if out.String() != tc.expectdOut {
				t.Errorf("Expected output %q got %q\n", tc.expectdOut, out.String())
			}
		})
	}
}

func TestIntegration(t *testing.T) {
	hosts := []string{
		"host1",
		"host2",
		"host3",
	}
	tf, cleanup := setup(t, hosts, false)
	defer cleanup()
	delHost := "host2"
	hostsEnd := []string{
		"host1",
		"host3",
	}

	var out bytes.Buffer

	expectedOut := ""
	for _, v := range hosts {
		expectedOut += fmt.Sprintf("Added host: %s\n", v)

	}
	expectedOut += strings.Join(hosts, "\n")
	expectedOut += fmt.Sprintln()
	expectedOut += fmt.Sprintf("Deleted host: %s\n", delHost)
	expectedOut += strings.Join(hostsEnd, "\n")
	expectedOut += fmt.Sprintln()

	if err := addAction(&out, tf, hosts); err != nil {
		t.Fatalf("Expected no error,got %q\n", err)
	}
	if err := listAction(&out, tf, nil); err != nil {
		t.Fatalf("Expected no error,got%q\n", err)
	}
	if err := deleteAction(&out, tf, []string{delHost}); err != nil {
		t.Fatalf("Expected no error ,got %q\n", err)
	}
	if err := listAction(&out, tf, nil); err != nil {
		t.Fatalf("Expected no error ,got%q\n", err)
	}

	if out.String() != expectedOut {
		t.Errorf("Expectd out %q ,got %q\n instead", expectedOut, out.String())
	}
}
