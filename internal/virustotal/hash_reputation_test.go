package virustotal

import (
	"reflect"
	"testing"
)

func TestCheckHashReputation(t *testing.T) {

	tests := []struct {
		name     string
		jsonData []byte
		expected []string
	}{
		{
			name:     "Hash Not Found",
			jsonData: []byte(`{"message":"hash_not_found"}`),
			expected: []string{
				"VirusTotal: hash not found",
			},
		},
		{
			name: "Malicious",
			jsonData: []byte(`{
				"data":{
					"attributes":{
						"last_analysis_stats":{
							"malicious":5,
							"suspicious":0,
							"harmless":10
						}
					}
				}
			}`),
			expected: []string{
				"VirusTotal malicious hash detected",
			},
		},
		{
			name: "Suspicious",
			jsonData: []byte(`{
				"data":{
					"attributes":{
						"last_analysis_stats":{
							"malicious":0,
							"suspicious":3,
							"harmless":20
						}
					}
				}
			}`),
			expected: []string{
				"VirusTotal suspicious hash detected",
			},
		},
		{
			name: "Harmless",
			jsonData: []byte(`{
				"data":{
					"attributes":{
						"last_analysis_stats":{
							"malicious":0,
							"suspicious":0,
							"harmless":70
						}
					}
				}
			}`),
			expected: []string{
				"VirusTotal hash marked harmless",
			},
		},
		{
			name:     "Invalid JSON",
			jsonData: []byte(`abc`),
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := CheckHashReputation(tc.jsonData)

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