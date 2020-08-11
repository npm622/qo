package commands

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)

// set of supported command args
const (
	ArgMonteCarloCount = "count"
)

// MonteCarlo is the monte carlo command
func MonteCarlo(c *cli.Context) error {
	start := time.Now()
	args := parseGlobalArgs(c)
	n := c.Int(ArgMonteCarloCount)

	as := make([]analytics, n)
	var avgTime float64
	for i := range as {
		s := time.Now()
		rs, rsErr := getSortedPlayerSalaries(args.url)
		if rsErr != nil {
			return rsErr
		}

		as[i] = analyzeSalaries(rs, false, false)
		avgTime += float64(time.Now().Sub(s).Milliseconds())
	}
	avgTime = avgTime / float64(len(as))

	var r results
	for _, a := range as {
		r.size += float64(a.size)
		r.qualifyingOffer.n++
		r.qualifyingOffer.tot += a.qualifyingOffer
		r.qualifyingOffer.totSq += (a.qualifyingOffer * a.qualifyingOffer)
	}
	r.size = r.size / float64(len(as))
	end := time.Now()

	fmt.Printf("simulations ran:\t%d @ %.3f records per set\n", n, r.size)
	defaultPrinter.Printf("qualifying offer:\t$%.2f", r.qualifyingOffer.avg())
	if args.verbose || args.ferbose {
		defaultPrinter.Printf(" (std dev: $%.3f)\n", r.qualifyingOffer.stdDev(true))
	} else {
		fmt.Println()
	}
	if args.verbose {
		fmt.Println()
		defaultPrinter.Printf("took %.3fs overall, %.3fs per simulation\n", float64(end.Sub(start).Milliseconds())/1000, avgTime/1000)
	}
	return nil
}

type results struct {
	size            float64
	qualifyingOffer result
}

type result struct {
	n     int64
	tot   float64
	totSq float64
}

func (r result) avg() float64 {
	return r.tot / float64(r.n)
}

func (r result) stdDev(finite bool) float64 {
	pop := r.n - 1
	if finite {
		pop = r.n
	}

	return (r.totSq - (r.tot*r.tot)/float64(r.n)) / float64(pop)
}
