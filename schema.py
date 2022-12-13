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

data_str_test = """"""

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


# Solution:


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
