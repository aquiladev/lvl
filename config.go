package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/btcsuite/btcutil"
	"github.com/jessevdk/go-flags"
)

const (
	defaultDataDirname = "data"
)

var (
	defaultHomeDir = btcutil.AppDataDir("lvl", false)
	defaultDataDir = filepath.Join(defaultHomeDir, defaultDataDirname)
)

type config struct {
	ShowVersion  bool    `short:"V" long:"version" description:"Display version information and exit."`
	DataDir      string  `short:"d" long:"datadir" description:"Directory to store data."`
	MaxDiskUsage float64 `short:"m" long:"maxdiskusage" description:"Max disk usage in percentage." default:"95.0"`
}

func loadConfig() (*config, error) {
	cfg := config{
		DataDir: defaultDataDir,
	}

	// Show the version and exit if the version flag was specified.
	appName := filepath.Base(os.Args[0])
	appName = strings.TrimSuffix(appName, filepath.Ext(appName))
	usageMessage := fmt.Sprintf("Use %s -h to show usage", appName)

	// Parse command line options.
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		if e, ok := err.(*flags.Error); !ok || e.Type != flags.ErrHelp {
			fmt.Fprintln(os.Stderr, usageMessage)
		}
		return nil, err
	}

	if cfg.ShowVersion {
		fmt.Println(appName, "version", version())
		os.Exit(0)
	}

	return &cfg, nil
}
