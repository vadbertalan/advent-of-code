# https://adventofcode.com/2022/day/22#part2

# (You guessed 113190.) too high
# (You guessed 113150.) too high

# Your puzzle answer was 34426.

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

data2 = {
    1: """...#..#.#........#.#..#......#..#.#....##.........
...#..........#.......................#...#.##....
........##....#..#....................#.......#...
..#..###..........................#..#...#.....##.
#..#........#............#..#........#..##.....#..
...................#...............#....#.........
..#......................#....#......#....#...##..
......#...................#....#..........#....#..
.................#..#.......#.#..#.#.....#.#......
....#......#........#...............#...#...#.....
........#.....#......#..............#......#......
........##...................#............##......
...........#......#.....#.#.....#................#
...................#..............#...............
......#.......#...................................
..............#.....#....#...#...........#..#.....
#..........#.......#......#......................#
....................##......#.....................
...#...#...##.#..........#...........#........#...
.#.....##...................#.......#.............
.....#..#..........#.........#....#.....#.........
...............#.#............#.#.#.....#.........
....#......#.....#..#......#...........#..........
..................#..................#......#.....
.#.#..........#.........#..............#......#.#.
..#..#............##............................##
...#...##...#....................#......#.........
.##.......#..............#..#..........#..........
#..#............#....#............#......#.......#
.......#.......#.#...............#.....#.#........
.....#.##........#................................
.....#.#....#....##...........................#...
.............#...................#......#.........
.........#..#.#.....##.........................#..
#..................#........#.......#..#.......#..
..........#..........................#............
................#.............#....#...........#..
#.#.......#..............#................#.#.....
....#...#..............#........#.....#...#.#..#..
...............#.#..............#....##.......#...
.................#........#..#............#.......
..#......#........#...............#....#..........
.#..............#...#..............#.......#......
...........#...#.................................#
............................#.#...................
......................#.....#.#.................#.
.#...#..........#...#.##.....#....................
...#.##.#.....#.#.#...............................
...#........................#.....#......#........
............#........##..#.....#..............#..#""",
    2: """.......##.#...#....#.............#........#.......
.............#.............#.......#...#..........
......#............#..#....#......................
........#..............##...#........#............
...#.#............##.......#.........###...#..#...
..........#.......#...........................#...
..........#.#...............................#.....
.........................#.......##....#.......#..
.#.................#............#......#...#......
#.........#......#.#.#.#....#.....................
...#................#.....##.....#................
.#.........#...........#...#.........#............
...#.................................#........##..
.#.....#............#.....#.....#..............#..
................#...............#.................
#........#...#.......#............................
..............#...###..#..........................
.........##...............#...................#...
#...............................#.......#..#......
.........##....#..................................
............#.......#......#....#...#...#..#.#...#
...#......#.................#.....#....#..........
........#...........................#............#
..#..#......#.............#...#..#................
.#..........................#....#................
..........##.#....#......#........................
..#..##.......#.......#....................#....#.
.....#................#...#........#..............
...............##.....................#......#....
.............#..#................................#
....#.....................#.....#.............#...
......#..........................#.............#..
....#....#............#...............#.#.........
...##.....................##...................#..
..................................##.....#..#.....
...............#......#.......#...#........#......
.................#......#....#...#....#...........
....................##.#.....#....#..#..##.......#
.#................................................
#..#..#...#.........#........#.....#..#...........
....#......#..................................#...
.....................#..........#.......#.........
............#..#...............#...............#..
....#..................#....#.......##..#........#
#................#....................#...#....#.#
......#......#.........................#..........
.................#....#.##.###.........#.##....#..
#............................##.......#..........#
..###..........#....#..#....#..#...............#..
.....#.....#.........#..#....................#....""",
    3: """.#.................#.......#....#.................
..#....#...............#......##......#...#.......
##.............................#.#.###..#...#.....
...#...............#...#.....#..#................#
.......................................#......#...
....................#...#......#...#...........#..
..#..#.....#......#.................#..#..........
......#..#.........#..................#.#....#....
.......#...........#...............#.........##.#.
#.#.......#....#.#..#...........................#.
................#...#......#.............#........
....#....................#........................
...#.#....................#.......#...............
........#.......#...................#............#
........#...........#.....#.........#.......#....#
....................##..............#.............
...........#..........................#...........
.............#...#..........................#.....
.......................#........##.......#........
.....#............................................
............#..#.........................#........
....#.........#...........##.....#......#....#.###
........#.......#.............#.......##....##....
......##.........##................#.#.........#..
.......##............#..................#.........
....#.....#.......#..#.......#.......#.....#......
.............#....................................
#........#........................................
....#..#.................#..............#........#
.............#....#.................#..#..........
.#.............#.......#......#.#..............#.#
.#.#....#............#.....#............##........
..#.........#....................#....#......#.#..
......#.............#......#...#......#...........
.....#......................................#...#.
..#............#......#......#....................
.......#........#.......................##.#.....#
.....#............................#....#.#........
.#....#.#..................#...............#......
#..................##.............................
#..#..............#..#................#..#........
......#....#...#......##..........................
#...#.........................................#.#.
........#....#....................................
.................#................................
......#......#.....................#..............
#.......#............#......#.....#..............#
.............#......#.....#...#........#..........
.....................#.....#..#......#............
.........#.......#.......#......................##""",
    4: """............................#...#............###..
..#............................#..................
..............#...#..................#.........#..
...............#......................##.####.....
#...#...........#.....................#...........
.....#......#.......................#..........#..
..........................#.......................
..........................#........#..##..........
#..................#...................#...#..#.#.
.#...#........#...................................
..................#..................#............
....#....#........................................
..........#...#......#.#.................#....#...
.........#...................#....#...............
.#..........#..................#.....##..#......#.
..........#..........#.......#.............#......
..................#.#................#...........#
.....#................#......#....................
...............#.........#.......#.........#......
...#...#.#...#........#.........#.........#.......
...............#.....................#............
.....#............#.....#..#............#...#..#..
.....#................#....#...............#......
.............#.........#......#...................
....................#.......#....................#
#........#..........#................#........#.#.
...........#............#..........#....#.........
.......#.....#...................#........#.....#.
......#...........................................
...............#...............##..............#..
.#.........##.................#......#............
......#...................................#...#...
...##...#.................#..........#............
......................#................#..........
..#.........#..#....#...#.......#.................
.#..#.##.#..............#.....##......#...........
...........#.....#............#...##.......#......
.......#...#...#..........#........#...........#.#
#..#.##.........#.#..........#.........#...#.....#
...#......#.......#...#....#..................#...
.....................................#.#..........
.....................#.........#.............#..#.
..#.#...#............................#............
..#.#...#...#...#........#...#..............#.#..#
.........#......#.................................
..........#...........#......................#....
.......#........#......#...........#....#.........
#.......#...#.....#........#......................
#...................#.......#.......#...#.........
..#...#.............##................#...........""",
    5: """.........#.#.........................#.#..........
...............#..............#........#..........
......#.#....................#...........#.....#..
............#.........##...........#.#............
...#.#.......#.....#...#...........#..........#...
..........#.............#.......#...#.............
..#.....................................#.........
#.....#.......#.........................##........
#..............#....#......#.......#.##..#........
................#.#.........#......#.##.#.........
.........#.............#....#.##......##..........
................#..................#..........##..
..#....#..#...#....................#.............#
#...#.......#..#...#......#.#...#.....#...........
......#..........................#................
..........#..#......#..#...........#....#.......#.
#.#........................##.....#...#......#....
#.........#....#...............#...........###....
.....#..........##....#.#........#................
.#...#..................................#.........
.....#..#...........................#.............
..........#.......................................
.......#.......###...#...#...##..................#
.......#......#...........#......#.......#........
.##....##..................#.#............##..#...
...................................#..............
...#....#..#............#.........................
......#...............................#........#..
.#..##.........#..#........#......................
.................#..#..#...............#..........
.......##....................#...............#....
.#.............................#....#.#...........
.......#.......#.....#..#.......#.........#....#..
...#.................#....#......#...#............
.......#....#...........#................##..#....
..............#......#...#...#....................
#...........#....##.....#.............##...#...#..
........#......#......##.#..#......#...#......#..#
..##...............#.#........#...................
#......#..........#...#....#......#.........#..#..
......................#......#.............#......
......................#...........................
.....#.............#....#..............#..........
#..............#........##.............#..........
#.......##..#....#.........#..#...#.......#.#.....
.#...##...........#...............................
.#..........#....#.....#...............#..........
.....#....................#.........#.#...........
...#......#.............#.......#.....#...........
.........#............#.#.....#.#.................""",
    6: """...#................#....#........#....#.......#.#
........#.......#..............#.#.#.....#.....#.#
.......#..#.........#.............#............#..
...#..............................#.......#.......
..#............#.....#...................#........
#......#......#.#....#...................##......#
...........#......#..........................#....
......#...............#..#....#............#......
..#.......#..#....#...............................
.....#...........................#................
#.....#..........#...............#............#...
.......#....##................#...............#...
.....................#.......#.....#..............
.....###..#........#.........##..#.......#........
..............................###.....#...........
......#......#..#.................................
.#.......#........#.#...............#.#...........
..........#.###.#..#........#........#..#.#.......
.#.#.......#.......#......#.....#.#...............
.............#...#........#.#......#..............
.#.#..............#..#..#.........................
.#..................#......#.......#.............#
......................#.#.........#.....#.......#.
.........#.....#..............##.#....#...........
.....#.............#....##........................
....................#.................#..........#
............#.................#..#....#...........
......#..#.............#...........##.............
..............................#...................
..#..............#.#..............#......#.......#
.............................#....#.......#.......
..............#...#.#......#..............#.......
#................................#................
..#......#.............#.................#...#....
............#...................#.......#.........
..#...#..#..##..#....#..........#.......#.........
.....#.............#.....#.....#.....#.........#..
.........#.....#...........#.......##......#......
....#.............#...#....#..#......#.........#..
........#........#.....#...#...........#.......#..
.....#............#.#....#...........##........#..
.#........#.......#...............................
...........#..#....#...............#..#...........
..##.........##....#.........##..#.............#..
.....#.#............####.#..............#...#.....
.....#...........#.............................#..
#..#.........#..............#...#...#..........#..
.....#.#........................#..........#...#..
...........#.............#................#.....#.
........#.#..#..............#..........#.........."""
}


