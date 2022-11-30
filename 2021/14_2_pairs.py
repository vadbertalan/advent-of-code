# --- Part Two - --

# The resulting polymer isn't nearly strong enough to reinforce the submarine. You'll need to run more steps of the pair insertion process
# a total of 40 steps should do it.

# In the above example, the most common element is B(occurring 2192039569602 times) and the least common element is H(occurring 3849876073 times)
# subtracting these produces 2188189693529.

# Apply 40 steps of pair insertion to the polymer template and find the most and least common elements in the result. What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?

# Your puzzle answer was 4110568157153.

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

res = 0

initial_pairs = []

i = 0
while i < len(data) - 1:
    initial_pairs.append(data[i:i+2])
    i += 1

print('initial pairs', initial_pairs)

pair_counts = Counter(initial_pairs)


char_counts = {char: 0 for char in set(
    [item for sublist in d.keys() for item in sublist])}

for char in data:
    char_counts[char] += 1
print('\nchar counts initially: ', char_counts)

x = 40
print()

step = 1
while step <= x:
    new_pair_counts = {}
    for pair in [*pair_counts.keys()]:
        for new_pair in [f'{pair[0]}{d[pair]}', f'{d[pair]}{pair[1]}']:
            if new_pair in new_pair_counts:
                new_pair_counts[new_pair] += pair_counts[pair]
            else:
                new_pair_counts[new_pair] = pair_counts[pair]

        char_counts[d[pair]] += pair_counts[pair]

    pair_counts = new_pair_counts

    step += 1

print('-------')
print('pair_counts: ', pair_counts)

print('-------')
print(char_counts)

counts = sorted(char_counts.values())

for c in counts:
    print(c)

print('-------')
print(counts[-1] - counts[0])
