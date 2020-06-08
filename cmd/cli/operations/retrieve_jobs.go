package operations

import "github.com/arun-spire/flink-deployer/cmd/cli/flink"

// RetrieveJobs executes the logic required for retrieving
// the jobs from a Flink cluster
func (o RealOperator) RetrieveJobs() ([]flink.Job, error) {
	return o.FlinkRestAPI.RetrieveJobs()
}
