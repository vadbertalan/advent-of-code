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

	def __eq__(self, other):
		return self.x == other.x and self.y == other.y

	def get_tuple(self):
		return self.x, self.y



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

data_str_test = """Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3"""

data_str = data_str_test if with_test_data else aocdata

lines = data_str.split('\n')


def calc_manhattan_distance(point1: Coord, point2: Coord):
	return abs(point1.x - point2.x) + abs(point1.y - point2.y)


sensors = []
beacons = []
sensor_dists = {}

for line in lines:
	line_split = line.split(' ')
	sensor_coord = Coord(int(line_split[2].split('=')[1].split(
		',')[0]), int(line_split[3].split('=')[1].split(':')[0]))
	beacon_coord = Coord(int(line_split[8].split('=')[1].split(
		',')[0]), int(line_split[9].split('=')[1]))

	sensors.append(sensor_coord)
	if beacon_coord not in beacons:
		beacons.append(beacon_coord)
	sensor_dists[sensor_coord.get_tuple()] = calc_manhattan_distance(
		sensor_coord, beacon_coord)

if printing_enabled:
	print(len(sensors))
	print(sensors)
	print(beacons)
	print(sensor_dists)
	print()

INFINITY = 99999999999
min_y = min_x = INFINITY
max_x = max_y = -INFINITY
for coord_it in [*sensors, *beacons]:
	if coord_it.x < min_x:
		min_x = coord_it.x
	if coord_it.y < min_y:
		min_y = coord_it.y
	if coord_it.x > max_x:
		max_x = coord_it.x
	if coord_it.y > max_y:
		max_y = coord_it.y


if printing_enabled:
	print()
	print(min_x, min_y, max_x, max_y)
	print()

# Solution:

X_MAX = 20 if with_test_data else 4000000
Y_MAX = 20 if with_test_data else 4000000

solution_coord = None
y = 0
while y <= Y_MAX and solution_coord is None:
	x = 0
	while x <= X_MAX:
		coord_it = Coord(x, y)

		if coord_it not in beacons:
			under_sensor_coverage = False
			for sensor_coord in sensors:
				dist = calc_manhattan_distance(coord_it, sensor_coord)
				coverage = sensor_dists[sensor_coord.get_tuple()]
				if dist <= coverage:
					under_sensor_coverage = True
					if printing_enabled:
						print(
							coord_it, f' | sensor {sensor_coord} with cov {sensor_dists[sensor_coord.get_tuple()]} -> dist {dist}')
					
					x += coverage + (sensor_coord.x - x) - abs(sensor_coord.y - y) 

					break
			
			if not under_sensor_coverage:
				print(f'NOT UNDER COV -> {coord_it}')
				solution_coord = coord_it
				break
			
		x += 1
		
	y += 1


print(solution_coord)

solution = solution_coord.x * 4000000 + solution_coord.y
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
