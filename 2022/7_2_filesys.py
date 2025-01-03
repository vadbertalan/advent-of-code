# --- Part Two - --

# Now, you're ready to choose a directory to delete.

# The total disk space available to the filesystem is 70000000. To run the update, you need unused space of at least 30000000. You need to find a directory you can delete that will free up enough space to run the update.

# In the example above, the total size of the outermost directory (and thus the total amount of used space) is 48381165
# this means that the size of the unused space must currently be 21618835, which isn't quite the 30000000 required by the update. Therefore, the update still requires a directory with total size of at least 8381165 to be deleted before it can run.

# To achieve this, you have the following options:

# Delete directory e, which would increase unused space by 584.
# Delete directory a, which would increase unused space by 94853.
# Delete directory d, which would increase unused space by 24933642.
# Delete directory /, which would increase unused space by 48381165.
# Directories e and a are both too small
# deleting them would not free up enough space. However, directories d and / are both big enough! Between these, choose the smallest: d, increasing unused space by 24933642.

# Find the smallest directory that, if deleted, would free up enough space on the filesystem to run the update. What is the total size of that directory?

# tried 50216456 but wrong

# Your puzzle answer was 12545514.

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

lines = data_str.split('\n')


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

space_needed_to_update = 30000000
space_available = 70000000 - dirs_and_sizes[0][1]
space_required = space_needed_to_update - space_available
print('space needed = ', space_needed_to_update)
print('space available = ', space_available)
print('space required = ', space_required)

dirs_and_sizes_filtered_and_sorted = sorted(filter(
    lambda dir: dir[1] >= space_required, dirs_and_sizes), key=lambda dir: dir[1])
print(dirs_and_sizes_filtered_and_sorted)
print(dirs_and_sizes_filtered_and_sorted[0][1])


stop = timeit.default_timer()
print(f'\n✨Time: {stop - start_time} s')
