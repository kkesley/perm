package perm

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//UploadRole to s3 input bucket
func UploadRole(bucket string, region string, folder string, policy []RawPolicy) error {
	//get the actions in bytes
	roleByte, err := json.Marshal(policy)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//upload the converted policy to an s3 bucket
	sess := session.Must(
		session.NewSession(&aws.Config{
			Region: aws.String(region),
		}),
	)
	svc := s3.New(sess)
	if _, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(folder + "/" + "role.json"),
		Body:   bytes.NewReader(roleByte),
	}); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
