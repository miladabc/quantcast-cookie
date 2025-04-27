package main

import (
	"fmt"
	"log"
	"os"

	"github.com/miladabc/quantcast-cookie/internal/cli"
	"github.com/miladabc/quantcast-cookie/internal/cookie"
)

func main() {
	query, err := cli.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("cli parsing: %v", err)
	}

	file, err := os.Open(query.CookieFilePath)
	if err != nil {
		log.Fatalf("opening cookies file: %v", err)
	}
	defer file.Close()

	mostActiveCookies, err := cookie.FindMostActive(file, query.CookieTimestamp)
	if err != nil {
		log.Fatalf("searching cookies: %v", err)
	}

	if len(mostActiveCookies) == 0 {
		log.Printf("no cookies found for `%s`", query.CookieTimestamp)
	}

	for _, cookie := range mostActiveCookies {
		fmt.Println(cookie)
	}
}
