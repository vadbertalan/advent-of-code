# https://adventofcode.com/2022/day/20

# (You guessed -11910070820275.) no help if it's too big or no :()
# (You guessed 2119870867636.) too low :O

# Your puzzle answer was 3760092545849.


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


data_str_test = """1
2
-3
3
-2
0
4"""


data_str = data_str_test if with_test_data else aocdata


class Nr():
    def __init__(self, nr, index, bignr) -> None:
        self.nr = nr
        self.bignr = bignr
        self.index = index
        self.next = None

    def set_next(self, next):
        self.next = next

    def set_prev(self, prev):
        self.prev = prev

    def __str__(self) -> str:
        return f'({self.nr}, {self.index} {self.bignr})'

    def __repr__(self) -> str:
        return f'({self.nr}, {self.index} {self.bignr})'


KEY = 811589153
data = []
i = 0
for line in data_str.split('\n'):
    nr = int(line) * KEY
    data.append(Nr(nr, i, nr))
    i += 1


N = len(data)
for nr in data:
    if nr.nr >= 0:
        nr.nr %= (N - 1)
    else:
        nr.nr = (-nr.nr) % (N - 1)
        nr.nr *= -1

for i in range(len(data) - 1):
    data[i].next = data[i + 1]
    data[i + 1].prev = data[i]
data[len(data) - 1].next = data[0]
data[0].prev = data[len(data) - 1]
first = data[0]

if printing_enabled:
    print(data)
    print(len(data))

# Solution:
solution = 0


def swap(node1, node2):
    node1.next = node2.next
    node2.next.prev = node1

    node1.prev.next = node2
    node2.prev = node1.prev

    node1.prev = node2
    node2.next = node1

    global N, first

    node1.index = (node1.index + 1) % N
    node2.index = (N - 1) if node2.index == 0 else node2.index - 1

    if node1 == first:
        first = node2
    elif node2 == first:
        first = node1


def move(node):
    it = node
    move_amount = abs(it.nr)  # THIS HAS BEEN CHANGED....
    # move_amount = abs(it.bignr)

    while move_amount > 0:
        if it.nr > 0:
            swap(it, it.next)
        else:
            swap(it.prev, it)

        move_amount -= 1


def print_nodes():
    global first
    it = first
    i = 0
    while i < N:
        print(it.bignr, end=' ')
        i += 1
        it = it.next
    print()


if printing_enabled:
    print_nodes()
print('----- start mixing ------')
for i in range(10):
    print('After mixing no. ', i + 1)
    for node in data:
        move(node)
    if printing_enabled:
        print_nodes()
print('----- end mixing ------')
if printing_enabled:
    print_nodes()

it = first
while it.bignr != 0:
    it = it.next
it = it.next

sol_arr = []
i = 1
while i <= 3000:
    if i == 1000:
        sol_arr.append(it.bignr)
    elif i == 2000:
        sol_arr.append(it.bignr)
    elif i == 3000:
        sol_arr.append(it.bignr)
    it = it.next
    i += 1

solution = sum(sol_arr)

print(sol_arr)
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
