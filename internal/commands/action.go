package commands

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

// Action is the action
func Action(c *cli.Context) error {
	args := parseGlobalArgs(c)

	rs, rsErr := getSortedPlayerSalaries(args.url)
	if rsErr != nil {
		return rsErr
	}

	a := analyzeSalaries(rs, args.verbose, args.ferbose)

	fmt.Printf("player salary records found:\t%d\n", len(rs))
	defaultPrinter.Printf("the qualifying offer is:\t$%.2f\n", a.qualifyingOffer)
	if args.verbose || args.ferbose {
		fmt.Println("---")
		defaultPrinter.Printf("net avg:\t$%.2f\n", a.netAvg)
		defaultPrinter.Printf("gross avg:\t$%.2f\n", a.grossAvg)

		if args.verbose {
			fmt.Println("---")
			fmt.Printf("%d players with 'no salary data': %s\n",
				len(a.noSalary),
				strings.Join(a.noSalary, ", "))
			fmt.Println("---")
			fmt.Printf("%d players with empty/zero salary: %s\n",
				len(a.zeroSalary),
				strings.Join(a.zeroSalary, ", "))
		}
	}
	return nil
}
