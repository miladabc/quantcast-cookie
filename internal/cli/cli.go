package cli

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// Query represents the parsed CLI arguments for the application.
// It contains the path to the cookie log file and the target timestamp.
type Query struct {
	CookieFilePath  string
	CookieTimestamp time.Time
}

// timezone defines the time location used for parsing timestamps.
// It is fixed to UTC.
var timezone = time.UTC

var (
	// ErrEmptyArg is returned when either the filename or date argument is missing.
	ErrEmptyArg = fmt.Errorf("empty filename or date")
	// ErrInvalidDate is returned when the provided date argument cannot be parsed.
	ErrInvalidDate = fmt.Errorf("invalid date")
)

// Parse processes CLI arguments to extract the cookie file path and target date.
//
// It expects two flags:
//
//	-f string: Path to the cookie log file.
//	-d string: Date to filter cookies in UTC timezone (format: YYYY-MM-DD).
//
// Parameters:
//   - args: a slice of strings representing the command-line arguments.
//
// Returns:
//   - Query: the parsed query containing the file path and target timestamp.
//   - error: if parsing fails or arguments are invalid.
func Parse(args []string) (Query, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.String("f", "", "Path to cookie log file")
	f.String("d", "", "Date to filter cookies in UTC timezone (YYYY-MM-DD)")

	err := f.Parse(args)
	if err != nil {
		f.Usage()
		return Query{}, fmt.Errorf("parsing flags: %w", err)
	}

	filename := f.Lookup("f").Value.String()
	date := f.Lookup("d").Value.String()

	if filename == "" || date == "" {
		f.Usage()
		return Query{}, ErrEmptyArg
	}

	targetDate, err := time.ParseInLocation(time.DateOnly, date, timezone)
	if err != nil {
		f.Usage()
		return Query{}, fmt.Errorf("%w: %w", ErrInvalidDate, err)
	}

	return Query{
		CookieFilePath:  filename,
		CookieTimestamp: targetDate,
	}, nil
}
