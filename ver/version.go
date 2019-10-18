package ver

import "fmt"

// Version info
var (
	Version   string
	BuildTime string
	Commit    string
)

// Print version info
func Print() {
	fmt.Printf("  Version:   %s\n", Version)
	fmt.Printf("  BuildTime: %s\n", BuildTime)
	fmt.Printf("  Commit:    %s\n", Commit)
}

// Info getter
func Info() map[string]string {
	return map[string]string{
		"version":   Version,
		"commit":    Commit,
		"buildtime": BuildTime,
	}
}
