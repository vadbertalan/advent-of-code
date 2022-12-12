# --- Part Two - --

# You're worried you might not ever get your items back. So worried, in fact, that your relief that a monkey's inspection didn't damage an item no longer causes your worry level to be divided by three.

# Unfortunately, that relief was all that was keeping your worry levels from reaching ridiculous levels. You'll need to find another way to keep your worry levels manageable.

# At this rate, you might be putting up with these monkeys for a very long time - possibly 10000 rounds!

# With these new rules, you can still figure out the monkey business after 10000 rounds. Using the same example above:

# == After round 1 ==
# Monkey 0 inspected items 2 times.
# Monkey 1 inspected items 4 times.
# Monkey 2 inspected items 3 times.
# Monkey 3 inspected items 6 times.

# == After round 20 ==
# Monkey 0 inspected items 99 times.
# Monkey 1 inspected items 97 times.
# Monkey 2 inspected items 8 times.
# Monkey 3 inspected items 103 times.

# == After round 1000 ==
# Monkey 0 inspected items 5204 times.
# Monkey 1 inspected items 4792 times.
# Monkey 2 inspected items 199 times.
# Monkey 3 inspected items 5192 times.

# == After round 2000 ==
# Monkey 0 inspected items 10419 times.
# Monkey 1 inspected items 9577 times.
# Monkey 2 inspected items 392 times.
# Monkey 3 inspected items 10391 times.

# == After round 3000 ==
# Monkey 0 inspected items 15638 times.
# Monkey 1 inspected items 14358 times.
# Monkey 2 inspected items 587 times.
# Monkey 3 inspected items 15593 times.

# == After round 4000 ==
# Monkey 0 inspected items 20858 times.
# Monkey 1 inspected items 19138 times.
# Monkey 2 inspected items 780 times.
# Monkey 3 inspected items 20797 times.

# == After round 5000 ==
# Monkey 0 inspected items 26075 times.
# Monkey 1 inspected items 23921 times.
# Monkey 2 inspected items 974 times.
# Monkey 3 inspected items 26000 times.

# == After round 6000 ==
# Monkey 0 inspected items 31294 times.
# Monkey 1 inspected items 28702 times.
# Monkey 2 inspected items 1165 times.
# Monkey 3 inspected items 31204 times.

# == After round 7000 ==
# Monkey 0 inspected items 36508 times.
# Monkey 1 inspected items 33488 times.
# Monkey 2 inspected items 1360 times.
# Monkey 3 inspected items 36400 times.

# == After round 8000 ==
# Monkey 0 inspected items 41728 times.
# Monkey 1 inspected items 38268 times.
# Monkey 2 inspected items 1553 times.
# Monkey 3 inspected items 41606 times.

# == After round 9000 ==
# Monkey 0 inspected items 46945 times.
# Monkey 1 inspected items 43051 times.
# Monkey 2 inspected items 1746 times.
# Monkey 3 inspected items 46807 times.

# == After round 10000 ==
# Monkey 0 inspected items 52166 times.
# Monkey 1 inspected items 47830 times.
# Monkey 2 inspected items 1938 times.
# Monkey 3 inspected items 52013 times.
# After 10000 rounds, the two most active monkeys inspected items 52166 and 52013 times. Multiplying these together, the level of monkey business in this situation is now 2713310158.

# Worry levels are no longer divided by three after each item is inspected
# you'll need to find another way to keep your worry levels manageable. Starting again from the initial state in your puzzle input, what is the level of monkey business after 10000 rounds?


# !!! Should have been resolved with multimod solution => (a + b) mod x = ((a mod x) + (b mod x)) mod x
# Thus the divisors should be multiplied and each worry level should be modded by it.


# Your puzzle answer was 17408399184.

from typing import Dict
from aocd import data as aocdata, lines, numbers, submit
from collections import Counter
import math
import sys
import timeit
import json
from functools import reduce
start_time = timeit.default_timer()


printing_enabled = sys.argv[1] == '-v' if len(sys.argv) == 2 else False

data_str1 = """Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1"""

data_str = aocdata

lines = data_str.split('\n')

res = 0

data = []
N = len(lines)

if printing_enabled:
    print(N)
    for line in lines:
        print(line)
    print()

