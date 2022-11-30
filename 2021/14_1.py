# --- Day 14: Extended Polymerization - --

# The incredible pressures at this depth are starting to put a strain on your submarine. The submarine has polymerization equipment that would produce suitable materials to reinforce the submarine, and the nearby volcanically-active caves should even have the necessary input elements in sufficient quantities.

# The submarine manual contains instructions for finding the optimal polymer formula
# specifically, it offers a polymer template and a list of pair insertion rules(your puzzle input). You just need to work out what polymer would result after repeating the pair insertion process a few times.

# For example:

# NNCB

# CH -> B
# HH -> N
# CB -> H
# NH -> C
# HB -> C
# HC -> B
# HN -> C
# NN -> C
# BH -> H
# NC -> B
# NB -> B
# BN -> B
# BB -> N
# BC -> B
# CC -> N
# CN -> C
# The first line is the polymer template - this is the starting point of the process.

# The following section defines the pair insertion rules. A rule like AB -> C means that when elements A and B are immediately adjacent, element C should be inserted between them. These insertions all happen simultaneously.

# So, starting with the polymer template NNCB, the first step simultaneously considers all three pairs:

# The first pair(NN) matches the rule NN -> C, so element C is inserted between the first N and the second N.
# The second pair(NC) matches the rule NC -> B, so element B is inserted between the N and the C.
# The third pair(CB) matches the rule CB -> H, so element H is inserted between the C and the B.
# Note that these pairs overlap: the second element of one pair is the first element of the next pair. Also, because all pairs are considered simultaneously, inserted elements are not considered to be part of a pair until the next step.

# After the first step of this process, the polymer becomes NCNBCHB.

# Here are the results of a few steps using the above rules:

# Template:     NNCB
# After step 1: NCNBCHB
# After step 2: NBCCNBBBCBHCB
# After step 3: NBBBCNCCNBBNBNBBCHBHHBCHB
# After step 4: NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB
# This polymer grows quickly. After step 5, it has length 97
# After step 10, it has length 3073. After step 10, B occurs 1749 times, C occurs 298 times, H occurs 161 times, and N occurs 865 times
# taking the quantity of the most common element(B, 1749) and subtracting the quantity of the least common element(H, 161) produces 1749 - 161 = 1588.

# Apply 10 steps of pair insertion to the polymer template and find the most and least common elements in the result. What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?

# Your puzzle answer was 3247.


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

data = data_str1.split('\n\n')[0].split('\n')[0]

instructions = list(map(lambda x: x.split(' -> '),
                    data_str1.split('\n\n')[1].split('\n')))

d = {}
for inst in instructions:
    d[inst[0]] = inst[1]

res = 0

print(data)
print(instructions)
print(d)

step = 1
while step <= 10:
    i = 0
    while i < len(data) - 1:
        substr = data[i:i + 2]
        if substr in d.keys():
            data = data.replace(
                substr, f'{substr[0]}{d[substr]}{substr[1]}', 1)
            i += 1
        i += 1
    print(f'after step {step}: {len(data)} - {Counter(data)}')
    step += 1

counts = []
for c in data:
    counts.append(data.count(c))

counts = list(set(counts))
counts = sorted(counts)

print(counts)

print(counts[-1] - counts[0])


print(res)
