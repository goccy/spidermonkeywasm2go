//go:build (!amd64 && !arm64) || purego

package p5

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

//go:linkname Fn33 github.com/goccy/spidermonkeywasm2go/p7.Fn33
func Fn33(m *base.Module, l0 int32) int64

//go:linkname Fn34 github.com/goccy/spidermonkeywasm2go/p7.Fn34
func Fn34(m *base.Module, l0 int32)

//go:linkname Fn36 github.com/goccy/spidermonkeywasm2go/p7.Fn36
func Fn36(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn41 github.com/goccy/spidermonkeywasm2go/p7.Fn41
func Fn41(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn42 github.com/goccy/spidermonkeywasm2go/p7.Fn42
func Fn42(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn73 github.com/goccy/spidermonkeywasm2go/p7.Fn73
func Fn73(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn74 github.com/goccy/spidermonkeywasm2go/p7.Fn74
func Fn74(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn75 github.com/goccy/spidermonkeywasm2go/p7.Fn75
func Fn75(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn78 github.com/goccy/spidermonkeywasm2go/p7.Fn78
func Fn78(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn80 github.com/goccy/spidermonkeywasm2go/p7.Fn80
func Fn80(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn81 github.com/goccy/spidermonkeywasm2go/p6.Fn81
func Fn81(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn83 github.com/goccy/spidermonkeywasm2go/p6.Fn83
func Fn83(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn85 github.com/goccy/spidermonkeywasm2go/p7.Fn85
func Fn85(m *base.Module, l0 int32)

//go:linkname Fn86 github.com/goccy/spidermonkeywasm2go/p7.Fn86
func Fn86(m *base.Module, l0 int32)

//go:linkname Fn87 github.com/goccy/spidermonkeywasm2go/p7.Fn87
func Fn87(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn88 github.com/goccy/spidermonkeywasm2go/p6.Fn88
func Fn88(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn89 github.com/goccy/spidermonkeywasm2go/p7.Fn89
func Fn89(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn92 github.com/goccy/spidermonkeywasm2go/p7.Fn92
func Fn92(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn102 github.com/goccy/spidermonkeywasm2go/p7.Fn102
func Fn102(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn105 github.com/goccy/spidermonkeywasm2go/p7.Fn105
func Fn105(m *base.Module, l0 int32)

//go:linkname Fn109 github.com/goccy/spidermonkeywasm2go/p7.Fn109
func Fn109(m *base.Module, l0 int32) int32

//go:linkname Fn118 github.com/goccy/spidermonkeywasm2go/p7.Fn118
func Fn118(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn120 github.com/goccy/spidermonkeywasm2go/p7.Fn120
func Fn120(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn123 github.com/goccy/spidermonkeywasm2go/p7.Fn123
func Fn123(m *base.Module, l0 int32)

//go:linkname Fn129 github.com/goccy/spidermonkeywasm2go/p6.Fn129
func Fn129(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn130 github.com/goccy/spidermonkeywasm2go/p7.Fn130
func Fn130(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn135 github.com/goccy/spidermonkeywasm2go/p7.Fn135
func Fn135(m *base.Module)

//go:linkname Fn138 github.com/goccy/spidermonkeywasm2go/p7.Fn138
func Fn138(m *base.Module, l0 int32)

//go:linkname Fn139 github.com/goccy/spidermonkeywasm2go/p7.Fn139
func Fn139(m *base.Module)

//go:linkname Fn140 github.com/goccy/spidermonkeywasm2go/p7.Fn140
func Fn140(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn146 github.com/goccy/spidermonkeywasm2go/p6.Fn146
func Fn146(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn147 github.com/goccy/spidermonkeywasm2go/p7.Fn147
func Fn147(m *base.Module, l0 int32)

//go:linkname Fn148 github.com/goccy/spidermonkeywasm2go/p7.Fn148
func Fn148(m *base.Module, l0 int32)

//go:linkname Fn151 github.com/goccy/spidermonkeywasm2go/p7.Fn151
func Fn151(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn152 github.com/goccy/spidermonkeywasm2go/p7.Fn152
func Fn152(m *base.Module, l0 int32) int32

//go:linkname Fn154 github.com/goccy/spidermonkeywasm2go/p7.Fn154
func Fn154(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn155 github.com/goccy/spidermonkeywasm2go/p0.Fn155
func Fn155(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn157 github.com/goccy/spidermonkeywasm2go/p7.Fn157
func Fn157(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn158 github.com/goccy/spidermonkeywasm2go/p7.Fn158
func Fn158(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn161 github.com/goccy/spidermonkeywasm2go/p3.Fn161
func Fn161(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn166 github.com/goccy/spidermonkeywasm2go/p4.Fn166
func Fn166(m *base.Module, l0 float64) float64

//go:linkname Fn167 github.com/goccy/spidermonkeywasm2go/p3.Fn167
func Fn167(m *base.Module, l0 float64) float64

//go:linkname Fn171 github.com/goccy/spidermonkeywasm2go/p3.Fn171
func Fn171(m *base.Module, l0 float64) float64

//go:linkname Fn172 github.com/goccy/spidermonkeywasm2go/p4.Fn172
func Fn172(m *base.Module, l0 int32) float64

//go:linkname Fn175 github.com/goccy/spidermonkeywasm2go/p3.Fn175
func Fn175(m *base.Module, l0 float64, l1 int32) int32

//go:linkname Fn184 github.com/goccy/spidermonkeywasm2go/p6.Fn184
func Fn184(m *base.Module, l0 int32) int32

//go:linkname Fn185 github.com/goccy/spidermonkeywasm2go/p7.Fn185
func Fn185(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn189 github.com/goccy/spidermonkeywasm2go/p3.Fn189
func Fn189(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn195 github.com/goccy/spidermonkeywasm2go/p4.Fn195
func Fn195(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn197 github.com/goccy/spidermonkeywasm2go/p7.Fn197
func Fn197(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn198 github.com/goccy/spidermonkeywasm2go/p7.Fn198
func Fn198(m *base.Module, l0 int32)

//go:linkname Fn202 github.com/goccy/spidermonkeywasm2go/p7.Fn202
func Fn202(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn207 github.com/goccy/spidermonkeywasm2go/p6.Fn207
func Fn207(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn217 github.com/goccy/spidermonkeywasm2go/p7.Fn217
func Fn217(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn234 github.com/goccy/spidermonkeywasm2go/p7.Fn234
func Fn234(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn235 github.com/goccy/spidermonkeywasm2go/p7.Fn235
func Fn235(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn236 github.com/goccy/spidermonkeywasm2go/p7.Fn236
func Fn236(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn245 github.com/goccy/spidermonkeywasm2go/p7.Fn245
func Fn245(m *base.Module, l0 int32) int32

//go:linkname Fn246 github.com/goccy/spidermonkeywasm2go/p7.Fn246
func Fn246(m *base.Module, l0 int32)

//go:linkname Fn249 github.com/goccy/spidermonkeywasm2go/p2.Fn249
func Fn249(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn250 github.com/goccy/spidermonkeywasm2go/p3.Fn250
func Fn250(m *base.Module, l0 int32) int32

//go:linkname Fn252 github.com/goccy/spidermonkeywasm2go/p6.Fn252
func Fn252(m *base.Module, l0 int32) int32

//go:linkname Fn253 github.com/goccy/spidermonkeywasm2go/p7.Fn253
func Fn253(m *base.Module, l0 int32) int32

//go:linkname Fn254 github.com/goccy/spidermonkeywasm2go/p7.Fn254
func Fn254(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn257 github.com/goccy/spidermonkeywasm2go/p6.Fn257
func Fn257(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn258 github.com/goccy/spidermonkeywasm2go/p6.Fn258
func Fn258(m *base.Module, l0 int32) int32

//go:linkname Fn259 github.com/goccy/spidermonkeywasm2go/p4.Fn259
func Fn259(m *base.Module, l0 int32) int32

//go:linkname Fn261 github.com/goccy/spidermonkeywasm2go/p7.Fn261
func Fn261(m *base.Module, l0 int32)

//go:linkname Fn289 github.com/goccy/spidermonkeywasm2go/p7.Fn289
func Fn289(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn304 github.com/goccy/spidermonkeywasm2go/p2.Fn304
func Fn304(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn340 github.com/goccy/spidermonkeywasm2go/p7.Fn340
func Fn340(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn342 github.com/goccy/spidermonkeywasm2go/p6.Fn342
func Fn342(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn367 github.com/goccy/spidermonkeywasm2go/p7.Fn367
func Fn367(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn368 github.com/goccy/spidermonkeywasm2go/p4.Fn368
func Fn368(m *base.Module, l0 int32) int32

//go:linkname Fn369 github.com/goccy/spidermonkeywasm2go/p7.Fn369
func Fn369(m *base.Module)

//go:linkname Fn370 github.com/goccy/spidermonkeywasm2go/p7.Fn370
func Fn370(m *base.Module)

//go:linkname Fn371 github.com/goccy/spidermonkeywasm2go/p7.Fn371
func Fn371(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn373 github.com/goccy/spidermonkeywasm2go/p7.Fn373
func Fn373(m *base.Module, l0 int32) int32

//go:linkname Fn375 github.com/goccy/spidermonkeywasm2go/p7.Fn375
func Fn375(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn378 github.com/goccy/spidermonkeywasm2go/p6.Fn378
func Fn378(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn384 github.com/goccy/spidermonkeywasm2go/p6.Fn384
func Fn384(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn387 github.com/goccy/spidermonkeywasm2go/p7.Fn387
func Fn387(m *base.Module, l0 int32)

//go:linkname Fn388 github.com/goccy/spidermonkeywasm2go/p7.Fn388
func Fn388(m *base.Module, l0 int32)

//go:linkname Fn389 github.com/goccy/spidermonkeywasm2go/p7.Fn389
func Fn389(m *base.Module, l0 int32)

//go:linkname Fn391 github.com/goccy/spidermonkeywasm2go/p7.Fn391
func Fn391(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn426 github.com/goccy/spidermonkeywasm2go/p7.Fn426
func Fn426(m *base.Module)

//go:linkname Fn493 github.com/goccy/spidermonkeywasm2go/p6.Fn493
func Fn493(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn500 github.com/goccy/spidermonkeywasm2go/p7.Fn500
func Fn500(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn501 github.com/goccy/spidermonkeywasm2go/p7.Fn501
func Fn501(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn502 github.com/goccy/spidermonkeywasm2go/p4.Fn502
func Fn502(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn508 github.com/goccy/spidermonkeywasm2go/p6.Fn508
func Fn508(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn514 github.com/goccy/spidermonkeywasm2go/p4.Fn514
func Fn514(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn543 github.com/goccy/spidermonkeywasm2go/p2.Fn543
func Fn543(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32

//go:linkname Fn546 github.com/goccy/spidermonkeywasm2go/p2.Fn546
func Fn546(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32

//go:linkname Fn549 github.com/goccy/spidermonkeywasm2go/p6.Fn549
func Fn549(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)

//go:linkname Fn553 github.com/goccy/spidermonkeywasm2go/p6.Fn553
func Fn553(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32)

//go:linkname Fn565 github.com/goccy/spidermonkeywasm2go/p7.Fn565
func Fn565(m *base.Module, l0 int32)

//go:linkname Fn652 github.com/goccy/spidermonkeywasm2go/p7.Fn652
func Fn652(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn653 github.com/goccy/spidermonkeywasm2go/p7.Fn653
func Fn653(m *base.Module)

//go:linkname Fn654 github.com/goccy/spidermonkeywasm2go/p7.Fn654
func Fn654(m *base.Module, l0 int32) int32

//go:linkname Fn655 github.com/goccy/spidermonkeywasm2go/p7.Fn655
func Fn655(m *base.Module, l0 int32)

//go:linkname Fn659 github.com/goccy/spidermonkeywasm2go/p6.Fn659
func Fn659(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn669 github.com/goccy/spidermonkeywasm2go/p7.Fn669
func Fn669(m *base.Module, l0 int32) int32

//go:linkname Fn671 github.com/goccy/spidermonkeywasm2go/p7.Fn671
func Fn671(m *base.Module, l0 int32)

//go:linkname Fn672 github.com/goccy/spidermonkeywasm2go/p7.Fn672
func Fn672(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn673 github.com/goccy/spidermonkeywasm2go/p7.Fn673
func Fn673(m *base.Module, l0 int32) int64

//go:linkname Fn674 github.com/goccy/spidermonkeywasm2go/p7.Fn674
func Fn674(m *base.Module, l0 int32) int32

//go:linkname Fn676 github.com/goccy/spidermonkeywasm2go/p7.Fn676
func Fn676(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn681 github.com/goccy/spidermonkeywasm2go/p6.Fn681
func Fn681(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn682 github.com/goccy/spidermonkeywasm2go/p6.Fn682
func Fn682(m *base.Module, l0 int32) int32

//go:linkname Fn685 github.com/goccy/spidermonkeywasm2go/p6.Fn685
func Fn685(m *base.Module, l0 int32)

//go:linkname Fn686 github.com/goccy/spidermonkeywasm2go/p7.Fn686
func Fn686(m *base.Module, l0 int32) int32

//go:linkname Fn691 github.com/goccy/spidermonkeywasm2go/p7.Fn691
func Fn691(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn692 github.com/goccy/spidermonkeywasm2go/p7.Fn692
func Fn692(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn694 github.com/goccy/spidermonkeywasm2go/p6.Fn694
func Fn694(m *base.Module, l0 int32) int64

//go:linkname Fn697 github.com/goccy/spidermonkeywasm2go/p7.Fn697
func Fn697(m *base.Module, l0 int32)

//go:linkname Fn707 github.com/goccy/spidermonkeywasm2go/p7.Fn707
func Fn707(m *base.Module, l0 int32) int32

//go:linkname Fn711 github.com/goccy/spidermonkeywasm2go/p7.Fn711
func Fn711(m *base.Module, l0 float64) float64

//go:linkname Fn715 github.com/goccy/spidermonkeywasm2go/p7.Fn715
func Fn715(m *base.Module, l0 float64) float64

//go:linkname Fn717 github.com/goccy/spidermonkeywasm2go/p7.Fn717
func Fn717(m *base.Module, l0 float64) float64

//go:linkname Fn720 github.com/goccy/spidermonkeywasm2go/p6.Fn720
func Fn720(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn723 github.com/goccy/spidermonkeywasm2go/p6.Fn723
func Fn723(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn724 github.com/goccy/spidermonkeywasm2go/p7.Fn724
func Fn724(m *base.Module, l0 int32)

//go:linkname Fn726 github.com/goccy/spidermonkeywasm2go/p6.Fn726
func Fn726(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn727 github.com/goccy/spidermonkeywasm2go/p7.Fn727
func Fn727(m *base.Module, l0 int32) int32

//go:linkname Fn734 github.com/goccy/spidermonkeywasm2go/p7.Fn734
func Fn734(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn737 github.com/goccy/spidermonkeywasm2go/p6.Fn737
func Fn737(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn738 github.com/goccy/spidermonkeywasm2go/p7.Fn738
func Fn738(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn746 github.com/goccy/spidermonkeywasm2go/p6.Fn746
func Fn746(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn750 github.com/goccy/spidermonkeywasm2go/p7.Fn750
func Fn750(m *base.Module) int32

//go:linkname Fn751 github.com/goccy/spidermonkeywasm2go/p7.Fn751
func Fn751(m *base.Module, l0 int32)

//go:linkname Fn754 github.com/goccy/spidermonkeywasm2go/p7.Fn754
func Fn754(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn755 github.com/goccy/spidermonkeywasm2go/p3.Fn755
func Fn755(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn762 github.com/goccy/spidermonkeywasm2go/p7.Fn762
func Fn762(m *base.Module, l0 int32) int32

//go:linkname Fn763 github.com/goccy/spidermonkeywasm2go/p7.Fn763
func Fn763(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn764 github.com/goccy/spidermonkeywasm2go/p6.Fn764
func Fn764(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn767 github.com/goccy/spidermonkeywasm2go/p7.Fn767
func Fn767(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn772 github.com/goccy/spidermonkeywasm2go/p7.Fn772
func Fn772(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn773 github.com/goccy/spidermonkeywasm2go/p6.Fn773
func Fn773(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn774 github.com/goccy/spidermonkeywasm2go/p7.Fn774
func Fn774(m *base.Module, l0 int32, l1 int32, l2 int32) int32

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

//go:linkname Fn789 github.com/goccy/spidermonkeywasm2go/p7.Fn789
func Fn789(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn792 github.com/goccy/spidermonkeywasm2go/p6.Fn792
func Fn792(m *base.Module, l0 int32)

//go:linkname Fn793 github.com/goccy/spidermonkeywasm2go/p7.Fn793
func Fn793(m *base.Module, l0 int32)

//go:linkname Fn795 github.com/goccy/spidermonkeywasm2go/p6.Fn795
func Fn795(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn797 github.com/goccy/spidermonkeywasm2go/p7.Fn797
func Fn797(m *base.Module, l0 int32) int32

//go:linkname Fn798 github.com/goccy/spidermonkeywasm2go/p7.Fn798
func Fn798(m *base.Module, l0 int32) int32

//go:linkname Fn799 github.com/goccy/spidermonkeywasm2go/p7.Fn799
func Fn799(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn800 github.com/goccy/spidermonkeywasm2go/p7.Fn800
func Fn800(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn803 github.com/goccy/spidermonkeywasm2go/p7.Fn803
func Fn803(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn805 github.com/goccy/spidermonkeywasm2go/p6.Fn805
func Fn805(m *base.Module, l0 int32) int32

//go:linkname Fn806 github.com/goccy/spidermonkeywasm2go/p6.Fn806
func Fn806(m *base.Module, l0 int32) int32

//go:linkname Fn808 github.com/goccy/spidermonkeywasm2go/p6.Fn808
func Fn808(m *base.Module, l0 int32) int32

//go:linkname Fn809 github.com/goccy/spidermonkeywasm2go/p6.Fn809
func Fn809(m *base.Module)

//go:linkname Fn811 github.com/goccy/spidermonkeywasm2go/p7.Fn811
func Fn811(m *base.Module)

//go:linkname Fn821 github.com/goccy/spidermonkeywasm2go/p7.Fn821
func Fn821(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn822 github.com/goccy/spidermonkeywasm2go/p7.Fn822
func Fn822(m *base.Module)

//go:linkname Fn824 github.com/goccy/spidermonkeywasm2go/p2.Fn824
func Fn824(m *base.Module, l0 int32) int32

//go:linkname Fn825 github.com/goccy/spidermonkeywasm2go/p4.Fn825
func Fn825(m *base.Module, l0 int32)

//go:linkname Fn826 github.com/goccy/spidermonkeywasm2go/p7.Fn826
func Fn826(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn828 github.com/goccy/spidermonkeywasm2go/p4.Fn828
func Fn828(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn830 github.com/goccy/spidermonkeywasm2go/p6.Fn830
func Fn830(m *base.Module, l0 int32) int32

//go:linkname Fn832 github.com/goccy/spidermonkeywasm2go/p0.Fn832
func Fn832(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn835 github.com/goccy/spidermonkeywasm2go/p6.Fn835
func Fn835(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn838 github.com/goccy/spidermonkeywasm2go/p0.Fn838
func Fn838(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn839 github.com/goccy/spidermonkeywasm2go/p0.Fn839
func Fn839(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn842 github.com/goccy/spidermonkeywasm2go/p0.Fn842
func Fn842(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn847 github.com/goccy/spidermonkeywasm2go/p7.Fn847
func Fn847(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn850 github.com/goccy/spidermonkeywasm2go/p0.Fn850
func Fn850(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn851 github.com/goccy/spidermonkeywasm2go/p0.Fn851
func Fn851(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn852 github.com/goccy/spidermonkeywasm2go/p0.Fn852
func Fn852(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn854 github.com/goccy/spidermonkeywasm2go/p0.Fn854
func Fn854(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn867 github.com/goccy/spidermonkeywasm2go/p6.Fn867
func Fn867(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn868 github.com/goccy/spidermonkeywasm2go/p6.Fn868
func Fn868(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn871 github.com/goccy/spidermonkeywasm2go/p0.Fn871
func Fn871(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn873 github.com/goccy/spidermonkeywasm2go/p6.Fn873
func Fn873(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn875 github.com/goccy/spidermonkeywasm2go/p0.Fn875
func Fn875(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn883 github.com/goccy/spidermonkeywasm2go/p0.Fn883
func Fn883(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn886 github.com/goccy/spidermonkeywasm2go/p0.Fn886
func Fn886(m *base.Module, l0 int32) int32

//go:linkname Fn887 github.com/goccy/spidermonkeywasm2go/p0.Fn887
func Fn887(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn890 github.com/goccy/spidermonkeywasm2go/p7.Fn890
func Fn890(m *base.Module, l0 int32) int32

//go:linkname Fn904 github.com/goccy/spidermonkeywasm2go/p6.Fn904
func Fn904(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn916 github.com/goccy/spidermonkeywasm2go/p6.Fn916
func Fn916(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn932 github.com/goccy/spidermonkeywasm2go/p7.Fn932
func Fn932(m *base.Module, l0 int32) int32

//go:linkname Fn937 github.com/goccy/spidermonkeywasm2go/p7.Fn937
func Fn937(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1057 github.com/goccy/spidermonkeywasm2go/p6.Fn1057
func Fn1057(m *base.Module, l0 int32) int32

//go:linkname Fn1086 github.com/goccy/spidermonkeywasm2go/p0.Fn1086
func Fn1086(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1087 github.com/goccy/spidermonkeywasm2go/p0.Fn1087
func Fn1087(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1088 github.com/goccy/spidermonkeywasm2go/p0.Fn1088
func Fn1088(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1090 github.com/goccy/spidermonkeywasm2go/p0.Fn1090
func Fn1090(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1108 github.com/goccy/spidermonkeywasm2go/p0.Fn1108
func Fn1108(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1109 github.com/goccy/spidermonkeywasm2go/p0.Fn1109
func Fn1109(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1117 github.com/goccy/spidermonkeywasm2go/p0.Fn1117
func Fn1117(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1120 github.com/goccy/spidermonkeywasm2go/p0.Fn1120
func Fn1120(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1129 github.com/goccy/spidermonkeywasm2go/p6.Fn1129
func Fn1129(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1132 github.com/goccy/spidermonkeywasm2go/p0.Fn1132
func Fn1132(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1138 github.com/goccy/spidermonkeywasm2go/p4.Fn1138
func Fn1138(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1162 github.com/goccy/spidermonkeywasm2go/p6.Fn1162
func Fn1162(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1164 github.com/goccy/spidermonkeywasm2go/p6.Fn1164
func Fn1164(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1172 github.com/goccy/spidermonkeywasm2go/p6.Fn1172
func Fn1172(m *base.Module, l0 int32)

//go:linkname Fn1177 github.com/goccy/spidermonkeywasm2go/p4.Fn1177
func Fn1177(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn1209 github.com/goccy/spidermonkeywasm2go/p6.Fn1209
func Fn1209(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1281 github.com/goccy/spidermonkeywasm2go/p4.Fn1281
func Fn1281(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1282 github.com/goccy/spidermonkeywasm2go/p6.Fn1282
func Fn1282(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1298 github.com/goccy/spidermonkeywasm2go/p0.Fn1298
func Fn1298(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1306 github.com/goccy/spidermonkeywasm2go/p6.Fn1306
func Fn1306(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1314 github.com/goccy/spidermonkeywasm2go/p6.Fn1314
func Fn1314(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1315 github.com/goccy/spidermonkeywasm2go/p6.Fn1315
func Fn1315(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1353 github.com/goccy/spidermonkeywasm2go/p7.Fn1353
func Fn1353(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1357 github.com/goccy/spidermonkeywasm2go/p3.Fn1357
func Fn1357(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1386 github.com/goccy/spidermonkeywasm2go/p7.Fn1386
func Fn1386(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1387 github.com/goccy/spidermonkeywasm2go/p4.Fn1387
func Fn1387(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1404 github.com/goccy/spidermonkeywasm2go/p0.Fn1404
func Fn1404(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1408 github.com/goccy/spidermonkeywasm2go/p7.Fn1408
func Fn1408(m *base.Module, l0 int32) int32

//go:linkname Fn1425 github.com/goccy/spidermonkeywasm2go/p7.Fn1425
func Fn1425(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1434 github.com/goccy/spidermonkeywasm2go/p6.Fn1434
func Fn1434(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn1436 github.com/goccy/spidermonkeywasm2go/p0.Fn1436
func Fn1436(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1442 github.com/goccy/spidermonkeywasm2go/p3.Fn1442
func Fn1442(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1443 github.com/goccy/spidermonkeywasm2go/p0.Fn1443
func Fn1443(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1444 github.com/goccy/spidermonkeywasm2go/p6.Fn1444
func Fn1444(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1449 github.com/goccy/spidermonkeywasm2go/p0.Fn1449
func Fn1449(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1454 github.com/goccy/spidermonkeywasm2go/p6.Fn1454
func Fn1454(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1456 github.com/goccy/spidermonkeywasm2go/p3.Fn1456
func Fn1456(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1457 github.com/goccy/spidermonkeywasm2go/p6.Fn1457
func Fn1457(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn1459 github.com/goccy/spidermonkeywasm2go/p6.Fn1459
func Fn1459(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1461 github.com/goccy/spidermonkeywasm2go/p6.Fn1461
func Fn1461(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1462 github.com/goccy/spidermonkeywasm2go/p7.Fn1462
func Fn1462(m *base.Module, l0 int32)

//go:linkname Fn1463 github.com/goccy/spidermonkeywasm2go/p6.Fn1463
func Fn1463(m *base.Module, l0 int32)

//go:linkname Fn1465 github.com/goccy/spidermonkeywasm2go/p6.Fn1465
func Fn1465(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1466 github.com/goccy/spidermonkeywasm2go/p7.Fn1466
func Fn1466(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1468 github.com/goccy/spidermonkeywasm2go/p6.Fn1468
func Fn1468(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1470 github.com/goccy/spidermonkeywasm2go/p7.Fn1470
func Fn1470(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1472 github.com/goccy/spidermonkeywasm2go/p0.Fn1472
func Fn1472(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1473 github.com/goccy/spidermonkeywasm2go/p0.Fn1473
func Fn1473(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1478 github.com/goccy/spidermonkeywasm2go/p7.Fn1478
func Fn1478(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1479 github.com/goccy/spidermonkeywasm2go/p7.Fn1479
func Fn1479(m *base.Module, l0 int32)

//go:linkname Fn1491 github.com/goccy/spidermonkeywasm2go/p6.Fn1491
func Fn1491(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1495 github.com/goccy/spidermonkeywasm2go/p0.Fn1495
func Fn1495(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1496 github.com/goccy/spidermonkeywasm2go/p7.Fn1496
func Fn1496(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1497 github.com/goccy/spidermonkeywasm2go/p7.Fn1497
func Fn1497(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1498 github.com/goccy/spidermonkeywasm2go/p0.Fn1498
func Fn1498(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1499 github.com/goccy/spidermonkeywasm2go/p7.Fn1499
func Fn1499(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1501 github.com/goccy/spidermonkeywasm2go/p7.Fn1501
func Fn1501(m *base.Module, l0 int32) int32

//go:linkname Fn1502 github.com/goccy/spidermonkeywasm2go/p7.Fn1502
func Fn1502(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1503 github.com/goccy/spidermonkeywasm2go/p7.Fn1503
func Fn1503(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1504 github.com/goccy/spidermonkeywasm2go/p7.Fn1504
func Fn1504(m *base.Module, l0 int32)

//go:linkname Fn1506 github.com/goccy/spidermonkeywasm2go/p6.Fn1506
func Fn1506(m *base.Module, l0 int32) int32

//go:linkname Fn1509 github.com/goccy/spidermonkeywasm2go/p7.Fn1509
func Fn1509(m *base.Module, l0 int32) int32

//go:linkname Fn1511 github.com/goccy/spidermonkeywasm2go/p7.Fn1511
func Fn1511(m *base.Module) float64

//go:linkname Fn1590 github.com/goccy/spidermonkeywasm2go/p3.Fn1590
func Fn1590(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1615 github.com/goccy/spidermonkeywasm2go/p6.Fn1615
func Fn1615(m *base.Module, l0 int32)

//go:linkname Fn1616 github.com/goccy/spidermonkeywasm2go/p2.Fn1616
func Fn1616(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1619 github.com/goccy/spidermonkeywasm2go/p4.Fn1619
func Fn1619(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1620 github.com/goccy/spidermonkeywasm2go/p0.Fn1620
func Fn1620(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1623 github.com/goccy/spidermonkeywasm2go/p6.Fn1623
func Fn1623(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1626 github.com/goccy/spidermonkeywasm2go/p6.Fn1626
func Fn1626(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1627 github.com/goccy/spidermonkeywasm2go/p7.Fn1627
func Fn1627(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1630 github.com/goccy/spidermonkeywasm2go/p6.Fn1630
func Fn1630(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1632 github.com/goccy/spidermonkeywasm2go/p0.Fn1632
func Fn1632(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1642 github.com/goccy/spidermonkeywasm2go/p6.Fn1642
func Fn1642(m *base.Module, l0 int32, l1 float64, l2 int32) int32

//go:linkname Fn1644 github.com/goccy/spidermonkeywasm2go/p0.Fn1644
func Fn1644(m *base.Module, l0 int32, l1 float64, l2 int32) int32

//go:linkname Fn1652 github.com/goccy/spidermonkeywasm2go/p0.Fn1652
func Fn1652(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1653 github.com/goccy/spidermonkeywasm2go/p6.Fn1653
func Fn1653(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn1656 github.com/goccy/spidermonkeywasm2go/p6.Fn1656
func Fn1656(m *base.Module, l0 int32, l1 int32, l2 int32) float64

//go:linkname Fn1660 github.com/goccy/spidermonkeywasm2go/p0.Fn1660
func Fn1660(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1663 github.com/goccy/spidermonkeywasm2go/p0.Fn1663
func Fn1663(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1664 github.com/goccy/spidermonkeywasm2go/p0.Fn1664
func Fn1664(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1672 github.com/goccy/spidermonkeywasm2go/p4.Fn1672
func Fn1672(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1715 github.com/goccy/spidermonkeywasm2go/p7.Fn1715
func Fn1715(m *base.Module, l0 int32)

//go:linkname Fn1717 github.com/goccy/spidermonkeywasm2go/p6.Fn1717
func Fn1717(m *base.Module, l0 int32)

//go:linkname Fn1718 github.com/goccy/spidermonkeywasm2go/p7.Fn1718
func Fn1718(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1720 github.com/goccy/spidermonkeywasm2go/p4.Fn1720
func Fn1720(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1721 github.com/goccy/spidermonkeywasm2go/p7.Fn1721
func Fn1721(m *base.Module, l0 int32)

//go:linkname Fn1722 github.com/goccy/spidermonkeywasm2go/p6.Fn1722
func Fn1722(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1724 github.com/goccy/spidermonkeywasm2go/p6.Fn1724
func Fn1724(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1744 github.com/goccy/spidermonkeywasm2go/p0.Fn1744
func Fn1744(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1750 github.com/goccy/spidermonkeywasm2go/p0.Fn1750
func Fn1750(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1753 github.com/goccy/spidermonkeywasm2go/p0.Fn1753
func Fn1753(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1759 github.com/goccy/spidermonkeywasm2go/p6.Fn1759
func Fn1759(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1762 github.com/goccy/spidermonkeywasm2go/p7.Fn1762
func Fn1762(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1771 github.com/goccy/spidermonkeywasm2go/p6.Fn1771
func Fn1771(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1786 github.com/goccy/spidermonkeywasm2go/p0.Fn1786
func Fn1786(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1790 github.com/goccy/spidermonkeywasm2go/p6.Fn1790
func Fn1790(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1805 github.com/goccy/spidermonkeywasm2go/p7.Fn1805
func Fn1805(m *base.Module, l0 int32) int32

//go:linkname Fn1806 github.com/goccy/spidermonkeywasm2go/p6.Fn1806
func Fn1806(m *base.Module, l0 int32) int32

//go:linkname Fn1807 github.com/goccy/spidermonkeywasm2go/p6.Fn1807
func Fn1807(m *base.Module, l0 int32) int32

//go:linkname Fn1827 github.com/goccy/spidermonkeywasm2go/p0.Fn1827
func Fn1827(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1846 github.com/goccy/spidermonkeywasm2go/p0.Fn1846
func Fn1846(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1847 github.com/goccy/spidermonkeywasm2go/p7.Fn1847
func Fn1847(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1849 github.com/goccy/spidermonkeywasm2go/p0.Fn1849
func Fn1849(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1855 github.com/goccy/spidermonkeywasm2go/p6.Fn1855
func Fn1855(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1867 github.com/goccy/spidermonkeywasm2go/p6.Fn1867
func Fn1867(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1868 github.com/goccy/spidermonkeywasm2go/p6.Fn1868
func Fn1868(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1870 github.com/goccy/spidermonkeywasm2go/p3.Fn1870
func Fn1870(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1878 github.com/goccy/spidermonkeywasm2go/p4.Fn1878
func Fn1878(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1881 github.com/goccy/spidermonkeywasm2go/p4.Fn1881
func Fn1881(m *base.Module, l0 int32)

//go:linkname Fn1883 github.com/goccy/spidermonkeywasm2go/p6.Fn1883
func Fn1883(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1884 github.com/goccy/spidermonkeywasm2go/p7.Fn1884
func Fn1884(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1885 github.com/goccy/spidermonkeywasm2go/p4.Fn1885
func Fn1885(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1887 github.com/goccy/spidermonkeywasm2go/p7.Fn1887
func Fn1887(m *base.Module, l0 int32) int32

//go:linkname Fn1890 github.com/goccy/spidermonkeywasm2go/p7.Fn1890
func Fn1890(m *base.Module, l0 int32) int32

//go:linkname Fn1904 github.com/goccy/spidermonkeywasm2go/p6.Fn1904
func Fn1904(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1906 github.com/goccy/spidermonkeywasm2go/p6.Fn1906
func Fn1906(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1911 github.com/goccy/spidermonkeywasm2go/p0.Fn1911
func Fn1911(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1933 github.com/goccy/spidermonkeywasm2go/p0.Fn1933
func Fn1933(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1938 github.com/goccy/spidermonkeywasm2go/p7.Fn1938
func Fn1938(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1940 github.com/goccy/spidermonkeywasm2go/p0.Fn1940
func Fn1940(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1953 github.com/goccy/spidermonkeywasm2go/p0.Fn1953
func Fn1953(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn1954 github.com/goccy/spidermonkeywasm2go/p6.Fn1954
func Fn1954(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1957 github.com/goccy/spidermonkeywasm2go/p0.Fn1957
func Fn1957(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1958 github.com/goccy/spidermonkeywasm2go/p0.Fn1958
func Fn1958(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1959 github.com/goccy/spidermonkeywasm2go/p0.Fn1959
func Fn1959(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1960 github.com/goccy/spidermonkeywasm2go/p0.Fn1960
func Fn1960(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1973 github.com/goccy/spidermonkeywasm2go/p0.Fn1973
func Fn1973(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1974 github.com/goccy/spidermonkeywasm2go/p0.Fn1974
func Fn1974(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1975 github.com/goccy/spidermonkeywasm2go/p0.Fn1975
func Fn1975(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1976 github.com/goccy/spidermonkeywasm2go/p0.Fn1976
func Fn1976(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1977 github.com/goccy/spidermonkeywasm2go/p0.Fn1977
func Fn1977(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1980 github.com/goccy/spidermonkeywasm2go/p0.Fn1980
func Fn1980(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1981 github.com/goccy/spidermonkeywasm2go/p0.Fn1981
func Fn1981(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2023 github.com/goccy/spidermonkeywasm2go/p7.Fn2023
func Fn2023(m *base.Module, l0 int32) int32

//go:linkname Fn2031 github.com/goccy/spidermonkeywasm2go/p7.Fn2031
func Fn2031(m *base.Module, l0 int32) int32

//go:linkname Fn2034 github.com/goccy/spidermonkeywasm2go/p6.Fn2034
func Fn2034(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2038 github.com/goccy/spidermonkeywasm2go/p7.Fn2038
func Fn2038(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2039 github.com/goccy/spidermonkeywasm2go/p6.Fn2039
func Fn2039(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2042 github.com/goccy/spidermonkeywasm2go/p7.Fn2042
func Fn2042(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2053 github.com/goccy/spidermonkeywasm2go/p0.Fn2053
func Fn2053(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2058 github.com/goccy/spidermonkeywasm2go/p0.Fn2058
func Fn2058(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2060 github.com/goccy/spidermonkeywasm2go/p6.Fn2060
func Fn2060(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2061 github.com/goccy/spidermonkeywasm2go/p6.Fn2061
func Fn2061(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2071 github.com/goccy/spidermonkeywasm2go/p7.Fn2071
func Fn2071(m *base.Module, l0 int32) int32

//go:linkname Fn2081 github.com/goccy/spidermonkeywasm2go/p2.Fn2081
func Fn2081(m *base.Module, l0 int32) int32

//go:linkname Fn2082 github.com/goccy/spidermonkeywasm2go/p7.Fn2082
func Fn2082(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2083 github.com/goccy/spidermonkeywasm2go/p6.Fn2083
func Fn2083(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2088 github.com/goccy/spidermonkeywasm2go/p6.Fn2088
func Fn2088(m *base.Module, l0 int32) int32

//go:linkname Fn2091 github.com/goccy/spidermonkeywasm2go/p4.Fn2091
func Fn2091(m *base.Module, l0 int32)

//go:linkname Fn2094 github.com/goccy/spidermonkeywasm2go/p6.Fn2094
func Fn2094(m *base.Module, l0 int32) int32

//go:linkname Fn2095 github.com/goccy/spidermonkeywasm2go/p6.Fn2095
func Fn2095(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn2098 github.com/goccy/spidermonkeywasm2go/p7.Fn2098
func Fn2098(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn2102 github.com/goccy/spidermonkeywasm2go/p7.Fn2102
func Fn2102(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2120 github.com/goccy/spidermonkeywasm2go/p4.Fn2120
func Fn2120(m *base.Module, l0 int32)

//go:linkname Fn2123 github.com/goccy/spidermonkeywasm2go/p6.Fn2123
func Fn2123(m *base.Module, l0 int32) int32

//go:linkname Fn2124 github.com/goccy/spidermonkeywasm2go/p6.Fn2124
func Fn2124(m *base.Module, l0 int32) int32

//go:linkname Fn2125 github.com/goccy/spidermonkeywasm2go/p0.Fn2125
func Fn2125(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn2142 github.com/goccy/spidermonkeywasm2go/p0.Fn2142
func Fn2142(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2183 github.com/goccy/spidermonkeywasm2go/p7.Fn2183
func Fn2183(m *base.Module, l0 int32)

//go:linkname Fn2184 github.com/goccy/spidermonkeywasm2go/p4.Fn2184
func Fn2184(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2185 github.com/goccy/spidermonkeywasm2go/p0.Fn2185
func Fn2185(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2187 github.com/goccy/spidermonkeywasm2go/p6.Fn2187
func Fn2187(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2188 github.com/goccy/spidermonkeywasm2go/p7.Fn2188
func Fn2188(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2190 github.com/goccy/spidermonkeywasm2go/p7.Fn2190
func Fn2190(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2196 github.com/goccy/spidermonkeywasm2go/p4.Fn2196
func Fn2196(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2197 github.com/goccy/spidermonkeywasm2go/p0.Fn2197
func Fn2197(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32

//go:linkname Fn2207 github.com/goccy/spidermonkeywasm2go/p6.Fn2207
func Fn2207(m *base.Module, l0 int32)

//go:linkname Fn2211 github.com/goccy/spidermonkeywasm2go/p6.Fn2211
func Fn2211(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2213 github.com/goccy/spidermonkeywasm2go/p4.Fn2213
func Fn2213(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2221 github.com/goccy/spidermonkeywasm2go/p6.Fn2221
func Fn2221(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2222 github.com/goccy/spidermonkeywasm2go/p4.Fn2222
func Fn2222(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2224 github.com/goccy/spidermonkeywasm2go/p6.Fn2224
func Fn2224(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2235 github.com/goccy/spidermonkeywasm2go/p6.Fn2235
func Fn2235(m *base.Module, l0 int32)

//go:linkname Fn2245 github.com/goccy/spidermonkeywasm2go/p0.Fn2245
func Fn2245(m *base.Module, l0 int32) int32

//go:linkname Fn2250 github.com/goccy/spidermonkeywasm2go/p0.Fn2250
func Fn2250(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2252 github.com/goccy/spidermonkeywasm2go/p0.Fn2252
func Fn2252(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2276 github.com/goccy/spidermonkeywasm2go/p6.Fn2276
func Fn2276(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2287 github.com/goccy/spidermonkeywasm2go/p3.Fn2287
func Fn2287(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2298 github.com/goccy/spidermonkeywasm2go/p0.Fn2298
func Fn2298(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2300 github.com/goccy/spidermonkeywasm2go/p0.Fn2300
func Fn2300(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn2303 github.com/goccy/spidermonkeywasm2go/p6.Fn2303
func Fn2303(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2304 github.com/goccy/spidermonkeywasm2go/p7.Fn2304
func Fn2304(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn2318 github.com/goccy/spidermonkeywasm2go/p0.Fn2318
func Fn2318(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2319 github.com/goccy/spidermonkeywasm2go/p6.Fn2319
func Fn2319(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2320 github.com/goccy/spidermonkeywasm2go/p0.Fn2320
func Fn2320(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2321 github.com/goccy/spidermonkeywasm2go/p0.Fn2321
func Fn2321(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2324 github.com/goccy/spidermonkeywasm2go/p0.Fn2324
func Fn2324(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2331 github.com/goccy/spidermonkeywasm2go/p7.Fn2331
func Fn2331(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2334 github.com/goccy/spidermonkeywasm2go/p0.Fn2334
func Fn2334(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2338 github.com/goccy/spidermonkeywasm2go/p6.Fn2338
func Fn2338(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2339 github.com/goccy/spidermonkeywasm2go/p6.Fn2339
func Fn2339(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2340 github.com/goccy/spidermonkeywasm2go/p6.Fn2340
func Fn2340(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2399 github.com/goccy/spidermonkeywasm2go/p4.Fn2399
func Fn2399(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2430 github.com/goccy/spidermonkeywasm2go/p6.Fn2430
func Fn2430(m *base.Module, l0 int32)

//go:linkname Fn2437 github.com/goccy/spidermonkeywasm2go/p7.Fn2437
func Fn2437(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2447 github.com/goccy/spidermonkeywasm2go/p0.Fn2447
func Fn2447(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2455 github.com/goccy/spidermonkeywasm2go/p6.Fn2455
func Fn2455(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2460 github.com/goccy/spidermonkeywasm2go/p0.Fn2460
func Fn2460(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2478 github.com/goccy/spidermonkeywasm2go/p4.Fn2478
func Fn2478(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2483 github.com/goccy/spidermonkeywasm2go/p6.Fn2483
func Fn2483(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2484 github.com/goccy/spidermonkeywasm2go/p4.Fn2484
func Fn2484(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2495 github.com/goccy/spidermonkeywasm2go/p6.Fn2495
func Fn2495(m *base.Module, l0 int32) int32

//go:linkname Fn2496 github.com/goccy/spidermonkeywasm2go/p6.Fn2496
func Fn2496(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2498 github.com/goccy/spidermonkeywasm2go/p6.Fn2498
func Fn2498(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn2505 github.com/goccy/spidermonkeywasm2go/p6.Fn2505
func Fn2505(m *base.Module, l0 int32) int32

//go:linkname Fn2527 github.com/goccy/spidermonkeywasm2go/p6.Fn2527
func Fn2527(m *base.Module, l0 int32)

//go:linkname Fn2531 github.com/goccy/spidermonkeywasm2go/p7.Fn2531
func Fn2531(m *base.Module, l0 int32)

//go:linkname Fn2536 github.com/goccy/spidermonkeywasm2go/p7.Fn2536
func Fn2536(m *base.Module, l0 int32)

//go:linkname Fn2551 github.com/goccy/spidermonkeywasm2go/p0.Fn2551
func Fn2551(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2552 github.com/goccy/spidermonkeywasm2go/p0.Fn2552
func Fn2552(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2571 github.com/goccy/spidermonkeywasm2go/p0.Fn2571
func Fn2571(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn2573 github.com/goccy/spidermonkeywasm2go/p0.Fn2573
func Fn2573(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2575 github.com/goccy/spidermonkeywasm2go/p3.Fn2575
func Fn2575(m *base.Module, l0 int32)

//go:linkname Fn2577 github.com/goccy/spidermonkeywasm2go/p0.Fn2577
func Fn2577(m *base.Module, l0 int32)

//go:linkname Fn2578 github.com/goccy/spidermonkeywasm2go/p0.Fn2578
func Fn2578(m *base.Module, l0 int32)

//go:linkname Fn2582 github.com/goccy/spidermonkeywasm2go/p0.Fn2582
func Fn2582(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn2584 github.com/goccy/spidermonkeywasm2go/p7.Fn2584
func Fn2584(m *base.Module, l0 int32)

//go:linkname Fn2589 github.com/goccy/spidermonkeywasm2go/p0.Fn2589
func Fn2589(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2593 github.com/goccy/spidermonkeywasm2go/p7.Fn2593
func Fn2593(m *base.Module, l0 int32)

//go:linkname Fn2606 github.com/goccy/spidermonkeywasm2go/p6.Fn2606
func Fn2606(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2607 github.com/goccy/spidermonkeywasm2go/p6.Fn2607
func Fn2607(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2609 github.com/goccy/spidermonkeywasm2go/p0.Fn2609
func Fn2609(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2611 github.com/goccy/spidermonkeywasm2go/p0.Fn2611
func Fn2611(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2617 github.com/goccy/spidermonkeywasm2go/p6.Fn2617
func Fn2617(m *base.Module, l0 int32) int32

//go:linkname Fn2618 github.com/goccy/spidermonkeywasm2go/p6.Fn2618
func Fn2618(m *base.Module, l0 int32) int32

//go:linkname Fn2620 github.com/goccy/spidermonkeywasm2go/p0.Fn2620
func Fn2620(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn2623 github.com/goccy/spidermonkeywasm2go/p0.Fn2623
func Fn2623(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2627 github.com/goccy/spidermonkeywasm2go/p6.Fn2627
func Fn2627(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2628 github.com/goccy/spidermonkeywasm2go/p6.Fn2628
func Fn2628(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2653 github.com/goccy/spidermonkeywasm2go/p6.Fn2653
func Fn2653(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2655 github.com/goccy/spidermonkeywasm2go/p7.Fn2655
func Fn2655(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2691 github.com/goccy/spidermonkeywasm2go/p6.Fn2691
func Fn2691(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2706 github.com/goccy/spidermonkeywasm2go/p6.Fn2706
func Fn2706(m *base.Module, l0 int32)

//go:linkname Fn2707 github.com/goccy/spidermonkeywasm2go/p7.Fn2707
func Fn2707(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2709 github.com/goccy/spidermonkeywasm2go/p0.Fn2709
func Fn2709(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2710 github.com/goccy/spidermonkeywasm2go/p7.Fn2710
func Fn2710(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2711 github.com/goccy/spidermonkeywasm2go/p7.Fn2711
func Fn2711(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2716 github.com/goccy/spidermonkeywasm2go/p0.Fn2716
func Fn2716(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2720 github.com/goccy/spidermonkeywasm2go/p0.Fn2720
func Fn2720(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2721 github.com/goccy/spidermonkeywasm2go/p0.Fn2721
func Fn2721(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2722 github.com/goccy/spidermonkeywasm2go/p0.Fn2722
func Fn2722(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2724 github.com/goccy/spidermonkeywasm2go/p0.Fn2724
func Fn2724(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2731 github.com/goccy/spidermonkeywasm2go/p4.Fn2731
func Fn2731(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2733 github.com/goccy/spidermonkeywasm2go/p6.Fn2733
func Fn2733(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2735 github.com/goccy/spidermonkeywasm2go/p0.Fn2735
func Fn2735(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2738 github.com/goccy/spidermonkeywasm2go/p0.Fn2738
func Fn2738(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2739 github.com/goccy/spidermonkeywasm2go/p0.Fn2739
func Fn2739(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2741 github.com/goccy/spidermonkeywasm2go/p0.Fn2741
func Fn2741(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2742 github.com/goccy/spidermonkeywasm2go/p0.Fn2742
func Fn2742(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2743 github.com/goccy/spidermonkeywasm2go/p0.Fn2743
func Fn2743(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2750 github.com/goccy/spidermonkeywasm2go/p6.Fn2750
func Fn2750(m *base.Module, l0 int32) int32

//go:linkname Fn2762 github.com/goccy/spidermonkeywasm2go/p7.Fn2762
func Fn2762(m *base.Module, l0 int32) int32

//go:linkname Fn2763 github.com/goccy/spidermonkeywasm2go/p6.Fn2763
func Fn2763(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2764 github.com/goccy/spidermonkeywasm2go/p7.Fn2764
func Fn2764(m *base.Module, l0 int32)

//go:linkname Fn2777 github.com/goccy/spidermonkeywasm2go/p6.Fn2777
func Fn2777(m *base.Module, l0 int32)

//go:linkname Fn2778 github.com/goccy/spidermonkeywasm2go/p7.Fn2778
func Fn2778(m *base.Module, l0 int32) int32

//go:linkname Fn2783 github.com/goccy/spidermonkeywasm2go/p6.Fn2783
func Fn2783(m *base.Module, l0 int32)

//go:linkname Fn2784 github.com/goccy/spidermonkeywasm2go/p7.Fn2784
func Fn2784(m *base.Module, l0 int32) int32

//go:linkname Fn2786 github.com/goccy/spidermonkeywasm2go/p7.Fn2786
func Fn2786(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2789 github.com/goccy/spidermonkeywasm2go/p6.Fn2789
func Fn2789(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2792 github.com/goccy/spidermonkeywasm2go/p6.Fn2792
func Fn2792(m *base.Module, l0 int32)

//go:linkname Fn2793 github.com/goccy/spidermonkeywasm2go/p7.Fn2793
func Fn2793(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2795 github.com/goccy/spidermonkeywasm2go/p6.Fn2795
func Fn2795(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2812 github.com/goccy/spidermonkeywasm2go/p6.Fn2812
func Fn2812(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2828 github.com/goccy/spidermonkeywasm2go/p2.Fn2828
func Fn2828(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2835 github.com/goccy/spidermonkeywasm2go/p6.Fn2835
func Fn2835(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2838 github.com/goccy/spidermonkeywasm2go/p4.Fn2838
func Fn2838(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2840 github.com/goccy/spidermonkeywasm2go/p2.Fn2840
func Fn2840(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2851 github.com/goccy/spidermonkeywasm2go/p0.Fn2851
func Fn2851(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2852 github.com/goccy/spidermonkeywasm2go/p6.Fn2852
func Fn2852(m *base.Module, l0 int32)

//go:linkname Fn2853 github.com/goccy/spidermonkeywasm2go/p6.Fn2853
func Fn2853(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2857 github.com/goccy/spidermonkeywasm2go/p6.Fn2857
func Fn2857(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2859 github.com/goccy/spidermonkeywasm2go/p6.Fn2859
func Fn2859(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2865 github.com/goccy/spidermonkeywasm2go/p6.Fn2865
func Fn2865(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2870 github.com/goccy/spidermonkeywasm2go/p0.Fn2870
func Fn2870(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2874 github.com/goccy/spidermonkeywasm2go/p0.Fn2874
func Fn2874(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2875 github.com/goccy/spidermonkeywasm2go/p4.Fn2875
func Fn2875(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2876 github.com/goccy/spidermonkeywasm2go/p0.Fn2876
func Fn2876(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2882 github.com/goccy/spidermonkeywasm2go/p0.Fn2882
func Fn2882(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2886 github.com/goccy/spidermonkeywasm2go/p0.Fn2886
func Fn2886(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2891 github.com/goccy/spidermonkeywasm2go/p6.Fn2891
func Fn2891(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2893 github.com/goccy/spidermonkeywasm2go/p4.Fn2893
func Fn2893(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2894 github.com/goccy/spidermonkeywasm2go/p6.Fn2894
func Fn2894(m *base.Module, l0 int32) int32

//go:linkname Fn2897 github.com/goccy/spidermonkeywasm2go/p6.Fn2897
func Fn2897(m *base.Module, l0 int32) int32

//go:linkname Fn2899 github.com/goccy/spidermonkeywasm2go/p6.Fn2899
func Fn2899(m *base.Module, l0 int32) int32

//go:linkname Fn2900 github.com/goccy/spidermonkeywasm2go/p6.Fn2900
func Fn2900(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2903 github.com/goccy/spidermonkeywasm2go/p6.Fn2903
func Fn2903(m *base.Module, l0 int32) int32

//go:linkname Fn2904 github.com/goccy/spidermonkeywasm2go/p6.Fn2904
func Fn2904(m *base.Module, l0 int32)

//go:linkname Fn2908 github.com/goccy/spidermonkeywasm2go/p7.Fn2908
func Fn2908(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2928 github.com/goccy/spidermonkeywasm2go/p0.Fn2928
func Fn2928(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2929 github.com/goccy/spidermonkeywasm2go/p0.Fn2929
func Fn2929(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2930 github.com/goccy/spidermonkeywasm2go/p0.Fn2930
func Fn2930(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2931 github.com/goccy/spidermonkeywasm2go/p0.Fn2931
func Fn2931(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2939 github.com/goccy/spidermonkeywasm2go/p7.Fn2939
func Fn2939(m *base.Module, l0 int32)

//go:linkname Fn2942 github.com/goccy/spidermonkeywasm2go/p7.Fn2942
func Fn2942(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2947 github.com/goccy/spidermonkeywasm2go/p0.Fn2947
func Fn2947(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2955 github.com/goccy/spidermonkeywasm2go/p6.Fn2955
func Fn2955(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2962 github.com/goccy/spidermonkeywasm2go/p0.Fn2962
func Fn2962(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn2963 github.com/goccy/spidermonkeywasm2go/p0.Fn2963
func Fn2963(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn2971 github.com/goccy/spidermonkeywasm2go/p7.Fn2971
func Fn2971(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2974 github.com/goccy/spidermonkeywasm2go/p6.Fn2974
func Fn2974(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2975 github.com/goccy/spidermonkeywasm2go/p0.Fn2975
func Fn2975(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2993 github.com/goccy/spidermonkeywasm2go/p6.Fn2993
func Fn2993(m *base.Module, l0 int32)

//go:linkname Fn2998 github.com/goccy/spidermonkeywasm2go/p6.Fn2998
func Fn2998(m *base.Module, l0 int32) int32

//go:linkname Fn3010 github.com/goccy/spidermonkeywasm2go/p6.Fn3010
func Fn3010(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3026 github.com/goccy/spidermonkeywasm2go/p6.Fn3026
func Fn3026(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3035 github.com/goccy/spidermonkeywasm2go/p0.Fn3035
func Fn3035(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3036 github.com/goccy/spidermonkeywasm2go/p7.Fn3036
func Fn3036(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3040 github.com/goccy/spidermonkeywasm2go/p6.Fn3040
func Fn3040(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3041 github.com/goccy/spidermonkeywasm2go/p0.Fn3041
func Fn3041(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3044 github.com/goccy/spidermonkeywasm2go/p6.Fn3044
func Fn3044(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3051 github.com/goccy/spidermonkeywasm2go/p6.Fn3051
func Fn3051(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3092 github.com/goccy/spidermonkeywasm2go/p0.Fn3092
func Fn3092(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3094 github.com/goccy/spidermonkeywasm2go/p6.Fn3094
func Fn3094(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3129 github.com/goccy/spidermonkeywasm2go/p0.Fn3129
func Fn3129(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3148 github.com/goccy/spidermonkeywasm2go/p6.Fn3148
func Fn3148(m *base.Module, l0 float64) float64

//go:linkname Fn3152 github.com/goccy/spidermonkeywasm2go/p7.Fn3152
func Fn3152(m *base.Module, l0 int32)

//go:linkname Fn3153 github.com/goccy/spidermonkeywasm2go/p0.Fn3153
func Fn3153(m *base.Module, l0 int32) int32

//go:linkname Fn3156 github.com/goccy/spidermonkeywasm2go/p7.Fn3156
func Fn3156(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3157 github.com/goccy/spidermonkeywasm2go/p7.Fn3157
func Fn3157(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3166 github.com/goccy/spidermonkeywasm2go/p0.Fn3166
func Fn3166(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3168 github.com/goccy/spidermonkeywasm2go/p0.Fn3168
func Fn3168(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3170 github.com/goccy/spidermonkeywasm2go/p6.Fn3170
func Fn3170(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3186 github.com/goccy/spidermonkeywasm2go/p0.Fn3186
func Fn3186(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3193 github.com/goccy/spidermonkeywasm2go/p6.Fn3193
func Fn3193(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3203 github.com/goccy/spidermonkeywasm2go/p0.Fn3203
func Fn3203(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3207 github.com/goccy/spidermonkeywasm2go/p0.Fn3207
func Fn3207(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3212 github.com/goccy/spidermonkeywasm2go/p6.Fn3212
func Fn3212(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3214 github.com/goccy/spidermonkeywasm2go/p6.Fn3214
func Fn3214(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3216 github.com/goccy/spidermonkeywasm2go/p6.Fn3216
func Fn3216(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3218 github.com/goccy/spidermonkeywasm2go/p6.Fn3218
func Fn3218(m *base.Module, l0 int32) int32

//go:linkname Fn3220 github.com/goccy/spidermonkeywasm2go/p6.Fn3220
func Fn3220(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3228 github.com/goccy/spidermonkeywasm2go/p0.Fn3228
func Fn3228(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3234 github.com/goccy/spidermonkeywasm2go/p6.Fn3234
func Fn3234(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3244 github.com/goccy/spidermonkeywasm2go/p0.Fn3244
func Fn3244(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3255 github.com/goccy/spidermonkeywasm2go/p0.Fn3255
func Fn3255(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3498 github.com/goccy/spidermonkeywasm2go/p7.Fn3498
func Fn3498(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3499 github.com/goccy/spidermonkeywasm2go/p7.Fn3499
func Fn3499(m *base.Module, l0 int32)

//go:linkname Fn3500 github.com/goccy/spidermonkeywasm2go/p4.Fn3500
func Fn3500(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3520 github.com/goccy/spidermonkeywasm2go/p7.Fn3520
func Fn3520(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3522 github.com/goccy/spidermonkeywasm2go/p2.Fn3522
func Fn3522(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3525 github.com/goccy/spidermonkeywasm2go/p6.Fn3525
func Fn3525(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3527 github.com/goccy/spidermonkeywasm2go/p7.Fn3527
func Fn3527(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3530 github.com/goccy/spidermonkeywasm2go/p6.Fn3530
func Fn3530(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3532 github.com/goccy/spidermonkeywasm2go/p7.Fn3532
func Fn3532(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3535 github.com/goccy/spidermonkeywasm2go/p7.Fn3535
func Fn3535(m *base.Module, l0 int32) int32

//go:linkname Fn3540 github.com/goccy/spidermonkeywasm2go/p0.Fn3540
func Fn3540(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3546 github.com/goccy/spidermonkeywasm2go/p0.Fn3546
func Fn3546(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3549 github.com/goccy/spidermonkeywasm2go/p0.Fn3549
func Fn3549(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3550 github.com/goccy/spidermonkeywasm2go/p0.Fn3550
func Fn3550(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3554 github.com/goccy/spidermonkeywasm2go/p0.Fn3554
func Fn3554(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3562 github.com/goccy/spidermonkeywasm2go/p6.Fn3562
func Fn3562(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3563 github.com/goccy/spidermonkeywasm2go/p0.Fn3563
func Fn3563(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3564 github.com/goccy/spidermonkeywasm2go/p0.Fn3564
func Fn3564(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3566 github.com/goccy/spidermonkeywasm2go/p7.Fn3566
func Fn3566(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3568 github.com/goccy/spidermonkeywasm2go/p6.Fn3568
func Fn3568(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3572 github.com/goccy/spidermonkeywasm2go/p6.Fn3572
func Fn3572(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3575 github.com/goccy/spidermonkeywasm2go/p6.Fn3575
func Fn3575(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3588 github.com/goccy/spidermonkeywasm2go/p6.Fn3588
func Fn3588(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3592 github.com/goccy/spidermonkeywasm2go/p4.Fn3592
func Fn3592(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3595 github.com/goccy/spidermonkeywasm2go/p7.Fn3595
func Fn3595(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3620 github.com/goccy/spidermonkeywasm2go/p6.Fn3620
func Fn3620(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3632 github.com/goccy/spidermonkeywasm2go/p0.Fn3632
func Fn3632(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3633 github.com/goccy/spidermonkeywasm2go/p0.Fn3633
func Fn3633(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3640 github.com/goccy/spidermonkeywasm2go/p6.Fn3640
func Fn3640(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3648 github.com/goccy/spidermonkeywasm2go/p0.Fn3648
func Fn3648(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3705 github.com/goccy/spidermonkeywasm2go/p3.Fn3705
func Fn3705(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3709 github.com/goccy/spidermonkeywasm2go/p0.Fn3709
func Fn3709(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3722 github.com/goccy/spidermonkeywasm2go/p6.Fn3722
func Fn3722(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3736 github.com/goccy/spidermonkeywasm2go/p4.Fn3736
func Fn3736(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3865 github.com/goccy/spidermonkeywasm2go/p0.Fn3865
func Fn3865(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3867 github.com/goccy/spidermonkeywasm2go/p0.Fn3867
func Fn3867(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3869 github.com/goccy/spidermonkeywasm2go/p0.Fn3869
func Fn3869(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3871 github.com/goccy/spidermonkeywasm2go/p0.Fn3871
func Fn3871(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3873 github.com/goccy/spidermonkeywasm2go/p0.Fn3873
func Fn3873(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3875 github.com/goccy/spidermonkeywasm2go/p0.Fn3875
func Fn3875(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3877 github.com/goccy/spidermonkeywasm2go/p0.Fn3877
func Fn3877(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3879 github.com/goccy/spidermonkeywasm2go/p0.Fn3879
func Fn3879(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3881 github.com/goccy/spidermonkeywasm2go/p0.Fn3881
func Fn3881(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3883 github.com/goccy/spidermonkeywasm2go/p0.Fn3883
func Fn3883(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3885 github.com/goccy/spidermonkeywasm2go/p0.Fn3885
func Fn3885(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3887 github.com/goccy/spidermonkeywasm2go/p0.Fn3887
func Fn3887(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3956 github.com/goccy/spidermonkeywasm2go/p4.Fn3956
func Fn3956(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3957 github.com/goccy/spidermonkeywasm2go/p4.Fn3957
func Fn3957(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3958 github.com/goccy/spidermonkeywasm2go/p4.Fn3958
func Fn3958(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3959 github.com/goccy/spidermonkeywasm2go/p4.Fn3959
func Fn3959(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3960 github.com/goccy/spidermonkeywasm2go/p4.Fn3960
func Fn3960(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3961 github.com/goccy/spidermonkeywasm2go/p4.Fn3961
func Fn3961(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3962 github.com/goccy/spidermonkeywasm2go/p4.Fn3962
func Fn3962(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3963 github.com/goccy/spidermonkeywasm2go/p0.Fn3963
func Fn3963(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn3964 github.com/goccy/spidermonkeywasm2go/p0.Fn3964
func Fn3964(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3965 github.com/goccy/spidermonkeywasm2go/p0.Fn3965
func Fn3965(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3968 github.com/goccy/spidermonkeywasm2go/p0.Fn3968
func Fn3968(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn3969 github.com/goccy/spidermonkeywasm2go/p0.Fn3969
func Fn3969(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3970 github.com/goccy/spidermonkeywasm2go/p0.Fn3970
func Fn3970(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3973 github.com/goccy/spidermonkeywasm2go/p0.Fn3973
func Fn3973(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn3974 github.com/goccy/spidermonkeywasm2go/p0.Fn3974
func Fn3974(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3975 github.com/goccy/spidermonkeywasm2go/p0.Fn3975
func Fn3975(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3978 github.com/goccy/spidermonkeywasm2go/p0.Fn3978
func Fn3978(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn3979 github.com/goccy/spidermonkeywasm2go/p0.Fn3979
func Fn3979(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3980 github.com/goccy/spidermonkeywasm2go/p0.Fn3980
func Fn3980(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3983 github.com/goccy/spidermonkeywasm2go/p0.Fn3983
func Fn3983(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn3984 github.com/goccy/spidermonkeywasm2go/p0.Fn3984
func Fn3984(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3985 github.com/goccy/spidermonkeywasm2go/p0.Fn3985
func Fn3985(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3988 github.com/goccy/spidermonkeywasm2go/p0.Fn3988
func Fn3988(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn3989 github.com/goccy/spidermonkeywasm2go/p0.Fn3989
func Fn3989(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3990 github.com/goccy/spidermonkeywasm2go/p0.Fn3990
func Fn3990(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3993 github.com/goccy/spidermonkeywasm2go/p0.Fn3993
func Fn3993(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn3994 github.com/goccy/spidermonkeywasm2go/p0.Fn3994
func Fn3994(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3995 github.com/goccy/spidermonkeywasm2go/p0.Fn3995
func Fn3995(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3998 github.com/goccy/spidermonkeywasm2go/p0.Fn3998
func Fn3998(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn3999 github.com/goccy/spidermonkeywasm2go/p0.Fn3999
func Fn3999(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4000 github.com/goccy/spidermonkeywasm2go/p0.Fn4000
func Fn4000(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4003 github.com/goccy/spidermonkeywasm2go/p0.Fn4003
func Fn4003(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn4004 github.com/goccy/spidermonkeywasm2go/p0.Fn4004
func Fn4004(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4005 github.com/goccy/spidermonkeywasm2go/p0.Fn4005
func Fn4005(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4008 github.com/goccy/spidermonkeywasm2go/p0.Fn4008
func Fn4008(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn4009 github.com/goccy/spidermonkeywasm2go/p0.Fn4009
func Fn4009(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4010 github.com/goccy/spidermonkeywasm2go/p0.Fn4010
func Fn4010(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4013 github.com/goccy/spidermonkeywasm2go/p0.Fn4013
func Fn4013(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn4014 github.com/goccy/spidermonkeywasm2go/p0.Fn4014
func Fn4014(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4015 github.com/goccy/spidermonkeywasm2go/p0.Fn4015
func Fn4015(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4018 github.com/goccy/spidermonkeywasm2go/p0.Fn4018
func Fn4018(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn4019 github.com/goccy/spidermonkeywasm2go/p0.Fn4019
func Fn4019(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4020 github.com/goccy/spidermonkeywasm2go/p0.Fn4020
func Fn4020(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4088 github.com/goccy/spidermonkeywasm2go/p0.Fn4088
func Fn4088(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4089 github.com/goccy/spidermonkeywasm2go/p3.Fn4089
func Fn4089(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn4090 github.com/goccy/spidermonkeywasm2go/p6.Fn4090
func Fn4090(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4092 github.com/goccy/spidermonkeywasm2go/p3.Fn4092
func Fn4092(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4109 github.com/goccy/spidermonkeywasm2go/p6.Fn4109
func Fn4109(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4116 github.com/goccy/spidermonkeywasm2go/p6.Fn4116
func Fn4116(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4130 github.com/goccy/spidermonkeywasm2go/p0.Fn4130
func Fn4130(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4176 github.com/goccy/spidermonkeywasm2go/p6.Fn4176
func Fn4176(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4192 github.com/goccy/spidermonkeywasm2go/p6.Fn4192
func Fn4192(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4193 github.com/goccy/spidermonkeywasm2go/p7.Fn4193
func Fn4193(m *base.Module, l0 int32)

//go:linkname Fn4205 github.com/goccy/spidermonkeywasm2go/p6.Fn4205
func Fn4205(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4212 github.com/goccy/spidermonkeywasm2go/p6.Fn4212
func Fn4212(m *base.Module, l0 int32) int32

//go:linkname Fn4220 github.com/goccy/spidermonkeywasm2go/p6.Fn4220
func Fn4220(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4222 github.com/goccy/spidermonkeywasm2go/p6.Fn4222
func Fn4222(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4223 github.com/goccy/spidermonkeywasm2go/p6.Fn4223
func Fn4223(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4225 github.com/goccy/spidermonkeywasm2go/p6.Fn4225
func Fn4225(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4229 github.com/goccy/spidermonkeywasm2go/p2.Fn4229
func Fn4229(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4241 github.com/goccy/spidermonkeywasm2go/p3.Fn4241
func Fn4241(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4242 github.com/goccy/spidermonkeywasm2go/p4.Fn4242
func Fn4242(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32) int32

//go:linkname Fn4244 github.com/goccy/spidermonkeywasm2go/p4.Fn4244
func Fn4244(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4269 github.com/goccy/spidermonkeywasm2go/p6.Fn4269
func Fn4269(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4287 github.com/goccy/spidermonkeywasm2go/p6.Fn4287
func Fn4287(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4289 github.com/goccy/spidermonkeywasm2go/p7.Fn4289
func Fn4289(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4307 github.com/goccy/spidermonkeywasm2go/p4.Fn4307
func Fn4307(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4322 github.com/goccy/spidermonkeywasm2go/p6.Fn4322
func Fn4322(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4328 github.com/goccy/spidermonkeywasm2go/p4.Fn4328
func Fn4328(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4337 github.com/goccy/spidermonkeywasm2go/p2.Fn4337
func Fn4337(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4403 github.com/goccy/spidermonkeywasm2go/p3.Fn4403
func Fn4403(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4405 github.com/goccy/spidermonkeywasm2go/p6.Fn4405
func Fn4405(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4410 github.com/goccy/spidermonkeywasm2go/p6.Fn4410
func Fn4410(m *base.Module, l0 int32) int32

//go:linkname Fn4426 github.com/goccy/spidermonkeywasm2go/p4.Fn4426
func Fn4426(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4453 github.com/goccy/spidermonkeywasm2go/p6.Fn4453
func Fn4453(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4454 github.com/goccy/spidermonkeywasm2go/p6.Fn4454
func Fn4454(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4457 github.com/goccy/spidermonkeywasm2go/p6.Fn4457
func Fn4457(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4464 github.com/goccy/spidermonkeywasm2go/p6.Fn4464
func Fn4464(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4495 github.com/goccy/spidermonkeywasm2go/p2.Fn4495
func Fn4495(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4500 github.com/goccy/spidermonkeywasm2go/p4.Fn4500
func Fn4500(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4506 github.com/goccy/spidermonkeywasm2go/p4.Fn4506
func Fn4506(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4513 github.com/goccy/spidermonkeywasm2go/p6.Fn4513
func Fn4513(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4515 github.com/goccy/spidermonkeywasm2go/p6.Fn4515
func Fn4515(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4520 github.com/goccy/spidermonkeywasm2go/p2.Fn4520
func Fn4520(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4567 github.com/goccy/spidermonkeywasm2go/p4.Fn4567
func Fn4567(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4568 github.com/goccy/spidermonkeywasm2go/p7.Fn4568
func Fn4568(m *base.Module, l0 int32) int32

//go:linkname Fn4573 github.com/goccy/spidermonkeywasm2go/p4.Fn4573
func Fn4573(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn4581 github.com/goccy/spidermonkeywasm2go/p4.Fn4581
func Fn4581(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4584 github.com/goccy/spidermonkeywasm2go/p6.Fn4584
func Fn4584(m *base.Module, l0 int32) int32

//go:linkname Fn4606 github.com/goccy/spidermonkeywasm2go/p6.Fn4606
func Fn4606(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4613 github.com/goccy/spidermonkeywasm2go/p6.Fn4613
func Fn4613(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32

//go:linkname Fn4622 github.com/goccy/spidermonkeywasm2go/p7.Fn4622
func Fn4622(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4646 github.com/goccy/spidermonkeywasm2go/p6.Fn4646
func Fn4646(m *base.Module, l0 int32)

//go:linkname Fn4684 github.com/goccy/spidermonkeywasm2go/p4.Fn4684
func Fn4684(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4689 github.com/goccy/spidermonkeywasm2go/p4.Fn4689
func Fn4689(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4708 github.com/goccy/spidermonkeywasm2go/p7.Fn4708
func Fn4708(m *base.Module, l0 int32)

//go:linkname Fn4724 github.com/goccy/spidermonkeywasm2go/p6.Fn4724
func Fn4724(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4728 github.com/goccy/spidermonkeywasm2go/p6.Fn4728
func Fn4728(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4736 github.com/goccy/spidermonkeywasm2go/p4.Fn4736
func Fn4736(m *base.Module, l0 int32)

//go:linkname Fn4754 github.com/goccy/spidermonkeywasm2go/p7.Fn4754
func Fn4754(m *base.Module, l0 int32)

//go:linkname Fn4757 github.com/goccy/spidermonkeywasm2go/p6.Fn4757
func Fn4757(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4764 github.com/goccy/spidermonkeywasm2go/p4.Fn4764
func Fn4764(m *base.Module, l0 int32)

//go:linkname Fn4791 github.com/goccy/spidermonkeywasm2go/p6.Fn4791
func Fn4791(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4792 github.com/goccy/spidermonkeywasm2go/p6.Fn4792
func Fn4792(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4793 github.com/goccy/spidermonkeywasm2go/p6.Fn4793
func Fn4793(m *base.Module, l0 int32)

//go:linkname Fn4796 github.com/goccy/spidermonkeywasm2go/p7.Fn4796
func Fn4796(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4812 github.com/goccy/spidermonkeywasm2go/p7.Fn4812
func Fn4812(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4813 github.com/goccy/spidermonkeywasm2go/p6.Fn4813
func Fn4813(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4817 github.com/goccy/spidermonkeywasm2go/p7.Fn4817
func Fn4817(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4862 github.com/goccy/spidermonkeywasm2go/p7.Fn4862
func Fn4862(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4867 github.com/goccy/spidermonkeywasm2go/p7.Fn4867
func Fn4867(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4939 github.com/goccy/spidermonkeywasm2go/p6.Fn4939
func Fn4939(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5125 github.com/goccy/spidermonkeywasm2go/p7.Fn5125
func Fn5125(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5130 github.com/goccy/spidermonkeywasm2go/p6.Fn5130
func Fn5130(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5136 github.com/goccy/spidermonkeywasm2go/p2.Fn5136
func Fn5136(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn5184 github.com/goccy/spidermonkeywasm2go/p6.Fn5184
func Fn5184(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5266 github.com/goccy/spidermonkeywasm2go/p6.Fn5266
func Fn5266(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5267 github.com/goccy/spidermonkeywasm2go/p6.Fn5267
func Fn5267(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5276 github.com/goccy/spidermonkeywasm2go/p0.Fn5276
func Fn5276(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5322 github.com/goccy/spidermonkeywasm2go/p6.Fn5322
func Fn5322(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5340 github.com/goccy/spidermonkeywasm2go/p6.Fn5340
func Fn5340(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5343 github.com/goccy/spidermonkeywasm2go/p6.Fn5343
func Fn5343(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5348 github.com/goccy/spidermonkeywasm2go/p6.Fn5348
func Fn5348(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5351 github.com/goccy/spidermonkeywasm2go/p6.Fn5351
func Fn5351(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5352 github.com/goccy/spidermonkeywasm2go/p4.Fn5352
func Fn5352(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5367 github.com/goccy/spidermonkeywasm2go/p7.Fn5367
func Fn5367(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5401 github.com/goccy/spidermonkeywasm2go/p6.Fn5401
func Fn5401(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5409 github.com/goccy/spidermonkeywasm2go/p4.Fn5409
func Fn5409(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5436 github.com/goccy/spidermonkeywasm2go/p6.Fn5436
func Fn5436(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5448 github.com/goccy/spidermonkeywasm2go/p6.Fn5448
func Fn5448(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5536 github.com/goccy/spidermonkeywasm2go/p7.Fn5536
func Fn5536(m *base.Module, l0 int32) int32

//go:linkname Fn5537 github.com/goccy/spidermonkeywasm2go/p7.Fn5537
func Fn5537(m *base.Module, l0 int32) int32

//go:linkname Fn5554 github.com/goccy/spidermonkeywasm2go/p6.Fn5554
func Fn5554(m *base.Module, l0 int32)

//go:linkname Fn5555 github.com/goccy/spidermonkeywasm2go/p6.Fn5555
func Fn5555(m *base.Module, l0 int32)

//go:linkname Fn5556 github.com/goccy/spidermonkeywasm2go/p6.Fn5556
func Fn5556(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5557 github.com/goccy/spidermonkeywasm2go/p7.Fn5557
func Fn5557(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5560 github.com/goccy/spidermonkeywasm2go/p7.Fn5560
func Fn5560(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5562 github.com/goccy/spidermonkeywasm2go/p4.Fn5562
func Fn5562(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5563 github.com/goccy/spidermonkeywasm2go/p0.Fn5563
func Fn5563(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5564 github.com/goccy/spidermonkeywasm2go/p0.Fn5564
func Fn5564(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5565 github.com/goccy/spidermonkeywasm2go/p6.Fn5565
func Fn5565(m *base.Module, l0 int32)

//go:linkname Fn5571 github.com/goccy/spidermonkeywasm2go/p7.Fn5571
func Fn5571(m *base.Module, l0 int32)

//go:linkname Fn5572 github.com/goccy/spidermonkeywasm2go/p4.Fn5572
func Fn5572(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5574 github.com/goccy/spidermonkeywasm2go/p7.Fn5574
func Fn5574(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5576 github.com/goccy/spidermonkeywasm2go/p4.Fn5576
func Fn5576(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5580 github.com/goccy/spidermonkeywasm2go/p6.Fn5580
func Fn5580(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5582 github.com/goccy/spidermonkeywasm2go/p6.Fn5582
func Fn5582(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5583 github.com/goccy/spidermonkeywasm2go/p6.Fn5583
func Fn5583(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5584 github.com/goccy/spidermonkeywasm2go/p6.Fn5584
func Fn5584(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5592 github.com/goccy/spidermonkeywasm2go/p6.Fn5592
func Fn5592(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5595 github.com/goccy/spidermonkeywasm2go/p6.Fn5595
func Fn5595(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5596 github.com/goccy/spidermonkeywasm2go/p3.Fn5596
func Fn5596(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5599 github.com/goccy/spidermonkeywasm2go/p6.Fn5599
func Fn5599(m *base.Module, l0 int32)

//go:linkname Fn5612 github.com/goccy/spidermonkeywasm2go/p6.Fn5612
func Fn5612(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5613 github.com/goccy/spidermonkeywasm2go/p6.Fn5613
func Fn5613(m *base.Module, l0 int32)

//go:linkname Fn5617 github.com/goccy/spidermonkeywasm2go/p6.Fn5617
func Fn5617(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn5618 github.com/goccy/spidermonkeywasm2go/p6.Fn5618
func Fn5618(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn5620 github.com/goccy/spidermonkeywasm2go/p6.Fn5620
func Fn5620(m *base.Module, l0 int32)

//go:linkname Fn5626 github.com/goccy/spidermonkeywasm2go/p6.Fn5626
func Fn5626(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5627 github.com/goccy/spidermonkeywasm2go/p4.Fn5627
func Fn5627(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5632 github.com/goccy/spidermonkeywasm2go/p6.Fn5632
func Fn5632(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5633 github.com/goccy/spidermonkeywasm2go/p6.Fn5633
func Fn5633(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5639 github.com/goccy/spidermonkeywasm2go/p6.Fn5639
func Fn5639(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5650 github.com/goccy/spidermonkeywasm2go/p6.Fn5650
func Fn5650(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5651 github.com/goccy/spidermonkeywasm2go/p6.Fn5651
func Fn5651(m *base.Module, l0 int32)

//go:linkname Fn5652 github.com/goccy/spidermonkeywasm2go/p6.Fn5652
func Fn5652(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5653 github.com/goccy/spidermonkeywasm2go/p6.Fn5653
func Fn5653(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5658 github.com/goccy/spidermonkeywasm2go/p6.Fn5658
func Fn5658(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5660 github.com/goccy/spidermonkeywasm2go/p7.Fn5660
func Fn5660(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5663 github.com/goccy/spidermonkeywasm2go/p6.Fn5663
func Fn5663(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5666 github.com/goccy/spidermonkeywasm2go/p6.Fn5666
func Fn5666(m *base.Module, l0 int32)

//go:linkname Fn5691 github.com/goccy/spidermonkeywasm2go/p7.Fn5691
func Fn5691(m *base.Module, l0 int32)

//go:linkname Fn5693 github.com/goccy/spidermonkeywasm2go/p6.Fn5693
func Fn5693(m *base.Module, l0 int32) int32

//go:linkname Fn5700 github.com/goccy/spidermonkeywasm2go/p4.Fn5700
func Fn5700(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5703 github.com/goccy/spidermonkeywasm2go/p6.Fn5703
func Fn5703(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5710 github.com/goccy/spidermonkeywasm2go/p7.Fn5710
func Fn5710(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5712 github.com/goccy/spidermonkeywasm2go/p6.Fn5712
func Fn5712(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5713 github.com/goccy/spidermonkeywasm2go/p7.Fn5713
func Fn5713(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5714 github.com/goccy/spidermonkeywasm2go/p6.Fn5714
func Fn5714(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5715 github.com/goccy/spidermonkeywasm2go/p6.Fn5715
func Fn5715(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5716 github.com/goccy/spidermonkeywasm2go/p6.Fn5716
func Fn5716(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5717 github.com/goccy/spidermonkeywasm2go/p6.Fn5717
func Fn5717(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5719 github.com/goccy/spidermonkeywasm2go/p6.Fn5719
func Fn5719(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5720 github.com/goccy/spidermonkeywasm2go/p7.Fn5720
func Fn5720(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5721 github.com/goccy/spidermonkeywasm2go/p6.Fn5721
func Fn5721(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5722 github.com/goccy/spidermonkeywasm2go/p6.Fn5722
func Fn5722(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5726 github.com/goccy/spidermonkeywasm2go/p7.Fn5726
func Fn5726(m *base.Module, l0 int32)

//go:linkname Fn5728 github.com/goccy/spidermonkeywasm2go/p6.Fn5728
func Fn5728(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5729 github.com/goccy/spidermonkeywasm2go/p6.Fn5729
func Fn5729(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5730 github.com/goccy/spidermonkeywasm2go/p6.Fn5730
func Fn5730(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5731 github.com/goccy/spidermonkeywasm2go/p6.Fn5731
func Fn5731(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5733 github.com/goccy/spidermonkeywasm2go/p6.Fn5733
func Fn5733(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5734 github.com/goccy/spidermonkeywasm2go/p6.Fn5734
func Fn5734(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5736 github.com/goccy/spidermonkeywasm2go/p3.Fn5736
func Fn5736(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5739 github.com/goccy/spidermonkeywasm2go/p6.Fn5739
func Fn5739(m *base.Module, l0 int32)

//go:linkname Fn5747 github.com/goccy/spidermonkeywasm2go/p7.Fn5747
func Fn5747(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5749 github.com/goccy/spidermonkeywasm2go/p6.Fn5749
func Fn5749(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5751 github.com/goccy/spidermonkeywasm2go/p6.Fn5751
func Fn5751(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5764 github.com/goccy/spidermonkeywasm2go/p6.Fn5764
func Fn5764(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5766 github.com/goccy/spidermonkeywasm2go/p3.Fn5766
func Fn5766(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5768 github.com/goccy/spidermonkeywasm2go/p6.Fn5768
func Fn5768(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5807 github.com/goccy/spidermonkeywasm2go/p4.Fn5807
func Fn5807(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5808 github.com/goccy/spidermonkeywasm2go/p3.Fn5808
func Fn5808(m *base.Module, l0 int32)

//go:linkname Fn5809 github.com/goccy/spidermonkeywasm2go/p6.Fn5809
func Fn5809(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5810 github.com/goccy/spidermonkeywasm2go/p6.Fn5810
func Fn5810(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5811 github.com/goccy/spidermonkeywasm2go/p6.Fn5811
func Fn5811(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5812 github.com/goccy/spidermonkeywasm2go/p6.Fn5812
func Fn5812(m *base.Module, l0 int32)

//go:linkname Fn5820 github.com/goccy/spidermonkeywasm2go/p7.Fn5820
func Fn5820(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5821 github.com/goccy/spidermonkeywasm2go/p7.Fn5821
func Fn5821(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5830 github.com/goccy/spidermonkeywasm2go/p7.Fn5830
func Fn5830(m *base.Module, l0 int32) int32

//go:linkname Fn5832 github.com/goccy/spidermonkeywasm2go/p7.Fn5832
func Fn5832(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5834 github.com/goccy/spidermonkeywasm2go/p6.Fn5834
func Fn5834(m *base.Module, l0 int32) int32

//go:linkname Fn5864 github.com/goccy/spidermonkeywasm2go/p6.Fn5864
func Fn5864(m *base.Module, l0 int32)

//go:linkname Fn5900 github.com/goccy/spidermonkeywasm2go/p6.Fn5900
func Fn5900(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5903 github.com/goccy/spidermonkeywasm2go/p4.Fn5903
func Fn5903(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5929 github.com/goccy/spidermonkeywasm2go/p6.Fn5929
func Fn5929(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5930 github.com/goccy/spidermonkeywasm2go/p6.Fn5930
func Fn5930(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5931 github.com/goccy/spidermonkeywasm2go/p4.Fn5931
func Fn5931(m *base.Module, l0 int32)

//go:linkname Fn5932 github.com/goccy/spidermonkeywasm2go/p7.Fn5932
func Fn5932(m *base.Module, l0 int32)

//go:linkname Fn5941 github.com/goccy/spidermonkeywasm2go/p7.Fn5941
func Fn5941(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5949 github.com/goccy/spidermonkeywasm2go/p6.Fn5949
func Fn5949(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5959 github.com/goccy/spidermonkeywasm2go/p0.Fn5959
func Fn5959(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5961 github.com/goccy/spidermonkeywasm2go/p0.Fn5961
func Fn5961(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5970 github.com/goccy/spidermonkeywasm2go/p7.Fn5970
func Fn5970(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5978 github.com/goccy/spidermonkeywasm2go/p6.Fn5978
func Fn5978(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6033 github.com/goccy/spidermonkeywasm2go/p4.Fn6033
func Fn6033(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6040 github.com/goccy/spidermonkeywasm2go/p6.Fn6040
func Fn6040(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6048 github.com/goccy/spidermonkeywasm2go/p7.Fn6048
func Fn6048(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6054 github.com/goccy/spidermonkeywasm2go/p7.Fn6054
func Fn6054(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6055 github.com/goccy/spidermonkeywasm2go/p6.Fn6055
func Fn6055(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6074 github.com/goccy/spidermonkeywasm2go/p3.Fn6074
func Fn6074(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6079 github.com/goccy/spidermonkeywasm2go/p7.Fn6079
func Fn6079(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6080 github.com/goccy/spidermonkeywasm2go/p7.Fn6080
func Fn6080(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6095 github.com/goccy/spidermonkeywasm2go/p7.Fn6095
func Fn6095(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6096 github.com/goccy/spidermonkeywasm2go/p6.Fn6096
func Fn6096(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6098 github.com/goccy/spidermonkeywasm2go/p6.Fn6098
func Fn6098(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6113 github.com/goccy/spidermonkeywasm2go/p7.Fn6113
func Fn6113(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6175 github.com/goccy/spidermonkeywasm2go/p6.Fn6175
func Fn6175(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6176 github.com/goccy/spidermonkeywasm2go/p3.Fn6176
func Fn6176(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6177 github.com/goccy/spidermonkeywasm2go/p6.Fn6177
func Fn6177(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6197 github.com/goccy/spidermonkeywasm2go/p6.Fn6197
func Fn6197(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn6209 github.com/goccy/spidermonkeywasm2go/p1.Fn6209
func Fn6209(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6211 github.com/goccy/spidermonkeywasm2go/p6.Fn6211
func Fn6211(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6218 github.com/goccy/spidermonkeywasm2go/p6.Fn6218
func Fn6218(m *base.Module, l0 int32) int32

//go:linkname Fn6222 github.com/goccy/spidermonkeywasm2go/p3.Fn6222
func Fn6222(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn6227 github.com/goccy/spidermonkeywasm2go/p3.Fn6227
func Fn6227(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6237 github.com/goccy/spidermonkeywasm2go/p6.Fn6237
func Fn6237(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6254 github.com/goccy/spidermonkeywasm2go/p6.Fn6254
func Fn6254(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6260 github.com/goccy/spidermonkeywasm2go/p6.Fn6260
func Fn6260(m *base.Module, l0 int32)

//go:linkname Fn6282 github.com/goccy/spidermonkeywasm2go/p3.Fn6282
func Fn6282(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6284 github.com/goccy/spidermonkeywasm2go/p2.Fn6284
func Fn6284(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn6334 github.com/goccy/spidermonkeywasm2go/p2.Fn6334
func Fn6334(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6336 github.com/goccy/spidermonkeywasm2go/p6.Fn6336
func Fn6336(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6346 github.com/goccy/spidermonkeywasm2go/p7.Fn6346
func Fn6346(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6654 github.com/goccy/spidermonkeywasm2go/p7.Fn6654
func Fn6654(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6655 github.com/goccy/spidermonkeywasm2go/p6.Fn6655
func Fn6655(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6656 github.com/goccy/spidermonkeywasm2go/p6.Fn6656
func Fn6656(m *base.Module, l0 int32) int32

//go:linkname Fn6669 github.com/goccy/spidermonkeywasm2go/p6.Fn6669
func Fn6669(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6678 github.com/goccy/spidermonkeywasm2go/p7.Fn6678
func Fn6678(m *base.Module, l0 int32)

//go:linkname Fn6679 github.com/goccy/spidermonkeywasm2go/p6.Fn6679
func Fn6679(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn6680 github.com/goccy/spidermonkeywasm2go/p1.Fn6680
func Fn6680(m *base.Module, l0 int32) int32

//go:linkname Fn6684 github.com/goccy/spidermonkeywasm2go/p7.Fn6684
func Fn6684(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6689 github.com/goccy/spidermonkeywasm2go/p6.Fn6689
func Fn6689(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6699 github.com/goccy/spidermonkeywasm2go/p7.Fn6699
func Fn6699(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6701 github.com/goccy/spidermonkeywasm2go/p7.Fn6701
func Fn6701(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6705 github.com/goccy/spidermonkeywasm2go/p7.Fn6705
func Fn6705(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6710 github.com/goccy/spidermonkeywasm2go/p7.Fn6710
func Fn6710(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6711 github.com/goccy/spidermonkeywasm2go/p7.Fn6711
func Fn6711(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6712 github.com/goccy/spidermonkeywasm2go/p7.Fn6712
func Fn6712(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6713 github.com/goccy/spidermonkeywasm2go/p6.Fn6713
func Fn6713(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6714 github.com/goccy/spidermonkeywasm2go/p4.Fn6714
func Fn6714(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn6715 github.com/goccy/spidermonkeywasm2go/p7.Fn6715
func Fn6715(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6723 github.com/goccy/spidermonkeywasm2go/p7.Fn6723
func Fn6723(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6754 github.com/goccy/spidermonkeywasm2go/p4.Fn6754
func Fn6754(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6761 github.com/goccy/spidermonkeywasm2go/p7.Fn6761
func Fn6761(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6762 github.com/goccy/spidermonkeywasm2go/p7.Fn6762
func Fn6762(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6763 github.com/goccy/spidermonkeywasm2go/p7.Fn6763
func Fn6763(m *base.Module, l0 int32, l1 float64) int32

//go:linkname Fn6764 github.com/goccy/spidermonkeywasm2go/p7.Fn6764
func Fn6764(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6765 github.com/goccy/spidermonkeywasm2go/p7.Fn6765
func Fn6765(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6766 github.com/goccy/spidermonkeywasm2go/p7.Fn6766
func Fn6766(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6767 github.com/goccy/spidermonkeywasm2go/p7.Fn6767
func Fn6767(m *base.Module, l0 int32) int32

//go:linkname Fn6773 github.com/goccy/spidermonkeywasm2go/p4.Fn6773
func Fn6773(m *base.Module, l0 int32) int32

//go:linkname Fn6774 github.com/goccy/spidermonkeywasm2go/p6.Fn6774
func Fn6774(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6780 github.com/goccy/spidermonkeywasm2go/p7.Fn6780
func Fn6780(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6784 github.com/goccy/spidermonkeywasm2go/p6.Fn6784
func Fn6784(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6787 github.com/goccy/spidermonkeywasm2go/p7.Fn6787
func Fn6787(m *base.Module, l0 int32)

//go:linkname Fn6788 github.com/goccy/spidermonkeywasm2go/p7.Fn6788
func Fn6788(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6790 github.com/goccy/spidermonkeywasm2go/p6.Fn6790
func Fn6790(m *base.Module, l0 int32) int32

//go:linkname Fn6854 github.com/goccy/spidermonkeywasm2go/p6.Fn6854
func Fn6854(m *base.Module, l0 int32) int32

//go:linkname Fn6855 github.com/goccy/spidermonkeywasm2go/p6.Fn6855
func Fn6855(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6856 github.com/goccy/spidermonkeywasm2go/p7.Fn6856
func Fn6856(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6862 github.com/goccy/spidermonkeywasm2go/p4.Fn6862
func Fn6862(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6865 github.com/goccy/spidermonkeywasm2go/p4.Fn6865
func Fn6865(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6866 github.com/goccy/spidermonkeywasm2go/p6.Fn6866
func Fn6866(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6868 github.com/goccy/spidermonkeywasm2go/p6.Fn6868
func Fn6868(m *base.Module, l0 int32)

//go:linkname Fn6872 github.com/goccy/spidermonkeywasm2go/p6.Fn6872
func Fn6872(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6873 github.com/goccy/spidermonkeywasm2go/p6.Fn6873
func Fn6873(m *base.Module, l0 int32)

//go:linkname Fn6885 github.com/goccy/spidermonkeywasm2go/p3.Fn6885
func Fn6885(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6886 github.com/goccy/spidermonkeywasm2go/p6.Fn6886
func Fn6886(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6898 github.com/goccy/spidermonkeywasm2go/p6.Fn6898
func Fn6898(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn6899 github.com/goccy/spidermonkeywasm2go/p7.Fn6899
func Fn6899(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6906 github.com/goccy/spidermonkeywasm2go/p6.Fn6906
func Fn6906(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6918 github.com/goccy/spidermonkeywasm2go/p6.Fn6918
func Fn6918(m *base.Module, l0 int32)

//go:linkname Fn6920 github.com/goccy/spidermonkeywasm2go/p6.Fn6920
func Fn6920(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

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

//go:linkname Fn6944 github.com/goccy/spidermonkeywasm2go/p6.Fn6944
func Fn6944(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6945 github.com/goccy/spidermonkeywasm2go/p6.Fn6945
func Fn6945(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6946 github.com/goccy/spidermonkeywasm2go/p3.Fn6946
func Fn6946(m *base.Module, l0 int32) int32

//go:linkname Fn6947 github.com/goccy/spidermonkeywasm2go/p3.Fn6947
func Fn6947(m *base.Module, l0 int32) int32

//go:linkname Fn6948 github.com/goccy/spidermonkeywasm2go/p6.Fn6948
func Fn6948(m *base.Module, l0 int32) int32

//go:linkname Fn6949 github.com/goccy/spidermonkeywasm2go/p4.Fn6949
func Fn6949(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6958 github.com/goccy/spidermonkeywasm2go/p6.Fn6958
func Fn6958(m *base.Module, l0 int32) int32

//go:linkname Fn6966 github.com/goccy/spidermonkeywasm2go/p0.Fn6966
func Fn6966(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6967 github.com/goccy/spidermonkeywasm2go/p7.Fn6967
func Fn6967(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6969 github.com/goccy/spidermonkeywasm2go/p7.Fn6969
func Fn6969(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6981 github.com/goccy/spidermonkeywasm2go/p6.Fn6981
func Fn6981(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6982 github.com/goccy/spidermonkeywasm2go/p6.Fn6982
func Fn6982(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6985 github.com/goccy/spidermonkeywasm2go/p7.Fn6985
func Fn6985(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6992 github.com/goccy/spidermonkeywasm2go/p6.Fn6992
func Fn6992(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6999 github.com/goccy/spidermonkeywasm2go/p6.Fn6999
func Fn6999(m *base.Module, l0 int32) int32

//go:linkname Fn7002 github.com/goccy/spidermonkeywasm2go/p6.Fn7002
func Fn7002(m *base.Module, l0 int32) int32

//go:linkname Fn7011 github.com/goccy/spidermonkeywasm2go/p6.Fn7011
func Fn7011(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7020 github.com/goccy/spidermonkeywasm2go/p7.Fn7020
func Fn7020(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7023 github.com/goccy/spidermonkeywasm2go/p6.Fn7023
func Fn7023(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7051 github.com/goccy/spidermonkeywasm2go/p6.Fn7051
func Fn7051(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7068 github.com/goccy/spidermonkeywasm2go/p6.Fn7068
func Fn7068(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7106 github.com/goccy/spidermonkeywasm2go/p7.Fn7106
func Fn7106(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7108 github.com/goccy/spidermonkeywasm2go/p6.Fn7108
func Fn7108(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7111 github.com/goccy/spidermonkeywasm2go/p6.Fn7111
func Fn7111(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7121 github.com/goccy/spidermonkeywasm2go/p6.Fn7121
func Fn7121(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7123 github.com/goccy/spidermonkeywasm2go/p4.Fn7123
func Fn7123(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7132 github.com/goccy/spidermonkeywasm2go/p6.Fn7132
func Fn7132(m *base.Module, l0 int32) int32

//go:linkname Fn7135 github.com/goccy/spidermonkeywasm2go/p4.Fn7135
func Fn7135(m *base.Module, l0 int32)

//go:linkname Fn7136 github.com/goccy/spidermonkeywasm2go/p6.Fn7136
func Fn7136(m *base.Module, l0 int32) int32

//go:linkname Fn7138 github.com/goccy/spidermonkeywasm2go/p7.Fn7138
func Fn7138(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7150 github.com/goccy/spidermonkeywasm2go/p6.Fn7150
func Fn7150(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn7151 github.com/goccy/spidermonkeywasm2go/p3.Fn7151
func Fn7151(m *base.Module, l0 int32)

//go:linkname Fn7155 github.com/goccy/spidermonkeywasm2go/p6.Fn7155
func Fn7155(m *base.Module, l0 int32) int32

//go:linkname Fn7157 github.com/goccy/spidermonkeywasm2go/p7.Fn7157
func Fn7157(m *base.Module, l0 int32) int32

//go:linkname Fn7164 github.com/goccy/spidermonkeywasm2go/p6.Fn7164
func Fn7164(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7169 github.com/goccy/spidermonkeywasm2go/p6.Fn7169
func Fn7169(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7181 github.com/goccy/spidermonkeywasm2go/p4.Fn7181
func Fn7181(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7185 github.com/goccy/spidermonkeywasm2go/p6.Fn7185
func Fn7185(m *base.Module, l0 int32) int32

//go:linkname Fn7187 github.com/goccy/spidermonkeywasm2go/p6.Fn7187
func Fn7187(m *base.Module, l0 int32) int32

//go:linkname Fn7193 github.com/goccy/spidermonkeywasm2go/p2.Fn7193
func Fn7193(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7203 github.com/goccy/spidermonkeywasm2go/p6.Fn7203
func Fn7203(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7222 github.com/goccy/spidermonkeywasm2go/p6.Fn7222
func Fn7222(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7223 github.com/goccy/spidermonkeywasm2go/p4.Fn7223
func Fn7223(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7230 github.com/goccy/spidermonkeywasm2go/p6.Fn7230
func Fn7230(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7255 github.com/goccy/spidermonkeywasm2go/p6.Fn7255
func Fn7255(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7298 github.com/goccy/spidermonkeywasm2go/p4.Fn7298
func Fn7298(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7308 github.com/goccy/spidermonkeywasm2go/p7.Fn7308
func Fn7308(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7336 github.com/goccy/spidermonkeywasm2go/p6.Fn7336
func Fn7336(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7347 github.com/goccy/spidermonkeywasm2go/p7.Fn7347
func Fn7347(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7351 github.com/goccy/spidermonkeywasm2go/p1.Fn7351
func Fn7351(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7366 github.com/goccy/spidermonkeywasm2go/p6.Fn7366
func Fn7366(m *base.Module, l0 int32) int32

//go:linkname Fn7368 github.com/goccy/spidermonkeywasm2go/p4.Fn7368
func Fn7368(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7378 github.com/goccy/spidermonkeywasm2go/p6.Fn7378
func Fn7378(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7383 github.com/goccy/spidermonkeywasm2go/p4.Fn7383
func Fn7383(m *base.Module, l0 int32) int32

//go:linkname Fn7414 github.com/goccy/spidermonkeywasm2go/p1.Fn7414
func Fn7414(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7423 github.com/goccy/spidermonkeywasm2go/p2.Fn7423
func Fn7423(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7425 github.com/goccy/spidermonkeywasm2go/p4.Fn7425
func Fn7425(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7428 github.com/goccy/spidermonkeywasm2go/p3.Fn7428
func Fn7428(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7431 github.com/goccy/spidermonkeywasm2go/p6.Fn7431
func Fn7431(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7432 github.com/goccy/spidermonkeywasm2go/p6.Fn7432
func Fn7432(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7437 github.com/goccy/spidermonkeywasm2go/p4.Fn7437
func Fn7437(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7442 github.com/goccy/spidermonkeywasm2go/p7.Fn7442
func Fn7442(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7445 github.com/goccy/spidermonkeywasm2go/p6.Fn7445
func Fn7445(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7457 github.com/goccy/spidermonkeywasm2go/p6.Fn7457
func Fn7457(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7459 github.com/goccy/spidermonkeywasm2go/p6.Fn7459
func Fn7459(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7461 github.com/goccy/spidermonkeywasm2go/p4.Fn7461
func Fn7461(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7462 github.com/goccy/spidermonkeywasm2go/p6.Fn7462
func Fn7462(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7463 github.com/goccy/spidermonkeywasm2go/p6.Fn7463
func Fn7463(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn7468 github.com/goccy/spidermonkeywasm2go/p6.Fn7468
func Fn7468(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7469 github.com/goccy/spidermonkeywasm2go/p7.Fn7469
func Fn7469(m *base.Module, l0 int32, l1 int32, l2 float64)

//go:linkname Fn7470 github.com/goccy/spidermonkeywasm2go/p6.Fn7470
func Fn7470(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn7472 github.com/goccy/spidermonkeywasm2go/p6.Fn7472
func Fn7472(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7475 github.com/goccy/spidermonkeywasm2go/p6.Fn7475
func Fn7475(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn7476 github.com/goccy/spidermonkeywasm2go/p6.Fn7476
func Fn7476(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7479 github.com/goccy/spidermonkeywasm2go/p6.Fn7479
func Fn7479(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn7482 github.com/goccy/spidermonkeywasm2go/p3.Fn7482
func Fn7482(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn7484 github.com/goccy/spidermonkeywasm2go/p6.Fn7484
func Fn7484(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7485 github.com/goccy/spidermonkeywasm2go/p4.Fn7485
func Fn7485(m *base.Module, l0 int32, l1 float64, l2 float64, l3 float64, l4 float64, l5 float64, l6 float64)

//go:linkname Fn7492 github.com/goccy/spidermonkeywasm2go/p4.Fn7492
func Fn7492(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7493 github.com/goccy/spidermonkeywasm2go/p2.Fn7493
func Fn7493(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7494 github.com/goccy/spidermonkeywasm2go/p6.Fn7494
func Fn7494(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7497 github.com/goccy/spidermonkeywasm2go/p6.Fn7497
func Fn7497(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn7505 github.com/goccy/spidermonkeywasm2go/p6.Fn7505
func Fn7505(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn7508 github.com/goccy/spidermonkeywasm2go/p6.Fn7508
func Fn7508(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7512 github.com/goccy/spidermonkeywasm2go/p7.Fn7512
func Fn7512(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7578 github.com/goccy/spidermonkeywasm2go/p4.Fn7578
func Fn7578(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7686 github.com/goccy/spidermonkeywasm2go/p6.Fn7686
func Fn7686(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn7694 github.com/goccy/spidermonkeywasm2go/p3.Fn7694
func Fn7694(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7695 github.com/goccy/spidermonkeywasm2go/p6.Fn7695
func Fn7695(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7696 github.com/goccy/spidermonkeywasm2go/p6.Fn7696
func Fn7696(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7703 github.com/goccy/spidermonkeywasm2go/p6.Fn7703
func Fn7703(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn7705 github.com/goccy/spidermonkeywasm2go/p6.Fn7705
func Fn7705(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7706 github.com/goccy/spidermonkeywasm2go/p4.Fn7706
func Fn7706(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7707 github.com/goccy/spidermonkeywasm2go/p7.Fn7707
func Fn7707(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7708 github.com/goccy/spidermonkeywasm2go/p7.Fn7708
func Fn7708(m *base.Module, l0 int32) int32

//go:linkname Fn7713 github.com/goccy/spidermonkeywasm2go/p6.Fn7713
func Fn7713(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7717 github.com/goccy/spidermonkeywasm2go/p4.Fn7717
func Fn7717(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7719 github.com/goccy/spidermonkeywasm2go/p6.Fn7719
func Fn7719(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7909 github.com/goccy/spidermonkeywasm2go/p6.Fn7909
func Fn7909(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7911 github.com/goccy/spidermonkeywasm2go/p6.Fn7911
func Fn7911(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7912 github.com/goccy/spidermonkeywasm2go/p6.Fn7912
func Fn7912(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7917 github.com/goccy/spidermonkeywasm2go/p6.Fn7917
func Fn7917(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7929 github.com/goccy/spidermonkeywasm2go/p6.Fn7929
func Fn7929(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7930 github.com/goccy/spidermonkeywasm2go/p6.Fn7930
func Fn7930(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7933 github.com/goccy/spidermonkeywasm2go/p6.Fn7933
func Fn7933(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7936 github.com/goccy/spidermonkeywasm2go/p6.Fn7936
func Fn7936(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7938 github.com/goccy/spidermonkeywasm2go/p4.Fn7938
func Fn7938(m *base.Module, l0 int32) int32

//go:linkname Fn7940 github.com/goccy/spidermonkeywasm2go/p6.Fn7940
func Fn7940(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7941 github.com/goccy/spidermonkeywasm2go/p6.Fn7941
func Fn7941(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7942 github.com/goccy/spidermonkeywasm2go/p4.Fn7942
func Fn7942(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7944 github.com/goccy/spidermonkeywasm2go/p6.Fn7944
func Fn7944(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7946 github.com/goccy/spidermonkeywasm2go/p6.Fn7946
func Fn7946(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7949 github.com/goccy/spidermonkeywasm2go/p4.Fn7949
func Fn7949(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7957 github.com/goccy/spidermonkeywasm2go/p6.Fn7957
func Fn7957(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn7958 github.com/goccy/spidermonkeywasm2go/p6.Fn7958
func Fn7958(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7961 github.com/goccy/spidermonkeywasm2go/p6.Fn7961
func Fn7961(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7966 github.com/goccy/spidermonkeywasm2go/p7.Fn7966
func Fn7966(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7972 github.com/goccy/spidermonkeywasm2go/p6.Fn7972
func Fn7972(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7976 github.com/goccy/spidermonkeywasm2go/p6.Fn7976
func Fn7976(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7980 github.com/goccy/spidermonkeywasm2go/p6.Fn7980
func Fn7980(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7981 github.com/goccy/spidermonkeywasm2go/p6.Fn7981
func Fn7981(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7982 github.com/goccy/spidermonkeywasm2go/p6.Fn7982
func Fn7982(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7983 github.com/goccy/spidermonkeywasm2go/p6.Fn7983
func Fn7983(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7984 github.com/goccy/spidermonkeywasm2go/p6.Fn7984
func Fn7984(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8085 github.com/goccy/spidermonkeywasm2go/p7.Fn8085
func Fn8085(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8090 github.com/goccy/spidermonkeywasm2go/p6.Fn8090
func Fn8090(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8100 github.com/goccy/spidermonkeywasm2go/p7.Fn8100
func Fn8100(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8102 github.com/goccy/spidermonkeywasm2go/p6.Fn8102
func Fn8102(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8103 github.com/goccy/spidermonkeywasm2go/p6.Fn8103
func Fn8103(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8107 github.com/goccy/spidermonkeywasm2go/p6.Fn8107
func Fn8107(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8110 github.com/goccy/spidermonkeywasm2go/p6.Fn8110
func Fn8110(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8137 github.com/goccy/spidermonkeywasm2go/p4.Fn8137
func Fn8137(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8142 github.com/goccy/spidermonkeywasm2go/p1.Fn8142
func Fn8142(m *base.Module, l0 int32) int32

//go:linkname Fn8145 github.com/goccy/spidermonkeywasm2go/p6.Fn8145
func Fn8145(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8146 github.com/goccy/spidermonkeywasm2go/p1.Fn8146
func Fn8146(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8147 github.com/goccy/spidermonkeywasm2go/p1.Fn8147
func Fn8147(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8150 github.com/goccy/spidermonkeywasm2go/p7.Fn8150
func Fn8150(m *base.Module, l0 int32) int32

//go:linkname Fn8152 github.com/goccy/spidermonkeywasm2go/p1.Fn8152
func Fn8152(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8153 github.com/goccy/spidermonkeywasm2go/p6.Fn8153
func Fn8153(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8154 github.com/goccy/spidermonkeywasm2go/p7.Fn8154
func Fn8154(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8159 github.com/goccy/spidermonkeywasm2go/p2.Fn8159
func Fn8159(m *base.Module) int32

//go:linkname Fn8163 github.com/goccy/spidermonkeywasm2go/p6.Fn8163
func Fn8163(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8164 github.com/goccy/spidermonkeywasm2go/p3.Fn8164
func Fn8164(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8182 github.com/goccy/spidermonkeywasm2go/p1.Fn8182
func Fn8182(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8198 github.com/goccy/spidermonkeywasm2go/p6.Fn8198
func Fn8198(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8200 github.com/goccy/spidermonkeywasm2go/p4.Fn8200
func Fn8200(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8208 github.com/goccy/spidermonkeywasm2go/p6.Fn8208
func Fn8208(m *base.Module, l0 int32) int32

//go:linkname Fn8252 github.com/goccy/spidermonkeywasm2go/p6.Fn8252
func Fn8252(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8253 github.com/goccy/spidermonkeywasm2go/p6.Fn8253
func Fn8253(m *base.Module, l0 int32) int32

//go:linkname Fn8256 github.com/goccy/spidermonkeywasm2go/p4.Fn8256
func Fn8256(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8258 github.com/goccy/spidermonkeywasm2go/p6.Fn8258
func Fn8258(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8265 github.com/goccy/spidermonkeywasm2go/p7.Fn8265
func Fn8265(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8266 github.com/goccy/spidermonkeywasm2go/p6.Fn8266
func Fn8266(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8268 github.com/goccy/spidermonkeywasm2go/p6.Fn8268
func Fn8268(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8273 github.com/goccy/spidermonkeywasm2go/p6.Fn8273
func Fn8273(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8275 github.com/goccy/spidermonkeywasm2go/p3.Fn8275
func Fn8275(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8279 github.com/goccy/spidermonkeywasm2go/p4.Fn8279
func Fn8279(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8287 github.com/goccy/spidermonkeywasm2go/p7.Fn8287
func Fn8287(m *base.Module) float64

//go:linkname Fn8291 github.com/goccy/spidermonkeywasm2go/p6.Fn8291
func Fn8291(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8298 github.com/goccy/spidermonkeywasm2go/p3.Fn8298
func Fn8298(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8308 github.com/goccy/spidermonkeywasm2go/p7.Fn8308
func Fn8308(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8311 github.com/goccy/spidermonkeywasm2go/p6.Fn8311
func Fn8311(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8312 github.com/goccy/spidermonkeywasm2go/p4.Fn8312
func Fn8312(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8328 github.com/goccy/spidermonkeywasm2go/p4.Fn8328
func Fn8328(m *base.Module, l0 int32) int32

//go:linkname Fn8335 github.com/goccy/spidermonkeywasm2go/p6.Fn8335
func Fn8335(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8347 github.com/goccy/spidermonkeywasm2go/p7.Fn8347
func Fn8347(m *base.Module, l0 int32)

//go:linkname Fn8348 github.com/goccy/spidermonkeywasm2go/p7.Fn8348
func Fn8348(m *base.Module, l0 int32)

//go:linkname Fn8350 github.com/goccy/spidermonkeywasm2go/p6.Fn8350
func Fn8350(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8353 github.com/goccy/spidermonkeywasm2go/p6.Fn8353
func Fn8353(m *base.Module, l0 int32)

//go:linkname Fn8355 github.com/goccy/spidermonkeywasm2go/p6.Fn8355
func Fn8355(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8357 github.com/goccy/spidermonkeywasm2go/p6.Fn8357
func Fn8357(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8378 github.com/goccy/spidermonkeywasm2go/p1.Fn8378
func Fn8378(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8380 github.com/goccy/spidermonkeywasm2go/p3.Fn8380
func Fn8380(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn8381 github.com/goccy/spidermonkeywasm2go/p7.Fn8381
func Fn8381(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8382 github.com/goccy/spidermonkeywasm2go/p7.Fn8382
func Fn8382(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8388 github.com/goccy/spidermonkeywasm2go/p1.Fn8388
func Fn8388(m *base.Module) int32

//go:linkname Fn8389 github.com/goccy/spidermonkeywasm2go/p6.Fn8389
func Fn8389(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8390 github.com/goccy/spidermonkeywasm2go/p1.Fn8390
func Fn8390(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8391 github.com/goccy/spidermonkeywasm2go/p7.Fn8391
func Fn8391(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8414 github.com/goccy/spidermonkeywasm2go/p6.Fn8414
func Fn8414(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8422 github.com/goccy/spidermonkeywasm2go/p6.Fn8422
func Fn8422(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8423 github.com/goccy/spidermonkeywasm2go/p6.Fn8423
func Fn8423(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8425 github.com/goccy/spidermonkeywasm2go/p6.Fn8425
func Fn8425(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8426 github.com/goccy/spidermonkeywasm2go/p4.Fn8426
func Fn8426(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8428 github.com/goccy/spidermonkeywasm2go/p7.Fn8428
func Fn8428(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8430 github.com/goccy/spidermonkeywasm2go/p6.Fn8430
func Fn8430(m *base.Module, l0 int32) int32

//go:linkname Fn8444 github.com/goccy/spidermonkeywasm2go/p6.Fn8444
func Fn8444(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8448 github.com/goccy/spidermonkeywasm2go/p6.Fn8448
func Fn8448(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8449 github.com/goccy/spidermonkeywasm2go/p6.Fn8449
func Fn8449(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8456 github.com/goccy/spidermonkeywasm2go/p6.Fn8456
func Fn8456(m *base.Module, l0 int32) int32

//go:linkname Fn8457 github.com/goccy/spidermonkeywasm2go/p6.Fn8457
func Fn8457(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8462 github.com/goccy/spidermonkeywasm2go/p1.Fn8462
func Fn8462(m *base.Module, l0 int32) int32

//go:linkname Fn8476 github.com/goccy/spidermonkeywasm2go/p4.Fn8476
func Fn8476(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8479 github.com/goccy/spidermonkeywasm2go/p6.Fn8479
func Fn8479(m *base.Module, l0 int32) int32

//go:linkname Fn8488 github.com/goccy/spidermonkeywasm2go/p4.Fn8488
func Fn8488(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8489 github.com/goccy/spidermonkeywasm2go/p6.Fn8489
func Fn8489(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8509 github.com/goccy/spidermonkeywasm2go/p6.Fn8509
func Fn8509(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8512 github.com/goccy/spidermonkeywasm2go/p7.Fn8512
func Fn8512(m *base.Module, l0 int32)

//go:linkname Fn8516 github.com/goccy/spidermonkeywasm2go/p6.Fn8516
func Fn8516(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8517 github.com/goccy/spidermonkeywasm2go/p6.Fn8517
func Fn8517(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8518 github.com/goccy/spidermonkeywasm2go/p6.Fn8518
func Fn8518(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8519 github.com/goccy/spidermonkeywasm2go/p3.Fn8519
func Fn8519(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8522 github.com/goccy/spidermonkeywasm2go/p7.Fn8522
func Fn8522(m *base.Module, l0 int32) int32

//go:linkname Fn8534 github.com/goccy/spidermonkeywasm2go/p6.Fn8534
func Fn8534(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8535 github.com/goccy/spidermonkeywasm2go/p6.Fn8535
func Fn8535(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8536 github.com/goccy/spidermonkeywasm2go/p6.Fn8536
func Fn8536(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8541 github.com/goccy/spidermonkeywasm2go/p6.Fn8541
func Fn8541(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8542 github.com/goccy/spidermonkeywasm2go/p7.Fn8542
func Fn8542(m *base.Module, l0 int32) int32

//go:linkname Fn8543 github.com/goccy/spidermonkeywasm2go/p6.Fn8543
func Fn8543(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8544 github.com/goccy/spidermonkeywasm2go/p7.Fn8544
func Fn8544(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8552 github.com/goccy/spidermonkeywasm2go/p7.Fn8552
func Fn8552(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8558 github.com/goccy/spidermonkeywasm2go/p4.Fn8558
func Fn8558(m *base.Module, l0 int32) int32

//go:linkname Fn8562 github.com/goccy/spidermonkeywasm2go/p6.Fn8562
func Fn8562(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8564 github.com/goccy/spidermonkeywasm2go/p6.Fn8564
func Fn8564(m *base.Module, l0 int32) int32

//go:linkname Fn8567 github.com/goccy/spidermonkeywasm2go/p7.Fn8567
func Fn8567(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8573 github.com/goccy/spidermonkeywasm2go/p4.Fn8573
func Fn8573(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8613 github.com/goccy/spidermonkeywasm2go/p7.Fn8613
func Fn8613(m *base.Module, l0 int32)

//go:linkname Fn8614 github.com/goccy/spidermonkeywasm2go/p7.Fn8614
func Fn8614(m *base.Module, l0 int32)

//go:linkname Fn8618 github.com/goccy/spidermonkeywasm2go/p6.Fn8618
func Fn8618(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8620 github.com/goccy/spidermonkeywasm2go/p6.Fn8620
func Fn8620(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8621 github.com/goccy/spidermonkeywasm2go/p1.Fn8621
func Fn8621(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8623 github.com/goccy/spidermonkeywasm2go/p6.Fn8623
func Fn8623(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8624 github.com/goccy/spidermonkeywasm2go/p1.Fn8624
func Fn8624(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8625 github.com/goccy/spidermonkeywasm2go/p6.Fn8625
func Fn8625(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8626 github.com/goccy/spidermonkeywasm2go/p1.Fn8626
func Fn8626(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8630 github.com/goccy/spidermonkeywasm2go/p6.Fn8630
func Fn8630(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8631 github.com/goccy/spidermonkeywasm2go/p1.Fn8631
func Fn8631(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8633 github.com/goccy/spidermonkeywasm2go/p1.Fn8633
func Fn8633(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8670 github.com/goccy/spidermonkeywasm2go/p6.Fn8670
func Fn8670(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8673 github.com/goccy/spidermonkeywasm2go/p6.Fn8673
func Fn8673(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn8678 github.com/goccy/spidermonkeywasm2go/p7.Fn8678
func Fn8678(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn8679 github.com/goccy/spidermonkeywasm2go/p4.Fn8679
func Fn8679(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8693 github.com/goccy/spidermonkeywasm2go/p4.Fn8693
func Fn8693(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8697 github.com/goccy/spidermonkeywasm2go/p3.Fn8697
func Fn8697(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn8704 github.com/goccy/spidermonkeywasm2go/p6.Fn8704
func Fn8704(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8711 github.com/goccy/spidermonkeywasm2go/p7.Fn8711
func Fn8711(m *base.Module, l0 int32)

//go:linkname Fn8714 github.com/goccy/spidermonkeywasm2go/p6.Fn8714
func Fn8714(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8716 github.com/goccy/spidermonkeywasm2go/p6.Fn8716
func Fn8716(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8722 github.com/goccy/spidermonkeywasm2go/p7.Fn8722
func Fn8722(m *base.Module, l0 int32)

//go:linkname Fn8723 github.com/goccy/spidermonkeywasm2go/p6.Fn8723
func Fn8723(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8730 github.com/goccy/spidermonkeywasm2go/p3.Fn8730
func Fn8730(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8738 github.com/goccy/spidermonkeywasm2go/p6.Fn8738
func Fn8738(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8739 github.com/goccy/spidermonkeywasm2go/p6.Fn8739
func Fn8739(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8740 github.com/goccy/spidermonkeywasm2go/p6.Fn8740
func Fn8740(m *base.Module, l0 int32) int32

//go:linkname Fn8743 github.com/goccy/spidermonkeywasm2go/p7.Fn8743
func Fn8743(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8744 github.com/goccy/spidermonkeywasm2go/p3.Fn8744
func Fn8744(m *base.Module, l0 int32) int32

//go:linkname Fn8749 github.com/goccy/spidermonkeywasm2go/p7.Fn8749
func Fn8749(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8751 github.com/goccy/spidermonkeywasm2go/p7.Fn8751
func Fn8751(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8759 github.com/goccy/spidermonkeywasm2go/p4.Fn8759
func Fn8759(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8761 github.com/goccy/spidermonkeywasm2go/p4.Fn8761
func Fn8761(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8762 github.com/goccy/spidermonkeywasm2go/p6.Fn8762
func Fn8762(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32) int32

//go:linkname Fn8778 github.com/goccy/spidermonkeywasm2go/p6.Fn8778
func Fn8778(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8781 github.com/goccy/spidermonkeywasm2go/p6.Fn8781
func Fn8781(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8783 github.com/goccy/spidermonkeywasm2go/p6.Fn8783
func Fn8783(m *base.Module, l0 int32, l1 float64, l2 float64, l3 int32) int32

//go:linkname Fn8832 github.com/goccy/spidermonkeywasm2go/p6.Fn8832
func Fn8832(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8844 github.com/goccy/spidermonkeywasm2go/p4.Fn8844
func Fn8844(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8855 github.com/goccy/spidermonkeywasm2go/p6.Fn8855
func Fn8855(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8902 github.com/goccy/spidermonkeywasm2go/p6.Fn8902
func Fn8902(m *base.Module, l0 int32) int32

//go:linkname Fn8909 github.com/goccy/spidermonkeywasm2go/p4.Fn8909
func Fn8909(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8910 github.com/goccy/spidermonkeywasm2go/p6.Fn8910
func Fn8910(m *base.Module, l0 float64) float64

//go:linkname Fn8915 github.com/goccy/spidermonkeywasm2go/p6.Fn8915
func Fn8915(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64

//go:linkname Fn8916 github.com/goccy/spidermonkeywasm2go/p6.Fn8916
func Fn8916(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32) int32

//go:linkname Fn8917 github.com/goccy/spidermonkeywasm2go/p6.Fn8917
func Fn8917(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8922 github.com/goccy/spidermonkeywasm2go/p7.Fn8922
func Fn8922(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8923 github.com/goccy/spidermonkeywasm2go/p7.Fn8923
func Fn8923(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8926 github.com/goccy/spidermonkeywasm2go/p6.Fn8926
func Fn8926(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8940 github.com/goccy/spidermonkeywasm2go/p6.Fn8940
func Fn8940(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn8941 github.com/goccy/spidermonkeywasm2go/p6.Fn8941
func Fn8941(m *base.Module, l0 float64) float64

//go:linkname Fn8942 github.com/goccy/spidermonkeywasm2go/p6.Fn8942
func Fn8942(m *base.Module, l0 int32, l1 int32, l2 float64, l3 float64, l4 int32) float64

//go:linkname Fn8945 github.com/goccy/spidermonkeywasm2go/p6.Fn8945
func Fn8945(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8947 github.com/goccy/spidermonkeywasm2go/p7.Fn8947
func Fn8947(m *base.Module, l0 int32)

//go:linkname Fn8982 github.com/goccy/spidermonkeywasm2go/p7.Fn8982
func Fn8982(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8984 github.com/goccy/spidermonkeywasm2go/p3.Fn8984
func Fn8984(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn9056 github.com/goccy/spidermonkeywasm2go/p2.Fn9056
func Fn9056(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9065 github.com/goccy/spidermonkeywasm2go/p4.Fn9065
func Fn9065(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9068 github.com/goccy/spidermonkeywasm2go/p6.Fn9068
func Fn9068(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9088 github.com/goccy/spidermonkeywasm2go/p2.Fn9088
func Fn9088(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9090 github.com/goccy/spidermonkeywasm2go/p4.Fn9090
func Fn9090(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9094 github.com/goccy/spidermonkeywasm2go/p4.Fn9094
func Fn9094(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9113 github.com/goccy/spidermonkeywasm2go/p6.Fn9113
func Fn9113(m *base.Module, l0 int32)

//go:linkname Fn9118 github.com/goccy/spidermonkeywasm2go/p6.Fn9118
func Fn9118(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9120 github.com/goccy/spidermonkeywasm2go/p6.Fn9120
func Fn9120(m *base.Module, l0 int32) int32

//go:linkname Fn9121 github.com/goccy/spidermonkeywasm2go/p6.Fn9121
func Fn9121(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9122 github.com/goccy/spidermonkeywasm2go/p6.Fn9122
func Fn9122(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9123 github.com/goccy/spidermonkeywasm2go/p6.Fn9123
func Fn9123(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9128 github.com/goccy/spidermonkeywasm2go/p6.Fn9128
func Fn9128(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9141 github.com/goccy/spidermonkeywasm2go/p6.Fn9141
func Fn9141(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9164 github.com/goccy/spidermonkeywasm2go/p7.Fn9164
func Fn9164(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn9165 github.com/goccy/spidermonkeywasm2go/p6.Fn9165
func Fn9165(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn9166 github.com/goccy/spidermonkeywasm2go/p6.Fn9166
func Fn9166(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9175 github.com/goccy/spidermonkeywasm2go/p6.Fn9175
func Fn9175(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9226 github.com/goccy/spidermonkeywasm2go/p6.Fn9226
func Fn9226(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9249 github.com/goccy/spidermonkeywasm2go/p6.Fn9249
func Fn9249(m *base.Module, l0 int32)

//go:linkname Fn9267 github.com/goccy/spidermonkeywasm2go/p6.Fn9267
func Fn9267(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn9268 github.com/goccy/spidermonkeywasm2go/p6.Fn9268
func Fn9268(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32)

//go:linkname Fn9269 github.com/goccy/spidermonkeywasm2go/p6.Fn9269
func Fn9269(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9276 github.com/goccy/spidermonkeywasm2go/p4.Fn9276
func Fn9276(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9277 github.com/goccy/spidermonkeywasm2go/p6.Fn9277
func Fn9277(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9279 github.com/goccy/spidermonkeywasm2go/p6.Fn9279
func Fn9279(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9281 github.com/goccy/spidermonkeywasm2go/p6.Fn9281
func Fn9281(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

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

//go:linkname Fn9292 github.com/goccy/spidermonkeywasm2go/p6.Fn9292
func Fn9292(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9296 github.com/goccy/spidermonkeywasm2go/p6.Fn9296
func Fn9296(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9297 github.com/goccy/spidermonkeywasm2go/p6.Fn9297
func Fn9297(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9301 github.com/goccy/spidermonkeywasm2go/p3.Fn9301
func Fn9301(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9302 github.com/goccy/spidermonkeywasm2go/p3.Fn9302
func Fn9302(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9303 github.com/goccy/spidermonkeywasm2go/p3.Fn9303
func Fn9303(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9304 github.com/goccy/spidermonkeywasm2go/p6.Fn9304
func Fn9304(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9311 github.com/goccy/spidermonkeywasm2go/p6.Fn9311
func Fn9311(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn9313 github.com/goccy/spidermonkeywasm2go/p6.Fn9313
func Fn9313(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9314 github.com/goccy/spidermonkeywasm2go/p6.Fn9314
func Fn9314(m *base.Module, l0 int32) int32

//go:linkname Fn9316 github.com/goccy/spidermonkeywasm2go/p6.Fn9316
func Fn9316(m *base.Module, l0 int32)

//go:linkname Fn9324 github.com/goccy/spidermonkeywasm2go/p6.Fn9324
func Fn9324(m *base.Module, l0 int32) int32

//go:linkname Fn9325 github.com/goccy/spidermonkeywasm2go/p3.Fn9325
func Fn9325(m *base.Module, l0 int32) int32

//go:linkname Fn9329 github.com/goccy/spidermonkeywasm2go/p3.Fn9329
func Fn9329(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9330 github.com/goccy/spidermonkeywasm2go/p4.Fn9330
func Fn9330(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9335 github.com/goccy/spidermonkeywasm2go/p4.Fn9335
func Fn9335(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9339 github.com/goccy/spidermonkeywasm2go/p6.Fn9339
func Fn9339(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9343 github.com/goccy/spidermonkeywasm2go/p3.Fn9343
func Fn9343(m *base.Module, l0 int32)

//go:linkname Fn9344 github.com/goccy/spidermonkeywasm2go/p4.Fn9344
func Fn9344(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9353 github.com/goccy/spidermonkeywasm2go/p6.Fn9353
func Fn9353(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9355 github.com/goccy/spidermonkeywasm2go/p7.Fn9355
func Fn9355(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn9356 github.com/goccy/spidermonkeywasm2go/p6.Fn9356
func Fn9356(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9369 github.com/goccy/spidermonkeywasm2go/p6.Fn9369
func Fn9369(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9390 github.com/goccy/spidermonkeywasm2go/p7.Fn9390
func Fn9390(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9448 github.com/goccy/spidermonkeywasm2go/p7.Fn9448
func Fn9448(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn9461 github.com/goccy/spidermonkeywasm2go/p3.Fn9461
func Fn9461(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9503 github.com/goccy/spidermonkeywasm2go/p6.Fn9503
func Fn9503(m *base.Module, l0 int32) int32

//go:linkname Fn9509 github.com/goccy/spidermonkeywasm2go/p6.Fn9509
func Fn9509(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9510 github.com/goccy/spidermonkeywasm2go/p6.Fn9510
func Fn9510(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9524 github.com/goccy/spidermonkeywasm2go/p6.Fn9524
func Fn9524(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9559 github.com/goccy/spidermonkeywasm2go/p7.Fn9559
func Fn9559(m *base.Module, l0 int32) int32

//go:linkname Fn9597 github.com/goccy/spidermonkeywasm2go/p3.Fn9597
func Fn9597(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9632 github.com/goccy/spidermonkeywasm2go/p4.Fn9632
func Fn9632(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9678 github.com/goccy/spidermonkeywasm2go/p4.Fn9678
func Fn9678(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9679 github.com/goccy/spidermonkeywasm2go/p4.Fn9679
func Fn9679(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn9680 github.com/goccy/spidermonkeywasm2go/p4.Fn9680
func Fn9680(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9683 github.com/goccy/spidermonkeywasm2go/p6.Fn9683
func Fn9683(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9684 github.com/goccy/spidermonkeywasm2go/p6.Fn9684
func Fn9684(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9694 github.com/goccy/spidermonkeywasm2go/p2.Fn9694
func Fn9694(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9731 github.com/goccy/spidermonkeywasm2go/p6.Fn9731
func Fn9731(m *base.Module) int32

//go:linkname Fn9746 github.com/goccy/spidermonkeywasm2go/p6.Fn9746
func Fn9746(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9755 github.com/goccy/spidermonkeywasm2go/p6.Fn9755
func Fn9755(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9759 github.com/goccy/spidermonkeywasm2go/p6.Fn9759
func Fn9759(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9788 github.com/goccy/spidermonkeywasm2go/p6.Fn9788
func Fn9788(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9839 github.com/goccy/spidermonkeywasm2go/p3.Fn9839
func Fn9839(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9840 github.com/goccy/spidermonkeywasm2go/p6.Fn9840
func Fn9840(m *base.Module, l0 int32) int32

//go:linkname Fn9841 github.com/goccy/spidermonkeywasm2go/p2.Fn9841
func Fn9841(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9854 github.com/goccy/spidermonkeywasm2go/p6.Fn9854
func Fn9854(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9932 github.com/goccy/spidermonkeywasm2go/p7.Fn9932
func Fn9932(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9946 github.com/goccy/spidermonkeywasm2go/p7.Fn9946
func Fn9946(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9947 github.com/goccy/spidermonkeywasm2go/p7.Fn9947
func Fn9947(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9953 github.com/goccy/spidermonkeywasm2go/p6.Fn9953
func Fn9953(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9957 github.com/goccy/spidermonkeywasm2go/p2.Fn9957
func Fn9957(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9960 github.com/goccy/spidermonkeywasm2go/p6.Fn9960
func Fn9960(m *base.Module, l0 int32)

//go:linkname Fn9962 github.com/goccy/spidermonkeywasm2go/p4.Fn9962
func Fn9962(m *base.Module, l0 int32)

//go:linkname Fn9995 github.com/goccy/spidermonkeywasm2go/p6.Fn9995
func Fn9995(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9996 github.com/goccy/spidermonkeywasm2go/p6.Fn9996
func Fn9996(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn9997 github.com/goccy/spidermonkeywasm2go/p6.Fn9997
func Fn9997(m *base.Module, l0 int32, l1 float64) int32

//go:linkname Fn9999 github.com/goccy/spidermonkeywasm2go/p4.Fn9999
func Fn9999(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 int32, l5 int32, l6 int32)

//go:linkname Fn10000 github.com/goccy/spidermonkeywasm2go/p3.Fn10000
func Fn10000(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10006 github.com/goccy/spidermonkeywasm2go/p6.Fn10006
func Fn10006(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10009 github.com/goccy/spidermonkeywasm2go/p4.Fn10009
func Fn10009(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn10022 github.com/goccy/spidermonkeywasm2go/p6.Fn10022
func Fn10022(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn10023 github.com/goccy/spidermonkeywasm2go/p6.Fn10023
func Fn10023(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 float64, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn10085 github.com/goccy/spidermonkeywasm2go/p6.Fn10085
func Fn10085(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10091 github.com/goccy/spidermonkeywasm2go/p6.Fn10091
func Fn10091(m *base.Module, l0 int32) int32

//go:linkname Fn10092 github.com/goccy/spidermonkeywasm2go/p6.Fn10092
func Fn10092(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10100 github.com/goccy/spidermonkeywasm2go/p6.Fn10100
func Fn10100(m *base.Module, l0 int32) int32

//go:linkname Fn10105 github.com/goccy/spidermonkeywasm2go/p1.Fn10105
func Fn10105(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn10114 github.com/goccy/spidermonkeywasm2go/p4.Fn10114
func Fn10114(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10117 github.com/goccy/spidermonkeywasm2go/p3.Fn10117
func Fn10117(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn10140 github.com/goccy/spidermonkeywasm2go/p6.Fn10140
func Fn10140(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10144 github.com/goccy/spidermonkeywasm2go/p7.Fn10144
func Fn10144(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10145 github.com/goccy/spidermonkeywasm2go/p7.Fn10145
func Fn10145(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10146 github.com/goccy/spidermonkeywasm2go/p7.Fn10146
func Fn10146(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10147 github.com/goccy/spidermonkeywasm2go/p4.Fn10147
func Fn10147(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10196 github.com/goccy/spidermonkeywasm2go/p3.Fn10196
func Fn10196(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10211 github.com/goccy/spidermonkeywasm2go/p6.Fn10211
func Fn10211(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn10212 github.com/goccy/spidermonkeywasm2go/p6.Fn10212
func Fn10212(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10219 github.com/goccy/spidermonkeywasm2go/p4.Fn10219
func Fn10219(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10235 github.com/goccy/spidermonkeywasm2go/p6.Fn10235
func Fn10235(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10236 github.com/goccy/spidermonkeywasm2go/p6.Fn10236
func Fn10236(m *base.Module, l0 int32)

//go:linkname Fn10240 github.com/goccy/spidermonkeywasm2go/p6.Fn10240
func Fn10240(m *base.Module, l0 int32) int32

//go:linkname Fn10242 github.com/goccy/spidermonkeywasm2go/p1.Fn10242
func Fn10242(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10247 github.com/goccy/spidermonkeywasm2go/p6.Fn10247
func Fn10247(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10250 github.com/goccy/spidermonkeywasm2go/p6.Fn10250
func Fn10250(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10283 github.com/goccy/spidermonkeywasm2go/p6.Fn10283
func Fn10283(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10303 github.com/goccy/spidermonkeywasm2go/p2.Fn10303
func Fn10303(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10328 github.com/goccy/spidermonkeywasm2go/p1.Fn10328
func Fn10328(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10331 github.com/goccy/spidermonkeywasm2go/p1.Fn10331
func Fn10331(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10338 github.com/goccy/spidermonkeywasm2go/p6.Fn10338
func Fn10338(m *base.Module, l0 int32) int32

//go:linkname Fn10372 github.com/goccy/spidermonkeywasm2go/p6.Fn10372
func Fn10372(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10403 github.com/goccy/spidermonkeywasm2go/p6.Fn10403
func Fn10403(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10423 github.com/goccy/spidermonkeywasm2go/p2.Fn10423
func Fn10423(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn10451 github.com/goccy/spidermonkeywasm2go/p2.Fn10451
func Fn10451(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10452 github.com/goccy/spidermonkeywasm2go/p6.Fn10452
func Fn10452(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10460 github.com/goccy/spidermonkeywasm2go/p2.Fn10460
func Fn10460(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn10476 github.com/goccy/spidermonkeywasm2go/p6.Fn10476
func Fn10476(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10478 github.com/goccy/spidermonkeywasm2go/p3.Fn10478
func Fn10478(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10509 github.com/goccy/spidermonkeywasm2go/p6.Fn10509
func Fn10509(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10547 github.com/goccy/spidermonkeywasm2go/p6.Fn10547
func Fn10547(m *base.Module, l0 int32) int32

//go:linkname Fn10550 github.com/goccy/spidermonkeywasm2go/p6.Fn10550
func Fn10550(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10551 github.com/goccy/spidermonkeywasm2go/p6.Fn10551
func Fn10551(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10552 github.com/goccy/spidermonkeywasm2go/p6.Fn10552
func Fn10552(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10567 github.com/goccy/spidermonkeywasm2go/p3.Fn10567
func Fn10567(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10573 github.com/goccy/spidermonkeywasm2go/p6.Fn10573
func Fn10573(m *base.Module, l0 int32) int32

//go:linkname Fn10578 github.com/goccy/spidermonkeywasm2go/p4.Fn10578
func Fn10578(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10579 github.com/goccy/spidermonkeywasm2go/p4.Fn10579
func Fn10579(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10581 github.com/goccy/spidermonkeywasm2go/p6.Fn10581
func Fn10581(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10588 github.com/goccy/spidermonkeywasm2go/p6.Fn10588
func Fn10588(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10598 github.com/goccy/spidermonkeywasm2go/p4.Fn10598
func Fn10598(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10617 github.com/goccy/spidermonkeywasm2go/p3.Fn10617
func Fn10617(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10620 github.com/goccy/spidermonkeywasm2go/p4.Fn10620
func Fn10620(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn10632 github.com/goccy/spidermonkeywasm2go/p4.Fn10632
func Fn10632(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10633 github.com/goccy/spidermonkeywasm2go/p4.Fn10633
func Fn10633(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10667 github.com/goccy/spidermonkeywasm2go/p6.Fn10667
func Fn10667(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10671 github.com/goccy/spidermonkeywasm2go/p7.Fn10671
func Fn10671(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10672 github.com/goccy/spidermonkeywasm2go/p4.Fn10672
func Fn10672(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10687 github.com/goccy/spidermonkeywasm2go/p6.Fn10687
func Fn10687(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10762 github.com/goccy/spidermonkeywasm2go/p4.Fn10762
func Fn10762(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10765 github.com/goccy/spidermonkeywasm2go/p6.Fn10765
func Fn10765(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10766 github.com/goccy/spidermonkeywasm2go/p4.Fn10766
func Fn10766(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10767 github.com/goccy/spidermonkeywasm2go/p4.Fn10767
func Fn10767(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10768 github.com/goccy/spidermonkeywasm2go/p4.Fn10768
func Fn10768(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10778 github.com/goccy/spidermonkeywasm2go/p6.Fn10778
func Fn10778(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10793 github.com/goccy/spidermonkeywasm2go/p7.Fn10793
func Fn10793(m *base.Module, l0 int32) int32

//go:linkname Fn10813 github.com/goccy/spidermonkeywasm2go/p6.Fn10813
func Fn10813(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10862 github.com/goccy/spidermonkeywasm2go/p6.Fn10862
func Fn10862(m *base.Module, l0 int32) int32

//go:linkname Fn10863 github.com/goccy/spidermonkeywasm2go/p4.Fn10863
func Fn10863(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10870 github.com/goccy/spidermonkeywasm2go/p4.Fn10870
func Fn10870(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10872 github.com/goccy/spidermonkeywasm2go/p7.Fn10872
func Fn10872(m *base.Module, l0 int32) int32

//go:linkname Fn10877 github.com/goccy/spidermonkeywasm2go/p7.Fn10877
func Fn10877(m *base.Module, l0 int32) int32

//go:linkname Fn10879 github.com/goccy/spidermonkeywasm2go/p6.Fn10879
func Fn10879(m *base.Module, l0 int32) int32

//go:linkname Fn10897 github.com/goccy/spidermonkeywasm2go/p3.Fn10897
func Fn10897(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10902 github.com/goccy/spidermonkeywasm2go/p6.Fn10902
func Fn10902(m *base.Module, l0 float64, l1 float64, l2 int32) float64

//go:linkname Fn10904 github.com/goccy/spidermonkeywasm2go/p6.Fn10904
func Fn10904(m *base.Module, l0 int32)

//go:linkname Fn10912 github.com/goccy/spidermonkeywasm2go/p4.Fn10912
func Fn10912(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10916 github.com/goccy/spidermonkeywasm2go/p6.Fn10916
func Fn10916(m *base.Module, l0 int32)

//go:linkname Fn10926 github.com/goccy/spidermonkeywasm2go/p7.Fn10926
func Fn10926(m *base.Module)

//go:linkname Fn10927 github.com/goccy/spidermonkeywasm2go/p2.Fn10927
func Fn10927(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10928 github.com/goccy/spidermonkeywasm2go/p2.Fn10928
func Fn10928(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10929 github.com/goccy/spidermonkeywasm2go/p2.Fn10929
func Fn10929(m *base.Module, l0 float64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn10936 github.com/goccy/spidermonkeywasm2go/p6.Fn10936
func Fn10936(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10937 github.com/goccy/spidermonkeywasm2go/p6.Fn10937
func Fn10937(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10938 github.com/goccy/spidermonkeywasm2go/p6.Fn10938
func Fn10938(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10939 github.com/goccy/spidermonkeywasm2go/p4.Fn10939
func Fn10939(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn10941 github.com/goccy/spidermonkeywasm2go/p3.Fn10941
func Fn10941(m *base.Module, l0 float64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn10947 github.com/goccy/spidermonkeywasm2go/p2.Fn10947
func Fn10947(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64

//go:linkname Fn10953 github.com/goccy/spidermonkeywasm2go/p2.Fn10953
func Fn10953(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64

//go:linkname Fn10957 github.com/goccy/spidermonkeywasm2go/p2.Fn10957
func Fn10957(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10960 github.com/goccy/spidermonkeywasm2go/p6.Fn10960
func Fn10960(m *base.Module, l0 int32, l1 int32, l2 int32)
