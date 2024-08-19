package parser

import "strings"

type AppArguments struct {
	Executable      string
	RunInBackground bool
}

func ParseArguments(args []string) AppArguments {
	appArgs := AppArguments{
		Executable:      args[0],
		RunInBackground: false,
	}
	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			continue
		}
		if arg == "-background" {
			appArgs.RunInBackground = true
		}
	}
	return appArgs
}