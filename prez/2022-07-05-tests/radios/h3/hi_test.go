package h3

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gitlab.com/AgentNemo/goradios"
)

type stationsFetcherMock struct {
	Stations    []goradios.Station
	wantErrJSON bool
}

func (sfm *stationsFetcherMock) FetchStations(by goradios.StationsBy, term string) {
	sfm.Stations = []goradios.Station{
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

func (sfm *stationsFetcherMock) MarshalJSON() ([]byte, error) {
	if sfm.wantErrJSON {
		return nil, fmt.Errorf("mock error")
	}

	return json.Marshal(sfm.Stations)
}

type faultyResponseWriter struct{}

func (frw faultyResponseWriter) Header() http.Header {
	m := make(map[string][]string)

	return m
}

func (frw faultyResponseWriter) Write([]byte) (int, error) {
	return 0, fmt.Errorf("mock error")
}

func (frw faultyResponseWriter) WriteHeader(statusCode int) {
}

func TestHi(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                 string
		faultyResponseWriter bool
		faultyStationFetcher bool
		expectedError        bool
	}{
		{
			name:          "happy path",
			expectedError: false,
		},
		{
			name:                 "error json",
			faultyStationFetcher: true,
			expectedError:        true,
		},
		{
			name:                 "error writer",
			faultyResponseWriter: true,
			expectedError:        true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest("GET", "/", nil)
			rec := httptest.NewRecorder()

			stationFetcher := &stationsFetcherMock{wantErrJSON: tt.faultyStationFetcher}

			if tt.faultyResponseWriter {
				Hi(stationFetcher)(&faultyResponseWriter{}, req)
			} else {
				Hi(stationFetcher)(rec, req)
			}

			if tt.expectedError {
				if rec.Result().StatusCode == http.StatusOK && !tt.faultyResponseWriter {
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

func BenchmarkHi(b *testing.B) {
	stationFetcher := &stationsFetcherMock{}

	hi := Hi(stationFetcher)

	for i := 0; i < b.N; i++ {
		b.StopTimer()

		req := httptest.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()

		b.StartTimer()

		hi(rec, req)
	}
}
