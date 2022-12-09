# --- Part Two - --

# A rope snaps! Suddenly, the river is getting a lot closer than you remember. The bridge is still there, but some of the ropes that broke are now whipping toward you as you fall through the air!

# The ropes are moving too quickly to grab
# you only have a few seconds to choose how to arch your body to avoid being hit. Fortunately, your simulation can be extended to support longer ropes.

# Rather than two knots, you now must simulate a rope consisting of ten knots. One knot is still the head of the rope and moves according to the series of motions. Each knot further down the rope follows the knot in front of it using the same rules as before.

# Using the same series of motions as the above example, but with the knots marked H, 1, 2, ..., 9, the motions now occur as follows:

# == Initial State ==

# ......
# ......
# ......
# ......
# H.....  (H covers 1, 2, 3, 4, 5, 6, 7, 8, 9, s)

# == R 4 ==

# ......
# ......
# ......
# ......
# 1H....  (1 covers 2, 3, 4, 5, 6, 7, 8, 9, s)

# ......
# ......
# ......
# ......
# 21H...  (2 covers 3, 4, 5, 6, 7, 8, 9, s)

# ......
# ......
# ......
# ......
# 321H..  (3 covers 4, 5, 6, 7, 8, 9, s)

# ......
# ......
# ......
# ......
# 4321H.  (4 covers 5, 6, 7, 8, 9, s)

# == U 4 ==

# ......
# ......
# ......
# ....H.
# 4321..  (4 covers 5, 6, 7, 8, 9, s)

# ......
# ......
# ....H.
# .4321.
# 5.....  (5 covers 6, 7, 8, 9, s)

# ......
# ....H.
# ....1.
# .432..
# 5.....  (5 covers 6, 7, 8, 9, s)

# ....H.
# ....1.
# ..432.
# .5....
# 6.....  (6 covers 7, 8, 9, s)

# == L 3 ==

# ...H..
# ....1.
# ..432.
# .5....
# 6.....  (6 covers 7, 8, 9, s)

# ..H1..
# ...2..
# ..43..
# .5....
# 6.....  (6 covers 7, 8, 9, s)

# .H1...
# ...2..
# ..43..
# .5....
# 6.....  (6 covers 7, 8, 9, s)

# == D 1 ==

# ..1...
# .H.2..
# ..43..
# .5....
# 6.....  (6 covers 7, 8, 9, s)

# == R 4 ==

# ..1...
# ..H2..
# ..43..
# .5....
# 6.....  (6 covers 7, 8, 9, s)

# ..1...
# ...H..  (H covers 2)
# ..43..
# .5....
# 6.....  (6 covers 7, 8, 9, s)

# ......
# ...1H.  (1 covers 2)
# ..43..
# .5....
# 6.....  (6 covers 7, 8, 9, s)

# ......
# ...21H
# ..43..
# .5....
# 6.....  (6 covers 7, 8, 9, s)

# == D 1 ==

# ......
# ...21.
# ..43.H
# .5....
# 6.....  (6 covers 7, 8, 9, s)

# == L 5 ==

# ......
# ...21.
# ..43H.
# .5....
# 6.....  (6 covers 7, 8, 9, s)

# ......
# ...21.
# ..4H..  (H covers 3)
# .5....
# 6.....  (6 covers 7, 8, 9, s)

# ......
# ...2..
# ..H1..  (H covers 4
#          1 covers 3)
# .5....
# 6.....  (6 covers 7, 8, 9, s)

# ......
# ...2..
# .H13..  (1 covers 4)
# .5....
# 6.....  (6 covers 7, 8, 9, s)

# ......
# ......
# H123..  (2 covers 4)
# .5....
# 6.....  (6 covers 7, 8, 9, s)

# == R 2 ==

# ......
# ......
# .H23..  (H covers 1
#          2 covers 4)
# .5....
# 6.....  (6 covers 7, 8, 9, s)

# ......
# ......
# .1H3..  (H covers 2, 4)
# .5....
# 6.....  (6 covers 7, 8, 9, s)
# Now, you need to keep track of the positions the new tail, 9, visits. In this example, the tail never moves, and so it only visits 1 position. However, be careful: more types of motion are possible than before, so you might want to visually compare your simulated rope to the one above.

# Here's a larger example:

# R 5
# U 8
# L 8
# D 3
# R 17
# D 10
# L 25
# U 20
# These motions occur as follows(individual steps are not shown):

