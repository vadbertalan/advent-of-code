# --- Part Two ---

# As you watch the crane operator expertly rearrange the crates, you notice the process isn't following your prediction.

# Some mud was covering the writing on the side of the crane, and you quickly wipe it away. The crane isn't a CrateMover 9000 - it's a CrateMover 9001.

# The CrateMover 9001 is notable for many new and exciting features: air conditioning, leather seats, an extra cup holder, and the ability to pick up and move multiple crates at once.

# Again considering the example above, the crates begin in the same configuration:

#     [D]    
# [N] [C]    
# [Z] [M] [P]
#  1   2   3 
# Moving a single crate from stack 2 to stack 1 behaves the same as before:

# [D]        
# [N] [C]    
# [Z] [M] [P]
#  1   2   3 
# However, the action of moving three crates from stack 1 to stack 3 means that those three moved crates stay in the same order, resulting in this new configuration:

#         [D]
#         [N]
#     [C] [Z]
#     [M] [P]
#  1   2   3
# Next, as both crates are moved from stack 2 to stack 1, they retain their order as well:

#         [D]
#         [N]
# [C]     [Z]
# [M]     [P]
#  1   2   3
# Finally, a single crate is still moved from stack 1 to stack 2, but now it's crate C that gets moved:

#         [D]
#         [N]
#         [Z]
# [M] [C] [P]
#  1   2   3
# In this example, the CrateMover 9001 has put the crates in a totally different order: MCD.

# Before the rearrangement process finishes, update your simulation so that the Elves know where they should stand to be ready to unload the final supplies. After the rearrangement procedure completes, what crate ends up on top of each stack?

# Your puzzle answer was PWPWHGFZS.

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

lines_stacks = input_stacks.split('\n')
lines_moves = input_moves.split('\n')

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
	popped = stacks[move[1] - 1][len(stacks[move[1] - 1]) - move[0]:]
	stacks[move[1] - 1] = stacks[move[1] - 1][:len(stacks[move[1] - 1]) - move[0]]
	stacks[move[2] - 1].extend(popped)
	print(stacks)

print(stacks)

res = ''
for stack in stacks:
	res += stack.pop()

print(res)

stop = timeit.default_timer()
print(f'\n✨Time: {stop - start_time} s')
