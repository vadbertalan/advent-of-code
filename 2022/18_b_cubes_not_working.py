# (You guessed 4044.) too high

from typing import Dict, List, Tuple
from aocd import data as aocdata, lines, numbers, submit
from collections import Counter
import math
import sys
import timeit
start_time = timeit.default_timer()

sys.setrecursionlimit(1500)

if len(sys.argv) > 2 or (len(sys.argv) == 2 and '-' not in sys.argv[1]):
    print(f'Usage: python {sys.argv[0]} [-vls] where:')
    print('\tv - verbose')
    print('\tl - use live data')
    print('\ts - submit at the end')
    exit(1)

printing_enabled = False
with_test_data = True
should_submit = False
testdata1 = False
testdata2 = False
testdata3 = False
if len(sys.argv) == 2:
    printing_enabled = 'v' in sys.argv[1]  # verbose
    with_test_data = 'l' not in sys.argv[1]  # live data
    should_submit = 's' in sys.argv[1]  # submit at the end
    testdata1 = '1' in sys.argv[1]
    testdata2 = '2' in sys.argv[1]
    testdata3 = '3' in sys.argv[1]

# answer 58
data_str_test1 = """2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5"""

# answer 42
data_str_test2 = """1,2,2 
4,2,2
2,1,2
3,1,2
2,3,2
3,3,2
2,2,3
3,2,3
2,2,1
3,2,1"""

# answer 126
data_str_test3 = """2,1,1
3,1,1
4,1,1
2,1,2
3,1,2
4,1,2
2,1,3
3,1,3
4,1,3
1,2,1
1,3,1
1,4,1
1,2,2
1,3,2
1,4,2
1,2,3
1,3,3
1,4,3
5,2,1
5,3,1
5,4,1
5,2,2
5,3,2
5,4,2
5,2,3
5,3,3
5,4,3
2,5,1
3,5,1
4,5,1
2,5,2
3,5,2
4,5,2
2,5,3
3,5,3
4,5,3
2,2,0
3,2,0
4,2,0
2,3,0
3,3,0
4,3,0
2,4,0
3,4,0
4,4,0
2,2,4
3,2,4
4,2,4
2,3,4
3,3,4
4,3,4
2,4,4
3,4,4
4,4,4"""

test_data = data_str_test3
if testdata1:
    test_data = data_str_test1
elif testdata2:
    test_data = data_str_test2
data_str = test_data if with_test_data else aocdata


