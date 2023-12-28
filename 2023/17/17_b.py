# https://adventofcode.com/2023/day/17

# Your puzzle answer was 925.

from heapq import heappush, heappop

data = open('17.in', 'r')

grid = [list(map(int, line.strip())) for line in data]

seen = set()
pq = [(0, 0, 0, 0, 0, 0)]

while pq:
    hl, r, c, dr, dc, n = heappop(pq)
    
    if r == len(grid) - 1 and c == len(grid[0]) - 1:
        print(hl)
        break

    if (r, c, dr, dc, n) in seen:
        continue

    seen.add((r, c, dr, dc, n))
    
    # if n < 4 and (dr, dc) != (0, 0):
    #     nr = r + dr
    #     nc = c + dc
    #     if 0 <= nr < len(grid) and 0 <= nc < len(grid[0]):
    #         heappush(pq, (hl + grid[nr][nc], nr, nc, dr, dc, n + 1))

    if n < 10 and (dr, dc) != (0, 0):
        nr = r + dr
        nc = c + dc
        if 0 <= nr < len(grid) and 0 <= nc < len(grid[0]):
            heappush(pq, (hl + grid[nr][nc], nr, nc, dr, dc, n + 1))
        if n < 4:
            continue

    for ndr, ndc in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
        if (ndr, ndc) != (dr, dc) and (ndr, ndc) != (-dr, -dc):
            nr = r + ndr
            nc = c + ndc
            if 0 <= nr < len(grid) and 0 <= nc < len(grid[0]):
                heappush(pq, (hl + grid[nr][nc], nr, nc, ndr, ndc, 1))