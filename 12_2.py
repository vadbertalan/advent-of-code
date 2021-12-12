# --- Part Two - --

# After reviewing the available paths, you realize you might have time to visit a single small cave twice. Specifically, big caves can be visited any number of times, a single small cave can be visited at most twice, and the remaining small caves can be visited at most once. However, the caves named start and end can only be visited exactly once each: once you leave the start cave, you may not return to it, and once you reach the end cave, the path must end immediately.

# Now, the 36 possible paths through the first example above are:

# start, A, b, A, b, A, c, A, end
# start, A, b, A, b, A, end
# start, A, b, A, b, end
# start, A, b, A, c, A, b, A, end
# start, A, b, A, c, A, b, end
# start, A, b, A, c, A, c, A, end
# start, A, b, A, c, A, end
# start, A, b, A, end
# start, A, b, d, b, A, c, A, end
# start, A, b, d, b, A, end
# start, A, b, d, b, end
# start, A, b, end
# start, A, c, A, b, A, b, A, end
# start, A, c, A, b, A, b, end
# start, A, c, A, b, A, c, A, end
# start, A, c, A, b, A, end
# start, A, c, A, b, d, b, A, end
# start, A, c, A, b, d, b, end
# start, A, c, A, b, end
# start, A, c, A, c, A, b, A, end
# start, A, c, A, c, A, b, end
# start, A, c, A, c, A, end
# start, A, c, A, end
# start, A, end
# start, b, A, b, A, c, A, end
# start, b, A, b, A, end
# start, b, A, b, end
# start, b, A, c, A, b, A, end
# start, b, A, c, A, b, end
# start, b, A, c, A, c, A, end
# start, b, A, c, A, end
# start, b, A, end
# start, b, d, b, A, c, A, end
# start, b, d, b, A, end
# start, b, d, b, end
# start, b, end
# The slightly larger example above now has 103 paths through it, and the even larger example now has 3509 paths through it.

# Given these new rules, how many paths through this cave system are there?

# Your puzzle answer was 83475.

import copy

data_str1 = """fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW"""

data_str = """re-js
qx-CG
start-js
start-bj
qx-ak
js-bj
ak-re
CG-ak
js-CG
bj-re
ak-lg
lg-CG
qx-re
WP-ak
WP-end
re-lg
end-ak
WP-re
bj-CG
qx-start
bj-WP
JG-lg
end-lg
lg-iw"""

lines = data_str.split('\n')


g = {}
seen = {}

for line in lines:
    [a, b] = line.split('-')
    # print(a, b, line.split('-'))
    if a not in g.keys():
        g[a] = [b]
    else:
        g[a].append(b)
    if b not in g.keys():
        g[b] = [a]
    else:
        g[b].append(a)
    seen[a] = False
    seen[b] = False

print(g)
print()

paths = []


def visit(g, node: str, path=[]):
    global res

    localpath = copy.deepcopy(path)
    localpath.append(node)

    if (node == 'end'):
        # print('-'.join(localpath))
        paths.append('-'.join(localpath))
        return

    for n in g[node]:
        if n not in path or n.isupper():
            visit(g, n, localpath)


# determining which nodes to duplicate
new_nodes = []
for node in g.keys():
    if node not in ['start', 'end'] and not node.isupper():
        new_nodes.append(node)

# duplicating node: both incoming and outgoing edges are considered
for n in new_nodes:
    g_copy = copy.deepcopy(g)

    g_copy[n * 2] = g_copy[n]

    for node in g.keys():
        if n in g_copy[node]:
            g_copy[node].append(n * 2)
    visit(g_copy, 'start')

    # converting paths A-b-A-bb to A-b-A-b in resulting array
    new_paths = []
    for path in paths:
        nodes = path.split('-')
        newpath = list(
            map(lambda node: node[:len(node) // 2] if node == n * 2 else node, nodes))
        new_paths.append('-'.join(newpath))
    paths = new_paths

print(len(paths))
a = set(paths)
print(len(a))