# == Initial State ==

# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ...........H..............  (H covers 1, 2, 3, 4, 5, 6, 7, 8, 9, s)
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................

# == R 5 ==

# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ...........54321H.........  (5 covers 6, 7, 8, 9, s)
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................

# == U 8 ==

# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ................H.........
# ................1.........
# ................2.........
# ................3.........
# ...............54.........
# ..............6...........
# .............7............
# ............8.............
# ...........9..............  (9 covers s)
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................

# == L 8 ==

# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ........H1234.............
# ............5.............
# ............6.............
# ............7.............
# ............8.............
# ............9.............
# ..........................
# ..........................
# ...........s..............
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................

# == D 3 ==

# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# .........2345.............
# ........1...6.............
# ........H...7.............
# ............8.............
# ............9.............
# ..........................
# ..........................
# ...........s..............
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................

# == R 17 ==

# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ................987654321H
# ..........................
# ..........................
# ..........................
# ..........................
# ...........s..............
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................

# == D 10 ==

# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ...........s.........98765
# .........................4
# .........................3
# .........................2
# .........................1
# .........................H

# == L 25 ==

# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ...........s..............
# ..........................
# ..........................
# ..........................
# ..........................
# H123456789................

# == U 20 ==

# H.........................
# 1.........................
# 2.........................
# 3.........................
# 4.........................
# 5.........................
# 6.........................
# 7.........................
# 8.........................
# 9.........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ...........s..............
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................

# Now, the tail(9) visits 36 positions(including s) at least once:

# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# ..........................
# # .........................
# # .............###.........
# # ............#...#........
# .  # ..........#.....#.......
# ..  # ..........#.....#......
# ...  # ........#.......#.....
# ....  # ......s.........#....
# .....  # ..............#.....
# ......  # ............#......
# .......  # ..........#.......
# ........  # ........#........
# .........  # .........
# Simulate your complete series of motions on a larger rope with ten knots. How many positions does the tail of the rope visit at least once?

# (You guessed 2425.)

# Your puzzle answer was 2443.

from collections import Counter
import math
import sys
import timeit
start_time = timeit.default_timer()

printing_enabled = sys.argv[1] == '-v' if len(sys.argv) == 2 else False

data_str1 = """R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2"""

data_str2 = """R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20"""

