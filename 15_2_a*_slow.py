# !!!!!   Very slow but works in ~ 1 hour. Second solution resolves under 32s

# --- Part Two - --

# Now that you know how to find low-risk paths in the cave, you can try to find your way out.

# The entire cave is actually five times larger in both dimensions than you thought
# the area you originally scanned is just one tile in a 5x5 tile area that forms the full map. Your original map tile repeats to the right and downward
# each time the tile repeats to the right or downward, all of its risk levels are 1 higher than the tile immediately up or left of it. However, risk levels above 9 wrap back around to 1. So, if your original map had some position with a risk level of 8, then that same position on each of the 25 total tiles would be as follows:

# 8 9 1 2 3
# 9 1 2 3 4
# 1 2 3 4 5
# 2 3 4 5 6
# 3 4 5 6 7
# Each single digit above corresponds to the example position with a value of 8 on the top-left tile. Because the full map is actually five times larger in both dimensions, that position appears a total of 25 times, once in each duplicated tile, with the values shown above.

# Here is the full five-times-as-large version of the first example above, with the original map in the top left corner highlighted:

# The total risk of this path is 315 (the starting position is still never entered, so its risk is not counted).

# Using the full map, what is the lowest total risk of any path from the top left to the bottom right?

# Your puzzle answer was 2821.


import pprint
import math
import timeit
start_time = timeit.default_timer()

pp = pprint.PrettyPrinter(indent=4)


data_str1 = """1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581"""

data_str1 = open('15_input.txt', 'r').read()

lines = data_str1.split('\n')

data = []
N = len(lines) * 5
M = len(lines[0]) * 5

print(N, M)


def print_data(mat):
    for row in mat:
        for num in row:
            print(f'{num: <4}', end='')
        print()
    print()
    print()


def euclidean_distance(point1, point2):
    return math.sqrt((point1[0] - point2[0]) ** 2 + (point1[1] - point2[1]) ** 2)


def manhattan_distance(point1, point2):
    return abs(point1[0] - point2[0]) + abs(point1[1] - point2[1])


def shift(subrow, i):
    return ''.join(list(map(lambda x: str((x + i) if x + i <= 9 else ((x + i) % 10) + 1), subrow)))


def gen_row(subrow, count=5):
    row = ''
    for i in range(count):
        row += shift(subrow, i)
    return row


for line in lines:
    subrow = list(map(int, line))
    data.append(gen_row(subrow, 5))

for j in range(1, 5):
    for line in lines:
        subrow = list(map(int, line))
        data.append(gen_row(list(map(int, shift(subrow, j))), 5))

extended_data = []
for line in data:
    extended_data.append(list(map(int, line)))

weights = [[manhattan_distance((x, y), (N - 1, M - 1))
            for y in range(M)] for x in range(N)]

MAX = N * M * 9 + 1 + weights[0][0] * 2
print(weights[0][0], MAX)
seen = [[False for y in range(M)] for x in range(N)]
shortest_dist = [[MAX for y in range(M)] for x in range(N)]
shortest_dist_weighted = [[MAX for y in range(M)] for x in range(N)]
parents = [[None for y in range(M)] for x in range(N)]
shortest_dist[0][0] = 0


def valid_coord(coord):
    return 0 <= coord[0] < N and 0 <= coord[1] < M


def get_cost_on_coord(coord):
    global extended_data
    return extended_data[coord[0]][coord[1]]


queue = set()


def get_cur_shortest_dist_coord_from_queue():
    global queue

    min_coords = min(
        queue, key=lambda coord: shortest_dist_weighted[coord[0]][coord[1]])
    queue.remove(min_coords)

    # for i in range(N):
    #     for j in range(M):
    #         if not seen[i][j] and shortest_dist_weighted[i][j] < min_item:
    #             min_coords = (i, j)
    #             min_item = shortest_dist_weighted[i][j]

    # if min_coords == None:
    #     raise 'Error in get min shortest dist'
    return min_coords


current_coord = (0, 0)
while current_coord != (N - 1, M - 1):
    seen[current_coord[0]][current_coord[1]] = True

    neighbors = {
        'right': [current_coord[0], current_coord[1] + 1],
        'down': [current_coord[0] + 1, current_coord[1]],
        'left': [current_coord[0], current_coord[1] - 1],
        'up': [current_coord[0] - 1, current_coord[1]]
    }

    # it would be enough to pass through only the coord
    dir_coords = list(filter(lambda item: valid_coord(
        item[1]) and not seen[item[1][0]][item[1][1]], neighbors.items()))
    for dc in dir_coords:
        if shortest_dist[dc[1][0]][dc[1][1]] > shortest_dist[current_coord[0]][current_coord[1]] + get_cost_on_coord(dc[1]):
            queue.add(tuple(dc[1]))
            shortest_dist[dc[1][0]][dc[1][1]] = shortest_dist[current_coord[0]
                                                              ][current_coord[1]] + get_cost_on_coord(dc[1])
            shortest_dist_weighted[dc[1][0]][dc[1][1]] = shortest_dist[current_coord[0]
                                                                       ][current_coord[1]] + get_cost_on_coord(dc[1]) + weights[dc[1][0]][dc[1][1]]
            parents[dc[1][0]][dc[1][1]] = current_coord

    current_coord = get_cur_shortest_dist_coord_from_queue()

for i in shortest_dist:
    for j in i:
        print(f'{round(j, 1): <3}', end=' ')
    print()
print()
for i in seen:
    print(list(map(int, i)))
print()
# pp.pprint(parents)

route = []
cost = 0
cur = (N - 1, M - 1)
while cur != (0, 0):
    route.append(cur)
    cost += get_cost_on_coord(cur)
    cur = parents[cur[0]][cur[1]]

route_mat = [['.' for y in range(M)] for x in range(N)]
for coord in route:
    route_mat[coord[0]][coord[1]] = str(get_cost_on_coord(coord))

for line in route_mat:
    print(''.join(line))

print(
    f'shortest dist from 0,0 to {N},{M} -> {cost}  |  {shortest_dist[N - 1][M - 1]}')

c = 0
for i in range(N):
    for j in range(M):
        c += 1 if seen[i][j] else 0
print(f'Total vertices: {N * M}')
print(f'Seen vertices: {c}')
print(f'Unseen vertices: {N * M - c}')

stop = timeit.default_timer()
print(f'\nTime: {stop - start_time} s')
