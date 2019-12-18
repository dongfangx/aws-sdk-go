//Package storagegateway provides gucumber integration tests suppport.
package storagegateway

import (
	"github.com/dongfangx/aws-sdk-go/internal/features/shared"
	"github.com/dongfangx/aws-sdk-go/service/storagegateway"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@storagegateway", func() {
		World["client"] = storagegateway.New(nil)
	})
}
