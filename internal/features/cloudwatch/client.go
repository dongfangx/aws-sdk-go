//Package cloudwatch provides gucumber integration tests suppport.
package cloudwatch

import (
	"github.com/dongfangx/aws-sdk-go/internal/features/shared"
	"github.com/dongfangx/aws-sdk-go/service/cloudwatch"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@cloudwatch", func() {
		World["client"] = cloudwatch.New(nil)
	})
}
