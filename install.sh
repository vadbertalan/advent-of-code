brew bundle --no-upgrade

echo "\nSetting up environment variables for aoc_tiles. Make sure to set AOC_SESSION in your shell profile.\n"

# Install pre-commit hooks for aoc_tiles
# https://github.com/LiquidFun/aoc_tiles
pre-commit install --hook-type post-commit

if [ -z "$AOC_SESSION" ]; then
  echo "Error: AOC_SESSION environment variable is not set. Please set it in your shell profile."
else
	mkdir -p .aoc_tiles
	echo $AOC_SESSION > .aoc_tiles/session.cookie
fi