# --- Day 7: No Space Left On Device - --

# You can hear birds chirping and raindrops hitting leaves as the expedition proceeds. Occasionally, you can even hear much louder sounds in the distance
# how big do the animals get out here, anyway?

# The device the Elves gave you has problems with more than just its communication system. You try to run a system update:

# $ system-update - -please - -pretty-please-with -sugar-on-top
# Error: No space left on device
# Perhaps you can delete some files to make space for the update?

# You browse around the filesystem to assess the situation and save the resulting terminal output(your puzzle input). For example:

# $ cd /
# $ ls
# dir a
# 14848514 b.txt
# 8504156 c.dat
# dir d
# $ cd a
# $ ls
# dir e
# 29116 f
# 2557 g
# 62596 h.lst
# $ cd e
# $ ls
# 584 i
# $ cd ..
# $ cd ..
# $ cd d
# $ ls
# 4060174 j
# 8033020 d.log
# 5626152 d.ext
# 7214296 k
# The filesystem consists of a tree of files(plain data) and directories(which can contain other directories or files). The outermost directory is called / . You can navigate around the filesystem, moving into or out of directories and listing the contents of the directory you're currently in .

# Within the terminal output, lines that begin with $ are commands you executed, very much like some modern computers:

# cd means change directory. This changes which directory is the current directory, but the specific result depends on the argument:
# cd x moves in one level: it looks in the current directory for the directory named x and makes it the current directory.
# cd .. moves out one level: it finds the directory that contains the current directory, then makes that directory the current directory.
# cd / switches the current directory to the outermost directory, / .
# ls means list. It prints out all of the files and directories immediately contained by the current directory:
# 123 abc means that the current directory contains a file named abc with size 123.
# dir xyz means that the current directory contains a directory named xyz.
# Given the commands and output in the example above, you can determine that the filesystem looks visually like this:

# - / (dir)
#  - a(dir)
#    - e(dir)
#      - i(file, size=584)
#     - f(file, size=29116)
#     - g(file, size=2557)
#     - h.lst(file, size=62596)
#   - b.txt(file, size=14848514)
#   - c.dat(file, size=8504156)
#   - d(dir)
#    - j (file, size=4060174)
#     - d.log(file, size=8033020)
#     - d.ext(file, size=5626152)
#     - k(file, size=7214296)
# Here, there are four directories: / (the outermost directory), a and d(which are in /), and e (which is in a). These directories also contain files of various sizes.

# Since the disk is full, your first step should probably be to find directories that are good candidates for deletion. To do this, you need to determine the total size of each directory. The total size of a directory is the sum of the sizes of the files it contains, directly or indirectly. (Directories themselves do not count as having any intrinsic size.)

# The total sizes of the directories above can be found as follows:

# The total size of directory e is 584 because it contains a single file i of size 584 and no other directories.
# The directory a has total size 94853 because it contains files f(size 29116), g (size 2557), and h.lst (size 62596), plus file i indirectly (a contains e which contains i).
# Directory d has total size 24933642.
# As the outermost directory, / contains every file. Its total size is 48381165, the sum of the size of every file.
# To begin, find all of the directories with a total size of at most 100000, then calculate the sum of their total sizes. In the example above, these directories are a and e
# the sum of their total sizes is 95437 (94853 + 584). (As in this example, this process can count files more than once!)

# Find all of the directories with a total size of at most 100000. What is the sum of the total sizes of those directories?

# Your puzzle answer was 1118405.

from collections import Counter
import math
import timeit
start_time = timeit.default_timer()

data_str1 = """$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k"""

