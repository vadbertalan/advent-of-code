# https://adventofcode.com/2022/day/22

# Your puzzle answer was 186128.

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

data_str_test = """        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5"""

data_str = data_str_test if with_test_data else aocdata

m, code_str = data_str.split('\n\n')
lines = m.split('\n')

data = []
N = len(lines)
M = len(lines[0])


for line in lines:
    data.append([ch for ch in line])

# parse code
instructions = []

move_amount_str = ''
dir = ''
for ch in code_str:
    if ch.isnumeric():
        move_amount_str += ch
    else:
        dir = ch
        move_amount = int(move_amount_str)
        instructions.append((move_amount, dir))
        move_amount_str = ''

last_move_amount = int(move_amount_str)

# get start coords
start_coord = None
for i, ch in enumerate(data[0]):
    if ch != ' ':
        start_coord = (0, i)
        break


if printing_enabled:
    print(N, M)
    for line in lines:
        print(line)
    for line in data:
        print(line)
    print(code_str)
    print(instructions)
    print('last_move_amount:', last_move_amount)
    print('start_coord:', start_coord)
    print()


# Solution:
NEXT_FACING_TURN_RIGHT = {
    'right': 'down',
    'down': 'left',
    'left': 'up',
    'up': 'right'
}
NEXT_FACING_TURN_LEFT = {
    'right': 'up',
    'up': 'left',
    'left': 'down',
    'down': 'right'
}
OPPOSITE_FACING = {
    'right': 'left',
    'up': 'down',
    'left': 'right',
    'down': 'up'
}


def change_facing(current_facing, left_or_right):
    assert (left_or_right in ['L', 'R'])
    assert (current_facing in ['left', 'down', 'right', 'up'])

    if left_or_right == 'R':
        return NEXT_FACING_TURN_RIGHT[current_facing]
    return NEXT_FACING_TURN_RIGHT[current_facing] if left_or_right == 'R' else NEXT_FACING_TURN_LEFT[current_facing]


def is_valid_coord(coord):
    global data
    return 0 <= coord[0] < N and 0 <= coord[1] < len(data[coord[0]])


def get_next_coord_with_facing(coord, facing):
    assert (facing in ['left', 'down', 'right', 'up'])

    if facing == 'right':
        return (coord[0], coord[1] + 1)
    if facing == 'down':
        return (coord[0] + 1, coord[1])
    if facing == 'left':
        return (coord[0], coord[1] - 1)
    # dir is 'up'
    return (coord[0] - 1, coord[1])


def get_wrapped_coord(coord, facing):
    # QUESTION: wrap till the end of the board??
    opposite_facing = OPPOSITE_FACING[facing]
    new_coord = coord
    isvalid = True
    while isvalid:
        next_coord = get_next_coord_with_facing(new_coord, opposite_facing)
        isvalid = is_valid_coord(
            next_coord) and get_char_on_coord(next_coord) != ' '
        if isvalid:
            new_coord = next_coord
    return new_coord


def move_once(coord, facing):
    new_coord = get_next_coord_with_facing(coord, facing)

    if not is_valid_coord(new_coord) or get_char_on_coord(new_coord) == ' ':
        new_coord = get_wrapped_coord(coord, facing)

    if get_char_on_coord(new_coord) == '#':
        return coord

    return new_coord


def get_char_on_coord(coord):
    global data
    return data[coord[0]][coord[1]]


coord = start_coord
facing = 'right'
for move_amount, dir in instructions:
    for _ in range(move_amount):
        coord = move_once(coord, facing)
    facing = change_facing(facing, dir)
    print(f'After move {move_amount} and turn {dir} =>', coord, facing)

# last move without turn
for _ in range(last_move_amount):
    coord = move_once(coord, facing)
print(f'After last move {last_move_amount} (no turn) =>', coord, facing)

points_of_dir = {
    'right': 0,
    'down': 1,
    'left': 2,
    'up': 3
}

coord = (coord[0] + 1, coord[1] + 1)
solution = 1000 * coord[0] + 4 * coord[1] + points_of_dir[facing]
print(coord, facing, solution)

if should_submit:
    response = input(f'Are you sure you want to submit {solution}? [y/n] -> ')
    if 'y' in response:
        print(f'✨ Submitting solution: {solution}')
        submit(solution)
    else:
        print(f'❌ Not submitting solution: {solution}')

stop = timeit.default_timer()
print(f'\n✨Time: {stop - start_time} s')
