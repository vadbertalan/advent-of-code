# https://adventofcode.com/2022/day/13

import functools
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

data_str_test = """[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]"""

data_str = data_str_test if with_test_data else aocdata

lines = data_str.split('\n')

items = list(map(eval, filter(lambda line: line != '', lines)))

res = 0

data = []


def cmp_ints(left, right) -> int:
    if left == right:
        return 0
    return -1 if left < right else 1


def cmp_lists(left, right) -> int:
    i = 0
    while i < len(left) and i < len(right):
        cmp_res = cmp(left[i], right[i])
        if cmp_res != 0:
            return cmp_res
        i += 1  # equal
    return 0 if len(left) == len(right) else (-1 if i < len(right) else 1)


def cmp_int_with_list(left, right, is_left_int=True) -> int:
    list_left = list_right = []

    if is_left_int:
        list_left = [left]
        list_right = right
    else:
        list_left = left
        list_right = [right]

    return cmp_lists(list_left, list_right)


def cmp(left, right) -> int:
    if type(left) == int:
        if type(right) == int:
            return cmp_ints(left, right)
        return cmp_int_with_list(left, right)
    if type(right) == int:
        return cmp_int_with_list(left, right, is_left_int=False)
    return cmp_lists(left, right)


# Solution:
# def get_in_order_indices(pairs) -> int:
# 	for i, pair in enumerate(pairs):
# 		left, right = pair
# 		if cmp(left, right) == -1:
# 			data.append(i + 1)

items = [*items, [[2]], [[6]]]
items.sort(key=functools.cmp_to_key(cmp))
# print(items)
for item in items:
	print(item)

res = (items.index([[2]]) + 1) * (items.index([[6]]) + 1)

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