data_str = """D 2
R 2
D 1
U 1
L 2
U 2
D 1
R 2
U 2
R 2
L 1
D 2
L 1
U 1
R 1
D 1
U 1
R 1
U 1
R 1
D 2
L 1
D 1
L 1
D 1
U 1
L 1
R 1
L 2
D 1
L 2
R 1
U 2
L 2
D 1
U 2
D 1
L 1
U 2
L 2
U 2
D 2
U 2
D 2
L 1
D 2
L 2
D 2
U 2
L 2
R 2
U 1
L 1
D 2
R 1
D 2
L 2
R 1
D 2
R 1
U 2
L 2
U 2
D 1
U 2
L 2
U 2
L 1
U 2
L 1
U 2
D 1
U 1
D 2
R 2
L 1
U 2
L 1
D 1
R 2
U 1
L 1
U 1
R 2
U 2
D 2
L 2
U 2
L 2
R 1
L 2
U 1
R 1
L 1
R 2
U 1
R 1
D 1
L 2
R 2
L 2
R 1
D 2
R 1
U 2
L 2
D 1
R 2
D 2
L 2
U 2
D 2
L 1
D 2
R 1
L 3
D 2
R 1
L 3
R 2
D 3
U 3
R 2
L 1
R 2
U 2
R 2
L 1
U 2
D 1
L 3
U 3
R 3
D 2
U 1
L 2
R 2
D 1
L 2
U 3
R 1
L 3
R 1
L 3
U 3
D 3
U 2
R 2
L 3
U 2
L 2
U 2
R 2
D 1
R 3
L 2
R 3
D 2
R 2
L 2
R 3
L 2
R 1
U 2
D 1
R 1
D 2
U 2
L 2
D 1
R 1
L 3
U 1
R 3
L 2
D 1
U 1
L 2
U 1
D 2
R 1
U 2
D 3
R 3
L 2
R 2
U 2
L 3
U 3
L 3
R 2
U 1
D 1
U 2
R 3
D 3
L 1
D 2
U 1
D 2
R 3
U 2
D 1
L 1
D 3
R 1
L 1
U 2
L 3
D 1
U 2
R 1
D 1
R 3
D 1
R 1
L 1
R 2
L 1
D 1
R 2
U 1
L 4
U 2
D 2
L 4
R 1
L 3
D 3
R 3
L 4
D 1
R 1
U 4
L 1
D 3
R 4
U 3
L 3
U 4
D 3
L 2
D 2
R 1
D 3
U 1
R 4
U 2
D 1
L 4
U 4
L 1
U 1
L 3
D 2
L 2
R 3
U 3
D 3
U 2
R 3
D 2
L 1
U 2
D 3
R 1
D 2
U 4
R 4
U 3
R 3
L 1
U 4
R 4
L 1
R 4
L 4
R 1
L 1
R 3
U 2
R 3
L 3
D 2
R 1
D 2
U 2
R 4
U 2
R 4
L 1
R 3
U 4
L 3
R 2
U 1
R 2
D 4
L 1
U 4
R 2
L 3
U 2
L 3
U 1
D 1
L 4
U 2
L 2
U 4
D 2
U 3
R 4
L 1
D 1
R 2
D 4
U 3
R 1
L 2
R 4
D 4
L 3
D 1
R 2
L 3
U 1
L 2
D 3
L 2
R 1
L 2
R 2
L 4
U 1
L 5
R 2
D 5
L 4
R 2
D 3
L 5
U 4
R 5
L 1
D 4
R 2
D 4
U 3
R 1
D 1
U 5
R 2
D 2
R 5
D 3
L 1
D 4
R 2
L 5
U 4
R 1
D 1
U 1
R 1
D 3
U 1
L 4
U 2
D 2
R 1
U 2
L 3
R 5
D 1
L 4
R 4
D 5
L 5
D 5
U 1
L 2
D 3
L 3
D 1
U 2
D 3
U 1
L 3
R 1
D 2
U 4
R 4
D 2
U 1
L 1
R 1
U 1
R 4
D 4
L 2
U 4
D 1
R 4
L 2
U 3
D 4
R 2
U 1
R 2
U 5
D 2
U 4
L 2
R 2
L 1
U 2
R 5
U 5
R 3
L 2
R 4
L 3
R 3
L 2
U 4
L 2
U 2
D 3
U 3
L 3
R 2
U 2
D 5
U 4
R 3
L 5
R 1
U 4
L 4
U 4
D 1
U 2
D 4
R 4
U 5
R 2
L 2
R 2
L 3
R 1
D 5
R 5
L 4
U 2
D 6
U 3
R 6
L 1
D 3
U 2
R 1
D 3
L 6
R 5
L 6
U 5
R 4
D 1
U 2
R 3
U 1
R 6
D 1
L 4
R 2
L 1
U 6
L 1
U 3
R 3
U 4
D 5
R 2
U 2
D 3
R 1
U 5
D 2
L 3
U 3
D 2
U 4
D 4
R 3
L 3
R 3
D 1
R 2
L 4
D 2
L 5
U 5
D 2
L 5
U 1
L 2
U 2
L 5
R 1
L 2
U 2
R 3
U 5
R 1
L 2
R 3
D 3
R 3
L 6
D 1
L 2
R 1
U 4
D 3
R 6
L 4
U 2
L 3
D 4
L 3
D 5
L 3
D 4
U 3
D 4
U 2
D 6
U 4
R 2
L 2
R 1
L 4
R 6
D 5
L 2
D 6
L 4
R 4
U 6
L 6
D 6
R 4
L 2
R 2
U 5
L 6
D 7
U 6
R 1
D 4
R 2
D 7
L 4
U 7
D 4
L 6
D 7
U 6
R 1
U 1
R 4
L 7
U 1
D 2
U 2
R 3
L 6
R 7
L 5
D 5
U 2
D 1
U 4
R 6
D 5
R 6
U 5
R 7
U 7
R 4
L 1
D 1
U 3
D 1
L 5
U 4
R 2
L 3
D 3
U 2
R 2
D 1
R 6
U 5
D 5
R 7
U 2
L 4
D 3
R 2
L 1
D 1
R 2
D 7
U 6
D 1
R 2
D 7
U 5
R 6
D 1
L 5
U 3
L 6
D 3
L 5
D 2
L 5
R 1
D 7
L 3
U 3
R 6
D 4
R 2
U 1
L 2
R 1
L 3
D 3
L 6
U 7
R 4
L 4
U 5
L 7
D 1
R 4
U 6
D 3
R 2
L 1
D 4
R 5
D 3
L 6
R 6
U 4
R 7
D 2
U 3
L 2
D 3
L 5
D 1
U 5
D 8
L 8
U 2
L 8
U 2
D 3
U 1
L 1
R 3
D 5
U 1
D 6
R 6
L 6
D 3
R 3
D 8
R 3
U 8
L 5
U 8
R 4
L 3
U 6
D 5
R 6
D 5
U 2
D 1
L 4
D 5
U 5
R 3
L 4
U 5
D 2
R 4
U 3
R 7
D 2
R 8
L 4
R 6
D 6
U 3
R 6
L 2
R 5
L 2
D 1
U 6
R 7
D 1
L 7
R 8
D 6
U 2
R 5
U 4
L 6
R 5
U 3
R 7
D 3
U 1
R 1
L 2
U 4
L 4
R 5
D 2
U 8
D 4
R 6
L 1
R 3
U 3
L 7
D 8
U 8
D 3
U 2
R 1
D 6
R 5
L 8
U 7
L 8
D 7
U 5
L 6
U 3
R 3
U 8
D 1
U 3
D 6
R 7
U 6
R 4
U 7
L 8
D 4
L 5
R 3
L 1
D 1
R 6
L 4
D 8
R 1
L 6
D 8
L 2
U 6
D 9
L 4
R 6
D 8
U 4
D 5
L 2
R 9
L 3
D 1
R 4
L 8
R 7
U 9
R 1
D 7
R 5
L 4
R 3
L 5
U 4
R 9
U 7
D 8
L 5
U 7
R 2
D 1
U 6
R 1
L 5
D 3
U 3
L 1
D 5
R 3
L 8
D 7
U 3
D 7
R 2
U 3
R 1
D 1
R 3
L 1
D 3
L 3
R 3
D 3
L 9
U 6
L 6
R 6
U 2
R 1
L 6
D 4
R 2
U 9
R 1
L 3
R 7
U 6
R 8
U 6
R 1
U 5
D 5
R 6
L 6
U 5
L 1
D 4
U 3
D 5
R 1
U 9
D 1
R 1
L 7
R 7
L 5
R 1
U 6
R 4
D 8
R 2
L 2
D 1
U 1
R 7
D 6
L 1
R 8
U 2
R 4
U 8
R 1
D 2
U 3
R 5
U 8
R 1
L 6
U 4
L 4
R 1
D 1
L 3
R 8
L 2
D 9
U 4
L 3
U 9
L 9
D 6
U 8
R 3
L 2
U 4
L 1
D 8
L 8
R 5
L 1
U 3
L 6
D 5
U 9
D 4
U 2
D 2
U 1
R 6
D 1
R 1
D 8
L 3
D 10
R 5
D 6
L 1
U 4
D 6
R 6
U 4
L 10
D 1
L 3
R 9
U 5
L 5
R 1
U 7
R 6
L 10
R 1
U 10
D 3
R 1
L 1
R 7
U 4
R 3
U 7
R 2
D 4
R 2
L 2
R 9
L 8
U 10
L 9
D 4
L 2
D 9
L 10
D 2
U 5
L 7
R 1
D 1
L 5
D 8
L 3
R 4
D 5
R 4
U 4
L 8
D 9
L 10
R 5
L 5
R 8
D 5
L 3
R 2
L 2
R 1
L 5
U 7
D 7
L 10
D 1
L 9
D 7
R 1
U 7
R 10
L 7
U 8
D 8
L 9
R 9
D 3
L 1
D 6
U 11
D 5
U 5
L 6
R 1
U 11
D 3
U 7
L 7
R 9
L 1
R 6
L 2
U 6
D 7
L 8
U 10
R 2
D 9
L 11
R 7
L 3
U 7
D 3
R 1
D 6
L 5
D 4
R 7
L 8
U 11
L 10
D 2
L 3
R 4
D 1
R 6
L 1
R 7
L 3
R 4
U 5
L 4
U 8
D 2
R 3
U 7
R 8
L 2
R 5
U 9
R 1
U 2
R 8
L 4
R 1
L 10
U 6
L 3
R 7
U 1
L 8
U 2
D 1
L 5
R 11
U 5
L 7
U 6
L 7
U 4
D 10
R 11
L 2
R 8
D 1
R 8
D 1
U 6
L 7
R 6
U 3
R 8
D 10
L 8
D 3
R 1
U 5
L 3
U 7
R 7
D 10
L 7
U 11
L 9
D 2
L 6
R 6
L 6
U 11
D 4
L 10
R 11
D 11
L 1
D 1
U 7
R 10
U 3
L 5
R 9
U 1
L 12
U 9
R 8
L 12
U 9
L 11
R 11
D 11
R 12
L 9
D 2
U 9
L 1
R 5
D 2
U 10
R 2
U 9
D 6
R 4
L 3
U 4
L 4
D 8
U 9
R 3
D 7
L 11
U 11
D 5
R 4
L 5
D 11
U 8
L 5
D 2
L 5
D 1
R 9
L 1
U 1
D 11
R 2
D 10
L 6
D 8
U 12
D 9
L 3
R 2
U 3
D 2
L 6
D 7
U 1
L 4
U 10
D 12
R 9
D 7
U 6
L 8
D 5
R 4
U 6
L 4
R 5
D 4
U 3
L 10
D 7
U 1
L 2
R 8
U 10
L 11
R 9
D 5
U 10
L 10
R 10
L 5
R 3
D 12
U 11
R 1
U 2
L 6
D 9
U 5
L 10
D 9
R 9
D 10
L 12
R 3
D 3
L 7
D 4
U 5
L 2
D 1
R 11
U 8
L 7
D 10
L 5
U 8
R 4
D 12
L 10
U 1
R 11
D 4
L 12
D 7
R 1
D 12
U 3
R 2
D 3
U 9
D 9
R 2
D 3
L 12
D 10
U 3
R 8
L 9
R 6
L 11
R 3
L 7
R 10
D 3
R 2
D 3
U 8
L 7
U 2
D 4
R 6
U 8
R 4
L 5
U 6
D 1
U 1
R 4
D 5
U 3
D 13
U 4
D 8
U 3
L 12
U 7
D 7
L 2
R 13
L 11
U 8
L 7
R 11
U 13
L 5
R 10
D 2
R 5
D 11
R 8
U 9
D 13
R 4
L 4
D 1
R 6
L 4
U 4
D 5
L 13
D 10
R 10
D 1
U 8
D 8
L 9
U 12
R 9
L 10
R 1
U 12
L 3
D 10
L 9
R 12
D 3
U 1
D 13
L 6
D 13
L 13
D 12
L 7
R 2
U 7
D 7
R 7
U 12
R 6
D 3
R 12
L 13
D 12
U 6
D 1
R 2
U 6
L 12
R 7
D 2
U 13
D 2
U 8
R 10
D 13
L 1
R 10
D 6
U 8
R 13
U 4
R 2
L 1
D 5
R 6
U 13
R 3
L 4
D 14
L 8
D 11
U 13
D 7
L 1
U 2
L 13
R 9
L 10
R 8
L 4
U 12
D 9
U 8
D 3
U 14
D 1
L 7
R 1
D 1
R 3
L 10
R 6
L 6
D 14
U 9
D 9
R 11
D 14
L 4
R 9
L 4
D 1
U 7
D 14
U 14
L 2
D 1
L 14
D 11
R 9
L 11
D 2
R 8
L 2
R 2
U 13
D 1
L 4
U 6
D 6
R 1
D 10
R 9
L 9
R 1
D 13
U 4
L 7
D 7
R 2
L 2
U 11
R 8
U 10
D 9
R 6
U 8
D 9
U 3
R 2
D 8
U 4
R 7
L 10
R 3
U 7
D 10
R 4
U 6
L 9
D 13
L 6
U 11
D 12
U 10
L 14
U 9
D 3
L 6
R 5
U 2
R 15
D 1
R 8
D 3
U 10
D 5
U 3
R 6
D 8
U 15
L 13
D 6
L 13
U 5
D 5
L 6
D 7
U 12
L 14
D 14
U 14
R 6
U 9
R 9
L 15
U 6
R 8
U 6
D 6
L 3
D 10
L 9
D 11
L 12
D 9
U 12
D 6
L 10
R 6
L 10
U 9
D 4
R 3
D 13
R 6
D 6
U 8
D 2
L 2
R 12
D 4
L 14
U 4
L 9
U 10
R 13
D 13
L 9
D 3
U 15
D 14
L 7
U 2
D 5
L 14
U 14
D 4
L 12
U 4
R 3
U 7
D 5
U 8
D 6
R 6
D 13
L 6
R 6
U 8
L 6
U 6
D 7
U 7
R 5
U 15
D 11
L 12
R 4
L 9
U 14
R 6
L 2
U 5
L 5
U 6
L 5
U 9
R 11
D 11
L 2
U 12
L 6
D 6
U 5
L 7
U 3
L 10
D 12
R 16
L 8
U 7
D 11
U 5
R 5
D 10
R 11
D 15
U 8
D 14
U 16
R 7
D 9
U 9
D 14
U 4
D 11
U 5
R 7
D 11
R 4
U 4
R 11
D 3
U 9
D 11
R 13
U 9
D 2
L 6
D 12
R 16
L 7
U 13
R 14
L 7
U 11
R 14
L 11
D 14
R 2
L 14
U 12
D 4
L 2
D 8
L 5
D 7
U 3
R 10
D 12
R 4
L 16
U 1
R 9
U 14
L 15
D 3
L 5
U 10
L 6
U 9
D 10
R 3
L 13
R 8
D 6
R 9
U 14
L 14
D 8
R 7
D 4
U 1
L 4
U 13
R 12
L 2
D 1
U 2
R 5
L 13
D 12
R 10
L 1
D 3
L 4
R 1
D 1
R 12
L 3
U 4
L 12
R 14
L 7
R 7
D 12
U 1
D 15
U 6
R 2
U 6
R 8
L 3
U 8
D 5
R 11
U 12
R 14
U 8
R 14
U 4
R 15
D 17
U 8
L 7
U 2
L 10
D 11
U 12
R 10
L 11
U 2
L 10
D 11
L 3
R 12
D 5
U 3
L 12
D 12
L 2
U 2
R 14
U 2
D 5
R 6
L 14
D 8
R 10
D 14
L 14
U 5
D 8
L 1
U 3
D 8
L 4
U 4
R 16
D 2
L 14
R 10
L 13
D 8
U 12
L 4
R 14
L 16
D 12
L 15
R 8
U 2
D 8
L 14
U 5
D 7
U 17
L 10
R 8
D 11
R 6
D 6
U 10
L 16
D 3
U 10
L 1
R 14
D 6
R 10
D 1
R 4
U 16
L 17
U 3
D 11
U 12
D 8
U 16
L 7
U 10
L 2
U 14
D 15
L 4
U 1
R 13
L 12
D 15
U 6
R 17
L 15
R 15
D 16
U 12
D 4
L 4
R 7
D 13
L 11
U 7
L 4
R 10
D 15
L 15
U 14
L 5
U 2
L 10
R 17
U 7
D 11
L 11
D 4
R 15
D 6
U 7
R 7
U 10
L 5
U 15
L 2
U 1
D 6
U 11
L 13
D 17
R 13
U 6
R 12
D 4
L 15
R 11
U 14
R 17
D 18
L 16
D 9
U 9
R 8
L 17
U 9
R 8
D 5
U 7
L 5
R 18
L 17
U 13
L 14
R 2
D 15
U 1
R 10
D 13
R 14
D 7
U 13
L 18
U 6
R 13
U 16
L 2
R 1
L 17
R 9
U 3
R 7
D 10
L 3
D 15
L 13
R 4
D 16
U 7
D 6
R 8
U 3
L 4
U 14
D 3
L 17
R 3
L 18
D 7
L 1
U 8
D 1
U 9
L 17
U 6
L 5
R 10
L 18
U 2
D 18
U 1
D 13
U 11
R 9
U 17
D 14
U 5
R 3
U 11
D 1
L 2
R 5
U 17
R 12
U 5
D 4
U 14
D 6
U 7
L 18
D 4
R 17
D 6
U 8
R 8
L 5
R 6
L 4
R 1
U 9
R 17
U 11
L 5
R 17
L 11
D 9
L 11
D 16
R 11
D 17
L 13
D 11
L 18
R 13
L 1
D 19
L 1
U 15
D 19
U 13
L 18
U 19
R 8
U 9
L 17
R 6
L 12
R 19
D 15
R 15
U 10
D 3
L 10
D 18
U 18
R 6
U 13
R 6
D 17
R 12
D 16
R 3
D 17
U 12
D 9
R 6
U 12
L 10
U 8
D 4
U 13
R 1
U 4
L 11
D 1
U 13
D 13
L 12
D 6
U 18
L 11
D 16
R 5
D 2
U 19
R 12
U 18
L 5
R 12
L 3
U 9
R 13
L 1
U 10
R 16
L 6
U 11
L 3
R 6
U 5
R 18
U 8
R 13
D 19
R 14
U 14
R 3
D 18
L 18
U 5
L 8
U 13
R 11
U 12
D 19
L 10
R 13
D 18
R 16
L 14
D 13
U 2
D 7
U 3
L 8
R 15
D 19"""

