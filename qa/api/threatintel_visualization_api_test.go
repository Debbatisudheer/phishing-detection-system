package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	iocapi "phishing-platform/internal/api/ioc"
)

func TestThreatIntelVisualizationAPI(t *testing.T) {

	setupDatabase()

	tests := []struct {
		name           string
		handler        http.HandlerFunc
		method         string
		url            string
		expectedStatus int
	}{
		{
			"IOC Graph GET",
			iocapi.IOCGraphHandler,
			http.MethodGet,
			"/api/ioc-graph",
			http.StatusOK,
		},
		{
			"IOC Graph Invalid Method",
			iocapi.IOCGraphHandler,
			http.MethodPost,
			"/api/ioc-graph",
			http.StatusMethodNotAllowed,
		},
		{
			"IOC Sources Missing IOC",
			iocapi.IOCSourcesHandler,
			http.MethodGet,
			"/api/ioc-sources",
			http.StatusBadRequest,
		},
		{
			"IOC Sources Invalid Method",
			iocapi.IOCSourcesHandler,
			http.MethodPost,
			"/api/ioc-sources?ioc=evil.com",
			http.StatusMethodNotAllowed,
		},
		{
			"IOC Trends GET",
			iocapi.IOCTrendsHandler,
			http.MethodGet,
			"/api/ioc-trends",
			http.StatusOK,
		},
		{
			"IOC Trends Invalid Method",
			iocapi.IOCTrendsHandler,
			http.MethodPost,
			"/api/ioc-trends",
			http.StatusMethodNotAllowed,
		},
		{
			"MITRE Heatmap GET",
			iocapi.MITREHeatmapHandler,
			http.MethodGet,
			"/api/mitre-heatmap",
			http.StatusOK,
		},
		{
			"MITRE Heatmap Invalid Method",
			iocapi.MITREHeatmapHandler,
			http.MethodPost,
			"/api/mitre-heatmap",
			http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			req := httptest.NewRequest(
				tc.method,
				tc.url,
				nil,
			)

			rec := httptest.NewRecorder()

			tc.handler(
				rec,
				req,
			)

			if rec.Code != tc.expectedStatus {

				t.Fatalf(
					"%s expected %d got %d",
					tc.name,
					tc.expectedStatus,
					rec.Code,
				)
			}
		})
	}
}
