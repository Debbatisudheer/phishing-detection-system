package virustotal

import (
	"reflect"
	"testing"
)

func TestCheckURLReputation(t *testing.T) {

	tests := []struct {
		name     string
		jsonData []byte
		expected []string
	}{
		{
			name: "Malicious URL",
			jsonData: []byte(`{
				"data":{
					"attributes":{
						"last_analysis_stats":{
							"malicious":10,
							"suspicious":0,
							"harmless":5
						}
					}
				}
			}`),
			expected: []string{
				"VirusTotal malicious URL detected",
			},
		},
		{
			name: "Suspicious URL",
			jsonData: []byte(`{
				"data":{
					"attributes":{
						"last_analysis_stats":{
							"malicious":0,
							"suspicious":2,
							"harmless":20
						}
					}
				}
			}`),
			expected: []string{
				"VirusTotal suspicious URL detected",
			},
		},
		{
			name: "Harmless URL",
			jsonData: []byte(`{
				"data":{
					"attributes":{
						"last_analysis_stats":{
							"malicious":0,
							"suspicious":0,
							"harmless":40
						}
					}
				}
			}`),
			expected: []string{},
		},
		{
			name:     "Invalid JSON",
			jsonData: []byte(`abc`),
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := CheckURLReputation(tc.jsonData)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf(
					"expected %v got %v",
					tc.expected,
					result,
				)
			}
		})
	}
}