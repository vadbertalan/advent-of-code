# https://adventofcode.com/2022/day/20
# Correct answer: 14888

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
    def __init__(self, nr, index) -> None:
        self.nr = nr
        self.index = index
        self.next = None

    def set_next(self, next):
        self.next = next
    
    def set_prev(self, prev):
        self.prev = prev

    def __str__(self) -> str:
        return f'({self.nr}, {self.index})'

    def __repr__(self) -> str:
        return f'({self.nr}, {self.index})'


data = []
nrs = []
i = 0
for line in data_str.split('\n'):
    nr = int(line)
    nrs.append(nr)
    data.append(Nr(nr, i))
    i += 1

for i in range(len(data) - 1):
    data[i].next = data[i + 1]
    data[i + 1].prev = data[i]
data[len(data) - 1].next = data[0]
data[0].prev = data[len(data) - 1]
first = data[0]

N = len(nrs)

if printing_enabled:
    print(data)
    print(len(data))
    print(nrs)
    print(len(nrs))





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
    # while it != node:
    #     it = it.next
    move_amount = abs(it.nr)
    # it.index += move_amount

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
        print(it.nr, end=' ')
        i += 1
        it = it.next
    print()

for node in data:
    move(node)




it = first
while it.nr != 0:
    it = it.next
it = it.next

sol_arr = []
i = 0
while i < 3000:
    if i == 999:
        sol_arr.append(it.nr)
    elif i == 1999:
        sol_arr.append(it.nr)
    elif i == 2999:
        sol_arr.append(it.nr)
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
