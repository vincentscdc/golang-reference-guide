package h4

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/goccy/go-json"
	"gitlab.com/AgentNemo/goradios"
	"go.uber.org/goleak"
)

type stationsFetcherMock struct {
	Stations    []goradios.Station
	wantErrJSON bool
}

func (sfm *stationsFetcherMock) FetchStations(by goradios.StationsBy, term string) {
	sfm.Stations = testListStations
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
		name                  string
		responseWriter        http.ResponseWriter
		stationFetcher        StationsFetcher
		expectedHTTPCodeError bool
		expectedWriterError   bool
	}{
		{
			name:           "happy path",
			responseWriter: httptest.NewRecorder(),
			stationFetcher: &stationsFetcherMock{},
		},
		{
			name:                  "error json",
			responseWriter:        httptest.NewRecorder(),
			stationFetcher:        &stationsFetcherMock{wantErrJSON: true},
			expectedHTTPCodeError: true,
		},
		{
			name:                "error writer",
			responseWriter:      &faultyResponseWriter{},
			stationFetcher:      &stationsFetcherMock{},
			expectedWriterError: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest("GET", "/", nil)

			var buf bytes.Buffer
			if tt.expectedWriterError {
				log.SetOutput(&buf)
				defer func() {
					log.SetOutput(os.Stderr)
				}()
			}

			Hi(tt.stationFetcher)(tt.responseWriter, req)

			if tt.expectedWriterError {
				receivedLog := buf.String()

				if !strings.Contains(receivedLog, "mock error") {
					t.Errorf("no log received in case of writer error: %s", receivedLog)
				}

				return
			}

			rw := tt.responseWriter.(*httptest.ResponseRecorder)

			if (rw.Result().StatusCode != http.StatusOK) != tt.expectedHTTPCodeError {
				t.Errorf("Hi() expected error %t, got status code %d", tt.expectedHTTPCodeError, rw.Result().StatusCode)
			}
		})
	}
}

func BenchmarkHi(b *testing.B) {
	stationFetcher := &stationsFetcherMock{}

	hi := Hi(stationFetcher)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()

		req := httptest.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()

		b.StartTimer()

		hi(rec, req)
	}
}

// go test -bench=Hi > old.txt
// go test -bench=Hi > new.txt
// benchstat old.txt new.txt

var testListStations []goradios.Station

// go test ./... -cover -leak
func TestMain(m *testing.M) {
	leak := flag.Bool("leak", false, "use leak detector")
	flag.Parse()

	if *leak {
		goleak.VerifyTestMain(m)

		return
	}

	testListStations = goradios.FetchAllStations()

	code := m.Run()

	os.Exit(code)
}

func TestStations_FetchStations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
	}{
		{
			name: "happy path",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &Stations{}
			s.FetchStations(goradios.StationsByCountry, "Singapore")

			if _, err := s.MarshalJSON(); err != nil {
				t.Errorf("could not marshal stations from live api")
			}
		})
	}
}
