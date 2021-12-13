# --- Day 13: Transparent Origami - --

# You reach another volcanically active part of the cave. It would be nice if you could do some kind of thermal imaging so you could tell ahead of time which caves are too hot to safely enter.

# Fortunately, the submarine seems to be equipped with a thermal camera! When you activate it, you are greeted with:

# Congratulations on your purchase! To activate this infrared thermal imaging
# camera system, please enter the code found on page 1 of the manual.
# Apparently, the Elves have never used this feature. To your surprise, you manage to find the manual
# as you go to open it, page 1 falls out. It's a large sheet of transparent paper! The transparent paper is marked with random dots and includes instructions on how to fold it up(your puzzle input). For example:

# 6, 10
# 0, 14
# 9, 10
# 0, 3
# 10, 4
# 4, 11
# 6, 0
# 6, 12
# 4, 1
# 0, 13
# 10, 12
# 3, 4
# 3, 0
# 8, 4
# 1, 10
# 2, 14
# 8, 10
# 9, 0

# fold along y = 7
# fold along x = 5
# # is a dot on the paper and . is an empty, unmarked position:
# The first section is a list of dots on the transparent paper. 0, 0 represents the top-left coordinate. The first value, x, increases to the right. The second value, y, increases downward. So, the coordinate 3, 0 is to the right of 0, 0, and the coordinate 0, 7 is below 0, 0. The coordinates in this example form the following pattern, where

# ...  # ..#..#.
# ....  # ......
# ...........
# #..........
# ...  # ....#.#
# ...........
# ...........
# ...........
# ...........
# ...........
# .  # ....#.##.
# ....  # ......
# ......  # ...#
# #..........
# #.#........
# Then, there is a list of fold instructions. Each instruction indicates a line on the transparent paper and wants you to fold the paper up(for horizontal y=... lines) or left(for vertical x=... lines). In this example, the first fold instruction is fold along y = 7, which designates the line formed by all of the positions where y is 7 (marked here with -):

# ...  # ..#..#.
# ....  # ......
# ...........
# #..........
# ...  # ....#.#
# ...........
# ...........
# -----------
# ...........
# ...........
# .  # ....#.##.
# ....  # ......
# ......  # ...#
# #..........
# #.#........
# Because this is a horizontal line, fold the bottom half up. Some of the dots might end up overlapping after the fold is complete, but dots will never appear exactly on a fold line. The result of doing this fold looks like this:

#     #.##..#..#.
#     #...#......
# ......  # ...#
# #...#......
# .  # .#..#.###
# ...........
# ...........
# Now, only 17 dots are visible.

# Notice, for example, the two dots in the bottom left corner before the transparent paper is folded
# after the fold is complete, those dots appear in the top left corner(at 0, 0 and 0, 1). Because the paper is transparent, the dot just below them in the result(at 0, 3) remains visible, as it can be seen through the transparent paper.

# Also notice that some dots can end up overlapping
# in this case, the dots merge together and become a single dot.

# The second fold instruction is fold along x = 5, which indicates this line:

#     #.##.|#..#.
#     #...#|.....
# .....|  # ...#
# #...#|.....
# .  # .#.|#.###
# ..... | .....
# ..... | .....
# Because this is a vertical line, fold left:

#     #####
#     #...#
#     #...#
#     #...#
#     #####
# .....
# .....
# The instructions made a square!

# The transparent paper is pretty big, so for now, focus on just completing the first fold. After the first fold in the example above, 17 dots are visible - dots that end up overlapping after the fold is completed count as a single dot.

# How many dots are visible after completing just the first fold instruction on your transparent paper?

# Your puzzle answer was 607.

data_str1 = """6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5"""

