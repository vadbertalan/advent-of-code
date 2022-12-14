# https://adventofcode.com/2022/day/14

from aocd import data as aocdata, lines, numbers, submit
from collections import Counter
import math
import sys
import timeit
start_time = timeit.default_timer()


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

data_str_test = """498,4 -> 498,6 -> 496,6
504,4 -> 502,4 -> 502,9 -> 494,9"""

data_str = data_str_test if with_test_data else aocdata

lines = data_str.split('\n')


res = 0

paths = []
N = len(lines)

# if printing_enabled:
# 	print(N)
# 	for line in lines:
# 		print(line)
# 	print()

for line in lines:
	paths.append(
		list(map(lambda coord: Coord(int(coord.split(',')[0]), int(coord.split(',')[1])), line.split(' -> '))))

if printing_enabled:
	print(paths)

# Solution:

min_y = 0
min_x = 99999999999  # infinity
max_x = max_y = -1
for path in paths:
	for coord in path:
		if coord.x < min_x:
			min_x = coord.x
		if coord.x > max_x:
			max_x = coord.x
		if coord.y > max_y:
			max_y = coord.y

if printing_enabled:
	print(min_x, min_y, max_x, max_y)

for path in paths:
	for coord in path:
		coord.x -= min_x

old_min_x = min_x
min_x = 0
max_x -= old_min_x

if printing_enabled:
	print(paths)
	print(min_x, min_y, max_x, max_y)

AIR = '.'
SAND = 'o'
ROCK = '#'
SAND_SOURCE = '+'


data = list([[AIR for x in range(max_x + 1)] for y in range(max_y + 1)])


def print_data():
	row_count = 0
	for row in data:
		print(f'{row_count:<3}', end=' | ')
		row_count += 1
		for num in row:
			print(f'{num:<3}', end='')
		print()
	col_count = 0
	print('      ', end='')
	for i in range(max_x + 1):
		print(f'{i:<3}', end='')
	print()


def draw_rock_line(coord1, coord2):
	global data
	on_or_minus1_coordY = 1 if coord1.y <= coord2.y else -1
	for i in range(coord1.y, coord2.y + on_or_minus1_coordY, on_or_minus1_coordY):
		on_or_minus1_coordX = 1 if coord1.x <= coord2.x else -1
		for j in range(coord1.x, coord2.x + on_or_minus1_coordX, on_or_minus1_coordX):
			data[i][j] = ROCK


for path in paths:
	for i in range(len(path) - 1):
		draw_rock_line(path[i], path[i + 1])


sand_source_coord = Coord(500 - old_min_x, 0)
data[sand_source_coord.y][sand_source_coord.x] = SAND_SOURCE

if printing_enabled:
	print()
	print_data()
	# exit()


def is_valid_coord(coord: Coord) -> bool:
	return 0 <= coord.x <= max_x and 0 <= coord.y <= max_y


def get(coord):
	return data[coord.y][coord.x]


def set(coord, value):
	data[coord.y][coord.x] = value


def get_below_coord(coord: Coord) -> Coord:
	coord = Coord(coord.x, coord.y + 1)
	return coord


def get_below_left_coord(coord: Coord) -> Coord:
	coord = Coord(coord.x - 1, coord.y + 1)
	return coord


def get_below_right_coord(coord: Coord) -> Coord:
	coord = Coord(coord.x + 1, coord.y + 1)
	return coord


ok = True
while ok:
	sand_coord = Coord(sand_source_coord.x, sand_source_coord.y)
	can_move = True
	while can_move:
		prev_coord = sand_coord
		
		below_coord = get_below_coord(sand_coord)
		if not is_valid_coord(below_coord):
			ok = False
			set(sand_coord, AIR)
			break
		if get(below_coord) == AIR:
			sand_coord = below_coord
		else:
			below_left_coord = get_below_left_coord(sand_coord)
			if not is_valid_coord(below_left_coord):
				ok = False
				set(sand_coord, AIR)
				break
			if get(below_left_coord) == AIR:
				sand_coord = below_left_coord
			else:
				below_right_coord = get_below_right_coord(sand_coord)
				if not is_valid_coord(below_right_coord):
					ok = False
					set(sand_coord, AIR)
					break
				if get(below_right_coord) == AIR:
					sand_coord = below_right_coord
				else:
					can_move = False
		if can_move:
			set(sand_coord, SAND)
			if get(prev_coord) != SAND_SOURCE:
				set(prev_coord, AIR)
	# print_data()

# if printing_enabled:
print_data()

for i in range(max_y + 1):
	for j in range(max_x + 1):
		if data[i][j] == SAND:
			res += 1


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
