# --- Part Two - --

# It seems like the individual flashes aren't bright enough to navigate. However, you might have a better option: the flashes seem to be synchronizing!

# In the example above, the first time all octopuses flash simultaneously is step 195:

# After step 193:
# 5877777777
# 8877777777
# 7777777777
# 7777777777
# 7777777777
# 7777777777
# 7777777777
# 7777777777
# 7777777777
# 7777777777

# After step 194:
# 6988888888
# 9988888888
# 8888888888
# 8888888888
# 8888888888
# 8888888888
# 8888888888
# 8888888888
# 8888888888
# 8888888888

# After step 195:
# 0000000000
# 0000000000
# 0000000000
# 0000000000
# 0000000000
# 0000000000
# 0000000000
# 0000000000
# 0000000000
# 0000000000
# If you can calculate the exact moments when the octopuses will all flash simultaneously, you should be able to navigate through the cavern. What is the first step during which all octopuses flash?

# Your puzzle answer was 422.

data_str1 = """5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526"""

data_str = """7232374314
8531113786
3411787828
5482241344
5856827742
7614532764
5311321758
1255116187
5821277714
2623834788"""

lines = data_str.split('\n')

res = 0

data = []
N = len(lines)
M = len(lines[0])
print(N, M)

for line in lines:
    data.append(list(map(int, line)))


def print_data():
    for row in data:
        for num in row:
            print(f'{num: < 3}', end='')
        print()
    print()


def flash(i, j):
    data[i][j] = 0

    adj_i = [i - 1, i - 1, i - 1, i, i + 1, i + 1, i + 1, i]
    adj_j = [j - 1, j, j + 1, j + 1, j + 1, j, j - 1, j - 1]
    neighbor_coords = list(
        filter(lambda p: 0 <= p[0] < N and 0 <= p[1] < M, list(zip(adj_i, adj_j))))

    for x, y in neighbor_coords:
        data[x][y] += 1 if data[x][y] != 0 else 0
        if data[x][y] > 9:
            flash(x, y)


while True:
    res += 1
    # print_data()
    for i in range(N):
        for j in range(M):
            data[i][j] += 1
    for i in range(N):
        for j in range(M):
            if data[i][j] > 9:
                flash(i, j)
    ex = True
    for row in data:
        if all([x == 0 for x in row]):
            continue
        ex = False
    if ex:
        break


print_data()

print(res)
