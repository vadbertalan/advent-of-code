# --- Day 5: Supply Stacks ---

# The expedition can depart as soon as the final supplies have been unloaded from the ships. Supplies are stored in stacks of marked crates, but because the needed supplies are buried under many other crates, the crates need to be rearranged.

# The ship has a giant cargo crane capable of moving crates between stacks. To ensure none of the crates get crushed or fall over, the crane operator will rearrange them in a series of carefully-planned steps. After the crates are rearranged, the desired crates will be at the top of each stack.

# The Elves don't want to interrupt the crane operator during this delicate procedure, but they forgot to ask her which crate will end up where, and they want to be ready to unload them as soon as possible so they can embark.

# They do, however, have a drawing of the starting stacks of crates and the rearrangement procedure (your puzzle input). For example:

#     [D]    
# [N] [C]    
# [Z] [M] [P]
#  1   2   3 

# move 1 from 2 to 1
# move 3 from 1 to 3
# move 2 from 2 to 1
# move 1 from 1 to 2
# In this example, there are three stacks of crates. Stack 1 contains two crates: crate Z is on the bottom, and crate N is on top. Stack 2 contains three crates; from bottom to top, they are crates M, C, and D. Finally, stack 3 contains a single crate, P.

# Then, the rearrangement procedure is given. In each step of the procedure, a quantity of crates is moved from one stack to a different stack. In the first step of the above rearrangement procedure, one crate is moved from stack 2 to stack 1, resulting in this configuration:

# [D]        
# [N] [C]    
# [Z] [M] [P]
#  1   2   3 
# In the second step, three crates are moved from stack 1 to stack 3. Crates are moved one at a time, so the first crate to be moved (D) ends up below the second and third crates:

#         [Z]
#         [N]
#     [C] [D]
#     [M] [P]
#  1   2   3
# Then, both crates are moved from stack 2 to stack 1. Again, because crates are moved one at a time, crate C ends up below crate M:

#         [Z]
#         [N]
# [M]     [D]
# [C]     [P]
#  1   2   3
# Finally, one crate is moved from stack 1 to stack 2:

#         [Z]
#         [N]
#         [D]
# [C] [M] [P]
#  1   2   3
# The Elves just need to know which crate will end up on top of each stack; in this example, the top crates are C in stack 1, M in stack 2, and Z in stack 3, so you should combine these together and give the Elves the message CMZ.

# After the rearrangement procedure completes, what crate ends up on top of each stack?

# Your puzzle answer was FWSHSPJWM.

from collections import Counter
import math
import timeit
start_time = timeit.default_timer()

input_stacks1 = """ZN
MCD
P"""
input_moves1 = """move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2"""

