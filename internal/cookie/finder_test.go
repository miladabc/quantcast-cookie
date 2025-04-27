package cookie

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFindMostActive(t *testing.T) {
	tests := []struct {
		name            string
		logs            string
		targetDate      string
		expectedCookies []string
	}{
		{
			name: "single most active cookie",
			logs: `cookie,timestamp
				cookie1,2018-12-09T14:19:00+00:00
				cookie2,2018-12-09T10:13:00+00:00
				cookie1,2018-12-09T06:19:00+00:00
				cookie3,2018-12-09T12:19:00+00:00`,
			targetDate:      "2018-12-09",
			expectedCookies: []string{"cookie1"},
		},
		{
			name: "multiple most active cookies",
			logs: `cookie,timestamp
				cookie1,2018-12-09T00:00:00+00:00
				cookie2,2018-12-09T01:00:00+00:00
				cookie1,2018-12-09T02:00:00+00:00
				cookie2,2018-12-09T03:00:00+00:00`,
			targetDate:      "2018-12-09",
			expectedCookies: []string{"cookie1", "cookie2"},
		},
		{
			name: "no cookies found for date",
			logs: `cookie,timestamp
				cookie1,2018-12-08T00:00:00+00:00
				cookie2,2018-12-08T01:00:00+00:00`,
			targetDate:      "2018-12-09",
			expectedCookies: nil,
		},
		{
			name: "invalid cookie lines are ignored",
			logs: `cookie,timestamp

				cookie1,2018-12-09T00:00:00+00:00
				invalid_line_without_comma
				cookie1,invalid_timestamp

				cookie2,2018-12-09T01:00:00+00:00`,
			targetDate:      "2018-12-09",
			expectedCookies: []string{"cookie1", "cookie2"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logs := strings.NewReader(tc.logs)
			targetDate, err := time.ParseInLocation(time.DateOnly, tc.targetDate, time.UTC)
			require.NoError(t, err)

			cookies, err := FindMostActive(logs, targetDate)
			require.NoError(t, err)
			require.ElementsMatch(t, cookies, tc.expectedCookies)
		})
	}
}
