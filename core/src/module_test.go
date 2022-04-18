package src

import (
	"testing"
)

func TestModuleName(t *testing.T) {
	if ProjectName() != "core" {
		t.Errorf("Project name `%s` incorrect", ProjectName())
	}
}
