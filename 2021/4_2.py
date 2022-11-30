# On the other hand, it might be wise to try a different strategy: let the giant squid win.

# You aren't sure how many bingo boards a giant squid could play at once, so rather than waste time counting its arms, the safe thing to do is to figure out which board will win last and choose that one. That way, no matter which boards it picks, it will win for sure.

# In the above example, the second board is the last to win, which happens after 13 is eventually called and its middle column is completely marked. If you were to keep playing until this point, the second board would have a sum of unmarked numbers equal to 148 for a final score of 148 * 13 = 1924.

# Figure out which board will win last. Once it wins, what would its final score be?

# Your puzzle answer was 4920.

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

winner_tables = set()


def check_for_bingo():
    for table_index, table in enumerate(tables):
        # this for could be made better, because i has to be iterated over the first row and first column only
        for i in range(0, BINGO_TABLE_SIZE):
            def get_marked_value(x): return x[1]
            # checking rows and columns
            if (i % BINGO_ROW_SIZE == 0 and all(map(get_marked_value, table[i: i + BINGO_ROW_SIZE]))) or (i < BINGO_ROW_SIZE and all(map(get_marked_value, [table[x] for x in range(i, BINGO_TABLE_SIZE, BINGO_ROW_SIZE)]))):
                if table_index not in winner_tables and len(winner_tables) == table_count - 1:
                    return table_index, i
                else:
                    winner_tables.add(table_index)
                    print(winner_tables)
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
    # if bingo happened and it's the last table
    if table_index != -1 and row_index != 1:
        print('Bingo happened at: ', table_index, row_index)

        print(tables[table_index])

        sum = 0
        for table_item in tables[table_index]:
            sum += int(table_item[0]) if not table_item[1] else 0
            print(table_item)
            print('sum ->', sum)

        print(sum, nr, sum * int(nr))

        break
