# https://adventofcode.com/2022/day/16

# (You guessed 1873.) too high, but correct for someone else
# (You guessed 1872.) your answer is too high

# Your puzzle answer was 1796.

import pprint
from typing import Dict, List
from aocd import data as aocdata, lines, numbers, submit
from collections import Counter
import math
import sys
from queue import PriorityQueue
import timeit
start_time = timeit.default_timer()

pp = pprint.PrettyPrinter(indent=4)


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


data_str_test = """Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II"""


data_str = data_str_test if with_test_data else aocdata
# print(data_str)


class Node(object):
    def __init__(self, neighbours, pressure, name):
        self.neighbours = neighbours
        self.pressure = pressure
        self.name = name

    def __str__(self):
        return f'{self.name} ({self.pressure}) -> {",".join(self.neighbours)}'

    def __repr__(self):
        return f'{self.name} ({self.pressure}) -> {",".join(self.neighbours)}'


class Valve(object):
    def __init__(self, pressure, name):
        self.isopen = False
        self.pressure = pressure
        self.name = name

    def __str__(self):
        return f'{self.name} ({self.pressure}) -> isopen = {self.isopen}'

    def __repr__(self):
        return f'{self.name} ({self.pressure}) -> isopen = {self.isopen}'


nodes = {}
for line in data_str.split('\n'):
    if printing_enabled:
        print(line)
    linesplit = line.split(' ')
    name = linesplit[1]
    pressure = int(linesplit[4].split('=')[1].split(';')[0])
    neighbours = list(map(lambda str: str.replace(',', ''),
                          line.split('to ')[1].split(' ')[1:]))
    nodes[name] = Node(neighbours, pressure, name)

valves = list(
    map(lambda x: Valve(x.pressure, x.name), nodes.values()))
valves_filtered = list(
    filter(lambda x: x.pressure > 0, valves))



def dijkstra_all(nodes: Dict[str, Node], cost_fn) -> List[List[int]]:
    node_count = len(nodes)
    MAX_DIST = node_count + 1
    shortest_distances = {}
    # shortest_distances = {node.name: {node.name: MAX_DIST for node in nodes}
    # 				 for node in nodes}
    # parents = [[None for x in range(M)] for y in range(N)]

    for starting_node in nodes.values():
        dists = {node_name: MAX_DIST for node_name in nodes}
        dists[starting_node.name] = 0

        visited_nodes = []

        pq = PriorityQueue()
        pq.put((0, starting_node.name))

        while not pq.empty():
            (dist, current_node_name) = pq.get()
            current_node = nodes[current_node_name]
            visited_nodes.append(current_node.name)

            for neighbour_node_name in current_node.neighbours:
                if neighbour_node_name not in visited_nodes:
                    new_distance_through_current_node = dists[current_node.name] + cost_fn(
                        nodes[neighbour_node_name], current_node)
                    if dists[neighbour_node_name] > new_distance_through_current_node:
                        pq.put(
                            (new_distance_through_current_node, neighbour_node_name))
                        dists[neighbour_node_name] = new_distance_through_current_node
                        # parents[neighbour_node] = ...

        shortest_distances[starting_node.name] = dists

    return shortest_distances


distinminutes = dijkstra_all(nodes, lambda _, __: 1)
# pp.pprint(distinminutes)


def traverse(currentnodename, closednodenames: List[str], minutesleft: int, gainperminute: int, route: List[str]) -> int:
    # print(currentnodename, minutesleft, gainperminute)

    if minutesleft <= 0:
        # print('route: ', route, gainperminute)
        return gainperminute

    openedvalvethisminute = False

    newgainperminute = gainperminute
    # when the starting node has 0 pressure or current node is closed already, then don't spend time to open valve
    if nodes[currentnodename].pressure > 0 and len(closednodenames) != 0:
        minutesleft -= 1
        openedvalvethisminute = True
        newgainperminute += nodes[currentnodename].pressure
        route.append(currentnodename)
        closednodenames.remove(currentnodename)

    if len(closednodenames) == 0:
        return gainperminute + traverse(currentnodename, [], minutesleft - 1, newgainperminute, route)

    pressures = []
    for targetnodename in closednodenames:
        minutesneededtoreach = distinminutes[currentnodename][targetnodename]
        # if minutesneededtoreach < minutesleft and not minutesneededtoreach + 2 < minutesleft:
        #     print('aaaa')
        if minutesneededtoreach + 2 < minutesleft:
            gain_on_the_route = newgainperminute * minutesneededtoreach
            pressures.append(gain_on_the_route + traverse(targetnodename, closednodenames.copy(),
                                                                                               minutesleft - minutesneededtoreach, newgainperminute, route.copy()))
    if len(pressures) > 0:
        return (gainperminute if openedvalvethisminute else 0) + max(pressures)

    return (gainperminute if openedvalvethisminute else 0) + traverse(currentnodename, [], minutesleft - 1, newgainperminute, route)


# Solution:
solution = traverse('AA', list(
    map(lambda valve: valve.name, valves_filtered)), 30, 0, [])


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