lines = data_str.split('\n')

data = []
N = len(lines)


class Coord():
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __str__(self):
        return f'({self.x}, {self.y})'

    def __repr__(self):
        return f'({self.x}, {self.y})'

    def get_tuple(self):
        return self.x, self.y


class Move():
    def __init__(self, dir, amount):
        self.dir = dir
        self.amount = amount


for line in lines:
    line_split = line.split(' ')
    data.append(Move(line_split[0], int(line_split[1])))


head_coords = Coord(0, 0)
tail_coords = [Coord(0, 0), Coord(0, 0), Coord(
    0, 0), Coord(0, 0), Coord(0, 0), Coord(0, 0), Coord(0, 0), Coord(0, 0), Coord(0, 0)]

tail_recorded_coords_set = set()  # or could have been = {*()}
tail_recorded_coords_set.add(tail_coords[8].get_tuple())


def print_matrix():
    if not printing_enabled:
        return
    for i in range(-15, 15):
        for j in range(-15, 15):
            if (-i, j) == head_coords.get_tuple():
                print('H', end=' ')
                continue

            found_tail = False
            for (idx, tailI) in enumerate(tail_coords):
                if (-i, j) == tailI.get_tuple():
                    print(idx + 1, end=' ')
                    found_tail = True
                    break

            if not found_tail:
                print('.', end=' ')
        print()
    print()


