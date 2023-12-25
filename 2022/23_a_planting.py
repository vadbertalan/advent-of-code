# https://adventofcode.com/2022/day/23

# Your puzzle answer was 3947.

from aocd import data as aocdata, lines, numbers, submit
from collections import Counter
import math
import sys
import timeit
start_time = timeit.default_timer()

ID = 0

NORTH = 'north'
EAST = 'east'
SOUTH = 'south'
WEST = 'west'


class Coord():
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __str__(self):
        return f'({self.x}, {self.y})'

    def __repr__(self):
        return f'({self.x}, {self.y})'

    def __eq__(self, other):
        return self.x == other.x and self.y == other.y

    def get_tuple(self):
        return self.x, self.y

    def __hash__(self):
        global ID
        id = ID
        ID += 1
        return id


class Elf():
    def __init__(self, coord):
        self.coord = coord

    def __str__(self):
        return self.coord.__str__()

    def __repr__(self):
        return self.coord.__repr__()


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

data_str_test = """..............
..............
.......#......
.....###.#....
...#...#.#....
....#...##....
...#.###......
...##.#.##....
....#..#......
..............
..............
.............."""

data_str = data_str_test if with_test_data else aocdata

lines = data_str.split('\n')

data = []
N = len(lines)
M = len(lines[0])

if printing_enabled:
    print(N, M)
    for line in lines:
        print(line)
    print()

elves = set()

for row, line in enumerate(lines):
    data.append([])
    for col, ch in enumerate(line):
        data[row].append(ch)
        if ch == '#':
            elves.add(Elf(Coord(row, col)))


def print_data():
    print('      ', end='')
    for i in range(len(data[0])):
        print(f'{i:<3}', end='')
    print()

    print('      ', end='')
    for _ in range(len(data[0] * 3)):
        print('-', end='')
    print()

    row_count = 0
    for row in data:
        print(f'{row_count:<3}', end=' | ')

        for value in row:
            print(f'{value:<3}', end='')

        print(f' | {row_count}')

        row_count += 1

    print('      ', end='')
    for _ in range(len(data[0] * 3)):
        print('-', end='')
    print()

    print('      ', end='')
    for i in range(len(data[0])):
        print(f'{i:<3}', end='')

    print()


if printing_enabled:
    print_data()
    print(elves)
    print('elf count:', len(elves))
    print()


# Solution:

# def get_data_on_coord(coord):
#     global data
#     return data[coord[0]][coord[1]]

def has_any_elf_coord(coord):
    for elf in elves:
        if coord == elf.coord:
            return True
    return False

def elf_has_neighbors(elf):
    neighbors = [
        Coord(elf.coord.x - 1, elf.coord.y),  # up
        Coord(elf.coord.x - 1, elf.coord.y + 1),  # up - right
        Coord(elf.coord.x, elf.coord.y + 1),  # right
        Coord(elf.coord.x + 1, elf.coord.y + 1),  # right - down
        Coord(elf.coord.x + 1, elf.coord.y),  # down
        Coord(elf.coord.x + 1, elf.coord.y - 1),  # down - left
        Coord(elf.coord.x, elf.coord.y - 1),  # left
        Coord(elf.coord.x - 1, elf.coord.y - 1),  # left - up
    ]

    for neighbor_coord in neighbors:
        if has_any_elf_coord(neighbor_coord):
            return True
    return False


def elf_has_neighbors_in_dir(elf, dir):
    if dir == NORTH:
        neighbors = [
            Coord(elf.coord.x - 1, elf.coord.y - 1),  # left - up
            Coord(elf.coord.x - 1, elf.coord.y),  # up
            Coord(elf.coord.x - 1, elf.coord.y + 1)  # up - right
        ]
    elif dir == EAST:
        neighbors = [
            Coord(elf.coord.x - 1, elf.coord.y + 1),  # up - right
            Coord(elf.coord.x, elf.coord.y + 1),  # right
            Coord(elf.coord.x + 1, elf.coord.y + 1),  # right - down
        ]
    elif dir == SOUTH:
        neighbors = [
            Coord(elf.coord.x + 1, elf.coord.y + 1),  # right - down
            Coord(elf.coord.x + 1, elf.coord.y),  # down
            Coord(elf.coord.x + 1, elf.coord.y - 1),  # down - left
        ]
    elif dir == WEST:
        neighbors = [
            Coord(elf.coord.x + 1, elf.coord.y - 1),  # down - left
            Coord(elf.coord.x, elf.coord.y - 1),  # left
            Coord(elf.coord.x - 1, elf.coord.y - 1),  # left - up
        ]
    else:
        print('ERROR: Unknown direction', dir)
    
    for neighbor_coord in neighbors:
        if has_any_elf_coord(neighbor_coord):
            return True
    return False

def get_coords_of_elf_in_dir(elf, dir) -> Coord:
    if dir == NORTH:
        return Coord(elf.coord.x - 1, elf.coord.y)
    if dir == EAST:
        return Coord(elf.coord.x, elf.coord.y + 1)
    if dir == SOUTH:
        return Coord(elf.coord.x + 1, elf.coord.y)
    if dir == WEST:
        return Coord(elf.coord.x, elf.coord.y - 1)
    else:
        print('ERROR: Unknown direction', dir)


dirs = [NORTH, SOUTH, WEST, EAST]

round = 0
while round < 10:
    proposals = []

    # proposing
    for elf in elves:
        if not elf_has_neighbors(elf):
            continue

        for dir in dirs:
            if not elf_has_neighbors_in_dir(elf, dir):
                proposals.append((elf, dir))
                break
    
    # checking whether other elves want to go there
    proposals = list(map(lambda p: (p[0], get_coords_of_elf_in_dir(p[0], p[1])), proposals))
    new_coords = list(map(lambda p: p[1], proposals))

    # don't move elves that want to go to the same coord
    for elf, new_coord in proposals.copy():
        if new_coords.count(new_coord) > 1:
            proposals = list(filter(lambda p: p[1] != new_coord, proposals))

    
    for elf, new_coord in proposals:
        # data[elf.coord.x][elf.coord.y] = '.'
        # data[new_coord.x][new_coord.y] = '#'
        elf.coord = new_coord

    popped = dirs.pop(0)
    dirs.append(popped)

    if printing_enabled:
        print()
        print(f'After round {round + 1}')
        print_data()

    round += 1

INFINITY = 99999999999
min_y = min_x = INFINITY
max_x = max_y = -INFINITY
for elf in elves:
    coord = elf.coord
    if coord.x < min_x:
        min_x = coord.x
    if coord.y < min_y:
        min_y = coord.y
    if coord.x > max_x:
        max_x = coord.x
    if coord.y > max_y:
        max_y = coord.y

if printing_enabled:
    print(min_x, min_y, max_x, max_y)
    print(abs(max_x - min_x), ' * ', abs(max_y - min_y), ' - ', len(elves))

solution = (abs(max_x - min_x) + 1) * (abs(max_y - min_y) + 1) - len(elves)
print(solution)

if should_submit:
    response = input(f'Are you sure you want to submit {solution}? [y/n] -> ')
    if 'y' in response:
        print(f'✨ Submitting solution: {solution}')
        submit(solution)
    else:
        print(f'❌ Not submitting solution: {solution}')

stop = timeit.default_timer()
print(f'\n✨ Time: {stop - start_time} s')