monkey_ids = []
current_monkey = None
d = {}


def get_add_fn(X):
    def ret(old):
        return old + X
    return ret


def get_mul_fn(X):
    def ret(old):
        return old * X
    return ret


def add_with_self(old):
    return old + old


def mul_with_self(old):
    return old * old


def prime_factors(num):
    ret = []

    while num % 2 == 0:
        ret.append(2)
        num = num / 2

    for i in range(3, int(math.sqrt(num)) + 1, 2):
        while num % i == 0:
            ret.append(i)
            num = num / i

    if num > 2:
        ret.append(num)

    return ret


def prime_factors_set(num) -> set:
    ret = set()

    while num % 2 == 0:
        ret.add(2)
        num = num / 2

    for i in range(3, int(math.sqrt(num)) + 1, 2):
        while num % i == 0:
            ret.add(i)
            num = num / i

    if num > 2:
        ret.add(num)

    return ret


PRIMES = [2, 3, 5, 7, 11, 13, 17, 19, 23]


def get_magic(op, nrs) -> list:
    # magic = {}
    # for prime in PRIMES:
    #     magic[prime] = number % prime == 0
    # return magic

    # return prime_factors(number)

    # return prime_factors_set(number)

    # magic = [None for _ in range(len(monkey_ids))]
    magic = []

    for monkey_id in monkey_ids:
        # magic.append(d[monkey_id]['operation'](
        #     number % d[monkey_id]['divbyamount']) % d[monkey_id]['divbyamount'])
        magic.append(op(nrs[monkey_id]) %
                     d[monkey_id]['divbyamount'])

    return magic


def merge_magics(magic1, magic2) -> set:
    # new_magic = {}
    # for key, value in magic1.items():
    #     new_magic[key] = value or magic2[key]
    # return new_magic

    # return [*magic1, *magic2]

    return magic1.union(magic2)


for line in lines:
    line = line.strip()
    if line.startswith('Monkey'):
        current_monkey = int(line.split(' ')[1][:-1])
        monkey_ids.append(current_monkey)
    elif line.startswith('Starting items'):
        # starting_items = list(
        #     map(lambda x: get_magic(int(x[:-1] if x[-1] == ',' else x)), line.split(' ')[2:]))
        starting_items = list(
            map(lambda x: int(x[:-1] if x[-1] == ',' else x), line.split(' ')[2:]))
        d[current_monkey] = {
            'startingitems': starting_items, 'inspectcount': 0}
    elif line.startswith('Operation'):
        op, amount = line.split(' ')[-2:]
        # amount = int(amount)
        fns = (add_with_self, mul_with_self) if amount == 'old' else (
            get_add_fn(int(amount)), get_mul_fn(int(amount)))
        d[current_monkey]['operation'] = fns[0] if op == '+' else fns[1] if op == '*' else lambda _: print(
            'ERROR: could not parse op')
        d[current_monkey]['op'] = op
        d[current_monkey]['operand'] = amount
    elif line.startswith('Test'):
        amount = int(line.split(' ')[-1])
        d[current_monkey]['divbyamount'] = amount
    elif line.startswith('If true'):
        throw_to = int(line.split(' ')[-1])
        d[current_monkey]['throwtoiftrue'] = throw_to
    elif line.startswith('If false'):
        throw_to = int(line.split(' ')[-1])
        d[current_monkey]['throwtoiffalse'] = throw_to


for monkey_id in monkey_ids:
    # d[monkey_id]['startingitems'] = list(
    #     map(lambda num: get_magic(num), d[monkey_id]['startingitems']))
    d[monkey_id]['startingitems'] = list(
        map(lambda num: [num % d[i]['divbyamount'] for i in range(len(monkey_ids))], d[monkey_id]['startingitems']))


def print_data():
    if printing_enabled:
        print(monkey_ids)
        for key, value in d.items():
            print(key, value['inspectcount'])
        print()


ROUNDS = 10000

# for round_id in range(ROUNDS):
#     for monkey_id in monkey_ids:
#         d[monkey_id]['inspectcount'] += len(d[monkey_id]['startingitems'])
#         for magic in d[monkey_id]['startingitems']:
#             if d[monkey_id]['op'] == '+':
#                 magic = get_magic(d[monkey_id]['operation'](reduce_to_1_item(magic)))

