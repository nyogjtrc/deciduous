package logging

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLogger(t *testing.T) {
	Convey("Given logging package", t, func() {
		Convey("When package init", func() {
			Convey("Then loggers should not be nil", func() {
				So(logger, ShouldNotBeEmpty)
				So(slogger, ShouldNotBeEmpty)
				So(L(), ShouldNotBeEmpty)
				So(S(), ShouldNotBeEmpty)
			})
		})
	})
}