def calc_manhattan_distance(point1: Coord, point2: Coord):
    return abs(point1.x - point2.x) + abs(point1.y - point2.y)


def calc_euclidean_distance(point1: Coord, point2: Coord):
    return math.sqrt(abs(point1.x - point2.x) ^ 2 + abs(point1.y - point2.y) ^ 2)


def get_closest_position_from_ref_to_subject(ref: Coord, subject: Coord) -> Coord:
    candidates = [Coord(subject.x + x, subject.y + y)
                  for (x, y) in [(+1, 0), (0, +1), (-1, 0), (0, -1)]]
    candidates_sorted_by_distances = sorted(
        candidates, key=lambda point: calc_euclidean_distance(ref, point))
    return candidates_sorted_by_distances[0]


def get_closest_position_from_ref_to_subject_diagonally(ref: Coord, subject: Coord) -> Coord:
    candidates = [Coord(subject.x + x, subject.y + y)
                  for (x, y) in [(+1, +1), (-1, +1), (-1, -1), (+1, -1)]]
    candidates_sorted_by_distances = sorted(
        candidates, key=lambda point: calc_euclidean_distance(ref, point))
    return candidates_sorted_by_distances[0]


# def adjust_tail_position_after_up(head_idx, tail_idx) -> bool:
#     if abs(tail_coords[head_idx].x - tail_coords[tail_idx].x) > 1:
#         if tail_coords[head_idx].y != tail_coords[tail_idx].y:
#             tail_coords[tail_idx].y = tail_coords[head_idx].y
#         tail_coords[tail_idx].x += 1
#         if tail_idx == 8:
#             tail_recorded_coords_set.add(tail_coords[8].get_tuple())
#         return True
#     return False


