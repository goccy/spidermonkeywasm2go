//go:build (!amd64 && !arm64) || purego

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

//go:linkname Fn98 github.com/goccy/spidermonkeywasm2go/p7.Fn98
func Fn98(m *base.Module, l0 int32) int32

//go:linkname Fn104 github.com/goccy/spidermonkeywasm2go/p7.Fn104
func Fn104(m *base.Module, l0 int32)

//go:linkname Fn105 github.com/goccy/spidermonkeywasm2go/p7.Fn105
func Fn105(m *base.Module, l0 int32)

//go:linkname Fn123 github.com/goccy/spidermonkeywasm2go/p7.Fn123
func Fn123(m *base.Module, l0 int32)

//go:linkname Fn127 github.com/goccy/spidermonkeywasm2go/p7.Fn127
func Fn127(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn139 github.com/goccy/spidermonkeywasm2go/p7.Fn139
func Fn139(m *base.Module)

//go:linkname Fn155 github.com/goccy/spidermonkeywasm2go/p0.Fn155
func Fn155(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn161 github.com/goccy/spidermonkeywasm2go/p3.Fn161
func Fn161(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn207 github.com/goccy/spidermonkeywasm2go/p6.Fn207
func Fn207(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn210 github.com/goccy/spidermonkeywasm2go/p3.Fn210
func Fn210(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn214 github.com/goccy/spidermonkeywasm2go/p6.Fn214
func Fn214(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn305 github.com/goccy/spidermonkeywasm2go/p3.Fn305
func Fn305(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn341 github.com/goccy/spidermonkeywasm2go/p7.Fn341
func Fn341(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn343 github.com/goccy/spidermonkeywasm2go/p7.Fn343
func Fn343(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn346 github.com/goccy/spidermonkeywasm2go/p7.Fn346
func Fn346(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn349 github.com/goccy/spidermonkeywasm2go/p7.Fn349
func Fn349(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn352 github.com/goccy/spidermonkeywasm2go/p7.Fn352
func Fn352(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn355 github.com/goccy/spidermonkeywasm2go/p7.Fn355
func Fn355(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn358 github.com/goccy/spidermonkeywasm2go/p7.Fn358
func Fn358(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn361 github.com/goccy/spidermonkeywasm2go/p7.Fn361
func Fn361(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn364 github.com/goccy/spidermonkeywasm2go/p7.Fn364
func Fn364(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn383 github.com/goccy/spidermonkeywasm2go/p5.Fn383
func Fn383(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn654 github.com/goccy/spidermonkeywasm2go/p7.Fn654
func Fn654(m *base.Module, l0 int32) int32

//go:linkname Fn655 github.com/goccy/spidermonkeywasm2go/p7.Fn655
func Fn655(m *base.Module, l0 int32)

//go:linkname Fn669 github.com/goccy/spidermonkeywasm2go/p7.Fn669
func Fn669(m *base.Module, l0 int32) int32

//go:linkname Fn672 github.com/goccy/spidermonkeywasm2go/p7.Fn672
func Fn672(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn682 github.com/goccy/spidermonkeywasm2go/p6.Fn682
func Fn682(m *base.Module, l0 int32) int32

//go:linkname Fn685 github.com/goccy/spidermonkeywasm2go/p6.Fn685
func Fn685(m *base.Module, l0 int32)

//go:linkname Fn687 github.com/goccy/spidermonkeywasm2go/p7.Fn687
func Fn687(m *base.Module, l0 int32) int32

//go:linkname Fn697 github.com/goccy/spidermonkeywasm2go/p7.Fn697
func Fn697(m *base.Module, l0 int32)

//go:linkname Fn698 github.com/goccy/spidermonkeywasm2go/p7.Fn698
func Fn698(m *base.Module, l0 int32) int32

//go:linkname Fn725 github.com/goccy/spidermonkeywasm2go/p6.Fn725
func Fn725(m *base.Module, l0 int32) int32

//go:linkname Fn733 github.com/goccy/spidermonkeywasm2go/p5.Fn733
func Fn733(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn734 github.com/goccy/spidermonkeywasm2go/p7.Fn734
func Fn734(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn746 github.com/goccy/spidermonkeywasm2go/p6.Fn746
func Fn746(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn751 github.com/goccy/spidermonkeywasm2go/p7.Fn751
func Fn751(m *base.Module, l0 int32)

//go:linkname Fn752 github.com/goccy/spidermonkeywasm2go/p7.Fn752
func Fn752(m *base.Module, l0 int32)

//go:linkname Fn753 github.com/goccy/spidermonkeywasm2go/p7.Fn753
func Fn753(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn762 github.com/goccy/spidermonkeywasm2go/p7.Fn762
func Fn762(m *base.Module, l0 int32) int32

//go:linkname Fn772 github.com/goccy/spidermonkeywasm2go/p7.Fn772
func Fn772(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn773 github.com/goccy/spidermonkeywasm2go/p6.Fn773
func Fn773(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn774 github.com/goccy/spidermonkeywasm2go/p7.Fn774
func Fn774(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn775 github.com/goccy/spidermonkeywasm2go/p7.Fn775
func Fn775(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn776 github.com/goccy/spidermonkeywasm2go/p7.Fn776
func Fn776(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn778 github.com/goccy/spidermonkeywasm2go/p7.Fn778
func Fn778(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn779 github.com/goccy/spidermonkeywasm2go/p6.Fn779
func Fn779(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn781 github.com/goccy/spidermonkeywasm2go/p7.Fn781
func Fn781(m *base.Module, l0 int32) int32

//go:linkname Fn782 github.com/goccy/spidermonkeywasm2go/p7.Fn782
func Fn782(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn783 github.com/goccy/spidermonkeywasm2go/p6.Fn783
func Fn783(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn785 github.com/goccy/spidermonkeywasm2go/p7.Fn785
func Fn785(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn786 github.com/goccy/spidermonkeywasm2go/p4.Fn786
func Fn786(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn787 github.com/goccy/spidermonkeywasm2go/p6.Fn787
func Fn787(m *base.Module, l0 int32) int32

//go:linkname Fn792 github.com/goccy/spidermonkeywasm2go/p6.Fn792
func Fn792(m *base.Module, l0 int32)

//go:linkname Fn793 github.com/goccy/spidermonkeywasm2go/p7.Fn793
func Fn793(m *base.Module, l0 int32)

//go:linkname Fn797 github.com/goccy/spidermonkeywasm2go/p7.Fn797
func Fn797(m *base.Module, l0 int32) int32

//go:linkname Fn806 github.com/goccy/spidermonkeywasm2go/p6.Fn806
func Fn806(m *base.Module, l0 int32) int32

//go:linkname Fn808 github.com/goccy/spidermonkeywasm2go/p6.Fn808
func Fn808(m *base.Module, l0 int32) int32

//go:linkname Fn824 github.com/goccy/spidermonkeywasm2go/p2.Fn824
func Fn824(m *base.Module, l0 int32) int32

//go:linkname Fn825 github.com/goccy/spidermonkeywasm2go/p4.Fn825
func Fn825(m *base.Module, l0 int32)

//go:linkname Fn826 github.com/goccy/spidermonkeywasm2go/p7.Fn826
func Fn826(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn827 github.com/goccy/spidermonkeywasm2go/p5.Fn827
func Fn827(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn830 github.com/goccy/spidermonkeywasm2go/p6.Fn830
func Fn830(m *base.Module, l0 int32) int32

//go:linkname Fn832 github.com/goccy/spidermonkeywasm2go/p0.Fn832
func Fn832(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn834 github.com/goccy/spidermonkeywasm2go/p5.Fn834
func Fn834(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn837 github.com/goccy/spidermonkeywasm2go/p0.Fn837
func Fn837(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn842 github.com/goccy/spidermonkeywasm2go/p0.Fn842
func Fn842(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn843 github.com/goccy/spidermonkeywasm2go/p0.Fn843
func Fn843(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn847 github.com/goccy/spidermonkeywasm2go/p7.Fn847
func Fn847(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn850 github.com/goccy/spidermonkeywasm2go/p0.Fn850
func Fn850(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn851 github.com/goccy/spidermonkeywasm2go/p0.Fn851
func Fn851(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn852 github.com/goccy/spidermonkeywasm2go/p0.Fn852
func Fn852(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn857 github.com/goccy/spidermonkeywasm2go/p3.Fn857
func Fn857(m *base.Module, l0 int32) int32

//go:linkname Fn858 github.com/goccy/spidermonkeywasm2go/p5.Fn858
func Fn858(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn859 github.com/goccy/spidermonkeywasm2go/p4.Fn859
func Fn859(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn862 github.com/goccy/spidermonkeywasm2go/p0.Fn862
func Fn862(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn863 github.com/goccy/spidermonkeywasm2go/p7.Fn863
func Fn863(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn864 github.com/goccy/spidermonkeywasm2go/p6.Fn864
func Fn864(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn867 github.com/goccy/spidermonkeywasm2go/p6.Fn867
func Fn867(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn868 github.com/goccy/spidermonkeywasm2go/p6.Fn868
func Fn868(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn875 github.com/goccy/spidermonkeywasm2go/p0.Fn875
func Fn875(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn883 github.com/goccy/spidermonkeywasm2go/p0.Fn883
func Fn883(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn885 github.com/goccy/spidermonkeywasm2go/p6.Fn885
func Fn885(m *base.Module, l0 int32) int32

//go:linkname Fn892 github.com/goccy/spidermonkeywasm2go/p5.Fn892
func Fn892(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn916 github.com/goccy/spidermonkeywasm2go/p6.Fn916
func Fn916(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn932 github.com/goccy/spidermonkeywasm2go/p7.Fn932
func Fn932(m *base.Module, l0 int32) int32

//go:linkname Fn1002 github.com/goccy/spidermonkeywasm2go/p5.Fn1002
func Fn1002(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1028 github.com/goccy/spidermonkeywasm2go/p5.Fn1028
func Fn1028(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1086 github.com/goccy/spidermonkeywasm2go/p0.Fn1086
func Fn1086(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1087 github.com/goccy/spidermonkeywasm2go/p0.Fn1087
func Fn1087(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1088 github.com/goccy/spidermonkeywasm2go/p0.Fn1088
func Fn1088(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1102 github.com/goccy/spidermonkeywasm2go/p0.Fn1102
func Fn1102(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1126 github.com/goccy/spidermonkeywasm2go/p6.Fn1126
func Fn1126(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1136 github.com/goccy/spidermonkeywasm2go/p6.Fn1136
func Fn1136(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1138 github.com/goccy/spidermonkeywasm2go/p4.Fn1138
func Fn1138(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1139 github.com/goccy/spidermonkeywasm2go/p5.Fn1139
func Fn1139(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1140 github.com/goccy/spidermonkeywasm2go/p5.Fn1140
func Fn1140(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1141 github.com/goccy/spidermonkeywasm2go/p5.Fn1141
func Fn1141(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1142 github.com/goccy/spidermonkeywasm2go/p5.Fn1142
func Fn1142(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1144 github.com/goccy/spidermonkeywasm2go/p6.Fn1144
func Fn1144(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1145 github.com/goccy/spidermonkeywasm2go/p6.Fn1145
func Fn1145(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1146 github.com/goccy/spidermonkeywasm2go/p6.Fn1146
func Fn1146(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1147 github.com/goccy/spidermonkeywasm2go/p5.Fn1147
func Fn1147(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1162 github.com/goccy/spidermonkeywasm2go/p6.Fn1162
func Fn1162(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1164 github.com/goccy/spidermonkeywasm2go/p6.Fn1164
func Fn1164(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1165 github.com/goccy/spidermonkeywasm2go/p6.Fn1165
func Fn1165(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1177 github.com/goccy/spidermonkeywasm2go/p4.Fn1177
func Fn1177(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn1193 github.com/goccy/spidermonkeywasm2go/p6.Fn1193
func Fn1193(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn1281 github.com/goccy/spidermonkeywasm2go/p4.Fn1281
func Fn1281(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1282 github.com/goccy/spidermonkeywasm2go/p6.Fn1282
func Fn1282(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1297 github.com/goccy/spidermonkeywasm2go/p6.Fn1297
func Fn1297(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1298 github.com/goccy/spidermonkeywasm2go/p0.Fn1298
func Fn1298(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1299 github.com/goccy/spidermonkeywasm2go/p0.Fn1299
func Fn1299(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1300 github.com/goccy/spidermonkeywasm2go/p4.Fn1300
func Fn1300(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1301 github.com/goccy/spidermonkeywasm2go/p5.Fn1301
func Fn1301(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1302 github.com/goccy/spidermonkeywasm2go/p5.Fn1302
func Fn1302(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1303 github.com/goccy/spidermonkeywasm2go/p4.Fn1303
func Fn1303(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1304 github.com/goccy/spidermonkeywasm2go/p2.Fn1304
func Fn1304(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1308 github.com/goccy/spidermonkeywasm2go/p3.Fn1308
func Fn1308(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1314 github.com/goccy/spidermonkeywasm2go/p6.Fn1314
func Fn1314(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1315 github.com/goccy/spidermonkeywasm2go/p6.Fn1315
func Fn1315(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1358 github.com/goccy/spidermonkeywasm2go/p6.Fn1358
func Fn1358(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1359 github.com/goccy/spidermonkeywasm2go/p6.Fn1359
func Fn1359(m *base.Module, l0 int32) int32

//go:linkname Fn1388 github.com/goccy/spidermonkeywasm2go/p6.Fn1388
func Fn1388(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1389 github.com/goccy/spidermonkeywasm2go/p6.Fn1389
func Fn1389(m *base.Module, l0 int32) int32

//go:linkname Fn1408 github.com/goccy/spidermonkeywasm2go/p7.Fn1408
func Fn1408(m *base.Module, l0 int32) int32

//go:linkname Fn1414 github.com/goccy/spidermonkeywasm2go/p0.Fn1414
func Fn1414(m *base.Module, l0 int32) int32

//go:linkname Fn1424 github.com/goccy/spidermonkeywasm2go/p7.Fn1424
func Fn1424(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1436 github.com/goccy/spidermonkeywasm2go/p0.Fn1436
func Fn1436(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1443 github.com/goccy/spidermonkeywasm2go/p0.Fn1443
func Fn1443(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1460 github.com/goccy/spidermonkeywasm2go/p5.Fn1460
func Fn1460(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1461 github.com/goccy/spidermonkeywasm2go/p6.Fn1461
func Fn1461(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1463 github.com/goccy/spidermonkeywasm2go/p6.Fn1463
func Fn1463(m *base.Module, l0 int32)

//go:linkname Fn1464 github.com/goccy/spidermonkeywasm2go/p5.Fn1464
func Fn1464(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1465 github.com/goccy/spidermonkeywasm2go/p6.Fn1465
func Fn1465(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1466 github.com/goccy/spidermonkeywasm2go/p7.Fn1466
func Fn1466(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1467 github.com/goccy/spidermonkeywasm2go/p5.Fn1467
func Fn1467(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1472 github.com/goccy/spidermonkeywasm2go/p0.Fn1472
func Fn1472(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1473 github.com/goccy/spidermonkeywasm2go/p0.Fn1473
func Fn1473(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1475 github.com/goccy/spidermonkeywasm2go/p7.Fn1475
func Fn1475(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1476 github.com/goccy/spidermonkeywasm2go/p7.Fn1476
func Fn1476(m *base.Module, l0 int32)

//go:linkname Fn1478 github.com/goccy/spidermonkeywasm2go/p7.Fn1478
func Fn1478(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1479 github.com/goccy/spidermonkeywasm2go/p7.Fn1479
func Fn1479(m *base.Module, l0 int32)

//go:linkname Fn1486 github.com/goccy/spidermonkeywasm2go/p7.Fn1486
func Fn1486(m *base.Module, l0 int32)

//go:linkname Fn1487 github.com/goccy/spidermonkeywasm2go/p7.Fn1487
func Fn1487(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1491 github.com/goccy/spidermonkeywasm2go/p6.Fn1491
func Fn1491(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1492 github.com/goccy/spidermonkeywasm2go/p7.Fn1492
func Fn1492(m *base.Module, l0 int32)

//go:linkname Fn1493 github.com/goccy/spidermonkeywasm2go/p6.Fn1493
func Fn1493(m *base.Module, l0 int32) int32

//go:linkname Fn1496 github.com/goccy/spidermonkeywasm2go/p7.Fn1496
func Fn1496(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1497 github.com/goccy/spidermonkeywasm2go/p7.Fn1497
func Fn1497(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1504 github.com/goccy/spidermonkeywasm2go/p7.Fn1504
func Fn1504(m *base.Module, l0 int32)

//go:linkname Fn1509 github.com/goccy/spidermonkeywasm2go/p7.Fn1509
func Fn1509(m *base.Module, l0 int32) int32

//go:linkname Fn1510 github.com/goccy/spidermonkeywasm2go/p5.Fn1510
func Fn1510(m *base.Module, l0 float64, l1 float64, l2 float64) float64

//go:linkname Fn1512 github.com/goccy/spidermonkeywasm2go/p5.Fn1512
func Fn1512(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn1516 github.com/goccy/spidermonkeywasm2go/p6.Fn1516
func Fn1516(m *base.Module, l0 int32, l1 float64) float64

//go:linkname Fn1526 github.com/goccy/spidermonkeywasm2go/p7.Fn1526
func Fn1526(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1527 github.com/goccy/spidermonkeywasm2go/p7.Fn1527
func Fn1527(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1528 github.com/goccy/spidermonkeywasm2go/p6.Fn1528
func Fn1528(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1529 github.com/goccy/spidermonkeywasm2go/p6.Fn1529
func Fn1529(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1620 github.com/goccy/spidermonkeywasm2go/p0.Fn1620
func Fn1620(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1627 github.com/goccy/spidermonkeywasm2go/p7.Fn1627
func Fn1627(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1643 github.com/goccy/spidermonkeywasm2go/p7.Fn1643
func Fn1643(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1653 github.com/goccy/spidermonkeywasm2go/p6.Fn1653
func Fn1653(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn1655 github.com/goccy/spidermonkeywasm2go/p5.Fn1655
func Fn1655(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn1657 github.com/goccy/spidermonkeywasm2go/p7.Fn1657
func Fn1657(m *base.Module, l0 int32) float64

//go:linkname Fn1658 github.com/goccy/spidermonkeywasm2go/p7.Fn1658
func Fn1658(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1660 github.com/goccy/spidermonkeywasm2go/p0.Fn1660
func Fn1660(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1663 github.com/goccy/spidermonkeywasm2go/p0.Fn1663
func Fn1663(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1760 github.com/goccy/spidermonkeywasm2go/p7.Fn1760
func Fn1760(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1762 github.com/goccy/spidermonkeywasm2go/p7.Fn1762
func Fn1762(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1805 github.com/goccy/spidermonkeywasm2go/p7.Fn1805
func Fn1805(m *base.Module, l0 int32) int32

//go:linkname Fn1806 github.com/goccy/spidermonkeywasm2go/p6.Fn1806
func Fn1806(m *base.Module, l0 int32) int32

//go:linkname Fn1807 github.com/goccy/spidermonkeywasm2go/p6.Fn1807
func Fn1807(m *base.Module, l0 int32) int32

//go:linkname Fn1817 github.com/goccy/spidermonkeywasm2go/p5.Fn1817
func Fn1817(m *base.Module, l0 int32)

//go:linkname Fn1849 github.com/goccy/spidermonkeywasm2go/p0.Fn1849
func Fn1849(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1870 github.com/goccy/spidermonkeywasm2go/p3.Fn1870
func Fn1870(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1887 github.com/goccy/spidermonkeywasm2go/p7.Fn1887
func Fn1887(m *base.Module, l0 int32) int32

//go:linkname Fn1890 github.com/goccy/spidermonkeywasm2go/p7.Fn1890
func Fn1890(m *base.Module, l0 int32) int32

//go:linkname Fn1891 github.com/goccy/spidermonkeywasm2go/p7.Fn1891
func Fn1891(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1933 github.com/goccy/spidermonkeywasm2go/p0.Fn1933
func Fn1933(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1940 github.com/goccy/spidermonkeywasm2go/p0.Fn1940
func Fn1940(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1961 github.com/goccy/spidermonkeywasm2go/p0.Fn1961
func Fn1961(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1973 github.com/goccy/spidermonkeywasm2go/p0.Fn1973
func Fn1973(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1974 github.com/goccy/spidermonkeywasm2go/p0.Fn1974
func Fn1974(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1977 github.com/goccy/spidermonkeywasm2go/p0.Fn1977
func Fn1977(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1981 github.com/goccy/spidermonkeywasm2go/p0.Fn1981
func Fn1981(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1989 github.com/goccy/spidermonkeywasm2go/p0.Fn1989
func Fn1989(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1992 github.com/goccy/spidermonkeywasm2go/p6.Fn1992
func Fn1992(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn1993 github.com/goccy/spidermonkeywasm2go/p0.Fn1993
func Fn1993(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1994 github.com/goccy/spidermonkeywasm2go/p5.Fn1994
func Fn1994(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1995 github.com/goccy/spidermonkeywasm2go/p7.Fn1995
func Fn1995(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2040 github.com/goccy/spidermonkeywasm2go/p0.Fn2040
func Fn2040(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2058 github.com/goccy/spidermonkeywasm2go/p0.Fn2058
func Fn2058(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2071 github.com/goccy/spidermonkeywasm2go/p7.Fn2071
func Fn2071(m *base.Module, l0 int32) int32

//go:linkname Fn2073 github.com/goccy/spidermonkeywasm2go/p6.Fn2073
func Fn2073(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2078 github.com/goccy/spidermonkeywasm2go/p0.Fn2078
func Fn2078(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2082 github.com/goccy/spidermonkeywasm2go/p7.Fn2082
func Fn2082(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2105 github.com/goccy/spidermonkeywasm2go/p6.Fn2105
func Fn2105(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2126 github.com/goccy/spidermonkeywasm2go/p4.Fn2126
func Fn2126(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2138 github.com/goccy/spidermonkeywasm2go/p6.Fn2138
func Fn2138(m *base.Module, l0 int32)

//go:linkname Fn2197 github.com/goccy/spidermonkeywasm2go/p0.Fn2197
func Fn2197(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32

//go:linkname Fn2212 github.com/goccy/spidermonkeywasm2go/p5.Fn2212
func Fn2212(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2213 github.com/goccy/spidermonkeywasm2go/p4.Fn2213
func Fn2213(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2214 github.com/goccy/spidermonkeywasm2go/p5.Fn2214
func Fn2214(m *base.Module, l0 int32)

//go:linkname Fn2224 github.com/goccy/spidermonkeywasm2go/p6.Fn2224
func Fn2224(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2242 github.com/goccy/spidermonkeywasm2go/p6.Fn2242
func Fn2242(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2254 github.com/goccy/spidermonkeywasm2go/p6.Fn2254
func Fn2254(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2284 github.com/goccy/spidermonkeywasm2go/p6.Fn2284
func Fn2284(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2294 github.com/goccy/spidermonkeywasm2go/p6.Fn2294
func Fn2294(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2295 github.com/goccy/spidermonkeywasm2go/p5.Fn2295
func Fn2295(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2296 github.com/goccy/spidermonkeywasm2go/p7.Fn2296
func Fn2296(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2298 github.com/goccy/spidermonkeywasm2go/p0.Fn2298
func Fn2298(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2300 github.com/goccy/spidermonkeywasm2go/p0.Fn2300
func Fn2300(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn2303 github.com/goccy/spidermonkeywasm2go/p6.Fn2303
func Fn2303(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2313 github.com/goccy/spidermonkeywasm2go/p0.Fn2313
func Fn2313(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2315 github.com/goccy/spidermonkeywasm2go/p7.Fn2315
func Fn2315(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2318 github.com/goccy/spidermonkeywasm2go/p0.Fn2318
func Fn2318(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2319 github.com/goccy/spidermonkeywasm2go/p6.Fn2319
func Fn2319(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2324 github.com/goccy/spidermonkeywasm2go/p0.Fn2324
func Fn2324(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2342 github.com/goccy/spidermonkeywasm2go/p5.Fn2342
func Fn2342(m *base.Module, l0 int32)

//go:linkname Fn2343 github.com/goccy/spidermonkeywasm2go/p7.Fn2343
func Fn2343(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2394 github.com/goccy/spidermonkeywasm2go/p6.Fn2394
func Fn2394(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2395 github.com/goccy/spidermonkeywasm2go/p6.Fn2395
func Fn2395(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2399 github.com/goccy/spidermonkeywasm2go/p4.Fn2399
func Fn2399(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2447 github.com/goccy/spidermonkeywasm2go/p0.Fn2447
func Fn2447(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2465 github.com/goccy/spidermonkeywasm2go/p0.Fn2465
func Fn2465(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2468 github.com/goccy/spidermonkeywasm2go/p6.Fn2468
func Fn2468(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2474 github.com/goccy/spidermonkeywasm2go/p4.Fn2474
func Fn2474(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2476 github.com/goccy/spidermonkeywasm2go/p6.Fn2476
func Fn2476(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2484 github.com/goccy/spidermonkeywasm2go/p4.Fn2484
func Fn2484(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2497 github.com/goccy/spidermonkeywasm2go/p2.Fn2497
func Fn2497(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2503 github.com/goccy/spidermonkeywasm2go/p5.Fn2503
func Fn2503(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2523 github.com/goccy/spidermonkeywasm2go/p0.Fn2523
func Fn2523(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2527 github.com/goccy/spidermonkeywasm2go/p6.Fn2527
func Fn2527(m *base.Module, l0 int32)

//go:linkname Fn2530 github.com/goccy/spidermonkeywasm2go/p0.Fn2530
func Fn2530(m *base.Module, l0 int32) int32

//go:linkname Fn2531 github.com/goccy/spidermonkeywasm2go/p7.Fn2531
func Fn2531(m *base.Module, l0 int32)

//go:linkname Fn2536 github.com/goccy/spidermonkeywasm2go/p7.Fn2536
func Fn2536(m *base.Module, l0 int32)

//go:linkname Fn2539 github.com/goccy/spidermonkeywasm2go/p5.Fn2539
func Fn2539(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2541 github.com/goccy/spidermonkeywasm2go/p0.Fn2541
func Fn2541(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2543 github.com/goccy/spidermonkeywasm2go/p0.Fn2543
func Fn2543(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2546 github.com/goccy/spidermonkeywasm2go/p0.Fn2546
func Fn2546(m *base.Module, l0 int32)

//go:linkname Fn2551 github.com/goccy/spidermonkeywasm2go/p0.Fn2551
func Fn2551(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2552 github.com/goccy/spidermonkeywasm2go/p0.Fn2552
func Fn2552(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2559 github.com/goccy/spidermonkeywasm2go/p5.Fn2559
func Fn2559(m *base.Module, l0 int32)

//go:linkname Fn2560 github.com/goccy/spidermonkeywasm2go/p5.Fn2560
func Fn2560(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2566 github.com/goccy/spidermonkeywasm2go/p0.Fn2566
func Fn2566(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2567 github.com/goccy/spidermonkeywasm2go/p0.Fn2567
func Fn2567(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2570 github.com/goccy/spidermonkeywasm2go/p0.Fn2570
func Fn2570(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2572 github.com/goccy/spidermonkeywasm2go/p0.Fn2572
func Fn2572(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2573 github.com/goccy/spidermonkeywasm2go/p0.Fn2573
func Fn2573(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2575 github.com/goccy/spidermonkeywasm2go/p3.Fn2575
func Fn2575(m *base.Module, l0 int32)

//go:linkname Fn2577 github.com/goccy/spidermonkeywasm2go/p0.Fn2577
func Fn2577(m *base.Module, l0 int32)

//go:linkname Fn2578 github.com/goccy/spidermonkeywasm2go/p0.Fn2578
func Fn2578(m *base.Module, l0 int32)

//go:linkname Fn2583 github.com/goccy/spidermonkeywasm2go/p6.Fn2583
func Fn2583(m *base.Module, l0 int32) int32

//go:linkname Fn2589 github.com/goccy/spidermonkeywasm2go/p0.Fn2589
func Fn2589(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2592 github.com/goccy/spidermonkeywasm2go/p4.Fn2592
func Fn2592(m *base.Module, l0 int32)

//go:linkname Fn2599 github.com/goccy/spidermonkeywasm2go/p6.Fn2599
func Fn2599(m *base.Module, l0 int32) int32

//go:linkname Fn2600 github.com/goccy/spidermonkeywasm2go/p0.Fn2600
func Fn2600(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2609 github.com/goccy/spidermonkeywasm2go/p0.Fn2609
func Fn2609(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2611 github.com/goccy/spidermonkeywasm2go/p0.Fn2611
func Fn2611(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2631 github.com/goccy/spidermonkeywasm2go/p6.Fn2631
func Fn2631(m *base.Module, l0 int32)

//go:linkname Fn2632 github.com/goccy/spidermonkeywasm2go/p0.Fn2632
func Fn2632(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2633 github.com/goccy/spidermonkeywasm2go/p7.Fn2633
func Fn2633(m *base.Module, l0 int32)

//go:linkname Fn2634 github.com/goccy/spidermonkeywasm2go/p0.Fn2634
func Fn2634(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2635 github.com/goccy/spidermonkeywasm2go/p0.Fn2635
func Fn2635(m *base.Module, l0 int32) int32

//go:linkname Fn2636 github.com/goccy/spidermonkeywasm2go/p0.Fn2636
func Fn2636(m *base.Module, l0 int32) int32

//go:linkname Fn2639 github.com/goccy/spidermonkeywasm2go/p0.Fn2639
func Fn2639(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2640 github.com/goccy/spidermonkeywasm2go/p0.Fn2640
func Fn2640(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2641 github.com/goccy/spidermonkeywasm2go/p0.Fn2641
func Fn2641(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2644 github.com/goccy/spidermonkeywasm2go/p7.Fn2644
func Fn2644(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2645 github.com/goccy/spidermonkeywasm2go/p3.Fn2645
func Fn2645(m *base.Module, l0 int32) int32

//go:linkname Fn2646 github.com/goccy/spidermonkeywasm2go/p2.Fn2646
func Fn2646(m *base.Module, l0 int32) int32

//go:linkname Fn2647 github.com/goccy/spidermonkeywasm2go/p6.Fn2647
func Fn2647(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2648 github.com/goccy/spidermonkeywasm2go/p3.Fn2648
func Fn2648(m *base.Module, l0 int32) int32

//go:linkname Fn2649 github.com/goccy/spidermonkeywasm2go/p2.Fn2649
func Fn2649(m *base.Module, l0 int32) int32

//go:linkname Fn2709 github.com/goccy/spidermonkeywasm2go/p0.Fn2709
func Fn2709(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2716 github.com/goccy/spidermonkeywasm2go/p0.Fn2716
func Fn2716(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2721 github.com/goccy/spidermonkeywasm2go/p0.Fn2721
func Fn2721(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2722 github.com/goccy/spidermonkeywasm2go/p0.Fn2722
func Fn2722(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2731 github.com/goccy/spidermonkeywasm2go/p4.Fn2731
func Fn2731(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2732 github.com/goccy/spidermonkeywasm2go/p4.Fn2732
func Fn2732(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2741 github.com/goccy/spidermonkeywasm2go/p0.Fn2741
func Fn2741(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2743 github.com/goccy/spidermonkeywasm2go/p0.Fn2743
func Fn2743(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2744 github.com/goccy/spidermonkeywasm2go/p0.Fn2744
func Fn2744(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2753 github.com/goccy/spidermonkeywasm2go/p4.Fn2753
func Fn2753(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2754 github.com/goccy/spidermonkeywasm2go/p6.Fn2754
func Fn2754(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2755 github.com/goccy/spidermonkeywasm2go/p0.Fn2755
func Fn2755(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2763 github.com/goccy/spidermonkeywasm2go/p6.Fn2763
func Fn2763(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2764 github.com/goccy/spidermonkeywasm2go/p7.Fn2764
func Fn2764(m *base.Module, l0 int32)

//go:linkname Fn2769 github.com/goccy/spidermonkeywasm2go/p6.Fn2769
func Fn2769(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2833 github.com/goccy/spidermonkeywasm2go/p5.Fn2833
func Fn2833(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2851 github.com/goccy/spidermonkeywasm2go/p0.Fn2851
func Fn2851(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2855 github.com/goccy/spidermonkeywasm2go/p5.Fn2855
func Fn2855(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2856 github.com/goccy/spidermonkeywasm2go/p5.Fn2856
func Fn2856(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2870 github.com/goccy/spidermonkeywasm2go/p0.Fn2870
func Fn2870(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2876 github.com/goccy/spidermonkeywasm2go/p0.Fn2876
func Fn2876(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2903 github.com/goccy/spidermonkeywasm2go/p6.Fn2903
func Fn2903(m *base.Module, l0 int32) int32

//go:linkname Fn2929 github.com/goccy/spidermonkeywasm2go/p0.Fn2929
func Fn2929(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2931 github.com/goccy/spidermonkeywasm2go/p0.Fn2931
func Fn2931(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2932 github.com/goccy/spidermonkeywasm2go/p0.Fn2932
func Fn2932(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2934 github.com/goccy/spidermonkeywasm2go/p4.Fn2934
func Fn2934(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2935 github.com/goccy/spidermonkeywasm2go/p0.Fn2935
func Fn2935(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2947 github.com/goccy/spidermonkeywasm2go/p0.Fn2947
func Fn2947(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2961 github.com/goccy/spidermonkeywasm2go/p0.Fn2961
func Fn2961(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn2993 github.com/goccy/spidermonkeywasm2go/p6.Fn2993
func Fn2993(m *base.Module, l0 int32)

//go:linkname Fn2995 github.com/goccy/spidermonkeywasm2go/p5.Fn2995
func Fn2995(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2999 github.com/goccy/spidermonkeywasm2go/p6.Fn2999
func Fn2999(m *base.Module, l0 int32) int32

//go:linkname Fn3034 github.com/goccy/spidermonkeywasm2go/p0.Fn3034
func Fn3034(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3035 github.com/goccy/spidermonkeywasm2go/p0.Fn3035
func Fn3035(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3036 github.com/goccy/spidermonkeywasm2go/p7.Fn3036
func Fn3036(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3040 github.com/goccy/spidermonkeywasm2go/p6.Fn3040
func Fn3040(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3149 github.com/goccy/spidermonkeywasm2go/p2.Fn3149
func Fn3149(m *base.Module, l0 int32) int32

//go:linkname Fn3151 github.com/goccy/spidermonkeywasm2go/p2.Fn3151
func Fn3151(m *base.Module, l0 int32)

//go:linkname Fn3152 github.com/goccy/spidermonkeywasm2go/p7.Fn3152
func Fn3152(m *base.Module, l0 int32)

//go:linkname Fn3153 github.com/goccy/spidermonkeywasm2go/p0.Fn3153
func Fn3153(m *base.Module, l0 int32) int32

//go:linkname Fn3156 github.com/goccy/spidermonkeywasm2go/p7.Fn3156
func Fn3156(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3159 github.com/goccy/spidermonkeywasm2go/p5.Fn3159
func Fn3159(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3166 github.com/goccy/spidermonkeywasm2go/p0.Fn3166
func Fn3166(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3179 github.com/goccy/spidermonkeywasm2go/p6.Fn3179
func Fn3179(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3180 github.com/goccy/spidermonkeywasm2go/p7.Fn3180
func Fn3180(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3181 github.com/goccy/spidermonkeywasm2go/p7.Fn3181
func Fn3181(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3183 github.com/goccy/spidermonkeywasm2go/p0.Fn3183
func Fn3183(m *base.Module, l0 int32) int32

//go:linkname Fn3194 github.com/goccy/spidermonkeywasm2go/p5.Fn3194
func Fn3194(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3205 github.com/goccy/spidermonkeywasm2go/p0.Fn3205
func Fn3205(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3207 github.com/goccy/spidermonkeywasm2go/p0.Fn3207
func Fn3207(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3209 github.com/goccy/spidermonkeywasm2go/p5.Fn3209
func Fn3209(m *base.Module, l0 int32)

//go:linkname Fn3212 github.com/goccy/spidermonkeywasm2go/p6.Fn3212
func Fn3212(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3213 github.com/goccy/spidermonkeywasm2go/p4.Fn3213
func Fn3213(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3228 github.com/goccy/spidermonkeywasm2go/p0.Fn3228
func Fn3228(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3239 github.com/goccy/spidermonkeywasm2go/p0.Fn3239
func Fn3239(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3256 github.com/goccy/spidermonkeywasm2go/p5.Fn3256
func Fn3256(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3457 github.com/goccy/spidermonkeywasm2go/p5.Fn3457
func Fn3457(m *base.Module, l0 int32) int32

//go:linkname Fn3458 github.com/goccy/spidermonkeywasm2go/p6.Fn3458
func Fn3458(m *base.Module, l0 float64) int32

//go:linkname Fn3459 github.com/goccy/spidermonkeywasm2go/p7.Fn3459
func Fn3459(m *base.Module, l0 float64) int32

//go:linkname Fn3460 github.com/goccy/spidermonkeywasm2go/p7.Fn3460
func Fn3460(m *base.Module, l0 float64) int32

//go:linkname Fn3461 github.com/goccy/spidermonkeywasm2go/p7.Fn3461
func Fn3461(m *base.Module, l0 float64) int32

//go:linkname Fn3462 github.com/goccy/spidermonkeywasm2go/p7.Fn3462
func Fn3462(m *base.Module, l0 float64) int32

//go:linkname Fn3480 github.com/goccy/spidermonkeywasm2go/p7.Fn3480
func Fn3480(m *base.Module, l0 int32)

//go:linkname Fn3492 github.com/goccy/spidermonkeywasm2go/p7.Fn3492
func Fn3492(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3496 github.com/goccy/spidermonkeywasm2go/p7.Fn3496
func Fn3496(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3522 github.com/goccy/spidermonkeywasm2go/p2.Fn3522
func Fn3522(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3523 github.com/goccy/spidermonkeywasm2go/p0.Fn3523
func Fn3523(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3525 github.com/goccy/spidermonkeywasm2go/p6.Fn3525
func Fn3525(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3528 github.com/goccy/spidermonkeywasm2go/p7.Fn3528
func Fn3528(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3529 github.com/goccy/spidermonkeywasm2go/p7.Fn3529
func Fn3529(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3530 github.com/goccy/spidermonkeywasm2go/p6.Fn3530
func Fn3530(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3531 github.com/goccy/spidermonkeywasm2go/p5.Fn3531
func Fn3531(m *base.Module, l0 int32) int32

//go:linkname Fn3532 github.com/goccy/spidermonkeywasm2go/p7.Fn3532
func Fn3532(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3535 github.com/goccy/spidermonkeywasm2go/p7.Fn3535
func Fn3535(m *base.Module, l0 int32) int32

//go:linkname Fn3536 github.com/goccy/spidermonkeywasm2go/p0.Fn3536
func Fn3536(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3539 github.com/goccy/spidermonkeywasm2go/p0.Fn3539
func Fn3539(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3540 github.com/goccy/spidermonkeywasm2go/p0.Fn3540
func Fn3540(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3550 github.com/goccy/spidermonkeywasm2go/p0.Fn3550
func Fn3550(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3562 github.com/goccy/spidermonkeywasm2go/p6.Fn3562
func Fn3562(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3564 github.com/goccy/spidermonkeywasm2go/p0.Fn3564
func Fn3564(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3568 github.com/goccy/spidermonkeywasm2go/p6.Fn3568
func Fn3568(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3572 github.com/goccy/spidermonkeywasm2go/p6.Fn3572
func Fn3572(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3574 github.com/goccy/spidermonkeywasm2go/p5.Fn3574
func Fn3574(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3575 github.com/goccy/spidermonkeywasm2go/p6.Fn3575
func Fn3575(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3576 github.com/goccy/spidermonkeywasm2go/p5.Fn3576
func Fn3576(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3579 github.com/goccy/spidermonkeywasm2go/p4.Fn3579
func Fn3579(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3582 github.com/goccy/spidermonkeywasm2go/p7.Fn3582
func Fn3582(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3583 github.com/goccy/spidermonkeywasm2go/p7.Fn3583
func Fn3583(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3584 github.com/goccy/spidermonkeywasm2go/p6.Fn3584
func Fn3584(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3585 github.com/goccy/spidermonkeywasm2go/p6.Fn3585
func Fn3585(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3586 github.com/goccy/spidermonkeywasm2go/p7.Fn3586
func Fn3586(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3587 github.com/goccy/spidermonkeywasm2go/p7.Fn3587
func Fn3587(m *base.Module, l0 int32, l1 float64) int32

//go:linkname Fn3592 github.com/goccy/spidermonkeywasm2go/p4.Fn3592
func Fn3592(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3593 github.com/goccy/spidermonkeywasm2go/p5.Fn3593
func Fn3593(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3594 github.com/goccy/spidermonkeywasm2go/p5.Fn3594
func Fn3594(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3595 github.com/goccy/spidermonkeywasm2go/p7.Fn3595
func Fn3595(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3596 github.com/goccy/spidermonkeywasm2go/p7.Fn3596
func Fn3596(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3597 github.com/goccy/spidermonkeywasm2go/p7.Fn3597
func Fn3597(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3598 github.com/goccy/spidermonkeywasm2go/p3.Fn3598
func Fn3598(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3602 github.com/goccy/spidermonkeywasm2go/p7.Fn3602
func Fn3602(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3603 github.com/goccy/spidermonkeywasm2go/p7.Fn3603
func Fn3603(m *base.Module, l0 int32) int32

//go:linkname Fn3642 github.com/goccy/spidermonkeywasm2go/p3.Fn3642
func Fn3642(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3651 github.com/goccy/spidermonkeywasm2go/p0.Fn3651
func Fn3651(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3652 github.com/goccy/spidermonkeywasm2go/p0.Fn3652
func Fn3652(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3684 github.com/goccy/spidermonkeywasm2go/p6.Fn3684
func Fn3684(m *base.Module, l0 int32)

//go:linkname Fn3709 github.com/goccy/spidermonkeywasm2go/p0.Fn3709
func Fn3709(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3722 github.com/goccy/spidermonkeywasm2go/p6.Fn3722
func Fn3722(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3736 github.com/goccy/spidermonkeywasm2go/p4.Fn3736
func Fn3736(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3740 github.com/goccy/spidermonkeywasm2go/p7.Fn3740
func Fn3740(m *base.Module, l0 float64) int32

//go:linkname Fn3754 github.com/goccy/spidermonkeywasm2go/p0.Fn3754
func Fn3754(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3755 github.com/goccy/spidermonkeywasm2go/p0.Fn3755
func Fn3755(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3756 github.com/goccy/spidermonkeywasm2go/p0.Fn3756
func Fn3756(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3757 github.com/goccy/spidermonkeywasm2go/p0.Fn3757
func Fn3757(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3758 github.com/goccy/spidermonkeywasm2go/p0.Fn3758
func Fn3758(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3759 github.com/goccy/spidermonkeywasm2go/p0.Fn3759
func Fn3759(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3760 github.com/goccy/spidermonkeywasm2go/p0.Fn3760
func Fn3760(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3761 github.com/goccy/spidermonkeywasm2go/p0.Fn3761
func Fn3761(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3762 github.com/goccy/spidermonkeywasm2go/p0.Fn3762
func Fn3762(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3763 github.com/goccy/spidermonkeywasm2go/p0.Fn3763
func Fn3763(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3764 github.com/goccy/spidermonkeywasm2go/p0.Fn3764
func Fn3764(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3765 github.com/goccy/spidermonkeywasm2go/p0.Fn3765
func Fn3765(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3766 github.com/goccy/spidermonkeywasm2go/p7.Fn3766
func Fn3766(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3767 github.com/goccy/spidermonkeywasm2go/p7.Fn3767
func Fn3767(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3768 github.com/goccy/spidermonkeywasm2go/p7.Fn3768
func Fn3768(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3769 github.com/goccy/spidermonkeywasm2go/p7.Fn3769
func Fn3769(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3770 github.com/goccy/spidermonkeywasm2go/p7.Fn3770
func Fn3770(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3771 github.com/goccy/spidermonkeywasm2go/p7.Fn3771
func Fn3771(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3772 github.com/goccy/spidermonkeywasm2go/p7.Fn3772
func Fn3772(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3773 github.com/goccy/spidermonkeywasm2go/p7.Fn3773
func Fn3773(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3774 github.com/goccy/spidermonkeywasm2go/p7.Fn3774
func Fn3774(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3775 github.com/goccy/spidermonkeywasm2go/p7.Fn3775
func Fn3775(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3776 github.com/goccy/spidermonkeywasm2go/p7.Fn3776
func Fn3776(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3777 github.com/goccy/spidermonkeywasm2go/p7.Fn3777
func Fn3777(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3778 github.com/goccy/spidermonkeywasm2go/p7.Fn3778
func Fn3778(m *base.Module, l0 int32) int32

//go:linkname Fn3779 github.com/goccy/spidermonkeywasm2go/p7.Fn3779
func Fn3779(m *base.Module, l0 int32) int32

//go:linkname Fn3780 github.com/goccy/spidermonkeywasm2go/p7.Fn3780
func Fn3780(m *base.Module, l0 int32) int32

//go:linkname Fn3781 github.com/goccy/spidermonkeywasm2go/p7.Fn3781
func Fn3781(m *base.Module, l0 int32) int32

//go:linkname Fn3782 github.com/goccy/spidermonkeywasm2go/p7.Fn3782
func Fn3782(m *base.Module, l0 int32) int32

//go:linkname Fn3783 github.com/goccy/spidermonkeywasm2go/p7.Fn3783
func Fn3783(m *base.Module, l0 int32) int32

//go:linkname Fn3784 github.com/goccy/spidermonkeywasm2go/p7.Fn3784
func Fn3784(m *base.Module, l0 int32) int32

//go:linkname Fn3785 github.com/goccy/spidermonkeywasm2go/p7.Fn3785
func Fn3785(m *base.Module, l0 int32) int32

//go:linkname Fn3786 github.com/goccy/spidermonkeywasm2go/p7.Fn3786
func Fn3786(m *base.Module, l0 int32) int32

//go:linkname Fn3787 github.com/goccy/spidermonkeywasm2go/p7.Fn3787
func Fn3787(m *base.Module, l0 int32) int32

//go:linkname Fn3788 github.com/goccy/spidermonkeywasm2go/p7.Fn3788
func Fn3788(m *base.Module, l0 int32) int32

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

//go:linkname Fn3843 github.com/goccy/spidermonkeywasm2go/p0.Fn3843
func Fn3843(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3844 github.com/goccy/spidermonkeywasm2go/p2.Fn3844
func Fn3844(m *base.Module, l0 int32) int32

//go:linkname Fn3852 github.com/goccy/spidermonkeywasm2go/p5.Fn3852
func Fn3852(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3866 github.com/goccy/spidermonkeywasm2go/p5.Fn3866
func Fn3866(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3868 github.com/goccy/spidermonkeywasm2go/p5.Fn3868
func Fn3868(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3870 github.com/goccy/spidermonkeywasm2go/p5.Fn3870
func Fn3870(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3872 github.com/goccy/spidermonkeywasm2go/p5.Fn3872
func Fn3872(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3874 github.com/goccy/spidermonkeywasm2go/p5.Fn3874
func Fn3874(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3876 github.com/goccy/spidermonkeywasm2go/p5.Fn3876
func Fn3876(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3878 github.com/goccy/spidermonkeywasm2go/p5.Fn3878
func Fn3878(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3880 github.com/goccy/spidermonkeywasm2go/p5.Fn3880
func Fn3880(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3882 github.com/goccy/spidermonkeywasm2go/p5.Fn3882
func Fn3882(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3884 github.com/goccy/spidermonkeywasm2go/p5.Fn3884
func Fn3884(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3886 github.com/goccy/spidermonkeywasm2go/p5.Fn3886
func Fn3886(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3888 github.com/goccy/spidermonkeywasm2go/p5.Fn3888
func Fn3888(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3896 github.com/goccy/spidermonkeywasm2go/p0.Fn3896
func Fn3896(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3898 github.com/goccy/spidermonkeywasm2go/p0.Fn3898
func Fn3898(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3900 github.com/goccy/spidermonkeywasm2go/p0.Fn3900
func Fn3900(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3902 github.com/goccy/spidermonkeywasm2go/p0.Fn3902
func Fn3902(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3907 github.com/goccy/spidermonkeywasm2go/p0.Fn3907
func Fn3907(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3909 github.com/goccy/spidermonkeywasm2go/p0.Fn3909
func Fn3909(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3911 github.com/goccy/spidermonkeywasm2go/p0.Fn3911
func Fn3911(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3916 github.com/goccy/spidermonkeywasm2go/p0.Fn3916
func Fn3916(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3936 github.com/goccy/spidermonkeywasm2go/p7.Fn3936
func Fn3936(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3939 github.com/goccy/spidermonkeywasm2go/p7.Fn3939
func Fn3939(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3942 github.com/goccy/spidermonkeywasm2go/p7.Fn3942
func Fn3942(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3945 github.com/goccy/spidermonkeywasm2go/p3.Fn3945
func Fn3945(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3946 github.com/goccy/spidermonkeywasm2go/p5.Fn3946
func Fn3946(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3947 github.com/goccy/spidermonkeywasm2go/p3.Fn3947
func Fn3947(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3948 github.com/goccy/spidermonkeywasm2go/p5.Fn3948
func Fn3948(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3949 github.com/goccy/spidermonkeywasm2go/p3.Fn3949
func Fn3949(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3950 github.com/goccy/spidermonkeywasm2go/p5.Fn3950
func Fn3950(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3951 github.com/goccy/spidermonkeywasm2go/p7.Fn3951
func Fn3951(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3954 github.com/goccy/spidermonkeywasm2go/p2.Fn3954
func Fn3954(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3955 github.com/goccy/spidermonkeywasm2go/p5.Fn3955
func Fn3955(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4025 github.com/goccy/spidermonkeywasm2go/p4.Fn4025
func Fn4025(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4026 github.com/goccy/spidermonkeywasm2go/p4.Fn4026
func Fn4026(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4104 github.com/goccy/spidermonkeywasm2go/p6.Fn4104
func Fn4104(m *base.Module, l0 int32) int32

//go:linkname Fn4105 github.com/goccy/spidermonkeywasm2go/p2.Fn4105
func Fn4105(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4106 github.com/goccy/spidermonkeywasm2go/p3.Fn4106
func Fn4106(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4107 github.com/goccy/spidermonkeywasm2go/p2.Fn4107
func Fn4107(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4108 github.com/goccy/spidermonkeywasm2go/p2.Fn4108
func Fn4108(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4128 github.com/goccy/spidermonkeywasm2go/p0.Fn4128
func Fn4128(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4185 github.com/goccy/spidermonkeywasm2go/p7.Fn4185
func Fn4185(m *base.Module, l0 int32)

//go:linkname Fn4186 github.com/goccy/spidermonkeywasm2go/p7.Fn4186
func Fn4186(m *base.Module, l0 int32)

//go:linkname Fn4216 github.com/goccy/spidermonkeywasm2go/p7.Fn4216
func Fn4216(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4227 github.com/goccy/spidermonkeywasm2go/p6.Fn4227
func Fn4227(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4229 github.com/goccy/spidermonkeywasm2go/p2.Fn4229
func Fn4229(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4231 github.com/goccy/spidermonkeywasm2go/p5.Fn4231
func Fn4231(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4241 github.com/goccy/spidermonkeywasm2go/p3.Fn4241
func Fn4241(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4242 github.com/goccy/spidermonkeywasm2go/p4.Fn4242
func Fn4242(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32) int32

//go:linkname Fn4243 github.com/goccy/spidermonkeywasm2go/p5.Fn4243
func Fn4243(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32) int32

//go:linkname Fn4245 github.com/goccy/spidermonkeywasm2go/p4.Fn4245
func Fn4245(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4252 github.com/goccy/spidermonkeywasm2go/p7.Fn4252
func Fn4252(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn4253 github.com/goccy/spidermonkeywasm2go/p7.Fn4253
func Fn4253(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4259 github.com/goccy/spidermonkeywasm2go/p5.Fn4259
func Fn4259(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4260 github.com/goccy/spidermonkeywasm2go/p6.Fn4260
func Fn4260(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4261 github.com/goccy/spidermonkeywasm2go/p5.Fn4261
func Fn4261(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4262 github.com/goccy/spidermonkeywasm2go/p4.Fn4262
func Fn4262(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4263 github.com/goccy/spidermonkeywasm2go/p6.Fn4263
func Fn4263(m *base.Module, l0 int32) int32

//go:linkname Fn4264 github.com/goccy/spidermonkeywasm2go/p4.Fn4264
func Fn4264(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4265 github.com/goccy/spidermonkeywasm2go/p6.Fn4265
func Fn4265(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4287 github.com/goccy/spidermonkeywasm2go/p6.Fn4287
func Fn4287(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4288 github.com/goccy/spidermonkeywasm2go/p7.Fn4288
func Fn4288(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4330 github.com/goccy/spidermonkeywasm2go/p4.Fn4330
func Fn4330(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4332 github.com/goccy/spidermonkeywasm2go/p5.Fn4332
func Fn4332(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4333 github.com/goccy/spidermonkeywasm2go/p5.Fn4333
func Fn4333(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4334 github.com/goccy/spidermonkeywasm2go/p6.Fn4334
func Fn4334(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4335 github.com/goccy/spidermonkeywasm2go/p5.Fn4335
func Fn4335(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4408 github.com/goccy/spidermonkeywasm2go/p5.Fn4408
func Fn4408(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4532 github.com/goccy/spidermonkeywasm2go/p2.Fn4532
func Fn4532(m *base.Module, l0 int32, l1 float64, l2 float64, l3 int32, l4 int32) int32

//go:linkname Fn4533 github.com/goccy/spidermonkeywasm2go/p7.Fn4533
func Fn4533(m *base.Module, l0 int32) int32

//go:linkname Fn4538 github.com/goccy/spidermonkeywasm2go/p6.Fn4538
func Fn4538(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4540 github.com/goccy/spidermonkeywasm2go/p4.Fn4540
func Fn4540(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn4542 github.com/goccy/spidermonkeywasm2go/p3.Fn4542
func Fn4542(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4543 github.com/goccy/spidermonkeywasm2go/p6.Fn4543
func Fn4543(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4547 github.com/goccy/spidermonkeywasm2go/p2.Fn4547
func Fn4547(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4557 github.com/goccy/spidermonkeywasm2go/p4.Fn4557
func Fn4557(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4561 github.com/goccy/spidermonkeywasm2go/p5.Fn4561
func Fn4561(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4564 github.com/goccy/spidermonkeywasm2go/p6.Fn4564
func Fn4564(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4566 github.com/goccy/spidermonkeywasm2go/p6.Fn4566
func Fn4566(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4569 github.com/goccy/spidermonkeywasm2go/p7.Fn4569
func Fn4569(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4583 github.com/goccy/spidermonkeywasm2go/p5.Fn4583
func Fn4583(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn4648 github.com/goccy/spidermonkeywasm2go/p5.Fn4648
func Fn4648(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4656 github.com/goccy/spidermonkeywasm2go/p4.Fn4656
func Fn4656(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4658 github.com/goccy/spidermonkeywasm2go/p6.Fn4658
func Fn4658(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4659 github.com/goccy/spidermonkeywasm2go/p3.Fn4659
func Fn4659(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4668 github.com/goccy/spidermonkeywasm2go/p6.Fn4668
func Fn4668(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4671 github.com/goccy/spidermonkeywasm2go/p0.Fn4671
func Fn4671(m *base.Module, l0 int32)

//go:linkname Fn4688 github.com/goccy/spidermonkeywasm2go/p0.Fn4688
func Fn4688(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4691 github.com/goccy/spidermonkeywasm2go/p6.Fn4691
func Fn4691(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4706 github.com/goccy/spidermonkeywasm2go/p4.Fn4706
func Fn4706(m *base.Module, l0 int32)

//go:linkname Fn4708 github.com/goccy/spidermonkeywasm2go/p7.Fn4708
func Fn4708(m *base.Module, l0 int32)

//go:linkname Fn4709 github.com/goccy/spidermonkeywasm2go/p6.Fn4709
func Fn4709(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4710 github.com/goccy/spidermonkeywasm2go/p7.Fn4710
func Fn4710(m *base.Module, l0 int32) int32

//go:linkname Fn4711 github.com/goccy/spidermonkeywasm2go/p7.Fn4711
func Fn4711(m *base.Module, l0 int32) int32

//go:linkname Fn4712 github.com/goccy/spidermonkeywasm2go/p7.Fn4712
func Fn4712(m *base.Module, l0 int32) int32

//go:linkname Fn4713 github.com/goccy/spidermonkeywasm2go/p7.Fn4713
func Fn4713(m *base.Module, l0 int32) int32

//go:linkname Fn4714 github.com/goccy/spidermonkeywasm2go/p7.Fn4714
func Fn4714(m *base.Module, l0 int32) int32

//go:linkname Fn4715 github.com/goccy/spidermonkeywasm2go/p7.Fn4715
func Fn4715(m *base.Module, l0 int32) int32

//go:linkname Fn4716 github.com/goccy/spidermonkeywasm2go/p7.Fn4716
func Fn4716(m *base.Module, l0 int32) int32

//go:linkname Fn4790 github.com/goccy/spidermonkeywasm2go/p5.Fn4790
func Fn4790(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4791 github.com/goccy/spidermonkeywasm2go/p6.Fn4791
func Fn4791(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4792 github.com/goccy/spidermonkeywasm2go/p6.Fn4792
func Fn4792(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4793 github.com/goccy/spidermonkeywasm2go/p6.Fn4793
func Fn4793(m *base.Module, l0 int32)

//go:linkname Fn4795 github.com/goccy/spidermonkeywasm2go/p0.Fn4795
func Fn4795(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4812 github.com/goccy/spidermonkeywasm2go/p7.Fn4812
func Fn4812(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4813 github.com/goccy/spidermonkeywasm2go/p6.Fn4813
func Fn4813(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4814 github.com/goccy/spidermonkeywasm2go/p0.Fn4814
func Fn4814(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4817 github.com/goccy/spidermonkeywasm2go/p7.Fn4817
func Fn4817(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4857 github.com/goccy/spidermonkeywasm2go/p0.Fn4857
func Fn4857(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4899 github.com/goccy/spidermonkeywasm2go/p5.Fn4899
func Fn4899(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4900 github.com/goccy/spidermonkeywasm2go/p5.Fn4900
func Fn4900(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn4903 github.com/goccy/spidermonkeywasm2go/p0.Fn4903
func Fn4903(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5124 github.com/goccy/spidermonkeywasm2go/p6.Fn5124
func Fn5124(m *base.Module, l0 int32) int32

//go:linkname Fn5137 github.com/goccy/spidermonkeywasm2go/p6.Fn5137
func Fn5137(m *base.Module, l0 int32) int32

//go:linkname Fn5138 github.com/goccy/spidermonkeywasm2go/p3.Fn5138
func Fn5138(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5139 github.com/goccy/spidermonkeywasm2go/p5.Fn5139
func Fn5139(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5140 github.com/goccy/spidermonkeywasm2go/p0.Fn5140
func Fn5140(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5142 github.com/goccy/spidermonkeywasm2go/p0.Fn5142
func Fn5142(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5143 github.com/goccy/spidermonkeywasm2go/p7.Fn5143
func Fn5143(m *base.Module, l0 int32)

//go:linkname Fn5273 github.com/goccy/spidermonkeywasm2go/p0.Fn5273
func Fn5273(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5275 github.com/goccy/spidermonkeywasm2go/p5.Fn5275
func Fn5275(m *base.Module, l0 int32)

//go:linkname Fn5276 github.com/goccy/spidermonkeywasm2go/p0.Fn5276
func Fn5276(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5316 github.com/goccy/spidermonkeywasm2go/p0.Fn5316
func Fn5316(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn5321 github.com/goccy/spidermonkeywasm2go/p7.Fn5321
func Fn5321(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5327 github.com/goccy/spidermonkeywasm2go/p6.Fn5327
func Fn5327(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn5328 github.com/goccy/spidermonkeywasm2go/p6.Fn5328
func Fn5328(m *base.Module, l0 int32) int32

//go:linkname Fn5337 github.com/goccy/spidermonkeywasm2go/p6.Fn5337
func Fn5337(m *base.Module, l0 int32) int32

//go:linkname Fn5341 github.com/goccy/spidermonkeywasm2go/p5.Fn5341
func Fn5341(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5342 github.com/goccy/spidermonkeywasm2go/p5.Fn5342
func Fn5342(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5348 github.com/goccy/spidermonkeywasm2go/p6.Fn5348
func Fn5348(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5352 github.com/goccy/spidermonkeywasm2go/p4.Fn5352
func Fn5352(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5356 github.com/goccy/spidermonkeywasm2go/p7.Fn5356
func Fn5356(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5357 github.com/goccy/spidermonkeywasm2go/p5.Fn5357
func Fn5357(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5376 github.com/goccy/spidermonkeywasm2go/p6.Fn5376
func Fn5376(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5391 github.com/goccy/spidermonkeywasm2go/p6.Fn5391
func Fn5391(m *base.Module, l0 int32) int32

//go:linkname Fn5392 github.com/goccy/spidermonkeywasm2go/p7.Fn5392
func Fn5392(m *base.Module, l0 int32) int32

//go:linkname Fn5393 github.com/goccy/spidermonkeywasm2go/p5.Fn5393
func Fn5393(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5394 github.com/goccy/spidermonkeywasm2go/p6.Fn5394
func Fn5394(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5395 github.com/goccy/spidermonkeywasm2go/p6.Fn5395
func Fn5395(m *base.Module, l0 int32) int32

//go:linkname Fn5396 github.com/goccy/spidermonkeywasm2go/p6.Fn5396
func Fn5396(m *base.Module, l0 int32) int32

//go:linkname Fn5397 github.com/goccy/spidermonkeywasm2go/p6.Fn5397
func Fn5397(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5398 github.com/goccy/spidermonkeywasm2go/p6.Fn5398
func Fn5398(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5399 github.com/goccy/spidermonkeywasm2go/p5.Fn5399
func Fn5399(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5400 github.com/goccy/spidermonkeywasm2go/p5.Fn5400
func Fn5400(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5402 github.com/goccy/spidermonkeywasm2go/p6.Fn5402
func Fn5402(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5405 github.com/goccy/spidermonkeywasm2go/p3.Fn5405
func Fn5405(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5410 github.com/goccy/spidermonkeywasm2go/p5.Fn5410
func Fn5410(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5418 github.com/goccy/spidermonkeywasm2go/p6.Fn5418
func Fn5418(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5427 github.com/goccy/spidermonkeywasm2go/p5.Fn5427
func Fn5427(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5497 github.com/goccy/spidermonkeywasm2go/p0.Fn5497
func Fn5497(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5538 github.com/goccy/spidermonkeywasm2go/p6.Fn5538
func Fn5538(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5539 github.com/goccy/spidermonkeywasm2go/p7.Fn5539
func Fn5539(m *base.Module, l0 int32) int32

//go:linkname Fn5557 github.com/goccy/spidermonkeywasm2go/p7.Fn5557
func Fn5557(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5559 github.com/goccy/spidermonkeywasm2go/p5.Fn5559
func Fn5559(m *base.Module, l0 int32) int32

//go:linkname Fn5563 github.com/goccy/spidermonkeywasm2go/p0.Fn5563
func Fn5563(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5564 github.com/goccy/spidermonkeywasm2go/p0.Fn5564
func Fn5564(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5565 github.com/goccy/spidermonkeywasm2go/p6.Fn5565
func Fn5565(m *base.Module, l0 int32)

//go:linkname Fn5570 github.com/goccy/spidermonkeywasm2go/p5.Fn5570
func Fn5570(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5592 github.com/goccy/spidermonkeywasm2go/p6.Fn5592
func Fn5592(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5596 github.com/goccy/spidermonkeywasm2go/p3.Fn5596
func Fn5596(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5609 github.com/goccy/spidermonkeywasm2go/p5.Fn5609
func Fn5609(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5610 github.com/goccy/spidermonkeywasm2go/p0.Fn5610
func Fn5610(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5641 github.com/goccy/spidermonkeywasm2go/p6.Fn5641
func Fn5641(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5643 github.com/goccy/spidermonkeywasm2go/p6.Fn5643
func Fn5643(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5644 github.com/goccy/spidermonkeywasm2go/p7.Fn5644
func Fn5644(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5645 github.com/goccy/spidermonkeywasm2go/p7.Fn5645
func Fn5645(m *base.Module, l0 int32)

//go:linkname Fn5647 github.com/goccy/spidermonkeywasm2go/p5.Fn5647
func Fn5647(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5651 github.com/goccy/spidermonkeywasm2go/p6.Fn5651
func Fn5651(m *base.Module, l0 int32)

//go:linkname Fn5655 github.com/goccy/spidermonkeywasm2go/p6.Fn5655
func Fn5655(m *base.Module, l0 int32)

//go:linkname Fn5658 github.com/goccy/spidermonkeywasm2go/p6.Fn5658
func Fn5658(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5660 github.com/goccy/spidermonkeywasm2go/p7.Fn5660
func Fn5660(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5679 github.com/goccy/spidermonkeywasm2go/p6.Fn5679
func Fn5679(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5682 github.com/goccy/spidermonkeywasm2go/p5.Fn5682
func Fn5682(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5694 github.com/goccy/spidermonkeywasm2go/p7.Fn5694
func Fn5694(m *base.Module, l0 int32) int32

//go:linkname Fn5726 github.com/goccy/spidermonkeywasm2go/p7.Fn5726
func Fn5726(m *base.Module, l0 int32)

//go:linkname Fn5739 github.com/goccy/spidermonkeywasm2go/p6.Fn5739
func Fn5739(m *base.Module, l0 int32)

//go:linkname Fn5742 github.com/goccy/spidermonkeywasm2go/p6.Fn5742
func Fn5742(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5809 github.com/goccy/spidermonkeywasm2go/p6.Fn5809
func Fn5809(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5810 github.com/goccy/spidermonkeywasm2go/p6.Fn5810
func Fn5810(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5811 github.com/goccy/spidermonkeywasm2go/p6.Fn5811
func Fn5811(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5812 github.com/goccy/spidermonkeywasm2go/p6.Fn5812
func Fn5812(m *base.Module, l0 int32)

//go:linkname Fn5826 github.com/goccy/spidermonkeywasm2go/p7.Fn5826
func Fn5826(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5827 github.com/goccy/spidermonkeywasm2go/p6.Fn5827
func Fn5827(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5828 github.com/goccy/spidermonkeywasm2go/p6.Fn5828
func Fn5828(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5836 github.com/goccy/spidermonkeywasm2go/p6.Fn5836
func Fn5836(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5839 github.com/goccy/spidermonkeywasm2go/p6.Fn5839
func Fn5839(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5840 github.com/goccy/spidermonkeywasm2go/p5.Fn5840
func Fn5840(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5841 github.com/goccy/spidermonkeywasm2go/p5.Fn5841
func Fn5841(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5842 github.com/goccy/spidermonkeywasm2go/p5.Fn5842
func Fn5842(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5843 github.com/goccy/spidermonkeywasm2go/p2.Fn5843
func Fn5843(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5844 github.com/goccy/spidermonkeywasm2go/p5.Fn5844
func Fn5844(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5845 github.com/goccy/spidermonkeywasm2go/p5.Fn5845
func Fn5845(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5846 github.com/goccy/spidermonkeywasm2go/p5.Fn5846
func Fn5846(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5847 github.com/goccy/spidermonkeywasm2go/p5.Fn5847
func Fn5847(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5848 github.com/goccy/spidermonkeywasm2go/p5.Fn5848
func Fn5848(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5849 github.com/goccy/spidermonkeywasm2go/p6.Fn5849
func Fn5849(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5850 github.com/goccy/spidermonkeywasm2go/p5.Fn5850
func Fn5850(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5851 github.com/goccy/spidermonkeywasm2go/p5.Fn5851
func Fn5851(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5852 github.com/goccy/spidermonkeywasm2go/p5.Fn5852
func Fn5852(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5854 github.com/goccy/spidermonkeywasm2go/p2.Fn5854
func Fn5854(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5855 github.com/goccy/spidermonkeywasm2go/p6.Fn5855
func Fn5855(m *base.Module, l0 int32)

//go:linkname Fn5864 github.com/goccy/spidermonkeywasm2go/p6.Fn5864
func Fn5864(m *base.Module, l0 int32)

//go:linkname Fn5871 github.com/goccy/spidermonkeywasm2go/p4.Fn5871
func Fn5871(m *base.Module, l0 int32)

//go:linkname Fn5874 github.com/goccy/spidermonkeywasm2go/p5.Fn5874
func Fn5874(m *base.Module, l0 int32) int32

//go:linkname Fn5931 github.com/goccy/spidermonkeywasm2go/p4.Fn5931
func Fn5931(m *base.Module, l0 int32)

//go:linkname Fn5932 github.com/goccy/spidermonkeywasm2go/p7.Fn5932
func Fn5932(m *base.Module, l0 int32)

//go:linkname Fn5942 github.com/goccy/spidermonkeywasm2go/p3.Fn5942
func Fn5942(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5943 github.com/goccy/spidermonkeywasm2go/p3.Fn5943
func Fn5943(m *base.Module, l0 int32) int32

//go:linkname Fn5944 github.com/goccy/spidermonkeywasm2go/p6.Fn5944
func Fn5944(m *base.Module, l0 int32) int32

//go:linkname Fn5947 github.com/goccy/spidermonkeywasm2go/p5.Fn5947
func Fn5947(m *base.Module, l0 int32)

//go:linkname Fn5948 github.com/goccy/spidermonkeywasm2go/p4.Fn5948
func Fn5948(m *base.Module, l0 int32)

//go:linkname Fn5957 github.com/goccy/spidermonkeywasm2go/p6.Fn5957
func Fn5957(m *base.Module, l0 int32)

//go:linkname Fn5960 github.com/goccy/spidermonkeywasm2go/p5.Fn5960
func Fn5960(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5966 github.com/goccy/spidermonkeywasm2go/p6.Fn5966
func Fn5966(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5968 github.com/goccy/spidermonkeywasm2go/p7.Fn5968
func Fn5968(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn6040 github.com/goccy/spidermonkeywasm2go/p6.Fn6040
func Fn6040(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6048 github.com/goccy/spidermonkeywasm2go/p7.Fn6048
func Fn6048(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6080 github.com/goccy/spidermonkeywasm2go/p7.Fn6080
func Fn6080(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6094 github.com/goccy/spidermonkeywasm2go/p7.Fn6094
func Fn6094(m *base.Module)

//go:linkname Fn6096 github.com/goccy/spidermonkeywasm2go/p6.Fn6096
func Fn6096(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6198 github.com/goccy/spidermonkeywasm2go/p6.Fn6198
func Fn6198(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6206 github.com/goccy/spidermonkeywasm2go/p7.Fn6206
func Fn6206(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6207 github.com/goccy/spidermonkeywasm2go/p6.Fn6207
func Fn6207(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6210 github.com/goccy/spidermonkeywasm2go/p6.Fn6210
func Fn6210(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6215 github.com/goccy/spidermonkeywasm2go/p6.Fn6215
func Fn6215(m *base.Module, l0 int32) int32

//go:linkname Fn6216 github.com/goccy/spidermonkeywasm2go/p5.Fn6216
func Fn6216(m *base.Module, l0 int32)

//go:linkname Fn6217 github.com/goccy/spidermonkeywasm2go/p5.Fn6217
func Fn6217(m *base.Module, l0 int32)

//go:linkname Fn6218 github.com/goccy/spidermonkeywasm2go/p6.Fn6218
func Fn6218(m *base.Module, l0 int32) int32

//go:linkname Fn6219 github.com/goccy/spidermonkeywasm2go/p6.Fn6219
func Fn6219(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6220 github.com/goccy/spidermonkeywasm2go/p2.Fn6220
func Fn6220(m *base.Module, l0 int32) int32

//go:linkname Fn6221 github.com/goccy/spidermonkeywasm2go/p4.Fn6221
func Fn6221(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6222 github.com/goccy/spidermonkeywasm2go/p3.Fn6222
func Fn6222(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn6223 github.com/goccy/spidermonkeywasm2go/p6.Fn6223
func Fn6223(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6224 github.com/goccy/spidermonkeywasm2go/p4.Fn6224
func Fn6224(m *base.Module, l0 int32) int32

//go:linkname Fn6225 github.com/goccy/spidermonkeywasm2go/p6.Fn6225
func Fn6225(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6226 github.com/goccy/spidermonkeywasm2go/p7.Fn6226
func Fn6226(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6227 github.com/goccy/spidermonkeywasm2go/p3.Fn6227
func Fn6227(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6228 github.com/goccy/spidermonkeywasm2go/p7.Fn6228
func Fn6228(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6229 github.com/goccy/spidermonkeywasm2go/p6.Fn6229
func Fn6229(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6230 github.com/goccy/spidermonkeywasm2go/p4.Fn6230
func Fn6230(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6231 github.com/goccy/spidermonkeywasm2go/p5.Fn6231
func Fn6231(m *base.Module, l0 int32)

//go:linkname Fn6232 github.com/goccy/spidermonkeywasm2go/p4.Fn6232
func Fn6232(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6233 github.com/goccy/spidermonkeywasm2go/p6.Fn6233
func Fn6233(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6234 github.com/goccy/spidermonkeywasm2go/p6.Fn6234
func Fn6234(m *base.Module, l0 int32) int32

//go:linkname Fn6235 github.com/goccy/spidermonkeywasm2go/p6.Fn6235
func Fn6235(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn6236 github.com/goccy/spidermonkeywasm2go/p5.Fn6236
func Fn6236(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6280 github.com/goccy/spidermonkeywasm2go/p5.Fn6280
func Fn6280(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6282 github.com/goccy/spidermonkeywasm2go/p3.Fn6282
func Fn6282(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6334 github.com/goccy/spidermonkeywasm2go/p2.Fn6334
func Fn6334(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6654 github.com/goccy/spidermonkeywasm2go/p7.Fn6654
func Fn6654(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6655 github.com/goccy/spidermonkeywasm2go/p6.Fn6655
func Fn6655(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6656 github.com/goccy/spidermonkeywasm2go/p6.Fn6656
func Fn6656(m *base.Module, l0 int32) int32

//go:linkname Fn6669 github.com/goccy/spidermonkeywasm2go/p6.Fn6669
func Fn6669(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6677 github.com/goccy/spidermonkeywasm2go/p7.Fn6677
func Fn6677(m *base.Module, l0 int32)

//go:linkname Fn6678 github.com/goccy/spidermonkeywasm2go/p7.Fn6678
func Fn6678(m *base.Module, l0 int32)

//go:linkname Fn6679 github.com/goccy/spidermonkeywasm2go/p6.Fn6679
func Fn6679(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn6681 github.com/goccy/spidermonkeywasm2go/p6.Fn6681
func Fn6681(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6682 github.com/goccy/spidermonkeywasm2go/p5.Fn6682
func Fn6682(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6683 github.com/goccy/spidermonkeywasm2go/p4.Fn6683
func Fn6683(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6684 github.com/goccy/spidermonkeywasm2go/p7.Fn6684
func Fn6684(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6685 github.com/goccy/spidermonkeywasm2go/p7.Fn6685
func Fn6685(m *base.Module, l0 int32)

//go:linkname Fn6686 github.com/goccy/spidermonkeywasm2go/p6.Fn6686
func Fn6686(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6687 github.com/goccy/spidermonkeywasm2go/p6.Fn6687
func Fn6687(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6688 github.com/goccy/spidermonkeywasm2go/p6.Fn6688
func Fn6688(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6689 github.com/goccy/spidermonkeywasm2go/p6.Fn6689
func Fn6689(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6690 github.com/goccy/spidermonkeywasm2go/p6.Fn6690
func Fn6690(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6691 github.com/goccy/spidermonkeywasm2go/p7.Fn6691
func Fn6691(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6692 github.com/goccy/spidermonkeywasm2go/p5.Fn6692
func Fn6692(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn6693 github.com/goccy/spidermonkeywasm2go/p6.Fn6693
func Fn6693(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6694 github.com/goccy/spidermonkeywasm2go/p4.Fn6694
func Fn6694(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6695 github.com/goccy/spidermonkeywasm2go/p6.Fn6695
func Fn6695(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn6696 github.com/goccy/spidermonkeywasm2go/p7.Fn6696
func Fn6696(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6697 github.com/goccy/spidermonkeywasm2go/p6.Fn6697
func Fn6697(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6698 github.com/goccy/spidermonkeywasm2go/p7.Fn6698
func Fn6698(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6699 github.com/goccy/spidermonkeywasm2go/p7.Fn6699
func Fn6699(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6700 github.com/goccy/spidermonkeywasm2go/p5.Fn6700
func Fn6700(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6701 github.com/goccy/spidermonkeywasm2go/p7.Fn6701
func Fn6701(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6702 github.com/goccy/spidermonkeywasm2go/p7.Fn6702
func Fn6702(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6703 github.com/goccy/spidermonkeywasm2go/p7.Fn6703
func Fn6703(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6704 github.com/goccy/spidermonkeywasm2go/p6.Fn6704
func Fn6704(m *base.Module, l0 int32) int32

//go:linkname Fn6705 github.com/goccy/spidermonkeywasm2go/p7.Fn6705
func Fn6705(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6706 github.com/goccy/spidermonkeywasm2go/p6.Fn6706
func Fn6706(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6707 github.com/goccy/spidermonkeywasm2go/p6.Fn6707
func Fn6707(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6708 github.com/goccy/spidermonkeywasm2go/p7.Fn6708
func Fn6708(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6710 github.com/goccy/spidermonkeywasm2go/p7.Fn6710
func Fn6710(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6712 github.com/goccy/spidermonkeywasm2go/p7.Fn6712
func Fn6712(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6713 github.com/goccy/spidermonkeywasm2go/p6.Fn6713
func Fn6713(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6715 github.com/goccy/spidermonkeywasm2go/p7.Fn6715
func Fn6715(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6722 github.com/goccy/spidermonkeywasm2go/p7.Fn6722
func Fn6722(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6723 github.com/goccy/spidermonkeywasm2go/p7.Fn6723
func Fn6723(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6724 github.com/goccy/spidermonkeywasm2go/p7.Fn6724
func Fn6724(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6726 github.com/goccy/spidermonkeywasm2go/p4.Fn6726
func Fn6726(m *base.Module, l0 int32) int32

//go:linkname Fn6727 github.com/goccy/spidermonkeywasm2go/p4.Fn6727
func Fn6727(m *base.Module, l0 int32) int32

//go:linkname Fn6728 github.com/goccy/spidermonkeywasm2go/p5.Fn6728
func Fn6728(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6729 github.com/goccy/spidermonkeywasm2go/p5.Fn6729
func Fn6729(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6730 github.com/goccy/spidermonkeywasm2go/p5.Fn6730
func Fn6730(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6731 github.com/goccy/spidermonkeywasm2go/p4.Fn6731
func Fn6731(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6732 github.com/goccy/spidermonkeywasm2go/p5.Fn6732
func Fn6732(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6733 github.com/goccy/spidermonkeywasm2go/p4.Fn6733
func Fn6733(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6734 github.com/goccy/spidermonkeywasm2go/p3.Fn6734
func Fn6734(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6735 github.com/goccy/spidermonkeywasm2go/p5.Fn6735
func Fn6735(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6736 github.com/goccy/spidermonkeywasm2go/p4.Fn6736
func Fn6736(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6737 github.com/goccy/spidermonkeywasm2go/p5.Fn6737
func Fn6737(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6738 github.com/goccy/spidermonkeywasm2go/p4.Fn6738
func Fn6738(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6739 github.com/goccy/spidermonkeywasm2go/p5.Fn6739
func Fn6739(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6740 github.com/goccy/spidermonkeywasm2go/p4.Fn6740
func Fn6740(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6741 github.com/goccy/spidermonkeywasm2go/p4.Fn6741
func Fn6741(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6742 github.com/goccy/spidermonkeywasm2go/p7.Fn6742
func Fn6742(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6743 github.com/goccy/spidermonkeywasm2go/p6.Fn6743
func Fn6743(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6744 github.com/goccy/spidermonkeywasm2go/p7.Fn6744
func Fn6744(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6745 github.com/goccy/spidermonkeywasm2go/p7.Fn6745
func Fn6745(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6746 github.com/goccy/spidermonkeywasm2go/p7.Fn6746
func Fn6746(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6747 github.com/goccy/spidermonkeywasm2go/p7.Fn6747
func Fn6747(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6751 github.com/goccy/spidermonkeywasm2go/p7.Fn6751
func Fn6751(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6754 github.com/goccy/spidermonkeywasm2go/p4.Fn6754
func Fn6754(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6755 github.com/goccy/spidermonkeywasm2go/p6.Fn6755
func Fn6755(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6756 github.com/goccy/spidermonkeywasm2go/p5.Fn6756
func Fn6756(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6757 github.com/goccy/spidermonkeywasm2go/p6.Fn6757
func Fn6757(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6758 github.com/goccy/spidermonkeywasm2go/p6.Fn6758
func Fn6758(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6759 github.com/goccy/spidermonkeywasm2go/p7.Fn6759
func Fn6759(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6760 github.com/goccy/spidermonkeywasm2go/p7.Fn6760
func Fn6760(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6761 github.com/goccy/spidermonkeywasm2go/p7.Fn6761
func Fn6761(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6762 github.com/goccy/spidermonkeywasm2go/p7.Fn6762
func Fn6762(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6766 github.com/goccy/spidermonkeywasm2go/p7.Fn6766
func Fn6766(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6767 github.com/goccy/spidermonkeywasm2go/p7.Fn6767
func Fn6767(m *base.Module, l0 int32) int32

//go:linkname Fn6768 github.com/goccy/spidermonkeywasm2go/p7.Fn6768
func Fn6768(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6769 github.com/goccy/spidermonkeywasm2go/p7.Fn6769
func Fn6769(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn6771 github.com/goccy/spidermonkeywasm2go/p7.Fn6771
func Fn6771(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6772 github.com/goccy/spidermonkeywasm2go/p7.Fn6772
func Fn6772(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6773 github.com/goccy/spidermonkeywasm2go/p4.Fn6773
func Fn6773(m *base.Module, l0 int32) int32

//go:linkname Fn6774 github.com/goccy/spidermonkeywasm2go/p6.Fn6774
func Fn6774(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6775 github.com/goccy/spidermonkeywasm2go/p7.Fn6775
func Fn6775(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6776 github.com/goccy/spidermonkeywasm2go/p7.Fn6776
func Fn6776(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6777 github.com/goccy/spidermonkeywasm2go/p7.Fn6777
func Fn6777(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6778 github.com/goccy/spidermonkeywasm2go/p7.Fn6778
func Fn6778(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6779 github.com/goccy/spidermonkeywasm2go/p7.Fn6779
func Fn6779(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6782 github.com/goccy/spidermonkeywasm2go/p5.Fn6782
func Fn6782(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6783 github.com/goccy/spidermonkeywasm2go/p5.Fn6783
func Fn6783(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6784 github.com/goccy/spidermonkeywasm2go/p6.Fn6784
func Fn6784(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6785 github.com/goccy/spidermonkeywasm2go/p6.Fn6785
func Fn6785(m *base.Module, l0 int32) int32

//go:linkname Fn6786 github.com/goccy/spidermonkeywasm2go/p6.Fn6786
func Fn6786(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6787 github.com/goccy/spidermonkeywasm2go/p7.Fn6787
func Fn6787(m *base.Module, l0 int32)

//go:linkname Fn6788 github.com/goccy/spidermonkeywasm2go/p7.Fn6788
func Fn6788(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6789 github.com/goccy/spidermonkeywasm2go/p6.Fn6789
func Fn6789(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6791 github.com/goccy/spidermonkeywasm2go/p7.Fn6791
func Fn6791(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6792 github.com/goccy/spidermonkeywasm2go/p7.Fn6792
func Fn6792(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6793 github.com/goccy/spidermonkeywasm2go/p7.Fn6793
func Fn6793(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6794 github.com/goccy/spidermonkeywasm2go/p7.Fn6794
func Fn6794(m *base.Module, l0 int32) int32

//go:linkname Fn6795 github.com/goccy/spidermonkeywasm2go/p6.Fn6795
func Fn6795(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6796 github.com/goccy/spidermonkeywasm2go/p6.Fn6796
func Fn6796(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6797 github.com/goccy/spidermonkeywasm2go/p7.Fn6797
func Fn6797(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6798 github.com/goccy/spidermonkeywasm2go/p4.Fn6798
func Fn6798(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6799 github.com/goccy/spidermonkeywasm2go/p6.Fn6799
func Fn6799(m *base.Module, l0 int32)

//go:linkname Fn6800 github.com/goccy/spidermonkeywasm2go/p7.Fn6800
func Fn6800(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6801 github.com/goccy/spidermonkeywasm2go/p7.Fn6801
func Fn6801(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6802 github.com/goccy/spidermonkeywasm2go/p7.Fn6802
func Fn6802(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6803 github.com/goccy/spidermonkeywasm2go/p7.Fn6803
func Fn6803(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6804 github.com/goccy/spidermonkeywasm2go/p6.Fn6804
func Fn6804(m *base.Module, l0 int32) int32

//go:linkname Fn6805 github.com/goccy/spidermonkeywasm2go/p6.Fn6805
func Fn6805(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6806 github.com/goccy/spidermonkeywasm2go/p7.Fn6806
func Fn6806(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6807 github.com/goccy/spidermonkeywasm2go/p7.Fn6807
func Fn6807(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6808 github.com/goccy/spidermonkeywasm2go/p7.Fn6808
func Fn6808(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6809 github.com/goccy/spidermonkeywasm2go/p7.Fn6809
func Fn6809(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6810 github.com/goccy/spidermonkeywasm2go/p7.Fn6810
func Fn6810(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6811 github.com/goccy/spidermonkeywasm2go/p7.Fn6811
func Fn6811(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6812 github.com/goccy/spidermonkeywasm2go/p7.Fn6812
func Fn6812(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6813 github.com/goccy/spidermonkeywasm2go/p7.Fn6813
func Fn6813(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6814 github.com/goccy/spidermonkeywasm2go/p7.Fn6814
func Fn6814(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6815 github.com/goccy/spidermonkeywasm2go/p7.Fn6815
func Fn6815(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6816 github.com/goccy/spidermonkeywasm2go/p6.Fn6816
func Fn6816(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6817 github.com/goccy/spidermonkeywasm2go/p7.Fn6817
func Fn6817(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6818 github.com/goccy/spidermonkeywasm2go/p7.Fn6818
func Fn6818(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6819 github.com/goccy/spidermonkeywasm2go/p7.Fn6819
func Fn6819(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6820 github.com/goccy/spidermonkeywasm2go/p6.Fn6820
func Fn6820(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6821 github.com/goccy/spidermonkeywasm2go/p7.Fn6821
func Fn6821(m *base.Module, l0 int32) int32

//go:linkname Fn6822 github.com/goccy/spidermonkeywasm2go/p7.Fn6822
func Fn6822(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6825 github.com/goccy/spidermonkeywasm2go/p5.Fn6825
func Fn6825(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6826 github.com/goccy/spidermonkeywasm2go/p3.Fn6826
func Fn6826(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6828 github.com/goccy/spidermonkeywasm2go/p6.Fn6828
func Fn6828(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6829 github.com/goccy/spidermonkeywasm2go/p5.Fn6829
func Fn6829(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6836 github.com/goccy/spidermonkeywasm2go/p5.Fn6836
func Fn6836(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6837 github.com/goccy/spidermonkeywasm2go/p7.Fn6837
func Fn6837(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6838 github.com/goccy/spidermonkeywasm2go/p7.Fn6838
func Fn6838(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6839 github.com/goccy/spidermonkeywasm2go/p7.Fn6839
func Fn6839(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6840 github.com/goccy/spidermonkeywasm2go/p7.Fn6840
func Fn6840(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6841 github.com/goccy/spidermonkeywasm2go/p7.Fn6841
func Fn6841(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6842 github.com/goccy/spidermonkeywasm2go/p7.Fn6842
func Fn6842(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6843 github.com/goccy/spidermonkeywasm2go/p7.Fn6843
func Fn6843(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6844 github.com/goccy/spidermonkeywasm2go/p7.Fn6844
func Fn6844(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6845 github.com/goccy/spidermonkeywasm2go/p7.Fn6845
func Fn6845(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6846 github.com/goccy/spidermonkeywasm2go/p6.Fn6846
func Fn6846(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6847 github.com/goccy/spidermonkeywasm2go/p6.Fn6847
func Fn6847(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6848 github.com/goccy/spidermonkeywasm2go/p6.Fn6848
func Fn6848(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6856 github.com/goccy/spidermonkeywasm2go/p7.Fn6856
func Fn6856(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6881 github.com/goccy/spidermonkeywasm2go/p7.Fn6881
func Fn6881(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6882 github.com/goccy/spidermonkeywasm2go/p5.Fn6882
func Fn6882(m *base.Module)

//go:linkname Fn6883 github.com/goccy/spidermonkeywasm2go/p7.Fn6883
func Fn6883(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6885 github.com/goccy/spidermonkeywasm2go/p3.Fn6885
func Fn6885(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6915 github.com/goccy/spidermonkeywasm2go/p6.Fn6915
func Fn6915(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6928 github.com/goccy/spidermonkeywasm2go/p6.Fn6928
func Fn6928(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6935 github.com/goccy/spidermonkeywasm2go/p7.Fn6935
func Fn6935(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6937 github.com/goccy/spidermonkeywasm2go/p0.Fn6937
func Fn6937(m *base.Module, l0 int32) int32

//go:linkname Fn6940 github.com/goccy/spidermonkeywasm2go/p0.Fn6940
func Fn6940(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6941 github.com/goccy/spidermonkeywasm2go/p0.Fn6941
func Fn6941(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6942 github.com/goccy/spidermonkeywasm2go/p0.Fn6942
func Fn6942(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6967 github.com/goccy/spidermonkeywasm2go/p7.Fn6967
func Fn6967(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6969 github.com/goccy/spidermonkeywasm2go/p7.Fn6969
func Fn6969(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6972 github.com/goccy/spidermonkeywasm2go/p6.Fn6972
func Fn6972(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6973 github.com/goccy/spidermonkeywasm2go/p4.Fn6973
func Fn6973(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7000 github.com/goccy/spidermonkeywasm2go/p7.Fn7000
func Fn7000(m *base.Module, l0 int32)

//go:linkname Fn7001 github.com/goccy/spidermonkeywasm2go/p7.Fn7001
func Fn7001(m *base.Module, l0 int32)

//go:linkname Fn7013 github.com/goccy/spidermonkeywasm2go/p5.Fn7013
func Fn7013(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7020 github.com/goccy/spidermonkeywasm2go/p7.Fn7020
func Fn7020(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7048 github.com/goccy/spidermonkeywasm2go/p6.Fn7048
func Fn7048(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7068 github.com/goccy/spidermonkeywasm2go/p6.Fn7068
func Fn7068(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7070 github.com/goccy/spidermonkeywasm2go/p4.Fn7070
func Fn7070(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7071 github.com/goccy/spidermonkeywasm2go/p2.Fn7071
func Fn7071(m *base.Module, l0 int32) int32

//go:linkname Fn7086 github.com/goccy/spidermonkeywasm2go/p5.Fn7086
func Fn7086(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7087 github.com/goccy/spidermonkeywasm2go/p5.Fn7087
func Fn7087(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7103 github.com/goccy/spidermonkeywasm2go/p6.Fn7103
func Fn7103(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7104 github.com/goccy/spidermonkeywasm2go/p6.Fn7104
func Fn7104(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7105 github.com/goccy/spidermonkeywasm2go/p6.Fn7105
func Fn7105(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7106 github.com/goccy/spidermonkeywasm2go/p7.Fn7106
func Fn7106(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7107 github.com/goccy/spidermonkeywasm2go/p6.Fn7107
func Fn7107(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7109 github.com/goccy/spidermonkeywasm2go/p5.Fn7109
func Fn7109(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7110 github.com/goccy/spidermonkeywasm2go/p5.Fn7110
func Fn7110(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn7117 github.com/goccy/spidermonkeywasm2go/p7.Fn7117
func Fn7117(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7118 github.com/goccy/spidermonkeywasm2go/p4.Fn7118
func Fn7118(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7128 github.com/goccy/spidermonkeywasm2go/p6.Fn7128
func Fn7128(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7130 github.com/goccy/spidermonkeywasm2go/p4.Fn7130
func Fn7130(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7132 github.com/goccy/spidermonkeywasm2go/p6.Fn7132
func Fn7132(m *base.Module, l0 int32) int32

//go:linkname Fn7133 github.com/goccy/spidermonkeywasm2go/p6.Fn7133
func Fn7133(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn7137 github.com/goccy/spidermonkeywasm2go/p5.Fn7137
func Fn7137(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn7138 github.com/goccy/spidermonkeywasm2go/p7.Fn7138
func Fn7138(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7165 github.com/goccy/spidermonkeywasm2go/p3.Fn7165
func Fn7165(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7166 github.com/goccy/spidermonkeywasm2go/p6.Fn7166
func Fn7166(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7167 github.com/goccy/spidermonkeywasm2go/p5.Fn7167
func Fn7167(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7168 github.com/goccy/spidermonkeywasm2go/p5.Fn7168
func Fn7168(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7169 github.com/goccy/spidermonkeywasm2go/p6.Fn7169
func Fn7169(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7170 github.com/goccy/spidermonkeywasm2go/p7.Fn7170
func Fn7170(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7171 github.com/goccy/spidermonkeywasm2go/p5.Fn7171
func Fn7171(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7172 github.com/goccy/spidermonkeywasm2go/p6.Fn7172
func Fn7172(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7173 github.com/goccy/spidermonkeywasm2go/p5.Fn7173
func Fn7173(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7174 github.com/goccy/spidermonkeywasm2go/p5.Fn7174
func Fn7174(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7175 github.com/goccy/spidermonkeywasm2go/p5.Fn7175
func Fn7175(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7176 github.com/goccy/spidermonkeywasm2go/p6.Fn7176
func Fn7176(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7177 github.com/goccy/spidermonkeywasm2go/p2.Fn7177
func Fn7177(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7181 github.com/goccy/spidermonkeywasm2go/p4.Fn7181
func Fn7181(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7182 github.com/goccy/spidermonkeywasm2go/p6.Fn7182
func Fn7182(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7183 github.com/goccy/spidermonkeywasm2go/p6.Fn7183
func Fn7183(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7187 github.com/goccy/spidermonkeywasm2go/p6.Fn7187
func Fn7187(m *base.Module, l0 int32) int32

//go:linkname Fn7189 github.com/goccy/spidermonkeywasm2go/p5.Fn7189
func Fn7189(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7191 github.com/goccy/spidermonkeywasm2go/p2.Fn7191
func Fn7191(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7205 github.com/goccy/spidermonkeywasm2go/p6.Fn7205
func Fn7205(m *base.Module, l0 int32) int32

//go:linkname Fn7206 github.com/goccy/spidermonkeywasm2go/p7.Fn7206
func Fn7206(m *base.Module, l0 int32)

//go:linkname Fn7207 github.com/goccy/spidermonkeywasm2go/p6.Fn7207
func Fn7207(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7214 github.com/goccy/spidermonkeywasm2go/p4.Fn7214
func Fn7214(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7215 github.com/goccy/spidermonkeywasm2go/p5.Fn7215
func Fn7215(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7217 github.com/goccy/spidermonkeywasm2go/p7.Fn7217
func Fn7217(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7222 github.com/goccy/spidermonkeywasm2go/p6.Fn7222
func Fn7222(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7240 github.com/goccy/spidermonkeywasm2go/p5.Fn7240
func Fn7240(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7255 github.com/goccy/spidermonkeywasm2go/p6.Fn7255
func Fn7255(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7264 github.com/goccy/spidermonkeywasm2go/p7.Fn7264
func Fn7264(m *base.Module, l0 int32)

//go:linkname Fn7335 github.com/goccy/spidermonkeywasm2go/p6.Fn7335
func Fn7335(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7352 github.com/goccy/spidermonkeywasm2go/p6.Fn7352
func Fn7352(m *base.Module, l0 int32) int32

//go:linkname Fn7357 github.com/goccy/spidermonkeywasm2go/p5.Fn7357
func Fn7357(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7361 github.com/goccy/spidermonkeywasm2go/p6.Fn7361
func Fn7361(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7362 github.com/goccy/spidermonkeywasm2go/p6.Fn7362
func Fn7362(m *base.Module, l0 int32)

//go:linkname Fn7363 github.com/goccy/spidermonkeywasm2go/p6.Fn7363
func Fn7363(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7366 github.com/goccy/spidermonkeywasm2go/p6.Fn7366
func Fn7366(m *base.Module, l0 int32) int32

//go:linkname Fn7368 github.com/goccy/spidermonkeywasm2go/p4.Fn7368
func Fn7368(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7371 github.com/goccy/spidermonkeywasm2go/p7.Fn7371
func Fn7371(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7372 github.com/goccy/spidermonkeywasm2go/p6.Fn7372
func Fn7372(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn7375 github.com/goccy/spidermonkeywasm2go/p6.Fn7375
func Fn7375(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7376 github.com/goccy/spidermonkeywasm2go/p6.Fn7376
func Fn7376(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7378 github.com/goccy/spidermonkeywasm2go/p6.Fn7378
func Fn7378(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7383 github.com/goccy/spidermonkeywasm2go/p4.Fn7383
func Fn7383(m *base.Module, l0 int32) int32

//go:linkname Fn7389 github.com/goccy/spidermonkeywasm2go/p6.Fn7389
func Fn7389(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64) int32

//go:linkname Fn7390 github.com/goccy/spidermonkeywasm2go/p5.Fn7390
func Fn7390(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7391 github.com/goccy/spidermonkeywasm2go/p7.Fn7391
func Fn7391(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7392 github.com/goccy/spidermonkeywasm2go/p7.Fn7392
func Fn7392(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7393 github.com/goccy/spidermonkeywasm2go/p7.Fn7393
func Fn7393(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7394 github.com/goccy/spidermonkeywasm2go/p7.Fn7394
func Fn7394(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7395 github.com/goccy/spidermonkeywasm2go/p7.Fn7395
func Fn7395(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7396 github.com/goccy/spidermonkeywasm2go/p6.Fn7396
func Fn7396(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7397 github.com/goccy/spidermonkeywasm2go/p6.Fn7397
func Fn7397(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn7398 github.com/goccy/spidermonkeywasm2go/p6.Fn7398
func Fn7398(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7399 github.com/goccy/spidermonkeywasm2go/p6.Fn7399
func Fn7399(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7400 github.com/goccy/spidermonkeywasm2go/p7.Fn7400
func Fn7400(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7401 github.com/goccy/spidermonkeywasm2go/p5.Fn7401
func Fn7401(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7402 github.com/goccy/spidermonkeywasm2go/p6.Fn7402
func Fn7402(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7403 github.com/goccy/spidermonkeywasm2go/p4.Fn7403
func Fn7403(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7404 github.com/goccy/spidermonkeywasm2go/p5.Fn7404
func Fn7404(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7405 github.com/goccy/spidermonkeywasm2go/p5.Fn7405
func Fn7405(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7406 github.com/goccy/spidermonkeywasm2go/p6.Fn7406
func Fn7406(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7407 github.com/goccy/spidermonkeywasm2go/p6.Fn7407
func Fn7407(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7408 github.com/goccy/spidermonkeywasm2go/p4.Fn7408
func Fn7408(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7409 github.com/goccy/spidermonkeywasm2go/p7.Fn7409
func Fn7409(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7410 github.com/goccy/spidermonkeywasm2go/p5.Fn7410
func Fn7410(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn7411 github.com/goccy/spidermonkeywasm2go/p5.Fn7411
func Fn7411(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7412 github.com/goccy/spidermonkeywasm2go/p4.Fn7412
func Fn7412(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7415 github.com/goccy/spidermonkeywasm2go/p6.Fn7415
func Fn7415(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7416 github.com/goccy/spidermonkeywasm2go/p6.Fn7416
func Fn7416(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7417 github.com/goccy/spidermonkeywasm2go/p5.Fn7417
func Fn7417(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7418 github.com/goccy/spidermonkeywasm2go/p6.Fn7418
func Fn7418(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7419 github.com/goccy/spidermonkeywasm2go/p3.Fn7419
func Fn7419(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7420 github.com/goccy/spidermonkeywasm2go/p5.Fn7420
func Fn7420(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7421 github.com/goccy/spidermonkeywasm2go/p5.Fn7421
func Fn7421(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7422 github.com/goccy/spidermonkeywasm2go/p6.Fn7422
func Fn7422(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7493 github.com/goccy/spidermonkeywasm2go/p2.Fn7493
func Fn7493(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8085 github.com/goccy/spidermonkeywasm2go/p7.Fn8085
func Fn8085(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8097 github.com/goccy/spidermonkeywasm2go/p5.Fn8097
func Fn8097(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8102 github.com/goccy/spidermonkeywasm2go/p6.Fn8102
func Fn8102(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8103 github.com/goccy/spidermonkeywasm2go/p6.Fn8103
func Fn8103(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8104 github.com/goccy/spidermonkeywasm2go/p7.Fn8104
func Fn8104(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8106 github.com/goccy/spidermonkeywasm2go/p6.Fn8106
func Fn8106(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8107 github.com/goccy/spidermonkeywasm2go/p6.Fn8107
func Fn8107(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8111 github.com/goccy/spidermonkeywasm2go/p3.Fn8111
func Fn8111(m *base.Module, l0 int32)

//go:linkname Fn8127 github.com/goccy/spidermonkeywasm2go/p6.Fn8127
func Fn8127(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8128 github.com/goccy/spidermonkeywasm2go/p6.Fn8128
func Fn8128(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8129 github.com/goccy/spidermonkeywasm2go/p6.Fn8129
func Fn8129(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8130 github.com/goccy/spidermonkeywasm2go/p6.Fn8130
func Fn8130(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8145 github.com/goccy/spidermonkeywasm2go/p6.Fn8145
func Fn8145(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8148 github.com/goccy/spidermonkeywasm2go/p7.Fn8148
func Fn8148(m *base.Module, l0 int32) int32

//go:linkname Fn8150 github.com/goccy/spidermonkeywasm2go/p7.Fn8150
func Fn8150(m *base.Module, l0 int32) int32

//go:linkname Fn8153 github.com/goccy/spidermonkeywasm2go/p6.Fn8153
func Fn8153(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8159 github.com/goccy/spidermonkeywasm2go/p2.Fn8159
func Fn8159(m *base.Module) int32

//go:linkname Fn8162 github.com/goccy/spidermonkeywasm2go/p6.Fn8162
func Fn8162(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8164 github.com/goccy/spidermonkeywasm2go/p3.Fn8164
func Fn8164(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8175 github.com/goccy/spidermonkeywasm2go/p4.Fn8175
func Fn8175(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8176 github.com/goccy/spidermonkeywasm2go/p7.Fn8176
func Fn8176(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8177 github.com/goccy/spidermonkeywasm2go/p7.Fn8177
func Fn8177(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8179 github.com/goccy/spidermonkeywasm2go/p7.Fn8179
func Fn8179(m *base.Module, l0 int32) int32

//go:linkname Fn8183 github.com/goccy/spidermonkeywasm2go/p6.Fn8183
func Fn8183(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32)

//go:linkname Fn8185 github.com/goccy/spidermonkeywasm2go/p5.Fn8185
func Fn8185(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8188 github.com/goccy/spidermonkeywasm2go/p5.Fn8188
func Fn8188(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8189 github.com/goccy/spidermonkeywasm2go/p4.Fn8189
func Fn8189(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn8190 github.com/goccy/spidermonkeywasm2go/p6.Fn8190
func Fn8190(m *base.Module, l0 int32) int32

//go:linkname Fn8194 github.com/goccy/spidermonkeywasm2go/p7.Fn8194
func Fn8194(m *base.Module, l0 int32) int32

//go:linkname Fn8196 github.com/goccy/spidermonkeywasm2go/p7.Fn8196
func Fn8196(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8198 github.com/goccy/spidermonkeywasm2go/p6.Fn8198
func Fn8198(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8199 github.com/goccy/spidermonkeywasm2go/p5.Fn8199
func Fn8199(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8201 github.com/goccy/spidermonkeywasm2go/p7.Fn8201
func Fn8201(m *base.Module, l0 int32) int32

//go:linkname Fn8202 github.com/goccy/spidermonkeywasm2go/p5.Fn8202
func Fn8202(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8203 github.com/goccy/spidermonkeywasm2go/p7.Fn8203
func Fn8203(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8262 github.com/goccy/spidermonkeywasm2go/p5.Fn8262
func Fn8262(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8285 github.com/goccy/spidermonkeywasm2go/p6.Fn8285
func Fn8285(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8315 github.com/goccy/spidermonkeywasm2go/p5.Fn8315
func Fn8315(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8316 github.com/goccy/spidermonkeywasm2go/p5.Fn8316
func Fn8316(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8328 github.com/goccy/spidermonkeywasm2go/p4.Fn8328
func Fn8328(m *base.Module, l0 int32) int32

//go:linkname Fn8340 github.com/goccy/spidermonkeywasm2go/p3.Fn8340
func Fn8340(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8348 github.com/goccy/spidermonkeywasm2go/p7.Fn8348
func Fn8348(m *base.Module, l0 int32)

//go:linkname Fn8350 github.com/goccy/spidermonkeywasm2go/p6.Fn8350
func Fn8350(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8351 github.com/goccy/spidermonkeywasm2go/p5.Fn8351
func Fn8351(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8352 github.com/goccy/spidermonkeywasm2go/p6.Fn8352
func Fn8352(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8353 github.com/goccy/spidermonkeywasm2go/p6.Fn8353
func Fn8353(m *base.Module, l0 int32)

//go:linkname Fn8355 github.com/goccy/spidermonkeywasm2go/p6.Fn8355
func Fn8355(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8356 github.com/goccy/spidermonkeywasm2go/p5.Fn8356
func Fn8356(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8357 github.com/goccy/spidermonkeywasm2go/p6.Fn8357
func Fn8357(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8369 github.com/goccy/spidermonkeywasm2go/p5.Fn8369
func Fn8369(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn8371 github.com/goccy/spidermonkeywasm2go/p6.Fn8371
func Fn8371(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8376 github.com/goccy/spidermonkeywasm2go/p5.Fn8376
func Fn8376(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8379 github.com/goccy/spidermonkeywasm2go/p6.Fn8379
func Fn8379(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8380 github.com/goccy/spidermonkeywasm2go/p3.Fn8380
func Fn8380(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn8381 github.com/goccy/spidermonkeywasm2go/p7.Fn8381
func Fn8381(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8382 github.com/goccy/spidermonkeywasm2go/p7.Fn8382
func Fn8382(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8383 github.com/goccy/spidermonkeywasm2go/p7.Fn8383
func Fn8383(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8385 github.com/goccy/spidermonkeywasm2go/p4.Fn8385
func Fn8385(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8386 github.com/goccy/spidermonkeywasm2go/p6.Fn8386
func Fn8386(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8398 github.com/goccy/spidermonkeywasm2go/p6.Fn8398
func Fn8398(m *base.Module, l0 int32) int32

//go:linkname Fn8399 github.com/goccy/spidermonkeywasm2go/p6.Fn8399
func Fn8399(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8401 github.com/goccy/spidermonkeywasm2go/p7.Fn8401
func Fn8401(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8403 github.com/goccy/spidermonkeywasm2go/p6.Fn8403
func Fn8403(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8404 github.com/goccy/spidermonkeywasm2go/p6.Fn8404
func Fn8404(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8405 github.com/goccy/spidermonkeywasm2go/p7.Fn8405
func Fn8405(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8406 github.com/goccy/spidermonkeywasm2go/p7.Fn8406
func Fn8406(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8407 github.com/goccy/spidermonkeywasm2go/p7.Fn8407
func Fn8407(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8408 github.com/goccy/spidermonkeywasm2go/p7.Fn8408
func Fn8408(m *base.Module, l0 int32) int32

//go:linkname Fn8409 github.com/goccy/spidermonkeywasm2go/p4.Fn8409
func Fn8409(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8410 github.com/goccy/spidermonkeywasm2go/p4.Fn8410
func Fn8410(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8411 github.com/goccy/spidermonkeywasm2go/p6.Fn8411
func Fn8411(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8412 github.com/goccy/spidermonkeywasm2go/p6.Fn8412
func Fn8412(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8422 github.com/goccy/spidermonkeywasm2go/p6.Fn8422
func Fn8422(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8425 github.com/goccy/spidermonkeywasm2go/p6.Fn8425
func Fn8425(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8427 github.com/goccy/spidermonkeywasm2go/p5.Fn8427
func Fn8427(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8430 github.com/goccy/spidermonkeywasm2go/p6.Fn8430
func Fn8430(m *base.Module, l0 int32) int32

//go:linkname Fn8436 github.com/goccy/spidermonkeywasm2go/p6.Fn8436
func Fn8436(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8452 github.com/goccy/spidermonkeywasm2go/p6.Fn8452
func Fn8452(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8456 github.com/goccy/spidermonkeywasm2go/p6.Fn8456
func Fn8456(m *base.Module, l0 int32) int32

//go:linkname Fn8466 github.com/goccy/spidermonkeywasm2go/p5.Fn8466
func Fn8466(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8482 github.com/goccy/spidermonkeywasm2go/p6.Fn8482
func Fn8482(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8483 github.com/goccy/spidermonkeywasm2go/p6.Fn8483
func Fn8483(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8484 github.com/goccy/spidermonkeywasm2go/p6.Fn8484
func Fn8484(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8489 github.com/goccy/spidermonkeywasm2go/p6.Fn8489
func Fn8489(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8490 github.com/goccy/spidermonkeywasm2go/p6.Fn8490
func Fn8490(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8502 github.com/goccy/spidermonkeywasm2go/p7.Fn8502
func Fn8502(m *base.Module, l0 int32)

//go:linkname Fn8506 github.com/goccy/spidermonkeywasm2go/p5.Fn8506
func Fn8506(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8510 github.com/goccy/spidermonkeywasm2go/p5.Fn8510
func Fn8510(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8511 github.com/goccy/spidermonkeywasm2go/p5.Fn8511
func Fn8511(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8512 github.com/goccy/spidermonkeywasm2go/p7.Fn8512
func Fn8512(m *base.Module, l0 int32)

//go:linkname Fn8515 github.com/goccy/spidermonkeywasm2go/p5.Fn8515
func Fn8515(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8516 github.com/goccy/spidermonkeywasm2go/p6.Fn8516
func Fn8516(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8517 github.com/goccy/spidermonkeywasm2go/p6.Fn8517
func Fn8517(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8518 github.com/goccy/spidermonkeywasm2go/p6.Fn8518
func Fn8518(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8519 github.com/goccy/spidermonkeywasm2go/p3.Fn8519
func Fn8519(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8520 github.com/goccy/spidermonkeywasm2go/p5.Fn8520
func Fn8520(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8522 github.com/goccy/spidermonkeywasm2go/p7.Fn8522
func Fn8522(m *base.Module, l0 int32) int32

//go:linkname Fn8524 github.com/goccy/spidermonkeywasm2go/p7.Fn8524
func Fn8524(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8526 github.com/goccy/spidermonkeywasm2go/p7.Fn8526
func Fn8526(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8530 github.com/goccy/spidermonkeywasm2go/p7.Fn8530
func Fn8530(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8532 github.com/goccy/spidermonkeywasm2go/p6.Fn8532
func Fn8532(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8540 github.com/goccy/spidermonkeywasm2go/p6.Fn8540
func Fn8540(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8541 github.com/goccy/spidermonkeywasm2go/p6.Fn8541
func Fn8541(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8544 github.com/goccy/spidermonkeywasm2go/p7.Fn8544
func Fn8544(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8546 github.com/goccy/spidermonkeywasm2go/p7.Fn8546
func Fn8546(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8567 github.com/goccy/spidermonkeywasm2go/p7.Fn8567
func Fn8567(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8613 github.com/goccy/spidermonkeywasm2go/p7.Fn8613
func Fn8613(m *base.Module, l0 int32)

//go:linkname Fn8614 github.com/goccy/spidermonkeywasm2go/p7.Fn8614
func Fn8614(m *base.Module, l0 int32)

//go:linkname Fn8615 github.com/goccy/spidermonkeywasm2go/p6.Fn8615
func Fn8615(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8616 github.com/goccy/spidermonkeywasm2go/p7.Fn8616
func Fn8616(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8617 github.com/goccy/spidermonkeywasm2go/p6.Fn8617
func Fn8617(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8623 github.com/goccy/spidermonkeywasm2go/p6.Fn8623
func Fn8623(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8625 github.com/goccy/spidermonkeywasm2go/p6.Fn8625
func Fn8625(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8627 github.com/goccy/spidermonkeywasm2go/p6.Fn8627
func Fn8627(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn8628 github.com/goccy/spidermonkeywasm2go/p6.Fn8628
func Fn8628(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8630 github.com/goccy/spidermonkeywasm2go/p6.Fn8630
func Fn8630(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8642 github.com/goccy/spidermonkeywasm2go/p2.Fn8642
func Fn8642(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn8646 github.com/goccy/spidermonkeywasm2go/p6.Fn8646
func Fn8646(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8647 github.com/goccy/spidermonkeywasm2go/p6.Fn8647
func Fn8647(m *base.Module, l0 int32)

//go:linkname Fn8653 github.com/goccy/spidermonkeywasm2go/p6.Fn8653
func Fn8653(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8662 github.com/goccy/spidermonkeywasm2go/p6.Fn8662
func Fn8662(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8668 github.com/goccy/spidermonkeywasm2go/p5.Fn8668
func Fn8668(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8669 github.com/goccy/spidermonkeywasm2go/p6.Fn8669
func Fn8669(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8670 github.com/goccy/spidermonkeywasm2go/p6.Fn8670
func Fn8670(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8671 github.com/goccy/spidermonkeywasm2go/p7.Fn8671
func Fn8671(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8672 github.com/goccy/spidermonkeywasm2go/p6.Fn8672
func Fn8672(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8679 github.com/goccy/spidermonkeywasm2go/p4.Fn8679
func Fn8679(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8693 github.com/goccy/spidermonkeywasm2go/p4.Fn8693
func Fn8693(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8697 github.com/goccy/spidermonkeywasm2go/p3.Fn8697
func Fn8697(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn8699 github.com/goccy/spidermonkeywasm2go/p5.Fn8699
func Fn8699(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8700 github.com/goccy/spidermonkeywasm2go/p6.Fn8700
func Fn8700(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8704 github.com/goccy/spidermonkeywasm2go/p6.Fn8704
func Fn8704(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8705 github.com/goccy/spidermonkeywasm2go/p5.Fn8705
func Fn8705(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8709 github.com/goccy/spidermonkeywasm2go/p7.Fn8709
func Fn8709(m *base.Module, l0 int32) int32

//go:linkname Fn8710 github.com/goccy/spidermonkeywasm2go/p7.Fn8710
func Fn8710(m *base.Module, l0 int32)

//go:linkname Fn8713 github.com/goccy/spidermonkeywasm2go/p6.Fn8713
func Fn8713(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8714 github.com/goccy/spidermonkeywasm2go/p6.Fn8714
func Fn8714(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8715 github.com/goccy/spidermonkeywasm2go/p6.Fn8715
func Fn8715(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8743 github.com/goccy/spidermonkeywasm2go/p7.Fn8743
func Fn8743(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8994 github.com/goccy/spidermonkeywasm2go/p4.Fn8994
func Fn8994(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8995 github.com/goccy/spidermonkeywasm2go/p6.Fn8995
func Fn8995(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9015 github.com/goccy/spidermonkeywasm2go/p5.Fn9015
func Fn9015(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9016 github.com/goccy/spidermonkeywasm2go/p7.Fn9016
func Fn9016(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9048 github.com/goccy/spidermonkeywasm2go/p6.Fn9048
func Fn9048(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9053 github.com/goccy/spidermonkeywasm2go/p3.Fn9053
func Fn9053(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9065 github.com/goccy/spidermonkeywasm2go/p4.Fn9065
func Fn9065(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9088 github.com/goccy/spidermonkeywasm2go/p2.Fn9088
func Fn9088(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9090 github.com/goccy/spidermonkeywasm2go/p4.Fn9090
func Fn9090(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9094 github.com/goccy/spidermonkeywasm2go/p4.Fn9094
func Fn9094(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9128 github.com/goccy/spidermonkeywasm2go/p6.Fn9128
func Fn9128(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9144 github.com/goccy/spidermonkeywasm2go/p5.Fn9144
func Fn9144(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9262 github.com/goccy/spidermonkeywasm2go/p3.Fn9262
func Fn9262(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9267 github.com/goccy/spidermonkeywasm2go/p6.Fn9267
func Fn9267(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn9268 github.com/goccy/spidermonkeywasm2go/p6.Fn9268
func Fn9268(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32)

//go:linkname Fn9285 github.com/goccy/spidermonkeywasm2go/p6.Fn9285
func Fn9285(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9286 github.com/goccy/spidermonkeywasm2go/p6.Fn9286
func Fn9286(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9287 github.com/goccy/spidermonkeywasm2go/p7.Fn9287
func Fn9287(m *base.Module, l0 int32) int32

//go:linkname Fn9288 github.com/goccy/spidermonkeywasm2go/p6.Fn9288
func Fn9288(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9289 github.com/goccy/spidermonkeywasm2go/p6.Fn9289
func Fn9289(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9293 github.com/goccy/spidermonkeywasm2go/p5.Fn9293
func Fn9293(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9325 github.com/goccy/spidermonkeywasm2go/p3.Fn9325
func Fn9325(m *base.Module, l0 int32) int32

//go:linkname Fn9336 github.com/goccy/spidermonkeywasm2go/p7.Fn9336
func Fn9336(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9339 github.com/goccy/spidermonkeywasm2go/p6.Fn9339
func Fn9339(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9343 github.com/goccy/spidermonkeywasm2go/p3.Fn9343
func Fn9343(m *base.Module, l0 int32)

//go:linkname Fn9344 github.com/goccy/spidermonkeywasm2go/p4.Fn9344
func Fn9344(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9345 github.com/goccy/spidermonkeywasm2go/p5.Fn9345
func Fn9345(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9350 github.com/goccy/spidermonkeywasm2go/p5.Fn9350
func Fn9350(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9353 github.com/goccy/spidermonkeywasm2go/p6.Fn9353
func Fn9353(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9356 github.com/goccy/spidermonkeywasm2go/p6.Fn9356
func Fn9356(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9357 github.com/goccy/spidermonkeywasm2go/p5.Fn9357
func Fn9357(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9358 github.com/goccy/spidermonkeywasm2go/p5.Fn9358
func Fn9358(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9359 github.com/goccy/spidermonkeywasm2go/p6.Fn9359
func Fn9359(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9361 github.com/goccy/spidermonkeywasm2go/p5.Fn9361
func Fn9361(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9362 github.com/goccy/spidermonkeywasm2go/p5.Fn9362
func Fn9362(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9415 github.com/goccy/spidermonkeywasm2go/p5.Fn9415
func Fn9415(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9416 github.com/goccy/spidermonkeywasm2go/p2.Fn9416
func Fn9416(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn9419 github.com/goccy/spidermonkeywasm2go/p5.Fn9419
func Fn9419(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9425 github.com/goccy/spidermonkeywasm2go/p5.Fn9425
func Fn9425(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9432 github.com/goccy/spidermonkeywasm2go/p4.Fn9432
func Fn9432(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn9475 github.com/goccy/spidermonkeywasm2go/p4.Fn9475
func Fn9475(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9476 github.com/goccy/spidermonkeywasm2go/p3.Fn9476
func Fn9476(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9509 github.com/goccy/spidermonkeywasm2go/p6.Fn9509
func Fn9509(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9512 github.com/goccy/spidermonkeywasm2go/p6.Fn9512
func Fn9512(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9518 github.com/goccy/spidermonkeywasm2go/p6.Fn9518
func Fn9518(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9520 github.com/goccy/spidermonkeywasm2go/p6.Fn9520
func Fn9520(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9550 github.com/goccy/spidermonkeywasm2go/p7.Fn9550
func Fn9550(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9551 github.com/goccy/spidermonkeywasm2go/p7.Fn9551
func Fn9551(m *base.Module, l0 int32)

//go:linkname Fn9552 github.com/goccy/spidermonkeywasm2go/p6.Fn9552
func Fn9552(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn9553 github.com/goccy/spidermonkeywasm2go/p6.Fn9553
func Fn9553(m *base.Module, l0 int32, l1 float64, l2 int32)

//go:linkname Fn9566 github.com/goccy/spidermonkeywasm2go/p5.Fn9566
func Fn9566(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn9568 github.com/goccy/spidermonkeywasm2go/p4.Fn9568
func Fn9568(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9582 github.com/goccy/spidermonkeywasm2go/p5.Fn9582
func Fn9582(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9595 github.com/goccy/spidermonkeywasm2go/p6.Fn9595
func Fn9595(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9614 github.com/goccy/spidermonkeywasm2go/p5.Fn9614
func Fn9614(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn9620 github.com/goccy/spidermonkeywasm2go/p4.Fn9620
func Fn9620(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9621 github.com/goccy/spidermonkeywasm2go/p5.Fn9621
func Fn9621(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9635 github.com/goccy/spidermonkeywasm2go/p2.Fn9635
func Fn9635(m *base.Module, l0 int32)

//go:linkname Fn9636 github.com/goccy/spidermonkeywasm2go/p3.Fn9636
func Fn9636(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9640 github.com/goccy/spidermonkeywasm2go/p7.Fn9640
func Fn9640(m *base.Module, l0 int32) int32

//go:linkname Fn9642 github.com/goccy/spidermonkeywasm2go/p6.Fn9642
func Fn9642(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9662 github.com/goccy/spidermonkeywasm2go/p7.Fn9662
func Fn9662(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9665 github.com/goccy/spidermonkeywasm2go/p6.Fn9665
func Fn9665(m *base.Module, l0 int32)

//go:linkname Fn9666 github.com/goccy/spidermonkeywasm2go/p6.Fn9666
func Fn9666(m *base.Module, l0 int32)

//go:linkname Fn9667 github.com/goccy/spidermonkeywasm2go/p6.Fn9667
func Fn9667(m *base.Module, l0 int32)

//go:linkname Fn9674 github.com/goccy/spidermonkeywasm2go/p6.Fn9674
func Fn9674(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9676 github.com/goccy/spidermonkeywasm2go/p6.Fn9676
func Fn9676(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9686 github.com/goccy/spidermonkeywasm2go/p5.Fn9686
func Fn9686(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9698 github.com/goccy/spidermonkeywasm2go/p4.Fn9698
func Fn9698(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9699 github.com/goccy/spidermonkeywasm2go/p4.Fn9699
func Fn9699(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9756 github.com/goccy/spidermonkeywasm2go/p7.Fn9756
func Fn9756(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9761 github.com/goccy/spidermonkeywasm2go/p7.Fn9761
func Fn9761(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10106 github.com/goccy/spidermonkeywasm2go/p6.Fn10106
func Fn10106(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10116 github.com/goccy/spidermonkeywasm2go/p7.Fn10116
func Fn10116(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10143 github.com/goccy/spidermonkeywasm2go/p5.Fn10143
func Fn10143(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10144 github.com/goccy/spidermonkeywasm2go/p7.Fn10144
func Fn10144(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10145 github.com/goccy/spidermonkeywasm2go/p7.Fn10145
func Fn10145(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10146 github.com/goccy/spidermonkeywasm2go/p7.Fn10146
func Fn10146(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10168 github.com/goccy/spidermonkeywasm2go/p6.Fn10168
func Fn10168(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10169 github.com/goccy/spidermonkeywasm2go/p5.Fn10169
func Fn10169(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10170 github.com/goccy/spidermonkeywasm2go/p6.Fn10170
func Fn10170(m *base.Module, l0 int32)

//go:linkname Fn10171 github.com/goccy/spidermonkeywasm2go/p6.Fn10171
func Fn10171(m *base.Module, l0 int32) int32

//go:linkname Fn10172 github.com/goccy/spidermonkeywasm2go/p7.Fn10172
func Fn10172(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10173 github.com/goccy/spidermonkeywasm2go/p6.Fn10173
func Fn10173(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10174 github.com/goccy/spidermonkeywasm2go/p4.Fn10174
func Fn10174(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10187 github.com/goccy/spidermonkeywasm2go/p3.Fn10187
func Fn10187(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10188 github.com/goccy/spidermonkeywasm2go/p5.Fn10188
func Fn10188(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10196 github.com/goccy/spidermonkeywasm2go/p3.Fn10196
func Fn10196(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10212 github.com/goccy/spidermonkeywasm2go/p6.Fn10212
func Fn10212(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10214 github.com/goccy/spidermonkeywasm2go/p4.Fn10214
func Fn10214(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10217 github.com/goccy/spidermonkeywasm2go/p5.Fn10217
func Fn10217(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10218 github.com/goccy/spidermonkeywasm2go/p5.Fn10218
func Fn10218(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10221 github.com/goccy/spidermonkeywasm2go/p6.Fn10221
func Fn10221(m *base.Module, l0 int32, l1 int64, l2 int64) int32

//go:linkname Fn10236 github.com/goccy/spidermonkeywasm2go/p6.Fn10236
func Fn10236(m *base.Module, l0 int32)

//go:linkname Fn10237 github.com/goccy/spidermonkeywasm2go/p5.Fn10237
func Fn10237(m *base.Module, l0 int32) int32

//go:linkname Fn10238 github.com/goccy/spidermonkeywasm2go/p5.Fn10238
func Fn10238(m *base.Module, l0 int32)

//go:linkname Fn10239 github.com/goccy/spidermonkeywasm2go/p5.Fn10239
func Fn10239(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10240 github.com/goccy/spidermonkeywasm2go/p6.Fn10240
func Fn10240(m *base.Module, l0 int32) int32

//go:linkname Fn10241 github.com/goccy/spidermonkeywasm2go/p3.Fn10241
func Fn10241(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10253 github.com/goccy/spidermonkeywasm2go/p6.Fn10253
func Fn10253(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10257 github.com/goccy/spidermonkeywasm2go/p7.Fn10257
func Fn10257(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10262 github.com/goccy/spidermonkeywasm2go/p7.Fn10262
func Fn10262(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10265 github.com/goccy/spidermonkeywasm2go/p7.Fn10265
func Fn10265(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10268 github.com/goccy/spidermonkeywasm2go/p7.Fn10268
func Fn10268(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10270 github.com/goccy/spidermonkeywasm2go/p7.Fn10270
func Fn10270(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10279 github.com/goccy/spidermonkeywasm2go/p5.Fn10279
func Fn10279(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn10281 github.com/goccy/spidermonkeywasm2go/p6.Fn10281
func Fn10281(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10290 github.com/goccy/spidermonkeywasm2go/p3.Fn10290
func Fn10290(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10301 github.com/goccy/spidermonkeywasm2go/p4.Fn10301
func Fn10301(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10332 github.com/goccy/spidermonkeywasm2go/p6.Fn10332
func Fn10332(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10333 github.com/goccy/spidermonkeywasm2go/p5.Fn10333
func Fn10333(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10334 github.com/goccy/spidermonkeywasm2go/p4.Fn10334
func Fn10334(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn10335 github.com/goccy/spidermonkeywasm2go/p6.Fn10335
func Fn10335(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10336 github.com/goccy/spidermonkeywasm2go/p5.Fn10336
func Fn10336(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10337 github.com/goccy/spidermonkeywasm2go/p6.Fn10337
func Fn10337(m *base.Module, l0 int32) int32

//go:linkname Fn10547 github.com/goccy/spidermonkeywasm2go/p6.Fn10547
func Fn10547(m *base.Module, l0 int32) int32

//go:linkname Fn10566 github.com/goccy/spidermonkeywasm2go/p7.Fn10566
func Fn10566(m *base.Module, l0 int32) int32

//go:linkname Fn10780 github.com/goccy/spidermonkeywasm2go/p6.Fn10780
func Fn10780(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10783 github.com/goccy/spidermonkeywasm2go/p6.Fn10783
func Fn10783(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10784 github.com/goccy/spidermonkeywasm2go/p6.Fn10784
func Fn10784(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn10806 github.com/goccy/spidermonkeywasm2go/p7.Fn10806
func Fn10806(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10811 github.com/goccy/spidermonkeywasm2go/p5.Fn10811
func Fn10811(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10816 github.com/goccy/spidermonkeywasm2go/p6.Fn10816
func Fn10816(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10820 github.com/goccy/spidermonkeywasm2go/p6.Fn10820
func Fn10820(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10821 github.com/goccy/spidermonkeywasm2go/p5.Fn10821
func Fn10821(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn10926 github.com/goccy/spidermonkeywasm2go/p7.Fn10926
func Fn10926(m *base.Module)

//go:linkname Fn10928 github.com/goccy/spidermonkeywasm2go/p2.Fn10928
func Fn10928(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10959 github.com/goccy/spidermonkeywasm2go/p6.Fn10959
func Fn10959(m *base.Module, l0 int32) int32

//go:linkname Fn10960 github.com/goccy/spidermonkeywasm2go/p6.Fn10960
func Fn10960(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10963 github.com/goccy/spidermonkeywasm2go/p7.Fn10963
func Fn10963(m *base.Module, l0 int32)
