# --- Day 16: Packet Decoder - --

# As you leave the cave and reach open waters, you receive a transmission from the Elves back on the ship.

# The transmission was sent using the Buoyancy Interchange Transmission System(BITS), a method of packing numeric expressions into a binary sequence. Your submarine's computer has saved the transmission in hexadecimal(your puzzle input).

# The first step of decoding the message is to convert the hexadecimal representation into binary. Each character of hexadecimal corresponds to four bits of binary data:

# 0 = 0000
# 1 = 0001
# 2 = 0010
# 3 = 0011
# 4 = 0100
# 5 = 0101
# 6 = 0110
# 7 = 0111
# 8 = 1000
# 9 = 1001
# A = 1010
# B = 1011
# C = 1100
# D = 1101
# E = 1110
# F = 1111
# The BITS transmission contains a single packet at its outermost layer which itself contains many other packets. The hexadecimal representation of this packet might encode a few extra 0 bits at the end
# these are not part of the transmission and should be ignored.

# Every packet begins with a standard header: the first three bits encode the packet version, and the next three bits encode the packet type ID. These two values are numbers
# all numbers encoded in any packet are represented as binary with the most significant bit first. For example, a version encoded as the binary sequence 100 represents the number 4.

# Packets with type ID 4 represent a literal value. Literal value packets encode a single binary number. To do this, the binary number is padded with leading zeroes until its length is a multiple of four bits, and then it is broken into groups of four bits. Each group is prefixed by a 1 bit except the last group, which is prefixed by a 0 bit. These groups of five bits immediately follow the packet header. For example, the hexadecimal string D2FE28 becomes:

# 110100101111111000101000
# VVVTTTAAAAABBBBBCCCCC
# Below each bit is a label indicating its purpose:

# The three bits labeled V(110) are the packet version, 6.
# The three bits labeled T(100) are the packet type ID, 4, which means the packet is a literal value.
# The five bits labeled A(10111) start with a 1 (not the last group, keep reading) and contain the first four bits of the number, 0111.
# The five bits labeled B(11110) start with a 1 (not the last group, keep reading) and contain four more bits of the number, 1110.
# The five bits labeled C(00101) start with a 0 (last group, end of packet) and contain the last four bits of the number, 0101.
# The three unlabeled 0 bits at the end are extra due to the hexadecimal representation and should be ignored.
# So, this packet represents a literal value with binary representation 011111100101, which is 2021 in decimal.

# Every other type of packet(any packet with a type ID other than 4) represent an operator that performs some calculation on one or more sub-packets contained within. Right now, the specific operations aren't important
# focus on parsing the hierarchy of sub-packets.

# An operator packet contains one or more packets. To indicate which subsequent binary data represents its sub-packets, an operator packet can use one of two modes indicated by the bit immediately after the packet header
# this is called the length type ID:

# If the length type ID is 0, then the next 15 bits are a number that represents the total length in bits of the sub-packets contained by this packet.
# If the length type ID is 1, then the next 11 bits are a number that represents the number of sub-packets immediately contained by this packet.
# Finally, after the length type ID bit and the 15-bit or 11-bit field, the sub-packets appear.

# For example, here is an operator packet(hexadecimal string 38006F45291200) with length type ID 0 that contains two sub-packets:

# 00111000000000000110111101000101001010010001001000000000
# VVVTTTILLLLLLLLLLLLLLLAAAAAAAAAAABBBBBBBBBBBBBBBB
# The three bits labeled V(001) are the packet version, 1.
# The three bits labeled T(110) are the packet type ID, 6, which means the packet is an operator.
# The bit labeled I(0) is the length type ID, which indicates that the length is a 15-bit number representing the number of bits in the sub-packets.
# The 15 bits labeled L(000000000011011) contain the length of the sub-packets in bits, 27.
# The 11 bits labeled A contain the first sub-packet, a literal value representing the number 10.
# The 16 bits labeled B contain the second sub-packet, a literal value representing the number 20.
# After reading 11 and 16 bits of sub-packet data, the total length indicated in L(27) is reached, and so parsing of this packet stops.

# As another example, here is an operator packet(hexadecimal string EE00D40C823060) with length type ID 1 that contains three sub-packets:

# 11101110000000001101010000001100100000100011000001100000
# VVVTTTILLLLLLLLLLLAAAAAAAAAAABBBBBBBBBBBCCCCCCCCCCC
# The three bits labeled V(111) are the packet version, 7.
# The three bits labeled T(011) are the packet type ID, 3, which means the packet is an operator.
# The bit labeled I(1) is the length type ID, which indicates that the length is a 11-bit number representing the number of sub-packets.
# The 11 bits labeled L(00000000011) contain the number of sub-packets, 3.
# The 11 bits labeled A contain the first sub-packet, a literal value representing the number 1.
# The 11 bits labeled B contain the second sub-packet, a literal value representing the number 2.
# The 11 bits labeled C contain the third sub-packet, a literal value representing the number 3.
# After reading 3 complete sub-packets, the number of sub-packets indicated in L(3) is reached, and so parsing of this packet stops.

# For now, parse the hierarchy of the packets throughout the transmission and add up all of the version numbers.

# Here are a few more examples of hexadecimal-encoded transmissions:

# 8A004A801A8002F478 represents an operator packet(version 4) which contains an operator packet(version 1) which contains an operator packet(version 5) which contains a literal value(version 6)
# this packet has a version sum of 16.
# 620080001611562C8802118E34 represents an operator packet(version 3) which contains two sub-packets
# each sub-packet is an operator packet that contains two literal values. This packet has a version sum of 12.
# C0015000016115A2E0802F182340 has the same structure as the previous example, but the outermost packet uses a different length type ID. This packet has a version sum of 23.
# A0016C880162017C3686B18A3D4780 is an operator packet that contains an operator packet that contains an operator packet that contains five literal values
# it has a version sum of 31.
# Decode the structure of your hexadecimal-encoded BITS transmission
# what do you get if you add up the version numbers in all packets?

# Your puzzle answer was 873.

# --- Part Two - --

# Now that you have the structure of your transmission decoded, you can calculate the value of the expression it represents.

# Literal values(type ID 4) represent a single number as described above. The remaining type IDs are more interesting:

# Packets with type ID 0 are sum packets - their value is the sum of the values of their sub-packets. If they only have a single sub-packet, their value is the value of the sub-packet.
# Packets with type ID 1 are product packets - their value is the result of multiplying together the values of their sub-packets. If they only have a single sub-packet, their value is the value of the sub-packet.
# Packets with type ID 2 are minimum packets - their value is the minimum of the values of their sub-packets.
# Packets with type ID 3 are maximum packets - their value is the maximum of the values of their sub-packets.
# Packets with type ID 5 are greater than packets - their value is 1 if the value of the first sub-packet is greater than the value of the second sub-packet
# otherwise, their value is 0. These packets always have exactly two sub-packets.
# Packets with type ID 6 are less than packets - their value is 1 if the value of the first sub-packet is less than the value of the second sub-packet
# otherwise, their value is 0. These packets always have exactly two sub-packets.
# Packets with type ID 7 are equal to packets - their value is 1 if the value of the first sub-packet is equal to the value of the second sub-packet
# otherwise, their value is 0. These packets always have exactly two sub-packets.
# Using these rules, you can now work out the value of the outermost packet in your BITS transmission.

# For example:

# C200B40A82 finds the sum of 1 and 2, resulting in the value 3.
# 04005AC33890 finds the product of 6 and 9, resulting in the value 54.
# 880086C3E88112 finds the minimum of 7, 8, and 9, resulting in the value 7.
# CE00C43D881120 finds the maximum of 7, 8, and 9, resulting in the value 9.
# D8005AC2A8F0 produces 1, because 5 is less than 15.
# F600BC2D8F produces 0, because 5 is not greater than 15.
# 9C005AC2F8F0 produces 0, because 5 is not equal to 15.
# 9C0141080250320F1802104A08 produces 1, because 1 + 3 = 2 * 2.
# What do you get if you evaluate the expression represented by your hexadecimal-encoded BITS transmission?

# Your puzzle answer was 402817863665.


from math import prod

