//go:build !amd64 && !arm64

package p6

import (
	base "github.com/goccy/spidermonkeywasm2go/base"
	_ "unsafe"
)

//go:linkname Fn27 github.com/goccy/spidermonkeywasm2go/p7.Fn27
func Fn27(m *base.Module, l0 int32) int32

//go:linkname Fn28 github.com/goccy/spidermonkeywasm2go/p7.Fn28
func Fn28(m *base.Module, l0 int32) int64

//go:linkname Fn29 github.com/goccy/spidermonkeywasm2go/p7.Fn29
func Fn29(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn30 github.com/goccy/spidermonkeywasm2go/p7.Fn30
func Fn30(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn31 github.com/goccy/spidermonkeywasm2go/p7.Fn31
func Fn31(m *base.Module, l0 int32)

//go:linkname Fn32 github.com/goccy/spidermonkeywasm2go/p7.Fn32
func Fn32(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn33 github.com/goccy/spidermonkeywasm2go/p7.Fn33
func Fn33(m *base.Module, l0 int32) int64

//go:linkname Fn34 github.com/goccy/spidermonkeywasm2go/p7.Fn34
func Fn34(m *base.Module, l0 int32)

//go:linkname Fn36 github.com/goccy/spidermonkeywasm2go/p7.Fn36
func Fn36(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn42 github.com/goccy/spidermonkeywasm2go/p7.Fn42
func Fn42(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn43 github.com/goccy/spidermonkeywasm2go/p7.Fn43
func Fn43(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn77 github.com/goccy/spidermonkeywasm2go/p7.Fn77
func Fn77(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn79 github.com/goccy/spidermonkeywasm2go/p7.Fn79
func Fn79(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn81 github.com/goccy/spidermonkeywasm2go/p5.Fn81
func Fn81(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn87 github.com/goccy/spidermonkeywasm2go/p5.Fn87
func Fn87(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn89 github.com/goccy/spidermonkeywasm2go/p5.Fn89
func Fn89(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn92 github.com/goccy/spidermonkeywasm2go/p7.Fn92
func Fn92(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn98 github.com/goccy/spidermonkeywasm2go/p7.Fn98
func Fn98(m *base.Module, l0 int32)

//go:linkname Fn102 github.com/goccy/spidermonkeywasm2go/p7.Fn102
func Fn102(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn104 github.com/goccy/spidermonkeywasm2go/p7.Fn104
func Fn104(m *base.Module, l0 int32)

//go:linkname Fn111 github.com/goccy/spidermonkeywasm2go/p7.Fn111
func Fn111(m *base.Module, l0 int32)

//go:linkname Fn116 github.com/goccy/spidermonkeywasm2go/p7.Fn116
func Fn116(m *base.Module, l0 int32) int32

//go:linkname Fn142 github.com/goccy/spidermonkeywasm2go/p7.Fn142
func Fn142(m *base.Module)

//go:linkname Fn144 github.com/goccy/spidermonkeywasm2go/p7.Fn144
func Fn144(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn146 github.com/goccy/spidermonkeywasm2go/p7.Fn146
func Fn146(m *base.Module)

//go:linkname Fn148 github.com/goccy/spidermonkeywasm2go/p0.Fn148
func Fn148(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn160 github.com/goccy/spidermonkeywasm2go/p7.Fn160
func Fn160(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn161 github.com/goccy/spidermonkeywasm2go/p7.Fn161
func Fn161(m *base.Module, l0 int32) int32

//go:linkname Fn165 github.com/goccy/spidermonkeywasm2go/p7.Fn165
func Fn165(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn166 github.com/goccy/spidermonkeywasm2go/p7.Fn166
func Fn166(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn168 github.com/goccy/spidermonkeywasm2go/p3.Fn168
func Fn168(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn171 github.com/goccy/spidermonkeywasm2go/p7.Fn171
func Fn171(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn174 github.com/goccy/spidermonkeywasm2go/p3.Fn174
func Fn174(m *base.Module, l0 float64) float64

//go:linkname Fn175 github.com/goccy/spidermonkeywasm2go/p5.Fn175
func Fn175(m *base.Module, l0 float64) float64

//go:linkname Fn176 github.com/goccy/spidermonkeywasm2go/p5.Fn176
func Fn176(m *base.Module, l0 float64) float64

//go:linkname Fn185 github.com/goccy/spidermonkeywasm2go/p5.Fn185
func Fn185(m *base.Module, l0 float64) float64

//go:linkname Fn192 github.com/goccy/spidermonkeywasm2go/p7.Fn192
func Fn192(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn195 github.com/goccy/spidermonkeywasm2go/p7.Fn195
func Fn195(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn196 github.com/goccy/spidermonkeywasm2go/p3.Fn196
func Fn196(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn197 github.com/goccy/spidermonkeywasm2go/p3.Fn197
func Fn197(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn198 github.com/goccy/spidermonkeywasm2go/p3.Fn198
func Fn198(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn204 github.com/goccy/spidermonkeywasm2go/p7.Fn204
func Fn204(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn205 github.com/goccy/spidermonkeywasm2go/p7.Fn205
func Fn205(m *base.Module, l0 int32)

//go:linkname Fn218 github.com/goccy/spidermonkeywasm2go/p7.Fn218
func Fn218(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn223 github.com/goccy/spidermonkeywasm2go/p4.Fn223
func Fn223(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn231 github.com/goccy/spidermonkeywasm2go/p4.Fn231
func Fn231(m *base.Module, l0 int64) int64

//go:linkname Fn232 github.com/goccy/spidermonkeywasm2go/p3.Fn232
func Fn232(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn233 github.com/goccy/spidermonkeywasm2go/p4.Fn233
func Fn233(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64)

//go:linkname Fn235 github.com/goccy/spidermonkeywasm2go/p5.Fn235
func Fn235(m *base.Module, l0 int64) int64

//go:linkname Fn236 github.com/goccy/spidermonkeywasm2go/p3.Fn236
func Fn236(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn237 github.com/goccy/spidermonkeywasm2go/p4.Fn237
func Fn237(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64)

//go:linkname Fn245 github.com/goccy/spidermonkeywasm2go/p4.Fn245
func Fn245(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn248 github.com/goccy/spidermonkeywasm2go/p7.Fn248
func Fn248(m *base.Module, l0 int32)

//go:linkname Fn252 github.com/goccy/spidermonkeywasm2go/p7.Fn252
func Fn252(m *base.Module, l0 int32) int32

//go:linkname Fn253 github.com/goccy/spidermonkeywasm2go/p7.Fn253
func Fn253(m *base.Module, l0 int32)

//go:linkname Fn256 github.com/goccy/spidermonkeywasm2go/p2.Fn256
func Fn256(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn258 github.com/goccy/spidermonkeywasm2go/p2.Fn258
func Fn258(m *base.Module, l0 int32) int32

//go:linkname Fn260 github.com/goccy/spidermonkeywasm2go/p7.Fn260
func Fn260(m *base.Module, l0 int32) int32

//go:linkname Fn261 github.com/goccy/spidermonkeywasm2go/p7.Fn261
func Fn261(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn262 github.com/goccy/spidermonkeywasm2go/p7.Fn262
func Fn262(m *base.Module, l0 int32) int32

//go:linkname Fn266 github.com/goccy/spidermonkeywasm2go/p4.Fn266
func Fn266(m *base.Module, l0 int32) int32

//go:linkname Fn267 github.com/goccy/spidermonkeywasm2go/p5.Fn267
func Fn267(m *base.Module, l0 int32) int32

//go:linkname Fn268 github.com/goccy/spidermonkeywasm2go/p7.Fn268
func Fn268(m *base.Module, l0 int32)

//go:linkname Fn275 github.com/goccy/spidermonkeywasm2go/p7.Fn275
func Fn275(m *base.Module, l0 int32)

//go:linkname Fn279 github.com/goccy/spidermonkeywasm2go/p7.Fn279
func Fn279(m *base.Module, l0 int32)

//go:linkname Fn295 github.com/goccy/spidermonkeywasm2go/p7.Fn295
func Fn295(m *base.Module, l0 int32)

//go:linkname Fn296 github.com/goccy/spidermonkeywasm2go/p7.Fn296
func Fn296(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn374 github.com/goccy/spidermonkeywasm2go/p7.Fn374
func Fn374(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn377 github.com/goccy/spidermonkeywasm2go/p7.Fn377
func Fn377(m *base.Module)

//go:linkname Fn378 github.com/goccy/spidermonkeywasm2go/p7.Fn378
func Fn378(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn379 github.com/goccy/spidermonkeywasm2go/p7.Fn379
func Fn379(m *base.Module)

//go:linkname Fn380 github.com/goccy/spidermonkeywasm2go/p7.Fn380
func Fn380(m *base.Module, l0 int32) int32

//go:linkname Fn382 github.com/goccy/spidermonkeywasm2go/p7.Fn382
func Fn382(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn390 github.com/goccy/spidermonkeywasm2go/p7.Fn390
func Fn390(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn391 github.com/goccy/spidermonkeywasm2go/p5.Fn391
func Fn391(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn395 github.com/goccy/spidermonkeywasm2go/p7.Fn395
func Fn395(m *base.Module, l0 int32)

//go:linkname Fn396 github.com/goccy/spidermonkeywasm2go/p7.Fn396
func Fn396(m *base.Module, l0 int32)

//go:linkname Fn397 github.com/goccy/spidermonkeywasm2go/p1.Fn397
func Fn397(m *base.Module, l0 int32)

//go:linkname Fn399 github.com/goccy/spidermonkeywasm2go/p7.Fn399
func Fn399(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn402 github.com/goccy/spidermonkeywasm2go/p7.Fn402
func Fn402(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn424 github.com/goccy/spidermonkeywasm2go/p7.Fn424
func Fn424(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn425 github.com/goccy/spidermonkeywasm2go/p7.Fn425
func Fn425(m *base.Module, l0 int32)

//go:linkname Fn434 github.com/goccy/spidermonkeywasm2go/p7.Fn434
func Fn434(m *base.Module)

//go:linkname Fn464 github.com/goccy/spidermonkeywasm2go/p5.Fn464
func Fn464(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn482 github.com/goccy/spidermonkeywasm2go/p5.Fn482
func Fn482(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn499 github.com/goccy/spidermonkeywasm2go/p5.Fn499
func Fn499(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn500 github.com/goccy/spidermonkeywasm2go/p5.Fn500
func Fn500(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn503 github.com/goccy/spidermonkeywasm2go/p5.Fn503
func Fn503(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32)

//go:linkname Fn515 github.com/goccy/spidermonkeywasm2go/p5.Fn515
func Fn515(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn532 github.com/goccy/spidermonkeywasm2go/p5.Fn532
func Fn532(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn540 github.com/goccy/spidermonkeywasm2go/p5.Fn540
func Fn540(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn568 github.com/goccy/spidermonkeywasm2go/p2.Fn568
func Fn568(m *base.Module) int32

//go:linkname Fn571 github.com/goccy/spidermonkeywasm2go/p7.Fn571
func Fn571(m *base.Module, l0 int32)

//go:linkname Fn573 github.com/goccy/spidermonkeywasm2go/p7.Fn573
func Fn573(m *base.Module, l0 int32)

//go:linkname Fn574 github.com/goccy/spidermonkeywasm2go/p7.Fn574
func Fn574(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn662 github.com/goccy/spidermonkeywasm2go/p7.Fn662
func Fn662(m *base.Module, l0 int32) int32

//go:linkname Fn663 github.com/goccy/spidermonkeywasm2go/p7.Fn663
func Fn663(m *base.Module, l0 int32)

//go:linkname Fn665 github.com/goccy/spidermonkeywasm2go/p7.Fn665
func Fn665(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn677 github.com/goccy/spidermonkeywasm2go/p7.Fn677
func Fn677(m *base.Module, l0 int32) int32

//go:linkname Fn679 github.com/goccy/spidermonkeywasm2go/p7.Fn679
func Fn679(m *base.Module, l0 int32)

//go:linkname Fn680 github.com/goccy/spidermonkeywasm2go/p7.Fn680
func Fn680(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn684 github.com/goccy/spidermonkeywasm2go/p7.Fn684
func Fn684(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn685 github.com/goccy/spidermonkeywasm2go/p7.Fn685
func Fn685(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn687 github.com/goccy/spidermonkeywasm2go/p7.Fn687
func Fn687(m *base.Module, l0 int32)

//go:linkname Fn692 github.com/goccy/spidermonkeywasm2go/p7.Fn692
func Fn692(m *base.Module, l0 int32)

//go:linkname Fn704 github.com/goccy/spidermonkeywasm2go/p4.Fn704
func Fn704(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn705 github.com/goccy/spidermonkeywasm2go/p7.Fn705
func Fn705(m *base.Module, l0 int32)

//go:linkname Fn710 github.com/goccy/spidermonkeywasm2go/p5.Fn710
func Fn710(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn714 github.com/goccy/spidermonkeywasm2go/p5.Fn714
func Fn714(m *base.Module) int32

//go:linkname Fn719 github.com/goccy/spidermonkeywasm2go/p7.Fn719
func Fn719(m *base.Module, l0 float64) float64

//go:linkname Fn723 github.com/goccy/spidermonkeywasm2go/p7.Fn723
func Fn723(m *base.Module, l0 float64) float64

//go:linkname Fn725 github.com/goccy/spidermonkeywasm2go/p7.Fn725
func Fn725(m *base.Module, l0 float64) float64

//go:linkname Fn729 github.com/goccy/spidermonkeywasm2go/p5.Fn729
func Fn729(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn732 github.com/goccy/spidermonkeywasm2go/p7.Fn732
func Fn732(m *base.Module, l0 int32)

//go:linkname Fn737 github.com/goccy/spidermonkeywasm2go/p7.Fn737
func Fn737(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn741 github.com/goccy/spidermonkeywasm2go/p5.Fn741
func Fn741(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn742 github.com/goccy/spidermonkeywasm2go/p7.Fn742
func Fn742(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn744 github.com/goccy/spidermonkeywasm2go/p7.Fn744
func Fn744(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn746 github.com/goccy/spidermonkeywasm2go/p7.Fn746
func Fn746(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn747 github.com/goccy/spidermonkeywasm2go/p7.Fn747
func Fn747(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn749 github.com/goccy/spidermonkeywasm2go/p7.Fn749
func Fn749(m *base.Module, l0 int32) int32

//go:linkname Fn750 github.com/goccy/spidermonkeywasm2go/p7.Fn750
func Fn750(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn751 github.com/goccy/spidermonkeywasm2go/p7.Fn751
func Fn751(m *base.Module, l0 int32) int32

//go:linkname Fn752 github.com/goccy/spidermonkeywasm2go/p7.Fn752
func Fn752(m *base.Module, l0 int32)

//go:linkname Fn753 github.com/goccy/spidermonkeywasm2go/p7.Fn753
func Fn753(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn755 github.com/goccy/spidermonkeywasm2go/p7.Fn755
func Fn755(m *base.Module, l0 int32) int32

//go:linkname Fn756 github.com/goccy/spidermonkeywasm2go/p7.Fn756
func Fn756(m *base.Module, l0 int32) int32

//go:linkname Fn758 github.com/goccy/spidermonkeywasm2go/p7.Fn758
func Fn758(m *base.Module) int32

//go:linkname Fn759 github.com/goccy/spidermonkeywasm2go/p7.Fn759
func Fn759(m *base.Module, l0 int32)

//go:linkname Fn762 github.com/goccy/spidermonkeywasm2go/p7.Fn762
func Fn762(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn765 github.com/goccy/spidermonkeywasm2go/p7.Fn765
func Fn765(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn775 github.com/goccy/spidermonkeywasm2go/p7.Fn775
func Fn775(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn777 github.com/goccy/spidermonkeywasm2go/p7.Fn777
func Fn777(m *base.Module, l0 int32) int32

//go:linkname Fn780 github.com/goccy/spidermonkeywasm2go/p7.Fn780
func Fn780(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn784 github.com/goccy/spidermonkeywasm2go/p7.Fn784
func Fn784(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn788 github.com/goccy/spidermonkeywasm2go/p7.Fn788
func Fn788(m *base.Module, l0 int32) int32

//go:linkname Fn789 github.com/goccy/spidermonkeywasm2go/p7.Fn789
func Fn789(m *base.Module, l0 int32) int32

//go:linkname Fn793 github.com/goccy/spidermonkeywasm2go/p7.Fn793
func Fn793(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn798 github.com/goccy/spidermonkeywasm2go/p7.Fn798
func Fn798(m *base.Module, l0 int32) int32

//go:linkname Fn804 github.com/goccy/spidermonkeywasm2go/p7.Fn804
func Fn804(m *base.Module, l0 int32) int32

//go:linkname Fn805 github.com/goccy/spidermonkeywasm2go/p7.Fn805
func Fn805(m *base.Module, l0 int32) int32

//go:linkname Fn806 github.com/goccy/spidermonkeywasm2go/p7.Fn806
func Fn806(m *base.Module, l0 int32) int32

//go:linkname Fn808 github.com/goccy/spidermonkeywasm2go/p7.Fn808
func Fn808(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn809 github.com/goccy/spidermonkeywasm2go/p7.Fn809
func Fn809(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn810 github.com/goccy/spidermonkeywasm2go/p5.Fn810
func Fn810(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn811 github.com/goccy/spidermonkeywasm2go/p7.Fn811
func Fn811(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn812 github.com/goccy/spidermonkeywasm2go/p4.Fn812
func Fn812(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn823 github.com/goccy/spidermonkeywasm2go/p2.Fn823
func Fn823(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn826 github.com/goccy/spidermonkeywasm2go/p7.Fn826
func Fn826(m *base.Module)

//go:linkname Fn827 github.com/goccy/spidermonkeywasm2go/p2.Fn827
func Fn827(m *base.Module, l0 int32, l1 int32, l2 int32) float64

//go:linkname Fn829 github.com/goccy/spidermonkeywasm2go/p7.Fn829
func Fn829(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn832 github.com/goccy/spidermonkeywasm2go/p2.Fn832
func Fn832(m *base.Module, l0 int32) int32

//go:linkname Fn833 github.com/goccy/spidermonkeywasm2go/p4.Fn833
func Fn833(m *base.Module, l0 int32)

//go:linkname Fn834 github.com/goccy/spidermonkeywasm2go/p7.Fn834
func Fn834(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn835 github.com/goccy/spidermonkeywasm2go/p5.Fn835
func Fn835(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn837 github.com/goccy/spidermonkeywasm2go/p5.Fn837
func Fn837(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn840 github.com/goccy/spidermonkeywasm2go/p0.Fn840
func Fn840(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn844 github.com/goccy/spidermonkeywasm2go/p4.Fn844
func Fn844(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn850 github.com/goccy/spidermonkeywasm2go/p0.Fn850
func Fn850(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn855 github.com/goccy/spidermonkeywasm2go/p7.Fn855
func Fn855(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn857 github.com/goccy/spidermonkeywasm2go/p0.Fn857
func Fn857(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn858 github.com/goccy/spidermonkeywasm2go/p0.Fn858
func Fn858(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn859 github.com/goccy/spidermonkeywasm2go/p0.Fn859
func Fn859(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn860 github.com/goccy/spidermonkeywasm2go/p0.Fn860
func Fn860(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn862 github.com/goccy/spidermonkeywasm2go/p0.Fn862
func Fn862(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn863 github.com/goccy/spidermonkeywasm2go/p0.Fn863
func Fn863(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn867 github.com/goccy/spidermonkeywasm2go/p4.Fn867
func Fn867(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn868 github.com/goccy/spidermonkeywasm2go/p7.Fn868
func Fn868(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn879 github.com/goccy/spidermonkeywasm2go/p0.Fn879
func Fn879(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn880 github.com/goccy/spidermonkeywasm2go/p5.Fn880
func Fn880(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn883 github.com/goccy/spidermonkeywasm2go/p0.Fn883
func Fn883(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn891 github.com/goccy/spidermonkeywasm2go/p0.Fn891
func Fn891(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn894 github.com/goccy/spidermonkeywasm2go/p0.Fn894
func Fn894(m *base.Module, l0 int32) int32

//go:linkname Fn895 github.com/goccy/spidermonkeywasm2go/p0.Fn895
func Fn895(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn899 github.com/goccy/spidermonkeywasm2go/p7.Fn899
func Fn899(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn900 github.com/goccy/spidermonkeywasm2go/p5.Fn900
func Fn900(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn944 github.com/goccy/spidermonkeywasm2go/p5.Fn944
func Fn944(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn945 github.com/goccy/spidermonkeywasm2go/p7.Fn945
func Fn945(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn951 github.com/goccy/spidermonkeywasm2go/p5.Fn951
func Fn951(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn956 github.com/goccy/spidermonkeywasm2go/p5.Fn956
func Fn956(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn961 github.com/goccy/spidermonkeywasm2go/p5.Fn961
func Fn961(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn972 github.com/goccy/spidermonkeywasm2go/p5.Fn972
func Fn972(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn976 github.com/goccy/spidermonkeywasm2go/p5.Fn976
func Fn976(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn980 github.com/goccy/spidermonkeywasm2go/p5.Fn980
func Fn980(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn984 github.com/goccy/spidermonkeywasm2go/p5.Fn984
func Fn984(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn988 github.com/goccy/spidermonkeywasm2go/p5.Fn988
func Fn988(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn991 github.com/goccy/spidermonkeywasm2go/p5.Fn991
func Fn991(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn994 github.com/goccy/spidermonkeywasm2go/p5.Fn994
func Fn994(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1010 github.com/goccy/spidermonkeywasm2go/p5.Fn1010
func Fn1010(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1090 github.com/goccy/spidermonkeywasm2go/p7.Fn1090
func Fn1090(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1094 github.com/goccy/spidermonkeywasm2go/p0.Fn1094
func Fn1094(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1095 github.com/goccy/spidermonkeywasm2go/p0.Fn1095
func Fn1095(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1098 github.com/goccy/spidermonkeywasm2go/p0.Fn1098
func Fn1098(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1114 github.com/goccy/spidermonkeywasm2go/p0.Fn1114
func Fn1114(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1116 github.com/goccy/spidermonkeywasm2go/p0.Fn1116
func Fn1116(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1128 github.com/goccy/spidermonkeywasm2go/p0.Fn1128
func Fn1128(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1140 github.com/goccy/spidermonkeywasm2go/p0.Fn1140
func Fn1140(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1142 github.com/goccy/spidermonkeywasm2go/p0.Fn1142
func Fn1142(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1146 github.com/goccy/spidermonkeywasm2go/p4.Fn1146
func Fn1146(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1150 github.com/goccy/spidermonkeywasm2go/p5.Fn1150
func Fn1150(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1155 github.com/goccy/spidermonkeywasm2go/p5.Fn1155
func Fn1155(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1171 github.com/goccy/spidermonkeywasm2go/p5.Fn1171
func Fn1171(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1185 github.com/goccy/spidermonkeywasm2go/p4.Fn1185
func Fn1185(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn1255 github.com/goccy/spidermonkeywasm2go/p7.Fn1255
func Fn1255(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1260 github.com/goccy/spidermonkeywasm2go/p4.Fn1260
func Fn1260(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1263 github.com/goccy/spidermonkeywasm2go/p7.Fn1263
func Fn1263(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1289 github.com/goccy/spidermonkeywasm2go/p4.Fn1289
func Fn1289(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1295 github.com/goccy/spidermonkeywasm2go/p4.Fn1295
func Fn1295(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1300 github.com/goccy/spidermonkeywasm2go/p0.Fn1300
func Fn1300(m *base.Module, l0 int32)

//go:linkname Fn1321 github.com/goccy/spidermonkeywasm2go/p4.Fn1321
func Fn1321(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1345 github.com/goccy/spidermonkeywasm2go/p3.Fn1345
func Fn1345(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1353 github.com/goccy/spidermonkeywasm2go/p5.Fn1353
func Fn1353(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1356 github.com/goccy/spidermonkeywasm2go/p3.Fn1356
func Fn1356(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1360 github.com/goccy/spidermonkeywasm2go/p3.Fn1360
func Fn1360(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1363 github.com/goccy/spidermonkeywasm2go/p4.Fn1363
func Fn1363(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1369 github.com/goccy/spidermonkeywasm2go/p4.Fn1369
func Fn1369(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1382 github.com/goccy/spidermonkeywasm2go/p4.Fn1382
func Fn1382(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1385 github.com/goccy/spidermonkeywasm2go/p4.Fn1385
func Fn1385(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1390 github.com/goccy/spidermonkeywasm2go/p5.Fn1390
func Fn1390(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1416 github.com/goccy/spidermonkeywasm2go/p7.Fn1416
func Fn1416(m *base.Module, l0 int32) int32

//go:linkname Fn1433 github.com/goccy/spidermonkeywasm2go/p7.Fn1433
func Fn1433(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1444 github.com/goccy/spidermonkeywasm2go/p0.Fn1444
func Fn1444(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1445 github.com/goccy/spidermonkeywasm2go/p5.Fn1445
func Fn1445(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1449 github.com/goccy/spidermonkeywasm2go/p5.Fn1449
func Fn1449(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1450 github.com/goccy/spidermonkeywasm2go/p3.Fn1450
func Fn1450(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1451 github.com/goccy/spidermonkeywasm2go/p0.Fn1451
func Fn1451(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1453 github.com/goccy/spidermonkeywasm2go/p5.Fn1453
func Fn1453(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1457 github.com/goccy/spidermonkeywasm2go/p0.Fn1457
func Fn1457(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1468 github.com/goccy/spidermonkeywasm2go/p5.Fn1468
func Fn1468(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1470 github.com/goccy/spidermonkeywasm2go/p7.Fn1470
func Fn1470(m *base.Module, l0 int32)

//go:linkname Fn1472 github.com/goccy/spidermonkeywasm2go/p5.Fn1472
func Fn1472(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1475 github.com/goccy/spidermonkeywasm2go/p5.Fn1475
func Fn1475(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1478 github.com/goccy/spidermonkeywasm2go/p7.Fn1478
func Fn1478(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1480 github.com/goccy/spidermonkeywasm2go/p0.Fn1480
func Fn1480(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1481 github.com/goccy/spidermonkeywasm2go/p0.Fn1481
func Fn1481(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1495 github.com/goccy/spidermonkeywasm2go/p7.Fn1495
func Fn1495(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1503 github.com/goccy/spidermonkeywasm2go/p0.Fn1503
func Fn1503(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1504 github.com/goccy/spidermonkeywasm2go/p7.Fn1504
func Fn1504(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1505 github.com/goccy/spidermonkeywasm2go/p7.Fn1505
func Fn1505(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1506 github.com/goccy/spidermonkeywasm2go/p0.Fn1506
func Fn1506(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1515 github.com/goccy/spidermonkeywasm2go/p5.Fn1515
func Fn1515(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1518 github.com/goccy/spidermonkeywasm2go/p7.Fn1518
func Fn1518(m *base.Module, l0 int32) int32

//go:linkname Fn1521 github.com/goccy/spidermonkeywasm2go/p5.Fn1521
func Fn1521(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn1522 github.com/goccy/spidermonkeywasm2go/p3.Fn1522
func Fn1522(m *base.Module, l0 int32)

//go:linkname Fn1530 github.com/goccy/spidermonkeywasm2go/p3.Fn1530
func Fn1530(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 int32, l5 int32) int32

//go:linkname Fn1532 github.com/goccy/spidermonkeywasm2go/p1.Fn1532
func Fn1532(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1591 github.com/goccy/spidermonkeywasm2go/p4.Fn1591
func Fn1591(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1606 github.com/goccy/spidermonkeywasm2go/p7.Fn1606
func Fn1606(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1628 github.com/goccy/spidermonkeywasm2go/p4.Fn1628
func Fn1628(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1636 github.com/goccy/spidermonkeywasm2go/p7.Fn1636
func Fn1636(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1664 github.com/goccy/spidermonkeywasm2go/p5.Fn1664
func Fn1664(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn1669 github.com/goccy/spidermonkeywasm2go/p0.Fn1669
func Fn1669(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1672 github.com/goccy/spidermonkeywasm2go/p0.Fn1672
func Fn1672(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1724 github.com/goccy/spidermonkeywasm2go/p7.Fn1724
func Fn1724(m *base.Module, l0 int32)

//go:linkname Fn1730 github.com/goccy/spidermonkeywasm2go/p7.Fn1730
func Fn1730(m *base.Module, l0 int32)

//go:linkname Fn1753 github.com/goccy/spidermonkeywasm2go/p0.Fn1753
func Fn1753(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1759 github.com/goccy/spidermonkeywasm2go/p0.Fn1759
func Fn1759(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1760 github.com/goccy/spidermonkeywasm2go/p0.Fn1760
func Fn1760(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1761 github.com/goccy/spidermonkeywasm2go/p7.Fn1761
func Fn1761(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1762 github.com/goccy/spidermonkeywasm2go/p0.Fn1762
func Fn1762(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1767 github.com/goccy/spidermonkeywasm2go/p0.Fn1767
func Fn1767(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1769 github.com/goccy/spidermonkeywasm2go/p7.Fn1769
func Fn1769(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1771 github.com/goccy/spidermonkeywasm2go/p7.Fn1771
func Fn1771(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1814 github.com/goccy/spidermonkeywasm2go/p7.Fn1814
func Fn1814(m *base.Module, l0 int32) int32

//go:linkname Fn1836 github.com/goccy/spidermonkeywasm2go/p0.Fn1836
func Fn1836(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1854 github.com/goccy/spidermonkeywasm2go/p3.Fn1854
func Fn1854(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1855 github.com/goccy/spidermonkeywasm2go/p0.Fn1855
func Fn1855(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1856 github.com/goccy/spidermonkeywasm2go/p7.Fn1856
func Fn1856(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1857 github.com/goccy/spidermonkeywasm2go/p5.Fn1857
func Fn1857(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1858 github.com/goccy/spidermonkeywasm2go/p0.Fn1858
func Fn1858(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1881 github.com/goccy/spidermonkeywasm2go/p7.Fn1881
func Fn1881(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1882 github.com/goccy/spidermonkeywasm2go/p7.Fn1882
func Fn1882(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1886 github.com/goccy/spidermonkeywasm2go/p4.Fn1886
func Fn1886(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1887 github.com/goccy/spidermonkeywasm2go/p4.Fn1887
func Fn1887(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1900 github.com/goccy/spidermonkeywasm2go/p7.Fn1900
func Fn1900(m *base.Module, l0 int32) int32

//go:linkname Fn1901 github.com/goccy/spidermonkeywasm2go/p7.Fn1901
func Fn1901(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1944 github.com/goccy/spidermonkeywasm2go/p0.Fn1944
func Fn1944(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1947 github.com/goccy/spidermonkeywasm2go/p0.Fn1947
func Fn1947(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1960 github.com/goccy/spidermonkeywasm2go/p0.Fn1960
func Fn1960(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1961 github.com/goccy/spidermonkeywasm2go/p0.Fn1961
func Fn1961(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1962 github.com/goccy/spidermonkeywasm2go/p0.Fn1962
func Fn1962(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1963 github.com/goccy/spidermonkeywasm2go/p0.Fn1963
func Fn1963(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1964 github.com/goccy/spidermonkeywasm2go/p0.Fn1964
func Fn1964(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn1966 github.com/goccy/spidermonkeywasm2go/p0.Fn1966
func Fn1966(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1968 github.com/goccy/spidermonkeywasm2go/p0.Fn1968
func Fn1968(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1969 github.com/goccy/spidermonkeywasm2go/p0.Fn1969
func Fn1969(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1970 github.com/goccy/spidermonkeywasm2go/p0.Fn1970
func Fn1970(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1971 github.com/goccy/spidermonkeywasm2go/p0.Fn1971
func Fn1971(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1972 github.com/goccy/spidermonkeywasm2go/p0.Fn1972
func Fn1972(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1973 github.com/goccy/spidermonkeywasm2go/p0.Fn1973
func Fn1973(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1975 github.com/goccy/spidermonkeywasm2go/p0.Fn1975
func Fn1975(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1976 github.com/goccy/spidermonkeywasm2go/p0.Fn1976
func Fn1976(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1977 github.com/goccy/spidermonkeywasm2go/p0.Fn1977
func Fn1977(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1978 github.com/goccy/spidermonkeywasm2go/p0.Fn1978
func Fn1978(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1982 github.com/goccy/spidermonkeywasm2go/p5.Fn1982
func Fn1982(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn1983 github.com/goccy/spidermonkeywasm2go/p5.Fn1983
func Fn1983(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn1986 github.com/goccy/spidermonkeywasm2go/p0.Fn1986
func Fn1986(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1988 github.com/goccy/spidermonkeywasm2go/p0.Fn1988
func Fn1988(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1992 github.com/goccy/spidermonkeywasm2go/p0.Fn1992
func Fn1992(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2002 github.com/goccy/spidermonkeywasm2go/p0.Fn2002
func Fn2002(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2013 github.com/goccy/spidermonkeywasm2go/p0.Fn2013
func Fn2013(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2034 github.com/goccy/spidermonkeywasm2go/p7.Fn2034
func Fn2034(m *base.Module, l0 int32) int32

//go:linkname Fn2044 github.com/goccy/spidermonkeywasm2go/p5.Fn2044
func Fn2044(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2049 github.com/goccy/spidermonkeywasm2go/p7.Fn2049
func Fn2049(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2053 github.com/goccy/spidermonkeywasm2go/p7.Fn2053
func Fn2053(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2062 github.com/goccy/spidermonkeywasm2go/p5.Fn2062
func Fn2062(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2064 github.com/goccy/spidermonkeywasm2go/p0.Fn2064
func Fn2064(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2069 github.com/goccy/spidermonkeywasm2go/p0.Fn2069
func Fn2069(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2082 github.com/goccy/spidermonkeywasm2go/p7.Fn2082
func Fn2082(m *base.Module, l0 int32) int32

//go:linkname Fn2093 github.com/goccy/spidermonkeywasm2go/p7.Fn2093
func Fn2093(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2102 github.com/goccy/spidermonkeywasm2go/p4.Fn2102
func Fn2102(m *base.Module, l0 int32)

//go:linkname Fn2103 github.com/goccy/spidermonkeywasm2go/p7.Fn2103
func Fn2103(m *base.Module, l0 int32) int32

//go:linkname Fn2109 github.com/goccy/spidermonkeywasm2go/p7.Fn2109
func Fn2109(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn2112 github.com/goccy/spidermonkeywasm2go/p5.Fn2112
func Fn2112(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2140 github.com/goccy/spidermonkeywasm2go/p4.Fn2140
func Fn2140(m *base.Module, l0 int32)

//go:linkname Fn2142 github.com/goccy/spidermonkeywasm2go/p4.Fn2142
func Fn2142(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2164 github.com/goccy/spidermonkeywasm2go/p5.Fn2164
func Fn2164(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2194 github.com/goccy/spidermonkeywasm2go/p7.Fn2194
func Fn2194(m *base.Module, l0 int32)

//go:linkname Fn2199 github.com/goccy/spidermonkeywasm2go/p7.Fn2199
func Fn2199(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2207 github.com/goccy/spidermonkeywasm2go/p4.Fn2207
func Fn2207(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2211 github.com/goccy/spidermonkeywasm2go/p0.Fn2211
func Fn2211(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2214 github.com/goccy/spidermonkeywasm2go/p4.Fn2214
func Fn2214(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2215 github.com/goccy/spidermonkeywasm2go/p5.Fn2215
func Fn2215(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2216 github.com/goccy/spidermonkeywasm2go/p5.Fn2216
func Fn2216(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn2247 github.com/goccy/spidermonkeywasm2go/p5.Fn2247
func Fn2247(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2249 github.com/goccy/spidermonkeywasm2go/p4.Fn2249
func Fn2249(m *base.Module, l0 int32) int32

//go:linkname Fn2256 github.com/goccy/spidermonkeywasm2go/p0.Fn2256
func Fn2256(m *base.Module, l0 int32) int32

//go:linkname Fn2261 github.com/goccy/spidermonkeywasm2go/p0.Fn2261
func Fn2261(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2263 github.com/goccy/spidermonkeywasm2go/p0.Fn2263
func Fn2263(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2290 github.com/goccy/spidermonkeywasm2go/p0.Fn2290
func Fn2290(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2298 github.com/goccy/spidermonkeywasm2go/p3.Fn2298
func Fn2298(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2300 github.com/goccy/spidermonkeywasm2go/p0.Fn2300
func Fn2300(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2302 github.com/goccy/spidermonkeywasm2go/p0.Fn2302
func Fn2302(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2309 github.com/goccy/spidermonkeywasm2go/p0.Fn2309
func Fn2309(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2315 github.com/goccy/spidermonkeywasm2go/p7.Fn2315
func Fn2315(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn2317 github.com/goccy/spidermonkeywasm2go/p0.Fn2317
func Fn2317(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2324 github.com/goccy/spidermonkeywasm2go/p0.Fn2324
func Fn2324(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2331 github.com/goccy/spidermonkeywasm2go/p0.Fn2331
func Fn2331(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2332 github.com/goccy/spidermonkeywasm2go/p0.Fn2332
func Fn2332(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2335 github.com/goccy/spidermonkeywasm2go/p0.Fn2335
func Fn2335(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2342 github.com/goccy/spidermonkeywasm2go/p7.Fn2342
func Fn2342(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2343 github.com/goccy/spidermonkeywasm2go/p0.Fn2343
func Fn2343(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2345 github.com/goccy/spidermonkeywasm2go/p0.Fn2345
func Fn2345(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2355 github.com/goccy/spidermonkeywasm2go/p7.Fn2355
func Fn2355(m *base.Module, l0 int32)

//go:linkname Fn2438 github.com/goccy/spidermonkeywasm2go/p5.Fn2438
func Fn2438(m *base.Module, l0 int32)

//go:linkname Fn2439 github.com/goccy/spidermonkeywasm2go/p5.Fn2439
func Fn2439(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2440 github.com/goccy/spidermonkeywasm2go/p5.Fn2440
func Fn2440(m *base.Module, l0 int32) int32

//go:linkname Fn2453 github.com/goccy/spidermonkeywasm2go/p4.Fn2453
func Fn2453(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2457 github.com/goccy/spidermonkeywasm2go/p5.Fn2457
func Fn2457(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2458 github.com/goccy/spidermonkeywasm2go/p0.Fn2458
func Fn2458(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2461 github.com/goccy/spidermonkeywasm2go/p5.Fn2461
func Fn2461(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2462 github.com/goccy/spidermonkeywasm2go/p5.Fn2462
func Fn2462(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2475 github.com/goccy/spidermonkeywasm2go/p5.Fn2475
func Fn2475(m *base.Module, l0 int32) int32

//go:linkname Fn2489 github.com/goccy/spidermonkeywasm2go/p4.Fn2489
func Fn2489(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2495 github.com/goccy/spidermonkeywasm2go/p4.Fn2495
func Fn2495(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2497 github.com/goccy/spidermonkeywasm2go/p4.Fn2497
func Fn2497(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2514 github.com/goccy/spidermonkeywasm2go/p5.Fn2514
func Fn2514(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2534 github.com/goccy/spidermonkeywasm2go/p0.Fn2534
func Fn2534(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2541 github.com/goccy/spidermonkeywasm2go/p0.Fn2541
func Fn2541(m *base.Module, l0 int32) int32

//go:linkname Fn2545 github.com/goccy/spidermonkeywasm2go/p0.Fn2545
func Fn2545(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2547 github.com/goccy/spidermonkeywasm2go/p7.Fn2547
func Fn2547(m *base.Module, l0 int32)

//go:linkname Fn2550 github.com/goccy/spidermonkeywasm2go/p5.Fn2550
func Fn2550(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2553 github.com/goccy/spidermonkeywasm2go/p0.Fn2553
func Fn2553(m *base.Module, l0 int32) int32

//go:linkname Fn2555 github.com/goccy/spidermonkeywasm2go/p0.Fn2555
func Fn2555(m *base.Module, l0 int32) int32

//go:linkname Fn2560 github.com/goccy/spidermonkeywasm2go/p0.Fn2560
func Fn2560(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2562 github.com/goccy/spidermonkeywasm2go/p0.Fn2562
func Fn2562(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2563 github.com/goccy/spidermonkeywasm2go/p0.Fn2563
func Fn2563(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2570 github.com/goccy/spidermonkeywasm2go/p5.Fn2570
func Fn2570(m *base.Module, l0 int32)

//go:linkname Fn2571 github.com/goccy/spidermonkeywasm2go/p5.Fn2571
func Fn2571(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2579 github.com/goccy/spidermonkeywasm2go/p2.Fn2579
func Fn2579(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2580 github.com/goccy/spidermonkeywasm2go/p2.Fn2580
func Fn2580(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2581 github.com/goccy/spidermonkeywasm2go/p0.Fn2581
func Fn2581(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2582 github.com/goccy/spidermonkeywasm2go/p0.Fn2582
func Fn2582(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn2583 github.com/goccy/spidermonkeywasm2go/p0.Fn2583
func Fn2583(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2584 github.com/goccy/spidermonkeywasm2go/p0.Fn2584
func Fn2584(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2588 github.com/goccy/spidermonkeywasm2go/p0.Fn2588
func Fn2588(m *base.Module, l0 int32)

//go:linkname Fn2589 github.com/goccy/spidermonkeywasm2go/p0.Fn2589
func Fn2589(m *base.Module, l0 int32)

//go:linkname Fn2595 github.com/goccy/spidermonkeywasm2go/p7.Fn2595
func Fn2595(m *base.Module, l0 int32)

//go:linkname Fn2600 github.com/goccy/spidermonkeywasm2go/p0.Fn2600
func Fn2600(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2611 github.com/goccy/spidermonkeywasm2go/p0.Fn2611
func Fn2611(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2613 github.com/goccy/spidermonkeywasm2go/p0.Fn2613
func Fn2613(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2631 github.com/goccy/spidermonkeywasm2go/p0.Fn2631
func Fn2631(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn2636 github.com/goccy/spidermonkeywasm2go/p0.Fn2636
func Fn2636(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2665 github.com/goccy/spidermonkeywasm2go/p5.Fn2665
func Fn2665(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2720 github.com/goccy/spidermonkeywasm2go/p0.Fn2720
func Fn2720(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2727 github.com/goccy/spidermonkeywasm2go/p0.Fn2727
func Fn2727(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2731 github.com/goccy/spidermonkeywasm2go/p0.Fn2731
func Fn2731(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2732 github.com/goccy/spidermonkeywasm2go/p0.Fn2732
func Fn2732(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2733 github.com/goccy/spidermonkeywasm2go/p0.Fn2733
func Fn2733(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2735 github.com/goccy/spidermonkeywasm2go/p0.Fn2735
func Fn2735(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2736 github.com/goccy/spidermonkeywasm2go/p4.Fn2736
func Fn2736(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2742 github.com/goccy/spidermonkeywasm2go/p4.Fn2742
func Fn2742(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2746 github.com/goccy/spidermonkeywasm2go/p0.Fn2746
func Fn2746(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2754 github.com/goccy/spidermonkeywasm2go/p0.Fn2754
func Fn2754(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2755 github.com/goccy/spidermonkeywasm2go/p0.Fn2755
func Fn2755(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2760 github.com/goccy/spidermonkeywasm2go/p5.Fn2760
func Fn2760(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2764 github.com/goccy/spidermonkeywasm2go/p4.Fn2764
func Fn2764(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2766 github.com/goccy/spidermonkeywasm2go/p0.Fn2766
func Fn2766(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2773 github.com/goccy/spidermonkeywasm2go/p7.Fn2773
func Fn2773(m *base.Module, l0 int32) int32

//go:linkname Fn2791 github.com/goccy/spidermonkeywasm2go/p7.Fn2791
func Fn2791(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2797 github.com/goccy/spidermonkeywasm2go/p7.Fn2797
func Fn2797(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2805 github.com/goccy/spidermonkeywasm2go/p5.Fn2805
func Fn2805(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2819 github.com/goccy/spidermonkeywasm2go/p7.Fn2819
func Fn2819(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2835 github.com/goccy/spidermonkeywasm2go/p5.Fn2835
func Fn2835(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2847 github.com/goccy/spidermonkeywasm2go/p3.Fn2847
func Fn2847(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2862 github.com/goccy/spidermonkeywasm2go/p0.Fn2862
func Fn2862(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2866 github.com/goccy/spidermonkeywasm2go/p5.Fn2866
func Fn2866(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2867 github.com/goccy/spidermonkeywasm2go/p5.Fn2867
func Fn2867(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2871 github.com/goccy/spidermonkeywasm2go/p5.Fn2871
func Fn2871(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2878 github.com/goccy/spidermonkeywasm2go/p4.Fn2878
func Fn2878(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2879 github.com/goccy/spidermonkeywasm2go/p4.Fn2879
func Fn2879(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2881 github.com/goccy/spidermonkeywasm2go/p0.Fn2881
func Fn2881(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2886 github.com/goccy/spidermonkeywasm2go/p4.Fn2886
func Fn2886(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2887 github.com/goccy/spidermonkeywasm2go/p0.Fn2887
func Fn2887(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2893 github.com/goccy/spidermonkeywasm2go/p0.Fn2893
func Fn2893(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2897 github.com/goccy/spidermonkeywasm2go/p0.Fn2897
func Fn2897(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2900 github.com/goccy/spidermonkeywasm2go/p4.Fn2900
func Fn2900(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2904 github.com/goccy/spidermonkeywasm2go/p4.Fn2904
func Fn2904(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2916 github.com/goccy/spidermonkeywasm2go/p7.Fn2916
func Fn2916(m *base.Module, l0 int32)

//go:linkname Fn2918 github.com/goccy/spidermonkeywasm2go/p5.Fn2918
func Fn2918(m *base.Module, l0 int32)

//go:linkname Fn2940 github.com/goccy/spidermonkeywasm2go/p0.Fn2940
func Fn2940(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2942 github.com/goccy/spidermonkeywasm2go/p0.Fn2942
func Fn2942(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2943 github.com/goccy/spidermonkeywasm2go/p0.Fn2943
func Fn2943(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2953 github.com/goccy/spidermonkeywasm2go/p7.Fn2953
func Fn2953(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2958 github.com/goccy/spidermonkeywasm2go/p0.Fn2958
func Fn2958(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3005 github.com/goccy/spidermonkeywasm2go/p3.Fn3005
func Fn3005(m *base.Module, l0 int32) int32

//go:linkname Fn3008 github.com/goccy/spidermonkeywasm2go/p4.Fn3008
func Fn3008(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3045 github.com/goccy/spidermonkeywasm2go/p0.Fn3045
func Fn3045(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3046 github.com/goccy/spidermonkeywasm2go/p0.Fn3046
func Fn3046(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3047 github.com/goccy/spidermonkeywasm2go/p7.Fn3047
func Fn3047(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3056 github.com/goccy/spidermonkeywasm2go/p0.Fn3056
func Fn3056(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3103 github.com/goccy/spidermonkeywasm2go/p0.Fn3103
func Fn3103(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3106 github.com/goccy/spidermonkeywasm2go/p4.Fn3106
func Fn3106(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3140 github.com/goccy/spidermonkeywasm2go/p0.Fn3140
func Fn3140(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3144 github.com/goccy/spidermonkeywasm2go/p0.Fn3144
func Fn3144(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3164 github.com/goccy/spidermonkeywasm2go/p0.Fn3164
func Fn3164(m *base.Module, l0 int32) int32

//go:linkname Fn3167 github.com/goccy/spidermonkeywasm2go/p7.Fn3167
func Fn3167(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3168 github.com/goccy/spidermonkeywasm2go/p7.Fn3168
func Fn3168(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3177 github.com/goccy/spidermonkeywasm2go/p0.Fn3177
func Fn3177(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3179 github.com/goccy/spidermonkeywasm2go/p0.Fn3179
func Fn3179(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3200 github.com/goccy/spidermonkeywasm2go/p5.Fn3200
func Fn3200(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3206 github.com/goccy/spidermonkeywasm2go/p5.Fn3206
func Fn3206(m *base.Module, l0 int32)

//go:linkname Fn3214 github.com/goccy/spidermonkeywasm2go/p0.Fn3214
func Fn3214(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3222 github.com/goccy/spidermonkeywasm2go/p5.Fn3222
func Fn3222(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3239 github.com/goccy/spidermonkeywasm2go/p0.Fn3239
func Fn3239(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3257 github.com/goccy/spidermonkeywasm2go/p5.Fn3257
func Fn3257(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3411 github.com/goccy/spidermonkeywasm2go/p0.Fn3411
func Fn3411(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3412 github.com/goccy/spidermonkeywasm2go/p0.Fn3412
func Fn3412(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3413 github.com/goccy/spidermonkeywasm2go/p0.Fn3413
func Fn3413(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3414 github.com/goccy/spidermonkeywasm2go/p0.Fn3414
func Fn3414(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3415 github.com/goccy/spidermonkeywasm2go/p0.Fn3415
func Fn3415(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3416 github.com/goccy/spidermonkeywasm2go/p0.Fn3416
func Fn3416(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3417 github.com/goccy/spidermonkeywasm2go/p0.Fn3417
func Fn3417(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3418 github.com/goccy/spidermonkeywasm2go/p0.Fn3418
func Fn3418(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3419 github.com/goccy/spidermonkeywasm2go/p0.Fn3419
func Fn3419(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3420 github.com/goccy/spidermonkeywasm2go/p0.Fn3420
func Fn3420(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3488 github.com/goccy/spidermonkeywasm2go/p5.Fn3488
func Fn3488(m *base.Module, l0 int32)

//go:linkname Fn3491 github.com/goccy/spidermonkeywasm2go/p7.Fn3491
func Fn3491(m *base.Module, l0 int32)

//go:linkname Fn3503 github.com/goccy/spidermonkeywasm2go/p7.Fn3503
func Fn3503(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3510 github.com/goccy/spidermonkeywasm2go/p7.Fn3510
func Fn3510(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3511 github.com/goccy/spidermonkeywasm2go/p7.Fn3511
func Fn3511(m *base.Module, l0 int32)

//go:linkname Fn3512 github.com/goccy/spidermonkeywasm2go/p4.Fn3512
func Fn3512(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3513 github.com/goccy/spidermonkeywasm2go/p5.Fn3513
func Fn3513(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3534 github.com/goccy/spidermonkeywasm2go/p2.Fn3534
func Fn3534(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3543 github.com/goccy/spidermonkeywasm2go/p5.Fn3543
func Fn3543(m *base.Module, l0 int32) int32

//go:linkname Fn3550 github.com/goccy/spidermonkeywasm2go/p0.Fn3550
func Fn3550(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3551 github.com/goccy/spidermonkeywasm2go/p0.Fn3551
func Fn3551(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3552 github.com/goccy/spidermonkeywasm2go/p0.Fn3552
func Fn3552(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3558 github.com/goccy/spidermonkeywasm2go/p0.Fn3558
func Fn3558(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3561 github.com/goccy/spidermonkeywasm2go/p0.Fn3561
func Fn3561(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3562 github.com/goccy/spidermonkeywasm2go/p0.Fn3562
func Fn3562(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3566 github.com/goccy/spidermonkeywasm2go/p0.Fn3566
func Fn3566(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3575 github.com/goccy/spidermonkeywasm2go/p0.Fn3575
func Fn3575(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3576 github.com/goccy/spidermonkeywasm2go/p0.Fn3576
func Fn3576(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3579 github.com/goccy/spidermonkeywasm2go/p2.Fn3579
func Fn3579(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3602 github.com/goccy/spidermonkeywasm2go/p5.Fn3602
func Fn3602(m *base.Module, l0 int32)

//go:linkname Fn3617 github.com/goccy/spidermonkeywasm2go/p0.Fn3617
func Fn3617(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3620 github.com/goccy/spidermonkeywasm2go/p0.Fn3620
func Fn3620(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3633 github.com/goccy/spidermonkeywasm2go/p5.Fn3633
func Fn3633(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3644 github.com/goccy/spidermonkeywasm2go/p0.Fn3644
func Fn3644(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3645 github.com/goccy/spidermonkeywasm2go/p0.Fn3645
func Fn3645(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3648 github.com/goccy/spidermonkeywasm2go/p0.Fn3648
func Fn3648(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3654 github.com/goccy/spidermonkeywasm2go/p3.Fn3654
func Fn3654(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3661 github.com/goccy/spidermonkeywasm2go/p0.Fn3661
func Fn3661(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3662 github.com/goccy/spidermonkeywasm2go/p0.Fn3662
func Fn3662(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3665 github.com/goccy/spidermonkeywasm2go/p0.Fn3665
func Fn3665(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3668 github.com/goccy/spidermonkeywasm2go/p0.Fn3668
func Fn3668(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3721 github.com/goccy/spidermonkeywasm2go/p0.Fn3721
func Fn3721(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3745 github.com/goccy/spidermonkeywasm2go/p0.Fn3745
func Fn3745(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3748 github.com/goccy/spidermonkeywasm2go/p4.Fn3748
func Fn3748(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3753 github.com/goccy/spidermonkeywasm2go/p0.Fn3753
func Fn3753(m *base.Module, l0 int32, l1 int32, l2 int32) int32

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

//go:linkname Fn3837 github.com/goccy/spidermonkeywasm2go/p3.Fn3837
func Fn3837(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3841 github.com/goccy/spidermonkeywasm2go/p4.Fn3841
func Fn3841(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3843 github.com/goccy/spidermonkeywasm2go/p4.Fn3843
func Fn3843(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3845 github.com/goccy/spidermonkeywasm2go/p4.Fn3845
func Fn3845(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3847 github.com/goccy/spidermonkeywasm2go/p2.Fn3847
func Fn3847(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3849 github.com/goccy/spidermonkeywasm2go/p2.Fn3849
func Fn3849(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3851 github.com/goccy/spidermonkeywasm2go/p4.Fn3851
func Fn3851(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3853 github.com/goccy/spidermonkeywasm2go/p2.Fn3853
func Fn3853(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3858 github.com/goccy/spidermonkeywasm2go/p3.Fn3858
func Fn3858(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3860 github.com/goccy/spidermonkeywasm2go/p3.Fn3860
func Fn3860(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3862 github.com/goccy/spidermonkeywasm2go/p3.Fn3862
func Fn3862(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3879 github.com/goccy/spidermonkeywasm2go/p0.Fn3879
func Fn3879(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4105 github.com/goccy/spidermonkeywasm2go/p3.Fn4105
func Fn4105(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4109 github.com/goccy/spidermonkeywasm2go/p3.Fn4109
func Fn4109(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4112 github.com/goccy/spidermonkeywasm2go/p5.Fn4112
func Fn4112(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4114 github.com/goccy/spidermonkeywasm2go/p2.Fn4114
func Fn4114(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4116 github.com/goccy/spidermonkeywasm2go/p4.Fn4116
func Fn4116(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4140 github.com/goccy/spidermonkeywasm2go/p7.Fn4140
func Fn4140(m *base.Module, l0 int32)

//go:linkname Fn4141 github.com/goccy/spidermonkeywasm2go/p0.Fn4141
func Fn4141(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4142 github.com/goccy/spidermonkeywasm2go/p0.Fn4142
func Fn4142(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4157 github.com/goccy/spidermonkeywasm2go/p5.Fn4157
func Fn4157(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4158 github.com/goccy/spidermonkeywasm2go/p4.Fn4158
func Fn4158(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4161 github.com/goccy/spidermonkeywasm2go/p4.Fn4161
func Fn4161(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4164 github.com/goccy/spidermonkeywasm2go/p4.Fn4164
func Fn4164(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4180 github.com/goccy/spidermonkeywasm2go/p4.Fn4180
func Fn4180(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4190 github.com/goccy/spidermonkeywasm2go/p3.Fn4190
func Fn4190(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4217 github.com/goccy/spidermonkeywasm2go/p4.Fn4217
func Fn4217(m *base.Module, l0 int32)

//go:linkname Fn4260 github.com/goccy/spidermonkeywasm2go/p5.Fn4260
func Fn4260(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4271 github.com/goccy/spidermonkeywasm2go/p1.Fn4271
func Fn4271(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4274 github.com/goccy/spidermonkeywasm2go/p5.Fn4274
func Fn4274(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4275 github.com/goccy/spidermonkeywasm2go/p4.Fn4275
func Fn4275(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4280 github.com/goccy/spidermonkeywasm2go/p7.Fn4280
func Fn4280(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4302 github.com/goccy/spidermonkeywasm2go/p7.Fn4302
func Fn4302(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4320 github.com/goccy/spidermonkeywasm2go/p4.Fn4320
func Fn4320(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4330 github.com/goccy/spidermonkeywasm2go/p7.Fn4330
func Fn4330(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4332 github.com/goccy/spidermonkeywasm2go/p4.Fn4332
func Fn4332(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4341 github.com/goccy/spidermonkeywasm2go/p4.Fn4341
func Fn4341(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4351 github.com/goccy/spidermonkeywasm2go/p5.Fn4351
func Fn4351(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4375 github.com/goccy/spidermonkeywasm2go/p4.Fn4375
func Fn4375(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4379 github.com/goccy/spidermonkeywasm2go/p5.Fn4379
func Fn4379(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4392 github.com/goccy/spidermonkeywasm2go/p5.Fn4392
func Fn4392(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4409 github.com/goccy/spidermonkeywasm2go/p7.Fn4409
func Fn4409(m *base.Module, l0 int32)

//go:linkname Fn4417 github.com/goccy/spidermonkeywasm2go/p4.Fn4417
func Fn4417(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4420 github.com/goccy/spidermonkeywasm2go/p3.Fn4420
func Fn4420(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4422 github.com/goccy/spidermonkeywasm2go/p5.Fn4422
func Fn4422(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4450 github.com/goccy/spidermonkeywasm2go/p4.Fn4450
func Fn4450(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4479 github.com/goccy/spidermonkeywasm2go/p7.Fn4479
func Fn4479(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4507 github.com/goccy/spidermonkeywasm2go/p2.Fn4507
func Fn4507(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4508 github.com/goccy/spidermonkeywasm2go/p2.Fn4508
func Fn4508(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4513 github.com/goccy/spidermonkeywasm2go/p4.Fn4513
func Fn4513(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4514 github.com/goccy/spidermonkeywasm2go/p7.Fn4514
func Fn4514(m *base.Module, l0 int32)

//go:linkname Fn4518 github.com/goccy/spidermonkeywasm2go/p5.Fn4518
func Fn4518(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4519 github.com/goccy/spidermonkeywasm2go/p4.Fn4519
func Fn4519(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4525 github.com/goccy/spidermonkeywasm2go/p4.Fn4525
func Fn4525(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4529 github.com/goccy/spidermonkeywasm2go/p5.Fn4529
func Fn4529(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4536 github.com/goccy/spidermonkeywasm2go/p4.Fn4536
func Fn4536(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4575 github.com/goccy/spidermonkeywasm2go/p3.Fn4575
func Fn4575(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4578 github.com/goccy/spidermonkeywasm2go/p3.Fn4578
func Fn4578(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4581 github.com/goccy/spidermonkeywasm2go/p7.Fn4581
func Fn4581(m *base.Module, l0 int32) int32

//go:linkname Fn4642 github.com/goccy/spidermonkeywasm2go/p7.Fn4642
func Fn4642(m *base.Module, l0 int32)

//go:linkname Fn4669 github.com/goccy/spidermonkeywasm2go/p4.Fn4669
func Fn4669(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4697 github.com/goccy/spidermonkeywasm2go/p4.Fn4697
func Fn4697(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4702 github.com/goccy/spidermonkeywasm2go/p4.Fn4702
func Fn4702(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4711 github.com/goccy/spidermonkeywasm2go/p5.Fn4711
func Fn4711(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4718 github.com/goccy/spidermonkeywasm2go/p5.Fn4718
func Fn4718(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4719 github.com/goccy/spidermonkeywasm2go/p4.Fn4719
func Fn4719(m *base.Module, l0 int32)

//go:linkname Fn4721 github.com/goccy/spidermonkeywasm2go/p7.Fn4721
func Fn4721(m *base.Module, l0 int32)

//go:linkname Fn4738 github.com/goccy/spidermonkeywasm2go/p7.Fn4738
func Fn4738(m *base.Module, l0 int32)

//go:linkname Fn4767 github.com/goccy/spidermonkeywasm2go/p7.Fn4767
func Fn4767(m *base.Module, l0 int32)

//go:linkname Fn4809 github.com/goccy/spidermonkeywasm2go/p7.Fn4809
func Fn4809(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4825 github.com/goccy/spidermonkeywasm2go/p7.Fn4825
func Fn4825(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4830 github.com/goccy/spidermonkeywasm2go/p7.Fn4830
func Fn4830(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4867 github.com/goccy/spidermonkeywasm2go/p5.Fn4867
func Fn4867(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4875 github.com/goccy/spidermonkeywasm2go/p7.Fn4875
func Fn4875(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4880 github.com/goccy/spidermonkeywasm2go/p7.Fn4880
func Fn4880(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4913 github.com/goccy/spidermonkeywasm2go/p5.Fn4913
func Fn4913(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn5071 github.com/goccy/spidermonkeywasm2go/p5.Fn5071
func Fn5071(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn5138 github.com/goccy/spidermonkeywasm2go/p7.Fn5138
func Fn5138(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5141 github.com/goccy/spidermonkeywasm2go/p7.Fn5141
func Fn5141(m *base.Module, l0 int32)

//go:linkname Fn5142 github.com/goccy/spidermonkeywasm2go/p7.Fn5142
func Fn5142(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5288 github.com/goccy/spidermonkeywasm2go/p5.Fn5288
func Fn5288(m *base.Module, l0 int32)

//go:linkname Fn5289 github.com/goccy/spidermonkeywasm2go/p0.Fn5289
func Fn5289(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5333 github.com/goccy/spidermonkeywasm2go/p7.Fn5333
func Fn5333(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5334 github.com/goccy/spidermonkeywasm2go/p7.Fn5334
func Fn5334(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5337 github.com/goccy/spidermonkeywasm2go/p5.Fn5337
func Fn5337(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn5351 github.com/goccy/spidermonkeywasm2go/p5.Fn5351
func Fn5351(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5352 github.com/goccy/spidermonkeywasm2go/p4.Fn5352
func Fn5352(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5360 github.com/goccy/spidermonkeywasm2go/p4.Fn5360
func Fn5360(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5365 github.com/goccy/spidermonkeywasm2go/p4.Fn5365
func Fn5365(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5388 github.com/goccy/spidermonkeywasm2go/p5.Fn5388
func Fn5388(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5399 github.com/goccy/spidermonkeywasm2go/p4.Fn5399
func Fn5399(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5423 github.com/goccy/spidermonkeywasm2go/p5.Fn5423
func Fn5423(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5436 github.com/goccy/spidermonkeywasm2go/p7.Fn5436
func Fn5436(m *base.Module, l0 int32)

//go:linkname Fn5440 github.com/goccy/spidermonkeywasm2go/p5.Fn5440
func Fn5440(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5452 github.com/goccy/spidermonkeywasm2go/p0.Fn5452
func Fn5452(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5457 github.com/goccy/spidermonkeywasm2go/p5.Fn5457
func Fn5457(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5469 github.com/goccy/spidermonkeywasm2go/p0.Fn5469
func Fn5469(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5477 github.com/goccy/spidermonkeywasm2go/p0.Fn5477
func Fn5477(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5570 github.com/goccy/spidermonkeywasm2go/p7.Fn5570
func Fn5570(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5573 github.com/goccy/spidermonkeywasm2go/p7.Fn5573
func Fn5573(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5575 github.com/goccy/spidermonkeywasm2go/p4.Fn5575
func Fn5575(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5576 github.com/goccy/spidermonkeywasm2go/p0.Fn5576
func Fn5576(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5577 github.com/goccy/spidermonkeywasm2go/p0.Fn5577
func Fn5577(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5580 github.com/goccy/spidermonkeywasm2go/p5.Fn5580
func Fn5580(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5583 github.com/goccy/spidermonkeywasm2go/p5.Fn5583
func Fn5583(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5585 github.com/goccy/spidermonkeywasm2go/p4.Fn5585
func Fn5585(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5587 github.com/goccy/spidermonkeywasm2go/p7.Fn5587
func Fn5587(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5588 github.com/goccy/spidermonkeywasm2go/p5.Fn5588
func Fn5588(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5589 github.com/goccy/spidermonkeywasm2go/p4.Fn5589
func Fn5589(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5599 github.com/goccy/spidermonkeywasm2go/p5.Fn5599
func Fn5599(m *base.Module, l0 int32)

//go:linkname Fn5604 github.com/goccy/spidermonkeywasm2go/p4.Fn5604
func Fn5604(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5607 github.com/goccy/spidermonkeywasm2go/p5.Fn5607
func Fn5607(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5616 github.com/goccy/spidermonkeywasm2go/p5.Fn5616
func Fn5616(m *base.Module, l0 int32) int32

//go:linkname Fn5651 github.com/goccy/spidermonkeywasm2go/p7.Fn5651
func Fn5651(m *base.Module, l0 float64) int32

//go:linkname Fn5660 github.com/goccy/spidermonkeywasm2go/p5.Fn5660
func Fn5660(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5672 github.com/goccy/spidermonkeywasm2go/p7.Fn5672
func Fn5672(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5675 github.com/goccy/spidermonkeywasm2go/p0.Fn5675
func Fn5675(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5688 github.com/goccy/spidermonkeywasm2go/p5.Fn5688
func Fn5688(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5695 github.com/goccy/spidermonkeywasm2go/p5.Fn5695
func Fn5695(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5700 github.com/goccy/spidermonkeywasm2go/p5.Fn5700
func Fn5700(m *base.Module, l0 int32)

//go:linkname Fn5704 github.com/goccy/spidermonkeywasm2go/p7.Fn5704
func Fn5704(m *base.Module, l0 int32)

//go:linkname Fn5708 github.com/goccy/spidermonkeywasm2go/p7.Fn5708
func Fn5708(m *base.Module, l0 int32) int32

//go:linkname Fn5714 github.com/goccy/spidermonkeywasm2go/p4.Fn5714
func Fn5714(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5716 github.com/goccy/spidermonkeywasm2go/p5.Fn5716
func Fn5716(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5717 github.com/goccy/spidermonkeywasm2go/p7.Fn5717
func Fn5717(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5720 github.com/goccy/spidermonkeywasm2go/p5.Fn5720
func Fn5720(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5722 github.com/goccy/spidermonkeywasm2go/p5.Fn5722
func Fn5722(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5723 github.com/goccy/spidermonkeywasm2go/p5.Fn5723
func Fn5723(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5732 github.com/goccy/spidermonkeywasm2go/p5.Fn5732
func Fn5732(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5734 github.com/goccy/spidermonkeywasm2go/p7.Fn5734
func Fn5734(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5738 github.com/goccy/spidermonkeywasm2go/p7.Fn5738
func Fn5738(m *base.Module, l0 int32)

//go:linkname Fn5740 github.com/goccy/spidermonkeywasm2go/p7.Fn5740
func Fn5740(m *base.Module, l0 int32)

//go:linkname Fn5749 github.com/goccy/spidermonkeywasm2go/p3.Fn5749
func Fn5749(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5750 github.com/goccy/spidermonkeywasm2go/p3.Fn5750
func Fn5750(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5760 github.com/goccy/spidermonkeywasm2go/p7.Fn5760
func Fn5760(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5762 github.com/goccy/spidermonkeywasm2go/p5.Fn5762
func Fn5762(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5779 github.com/goccy/spidermonkeywasm2go/p5.Fn5779
func Fn5779(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5780 github.com/goccy/spidermonkeywasm2go/p3.Fn5780
func Fn5780(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5784 github.com/goccy/spidermonkeywasm2go/p5.Fn5784
func Fn5784(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5843 github.com/goccy/spidermonkeywasm2go/p7.Fn5843
func Fn5843(m *base.Module, l0 int32) int32

//go:linkname Fn5868 github.com/goccy/spidermonkeywasm2go/p2.Fn5868
func Fn5868(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5871 github.com/goccy/spidermonkeywasm2go/p3.Fn5871
func Fn5871(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5900 github.com/goccy/spidermonkeywasm2go/p5.Fn5900
func Fn5900(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5901 github.com/goccy/spidermonkeywasm2go/p5.Fn5901
func Fn5901(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5903 github.com/goccy/spidermonkeywasm2go/p3.Fn5903
func Fn5903(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5908 github.com/goccy/spidermonkeywasm2go/p5.Fn5908
func Fn5908(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5915 github.com/goccy/spidermonkeywasm2go/p4.Fn5915
func Fn5915(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5922 github.com/goccy/spidermonkeywasm2go/p7.Fn5922
func Fn5922(m *base.Module, l0 int32)

//go:linkname Fn5945 github.com/goccy/spidermonkeywasm2go/p4.Fn5945
func Fn5945(m *base.Module, l0 int32)

//go:linkname Fn5946 github.com/goccy/spidermonkeywasm2go/p7.Fn5946
func Fn5946(m *base.Module, l0 int32)

//go:linkname Fn5954 github.com/goccy/spidermonkeywasm2go/p5.Fn5954
func Fn5954(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5955 github.com/goccy/spidermonkeywasm2go/p7.Fn5955
func Fn5955(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5960 github.com/goccy/spidermonkeywasm2go/p2.Fn5960
func Fn5960(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5983 github.com/goccy/spidermonkeywasm2go/p7.Fn5983
func Fn5983(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5984 github.com/goccy/spidermonkeywasm2go/p7.Fn5984
func Fn5984(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5990 github.com/goccy/spidermonkeywasm2go/p5.Fn5990
func Fn5990(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6053 github.com/goccy/spidermonkeywasm2go/p5.Fn6053
func Fn6053(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6055 github.com/goccy/spidermonkeywasm2go/p5.Fn6055
func Fn6055(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6058 github.com/goccy/spidermonkeywasm2go/p7.Fn6058
func Fn6058(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6061 github.com/goccy/spidermonkeywasm2go/p3.Fn6061
func Fn6061(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6062 github.com/goccy/spidermonkeywasm2go/p7.Fn6062
func Fn6062(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6063 github.com/goccy/spidermonkeywasm2go/p7.Fn6063
func Fn6063(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6067 github.com/goccy/spidermonkeywasm2go/p7.Fn6067
func Fn6067(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6070 github.com/goccy/spidermonkeywasm2go/p7.Fn6070
func Fn6070(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6093 github.com/goccy/spidermonkeywasm2go/p7.Fn6093
func Fn6093(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6094 github.com/goccy/spidermonkeywasm2go/p7.Fn6094
func Fn6094(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6106 github.com/goccy/spidermonkeywasm2go/p7.Fn6106
func Fn6106(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6108 github.com/goccy/spidermonkeywasm2go/p7.Fn6108
func Fn6108(m *base.Module)

//go:linkname Fn6109 github.com/goccy/spidermonkeywasm2go/p7.Fn6109
func Fn6109(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6117 github.com/goccy/spidermonkeywasm2go/p2.Fn6117
func Fn6117(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6122 github.com/goccy/spidermonkeywasm2go/p7.Fn6122
func Fn6122(m *base.Module)

//go:linkname Fn6190 github.com/goccy/spidermonkeywasm2go/p3.Fn6190
func Fn6190(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6203 github.com/goccy/spidermonkeywasm2go/p3.Fn6203
func Fn6203(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6210 github.com/goccy/spidermonkeywasm2go/p5.Fn6210
func Fn6210(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6226 github.com/goccy/spidermonkeywasm2go/p2.Fn6226
func Fn6226(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6230 github.com/goccy/spidermonkeywasm2go/p5.Fn6230
func Fn6230(m *base.Module, l0 int32)

//go:linkname Fn6231 github.com/goccy/spidermonkeywasm2go/p5.Fn6231
func Fn6231(m *base.Module, l0 int32)

//go:linkname Fn6245 github.com/goccy/spidermonkeywasm2go/p5.Fn6245
func Fn6245(m *base.Module, l0 int32)

//go:linkname Fn6263 github.com/goccy/spidermonkeywasm2go/p2.Fn6263
func Fn6263(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn6269 github.com/goccy/spidermonkeywasm2go/p7.Fn6269
func Fn6269(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6273 github.com/goccy/spidermonkeywasm2go/p7.Fn6273
func Fn6273(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6277 github.com/goccy/spidermonkeywasm2go/p4.Fn6277
func Fn6277(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn6279 github.com/goccy/spidermonkeywasm2go/p4.Fn6279
func Fn6279(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6294 github.com/goccy/spidermonkeywasm2go/p5.Fn6294
func Fn6294(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6298 github.com/goccy/spidermonkeywasm2go/p2.Fn6298
func Fn6298(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn6362 github.com/goccy/spidermonkeywasm2go/p3.Fn6362
func Fn6362(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6666 github.com/goccy/spidermonkeywasm2go/p5.Fn6666
func Fn6666(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6668 github.com/goccy/spidermonkeywasm2go/p7.Fn6668
func Fn6668(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6687 github.com/goccy/spidermonkeywasm2go/p5.Fn6687
func Fn6687(m *base.Module, l0 int32)

//go:linkname Fn6696 github.com/goccy/spidermonkeywasm2go/p5.Fn6696
func Fn6696(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6724 github.com/goccy/spidermonkeywasm2go/p7.Fn6724
func Fn6724(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6725 github.com/goccy/spidermonkeywasm2go/p7.Fn6725
func Fn6725(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6729 github.com/goccy/spidermonkeywasm2go/p7.Fn6729
func Fn6729(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6730 github.com/goccy/spidermonkeywasm2go/p1.Fn6730
func Fn6730(m *base.Module, l0 int32) int32

//go:linkname Fn6734 github.com/goccy/spidermonkeywasm2go/p7.Fn6734
func Fn6734(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6768 github.com/goccy/spidermonkeywasm2go/p4.Fn6768
func Fn6768(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6775 github.com/goccy/spidermonkeywasm2go/p7.Fn6775
func Fn6775(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6780 github.com/goccy/spidermonkeywasm2go/p7.Fn6780
func Fn6780(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6781 github.com/goccy/spidermonkeywasm2go/p7.Fn6781
func Fn6781(m *base.Module, l0 int32) int32

//go:linkname Fn6787 github.com/goccy/spidermonkeywasm2go/p4.Fn6787
func Fn6787(m *base.Module, l0 int32) int32

//go:linkname Fn6792 github.com/goccy/spidermonkeywasm2go/p7.Fn6792
func Fn6792(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6794 github.com/goccy/spidermonkeywasm2go/p7.Fn6794
func Fn6794(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6816 github.com/goccy/spidermonkeywasm2go/p7.Fn6816
func Fn6816(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6864 github.com/goccy/spidermonkeywasm2go/p7.Fn6864
func Fn6864(m *base.Module, l0 int32, l1 int32, l2 int64)

//go:linkname Fn6866 github.com/goccy/spidermonkeywasm2go/p5.Fn6866
func Fn6866(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6876 github.com/goccy/spidermonkeywasm2go/p4.Fn6876
func Fn6876(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6879 github.com/goccy/spidermonkeywasm2go/p4.Fn6879
func Fn6879(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6899 github.com/goccy/spidermonkeywasm2go/p3.Fn6899
func Fn6899(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6902 github.com/goccy/spidermonkeywasm2go/p0.Fn6902
func Fn6902(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6903 github.com/goccy/spidermonkeywasm2go/p5.Fn6903
func Fn6903(m *base.Module, l0 int32, l1 int32, l2 int32) int64

//go:linkname Fn6904 github.com/goccy/spidermonkeywasm2go/p4.Fn6904
func Fn6904(m *base.Module, l0 int32)

//go:linkname Fn6914 github.com/goccy/spidermonkeywasm2go/p7.Fn6914
func Fn6914(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6917 github.com/goccy/spidermonkeywasm2go/p7.Fn6917
func Fn6917(m *base.Module, l0 int32) int32

//go:linkname Fn6922 github.com/goccy/spidermonkeywasm2go/p7.Fn6922
func Fn6922(m *base.Module, l0 int32) int32

//go:linkname Fn6950 github.com/goccy/spidermonkeywasm2go/p7.Fn6950
func Fn6950(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6952 github.com/goccy/spidermonkeywasm2go/p0.Fn6952
func Fn6952(m *base.Module, l0 int32) int32

//go:linkname Fn6955 github.com/goccy/spidermonkeywasm2go/p0.Fn6955
func Fn6955(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6956 github.com/goccy/spidermonkeywasm2go/p0.Fn6956
func Fn6956(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6961 github.com/goccy/spidermonkeywasm2go/p3.Fn6961
func Fn6961(m *base.Module, l0 int32) int32

//go:linkname Fn6962 github.com/goccy/spidermonkeywasm2go/p3.Fn6962
func Fn6962(m *base.Module, l0 int32) int32

//go:linkname Fn6970 github.com/goccy/spidermonkeywasm2go/p7.Fn6970
func Fn6970(m *base.Module, l0 int32)

//go:linkname Fn6974 github.com/goccy/spidermonkeywasm2go/p5.Fn6974
func Fn6974(m *base.Module, l0 int32) int32

//go:linkname Fn6982 github.com/goccy/spidermonkeywasm2go/p7.Fn6982
func Fn6982(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6983 github.com/goccy/spidermonkeywasm2go/p7.Fn6983
func Fn6983(m *base.Module, l0 int32) int32

//go:linkname Fn7000 github.com/goccy/spidermonkeywasm2go/p7.Fn7000
func Fn7000(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7013 github.com/goccy/spidermonkeywasm2go/p5.Fn7013
func Fn7013(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7015 github.com/goccy/spidermonkeywasm2go/p7.Fn7015
func Fn7015(m *base.Module, l0 int32)

//go:linkname Fn7035 github.com/goccy/spidermonkeywasm2go/p7.Fn7035
func Fn7035(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7050 github.com/goccy/spidermonkeywasm2go/p7.Fn7050
func Fn7050(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7052 github.com/goccy/spidermonkeywasm2go/p7.Fn7052
func Fn7052(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7085 github.com/goccy/spidermonkeywasm2go/p4.Fn7085
func Fn7085(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7086 github.com/goccy/spidermonkeywasm2go/p2.Fn7086
func Fn7086(m *base.Module, l0 int32) int32

//go:linkname Fn7100 github.com/goccy/spidermonkeywasm2go/p5.Fn7100
func Fn7100(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7103 github.com/goccy/spidermonkeywasm2go/p5.Fn7103
func Fn7103(m *base.Module, l0 int32) int32

//go:linkname Fn7121 github.com/goccy/spidermonkeywasm2go/p7.Fn7121
func Fn7121(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7128 github.com/goccy/spidermonkeywasm2go/p5.Fn7128
func Fn7128(m *base.Module, l0 int32) int32

//go:linkname Fn7149 github.com/goccy/spidermonkeywasm2go/p5.Fn7149
func Fn7149(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7150 github.com/goccy/spidermonkeywasm2go/p4.Fn7150
func Fn7150(m *base.Module, l0 int32)

//go:linkname Fn7152 github.com/goccy/spidermonkeywasm2go/p5.Fn7152
func Fn7152(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn7153 github.com/goccy/spidermonkeywasm2go/p7.Fn7153
func Fn7153(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7171 github.com/goccy/spidermonkeywasm2go/p7.Fn7171
func Fn7171(m *base.Module, l0 int32)

//go:linkname Fn7174 github.com/goccy/spidermonkeywasm2go/p7.Fn7174
func Fn7174(m *base.Module, l0 int32)

//go:linkname Fn7192 github.com/goccy/spidermonkeywasm2go/p2.Fn7192
func Fn7192(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7204 github.com/goccy/spidermonkeywasm2go/p5.Fn7204
func Fn7204(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7209 github.com/goccy/spidermonkeywasm2go/p4.Fn7209
func Fn7209(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7310 github.com/goccy/spidermonkeywasm2go/p5.Fn7310
func Fn7310(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7318 github.com/goccy/spidermonkeywasm2go/p5.Fn7318
func Fn7318(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7398 github.com/goccy/spidermonkeywasm2go/p4.Fn7398
func Fn7398(m *base.Module, l0 int32) int32

//go:linkname Fn7401 github.com/goccy/spidermonkeywasm2go/p5.Fn7401
func Fn7401(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7405 github.com/goccy/spidermonkeywasm2go/p5.Fn7405
func Fn7405(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7428 github.com/goccy/spidermonkeywasm2go/p5.Fn7428
func Fn7428(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7440 github.com/goccy/spidermonkeywasm2go/p4.Fn7440
func Fn7440(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7443 github.com/goccy/spidermonkeywasm2go/p3.Fn7443
func Fn7443(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7445 github.com/goccy/spidermonkeywasm2go/p5.Fn7445
func Fn7445(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7448 github.com/goccy/spidermonkeywasm2go/p5.Fn7448
func Fn7448(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7453 github.com/goccy/spidermonkeywasm2go/p3.Fn7453
func Fn7453(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7456 github.com/goccy/spidermonkeywasm2go/p5.Fn7456
func Fn7456(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7457 github.com/goccy/spidermonkeywasm2go/p7.Fn7457
func Fn7457(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7458 github.com/goccy/spidermonkeywasm2go/p7.Fn7458
func Fn7458(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7464 github.com/goccy/spidermonkeywasm2go/p3.Fn7464
func Fn7464(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7471 github.com/goccy/spidermonkeywasm2go/p5.Fn7471
func Fn7471(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7473 github.com/goccy/spidermonkeywasm2go/p5.Fn7473
func Fn7473(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7482 github.com/goccy/spidermonkeywasm2go/p3.Fn7482
func Fn7482(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7484 github.com/goccy/spidermonkeywasm2go/p7.Fn7484
func Fn7484(m *base.Module, l0 int32, l1 int32, l2 float64)

//go:linkname Fn7486 github.com/goccy/spidermonkeywasm2go/p5.Fn7486
func Fn7486(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7489 github.com/goccy/spidermonkeywasm2go/p5.Fn7489
func Fn7489(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn7492 github.com/goccy/spidermonkeywasm2go/p3.Fn7492
func Fn7492(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7507 github.com/goccy/spidermonkeywasm2go/p4.Fn7507
func Fn7507(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7513 github.com/goccy/spidermonkeywasm2go/p3.Fn7513
func Fn7513(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) int32

//go:linkname Fn7566 github.com/goccy/spidermonkeywasm2go/p3.Fn7566
func Fn7566(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7593 github.com/goccy/spidermonkeywasm2go/p4.Fn7593
func Fn7593(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7626 github.com/goccy/spidermonkeywasm2go/p3.Fn7626
func Fn7626(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7699 github.com/goccy/spidermonkeywasm2go/p5.Fn7699
func Fn7699(m *base.Module, l0 int32, l1 float64, l2 float64, l3 float64, l4 float64, l5 float64, l6 float64) int32

//go:linkname Fn7709 github.com/goccy/spidermonkeywasm2go/p3.Fn7709
func Fn7709(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7725 github.com/goccy/spidermonkeywasm2go/p3.Fn7725
func Fn7725(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7726 github.com/goccy/spidermonkeywasm2go/p5.Fn7726
func Fn7726(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7729 github.com/goccy/spidermonkeywasm2go/p5.Fn7729
func Fn7729(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7730 github.com/goccy/spidermonkeywasm2go/p5.Fn7730
func Fn7730(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7739 github.com/goccy/spidermonkeywasm2go/p3.Fn7739
func Fn7739(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7821 github.com/goccy/spidermonkeywasm2go/p3.Fn7821
func Fn7821(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7879 github.com/goccy/spidermonkeywasm2go/p3.Fn7879
func Fn7879(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7922 github.com/goccy/spidermonkeywasm2go/p3.Fn7922
func Fn7922(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7938 github.com/goccy/spidermonkeywasm2go/p5.Fn7938
func Fn7938(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7947 github.com/goccy/spidermonkeywasm2go/p5.Fn7947
func Fn7947(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7950 github.com/goccy/spidermonkeywasm2go/p5.Fn7950
func Fn7950(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7957 github.com/goccy/spidermonkeywasm2go/p4.Fn7957
func Fn7957(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7960 github.com/goccy/spidermonkeywasm2go/p5.Fn7960
func Fn7960(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7962 github.com/goccy/spidermonkeywasm2go/p5.Fn7962
func Fn7962(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7963 github.com/goccy/spidermonkeywasm2go/p5.Fn7963
func Fn7963(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7967 github.com/goccy/spidermonkeywasm2go/p4.Fn7967
func Fn7967(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7968 github.com/goccy/spidermonkeywasm2go/p5.Fn7968
func Fn7968(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7974 github.com/goccy/spidermonkeywasm2go/p4.Fn7974
func Fn7974(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7978 github.com/goccy/spidermonkeywasm2go/p5.Fn7978
func Fn7978(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7979 github.com/goccy/spidermonkeywasm2go/p4.Fn7979
func Fn7979(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn7981 github.com/goccy/spidermonkeywasm2go/p7.Fn7981
func Fn7981(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7984 github.com/goccy/spidermonkeywasm2go/p4.Fn7984
func Fn7984(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8001 github.com/goccy/spidermonkeywasm2go/p3.Fn8001
func Fn8001(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8100 github.com/goccy/spidermonkeywasm2go/p7.Fn8100
func Fn8100(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8119 github.com/goccy/spidermonkeywasm2go/p7.Fn8119
func Fn8119(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8149 github.com/goccy/spidermonkeywasm2go/p7.Fn8149
func Fn8149(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8151 github.com/goccy/spidermonkeywasm2go/p5.Fn8151
func Fn8151(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8152 github.com/goccy/spidermonkeywasm2go/p4.Fn8152
func Fn8152(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8157 github.com/goccy/spidermonkeywasm2go/p1.Fn8157
func Fn8157(m *base.Module, l0 int32) int32

//go:linkname Fn8161 github.com/goccy/spidermonkeywasm2go/p1.Fn8161
func Fn8161(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8162 github.com/goccy/spidermonkeywasm2go/p1.Fn8162
func Fn8162(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8167 github.com/goccy/spidermonkeywasm2go/p1.Fn8167
func Fn8167(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8174 github.com/goccy/spidermonkeywasm2go/p2.Fn8174
func Fn8174(m *base.Module) int32

//go:linkname Fn8179 github.com/goccy/spidermonkeywasm2go/p3.Fn8179
func Fn8179(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8215 github.com/goccy/spidermonkeywasm2go/p4.Fn8215
func Fn8215(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8270 github.com/goccy/spidermonkeywasm2go/p5.Fn8270
func Fn8270(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8271 github.com/goccy/spidermonkeywasm2go/p4.Fn8271
func Fn8271(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8272 github.com/goccy/spidermonkeywasm2go/p5.Fn8272
func Fn8272(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8277 github.com/goccy/spidermonkeywasm2go/p5.Fn8277
func Fn8277(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8278 github.com/goccy/spidermonkeywasm2go/p2.Fn8278
func Fn8278(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8282 github.com/goccy/spidermonkeywasm2go/p5.Fn8282
func Fn8282(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8287 github.com/goccy/spidermonkeywasm2go/p5.Fn8287
func Fn8287(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8290 github.com/goccy/spidermonkeywasm2go/p3.Fn8290
func Fn8290(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8294 github.com/goccy/spidermonkeywasm2go/p4.Fn8294
func Fn8294(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8305 github.com/goccy/spidermonkeywasm2go/p5.Fn8305
func Fn8305(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8327 github.com/goccy/spidermonkeywasm2go/p4.Fn8327
func Fn8327(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8328 github.com/goccy/spidermonkeywasm2go/p4.Fn8328
func Fn8328(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8330 github.com/goccy/spidermonkeywasm2go/p5.Fn8330
func Fn8330(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8331 github.com/goccy/spidermonkeywasm2go/p5.Fn8331
func Fn8331(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8343 github.com/goccy/spidermonkeywasm2go/p4.Fn8343
func Fn8343(m *base.Module, l0 int32) int32

//go:linkname Fn8355 github.com/goccy/spidermonkeywasm2go/p3.Fn8355
func Fn8355(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8362 github.com/goccy/spidermonkeywasm2go/p7.Fn8362
func Fn8362(m *base.Module, l0 int32)

//go:linkname Fn8369 github.com/goccy/spidermonkeywasm2go/p5.Fn8369
func Fn8369(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8371 github.com/goccy/spidermonkeywasm2go/p5.Fn8371
func Fn8371(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8384 github.com/goccy/spidermonkeywasm2go/p5.Fn8384
func Fn8384(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn8391 github.com/goccy/spidermonkeywasm2go/p5.Fn8391
func Fn8391(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8393 github.com/goccy/spidermonkeywasm2go/p1.Fn8393
func Fn8393(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8406 github.com/goccy/spidermonkeywasm2go/p7.Fn8406
func Fn8406(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8431 github.com/goccy/spidermonkeywasm2go/p5.Fn8431
func Fn8431(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8439 github.com/goccy/spidermonkeywasm2go/p5.Fn8439
func Fn8439(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8442 github.com/goccy/spidermonkeywasm2go/p5.Fn8442
func Fn8442(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8458 github.com/goccy/spidermonkeywasm2go/p4.Fn8458
func Fn8458(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8468 github.com/goccy/spidermonkeywasm2go/p5.Fn8468
func Fn8468(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8474 github.com/goccy/spidermonkeywasm2go/p5.Fn8474
func Fn8474(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8477 github.com/goccy/spidermonkeywasm2go/p1.Fn8477
func Fn8477(m *base.Module, l0 int32) int32

//go:linkname Fn8483 github.com/goccy/spidermonkeywasm2go/p3.Fn8483
func Fn8483(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8506 github.com/goccy/spidermonkeywasm2go/p2.Fn8506
func Fn8506(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn8525 github.com/goccy/spidermonkeywasm2go/p5.Fn8525
func Fn8525(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8526 github.com/goccy/spidermonkeywasm2go/p5.Fn8526
func Fn8526(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8529 github.com/goccy/spidermonkeywasm2go/p7.Fn8529
func Fn8529(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8530 github.com/goccy/spidermonkeywasm2go/p5.Fn8530
func Fn8530(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8535 github.com/goccy/spidermonkeywasm2go/p5.Fn8535
func Fn8535(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8557 github.com/goccy/spidermonkeywasm2go/p7.Fn8557
func Fn8557(m *base.Module, l0 int32) int32

//go:linkname Fn8567 github.com/goccy/spidermonkeywasm2go/p7.Fn8567
func Fn8567(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8573 github.com/goccy/spidermonkeywasm2go/p4.Fn8573
func Fn8573(m *base.Module, l0 int32) int32

//go:linkname Fn8576 github.com/goccy/spidermonkeywasm2go/p7.Fn8576
func Fn8576(m *base.Module, l0 int32) int32

//go:linkname Fn8628 github.com/goccy/spidermonkeywasm2go/p7.Fn8628
func Fn8628(m *base.Module, l0 int32)

//go:linkname Fn8629 github.com/goccy/spidermonkeywasm2go/p7.Fn8629
func Fn8629(m *base.Module, l0 int32)

//go:linkname Fn8634 github.com/goccy/spidermonkeywasm2go/p5.Fn8634
func Fn8634(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8637 github.com/goccy/spidermonkeywasm2go/p1.Fn8637
func Fn8637(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn8639 github.com/goccy/spidermonkeywasm2go/p1.Fn8639
func Fn8639(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8641 github.com/goccy/spidermonkeywasm2go/p1.Fn8641
func Fn8641(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8644 github.com/goccy/spidermonkeywasm2go/p5.Fn8644
func Fn8644(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8646 github.com/goccy/spidermonkeywasm2go/p1.Fn8646
func Fn8646(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8648 github.com/goccy/spidermonkeywasm2go/p1.Fn8648
func Fn8648(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8657 github.com/goccy/spidermonkeywasm2go/p2.Fn8657
func Fn8657(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn8683 github.com/goccy/spidermonkeywasm2go/p5.Fn8683
func Fn8683(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8689 github.com/goccy/spidermonkeywasm2go/p7.Fn8689
func Fn8689(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn8692 github.com/goccy/spidermonkeywasm2go/p3.Fn8692
func Fn8692(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn8694 github.com/goccy/spidermonkeywasm2go/p4.Fn8694
func Fn8694(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8708 github.com/goccy/spidermonkeywasm2go/p4.Fn8708
func Fn8708(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8720 github.com/goccy/spidermonkeywasm2go/p5.Fn8720
func Fn8720(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8723 github.com/goccy/spidermonkeywasm2go/p7.Fn8723
func Fn8723(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8726 github.com/goccy/spidermonkeywasm2go/p7.Fn8726
func Fn8726(m *base.Module, l0 int32)

//go:linkname Fn8743 github.com/goccy/spidermonkeywasm2go/p5.Fn8743
func Fn8743(m *base.Module, l0 int32) int32

//go:linkname Fn8744 github.com/goccy/spidermonkeywasm2go/p5.Fn8744
func Fn8744(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8745 github.com/goccy/spidermonkeywasm2go/p3.Fn8745
func Fn8745(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8757 github.com/goccy/spidermonkeywasm2go/p5.Fn8757
func Fn8757(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8759 github.com/goccy/spidermonkeywasm2go/p3.Fn8759
func Fn8759(m *base.Module, l0 int32) int32

//go:linkname Fn8765 github.com/goccy/spidermonkeywasm2go/p7.Fn8765
func Fn8765(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8766 github.com/goccy/spidermonkeywasm2go/p7.Fn8766
func Fn8766(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8776 github.com/goccy/spidermonkeywasm2go/p4.Fn8776
func Fn8776(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8799 github.com/goccy/spidermonkeywasm2go/p4.Fn8799
func Fn8799(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8844 github.com/goccy/spidermonkeywasm2go/p5.Fn8844
func Fn8844(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8858 github.com/goccy/spidermonkeywasm2go/p5.Fn8858
func Fn8858(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8859 github.com/goccy/spidermonkeywasm2go/p4.Fn8859
func Fn8859(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8924 github.com/goccy/spidermonkeywasm2go/p4.Fn8924
func Fn8924(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8933 github.com/goccy/spidermonkeywasm2go/p5.Fn8933
func Fn8933(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8936 github.com/goccy/spidermonkeywasm2go/p5.Fn8936
func Fn8936(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8937 github.com/goccy/spidermonkeywasm2go/p7.Fn8937
func Fn8937(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8962 github.com/goccy/spidermonkeywasm2go/p7.Fn8962
func Fn8962(m *base.Module, l0 int32)

//go:linkname Fn9011 github.com/goccy/spidermonkeywasm2go/p4.Fn9011
func Fn9011(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9022 github.com/goccy/spidermonkeywasm2go/p5.Fn9022
func Fn9022(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32)

//go:linkname Fn9024 github.com/goccy/spidermonkeywasm2go/p5.Fn9024
func Fn9024(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9071 github.com/goccy/spidermonkeywasm2go/p2.Fn9071
func Fn9071(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9072 github.com/goccy/spidermonkeywasm2go/p3.Fn9072
func Fn9072(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn9080 github.com/goccy/spidermonkeywasm2go/p4.Fn9080
func Fn9080(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9086 github.com/goccy/spidermonkeywasm2go/p3.Fn9086
func Fn9086(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9129 github.com/goccy/spidermonkeywasm2go/p5.Fn9129
func Fn9129(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9132 github.com/goccy/spidermonkeywasm2go/p5.Fn9132
func Fn9132(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9159 github.com/goccy/spidermonkeywasm2go/p5.Fn9159
func Fn9159(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9171 github.com/goccy/spidermonkeywasm2go/p5.Fn9171
func Fn9171(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9182 github.com/goccy/spidermonkeywasm2go/p7.Fn9182
func Fn9182(m *base.Module, l0 float64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn9277 github.com/goccy/spidermonkeywasm2go/p3.Fn9277
func Fn9277(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9291 github.com/goccy/spidermonkeywasm2go/p4.Fn9291
func Fn9291(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9295 github.com/goccy/spidermonkeywasm2go/p5.Fn9295
func Fn9295(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9302 github.com/goccy/spidermonkeywasm2go/p7.Fn9302
func Fn9302(m *base.Module, l0 int32) int32

//go:linkname Fn9304 github.com/goccy/spidermonkeywasm2go/p7.Fn9304
func Fn9304(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9308 github.com/goccy/spidermonkeywasm2go/p5.Fn9308
func Fn9308(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9313 github.com/goccy/spidermonkeywasm2go/p5.Fn9313
func Fn9313(m *base.Module, l0 int32)

//go:linkname Fn9317 github.com/goccy/spidermonkeywasm2go/p3.Fn9317
func Fn9317(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9321 github.com/goccy/spidermonkeywasm2go/p5.Fn9321
func Fn9321(m *base.Module, l0 int32) float64

//go:linkname Fn9340 github.com/goccy/spidermonkeywasm2go/p3.Fn9340
func Fn9340(m *base.Module, l0 int32) int32

//go:linkname Fn9343 github.com/goccy/spidermonkeywasm2go/p5.Fn9343
func Fn9343(m *base.Module, l0 int32)

//go:linkname Fn9344 github.com/goccy/spidermonkeywasm2go/p3.Fn9344
func Fn9344(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9358 github.com/goccy/spidermonkeywasm2go/p3.Fn9358
func Fn9358(m *base.Module, l0 int32)

//go:linkname Fn9359 github.com/goccy/spidermonkeywasm2go/p4.Fn9359
func Fn9359(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9360 github.com/goccy/spidermonkeywasm2go/p5.Fn9360
func Fn9360(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9365 github.com/goccy/spidermonkeywasm2go/p5.Fn9365
func Fn9365(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9366 github.com/goccy/spidermonkeywasm2go/p5.Fn9366
func Fn9366(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9405 github.com/goccy/spidermonkeywasm2go/p7.Fn9405
func Fn9405(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9468 github.com/goccy/spidermonkeywasm2go/p3.Fn9468
func Fn9468(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9502 github.com/goccy/spidermonkeywasm2go/p5.Fn9502
func Fn9502(m *base.Module, l0 int32)

//go:linkname Fn9530 github.com/goccy/spidermonkeywasm2go/p5.Fn9530
func Fn9530(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn9531 github.com/goccy/spidermonkeywasm2go/p5.Fn9531
func Fn9531(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9558 github.com/goccy/spidermonkeywasm2go/p4.Fn9558
func Fn9558(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9574 github.com/goccy/spidermonkeywasm2go/p7.Fn9574
func Fn9574(m *base.Module, l0 int32) int32

//go:linkname Fn9597 github.com/goccy/spidermonkeywasm2go/p5.Fn9597
func Fn9597(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9611 github.com/goccy/spidermonkeywasm2go/p3.Fn9611
func Fn9611(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9628 github.com/goccy/spidermonkeywasm2go/p2.Fn9628
func Fn9628(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9629 github.com/goccy/spidermonkeywasm2go/p5.Fn9629
func Fn9629(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn9630 github.com/goccy/spidermonkeywasm2go/p7.Fn9630
func Fn9630(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9632 github.com/goccy/spidermonkeywasm2go/p5.Fn9632
func Fn9632(m *base.Module, l0 int32)

//go:linkname Fn9637 github.com/goccy/spidermonkeywasm2go/p7.Fn9637
func Fn9637(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9638 github.com/goccy/spidermonkeywasm2go/p5.Fn9638
func Fn9638(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9647 github.com/goccy/spidermonkeywasm2go/p4.Fn9647
func Fn9647(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9654 github.com/goccy/spidermonkeywasm2go/p2.Fn9654
func Fn9654(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9690 github.com/goccy/spidermonkeywasm2go/p7.Fn9690
func Fn9690(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9693 github.com/goccy/spidermonkeywasm2go/p4.Fn9693
func Fn9693(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9697 github.com/goccy/spidermonkeywasm2go/p5.Fn9697
func Fn9697(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9701 github.com/goccy/spidermonkeywasm2go/p5.Fn9701
func Fn9701(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9704 github.com/goccy/spidermonkeywasm2go/p5.Fn9704
func Fn9704(m *base.Module, l0 int32, l1 float64, l2 float64, l3 int32, l4 int32)

//go:linkname Fn9708 github.com/goccy/spidermonkeywasm2go/p7.Fn9708
func Fn9708(m *base.Module, l0 int32)

//go:linkname Fn9709 github.com/goccy/spidermonkeywasm2go/p2.Fn9709
func Fn9709(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9713 github.com/goccy/spidermonkeywasm2go/p4.Fn9713
func Fn9713(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9734 github.com/goccy/spidermonkeywasm2go/p7.Fn9734
func Fn9734(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9735 github.com/goccy/spidermonkeywasm2go/p3.Fn9735
func Fn9735(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9768 github.com/goccy/spidermonkeywasm2go/p2.Fn9768
func Fn9768(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9772 github.com/goccy/spidermonkeywasm2go/p7.Fn9772
func Fn9772(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9777 github.com/goccy/spidermonkeywasm2go/p5.Fn9777
func Fn9777(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9783 github.com/goccy/spidermonkeywasm2go/p7.Fn9783
func Fn9783(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9816 github.com/goccy/spidermonkeywasm2go/p4.Fn9816
func Fn9816(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn9844 github.com/goccy/spidermonkeywasm2go/p4.Fn9844
func Fn9844(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9856 github.com/goccy/spidermonkeywasm2go/p2.Fn9856
func Fn9856(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9947 github.com/goccy/spidermonkeywasm2go/p7.Fn9947
func Fn9947(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9948 github.com/goccy/spidermonkeywasm2go/p3.Fn9948
func Fn9948(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9977 github.com/goccy/spidermonkeywasm2go/p4.Fn9977
func Fn9977(m *base.Module, l0 int32)

//go:linkname Fn9991 github.com/goccy/spidermonkeywasm2go/p5.Fn9991
func Fn9991(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10013 github.com/goccy/spidermonkeywasm2go/p5.Fn10013
func Fn10013(m *base.Module, l0 int32, l1 float64) int32

//go:linkname Fn10014 github.com/goccy/spidermonkeywasm2go/p4.Fn10014
func Fn10014(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 int32, l5 int32, l6 int32)

//go:linkname Fn10023 github.com/goccy/spidermonkeywasm2go/p4.Fn10023
func Fn10023(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn10024 github.com/goccy/spidermonkeywasm2go/p4.Fn10024
func Fn10024(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn10025 github.com/goccy/spidermonkeywasm2go/p4.Fn10025
func Fn10025(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn10026 github.com/goccy/spidermonkeywasm2go/p3.Fn10026
func Fn10026(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10087 github.com/goccy/spidermonkeywasm2go/p4.Fn10087
func Fn10087(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn10097 github.com/goccy/spidermonkeywasm2go/p5.Fn10097
func Fn10097(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10098 github.com/goccy/spidermonkeywasm2go/p7.Fn10098
func Fn10098(m *base.Module, l0 int32) int32

//go:linkname Fn10125 github.com/goccy/spidermonkeywasm2go/p2.Fn10125
func Fn10125(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10127 github.com/goccy/spidermonkeywasm2go/p5.Fn10127
func Fn10127(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10145 github.com/goccy/spidermonkeywasm2go/p7.Fn10145
func Fn10145(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10146 github.com/goccy/spidermonkeywasm2go/p3.Fn10146
func Fn10146(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10158 github.com/goccy/spidermonkeywasm2go/p5.Fn10158
func Fn10158(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10169 github.com/goccy/spidermonkeywasm2go/p5.Fn10169
func Fn10169(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10189 github.com/goccy/spidermonkeywasm2go/p4.Fn10189
func Fn10189(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10193 github.com/goccy/spidermonkeywasm2go/p5.Fn10193
func Fn10193(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10202 github.com/goccy/spidermonkeywasm2go/p3.Fn10202
func Fn10202(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10214 github.com/goccy/spidermonkeywasm2go/p5.Fn10214
func Fn10214(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10215 github.com/goccy/spidermonkeywasm2go/p5.Fn10215
func Fn10215(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10252 github.com/goccy/spidermonkeywasm2go/p5.Fn10252
func Fn10252(m *base.Module, l0 int32) int32

//go:linkname Fn10340 github.com/goccy/spidermonkeywasm2go/p2.Fn10340
func Fn10340(m *base.Module, l0 int32)

//go:linkname Fn10343 github.com/goccy/spidermonkeywasm2go/p1.Fn10343
func Fn10343(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10345 github.com/goccy/spidermonkeywasm2go/p5.Fn10345
func Fn10345(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10417 github.com/goccy/spidermonkeywasm2go/p5.Fn10417
func Fn10417(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10440 github.com/goccy/spidermonkeywasm2go/p5.Fn10440
func Fn10440(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10457 github.com/goccy/spidermonkeywasm2go/p3.Fn10457
func Fn10457(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10474 github.com/goccy/spidermonkeywasm2go/p5.Fn10474
func Fn10474(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10577 github.com/goccy/spidermonkeywasm2go/p5.Fn10577
func Fn10577(m *base.Module) int32

//go:linkname Fn10582 github.com/goccy/spidermonkeywasm2go/p3.Fn10582
func Fn10582(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10584 github.com/goccy/spidermonkeywasm2go/p5.Fn10584
func Fn10584(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10586 github.com/goccy/spidermonkeywasm2go/p5.Fn10586
func Fn10586(m *base.Module) int32

//go:linkname Fn10614 github.com/goccy/spidermonkeywasm2go/p5.Fn10614
func Fn10614(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10621 github.com/goccy/spidermonkeywasm2go/p4.Fn10621
func Fn10621(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10687 github.com/goccy/spidermonkeywasm2go/p4.Fn10687
func Fn10687(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10688 github.com/goccy/spidermonkeywasm2go/p5.Fn10688
func Fn10688(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn10700 github.com/goccy/spidermonkeywasm2go/p5.Fn10700
func Fn10700(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10701 github.com/goccy/spidermonkeywasm2go/p5.Fn10701
func Fn10701(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10713 github.com/goccy/spidermonkeywasm2go/p5.Fn10713
func Fn10713(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10716 github.com/goccy/spidermonkeywasm2go/p4.Fn10716
func Fn10716(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10733 github.com/goccy/spidermonkeywasm2go/p4.Fn10733
func Fn10733(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10792 github.com/goccy/spidermonkeywasm2go/p4.Fn10792
func Fn10792(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn10813 github.com/goccy/spidermonkeywasm2go/p5.Fn10813
func Fn10813(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10819 github.com/goccy/spidermonkeywasm2go/p5.Fn10819
func Fn10819(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn10824 github.com/goccy/spidermonkeywasm2go/p3.Fn10824
func Fn10824(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32)

//go:linkname Fn10825 github.com/goccy/spidermonkeywasm2go/p7.Fn10825
func Fn10825(m *base.Module, l0 int32)

//go:linkname Fn10827 github.com/goccy/spidermonkeywasm2go/p5.Fn10827
func Fn10827(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10830 github.com/goccy/spidermonkeywasm2go/p4.Fn10830
func Fn10830(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10858 github.com/goccy/spidermonkeywasm2go/p5.Fn10858
func Fn10858(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10885 github.com/goccy/spidermonkeywasm2go/p4.Fn10885
func Fn10885(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10889 github.com/goccy/spidermonkeywasm2go/p4.Fn10889
func Fn10889(m *base.Module, l0 int32) int32

//go:linkname Fn10891 github.com/goccy/spidermonkeywasm2go/p5.Fn10891
func Fn10891(m *base.Module)

//go:linkname Fn10893 github.com/goccy/spidermonkeywasm2go/p5.Fn10893
func Fn10893(m *base.Module, l0 int32) int32

//go:linkname Fn10901 github.com/goccy/spidermonkeywasm2go/p4.Fn10901
func Fn10901(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn10929 github.com/goccy/spidermonkeywasm2go/p2.Fn10929
func Fn10929(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10941 github.com/goccy/spidermonkeywasm2go/p7.Fn10941
func Fn10941(m *base.Module)

//go:linkname Fn10942 github.com/goccy/spidermonkeywasm2go/p2.Fn10942
func Fn10942(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10943 github.com/goccy/spidermonkeywasm2go/p2.Fn10943
func Fn10943(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10944 github.com/goccy/spidermonkeywasm2go/p2.Fn10944
func Fn10944(m *base.Module, l0 float64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn10947 github.com/goccy/spidermonkeywasm2go/p5.Fn10947
func Fn10947(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10962 github.com/goccy/spidermonkeywasm2go/p2.Fn10962
func Fn10962(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64

//go:linkname Fn10968 github.com/goccy/spidermonkeywasm2go/p2.Fn10968
func Fn10968(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64

//go:linkname Fn10972 github.com/goccy/spidermonkeywasm2go/p2.Fn10972
func Fn10972(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10976 github.com/goccy/spidermonkeywasm2go/p7.Fn10976
func Fn10976(m *base.Module, l0 int32)
