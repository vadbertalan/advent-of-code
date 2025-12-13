# This script sets up the environment for Advent of Code challenges.
# It accepts the following command-line arguments:
#   -y, --year <YEAR> : Specifies the year of the challenge.
#   -d, --day <DAY>   : Specifies the day of the challenge.
#   -h, --help        : Displays usage information.
#
# Example usage:
#   ./solvego.sh -y 2023 -d 1
#
# The script processes the provided arguments and passes them to the Go program `setup.go`.
# If no arguments are provided, the script will end up calling `go run setup.go` without any arguments,
# which will set up the environment for the current year and day.

while [[ $# -gt 0 ]]; do
  case $1 in
    -y|--year)
      YEAR="$2"
      shift # past argument
      shift # past value
      ;;
    -d|--day)
      DAY="$2"
      shift # past argument
      shift # past value
      ;;
    -h|--help)
      echo "Usage: $0 [-y|--year <YEAR>] [-d|--day <DAY>]"
	  exit 0
      ;;
    -*|--*)
      echo "Unknown option $1"
      exit 1
      ;;
  esac
done

if [ ! -z $DAY ]
then
	DAY_ARG="-d $DAY"
fi

if [ ! -z $YEAR ]
then
	YEAR_ARG="-y $YEAR"
fi

go run setup.go $YEAR_ARG $DAY_ARG
