//Package rds provides gucumber integration tests suppport.
package rds

import (
	"github.com/dongfangx/aws-sdk-go/internal/features/shared"
	"github.com/dongfangx/aws-sdk-go/service/rds"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@rds", func() {
		World["client"] = rds.New(nil)
	})
}