# def adjust_tail_position_after_right(head_idx, tail_idx) -> bool:
#     if abs(tail_coords[head_idx].y - tail_coords[tail_idx].y) > 1:
#         if tail_coords[head_idx].x != tail_coords[tail_idx].x:
#             tail_coords[tail_idx].x = tail_coords[head_idx].x
#         tail_coords[tail_idx].y += 1
#         if tail_idx == 8:
#             tail_recorded_coords_set.add(tail_coords[8].get_tuple())
#         return True
#     return False


# def adjust_tail_position_after_down(head_idx, tail_idx) -> bool:
#     if abs(tail_coords[head_idx].x - tail_coords[tail_idx].x) > 1:
#         if tail_coords[head_idx].y != tail_coords[tail_idx].y:
#             tail_coords[tail_idx].y = tail_coords[head_idx].y
#         tail_coords[tail_idx].x -= 1
#         if tail_idx == 8:
#             tail_recorded_coords_set.add(tail_coords[8].get_tuple())
#         return True
#     return False


# def adjust_tail_position_after_left(head_idx, tail_idx) -> bool:
#     if abs(tail_coords[head_idx].y - tail_coords[tail_idx].y) > 1:
#         if tail_coords[head_idx].x != tail_coords[tail_idx].x:
#             tail_coords[tail_idx].x = tail_coords[head_idx].x
#         tail_coords[tail_idx].y -= 1
#         if tail_idx == 8:
#             tail_recorded_coords_set.add(tail_coords[8].get_tuple())
#         return True
#     return False