data_str = """$ cd /
$ ls
dir bqm
dir ctztn
dir dbclg
dir fhndmnt
dir gczqbh
276177 hvbf.lvm
dir lnsgbqp
dir pblb
dir pwfs
209719 rtv.cjj
192236 vmtnnfv.mdq
dir vmvpf
dir wjgh
dir wjqsq
$ cd bqm
$ ls
133711 vqv
263237 wwlv.vgv
$ cd ..
$ cd ctztn
$ ls
dir gpgbfzbs
dir hfvqpt
55466 jshtffs
dir rgzfgz
115519 wdzh.szq
dir wmfvclz
221554 zgwwgps.gfn
$ cd gpgbfzbs
$ ls
354710 tlvmh.ghp
$ cd ..
$ cd hfvqpt
$ ls
dir cggfgt
dir jshtffs
175230 nnn.chs
13406 plqjpqss
dir pnwcq
dir str
$ cd cggfgt
$ ls
109434 fhww
dir glnrg
dir mptpfvlh
305315 nlhfgpr.vnv
297519 tlvmh.ghp
dir vmvpf
$ cd glnrg
$ ls
345026 qzfpwv
196361 rgzfgz
$ cd ..
$ cd mptpfvlh
$ ls
dir lsbjp
121174 vlbbbgnn
$ cd lsbjp
$ ls
dir jqclpq
359939 nlhfgpr
dir whp
$ cd jqclpq
$ ls
189319 mnhnclpt.qsv
$ cd ..
$ cd whp
$ ls
dir lgdlztb
$ cd lgdlztb
$ ls
94150 wdzh.szq
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd vmvpf
$ ls
dir pfzd
66706 rgzfgz
dir zrc
$ cd pfzd
$ ls
345595 bpjpm
dir ljtdcgt
dir plltpvgv
$ cd ljtdcgt
$ ls
134178 nlhfgpr
$ cd ..
$ cd plltpvgv
$ ls
355576 rqj
$ cd ..
$ cd ..
$ cd zrc
$ ls
45363 vqv
$ cd ..
$ cd ..
$ cd ..
$ cd jshtffs
$ ls
dir bwbwldd
124375 fwvcbcpz.lvr
dir jctctjj
dir jshtffs
dir sngchq
$ cd bwbwldd
$ ls
144330 vthsfg
$ cd ..
$ cd jctctjj
$ ls
dir flhrmt
dir rgzfgz
dir vmtnnfv
$ cd flhrmt
$ ls
132008 dgrt
263345 nlhfgpr.hss
$ cd ..
$ cd rgzfgz
$ ls
214808 mcqlbzs.gwl
120157 nlhfgpr.hcs
$ cd ..
$ cd vmtnnfv
$ ls
dir rgzfgz
52978 vlbbbgnn
$ cd rgzfgz
$ ls
177061 nnc
$ cd ..
$ cd ..
$ cd ..
$ cd jshtffs
$ ls
157488 bbsfrg.bbr
166412 dhwbvggg
$ cd ..
$ cd sngchq
$ ls
dir bvcsvv
dir hhd
$ cd bvcsvv
$ ls
dir gbmccf
5663 vmtnnfv
278760 vqv
323147 wjlm.cbw
$ cd gbmccf
$ ls
dir fttnqbwp
$ cd fttnqbwp
$ ls
dir fhs
$ cd fhs
$ ls
137718 rgzfgz
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd hhd
$ ls
350019 gphghhdh.rrb
$ cd ..
$ cd ..
$ cd ..
$ cd pnwcq
$ ls
dir qtqqdd
$ cd qtqqdd
$ ls
62319 tlvmh.ghp
$ cd ..
$ cd ..
$ cd str
$ ls
334780 tlvmh.ghp
dir vcdcgbtr
211478 vvh.tzh
dir zlsv
$ cd vcdcgbtr
$ ls
276438 zhh.dsw
$ cd ..
$ cd zlsv
$ ls
315890 vmvpf.tcs
$ cd ..
$ cd ..
$ cd ..
$ cd rgzfgz
$ ls
dir qwzfs
$ cd qwzfs
$ ls
dir lzl
$ cd lzl
$ ls
dir gpgccbc
259493 vmvpf.dgz
$ cd gpgccbc
$ ls
124141 tlvmh.ghp
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd wmfvclz
$ ls
25819 fmqtjdfd.dwp
253619 qccd.gnd
121087 sjsz.rtg
350518 tgs
dir zmft
$ cd zmft
$ ls
248086 vlbbbgnn
$ cd ..
$ cd ..
$ cd ..
$ cd dbclg
$ ls
dir dgm
dir rwdfw
182454 slctrj
103919 tzp.dvf
311601 zlbmvfq.qsh
$ cd dgm
$ ls
dir fflb
dir jshtffs
26424 shzfvbs.jrj
$ cd fflb
$ ls
154495 ngzgfpt
$ cd ..
$ cd jshtffs
$ ls
11509 fwvcbcpz.lvr
169064 hzhbhgjn.dvc
38675 jwcnwm.prf
dir mjb
dir nwwcb
300017 qtrqn.qdv
dir rfjdh
173691 rgzfgz.wgz
$ cd mjb
$ ls
dir rgzfgz
$ cd rgzfgz
$ ls
192387 ssm
$ cd ..
$ cd ..
$ cd nwwcb
$ ls
312609 ggvwvdw.hhg
107482 ppqmjzh
204253 vlbbbgnn
$ cd ..
$ cd rfjdh
$ ls
305213 smf.jgv
$ cd ..
$ cd ..
$ cd ..
$ cd rwdfw
$ ls
dir hbsbctsr
145955 tqj
dir vvthg
$ cd hbsbctsr
$ ls
dir shqdr
$ cd shqdr
$ ls
101441 tlvmh.ghp
$ cd ..
$ cd ..
$ cd vvthg
$ ls
344816 wdzh.szq
$ cd ..
$ cd ..
$ cd ..
$ cd fhndmnt
$ ls
dir lstcq
dir nvlncq
dir rgzfgz
dir ssqs
dir vzpnjpsv
$ cd lstcq
$ ls
24507 jshtffs.nrf
34332 lvqdthm.gwh
313126 szfqhv.jtv
274515 wmwpglh
$ cd ..
$ cd nvlncq
$ ls
dir vmtnnfv
$ cd vmtnnfv
$ ls
104013 wdzh.szq
$ cd ..
$ cd ..
$ cd rgzfgz
$ ls
311945 nlhfgpr.fcv
$ cd ..
$ cd ssqs
$ ls
104254 vlbbbgnn
207441 vmpd.tdt
$ cd ..
$ cd vzpnjpsv
$ ls
253593 plfsrlr
$ cd ..
$ cd ..
$ cd gczqbh
$ ls
dir nlhfgpr
dir rgzfgz
dir vmtnnfv
$ cd nlhfgpr
$ ls
dir bwpd
86459 vmvpf.mwh
$ cd bwpd
$ ls
345948 rgzfgz.qnp
$ cd ..
$ cd ..
$ cd rgzfgz
$ ls
dir rgzfgz
dir sdsqqq
$ cd rgzfgz
$ ls
dir mprdrmmz
331389 rbcrg.chs
191328 rpvpnprr.smn
$ cd mprdrmmz
$ ls
dir ncljgf
$ cd ncljgf
$ ls
309996 bjwvhw.ltr
$ cd ..
$ cd ..
$ cd ..
$ cd sdsqqq
$ ls
95157 qmwcb.wsm
$ cd ..
$ cd ..
$ cd vmtnnfv
$ ls
dir gbntbhrj
dir nlhfgpr
283985 ptl
dir tnljzft
38026 vlbbbgnn
$ cd gbntbhrj
$ ls
117782 jshtffs.fpb
9101 rgzfgz.bdp
300155 vmtnnfv.wcq
$ cd ..
$ cd nlhfgpr
$ ls
dir cpglmt
65206 cwmpnz.czl
dir gcrntzb
$ cd cpglmt
$ ls
288805 mghgq.msj
$ cd ..
$ cd gcrntzb
$ ls
126442 tlvmh.ghp
$ cd ..
$ cd ..
$ cd tnljzft
$ ls
278719 ccwfvtgh
$ cd ..
$ cd ..
$ cd ..
$ cd lnsgbqp
$ ls
dir lqltnz
dir vmtnnfv
69119 vmvpf.rgd
$ cd lqltnz
$ ls
dir vmtnnfv
dir whvdf
$ cd vmtnnfv
$ ls
3012 vmtnnfv.zlv
$ cd ..
$ cd whvdf
$ ls
dir nlhfgpr
115838 rwqbdmfb
298358 tlvmh.ghp
dir vmtnnfv
335379 vmtnnfv.ltj
129665 zglmsgf
$ cd nlhfgpr
$ ls
12607 fwvcbcpz.lvr
dir lrbrcz
dir nvqzmflc
3793 tlvmh.ghp
39593 vqv
$ cd lrbrcz
$ ls
195996 vmtnnfv.hlb
$ cd ..
$ cd nvqzmflc
$ ls
190198 vmvpf
$ cd ..
$ cd ..
$ cd vmtnnfv
$ ls
359558 rdjqs
$ cd ..
$ cd ..
$ cd ..
$ cd vmtnnfv
$ ls
dir clwrz
362269 nphz.plr
dir rqbq
188843 tlvmh.ghp
dir wsrdwm
$ cd clwrz
$ ls
233770 wdzh.szq
$ cd ..
$ cd rqbq
$ ls
183770 jtnlfd
71164 lwgqzccd.dgb
$ cd ..
$ cd wsrdwm
$ ls
266826 bqg
dir nlhfgpr
53502 nlhfgpr.slz
320422 qzsfvzl.wbb
dir vhq
dir vmvpf
$ cd nlhfgpr
$ ls
dir ddlplnhp
dir nbdj
$ cd ddlplnhp
$ ls
dir czp
dir dmchbv
dir nlj
299460 tlvmh.ghp
dir wrqqln
$ cd czp
$ ls
286886 tlvmh.ghp
98020 vlln
$ cd ..
$ cd dmchbv
$ ls
239186 ppjcgwq
$ cd ..
$ cd nlj
$ ls
310042 nsgvvfcw.vfm
70206 vlbbbgnn
$ cd ..
$ cd wrqqln
$ ls
dir bdprgbp
$ cd bdprgbp
$ ls
42055 hjwrqzf.cqj
$ cd ..
$ cd ..
$ cd ..
$ cd nbdj
$ ls
dir cdjd
$ cd cdjd
$ ls
330601 bzhvfzh.ldp
14662 hjhdwz.qrt
188113 pct
$ cd ..
$ cd ..
$ cd ..
$ cd vhq
$ ls
348147 qsgjjzn.nft
$ cd ..
$ cd vmvpf
$ ls
192773 wdzh.szq
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd pblb
$ ls
dir jshtffs
dir pwgtvhm
$ cd jshtffs
$ ls
dir vtthv
$ cd vtthv
$ ls
39542 fwfz.gll
$ cd ..
$ cd ..
$ cd pwgtvhm
$ ls
271449 hvfq
$ cd ..
$ cd ..
$ cd pwfs
$ ls
dir dnnzr
dir ncjzvsph
dir nlhfgpr
$ cd dnnzr
$ ls
337565 vlbbbgnn
345668 vqv
$ cd ..
$ cd ncjzvsph
$ ls
290438 fdtftjj.sdt
170963 tlvmh.ghp
56596 vlbbbgnn
72923 vmvpf.ldg
$ cd ..
$ cd nlhfgpr
$ ls
321213 vqv
$ cd ..
$ cd ..
$ cd vmvpf
$ ls
dir gwbt
dir vmvpf
dir vplm
$ cd gwbt
$ ls
209996 jrzlrp
23503 vmvpf.hwn
173320 vqv
$ cd ..
$ cd vmvpf
$ ls
299874 hqwppnws
177211 pcqdrn.wqd
266011 pjllm
227969 wdzh.szq
$ cd ..
$ cd vplm
$ ls
dir fbl
354786 jdb
273253 nncwhbq.vzn
dir rgzfgz
$ cd fbl
$ ls
38301 vqv
256284 wdzh.szq
$ cd ..
$ cd rgzfgz
$ ls
60018 rdsrl.whb
dir rgzfgz
dir wgvw
$ cd rgzfgz
$ ls
dir gmcqfqrn
$ cd gmcqfqrn
$ ls
337129 rgzfgz.pjs
$ cd ..
$ cd ..
$ cd wgvw
$ ls
14727 zdhwsm
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd wjgh
$ ls
dir ccgddrt
dir cqh
60305 dtjctmsw.qgp
dir hwr
dir jshtffs
335879 smfjq
344915 vlbbbgnn
63031 wdzh.szq
dir zbmjg
$ cd ccgddrt
$ ls
148368 dpnlbg.fbb
dir vmvpf
$ cd vmvpf
$ ls
dir ptwrssn
340536 tlvmh.ghp
$ cd ptwrssn
$ ls
119875 jshtffs
$ cd ..
$ cd ..
$ cd ..
$ cd cqh
$ ls
dir blvcjl
158276 fjv
279848 fnh.rdf
dir vmtnnfv
$ cd blvcjl
$ ls
dir cgsznlm
47381 gtcv
227459 rgzfgz.wvr
dir tgcqhcv
dir vff
dir vmtnnfv
dir vmvpf
75937 vmvpf.psj
29291 wdzh.szq
dir zqmsn
$ cd cgsznlm
$ ls
91350 rpmwrtm.dlr
$ cd ..
$ cd tgcqhcv
$ ls
dir hnrhdvff
81404 jshtffs.msp
158082 jshtffs.szt
dir nhmqcm
dir rvfsl
dir smcrnh
$ cd hnrhdvff
$ ls
154674 htwms.nqc
dir rslbgns
192997 vlbbbgnn
dir vmvpf
230969 vqv
195769 ztfg.mfd
$ cd rslbgns
$ ls
236439 dwbjh.whr
dir gzph
279480 tlvmh.ghp
107447 vqv
$ cd gzph
$ ls
46947 tdnmmqm
$ cd ..
$ cd ..
$ cd vmvpf
$ ls
285442 sftfzz.pzn
292269 vmtnnfv.bbc
$ cd ..
$ cd ..
$ cd nhmqcm
$ ls
219245 hwpgpdm.qrf
130169 rsmdgwjz.rph
$ cd ..
$ cd rvfsl
$ ls
263768 ctprwbl.tjt
dir jshtffs
dir rgzfgz
291351 vpcd.hvf
$ cd jshtffs
$ ls
9465 fwvcbcpz.lvr
dir htb
dir jshtffs
202004 nrgq.cwj
dir vmvpf
$ cd htb
$ ls
323938 gwhlttb
75918 llmlgmqz.zsp
dir nvfv
117786 qjjv
167090 vmvpf.qvd
$ cd nvfv
$ ls
135820 tlvmh.ghp
$ cd ..
$ cd ..
$ cd jshtffs
$ ls
240326 fhr.jbf
$ cd ..
$ cd vmvpf
$ ls
154309 fwvcbcpz.lvr
$ cd ..
$ cd ..
$ cd rgzfgz
$ ls
dir jjfdn
dir nlhfgpr
$ cd jjfdn
$ ls
46675 vfwhjfl.zrz
$ cd ..
$ cd nlhfgpr
$ ls
26537 qwzcnfz
$ cd ..
$ cd ..
$ cd ..
$ cd smcrnh
$ ls
331015 vpttwwc.jgr
$ cd ..
$ cd ..
$ cd vff
$ ls
68000 ntbfs.znj
265866 rgzfgz.lsl
53570 vmvpf.ncb
74229 vqv
96730 zwq.zdb
$ cd ..
$ cd vmtnnfv
$ ls
dir rgzfgz
dir zmr
$ cd rgzfgz
$ ls
7492 fwvcbcpz.lvr
dir gjg
dir sfrlhmm
$ cd gjg
$ ls
dir dqjfz
59560 vqv
33788 wdzh.szq
$ cd dqjfz
$ ls
dir zhtbsq
$ cd zhtbsq
$ ls
222414 cgsjvqvd.rsd
$ cd ..
$ cd ..
$ cd ..
$ cd sfrlhmm
$ ls
94423 rgzfgz
$ cd ..
$ cd ..
$ cd zmr
$ ls
269611 gvhrsmt.bdt
dir vmtnnfv
$ cd vmtnnfv
$ ls
dir rgzfgz
$ cd rgzfgz
$ ls
355703 fwvcbcpz.lvr
286378 pcf.dss
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd vmvpf
$ ls
336533 fwvcbcpz.lvr
80609 hsfnp
9548 jld.rvt
dir jshtffs
219383 rjtrbjl.vbc
dir vmtnnfv
182036 vmvpf
$ cd jshtffs
$ ls
dir jshtffs
dir rtnwnqcb
358722 szfhzcj.tjt
167284 tlvmh.ghp
$ cd jshtffs
$ ls
68702 blcgmn
299911 jjd.pvb
243135 wdzh.szq
$ cd ..
$ cd rtnwnqcb
$ ls
128787 jmvlgd.bwv
268582 tdpqcqfs.nzp
348305 vmvpf.bnl
$ cd ..
$ cd ..
$ cd vmtnnfv
$ ls
dir nlhfgpr
241246 tlvmh.ghp
$ cd nlhfgpr
$ ls
292694 tlvmh.ghp
$ cd ..
$ cd ..
$ cd ..
$ cd zqmsn
$ ls
244052 mmthhf.wmd
144084 psspc.czf
dir vflpgd
$ cd vflpgd
$ ls
22881 hsbbcdp.psd
dir rgzfgz
dir tmz
246706 vmvpf
$ cd rgzfgz
$ ls
346162 jshtffs.hqw
dir lhmh
dir pfwbz
87008 tzddth.whp
$ cd lhmh
$ ls
dir nlhfgpr
$ cd nlhfgpr
$ ls
167305 tlvmh.ghp
dir vnsdh
$ cd vnsdh
$ ls
22714 tlvmh.ghp
$ cd ..
$ cd ..
$ cd ..
$ cd pfwbz
$ ls
163433 bzttjhqz.hpz
198604 vlbbbgnn
$ cd ..
$ cd ..
$ cd tmz
$ ls
300571 bmrs
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd vmtnnfv
$ ls
dir llwmdf
dir pvzdlsgp
48025 zvvrwh.dhw
$ cd llwmdf
$ ls
236755 hbfdll
$ cd ..
$ cd pvzdlsgp
$ ls
dir drjbqvn
102548 fwvcbcpz.lvr
300337 gcv.bnn
dir grgc
27918 mtnvzzc.vlj
dir nlhfgpr
92639 wdzh.szq
152813 wmnpg
$ cd drjbqvn
$ ls
292541 jshtffs.gjg
84918 vqv
355763 zqsmd
$ cd ..
$ cd grgc
$ ls
45875 nlhfgpr.rnf
294562 slshp
$ cd ..
$ cd nlhfgpr
$ ls
310430 fpglshq
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd hwr
$ ls
dir jshtffs
dir srrgj
$ cd jshtffs
$ ls
275253 fwvcbcpz.lvr
262142 vlqflmdn.gtd
dir zrm
$ cd zrm
$ ls
103078 jshtffs.ncf
$ cd ..
$ cd ..
$ cd srrgj
$ ls
328221 fqdmfshj.ftn
107311 lpdrrw
214681 ltq.gmq
321175 pzd
$ cd ..
$ cd ..
$ cd jshtffs
$ ls
dir ngbb
$ cd ngbb
$ ls
218170 lvmj.sjg
123373 nwv
326485 rgzfgz.gjd
192138 vlbbbgnn
dir vmtnnfv
179027 vptgrfr
$ cd vmtnnfv
$ ls
206046 vnlsclvh
$ cd ..
$ cd ..
$ cd ..
$ cd zbmjg
$ ls
30152 fwvcbcpz.lvr
$ cd ..
$ cd ..
$ cd wjqsq
$ ls
23196 cmvfcwp.nfz
dir lnrzw
316798 nlhfgpr.spt
dir ntrngh
284657 tlvmh.ghp
50758 vmvpf.msg
$ cd lnrzw
$ ls
dir crc
dir ddhnvwqt
dir gdhfqlln
dir vmtnnfv
$ cd crc
$ ls
296423 vmtnnfv.prq
26320 vqv
$ cd ..
$ cd ddhnvwqt
$ ls
173923 vqv
$ cd ..
$ cd gdhfqlln
$ ls
275288 vqv
$ cd ..
$ cd vmtnnfv
$ ls
52004 pgtmmtrf.vfc
$ cd ..
$ cd ..
$ cd ntrngh
$ ls
dir czt
185012 vlbbbgnn
$ cd czt
$ ls
355964 cvnslv
143166 mwbcbphd.prr
40666 rgzfgz"""

