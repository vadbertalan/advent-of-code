# (You guessed 8506900824105.) too high
# (You guessed 8506900824111.) too high
# (You guessed 8506900824111.264.)

from aocd import data as aocdata, lines, numbers, submit
from collections import Counter
import math
import sys
import timeit
import pprint
pp = pprint.PrettyPrinter(indent=4)

start_time = timeit.default_timer()


if len(sys.argv) > 2 or (len(sys.argv) == 2 and '-' not in sys.argv[1]):
    print(f'Usage: python {sys.argv[0]} [-vls] where:')
    print('\tv - verbose')
    print('\tl - use live data')
    print('\ts - submit at the end')
    exit(1)

printing_enabled = False
with_test_data = True
should_submit = False
if len(sys.argv) == 2:
    printing_enabled = 'v' in sys.argv[1]  # verbose
    with_test_data = 'l' not in sys.argv[1]  # live data
    should_submit = 's' in sys.argv[1]  # submit at the end

# answer 152
data_str_test = """root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: dvpt - humn
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32"""


data_str = data_str_test if with_test_data else aocdata


def add(a, b): return a + b
def sub(a, b): return a - b
def mul(a, b): return a * b
def div(a, b): return a / b


def eq(a, b):
    # if a == ['X']:
    # 	return b
    # if b == ['X']:
    # 	return a
    if type(a) == list:
        return b
    if type(b) == list:
        return a
    print('ERROR: a or b should be ["X"]')


class Node:
    def __init__(self, name, op, op1, op2):
        self.name = name
        self.op = op
        self.operand1 = op1
        self.operand2 = op2

    def __str__(self):
        return self.name

    def __repr__(self):
        return self.name


nodes = []
yelled_nrs = {}
for line in data_str.split('\n'):
    line_split = line.split(': ')
    name = line_split[0]
    rest = line_split[1]

    if name == 'humn':
        yelled_nrs['humn'] = ['X']
        continue

    if rest.isnumeric():
        yelled_nrs[name] = int(rest)
    else:
        yelled_nrs[name] = None

        rest_split = rest.split(' ')
        op1, op_str, op2 = rest_split

        if op_str == '+':
            op = add
        elif op_str == '-':
            op = sub
        elif op_str == '*':
            op = mul
        elif op_str == '/':
            op = div

        if name == 'root':
            op = eq

        nodes.append(Node(name, op, op1, op2))


if printing_enabled:
    print(nodes)
    print(yelled_nrs)


def try_to_solve_node(node) -> int:
    global nodes, yelled_nrs

    result = None

    if type(node.operand1) == int and type(node.operand2) == int:
        result = int(node.op(node.operand1, node.operand2))
    elif type(node.operand1) == int and type(node.operand2) == list or type(node.operand1) == list and type(node.operand2) == int:
        result = [node.operand1, node.op, node.operand2]

    if result is not None:
        if printing_enabled:
            print(
                f'Solving node {node.name} = {node.operand1} op {node.operand2} -> {result}')
        yelled_nrs[node.name] = result
        nodes.remove(node)


# Solution:
while yelled_nrs['root'] is None:
    for node in nodes.copy():
        if type(node.operand1) == str:
            if yelled_nrs[node.operand1] is not None:
                node.operand1 = yelled_nrs[node.operand1]
        if type(node.operand2) == str:
            if yelled_nrs[node.operand2] is not None:
                node.operand2 = yelled_nrs[node.operand2]
        try_to_solve_node(node)

print(list(filter(lambda x: x[1] is None, yelled_nrs.items())))
print(nodes)

equation = yelled_nrs['root']
print()
X = equation[1](equation[0], equation[2])  # '=' op
print(X)
print()

fn_repr = {
    add: '+',
    sub: '-',
    mul: '*',
    div: '/'
}

def print_subeq(subeq):
    print('(', end='')
    for x in subeq:
        if type(x) == list:
            print_subeq(x)
        elif type(x) == int or x == 'X':
            print(x, end=' ')
        else:
            print(fn_repr[x], end=' ')
        
    print(')', end=' ')


equation = equation[0] if equation[2] == X else equation[2]
# pp.pprint(equation)
print_subeq(equation)
print()

def get_opposite_op(op):
    if op == add:
        return sub
    if op == sub:
        return add
    if op == mul:
        return div
    if op == div:
        return mul


while equation != ['X']:
    opposite_op = get_opposite_op(equation[1])

    should_negate = False

    if type(equation[0]) == list:
        operand = equation[2]
        next_equation = equation[0]
    else:
        operand = equation[0]
        next_equation = equation[2]

        if equation[1] == sub:
            opposite_op = sub
            should_negate = True
    print(f'{X} {fn_repr[opposite_op]} {operand} = {opposite_op(X, operand)}')

    X = opposite_op(X, operand)
    
    if should_negate:
        X *= -1

    equation = next_equation

print()
print(equation, ' = ', X)
solution = int(X)
print(solution)

if should_submit:
    response = input(f'Are you sure you want to submit {solution}? [y/n] -> ')
    if 'y' in response:
        print(f'✨ Submitting solution: {solution}')
        submit(solution)
    else:
        print(f'❌ Not submitting solution: {solution}')

stop = timeit.default_timer()
print(f'\n✨Time: {stop - start_time} s')
