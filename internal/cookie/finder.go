package cookie

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

const fileHeader = "cookie,timestamp"

// FindMostActive reads cookies from the provided io.Reader,
// filters them by the specified target date, and returns the
// most active cookies for that day.
//
// A cookie is considered "active" if it appears most frequently
// on the specified date. If multiple cookies have the same highest
// frequency, all of them are returned.
//
// Parameters:
//   - cookies: an io.Reader containing cookie logs, each line in "cookie,timestamp" format.
//   - targetDate: the date for which to find the most active cookies.
//
// Returns:
//   - []string: list of most active cookie values for the target date.
//   - error: if there is an error reading from the input.
func FindMostActive(cookies io.Reader, targetDate time.Time) ([]string, error) {
	scanner := bufio.NewScanner(cookies)
	activeCookies := make(map[string]int)
	mostActive := 0

	for scanner.Scan() {
		cookieLog := strings.TrimSpace(scanner.Text())

		if cookieLog == "" || cookieLog == fileHeader {
			continue
		}

		cookie, err := Parse(cookieLog, targetDate.Location())
		if err != nil {
			log.Printf("ignoring invalid cookie: `%s`: %s", cookieLog, err)
			continue
		}

		if cookie.Timestamp.Before(targetDate) {
			break
		}

		if sameDay(cookie.Timestamp, targetDate) {
			activeCookies[cookie.Value]++

			if activeCookies[cookie.Value] > mostActive {
				mostActive = activeCookies[cookie.Value]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading cookies: %w", err)
	}

	var mostActiveCookies []string

	for cookie, count := range activeCookies {
		if count == mostActive {
			mostActiveCookies = append(mostActiveCookies, cookie)
		}
	}

	return mostActiveCookies, nil
}

func sameDay(a, b time.Time) bool {
	return a.Year() == b.Year() && a.Month() == b.Month() && a.Day() == b.Day()
}