def adjust_tail_position(head_idx, tail_idx):
    if abs(tail_coords[head_idx].x - tail_coords[tail_idx].x) == 2 and abs(tail_coords[head_idx].y - tail_coords[tail_idx].y) == 2:
        tail_coords[tail_idx] = get_closest_position_from_ref_to_subject_diagonally(tail_coords[tail_idx], tail_coords[head_idx])
        if tail_idx == 8:
            tail_recorded_coords_set.add(tail_coords[8].get_tuple())
        return True
    if abs(tail_coords[head_idx].x - tail_coords[tail_idx].x) > 1 or abs(tail_coords[head_idx].y - tail_coords[tail_idx].y) > 1:
        tail_coords[tail_idx] = get_closest_position_from_ref_to_subject(tail_coords[tail_idx], tail_coords[head_idx])
        if tail_idx == 8:
            tail_recorded_coords_set.add(tail_coords[8].get_tuple())
        return True
        
    return False


def move_up_amount(amount):
    for _ in range(amount):
        head_coords.x += 1
        if abs(head_coords.x - tail_coords[0].x) > 1:
            if head_coords.y != tail_coords[0].y:
                tail_coords[0].y = head_coords.y
            tail_coords[0].x += 1

            for i in range(1, 9):
                was_change = adjust_tail_position(i - 1, i)
                if not was_change:
                    break
        print_matrix()

