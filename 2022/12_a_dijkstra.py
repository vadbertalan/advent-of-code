from aocd import data as aocdata, lines, numbers, submit
from collections import Counter
import math
import sys
import timeit
from utils import Coord
start_time = timeit.default_timer()


printing_enabled = sys.argv[1] == '-v' if len(sys.argv) == 2 else False

data_str1 = """Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi"""

data_str = aocdata

lines = data_str1.split('\n')

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
    data.append(line)


def print_data():
    for row in data:
        for num in row:
            print(f'{num:<3}', end='')
        print()
    print()


if printing_enabled:
    print_data()
    print()


MAX = N * M * 9 + 1
seen = [[False for x in range(M)] for y in range(N)]
shortest_dist = [[MAX for x in range(M)] for y in range(N)]
parents = [[None for x in range(M)] for y in range(N)]
shortest_dist[0][0] = 0


def valid_coord(coord):
    return 0 <= coord.x < N and 0 <= coord.y < M


def get_data_on_coord(coord):
    global data
    return data[coord.x][coord.y]


def get_min_shortest_dist_coord():
    min_item = MAX
    min_coords = None

    for i in range(N):
        for j in range(M):
            if not seen[i][j] and shortest_dist[i][j] < min_item:
                min_coords = (i, j)
                min_item = shortest_dist[i][j]

    if min_coords == None:
        raise 'Error in get min shortest dist'
    return min_coords

# get start and end coordinates
start_coord = None
end_coord = None
for i in range(N):
    for j in range(M):
        if data[i][j] == 'S':
            start_coord = Coord(i, j)
        elif data[i][j] == 'E':
            end_coord = Coord(i, j)
    if start_coord != None and end_coord != None:
        break
if start_coord == None or end_coord == None:
    print('ERROR: did not find start and end coordinates')


current_coord = start_coord
while current_coord != end_coord:
    seen[current_coord.x][current_coord.y] = True

    neighbors = {
        'up': [current_coord.x - 1, current_coord.y],
        'right': [current_coord.x, current_coord.y + 1],
        'down': [current_coord.x + 1, current_coord.y],
        'left': [current_coord.x, current_coord.y - 1],
    }

    dir_coords = list(filter(lambda item: valid_coord(
        item[1]) and not seen[item[1].x][item[1].y], neighbors.items()))
    for _, new_coord in dir_coords:
        if shortest_dist[new_coord.x][new_coord.y] > shortest_dist[current_coord.x][current_coord.y] + get_data_on_coord(new_coord[1]):
            shortest_dist[new_coord.x][new_coord.y] = shortest_dist[current_coord.x
                                                              ][current_coord.y] + get_data_on_coord(new_coord.y)
            parents[new_coord.x][new_coord.y] = current_coord

    min_coord = get_min_shortest_dist_coord()
    current_coord = min_coord

print(res)

stop = timeit.default_timer()
print(f'\n✨Time: {stop - start_time} s')
