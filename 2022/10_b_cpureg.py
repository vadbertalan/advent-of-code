# --- Part Two - --

# It seems like the X register controls the horizontal position of a sprite. Specifically, the sprite is 3 pixels wide, and the X register sets the horizontal position of the middle of that sprite. (In this system, there is no such thing as "vertical position": if the sprite's horizontal position puts its pixels where the CRT is currently drawing, then those pixels will be drawn.)

# You count the pixels on the CRT: 40 wide and 6 high. This CRT screen draws the top row of pixels left-to-right, then the row below that, and so on. The left-most pixel in each row is in position 0, and the right-most pixel in each row is in position 39.

# # , here are the cycles during which the first and last pixel in each row are drawn:
# Like the CPU, the CRT is tied closely to the clock circuit: the CRT draws a single pixel during each cycle. Representing each pixel of the screen as a

# Cycle   1 ->  # <- Cycle  40
# Cycle  41 ->  # <- Cycle  80
# Cycle  81 ->  # <- Cycle 120
# Cycle 121 ->  # <- Cycle 160
# Cycle 161 ->  # <- Cycle 200
# Cycle 201 ->  # <- Cycle 240
# So, by carefully timing the CPU instructions and the CRT drawing operations, you should be able to determine whether the sprite is visible the instant each pixel is drawn. If the sprite is positioned such that one of its three pixels is the pixel currently being drawn, the screen produces a lit pixel(  # ); otherwise, the screen leaves the pixel dark (.).

#     The first few pixels from the larger example above are drawn as follows:

#     Sprite position:  # .....................................

#     Start cycle   1: begin executing addx 15
#     During cycle  1: CRT draws pixel in position 0
#     Current CRT row:

#     During cycle  2: CRT draws pixel in position 1
#     Current CRT row:
#     End of cycle  2: finish executing addx 15 (Register X is now 16)
#     Sprite position: ...............  # ......................

#     Start cycle   3: begin executing addx - 11
#     During cycle  3: CRT draws pixel in position 2
#     Current CRT row:  # .

#     During cycle  4: CRT draws pixel in position 3
#     Current CRT row:  # ..
#     End of cycle  4: finish executing addx - 11 (Register X is now 5)
#     Sprite position: ....  # .................................

#     Start cycle   5: begin executing addx 6
#     During cycle  5: CRT draws pixel in position 4
#     Current CRT row:  # ..#

#     During cycle  6: CRT draws pixel in position 5
#     Current CRT row:  # ..##
#     End of cycle  6: finish executing addx 6 (Register X is now 11)
#     Sprite position: ..........  # ...........................

#     Start cycle   7: begin executing addx - 3
#     During cycle  7: CRT draws pixel in position 6
#     Current CRT row:  # ..##.

#     During cycle  8: CRT draws pixel in position 7
#     Current CRT row:  # ..##..
#     End of cycle  8: finish executing addx - 3 (Register X is now 8)
#     Sprite position: .......  # ..............................

#     Start cycle   9: begin executing addx 5
#     During cycle  9: CRT draws pixel in position 8
#     Current CRT row:  # ..##..#

#     During cycle 10: CRT draws pixel in position 9
#     Current CRT row:  # ..##..##
#     End of cycle 10: finish executing addx 5 (Register X is now 13)
#     Sprite position: ............  # .........................

#     Start cycle  11: begin executing addx - 1
#     During cycle 11: CRT draws pixel in position 10
#     Current CRT row:  # ..##..##.

#     During cycle 12: CRT draws pixel in position 11
#     Current CRT row:  # ..##..##..
#     End of cycle 12: finish executing addx - 1 (Register X is now 12)
#     Sprite position: ...........  # ..........................

#     Start cycle  13: begin executing addx - 8
#     During cycle 13: CRT draws pixel in position 12
#     Current CRT row:  # ..##..##..#

#     During cycle 14: CRT draws pixel in position 13
#     Current CRT row:  # ..##..##..##
#     End of cycle 14: finish executing addx - 8 (Register X is now 4)
#     Sprite position: ...  # ..................................

#     Start cycle  15: begin executing addx 13
#     During cycle 15: CRT draws pixel in position 14
#     Current CRT row:  # ..##..##..##.

#     During cycle 16: CRT draws pixel in position 15
#     Current CRT row:  # ..##..##..##..
#     End of cycle 16: finish executing addx 13 (Register X is now 17)
#     Sprite position: ................  # .....................

#     Start cycle  17: begin executing addx 4
#     During cycle 17: CRT draws pixel in position 16
#     Current CRT row:  # ..##..##..##..#

