# --- Day 12: Hill Climbing Algorithm - --

# You try contacting the Elves using your handheld device, but the river you're following must be too low to get a decent signal.

# You ask the device for a heightmap of the surrounding area(your puzzle input). The heightmap shows the local area from above broken into a grid
# the elevation of each square of the grid is given by a single lowercase letter, where a is the lowest elevation, b is the next-lowest, and so on up to the highest elevation, z.

# Also included on the heightmap are marks for your current position(S) and the location that should get the best signal(E). Your current position(S) has elevation a, and the location that should get the best signal(E) has elevation z.

# You'd like to reach E, but to save energy, you should do it in as few steps as possible. During each step, you can move exactly one square up, down, left, or right. To avoid needing to get out your climbing gear, the elevation of the destination square can be at most one higher than the elevation of your current square
# that is , if your current elevation is m, you could step to elevation n, but not to elevation o. (This also means that the elevation of the destination square can be much lower than the elevation of your current square.)

# For example:

# Sabqponm
# abcryxxl
# accszExk
# acctuvwj
# abdefghi
# Here, you start in the top-left corner
# your goal is near the middle. You could start by moving down or right, but eventually you'll need to head toward the e at the bottom. From there, you can spiral around to the goal:

# v..v << <<
# >v.vv << ^
# . > vv > E ^ ^
# ..v >> > ^ ^
# ..>>>> > ^
# In the above diagram, the symbols indicate whether the path exits each square moving up (^), down (v), left ( < ), or right ( > ). The location that should get the best signal is still E, and . marks unvisited squares.

# This path reaches the goal in 31 steps, the fewest possible.

# What is the fewest steps required to move from your current position to the location that should get the best signal?

# Your puzzle answer was 504.

from aocd import data as aocdata, lines, numbers, submit
from collections import Counter
import math
import sys
import timeit
start_time = timeit.default_timer()

if len(sys.argv) > 2 or (len(sys.argv) == 2 and '-' not in sys.argv[1]):
    print(f'Usage: python {sys.argv[0]} [-vls] where:')
    print('\tv - verbose')
    print('\tl - use live data')
    print('\ts - submit at the end')
    exit(1)

printing_enabled = False
with_test_data = True
should_submit = False
if len(sys.argv) == 2:
    printing_enabled = 'v' in sys.argv[1]  # verbose
    with_test_data = 'l' not in sys.argv[1]  # live data
    should_submit = 's' in sys.argv[1]  # submit at the end


class Coord():
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __str__(self):
        return f'({self.x}, {self.y})'

    def __repr__(self):
        return f'({self.x}, {self.y})'

    def get_tuple(self):
        return self.x, self.y

    def __eq__(self, other):
        return self.x == other.x and self.y == other.y


data_str_test = """Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi"""

data_str = data_str_test if with_test_data else aocdata

lines = data_str.split('\n')

res = 0

data = []
N = len(lines)
M = len(lines[0])

if printing_enabled:
    print(N, M)
    for line in lines:
        print(line)
    print()

for line in lines:
    data.append(list(map(lambda char: chr(ord(char)), line)))


def print_data(m=data):
    for row in m:
        for num in row:
            print(f'{num:<5}', end='')
        print()
    print()


if printing_enabled:
    print_data()
    print()

# get start and end coordinates
start_coord = None
end_coord = None
for i in range(N):
    for j in range(M):
        if data[i][j] == 'S':
            start_coord = Coord(i, j)
            continue
        if data[i][j] == 'E':
            end_coord = Coord(i, j)
    if type(start_coord) == Coord and type(end_coord) == Coord:
        break
if start_coord is None or end_coord is None:
    print('ERROR: did not find start and end coordinates')

data[start_coord.x][start_coord.y] = 'a'
data[end_coord.x][end_coord.y] = 'z'

if printing_enabled:
    print(start_coord, end_coord)

MAX = N * M + 1  # there are row count * col count squares
seen = [[False for x in range(M)] for y in range(N)]
shortest_dist = [[MAX for x in range(M)] for y in range(N)]
parents = [[None for x in range(M)] for y in range(N)]
shortest_dist[start_coord.x][start_coord.y] = 0


def valid_coord(coord: Coord) -> bool:
    return 0 <= coord.x < N and 0 <= coord.y < M


# def get_weight_of_coord(from_coord: Coord, new_coord: Coord) -> int:
#     return 1 if data[new_coord.x][new_coord.y] == data[from_coord.x][from_coord.y] or data[new_coord.x][new_coord.y] == chr(ord(data[from_coord.x][from_coord.y]) + 1) else MAX


def can_step_there_from_here(from_coord: Coord, new_coord: Coord) -> bool:
    return data[new_coord.x][new_coord.y] <= chr(ord(data[from_coord.x][from_coord.y]) + 1)

def get_min_shortest_dist_coord() -> Coord:
    min_cost = MAX
    min_coord = None

    for i in range(N):
        for j in range(M):
            if not seen[i][j] and shortest_dist[i][j] < min_cost:
                min_coord = Coord(i, j)
                min_cost = shortest_dist[i][j]

    if min_coord is None:
        print('ERROR: did not find min coord', min_cost, i, j)
        print_data(shortest_dist)
        print_data(seen)
        raise 'Error in get min shortest dist'

    return min_coord


current_coord = start_coord
while current_coord != end_coord:
    seen[current_coord.x][current_coord.y] = True

    neighbors = [
        Coord(current_coord.x - 1, current_coord.y),  # up
        Coord(current_coord.x, current_coord.y + 1),  # right
        Coord(current_coord.x + 1, current_coord.y),  # down
        Coord(current_coord.x, current_coord.y - 1),  # left
    ]

    dir_coords = list(filter(lambda coord: valid_coord(
        coord) and not seen[coord.x][coord.y] and can_step_there_from_here(current_coord, coord), neighbors))

    for new_coord in dir_coords:
        new_distance_through_current_coord = shortest_dist[current_coord.x][current_coord.y] + 1  # weight is 1 (and not MAX) as it can step there for sure
        if shortest_dist[new_coord.x][new_coord.y] > new_distance_through_current_coord:
            shortest_dist[new_coord.x][new_coord.y] = new_distance_through_current_coord
            parents[new_coord.x][new_coord.y] = current_coord

    min_coord = get_min_shortest_dist_coord()
    current_coord = min_coord


if printing_enabled:
    print_data(shortest_dist)
    print_data(seen)

res = shortest_dist[end_coord.x][end_coord.y]
print(res)

if should_submit:
    print(f'submitting solution: {res}')
    submit(res)

stop = timeit.default_timer()
print(f'\n✨Time: {stop - start_time} s')