data_str = """1302,268
872,390
164,653
1257,768
649,367
314,747
239,246
572,431
610,427
495,876
321,8
457,553
1174,140
470,728
216,266
142,502
611,831
475,630
216,504
845,23
99,427
1037,334
381,112
616,859
982,876
561,322
249,99
987,652
793,729
606,579
1017,140
679,81
544,0
1274,259
1235,246
1200,266
903,588
807,103
17,455
527,829
512,144
523,37
976,140
662,490
570,180
1159,791
816,271
740,266
149,805
570,266
410,92
1235,462
363,45
1004,362
331,204
724,462
845,247
395,138
758,816
939,511
472,189
184,750
1295,134
513,745
735,140
1179,37
1009,135
331,690
704,355
502,498
274,614
308,427
216,180
5,864
1287,305
1046,859
807,327
117,229
326,18
1093,840
572,802
1253,683
262,721
786,679
987,764
1061,99
610,467
473,775
53,560
1126,840
267,885
792,82
631,253
704,539
0,360
1216,14
60,428
552,485
1295,312
199,437
206,91
1262,355
989,310
1292,733
982,18
1044,416
407,588
1126,40
556,609
1208,814
766,378
1049,259
23,589
1101,480
1223,44
681,324
1206,544
306,114
1146,653
503,327
1293,439
853,553
463,376
959,529
490,260
783,491
596,208
750,880
447,287
15,299
431,434
131,37
934,763
947,334
73,521
8,178
935,278
136,562
1009,725
740,180
648,490
415,716
311,376
730,801
115,392
952,145
1076,732
1171,365
1245,850
792,562
32,528
465,247
771,112
353,583
1235,866
246,802
922,245
1250,390
1140,728
102,78
422,14
487,242
1186,747
597,659
539,502
596,462
1250,652
438,504
689,306
75,246
699,831
328,86
184,502
552,814
703,40
351,164
977,397
638,173
808,498
691,437
0,260
480,274
725,829
115,502
393,655
875,471
967,749
115,54
1278,366
704,864
816,607
435,378
1217,572
527,491
592,35
430,256
1119,152
1026,399
1071,178
196,782
136,82
1019,605
103,885
430,394
713,235
142,392
274,448
189,647
346,285
734,273
468,504
539,560
991,728
1243,532
206,763
455,894
1231,588
320,364
363,114
15,312
1285,311
815,876
773,99
448,383
117,99
488,523
838,189
103,457
982,428
1190,131
552,816
797,689
388,245
207,831
1151,876
1288,56
1299,712
246,702
199,289
420,162
952,222
1216,674
758,500
315,227
494,175
80,683
261,259
1124,56
199,448
189,23
1164,273
813,79
1156,738
552,500
716,166
151,791
488,607
127,558
166,523
982,466
141,689
1131,626
868,366
442,427
167,169
1094,448
67,532
662,404
642,539
1009,759
1160,544
924,583
552,526
909,712
907,866
147,268
88,812
380,714
321,30
93,322
104,350
311,33
1143,690
1174,82
1248,859
216,740
25,324
35,347
964,161
258,450
435,277
239,268
487,169
967,319
606,864
930,894
534,691
900,463
1310,534
646,626
18,691
223,605
75,169
1144,623
749,322
283,511
1272,35
1310,708
196,231
606,366
492,241
1128,705
341,537
176,450
1310,260
1144,355
1203,337
977,721
321,702
171,633
154,604
542,481
306,810
259,756
222,798
375,165
718,467
77,771
1151,18
1208,409
130,231
818,241
959,753
1093,56
1248,203
964,61
773,638
315,219
631,753
1144,523
1143,242
217,838
1298,763
405,771
407,754
1248,467
979,204
1290,207
683,422
1258,796
1036,796
590,497
331,242
119,595
566,82
1158,498
759,12
1255,535
455,885
1287,813
346,609
754,609
813,169
1302,716
1169,689
915,756
1159,103
1007,365
885,149
185,205
483,709
1146,634
689,588
1305,864
997,537
1061,795
1009,583
3,729
688,560
469,689
979,242
562,63
49,54
1111,894
606,752
303,365
694,35
480,620
25,886
94,462
977,49
217,86
1052,674
503,119
1235,28
813,290
520,882
649,30
261,689
213,168
171,185
497,815
75,28
949,677
1049,595
246,876
8,268
1141,277
850,224
411,735
78,621
27,248
503,103
1292,691
371,511
151,103
1134,831
177,253
11,787
1001,7
458,231
1119,742
651,242
157,613
995,103
124,147
729,854
731,648
482,725
348,431
1076,511
216,448
403,14
261,595
12,819
1279,821
64,497
638,721
1144,47
15,134
259,585
179,268
964,285
1086,534
930,0
480,572
1067,735
1104,355
94,880
181,567
629,501
257,124
629,214
560,880
1017,588
611,299
730,129
947,114
266,478
55,535
753,575
977,254
147,299
1305,31
792,754
594,166
1275,224
380,0
700,467
246,192
146,621
768,316
12,399
284,847
147,595
1287,641
817,658
99,467
535,147
1257,126
947,45
102,857
691,65
999,767
518,562
351,529
642,75
1275,347
494,287
872,628
1203,221
1295,299
914,2
386,0
1133,253
1163,299
1032,416
1195,112
661,845
658,770
701,390
997,810
957,103
323,652
574,130
750,220
60,142
689,712
656,360
749,236
216,628
753,149
681,680
371,75
18,257
333,456
209,480
333,886
320,754
1032,780
534,651
32,366
1232,621
775,63
524,679
758,409
872,446
358,145
1285,324
1066,431
1066,463
813,255
924,311
306,362
683,24
1125,205
812,68
681,886
274,796
542,126
294,334
1265,56
907,880
776,691
560,462
557,121
691,457
309,887
410,687
679,641
311,518
890,511
152,498
807,567
1226,714
156,68
823,617
438,154
1053,829
478,831
181,327
704,752
408,18
828,169
388,30
167,725
547,337
112,894
104,96
840,502
798,502
930,98
498,168
947,849
893,432
959,365
375,726
261,207
651,725
1295,760
182,705
336,738
1261,54
1223,850
907,656
1031,616
246,18
544,628
736,683
1310,360
212,467
1139,633
1141,290
186,56
1226,180
62,859
191,742
497,290
991,166
930,390
169,277
570,714
194,224
1285,772
60,80
607,40
1104,383
93,572
305,381
1044,142
463,742
1285,456
113,705
855,885
709,544
1198,0
1212,609
358,672
905,123
303,267
734,852
1004,738
1146,260
585,65
855,894
115,782
251,609
934,150
962,431
1088,96
448,511
1021,567
537,638
977,310
455,157
62,203
1258,628
184,816
1009,169
633,513
537,99
838,145
1278,528
442,19
258,444
800,413
962,532
542,768
949,217
999,376
485,290
469,205
823,242
1235,277
725,65
825,277
735,306
803,268
79,306
1001,887
223,882
823,277
27,198
381,782
913,322
1235,725
376,416
249,795
643,623
1125,653
649,79
1237,521
840,392
999,319
62,427
1032,478
420,732
962,362
873,392
341,84
330,770
217,392
847,742
674,637
738,431
852,112
387,625
518,332
589,492
480,322
816,287
1307,617
841,689
668,399
840,78
1036,0
328,466
319,728
1005,381
446,355
328,876
185,653
1115,665
278,780
217,56
664,626
373,627
848,854
30,525
244,431
817,236
498,826
1043,45
162,241
229,297

fold along x=655
fold along y=447
fold along x=327
fold along y=223
fold along x=163
fold along y=111
fold along x=81
fold along y=55
fold along x=40
fold along y=27
fold along y=13
fold along y=6"""

dots = data_str.split('\n\n')[0].split('\n')
instructions = data_str.split('\n\n')[1].split('\n')

dots = list(map(lambda x: list(map(int, x.split(','))), dots))
instructions = list(
    map(lambda x: (x.split('=')[0][-1], int(x.split('=')[1])), instructions))

# print(dots)
# print(instructions)


for inst in instructions:
    if inst[0] == 'y':  # horizontal fold (up)
        for i in range(len(dots)):
            if dots[i][1] > inst[1]:
                dots[i][1] -= 2 * (dots[i][1] - inst[1])
    else:  # vertical fold (left)
        for i in range(len(dots)):
            if dots[i][0] > inst[1]:
                dots[i][0] -= 2 * (dots[i][0] - inst[1])
    # just the first is needed
    break


print(len(dots))
dots = list(set(map(tuple, dots)))
print(len(dots))
