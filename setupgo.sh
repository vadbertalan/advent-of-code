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
