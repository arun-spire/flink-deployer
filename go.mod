module flink-deployer

go 1.13

replace github.com/arun-spire/flink-deployer => ./

require (
	github.com/arun-spire/flink-deployer v0.0.0-20191209121213-b2d81476874b
	github.com/aws/aws-sdk-go v1.29.34
	github.com/bmatcuk/doublestar v1.2.2
	github.com/bsm/bfs v0.10.4
	github.com/bsm/bfs/bfss3 v0.0.0-20200526081558-b611316062db
	github.com/cenkalti/backoff v2.0.0+incompatible
	github.com/hashicorp/go-retryablehttp v0.0.0-20180718195005-e651d75abec6
	github.com/stretchr/testify v1.5.1
	github.com/urfave/cli v1.20.0
)
