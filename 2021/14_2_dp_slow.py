from collections import Counter

data_str1 = """NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C"""

data_str = """CKFFSCFSCBCKBPBCSPKP

NS -> P
KV -> B
FV -> S
BB -> V
CF -> O
CK -> N
BC -> B
PV -> N
KO -> C
CO -> O
HP -> P
HO -> P
OV -> O
VO -> C
SP -> P
BV -> H
CB -> F
SF -> H
ON -> O
KK -> V
HC -> N
FH -> P
OO -> P
VC -> F
VP -> N
FO -> F
CP -> C
SV -> S
PF -> O
OF -> H
BN -> V
SC -> V
SB -> O
NC -> P
CN -> K
BP -> O
PC -> H
PS -> C
NB -> K
VB -> P
HS -> V
BO -> K
NV -> B
PK -> K
SN -> H
OB -> C
BK -> S
KH -> P
BS -> S
HV -> O
FN -> F
FS -> N
FP -> F
PO -> B
NP -> O
FF -> H
PN -> K
HF -> H
VK -> K
NF -> K
PP -> H
PH -> B
SK -> P
HN -> B
VS -> V
VN -> N
KB -> O
KC -> O
KP -> C
OS -> O
SO -> O
VH -> C
OK -> B
HH -> B
OC -> P
CV -> N
SH -> O
HK -> N
NO -> F
VF -> S
NN -> O
FK -> V
HB -> O
SS -> O
FB -> B
KS -> O
CC -> S
KF -> V
VV -> S
OP -> H
KN -> F
CS -> H
CH -> P
BF -> F
NH -> O
NK -> C
OH -> C
BH -> O
FC -> V
PB -> B"""

dat = data_str
data = dat.split('\n\n')[0].split('\n')[0]
instructions = list(map(lambda x: x.split(' -> '),
                    dat.split('\n\n')[1].split('\n')))

d = {}
for inst in instructions:
    d[inst[0]] = inst[1]

pairs = []
i = 0
while i < len(data) - 1:
    pairs.append(data[i:i+2])
    i += 1
print(pairs)

dp = {}


def substitute(str_of_len_2, c):
    return f'{str_of_len_2[0]}{c}{str_of_len_2[1]}'


def calc(substr, steps_left):
    global dp

    # print(f'calcing {substr} {steps_left}')
    if steps_left == 0 or substr not in d.keys():
        return substr

    key = f'{substr}-{steps_left}'
    if key in dp.keys():
        # print(f'Using DP - dp[{key}]')
        return dp[key]

    sub = substitute(substr, d[substr])
    dp[key] = calc(sub[:2], steps_left - 1)[:-1] + \
        calc(sub[1:], steps_left - 1)
    return dp[key]


x = 20

results = list(map(lambda pair: calc(pair, x), pairs))

data = ''.join(list(map(lambda str: str[:-1], results[:-1]))) + results[-1]

print('Data is ready')

counts = Counter(data)

print(counts)
counts = sorted(counts.values())

for c in counts:
    print(c)

print('----- bad solution -----')
print(counts[-1] - counts[0])
