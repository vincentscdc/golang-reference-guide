package h2

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gitlab.com/AgentNemo/goradios"
)

type stationsFetcherMock struct{}

func (sfm *stationsFetcherMock) FetchStations(by goradios.StationsBy, term string) []goradios.Station {
	return []goradios.Station{
		{
			Name: "Radio 1",
			URL:  "http://radio1.com",
		},
		{
			Name: "Radio 2",
			URL:  "http://radio2.com",
		},
	}
}

func TestHi(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		expectError bool
	}{
		{
			name:        "Test Hi",
			expectError: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest("GET", "/", nil)
			rec := httptest.NewRecorder()
			Hi(&stationsFetcherMock{})(rec, req)

			if tt.expectError {
				if rec.Result().StatusCode == http.StatusOK {
					t.Errorf("Hi() expected status code 200, got %v", rec.Result().StatusCode)
				}

				return
			}

			if rec.Result().StatusCode != http.StatusOK {
				t.Errorf("Hi() expected status code 200, got %v", rec.Result().StatusCode)
			}
		})
	}
}
