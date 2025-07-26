package chatcontext

import (
	"fmt"
	"runtime"
	"slices"
)

const minimalContext = "Your responses should contain only the command itself, without explanations or additional text. " +
	"If multiple commands are required, provide them on separate lines."

const verboseContext = "You should provide detailed responses that include explanations. " +
	"For each shell command you generate, explain what the command does and how it works." +
	"The response should display nicely in the terminal. " +
	"Assume the user is already working within the terminal, so do not include instructions relating to opening the terminal application and copying the command. " +
	"When returning the shell command, do not format in code blocks and return as plain text instead."

var verbosityContext = map[Verbosity]string{
	VerbosityMinimal: minimalContext,
	VerbosityVerbose: verboseContext,
}

type Verbosity int

const (
	VerbosityMinimal Verbosity = 0
	VerbosityVerbose Verbosity = 1
)

type OS string

const (
	OSDarwin  OS = "darwin"
	OSLinux   OS = "linux"
	OSWindows OS = "windows"
)

func BuildContext(getOS OSGetter, verbosity Verbosity) string {
	contextOS := OSLinux
	systemOS := getOS()
	if slices.Contains([]OS{OSDarwin, OSLinux, OSWindows}, OS(systemOS)) {
		contextOS = OS(systemOS)
	}

	return fmt.Sprintf(
		"You are an AI assistant that translates natural language prompts into precise shell commands. "+
			"Use commonly available command-line tools unless otherwise specified. "+
			"Assume the user is running commands on the %s operating system. "+
			"%s",
		contextOS,
		verbosityContext[verbosity],
	)
}

type OSGetter func() string

func GetOS() string {
	return runtime.GOOS
}
