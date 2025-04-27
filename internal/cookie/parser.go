package cookie

import (
	"fmt"
	"strings"
	"time"
)

// Cookie represents a cookie with its value and associated timestamp.
type Cookie struct {
	// Value is the unique identifier of the cookie.
	Value string
	// Timestamp is the time when the cookie was active.
	Timestamp time.Time
}

var (
	// ErrInvalidCookieString is returned when a cookie string is malformed.
	ErrInvalidCookieString = fmt.Errorf("invalid cookie string")
	// ErrInvalidCookieTime is returned when a cookie timestamp cannot be parsed.
	ErrInvalidCookieTime = fmt.Errorf("invalid cookie time")
)

// Parse parses a single line of cookie data into a Cookie struct.
//
// The expected input format is a comma-separated string: "cookie,timestamp".
// The timestamp should follow RFC3339 format (e.g., "2018-12-09T14:19:00+00:00").
//
// It returns a Cookie struct on success, or an error if the input is invalid.
//
// Parameters:
//   - cookie: a string representing the raw cookie log entry.
//   - timezone: the time zone to interpret the timestamp.
//
// Returns:
//   - Cookie: the parsed cookie data.
//   - error: if parsing fails due to format issues or invalid timestamp.
func Parse(cookie string, timezone *time.Location) (Cookie, error) {
	parts := strings.Split(cookie, ",")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return Cookie{}, fmt.Errorf("`%s`: %w", cookie, ErrInvalidCookieString)
	}

	timestamp, err := time.ParseInLocation(time.RFC3339, parts[1], timezone)
	if err != nil {
		return Cookie{}, fmt.Errorf("%w: %w", err, ErrInvalidCookieTime)
	}

	return Cookie{
		Value:     parts[0],
		Timestamp: timestamp,
	}, nil
}
