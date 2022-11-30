# <!-- You're already almost 1.5km (almost a mile) below the surface of the ocean, already so deep that you can't see any sunlight. What you can see, however, is a giant squid that has attached itself to the outside of your submarine.

# Maybe it wants to play bingo?

# Bingo is played on a set of boards each consisting of a 5x5 grid of numbers. Numbers are chosen at random, and the chosen number is marked on all boards on which it appears. (Numbers may not appear on all boards.) If all numbers in any row or any column of a board are marked, that board wins. (Diagonals don't count.)

# The submarine has a bingo subsystem to help passengers (currently, you and the giant squid) pass the time. It automatically generates a random order in which to draw numbers and a random set of boards (your puzzle input). For example:

# 7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

# 22 13 17 11  0
#  8  2 23  4 24
# 21  9 14 16  7
#  6 10  3 18  5
#  1 12 20 15 19

#  3 15  0  2 22
#  9 18 13 17  5
# 19  8  7 25 23
# 20 11 10 24  4
# 14 21 16 12  6

# 14 21 17 24  4
# 10 16 15  9 19
# 18  8 23 26 20
# 22 11 13  6  5
#  2  0 12  3  7
# After the first five numbers are drawn (7, 4, 9, 5, and 11), there are no winners, but the boards are marked as follows (shown here adjacent to each other to save space):

# 22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
#  8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
# 21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
#  6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
#  1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
# After the next six numbers are drawn (17, 23, 2, 0, 14, and 21), there are still no winners:

# 22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
#  8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
# 21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
#  6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
#  1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
# Finally, 24 is drawn:

# 22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
#  8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
# 21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
#  6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
#  1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
# At this point, the third board wins because it has at least one complete row or column of marked numbers (in this case, the entire top row is marked: 14 21 17 24 4).

# The score of the winning board can now be calculated. Start by finding the sum of all unmarked numbers on that board; in this case, the sum is 188. Then, multiply that sum by the number that was just called when the board won, 24, to get the final score, 188 * 24 = 4512.

# To guarantee victory against the giant squid, figure out which board will win first. What will your final score be if you choose that board? -->

# Your puzzle answer was 32844.

numbers_raw = """7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1"""

tables_raw = """22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
"""

BINGO_TABLE_SIZE = 25
BINGO_ROW_SIZE = 5

numbers = numbers_raw.split(',')

print(numbers)
print(len(numbers))

tables_numbers = tables_raw.split()

table_count = len(tables_numbers) // BINGO_TABLE_SIZE

tables = []
c = 0
t = 0
for i in range(1, table_count + 1):
    tables.append([[tables_numbers[x], False]
                  for x in range(c, i * BINGO_TABLE_SIZE)])
    c += BINGO_TABLE_SIZE
    print(tables[i - 1])
    print(len(tables[i - 1]))

# check for bingo


def check_for_bingo():
    for table_index, table in enumerate(tables):
        # this for could be made better, because i has to be iterated over the first row and first column only
        for i in range(0, BINGO_TABLE_SIZE):
            def get_marked_value(x): return x[1]
            # checking rows and columns
            if (i % BINGO_ROW_SIZE == 0 and all(map(get_marked_value, table[i: i + BINGO_ROW_SIZE]))) or (i < BINGO_ROW_SIZE and all(map(get_marked_value, [table[x] for x in range(i, BINGO_TABLE_SIZE, BINGO_ROW_SIZE)]))):
                return table_index, i
    return -1, -1


# process numbers
for nr in numbers:
    for i in range(table_count):
        for j in range(len(tables[i])):
            # marking number True
            if tables[i][j][0] == nr:
                tables[i][j][1] = True
                break  # break here because numbers in tables are distinct

    table_index, row_index = check_for_bingo()
    if table_index != -1 and row_index != 1:  # if bingo happened
        print('Bingo happened at: ', table_index, row_index)

        print(tables[table_index])

        sum = 0
        for table_item in tables[table_index]:
            sum += int(table_item[0]) if not table_item[1] else 0
            print(table_item)
            print('sum ->', sum)

        print(sum, nr, sum * int(nr))

        break
