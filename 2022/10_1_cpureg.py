# --- Day 10: Cathode-Ray Tube - --

# You avoid the ropes, plunge into the river, and swim to shore.

# The Elves yell something about meeting back up with them upriver, but the river is too loud to tell exactly what they're saying. They finish crossing the bridge and disappear from view.

# Situations like this must be why the Elves prioritized getting the communication system on your handheld device working. You pull it out of your pack, but the amount of water slowly draining from a big crack in its screen tells you it probably won't be of much immediate use.

# Unless, that is, you can design a replacement for the device's video system! It seems to be some kind of cathode-ray tube screen and simple CPU that are both driven by a precise clock circuit. The clock circuit ticks at a constant rate
# each tick is called a cycle.

# Start by figuring out the signal being sent by the CPU. The CPU has a single register, X, which starts with the value 1. It supports only two instructions:

# addx V takes two cycles to complete. After two cycles, the X register is increased by the value V. (V can be negative.)
# noop takes one cycle to complete. It has no other effect.
# The CPU uses these instructions in a program(your puzzle input) to, somehow, tell the screen what to draw.

# Consider the following small program:

# noop
# addx 3
# addx - 5
# Execution of this program proceeds as follows:

# At the start of the first cycle, the noop instruction begins execution. During the first cycle, X is 1. After the first cycle, the noop instruction finishes execution, doing nothing.
# At the start of the second cycle, the addx 3 instruction begins execution. During the second cycle, X is still 1.
# During the third cycle, X is still 1. After the third cycle, the addx 3 instruction finishes execution, setting X to 4.
# At the start of the fourth cycle, the addx - 5 instruction begins execution. During the fourth cycle, X is still 4.
# During the fifth cycle, X is still 4. After the fifth cycle, the addx - 5 instruction finishes execution, setting X to - 1.
# Maybe you can learn something by looking at the value of the X register throughout execution. For now, consider the signal strength (the cycle number multiplied by the value of the X register) during the 20th cycle and every 40 cycles after that (that is , during the 20th, 60th, 100th, 140th, 180th, and 220th cycles).

# For example, consider this larger program:

# addx 15
# addx - 11
# addx 6
# addx - 3
# addx 5
# addx - 1
# addx - 8
# addx 13
# addx 4
# noop
# addx - 1
# addx 5
# addx - 1
# addx 5
# addx - 1
# addx 5
# addx - 1
# addx 5
# addx - 1
# addx - 35
# addx 1
# addx 24
# addx - 19
# addx 1
# addx 16
# addx - 11
# noop
# noop
# addx 21
# addx - 15
# noop
# noop
# addx - 3
# addx 9
# addx 1
# addx - 3
# addx 8
# addx 1
# addx 5
# noop
# noop
# noop
# noop
# noop
# addx - 36
# noop
# addx 1
# addx 7
# noop
# noop
# noop
# addx 2
# addx 6
# noop
# noop
# noop
# noop
# noop
# addx 1
# noop
# noop
# addx 7
# addx 1
# noop
# addx - 13
# addx 13
# addx 7
# noop
# addx 1
# addx - 33
# noop
# noop
# noop
# addx 2
# noop
# noop
# noop
# addx 8
# noop
# addx - 1
# addx 2
# addx 1
# noop
# addx 17
# addx - 9
# addx 1
# addx 1
# addx - 3
# addx 11
# noop
# noop
# addx 1
# noop
# addx 1
# noop
# noop
# addx - 13
# addx - 19
# addx 1
# addx 3
# addx 26
# addx - 30
# addx 12
# addx - 1
# addx 3
# addx 1
# noop
# noop
# noop
# addx - 9
# addx 18
# addx 1
# addx 2
# noop
# noop
# addx 9
# noop
# noop
# noop
# addx - 1
# addx 2
# addx - 37
# addx 1
# addx 3
# noop
# addx 15
# addx - 21
# addx 22
# addx - 6
# addx 1
# noop
# addx 2
# addx 1
# noop
# addx - 10
# noop
# noop
# addx 20
# addx 1
# addx 2
# addx 2
# addx - 6
# addx - 11
# noop
# noop
# noop
# The interesting signal strengths can be determined as follows:

# During the 20th cycle, register X has the value 21, so the signal strength is 20 * 21 = 420. (The 20th cycle occurs in the middle of the second addx - 1, so the value of register X is the starting value, 1, plus all of the other addx values up to that point: 1 + 15 - 11 + 6 - 3 + 5 - 1 - 8 + 13 + 4=21.)
# During the 60th cycle, register X has the value 19, so the signal strength is 60 * 19 = 1140.
# During the 100th cycle, register X has the value 18, so the signal strength is 100 * 18 = 1800.
# During the 140th cycle, register X has the value 21, so the signal strength is 140 * 21 = 2940.
# During the 180th cycle, register X has the value 16, so the signal strength is 180 * 16 = 2880.
# During the 220th cycle, register X has the value 18, so the signal strength is 220 * 18 = 3960.
# The sum of these signal strengths is 13140.

# Find the signal strength during the 20th, 60th, 100th, 140th, 180th, and 220th cycles. What is the sum of these six signal strengths?

# (You guessed 19220.) too high

# Your puzzle answer was 17940.

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


res = 0
cycle = 0
value = 1

cycle_milestones = [20, 60, 100, 140, 180, 220]


def check_cycle():
    global res
    if cycle in cycle_milestones:
        if printing_enabled:
            print(cycle, value, cycle * value)
        res += cycle * value


for line in lines:
    line_split = line.split(' ')
    # print(line_split)
    # data.append(Op(*line_split))
    op = Op(*line_split)
    if op.instruction == 'noop':
        cycle += 1
        check_cycle()
        continue
    if op.instruction == 'addx':
        for _ in range(2):
            cycle += 1
            check_cycle()
        value += op.amount
    else:
        print('ERROR: unknown instruction')


print(res)

stop = timeit.default_timer()
print(f'\n✨Time: {stop - start_time} s')
