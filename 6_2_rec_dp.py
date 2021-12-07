# --- Part Two - --

# Suppose the lanternfish live forever and have unlimited food and space. Would they take over the entire ocean?

# After 256 days in the example above, there would be a total of 26984457539 lanternfish!

# How many lanternfish would there be after 256 days?

# Your puzzle answer was 1732731810807.

data_str = """3,4,3,1,2"""
data_str = """1,3,4,1,5,2,1,1,1,1,5,1,5,1,1,1,1,3,1,1,1,1,1,1,1,2,1,5,1,1,1,1,1,4,4,1,1,4,1,1,2,3,1,5,1,4,1,2,4,1,1,1,1,1,1,1,1,2,5,3,3,5,1,1,1,1,4,1,1,3,1,1,1,2,3,4,1,1,5,1,1,1,1,1,2,1,3,1,3,1,2,5,1,1,1,1,5,1,5,5,1,1,1,1,3,4,4,4,1,5,1,1,4,4,1,1,1,1,3,1,1,1,1,1,1,3,2,1,4,1,1,4,1,5,5,1,2,2,1,5,4,2,1,1,5,1,5,1,3,1,1,1,1,1,4,1,2,1,1,5,1,1,4,1,4,5,3,5,5,1,2,1,1,1,1,1,3,5,1,2,1,2,1,3,1,1,1,1,1,4,5,4,1,3,3,1,1,1,1,1,1,1,1,1,5,1,1,1,5,1,1,4,1,5,2,4,1,1,1,2,1,1,4,4,1,2,1,1,1,1,5,3,1,1,1,1,4,1,4,1,1,1,1,1,1,3,1,1,2,1,1,1,1,1,2,1,1,1,1,1,1,1,2,1,1,1,1,1,1,4,1,1,1,1,1,1,1,1,1,1,1,1,1,1,2,1,1,2,5,1,2,1,1,1,1,1,1,1,1,1"""

data = list(map(int, data_str.split(',')))

total_days = 256

dpm = {}  # DP matrix


def spawn(timer, days_passed, parent, print_indent=''):
    global total_days

    if days_passed > total_days:
        return 0

    # print(
    #     f'{print_indent}Spawning fish number {c} on day {days_passed} with timer {timer}. Parent is {parent}')
    day = days_passed
    local_timer = timer
    c = 0
    while (day <= total_days):
        day += local_timer

        dpm_key = f'{str(day)}-{str(timer)}'
        if dpm_key in dpm:
            children_count = dpm[dpm_key]
        else:
            children_count = spawn(8, day + 1, c, print_indent + '  ')
            dpm[dpm_key] = children_count

        c += children_count
        local_timer = 7

    return 1 + c


counts = list(map(lambda timer: spawn(timer, 0, -9999), data))
count = sum(counts)

print(count)
