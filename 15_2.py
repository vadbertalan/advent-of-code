# --- Day 15: Chiton ---
# You've almost reached the exit of the cave, but the walls are getting closer together. Your submarine can barely still fit, though; the main problem is that the walls of the cave are covered in chitons, and it would be best not to bump any of them.

# The cavern is large, but has a very low ceiling, restricting your motion to two dimensions. The shape of the cavern resembles a square; a quick scan of chiton density produces a map of risk level throughout the cave (your puzzle input). For example:

# 1163751742
# 1381373672
# 2136511328
# 3694931569
# 7463417111
# 1319128137
# 1359912421
# 3125421639
# 1293138521
# 2311944581
# You start in the top left position, your destination is the bottom right position, and you cannot move diagonally. The number at each position is its risk level; to determine the total risk of an entire path, add up the risk levels of each position you enter (that is, don't count the risk level of your starting position unless you enter it; leaving it adds no risk to your total).

# Your goal is to find a path with the lowest total risk. In this example, a path with the lowest total risk is highlighted here:

# 1163751742
# 1381373672
# 2136511328
# 3694931569
# 7463417111
# 1319128137
# 1359912421
# 3125421639
# 1293138521
# 2311944581
# The total risk of this path is 40 (the starting position is never entered, so its risk is not counted).

# What is the lowest total risk of any path from the top left to the bottom right?

# Your puzzle answer was 487.

import math
import pprint
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

# data_str1 = open('15_input.txt', 'r').read()
lines = data_str1.split('\n')

data = []
N = len(lines) * 5
M = len(lines[0]) * 5

print(N, M)

def euclidean_distance(point1, point2):
    return math.sqrt((point1[0] - point2[0]) ** 2 + (point1[1] - point2[1]) ** 2)

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


MAX = N * M * 9 + 1
seen = [[False for y in range(M)] for x in range(N)]
shortest_dist = [[MAX for y in range(M)] for x in range(N)]
parents = [[None for y in range(M)] for x in range(N)]
shortest_dist[0][0] = 0

# euclidean A* weight
weights = [[euclidean_distance((x, y), (N - 1, M - 1)) for y in range(M)] for x in range(N)]


def valid_coord(coord):
    return 0 <= coord[0] < N and 0 <= coord[1] < M

def get_data_on_coord(coord): 
    global extended_data
    return extended_data[coord[0]][coord[1]]

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

current_coord = (0, 0)
while current_coord != (N - 1, M - 1): 
    seen[current_coord[0]][current_coord[1]] = True

    neighbors = {
        'right': [current_coord[0], current_coord[1] + 1],
        'down': [current_coord[0] + 1, current_coord[1]],
        'left': [current_coord[0], current_coord[1] - 1],
        'up': [current_coord[0] - 1, current_coord[1]]
    }

    dir_coords = list(filter(lambda item: valid_coord(item[1]) and not seen[item[1][0]][item[1][1]], neighbors.items()))
    for dc in dir_coords:
        if shortest_dist[dc[1][0]][dc[1][1]] > shortest_dist[current_coord[0]][current_coord[1]] + get_data_on_coord(dc[1]):
            shortest_dist[dc[1][0]][dc[1][1]] = shortest_dist[current_coord[0]][current_coord[1]] + get_data_on_coord(dc[1])
            parents[dc[1][0]][dc[1][1]] = current_coord

    min_coord = get_min_shortest_dist_coord()
    current_coord = min_coord

# pp.pprint(shortest_dist)
# pp.pprint(seen)
# pp.pprint(parents)

route = []
cur = (N - 1, M - 1)
while cur != (0, 0): 
    route.append(cur)
    cur = parents[cur[0]][cur[1]]

route_mat = [['.' for y in range(M)] for x in range(N)]
for coord in route:
    route_mat[coord[0]][coord[1]] = str(get_data_on_coord(coord))

for line in route_mat:
    print(''.join(line))

print(f'shortest dist from 0,0 to {N},{M} -> {shortest_dist[N - 1][M - 1]}')

