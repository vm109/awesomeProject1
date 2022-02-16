package s3

import (
	"context"
	"fmt"
	"testing"
)

func TestFindBucketRegion(t *testing.T) {
	ctx := context.Background()
	err := FindBucketRegion(ctx, "hjdfsb")
	fmt.Println(err)
}
