//Package efs provides gucumber integration tests suppport.
package efs

import (
	"github.com/dongfangx/aws-sdk-go/aws"
	"github.com/dongfangx/aws-sdk-go/internal/features/shared"
	"github.com/dongfangx/aws-sdk-go/service/efs"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@efs", func() {
		// FIXME remove custom region
		World["client"] = efs.New(&aws.Config{Region: "us-west-2"})
	})
}
