package pipelinerun

import (
	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/apis"

	pipelinev1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
)

// State represents the state of a Pipeline.
type State int

const (
	Pending State = iota
	Failed
	Successful
	Error
)

func (s State) String() string {
	names := [...]string{
		"Pending",
		"Failed",
		"Successful",
		"Error"}
	return names[s]
}

// getPipelineRunState returns whether or not a PipelineRun was successful or
// not.
//
// It can return a Pending result if the task has not yet completed.
// TODO: will likely need to work out if a task was killed OOM.
func getPipelineRunState(p *pipelinev1.PipelineRun) State {
	for _, c := range p.Status.Conditions {
		if c.Type == apis.ConditionSucceeded {
			switch c.Status {
			case
				corev1.ConditionFalse:
				return Failed
			case corev1.ConditionTrue:
				return Successful
			case corev1.ConditionUnknown:
				return Pending
			}
		}
	}
	return Pending
}
