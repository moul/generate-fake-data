package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"

	gofakeit "github.com/brianvoe/gofakeit/v5"
	"github.com/peterbourgon/ff/v3/ffcli"
	"moul.io/srand"
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	// init dictionaries
	var (
		availableDicts []string
		dictMap        map[string]func() string
	)
	{
		dictMap = map[string]func() string{
			"quote":       gofakeit.Quote,
			"phrase":      gofakeit.Phrase,
			"beer":        gofakeit.BeerName,
			"uuid":        gofakeit.UUID,
			"question":    gofakeit.Question,
			"hacker":      gofakeit.HackerPhrase,
			"hipster":     func() string { return gofakeit.HipsterSentence(8) },
			"lorem-ipsum": func() string { return gofakeit.LoremIpsumSentence(8) },
			"apache-log": func() string {
				return fmt.Sprintf(
					`%s - - [%s] %d "%s %s/%s HTTP/1.1" %d 17 "-" "%q"`,
					gofakeit.IPv4Address(),
					gofakeit.Date(),
					rand.Intn(99999999),
					gofakeit.HTTPMethod(),
					gofakeit.LoremIpsumWord(),
					gofakeit.LoremIpsumWord(),
					gofakeit.HTTPStatusCode(),
					gofakeit.UserAgent(),
				)
			},
			"address": func() string {
				address := gofakeit.Address()
				return fmt.Sprintf("%s, %s", address.Address, address.Country)
			},
			"random": func() string {
				for {
					choice := availableDicts[rand.Intn(len(availableDicts))]
					if choice == "random" {
						continue
					}
					return dictMap[choice]()
				}
			},
		}
		availableDicts = make([]string, len(dictMap))
		i := 0
		for dict := range dictMap {
			availableDicts[i] = dict
			i++
		}
		sort.Strings(availableDicts)
	}

	// CLI configuration
	var (
		fs       = flag.NewFlagSet("generate-fake-data", flag.ExitOnError)
		seed     = fs.Int64("seed", 0, "seed random")
		lines    = fs.Int("lines", 42, "amount of lines to generate")
		sleepMin = fs.Duration("sleep-min", 0, "minimum sleep duration between two lines")
		sleepMax = fs.Duration("sleep-max", 300*time.Millisecond, "maximum sleep duration between two lines")
		dict     = fs.String("dict", "random", "available: "+strings.Join(availableDicts, ", "))
		noStdout = fs.Bool("no-stdout", false, "disable writing to stdout (incompatible with no-stderr)")
		noStderr = fs.Bool("no-stderr", false, "disable writing to stderr (incompatible with no-stdout)")
	)

	root := &ffcli.Command{
		Name:    "generate-fake-data",
		FlagSet: fs,
		Exec: func(ctx context.Context, args []string) error {
			if _, ok := dictMap[*dict]; !ok {
				return flag.ErrHelp
			}
			if *noStdout && *noStderr {
				return flag.ErrHelp
			}

			if *seed != 0 {
				rand.Seed(*seed)
			} else {
				rand.Seed(srand.Fast())
			}

			for i := 0; i < *lines; i++ {
				line := dictMap[*dict]()
				switch {
				case *noStdout:
					fmt.Fprintln(os.Stderr, line)
				case *noStderr:
					fmt.Fprintln(os.Stdout, line)
				default:
					if rand.Intn(2) == 0 {
						fmt.Fprintln(os.Stderr, line)
					} else {
						fmt.Fprintln(os.Stdout, line)
					}
				}

				if *sleepMin >= 0 && *sleepMax > *sleepMin {
					sleepDuration := float64(*sleepMin) + rand.Float64()*(float64(*sleepMax)-float64(*sleepMin))
					time.Sleep(time.Duration(sleepDuration))
				}
			}
			return nil
		},
	}
	return root.ParseAndRun(context.Background(), args[1:])
}