lines = data_str1.split('\n')


class File(object):
    def __init__(self, size: int, name: str):
        self.size = size
        self.name = name

    def get_size(self):
        return self.size


class Dir(object):
    def __init__(self, name: str, parent):
        self.children = []
        self.name = name
        self.parent = parent

    def get_size(self):
        return sum(map(lambda child: child.get_size(), self.children))


tree = Dir('/', None)
it = tree

iscmd = True

for line in lines[1:]:
    print(line)
    cmd_split = line.split(' ')
    print(cmd_split)

    if not iscmd:
        if cmd_split[0] == '$':
            iscmd = True
        else:
            if cmd_split[0] == 'dir':
                it.children.append(Dir(cmd_split[1], it))
            elif cmd_split[0].isnumeric():
                it.children.append(File(int(cmd_split[0]), cmd_split[1]))
            else:
                print('ERROR: not a file or a dir')
            continue

    cmd = cmd_split[1]

    if cmd == 'ls':
        iscmd = False
    elif cmd == 'cd':
        if (cmd_split[2] == '..'):
            if it.parent == None:
                print(f'ERROR: {it.name} has no parent (prolly it is /)')
            it = it.parent
            continue

        for child in it.children:
            if child.name == cmd_split[2]:
                it = child
                break
            print('ERROR: file not found', cmd_split[2])
    else:
        print('ERROR: cmd not found', cmd)
    print()