#     During cycle 18: CRT draws pixel in position 17
#     Current CRT row:  # ..##..##..##..##
#     End of cycle 18: finish executing addx 4 (Register X is now 21)
#     Sprite position: ....................  # .................

#     Start cycle  19: begin executing noop
#     During cycle 19: CRT draws pixel in position 18
#     Current CRT row:  # ..##..##..##..##.
#     End of cycle 19: finish executing noop

#     Start cycle  20: begin executing addx - 1
#     During cycle 20: CRT draws pixel in position 19
#     Current CRT row:  # ..##..##..##..##..

#     During cycle 21: CRT draws pixel in position 20
#     Current CRT row:  # ..##..##..##..##..#
#     End of cycle 21: finish executing addx - 1 (Register X is now 20)
#     Sprite position: ...................  # ..................
#     Allowing the program to run to completion causes the CRT to produce the following image:

#     # ..##..##..##..##..##..##..##..##..##..
#     # ...###...###...###...###...###...###.
#     # ....####....####....####....####....
#     # .....#####.....#####.....#####.....
#     ###### ......######......######......####
#     # .......#######.......#######.....
#     Render the image given by your program. What eight capital letters appear on your CRT?

#     Your puzzle answer was ZCBAJFJZ.

from collections import Counter
import math
import sys
import timeit
start_time = timeit.default_timer()

printing_enabled = sys.argv[1] == '-v' if len(sys.argv) == 2 else False

data_str1 = """noop
addx 3
addx -5"""

data_str2 = """addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop"""

data_str = """addx 1
noop
addx 5
addx -1
noop
addx 5
noop
addx -2
addx 8
addx -1
addx 7
noop
addx -1
addx 4
noop
addx 1
noop
noop
addx 6
addx -1
addx 3
addx 2
addx -5
addx -27
addx -3
addx 2
noop
addx 3
addx 2
addx 5
addx 2
addx 3
noop
addx 5
noop
noop
addx -2
addx 2
noop
addx -13
addx 23
noop
noop
addx -9
addx -18
addx 30
noop
addx -34
addx 3
addx -2
noop
addx 1
addx 6
noop
addx 28
addx -25
addx 5
addx 5
addx -11
addx 9
addx 4
noop
addx -26
addx 34
noop
addx -2
noop
noop
addx 4
addx -34
noop
addx 11
addx -7
addx 7
addx -9
addx 5
addx 5
addx 2
addx 1
noop
noop
noop
addx 22
addx -17
addx 2
noop
addx -19
addx 29
noop
addx -2
noop
addx 3
noop
noop
addx -36
addx 7
noop
noop
addx 3
addx -2
addx 2
addx 5
addx 2
addx 3
noop
addx 2
addx 11
addx -10
addx 2
addx 7
noop
addx -2
addx 5
addx 2
addx -36
addx 1
addx -1
addx 3
addx 4
addx -1
addx 5
noop
noop
noop
noop
noop
addx 3
addx 5
addx 15
addx -13
addx 6
addx -3
addx -1
addx 9
addx -1
addx 5
noop
addx 1
noop
noop
noop
noop"""

lines = data_str.split('\n')


data = []
N = len(lines)

if printing_enabled:
    print(N)
    for line in lines:
        print(line)
    print()


class Op():
    def __init__(self, instruction, amount=None):
        self.instruction = instruction
        self.amount = int(amount) if amount is not None else None

    # def __init__(self, instruction):
    #     self.instruction = instruction

    def __str__(self):
        return f'{self.instruction}{f" {self.amount}" if self.amount is not None else ""}'

    def __repr__(self):
        return f'{self.instruction}{f" {self.amount}" if self.amount is not None else ""}'


cycle_milestones = [40, 80, 120, 160, 200, 240]
cycle = 1
value = 1


def get_sprite_pos(value):
    return value, value + 1, value + 2


def print_and_advance_cycle():
    global cycle

    print('#' if cycle % 40 in get_sprite_pos(value) else '.', end='')
    if cycle in cycle_milestones:
        print()
    cycle += 1


for line in lines:
    line_split = line.split(' ')
    op = Op(*line_split)
    if op.instruction == 'noop':
        print_and_advance_cycle()
        continue
    if op.instruction == 'addx':
        for _ in range(2):
            print_and_advance_cycle()
        value += op.amount
    else:
        print('ERROR: unknown instruction')


stop = timeit.default_timer()
print(f'\n✨Time: {stop - start_time} s')
