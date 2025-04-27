package cookie

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name           string
		cookieLog      string
		expectedCookie Cookie
		expectedErr    error
	}{
		{
			"empty log",
			"",
			Cookie{},
			ErrInvalidCookieString,
		},
		{
			"no cookie timestamp",
			"AtY0laUfhglK3lC7,",
			Cookie{},
			ErrInvalidCookieString,
		},
		{
			"no cookie value",
			",2018-12-07T23:30:00+00:00",
			Cookie{},
			ErrInvalidCookieString,
		},
		{
			"invalid cookie timestamp",
			"AtY0laUfhglK3lC7,2018-12-07 23:30:00+00:00",
			Cookie{},
			ErrInvalidCookieTime,
		},
		{
			"extra fields",
			"AtY0laUfhglK3lC7,2018-12-07 23:30:00+00:00,unexpected",
			Cookie{},
			ErrInvalidCookieString,
		},
		{
			"valid cookie log",
			"AtY0laUfhglK3lC7,2018-12-07T23:30:00+00:00",
			Cookie{
				Value:     "AtY0laUfhglK3lC7",
				Timestamp: time.Date(2018, 12, 7, 23, 30, 0, 0, time.UTC),
			},
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cookie, err := Parse(tc.cookieLog, time.UTC)
			require.ErrorIs(t, err, tc.expectedErr)
			require.Equal(t, tc.expectedCookie, cookie)
		})
	}
}
