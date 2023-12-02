## Session for aocd package

Session token needed for the package to work: `export AOC_SESSION=ABCD1324`, get that from browser cookies. (More info: https://github.com/wimglenn/advent-of-code-wim/issues/1)

## Solving with Python

The package also introspects the current day from the path/name of the script that contains the solution, so be careful with namings: `10_1_asd.py` is not good, use `10_a_asd.py` instead.

## Solving with Go

You can setup an initial workspace for Go for the current day by running `./setupgo.sh` and then navigate into the create dir. Be sure that the file has the right permissions: `chmod 777 setupgo.sh`.
