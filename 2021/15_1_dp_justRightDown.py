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

# data_str1 = open('15_input.txt', 'r').read()

lines = data_str1.split('\n')

data = []
N = len(lines)
M = len(lines[0])
print(N, M)

for line in lines:
    data.append(list(map(int, line)))


def valid_coords(coord):
    global N, M 
    return 0 <= coord[0] < N and 0 <= coord[1] < M

def get(coord): 
    global data
    return data[coord[0]][coord[1]]

dirs = ['right', 'down', 'left', 'up']
op_dirs = {
    'right': 'left',
    'left': 'right',
    'up': 'down',
    'down': 'up'
}
dp = {}
seen = [[False for x in range(M)] for y in range(N)]

def move(from_pos, to_dir): 
    global N, M, dp

    if from_pos == [N - 1, M - 1]:
        return get(from_pos)

    key = f'{from_pos[0]}-{from_pos[1]}'
    if key in dp:
        return dp[key]

    d = {
        'right': [from_pos[0], from_pos[1] + 1],
        'down': [from_pos[0] + 1, from_pos[1]],
    }

    dir_coords = list(filter(lambda item: valid_coords(item[1]), d.items()))
    
    dir_costs = list(map(lambda dir_coord: move(dir_coord[1], dir_coord[0]), dir_coords))

    min_cost = min(dir_costs)

    dp[key] = get(from_pos) + min_cost
    return dp[key]


print(move([0, 0], 'down') - data[0][0])
