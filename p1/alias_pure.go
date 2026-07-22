//go:build !amd64 && !arm64

package p1

import (
	base "github.com/goccy/spidermonkeywasm2go/base"
	_ "unsafe"
)

//go:linkname Fn27 github.com/goccy/spidermonkeywasm2go/p7.Fn27
func Fn27(m *base.Module, l0 int32) int32

//go:linkname Fn28 github.com/goccy/spidermonkeywasm2go/p7.Fn28
func Fn28(m *base.Module, l0 int32) int64

//go:linkname Fn31 github.com/goccy/spidermonkeywasm2go/p7.Fn31
func Fn31(m *base.Module, l0 int32)

//go:linkname Fn32 github.com/goccy/spidermonkeywasm2go/p7.Fn32
func Fn32(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn33 github.com/goccy/spidermonkeywasm2go/p7.Fn33
func Fn33(m *base.Module, l0 int32) int64

//go:linkname Fn34 github.com/goccy/spidermonkeywasm2go/p7.Fn34
func Fn34(m *base.Module, l0 int32)

//go:linkname Fn111 github.com/goccy/spidermonkeywasm2go/p7.Fn111
func Fn111(m *base.Module, l0 int32)

//go:linkname Fn112 github.com/goccy/spidermonkeywasm2go/p7.Fn112
func Fn112(m *base.Module, l0 int32)

//go:linkname Fn130 github.com/goccy/spidermonkeywasm2go/p7.Fn130
func Fn130(m *base.Module, l0 int32)

//go:linkname Fn134 github.com/goccy/spidermonkeywasm2go/p7.Fn134
func Fn134(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn146 github.com/goccy/spidermonkeywasm2go/p7.Fn146
func Fn146(m *base.Module)

//go:linkname Fn148 github.com/goccy/spidermonkeywasm2go/p0.Fn148
func Fn148(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn168 github.com/goccy/spidermonkeywasm2go/p3.Fn168
func Fn168(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn214 github.com/goccy/spidermonkeywasm2go/p6.Fn214
func Fn214(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn217 github.com/goccy/spidermonkeywasm2go/p3.Fn217
func Fn217(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn221 github.com/goccy/spidermonkeywasm2go/p6.Fn221
func Fn221(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn312 github.com/goccy/spidermonkeywasm2go/p3.Fn312
func Fn312(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn350 github.com/goccy/spidermonkeywasm2go/p7.Fn350
func Fn350(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn353 github.com/goccy/spidermonkeywasm2go/p7.Fn353
func Fn353(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn356 github.com/goccy/spidermonkeywasm2go/p7.Fn356
func Fn356(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn359 github.com/goccy/spidermonkeywasm2go/p7.Fn359
func Fn359(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn362 github.com/goccy/spidermonkeywasm2go/p7.Fn362
func Fn362(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn365 github.com/goccy/spidermonkeywasm2go/p7.Fn365
func Fn365(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn368 github.com/goccy/spidermonkeywasm2go/p7.Fn368
func Fn368(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn371 github.com/goccy/spidermonkeywasm2go/p7.Fn371
func Fn371(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn391 github.com/goccy/spidermonkeywasm2go/p5.Fn391
func Fn391(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn662 github.com/goccy/spidermonkeywasm2go/p7.Fn662
func Fn662(m *base.Module, l0 int32) int32

//go:linkname Fn663 github.com/goccy/spidermonkeywasm2go/p7.Fn663
func Fn663(m *base.Module, l0 int32)

//go:linkname Fn677 github.com/goccy/spidermonkeywasm2go/p7.Fn677
func Fn677(m *base.Module, l0 int32) int32

//go:linkname Fn680 github.com/goccy/spidermonkeywasm2go/p7.Fn680
func Fn680(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn690 github.com/goccy/spidermonkeywasm2go/p6.Fn690
func Fn690(m *base.Module, l0 int32) int32

//go:linkname Fn693 github.com/goccy/spidermonkeywasm2go/p6.Fn693
func Fn693(m *base.Module, l0 int32)

//go:linkname Fn705 github.com/goccy/spidermonkeywasm2go/p7.Fn705
func Fn705(m *base.Module, l0 int32)

//go:linkname Fn733 github.com/goccy/spidermonkeywasm2go/p6.Fn733
func Fn733(m *base.Module, l0 int32) int32

//go:linkname Fn741 github.com/goccy/spidermonkeywasm2go/p5.Fn741
func Fn741(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn742 github.com/goccy/spidermonkeywasm2go/p7.Fn742
func Fn742(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn754 github.com/goccy/spidermonkeywasm2go/p6.Fn754
func Fn754(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn759 github.com/goccy/spidermonkeywasm2go/p7.Fn759
func Fn759(m *base.Module, l0 int32)

//go:linkname Fn760 github.com/goccy/spidermonkeywasm2go/p7.Fn760
func Fn760(m *base.Module, l0 int32)

//go:linkname Fn780 github.com/goccy/spidermonkeywasm2go/p7.Fn780
func Fn780(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn783 github.com/goccy/spidermonkeywasm2go/p7.Fn783
func Fn783(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn784 github.com/goccy/spidermonkeywasm2go/p7.Fn784
func Fn784(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn789 github.com/goccy/spidermonkeywasm2go/p7.Fn789
func Fn789(m *base.Module, l0 int32) int32

//go:linkname Fn793 github.com/goccy/spidermonkeywasm2go/p7.Fn793
func Fn793(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn794 github.com/goccy/spidermonkeywasm2go/p4.Fn794
func Fn794(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn795 github.com/goccy/spidermonkeywasm2go/p6.Fn795
func Fn795(m *base.Module, l0 int32) int32

//go:linkname Fn800 github.com/goccy/spidermonkeywasm2go/p6.Fn800
func Fn800(m *base.Module, l0 int32)

//go:linkname Fn805 github.com/goccy/spidermonkeywasm2go/p7.Fn805
func Fn805(m *base.Module, l0 int32) int32

//go:linkname Fn814 github.com/goccy/spidermonkeywasm2go/p6.Fn814
func Fn814(m *base.Module, l0 int32) int32

//go:linkname Fn816 github.com/goccy/spidermonkeywasm2go/p6.Fn816
func Fn816(m *base.Module, l0 int32) int32

//go:linkname Fn832 github.com/goccy/spidermonkeywasm2go/p2.Fn832
func Fn832(m *base.Module, l0 int32) int32

//go:linkname Fn833 github.com/goccy/spidermonkeywasm2go/p4.Fn833
func Fn833(m *base.Module, l0 int32)

//go:linkname Fn834 github.com/goccy/spidermonkeywasm2go/p7.Fn834
func Fn834(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn835 github.com/goccy/spidermonkeywasm2go/p5.Fn835
func Fn835(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn838 github.com/goccy/spidermonkeywasm2go/p6.Fn838
func Fn838(m *base.Module, l0 int32) int32

//go:linkname Fn840 github.com/goccy/spidermonkeywasm2go/p0.Fn840
func Fn840(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn842 github.com/goccy/spidermonkeywasm2go/p5.Fn842
func Fn842(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn845 github.com/goccy/spidermonkeywasm2go/p0.Fn845
func Fn845(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn850 github.com/goccy/spidermonkeywasm2go/p0.Fn850
func Fn850(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn851 github.com/goccy/spidermonkeywasm2go/p0.Fn851
func Fn851(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn855 github.com/goccy/spidermonkeywasm2go/p7.Fn855
func Fn855(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn858 github.com/goccy/spidermonkeywasm2go/p0.Fn858
func Fn858(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn859 github.com/goccy/spidermonkeywasm2go/p0.Fn859
func Fn859(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn860 github.com/goccy/spidermonkeywasm2go/p0.Fn860
func Fn860(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn865 github.com/goccy/spidermonkeywasm2go/p3.Fn865
func Fn865(m *base.Module, l0 int32) int32

//go:linkname Fn866 github.com/goccy/spidermonkeywasm2go/p5.Fn866
func Fn866(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn867 github.com/goccy/spidermonkeywasm2go/p4.Fn867
func Fn867(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn870 github.com/goccy/spidermonkeywasm2go/p0.Fn870
func Fn870(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn871 github.com/goccy/spidermonkeywasm2go/p7.Fn871
func Fn871(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn872 github.com/goccy/spidermonkeywasm2go/p6.Fn872
func Fn872(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn875 github.com/goccy/spidermonkeywasm2go/p6.Fn875
func Fn875(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn876 github.com/goccy/spidermonkeywasm2go/p6.Fn876
func Fn876(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn883 github.com/goccy/spidermonkeywasm2go/p0.Fn883
func Fn883(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn891 github.com/goccy/spidermonkeywasm2go/p0.Fn891
func Fn891(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn893 github.com/goccy/spidermonkeywasm2go/p6.Fn893
func Fn893(m *base.Module, l0 int32) int32

//go:linkname Fn900 github.com/goccy/spidermonkeywasm2go/p5.Fn900
func Fn900(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn924 github.com/goccy/spidermonkeywasm2go/p6.Fn924
func Fn924(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn940 github.com/goccy/spidermonkeywasm2go/p7.Fn940
func Fn940(m *base.Module, l0 int32) int32

//go:linkname Fn1010 github.com/goccy/spidermonkeywasm2go/p5.Fn1010
func Fn1010(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1036 github.com/goccy/spidermonkeywasm2go/p5.Fn1036
func Fn1036(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1094 github.com/goccy/spidermonkeywasm2go/p0.Fn1094
func Fn1094(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1095 github.com/goccy/spidermonkeywasm2go/p0.Fn1095
func Fn1095(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1096 github.com/goccy/spidermonkeywasm2go/p0.Fn1096
func Fn1096(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1110 github.com/goccy/spidermonkeywasm2go/p0.Fn1110
func Fn1110(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1134 github.com/goccy/spidermonkeywasm2go/p6.Fn1134
func Fn1134(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1144 github.com/goccy/spidermonkeywasm2go/p6.Fn1144
func Fn1144(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1146 github.com/goccy/spidermonkeywasm2go/p4.Fn1146
func Fn1146(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1147 github.com/goccy/spidermonkeywasm2go/p5.Fn1147
func Fn1147(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1148 github.com/goccy/spidermonkeywasm2go/p5.Fn1148
func Fn1148(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1149 github.com/goccy/spidermonkeywasm2go/p5.Fn1149
func Fn1149(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1150 github.com/goccy/spidermonkeywasm2go/p5.Fn1150
func Fn1150(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1152 github.com/goccy/spidermonkeywasm2go/p6.Fn1152
func Fn1152(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1153 github.com/goccy/spidermonkeywasm2go/p6.Fn1153
func Fn1153(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1154 github.com/goccy/spidermonkeywasm2go/p6.Fn1154
func Fn1154(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1155 github.com/goccy/spidermonkeywasm2go/p5.Fn1155
func Fn1155(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1170 github.com/goccy/spidermonkeywasm2go/p6.Fn1170
func Fn1170(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1172 github.com/goccy/spidermonkeywasm2go/p6.Fn1172
func Fn1172(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1173 github.com/goccy/spidermonkeywasm2go/p6.Fn1173
func Fn1173(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1185 github.com/goccy/spidermonkeywasm2go/p4.Fn1185
func Fn1185(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn1201 github.com/goccy/spidermonkeywasm2go/p6.Fn1201
func Fn1201(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn1289 github.com/goccy/spidermonkeywasm2go/p4.Fn1289
func Fn1289(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1306 github.com/goccy/spidermonkeywasm2go/p0.Fn1306
func Fn1306(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1307 github.com/goccy/spidermonkeywasm2go/p0.Fn1307
func Fn1307(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1308 github.com/goccy/spidermonkeywasm2go/p4.Fn1308
func Fn1308(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1309 github.com/goccy/spidermonkeywasm2go/p5.Fn1309
func Fn1309(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1310 github.com/goccy/spidermonkeywasm2go/p5.Fn1310
func Fn1310(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1311 github.com/goccy/spidermonkeywasm2go/p4.Fn1311
func Fn1311(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1312 github.com/goccy/spidermonkeywasm2go/p2.Fn1312
func Fn1312(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1316 github.com/goccy/spidermonkeywasm2go/p3.Fn1316
func Fn1316(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1322 github.com/goccy/spidermonkeywasm2go/p6.Fn1322
func Fn1322(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1323 github.com/goccy/spidermonkeywasm2go/p6.Fn1323
func Fn1323(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1366 github.com/goccy/spidermonkeywasm2go/p6.Fn1366
func Fn1366(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1367 github.com/goccy/spidermonkeywasm2go/p6.Fn1367
func Fn1367(m *base.Module, l0 int32) int32

//go:linkname Fn1396 github.com/goccy/spidermonkeywasm2go/p6.Fn1396
func Fn1396(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1397 github.com/goccy/spidermonkeywasm2go/p6.Fn1397
func Fn1397(m *base.Module, l0 int32) int32

//go:linkname Fn1416 github.com/goccy/spidermonkeywasm2go/p7.Fn1416
func Fn1416(m *base.Module, l0 int32) int32

//go:linkname Fn1422 github.com/goccy/spidermonkeywasm2go/p0.Fn1422
func Fn1422(m *base.Module, l0 int32) int32

//go:linkname Fn1432 github.com/goccy/spidermonkeywasm2go/p7.Fn1432
func Fn1432(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1444 github.com/goccy/spidermonkeywasm2go/p0.Fn1444
func Fn1444(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1451 github.com/goccy/spidermonkeywasm2go/p0.Fn1451
func Fn1451(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1468 github.com/goccy/spidermonkeywasm2go/p5.Fn1468
func Fn1468(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1469 github.com/goccy/spidermonkeywasm2go/p6.Fn1469
func Fn1469(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1471 github.com/goccy/spidermonkeywasm2go/p6.Fn1471
func Fn1471(m *base.Module, l0 int32)

//go:linkname Fn1472 github.com/goccy/spidermonkeywasm2go/p5.Fn1472
func Fn1472(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1473 github.com/goccy/spidermonkeywasm2go/p6.Fn1473
func Fn1473(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1475 github.com/goccy/spidermonkeywasm2go/p5.Fn1475
func Fn1475(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1480 github.com/goccy/spidermonkeywasm2go/p0.Fn1480
func Fn1480(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1481 github.com/goccy/spidermonkeywasm2go/p0.Fn1481
func Fn1481(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1495 github.com/goccy/spidermonkeywasm2go/p7.Fn1495
func Fn1495(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1500 github.com/goccy/spidermonkeywasm2go/p7.Fn1500
func Fn1500(m *base.Module, l0 int32)

//go:linkname Fn1501 github.com/goccy/spidermonkeywasm2go/p6.Fn1501
func Fn1501(m *base.Module, l0 int32) int32

//go:linkname Fn1504 github.com/goccy/spidermonkeywasm2go/p7.Fn1504
func Fn1504(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1505 github.com/goccy/spidermonkeywasm2go/p7.Fn1505
func Fn1505(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1517 github.com/goccy/spidermonkeywasm2go/p5.Fn1517
func Fn1517(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1518 github.com/goccy/spidermonkeywasm2go/p7.Fn1518
func Fn1518(m *base.Module, l0 int32) int32

//go:linkname Fn1519 github.com/goccy/spidermonkeywasm2go/p5.Fn1519
func Fn1519(m *base.Module, l0 float64, l1 float64, l2 float64) float64

//go:linkname Fn1521 github.com/goccy/spidermonkeywasm2go/p5.Fn1521
func Fn1521(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn1525 github.com/goccy/spidermonkeywasm2go/p6.Fn1525
func Fn1525(m *base.Module, l0 int32, l1 float64) float64

//go:linkname Fn1629 github.com/goccy/spidermonkeywasm2go/p0.Fn1629
func Fn1629(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1636 github.com/goccy/spidermonkeywasm2go/p7.Fn1636
func Fn1636(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1652 github.com/goccy/spidermonkeywasm2go/p7.Fn1652
func Fn1652(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1662 github.com/goccy/spidermonkeywasm2go/p6.Fn1662
func Fn1662(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn1664 github.com/goccy/spidermonkeywasm2go/p5.Fn1664
func Fn1664(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn1666 github.com/goccy/spidermonkeywasm2go/p7.Fn1666
func Fn1666(m *base.Module, l0 int32) float64

//go:linkname Fn1667 github.com/goccy/spidermonkeywasm2go/p7.Fn1667
func Fn1667(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1669 github.com/goccy/spidermonkeywasm2go/p0.Fn1669
func Fn1669(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1672 github.com/goccy/spidermonkeywasm2go/p0.Fn1672
func Fn1672(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1769 github.com/goccy/spidermonkeywasm2go/p7.Fn1769
func Fn1769(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1771 github.com/goccy/spidermonkeywasm2go/p7.Fn1771
func Fn1771(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1814 github.com/goccy/spidermonkeywasm2go/p7.Fn1814
func Fn1814(m *base.Module, l0 int32) int32

//go:linkname Fn1815 github.com/goccy/spidermonkeywasm2go/p6.Fn1815
func Fn1815(m *base.Module, l0 int32) int32

//go:linkname Fn1816 github.com/goccy/spidermonkeywasm2go/p6.Fn1816
func Fn1816(m *base.Module, l0 int32) int32

//go:linkname Fn1858 github.com/goccy/spidermonkeywasm2go/p0.Fn1858
func Fn1858(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1879 github.com/goccy/spidermonkeywasm2go/p3.Fn1879
func Fn1879(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1894 github.com/goccy/spidermonkeywasm2go/p7.Fn1894
func Fn1894(m *base.Module, l0 int32) int32

//go:linkname Fn1900 github.com/goccy/spidermonkeywasm2go/p7.Fn1900
func Fn1900(m *base.Module, l0 int32) int32

//go:linkname Fn1901 github.com/goccy/spidermonkeywasm2go/p7.Fn1901
func Fn1901(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1944 github.com/goccy/spidermonkeywasm2go/p0.Fn1944
func Fn1944(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1951 github.com/goccy/spidermonkeywasm2go/p0.Fn1951
func Fn1951(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1972 github.com/goccy/spidermonkeywasm2go/p0.Fn1972
func Fn1972(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1984 github.com/goccy/spidermonkeywasm2go/p0.Fn1984
func Fn1984(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1985 github.com/goccy/spidermonkeywasm2go/p0.Fn1985
func Fn1985(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1988 github.com/goccy/spidermonkeywasm2go/p0.Fn1988
func Fn1988(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2000 github.com/goccy/spidermonkeywasm2go/p0.Fn2000
func Fn2000(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2003 github.com/goccy/spidermonkeywasm2go/p6.Fn2003
func Fn2003(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn2004 github.com/goccy/spidermonkeywasm2go/p0.Fn2004
func Fn2004(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2005 github.com/goccy/spidermonkeywasm2go/p5.Fn2005
func Fn2005(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2006 github.com/goccy/spidermonkeywasm2go/p7.Fn2006
func Fn2006(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2051 github.com/goccy/spidermonkeywasm2go/p0.Fn2051
func Fn2051(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2069 github.com/goccy/spidermonkeywasm2go/p0.Fn2069
func Fn2069(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2082 github.com/goccy/spidermonkeywasm2go/p7.Fn2082
func Fn2082(m *base.Module, l0 int32) int32

//go:linkname Fn2084 github.com/goccy/spidermonkeywasm2go/p6.Fn2084
func Fn2084(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2089 github.com/goccy/spidermonkeywasm2go/p0.Fn2089
func Fn2089(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2093 github.com/goccy/spidermonkeywasm2go/p7.Fn2093
func Fn2093(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2116 github.com/goccy/spidermonkeywasm2go/p6.Fn2116
func Fn2116(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2137 github.com/goccy/spidermonkeywasm2go/p4.Fn2137
func Fn2137(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2149 github.com/goccy/spidermonkeywasm2go/p6.Fn2149
func Fn2149(m *base.Module, l0 int32)

//go:linkname Fn2208 github.com/goccy/spidermonkeywasm2go/p0.Fn2208
func Fn2208(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32

//go:linkname Fn2223 github.com/goccy/spidermonkeywasm2go/p5.Fn2223
func Fn2223(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2224 github.com/goccy/spidermonkeywasm2go/p4.Fn2224
func Fn2224(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2225 github.com/goccy/spidermonkeywasm2go/p5.Fn2225
func Fn2225(m *base.Module, l0 int32)

//go:linkname Fn2253 github.com/goccy/spidermonkeywasm2go/p6.Fn2253
func Fn2253(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2265 github.com/goccy/spidermonkeywasm2go/p6.Fn2265
func Fn2265(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2295 github.com/goccy/spidermonkeywasm2go/p6.Fn2295
func Fn2295(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2305 github.com/goccy/spidermonkeywasm2go/p6.Fn2305
func Fn2305(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2306 github.com/goccy/spidermonkeywasm2go/p5.Fn2306
func Fn2306(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2307 github.com/goccy/spidermonkeywasm2go/p7.Fn2307
func Fn2307(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2309 github.com/goccy/spidermonkeywasm2go/p0.Fn2309
func Fn2309(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2311 github.com/goccy/spidermonkeywasm2go/p0.Fn2311
func Fn2311(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn2314 github.com/goccy/spidermonkeywasm2go/p6.Fn2314
func Fn2314(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2324 github.com/goccy/spidermonkeywasm2go/p0.Fn2324
func Fn2324(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2326 github.com/goccy/spidermonkeywasm2go/p7.Fn2326
func Fn2326(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2329 github.com/goccy/spidermonkeywasm2go/p0.Fn2329
func Fn2329(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2330 github.com/goccy/spidermonkeywasm2go/p6.Fn2330
func Fn2330(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2335 github.com/goccy/spidermonkeywasm2go/p0.Fn2335
func Fn2335(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2353 github.com/goccy/spidermonkeywasm2go/p5.Fn2353
func Fn2353(m *base.Module, l0 int32)

//go:linkname Fn2354 github.com/goccy/spidermonkeywasm2go/p7.Fn2354
func Fn2354(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2405 github.com/goccy/spidermonkeywasm2go/p6.Fn2405
func Fn2405(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2406 github.com/goccy/spidermonkeywasm2go/p6.Fn2406
func Fn2406(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2410 github.com/goccy/spidermonkeywasm2go/p4.Fn2410
func Fn2410(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2458 github.com/goccy/spidermonkeywasm2go/p0.Fn2458
func Fn2458(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2476 github.com/goccy/spidermonkeywasm2go/p0.Fn2476
func Fn2476(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2479 github.com/goccy/spidermonkeywasm2go/p6.Fn2479
func Fn2479(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2485 github.com/goccy/spidermonkeywasm2go/p4.Fn2485
func Fn2485(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2487 github.com/goccy/spidermonkeywasm2go/p6.Fn2487
func Fn2487(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2495 github.com/goccy/spidermonkeywasm2go/p4.Fn2495
func Fn2495(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2508 github.com/goccy/spidermonkeywasm2go/p2.Fn2508
func Fn2508(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2514 github.com/goccy/spidermonkeywasm2go/p5.Fn2514
func Fn2514(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2534 github.com/goccy/spidermonkeywasm2go/p0.Fn2534
func Fn2534(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2538 github.com/goccy/spidermonkeywasm2go/p6.Fn2538
func Fn2538(m *base.Module, l0 int32)

//go:linkname Fn2541 github.com/goccy/spidermonkeywasm2go/p0.Fn2541
func Fn2541(m *base.Module, l0 int32) int32

//go:linkname Fn2547 github.com/goccy/spidermonkeywasm2go/p7.Fn2547
func Fn2547(m *base.Module, l0 int32)

//go:linkname Fn2550 github.com/goccy/spidermonkeywasm2go/p5.Fn2550
func Fn2550(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2552 github.com/goccy/spidermonkeywasm2go/p0.Fn2552
func Fn2552(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2554 github.com/goccy/spidermonkeywasm2go/p0.Fn2554
func Fn2554(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2557 github.com/goccy/spidermonkeywasm2go/p0.Fn2557
func Fn2557(m *base.Module, l0 int32)

//go:linkname Fn2562 github.com/goccy/spidermonkeywasm2go/p0.Fn2562
func Fn2562(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2563 github.com/goccy/spidermonkeywasm2go/p0.Fn2563
func Fn2563(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2570 github.com/goccy/spidermonkeywasm2go/p5.Fn2570
func Fn2570(m *base.Module, l0 int32)

//go:linkname Fn2571 github.com/goccy/spidermonkeywasm2go/p5.Fn2571
func Fn2571(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2577 github.com/goccy/spidermonkeywasm2go/p0.Fn2577
func Fn2577(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2578 github.com/goccy/spidermonkeywasm2go/p0.Fn2578
func Fn2578(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2581 github.com/goccy/spidermonkeywasm2go/p0.Fn2581
func Fn2581(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2583 github.com/goccy/spidermonkeywasm2go/p0.Fn2583
func Fn2583(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2584 github.com/goccy/spidermonkeywasm2go/p0.Fn2584
func Fn2584(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2586 github.com/goccy/spidermonkeywasm2go/p3.Fn2586
func Fn2586(m *base.Module, l0 int32)

//go:linkname Fn2588 github.com/goccy/spidermonkeywasm2go/p0.Fn2588
func Fn2588(m *base.Module, l0 int32)

//go:linkname Fn2589 github.com/goccy/spidermonkeywasm2go/p0.Fn2589
func Fn2589(m *base.Module, l0 int32)

//go:linkname Fn2594 github.com/goccy/spidermonkeywasm2go/p6.Fn2594
func Fn2594(m *base.Module, l0 int32) int32

//go:linkname Fn2600 github.com/goccy/spidermonkeywasm2go/p0.Fn2600
func Fn2600(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2603 github.com/goccy/spidermonkeywasm2go/p4.Fn2603
func Fn2603(m *base.Module, l0 int32)

//go:linkname Fn2610 github.com/goccy/spidermonkeywasm2go/p6.Fn2610
func Fn2610(m *base.Module, l0 int32) int32

//go:linkname Fn2611 github.com/goccy/spidermonkeywasm2go/p0.Fn2611
func Fn2611(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2620 github.com/goccy/spidermonkeywasm2go/p0.Fn2620
func Fn2620(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2622 github.com/goccy/spidermonkeywasm2go/p0.Fn2622
func Fn2622(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2642 github.com/goccy/spidermonkeywasm2go/p6.Fn2642
func Fn2642(m *base.Module, l0 int32)

//go:linkname Fn2643 github.com/goccy/spidermonkeywasm2go/p0.Fn2643
func Fn2643(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2644 github.com/goccy/spidermonkeywasm2go/p7.Fn2644
func Fn2644(m *base.Module, l0 int32)

//go:linkname Fn2645 github.com/goccy/spidermonkeywasm2go/p0.Fn2645
func Fn2645(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2646 github.com/goccy/spidermonkeywasm2go/p0.Fn2646
func Fn2646(m *base.Module, l0 int32) int32

//go:linkname Fn2647 github.com/goccy/spidermonkeywasm2go/p0.Fn2647
func Fn2647(m *base.Module, l0 int32) int32

//go:linkname Fn2650 github.com/goccy/spidermonkeywasm2go/p0.Fn2650
func Fn2650(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2651 github.com/goccy/spidermonkeywasm2go/p0.Fn2651
func Fn2651(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2652 github.com/goccy/spidermonkeywasm2go/p0.Fn2652
func Fn2652(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2655 github.com/goccy/spidermonkeywasm2go/p7.Fn2655
func Fn2655(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2656 github.com/goccy/spidermonkeywasm2go/p3.Fn2656
func Fn2656(m *base.Module, l0 int32) int32

//go:linkname Fn2657 github.com/goccy/spidermonkeywasm2go/p2.Fn2657
func Fn2657(m *base.Module, l0 int32) int32

//go:linkname Fn2658 github.com/goccy/spidermonkeywasm2go/p6.Fn2658
func Fn2658(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2659 github.com/goccy/spidermonkeywasm2go/p3.Fn2659
func Fn2659(m *base.Module, l0 int32) int32

//go:linkname Fn2660 github.com/goccy/spidermonkeywasm2go/p2.Fn2660
func Fn2660(m *base.Module, l0 int32) int32

//go:linkname Fn2720 github.com/goccy/spidermonkeywasm2go/p0.Fn2720
func Fn2720(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2727 github.com/goccy/spidermonkeywasm2go/p0.Fn2727
func Fn2727(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2732 github.com/goccy/spidermonkeywasm2go/p0.Fn2732
func Fn2732(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2733 github.com/goccy/spidermonkeywasm2go/p0.Fn2733
func Fn2733(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2742 github.com/goccy/spidermonkeywasm2go/p4.Fn2742
func Fn2742(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2743 github.com/goccy/spidermonkeywasm2go/p4.Fn2743
func Fn2743(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2752 github.com/goccy/spidermonkeywasm2go/p0.Fn2752
func Fn2752(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2754 github.com/goccy/spidermonkeywasm2go/p0.Fn2754
func Fn2754(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2755 github.com/goccy/spidermonkeywasm2go/p0.Fn2755
func Fn2755(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2764 github.com/goccy/spidermonkeywasm2go/p4.Fn2764
func Fn2764(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2765 github.com/goccy/spidermonkeywasm2go/p6.Fn2765
func Fn2765(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2766 github.com/goccy/spidermonkeywasm2go/p0.Fn2766
func Fn2766(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2774 github.com/goccy/spidermonkeywasm2go/p6.Fn2774
func Fn2774(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2775 github.com/goccy/spidermonkeywasm2go/p7.Fn2775
func Fn2775(m *base.Module, l0 int32)

//go:linkname Fn2780 github.com/goccy/spidermonkeywasm2go/p6.Fn2780
func Fn2780(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2844 github.com/goccy/spidermonkeywasm2go/p5.Fn2844
func Fn2844(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2862 github.com/goccy/spidermonkeywasm2go/p0.Fn2862
func Fn2862(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2866 github.com/goccy/spidermonkeywasm2go/p5.Fn2866
func Fn2866(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2867 github.com/goccy/spidermonkeywasm2go/p5.Fn2867
func Fn2867(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2881 github.com/goccy/spidermonkeywasm2go/p0.Fn2881
func Fn2881(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2887 github.com/goccy/spidermonkeywasm2go/p0.Fn2887
func Fn2887(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2914 github.com/goccy/spidermonkeywasm2go/p6.Fn2914
func Fn2914(m *base.Module, l0 int32) int32

//go:linkname Fn2940 github.com/goccy/spidermonkeywasm2go/p0.Fn2940
func Fn2940(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2942 github.com/goccy/spidermonkeywasm2go/p0.Fn2942
func Fn2942(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2943 github.com/goccy/spidermonkeywasm2go/p0.Fn2943
func Fn2943(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2945 github.com/goccy/spidermonkeywasm2go/p4.Fn2945
func Fn2945(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2946 github.com/goccy/spidermonkeywasm2go/p0.Fn2946
func Fn2946(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2958 github.com/goccy/spidermonkeywasm2go/p0.Fn2958
func Fn2958(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2972 github.com/goccy/spidermonkeywasm2go/p0.Fn2972
func Fn2972(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3004 github.com/goccy/spidermonkeywasm2go/p6.Fn3004
func Fn3004(m *base.Module, l0 int32)

//go:linkname Fn3006 github.com/goccy/spidermonkeywasm2go/p5.Fn3006
func Fn3006(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3010 github.com/goccy/spidermonkeywasm2go/p6.Fn3010
func Fn3010(m *base.Module, l0 int32) int32

//go:linkname Fn3045 github.com/goccy/spidermonkeywasm2go/p0.Fn3045
func Fn3045(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3046 github.com/goccy/spidermonkeywasm2go/p0.Fn3046
func Fn3046(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3047 github.com/goccy/spidermonkeywasm2go/p7.Fn3047
func Fn3047(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3051 github.com/goccy/spidermonkeywasm2go/p6.Fn3051
func Fn3051(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3160 github.com/goccy/spidermonkeywasm2go/p2.Fn3160
func Fn3160(m *base.Module, l0 int32) int32

//go:linkname Fn3162 github.com/goccy/spidermonkeywasm2go/p2.Fn3162
func Fn3162(m *base.Module, l0 int32)

//go:linkname Fn3163 github.com/goccy/spidermonkeywasm2go/p7.Fn3163
func Fn3163(m *base.Module, l0 int32)

//go:linkname Fn3164 github.com/goccy/spidermonkeywasm2go/p0.Fn3164
func Fn3164(m *base.Module, l0 int32) int32

//go:linkname Fn3167 github.com/goccy/spidermonkeywasm2go/p7.Fn3167
func Fn3167(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3170 github.com/goccy/spidermonkeywasm2go/p5.Fn3170
func Fn3170(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3177 github.com/goccy/spidermonkeywasm2go/p0.Fn3177
func Fn3177(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3190 github.com/goccy/spidermonkeywasm2go/p6.Fn3190
func Fn3190(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3191 github.com/goccy/spidermonkeywasm2go/p7.Fn3191
func Fn3191(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3192 github.com/goccy/spidermonkeywasm2go/p7.Fn3192
func Fn3192(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3194 github.com/goccy/spidermonkeywasm2go/p0.Fn3194
func Fn3194(m *base.Module, l0 int32) int32

//go:linkname Fn3205 github.com/goccy/spidermonkeywasm2go/p5.Fn3205
func Fn3205(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3216 github.com/goccy/spidermonkeywasm2go/p0.Fn3216
func Fn3216(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3218 github.com/goccy/spidermonkeywasm2go/p0.Fn3218
func Fn3218(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3223 github.com/goccy/spidermonkeywasm2go/p6.Fn3223
func Fn3223(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3224 github.com/goccy/spidermonkeywasm2go/p4.Fn3224
func Fn3224(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3239 github.com/goccy/spidermonkeywasm2go/p0.Fn3239
func Fn3239(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3250 github.com/goccy/spidermonkeywasm2go/p0.Fn3250
func Fn3250(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3267 github.com/goccy/spidermonkeywasm2go/p5.Fn3267
func Fn3267(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3471 github.com/goccy/spidermonkeywasm2go/p7.Fn3471
func Fn3471(m *base.Module, l0 float64) int32

//go:linkname Fn3473 github.com/goccy/spidermonkeywasm2go/p7.Fn3473
func Fn3473(m *base.Module, l0 float64) int32

//go:linkname Fn3491 github.com/goccy/spidermonkeywasm2go/p7.Fn3491
func Fn3491(m *base.Module, l0 int32)

//go:linkname Fn3503 github.com/goccy/spidermonkeywasm2go/p7.Fn3503
func Fn3503(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3507 github.com/goccy/spidermonkeywasm2go/p7.Fn3507
func Fn3507(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3509 github.com/goccy/spidermonkeywasm2go/p7.Fn3509
func Fn3509(m *base.Module, l0 int32) int32

//go:linkname Fn3534 github.com/goccy/spidermonkeywasm2go/p2.Fn3534
func Fn3534(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3535 github.com/goccy/spidermonkeywasm2go/p0.Fn3535
func Fn3535(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3537 github.com/goccy/spidermonkeywasm2go/p6.Fn3537
func Fn3537(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3540 github.com/goccy/spidermonkeywasm2go/p7.Fn3540
func Fn3540(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3541 github.com/goccy/spidermonkeywasm2go/p7.Fn3541
func Fn3541(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3543 github.com/goccy/spidermonkeywasm2go/p5.Fn3543
func Fn3543(m *base.Module, l0 int32) int32

//go:linkname Fn3544 github.com/goccy/spidermonkeywasm2go/p7.Fn3544
func Fn3544(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3548 github.com/goccy/spidermonkeywasm2go/p0.Fn3548
func Fn3548(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3551 github.com/goccy/spidermonkeywasm2go/p0.Fn3551
func Fn3551(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3552 github.com/goccy/spidermonkeywasm2go/p0.Fn3552
func Fn3552(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3562 github.com/goccy/spidermonkeywasm2go/p0.Fn3562
func Fn3562(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3574 github.com/goccy/spidermonkeywasm2go/p6.Fn3574
func Fn3574(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3576 github.com/goccy/spidermonkeywasm2go/p0.Fn3576
func Fn3576(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3580 github.com/goccy/spidermonkeywasm2go/p6.Fn3580
func Fn3580(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3584 github.com/goccy/spidermonkeywasm2go/p6.Fn3584
func Fn3584(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3586 github.com/goccy/spidermonkeywasm2go/p5.Fn3586
func Fn3586(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3587 github.com/goccy/spidermonkeywasm2go/p6.Fn3587
func Fn3587(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3588 github.com/goccy/spidermonkeywasm2go/p5.Fn3588
func Fn3588(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3591 github.com/goccy/spidermonkeywasm2go/p4.Fn3591
func Fn3591(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3594 github.com/goccy/spidermonkeywasm2go/p7.Fn3594
func Fn3594(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3595 github.com/goccy/spidermonkeywasm2go/p7.Fn3595
func Fn3595(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3596 github.com/goccy/spidermonkeywasm2go/p6.Fn3596
func Fn3596(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3597 github.com/goccy/spidermonkeywasm2go/p6.Fn3597
func Fn3597(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3598 github.com/goccy/spidermonkeywasm2go/p7.Fn3598
func Fn3598(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3599 github.com/goccy/spidermonkeywasm2go/p7.Fn3599
func Fn3599(m *base.Module, l0 int32, l1 float64) int32

//go:linkname Fn3604 github.com/goccy/spidermonkeywasm2go/p4.Fn3604
func Fn3604(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3605 github.com/goccy/spidermonkeywasm2go/p5.Fn3605
func Fn3605(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3606 github.com/goccy/spidermonkeywasm2go/p5.Fn3606
func Fn3606(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3607 github.com/goccy/spidermonkeywasm2go/p7.Fn3607
func Fn3607(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3608 github.com/goccy/spidermonkeywasm2go/p7.Fn3608
func Fn3608(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3609 github.com/goccy/spidermonkeywasm2go/p7.Fn3609
func Fn3609(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3610 github.com/goccy/spidermonkeywasm2go/p3.Fn3610
func Fn3610(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3614 github.com/goccy/spidermonkeywasm2go/p7.Fn3614
func Fn3614(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3615 github.com/goccy/spidermonkeywasm2go/p7.Fn3615
func Fn3615(m *base.Module, l0 int32) int32

//go:linkname Fn3654 github.com/goccy/spidermonkeywasm2go/p3.Fn3654
func Fn3654(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3663 github.com/goccy/spidermonkeywasm2go/p0.Fn3663
func Fn3663(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3664 github.com/goccy/spidermonkeywasm2go/p0.Fn3664
func Fn3664(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3696 github.com/goccy/spidermonkeywasm2go/p6.Fn3696
func Fn3696(m *base.Module, l0 int32)

//go:linkname Fn3721 github.com/goccy/spidermonkeywasm2go/p0.Fn3721
func Fn3721(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3734 github.com/goccy/spidermonkeywasm2go/p6.Fn3734
func Fn3734(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3748 github.com/goccy/spidermonkeywasm2go/p4.Fn3748
func Fn3748(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3766 github.com/goccy/spidermonkeywasm2go/p0.Fn3766
func Fn3766(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3767 github.com/goccy/spidermonkeywasm2go/p0.Fn3767
func Fn3767(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3768 github.com/goccy/spidermonkeywasm2go/p0.Fn3768
func Fn3768(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3769 github.com/goccy/spidermonkeywasm2go/p0.Fn3769
func Fn3769(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3770 github.com/goccy/spidermonkeywasm2go/p0.Fn3770
func Fn3770(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3771 github.com/goccy/spidermonkeywasm2go/p0.Fn3771
func Fn3771(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3772 github.com/goccy/spidermonkeywasm2go/p0.Fn3772
func Fn3772(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3773 github.com/goccy/spidermonkeywasm2go/p0.Fn3773
func Fn3773(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3774 github.com/goccy/spidermonkeywasm2go/p0.Fn3774
func Fn3774(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3775 github.com/goccy/spidermonkeywasm2go/p0.Fn3775
func Fn3775(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3776 github.com/goccy/spidermonkeywasm2go/p0.Fn3776
func Fn3776(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3777 github.com/goccy/spidermonkeywasm2go/p0.Fn3777
func Fn3777(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3778 github.com/goccy/spidermonkeywasm2go/p7.Fn3778
func Fn3778(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3779 github.com/goccy/spidermonkeywasm2go/p7.Fn3779
func Fn3779(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3780 github.com/goccy/spidermonkeywasm2go/p7.Fn3780
func Fn3780(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3781 github.com/goccy/spidermonkeywasm2go/p7.Fn3781
func Fn3781(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3782 github.com/goccy/spidermonkeywasm2go/p7.Fn3782
func Fn3782(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3783 github.com/goccy/spidermonkeywasm2go/p7.Fn3783
func Fn3783(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3784 github.com/goccy/spidermonkeywasm2go/p7.Fn3784
func Fn3784(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3785 github.com/goccy/spidermonkeywasm2go/p7.Fn3785
func Fn3785(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3786 github.com/goccy/spidermonkeywasm2go/p7.Fn3786
func Fn3786(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3787 github.com/goccy/spidermonkeywasm2go/p7.Fn3787
func Fn3787(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3788 github.com/goccy/spidermonkeywasm2go/p7.Fn3788
func Fn3788(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3789 github.com/goccy/spidermonkeywasm2go/p7.Fn3789
func Fn3789(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3790 github.com/goccy/spidermonkeywasm2go/p7.Fn3790
func Fn3790(m *base.Module, l0 int32) int32

//go:linkname Fn3791 github.com/goccy/spidermonkeywasm2go/p7.Fn3791
func Fn3791(m *base.Module, l0 int32) int32

//go:linkname Fn3792 github.com/goccy/spidermonkeywasm2go/p7.Fn3792
func Fn3792(m *base.Module, l0 int32) int32

//go:linkname Fn3793 github.com/goccy/spidermonkeywasm2go/p7.Fn3793
func Fn3793(m *base.Module, l0 int32) int32

//go:linkname Fn3794 github.com/goccy/spidermonkeywasm2go/p7.Fn3794
func Fn3794(m *base.Module, l0 int32) int32

//go:linkname Fn3795 github.com/goccy/spidermonkeywasm2go/p7.Fn3795
func Fn3795(m *base.Module, l0 int32) int32

//go:linkname Fn3796 github.com/goccy/spidermonkeywasm2go/p7.Fn3796
func Fn3796(m *base.Module, l0 int32) int32

//go:linkname Fn3797 github.com/goccy/spidermonkeywasm2go/p7.Fn3797
func Fn3797(m *base.Module, l0 int32) int32

//go:linkname Fn3798 github.com/goccy/spidermonkeywasm2go/p7.Fn3798
func Fn3798(m *base.Module, l0 int32) int32

//go:linkname Fn3799 github.com/goccy/spidermonkeywasm2go/p7.Fn3799
func Fn3799(m *base.Module, l0 int32) int32

//go:linkname Fn3800 github.com/goccy/spidermonkeywasm2go/p7.Fn3800
func Fn3800(m *base.Module, l0 int32) int32

//go:linkname Fn3801 github.com/goccy/spidermonkeywasm2go/p7.Fn3801
func Fn3801(m *base.Module, l0 int32) int32

//go:linkname Fn3802 github.com/goccy/spidermonkeywasm2go/p7.Fn3802
func Fn3802(m *base.Module, l0 int32) int32

//go:linkname Fn3803 github.com/goccy/spidermonkeywasm2go/p7.Fn3803
func Fn3803(m *base.Module, l0 int32) int32

//go:linkname Fn3804 github.com/goccy/spidermonkeywasm2go/p7.Fn3804
func Fn3804(m *base.Module, l0 int32) int32

//go:linkname Fn3805 github.com/goccy/spidermonkeywasm2go/p7.Fn3805
func Fn3805(m *base.Module, l0 int32) int32

//go:linkname Fn3806 github.com/goccy/spidermonkeywasm2go/p7.Fn3806
func Fn3806(m *base.Module, l0 int32) int32

//go:linkname Fn3807 github.com/goccy/spidermonkeywasm2go/p7.Fn3807
func Fn3807(m *base.Module, l0 int32) int32

//go:linkname Fn3808 github.com/goccy/spidermonkeywasm2go/p7.Fn3808
func Fn3808(m *base.Module, l0 int32) int32

//go:linkname Fn3809 github.com/goccy/spidermonkeywasm2go/p7.Fn3809
func Fn3809(m *base.Module, l0 int32) int32

//go:linkname Fn3810 github.com/goccy/spidermonkeywasm2go/p7.Fn3810
func Fn3810(m *base.Module, l0 int32) int32

//go:linkname Fn3811 github.com/goccy/spidermonkeywasm2go/p7.Fn3811
func Fn3811(m *base.Module, l0 int32) int32

//go:linkname Fn3812 github.com/goccy/spidermonkeywasm2go/p7.Fn3812
func Fn3812(m *base.Module, l0 int32) int32

//go:linkname Fn3813 github.com/goccy/spidermonkeywasm2go/p7.Fn3813
func Fn3813(m *base.Module, l0 int32) int32

//go:linkname Fn3855 github.com/goccy/spidermonkeywasm2go/p0.Fn3855
func Fn3855(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3856 github.com/goccy/spidermonkeywasm2go/p2.Fn3856
func Fn3856(m *base.Module, l0 int32) int32

//go:linkname Fn3877 github.com/goccy/spidermonkeywasm2go/p0.Fn3877
func Fn3877(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3878 github.com/goccy/spidermonkeywasm2go/p5.Fn3878
func Fn3878(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3879 github.com/goccy/spidermonkeywasm2go/p0.Fn3879
func Fn3879(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3880 github.com/goccy/spidermonkeywasm2go/p5.Fn3880
func Fn3880(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3881 github.com/goccy/spidermonkeywasm2go/p0.Fn3881
func Fn3881(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3882 github.com/goccy/spidermonkeywasm2go/p5.Fn3882
func Fn3882(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3883 github.com/goccy/spidermonkeywasm2go/p0.Fn3883
func Fn3883(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3884 github.com/goccy/spidermonkeywasm2go/p5.Fn3884
func Fn3884(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3885 github.com/goccy/spidermonkeywasm2go/p0.Fn3885
func Fn3885(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3886 github.com/goccy/spidermonkeywasm2go/p5.Fn3886
func Fn3886(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3887 github.com/goccy/spidermonkeywasm2go/p0.Fn3887
func Fn3887(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3888 github.com/goccy/spidermonkeywasm2go/p5.Fn3888
func Fn3888(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3889 github.com/goccy/spidermonkeywasm2go/p0.Fn3889
func Fn3889(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3890 github.com/goccy/spidermonkeywasm2go/p5.Fn3890
func Fn3890(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3891 github.com/goccy/spidermonkeywasm2go/p0.Fn3891
func Fn3891(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3892 github.com/goccy/spidermonkeywasm2go/p5.Fn3892
func Fn3892(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3893 github.com/goccy/spidermonkeywasm2go/p0.Fn3893
func Fn3893(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3894 github.com/goccy/spidermonkeywasm2go/p5.Fn3894
func Fn3894(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3895 github.com/goccy/spidermonkeywasm2go/p0.Fn3895
func Fn3895(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3896 github.com/goccy/spidermonkeywasm2go/p5.Fn3896
func Fn3896(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3897 github.com/goccy/spidermonkeywasm2go/p0.Fn3897
func Fn3897(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3898 github.com/goccy/spidermonkeywasm2go/p5.Fn3898
func Fn3898(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3899 github.com/goccy/spidermonkeywasm2go/p0.Fn3899
func Fn3899(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3900 github.com/goccy/spidermonkeywasm2go/p5.Fn3900
func Fn3900(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3908 github.com/goccy/spidermonkeywasm2go/p0.Fn3908
func Fn3908(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3911 github.com/goccy/spidermonkeywasm2go/p0.Fn3911
func Fn3911(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3913 github.com/goccy/spidermonkeywasm2go/p0.Fn3913
func Fn3913(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3915 github.com/goccy/spidermonkeywasm2go/p0.Fn3915
func Fn3915(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3920 github.com/goccy/spidermonkeywasm2go/p0.Fn3920
func Fn3920(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3922 github.com/goccy/spidermonkeywasm2go/p0.Fn3922
func Fn3922(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3924 github.com/goccy/spidermonkeywasm2go/p0.Fn3924
func Fn3924(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3929 github.com/goccy/spidermonkeywasm2go/p0.Fn3929
func Fn3929(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3949 github.com/goccy/spidermonkeywasm2go/p7.Fn3949
func Fn3949(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3952 github.com/goccy/spidermonkeywasm2go/p7.Fn3952
func Fn3952(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3955 github.com/goccy/spidermonkeywasm2go/p7.Fn3955
func Fn3955(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3958 github.com/goccy/spidermonkeywasm2go/p3.Fn3958
func Fn3958(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3959 github.com/goccy/spidermonkeywasm2go/p5.Fn3959
func Fn3959(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3960 github.com/goccy/spidermonkeywasm2go/p3.Fn3960
func Fn3960(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3961 github.com/goccy/spidermonkeywasm2go/p5.Fn3961
func Fn3961(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3962 github.com/goccy/spidermonkeywasm2go/p3.Fn3962
func Fn3962(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3963 github.com/goccy/spidermonkeywasm2go/p5.Fn3963
func Fn3963(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3964 github.com/goccy/spidermonkeywasm2go/p7.Fn3964
func Fn3964(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3967 github.com/goccy/spidermonkeywasm2go/p2.Fn3967
func Fn3967(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3968 github.com/goccy/spidermonkeywasm2go/p5.Fn3968
func Fn3968(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4038 github.com/goccy/spidermonkeywasm2go/p4.Fn4038
func Fn4038(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4039 github.com/goccy/spidermonkeywasm2go/p4.Fn4039
func Fn4039(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4117 github.com/goccy/spidermonkeywasm2go/p6.Fn4117
func Fn4117(m *base.Module, l0 int32) int32

//go:linkname Fn4118 github.com/goccy/spidermonkeywasm2go/p2.Fn4118
func Fn4118(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4119 github.com/goccy/spidermonkeywasm2go/p3.Fn4119
func Fn4119(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4120 github.com/goccy/spidermonkeywasm2go/p2.Fn4120
func Fn4120(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4121 github.com/goccy/spidermonkeywasm2go/p2.Fn4121
func Fn4121(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4141 github.com/goccy/spidermonkeywasm2go/p0.Fn4141
func Fn4141(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4198 github.com/goccy/spidermonkeywasm2go/p7.Fn4198
func Fn4198(m *base.Module, l0 int32)

//go:linkname Fn4199 github.com/goccy/spidermonkeywasm2go/p7.Fn4199
func Fn4199(m *base.Module, l0 int32)

//go:linkname Fn4229 github.com/goccy/spidermonkeywasm2go/p7.Fn4229
func Fn4229(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4240 github.com/goccy/spidermonkeywasm2go/p6.Fn4240
func Fn4240(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4242 github.com/goccy/spidermonkeywasm2go/p2.Fn4242
func Fn4242(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4244 github.com/goccy/spidermonkeywasm2go/p5.Fn4244
func Fn4244(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4254 github.com/goccy/spidermonkeywasm2go/p3.Fn4254
func Fn4254(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4255 github.com/goccy/spidermonkeywasm2go/p4.Fn4255
func Fn4255(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32) int32

//go:linkname Fn4256 github.com/goccy/spidermonkeywasm2go/p5.Fn4256
func Fn4256(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32) int32

//go:linkname Fn4258 github.com/goccy/spidermonkeywasm2go/p4.Fn4258
func Fn4258(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4265 github.com/goccy/spidermonkeywasm2go/p7.Fn4265
func Fn4265(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn4266 github.com/goccy/spidermonkeywasm2go/p7.Fn4266
func Fn4266(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4272 github.com/goccy/spidermonkeywasm2go/p5.Fn4272
func Fn4272(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4273 github.com/goccy/spidermonkeywasm2go/p6.Fn4273
func Fn4273(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4274 github.com/goccy/spidermonkeywasm2go/p5.Fn4274
func Fn4274(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4275 github.com/goccy/spidermonkeywasm2go/p4.Fn4275
func Fn4275(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4276 github.com/goccy/spidermonkeywasm2go/p6.Fn4276
func Fn4276(m *base.Module, l0 int32) int32

//go:linkname Fn4277 github.com/goccy/spidermonkeywasm2go/p4.Fn4277
func Fn4277(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4278 github.com/goccy/spidermonkeywasm2go/p6.Fn4278
func Fn4278(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4300 github.com/goccy/spidermonkeywasm2go/p6.Fn4300
func Fn4300(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4301 github.com/goccy/spidermonkeywasm2go/p7.Fn4301
func Fn4301(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4421 github.com/goccy/spidermonkeywasm2go/p5.Fn4421
func Fn4421(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4545 github.com/goccy/spidermonkeywasm2go/p2.Fn4545
func Fn4545(m *base.Module, l0 int32, l1 float64, l2 float64, l3 int32, l4 int32) int32

//go:linkname Fn4546 github.com/goccy/spidermonkeywasm2go/p7.Fn4546
func Fn4546(m *base.Module, l0 int32) int32

//go:linkname Fn4551 github.com/goccy/spidermonkeywasm2go/p6.Fn4551
func Fn4551(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4553 github.com/goccy/spidermonkeywasm2go/p4.Fn4553
func Fn4553(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn4555 github.com/goccy/spidermonkeywasm2go/p3.Fn4555
func Fn4555(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4556 github.com/goccy/spidermonkeywasm2go/p6.Fn4556
func Fn4556(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4560 github.com/goccy/spidermonkeywasm2go/p2.Fn4560
func Fn4560(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4570 github.com/goccy/spidermonkeywasm2go/p4.Fn4570
func Fn4570(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4577 github.com/goccy/spidermonkeywasm2go/p6.Fn4577
func Fn4577(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4661 github.com/goccy/spidermonkeywasm2go/p5.Fn4661
func Fn4661(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4669 github.com/goccy/spidermonkeywasm2go/p4.Fn4669
func Fn4669(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4671 github.com/goccy/spidermonkeywasm2go/p6.Fn4671
func Fn4671(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4672 github.com/goccy/spidermonkeywasm2go/p3.Fn4672
func Fn4672(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4681 github.com/goccy/spidermonkeywasm2go/p6.Fn4681
func Fn4681(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4684 github.com/goccy/spidermonkeywasm2go/p0.Fn4684
func Fn4684(m *base.Module, l0 int32)

//go:linkname Fn4701 github.com/goccy/spidermonkeywasm2go/p0.Fn4701
func Fn4701(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4704 github.com/goccy/spidermonkeywasm2go/p6.Fn4704
func Fn4704(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4719 github.com/goccy/spidermonkeywasm2go/p4.Fn4719
func Fn4719(m *base.Module, l0 int32)

//go:linkname Fn4721 github.com/goccy/spidermonkeywasm2go/p7.Fn4721
func Fn4721(m *base.Module, l0 int32)

//go:linkname Fn4722 github.com/goccy/spidermonkeywasm2go/p6.Fn4722
func Fn4722(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4723 github.com/goccy/spidermonkeywasm2go/p7.Fn4723
func Fn4723(m *base.Module, l0 int32) int32

//go:linkname Fn4724 github.com/goccy/spidermonkeywasm2go/p7.Fn4724
func Fn4724(m *base.Module, l0 int32) int32

//go:linkname Fn4725 github.com/goccy/spidermonkeywasm2go/p7.Fn4725
func Fn4725(m *base.Module, l0 int32) int32

//go:linkname Fn4726 github.com/goccy/spidermonkeywasm2go/p7.Fn4726
func Fn4726(m *base.Module, l0 int32) int32

//go:linkname Fn4727 github.com/goccy/spidermonkeywasm2go/p7.Fn4727
func Fn4727(m *base.Module, l0 int32) int32

//go:linkname Fn4728 github.com/goccy/spidermonkeywasm2go/p7.Fn4728
func Fn4728(m *base.Module, l0 int32) int32

//go:linkname Fn4729 github.com/goccy/spidermonkeywasm2go/p7.Fn4729
func Fn4729(m *base.Module, l0 int32) int32

//go:linkname Fn4803 github.com/goccy/spidermonkeywasm2go/p5.Fn4803
func Fn4803(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4805 github.com/goccy/spidermonkeywasm2go/p6.Fn4805
func Fn4805(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4806 github.com/goccy/spidermonkeywasm2go/p6.Fn4806
func Fn4806(m *base.Module, l0 int32)

//go:linkname Fn4808 github.com/goccy/spidermonkeywasm2go/p0.Fn4808
func Fn4808(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4825 github.com/goccy/spidermonkeywasm2go/p7.Fn4825
func Fn4825(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4826 github.com/goccy/spidermonkeywasm2go/p6.Fn4826
func Fn4826(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4827 github.com/goccy/spidermonkeywasm2go/p0.Fn4827
func Fn4827(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4830 github.com/goccy/spidermonkeywasm2go/p7.Fn4830
func Fn4830(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4870 github.com/goccy/spidermonkeywasm2go/p0.Fn4870
func Fn4870(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4913 github.com/goccy/spidermonkeywasm2go/p5.Fn4913
func Fn4913(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn4916 github.com/goccy/spidermonkeywasm2go/p0.Fn4916
func Fn4916(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5137 github.com/goccy/spidermonkeywasm2go/p6.Fn5137
func Fn5137(m *base.Module, l0 int32) int32

//go:linkname Fn5150 github.com/goccy/spidermonkeywasm2go/p6.Fn5150
func Fn5150(m *base.Module, l0 int32) int32

//go:linkname Fn5151 github.com/goccy/spidermonkeywasm2go/p3.Fn5151
func Fn5151(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5152 github.com/goccy/spidermonkeywasm2go/p5.Fn5152
func Fn5152(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5153 github.com/goccy/spidermonkeywasm2go/p0.Fn5153
func Fn5153(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5155 github.com/goccy/spidermonkeywasm2go/p0.Fn5155
func Fn5155(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5156 github.com/goccy/spidermonkeywasm2go/p7.Fn5156
func Fn5156(m *base.Module, l0 int32)

//go:linkname Fn5286 github.com/goccy/spidermonkeywasm2go/p0.Fn5286
func Fn5286(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5288 github.com/goccy/spidermonkeywasm2go/p5.Fn5288
func Fn5288(m *base.Module, l0 int32)

//go:linkname Fn5289 github.com/goccy/spidermonkeywasm2go/p0.Fn5289
func Fn5289(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5329 github.com/goccy/spidermonkeywasm2go/p0.Fn5329
func Fn5329(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn5334 github.com/goccy/spidermonkeywasm2go/p7.Fn5334
func Fn5334(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5341 github.com/goccy/spidermonkeywasm2go/p6.Fn5341
func Fn5341(m *base.Module, l0 int32) int32

//go:linkname Fn5354 github.com/goccy/spidermonkeywasm2go/p5.Fn5354
func Fn5354(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5355 github.com/goccy/spidermonkeywasm2go/p5.Fn5355
func Fn5355(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5361 github.com/goccy/spidermonkeywasm2go/p6.Fn5361
func Fn5361(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5365 github.com/goccy/spidermonkeywasm2go/p4.Fn5365
func Fn5365(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5369 github.com/goccy/spidermonkeywasm2go/p7.Fn5369
func Fn5369(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5370 github.com/goccy/spidermonkeywasm2go/p5.Fn5370
func Fn5370(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5389 github.com/goccy/spidermonkeywasm2go/p6.Fn5389
func Fn5389(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5404 github.com/goccy/spidermonkeywasm2go/p6.Fn5404
func Fn5404(m *base.Module, l0 int32) int32

//go:linkname Fn5405 github.com/goccy/spidermonkeywasm2go/p7.Fn5405
func Fn5405(m *base.Module, l0 int32) int32

//go:linkname Fn5406 github.com/goccy/spidermonkeywasm2go/p5.Fn5406
func Fn5406(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5408 github.com/goccy/spidermonkeywasm2go/p6.Fn5408
func Fn5408(m *base.Module, l0 int32) int32

//go:linkname Fn5409 github.com/goccy/spidermonkeywasm2go/p6.Fn5409
func Fn5409(m *base.Module, l0 int32) int32

//go:linkname Fn5410 github.com/goccy/spidermonkeywasm2go/p6.Fn5410
func Fn5410(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5411 github.com/goccy/spidermonkeywasm2go/p6.Fn5411
func Fn5411(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5412 github.com/goccy/spidermonkeywasm2go/p5.Fn5412
func Fn5412(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5413 github.com/goccy/spidermonkeywasm2go/p5.Fn5413
func Fn5413(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5415 github.com/goccy/spidermonkeywasm2go/p6.Fn5415
func Fn5415(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5418 github.com/goccy/spidermonkeywasm2go/p3.Fn5418
func Fn5418(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5423 github.com/goccy/spidermonkeywasm2go/p5.Fn5423
func Fn5423(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5431 github.com/goccy/spidermonkeywasm2go/p6.Fn5431
func Fn5431(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5440 github.com/goccy/spidermonkeywasm2go/p5.Fn5440
func Fn5440(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5510 github.com/goccy/spidermonkeywasm2go/p0.Fn5510
func Fn5510(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5570 github.com/goccy/spidermonkeywasm2go/p7.Fn5570
func Fn5570(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5572 github.com/goccy/spidermonkeywasm2go/p5.Fn5572
func Fn5572(m *base.Module, l0 int32) int32

//go:linkname Fn5576 github.com/goccy/spidermonkeywasm2go/p0.Fn5576
func Fn5576(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5577 github.com/goccy/spidermonkeywasm2go/p0.Fn5577
func Fn5577(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5578 github.com/goccy/spidermonkeywasm2go/p6.Fn5578
func Fn5578(m *base.Module, l0 int32)

//go:linkname Fn5583 github.com/goccy/spidermonkeywasm2go/p5.Fn5583
func Fn5583(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5605 github.com/goccy/spidermonkeywasm2go/p6.Fn5605
func Fn5605(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5609 github.com/goccy/spidermonkeywasm2go/p3.Fn5609
func Fn5609(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5622 github.com/goccy/spidermonkeywasm2go/p5.Fn5622
func Fn5622(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5623 github.com/goccy/spidermonkeywasm2go/p0.Fn5623
func Fn5623(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5654 github.com/goccy/spidermonkeywasm2go/p6.Fn5654
func Fn5654(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5656 github.com/goccy/spidermonkeywasm2go/p6.Fn5656
func Fn5656(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5657 github.com/goccy/spidermonkeywasm2go/p7.Fn5657
func Fn5657(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5658 github.com/goccy/spidermonkeywasm2go/p7.Fn5658
func Fn5658(m *base.Module, l0 int32)

//go:linkname Fn5660 github.com/goccy/spidermonkeywasm2go/p5.Fn5660
func Fn5660(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5664 github.com/goccy/spidermonkeywasm2go/p6.Fn5664
func Fn5664(m *base.Module, l0 int32)

//go:linkname Fn5668 github.com/goccy/spidermonkeywasm2go/p6.Fn5668
func Fn5668(m *base.Module, l0 int32)

//go:linkname Fn5671 github.com/goccy/spidermonkeywasm2go/p6.Fn5671
func Fn5671(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5673 github.com/goccy/spidermonkeywasm2go/p7.Fn5673
func Fn5673(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5692 github.com/goccy/spidermonkeywasm2go/p6.Fn5692
func Fn5692(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5695 github.com/goccy/spidermonkeywasm2go/p5.Fn5695
func Fn5695(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5708 github.com/goccy/spidermonkeywasm2go/p7.Fn5708
func Fn5708(m *base.Module, l0 int32) int32

//go:linkname Fn5740 github.com/goccy/spidermonkeywasm2go/p7.Fn5740
func Fn5740(m *base.Module, l0 int32)

//go:linkname Fn5753 github.com/goccy/spidermonkeywasm2go/p6.Fn5753
func Fn5753(m *base.Module, l0 int32)

//go:linkname Fn5756 github.com/goccy/spidermonkeywasm2go/p6.Fn5756
func Fn5756(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5823 github.com/goccy/spidermonkeywasm2go/p6.Fn5823
func Fn5823(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5824 github.com/goccy/spidermonkeywasm2go/p6.Fn5824
func Fn5824(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5825 github.com/goccy/spidermonkeywasm2go/p6.Fn5825
func Fn5825(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5826 github.com/goccy/spidermonkeywasm2go/p6.Fn5826
func Fn5826(m *base.Module, l0 int32)

//go:linkname Fn5850 github.com/goccy/spidermonkeywasm2go/p6.Fn5850
func Fn5850(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5853 github.com/goccy/spidermonkeywasm2go/p6.Fn5853
func Fn5853(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5854 github.com/goccy/spidermonkeywasm2go/p5.Fn5854
func Fn5854(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5855 github.com/goccy/spidermonkeywasm2go/p5.Fn5855
func Fn5855(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5856 github.com/goccy/spidermonkeywasm2go/p5.Fn5856
func Fn5856(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5857 github.com/goccy/spidermonkeywasm2go/p2.Fn5857
func Fn5857(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5858 github.com/goccy/spidermonkeywasm2go/p5.Fn5858
func Fn5858(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5859 github.com/goccy/spidermonkeywasm2go/p5.Fn5859
func Fn5859(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5860 github.com/goccy/spidermonkeywasm2go/p5.Fn5860
func Fn5860(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5861 github.com/goccy/spidermonkeywasm2go/p5.Fn5861
func Fn5861(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5862 github.com/goccy/spidermonkeywasm2go/p5.Fn5862
func Fn5862(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5863 github.com/goccy/spidermonkeywasm2go/p6.Fn5863
func Fn5863(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5864 github.com/goccy/spidermonkeywasm2go/p5.Fn5864
func Fn5864(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5865 github.com/goccy/spidermonkeywasm2go/p5.Fn5865
func Fn5865(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5866 github.com/goccy/spidermonkeywasm2go/p5.Fn5866
func Fn5866(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5868 github.com/goccy/spidermonkeywasm2go/p2.Fn5868
func Fn5868(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5869 github.com/goccy/spidermonkeywasm2go/p6.Fn5869
func Fn5869(m *base.Module, l0 int32)

//go:linkname Fn5885 github.com/goccy/spidermonkeywasm2go/p4.Fn5885
func Fn5885(m *base.Module, l0 int32)

//go:linkname Fn5888 github.com/goccy/spidermonkeywasm2go/p5.Fn5888
func Fn5888(m *base.Module, l0 int32) int32

//go:linkname Fn5945 github.com/goccy/spidermonkeywasm2go/p4.Fn5945
func Fn5945(m *base.Module, l0 int32)

//go:linkname Fn5946 github.com/goccy/spidermonkeywasm2go/p7.Fn5946
func Fn5946(m *base.Module, l0 int32)

//go:linkname Fn5956 github.com/goccy/spidermonkeywasm2go/p3.Fn5956
func Fn5956(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5957 github.com/goccy/spidermonkeywasm2go/p3.Fn5957
func Fn5957(m *base.Module, l0 int32) int32

//go:linkname Fn5958 github.com/goccy/spidermonkeywasm2go/p6.Fn5958
func Fn5958(m *base.Module, l0 int32) int32

//go:linkname Fn5961 github.com/goccy/spidermonkeywasm2go/p5.Fn5961
func Fn5961(m *base.Module, l0 int32)

//go:linkname Fn5962 github.com/goccy/spidermonkeywasm2go/p4.Fn5962
func Fn5962(m *base.Module, l0 int32)

//go:linkname Fn5971 github.com/goccy/spidermonkeywasm2go/p6.Fn5971
func Fn5971(m *base.Module, l0 int32)

//go:linkname Fn5974 github.com/goccy/spidermonkeywasm2go/p5.Fn5974
func Fn5974(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5980 github.com/goccy/spidermonkeywasm2go/p6.Fn5980
func Fn5980(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5982 github.com/goccy/spidermonkeywasm2go/p7.Fn5982
func Fn5982(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn6062 github.com/goccy/spidermonkeywasm2go/p7.Fn6062
func Fn6062(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6094 github.com/goccy/spidermonkeywasm2go/p7.Fn6094
func Fn6094(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6108 github.com/goccy/spidermonkeywasm2go/p7.Fn6108
func Fn6108(m *base.Module)

//go:linkname Fn6110 github.com/goccy/spidermonkeywasm2go/p6.Fn6110
func Fn6110(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6212 github.com/goccy/spidermonkeywasm2go/p6.Fn6212
func Fn6212(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6220 github.com/goccy/spidermonkeywasm2go/p7.Fn6220
func Fn6220(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6221 github.com/goccy/spidermonkeywasm2go/p6.Fn6221
func Fn6221(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6224 github.com/goccy/spidermonkeywasm2go/p6.Fn6224
func Fn6224(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6229 github.com/goccy/spidermonkeywasm2go/p6.Fn6229
func Fn6229(m *base.Module, l0 int32) int32

//go:linkname Fn6230 github.com/goccy/spidermonkeywasm2go/p5.Fn6230
func Fn6230(m *base.Module, l0 int32)

//go:linkname Fn6231 github.com/goccy/spidermonkeywasm2go/p5.Fn6231
func Fn6231(m *base.Module, l0 int32)

//go:linkname Fn6232 github.com/goccy/spidermonkeywasm2go/p6.Fn6232
func Fn6232(m *base.Module, l0 int32) int32

//go:linkname Fn6233 github.com/goccy/spidermonkeywasm2go/p6.Fn6233
func Fn6233(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6234 github.com/goccy/spidermonkeywasm2go/p2.Fn6234
func Fn6234(m *base.Module, l0 int32) int32

//go:linkname Fn6236 github.com/goccy/spidermonkeywasm2go/p3.Fn6236
func Fn6236(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn6237 github.com/goccy/spidermonkeywasm2go/p6.Fn6237
func Fn6237(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6238 github.com/goccy/spidermonkeywasm2go/p4.Fn6238
func Fn6238(m *base.Module, l0 int32) int32

//go:linkname Fn6239 github.com/goccy/spidermonkeywasm2go/p6.Fn6239
func Fn6239(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6240 github.com/goccy/spidermonkeywasm2go/p7.Fn6240
func Fn6240(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6241 github.com/goccy/spidermonkeywasm2go/p3.Fn6241
func Fn6241(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6242 github.com/goccy/spidermonkeywasm2go/p7.Fn6242
func Fn6242(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6243 github.com/goccy/spidermonkeywasm2go/p6.Fn6243
func Fn6243(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6245 github.com/goccy/spidermonkeywasm2go/p5.Fn6245
func Fn6245(m *base.Module, l0 int32)

//go:linkname Fn6246 github.com/goccy/spidermonkeywasm2go/p4.Fn6246
func Fn6246(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6247 github.com/goccy/spidermonkeywasm2go/p6.Fn6247
func Fn6247(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6249 github.com/goccy/spidermonkeywasm2go/p6.Fn6249
func Fn6249(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn6250 github.com/goccy/spidermonkeywasm2go/p5.Fn6250
func Fn6250(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6294 github.com/goccy/spidermonkeywasm2go/p5.Fn6294
func Fn6294(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6296 github.com/goccy/spidermonkeywasm2go/p3.Fn6296
func Fn6296(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6348 github.com/goccy/spidermonkeywasm2go/p2.Fn6348
func Fn6348(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6668 github.com/goccy/spidermonkeywasm2go/p7.Fn6668
func Fn6668(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6669 github.com/goccy/spidermonkeywasm2go/p6.Fn6669
func Fn6669(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6670 github.com/goccy/spidermonkeywasm2go/p6.Fn6670
func Fn6670(m *base.Module, l0 int32) int32

//go:linkname Fn6683 github.com/goccy/spidermonkeywasm2go/p6.Fn6683
func Fn6683(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6691 github.com/goccy/spidermonkeywasm2go/p7.Fn6691
func Fn6691(m *base.Module, l0 int32)

//go:linkname Fn6692 github.com/goccy/spidermonkeywasm2go/p7.Fn6692
func Fn6692(m *base.Module, l0 int32)

//go:linkname Fn6693 github.com/goccy/spidermonkeywasm2go/p6.Fn6693
func Fn6693(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn6695 github.com/goccy/spidermonkeywasm2go/p6.Fn6695
func Fn6695(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6696 github.com/goccy/spidermonkeywasm2go/p5.Fn6696
func Fn6696(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6697 github.com/goccy/spidermonkeywasm2go/p4.Fn6697
func Fn6697(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6698 github.com/goccy/spidermonkeywasm2go/p7.Fn6698
func Fn6698(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6699 github.com/goccy/spidermonkeywasm2go/p7.Fn6699
func Fn6699(m *base.Module, l0 int32)

//go:linkname Fn6700 github.com/goccy/spidermonkeywasm2go/p6.Fn6700
func Fn6700(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6701 github.com/goccy/spidermonkeywasm2go/p6.Fn6701
func Fn6701(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6702 github.com/goccy/spidermonkeywasm2go/p6.Fn6702
func Fn6702(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6703 github.com/goccy/spidermonkeywasm2go/p6.Fn6703
func Fn6703(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6704 github.com/goccy/spidermonkeywasm2go/p6.Fn6704
func Fn6704(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6705 github.com/goccy/spidermonkeywasm2go/p7.Fn6705
func Fn6705(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6706 github.com/goccy/spidermonkeywasm2go/p5.Fn6706
func Fn6706(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn6707 github.com/goccy/spidermonkeywasm2go/p6.Fn6707
func Fn6707(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6708 github.com/goccy/spidermonkeywasm2go/p4.Fn6708
func Fn6708(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6709 github.com/goccy/spidermonkeywasm2go/p6.Fn6709
func Fn6709(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn6710 github.com/goccy/spidermonkeywasm2go/p7.Fn6710
func Fn6710(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6711 github.com/goccy/spidermonkeywasm2go/p6.Fn6711
func Fn6711(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6712 github.com/goccy/spidermonkeywasm2go/p7.Fn6712
func Fn6712(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6713 github.com/goccy/spidermonkeywasm2go/p7.Fn6713
func Fn6713(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6714 github.com/goccy/spidermonkeywasm2go/p5.Fn6714
func Fn6714(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6715 github.com/goccy/spidermonkeywasm2go/p7.Fn6715
func Fn6715(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6716 github.com/goccy/spidermonkeywasm2go/p7.Fn6716
func Fn6716(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6717 github.com/goccy/spidermonkeywasm2go/p7.Fn6717
func Fn6717(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6718 github.com/goccy/spidermonkeywasm2go/p6.Fn6718
func Fn6718(m *base.Module, l0 int32) int32

//go:linkname Fn6719 github.com/goccy/spidermonkeywasm2go/p7.Fn6719
func Fn6719(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6721 github.com/goccy/spidermonkeywasm2go/p6.Fn6721
func Fn6721(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6722 github.com/goccy/spidermonkeywasm2go/p7.Fn6722
func Fn6722(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6724 github.com/goccy/spidermonkeywasm2go/p7.Fn6724
func Fn6724(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6726 github.com/goccy/spidermonkeywasm2go/p7.Fn6726
func Fn6726(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6727 github.com/goccy/spidermonkeywasm2go/p6.Fn6727
func Fn6727(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6729 github.com/goccy/spidermonkeywasm2go/p7.Fn6729
func Fn6729(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6736 github.com/goccy/spidermonkeywasm2go/p7.Fn6736
func Fn6736(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6737 github.com/goccy/spidermonkeywasm2go/p7.Fn6737
func Fn6737(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6738 github.com/goccy/spidermonkeywasm2go/p7.Fn6738
func Fn6738(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6740 github.com/goccy/spidermonkeywasm2go/p4.Fn6740
func Fn6740(m *base.Module, l0 int32) int32

//go:linkname Fn6741 github.com/goccy/spidermonkeywasm2go/p4.Fn6741
func Fn6741(m *base.Module, l0 int32) int32

//go:linkname Fn6742 github.com/goccy/spidermonkeywasm2go/p5.Fn6742
func Fn6742(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6743 github.com/goccy/spidermonkeywasm2go/p5.Fn6743
func Fn6743(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6744 github.com/goccy/spidermonkeywasm2go/p5.Fn6744
func Fn6744(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6745 github.com/goccy/spidermonkeywasm2go/p4.Fn6745
func Fn6745(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6746 github.com/goccy/spidermonkeywasm2go/p5.Fn6746
func Fn6746(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6747 github.com/goccy/spidermonkeywasm2go/p4.Fn6747
func Fn6747(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6748 github.com/goccy/spidermonkeywasm2go/p3.Fn6748
func Fn6748(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6749 github.com/goccy/spidermonkeywasm2go/p5.Fn6749
func Fn6749(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6750 github.com/goccy/spidermonkeywasm2go/p4.Fn6750
func Fn6750(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6751 github.com/goccy/spidermonkeywasm2go/p5.Fn6751
func Fn6751(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6752 github.com/goccy/spidermonkeywasm2go/p4.Fn6752
func Fn6752(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6753 github.com/goccy/spidermonkeywasm2go/p5.Fn6753
func Fn6753(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6754 github.com/goccy/spidermonkeywasm2go/p4.Fn6754
func Fn6754(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6755 github.com/goccy/spidermonkeywasm2go/p4.Fn6755
func Fn6755(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6756 github.com/goccy/spidermonkeywasm2go/p7.Fn6756
func Fn6756(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6757 github.com/goccy/spidermonkeywasm2go/p6.Fn6757
func Fn6757(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6759 github.com/goccy/spidermonkeywasm2go/p7.Fn6759
func Fn6759(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6760 github.com/goccy/spidermonkeywasm2go/p7.Fn6760
func Fn6760(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6761 github.com/goccy/spidermonkeywasm2go/p7.Fn6761
func Fn6761(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6765 github.com/goccy/spidermonkeywasm2go/p7.Fn6765
func Fn6765(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6768 github.com/goccy/spidermonkeywasm2go/p4.Fn6768
func Fn6768(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6769 github.com/goccy/spidermonkeywasm2go/p6.Fn6769
func Fn6769(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6770 github.com/goccy/spidermonkeywasm2go/p5.Fn6770
func Fn6770(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6771 github.com/goccy/spidermonkeywasm2go/p6.Fn6771
func Fn6771(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6773 github.com/goccy/spidermonkeywasm2go/p7.Fn6773
func Fn6773(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6774 github.com/goccy/spidermonkeywasm2go/p7.Fn6774
func Fn6774(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6775 github.com/goccy/spidermonkeywasm2go/p7.Fn6775
func Fn6775(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6776 github.com/goccy/spidermonkeywasm2go/p7.Fn6776
func Fn6776(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6780 github.com/goccy/spidermonkeywasm2go/p7.Fn6780
func Fn6780(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6781 github.com/goccy/spidermonkeywasm2go/p7.Fn6781
func Fn6781(m *base.Module, l0 int32) int32

//go:linkname Fn6782 github.com/goccy/spidermonkeywasm2go/p7.Fn6782
func Fn6782(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6785 github.com/goccy/spidermonkeywasm2go/p7.Fn6785
func Fn6785(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6786 github.com/goccy/spidermonkeywasm2go/p7.Fn6786
func Fn6786(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6787 github.com/goccy/spidermonkeywasm2go/p4.Fn6787
func Fn6787(m *base.Module, l0 int32) int32

//go:linkname Fn6788 github.com/goccy/spidermonkeywasm2go/p6.Fn6788
func Fn6788(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6789 github.com/goccy/spidermonkeywasm2go/p7.Fn6789
func Fn6789(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6790 github.com/goccy/spidermonkeywasm2go/p7.Fn6790
func Fn6790(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6791 github.com/goccy/spidermonkeywasm2go/p7.Fn6791
func Fn6791(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6792 github.com/goccy/spidermonkeywasm2go/p7.Fn6792
func Fn6792(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6793 github.com/goccy/spidermonkeywasm2go/p7.Fn6793
func Fn6793(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6796 github.com/goccy/spidermonkeywasm2go/p5.Fn6796
func Fn6796(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6797 github.com/goccy/spidermonkeywasm2go/p5.Fn6797
func Fn6797(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6798 github.com/goccy/spidermonkeywasm2go/p6.Fn6798
func Fn6798(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6799 github.com/goccy/spidermonkeywasm2go/p6.Fn6799
func Fn6799(m *base.Module, l0 int32) int32

//go:linkname Fn6800 github.com/goccy/spidermonkeywasm2go/p6.Fn6800
func Fn6800(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6801 github.com/goccy/spidermonkeywasm2go/p7.Fn6801
func Fn6801(m *base.Module, l0 int32)

//go:linkname Fn6802 github.com/goccy/spidermonkeywasm2go/p7.Fn6802
func Fn6802(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6803 github.com/goccy/spidermonkeywasm2go/p6.Fn6803
func Fn6803(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6805 github.com/goccy/spidermonkeywasm2go/p7.Fn6805
func Fn6805(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6806 github.com/goccy/spidermonkeywasm2go/p7.Fn6806
func Fn6806(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6807 github.com/goccy/spidermonkeywasm2go/p7.Fn6807
func Fn6807(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6809 github.com/goccy/spidermonkeywasm2go/p6.Fn6809
func Fn6809(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6810 github.com/goccy/spidermonkeywasm2go/p6.Fn6810
func Fn6810(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6811 github.com/goccy/spidermonkeywasm2go/p7.Fn6811
func Fn6811(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6812 github.com/goccy/spidermonkeywasm2go/p4.Fn6812
func Fn6812(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6813 github.com/goccy/spidermonkeywasm2go/p6.Fn6813
func Fn6813(m *base.Module, l0 int32)

//go:linkname Fn6814 github.com/goccy/spidermonkeywasm2go/p7.Fn6814
func Fn6814(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6815 github.com/goccy/spidermonkeywasm2go/p7.Fn6815
func Fn6815(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6816 github.com/goccy/spidermonkeywasm2go/p7.Fn6816
func Fn6816(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6817 github.com/goccy/spidermonkeywasm2go/p7.Fn6817
func Fn6817(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6818 github.com/goccy/spidermonkeywasm2go/p6.Fn6818
func Fn6818(m *base.Module, l0 int32) int32

//go:linkname Fn6819 github.com/goccy/spidermonkeywasm2go/p6.Fn6819
func Fn6819(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6820 github.com/goccy/spidermonkeywasm2go/p7.Fn6820
func Fn6820(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6821 github.com/goccy/spidermonkeywasm2go/p7.Fn6821
func Fn6821(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6822 github.com/goccy/spidermonkeywasm2go/p7.Fn6822
func Fn6822(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6823 github.com/goccy/spidermonkeywasm2go/p7.Fn6823
func Fn6823(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6824 github.com/goccy/spidermonkeywasm2go/p7.Fn6824
func Fn6824(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6825 github.com/goccy/spidermonkeywasm2go/p7.Fn6825
func Fn6825(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6826 github.com/goccy/spidermonkeywasm2go/p7.Fn6826
func Fn6826(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6827 github.com/goccy/spidermonkeywasm2go/p7.Fn6827
func Fn6827(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6828 github.com/goccy/spidermonkeywasm2go/p7.Fn6828
func Fn6828(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6829 github.com/goccy/spidermonkeywasm2go/p7.Fn6829
func Fn6829(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6830 github.com/goccy/spidermonkeywasm2go/p6.Fn6830
func Fn6830(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6831 github.com/goccy/spidermonkeywasm2go/p7.Fn6831
func Fn6831(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6832 github.com/goccy/spidermonkeywasm2go/p7.Fn6832
func Fn6832(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6833 github.com/goccy/spidermonkeywasm2go/p7.Fn6833
func Fn6833(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6834 github.com/goccy/spidermonkeywasm2go/p6.Fn6834
func Fn6834(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6835 github.com/goccy/spidermonkeywasm2go/p7.Fn6835
func Fn6835(m *base.Module, l0 int32) int32

//go:linkname Fn6836 github.com/goccy/spidermonkeywasm2go/p7.Fn6836
func Fn6836(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6839 github.com/goccy/spidermonkeywasm2go/p5.Fn6839
func Fn6839(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6840 github.com/goccy/spidermonkeywasm2go/p3.Fn6840
func Fn6840(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6842 github.com/goccy/spidermonkeywasm2go/p6.Fn6842
func Fn6842(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6843 github.com/goccy/spidermonkeywasm2go/p5.Fn6843
func Fn6843(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6850 github.com/goccy/spidermonkeywasm2go/p5.Fn6850
func Fn6850(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6851 github.com/goccy/spidermonkeywasm2go/p7.Fn6851
func Fn6851(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6852 github.com/goccy/spidermonkeywasm2go/p7.Fn6852
func Fn6852(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6853 github.com/goccy/spidermonkeywasm2go/p7.Fn6853
func Fn6853(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6854 github.com/goccy/spidermonkeywasm2go/p7.Fn6854
func Fn6854(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6855 github.com/goccy/spidermonkeywasm2go/p7.Fn6855
func Fn6855(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6856 github.com/goccy/spidermonkeywasm2go/p7.Fn6856
func Fn6856(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6857 github.com/goccy/spidermonkeywasm2go/p7.Fn6857
func Fn6857(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6858 github.com/goccy/spidermonkeywasm2go/p7.Fn6858
func Fn6858(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6859 github.com/goccy/spidermonkeywasm2go/p7.Fn6859
func Fn6859(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6860 github.com/goccy/spidermonkeywasm2go/p6.Fn6860
func Fn6860(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6861 github.com/goccy/spidermonkeywasm2go/p6.Fn6861
func Fn6861(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6862 github.com/goccy/spidermonkeywasm2go/p6.Fn6862
func Fn6862(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6870 github.com/goccy/spidermonkeywasm2go/p7.Fn6870
func Fn6870(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6895 github.com/goccy/spidermonkeywasm2go/p7.Fn6895
func Fn6895(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6896 github.com/goccy/spidermonkeywasm2go/p5.Fn6896
func Fn6896(m *base.Module)

//go:linkname Fn6897 github.com/goccy/spidermonkeywasm2go/p7.Fn6897
func Fn6897(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6899 github.com/goccy/spidermonkeywasm2go/p3.Fn6899
func Fn6899(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6930 github.com/goccy/spidermonkeywasm2go/p6.Fn6930
func Fn6930(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6943 github.com/goccy/spidermonkeywasm2go/p6.Fn6943
func Fn6943(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6950 github.com/goccy/spidermonkeywasm2go/p7.Fn6950
func Fn6950(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6952 github.com/goccy/spidermonkeywasm2go/p0.Fn6952
func Fn6952(m *base.Module, l0 int32) int32

//go:linkname Fn6955 github.com/goccy/spidermonkeywasm2go/p0.Fn6955
func Fn6955(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6956 github.com/goccy/spidermonkeywasm2go/p0.Fn6956
func Fn6956(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6957 github.com/goccy/spidermonkeywasm2go/p0.Fn6957
func Fn6957(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6982 github.com/goccy/spidermonkeywasm2go/p7.Fn6982
func Fn6982(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6984 github.com/goccy/spidermonkeywasm2go/p7.Fn6984
func Fn6984(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6988 github.com/goccy/spidermonkeywasm2go/p4.Fn6988
func Fn6988(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7015 github.com/goccy/spidermonkeywasm2go/p7.Fn7015
func Fn7015(m *base.Module, l0 int32)

//go:linkname Fn7016 github.com/goccy/spidermonkeywasm2go/p7.Fn7016
func Fn7016(m *base.Module, l0 int32)

//go:linkname Fn7028 github.com/goccy/spidermonkeywasm2go/p5.Fn7028
func Fn7028(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7035 github.com/goccy/spidermonkeywasm2go/p7.Fn7035
func Fn7035(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7083 github.com/goccy/spidermonkeywasm2go/p6.Fn7083
func Fn7083(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7085 github.com/goccy/spidermonkeywasm2go/p4.Fn7085
func Fn7085(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7086 github.com/goccy/spidermonkeywasm2go/p2.Fn7086
func Fn7086(m *base.Module, l0 int32) int32

//go:linkname Fn7101 github.com/goccy/spidermonkeywasm2go/p5.Fn7101
func Fn7101(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7102 github.com/goccy/spidermonkeywasm2go/p5.Fn7102
func Fn7102(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7118 github.com/goccy/spidermonkeywasm2go/p6.Fn7118
func Fn7118(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7119 github.com/goccy/spidermonkeywasm2go/p6.Fn7119
func Fn7119(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7121 github.com/goccy/spidermonkeywasm2go/p7.Fn7121
func Fn7121(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7122 github.com/goccy/spidermonkeywasm2go/p6.Fn7122
func Fn7122(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7125 github.com/goccy/spidermonkeywasm2go/p5.Fn7125
func Fn7125(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn7133 github.com/goccy/spidermonkeywasm2go/p4.Fn7133
func Fn7133(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7143 github.com/goccy/spidermonkeywasm2go/p6.Fn7143
func Fn7143(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7145 github.com/goccy/spidermonkeywasm2go/p4.Fn7145
func Fn7145(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7147 github.com/goccy/spidermonkeywasm2go/p6.Fn7147
func Fn7147(m *base.Module, l0 int32) int32

//go:linkname Fn7148 github.com/goccy/spidermonkeywasm2go/p6.Fn7148
func Fn7148(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn7152 github.com/goccy/spidermonkeywasm2go/p5.Fn7152
func Fn7152(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn7153 github.com/goccy/spidermonkeywasm2go/p7.Fn7153
func Fn7153(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7180 github.com/goccy/spidermonkeywasm2go/p3.Fn7180
func Fn7180(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7181 github.com/goccy/spidermonkeywasm2go/p6.Fn7181
func Fn7181(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7182 github.com/goccy/spidermonkeywasm2go/p5.Fn7182
func Fn7182(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7183 github.com/goccy/spidermonkeywasm2go/p5.Fn7183
func Fn7183(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7184 github.com/goccy/spidermonkeywasm2go/p6.Fn7184
func Fn7184(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7185 github.com/goccy/spidermonkeywasm2go/p7.Fn7185
func Fn7185(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7186 github.com/goccy/spidermonkeywasm2go/p5.Fn7186
func Fn7186(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7187 github.com/goccy/spidermonkeywasm2go/p6.Fn7187
func Fn7187(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7188 github.com/goccy/spidermonkeywasm2go/p5.Fn7188
func Fn7188(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7189 github.com/goccy/spidermonkeywasm2go/p5.Fn7189
func Fn7189(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7190 github.com/goccy/spidermonkeywasm2go/p5.Fn7190
func Fn7190(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7191 github.com/goccy/spidermonkeywasm2go/p6.Fn7191
func Fn7191(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7192 github.com/goccy/spidermonkeywasm2go/p2.Fn7192
func Fn7192(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7196 github.com/goccy/spidermonkeywasm2go/p4.Fn7196
func Fn7196(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7197 github.com/goccy/spidermonkeywasm2go/p6.Fn7197
func Fn7197(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7198 github.com/goccy/spidermonkeywasm2go/p6.Fn7198
func Fn7198(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7202 github.com/goccy/spidermonkeywasm2go/p6.Fn7202
func Fn7202(m *base.Module, l0 int32) int32

//go:linkname Fn7204 github.com/goccy/spidermonkeywasm2go/p5.Fn7204
func Fn7204(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7206 github.com/goccy/spidermonkeywasm2go/p2.Fn7206
func Fn7206(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7220 github.com/goccy/spidermonkeywasm2go/p6.Fn7220
func Fn7220(m *base.Module, l0 int32) int32

//go:linkname Fn7221 github.com/goccy/spidermonkeywasm2go/p7.Fn7221
func Fn7221(m *base.Module, l0 int32)

//go:linkname Fn7222 github.com/goccy/spidermonkeywasm2go/p6.Fn7222
func Fn7222(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7229 github.com/goccy/spidermonkeywasm2go/p4.Fn7229
func Fn7229(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7230 github.com/goccy/spidermonkeywasm2go/p5.Fn7230
func Fn7230(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7232 github.com/goccy/spidermonkeywasm2go/p7.Fn7232
func Fn7232(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7237 github.com/goccy/spidermonkeywasm2go/p6.Fn7237
func Fn7237(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7255 github.com/goccy/spidermonkeywasm2go/p5.Fn7255
func Fn7255(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7270 github.com/goccy/spidermonkeywasm2go/p6.Fn7270
func Fn7270(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7279 github.com/goccy/spidermonkeywasm2go/p7.Fn7279
func Fn7279(m *base.Module, l0 int32)

//go:linkname Fn7350 github.com/goccy/spidermonkeywasm2go/p6.Fn7350
func Fn7350(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7367 github.com/goccy/spidermonkeywasm2go/p6.Fn7367
func Fn7367(m *base.Module, l0 int32) int32

//go:linkname Fn7372 github.com/goccy/spidermonkeywasm2go/p5.Fn7372
func Fn7372(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7376 github.com/goccy/spidermonkeywasm2go/p6.Fn7376
func Fn7376(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7377 github.com/goccy/spidermonkeywasm2go/p6.Fn7377
func Fn7377(m *base.Module, l0 int32)

//go:linkname Fn7378 github.com/goccy/spidermonkeywasm2go/p6.Fn7378
func Fn7378(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7381 github.com/goccy/spidermonkeywasm2go/p6.Fn7381
func Fn7381(m *base.Module, l0 int32) int32

//go:linkname Fn7383 github.com/goccy/spidermonkeywasm2go/p4.Fn7383
func Fn7383(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7386 github.com/goccy/spidermonkeywasm2go/p7.Fn7386
func Fn7386(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7387 github.com/goccy/spidermonkeywasm2go/p6.Fn7387
func Fn7387(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn7390 github.com/goccy/spidermonkeywasm2go/p6.Fn7390
func Fn7390(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7391 github.com/goccy/spidermonkeywasm2go/p6.Fn7391
func Fn7391(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7398 github.com/goccy/spidermonkeywasm2go/p4.Fn7398
func Fn7398(m *base.Module, l0 int32) int32

//go:linkname Fn7404 github.com/goccy/spidermonkeywasm2go/p6.Fn7404
func Fn7404(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64) int32

//go:linkname Fn7405 github.com/goccy/spidermonkeywasm2go/p5.Fn7405
func Fn7405(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7406 github.com/goccy/spidermonkeywasm2go/p7.Fn7406
func Fn7406(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7407 github.com/goccy/spidermonkeywasm2go/p7.Fn7407
func Fn7407(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7408 github.com/goccy/spidermonkeywasm2go/p7.Fn7408
func Fn7408(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7409 github.com/goccy/spidermonkeywasm2go/p7.Fn7409
func Fn7409(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7410 github.com/goccy/spidermonkeywasm2go/p7.Fn7410
func Fn7410(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7411 github.com/goccy/spidermonkeywasm2go/p6.Fn7411
func Fn7411(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7412 github.com/goccy/spidermonkeywasm2go/p6.Fn7412
func Fn7412(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn7413 github.com/goccy/spidermonkeywasm2go/p6.Fn7413
func Fn7413(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7414 github.com/goccy/spidermonkeywasm2go/p6.Fn7414
func Fn7414(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7415 github.com/goccy/spidermonkeywasm2go/p7.Fn7415
func Fn7415(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7416 github.com/goccy/spidermonkeywasm2go/p5.Fn7416
func Fn7416(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7417 github.com/goccy/spidermonkeywasm2go/p6.Fn7417
func Fn7417(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7418 github.com/goccy/spidermonkeywasm2go/p4.Fn7418
func Fn7418(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7419 github.com/goccy/spidermonkeywasm2go/p5.Fn7419
func Fn7419(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7420 github.com/goccy/spidermonkeywasm2go/p5.Fn7420
func Fn7420(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7421 github.com/goccy/spidermonkeywasm2go/p6.Fn7421
func Fn7421(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7422 github.com/goccy/spidermonkeywasm2go/p6.Fn7422
func Fn7422(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7423 github.com/goccy/spidermonkeywasm2go/p4.Fn7423
func Fn7423(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7424 github.com/goccy/spidermonkeywasm2go/p7.Fn7424
func Fn7424(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7425 github.com/goccy/spidermonkeywasm2go/p5.Fn7425
func Fn7425(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn7426 github.com/goccy/spidermonkeywasm2go/p5.Fn7426
func Fn7426(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7427 github.com/goccy/spidermonkeywasm2go/p4.Fn7427
func Fn7427(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7430 github.com/goccy/spidermonkeywasm2go/p6.Fn7430
func Fn7430(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7431 github.com/goccy/spidermonkeywasm2go/p6.Fn7431
func Fn7431(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7432 github.com/goccy/spidermonkeywasm2go/p5.Fn7432
func Fn7432(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7433 github.com/goccy/spidermonkeywasm2go/p6.Fn7433
func Fn7433(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7434 github.com/goccy/spidermonkeywasm2go/p3.Fn7434
func Fn7434(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7435 github.com/goccy/spidermonkeywasm2go/p6.Fn7435
func Fn7435(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7436 github.com/goccy/spidermonkeywasm2go/p5.Fn7436
func Fn7436(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7437 github.com/goccy/spidermonkeywasm2go/p6.Fn7437
func Fn7437(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7508 github.com/goccy/spidermonkeywasm2go/p2.Fn7508
func Fn7508(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8100 github.com/goccy/spidermonkeywasm2go/p7.Fn8100
func Fn8100(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8112 github.com/goccy/spidermonkeywasm2go/p5.Fn8112
func Fn8112(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8117 github.com/goccy/spidermonkeywasm2go/p6.Fn8117
func Fn8117(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8118 github.com/goccy/spidermonkeywasm2go/p6.Fn8118
func Fn8118(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8119 github.com/goccy/spidermonkeywasm2go/p7.Fn8119
func Fn8119(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8121 github.com/goccy/spidermonkeywasm2go/p6.Fn8121
func Fn8121(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8122 github.com/goccy/spidermonkeywasm2go/p6.Fn8122
func Fn8122(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8126 github.com/goccy/spidermonkeywasm2go/p3.Fn8126
func Fn8126(m *base.Module, l0 int32)

//go:linkname Fn8142 github.com/goccy/spidermonkeywasm2go/p6.Fn8142
func Fn8142(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8143 github.com/goccy/spidermonkeywasm2go/p6.Fn8143
func Fn8143(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8144 github.com/goccy/spidermonkeywasm2go/p6.Fn8144
func Fn8144(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8145 github.com/goccy/spidermonkeywasm2go/p6.Fn8145
func Fn8145(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8160 github.com/goccy/spidermonkeywasm2go/p6.Fn8160
func Fn8160(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8163 github.com/goccy/spidermonkeywasm2go/p7.Fn8163
func Fn8163(m *base.Module, l0 int32) int32

//go:linkname Fn8165 github.com/goccy/spidermonkeywasm2go/p7.Fn8165
func Fn8165(m *base.Module, l0 int32) int32

//go:linkname Fn8168 github.com/goccy/spidermonkeywasm2go/p6.Fn8168
func Fn8168(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8174 github.com/goccy/spidermonkeywasm2go/p2.Fn8174
func Fn8174(m *base.Module) int32

//go:linkname Fn8177 github.com/goccy/spidermonkeywasm2go/p6.Fn8177
func Fn8177(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8179 github.com/goccy/spidermonkeywasm2go/p3.Fn8179
func Fn8179(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8190 github.com/goccy/spidermonkeywasm2go/p4.Fn8190
func Fn8190(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8191 github.com/goccy/spidermonkeywasm2go/p7.Fn8191
func Fn8191(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8192 github.com/goccy/spidermonkeywasm2go/p7.Fn8192
func Fn8192(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8194 github.com/goccy/spidermonkeywasm2go/p7.Fn8194
func Fn8194(m *base.Module, l0 int32) int32

//go:linkname Fn8198 github.com/goccy/spidermonkeywasm2go/p6.Fn8198
func Fn8198(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32)

//go:linkname Fn8200 github.com/goccy/spidermonkeywasm2go/p5.Fn8200
func Fn8200(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8203 github.com/goccy/spidermonkeywasm2go/p5.Fn8203
func Fn8203(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8204 github.com/goccy/spidermonkeywasm2go/p4.Fn8204
func Fn8204(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn8205 github.com/goccy/spidermonkeywasm2go/p6.Fn8205
func Fn8205(m *base.Module, l0 int32) int32

//go:linkname Fn8209 github.com/goccy/spidermonkeywasm2go/p7.Fn8209
func Fn8209(m *base.Module, l0 int32) int32

//go:linkname Fn8213 github.com/goccy/spidermonkeywasm2go/p6.Fn8213
func Fn8213(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8214 github.com/goccy/spidermonkeywasm2go/p5.Fn8214
func Fn8214(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8217 github.com/goccy/spidermonkeywasm2go/p5.Fn8217
func Fn8217(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8218 github.com/goccy/spidermonkeywasm2go/p7.Fn8218
func Fn8218(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8277 github.com/goccy/spidermonkeywasm2go/p5.Fn8277
func Fn8277(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8300 github.com/goccy/spidermonkeywasm2go/p6.Fn8300
func Fn8300(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8330 github.com/goccy/spidermonkeywasm2go/p5.Fn8330
func Fn8330(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8331 github.com/goccy/spidermonkeywasm2go/p5.Fn8331
func Fn8331(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8343 github.com/goccy/spidermonkeywasm2go/p4.Fn8343
func Fn8343(m *base.Module, l0 int32) int32

//go:linkname Fn8355 github.com/goccy/spidermonkeywasm2go/p3.Fn8355
func Fn8355(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8363 github.com/goccy/spidermonkeywasm2go/p7.Fn8363
func Fn8363(m *base.Module, l0 int32)

//go:linkname Fn8365 github.com/goccy/spidermonkeywasm2go/p6.Fn8365
func Fn8365(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8366 github.com/goccy/spidermonkeywasm2go/p5.Fn8366
func Fn8366(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8367 github.com/goccy/spidermonkeywasm2go/p6.Fn8367
func Fn8367(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8368 github.com/goccy/spidermonkeywasm2go/p6.Fn8368
func Fn8368(m *base.Module, l0 int32)

//go:linkname Fn8370 github.com/goccy/spidermonkeywasm2go/p6.Fn8370
func Fn8370(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8371 github.com/goccy/spidermonkeywasm2go/p5.Fn8371
func Fn8371(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8372 github.com/goccy/spidermonkeywasm2go/p6.Fn8372
func Fn8372(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8384 github.com/goccy/spidermonkeywasm2go/p5.Fn8384
func Fn8384(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn8386 github.com/goccy/spidermonkeywasm2go/p6.Fn8386
func Fn8386(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8391 github.com/goccy/spidermonkeywasm2go/p5.Fn8391
func Fn8391(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8394 github.com/goccy/spidermonkeywasm2go/p6.Fn8394
func Fn8394(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8395 github.com/goccy/spidermonkeywasm2go/p3.Fn8395
func Fn8395(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn8396 github.com/goccy/spidermonkeywasm2go/p7.Fn8396
func Fn8396(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8397 github.com/goccy/spidermonkeywasm2go/p7.Fn8397
func Fn8397(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8398 github.com/goccy/spidermonkeywasm2go/p7.Fn8398
func Fn8398(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8400 github.com/goccy/spidermonkeywasm2go/p4.Fn8400
func Fn8400(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8401 github.com/goccy/spidermonkeywasm2go/p6.Fn8401
func Fn8401(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8413 github.com/goccy/spidermonkeywasm2go/p6.Fn8413
func Fn8413(m *base.Module, l0 int32) int32

//go:linkname Fn8414 github.com/goccy/spidermonkeywasm2go/p6.Fn8414
func Fn8414(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8418 github.com/goccy/spidermonkeywasm2go/p6.Fn8418
func Fn8418(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8419 github.com/goccy/spidermonkeywasm2go/p6.Fn8419
func Fn8419(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8420 github.com/goccy/spidermonkeywasm2go/p7.Fn8420
func Fn8420(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8421 github.com/goccy/spidermonkeywasm2go/p7.Fn8421
func Fn8421(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8422 github.com/goccy/spidermonkeywasm2go/p7.Fn8422
func Fn8422(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8423 github.com/goccy/spidermonkeywasm2go/p7.Fn8423
func Fn8423(m *base.Module, l0 int32) int32

//go:linkname Fn8424 github.com/goccy/spidermonkeywasm2go/p4.Fn8424
func Fn8424(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8425 github.com/goccy/spidermonkeywasm2go/p4.Fn8425
func Fn8425(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8426 github.com/goccy/spidermonkeywasm2go/p6.Fn8426
func Fn8426(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8427 github.com/goccy/spidermonkeywasm2go/p6.Fn8427
func Fn8427(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8437 github.com/goccy/spidermonkeywasm2go/p6.Fn8437
func Fn8437(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8440 github.com/goccy/spidermonkeywasm2go/p6.Fn8440
func Fn8440(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8442 github.com/goccy/spidermonkeywasm2go/p5.Fn8442
func Fn8442(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8445 github.com/goccy/spidermonkeywasm2go/p6.Fn8445
func Fn8445(m *base.Module, l0 int32) int32

//go:linkname Fn8451 github.com/goccy/spidermonkeywasm2go/p6.Fn8451
func Fn8451(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8467 github.com/goccy/spidermonkeywasm2go/p6.Fn8467
func Fn8467(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8471 github.com/goccy/spidermonkeywasm2go/p6.Fn8471
func Fn8471(m *base.Module, l0 int32) int32

//go:linkname Fn8497 github.com/goccy/spidermonkeywasm2go/p6.Fn8497
func Fn8497(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8499 github.com/goccy/spidermonkeywasm2go/p6.Fn8499
func Fn8499(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8504 github.com/goccy/spidermonkeywasm2go/p6.Fn8504
func Fn8504(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8505 github.com/goccy/spidermonkeywasm2go/p6.Fn8505
func Fn8505(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8517 github.com/goccy/spidermonkeywasm2go/p7.Fn8517
func Fn8517(m *base.Module, l0 int32)

//go:linkname Fn8521 github.com/goccy/spidermonkeywasm2go/p5.Fn8521
func Fn8521(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8525 github.com/goccy/spidermonkeywasm2go/p5.Fn8525
func Fn8525(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8526 github.com/goccy/spidermonkeywasm2go/p5.Fn8526
func Fn8526(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8527 github.com/goccy/spidermonkeywasm2go/p7.Fn8527
func Fn8527(m *base.Module, l0 int32)

//go:linkname Fn8530 github.com/goccy/spidermonkeywasm2go/p5.Fn8530
func Fn8530(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8532 github.com/goccy/spidermonkeywasm2go/p6.Fn8532
func Fn8532(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8533 github.com/goccy/spidermonkeywasm2go/p6.Fn8533
func Fn8533(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8534 github.com/goccy/spidermonkeywasm2go/p3.Fn8534
func Fn8534(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8535 github.com/goccy/spidermonkeywasm2go/p5.Fn8535
func Fn8535(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8537 github.com/goccy/spidermonkeywasm2go/p7.Fn8537
func Fn8537(m *base.Module, l0 int32) int32

//go:linkname Fn8539 github.com/goccy/spidermonkeywasm2go/p7.Fn8539
func Fn8539(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8555 github.com/goccy/spidermonkeywasm2go/p6.Fn8555
func Fn8555(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8556 github.com/goccy/spidermonkeywasm2go/p6.Fn8556
func Fn8556(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8559 github.com/goccy/spidermonkeywasm2go/p7.Fn8559
func Fn8559(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8561 github.com/goccy/spidermonkeywasm2go/p7.Fn8561
func Fn8561(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8582 github.com/goccy/spidermonkeywasm2go/p7.Fn8582
func Fn8582(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8628 github.com/goccy/spidermonkeywasm2go/p7.Fn8628
func Fn8628(m *base.Module, l0 int32)

//go:linkname Fn8629 github.com/goccy/spidermonkeywasm2go/p7.Fn8629
func Fn8629(m *base.Module, l0 int32)

//go:linkname Fn8630 github.com/goccy/spidermonkeywasm2go/p6.Fn8630
func Fn8630(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8631 github.com/goccy/spidermonkeywasm2go/p7.Fn8631
func Fn8631(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8638 github.com/goccy/spidermonkeywasm2go/p6.Fn8638
func Fn8638(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8640 github.com/goccy/spidermonkeywasm2go/p6.Fn8640
func Fn8640(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8642 github.com/goccy/spidermonkeywasm2go/p6.Fn8642
func Fn8642(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn8643 github.com/goccy/spidermonkeywasm2go/p6.Fn8643
func Fn8643(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8645 github.com/goccy/spidermonkeywasm2go/p6.Fn8645
func Fn8645(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8657 github.com/goccy/spidermonkeywasm2go/p2.Fn8657
func Fn8657(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn8661 github.com/goccy/spidermonkeywasm2go/p6.Fn8661
func Fn8661(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8662 github.com/goccy/spidermonkeywasm2go/p6.Fn8662
func Fn8662(m *base.Module, l0 int32)

//go:linkname Fn8683 github.com/goccy/spidermonkeywasm2go/p5.Fn8683
func Fn8683(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8685 github.com/goccy/spidermonkeywasm2go/p6.Fn8685
func Fn8685(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8687 github.com/goccy/spidermonkeywasm2go/p6.Fn8687
func Fn8687(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8694 github.com/goccy/spidermonkeywasm2go/p4.Fn8694
func Fn8694(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8708 github.com/goccy/spidermonkeywasm2go/p4.Fn8708
func Fn8708(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8720 github.com/goccy/spidermonkeywasm2go/p5.Fn8720
func Fn8720(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8724 github.com/goccy/spidermonkeywasm2go/p7.Fn8724
func Fn8724(m *base.Module, l0 int32) int32

//go:linkname Fn8725 github.com/goccy/spidermonkeywasm2go/p7.Fn8725
func Fn8725(m *base.Module, l0 int32)

//go:linkname Fn8728 github.com/goccy/spidermonkeywasm2go/p6.Fn8728
func Fn8728(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8729 github.com/goccy/spidermonkeywasm2go/p6.Fn8729
func Fn8729(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8730 github.com/goccy/spidermonkeywasm2go/p6.Fn8730
func Fn8730(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8758 github.com/goccy/spidermonkeywasm2go/p7.Fn8758
func Fn8758(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9009 github.com/goccy/spidermonkeywasm2go/p4.Fn9009
func Fn9009(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn9010 github.com/goccy/spidermonkeywasm2go/p6.Fn9010
func Fn9010(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9030 github.com/goccy/spidermonkeywasm2go/p5.Fn9030
func Fn9030(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9031 github.com/goccy/spidermonkeywasm2go/p7.Fn9031
func Fn9031(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9063 github.com/goccy/spidermonkeywasm2go/p6.Fn9063
func Fn9063(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9068 github.com/goccy/spidermonkeywasm2go/p3.Fn9068
func Fn9068(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9080 github.com/goccy/spidermonkeywasm2go/p4.Fn9080
func Fn9080(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9103 github.com/goccy/spidermonkeywasm2go/p2.Fn9103
func Fn9103(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9105 github.com/goccy/spidermonkeywasm2go/p4.Fn9105
func Fn9105(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9109 github.com/goccy/spidermonkeywasm2go/p4.Fn9109
func Fn9109(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9143 github.com/goccy/spidermonkeywasm2go/p6.Fn9143
func Fn9143(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9159 github.com/goccy/spidermonkeywasm2go/p5.Fn9159
func Fn9159(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9277 github.com/goccy/spidermonkeywasm2go/p3.Fn9277
func Fn9277(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9300 github.com/goccy/spidermonkeywasm2go/p6.Fn9300
func Fn9300(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9301 github.com/goccy/spidermonkeywasm2go/p6.Fn9301
func Fn9301(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9302 github.com/goccy/spidermonkeywasm2go/p7.Fn9302
func Fn9302(m *base.Module, l0 int32) int32

//go:linkname Fn9303 github.com/goccy/spidermonkeywasm2go/p6.Fn9303
func Fn9303(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9304 github.com/goccy/spidermonkeywasm2go/p7.Fn9304
func Fn9304(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9308 github.com/goccy/spidermonkeywasm2go/p5.Fn9308
func Fn9308(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9340 github.com/goccy/spidermonkeywasm2go/p3.Fn9340
func Fn9340(m *base.Module, l0 int32) int32

//go:linkname Fn9351 github.com/goccy/spidermonkeywasm2go/p7.Fn9351
func Fn9351(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9354 github.com/goccy/spidermonkeywasm2go/p6.Fn9354
func Fn9354(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9358 github.com/goccy/spidermonkeywasm2go/p3.Fn9358
func Fn9358(m *base.Module, l0 int32)

//go:linkname Fn9359 github.com/goccy/spidermonkeywasm2go/p4.Fn9359
func Fn9359(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9360 github.com/goccy/spidermonkeywasm2go/p5.Fn9360
func Fn9360(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9365 github.com/goccy/spidermonkeywasm2go/p5.Fn9365
func Fn9365(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9368 github.com/goccy/spidermonkeywasm2go/p6.Fn9368
func Fn9368(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9371 github.com/goccy/spidermonkeywasm2go/p6.Fn9371
func Fn9371(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9372 github.com/goccy/spidermonkeywasm2go/p5.Fn9372
func Fn9372(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9373 github.com/goccy/spidermonkeywasm2go/p5.Fn9373
func Fn9373(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9374 github.com/goccy/spidermonkeywasm2go/p6.Fn9374
func Fn9374(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9376 github.com/goccy/spidermonkeywasm2go/p5.Fn9376
func Fn9376(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9377 github.com/goccy/spidermonkeywasm2go/p5.Fn9377
func Fn9377(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9430 github.com/goccy/spidermonkeywasm2go/p5.Fn9430
func Fn9430(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9431 github.com/goccy/spidermonkeywasm2go/p2.Fn9431
func Fn9431(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn9434 github.com/goccy/spidermonkeywasm2go/p5.Fn9434
func Fn9434(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9440 github.com/goccy/spidermonkeywasm2go/p5.Fn9440
func Fn9440(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9447 github.com/goccy/spidermonkeywasm2go/p4.Fn9447
func Fn9447(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn9490 github.com/goccy/spidermonkeywasm2go/p4.Fn9490
func Fn9490(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9491 github.com/goccy/spidermonkeywasm2go/p3.Fn9491
func Fn9491(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9524 github.com/goccy/spidermonkeywasm2go/p6.Fn9524
func Fn9524(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9527 github.com/goccy/spidermonkeywasm2go/p6.Fn9527
func Fn9527(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9533 github.com/goccy/spidermonkeywasm2go/p6.Fn9533
func Fn9533(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9535 github.com/goccy/spidermonkeywasm2go/p6.Fn9535
func Fn9535(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9565 github.com/goccy/spidermonkeywasm2go/p7.Fn9565
func Fn9565(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9566 github.com/goccy/spidermonkeywasm2go/p7.Fn9566
func Fn9566(m *base.Module, l0 int32)

//go:linkname Fn9567 github.com/goccy/spidermonkeywasm2go/p6.Fn9567
func Fn9567(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn9568 github.com/goccy/spidermonkeywasm2go/p6.Fn9568
func Fn9568(m *base.Module, l0 int32, l1 float64, l2 int32)

//go:linkname Fn9581 github.com/goccy/spidermonkeywasm2go/p5.Fn9581
func Fn9581(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn9583 github.com/goccy/spidermonkeywasm2go/p4.Fn9583
func Fn9583(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9597 github.com/goccy/spidermonkeywasm2go/p5.Fn9597
func Fn9597(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9610 github.com/goccy/spidermonkeywasm2go/p6.Fn9610
func Fn9610(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9629 github.com/goccy/spidermonkeywasm2go/p5.Fn9629
func Fn9629(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn9635 github.com/goccy/spidermonkeywasm2go/p4.Fn9635
func Fn9635(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9636 github.com/goccy/spidermonkeywasm2go/p5.Fn9636
func Fn9636(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9650 github.com/goccy/spidermonkeywasm2go/p2.Fn9650
func Fn9650(m *base.Module, l0 int32)

//go:linkname Fn9651 github.com/goccy/spidermonkeywasm2go/p3.Fn9651
func Fn9651(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9655 github.com/goccy/spidermonkeywasm2go/p7.Fn9655
func Fn9655(m *base.Module, l0 int32) int32

//go:linkname Fn9657 github.com/goccy/spidermonkeywasm2go/p6.Fn9657
func Fn9657(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9677 github.com/goccy/spidermonkeywasm2go/p7.Fn9677
func Fn9677(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9680 github.com/goccy/spidermonkeywasm2go/p6.Fn9680
func Fn9680(m *base.Module, l0 int32)

//go:linkname Fn9681 github.com/goccy/spidermonkeywasm2go/p6.Fn9681
func Fn9681(m *base.Module, l0 int32)

//go:linkname Fn9682 github.com/goccy/spidermonkeywasm2go/p6.Fn9682
func Fn9682(m *base.Module, l0 int32)

//go:linkname Fn9689 github.com/goccy/spidermonkeywasm2go/p6.Fn9689
func Fn9689(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9691 github.com/goccy/spidermonkeywasm2go/p6.Fn9691
func Fn9691(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9701 github.com/goccy/spidermonkeywasm2go/p5.Fn9701
func Fn9701(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9713 github.com/goccy/spidermonkeywasm2go/p4.Fn9713
func Fn9713(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9714 github.com/goccy/spidermonkeywasm2go/p4.Fn9714
func Fn9714(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9771 github.com/goccy/spidermonkeywasm2go/p7.Fn9771
func Fn9771(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9776 github.com/goccy/spidermonkeywasm2go/p7.Fn9776
func Fn9776(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10121 github.com/goccy/spidermonkeywasm2go/p6.Fn10121
func Fn10121(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10158 github.com/goccy/spidermonkeywasm2go/p5.Fn10158
func Fn10158(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10184 github.com/goccy/spidermonkeywasm2go/p5.Fn10184
func Fn10184(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10187 github.com/goccy/spidermonkeywasm2go/p7.Fn10187
func Fn10187(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10188 github.com/goccy/spidermonkeywasm2go/p6.Fn10188
func Fn10188(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10189 github.com/goccy/spidermonkeywasm2go/p4.Fn10189
func Fn10189(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10202 github.com/goccy/spidermonkeywasm2go/p3.Fn10202
func Fn10202(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10203 github.com/goccy/spidermonkeywasm2go/p5.Fn10203
func Fn10203(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10211 github.com/goccy/spidermonkeywasm2go/p3.Fn10211
func Fn10211(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10227 github.com/goccy/spidermonkeywasm2go/p6.Fn10227
func Fn10227(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10229 github.com/goccy/spidermonkeywasm2go/p4.Fn10229
func Fn10229(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10232 github.com/goccy/spidermonkeywasm2go/p5.Fn10232
func Fn10232(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10233 github.com/goccy/spidermonkeywasm2go/p5.Fn10233
func Fn10233(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10251 github.com/goccy/spidermonkeywasm2go/p6.Fn10251
func Fn10251(m *base.Module, l0 int32)

//go:linkname Fn10252 github.com/goccy/spidermonkeywasm2go/p5.Fn10252
func Fn10252(m *base.Module, l0 int32) int32

//go:linkname Fn10253 github.com/goccy/spidermonkeywasm2go/p5.Fn10253
func Fn10253(m *base.Module, l0 int32)

//go:linkname Fn10254 github.com/goccy/spidermonkeywasm2go/p5.Fn10254
func Fn10254(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10255 github.com/goccy/spidermonkeywasm2go/p6.Fn10255
func Fn10255(m *base.Module, l0 int32) int32

//go:linkname Fn10256 github.com/goccy/spidermonkeywasm2go/p3.Fn10256
func Fn10256(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10268 github.com/goccy/spidermonkeywasm2go/p6.Fn10268
func Fn10268(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10272 github.com/goccy/spidermonkeywasm2go/p7.Fn10272
func Fn10272(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10277 github.com/goccy/spidermonkeywasm2go/p7.Fn10277
func Fn10277(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10280 github.com/goccy/spidermonkeywasm2go/p7.Fn10280
func Fn10280(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10283 github.com/goccy/spidermonkeywasm2go/p7.Fn10283
func Fn10283(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10285 github.com/goccy/spidermonkeywasm2go/p7.Fn10285
func Fn10285(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10294 github.com/goccy/spidermonkeywasm2go/p5.Fn10294
func Fn10294(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn10296 github.com/goccy/spidermonkeywasm2go/p6.Fn10296
func Fn10296(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10305 github.com/goccy/spidermonkeywasm2go/p3.Fn10305
func Fn10305(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10316 github.com/goccy/spidermonkeywasm2go/p3.Fn10316
func Fn10316(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10347 github.com/goccy/spidermonkeywasm2go/p6.Fn10347
func Fn10347(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10348 github.com/goccy/spidermonkeywasm2go/p5.Fn10348
func Fn10348(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10349 github.com/goccy/spidermonkeywasm2go/p4.Fn10349
func Fn10349(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn10350 github.com/goccy/spidermonkeywasm2go/p6.Fn10350
func Fn10350(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10351 github.com/goccy/spidermonkeywasm2go/p5.Fn10351
func Fn10351(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10352 github.com/goccy/spidermonkeywasm2go/p6.Fn10352
func Fn10352(m *base.Module, l0 int32) int32

//go:linkname Fn10562 github.com/goccy/spidermonkeywasm2go/p6.Fn10562
func Fn10562(m *base.Module, l0 int32) int32

//go:linkname Fn10581 github.com/goccy/spidermonkeywasm2go/p7.Fn10581
func Fn10581(m *base.Module, l0 int32) int32

//go:linkname Fn10795 github.com/goccy/spidermonkeywasm2go/p6.Fn10795
func Fn10795(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10798 github.com/goccy/spidermonkeywasm2go/p6.Fn10798
func Fn10798(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10799 github.com/goccy/spidermonkeywasm2go/p6.Fn10799
func Fn10799(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn10821 github.com/goccy/spidermonkeywasm2go/p7.Fn10821
func Fn10821(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10826 github.com/goccy/spidermonkeywasm2go/p5.Fn10826
func Fn10826(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10831 github.com/goccy/spidermonkeywasm2go/p6.Fn10831
func Fn10831(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10835 github.com/goccy/spidermonkeywasm2go/p6.Fn10835
func Fn10835(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10836 github.com/goccy/spidermonkeywasm2go/p5.Fn10836
func Fn10836(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn10941 github.com/goccy/spidermonkeywasm2go/p7.Fn10941
func Fn10941(m *base.Module)

//go:linkname Fn10943 github.com/goccy/spidermonkeywasm2go/p2.Fn10943
func Fn10943(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10975 github.com/goccy/spidermonkeywasm2go/p6.Fn10975
func Fn10975(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10978 github.com/goccy/spidermonkeywasm2go/p7.Fn10978
func Fn10978(m *base.Module, l0 int32)
