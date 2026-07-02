package integration

import (
	"os"
	"testing"

	"phishing-platform/internal/sigma"
)

func TestSigmaPipeline(t *testing.T) {

	file := "sigma_rule.yml"

	err := sigma.GenerateRule(
		"attacker@example.com",
		670,
		"T1566 - Phishing",
		file,
	)

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(file)

	info, err := os.Stat(file)

	if err != nil {
		t.Fatal(err)
	}

	if info.Size() == 0 {
		t.Fatal("Sigma rule is empty")
	}

	t.Log("========= SIGMA PIPELINE =========")
	t.Log("File:", file)
	t.Log("Size:", info.Size())
	t.Log("=================================")
}