def print_dir(dir_or_file, indent=0):
    if type(dir_or_file) == Dir:
        print(indent * ' ', '- ', dir_or_file.name, '(dir)')
        for child in dir_or_file.children:
            print_dir(child, indent + 2)
    elif type(dir_or_file) == File:
        print(indent * ' ', '- ', dir_or_file.name,
              '(file, ', dir_or_file.size, ')')
    else:
        print('ERROR: type not known', type(dir_or_file), dir_or_file)

dirs_and_sizes = []

def collect_dir_sizes(dir_or_file):
    if type(dir_or_file) == Dir:
        dirs_and_sizes.append([dir_or_file.name, dir_or_file.get_size()])
        for child in dir_or_file.children:
            collect_dir_sizes(child)
    elif type(dir_or_file) == File:
        return
    else:
        print('ERROR: type not known', type(dir_or_file), dir_or_file)


print_dir(tree)
collect_dir_sizes(tree)
print('dirs_and_sizes', dirs_and_sizes)

dirs_and_sizes_filtered = list(filter(lambda dir: dir[0] !=
                                      '/' and dir[1] <= 100000, dirs_and_sizes))

print(dirs_and_sizes_filtered)

print(sum(map(lambda dir: dir[1], dirs_and_sizes_filtered)))

stop = timeit.default_timer()
print(f'\n✨Time: {stop - start_time} s')
