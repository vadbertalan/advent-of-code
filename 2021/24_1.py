import timeit
start_time = timeit.default_timer()

input1 = '''inp w
mul x 0
add x z
mod x 26
div z 1
add x 10
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 13
mul y x
add z y'''

memory = {}


def reset_memory():
    global memory
    memory = {
        'x': 0,
        'y': 0,
        'z': 0,
        'w': 0,
    }


def parse_var(var):
    try:
        return int(var)
    except:
        return memory[var]


def parse_com(line, buf):
    global memory

    tokens = line.split(' ')
    if len(tokens) == 2:
        com, op1 = tokens
    else:
        com, op1, op2 = tokens

    if com == 'inp':
        memory[op1] = int(buf.pop(0))
    elif com == 'add':
        memory[op1] += parse_var(op2)
    elif com == 'div':
        memory[op1] //= parse_var(op2)
    elif com == 'mul':
        memory[op1] *= parse_var(op2)
    elif com == 'mod':
        memory[op1] %= parse_var(op2)
    elif com == 'eql':
        memory[op1] = int(memory[op1] == (parse_var(op2)))


alu = open('24.alu', 'r').read()
alu_commands = input1.split('\n')


def run_alu(buffer):
    global memory

    reset_memory()

    for com in alu_commands:
        parse_com(com, buffer)

    return memory.copy()


# model_nr = int('9' * 14)
model_nr = 9
while model_nr > 0:
    print(f'Running ALU on {model_nr}')
    buffer = list(str(model_nr))
    # if '0' in buffer:
    #     model_nr -= 1
    #     continue
    res = run_alu(buffer)
    print(res)
    print()
    # break
    # if res['z'] == 0:
    #     print(model_nr)
    #     break
    model_nr -= 1


stop = timeit.default_timer()
print(f'\n✨ Time: {stop - start_time} s')
