# --- Part Two - --

# Through a little deduction, you should now be able to determine the remaining digits. Consider again the first example above:

# acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
# cdfeb fcadb cdfeb cdbaf
# After some careful analysis, the mapping between signal wires and segments only make sense in the following configuration:

#  dddd
# e    a
# e    a
# ffff
# g    b
# g    b
# cccc
# So, the unique signal patterns would correspond to the following digits:

# acedgfb: 8
# cdfbe: 5
# gcdfa: 2
# fbcad: 3
# dab: 7
# cefabd: 9
# cdfgeb: 6
# eafb: 4
# cagedb: 0
# ab: 1
# Then, the four digits of the output value can be decoded:

# cdfeb: 5
# fcadb: 3
# cdfeb: 5
# cdbaf: 3
# Therefore, the output value for this entry is 5353.

# Following this same process for each entry in the second, larger example above, the output value of each entry can be determined:

# fdgacbe cefdb cefbgd gcbe: 8394
# fcgedb cgb dgebacf gc: 9781
# cg cg fdcagb cbg: 1197
# efabcd cedba gadfec cb: 9361
# gecf egdcabf bgf bfgea: 4873
# gebdcfa ecba ca fadegcb: 8418
# cefg dcbef fcge gbcadfe: 4548
# ed bcgafe cdgba cbgef: 1625
# gbdfcae bgc cg cgb: 8717
# fgae cfgab fg bagce: 4315
# Adding all of the output values in this larger example produces 61229.

# For each entry, determine all of the wire/segment connections and decode the four-digit output values. What do you get if you add up all of the output values?

# Your puzzle answer was 1024649.

