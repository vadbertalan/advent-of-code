# --- Part Two - --

# Now that you're warmed up, it's time to play the real game.

# A second compartment opens, this time labeled Dirac dice. Out of it falls a single three-sided die.

# As you experiment with the die, you feel a little strange. An informational brochure in the compartment explains that this is a quantum die: when you roll it, the universe splits into multiple copies, one copy for each possible outcome of the die. In this case, rolling the die always splits the universe into three copies: one where the outcome of the roll was 1, one where it was 2, and one where it was 3.

# The game is played the same as before, although to prevent things from getting too far out of hand, the game now ends when either player's score reaches at least 21.

# Using the same starting positions as in the example above, player 1 wins in 444356092776315 universes, while player 2 merely wins in 341960390180808 universes.

# Using your given starting positions, determine every possible outcome. Find the player that wins in more universes
# in how many universes does that player win?

# Your puzzle answer was 265845890886828.

import timeit
from collections import Counter
start_time = timeit.default_timer()


data_str1 = """Player 1 starting position: 4
Player 2 starting position: 8"""

data_str = """Player 1 starting position: 5
Player 2 starting position: 10"""

lines = data_str.split('\n')

w1 = 0
w2 = 0

starting_pos1 = int(lines[0][-2:])
starting_pos2 = int(lines[1][-2:])

possible_amounts = {
    3: 1,
    4: 3,
    5: 6,
    6: 7,
    7: 6,
    8: 3,
    9: 1
}


def play(p1, p2, pos1, pos2, turn, nr_of_universes):
    # print(
    #     f'p - ({p1}, {p2})   pos - ({pos1}, {pos2})   turn - {turn}   universes - {nr_of_universes}')
    global w1, w2, possible_amounts
    if p1 >= 21:
        w1 += nr_of_universes
        return
    if p2 >= 21:
        w2 += nr_of_universes
        return

    new_turn = 1 if turn == 2 else 2

    for amount, count in possible_amounts.items():
        if turn == 1:
            new_pos = pos1 + amount
            new_pos = new_pos if new_pos <= 10 else (new_pos % 10)
            play(p1 + new_pos, p2, new_pos, pos2,
                 new_turn, nr_of_universes * count)
        elif turn == 2:
            new_pos = pos2 + amount
            new_pos = new_pos if new_pos <= 10 else (new_pos % 10)
            play(p1, p2 + new_pos, pos1, new_pos,
                 new_turn, nr_of_universes * count)


play(0, 0, starting_pos1, starting_pos2, 1, 1)

print(f'w1 -> {w1}')
print(f'w2 -> {w2}')

print(max(w1, w2))

stop = timeit.default_timer()
print(f'\nTime: {stop - start_time} s')
