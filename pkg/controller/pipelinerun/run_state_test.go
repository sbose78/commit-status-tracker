package pipelinerun

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/apis"

	pipelinev1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	tb "github.com/tektoncd/pipeline/test/builder"
)

func TestGetPipelineRunStatus(t *testing.T) {
	statusTests := []struct {
		conditionType   apis.ConditionType
		conditionStatus corev1.ConditionStatus
		want            State
	}{
		{apis.ConditionSucceeded, corev1.ConditionTrue, Successful},
		{apis.ConditionSucceeded, corev1.ConditionUnknown, Pending},
		{apis.ConditionSucceeded, corev1.ConditionFalse, Failed},
	}

	for _, tt := range statusTests {
		s := getPipelineRunState(makePipelineRunWithCondition(tt.conditionType, tt.conditionStatus))
		if s != tt.want {
			t.Errorf("getPipelineRunState(%s) got %v, want %v", tt.conditionStatus, s, tt.want)
		}
	}
}

func makePipelineRunWithCondition(s apis.ConditionType, c corev1.ConditionStatus) *pipelinev1.PipelineRun {
	return tb.PipelineRun(pipelineRunName, testNamespace, tb.PipelineRunSpec(
		"tomatoes",
	), tb.PipelineRunStatus(tb.PipelineRunStatusCondition(
		apis.Condition{Type: s, Status: c}),
		tb.PipelineRunTaskRunsStatus("trname", &pipelinev1.PipelineRunTaskRunStatus{
			PipelineTaskName: "task-1",
		})))
}
