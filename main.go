package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/sters/go-project-boilerplate/boilerplate"
)

//nolint:gochecknoglobals
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func getVersion() string {
	if version != "dev" {
		return version
	}

	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "" && info.Main.Version != "(devel)" {
			return info.Main.Version
		}
	}

	return version
}

func getCommit() string {
	if commit != "none" {
		return commit
	}

	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				if len(setting.Value) > 7 {
					return setting.Value[:7]
				}

				return setting.Value
			}
		}
	}

	return commit
}

func getDate() string {
	if date != "unknown" {
		return date
	}

	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.time" {
				if t, err := time.Parse(time.RFC3339, setting.Value); err == nil {
					return t.UTC().Format("2006-01-02T15:04:05Z")
				}

				return setting.Value
			}
		}
	}

	return date
}

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "show version information")
	flag.BoolVar(&showVersion, "v", false, "show version information (short)")
	flag.Parse()

	if showVersion {
		fmt.Printf("Version:    %s\n", getVersion())
		fmt.Printf("Commit:     %s\n", getCommit())
		fmt.Printf("Built:      %s\n", getDate())
		fmt.Printf("Go version: %s\n", runtime.Version())
		fmt.Printf("OS/Arch:    %s/%s\n", runtime.GOOS, runtime.GOARCH)
		os.Exit(0)
	}

	fmt.Printf("%s World!", boilerplate.Hello())
}
