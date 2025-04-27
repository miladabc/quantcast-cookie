package cli

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		expectedQuery Query
		expectErr     error
	}{
		{
			name: "valid arguments",
			args: []string{"-f", "cookie_log.csv", "-d", "2018-12-09"},
			expectedQuery: Query{
				CookieFilePath:  "cookie_log.csv",
				CookieTimestamp: time.Date(2018, 12, 9, 0, 0, 0, 0, timezone),
			},
			expectErr: nil,
		},
		{
			name:          "missing file argument",
			args:          []string{"-d", "2018-12-09"},
			expectedQuery: Query{},
			expectErr:     ErrEmptyArg,
		},
		{
			name:          "missing date argument",
			args:          []string{"-f", "cookie_log.csv"},
			expectedQuery: Query{},
			expectErr:     ErrEmptyArg,
		},
		{
			name:          "invalid date format",
			args:          []string{"-f", "cookie_log.csv", "-d", "09-12-2018"},
			expectedQuery: Query{},
			expectErr:     ErrInvalidDate,
		},
		{
			name:          "no arguments",
			args:          []string{"cmd"},
			expectedQuery: Query{},
			expectErr:     ErrEmptyArg,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			query, err := Parse(tc.args)
			require.ErrorIs(t, err, tc.expectErr)
			require.Equal(t, tc.expectedQuery, query)
		})
	}
}
