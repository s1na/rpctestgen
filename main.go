package main

import (
	"context"
	"fmt"
	"os"
	"regexp"

	"github.com/alexflint/go-arg"
)

const (
	HOST        string = "127.0.0.1"
	PORT        string = "13375"
	NETWORKPORT string = "13376"
	AUTHPORT    string = "13377"
)

type Args struct {
	ClientType  string `arg:"--client" help:"client type" default:"geth"`
	ClientBin   string `arg:"--bin" help:"path to client binary" default:"geth"`
	OutDir      string `arg:"--out" help:"directory where test fixtures will be written" default:"tests"`
	ChainDir    string `arg:"--chain" help:"path to directory with chain.rlp and genesis.json"`
	Verbose     bool   `arg:"-v,--verbose" help:"verbosity level of rpctestgen"`
	LogLevel    string `arg:"--loglevel" help:"log level of client" default:"info"`
	TestsRegexp string `arg:"--tests" help:"regex of tests to fill" default:".*"`

	tests       *regexp.Regexp
	logLevelInt int
}

type ArgsKey struct{}

var ARGS = ArgsKey{}

func main() {
	var args Args
	arg.MustParse(&args)

	lvl, err := loglevelToInt(args.LogLevel)
	if err != nil {
		exit(err)
	}
	args.logLevelInt = lvl

	if args.tests, err = regexp.Compile(args.TestsRegexp); err != nil {
		exit(err)
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, ARGS, &args)

	if err := runGenerator(ctx); err != nil {
		exit(err)
	}
}

func exit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func loglevelToInt(lvl string) (int, error) {
	switch lvl {
	case "err":
		return 1, nil
	case "warn":
		return 2, nil
	case "info":
		return 3, nil
	case "debug":
		return 4, nil
	case "trace":
		return 5, nil
	default:
		return 0, fmt.Errorf("unknown log level: %s", lvl)
	}
}