def move_right_amount(amount):
    for _ in range(amount):
        head_coords.y += 1
        if abs(head_coords.y - tail_coords[0].y) > 1:
            if head_coords.x != tail_coords[0].x:
                tail_coords[0].x = head_coords.x
            tail_coords[0].y += 1

            for i in range(1, 9):
                was_change = adjust_tail_position(i - 1, i)
                if not was_change:
                    break
        print_matrix()

def move_down_amount(amount):
    for _ in range(amount):
        head_coords.x -= 1
        if abs(head_coords.x - tail_coords[0].x) > 1:
            if head_coords.y != tail_coords[0].y:
                tail_coords[0].y = head_coords.y
            tail_coords[0].x -= 1

            for i in range(1, 9):
                was_change = adjust_tail_position(i - 1, i)
                if not was_change:
                    break
        print_matrix()

def move_left_amount(amount):
    for _ in range(amount):
        head_coords.y -= 1
        if abs(head_coords.y - tail_coords[0].y) > 1:
            if head_coords.x != tail_coords[0].x:
                tail_coords[0].x = head_coords.x
            tail_coords[0].y -= 1

            for i in range(1, 9):
                was_change = adjust_tail_position(i - 1, i)
                if not was_change:
                    break
        print_matrix()

for move in data:
    print(f'Moving {move.dir} {move.amount}')
    if move.dir == 'U':
        move_up_amount(move.amount)
    elif move.dir == 'R':
        move_right_amount(move.amount)
    elif move.dir == 'D':
        move_down_amount(move.amount)
    elif move.dir == 'L':
        move_left_amount(move.amount)
    else:
        print('ERROR: Unknown direction')

    print(head_coords)
    print(tail_coords)
    print()

print(len(tail_recorded_coords_set))

for i in range(-15, 15):
    for j in range(-15, 15):
        if (-i, j) in tail_recorded_coords_set:
            print('#', end=' ')
        else:
            print('.', end=' ')
    print()

stop = timeit.default_timer()
print(f'\n✨Time: {stop - start_time} s')
