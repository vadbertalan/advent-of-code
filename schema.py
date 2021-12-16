import math
import timeit
from collections import Counter
start_time = timeit.default_timer()


data_str1 = """"""

data_str = """"""

lines = data_str1.split('\n')

res = 0

for line in lines:
    print(line)

data = []
N = len(lines)
M = len(lines[0])
print(N, M)

for line in lines:
    data.append(list(map(int, line)))


def print_data():
    for row in data:
        for num in row:
            print(f'{num: < 3}', end='')
        print()
    print()


print_data()

print(res)

stop = timeit.default_timer()
print(f'\nTime: {stop - start_time} s')