def print_data():
    max_y = max(map(len, data))
    print('      ', end='')
    for i in range(max_y):
        print(f'{i:<3}', end='')
    print()

    print('      ', end='')
    for _ in range(max_y * 3):
        print('-', end='')
    print()

    row_count = 0
    for row in data:
        print(f'{row_count:<3}', end=' | ')

        for value in row:
            print(f'{value:<3}', end='')

        print(f' | {row_count}')

        row_count += 1

    print('      ', end='')
    for _ in range(max_y * 3):
        print('-', end='')
    print()

    print('      ', end='')
    for i in range(max_y):
        print(f'{i:<3}', end='')

    print()


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
    print_data()
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


def get_wrapped_coord_and_facing(coord, facing):
    # 1 - 6 edge
    if facing == 'up' and coord[0] == 0 and 50 <= coord[1] < 100:
        return (coord[1] + 100, 0), 'right'
    if facing == 'left' and 150 <= coord[0] < 200 and coord[1] == 0:
        return (0, coord[0] - 100), 'down'

    # 2 - 6 edge
    if facing == 'up' and coord[0] == 0 and 100 <= coord[1] < 150:
        return (199, coord[1] - 100), 'up'
    if facing == 'down' and coord[0] == 199 and 0 <= coord[1] < 50:
        return (0, coord[1] + 100), 'down'

    # 4 - 6 edge
    if facing == 'down' and coord[0] == 149 and 50 <= coord[1] < 100:
        return (coord[1] + 100, 49), 'left'
    if facing == 'right' and 150 <= coord[0] < 200 and coord[1] == 49:
        return (149, coord[0] - 100), 'up'

    # 5 - 6 edge implicit

    # 1 - 5 edge
    if facing == 'left' and 0 <= coord[0] < 50 and coord[1] == 50:
        return (149 - coord[0], 0), 'right'
    if facing == 'left' and 100 <= coord[0] < 150 and coord[1] == 0:
        return (149 - coord[0], 50), 'right'

    # 1 - 2 edge implicit
    
    # 1 - 3 edge implicit

    # 3 - 5 edge
    if facing == 'left' and 50 <= coord[0] < 100 and coord[1] == 50:
        return (100, coord[0] - 50), 'down'
    if facing == 'up' and coord[0] == 100 and 0 <= coord[1] < 50:
        return (coord[1] + 50, 50),  'right'

    # 2 - 3 edge
    if facing == 'down' and coord[0] == 49 and 100 <= coord[1] < 150:
        return (coord[1] - 50, 99), 'left'
    if facing == 'right' and 50 <= coord[0] < 100 and coord[1] == 99:
        return (49, coord[0] + 50), 'up'

    # 2 - 4 edge
    if facing == 'right' and 0 <= coord[0] < 50 and coord[1] == 149:
        return (149 - coord[0], 99), 'left'
    if facing == 'right' and 100 <= coord[0] < 150 and coord[1] == 99:
        return (149 - coord[0], 149), 'left'

    # 3 - 4 edge implicit

    # 4 - 5 edge implicit

    print('ERROR: No case entered in wrapping', coord, facing)


def move_once(coord, facing):
    new_coord = get_next_coord_with_facing(coord, facing)
    new_facing = facing

    if not is_valid_coord(new_coord) or get_char_on_coord(new_coord) == ' ':
        new_coord, new_facing = get_wrapped_coord_and_facing(coord, facing)

    if get_char_on_coord(new_coord) == '#':
        return coord, facing

    return new_coord, new_facing


def get_char_on_coord(coord):
    global data
    return data[coord[0]][coord[1]]


coord = start_coord
facing = 'right'
for move_amount, dir in instructions:
    for _ in range(move_amount):
        coord, facing = move_once(coord, facing)
    facing = change_facing(facing, dir)
    print(f'After move {move_amount} and turn {dir} =>', coord, facing)

# last move without turn
for _ in range(last_move_amount):
    coord, facing = move_once(coord, facing)
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
print(f'\nTime: {stop - start_time} s')