input_stacks = """SZPDLBFC
NVGPHWB
FWBJG
GJNFLWCS
WJLTPMSH
BCWGFS
HTPMQBW
FSWT
NCR"""
input_moves = """move 2 from 5 to 9
move 3 from 1 to 7
move 2 from 3 to 9
move 6 from 9 to 5
move 2 from 3 to 8
move 9 from 7 to 8
move 15 from 8 to 9
move 3 from 1 to 6
move 6 from 4 to 2
move 6 from 5 to 6
move 1 from 4 to 2
move 14 from 6 to 2
move 2 from 1 to 5
move 1 from 7 to 3
move 1 from 4 to 8
move 2 from 5 to 6
move 25 from 2 to 4
move 2 from 6 to 4
move 1 from 8 to 1
move 2 from 9 to 1
move 1 from 6 to 1
move 2 from 1 to 7
move 1 from 7 to 3
move 2 from 1 to 8
move 1 from 2 to 6
move 1 from 3 to 8
move 4 from 5 to 6
move 1 from 5 to 3
move 1 from 9 to 6
move 2 from 3 to 4
move 1 from 2 to 6
move 12 from 9 to 7
move 1 from 9 to 1
move 1 from 5 to 8
move 1 from 3 to 8
move 28 from 4 to 5
move 1 from 4 to 3
move 1 from 2 to 6
move 1 from 3 to 9
move 12 from 7 to 2
move 1 from 9 to 6
move 6 from 6 to 4
move 1 from 7 to 4
move 1 from 1 to 2
move 28 from 5 to 1
move 2 from 2 to 8
move 3 from 8 to 2
move 7 from 4 to 1
move 4 from 8 to 6
move 9 from 2 to 8
move 7 from 6 to 5
move 3 from 5 to 9
move 1 from 9 to 7
move 1 from 7 to 1
move 5 from 8 to 4
move 4 from 1 to 9
move 6 from 9 to 4
move 5 from 1 to 5
move 5 from 2 to 3
move 4 from 8 to 2
move 5 from 1 to 4
move 4 from 5 to 9
move 9 from 4 to 9
move 10 from 9 to 8
move 1 from 9 to 1
move 2 from 2 to 8
move 4 from 3 to 8
move 1 from 2 to 3
move 2 from 9 to 2
move 1 from 2 to 6
move 4 from 4 to 3
move 3 from 5 to 1
move 12 from 1 to 4
move 1 from 5 to 3
move 1 from 5 to 3
move 5 from 8 to 5
move 7 from 8 to 5
move 8 from 3 to 4
move 1 from 5 to 1
move 1 from 6 to 7
move 2 from 1 to 6
move 8 from 5 to 9
move 2 from 5 to 1
move 9 from 1 to 4
move 20 from 4 to 2
move 1 from 5 to 2
move 4 from 4 to 2
move 5 from 9 to 2
move 2 from 8 to 9
move 23 from 2 to 4
move 2 from 2 to 5
move 5 from 1 to 2
move 28 from 4 to 3
move 2 from 8 to 1
move 2 from 5 to 7
move 1 from 6 to 9
move 1 from 4 to 8
move 1 from 8 to 9
move 1 from 4 to 6
move 2 from 7 to 2
move 13 from 3 to 4
move 5 from 9 to 7
move 1 from 9 to 6
move 14 from 2 to 6
move 1 from 4 to 1
move 10 from 3 to 2
move 1 from 6 to 9
move 2 from 3 to 2
move 3 from 1 to 9
move 1 from 3 to 5
move 3 from 9 to 3
move 6 from 7 to 4
move 1 from 9 to 4
move 1 from 9 to 2
move 1 from 5 to 3
move 5 from 3 to 1
move 17 from 4 to 7
move 2 from 2 to 8
move 1 from 3 to 9
move 1 from 8 to 2
move 1 from 9 to 6
move 4 from 6 to 2
move 10 from 6 to 5
move 4 from 1 to 5
move 15 from 2 to 9
move 1 from 8 to 6
move 1 from 2 to 8
move 6 from 9 to 2
move 3 from 4 to 8
move 11 from 7 to 1
move 6 from 9 to 6
move 1 from 6 to 2
move 3 from 9 to 3
move 6 from 2 to 7
move 6 from 7 to 8
move 7 from 1 to 9
move 4 from 1 to 6
move 2 from 1 to 2
move 4 from 6 to 7
move 1 from 2 to 9
move 1 from 2 to 3
move 1 from 2 to 1
move 6 from 8 to 4
move 2 from 6 to 7
move 13 from 5 to 9
move 1 from 5 to 4
move 3 from 4 to 7
move 1 from 1 to 7
move 14 from 9 to 2
move 2 from 9 to 3
move 3 from 8 to 5
move 4 from 3 to 4
move 8 from 4 to 1
move 7 from 1 to 9
move 5 from 6 to 9
move 4 from 9 to 2
move 1 from 1 to 9
move 17 from 2 to 4
move 1 from 6 to 3
move 4 from 7 to 5
move 5 from 7 to 5
move 1 from 6 to 4
move 1 from 8 to 3
move 5 from 7 to 1
move 2 from 7 to 6
move 2 from 3 to 6
move 1 from 2 to 9
move 7 from 9 to 6
move 2 from 3 to 7
move 8 from 6 to 4
move 3 from 9 to 2
move 1 from 6 to 4
move 26 from 4 to 8
move 2 from 7 to 8
move 5 from 5 to 9
move 2 from 6 to 7
move 4 from 9 to 1
move 2 from 7 to 5
move 14 from 8 to 6
move 3 from 2 to 8
move 3 from 6 to 8
move 3 from 6 to 1
move 10 from 8 to 4
move 5 from 9 to 4
move 3 from 8 to 5
move 1 from 8 to 2
move 12 from 4 to 8
move 1 from 9 to 3
move 6 from 6 to 4
move 6 from 8 to 2
move 1 from 3 to 8
move 1 from 8 to 4
move 10 from 1 to 9
move 2 from 1 to 3
move 7 from 4 to 9
move 1 from 2 to 1
move 11 from 8 to 9
move 1 from 3 to 9
move 2 from 2 to 7
move 1 from 3 to 6
move 2 from 7 to 9
move 2 from 4 to 6
move 4 from 6 to 4
move 2 from 2 to 8
move 2 from 8 to 4
move 1 from 1 to 7
move 2 from 2 to 8
move 9 from 5 to 2
move 3 from 5 to 9
move 1 from 8 to 3
move 30 from 9 to 7
move 1 from 6 to 2
move 7 from 4 to 8
move 13 from 7 to 2
move 8 from 7 to 4
move 2 from 4 to 8
move 8 from 8 to 1
move 1 from 8 to 3
move 2 from 8 to 9
move 1 from 3 to 7
move 5 from 7 to 6
move 1 from 3 to 1
move 7 from 4 to 8
move 20 from 2 to 6
move 2 from 2 to 7
move 1 from 9 to 5
move 4 from 7 to 6
move 3 from 7 to 8
move 1 from 7 to 2
move 7 from 8 to 6
move 3 from 6 to 7
move 4 from 9 to 1
move 1 from 2 to 6
move 1 from 9 to 7
move 1 from 2 to 8
move 1 from 7 to 6
move 3 from 6 to 3
move 4 from 8 to 1
move 8 from 6 to 4
move 3 from 7 to 2
move 1 from 3 to 2
move 1 from 4 to 5
move 2 from 3 to 5
move 1 from 4 to 6
move 4 from 1 to 5
move 4 from 2 to 9
move 2 from 1 to 6
move 4 from 9 to 2
move 3 from 2 to 8
move 2 from 8 to 4
move 13 from 6 to 1
move 4 from 5 to 2
move 14 from 6 to 3
move 1 from 2 to 7
move 2 from 2 to 4
move 1 from 8 to 6
move 1 from 6 to 3
move 1 from 7 to 4
move 1 from 2 to 3
move 1 from 2 to 6
move 11 from 4 to 6
move 2 from 5 to 4
move 1 from 5 to 6
move 12 from 3 to 6
move 1 from 3 to 7
move 1 from 5 to 7
move 3 from 3 to 6
move 2 from 7 to 5
move 2 from 5 to 2
move 8 from 6 to 7
move 24 from 1 to 3
move 1 from 4 to 6
move 10 from 3 to 1
move 6 from 1 to 8
move 1 from 6 to 3
move 1 from 4 to 2
move 1 from 3 to 1
move 2 from 2 to 1
move 1 from 7 to 6
move 2 from 7 to 5
move 4 from 3 to 7
move 1 from 2 to 3
move 6 from 1 to 6
move 3 from 7 to 5
move 4 from 7 to 8
move 1 from 1 to 2
move 1 from 2 to 7
move 8 from 3 to 4
move 3 from 4 to 7
move 6 from 8 to 6
move 2 from 3 to 2
move 1 from 3 to 9
move 5 from 5 to 1
move 2 from 8 to 2
move 1 from 9 to 2
move 4 from 1 to 3
move 3 from 2 to 9
move 1 from 1 to 2
move 2 from 9 to 7
move 2 from 2 to 9
move 8 from 7 to 5
move 33 from 6 to 5
move 20 from 5 to 9
move 21 from 5 to 7
move 17 from 7 to 6
move 10 from 6 to 9
move 5 from 4 to 7
move 2 from 3 to 9
move 1 from 2 to 3
move 2 from 7 to 3
move 3 from 9 to 5
move 23 from 9 to 7
move 8 from 9 to 6
move 1 from 9 to 1
move 1 from 5 to 3
move 1 from 8 to 9
move 5 from 6 to 8
move 1 from 9 to 6
move 18 from 7 to 2
move 6 from 7 to 4
move 6 from 4 to 8
move 5 from 7 to 4
move 6 from 6 to 3
move 1 from 4 to 2
move 10 from 2 to 1
move 1 from 2 to 4
move 7 from 1 to 6
move 1 from 7 to 1
move 11 from 6 to 2
move 1 from 6 to 8
move 12 from 3 to 1
move 8 from 1 to 8
move 2 from 5 to 2
move 12 from 8 to 6
move 15 from 2 to 4
move 7 from 4 to 5
move 4 from 5 to 9
move 4 from 9 to 4
move 5 from 4 to 6
move 2 from 5 to 2
move 1 from 2 to 5
move 2 from 5 to 4
move 2 from 1 to 3
move 4 from 1 to 5
move 2 from 8 to 4
move 5 from 2 to 9
move 17 from 6 to 8
move 1 from 3 to 2
move 2 from 5 to 4
move 1 from 3 to 8
move 1 from 1 to 6
move 2 from 5 to 6
move 3 from 9 to 5
move 1 from 5 to 1
move 3 from 1 to 8
move 26 from 8 to 4
move 1 from 5 to 3
move 3 from 2 to 7
move 1 from 5 to 7
move 21 from 4 to 9
move 19 from 4 to 5
move 3 from 4 to 3
move 2 from 7 to 5
move 1 from 8 to 2
move 1 from 6 to 2
move 1 from 8 to 9
move 1 from 6 to 7
move 1 from 2 to 4
move 1 from 4 to 7
move 1 from 2 to 7
move 1 from 7 to 1
move 1 from 1 to 6
move 1 from 3 to 5
move 2 from 6 to 3
move 13 from 5 to 8
move 1 from 4 to 2
move 3 from 5 to 4
move 5 from 5 to 4
move 5 from 8 to 9
move 9 from 9 to 3
move 2 from 7 to 1
move 6 from 4 to 2
move 8 from 9 to 4
move 1 from 2 to 7
move 12 from 9 to 8
move 1 from 4 to 2
move 3 from 7 to 3
move 11 from 8 to 5
move 5 from 8 to 6
move 3 from 6 to 5
move 2 from 4 to 1
move 13 from 5 to 3
move 1 from 1 to 7
move 2 from 1 to 8
move 3 from 4 to 9
move 1 from 1 to 7
move 1 from 2 to 4
move 2 from 7 to 3
move 1 from 5 to 3
move 4 from 4 to 2
move 1 from 4 to 9
move 30 from 3 to 2
move 1 from 9 to 7
move 6 from 8 to 6
move 1 from 7 to 6
move 1 from 5 to 1
move 1 from 3 to 5
move 30 from 2 to 3
move 1 from 1 to 9
move 2 from 9 to 2
move 9 from 6 to 9
move 2 from 2 to 9
move 1 from 5 to 1
move 5 from 9 to 7
move 8 from 2 to 5
move 1 from 1 to 9
move 3 from 9 to 1
move 5 from 3 to 6
move 8 from 5 to 9
move 13 from 3 to 9
move 3 from 1 to 7
move 5 from 7 to 9
move 17 from 9 to 6
move 1 from 7 to 6
move 6 from 3 to 9
move 1 from 2 to 1
move 2 from 7 to 1
move 1 from 2 to 5
move 21 from 9 to 2
move 4 from 3 to 6
move 6 from 6 to 5
move 7 from 5 to 9
move 2 from 3 to 8
move 3 from 1 to 3
move 4 from 6 to 5
move 1 from 8 to 1
move 1 from 8 to 2
move 4 from 5 to 2
move 4 from 9 to 1
move 4 from 3 to 5
move 2 from 1 to 7
move 1 from 7 to 4
move 3 from 9 to 5
move 25 from 2 to 9
move 18 from 9 to 1
move 1 from 4 to 5
move 1 from 3 to 8
move 4 from 5 to 6
move 2 from 9 to 3
move 17 from 1 to 5
move 1 from 2 to 7
move 2 from 3 to 5
move 3 from 1 to 8
move 5 from 9 to 2
move 4 from 8 to 9
move 12 from 5 to 2
move 1 from 1 to 8
move 3 from 9 to 5
move 1 from 8 to 2
move 2 from 7 to 2
move 1 from 9 to 5
move 9 from 5 to 2
move 6 from 6 to 2
move 15 from 6 to 2
move 5 from 5 to 9
move 1 from 5 to 9
move 3 from 9 to 2
move 3 from 9 to 1
move 1 from 1 to 9
move 1 from 9 to 1
move 19 from 2 to 8
move 2 from 1 to 9
move 33 from 2 to 6
move 4 from 6 to 4
move 1 from 2 to 6
move 1 from 9 to 8
move 3 from 4 to 8
move 18 from 8 to 3
move 1 from 4 to 9
move 10 from 3 to 9
move 1 from 1 to 4
move 24 from 6 to 3
move 1 from 4 to 3
move 2 from 8 to 7
move 8 from 9 to 3
move 5 from 6 to 7
move 35 from 3 to 2
move 7 from 7 to 1
move 3 from 1 to 3
move 33 from 2 to 6
move 6 from 3 to 7
move 5 from 7 to 3
move 1 from 1 to 4
move 1 from 7 to 8
move 1 from 4 to 8
move 1 from 3 to 2
move 30 from 6 to 5
move 2 from 1 to 6
move 5 from 8 to 1
move 1 from 9 to 2
move 2 from 6 to 4
move 4 from 1 to 7
move 21 from 5 to 8"""

lines_stacks = input_stacks1.split('\n')
lines_moves = input_moves1.split('\n')

stacks = []

# parsing logic is omitted (hard coded values used)
for line in lines_stacks:
	# set = {*()} # empty set, not dictionary
	stack = []
	for letter in line:
		stack.append(letter)
	stacks.append(stack)

moves = []

for line in lines_moves:
	split_line = line.split(' ')
	move = list(map(int, [split_line[1], split_line[3], split_line[5]]))
	moves.append(move)

print(stacks)
print(moves)

for move in moves:
	print(f'move {move[0]} from {move[1]} to {move[2]}')
	for i in range(move[0]):
		popped = stacks[move[1] - 1].pop()
		stacks[move[2] - 1].append(popped)
	print(stacks)

print(stacks)

res = ''
for stack in stacks:
	res += stack.pop()

print(res)

stop = timeit.default_timer()
print(f'\n✨Time: {stop - start_time} s')
