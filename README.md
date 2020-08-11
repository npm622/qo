# a baseball qualifying offer tool

this repo provides access to scrape player salary data from an HTML data table and calculate the expected qualifying offer.

## how to run: cli

### the one-liner

please note: this will only working in a linux-like environment with access to cURL

```cmd
curl -o bqo https://raw.githubusercontent.com/npm622/qo/master/bin/bqo && chmod +x ./bqo && ./bqo
```

### from source

to run from source, please make sure you have both git and golang installed.

first, clone the repository and run `go mod download`.  from there, you should be able to run the program via `go run cmd/cli/main.go`.

you can always run `go build -o bqo ./cmd/cli/main.go` to build the executable yourself.

### advanced usage

the tool's help text prints out helpful information with respect to the various arguments and commands.

you can control the verbosity of the output with either of the `-v` or `-f` flags.  the default action is to calculate the qualifying offer one time.

due to the nature of the player salary data containing errors and missing data, one may wish to calculate the average qualifying offer over many calculations.  to do so, simply run:

```cmd
./bqo monte-carlo
```

the same verbosity flags still apply here (for different data points), and you may also specify the number of iterations to use via `-n`.