data_str1 = """9C0141080250320F1802104A08"""
data_str1 = """A20D5080210CE4BB9BAFB001BD14A4574C014C004AE46A9B2E27297EECF0C013F00564776D7E3A825CAB8CD47B6C537DB99CD746674C1000D29BBC5AC80442966FB004C401F8771B61D8803D0B22E4682010EE7E59ACE5BC086003E3270AE4024E15C8010073B2FAD98E004333F9957BCB602E7024C01197AD452C01295CE2DC9934928B005DD258A6637F534CB3D89A944230043801A596B234B7E58509E88798029600BCF5B3BA114F5B3BA10C9E77BAF20FA4016FCDD13340118B929DD4FD54E60327C00BEB7002080AA850031400D002369400B10034400F30021400F20157D804AD400FE00034E000A6D001EB2004E5C00B9AE3AC3C300470029091ACADBFA048D656DFD126792187008635CD736B3231A51BA5EBDF42D4D299804F26B33C872E213C840022EC9C21FFB34EDE7C559C8964B43F8AD77570200FC66697AFEB6C757AC0179AB641E6AD9022006065CEA714A4D24C0179F8E795D3078026200FC118EB1B40010A8D11EA27100990200C45A83F12C401A8611D60A0803B1723542889537EFB24D6E0844004248B1980292D608D00423F49F9908049798B4452C0131006230C14868200FC668B50650043196A7F95569CF6B663341535DCFE919C464400A96DCE1C6B96D5EEFE60096006A400087C1E8610A4401887D1863AC99F9802DC00D34B5BCD72D6F36CB6E7D95EBC600013A88010A8271B6281803B12E124633006A2AC3A8AC600BCD07C9851008712DEAE83A802929DC51EE5EF5AE61BCD0648028596129C3B98129E5A9A329ADD62CCE0164DDF2F9343135CCE2137094A620E53FACF37299F0007392A0B2A7F0BA5F61B3349F3DFAEDE8C01797BD3F8BC48740140004322246A8A2200CC678651AA46F09AEB80191940029A9A9546E79764F7C9D608EA0174B63F815922999A84CE7F95C954D7FD9E0890047D2DC13B0042488259F4C0159922B0046565833828A00ACCD63D189D4983E800AFC955F211C700"""


def pad_convert(hexa):
    h = bin(int(hexa, 16))[2:]
    return '0' * (4 - len(h)) + h


bin_data = ''.join(list(map(pad_convert, data_str1)))

print(bin_data)

res = 0


def calc(nrs, type):
    if type == 0:
        return sum(nrs)
    if type == 1:
        return prod(nrs)
    if type == 2:
        return min(nrs)
    if type == 3:
        return max(nrs)
    if type == 5:
        return int(nrs[0] > nrs[1])
    if type == 6:
        return int(nrs[0] < nrs[1])
    if type == 7:
        return int(nrs[0] == nrs[1])


def parse(binary):
    global res

    if len(binary) <= 7:
        return ('', -1)

    version = int(binary[0:3], 2)
    print(f'Version {version}')
    res += version

    type = int(binary[3:6], 2)

    i = 6
    if type == 4:  # literal type
        b = ''
        while binary[i] == '1':
            b += binary[i + 1: i + 5]
            i += 5
        # marker bit is 0 now
        b += binary[i + 1:i+5]
        nr = int(b, 2)
        print('int b ->', nr)

        return (binary[i + 5:], nr)
    else:  # operator
        length_type = int(binary[6], 2)
        i = 7
        nrs = []
        if length_type == 0:
            bit_count_left = int(binary[i: i + 15], 2)
            print(
                f'Length Type id is {length_type} -> bit left count: {bit_count_left}')
            b = binary[i + 15: i + 15 + bit_count_left]
            while len(b) > 0:
                b, nr = parse(b)
                if b != '' or nr != -1:
                    nrs.append(nr)

            result = calc(nrs, type)
            return (binary[i + 15 + bit_count_left:], result)
        elif length_type == 1:
            subpacket_count = int(binary[i: i + 11], 2)
            print(
                f'Length Type id is {length_type} -> sub packet count: {subpacket_count}')
            c = 0
            b = binary[i + 11:]
            while c < subpacket_count:
                b, nr = parse(b)
                nrs.append(nr)
                c += 1

            result = calc(nrs, type)
            return (b, result)


print()
print(parse(bin_data)[1])

print('version sum -> ', res)
