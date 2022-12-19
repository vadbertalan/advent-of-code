# https://adventofcode.com/2022/day/18

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


data_str_test = """2,2,2
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


data_str = data_str_test if with_test_data else aocdata


class Cube(object):
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


cubes=[]
for line in data_str.split('\n'):
    cubes.append(Cube(*map(int, line.split(','))))  # start here ...


if printing_enabled:
    print(cubes)
    print()

print(cubes[0].has_common_side(cubes[1]))

# Solution:
solution=0
connected_sides = 0

for i in range(len(cubes) - 1):
    for j in range(i + 1, len(cubes)):
        connected_sides += int(cubes[i].has_common_side(cubes[j]))

print(connected_sides)
solution = len(cubes) * 6 - 2 * connected_sides

print(solution)

if should_submit:
    response=input(f'Are you sure you want to submit {solution}? [y/n] -> ')
    if 'y' in response:
        print(f'✨ Submitting solution: {solution}')
        submit(solution)
    else:
        print(f'❌ Not submitting solution: {solution}')

stop=timeit.default_timer()
print(f'\n✨Time: {stop - start_time} s')