data_str = """dcagfb cbegda cabdf bagedf dcb cd dcbegaf aebcf bgdaf dgcf | dacebgf cbafd dbacf cdgf
dfgbc ceagbd egb cgdae eafcgd ecbgd eb aefbgcd bfaceg deba | baed ecabgd aebd cadfebg
acbd bagfec bedcf efbca dbfecag gadcfe dc dbcaef cde fbegd | deacfg bfagcde bgedf cde
dgec dfc fabgd bfcgae dfbcg baegdcf cd fbdace cdbefg becgf | cedg adfbecg badgf dgbfc
be bacedg dfacb fbge deagcf eba egacf gcebfa febac dbgacfe | befg be egbf aedgcbf
gafbd gbafce agced fdgae bcedgaf gdeacb efdc eaf fe cgeadf | cgdae fdec fcdeabg ef
cdaefgb cfadbe gc afgcbe afedcg faecb cabg egc bfgde bfgce | fbaedc gc abgc bfcea
egcafb adgecf caebf fdgbaec fbaed cebg facbgd cfe fgabc ec | fbgca deagfc ce ce
fcdab deb cdbfea fecb befcdga abecd gdcae efbgda fcdgba be | cfbad dacge gbdeaf bed
afdecg gfcadb ecadfbg facgd ed egda efbgc ecd cdegf bfadec | becgf fcbeg de cgefb
caedbg gcfdba bfdae ge beagfdc bedag gcabd egac gcefbd deg | fdbecg acge bgfacd abefd
befad ecgad adgefcb febg gb afdcbe dcbfag beafgd dbg eagdb | feabd aedfb gbd aebdcf
bcfde acbfed dcaf dbcae gacbfe fc dgbcae egbacfd fcb edbfg | gbdfe gbeafc bcf cebfga
cgbade adcbgfe fb efgbad gbcfde aecfd edgba afbg bfeda fbe | ecafbdg bfdea fgab edfbgc
cbfa gefbd bcfedga bagefc ab bga ecdgab gbaef gfacde fcage | acdegf gab cegadb gab
cafbeg dfaebc fdca dgcbe eacbf fde debgfca egabfd dfceb fd | df eabcf egbdc dfe
egcdfa ebdac fdaecbg ed fdacb agecb cbdgaf ecd efdb dbaecf | dcabf cbeag ebdf ecgba
gacdebf decfbg afbdg dcagbf cg agdbc agcf ebcad abgedf bcg | badec dcbga cgfa abcde
fgcad dfebga gdbef ec ebcd fbcgea ceg fgedc defgbac dcfgbe | ce gbdfe fbcgdea fcgbdea
cgbef afcbeg dgbec ebafcd cgbfeda adbge cd ecd gfcd fgbdec | edc gcebd fegacdb cd
bafdcg egfdc bcfge dgeabf cd cgbadfe agefd facegd eadc cgd | cgd dcg defbagc cfdbega
degfcb fdg edabgf cadge fg gfcb cdeafbg bcefd dgfec beadcf | dgcbfe cabfed fecbd cgdfbea
dbfaec bfcgae cdb cdaf defbcg abgde ceabf adceb decgafb cd | fcda bdc dc afgecb
cead efcdg adbcfeg fagdc def acegfd fcbagd cbgef ed fgdaeb | ecad fcaedg cgadf cfaged
egdbfa cbgefad dgacbf cb dfbgec gefbd cbf gefbc gcefa bdec | ebcd ecgdbf fgebc cfb
fagbe acbged bf eagcf dgaebcf gbf febadg deabg cdagbf fbed | fbcdag fgeca abgdcf cdbage
baefc febgadc fcag dbgec feg gacfbe gfecb afedgb fg edafcb | aegdfb feabc agbfed fabged
gfedcb eadgcf fed befdagc gaedf bgdea bdafcg fe acdfg feca | bcfdga cagfd afce efca
cbfed degfba fgdbac acdg fdcab ac acgebf agdfb bac gcbfdea | cbdegfa caefgbd ca gdbcafe
dbcgea dc dcg cadf abgcdf efgdb afgbc gefacdb fbcdg agebfc | dabgce gbceda dgc dc
ef dgfbc cefbg gabecfd agceb feg gbcadf dfbe fcdgeb dfagec | gfe cafegd dfeb fe
ebgdc egdacb dfbcega dfe dbfge bagef ecdbfg bdcf df fgecad | dbgce fd cfdb edf
bdage ab fbgaecd aefb bfadeg gefad cdgeb agcbdf gba fgcdae | cgdbfa befa dgefca acebfdg
aecb dbgecf eagfcb ecdbfag fdegba geacf gebfa cge cgdfa ec | acgfe gec cabe fgcda
acd cafg abedcgf cfebda efcbgd cegfd ca cdgafe agcde bgeda | egcfbda gcfdae gefacd dca
dcefga gdefa db fbaed abefdgc faebc fbgaed bdf gbed dcafgb | bfd db db gbdfae
befac abcd bdcfe ecd cd baegfc edfgca adfbce aebcdgf edbfg | cgedfab feabc cd dabc
eacdgf cebgdfa acgebd ecdba dbcfea fdgcb dag ga acgdb abge | ag ga adg ag
gca fagbe dfageb bagce ecgfab gfce gc dcfeabg fadbcg bedac | afbdceg gadfeb cga acbdfg
gbad febac dea gadbecf fdegbc edfacg da bdecg cdbea bdcgae | dbag cdeabg gcdaeb adbg
cdgfeb cbeaf adef cfa fa aecbg fedcb acdebf bdcfga cdefbga | fdea eadf efad caefb
aedgb df cgbaed gdef bfcgda bfd efbac dgfeab debfa aegcbdf | acbfe egfd afebgd gefd
badcfe fcbg bf dfecbga faged gfadbc dfb dfgab abdcg dbceag | bf fb dafbgec fbd
bdae decgbf afcbge bdecfa afcbd edgfcab ebcaf bd dfacg dbc | ebda db cbd cbd
acfgde ae eac agbdec gadcb ebgcafd baed gfcbe cgabe bfgadc | cea fecdga ebgdac eac
cbedga cfb adcbe bfae afcdgeb cgdaf fcedab bf fcbdeg dacbf | dfceba fcdga fbc fcebda
cdafg dfgeac fc dbgcef gadbc gacfebd dcf edfgab gefad caef | fadgc facged adgefb ecgbdaf
gdcbf eabdgf gaefd fbdge afcedg bacgfe fagbedc dbea efb be | ecagfb dgaef be eb
gbecad faced abgedf cd cbfd abedcf dacgbfe cde deabf agcfe | cd edbfac cgabed daebfg
cgebaf dacbge gafed eadbgf dgcef ad fadgbce dea bgafe dafb | fegad dgeaf ad gdfae
ecfadb eagdbc afbcdge abfcgd cf bgedf agfc cfd gbfdc bcagd | fcga fcd befgd gdbfc
afcbg gfeb gfcad ecgdba cbf cafbeg bacge fb ebdfac baedcfg | fbc gabce egabc fcadg
agd gadbc da agcfed gcebad dgfcb ecabg ebcagdf eadb becfag | fgbeadc bcega geacb gceab
gaebf gfdbce aefbgd fde aefcbg ed dcaebfg agde aefdb acfdb | efd agcdebf cbgefd ebgfa
cdgfe ecbafd ebgca ad gedbcfa cad adbg cedabg egcad fcebga | gbedac adgb ecdgab daceg
efcb fcdag begdfa egcdf edf fe gfcdeb cdbega geacdbf becdg | dagfc gbadef decbg fe
bgaed agfdc fdgea fae fabcgd gcfe fe decfga afgedcb facbed | gdafe ef eacfbgd fe
ebgfca def fd ecgfbd efbda eabdg edgcfba afdc ecbfa bfdaec | abcefd fd dafc debcagf
eacbg efbdgca fdbga gde abecdg beagd cdeafg de cebd ecagbf | gaedbc gdcbafe gaecdf bgdeac
bafce cedb edacgfb fbgac fbe aedfgb be cbedfa cdefa eacfgd | be acbfg gfbac edbc
dgba cbgfdea gbe fbaegc cbedga fegdc decafb cdbge bg daecb | cadfbe bgad beadgc bgedc
abd fdbcae aedfb daec ad bacfdg dgefcb agbef cbefd cbagdef | ad dab acde da
cdagbe dagbe fageb gdfbeca gdcba edac dgcfba egd egcfbd ed | fbgcad cafdbg de facbged
dbecf egadfb cfagde da fcebga cagef dae fadec dagc dfbcaeg | da dgac bfedc ad
efa adcf fgcde acgeb bgedfa efcga degfbc aegcfdb fa efdgac | cfgea fedgc ecagf efdacgb
fecga eafgd acgefb cf gceab fbcdag gecdafb cabedg fbec gcf | agcbfd gfc cgf ceagf
egcaf ecdafgb dfgea cbafe dfebgc bgac cbagef cg gfc ecdfab | edbfgca edafg cgf gfc
afdcgb decgb fbdegca fbgae fdbaeg dgfbe gbafce df fdb efda | gcbed dafe gbfdca aefgb
gaebd baefg cfadbg fb gaedfbc gadbef cfgea bfde baf gacdeb | abf cdbagf aecbdg fagec
cfdb dcg geadbf febdg edfabcg dc bgecd eabcg fcgead cgbfde | dgc dc bdfge dcfebga
dfbcaeg fdgcae gfdab bagc gabcdf fecbgd cadgf gb fbade bdg | bdg gacb bgac ecfbdg
eg befdg fceabd badfg dbecf egbcfa bedgcf gbacefd egb dceg | beg faecdb gfdbe dbacef
cdagef edfca dgfa fca fa agedc gfcabe bdfec bdcage fbcaedg | caedbfg ecafgbd fdcega decgfba
gfc cfeb cafgdb fc aedgbfc cgdfe degfb aedcg fbegcd dfgbea | acdge fedbg agedc ebgdfa
aegdcb gde dgbfca eg bdecfga fedbgc gcbfd fgecd bfge acefd | ge debcag bgedcf bfeg
gbd eagbc cbefagd gcde aefdb gcebad cgdafb dg aegbd facbge | cdeg cged ebdaf gbd
eagbcf cba aebd cdfga cdegbf afcbedg dcbfe ab adfcb acfbde | bcfeag fgdcbe daeb bdae
gc dabcfg dcfageb acg ecabf bgadef facge gfeda ecdg fagced | gcde efdga agc fgaed
bcgdea bacfdg aeg cdabg dfeagc befgd aebc dabfgec ae degba | ebfdg egadcf aceb gbdac
ceadf defagb fg agbed dfgaebc bfaecg eacdbg afgde efg fgbd | gf egbda fabegc badfgec
begcfd fdecb fbdga ca bace cfadb egfdca cda cfaedb efabdcg | gacefd fdgba fcdebg bcae
bcga agbefc fageb ca afc debfc dafegc fdgeab abfdgec fecab | gdafceb befac fdecga ac
fce adcfg cgdefa bgcfda fdae acegf fe cgbae adgfceb cdegfb | gedacbf fdebgac deaf afde
agcf dbegf bga cfeba gfdebca fgaeb cbdefa ag egadcb cbfega | fbdge gcabef cagf ebdafc
aedg ge decbga bafce ecg gcdab degfcab gcfdeb ceabg abdcfg | cbfea ecg bafec eacfb
gfdceb cgb bcdfe gebdc egcabdf egcf bafcgd bgdae edabfc cg | gbc bgdfca cgfe gc
bg bedcf dgfaec fdbgc gcbafed fbg bfaged gadcf cgab acbgfd | agfbde bg bgcdf fgb
gafcb cdfgab abfcge cbgde bdcfae ea faeg dbgcafe ace agebc | decbaf fbacde bcagf ace
bgfdac cgabdef efdca fagce febagc cd fcd dfeacg gdce aefdb | dc dgcfba ecdagf gfacbe
bde eb egfbdc ebgfd fbdgc daebgc cagfbd gfbceda efbc aedfg | cgebadf becf cfbdg cfbedg
ebcdfag fadec daceg gbecaf eabcfd fd fad aebfc fcgbad efbd | dcefbag fceagb fdecabg adf
efgd fd bfagcd acefb febad gdbaef edcgba badge dbgaefc daf | eafbd gbcdfa gdfe edgba
cefdag bfagec bedcg da acd bcfead bdgcefa gcaef acdeg agfd | geacbf faegdc fcdgbea eacgfdb
bcedgaf bgfae cab bdcg cb gcade egdacf abdegc agbec becfad | gfcadbe adgbec gebfcda cfaedg
gf edbcafg cfge gdafe gdf cdabfg eabdg dgefac acfed fbceda | geabd cbgdfa ceadbf agbfdc
fcbgde fbged cgfdb gdebcaf cfge ebg deabf baedcg ge dfabcg | beadf ge dcbgef cfeg
eafdcb gcdfbea eag dgcba gacdef agebd dbefag deabf ge fgbe | fcbdea bgcdeaf dbgac fagcebd
dc fgebad ecd aedcfb cdga ebcagd dfbaecg adebg debcg gbecf | bcegd gcdeb dacg eagdb
cad feabgd fadegc bgcdfea gbcfa bgcda bgaced decb dc ebadg | acd adfbge cd dc
af cfbge fadecg afbgc dbaf gbdfac bdgeacf dbcga cfa gebcda | fa dabcg dafb gbceda
fbgdea gcefa agfdb fecdba egbd fdgea de bagfcd gdcbafe dae | ed eda abdefc fadgeb
fdbcga fdbae bfa adfcge gcbfaed bfeg fdage eacbd fgbaed fb | bdcfage fbeg eadfbg gbef
cfdbga fdeab gefbad bcaefd egad fdaegbc ag gcefb afg agefb | abfgecd egda bgcfe adefb
fbeda ae dae bdgfe gcadbf beac fedabc badfc ceagfd dgcbafe | bfcad gfabdc aed fbdac
fgcaed bagcef bdfc edabg efacd ebcadf fb afb abefd agfbdce | cfdb fb abf abf
fed bfaedc fceab edabg bfdea ecdagf begfdac df fceabg bfcd | cebgfa df eagbd gafced
egfdc fcg cadg dfebc dfebcga cg dgabef ebgacf dfgea gcfade | fbedc gcfed dcag fgc
fcbde fbeacgd afedb da fcegda dabg ebgaf dgfaeb dea febagc | da bfdce egfacb ad
fcdeg cdbe fgcab dceagf febcg eb cebdgf geb edfgab dbgcaef | abgcfed bdec gdcfe cfgdbe
degfb ebdacf edbfcg agdfce dbf bd bgcd dgfce gfdcbae abegf | gebfcd fgeab gecdfb gdafcbe
egacd abedg decfgba dagceb ca cabe cag dagfcb badfeg cegfd | gca caeb cfgde ca
gfa ebafd fg gafdbec gaecfb gbdfa edfg gaedbf bdgca ecadfb | fbagec fecabd caegbf fg
dfgce cgfbe cefad bfdace fcaged cdg badceg fdceagb dg afdg | gcfbe dg gadf edfabc
dfgbc cbfeda bdgcfe fdb efbg bf gcedb fdaecgb gdecba fgadc | bgdafec bgcde fbgdc agdfc
fgecd dbegfac adbfgc fa bdcage acegb eacgf efab fca fceabg | gafce af gdcbea af
fgdacb fbdec adb gfda acbdegf edbcga cfgba da cbgefa cadfb | gbcfea bgcfda da dab
fbcdg abfcdg cfegbda ed dbcaef gedcfb aecgb dcegb ced egdf | cabfgde egcdfb fbcade dgcfab
dbec fcd cegbdf dc fdacgeb abdfg baefcg gbfcd dafgce gfbec | dcf cfd fcd agefcb
fcdeg gdefb cfgadbe fb bef bfcg edfacb fecdgb gcdfea gadeb | efb feb fbcg dabge
acebd dacfg bdf agdcbfe gcafdb fcegad afgdbe bcfg bfcad bf | dbf bdaec bf febacdg
acdegf bdfa afdebg febcga gaefb fd fed dgefcab bedgf edbcg | efgdab aecdgf cedgbfa def
ecdfabg dfge cedfba fcadgb fgabd ebgfad fea egafb bcgae ef | efgdbca eagcfdb defg bfgda
efacgd bfda bcadg gab gbceadf ab cfbdga gadfc bcgfae dgbec | fgecad dbaf ba abfgdec
bcgaed fbgc edcfg dgbecf fgdae dgc cg abdcef gfcbdea ebfdc | cg cdfeb cg edbcfag
cagb fgcdeab gdebaf fcegd ecabfd acfeg ag afceb fga acgbfe | ga agf bcadfe fag
fbgde efbagc ecgd efdgabc ebdcf fdcgbe abfdg fcbead egb eg | dcge cedbgf fbgad acegfbd
febadcg cfdbge bgcef gecd ed gbdaf bfdeg deb gfabce fcbead | efdbca ed fgcbe cebdaf
eagbc cgde cafbg abdge fbdgae gacedb aec ec acefbgd abcefd | bgacf cgbfeda agbed dgec
efgbdca gfdb cadbfg dag dg degfac bcdaef gdacb bfcad agebc | gda dfagecb gcfdea dag
gefab bfde acbfgde gfdac ebfgad bfgeca aegbcd eda gfade de | gafbce gcfeabd aefgd fgbcae
gaecfb cdebgf fec ebcafgd bgecad cf debgc bfead ebfdc cgdf | cdbfe gfcd egbdca edfcgb
cfdeba bfgd gcedfa gbdae dgefcab cbega dg febadg agd bdaef | dfgb adg gecdfa fcebda
dea bgdeaf da befdc gdcfaeb ecdbga abecd gacd bgafce eagcb | aedcb facebg abfecdg bagefdc
gcebf dfbae dbgcfe egdc dgebcaf febcd bcd cagfeb fdgcba dc | bcd efcbd daefb agbdfce
fa acged bdfec afeg gcfadb bfdagce ecgafd caf acdfe dgaecb | debfc af afge adegc
bacfdg fadbe egbdfc ba bage acefd baf bgdafe fgebd gdacbef | dbfage bgedf aedgbf acedf
acgb dgcbfa dba edgbfac acdefb begdf bdgfa dceafg dcagf ba | abd dfbge abd gacb
gcdea gc agfdce egdfa egcf fbeagd aecbd dgc fcebgad cagfbd | acebd egcf defag gfadceb
fab begadc ebcad dfbea gfcdab acdfgbe fa fedgb adcfbe feca | cbgead bfcdega fba gbadce
gb efcagd fgcad cfbg bedca abg fcgbda dgbcafe acbgd bdafeg | abcdgf bcgf cgfb bg
fbcagd fcdae beacgf gf agf baecdg fecag bgfe cagdbef gbcea | edbagfc ebfg gebf bgcae
cdgfea bdeacf aefcgb dbca beafc bcgdefa ebdgf ad defab ade | dbcgefa da ade bacd
bcgaedf badecg cfegb afceg dcbf egdbfa gcedb gbf bf fgcdbe | bagfed bdagfe fdbc gefcbd
cegba bgeda cgefa cbde fegadb dagcbe acfgdb cb cbg ecbdafg | cdegabf ebgda bfgade ebgac
afbdg gdabec gcafdeb fc cfdeba bcf fgebdc cefg gcdbf ecdbg | fecg cdbefg gdbfc gfce
bgdfa egbf bcdgfa fe dfcbea aef feadgb edcag dagef afdcegb | fea cbedfa acdfgeb dfgeab
eacgfd caegdb gbfecad bdfce edcfa dgaf fea fa egdac aegbcf | af fa edbcf gdeacbf
baefg fedabcg gbfac gacbfe cb dfageb dcfag cgb bcea cfdegb | cabe eabc gfacd fgdcabe
fgdeac cgd ebdfga agefdbc bfcegd cdgbf cd cfabg dceb befgd | cgd cfgab fdgeb efbgd
aedgbf cgebda dbcagf gbf gbefd gdecf gdfabec fb eabdg feab | bf fbea cefdagb adgbe
fcgea fbgedca gbecfa efbadc efgd fcagd df cdfaeg fdc acgdb | efgac fcd dfceba fegac
egdfac cab cadbf gbdaf bc eadcf dcebagf dbecag fbdace cebf | cedabfg cgbaed bc efacbdg
be abcged afcgb ebg bgefc gebfcd cgefd efdb gdeafc daecgbf | gabedc dbeafcg debf febcadg
dgebafc cefdbg db edfb fcgdb fdceg dgecfa bdc abcfg baecdg | fegdca baecdfg bcd bd
fbgead cfab cdbgefa agcbdf aedgc fgcbd fa fagcd ecgbdf gaf | adgcebf gaf gaf af
debagfc eb fcedba ebc cgdab efdgac befa eadcf edabc gdbefc | bcdagfe fbae ecadb eabf
gaedcb bfcgd bcegf bagdcf fd aedfgbc bfd cbdga cabefd adfg | ecfbda cbfgda cfbadg fbgcd
agdcfbe bcfega cbeag ge dacfeb facged bcaef gbdac gce efbg | feabc acfbdeg aefbc gce
febcdg egdacb bedcf debfacg abdfce ba bae fbeda gdfea acfb | eadbf ab ab eab
cdfbg dfbcae dcabe egba ag cgdfea abdcg aefcdgb cga decagb | defabc edacb cbfgd ag
fgac dfgec eag cfbdge bacde ag abdgef bagfdec dcgfae agcde | aecbdgf gcaf efcgd befcdg
de gecad fbcgae gcafd dge edfbcg eadb eacfdbg bcedag acbge | gabce fcegba bgfcead ed
gbacd dbefcg fdcgae edgca afdeg aecf egc gfaedb dcfagbe ce | feac acegd caefbdg aecf
egacbd aedbc dagbc fcdagb ec bacefdg ace dcge ecfagb febda | afecgb fgdbac bcade cea
ceadgb cgaedf bfdge cfg cf fdgbc dgabc cfdabg gdcebaf bacf | gebdf fdeacg dafbgec cgf
bgfda gdcfae afce bfecgd ac acdfg afebdgc dgefc bdcaeg cag | gdbfec gcedf gfdaceb cegdab
fa fgeabdc bfga faedb fad ebgdfa daceb fgbcde fgdeb dafcge | efdab gcdfeb fa gabf
ebfc fb aedbgc ebagdf fab cgfeab gcdfeba acbeg abcfg gfdca | gbfac bgfeac bfa bcfga
fagdec gadebc cg cabfegd agdfb cbge dbeac cdfabe adcbg cgd | acgbed dgfeabc fdeacb dgbacef
fgbdc gbedaf bdgaf dfc cbfadg gfceb degcfa dc bacd gbafcde | dcf egfdca cd cdegaf
fcb aebf bf cafgb efgabc cgadbe egbadfc egbca gcfebd acdgf | beacgfd gebfcd dfbgec bf
bfgcd deb bfacdg ed fcebgd baefg dbaecfg efdc dabecg fdgbe | de gebaf dbcgfe bde
facgdbe gda ebdfag gabdf da cagfeb gebaf deaf dbagec dcfbg | egfba dgebca aefd cgdbf
deabg eg gdfacb bdfea egd gebfdca ecdgab acge dacbg cefgbd | eg dgcbef agcdb fadbceg
cafbged feacb gbafd dc gfdc agdcbf fcabd deabgf cad acgedb | fdcba egbfdac cad dca
dfagce fbd eabdfc gdfeb ebgfa dfabcge fdgce cedfbg bcgd bd | cdbg agbcefd db db
fdagec fcg gf gbace efbgacd ecgbfd dfcea caefg dagf dbecaf | acdegfb fg gbdfaec fadg
edfag bgfced dga ag cfbgdea bfdcag bfged decfa afedbg bgae | abeg dgaef ag edfgb
dcaegb fageb cdga begda adbce gd dge gbecafd cabfed cgfebd | gdaebc ged egd dcga
cfb agebfc fegb fb adcef afbec ebgdca bafdgc gbcea fdbeagc | ebfg gabce fbc agfdceb
fdcega cbgfea febg agdebfc cfe afecb gbcae ef bcdfa acgedb | fbgdeac egbf acfgeb bfge
fbadge dfgebc fbgacde bg fdage afdcge cedba agbf bge dgbea | adgbe fgba gabf fceagd
cfbda bgedca fgdcba abcfe dc agbfd dbc fdgc facbged dagfeb | cd fbcda fdcab fdgc
cdefa fbgdce aebdg geacd ebcdgfa gc cdg baegdc febdag gbac | fdeca cg dgfbcae bcgfde
gcadeb ef ecbfga acbgfed fbdga bcade dafcbe dcfe deabf feb | fdegacb agdbf fgbedac fcdbaeg
adgc fac ca gdfcae cfdegb gcdef adcfbe dgefcba cagef abgef | gcadbfe feabdc cdefg bgdeafc
cfdeba dega fga gadfec cdfabeg aecbfg cbdfg ag faedc dfcag | ag ga acfed cebfad
efcgb feagdc gdfbea fadc ac cea gadbfce cfgae aedgbc geafd | fcgeabd cafeg cbadge ca
ebfda cfgbea cgdbef fbadgec dgcb cedfga dg fdbeg efcgb gfd | fgdeb bfgcde ebfdg dg
dcafg facgdb bgfa dbacf ba gdcefa ecfdb cfegdba bac acedbg | gcbdaf fbadc begdac dgcbaf
eagf agedbcf acdegb abg gcfba fcegb befgcd fbcda ag ebcfag | dbcegfa bfecg cafdebg faeg
fbgcd cba fgacdb gabcdef fbgdec gdabc gacf bafdec begad ac | eacbdf badge acb becfad
eacbd dabfec dbgfeac ag eagdb fcdbag bga agec dbfge bacged | agecdb bag gebcdaf cdbea
bfgcd bcdegfa fdabgc da efgca cfagd efbdca cda dgfecb bgda | acdebf badfcg abdg cafgd
efcbgd cadefg cfegd dcgba cafedgb cagfd gaf acef fa bedafg | agf fcedg gefdcb cfdge
afdceg dcaf fd dfg aefgc gdbecf cafbdeg dbage ecagfb deafg | adefgc bfdecga fdg cefag
bfdacg dagbf abdefcg acfg fbcda beafdc egdab cfdegb gfd gf | bcafgd efgcabd cfga dcaegfb
dbage dfagbe cadfeb gfeb bae fabgdc be bafdg cdgea bgdfaec | fedacbg eb dfcbea eab
gebfd cfedga begfa dagbfc fgeadcb ecdgbf bced edfgc dgb db | dbec gadcfb db aefbg
dbcg fbagc bcf bc cdaebgf aegbdf faecg fdcaeb cdafbg afbgd | adfbg fadebgc dgbafc febagd"""

