package engine

import (
	"github.com/rancher/rancher/pkg/controllers/user/pipeline/engine/jenkins"
	"github.com/rancher/rancher/pkg/controllers/user/pipeline/utils"
	"github.com/rancher/types/apis/management.cattle.io/v3"
	"github.com/rancher/types/config"
)

type PipelineEngine interface {
	RunPipeline(pipeline *v3.Pipeline, triggerType string) error
	RerunHistory(execution *v3.PipelineExecution) error
	StopHistory(execution *v3.PipelineExecution) error
	GetStepLog(execution *v3.PipelineExecution, stage int, step int) (string, error)
	OnHistoryCompelte(execution *v3.PipelineExecution)
	SyncExecution(execution *v3.PipelineExecution) (bool, error)
}

func New(cluster *config.UserContext) (PipelineEngine, error) {

	url, err := utils.GetJenkinsURL(cluster)
	if err != nil {
		return nil, err
	}
	user := "admin"
	token := "admin"
	client, err := jenkins.New(url, user, token)
	if err != nil {
		return nil, err
	}
	engine := &jenkins.JenkinsEngine{
		Client:  client,
		Cluster: cluster,
	}
	return engine, nil
}