package s3

import (
	"awesomeProject1/awsutil"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func FindBucketRegion(ctx context.Context, bucket_name string)(err error){
	var region_hint string = "us-west-1"
	sess, err := awsutil.CreateNewSession("",  "publishing");

	if err != nil {
		return
	}
	region, err := s3manager.GetBucketRegion(ctx, sess, bucket_name, region_hint);
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "NotFound" {
			fmt.Fprintf(os.Stderr, "unable to find bucket %s's region not found\n", bucket_name)
		}
		return err
	}
	fmt.Printf("Bucket %s is in %s region\n", bucket_name, region)
	return err
}
