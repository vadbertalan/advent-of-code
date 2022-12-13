# --- Part Two - --

# As you walk up the hill, you suspect that the Elves will want to turn this into a hiking trail. The beginning isn't very scenic, though
# perhaps you can find a better starting point.

# To maximize exercise while hiking, the trail should start as low as possible: elevation a. The goal is still the square marked E. However, the trail should still be direct, taking the fewest steps to reach its goal. So, you'll need to find the shortest path from any square at elevation a to the square marked E.

# Again consider the example from above:

# Sabqponm
# abcryxxl
# accszExk
# acctuvwj
# abdefghi
# Now, there are six choices for starting position(five marked a, plus the square marked S that counts as being at elevation a). If you start at the bottom-left square, you can reach the goal most quickly:

# ...v << <<
# ...vv << ^
# ...v > E ^ ^
# . > v >> > ^ ^
# > ^ >>>> > ^
# This path reaches the goal in only 29 steps, the fewest possible.

# What is the fewest steps required to move starting from any square with elevation a to the location that should get the best signal?

# Your puzzle answer was 500.

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
start_coords = []
end_coord = None
for i in range(N):
    for j in range(M):
        if data[i][j] == 'a' or data[i][j] == 'S':
            start_coord = Coord(i, j)
            data[start_coord.x][start_coord.y] = 'a'
            start_coords.append(start_coord)
            continue
        if data[i][j] == 'E':
            end_coord = Coord(i, j)
            data[end_coord.x][end_coord.y] = 'z'

if start_coords == [] or end_coord is None:
    print('ERROR: did not find start and/or end coordinates')


MAX = N * M + 1  # there are row count * col count squares

seen = []
shortest_distances = []
parents = []


def init():
    global seen, shortest_distances, parents

    seen = [[False for x in range(M)] for y in range(N)]
    shortest_distances = [[MAX for x in range(M)] for y in range(N)]
    parents = [[None for x in range(M)] for y in range(N)]


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
            if not seen[i][j] and shortest_distances[i][j] < min_cost:
                min_coord = Coord(i, j)
                min_cost = shortest_distances[i][j]

    if min_coord is None:
        print('\tPossibly an island found because no min coord found', min_cost, i, j)
        return None
        # print_data(shortest_distances)
        # print_data(seen)
        # raise 'Error in get min shortest dist'

    return min_coord


print(f'Found {len(start_coords)} candidates for starting coords. Hold on!')

minimum_cost = MAX
candidate_id = 1
for start_coord in start_coords:
    print(f'*** {candidate_id} ***  Trying start coordinates {start_coord}')
    candidate_id += 1

    init()

    shortest_distances[start_coord.x][start_coord.y] = 0

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
            # weight is 1 (and not MAX) as it can step there for sure
            new_distance_through_current_coord = shortest_distances[
                current_coord.x][current_coord.y] + 1
            if shortest_distances[new_coord.x][new_coord.y] > new_distance_through_current_coord:
                shortest_distances[new_coord.x][new_coord.y] = new_distance_through_current_coord
                parents[new_coord.x][new_coord.y] = current_coord

        min_coord = get_min_shortest_dist_coord()
        if min_coord is None:
            break

        current_coord = min_coord

    if shortest_distances[end_coord.x][end_coord.y] < minimum_cost:
        print(
            f'\t Updating minimum cost from {minimum_cost} to {shortest_distances[end_coord.x][end_coord.y]}')
        minimum_cost = shortest_distances[end_coord.x][end_coord.y]


# if printing_enabled:
#     print_data(shortest_distances)
#     print_data(seen)

res = minimum_cost
print(res)

if should_submit:
    response = input(f'Are you sure you want to submit {res}? [y/n] -> ')
    if 'y' in response:
        print(f'✨ Submitting solution: {res}')
        submit(res)
    else:
        print(f'❌ Not submitting solution: {res}')


stop = timeit.default_timer()
print(f'\n✨Time: {stop - start_time} s')
