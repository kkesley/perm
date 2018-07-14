package perm

import (
	"encoding/json"

	parser "github.com/kkesley/s3-parser"
)

//DownloadRole to s3 input bucket
func DownloadRole(bucket string, region string, folder string) ([]RawPolicy, error) {
	//download the bytes from s3
	roleByte, err := parser.GetS3DocumentDefault(region, bucket, folder+"/role.json")
	if err != nil {
		return make([]RawPolicy, 0), err
	}
	var policies []RawPolicy
	if err := json.Unmarshal(roleByte, &policies); err != nil {
		return make([]RawPolicy, 0), err
	}
	return policies, nil
}
