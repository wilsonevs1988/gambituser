package awsgo

import (
	"context"
	"fmt"
	"gambituser/src/app/constants"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	Ctx context.Context
	Cfg aws.Config
	err error
)

func InitAws() {
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion(constants.RegionAws))
	if err != nil {
		panic(fmt.Sprintf(constants.ErrorInitAws, err.Error()))
	}
}
