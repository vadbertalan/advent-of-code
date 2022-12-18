from typing import List
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


data_str_test = """>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"""


jets = data_str_test if with_test_data else aocdata

rocks = [
        ['..@@@@.'],
        [
            '...@...',
            '..@@@..',
            '...@...'
        ],
    [
            '....@..',
            '....@..',
            '..@@@..'
    ],
    [
            '..@....',
            '..@....',
            '..@....',
            '..@....',
    ],
    [
            '..@@...',
            '..@@...'
    ]
]

AIR = '.'
FALLING_ROCK = '@'
STATIC_ROCK = '#'

chamber = ['-------']
rock_count = 0
jet_count = 0


def print_chamber():
    for row in chamber:
        print('|', end='')
        print(row, end='')
        print('|')


def rock_shifted_left(rock: List[str]) -> List[str]:
    shifted_rock = rock.copy()
    shifted_rock_temp = []

    for row_idx, row in enumerate(shifted_rock):
        modified_row = [*row]

        moved_first_rock_in_row = False
        for i in range(1, len(row)):
            if row[i] == AIR:  # if not air, it has to be moved left
                if not moved_first_rock_in_row:
                    if row[i - 1] != AIR:  # there's something in the way
                        return shifted_rock
                    moved_first_rock_in_row = True

                modified_row[i - 1] = row[i]
                modified_row[i] = AIR

        shifted_rock_temp.append(''.join(modified_row))

    return shifted_rock_temp


print(rock_shifted_left(['...@...',
                         '..@@@..',
                         '...@...']))
exit()


def rock_shifted_left(rock: List[str]) -> List[str]:
    shifted_rock = rock.copy()
    shifted_rock_temp = []

    for row_idx, row in enumerate(shifted_rock):
        modified_row = [*row]

        moved_first_rock_in_row = False
        for i in range(1, len(row)):
            if row[i] == AIR:  # if not air, it has to be moved left
                if not moved_first_rock_in_row:
                    if row[i - 1] != AIR:  # there's something in the way
                        return shifted_rock
                    moved_first_rock_in_row = True

                modified_row[i - 1] = row[i]
                modified_row[i] = AIR

        shifted_rock_temp.append(''.join(modified_row))

    return shifted_rock_temp


while rock_count < 2020:
    current_rock = rocks[rock_count % len(rocks)].copy()

    air_below_falling_rock = [AIR * 7, AIR * 7, AIR * 7]

    chamber = [*current_rock, *air_below_falling_rock, *chamber]

    print_chamber()

    falling_rock_has_to_stop = False
    while not falling_rock_has_to_stop:
        while len(air_below_falling_rock) > 0:
            jet = jets[jet_count % len(jets)]
            current_rock_shifted = rock_shifted_right(
                current_rock) if jet == '>' else rock_shifted_left(current_rock)

            air_below_falling_rock = air_below_falling_rock[1:]

            chamber = [*current_rock_shifted,
                       *air_below_falling_rock,
                       *chamber]

if printing_enabled:
    print_chamber()


# Solution:
solution = 0


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
