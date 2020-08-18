# phillies baseball r&d questionnaire

## analyze python code: is_palindrome

```py
def is_palindrone(s):
    r=""
    for c in s:
        r = c +r
    for x in range(0, len(s)):
        if s[x] == r[x]:
            x = True
        else:
            return False
    return x
```

there are a few areas we could improve here.  for one, allocating an array of booleans for each character check is unnecessary, as it is simply the case that when a check fails, we can short-circuit our function and return false.  if we never short-circuit, we can simply return true knowing all of our checks passed.

there is another unnecessary allocation here in the `r` string.  creating a reverse of the input is actually not needed, one can simply compare the "i-th" index character with its mirror opposite: `len(s)-i-1`.  We know that if all of these "mirror checks" pass, then our string must be a palindrome (regardless of the input having even or odd length).  So in addition to not needing to allocate the `r` string, we can reduce our number of for loops to one and our iterations to at most `Math.floor(len(s)/2)`.

the resulting code would look like:

```py
def is_palindrone(s):
    for x in range(0, int(len(s)/2)):
        if s[x] == s[len(s)-x-1]:
            return False
    return True
```

## a baseball qualifying offer tool

this repo provides access to scrape player salary data from an HTML data table and calculate the expected qualifying offer.

### how to run

#### the one-liner

please note: this will only working in a linux-like environment with access to cURL

```cmd
curl -o bqo https://raw.githubusercontent.com/npm622/qo/master/bin/bqo && chmod +x ./bqo && ./bqo
```

#### from source

to run from source, please make sure you have both git and golang installed.

first, clone the repository and run `go mod download`.  from there, you should be able to run the program via `go run cmd/cli/main.go`.

you can always run `go build -o bqo ./cmd/cli/main.go` to build the executable yourself.

#### advanced usage

the tool's help text prints out helpful information with respect to the various arguments and commands.

you can control the verbosity of the output with either of the `-v` or `-f` flags.  the default action is to calculate the qualifying offer one time.

due to the nature of the player salary data containing errors and missing data, one may wish to calculate the average qualifying offer over many calculations.  to do so, simply run:

```cmd
./bqo monte-carlo
```

the same verbosity flags still apply here (for different data points), and you may also specify the number of iterations to use via `-n`.