package awsModule

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/joho/godotenv"
)

var (
	_           = godotenv.Load()
	cfg, cfgErr = config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("ap-southeast-2"),
		config.WithSharedConfigProfile(
			os.Getenv("AWS_CREDENTIAL_PROFILE")))
)
