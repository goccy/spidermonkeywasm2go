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

//go:linkname Fn105 github.com/goccy/spidermonkeywasm2go/p7.Fn105
func Fn105(m *base.Module, l0 int32) int32

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

//go:linkname Fn348 github.com/goccy/spidermonkeywasm2go/p7.Fn348
func Fn348(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)

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

//go:linkname Fn695 github.com/goccy/spidermonkeywasm2go/p7.Fn695
func Fn695(m *base.Module, l0 int32) int32

//go:linkname Fn705 github.com/goccy/spidermonkeywasm2go/p7.Fn705
func Fn705(m *base.Module, l0 int32)

//go:linkname Fn706 github.com/goccy/spidermonkeywasm2go/p7.Fn706
func Fn706(m *base.Module, l0 int32) int32

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

//go:linkname Fn761 github.com/goccy/spidermonkeywasm2go/p7.Fn761
func Fn761(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn770 github.com/goccy/spidermonkeywasm2go/p7.Fn770
func Fn770(m *base.Module, l0 int32) int32

//go:linkname Fn780 github.com/goccy/spidermonkeywasm2go/p7.Fn780
func Fn780(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn781 github.com/goccy/spidermonkeywasm2go/p6.Fn781
func Fn781(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn782 github.com/goccy/spidermonkeywasm2go/p7.Fn782
func Fn782(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn783 github.com/goccy/spidermonkeywasm2go/p7.Fn783
func Fn783(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn784 github.com/goccy/spidermonkeywasm2go/p7.Fn784
func Fn784(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn786 github.com/goccy/spidermonkeywasm2go/p7.Fn786
func Fn786(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn787 github.com/goccy/spidermonkeywasm2go/p6.Fn787
func Fn787(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn789 github.com/goccy/spidermonkeywasm2go/p7.Fn789
func Fn789(m *base.Module, l0 int32) int32

//go:linkname Fn790 github.com/goccy/spidermonkeywasm2go/p7.Fn790
func Fn790(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn791 github.com/goccy/spidermonkeywasm2go/p6.Fn791
func Fn791(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn793 github.com/goccy/spidermonkeywasm2go/p7.Fn793
func Fn793(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn794 github.com/goccy/spidermonkeywasm2go/p4.Fn794
func Fn794(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn795 github.com/goccy/spidermonkeywasm2go/p6.Fn795
func Fn795(m *base.Module, l0 int32) int32

//go:linkname Fn800 github.com/goccy/spidermonkeywasm2go/p6.Fn800
func Fn800(m *base.Module, l0 int32)

//go:linkname Fn801 github.com/goccy/spidermonkeywasm2go/p7.Fn801
func Fn801(m *base.Module, l0 int32)

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

//go:linkname Fn1290 github.com/goccy/spidermonkeywasm2go/p6.Fn1290
func Fn1290(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1305 github.com/goccy/spidermonkeywasm2go/p6.Fn1305
func Fn1305(m *base.Module, l0 int32, l1 int32, l2 int32)

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

//go:linkname Fn1474 github.com/goccy/spidermonkeywasm2go/p7.Fn1474
func Fn1474(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1475 github.com/goccy/spidermonkeywasm2go/p5.Fn1475
func Fn1475(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1480 github.com/goccy/spidermonkeywasm2go/p0.Fn1480
func Fn1480(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1481 github.com/goccy/spidermonkeywasm2go/p0.Fn1481
func Fn1481(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1483 github.com/goccy/spidermonkeywasm2go/p7.Fn1483
func Fn1483(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1484 github.com/goccy/spidermonkeywasm2go/p7.Fn1484
func Fn1484(m *base.Module, l0 int32)

//go:linkname Fn1486 github.com/goccy/spidermonkeywasm2go/p7.Fn1486
func Fn1486(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1487 github.com/goccy/spidermonkeywasm2go/p7.Fn1487
func Fn1487(m *base.Module, l0 int32)

//go:linkname Fn1494 github.com/goccy/spidermonkeywasm2go/p7.Fn1494
func Fn1494(m *base.Module, l0 int32)

//go:linkname Fn1495 github.com/goccy/spidermonkeywasm2go/p7.Fn1495
func Fn1495(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1499 github.com/goccy/spidermonkeywasm2go/p6.Fn1499
func Fn1499(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1500 github.com/goccy/spidermonkeywasm2go/p7.Fn1500
func Fn1500(m *base.Module, l0 int32)

//go:linkname Fn1501 github.com/goccy/spidermonkeywasm2go/p6.Fn1501
func Fn1501(m *base.Module, l0 int32) int32

//go:linkname Fn1504 github.com/goccy/spidermonkeywasm2go/p7.Fn1504
func Fn1504(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1505 github.com/goccy/spidermonkeywasm2go/p7.Fn1505
func Fn1505(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1512 github.com/goccy/spidermonkeywasm2go/p7.Fn1512
func Fn1512(m *base.Module, l0 int32)

//go:linkname Fn1517 github.com/goccy/spidermonkeywasm2go/p7.Fn1517
func Fn1517(m *base.Module, l0 int32) int32

//go:linkname Fn1518 github.com/goccy/spidermonkeywasm2go/p5.Fn1518
func Fn1518(m *base.Module, l0 float64, l1 float64, l2 float64) float64

//go:linkname Fn1520 github.com/goccy/spidermonkeywasm2go/p5.Fn1520
func Fn1520(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn1524 github.com/goccy/spidermonkeywasm2go/p6.Fn1524
func Fn1524(m *base.Module, l0 int32, l1 float64) float64

//go:linkname Fn1534 github.com/goccy/spidermonkeywasm2go/p7.Fn1534
func Fn1534(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1535 github.com/goccy/spidermonkeywasm2go/p7.Fn1535
func Fn1535(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1536 github.com/goccy/spidermonkeywasm2go/p6.Fn1536
func Fn1536(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1537 github.com/goccy/spidermonkeywasm2go/p6.Fn1537
func Fn1537(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1628 github.com/goccy/spidermonkeywasm2go/p0.Fn1628
func Fn1628(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1635 github.com/goccy/spidermonkeywasm2go/p7.Fn1635
func Fn1635(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1651 github.com/goccy/spidermonkeywasm2go/p7.Fn1651
func Fn1651(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1661 github.com/goccy/spidermonkeywasm2go/p6.Fn1661
func Fn1661(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn1663 github.com/goccy/spidermonkeywasm2go/p5.Fn1663
func Fn1663(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn1665 github.com/goccy/spidermonkeywasm2go/p7.Fn1665
func Fn1665(m *base.Module, l0 int32) float64

//go:linkname Fn1666 github.com/goccy/spidermonkeywasm2go/p7.Fn1666
func Fn1666(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1668 github.com/goccy/spidermonkeywasm2go/p0.Fn1668
func Fn1668(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1671 github.com/goccy/spidermonkeywasm2go/p0.Fn1671
func Fn1671(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1768 github.com/goccy/spidermonkeywasm2go/p7.Fn1768
func Fn1768(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1770 github.com/goccy/spidermonkeywasm2go/p7.Fn1770
func Fn1770(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1813 github.com/goccy/spidermonkeywasm2go/p7.Fn1813
func Fn1813(m *base.Module, l0 int32) int32

//go:linkname Fn1814 github.com/goccy/spidermonkeywasm2go/p6.Fn1814
func Fn1814(m *base.Module, l0 int32) int32

//go:linkname Fn1815 github.com/goccy/spidermonkeywasm2go/p6.Fn1815
func Fn1815(m *base.Module, l0 int32) int32

//go:linkname Fn1825 github.com/goccy/spidermonkeywasm2go/p5.Fn1825
func Fn1825(m *base.Module, l0 int32)

//go:linkname Fn1857 github.com/goccy/spidermonkeywasm2go/p0.Fn1857
func Fn1857(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1878 github.com/goccy/spidermonkeywasm2go/p3.Fn1878
func Fn1878(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1893 github.com/goccy/spidermonkeywasm2go/p7.Fn1893
func Fn1893(m *base.Module, l0 int32) int32

//go:linkname Fn1896 github.com/goccy/spidermonkeywasm2go/p7.Fn1896
func Fn1896(m *base.Module, l0 int32) int32

//go:linkname Fn1899 github.com/goccy/spidermonkeywasm2go/p7.Fn1899
func Fn1899(m *base.Module, l0 int32) int32

//go:linkname Fn1900 github.com/goccy/spidermonkeywasm2go/p7.Fn1900
func Fn1900(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1943 github.com/goccy/spidermonkeywasm2go/p0.Fn1943
func Fn1943(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1950 github.com/goccy/spidermonkeywasm2go/p0.Fn1950
func Fn1950(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1971 github.com/goccy/spidermonkeywasm2go/p0.Fn1971
func Fn1971(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1983 github.com/goccy/spidermonkeywasm2go/p0.Fn1983
func Fn1983(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1984 github.com/goccy/spidermonkeywasm2go/p0.Fn1984
func Fn1984(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1987 github.com/goccy/spidermonkeywasm2go/p0.Fn1987
func Fn1987(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1999 github.com/goccy/spidermonkeywasm2go/p0.Fn1999
func Fn1999(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2002 github.com/goccy/spidermonkeywasm2go/p6.Fn2002
func Fn2002(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn2003 github.com/goccy/spidermonkeywasm2go/p0.Fn2003
func Fn2003(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2004 github.com/goccy/spidermonkeywasm2go/p5.Fn2004
func Fn2004(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2005 github.com/goccy/spidermonkeywasm2go/p7.Fn2005
func Fn2005(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2050 github.com/goccy/spidermonkeywasm2go/p0.Fn2050
func Fn2050(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2068 github.com/goccy/spidermonkeywasm2go/p0.Fn2068
func Fn2068(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2081 github.com/goccy/spidermonkeywasm2go/p7.Fn2081
func Fn2081(m *base.Module, l0 int32) int32

//go:linkname Fn2083 github.com/goccy/spidermonkeywasm2go/p6.Fn2083
func Fn2083(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2088 github.com/goccy/spidermonkeywasm2go/p0.Fn2088
func Fn2088(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2092 github.com/goccy/spidermonkeywasm2go/p7.Fn2092
func Fn2092(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2115 github.com/goccy/spidermonkeywasm2go/p6.Fn2115
func Fn2115(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2136 github.com/goccy/spidermonkeywasm2go/p4.Fn2136
func Fn2136(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2148 github.com/goccy/spidermonkeywasm2go/p6.Fn2148
func Fn2148(m *base.Module, l0 int32)

//go:linkname Fn2207 github.com/goccy/spidermonkeywasm2go/p0.Fn2207
func Fn2207(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32

//go:linkname Fn2222 github.com/goccy/spidermonkeywasm2go/p5.Fn2222
func Fn2222(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2223 github.com/goccy/spidermonkeywasm2go/p4.Fn2223
func Fn2223(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2224 github.com/goccy/spidermonkeywasm2go/p5.Fn2224
func Fn2224(m *base.Module, l0 int32)

//go:linkname Fn2234 github.com/goccy/spidermonkeywasm2go/p6.Fn2234
func Fn2234(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2252 github.com/goccy/spidermonkeywasm2go/p6.Fn2252
func Fn2252(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2264 github.com/goccy/spidermonkeywasm2go/p6.Fn2264
func Fn2264(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2294 github.com/goccy/spidermonkeywasm2go/p6.Fn2294
func Fn2294(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2304 github.com/goccy/spidermonkeywasm2go/p6.Fn2304
func Fn2304(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2305 github.com/goccy/spidermonkeywasm2go/p5.Fn2305
func Fn2305(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2306 github.com/goccy/spidermonkeywasm2go/p7.Fn2306
func Fn2306(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2308 github.com/goccy/spidermonkeywasm2go/p0.Fn2308
func Fn2308(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2310 github.com/goccy/spidermonkeywasm2go/p0.Fn2310
func Fn2310(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn2313 github.com/goccy/spidermonkeywasm2go/p6.Fn2313
func Fn2313(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2323 github.com/goccy/spidermonkeywasm2go/p0.Fn2323
func Fn2323(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2325 github.com/goccy/spidermonkeywasm2go/p7.Fn2325
func Fn2325(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2328 github.com/goccy/spidermonkeywasm2go/p0.Fn2328
func Fn2328(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2329 github.com/goccy/spidermonkeywasm2go/p6.Fn2329
func Fn2329(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2334 github.com/goccy/spidermonkeywasm2go/p0.Fn2334
func Fn2334(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2352 github.com/goccy/spidermonkeywasm2go/p5.Fn2352
func Fn2352(m *base.Module, l0 int32)

//go:linkname Fn2353 github.com/goccy/spidermonkeywasm2go/p7.Fn2353
func Fn2353(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2404 github.com/goccy/spidermonkeywasm2go/p6.Fn2404
func Fn2404(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2405 github.com/goccy/spidermonkeywasm2go/p6.Fn2405
func Fn2405(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2409 github.com/goccy/spidermonkeywasm2go/p4.Fn2409
func Fn2409(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2457 github.com/goccy/spidermonkeywasm2go/p0.Fn2457
func Fn2457(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2475 github.com/goccy/spidermonkeywasm2go/p0.Fn2475
func Fn2475(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2478 github.com/goccy/spidermonkeywasm2go/p6.Fn2478
func Fn2478(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2484 github.com/goccy/spidermonkeywasm2go/p4.Fn2484
func Fn2484(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2486 github.com/goccy/spidermonkeywasm2go/p6.Fn2486
func Fn2486(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2494 github.com/goccy/spidermonkeywasm2go/p4.Fn2494
func Fn2494(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2507 github.com/goccy/spidermonkeywasm2go/p2.Fn2507
func Fn2507(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2513 github.com/goccy/spidermonkeywasm2go/p5.Fn2513
func Fn2513(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2533 github.com/goccy/spidermonkeywasm2go/p0.Fn2533
func Fn2533(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2537 github.com/goccy/spidermonkeywasm2go/p6.Fn2537
func Fn2537(m *base.Module, l0 int32)

//go:linkname Fn2540 github.com/goccy/spidermonkeywasm2go/p0.Fn2540
func Fn2540(m *base.Module, l0 int32) int32

//go:linkname Fn2541 github.com/goccy/spidermonkeywasm2go/p7.Fn2541
func Fn2541(m *base.Module, l0 int32)

//go:linkname Fn2546 github.com/goccy/spidermonkeywasm2go/p7.Fn2546
func Fn2546(m *base.Module, l0 int32)

//go:linkname Fn2549 github.com/goccy/spidermonkeywasm2go/p5.Fn2549
func Fn2549(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2551 github.com/goccy/spidermonkeywasm2go/p0.Fn2551
func Fn2551(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2553 github.com/goccy/spidermonkeywasm2go/p0.Fn2553
func Fn2553(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2556 github.com/goccy/spidermonkeywasm2go/p0.Fn2556
func Fn2556(m *base.Module, l0 int32)

//go:linkname Fn2561 github.com/goccy/spidermonkeywasm2go/p0.Fn2561
func Fn2561(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2562 github.com/goccy/spidermonkeywasm2go/p0.Fn2562
func Fn2562(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2569 github.com/goccy/spidermonkeywasm2go/p5.Fn2569
func Fn2569(m *base.Module, l0 int32)

//go:linkname Fn2570 github.com/goccy/spidermonkeywasm2go/p5.Fn2570
func Fn2570(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2576 github.com/goccy/spidermonkeywasm2go/p0.Fn2576
func Fn2576(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2577 github.com/goccy/spidermonkeywasm2go/p0.Fn2577
func Fn2577(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2580 github.com/goccy/spidermonkeywasm2go/p0.Fn2580
func Fn2580(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2582 github.com/goccy/spidermonkeywasm2go/p0.Fn2582
func Fn2582(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2583 github.com/goccy/spidermonkeywasm2go/p0.Fn2583
func Fn2583(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2585 github.com/goccy/spidermonkeywasm2go/p3.Fn2585
func Fn2585(m *base.Module, l0 int32)

//go:linkname Fn2587 github.com/goccy/spidermonkeywasm2go/p0.Fn2587
func Fn2587(m *base.Module, l0 int32)

//go:linkname Fn2588 github.com/goccy/spidermonkeywasm2go/p0.Fn2588
func Fn2588(m *base.Module, l0 int32)

//go:linkname Fn2593 github.com/goccy/spidermonkeywasm2go/p6.Fn2593
func Fn2593(m *base.Module, l0 int32) int32

//go:linkname Fn2599 github.com/goccy/spidermonkeywasm2go/p0.Fn2599
func Fn2599(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2602 github.com/goccy/spidermonkeywasm2go/p4.Fn2602
func Fn2602(m *base.Module, l0 int32)

//go:linkname Fn2609 github.com/goccy/spidermonkeywasm2go/p6.Fn2609
func Fn2609(m *base.Module, l0 int32) int32

//go:linkname Fn2610 github.com/goccy/spidermonkeywasm2go/p0.Fn2610
func Fn2610(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2619 github.com/goccy/spidermonkeywasm2go/p0.Fn2619
func Fn2619(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2621 github.com/goccy/spidermonkeywasm2go/p0.Fn2621
func Fn2621(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2641 github.com/goccy/spidermonkeywasm2go/p6.Fn2641
func Fn2641(m *base.Module, l0 int32)

//go:linkname Fn2642 github.com/goccy/spidermonkeywasm2go/p0.Fn2642
func Fn2642(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2643 github.com/goccy/spidermonkeywasm2go/p7.Fn2643
func Fn2643(m *base.Module, l0 int32)

//go:linkname Fn2644 github.com/goccy/spidermonkeywasm2go/p0.Fn2644
func Fn2644(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2645 github.com/goccy/spidermonkeywasm2go/p0.Fn2645
func Fn2645(m *base.Module, l0 int32) int32

//go:linkname Fn2646 github.com/goccy/spidermonkeywasm2go/p0.Fn2646
func Fn2646(m *base.Module, l0 int32) int32

//go:linkname Fn2649 github.com/goccy/spidermonkeywasm2go/p0.Fn2649
func Fn2649(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2650 github.com/goccy/spidermonkeywasm2go/p0.Fn2650
func Fn2650(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2651 github.com/goccy/spidermonkeywasm2go/p0.Fn2651
func Fn2651(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2654 github.com/goccy/spidermonkeywasm2go/p7.Fn2654
func Fn2654(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2655 github.com/goccy/spidermonkeywasm2go/p3.Fn2655
func Fn2655(m *base.Module, l0 int32) int32

//go:linkname Fn2656 github.com/goccy/spidermonkeywasm2go/p2.Fn2656
func Fn2656(m *base.Module, l0 int32) int32

//go:linkname Fn2657 github.com/goccy/spidermonkeywasm2go/p6.Fn2657
func Fn2657(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2658 github.com/goccy/spidermonkeywasm2go/p3.Fn2658
func Fn2658(m *base.Module, l0 int32) int32

//go:linkname Fn2659 github.com/goccy/spidermonkeywasm2go/p2.Fn2659
func Fn2659(m *base.Module, l0 int32) int32

//go:linkname Fn2719 github.com/goccy/spidermonkeywasm2go/p0.Fn2719
func Fn2719(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2726 github.com/goccy/spidermonkeywasm2go/p0.Fn2726
func Fn2726(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2731 github.com/goccy/spidermonkeywasm2go/p0.Fn2731
func Fn2731(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2732 github.com/goccy/spidermonkeywasm2go/p0.Fn2732
func Fn2732(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2741 github.com/goccy/spidermonkeywasm2go/p4.Fn2741
func Fn2741(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2742 github.com/goccy/spidermonkeywasm2go/p4.Fn2742
func Fn2742(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2751 github.com/goccy/spidermonkeywasm2go/p0.Fn2751
func Fn2751(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2753 github.com/goccy/spidermonkeywasm2go/p0.Fn2753
func Fn2753(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2754 github.com/goccy/spidermonkeywasm2go/p0.Fn2754
func Fn2754(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2763 github.com/goccy/spidermonkeywasm2go/p4.Fn2763
func Fn2763(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2764 github.com/goccy/spidermonkeywasm2go/p6.Fn2764
func Fn2764(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2765 github.com/goccy/spidermonkeywasm2go/p0.Fn2765
func Fn2765(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2773 github.com/goccy/spidermonkeywasm2go/p6.Fn2773
func Fn2773(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2774 github.com/goccy/spidermonkeywasm2go/p7.Fn2774
func Fn2774(m *base.Module, l0 int32)

//go:linkname Fn2779 github.com/goccy/spidermonkeywasm2go/p6.Fn2779
func Fn2779(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2843 github.com/goccy/spidermonkeywasm2go/p5.Fn2843
func Fn2843(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2861 github.com/goccy/spidermonkeywasm2go/p0.Fn2861
func Fn2861(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2865 github.com/goccy/spidermonkeywasm2go/p5.Fn2865
func Fn2865(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2866 github.com/goccy/spidermonkeywasm2go/p5.Fn2866
func Fn2866(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2880 github.com/goccy/spidermonkeywasm2go/p0.Fn2880
func Fn2880(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2886 github.com/goccy/spidermonkeywasm2go/p0.Fn2886
func Fn2886(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2913 github.com/goccy/spidermonkeywasm2go/p6.Fn2913
func Fn2913(m *base.Module, l0 int32) int32

//go:linkname Fn2939 github.com/goccy/spidermonkeywasm2go/p0.Fn2939
func Fn2939(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2941 github.com/goccy/spidermonkeywasm2go/p0.Fn2941
func Fn2941(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2942 github.com/goccy/spidermonkeywasm2go/p0.Fn2942
func Fn2942(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2944 github.com/goccy/spidermonkeywasm2go/p4.Fn2944
func Fn2944(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2945 github.com/goccy/spidermonkeywasm2go/p0.Fn2945
func Fn2945(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2957 github.com/goccy/spidermonkeywasm2go/p0.Fn2957
func Fn2957(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2971 github.com/goccy/spidermonkeywasm2go/p0.Fn2971
func Fn2971(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3003 github.com/goccy/spidermonkeywasm2go/p6.Fn3003
func Fn3003(m *base.Module, l0 int32)

//go:linkname Fn3005 github.com/goccy/spidermonkeywasm2go/p5.Fn3005
func Fn3005(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3009 github.com/goccy/spidermonkeywasm2go/p6.Fn3009
func Fn3009(m *base.Module, l0 int32) int32

//go:linkname Fn3044 github.com/goccy/spidermonkeywasm2go/p0.Fn3044
func Fn3044(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3045 github.com/goccy/spidermonkeywasm2go/p0.Fn3045
func Fn3045(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3046 github.com/goccy/spidermonkeywasm2go/p7.Fn3046
func Fn3046(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3050 github.com/goccy/spidermonkeywasm2go/p6.Fn3050
func Fn3050(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3159 github.com/goccy/spidermonkeywasm2go/p2.Fn3159
func Fn3159(m *base.Module, l0 int32) int32

//go:linkname Fn3161 github.com/goccy/spidermonkeywasm2go/p2.Fn3161
func Fn3161(m *base.Module, l0 int32)

//go:linkname Fn3162 github.com/goccy/spidermonkeywasm2go/p7.Fn3162
func Fn3162(m *base.Module, l0 int32)

//go:linkname Fn3163 github.com/goccy/spidermonkeywasm2go/p0.Fn3163
func Fn3163(m *base.Module, l0 int32) int32

//go:linkname Fn3166 github.com/goccy/spidermonkeywasm2go/p7.Fn3166
func Fn3166(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3169 github.com/goccy/spidermonkeywasm2go/p5.Fn3169
func Fn3169(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3176 github.com/goccy/spidermonkeywasm2go/p0.Fn3176
func Fn3176(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3189 github.com/goccy/spidermonkeywasm2go/p6.Fn3189
func Fn3189(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3190 github.com/goccy/spidermonkeywasm2go/p7.Fn3190
func Fn3190(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3191 github.com/goccy/spidermonkeywasm2go/p7.Fn3191
func Fn3191(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3193 github.com/goccy/spidermonkeywasm2go/p0.Fn3193
func Fn3193(m *base.Module, l0 int32) int32

//go:linkname Fn3204 github.com/goccy/spidermonkeywasm2go/p5.Fn3204
func Fn3204(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3215 github.com/goccy/spidermonkeywasm2go/p0.Fn3215
func Fn3215(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3217 github.com/goccy/spidermonkeywasm2go/p0.Fn3217
func Fn3217(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3219 github.com/goccy/spidermonkeywasm2go/p5.Fn3219
func Fn3219(m *base.Module, l0 int32)

//go:linkname Fn3222 github.com/goccy/spidermonkeywasm2go/p6.Fn3222
func Fn3222(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3223 github.com/goccy/spidermonkeywasm2go/p4.Fn3223
func Fn3223(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3238 github.com/goccy/spidermonkeywasm2go/p0.Fn3238
func Fn3238(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3249 github.com/goccy/spidermonkeywasm2go/p0.Fn3249
func Fn3249(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3266 github.com/goccy/spidermonkeywasm2go/p5.Fn3266
func Fn3266(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3467 github.com/goccy/spidermonkeywasm2go/p5.Fn3467
func Fn3467(m *base.Module, l0 int32) int32

//go:linkname Fn3468 github.com/goccy/spidermonkeywasm2go/p6.Fn3468
func Fn3468(m *base.Module, l0 float64) int32

//go:linkname Fn3469 github.com/goccy/spidermonkeywasm2go/p7.Fn3469
func Fn3469(m *base.Module, l0 float64) int32

//go:linkname Fn3470 github.com/goccy/spidermonkeywasm2go/p7.Fn3470
func Fn3470(m *base.Module, l0 float64) int32

//go:linkname Fn3471 github.com/goccy/spidermonkeywasm2go/p7.Fn3471
func Fn3471(m *base.Module, l0 float64) int32

//go:linkname Fn3472 github.com/goccy/spidermonkeywasm2go/p7.Fn3472
func Fn3472(m *base.Module, l0 float64) int32

//go:linkname Fn3490 github.com/goccy/spidermonkeywasm2go/p7.Fn3490
func Fn3490(m *base.Module, l0 int32)

//go:linkname Fn3502 github.com/goccy/spidermonkeywasm2go/p7.Fn3502
func Fn3502(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3506 github.com/goccy/spidermonkeywasm2go/p7.Fn3506
func Fn3506(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3508 github.com/goccy/spidermonkeywasm2go/p7.Fn3508
func Fn3508(m *base.Module, l0 int32) int32

//go:linkname Fn3533 github.com/goccy/spidermonkeywasm2go/p2.Fn3533
func Fn3533(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3534 github.com/goccy/spidermonkeywasm2go/p0.Fn3534
func Fn3534(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3536 github.com/goccy/spidermonkeywasm2go/p6.Fn3536
func Fn3536(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3539 github.com/goccy/spidermonkeywasm2go/p7.Fn3539
func Fn3539(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3540 github.com/goccy/spidermonkeywasm2go/p7.Fn3540
func Fn3540(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3541 github.com/goccy/spidermonkeywasm2go/p6.Fn3541
func Fn3541(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3542 github.com/goccy/spidermonkeywasm2go/p5.Fn3542
func Fn3542(m *base.Module, l0 int32) int32

//go:linkname Fn3543 github.com/goccy/spidermonkeywasm2go/p7.Fn3543
func Fn3543(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3546 github.com/goccy/spidermonkeywasm2go/p7.Fn3546
func Fn3546(m *base.Module, l0 int32) int32

//go:linkname Fn3547 github.com/goccy/spidermonkeywasm2go/p0.Fn3547
func Fn3547(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3550 github.com/goccy/spidermonkeywasm2go/p0.Fn3550
func Fn3550(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3551 github.com/goccy/spidermonkeywasm2go/p0.Fn3551
func Fn3551(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3561 github.com/goccy/spidermonkeywasm2go/p0.Fn3561
func Fn3561(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3573 github.com/goccy/spidermonkeywasm2go/p6.Fn3573
func Fn3573(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3575 github.com/goccy/spidermonkeywasm2go/p0.Fn3575
func Fn3575(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3579 github.com/goccy/spidermonkeywasm2go/p6.Fn3579
func Fn3579(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3583 github.com/goccy/spidermonkeywasm2go/p6.Fn3583
func Fn3583(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3585 github.com/goccy/spidermonkeywasm2go/p5.Fn3585
func Fn3585(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3586 github.com/goccy/spidermonkeywasm2go/p6.Fn3586
func Fn3586(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3587 github.com/goccy/spidermonkeywasm2go/p5.Fn3587
func Fn3587(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3590 github.com/goccy/spidermonkeywasm2go/p4.Fn3590
func Fn3590(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3593 github.com/goccy/spidermonkeywasm2go/p7.Fn3593
func Fn3593(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3594 github.com/goccy/spidermonkeywasm2go/p7.Fn3594
func Fn3594(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3595 github.com/goccy/spidermonkeywasm2go/p6.Fn3595
func Fn3595(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3596 github.com/goccy/spidermonkeywasm2go/p6.Fn3596
func Fn3596(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3597 github.com/goccy/spidermonkeywasm2go/p7.Fn3597
func Fn3597(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3598 github.com/goccy/spidermonkeywasm2go/p7.Fn3598
func Fn3598(m *base.Module, l0 int32, l1 float64) int32

//go:linkname Fn3603 github.com/goccy/spidermonkeywasm2go/p4.Fn3603
func Fn3603(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3604 github.com/goccy/spidermonkeywasm2go/p5.Fn3604
func Fn3604(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3605 github.com/goccy/spidermonkeywasm2go/p5.Fn3605
func Fn3605(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3606 github.com/goccy/spidermonkeywasm2go/p7.Fn3606
func Fn3606(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3607 github.com/goccy/spidermonkeywasm2go/p7.Fn3607
func Fn3607(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3608 github.com/goccy/spidermonkeywasm2go/p7.Fn3608
func Fn3608(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3609 github.com/goccy/spidermonkeywasm2go/p3.Fn3609
func Fn3609(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3613 github.com/goccy/spidermonkeywasm2go/p7.Fn3613
func Fn3613(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3614 github.com/goccy/spidermonkeywasm2go/p7.Fn3614
func Fn3614(m *base.Module, l0 int32) int32

//go:linkname Fn3653 github.com/goccy/spidermonkeywasm2go/p3.Fn3653
func Fn3653(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3662 github.com/goccy/spidermonkeywasm2go/p0.Fn3662
func Fn3662(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3663 github.com/goccy/spidermonkeywasm2go/p0.Fn3663
func Fn3663(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3695 github.com/goccy/spidermonkeywasm2go/p6.Fn3695
func Fn3695(m *base.Module, l0 int32)

//go:linkname Fn3720 github.com/goccy/spidermonkeywasm2go/p0.Fn3720
func Fn3720(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3733 github.com/goccy/spidermonkeywasm2go/p6.Fn3733
func Fn3733(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3747 github.com/goccy/spidermonkeywasm2go/p4.Fn3747
func Fn3747(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3751 github.com/goccy/spidermonkeywasm2go/p7.Fn3751
func Fn3751(m *base.Module, l0 float64) int32

//go:linkname Fn3765 github.com/goccy/spidermonkeywasm2go/p0.Fn3765
func Fn3765(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

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

//go:linkname Fn3777 github.com/goccy/spidermonkeywasm2go/p7.Fn3777
func Fn3777(m *base.Module, l0 int32, l1 int32) int32

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
func Fn3789(m *base.Module, l0 int32) int32

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

//go:linkname Fn3854 github.com/goccy/spidermonkeywasm2go/p0.Fn3854
func Fn3854(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3855 github.com/goccy/spidermonkeywasm2go/p2.Fn3855
func Fn3855(m *base.Module, l0 int32) int32

//go:linkname Fn3863 github.com/goccy/spidermonkeywasm2go/p5.Fn3863
func Fn3863(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3877 github.com/goccy/spidermonkeywasm2go/p5.Fn3877
func Fn3877(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3879 github.com/goccy/spidermonkeywasm2go/p5.Fn3879
func Fn3879(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3881 github.com/goccy/spidermonkeywasm2go/p5.Fn3881
func Fn3881(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3883 github.com/goccy/spidermonkeywasm2go/p5.Fn3883
func Fn3883(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3885 github.com/goccy/spidermonkeywasm2go/p5.Fn3885
func Fn3885(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3887 github.com/goccy/spidermonkeywasm2go/p5.Fn3887
func Fn3887(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3889 github.com/goccy/spidermonkeywasm2go/p5.Fn3889
func Fn3889(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3891 github.com/goccy/spidermonkeywasm2go/p5.Fn3891
func Fn3891(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3893 github.com/goccy/spidermonkeywasm2go/p5.Fn3893
func Fn3893(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3895 github.com/goccy/spidermonkeywasm2go/p5.Fn3895
func Fn3895(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3897 github.com/goccy/spidermonkeywasm2go/p5.Fn3897
func Fn3897(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3899 github.com/goccy/spidermonkeywasm2go/p5.Fn3899
func Fn3899(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3907 github.com/goccy/spidermonkeywasm2go/p0.Fn3907
func Fn3907(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3910 github.com/goccy/spidermonkeywasm2go/p0.Fn3910
func Fn3910(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3912 github.com/goccy/spidermonkeywasm2go/p0.Fn3912
func Fn3912(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3914 github.com/goccy/spidermonkeywasm2go/p0.Fn3914
func Fn3914(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3919 github.com/goccy/spidermonkeywasm2go/p0.Fn3919
func Fn3919(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3921 github.com/goccy/spidermonkeywasm2go/p0.Fn3921
func Fn3921(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3923 github.com/goccy/spidermonkeywasm2go/p0.Fn3923
func Fn3923(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3928 github.com/goccy/spidermonkeywasm2go/p0.Fn3928
func Fn3928(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3948 github.com/goccy/spidermonkeywasm2go/p7.Fn3948
func Fn3948(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3951 github.com/goccy/spidermonkeywasm2go/p7.Fn3951
func Fn3951(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3954 github.com/goccy/spidermonkeywasm2go/p7.Fn3954
func Fn3954(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3957 github.com/goccy/spidermonkeywasm2go/p3.Fn3957
func Fn3957(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3958 github.com/goccy/spidermonkeywasm2go/p5.Fn3958
func Fn3958(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3959 github.com/goccy/spidermonkeywasm2go/p3.Fn3959
func Fn3959(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3960 github.com/goccy/spidermonkeywasm2go/p5.Fn3960
func Fn3960(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3961 github.com/goccy/spidermonkeywasm2go/p3.Fn3961
func Fn3961(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3962 github.com/goccy/spidermonkeywasm2go/p5.Fn3962
func Fn3962(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3963 github.com/goccy/spidermonkeywasm2go/p7.Fn3963
func Fn3963(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3966 github.com/goccy/spidermonkeywasm2go/p2.Fn3966
func Fn3966(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3967 github.com/goccy/spidermonkeywasm2go/p5.Fn3967
func Fn3967(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4037 github.com/goccy/spidermonkeywasm2go/p4.Fn4037
func Fn4037(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4038 github.com/goccy/spidermonkeywasm2go/p4.Fn4038
func Fn4038(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4116 github.com/goccy/spidermonkeywasm2go/p6.Fn4116
func Fn4116(m *base.Module, l0 int32) int32

//go:linkname Fn4117 github.com/goccy/spidermonkeywasm2go/p2.Fn4117
func Fn4117(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4118 github.com/goccy/spidermonkeywasm2go/p3.Fn4118
func Fn4118(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4119 github.com/goccy/spidermonkeywasm2go/p2.Fn4119
func Fn4119(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4120 github.com/goccy/spidermonkeywasm2go/p2.Fn4120
func Fn4120(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4140 github.com/goccy/spidermonkeywasm2go/p0.Fn4140
func Fn4140(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4197 github.com/goccy/spidermonkeywasm2go/p7.Fn4197
func Fn4197(m *base.Module, l0 int32)

//go:linkname Fn4198 github.com/goccy/spidermonkeywasm2go/p7.Fn4198
func Fn4198(m *base.Module, l0 int32)

//go:linkname Fn4228 github.com/goccy/spidermonkeywasm2go/p7.Fn4228
func Fn4228(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4239 github.com/goccy/spidermonkeywasm2go/p6.Fn4239
func Fn4239(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4241 github.com/goccy/spidermonkeywasm2go/p2.Fn4241
func Fn4241(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4243 github.com/goccy/spidermonkeywasm2go/p5.Fn4243
func Fn4243(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4253 github.com/goccy/spidermonkeywasm2go/p3.Fn4253
func Fn4253(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4254 github.com/goccy/spidermonkeywasm2go/p4.Fn4254
func Fn4254(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32) int32

//go:linkname Fn4255 github.com/goccy/spidermonkeywasm2go/p5.Fn4255
func Fn4255(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32) int32

//go:linkname Fn4257 github.com/goccy/spidermonkeywasm2go/p4.Fn4257
func Fn4257(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4264 github.com/goccy/spidermonkeywasm2go/p7.Fn4264
func Fn4264(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn4265 github.com/goccy/spidermonkeywasm2go/p7.Fn4265
func Fn4265(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4271 github.com/goccy/spidermonkeywasm2go/p5.Fn4271
func Fn4271(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4272 github.com/goccy/spidermonkeywasm2go/p6.Fn4272
func Fn4272(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4273 github.com/goccy/spidermonkeywasm2go/p5.Fn4273
func Fn4273(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4274 github.com/goccy/spidermonkeywasm2go/p4.Fn4274
func Fn4274(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4275 github.com/goccy/spidermonkeywasm2go/p6.Fn4275
func Fn4275(m *base.Module, l0 int32) int32

//go:linkname Fn4276 github.com/goccy/spidermonkeywasm2go/p4.Fn4276
func Fn4276(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4277 github.com/goccy/spidermonkeywasm2go/p6.Fn4277
func Fn4277(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4299 github.com/goccy/spidermonkeywasm2go/p6.Fn4299
func Fn4299(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4300 github.com/goccy/spidermonkeywasm2go/p7.Fn4300
func Fn4300(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4420 github.com/goccy/spidermonkeywasm2go/p5.Fn4420
func Fn4420(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4544 github.com/goccy/spidermonkeywasm2go/p2.Fn4544
func Fn4544(m *base.Module, l0 int32, l1 float64, l2 float64, l3 int32, l4 int32) int32

//go:linkname Fn4545 github.com/goccy/spidermonkeywasm2go/p7.Fn4545
func Fn4545(m *base.Module, l0 int32) int32

//go:linkname Fn4550 github.com/goccy/spidermonkeywasm2go/p6.Fn4550
func Fn4550(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4552 github.com/goccy/spidermonkeywasm2go/p4.Fn4552
func Fn4552(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn4554 github.com/goccy/spidermonkeywasm2go/p3.Fn4554
func Fn4554(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4555 github.com/goccy/spidermonkeywasm2go/p6.Fn4555
func Fn4555(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4559 github.com/goccy/spidermonkeywasm2go/p2.Fn4559
func Fn4559(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4569 github.com/goccy/spidermonkeywasm2go/p4.Fn4569
func Fn4569(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4573 github.com/goccy/spidermonkeywasm2go/p5.Fn4573
func Fn4573(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4576 github.com/goccy/spidermonkeywasm2go/p6.Fn4576
func Fn4576(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4578 github.com/goccy/spidermonkeywasm2go/p6.Fn4578
func Fn4578(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4660 github.com/goccy/spidermonkeywasm2go/p5.Fn4660
func Fn4660(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4668 github.com/goccy/spidermonkeywasm2go/p4.Fn4668
func Fn4668(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4670 github.com/goccy/spidermonkeywasm2go/p6.Fn4670
func Fn4670(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4671 github.com/goccy/spidermonkeywasm2go/p3.Fn4671
func Fn4671(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4680 github.com/goccy/spidermonkeywasm2go/p6.Fn4680
func Fn4680(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4683 github.com/goccy/spidermonkeywasm2go/p0.Fn4683
func Fn4683(m *base.Module, l0 int32)

//go:linkname Fn4700 github.com/goccy/spidermonkeywasm2go/p0.Fn4700
func Fn4700(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4703 github.com/goccy/spidermonkeywasm2go/p6.Fn4703
func Fn4703(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4718 github.com/goccy/spidermonkeywasm2go/p4.Fn4718
func Fn4718(m *base.Module, l0 int32)

//go:linkname Fn4720 github.com/goccy/spidermonkeywasm2go/p7.Fn4720
func Fn4720(m *base.Module, l0 int32)

//go:linkname Fn4721 github.com/goccy/spidermonkeywasm2go/p6.Fn4721
func Fn4721(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4722 github.com/goccy/spidermonkeywasm2go/p7.Fn4722
func Fn4722(m *base.Module, l0 int32) int32

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

//go:linkname Fn4802 github.com/goccy/spidermonkeywasm2go/p5.Fn4802
func Fn4802(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4803 github.com/goccy/spidermonkeywasm2go/p6.Fn4803
func Fn4803(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4804 github.com/goccy/spidermonkeywasm2go/p6.Fn4804
func Fn4804(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4805 github.com/goccy/spidermonkeywasm2go/p6.Fn4805
func Fn4805(m *base.Module, l0 int32)

//go:linkname Fn4807 github.com/goccy/spidermonkeywasm2go/p0.Fn4807
func Fn4807(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4824 github.com/goccy/spidermonkeywasm2go/p7.Fn4824
func Fn4824(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4825 github.com/goccy/spidermonkeywasm2go/p6.Fn4825
func Fn4825(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4826 github.com/goccy/spidermonkeywasm2go/p0.Fn4826
func Fn4826(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4829 github.com/goccy/spidermonkeywasm2go/p7.Fn4829
func Fn4829(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4869 github.com/goccy/spidermonkeywasm2go/p0.Fn4869
func Fn4869(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4911 github.com/goccy/spidermonkeywasm2go/p5.Fn4911
func Fn4911(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4912 github.com/goccy/spidermonkeywasm2go/p5.Fn4912
func Fn4912(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn4915 github.com/goccy/spidermonkeywasm2go/p0.Fn4915
func Fn4915(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5136 github.com/goccy/spidermonkeywasm2go/p6.Fn5136
func Fn5136(m *base.Module, l0 int32) int32

//go:linkname Fn5149 github.com/goccy/spidermonkeywasm2go/p6.Fn5149
func Fn5149(m *base.Module, l0 int32) int32

//go:linkname Fn5150 github.com/goccy/spidermonkeywasm2go/p3.Fn5150
func Fn5150(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5151 github.com/goccy/spidermonkeywasm2go/p5.Fn5151
func Fn5151(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5152 github.com/goccy/spidermonkeywasm2go/p0.Fn5152
func Fn5152(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5154 github.com/goccy/spidermonkeywasm2go/p0.Fn5154
func Fn5154(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5155 github.com/goccy/spidermonkeywasm2go/p7.Fn5155
func Fn5155(m *base.Module, l0 int32)

//go:linkname Fn5285 github.com/goccy/spidermonkeywasm2go/p0.Fn5285
func Fn5285(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5287 github.com/goccy/spidermonkeywasm2go/p5.Fn5287
func Fn5287(m *base.Module, l0 int32)

//go:linkname Fn5288 github.com/goccy/spidermonkeywasm2go/p0.Fn5288
func Fn5288(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5328 github.com/goccy/spidermonkeywasm2go/p0.Fn5328
func Fn5328(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn5333 github.com/goccy/spidermonkeywasm2go/p7.Fn5333
func Fn5333(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5339 github.com/goccy/spidermonkeywasm2go/p6.Fn5339
func Fn5339(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn5340 github.com/goccy/spidermonkeywasm2go/p6.Fn5340
func Fn5340(m *base.Module, l0 int32) int32

//go:linkname Fn5349 github.com/goccy/spidermonkeywasm2go/p6.Fn5349
func Fn5349(m *base.Module, l0 int32) int32

//go:linkname Fn5353 github.com/goccy/spidermonkeywasm2go/p5.Fn5353
func Fn5353(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5354 github.com/goccy/spidermonkeywasm2go/p5.Fn5354
func Fn5354(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5360 github.com/goccy/spidermonkeywasm2go/p6.Fn5360
func Fn5360(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5364 github.com/goccy/spidermonkeywasm2go/p4.Fn5364
func Fn5364(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5368 github.com/goccy/spidermonkeywasm2go/p7.Fn5368
func Fn5368(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5369 github.com/goccy/spidermonkeywasm2go/p5.Fn5369
func Fn5369(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5388 github.com/goccy/spidermonkeywasm2go/p6.Fn5388
func Fn5388(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5403 github.com/goccy/spidermonkeywasm2go/p6.Fn5403
func Fn5403(m *base.Module, l0 int32) int32

//go:linkname Fn5404 github.com/goccy/spidermonkeywasm2go/p7.Fn5404
func Fn5404(m *base.Module, l0 int32) int32

//go:linkname Fn5405 github.com/goccy/spidermonkeywasm2go/p5.Fn5405
func Fn5405(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5406 github.com/goccy/spidermonkeywasm2go/p6.Fn5406
func Fn5406(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5407 github.com/goccy/spidermonkeywasm2go/p6.Fn5407
func Fn5407(m *base.Module, l0 int32) int32

//go:linkname Fn5408 github.com/goccy/spidermonkeywasm2go/p6.Fn5408
func Fn5408(m *base.Module, l0 int32) int32

//go:linkname Fn5409 github.com/goccy/spidermonkeywasm2go/p6.Fn5409
func Fn5409(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5410 github.com/goccy/spidermonkeywasm2go/p6.Fn5410
func Fn5410(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5411 github.com/goccy/spidermonkeywasm2go/p5.Fn5411
func Fn5411(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5412 github.com/goccy/spidermonkeywasm2go/p5.Fn5412
func Fn5412(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5414 github.com/goccy/spidermonkeywasm2go/p6.Fn5414
func Fn5414(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5417 github.com/goccy/spidermonkeywasm2go/p3.Fn5417
func Fn5417(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5422 github.com/goccy/spidermonkeywasm2go/p5.Fn5422
func Fn5422(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5430 github.com/goccy/spidermonkeywasm2go/p6.Fn5430
func Fn5430(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5439 github.com/goccy/spidermonkeywasm2go/p5.Fn5439
func Fn5439(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5509 github.com/goccy/spidermonkeywasm2go/p0.Fn5509
func Fn5509(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5550 github.com/goccy/spidermonkeywasm2go/p6.Fn5550
func Fn5550(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5551 github.com/goccy/spidermonkeywasm2go/p7.Fn5551
func Fn5551(m *base.Module, l0 int32) int32

//go:linkname Fn5569 github.com/goccy/spidermonkeywasm2go/p7.Fn5569
func Fn5569(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5571 github.com/goccy/spidermonkeywasm2go/p5.Fn5571
func Fn5571(m *base.Module, l0 int32) int32

//go:linkname Fn5575 github.com/goccy/spidermonkeywasm2go/p0.Fn5575
func Fn5575(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5576 github.com/goccy/spidermonkeywasm2go/p0.Fn5576
func Fn5576(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5577 github.com/goccy/spidermonkeywasm2go/p6.Fn5577
func Fn5577(m *base.Module, l0 int32)

//go:linkname Fn5582 github.com/goccy/spidermonkeywasm2go/p5.Fn5582
func Fn5582(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5604 github.com/goccy/spidermonkeywasm2go/p6.Fn5604
func Fn5604(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5608 github.com/goccy/spidermonkeywasm2go/p3.Fn5608
func Fn5608(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5621 github.com/goccy/spidermonkeywasm2go/p5.Fn5621
func Fn5621(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5622 github.com/goccy/spidermonkeywasm2go/p0.Fn5622
func Fn5622(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5653 github.com/goccy/spidermonkeywasm2go/p6.Fn5653
func Fn5653(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5655 github.com/goccy/spidermonkeywasm2go/p6.Fn5655
func Fn5655(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5656 github.com/goccy/spidermonkeywasm2go/p7.Fn5656
func Fn5656(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5657 github.com/goccy/spidermonkeywasm2go/p7.Fn5657
func Fn5657(m *base.Module, l0 int32)

//go:linkname Fn5659 github.com/goccy/spidermonkeywasm2go/p5.Fn5659
func Fn5659(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5663 github.com/goccy/spidermonkeywasm2go/p6.Fn5663
func Fn5663(m *base.Module, l0 int32)

//go:linkname Fn5667 github.com/goccy/spidermonkeywasm2go/p6.Fn5667
func Fn5667(m *base.Module, l0 int32)

//go:linkname Fn5670 github.com/goccy/spidermonkeywasm2go/p6.Fn5670
func Fn5670(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5672 github.com/goccy/spidermonkeywasm2go/p7.Fn5672
func Fn5672(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5691 github.com/goccy/spidermonkeywasm2go/p6.Fn5691
func Fn5691(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5694 github.com/goccy/spidermonkeywasm2go/p5.Fn5694
func Fn5694(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5706 github.com/goccy/spidermonkeywasm2go/p7.Fn5706
func Fn5706(m *base.Module, l0 int32) int32

//go:linkname Fn5738 github.com/goccy/spidermonkeywasm2go/p7.Fn5738
func Fn5738(m *base.Module, l0 int32)

//go:linkname Fn5751 github.com/goccy/spidermonkeywasm2go/p6.Fn5751
func Fn5751(m *base.Module, l0 int32)

//go:linkname Fn5754 github.com/goccy/spidermonkeywasm2go/p6.Fn5754
func Fn5754(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5821 github.com/goccy/spidermonkeywasm2go/p6.Fn5821
func Fn5821(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5822 github.com/goccy/spidermonkeywasm2go/p6.Fn5822
func Fn5822(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5823 github.com/goccy/spidermonkeywasm2go/p6.Fn5823
func Fn5823(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5824 github.com/goccy/spidermonkeywasm2go/p6.Fn5824
func Fn5824(m *base.Module, l0 int32)

//go:linkname Fn5838 github.com/goccy/spidermonkeywasm2go/p7.Fn5838
func Fn5838(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5839 github.com/goccy/spidermonkeywasm2go/p6.Fn5839
func Fn5839(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5840 github.com/goccy/spidermonkeywasm2go/p6.Fn5840
func Fn5840(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5848 github.com/goccy/spidermonkeywasm2go/p6.Fn5848
func Fn5848(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5851 github.com/goccy/spidermonkeywasm2go/p6.Fn5851
func Fn5851(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5852 github.com/goccy/spidermonkeywasm2go/p5.Fn5852
func Fn5852(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5853 github.com/goccy/spidermonkeywasm2go/p5.Fn5853
func Fn5853(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5854 github.com/goccy/spidermonkeywasm2go/p5.Fn5854
func Fn5854(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5855 github.com/goccy/spidermonkeywasm2go/p2.Fn5855
func Fn5855(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5856 github.com/goccy/spidermonkeywasm2go/p5.Fn5856
func Fn5856(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5857 github.com/goccy/spidermonkeywasm2go/p5.Fn5857
func Fn5857(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5858 github.com/goccy/spidermonkeywasm2go/p5.Fn5858
func Fn5858(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5859 github.com/goccy/spidermonkeywasm2go/p5.Fn5859
func Fn5859(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5860 github.com/goccy/spidermonkeywasm2go/p5.Fn5860
func Fn5860(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5861 github.com/goccy/spidermonkeywasm2go/p6.Fn5861
func Fn5861(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5862 github.com/goccy/spidermonkeywasm2go/p5.Fn5862
func Fn5862(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5863 github.com/goccy/spidermonkeywasm2go/p5.Fn5863
func Fn5863(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5864 github.com/goccy/spidermonkeywasm2go/p5.Fn5864
func Fn5864(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5866 github.com/goccy/spidermonkeywasm2go/p2.Fn5866
func Fn5866(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5867 github.com/goccy/spidermonkeywasm2go/p6.Fn5867
func Fn5867(m *base.Module, l0 int32)

//go:linkname Fn5876 github.com/goccy/spidermonkeywasm2go/p6.Fn5876
func Fn5876(m *base.Module, l0 int32)

//go:linkname Fn5883 github.com/goccy/spidermonkeywasm2go/p4.Fn5883
func Fn5883(m *base.Module, l0 int32)

//go:linkname Fn5886 github.com/goccy/spidermonkeywasm2go/p5.Fn5886
func Fn5886(m *base.Module, l0 int32) int32

//go:linkname Fn5943 github.com/goccy/spidermonkeywasm2go/p4.Fn5943
func Fn5943(m *base.Module, l0 int32)

//go:linkname Fn5944 github.com/goccy/spidermonkeywasm2go/p7.Fn5944
func Fn5944(m *base.Module, l0 int32)

//go:linkname Fn5954 github.com/goccy/spidermonkeywasm2go/p3.Fn5954
func Fn5954(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5955 github.com/goccy/spidermonkeywasm2go/p3.Fn5955
func Fn5955(m *base.Module, l0 int32) int32

//go:linkname Fn5956 github.com/goccy/spidermonkeywasm2go/p6.Fn5956
func Fn5956(m *base.Module, l0 int32) int32

//go:linkname Fn5959 github.com/goccy/spidermonkeywasm2go/p5.Fn5959
func Fn5959(m *base.Module, l0 int32)

//go:linkname Fn5960 github.com/goccy/spidermonkeywasm2go/p4.Fn5960
func Fn5960(m *base.Module, l0 int32)

//go:linkname Fn5969 github.com/goccy/spidermonkeywasm2go/p6.Fn5969
func Fn5969(m *base.Module, l0 int32)

//go:linkname Fn5972 github.com/goccy/spidermonkeywasm2go/p5.Fn5972
func Fn5972(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5978 github.com/goccy/spidermonkeywasm2go/p6.Fn5978
func Fn5978(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5980 github.com/goccy/spidermonkeywasm2go/p7.Fn5980
func Fn5980(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn6052 github.com/goccy/spidermonkeywasm2go/p6.Fn6052
func Fn6052(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6060 github.com/goccy/spidermonkeywasm2go/p7.Fn6060
func Fn6060(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6092 github.com/goccy/spidermonkeywasm2go/p7.Fn6092
func Fn6092(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6106 github.com/goccy/spidermonkeywasm2go/p7.Fn6106
func Fn6106(m *base.Module)

//go:linkname Fn6108 github.com/goccy/spidermonkeywasm2go/p6.Fn6108
func Fn6108(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6210 github.com/goccy/spidermonkeywasm2go/p6.Fn6210
func Fn6210(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6218 github.com/goccy/spidermonkeywasm2go/p7.Fn6218
func Fn6218(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6219 github.com/goccy/spidermonkeywasm2go/p6.Fn6219
func Fn6219(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6222 github.com/goccy/spidermonkeywasm2go/p6.Fn6222
func Fn6222(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6227 github.com/goccy/spidermonkeywasm2go/p6.Fn6227
func Fn6227(m *base.Module, l0 int32) int32

//go:linkname Fn6228 github.com/goccy/spidermonkeywasm2go/p5.Fn6228
func Fn6228(m *base.Module, l0 int32)

//go:linkname Fn6229 github.com/goccy/spidermonkeywasm2go/p5.Fn6229
func Fn6229(m *base.Module, l0 int32)

//go:linkname Fn6230 github.com/goccy/spidermonkeywasm2go/p6.Fn6230
func Fn6230(m *base.Module, l0 int32) int32

//go:linkname Fn6231 github.com/goccy/spidermonkeywasm2go/p6.Fn6231
func Fn6231(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6232 github.com/goccy/spidermonkeywasm2go/p2.Fn6232
func Fn6232(m *base.Module, l0 int32) int32

//go:linkname Fn6233 github.com/goccy/spidermonkeywasm2go/p4.Fn6233
func Fn6233(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6234 github.com/goccy/spidermonkeywasm2go/p3.Fn6234
func Fn6234(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn6235 github.com/goccy/spidermonkeywasm2go/p6.Fn6235
func Fn6235(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6236 github.com/goccy/spidermonkeywasm2go/p4.Fn6236
func Fn6236(m *base.Module, l0 int32) int32

//go:linkname Fn6237 github.com/goccy/spidermonkeywasm2go/p6.Fn6237
func Fn6237(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6238 github.com/goccy/spidermonkeywasm2go/p7.Fn6238
func Fn6238(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6239 github.com/goccy/spidermonkeywasm2go/p3.Fn6239
func Fn6239(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6240 github.com/goccy/spidermonkeywasm2go/p7.Fn6240
func Fn6240(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6241 github.com/goccy/spidermonkeywasm2go/p6.Fn6241
func Fn6241(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6242 github.com/goccy/spidermonkeywasm2go/p4.Fn6242
func Fn6242(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6243 github.com/goccy/spidermonkeywasm2go/p5.Fn6243
func Fn6243(m *base.Module, l0 int32)

//go:linkname Fn6244 github.com/goccy/spidermonkeywasm2go/p4.Fn6244
func Fn6244(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6245 github.com/goccy/spidermonkeywasm2go/p6.Fn6245
func Fn6245(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6246 github.com/goccy/spidermonkeywasm2go/p6.Fn6246
func Fn6246(m *base.Module, l0 int32) int32

//go:linkname Fn6247 github.com/goccy/spidermonkeywasm2go/p6.Fn6247
func Fn6247(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn6248 github.com/goccy/spidermonkeywasm2go/p5.Fn6248
func Fn6248(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6292 github.com/goccy/spidermonkeywasm2go/p5.Fn6292
func Fn6292(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6294 github.com/goccy/spidermonkeywasm2go/p3.Fn6294
func Fn6294(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6346 github.com/goccy/spidermonkeywasm2go/p2.Fn6346
func Fn6346(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6666 github.com/goccy/spidermonkeywasm2go/p7.Fn6666
func Fn6666(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6667 github.com/goccy/spidermonkeywasm2go/p6.Fn6667
func Fn6667(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6668 github.com/goccy/spidermonkeywasm2go/p6.Fn6668
func Fn6668(m *base.Module, l0 int32) int32

//go:linkname Fn6681 github.com/goccy/spidermonkeywasm2go/p6.Fn6681
func Fn6681(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6689 github.com/goccy/spidermonkeywasm2go/p7.Fn6689
func Fn6689(m *base.Module, l0 int32)

//go:linkname Fn6690 github.com/goccy/spidermonkeywasm2go/p7.Fn6690
func Fn6690(m *base.Module, l0 int32)

//go:linkname Fn6691 github.com/goccy/spidermonkeywasm2go/p6.Fn6691
func Fn6691(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn6693 github.com/goccy/spidermonkeywasm2go/p6.Fn6693
func Fn6693(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6694 github.com/goccy/spidermonkeywasm2go/p5.Fn6694
func Fn6694(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6695 github.com/goccy/spidermonkeywasm2go/p4.Fn6695
func Fn6695(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6696 github.com/goccy/spidermonkeywasm2go/p7.Fn6696
func Fn6696(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6697 github.com/goccy/spidermonkeywasm2go/p7.Fn6697
func Fn6697(m *base.Module, l0 int32)

//go:linkname Fn6698 github.com/goccy/spidermonkeywasm2go/p6.Fn6698
func Fn6698(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6699 github.com/goccy/spidermonkeywasm2go/p6.Fn6699
func Fn6699(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6700 github.com/goccy/spidermonkeywasm2go/p6.Fn6700
func Fn6700(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6701 github.com/goccy/spidermonkeywasm2go/p6.Fn6701
func Fn6701(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6702 github.com/goccy/spidermonkeywasm2go/p6.Fn6702
func Fn6702(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6703 github.com/goccy/spidermonkeywasm2go/p7.Fn6703
func Fn6703(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6704 github.com/goccy/spidermonkeywasm2go/p5.Fn6704
func Fn6704(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn6705 github.com/goccy/spidermonkeywasm2go/p6.Fn6705
func Fn6705(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6706 github.com/goccy/spidermonkeywasm2go/p4.Fn6706
func Fn6706(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6707 github.com/goccy/spidermonkeywasm2go/p6.Fn6707
func Fn6707(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn6708 github.com/goccy/spidermonkeywasm2go/p7.Fn6708
func Fn6708(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6709 github.com/goccy/spidermonkeywasm2go/p6.Fn6709
func Fn6709(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6710 github.com/goccy/spidermonkeywasm2go/p7.Fn6710
func Fn6710(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6711 github.com/goccy/spidermonkeywasm2go/p7.Fn6711
func Fn6711(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6712 github.com/goccy/spidermonkeywasm2go/p5.Fn6712
func Fn6712(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6713 github.com/goccy/spidermonkeywasm2go/p7.Fn6713
func Fn6713(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6714 github.com/goccy/spidermonkeywasm2go/p7.Fn6714
func Fn6714(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6715 github.com/goccy/spidermonkeywasm2go/p7.Fn6715
func Fn6715(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6716 github.com/goccy/spidermonkeywasm2go/p6.Fn6716
func Fn6716(m *base.Module, l0 int32) int32

//go:linkname Fn6717 github.com/goccy/spidermonkeywasm2go/p7.Fn6717
func Fn6717(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6718 github.com/goccy/spidermonkeywasm2go/p6.Fn6718
func Fn6718(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6719 github.com/goccy/spidermonkeywasm2go/p6.Fn6719
func Fn6719(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6720 github.com/goccy/spidermonkeywasm2go/p7.Fn6720
func Fn6720(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6722 github.com/goccy/spidermonkeywasm2go/p7.Fn6722
func Fn6722(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6724 github.com/goccy/spidermonkeywasm2go/p7.Fn6724
func Fn6724(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6725 github.com/goccy/spidermonkeywasm2go/p6.Fn6725
func Fn6725(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6727 github.com/goccy/spidermonkeywasm2go/p7.Fn6727
func Fn6727(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6734 github.com/goccy/spidermonkeywasm2go/p7.Fn6734
func Fn6734(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6735 github.com/goccy/spidermonkeywasm2go/p7.Fn6735
func Fn6735(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6736 github.com/goccy/spidermonkeywasm2go/p7.Fn6736
func Fn6736(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6738 github.com/goccy/spidermonkeywasm2go/p4.Fn6738
func Fn6738(m *base.Module, l0 int32) int32

//go:linkname Fn6739 github.com/goccy/spidermonkeywasm2go/p4.Fn6739
func Fn6739(m *base.Module, l0 int32) int32

//go:linkname Fn6740 github.com/goccy/spidermonkeywasm2go/p5.Fn6740
func Fn6740(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6741 github.com/goccy/spidermonkeywasm2go/p5.Fn6741
func Fn6741(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6742 github.com/goccy/spidermonkeywasm2go/p5.Fn6742
func Fn6742(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6743 github.com/goccy/spidermonkeywasm2go/p4.Fn6743
func Fn6743(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6744 github.com/goccy/spidermonkeywasm2go/p5.Fn6744
func Fn6744(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6745 github.com/goccy/spidermonkeywasm2go/p4.Fn6745
func Fn6745(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6746 github.com/goccy/spidermonkeywasm2go/p3.Fn6746
func Fn6746(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6747 github.com/goccy/spidermonkeywasm2go/p5.Fn6747
func Fn6747(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6748 github.com/goccy/spidermonkeywasm2go/p4.Fn6748
func Fn6748(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6749 github.com/goccy/spidermonkeywasm2go/p5.Fn6749
func Fn6749(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6750 github.com/goccy/spidermonkeywasm2go/p4.Fn6750
func Fn6750(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6751 github.com/goccy/spidermonkeywasm2go/p5.Fn6751
func Fn6751(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6752 github.com/goccy/spidermonkeywasm2go/p4.Fn6752
func Fn6752(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6753 github.com/goccy/spidermonkeywasm2go/p4.Fn6753
func Fn6753(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6754 github.com/goccy/spidermonkeywasm2go/p7.Fn6754
func Fn6754(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6755 github.com/goccy/spidermonkeywasm2go/p6.Fn6755
func Fn6755(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6756 github.com/goccy/spidermonkeywasm2go/p7.Fn6756
func Fn6756(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6757 github.com/goccy/spidermonkeywasm2go/p7.Fn6757
func Fn6757(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6758 github.com/goccy/spidermonkeywasm2go/p7.Fn6758
func Fn6758(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6759 github.com/goccy/spidermonkeywasm2go/p7.Fn6759
func Fn6759(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6763 github.com/goccy/spidermonkeywasm2go/p7.Fn6763
func Fn6763(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6766 github.com/goccy/spidermonkeywasm2go/p4.Fn6766
func Fn6766(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6767 github.com/goccy/spidermonkeywasm2go/p6.Fn6767
func Fn6767(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6768 github.com/goccy/spidermonkeywasm2go/p5.Fn6768
func Fn6768(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6769 github.com/goccy/spidermonkeywasm2go/p6.Fn6769
func Fn6769(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6770 github.com/goccy/spidermonkeywasm2go/p6.Fn6770
func Fn6770(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6771 github.com/goccy/spidermonkeywasm2go/p7.Fn6771
func Fn6771(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6772 github.com/goccy/spidermonkeywasm2go/p7.Fn6772
func Fn6772(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6773 github.com/goccy/spidermonkeywasm2go/p7.Fn6773
func Fn6773(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6774 github.com/goccy/spidermonkeywasm2go/p7.Fn6774
func Fn6774(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6778 github.com/goccy/spidermonkeywasm2go/p7.Fn6778
func Fn6778(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6779 github.com/goccy/spidermonkeywasm2go/p7.Fn6779
func Fn6779(m *base.Module, l0 int32) int32

//go:linkname Fn6780 github.com/goccy/spidermonkeywasm2go/p7.Fn6780
func Fn6780(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6781 github.com/goccy/spidermonkeywasm2go/p7.Fn6781
func Fn6781(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn6783 github.com/goccy/spidermonkeywasm2go/p7.Fn6783
func Fn6783(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6784 github.com/goccy/spidermonkeywasm2go/p7.Fn6784
func Fn6784(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6785 github.com/goccy/spidermonkeywasm2go/p4.Fn6785
func Fn6785(m *base.Module, l0 int32) int32

//go:linkname Fn6786 github.com/goccy/spidermonkeywasm2go/p6.Fn6786
func Fn6786(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6787 github.com/goccy/spidermonkeywasm2go/p7.Fn6787
func Fn6787(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6788 github.com/goccy/spidermonkeywasm2go/p7.Fn6788
func Fn6788(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6789 github.com/goccy/spidermonkeywasm2go/p7.Fn6789
func Fn6789(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6790 github.com/goccy/spidermonkeywasm2go/p7.Fn6790
func Fn6790(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6791 github.com/goccy/spidermonkeywasm2go/p7.Fn6791
func Fn6791(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6794 github.com/goccy/spidermonkeywasm2go/p5.Fn6794
func Fn6794(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6795 github.com/goccy/spidermonkeywasm2go/p5.Fn6795
func Fn6795(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6796 github.com/goccy/spidermonkeywasm2go/p6.Fn6796
func Fn6796(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6797 github.com/goccy/spidermonkeywasm2go/p6.Fn6797
func Fn6797(m *base.Module, l0 int32) int32

//go:linkname Fn6798 github.com/goccy/spidermonkeywasm2go/p6.Fn6798
func Fn6798(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6799 github.com/goccy/spidermonkeywasm2go/p7.Fn6799
func Fn6799(m *base.Module, l0 int32)

//go:linkname Fn6800 github.com/goccy/spidermonkeywasm2go/p7.Fn6800
func Fn6800(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6801 github.com/goccy/spidermonkeywasm2go/p6.Fn6801
func Fn6801(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6803 github.com/goccy/spidermonkeywasm2go/p7.Fn6803
func Fn6803(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6804 github.com/goccy/spidermonkeywasm2go/p7.Fn6804
func Fn6804(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6805 github.com/goccy/spidermonkeywasm2go/p7.Fn6805
func Fn6805(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6806 github.com/goccy/spidermonkeywasm2go/p7.Fn6806
func Fn6806(m *base.Module, l0 int32) int32

//go:linkname Fn6807 github.com/goccy/spidermonkeywasm2go/p3.Fn6807
func Fn6807(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6808 github.com/goccy/spidermonkeywasm2go/p6.Fn6808
func Fn6808(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6809 github.com/goccy/spidermonkeywasm2go/p7.Fn6809
func Fn6809(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6810 github.com/goccy/spidermonkeywasm2go/p4.Fn6810
func Fn6810(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6811 github.com/goccy/spidermonkeywasm2go/p6.Fn6811
func Fn6811(m *base.Module, l0 int32)

//go:linkname Fn6812 github.com/goccy/spidermonkeywasm2go/p7.Fn6812
func Fn6812(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6813 github.com/goccy/spidermonkeywasm2go/p7.Fn6813
func Fn6813(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6814 github.com/goccy/spidermonkeywasm2go/p7.Fn6814
func Fn6814(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6815 github.com/goccy/spidermonkeywasm2go/p7.Fn6815
func Fn6815(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6816 github.com/goccy/spidermonkeywasm2go/p6.Fn6816
func Fn6816(m *base.Module, l0 int32) int32

//go:linkname Fn6817 github.com/goccy/spidermonkeywasm2go/p6.Fn6817
func Fn6817(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6818 github.com/goccy/spidermonkeywasm2go/p7.Fn6818
func Fn6818(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6819 github.com/goccy/spidermonkeywasm2go/p7.Fn6819
func Fn6819(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6820 github.com/goccy/spidermonkeywasm2go/p7.Fn6820
func Fn6820(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6821 github.com/goccy/spidermonkeywasm2go/p7.Fn6821
func Fn6821(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6822 github.com/goccy/spidermonkeywasm2go/p7.Fn6822
func Fn6822(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6823 github.com/goccy/spidermonkeywasm2go/p7.Fn6823
func Fn6823(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6824 github.com/goccy/spidermonkeywasm2go/p7.Fn6824
func Fn6824(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6825 github.com/goccy/spidermonkeywasm2go/p7.Fn6825
func Fn6825(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6826 github.com/goccy/spidermonkeywasm2go/p7.Fn6826
func Fn6826(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6827 github.com/goccy/spidermonkeywasm2go/p7.Fn6827
func Fn6827(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6828 github.com/goccy/spidermonkeywasm2go/p6.Fn6828
func Fn6828(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6829 github.com/goccy/spidermonkeywasm2go/p7.Fn6829
func Fn6829(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6830 github.com/goccy/spidermonkeywasm2go/p7.Fn6830
func Fn6830(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6831 github.com/goccy/spidermonkeywasm2go/p7.Fn6831
func Fn6831(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6832 github.com/goccy/spidermonkeywasm2go/p6.Fn6832
func Fn6832(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6833 github.com/goccy/spidermonkeywasm2go/p7.Fn6833
func Fn6833(m *base.Module, l0 int32) int32

//go:linkname Fn6834 github.com/goccy/spidermonkeywasm2go/p7.Fn6834
func Fn6834(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6837 github.com/goccy/spidermonkeywasm2go/p5.Fn6837
func Fn6837(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6838 github.com/goccy/spidermonkeywasm2go/p3.Fn6838
func Fn6838(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6840 github.com/goccy/spidermonkeywasm2go/p6.Fn6840
func Fn6840(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6841 github.com/goccy/spidermonkeywasm2go/p5.Fn6841
func Fn6841(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6848 github.com/goccy/spidermonkeywasm2go/p5.Fn6848
func Fn6848(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6849 github.com/goccy/spidermonkeywasm2go/p7.Fn6849
func Fn6849(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6850 github.com/goccy/spidermonkeywasm2go/p7.Fn6850
func Fn6850(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6851 github.com/goccy/spidermonkeywasm2go/p7.Fn6851
func Fn6851(m *base.Module, l0 int32, l1 int32, l2 int32)

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

//go:linkname Fn6858 github.com/goccy/spidermonkeywasm2go/p6.Fn6858
func Fn6858(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6859 github.com/goccy/spidermonkeywasm2go/p6.Fn6859
func Fn6859(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6860 github.com/goccy/spidermonkeywasm2go/p6.Fn6860
func Fn6860(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6868 github.com/goccy/spidermonkeywasm2go/p7.Fn6868
func Fn6868(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6893 github.com/goccy/spidermonkeywasm2go/p7.Fn6893
func Fn6893(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6894 github.com/goccy/spidermonkeywasm2go/p5.Fn6894
func Fn6894(m *base.Module)

//go:linkname Fn6895 github.com/goccy/spidermonkeywasm2go/p7.Fn6895
func Fn6895(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6897 github.com/goccy/spidermonkeywasm2go/p3.Fn6897
func Fn6897(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6927 github.com/goccy/spidermonkeywasm2go/p6.Fn6927
func Fn6927(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6940 github.com/goccy/spidermonkeywasm2go/p6.Fn6940
func Fn6940(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6947 github.com/goccy/spidermonkeywasm2go/p7.Fn6947
func Fn6947(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6949 github.com/goccy/spidermonkeywasm2go/p0.Fn6949
func Fn6949(m *base.Module, l0 int32) int32

//go:linkname Fn6952 github.com/goccy/spidermonkeywasm2go/p0.Fn6952
func Fn6952(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6953 github.com/goccy/spidermonkeywasm2go/p0.Fn6953
func Fn6953(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6954 github.com/goccy/spidermonkeywasm2go/p0.Fn6954
func Fn6954(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6979 github.com/goccy/spidermonkeywasm2go/p7.Fn6979
func Fn6979(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6981 github.com/goccy/spidermonkeywasm2go/p7.Fn6981
func Fn6981(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6984 github.com/goccy/spidermonkeywasm2go/p6.Fn6984
func Fn6984(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6985 github.com/goccy/spidermonkeywasm2go/p4.Fn6985
func Fn6985(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7012 github.com/goccy/spidermonkeywasm2go/p7.Fn7012
func Fn7012(m *base.Module, l0 int32)

//go:linkname Fn7013 github.com/goccy/spidermonkeywasm2go/p7.Fn7013
func Fn7013(m *base.Module, l0 int32)

//go:linkname Fn7025 github.com/goccy/spidermonkeywasm2go/p5.Fn7025
func Fn7025(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7032 github.com/goccy/spidermonkeywasm2go/p7.Fn7032
func Fn7032(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7060 github.com/goccy/spidermonkeywasm2go/p6.Fn7060
func Fn7060(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7080 github.com/goccy/spidermonkeywasm2go/p6.Fn7080
func Fn7080(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7082 github.com/goccy/spidermonkeywasm2go/p4.Fn7082
func Fn7082(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7083 github.com/goccy/spidermonkeywasm2go/p2.Fn7083
func Fn7083(m *base.Module, l0 int32) int32

//go:linkname Fn7098 github.com/goccy/spidermonkeywasm2go/p5.Fn7098
func Fn7098(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7099 github.com/goccy/spidermonkeywasm2go/p5.Fn7099
func Fn7099(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7115 github.com/goccy/spidermonkeywasm2go/p6.Fn7115
func Fn7115(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7116 github.com/goccy/spidermonkeywasm2go/p6.Fn7116
func Fn7116(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7117 github.com/goccy/spidermonkeywasm2go/p6.Fn7117
func Fn7117(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7118 github.com/goccy/spidermonkeywasm2go/p7.Fn7118
func Fn7118(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7119 github.com/goccy/spidermonkeywasm2go/p6.Fn7119
func Fn7119(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7121 github.com/goccy/spidermonkeywasm2go/p5.Fn7121
func Fn7121(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7122 github.com/goccy/spidermonkeywasm2go/p5.Fn7122
func Fn7122(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn7129 github.com/goccy/spidermonkeywasm2go/p7.Fn7129
func Fn7129(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7130 github.com/goccy/spidermonkeywasm2go/p4.Fn7130
func Fn7130(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7140 github.com/goccy/spidermonkeywasm2go/p6.Fn7140
func Fn7140(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7142 github.com/goccy/spidermonkeywasm2go/p4.Fn7142
func Fn7142(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7144 github.com/goccy/spidermonkeywasm2go/p6.Fn7144
func Fn7144(m *base.Module, l0 int32) int32

//go:linkname Fn7145 github.com/goccy/spidermonkeywasm2go/p6.Fn7145
func Fn7145(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn7149 github.com/goccy/spidermonkeywasm2go/p5.Fn7149
func Fn7149(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn7150 github.com/goccy/spidermonkeywasm2go/p7.Fn7150
func Fn7150(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7177 github.com/goccy/spidermonkeywasm2go/p3.Fn7177
func Fn7177(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7178 github.com/goccy/spidermonkeywasm2go/p6.Fn7178
func Fn7178(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7179 github.com/goccy/spidermonkeywasm2go/p5.Fn7179
func Fn7179(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7180 github.com/goccy/spidermonkeywasm2go/p5.Fn7180
func Fn7180(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7181 github.com/goccy/spidermonkeywasm2go/p6.Fn7181
func Fn7181(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7182 github.com/goccy/spidermonkeywasm2go/p7.Fn7182
func Fn7182(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7183 github.com/goccy/spidermonkeywasm2go/p5.Fn7183
func Fn7183(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7184 github.com/goccy/spidermonkeywasm2go/p6.Fn7184
func Fn7184(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7185 github.com/goccy/spidermonkeywasm2go/p5.Fn7185
func Fn7185(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7186 github.com/goccy/spidermonkeywasm2go/p5.Fn7186
func Fn7186(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7187 github.com/goccy/spidermonkeywasm2go/p5.Fn7187
func Fn7187(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7188 github.com/goccy/spidermonkeywasm2go/p6.Fn7188
func Fn7188(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7189 github.com/goccy/spidermonkeywasm2go/p2.Fn7189
func Fn7189(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7193 github.com/goccy/spidermonkeywasm2go/p4.Fn7193
func Fn7193(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7194 github.com/goccy/spidermonkeywasm2go/p6.Fn7194
func Fn7194(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7195 github.com/goccy/spidermonkeywasm2go/p6.Fn7195
func Fn7195(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7199 github.com/goccy/spidermonkeywasm2go/p6.Fn7199
func Fn7199(m *base.Module, l0 int32) int32

//go:linkname Fn7201 github.com/goccy/spidermonkeywasm2go/p5.Fn7201
func Fn7201(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7203 github.com/goccy/spidermonkeywasm2go/p2.Fn7203
func Fn7203(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7217 github.com/goccy/spidermonkeywasm2go/p6.Fn7217
func Fn7217(m *base.Module, l0 int32) int32

//go:linkname Fn7218 github.com/goccy/spidermonkeywasm2go/p7.Fn7218
func Fn7218(m *base.Module, l0 int32)

//go:linkname Fn7219 github.com/goccy/spidermonkeywasm2go/p6.Fn7219
func Fn7219(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7226 github.com/goccy/spidermonkeywasm2go/p4.Fn7226
func Fn7226(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7227 github.com/goccy/spidermonkeywasm2go/p5.Fn7227
func Fn7227(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7229 github.com/goccy/spidermonkeywasm2go/p7.Fn7229
func Fn7229(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7234 github.com/goccy/spidermonkeywasm2go/p6.Fn7234
func Fn7234(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7252 github.com/goccy/spidermonkeywasm2go/p5.Fn7252
func Fn7252(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7267 github.com/goccy/spidermonkeywasm2go/p6.Fn7267
func Fn7267(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7276 github.com/goccy/spidermonkeywasm2go/p7.Fn7276
func Fn7276(m *base.Module, l0 int32)

//go:linkname Fn7347 github.com/goccy/spidermonkeywasm2go/p6.Fn7347
func Fn7347(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7364 github.com/goccy/spidermonkeywasm2go/p6.Fn7364
func Fn7364(m *base.Module, l0 int32) int32

//go:linkname Fn7369 github.com/goccy/spidermonkeywasm2go/p5.Fn7369
func Fn7369(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7373 github.com/goccy/spidermonkeywasm2go/p6.Fn7373
func Fn7373(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7374 github.com/goccy/spidermonkeywasm2go/p6.Fn7374
func Fn7374(m *base.Module, l0 int32)

//go:linkname Fn7375 github.com/goccy/spidermonkeywasm2go/p6.Fn7375
func Fn7375(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7378 github.com/goccy/spidermonkeywasm2go/p6.Fn7378
func Fn7378(m *base.Module, l0 int32) int32

//go:linkname Fn7380 github.com/goccy/spidermonkeywasm2go/p4.Fn7380
func Fn7380(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7383 github.com/goccy/spidermonkeywasm2go/p7.Fn7383
func Fn7383(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7384 github.com/goccy/spidermonkeywasm2go/p6.Fn7384
func Fn7384(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn7387 github.com/goccy/spidermonkeywasm2go/p6.Fn7387
func Fn7387(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7388 github.com/goccy/spidermonkeywasm2go/p6.Fn7388
func Fn7388(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7390 github.com/goccy/spidermonkeywasm2go/p6.Fn7390
func Fn7390(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7395 github.com/goccy/spidermonkeywasm2go/p4.Fn7395
func Fn7395(m *base.Module, l0 int32) int32

//go:linkname Fn7401 github.com/goccy/spidermonkeywasm2go/p6.Fn7401
func Fn7401(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64) int32

//go:linkname Fn7402 github.com/goccy/spidermonkeywasm2go/p5.Fn7402
func Fn7402(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7403 github.com/goccy/spidermonkeywasm2go/p7.Fn7403
func Fn7403(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7404 github.com/goccy/spidermonkeywasm2go/p7.Fn7404
func Fn7404(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7405 github.com/goccy/spidermonkeywasm2go/p7.Fn7405
func Fn7405(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7406 github.com/goccy/spidermonkeywasm2go/p7.Fn7406
func Fn7406(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7407 github.com/goccy/spidermonkeywasm2go/p7.Fn7407
func Fn7407(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7408 github.com/goccy/spidermonkeywasm2go/p6.Fn7408
func Fn7408(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7409 github.com/goccy/spidermonkeywasm2go/p6.Fn7409
func Fn7409(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn7410 github.com/goccy/spidermonkeywasm2go/p6.Fn7410
func Fn7410(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7411 github.com/goccy/spidermonkeywasm2go/p6.Fn7411
func Fn7411(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7412 github.com/goccy/spidermonkeywasm2go/p7.Fn7412
func Fn7412(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7413 github.com/goccy/spidermonkeywasm2go/p5.Fn7413
func Fn7413(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7414 github.com/goccy/spidermonkeywasm2go/p6.Fn7414
func Fn7414(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7415 github.com/goccy/spidermonkeywasm2go/p4.Fn7415
func Fn7415(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7416 github.com/goccy/spidermonkeywasm2go/p5.Fn7416
func Fn7416(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7417 github.com/goccy/spidermonkeywasm2go/p5.Fn7417
func Fn7417(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7418 github.com/goccy/spidermonkeywasm2go/p6.Fn7418
func Fn7418(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7419 github.com/goccy/spidermonkeywasm2go/p6.Fn7419
func Fn7419(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7420 github.com/goccy/spidermonkeywasm2go/p4.Fn7420
func Fn7420(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7421 github.com/goccy/spidermonkeywasm2go/p7.Fn7421
func Fn7421(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7422 github.com/goccy/spidermonkeywasm2go/p5.Fn7422
func Fn7422(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn7423 github.com/goccy/spidermonkeywasm2go/p5.Fn7423
func Fn7423(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7424 github.com/goccy/spidermonkeywasm2go/p4.Fn7424
func Fn7424(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7427 github.com/goccy/spidermonkeywasm2go/p6.Fn7427
func Fn7427(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7428 github.com/goccy/spidermonkeywasm2go/p6.Fn7428
func Fn7428(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7429 github.com/goccy/spidermonkeywasm2go/p5.Fn7429
func Fn7429(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7430 github.com/goccy/spidermonkeywasm2go/p6.Fn7430
func Fn7430(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7431 github.com/goccy/spidermonkeywasm2go/p3.Fn7431
func Fn7431(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7432 github.com/goccy/spidermonkeywasm2go/p6.Fn7432
func Fn7432(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7433 github.com/goccy/spidermonkeywasm2go/p5.Fn7433
func Fn7433(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7434 github.com/goccy/spidermonkeywasm2go/p6.Fn7434
func Fn7434(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7505 github.com/goccy/spidermonkeywasm2go/p2.Fn7505
func Fn7505(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8097 github.com/goccy/spidermonkeywasm2go/p7.Fn8097
func Fn8097(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8109 github.com/goccy/spidermonkeywasm2go/p5.Fn8109
func Fn8109(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8114 github.com/goccy/spidermonkeywasm2go/p6.Fn8114
func Fn8114(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8115 github.com/goccy/spidermonkeywasm2go/p6.Fn8115
func Fn8115(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8116 github.com/goccy/spidermonkeywasm2go/p7.Fn8116
func Fn8116(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8118 github.com/goccy/spidermonkeywasm2go/p6.Fn8118
func Fn8118(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8119 github.com/goccy/spidermonkeywasm2go/p6.Fn8119
func Fn8119(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8123 github.com/goccy/spidermonkeywasm2go/p3.Fn8123
func Fn8123(m *base.Module, l0 int32)

//go:linkname Fn8139 github.com/goccy/spidermonkeywasm2go/p6.Fn8139
func Fn8139(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8140 github.com/goccy/spidermonkeywasm2go/p6.Fn8140
func Fn8140(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8141 github.com/goccy/spidermonkeywasm2go/p6.Fn8141
func Fn8141(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8142 github.com/goccy/spidermonkeywasm2go/p6.Fn8142
func Fn8142(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8157 github.com/goccy/spidermonkeywasm2go/p6.Fn8157
func Fn8157(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8160 github.com/goccy/spidermonkeywasm2go/p7.Fn8160
func Fn8160(m *base.Module, l0 int32) int32

//go:linkname Fn8162 github.com/goccy/spidermonkeywasm2go/p7.Fn8162
func Fn8162(m *base.Module, l0 int32) int32

//go:linkname Fn8165 github.com/goccy/spidermonkeywasm2go/p6.Fn8165
func Fn8165(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8171 github.com/goccy/spidermonkeywasm2go/p2.Fn8171
func Fn8171(m *base.Module) int32

//go:linkname Fn8174 github.com/goccy/spidermonkeywasm2go/p6.Fn8174
func Fn8174(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8176 github.com/goccy/spidermonkeywasm2go/p3.Fn8176
func Fn8176(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8187 github.com/goccy/spidermonkeywasm2go/p4.Fn8187
func Fn8187(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8188 github.com/goccy/spidermonkeywasm2go/p7.Fn8188
func Fn8188(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8189 github.com/goccy/spidermonkeywasm2go/p7.Fn8189
func Fn8189(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8191 github.com/goccy/spidermonkeywasm2go/p7.Fn8191
func Fn8191(m *base.Module, l0 int32) int32

//go:linkname Fn8195 github.com/goccy/spidermonkeywasm2go/p6.Fn8195
func Fn8195(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32)

//go:linkname Fn8197 github.com/goccy/spidermonkeywasm2go/p5.Fn8197
func Fn8197(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8200 github.com/goccy/spidermonkeywasm2go/p5.Fn8200
func Fn8200(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8201 github.com/goccy/spidermonkeywasm2go/p4.Fn8201
func Fn8201(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn8202 github.com/goccy/spidermonkeywasm2go/p6.Fn8202
func Fn8202(m *base.Module, l0 int32) int32

//go:linkname Fn8206 github.com/goccy/spidermonkeywasm2go/p7.Fn8206
func Fn8206(m *base.Module, l0 int32) int32

//go:linkname Fn8208 github.com/goccy/spidermonkeywasm2go/p7.Fn8208
func Fn8208(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8210 github.com/goccy/spidermonkeywasm2go/p6.Fn8210
func Fn8210(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8211 github.com/goccy/spidermonkeywasm2go/p5.Fn8211
func Fn8211(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8213 github.com/goccy/spidermonkeywasm2go/p7.Fn8213
func Fn8213(m *base.Module, l0 int32) int32

//go:linkname Fn8214 github.com/goccy/spidermonkeywasm2go/p5.Fn8214
func Fn8214(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8215 github.com/goccy/spidermonkeywasm2go/p7.Fn8215
func Fn8215(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8274 github.com/goccy/spidermonkeywasm2go/p5.Fn8274
func Fn8274(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8297 github.com/goccy/spidermonkeywasm2go/p6.Fn8297
func Fn8297(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8327 github.com/goccy/spidermonkeywasm2go/p5.Fn8327
func Fn8327(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8328 github.com/goccy/spidermonkeywasm2go/p5.Fn8328
func Fn8328(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8340 github.com/goccy/spidermonkeywasm2go/p4.Fn8340
func Fn8340(m *base.Module, l0 int32) int32

//go:linkname Fn8352 github.com/goccy/spidermonkeywasm2go/p3.Fn8352
func Fn8352(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8360 github.com/goccy/spidermonkeywasm2go/p7.Fn8360
func Fn8360(m *base.Module, l0 int32)

//go:linkname Fn8362 github.com/goccy/spidermonkeywasm2go/p6.Fn8362
func Fn8362(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8363 github.com/goccy/spidermonkeywasm2go/p5.Fn8363
func Fn8363(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8364 github.com/goccy/spidermonkeywasm2go/p6.Fn8364
func Fn8364(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8365 github.com/goccy/spidermonkeywasm2go/p6.Fn8365
func Fn8365(m *base.Module, l0 int32)

//go:linkname Fn8367 github.com/goccy/spidermonkeywasm2go/p6.Fn8367
func Fn8367(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8368 github.com/goccy/spidermonkeywasm2go/p5.Fn8368
func Fn8368(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8369 github.com/goccy/spidermonkeywasm2go/p6.Fn8369
func Fn8369(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8381 github.com/goccy/spidermonkeywasm2go/p5.Fn8381
func Fn8381(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn8383 github.com/goccy/spidermonkeywasm2go/p6.Fn8383
func Fn8383(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8388 github.com/goccy/spidermonkeywasm2go/p5.Fn8388
func Fn8388(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8391 github.com/goccy/spidermonkeywasm2go/p6.Fn8391
func Fn8391(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8392 github.com/goccy/spidermonkeywasm2go/p3.Fn8392
func Fn8392(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn8393 github.com/goccy/spidermonkeywasm2go/p7.Fn8393
func Fn8393(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8394 github.com/goccy/spidermonkeywasm2go/p7.Fn8394
func Fn8394(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8395 github.com/goccy/spidermonkeywasm2go/p7.Fn8395
func Fn8395(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8397 github.com/goccy/spidermonkeywasm2go/p4.Fn8397
func Fn8397(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8398 github.com/goccy/spidermonkeywasm2go/p6.Fn8398
func Fn8398(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8410 github.com/goccy/spidermonkeywasm2go/p6.Fn8410
func Fn8410(m *base.Module, l0 int32) int32

//go:linkname Fn8411 github.com/goccy/spidermonkeywasm2go/p6.Fn8411
func Fn8411(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8413 github.com/goccy/spidermonkeywasm2go/p7.Fn8413
func Fn8413(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8415 github.com/goccy/spidermonkeywasm2go/p6.Fn8415
func Fn8415(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8416 github.com/goccy/spidermonkeywasm2go/p6.Fn8416
func Fn8416(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8417 github.com/goccy/spidermonkeywasm2go/p7.Fn8417
func Fn8417(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8418 github.com/goccy/spidermonkeywasm2go/p7.Fn8418
func Fn8418(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8419 github.com/goccy/spidermonkeywasm2go/p7.Fn8419
func Fn8419(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8420 github.com/goccy/spidermonkeywasm2go/p7.Fn8420
func Fn8420(m *base.Module, l0 int32) int32

//go:linkname Fn8421 github.com/goccy/spidermonkeywasm2go/p4.Fn8421
func Fn8421(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8422 github.com/goccy/spidermonkeywasm2go/p4.Fn8422
func Fn8422(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8423 github.com/goccy/spidermonkeywasm2go/p6.Fn8423
func Fn8423(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8424 github.com/goccy/spidermonkeywasm2go/p6.Fn8424
func Fn8424(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8434 github.com/goccy/spidermonkeywasm2go/p6.Fn8434
func Fn8434(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8437 github.com/goccy/spidermonkeywasm2go/p6.Fn8437
func Fn8437(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8439 github.com/goccy/spidermonkeywasm2go/p5.Fn8439
func Fn8439(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8442 github.com/goccy/spidermonkeywasm2go/p6.Fn8442
func Fn8442(m *base.Module, l0 int32) int32

//go:linkname Fn8448 github.com/goccy/spidermonkeywasm2go/p6.Fn8448
func Fn8448(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8464 github.com/goccy/spidermonkeywasm2go/p6.Fn8464
func Fn8464(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8468 github.com/goccy/spidermonkeywasm2go/p6.Fn8468
func Fn8468(m *base.Module, l0 int32) int32

//go:linkname Fn8478 github.com/goccy/spidermonkeywasm2go/p5.Fn8478
func Fn8478(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8494 github.com/goccy/spidermonkeywasm2go/p6.Fn8494
func Fn8494(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8495 github.com/goccy/spidermonkeywasm2go/p6.Fn8495
func Fn8495(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8496 github.com/goccy/spidermonkeywasm2go/p6.Fn8496
func Fn8496(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8501 github.com/goccy/spidermonkeywasm2go/p6.Fn8501
func Fn8501(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8502 github.com/goccy/spidermonkeywasm2go/p6.Fn8502
func Fn8502(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8514 github.com/goccy/spidermonkeywasm2go/p7.Fn8514
func Fn8514(m *base.Module, l0 int32)

//go:linkname Fn8518 github.com/goccy/spidermonkeywasm2go/p5.Fn8518
func Fn8518(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8522 github.com/goccy/spidermonkeywasm2go/p5.Fn8522
func Fn8522(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8523 github.com/goccy/spidermonkeywasm2go/p5.Fn8523
func Fn8523(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8524 github.com/goccy/spidermonkeywasm2go/p7.Fn8524
func Fn8524(m *base.Module, l0 int32)

//go:linkname Fn8527 github.com/goccy/spidermonkeywasm2go/p5.Fn8527
func Fn8527(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8528 github.com/goccy/spidermonkeywasm2go/p7.Fn8528
func Fn8528(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8529 github.com/goccy/spidermonkeywasm2go/p6.Fn8529
func Fn8529(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8530 github.com/goccy/spidermonkeywasm2go/p6.Fn8530
func Fn8530(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8531 github.com/goccy/spidermonkeywasm2go/p3.Fn8531
func Fn8531(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8532 github.com/goccy/spidermonkeywasm2go/p5.Fn8532
func Fn8532(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8534 github.com/goccy/spidermonkeywasm2go/p7.Fn8534
func Fn8534(m *base.Module, l0 int32) int32

//go:linkname Fn8536 github.com/goccy/spidermonkeywasm2go/p7.Fn8536
func Fn8536(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8538 github.com/goccy/spidermonkeywasm2go/p7.Fn8538
func Fn8538(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8542 github.com/goccy/spidermonkeywasm2go/p7.Fn8542
func Fn8542(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8544 github.com/goccy/spidermonkeywasm2go/p6.Fn8544
func Fn8544(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8552 github.com/goccy/spidermonkeywasm2go/p6.Fn8552
func Fn8552(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8553 github.com/goccy/spidermonkeywasm2go/p6.Fn8553
func Fn8553(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8556 github.com/goccy/spidermonkeywasm2go/p7.Fn8556
func Fn8556(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8558 github.com/goccy/spidermonkeywasm2go/p7.Fn8558
func Fn8558(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8579 github.com/goccy/spidermonkeywasm2go/p7.Fn8579
func Fn8579(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8625 github.com/goccy/spidermonkeywasm2go/p7.Fn8625
func Fn8625(m *base.Module, l0 int32)

//go:linkname Fn8626 github.com/goccy/spidermonkeywasm2go/p7.Fn8626
func Fn8626(m *base.Module, l0 int32)

//go:linkname Fn8627 github.com/goccy/spidermonkeywasm2go/p6.Fn8627
func Fn8627(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8628 github.com/goccy/spidermonkeywasm2go/p7.Fn8628
func Fn8628(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8629 github.com/goccy/spidermonkeywasm2go/p6.Fn8629
func Fn8629(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8635 github.com/goccy/spidermonkeywasm2go/p6.Fn8635
func Fn8635(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8637 github.com/goccy/spidermonkeywasm2go/p6.Fn8637
func Fn8637(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8639 github.com/goccy/spidermonkeywasm2go/p6.Fn8639
func Fn8639(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn8640 github.com/goccy/spidermonkeywasm2go/p6.Fn8640
func Fn8640(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8642 github.com/goccy/spidermonkeywasm2go/p6.Fn8642
func Fn8642(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8654 github.com/goccy/spidermonkeywasm2go/p2.Fn8654
func Fn8654(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn8658 github.com/goccy/spidermonkeywasm2go/p6.Fn8658
func Fn8658(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8659 github.com/goccy/spidermonkeywasm2go/p6.Fn8659
func Fn8659(m *base.Module, l0 int32)

//go:linkname Fn8665 github.com/goccy/spidermonkeywasm2go/p6.Fn8665
func Fn8665(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8674 github.com/goccy/spidermonkeywasm2go/p6.Fn8674
func Fn8674(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8680 github.com/goccy/spidermonkeywasm2go/p5.Fn8680
func Fn8680(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8681 github.com/goccy/spidermonkeywasm2go/p6.Fn8681
func Fn8681(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8682 github.com/goccy/spidermonkeywasm2go/p6.Fn8682
func Fn8682(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8683 github.com/goccy/spidermonkeywasm2go/p7.Fn8683
func Fn8683(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8684 github.com/goccy/spidermonkeywasm2go/p6.Fn8684
func Fn8684(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8691 github.com/goccy/spidermonkeywasm2go/p4.Fn8691
func Fn8691(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8705 github.com/goccy/spidermonkeywasm2go/p4.Fn8705
func Fn8705(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8709 github.com/goccy/spidermonkeywasm2go/p3.Fn8709
func Fn8709(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn8711 github.com/goccy/spidermonkeywasm2go/p5.Fn8711
func Fn8711(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8712 github.com/goccy/spidermonkeywasm2go/p6.Fn8712
func Fn8712(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8716 github.com/goccy/spidermonkeywasm2go/p6.Fn8716
func Fn8716(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8717 github.com/goccy/spidermonkeywasm2go/p5.Fn8717
func Fn8717(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8721 github.com/goccy/spidermonkeywasm2go/p7.Fn8721
func Fn8721(m *base.Module, l0 int32) int32

//go:linkname Fn8722 github.com/goccy/spidermonkeywasm2go/p7.Fn8722
func Fn8722(m *base.Module, l0 int32)

//go:linkname Fn8725 github.com/goccy/spidermonkeywasm2go/p6.Fn8725
func Fn8725(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8726 github.com/goccy/spidermonkeywasm2go/p6.Fn8726
func Fn8726(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8727 github.com/goccy/spidermonkeywasm2go/p6.Fn8727
func Fn8727(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8755 github.com/goccy/spidermonkeywasm2go/p7.Fn8755
func Fn8755(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9006 github.com/goccy/spidermonkeywasm2go/p4.Fn9006
func Fn9006(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn9007 github.com/goccy/spidermonkeywasm2go/p6.Fn9007
func Fn9007(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9027 github.com/goccy/spidermonkeywasm2go/p5.Fn9027
func Fn9027(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9028 github.com/goccy/spidermonkeywasm2go/p7.Fn9028
func Fn9028(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9060 github.com/goccy/spidermonkeywasm2go/p6.Fn9060
func Fn9060(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9065 github.com/goccy/spidermonkeywasm2go/p3.Fn9065
func Fn9065(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9075 github.com/goccy/spidermonkeywasm2go/p4.Fn9075
func Fn9075(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9077 github.com/goccy/spidermonkeywasm2go/p4.Fn9077
func Fn9077(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9078 github.com/goccy/spidermonkeywasm2go/p5.Fn9078
func Fn9078(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9080 github.com/goccy/spidermonkeywasm2go/p6.Fn9080
func Fn9080(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9100 github.com/goccy/spidermonkeywasm2go/p2.Fn9100
func Fn9100(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9102 github.com/goccy/spidermonkeywasm2go/p4.Fn9102
func Fn9102(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9106 github.com/goccy/spidermonkeywasm2go/p4.Fn9106
func Fn9106(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9140 github.com/goccy/spidermonkeywasm2go/p6.Fn9140
func Fn9140(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9156 github.com/goccy/spidermonkeywasm2go/p5.Fn9156
func Fn9156(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9274 github.com/goccy/spidermonkeywasm2go/p3.Fn9274
func Fn9274(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9297 github.com/goccy/spidermonkeywasm2go/p6.Fn9297
func Fn9297(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9298 github.com/goccy/spidermonkeywasm2go/p6.Fn9298
func Fn9298(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9299 github.com/goccy/spidermonkeywasm2go/p7.Fn9299
func Fn9299(m *base.Module, l0 int32) int32

//go:linkname Fn9300 github.com/goccy/spidermonkeywasm2go/p6.Fn9300
func Fn9300(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9301 github.com/goccy/spidermonkeywasm2go/p7.Fn9301
func Fn9301(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9305 github.com/goccy/spidermonkeywasm2go/p5.Fn9305
func Fn9305(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9337 github.com/goccy/spidermonkeywasm2go/p3.Fn9337
func Fn9337(m *base.Module, l0 int32) int32

//go:linkname Fn9348 github.com/goccy/spidermonkeywasm2go/p7.Fn9348
func Fn9348(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9351 github.com/goccy/spidermonkeywasm2go/p6.Fn9351
func Fn9351(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9355 github.com/goccy/spidermonkeywasm2go/p3.Fn9355
func Fn9355(m *base.Module, l0 int32)

//go:linkname Fn9356 github.com/goccy/spidermonkeywasm2go/p4.Fn9356
func Fn9356(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9357 github.com/goccy/spidermonkeywasm2go/p5.Fn9357
func Fn9357(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9362 github.com/goccy/spidermonkeywasm2go/p5.Fn9362
func Fn9362(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9365 github.com/goccy/spidermonkeywasm2go/p6.Fn9365
func Fn9365(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9368 github.com/goccy/spidermonkeywasm2go/p6.Fn9368
func Fn9368(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9369 github.com/goccy/spidermonkeywasm2go/p5.Fn9369
func Fn9369(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9370 github.com/goccy/spidermonkeywasm2go/p5.Fn9370
func Fn9370(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9371 github.com/goccy/spidermonkeywasm2go/p6.Fn9371
func Fn9371(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9373 github.com/goccy/spidermonkeywasm2go/p5.Fn9373
func Fn9373(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9374 github.com/goccy/spidermonkeywasm2go/p5.Fn9374
func Fn9374(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9427 github.com/goccy/spidermonkeywasm2go/p5.Fn9427
func Fn9427(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9428 github.com/goccy/spidermonkeywasm2go/p2.Fn9428
func Fn9428(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn9431 github.com/goccy/spidermonkeywasm2go/p5.Fn9431
func Fn9431(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9437 github.com/goccy/spidermonkeywasm2go/p5.Fn9437
func Fn9437(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9444 github.com/goccy/spidermonkeywasm2go/p4.Fn9444
func Fn9444(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn9487 github.com/goccy/spidermonkeywasm2go/p4.Fn9487
func Fn9487(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9488 github.com/goccy/spidermonkeywasm2go/p3.Fn9488
func Fn9488(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9521 github.com/goccy/spidermonkeywasm2go/p6.Fn9521
func Fn9521(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9524 github.com/goccy/spidermonkeywasm2go/p6.Fn9524
func Fn9524(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9530 github.com/goccy/spidermonkeywasm2go/p6.Fn9530
func Fn9530(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9532 github.com/goccy/spidermonkeywasm2go/p6.Fn9532
func Fn9532(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9562 github.com/goccy/spidermonkeywasm2go/p7.Fn9562
func Fn9562(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9563 github.com/goccy/spidermonkeywasm2go/p7.Fn9563
func Fn9563(m *base.Module, l0 int32)

//go:linkname Fn9564 github.com/goccy/spidermonkeywasm2go/p6.Fn9564
func Fn9564(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn9565 github.com/goccy/spidermonkeywasm2go/p6.Fn9565
func Fn9565(m *base.Module, l0 int32, l1 float64, l2 int32)

//go:linkname Fn9578 github.com/goccy/spidermonkeywasm2go/p5.Fn9578
func Fn9578(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn9580 github.com/goccy/spidermonkeywasm2go/p4.Fn9580
func Fn9580(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9594 github.com/goccy/spidermonkeywasm2go/p5.Fn9594
func Fn9594(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9607 github.com/goccy/spidermonkeywasm2go/p6.Fn9607
func Fn9607(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9626 github.com/goccy/spidermonkeywasm2go/p5.Fn9626
func Fn9626(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn9632 github.com/goccy/spidermonkeywasm2go/p4.Fn9632
func Fn9632(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9633 github.com/goccy/spidermonkeywasm2go/p5.Fn9633
func Fn9633(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9647 github.com/goccy/spidermonkeywasm2go/p2.Fn9647
func Fn9647(m *base.Module, l0 int32)

//go:linkname Fn9648 github.com/goccy/spidermonkeywasm2go/p3.Fn9648
func Fn9648(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9652 github.com/goccy/spidermonkeywasm2go/p7.Fn9652
func Fn9652(m *base.Module, l0 int32) int32

//go:linkname Fn9654 github.com/goccy/spidermonkeywasm2go/p6.Fn9654
func Fn9654(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9674 github.com/goccy/spidermonkeywasm2go/p7.Fn9674
func Fn9674(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9677 github.com/goccy/spidermonkeywasm2go/p6.Fn9677
func Fn9677(m *base.Module, l0 int32)

//go:linkname Fn9678 github.com/goccy/spidermonkeywasm2go/p6.Fn9678
func Fn9678(m *base.Module, l0 int32)

//go:linkname Fn9679 github.com/goccy/spidermonkeywasm2go/p6.Fn9679
func Fn9679(m *base.Module, l0 int32)

//go:linkname Fn9686 github.com/goccy/spidermonkeywasm2go/p6.Fn9686
func Fn9686(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9688 github.com/goccy/spidermonkeywasm2go/p6.Fn9688
func Fn9688(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9698 github.com/goccy/spidermonkeywasm2go/p5.Fn9698
func Fn9698(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9710 github.com/goccy/spidermonkeywasm2go/p4.Fn9710
func Fn9710(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9711 github.com/goccy/spidermonkeywasm2go/p4.Fn9711
func Fn9711(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9768 github.com/goccy/spidermonkeywasm2go/p7.Fn9768
func Fn9768(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9773 github.com/goccy/spidermonkeywasm2go/p7.Fn9773
func Fn9773(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10118 github.com/goccy/spidermonkeywasm2go/p6.Fn10118
func Fn10118(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10128 github.com/goccy/spidermonkeywasm2go/p7.Fn10128
func Fn10128(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10155 github.com/goccy/spidermonkeywasm2go/p5.Fn10155
func Fn10155(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10156 github.com/goccy/spidermonkeywasm2go/p7.Fn10156
func Fn10156(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10157 github.com/goccy/spidermonkeywasm2go/p7.Fn10157
func Fn10157(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10158 github.com/goccy/spidermonkeywasm2go/p7.Fn10158
func Fn10158(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10180 github.com/goccy/spidermonkeywasm2go/p6.Fn10180
func Fn10180(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10181 github.com/goccy/spidermonkeywasm2go/p5.Fn10181
func Fn10181(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10182 github.com/goccy/spidermonkeywasm2go/p6.Fn10182
func Fn10182(m *base.Module, l0 int32)

//go:linkname Fn10183 github.com/goccy/spidermonkeywasm2go/p6.Fn10183
func Fn10183(m *base.Module, l0 int32) int32

//go:linkname Fn10184 github.com/goccy/spidermonkeywasm2go/p7.Fn10184
func Fn10184(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10185 github.com/goccy/spidermonkeywasm2go/p6.Fn10185
func Fn10185(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10186 github.com/goccy/spidermonkeywasm2go/p4.Fn10186
func Fn10186(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10199 github.com/goccy/spidermonkeywasm2go/p3.Fn10199
func Fn10199(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10200 github.com/goccy/spidermonkeywasm2go/p5.Fn10200
func Fn10200(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10208 github.com/goccy/spidermonkeywasm2go/p3.Fn10208
func Fn10208(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10224 github.com/goccy/spidermonkeywasm2go/p6.Fn10224
func Fn10224(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10226 github.com/goccy/spidermonkeywasm2go/p4.Fn10226
func Fn10226(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10229 github.com/goccy/spidermonkeywasm2go/p5.Fn10229
func Fn10229(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10230 github.com/goccy/spidermonkeywasm2go/p5.Fn10230
func Fn10230(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10233 github.com/goccy/spidermonkeywasm2go/p6.Fn10233
func Fn10233(m *base.Module, l0 int32, l1 int64, l2 int64) int32

//go:linkname Fn10248 github.com/goccy/spidermonkeywasm2go/p6.Fn10248
func Fn10248(m *base.Module, l0 int32)

//go:linkname Fn10249 github.com/goccy/spidermonkeywasm2go/p5.Fn10249
func Fn10249(m *base.Module, l0 int32) int32

//go:linkname Fn10250 github.com/goccy/spidermonkeywasm2go/p5.Fn10250
func Fn10250(m *base.Module, l0 int32)

//go:linkname Fn10251 github.com/goccy/spidermonkeywasm2go/p5.Fn10251
func Fn10251(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10252 github.com/goccy/spidermonkeywasm2go/p6.Fn10252
func Fn10252(m *base.Module, l0 int32) int32

//go:linkname Fn10253 github.com/goccy/spidermonkeywasm2go/p3.Fn10253
func Fn10253(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10265 github.com/goccy/spidermonkeywasm2go/p6.Fn10265
func Fn10265(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10269 github.com/goccy/spidermonkeywasm2go/p7.Fn10269
func Fn10269(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10274 github.com/goccy/spidermonkeywasm2go/p7.Fn10274
func Fn10274(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10277 github.com/goccy/spidermonkeywasm2go/p7.Fn10277
func Fn10277(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10280 github.com/goccy/spidermonkeywasm2go/p7.Fn10280
func Fn10280(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10282 github.com/goccy/spidermonkeywasm2go/p7.Fn10282
func Fn10282(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10291 github.com/goccy/spidermonkeywasm2go/p5.Fn10291
func Fn10291(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn10293 github.com/goccy/spidermonkeywasm2go/p6.Fn10293
func Fn10293(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10302 github.com/goccy/spidermonkeywasm2go/p3.Fn10302
func Fn10302(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10313 github.com/goccy/spidermonkeywasm2go/p3.Fn10313
func Fn10313(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10344 github.com/goccy/spidermonkeywasm2go/p6.Fn10344
func Fn10344(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10345 github.com/goccy/spidermonkeywasm2go/p5.Fn10345
func Fn10345(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10346 github.com/goccy/spidermonkeywasm2go/p4.Fn10346
func Fn10346(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn10347 github.com/goccy/spidermonkeywasm2go/p6.Fn10347
func Fn10347(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10348 github.com/goccy/spidermonkeywasm2go/p5.Fn10348
func Fn10348(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10349 github.com/goccy/spidermonkeywasm2go/p6.Fn10349
func Fn10349(m *base.Module, l0 int32) int32

//go:linkname Fn10559 github.com/goccy/spidermonkeywasm2go/p6.Fn10559
func Fn10559(m *base.Module, l0 int32) int32

//go:linkname Fn10578 github.com/goccy/spidermonkeywasm2go/p7.Fn10578
func Fn10578(m *base.Module, l0 int32) int32

//go:linkname Fn10792 github.com/goccy/spidermonkeywasm2go/p6.Fn10792
func Fn10792(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10795 github.com/goccy/spidermonkeywasm2go/p6.Fn10795
func Fn10795(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10796 github.com/goccy/spidermonkeywasm2go/p6.Fn10796
func Fn10796(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn10818 github.com/goccy/spidermonkeywasm2go/p7.Fn10818
func Fn10818(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10823 github.com/goccy/spidermonkeywasm2go/p5.Fn10823
func Fn10823(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10828 github.com/goccy/spidermonkeywasm2go/p6.Fn10828
func Fn10828(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10832 github.com/goccy/spidermonkeywasm2go/p6.Fn10832
func Fn10832(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10833 github.com/goccy/spidermonkeywasm2go/p5.Fn10833
func Fn10833(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn10938 github.com/goccy/spidermonkeywasm2go/p7.Fn10938
func Fn10938(m *base.Module)

//go:linkname Fn10940 github.com/goccy/spidermonkeywasm2go/p2.Fn10940
func Fn10940(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10971 github.com/goccy/spidermonkeywasm2go/p6.Fn10971
func Fn10971(m *base.Module, l0 int32) int32

//go:linkname Fn10972 github.com/goccy/spidermonkeywasm2go/p6.Fn10972
func Fn10972(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10975 github.com/goccy/spidermonkeywasm2go/p7.Fn10975
func Fn10975(m *base.Module, l0 int32)
