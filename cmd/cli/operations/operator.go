package operations

import (
	"github.com/arun-spire/flink-deployer/cmd/cli/flink"
	"github.com/bsm/bfs"
)

// Operator is an interface which contains all the functionality
// that the deployer exposes
type Operator interface {
	Deploy(d Deploy) error
	Update(u UpdateJob) error
	RetrieveJobs() ([]flink.Job, error)
	Terminate(t TerminateJob) error
}

// RealOperator is the Operator used in the production code
type RealOperator struct {
	Filesystem   bfs.Bucket
	FlinkRestAPI flink.FlinkRestAPI
}