lines = data_str.split('\n')
displays = map(lambda x: x.split('|')[1], lines)

sum = 0

lengths1478 = [2, 4, 3, 7]

for line in lines:
    numbers = line.split(' | ')[0].split(' ')

    one = list(filter(lambda number: len(number) == 2, numbers))[0]
    seven = list(filter(lambda number: len(number) == 3, numbers))[0]
    four = list(filter(lambda number: len(number) == 4, numbers))[0]
    eight = list(filter(lambda number: len(number) == 7, numbers))[0]

    one = set(one)
    seven = set(seven)
    four = set(four)
    eight = set(eight)
    zero = {}
    two = {}
    three = {}
    five = {}
    six = {}
    nine = {}

    top = seven.difference(one).pop()

    mid_top_left_candidates = four.difference(one)

    top_left = bot_left = top_right = bot_right = mid = -1

    for nr in filter(lambda x: len(x) == 6, numbers):
        diff = eight.difference(nr)
        if diff.issubset(mid_top_left_candidates):
            mid = diff.pop()
            mid_top_left_candidates.remove(mid)
            top_left = mid_top_left_candidates.pop()
            zero = set(nr)
        elif diff.issubset(one):
            top_right = diff.pop()
            bot_right = one.difference(top_right).pop()
            six = set(nr)
        else:
            bot_left = diff.pop()
            nine = set(nr)

    bot = eight.difference(top, mid, top_right, top_left,
                           bot_right, bot_left).pop()

    two = {top, top_right, mid, bot_left, bot}
    three = {top, mid, bot, bot_right, top_right}
    five = {top, mid, bot, top_left, bot_right}

    # print(zero, one, two, three, four, five, six, seven, eight, nine)
    # print(numbers)
    # print(top, mid, bot)
    # print(top_left, top_right, bot_left, bot_right)

    display_numbers = list(map(set, line.split(' | ')[1].split(' ')))
    # print(display_numbers)

    d = {}
    for idx, var in enumerate([zero, one, two, three, four, five, six, seven, eight, nine]):
        d[idx] = var

    digits = []
    for display_number in display_numbers:
        for key, value in d.items():
            if value == display_number:
                digits.append(key)

    nr = int(''.join(map(str, digits)))
    print(nr)
    sum += nr


print(sum)
