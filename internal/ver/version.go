// Package ver provides build version info
//
// reference:
// - https://github.com/docker/cli/blob/master/cli/version/version.go
// - https://stackoverflow.com/questions/47509272/how-to-set-package-variable-using-ldflags-x-in-golang-build
package ver

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

// Version info
// overwrite by ldflags
var (
	Version   string = "unknown"
	BuildTime string = "unknown"
	Commit    string = "unknown"
)

// Print write version info to standard output
func Print() {
	fmt.Printf(
		"  Version: %s\n  GoVersion: %s\n  BuildTime: %s\n  Commit: %s\n",
		Version,
		runtime.Version(),
		BuildTime,
		Commit,
	)
}

// Maps of version info
func Maps() map[string]string {
	return map[string]string{
		"version":    Version,
		"go_version": runtime.Version(),
		"build_time": BuildTime,
		"commit":     Commit,
	}
}

// Router for gin
func Router(r *gin.Engine) {
	r.GET("/api/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, Maps())
	})
}