#             elif d[monkey_id]['op'] == '*':
#                 if d[monkey_id]['operand'] != 'old':
#                     magic = merge_magics(magic, get_magic(int(d[monkey_id]['operand'])))

#             else:
#                 print(f'ERROR: invalid op {d[monkey_id]["op"]}')
#                 exit()

#             # worry = math.floor(worry / 3)
#             # print(magic)
#             if magic[d[monkey_id]['divbyamount']]:
#                 d[d[monkey_id]['throwtoiftrue']]['startingitems'].append(magic)
#             else:
#                 d[d[monkey_id]['throwtoiffalse']
#                   ]['startingitems'].append(magic)
#         d[monkey_id]['startingitems'] = []
#     print(f'\n === Round {round_id + 1} ===')
#     print_data()


# for round_id in range(ROUNDS):
#     for monkey_id in monkey_ids:
#         d[monkey_id]['inspectcount'] += len(d[monkey_id]['startingitems'])
#         mul_magic = get_magic(
#             int(d[monkey_id]['operand'])) if d[monkey_id]['operand'] != 'old' else None
#         for magic in d[monkey_id]['startingitems']:
#             if d[monkey_id]['op'] == '+':
#                 magic = get_magic(d[monkey_id]['operation'](
#                     reduce(lambda x, y: x * y, magic)))

#             elif d[monkey_id]['op'] == '*':
#                 if d[monkey_id]['operand'] != 'old':
#                     magic = merge_magics(magic, mul_magic)
#                 # magic = merge_magics(magic, mul_magic if d[monkey_id]['operand'] != 'old' else magic)

#             else:
#                 print(f'ERROR: invalid op {d[monkey_id]["op"]}')
#                 exit()

#             # worry = math.floor(worry / 3)
#             # print(magic)
#             if d[monkey_id]['divbyamount'] in magic:
#                 d[d[monkey_id]['throwtoiftrue']]['startingitems'].append(magic)
#             else:
#                 d[d[monkey_id]['throwtoiffalse']
#                   ]['startingitems'].append(magic)
#         d[monkey_id]['startingitems'] = []
#     print(f'\n === Round {round_id + 1} ===')
#     print_data()

# worries = []
# for current_monkey in monkey_ids:
#     worries.extend(map(lambda worry: (current_monkey, worry), d[current_monkey]['startingitems']))

# for current_monkey, worry in worries:
#     print('Monkey {current_monkey} and worry {worry}')
#     for round_id in range(ROUNDS):
#         print(f'Monkey {current_monkey} and worry {worry}')
#         print(f'\n \t=== Round {round_id + 1} ===')

#         d[current_monkey]['inspectcount'] += 1
#         worry = d[current_monkey]['operation'](worry)

#         if worry % d[current_monkey]['divbyamount'] == 0:
#             current_monkey = d[current_monkey]['throwtoiftrue']
#         else:
#             current_monkey = d[current_monkey]['throwtoiffalse']

for round_id in range(ROUNDS):
    for monkey_id in monkey_ids:
        d[monkey_id]['inspectcount'] += len(d[monkey_id]['startingitems'])
        for magic in d[monkey_id]['startingitems']:
            r = magic[monkey_id]
            # r = d[monkey_id]['operation'](
            #     r % d[monkey_id]['divbyamount'])
            # r %= d[monkey_id]['divbyamount']
            new_magic = get_magic(d[monkey_id]['operation'], magic)
            if new_magic[monkey_id] == 0:
                d[d[monkey_id]['throwtoiftrue']
                  ]['startingitems'].append(new_magic)
            else:
                d[d[monkey_id]['throwtoiffalse']
                  ]['startingitems'].append(new_magic)
        d[monkey_id]['startingitems'] = []
    print(f'\n === Round {round_id + 1} ===')
    print_data()

inspects = (list(map(lambda x: (x, d[x]['inspectcount']), d.keys())))
print(inspects)
top2 = list(map(lambda x: x[1], sorted(
    inspects, key=lambda x: x[1], reverse=True)[:2]))
print(top2)
print(top2[0] * top2[1])

stop = timeit.default_timer()
print(f'\n✨Time: {stop - start_time} s')
