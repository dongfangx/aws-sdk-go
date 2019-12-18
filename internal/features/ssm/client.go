//Package ssm provides gucumber integration tests suppport.
package ssm

import (
	"github.com/dongfangx/aws-sdk-go/internal/features/shared"
	"github.com/dongfangx/aws-sdk-go/service/ssm"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@ssm", func() {
		World["client"] = ssm.New(nil)
	})
}
