# --- Part Two - --

# Maybe a fancy trick shot isn't the best idea
# after all, you only have one probe, so you had better not miss.

# To get the best idea of what your options are for launching the probe, you need to find every initial velocity that causes the probe to eventually be within the target area after any step.

# In the above example, there are 112 different initial velocity values that meet these criteria:

# 23, -10  25, -9   27, -5   29, -6   22, -6   21, -7   9, 0     27, -7   24, -5
# 25, -7   26, -6   25, -5   6, 8     11, -2   20, -5   29, -10  6, 3     28, -7
# 8, 0     30, -6   29, -8   20, -10  6, 7     6, 4     6, 1     14, -4   21, -6
# 26, -10  7, -1    7, 7     8, -1    21, -9   6, 2     20, -7   30, -10  14, -3
# 20, -8   13, -2   7, 3     28, -8   29, -9   15, -3   22, -5   26, -8   25, -8
# 25, -6   15, -4   9, -2    15, -2   12, -2   28, -9   12, -3   24, -6   23, -7
# 25, -10  7, 8     11, -3   26, -7   7, 1     23, -9   6, 0     22, -10  27, -6
# 8, 1     22, -8   13, -4   7, 6     28, -6   11, -4   12, -4   26, -9   7, 4
# 24, -10  23, -8   30, -8   7, 0     9, -1    10, -1   26, -5   22, -9   6, 5
# 7, 5     23, -6   28, -10  10, -2   11, -1   20, -9   14, -2   29, -7   13, -3
# 23, -5   24, -8   27, -9   30, -7   28, -5   21, -10  7, 9     6, 6     21, -5
# 27, -10  7, 2     30, -9   21, -8   22, -7   24, -9   20, -6   6, 9     29, -5
# 8, -2    27, -8   30, -5   24, -7
# How many distinct initial velocity values cause the probe to be within the target area after any step?

# Your puzzle answer was 5200.

input = """target area: x=20..30, y=-10..-5"""
input = """target area: x=57..116, y=-198..-148"""

xs = list(map(int, input.split('x=')[1].split(',')[0].split('..')))
ys = list(map(int, input.split('y=')[1].split('..')))

print(xs, ys)

# finding the initial X velocity
vel_x = 1
res = vel_x * (vel_x + 1) / 2
while res < xs[0]:
    vel_x += 1
    res = vel_x * (vel_x + 1) / 2

# brute forcing velocity X and Y. X starts from its min possible value
c = 0
while True:
    vel_y = ys[0]
    while True:
        vel_x_temp = vel_x
        vel_y_temp = vel_y
        y = 0
        x = 0
        # print(f'trying throw with vel({vel_x_temp} {vel_y_temp}) pos({x} {y})')
        while (not (ys[0] <= y <= ys[1] and xs[0] <= x <= xs[1])) and (not y < ys[0] and not x > xs[1]):
            y += vel_y_temp
            vel_y_temp -= 1

            x += vel_x_temp
            vel_x_temp += -1 if vel_x_temp > 0 else 1 if vel_x_temp < 0 else 0
        if ys[0] <= y <= ys[1] and xs[0] <= x <= xs[1]:
            print(
                f'{vel_x},{vel_y}')
            c += 1

        # stopping condition is when the vel_y_temp passed ys[0]
        if vel_y_temp < ys[0] and y < ys[0] and -vel_y < ys[0] or x > xs[1]:
            break
        vel_y += 1
    if vel_x > xs[1]:
        break
    vel_x += 1

print(c)
