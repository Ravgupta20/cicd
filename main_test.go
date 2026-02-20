package main

import (
	"testing"
)

func TestGenerateWorkflow(t *testing.T) {
	workflow := GenerateWorkflow()

	if workflow.Name != "Basic C" {
		t.Errorf("Missing workflow name , got '%s'", workflow.Name)
	}

	if len(workflow.On) == 0 {
		t.Errorf("Missing workflow trigger, got '%v'", workflow.On)
	}

	buildJob, ok := workflow.Jobs["build"]

	if !ok {
		t.Fatalf("Missing workflow 'build' job .")
	}

	if buildJob.RunsOn == "" {
		t.Errorf("Missing RunsOn OS.")
	}
}