class Coord3(object):
    def __init__(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z

    def __str__(self):
        return f'({self.x}, {self.y}, {self.z})'

    def __repr__(self):
        return f'({self.x}, {self.y}, {self.z})'

    def get_tuple(self):
        return self.x, self.y, self.z

    def has_common_side(self, other):
        pairs = list(zip(self.get_tuple(), other.get_tuple()))
        if len(list(filter(lambda pair: pair[0] == pair[1], pairs))) != 2:
            return False

        not_equal_coord = list(filter(
            lambda pair: pair[0] != pair[1], pairs))[0]  # has 1 len for sure
        return abs(not_equal_coord[0] - not_equal_coord[1]) == 1


def get_coord_of_tuple(tuple3):
    return Coord3(tuple3[0], tuple3[1], tuple3[2])


cubes = []
cubes_coord_map = {}
for line in data_str.split('\n'):
    if line == '':
        continue
    coord3 = Coord3(*map(int, line.split(',')))
    cubes.append(coord3)
    cubes_coord_map[coord3.get_tuple()] = True


INFINITY = 99999999999
min_x = min_y = min_z = INFINITY
max_x = max_y = max_z = -INFINITY

for cube_tuple in cubes:
    if cube_tuple.x < min_x:
        min_x = cube_tuple.x
    if cube_tuple.x > max_x:
        max_x = cube_tuple.x

    if cube_tuple.y < min_y:
        min_y = cube_tuple.y
    if cube_tuple.y > max_y:
        max_y = cube_tuple.y

    if cube_tuple.z < min_z:
        min_z = cube_tuple.z
    if cube_tuple.z > max_z:
        max_z = cube_tuple.z

if printing_enabled:
    print('number of cubes:', len(cubes))
    print(min_x, min_y, min_z)
    print(max_x, max_y, max_z)


# Solution:
wall_coords_map = {}

# Idea: build a cage around the droplet

# left wall
for y in range(min_y, max_y + 1):
    for z in range(min_z, max_z + 1):
        coord = Coord3(min_x - 1, y, z)
        wall_coords_map[coord.get_tuple()] = True

# right wall
for y in range(min_y, max_y + 1):
    for z in range(min_z, max_z + 1):
        coord = Coord3(max_x + 1, y, z)
        wall_coords_map[coord.get_tuple()] = True

# front wall
for x in range(min_x, max_x + 1):
    for z in range(min_z, max_z + 1):
        coord = Coord3(x, max_y + 1, z)
        wall_coords_map[coord.get_tuple()] = True

# back wall
for x in range(min_x, max_x + 1):
    for z in range(min_z, max_z + 1):
        coord = Coord3(x, min_y - 1, z)
        wall_coords_map[coord.get_tuple()] = True

# up wall
for x in range(min_x, max_x + 1):
    for y in range(min_y, max_y + 1):
        coord = Coord3(x, y, max_z + 1)
        wall_coords_map[coord.get_tuple()] = True

# down wall
for x in range(min_x, max_x + 1):
    for y in range(min_y, max_y + 1):
        coord = Coord3(x, y, min_z - 1)
        wall_coords_map[coord.get_tuple()] = True

surrounded_cube_map = set()
visited_coords_map = set()


def can_reach_wall_of_cage(coord: Coord3):
    global visited_coords_map, surrounded_cube_map

    visited_coords_map.add(coord.get_tuple())
    print(coord)

    if coord.get_tuple() in surrounded_cube_map:
        return False

    surrounding_coords = [
        Coord3(coord.x, coord.y, coord.z + 1),  # up
        Coord3(coord.x + 1, coord.y, coord.z),  # right
        Coord3(coord.x, coord.y, coord.z - 1),  # down
        Coord3(coord.x - 1, coord.y, coord.z),  # left
        Coord3(coord.x, coord.y - 1, coord.z),  # behind
        Coord3(coord.x, coord.y + 1, coord.z),  # front
    ]

    # filter destination coords that are cubes or already visited
    eligible_surrounding_coords = list(filter(lambda coord: coord.get_tuple(
    ) not in cubes_coord_map and coord.get_tuple() not in visited_coords_map, surrounding_coords))

    # if any of the destination coords are walls return True
    if any(map(lambda coord: coord.get_tuple() in wall_coords_map, eligible_surrounding_coords)):
        return True

    # if any of destination coords are surrounded for sure, then this coord is surrounded too
    if any(map(lambda coord: coord.get_tuple() in surrounded_cube_map, surrounding_coords)):
        # surrounded_cube_map.add(coord.get_tuple())
        if coord.get_tuple() == (3, 9, 4):
            print(coord)
        print(coord)
        return False

    # each valid surrounding coord will be traversed
    for surrounding_coord in eligible_surrounding_coords:
        if surrounding_coord.get_tuple() not in visited_coords_map:
            if can_reach_wall_of_cage(surrounding_coord):
                return True

    return False


def is_coord_surrounded(coord: Coord3):
    global surrounded_cube_map
    # visited_coords_map = {}
    is_surrounded = not can_reach_wall_of_cage(coord)
    if is_surrounded:
        surrounded_cube_map.add(coord.get_tuple())
    return is_surrounded


def get_common_side_count(coords: List[Coord3]) -> int:
    count = 0
    for i in range(len(coords) - 1):
        for j in range(i + 1, len(coords)):
            if coords[i].has_common_side(coords[j]):
                count += 1
    return count


# count open sides and subtract common sides
open_sides = len(cubes) * 6 - 2 * get_common_side_count(cubes)
if printing_enabled:
    print('open sides without counting inner air:', open_sides)


# see each air cube, whether it's a surrounded cube
surrounded_cube_count = 0
for x in range(min_x, max_x + 1):
    for y in range(min_y, max_y + 1):
        for z in range(min_z, max_z + 1):
            coord = Coord3(x, y, z)
            if coord.get_tuple() == (1, 7, 8):
                print('asd')
            if coord.get_tuple() not in cubes_coord_map and is_coord_surrounded(coord):
                # if printing_enabled:
                #     print(f'found surrounded cube: {coord}')
                surrounded_cube_count += 1

if printing_enabled:
    print('total count of surrounded cubes:', surrounded_cube_count)
    print('surrounded cubes:', surrounded_cube_map)

for cube_tuple in surrounded_cube_map:
    if cube_tuple in cubes_coord_map:
        print(cube_tuple, 'is a bad boy')

inner_air_common_sides = get_common_side_count(
    list(map(get_coord_of_tuple, list(surrounded_cube_map))))
print(f'inner_air_common_sides:', inner_air_common_sides)
solution = open_sides - (6 * surrounded_cube_count -
                         2 * inner_air_common_sides)

print(solution)

if should_submit:
    response = input(f'Are you sure you want to submit {solution}? [y/n] -> ')
    if 'y' in response:
        print(f'✨ Submitting solution: {solution}')
        submit(solution)
    else:
        print(f'❌ Not submitting solution: {solution}')

stop = timeit.default_timer()
print(f'\n✨Time: {stop - start_time} s')
