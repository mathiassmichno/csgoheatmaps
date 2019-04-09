package common

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
    "strings"

	com "github.com/markus-wa/demoinfocs-golang/common"
)

// DemoPathFromArgs returns the value of the -demo command line flag.
// Panics if an error occurs.
func DemoPathFromArgs() string {
	fl := new(flag.FlagSet)

	demPathPtr := fl.String("demo", "", "Demo file `path`")

	err := fl.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}

	demPath := *demPathPtr

	return demPath
}

func OptionsFromArgs() (string, int64, com.Team) {
	fl := new(flag.FlagSet)

    var demoPath string
    var teamStr string
    var steamId int64
	fl.StringVar(&demoPath, "demo", "", "Demo file `path`")
	fl.StringVar(&teamStr, "team", "", "Team T or CT")
	fl.Int64Var(&steamId, "steamid", 0, "steamid")

    var team com.Team

    if strings.ToLower(teamStr) == "ct" {
        team = com.TeamCounterTerrorists
    } else if strings.ToLower(teamStr) == "t" {
        team = com.TeamTerrorists
    }

	err := fl.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}

	return demoPath, steamId, team
}

// RedirectStdout redirects standard output to dev null.
// Panics if an error occurs.
func RedirectStdout(f func()) {
	// Redirect stdout, the resulting image is written to this
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	os.Stdout = w

	// Discard the output in a separate goroutine so writing to stdout can't block indefinitely
	go func() {
		for err := error(nil); err == nil; _, err = io.Copy(ioutil.Discard, r) {
		}
	}()

	f()

	os.Stdout = old
}
