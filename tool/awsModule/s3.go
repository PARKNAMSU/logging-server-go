package awsModule

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"logFile.com/log-file-go/tool/common"
)

var (
	s3Client = s3.NewFromConfig(cfg)
	uploader = manager.NewUploader(s3Client)
)

func UploadS3(bucket string, loc string, body string, contentType string) (*manager.UploadOutput, error) {
	if common.ErrorLogging(cfgErr) {
		return nil, cfgErr
	}
	ioReader := strings.NewReader(body)
	upload := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(loc),
		Body:        ioReader,
		ContentType: aws.String(contentType),
	}

	result, err := uploader.Upload(context.TODO(), upload)

	if common.ErrorLogging(cfgErr) {
		return nil, err
	}

	return result, nil
}

func GetObjectListS3(bucket string, loc string) (*s3.ListObjectsV2Output, error) {
	if common.ErrorLogging(cfgErr) {
		return nil, cfgErr
	}
	input := &s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Delimiter: aws.String("/"),
		Prefix:    aws.String(loc + "/"),
	}
	output, err := s3Client.ListObjectsV2(context.TODO(), input)
	if common.ErrorLogging(err) {
		return nil, err
	}
	return output, nil
}

func GetObjectS3(bucket string, loc string) (*s3.GetObjectOutput, error) {
	if common.ErrorLogging(cfgErr) {
		return nil, cfgErr
	}
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(loc),
	}
	output, err := s3Client.GetObject(context.TODO(), input)

	if common.ErrorLogging(err) {
		return nil, err
	}
	return output, nil
}
