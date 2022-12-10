from collections import Counter
import math
import sys
import timeit
start_time = timeit.default_timer()

from aocd import data as aocdata, lines, numbers, submit

printing_enabled = sys.argv[1] == '-v' if len(sys.argv) == 2 else False

data_str1 = """"""

data_str = aocdata

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
    data.append(list(map(int, line)))


def print_data():
    for row in data:
        for num in row:
            print(f'{num: < 3}', end='')
        print()
    print()


if printing_enabled:
    print_data()
    print()

print(res)

stop = timeit.default_timer()
print(f'\n✨Time: {stop - start_time} s')
