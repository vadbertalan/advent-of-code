# this solution is not accurate, because it only allows the user to

data_str1 = """1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581"""

data_str1 = open('15_input.txt', 'r').read()

lines = data_str1.split('\n')

data = []
N = len(lines) * 5
M = len(lines[0]) * 5


def shift(subrow, i):
    return ''.join(list(map(lambda x: str((x + i) if x + i <= 9 else ((x + i) % 10) + 1), subrow)))


def gen_row(subrow, count=5):
    row = ''
    for i in range(count):
        row += shift(subrow, i)
    return row


for line in lines:
    subrow = list(map(int, line))
    data.append(gen_row(subrow, 5))

for j in range(1, 5):
    for line in lines:
        subrow = list(map(int, line))
        data.append(gen_row(list(map(int, shift(subrow, j))), 5))

extended_data = []
for line in data:
    extended_data.append(list(map(int, line)))

print(N, M)


def print_data(mat):
    for row in mat:
        for num in row:
            print(f'{num: <4}', end='')
        print()
    print()
    print()


sum_mat = [[0 for _ in range(M)] for _ in range(N)]

# first row
for i in range(1, M):
    sum_mat[0][i] = sum_mat[0][i - 1] + extended_data[0][i]

# first column
for i in range(1, N):
    sum_mat[i][0] = sum_mat[i - 1][0] + extended_data[i][0]

for i in range(1, N):
    for j in range(1, M):
        val_up = sum_mat[i - 1][j]
        val_left = sum_mat[i][j - 1]
        sum_mat[i][j] = min([val_up, val_left]) + extended_data[i][j]

print_data(sum_mat)
print(sum_mat[N - 1][M - 1])
