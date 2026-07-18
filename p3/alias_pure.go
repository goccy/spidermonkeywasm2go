//go:build (!amd64 && !arm64) || purego

package p3

import (
	base "github.com/goccy/spidermonkeywasm2go/base"
	_ "unsafe"
)

//go:linkname Fn36 github.com/goccy/spidermonkeywasm2go/p7.Fn36
func Fn36(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn95 github.com/goccy/spidermonkeywasm2go/p7.Fn95
func Fn95(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn96 github.com/goccy/spidermonkeywasm2go/p7.Fn96
func Fn96(m *base.Module, l0 int64)

//go:linkname Fn104 github.com/goccy/spidermonkeywasm2go/p7.Fn104
func Fn104(m *base.Module, l0 int32)

//go:linkname Fn109 github.com/goccy/spidermonkeywasm2go/p7.Fn109
func Fn109(m *base.Module, l0 int32) int32

//go:linkname Fn110 github.com/goccy/spidermonkeywasm2go/p6.Fn110
func Fn110(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn111 github.com/goccy/spidermonkeywasm2go/p7.Fn111
func Fn111(m *base.Module, l0 int32)

//go:linkname Fn117 github.com/goccy/spidermonkeywasm2go/p7.Fn117
func Fn117(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn118 github.com/goccy/spidermonkeywasm2go/p7.Fn118
func Fn118(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn119 github.com/goccy/spidermonkeywasm2go/p7.Fn119
func Fn119(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn120 github.com/goccy/spidermonkeywasm2go/p7.Fn120
func Fn120(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn138 github.com/goccy/spidermonkeywasm2go/p7.Fn138
func Fn138(m *base.Module, l0 int32)

//go:linkname Fn139 github.com/goccy/spidermonkeywasm2go/p7.Fn139
func Fn139(m *base.Module)

//go:linkname Fn155 github.com/goccy/spidermonkeywasm2go/p0.Fn155
func Fn155(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn159 github.com/goccy/spidermonkeywasm2go/p7.Fn159
func Fn159(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn162 github.com/goccy/spidermonkeywasm2go/p6.Fn162
func Fn162(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn166 github.com/goccy/spidermonkeywasm2go/p4.Fn166
func Fn166(m *base.Module, l0 float64) float64

//go:linkname Fn168 github.com/goccy/spidermonkeywasm2go/p5.Fn168
func Fn168(m *base.Module, l0 float64) float64

//go:linkname Fn169 github.com/goccy/spidermonkeywasm2go/p5.Fn169
func Fn169(m *base.Module, l0 float64) float64

//go:linkname Fn170 github.com/goccy/spidermonkeywasm2go/p5.Fn170
func Fn170(m *base.Module, l0 float64) int32

//go:linkname Fn172 github.com/goccy/spidermonkeywasm2go/p4.Fn172
func Fn172(m *base.Module, l0 int32) float64

//go:linkname Fn173 github.com/goccy/spidermonkeywasm2go/p5.Fn173
func Fn173(m *base.Module, l0 int32) int64

//go:linkname Fn174 github.com/goccy/spidermonkeywasm2go/p5.Fn174
func Fn174(m *base.Module, l0 float64) float64

//go:linkname Fn176 github.com/goccy/spidermonkeywasm2go/p6.Fn176
func Fn176(m *base.Module, l0 float64, l1 float64, l2 float64) float64

//go:linkname Fn177 github.com/goccy/spidermonkeywasm2go/p6.Fn177
func Fn177(m *base.Module, l0 float64) float64

//go:linkname Fn178 github.com/goccy/spidermonkeywasm2go/p5.Fn178
func Fn178(m *base.Module, l0 float64) float64

//go:linkname Fn180 github.com/goccy/spidermonkeywasm2go/p6.Fn180
func Fn180(m *base.Module, l0 float64) float64

//go:linkname Fn181 github.com/goccy/spidermonkeywasm2go/p6.Fn181
func Fn181(m *base.Module, l0 float64) float64

//go:linkname Fn182 github.com/goccy/spidermonkeywasm2go/p6.Fn182
func Fn182(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn183 github.com/goccy/spidermonkeywasm2go/p6.Fn183
func Fn183(m *base.Module, l0 int32, l1 int32, l2 int32) int64

//go:linkname Fn185 github.com/goccy/spidermonkeywasm2go/p7.Fn185
func Fn185(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn198 github.com/goccy/spidermonkeywasm2go/p7.Fn198
func Fn198(m *base.Module, l0 int32)

//go:linkname Fn202 github.com/goccy/spidermonkeywasm2go/p7.Fn202
func Fn202(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn206 github.com/goccy/spidermonkeywasm2go/p6.Fn206
func Fn206(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn208 github.com/goccy/spidermonkeywasm2go/p6.Fn208
func Fn208(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn214 github.com/goccy/spidermonkeywasm2go/p6.Fn214
func Fn214(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn217 github.com/goccy/spidermonkeywasm2go/p7.Fn217
func Fn217(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn224 github.com/goccy/spidermonkeywasm2go/p4.Fn224
func Fn224(m *base.Module, l0 int64) int64

//go:linkname Fn228 github.com/goccy/spidermonkeywasm2go/p5.Fn228
func Fn228(m *base.Module, l0 int64) int64

//go:linkname Fn232 github.com/goccy/spidermonkeywasm2go/p5.Fn232
func Fn232(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn233 github.com/goccy/spidermonkeywasm2go/p6.Fn233
func Fn233(m *base.Module, l0 int32) int64

//go:linkname Fn234 github.com/goccy/spidermonkeywasm2go/p7.Fn234
func Fn234(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn235 github.com/goccy/spidermonkeywasm2go/p7.Fn235
func Fn235(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn236 github.com/goccy/spidermonkeywasm2go/p7.Fn236
func Fn236(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn242 github.com/goccy/spidermonkeywasm2go/p6.Fn242
func Fn242(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn245 github.com/goccy/spidermonkeywasm2go/p7.Fn245
func Fn245(m *base.Module, l0 int32) int32

//go:linkname Fn246 github.com/goccy/spidermonkeywasm2go/p7.Fn246
func Fn246(m *base.Module, l0 int32)

//go:linkname Fn252 github.com/goccy/spidermonkeywasm2go/p6.Fn252
func Fn252(m *base.Module, l0 int32) int32

//go:linkname Fn253 github.com/goccy/spidermonkeywasm2go/p7.Fn253
func Fn253(m *base.Module, l0 int32) int32

//go:linkname Fn254 github.com/goccy/spidermonkeywasm2go/p7.Fn254
func Fn254(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn258 github.com/goccy/spidermonkeywasm2go/p6.Fn258
func Fn258(m *base.Module, l0 int32) int32

//go:linkname Fn261 github.com/goccy/spidermonkeywasm2go/p7.Fn261
func Fn261(m *base.Module, l0 int32)

//go:linkname Fn266 github.com/goccy/spidermonkeywasm2go/p6.Fn266
func Fn266(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn270 github.com/goccy/spidermonkeywasm2go/p7.Fn270
func Fn270(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn271 github.com/goccy/spidermonkeywasm2go/p7.Fn271
func Fn271(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn273 github.com/goccy/spidermonkeywasm2go/p6.Fn273
func Fn273(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn278 github.com/goccy/spidermonkeywasm2go/p7.Fn278
func Fn278(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn279 github.com/goccy/spidermonkeywasm2go/p7.Fn279
func Fn279(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn280 github.com/goccy/spidermonkeywasm2go/p7.Fn280
func Fn280(m *base.Module, l0 int32) int32

//go:linkname Fn281 github.com/goccy/spidermonkeywasm2go/p7.Fn281
func Fn281(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn283 github.com/goccy/spidermonkeywasm2go/p7.Fn283
func Fn283(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn291 github.com/goccy/spidermonkeywasm2go/p6.Fn291
func Fn291(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn295 github.com/goccy/spidermonkeywasm2go/p5.Fn295
func Fn295(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn299 github.com/goccy/spidermonkeywasm2go/p6.Fn299
func Fn299(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32

//go:linkname Fn313 github.com/goccy/spidermonkeywasm2go/p7.Fn313
func Fn313(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn315 github.com/goccy/spidermonkeywasm2go/p7.Fn315
func Fn315(m *base.Module)

//go:linkname Fn340 github.com/goccy/spidermonkeywasm2go/p7.Fn340
func Fn340(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn341 github.com/goccy/spidermonkeywasm2go/p7.Fn341
func Fn341(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn342 github.com/goccy/spidermonkeywasm2go/p6.Fn342
func Fn342(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn345 github.com/goccy/spidermonkeywasm2go/p5.Fn345
func Fn345(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn348 github.com/goccy/spidermonkeywasm2go/p5.Fn348
func Fn348(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn351 github.com/goccy/spidermonkeywasm2go/p5.Fn351
func Fn351(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn354 github.com/goccy/spidermonkeywasm2go/p5.Fn354
func Fn354(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn357 github.com/goccy/spidermonkeywasm2go/p5.Fn357
func Fn357(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn360 github.com/goccy/spidermonkeywasm2go/p5.Fn360
func Fn360(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn363 github.com/goccy/spidermonkeywasm2go/p5.Fn363
func Fn363(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn366 github.com/goccy/spidermonkeywasm2go/p5.Fn366
func Fn366(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn367 github.com/goccy/spidermonkeywasm2go/p7.Fn367
func Fn367(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn373 github.com/goccy/spidermonkeywasm2go/p7.Fn373
func Fn373(m *base.Module, l0 int32) int32

//go:linkname Fn375 github.com/goccy/spidermonkeywasm2go/p7.Fn375
func Fn375(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn383 github.com/goccy/spidermonkeywasm2go/p5.Fn383
func Fn383(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn384 github.com/goccy/spidermonkeywasm2go/p6.Fn384
func Fn384(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn388 github.com/goccy/spidermonkeywasm2go/p7.Fn388
func Fn388(m *base.Module, l0 int32)

//go:linkname Fn389 github.com/goccy/spidermonkeywasm2go/p7.Fn389
func Fn389(m *base.Module, l0 int32)

//go:linkname Fn391 github.com/goccy/spidermonkeywasm2go/p7.Fn391
func Fn391(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn426 github.com/goccy/spidermonkeywasm2go/p7.Fn426
func Fn426(m *base.Module)

//go:linkname Fn456 github.com/goccy/spidermonkeywasm2go/p5.Fn456
func Fn456(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn474 github.com/goccy/spidermonkeywasm2go/p5.Fn474
func Fn474(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn476 github.com/goccy/spidermonkeywasm2go/p7.Fn476
func Fn476(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn517 github.com/goccy/spidermonkeywasm2go/p5.Fn517
func Fn517(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn524 github.com/goccy/spidermonkeywasm2go/p5.Fn524
func Fn524(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn526 github.com/goccy/spidermonkeywasm2go/p5.Fn526
func Fn526(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn532 github.com/goccy/spidermonkeywasm2go/p5.Fn532
func Fn532(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn565 github.com/goccy/spidermonkeywasm2go/p7.Fn565
func Fn565(m *base.Module, l0 int32)

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

//go:linkname Fn682 github.com/goccy/spidermonkeywasm2go/p6.Fn682
func Fn682(m *base.Module, l0 int32) int32

//go:linkname Fn684 github.com/goccy/spidermonkeywasm2go/p7.Fn684
func Fn684(m *base.Module, l0 int32)

//go:linkname Fn686 github.com/goccy/spidermonkeywasm2go/p7.Fn686
func Fn686(m *base.Module, l0 int32) int32

//go:linkname Fn713 github.com/goccy/spidermonkeywasm2go/p7.Fn713
func Fn713(m *base.Module, l0 float32) float32

//go:linkname Fn714 github.com/goccy/spidermonkeywasm2go/p7.Fn714
func Fn714(m *base.Module, l0 float64, l1 int32) float64

//go:linkname Fn719 github.com/goccy/spidermonkeywasm2go/p7.Fn719
func Fn719(m *base.Module, l0 float32) float32

//go:linkname Fn720 github.com/goccy/spidermonkeywasm2go/p6.Fn720
func Fn720(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn724 github.com/goccy/spidermonkeywasm2go/p7.Fn724
func Fn724(m *base.Module, l0 int32)

//go:linkname Fn725 github.com/goccy/spidermonkeywasm2go/p6.Fn725
func Fn725(m *base.Module, l0 int32) int32

//go:linkname Fn738 github.com/goccy/spidermonkeywasm2go/p7.Fn738
func Fn738(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn741 github.com/goccy/spidermonkeywasm2go/p7.Fn741
func Fn741(m *base.Module, l0 int32) int32

//go:linkname Fn751 github.com/goccy/spidermonkeywasm2go/p7.Fn751
func Fn751(m *base.Module, l0 int32)

//go:linkname Fn768 github.com/goccy/spidermonkeywasm2go/p7.Fn768
func Fn768(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn769 github.com/goccy/spidermonkeywasm2go/p7.Fn769
func Fn769(m *base.Module, l0 int32) int32

//go:linkname Fn770 github.com/goccy/spidermonkeywasm2go/p4.Fn770
func Fn770(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

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

//go:linkname Fn790 github.com/goccy/spidermonkeywasm2go/p7.Fn790
func Fn790(m *base.Module, l0 int32) int32

//go:linkname Fn791 github.com/goccy/spidermonkeywasm2go/p7.Fn791
func Fn791(m *base.Module, l0 int32)

//go:linkname Fn796 github.com/goccy/spidermonkeywasm2go/p7.Fn796
func Fn796(m *base.Module, l0 int32) int32

//go:linkname Fn797 github.com/goccy/spidermonkeywasm2go/p7.Fn797
func Fn797(m *base.Module, l0 int32) int32

//go:linkname Fn803 github.com/goccy/spidermonkeywasm2go/p7.Fn803
func Fn803(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn806 github.com/goccy/spidermonkeywasm2go/p6.Fn806
func Fn806(m *base.Module, l0 int32) int32

//go:linkname Fn808 github.com/goccy/spidermonkeywasm2go/p6.Fn808
func Fn808(m *base.Module, l0 int32) int32

//go:linkname Fn818 github.com/goccy/spidermonkeywasm2go/p7.Fn818
func Fn818(m *base.Module)

//go:linkname Fn819 github.com/goccy/spidermonkeywasm2go/p2.Fn819
func Fn819(m *base.Module, l0 int32, l1 int32, l2 int32) float64

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

//go:linkname Fn833 github.com/goccy/spidermonkeywasm2go/p0.Fn833
func Fn833(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn834 github.com/goccy/spidermonkeywasm2go/p5.Fn834
func Fn834(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn835 github.com/goccy/spidermonkeywasm2go/p6.Fn835
func Fn835(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

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

//go:linkname Fn855 github.com/goccy/spidermonkeywasm2go/p0.Fn855
func Fn855(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn856 github.com/goccy/spidermonkeywasm2go/p6.Fn856
func Fn856(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn858 github.com/goccy/spidermonkeywasm2go/p5.Fn858
func Fn858(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn859 github.com/goccy/spidermonkeywasm2go/p4.Fn859
func Fn859(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn864 github.com/goccy/spidermonkeywasm2go/p6.Fn864
func Fn864(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn868 github.com/goccy/spidermonkeywasm2go/p6.Fn868
func Fn868(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn869 github.com/goccy/spidermonkeywasm2go/p6.Fn869
func Fn869(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn871 github.com/goccy/spidermonkeywasm2go/p0.Fn871
func Fn871(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn872 github.com/goccy/spidermonkeywasm2go/p5.Fn872
func Fn872(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn873 github.com/goccy/spidermonkeywasm2go/p6.Fn873
func Fn873(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn875 github.com/goccy/spidermonkeywasm2go/p0.Fn875
func Fn875(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn876 github.com/goccy/spidermonkeywasm2go/p0.Fn876
func Fn876(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn877 github.com/goccy/spidermonkeywasm2go/p0.Fn877
func Fn877(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn885 github.com/goccy/spidermonkeywasm2go/p6.Fn885
func Fn885(m *base.Module, l0 int32) int32

//go:linkname Fn887 github.com/goccy/spidermonkeywasm2go/p0.Fn887
func Fn887(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn888 github.com/goccy/spidermonkeywasm2go/p6.Fn888
func Fn888(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn890 github.com/goccy/spidermonkeywasm2go/p7.Fn890
func Fn890(m *base.Module, l0 int32) int32

//go:linkname Fn892 github.com/goccy/spidermonkeywasm2go/p5.Fn892
func Fn892(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn904 github.com/goccy/spidermonkeywasm2go/p6.Fn904
func Fn904(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn916 github.com/goccy/spidermonkeywasm2go/p6.Fn916
func Fn916(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn932 github.com/goccy/spidermonkeywasm2go/p7.Fn932
func Fn932(m *base.Module, l0 int32) int32

//go:linkname Fn933 github.com/goccy/spidermonkeywasm2go/p5.Fn933
func Fn933(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1001 github.com/goccy/spidermonkeywasm2go/p7.Fn1001
func Fn1001(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1034 github.com/goccy/spidermonkeywasm2go/p6.Fn1034
func Fn1034(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64, l5 int32, l6 int32) int32

//go:linkname Fn1035 github.com/goccy/spidermonkeywasm2go/p5.Fn1035
func Fn1035(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1037 github.com/goccy/spidermonkeywasm2go/p6.Fn1037
func Fn1037(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1041 github.com/goccy/spidermonkeywasm2go/p6.Fn1041
func Fn1041(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1042 github.com/goccy/spidermonkeywasm2go/p6.Fn1042
func Fn1042(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1080 github.com/goccy/spidermonkeywasm2go/p0.Fn1080
func Fn1080(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1082 github.com/goccy/spidermonkeywasm2go/p6.Fn1082
func Fn1082(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1084 github.com/goccy/spidermonkeywasm2go/p0.Fn1084
func Fn1084(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1086 github.com/goccy/spidermonkeywasm2go/p0.Fn1086
func Fn1086(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1087 github.com/goccy/spidermonkeywasm2go/p0.Fn1087
func Fn1087(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1088 github.com/goccy/spidermonkeywasm2go/p0.Fn1088
func Fn1088(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1089 github.com/goccy/spidermonkeywasm2go/p0.Fn1089
func Fn1089(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1091 github.com/goccy/spidermonkeywasm2go/p0.Fn1091
func Fn1091(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1092 github.com/goccy/spidermonkeywasm2go/p6.Fn1092
func Fn1092(m *base.Module, l0 int32) int32

//go:linkname Fn1100 github.com/goccy/spidermonkeywasm2go/p0.Fn1100
func Fn1100(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1108 github.com/goccy/spidermonkeywasm2go/p0.Fn1108
func Fn1108(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1109 github.com/goccy/spidermonkeywasm2go/p0.Fn1109
func Fn1109(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1117 github.com/goccy/spidermonkeywasm2go/p0.Fn1117
func Fn1117(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1119 github.com/goccy/spidermonkeywasm2go/p0.Fn1119
func Fn1119(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1120 github.com/goccy/spidermonkeywasm2go/p0.Fn1120
func Fn1120(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1122 github.com/goccy/spidermonkeywasm2go/p0.Fn1122
func Fn1122(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1128 github.com/goccy/spidermonkeywasm2go/p5.Fn1128
func Fn1128(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1129 github.com/goccy/spidermonkeywasm2go/p6.Fn1129
func Fn1129(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1131 github.com/goccy/spidermonkeywasm2go/p0.Fn1131
func Fn1131(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1132 github.com/goccy/spidermonkeywasm2go/p0.Fn1132
func Fn1132(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1134 github.com/goccy/spidermonkeywasm2go/p0.Fn1134
func Fn1134(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1136 github.com/goccy/spidermonkeywasm2go/p6.Fn1136
func Fn1136(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1137 github.com/goccy/spidermonkeywasm2go/p0.Fn1137
func Fn1137(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1162 github.com/goccy/spidermonkeywasm2go/p6.Fn1162
func Fn1162(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1163 github.com/goccy/spidermonkeywasm2go/p5.Fn1163
func Fn1163(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1164 github.com/goccy/spidermonkeywasm2go/p6.Fn1164
func Fn1164(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1177 github.com/goccy/spidermonkeywasm2go/p4.Fn1177
func Fn1177(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn1270 github.com/goccy/spidermonkeywasm2go/p0.Fn1270
func Fn1270(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1273 github.com/goccy/spidermonkeywasm2go/p6.Fn1273
func Fn1273(m *base.Module, l0 int32) int32

//go:linkname Fn1274 github.com/goccy/spidermonkeywasm2go/p6.Fn1274
func Fn1274(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1281 github.com/goccy/spidermonkeywasm2go/p4.Fn1281
func Fn1281(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1282 github.com/goccy/spidermonkeywasm2go/p6.Fn1282
func Fn1282(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1297 github.com/goccy/spidermonkeywasm2go/p6.Fn1297
func Fn1297(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1298 github.com/goccy/spidermonkeywasm2go/p0.Fn1298
func Fn1298(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1312 github.com/goccy/spidermonkeywasm2go/p5.Fn1312
func Fn1312(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1314 github.com/goccy/spidermonkeywasm2go/p6.Fn1314
func Fn1314(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1315 github.com/goccy/spidermonkeywasm2go/p6.Fn1315
func Fn1315(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1353 github.com/goccy/spidermonkeywasm2go/p7.Fn1353
func Fn1353(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1354 github.com/goccy/spidermonkeywasm2go/p6.Fn1354
func Fn1354(m *base.Module, l0 int32) int32

//go:linkname Fn1355 github.com/goccy/spidermonkeywasm2go/p4.Fn1355
func Fn1355(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1356 github.com/goccy/spidermonkeywasm2go/p6.Fn1356
func Fn1356(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1404 github.com/goccy/spidermonkeywasm2go/p0.Fn1404
func Fn1404(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1407 github.com/goccy/spidermonkeywasm2go/p0.Fn1407
func Fn1407(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1408 github.com/goccy/spidermonkeywasm2go/p7.Fn1408
func Fn1408(m *base.Module, l0 int32) int32

//go:linkname Fn1434 github.com/goccy/spidermonkeywasm2go/p6.Fn1434
func Fn1434(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn1436 github.com/goccy/spidermonkeywasm2go/p0.Fn1436
func Fn1436(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1443 github.com/goccy/spidermonkeywasm2go/p0.Fn1443
func Fn1443(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1444 github.com/goccy/spidermonkeywasm2go/p6.Fn1444
func Fn1444(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1445 github.com/goccy/spidermonkeywasm2go/p5.Fn1445
func Fn1445(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1447 github.com/goccy/spidermonkeywasm2go/p7.Fn1447
func Fn1447(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1449 github.com/goccy/spidermonkeywasm2go/p0.Fn1449
func Fn1449(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1451 github.com/goccy/spidermonkeywasm2go/p5.Fn1451
func Fn1451(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1457 github.com/goccy/spidermonkeywasm2go/p6.Fn1457
func Fn1457(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn1460 github.com/goccy/spidermonkeywasm2go/p5.Fn1460
func Fn1460(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1461 github.com/goccy/spidermonkeywasm2go/p6.Fn1461
func Fn1461(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1462 github.com/goccy/spidermonkeywasm2go/p7.Fn1462
func Fn1462(m *base.Module, l0 int32)

//go:linkname Fn1463 github.com/goccy/spidermonkeywasm2go/p6.Fn1463
func Fn1463(m *base.Module, l0 int32)

//go:linkname Fn1464 github.com/goccy/spidermonkeywasm2go/p5.Fn1464
func Fn1464(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1465 github.com/goccy/spidermonkeywasm2go/p6.Fn1465
func Fn1465(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1470 github.com/goccy/spidermonkeywasm2go/p7.Fn1470
func Fn1470(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1472 github.com/goccy/spidermonkeywasm2go/p0.Fn1472
func Fn1472(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1473 github.com/goccy/spidermonkeywasm2go/p0.Fn1473
func Fn1473(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1474 github.com/goccy/spidermonkeywasm2go/p1.Fn1474
func Fn1474(m *base.Module, l0 int32, l1 int32) int32

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

//go:linkname Fn1493 github.com/goccy/spidermonkeywasm2go/p6.Fn1493
func Fn1493(m *base.Module, l0 int32) int32

//go:linkname Fn1495 github.com/goccy/spidermonkeywasm2go/p0.Fn1495
func Fn1495(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1497 github.com/goccy/spidermonkeywasm2go/p7.Fn1497
func Fn1497(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1498 github.com/goccy/spidermonkeywasm2go/p0.Fn1498
func Fn1498(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1502 github.com/goccy/spidermonkeywasm2go/p7.Fn1502
func Fn1502(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1504 github.com/goccy/spidermonkeywasm2go/p7.Fn1504
func Fn1504(m *base.Module, l0 int32)

//go:linkname Fn1508 github.com/goccy/spidermonkeywasm2go/p7.Fn1508
func Fn1508(m *base.Module, l0 int32)

//go:linkname Fn1509 github.com/goccy/spidermonkeywasm2go/p7.Fn1509
func Fn1509(m *base.Module, l0 int32) int32

//go:linkname Fn1510 github.com/goccy/spidermonkeywasm2go/p5.Fn1510
func Fn1510(m *base.Module, l0 float64, l1 float64, l2 float64) float64

//go:linkname Fn1516 github.com/goccy/spidermonkeywasm2go/p6.Fn1516
func Fn1516(m *base.Module, l0 int32, l1 float64) float64

//go:linkname Fn1522 github.com/goccy/spidermonkeywasm2go/p6.Fn1522
func Fn1522(m *base.Module, l0 int32, l1 int32, l2 float64) int32

//go:linkname Fn1523 github.com/goccy/spidermonkeywasm2go/p1.Fn1523
func Fn1523(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1524 github.com/goccy/spidermonkeywasm2go/p7.Fn1524
func Fn1524(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1525 github.com/goccy/spidermonkeywasm2go/p6.Fn1525
func Fn1525(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1582 github.com/goccy/spidermonkeywasm2go/p4.Fn1582
func Fn1582(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1595 github.com/goccy/spidermonkeywasm2go/p6.Fn1595
func Fn1595(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1597 github.com/goccy/spidermonkeywasm2go/p7.Fn1597
func Fn1597(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1619 github.com/goccy/spidermonkeywasm2go/p4.Fn1619
func Fn1619(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1620 github.com/goccy/spidermonkeywasm2go/p0.Fn1620
func Fn1620(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1621 github.com/goccy/spidermonkeywasm2go/p0.Fn1621
func Fn1621(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1622 github.com/goccy/spidermonkeywasm2go/p6.Fn1622
func Fn1622(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1623 github.com/goccy/spidermonkeywasm2go/p6.Fn1623
func Fn1623(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1627 github.com/goccy/spidermonkeywasm2go/p7.Fn1627
func Fn1627(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1632 github.com/goccy/spidermonkeywasm2go/p0.Fn1632
func Fn1632(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1642 github.com/goccy/spidermonkeywasm2go/p6.Fn1642
func Fn1642(m *base.Module, l0 int32, l1 float64, l2 int32) int32

//go:linkname Fn1643 github.com/goccy/spidermonkeywasm2go/p7.Fn1643
func Fn1643(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1660 github.com/goccy/spidermonkeywasm2go/p0.Fn1660
func Fn1660(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1662 github.com/goccy/spidermonkeywasm2go/p0.Fn1662
func Fn1662(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1749 github.com/goccy/spidermonkeywasm2go/p7.Fn1749
func Fn1749(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1760 github.com/goccy/spidermonkeywasm2go/p7.Fn1760
func Fn1760(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1762 github.com/goccy/spidermonkeywasm2go/p7.Fn1762
func Fn1762(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1771 github.com/goccy/spidermonkeywasm2go/p6.Fn1771
func Fn1771(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1772 github.com/goccy/spidermonkeywasm2go/p7.Fn1772
func Fn1772(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1778 github.com/goccy/spidermonkeywasm2go/p6.Fn1778
func Fn1778(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1781 github.com/goccy/spidermonkeywasm2go/p6.Fn1781
func Fn1781(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1782 github.com/goccy/spidermonkeywasm2go/p7.Fn1782
func Fn1782(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1806 github.com/goccy/spidermonkeywasm2go/p6.Fn1806
func Fn1806(m *base.Module, l0 int32) int32

//go:linkname Fn1807 github.com/goccy/spidermonkeywasm2go/p6.Fn1807
func Fn1807(m *base.Module, l0 int32) int32

//go:linkname Fn1808 github.com/goccy/spidermonkeywasm2go/p6.Fn1808
func Fn1808(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1815 github.com/goccy/spidermonkeywasm2go/p5.Fn1815
func Fn1815(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1818 github.com/goccy/spidermonkeywasm2go/p5.Fn1818
func Fn1818(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1819 github.com/goccy/spidermonkeywasm2go/p0.Fn1819
func Fn1819(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1846 github.com/goccy/spidermonkeywasm2go/p0.Fn1846
func Fn1846(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1847 github.com/goccy/spidermonkeywasm2go/p7.Fn1847
func Fn1847(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1848 github.com/goccy/spidermonkeywasm2go/p5.Fn1848
func Fn1848(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1849 github.com/goccy/spidermonkeywasm2go/p0.Fn1849
func Fn1849(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1854 github.com/goccy/spidermonkeywasm2go/p5.Fn1854
func Fn1854(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn1866 github.com/goccy/spidermonkeywasm2go/p6.Fn1866
func Fn1866(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1869 github.com/goccy/spidermonkeywasm2go/p6.Fn1869
func Fn1869(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1872 github.com/goccy/spidermonkeywasm2go/p6.Fn1872
func Fn1872(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1873 github.com/goccy/spidermonkeywasm2go/p7.Fn1873
func Fn1873(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1887 github.com/goccy/spidermonkeywasm2go/p7.Fn1887
func Fn1887(m *base.Module, l0 int32) int32

//go:linkname Fn1888 github.com/goccy/spidermonkeywasm2go/p6.Fn1888
func Fn1888(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1889 github.com/goccy/spidermonkeywasm2go/p7.Fn1889
func Fn1889(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1890 github.com/goccy/spidermonkeywasm2go/p7.Fn1890
func Fn1890(m *base.Module, l0 int32) int32

//go:linkname Fn1891 github.com/goccy/spidermonkeywasm2go/p7.Fn1891
func Fn1891(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1892 github.com/goccy/spidermonkeywasm2go/p5.Fn1892
func Fn1892(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1893 github.com/goccy/spidermonkeywasm2go/p6.Fn1893
func Fn1893(m *base.Module, l0 int32) int32

//go:linkname Fn1912 github.com/goccy/spidermonkeywasm2go/p4.Fn1912
func Fn1912(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1953 github.com/goccy/spidermonkeywasm2go/p0.Fn1953
func Fn1953(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn1968 github.com/goccy/spidermonkeywasm2go/p7.Fn1968
func Fn1968(m *base.Module, l0 int32) int64

//go:linkname Fn1969 github.com/goccy/spidermonkeywasm2go/p7.Fn1969
func Fn1969(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1975 github.com/goccy/spidermonkeywasm2go/p0.Fn1975
func Fn1975(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1977 github.com/goccy/spidermonkeywasm2go/p0.Fn1977
func Fn1977(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1993 github.com/goccy/spidermonkeywasm2go/p0.Fn1993
func Fn1993(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1998 github.com/goccy/spidermonkeywasm2go/p5.Fn1998
func Fn1998(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2024 github.com/goccy/spidermonkeywasm2go/p7.Fn2024
func Fn2024(m *base.Module, l0 int32) int32

//go:linkname Fn2031 github.com/goccy/spidermonkeywasm2go/p7.Fn2031
func Fn2031(m *base.Module, l0 int32) int32

//go:linkname Fn2047 github.com/goccy/spidermonkeywasm2go/p7.Fn2047
func Fn2047(m *base.Module, l0 int32) int32

//go:linkname Fn2048 github.com/goccy/spidermonkeywasm2go/p6.Fn2048
func Fn2048(m *base.Module, l0 int32)

//go:linkname Fn2049 github.com/goccy/spidermonkeywasm2go/p6.Fn2049
func Fn2049(m *base.Module, l0 int32)

//go:linkname Fn2052 github.com/goccy/spidermonkeywasm2go/p7.Fn2052
func Fn2052(m *base.Module, l0 int32) int32

//go:linkname Fn2053 github.com/goccy/spidermonkeywasm2go/p0.Fn2053
func Fn2053(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2054 github.com/goccy/spidermonkeywasm2go/p0.Fn2054
func Fn2054(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2055 github.com/goccy/spidermonkeywasm2go/p7.Fn2055
func Fn2055(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2056 github.com/goccy/spidermonkeywasm2go/p0.Fn2056
func Fn2056(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2058 github.com/goccy/spidermonkeywasm2go/p0.Fn2058
func Fn2058(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2065 github.com/goccy/spidermonkeywasm2go/p0.Fn2065
func Fn2065(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2073 github.com/goccy/spidermonkeywasm2go/p6.Fn2073
func Fn2073(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2082 github.com/goccy/spidermonkeywasm2go/p7.Fn2082
func Fn2082(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2084 github.com/goccy/spidermonkeywasm2go/p6.Fn2084
func Fn2084(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2089 github.com/goccy/spidermonkeywasm2go/p6.Fn2089
func Fn2089(m *base.Module, l0 int32)

//go:linkname Fn2091 github.com/goccy/spidermonkeywasm2go/p4.Fn2091
func Fn2091(m *base.Module, l0 int32)

//go:linkname Fn2092 github.com/goccy/spidermonkeywasm2go/p7.Fn2092
func Fn2092(m *base.Module, l0 int32) int32

//go:linkname Fn2095 github.com/goccy/spidermonkeywasm2go/p6.Fn2095
func Fn2095(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn2098 github.com/goccy/spidermonkeywasm2go/p7.Fn2098
func Fn2098(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn2099 github.com/goccy/spidermonkeywasm2go/p4.Fn2099
func Fn2099(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn2102 github.com/goccy/spidermonkeywasm2go/p7.Fn2102
func Fn2102(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2113 github.com/goccy/spidermonkeywasm2go/p5.Fn2113
func Fn2113(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2138 github.com/goccy/spidermonkeywasm2go/p6.Fn2138
func Fn2138(m *base.Module, l0 int32)

//go:linkname Fn2142 github.com/goccy/spidermonkeywasm2go/p0.Fn2142
func Fn2142(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2182 github.com/goccy/spidermonkeywasm2go/p6.Fn2182
func Fn2182(m *base.Module, l0 int32) int32

//go:linkname Fn2187 github.com/goccy/spidermonkeywasm2go/p6.Fn2187
func Fn2187(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2189 github.com/goccy/spidermonkeywasm2go/p6.Fn2189
func Fn2189(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2219 github.com/goccy/spidermonkeywasm2go/p6.Fn2219
func Fn2219(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2220 github.com/goccy/spidermonkeywasm2go/p0.Fn2220
func Fn2220(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2224 github.com/goccy/spidermonkeywasm2go/p6.Fn2224
func Fn2224(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2245 github.com/goccy/spidermonkeywasm2go/p0.Fn2245
func Fn2245(m *base.Module, l0 int32) int32

//go:linkname Fn2246 github.com/goccy/spidermonkeywasm2go/p6.Fn2246
func Fn2246(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2250 github.com/goccy/spidermonkeywasm2go/p0.Fn2250
func Fn2250(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2251 github.com/goccy/spidermonkeywasm2go/p0.Fn2251
func Fn2251(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2252 github.com/goccy/spidermonkeywasm2go/p0.Fn2252
func Fn2252(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2256 github.com/goccy/spidermonkeywasm2go/p6.Fn2256
func Fn2256(m *base.Module, l0 int32)

//go:linkname Fn2267 github.com/goccy/spidermonkeywasm2go/p6.Fn2267
func Fn2267(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2268 github.com/goccy/spidermonkeywasm2go/p5.Fn2268
func Fn2268(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2283 github.com/goccy/spidermonkeywasm2go/p0.Fn2283
func Fn2283(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2291 github.com/goccy/spidermonkeywasm2go/p0.Fn2291
func Fn2291(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2294 github.com/goccy/spidermonkeywasm2go/p6.Fn2294
func Fn2294(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2304 github.com/goccy/spidermonkeywasm2go/p7.Fn2304
func Fn2304(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn2308 github.com/goccy/spidermonkeywasm2go/p0.Fn2308
func Fn2308(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2327 github.com/goccy/spidermonkeywasm2go/p0.Fn2327
func Fn2327(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2328 github.com/goccy/spidermonkeywasm2go/p7.Fn2328
func Fn2328(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2330 github.com/goccy/spidermonkeywasm2go/p5.Fn2330
func Fn2330(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2333 github.com/goccy/spidermonkeywasm2go/p7.Fn2333
func Fn2333(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2334 github.com/goccy/spidermonkeywasm2go/p0.Fn2334
func Fn2334(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2403 github.com/goccy/spidermonkeywasm2go/p6.Fn2403
func Fn2403(m *base.Module, l0 int32)

//go:linkname Fn2427 github.com/goccy/spidermonkeywasm2go/p5.Fn2427
func Fn2427(m *base.Module, l0 int32)

//go:linkname Fn2429 github.com/goccy/spidermonkeywasm2go/p5.Fn2429
func Fn2429(m *base.Module, l0 int32) int32

//go:linkname Fn2432 github.com/goccy/spidermonkeywasm2go/p6.Fn2432
func Fn2432(m *base.Module, l0 int32) int32

//go:linkname Fn2437 github.com/goccy/spidermonkeywasm2go/p7.Fn2437
func Fn2437(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2438 github.com/goccy/spidermonkeywasm2go/p0.Fn2438
func Fn2438(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2439 github.com/goccy/spidermonkeywasm2go/p6.Fn2439
func Fn2439(m *base.Module, l0 int32) int32

//go:linkname Fn2441 github.com/goccy/spidermonkeywasm2go/p6.Fn2441
func Fn2441(m *base.Module, l0 int32)

//go:linkname Fn2447 github.com/goccy/spidermonkeywasm2go/p0.Fn2447
func Fn2447(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2451 github.com/goccy/spidermonkeywasm2go/p5.Fn2451
func Fn2451(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2478 github.com/goccy/spidermonkeywasm2go/p4.Fn2478
func Fn2478(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2484 github.com/goccy/spidermonkeywasm2go/p4.Fn2484
func Fn2484(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2486 github.com/goccy/spidermonkeywasm2go/p4.Fn2486
func Fn2486(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2497 github.com/goccy/spidermonkeywasm2go/p2.Fn2497
func Fn2497(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2503 github.com/goccy/spidermonkeywasm2go/p5.Fn2503
func Fn2503(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2504 github.com/goccy/spidermonkeywasm2go/p6.Fn2504
func Fn2504(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2523 github.com/goccy/spidermonkeywasm2go/p0.Fn2523
func Fn2523(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2527 github.com/goccy/spidermonkeywasm2go/p6.Fn2527
func Fn2527(m *base.Module, l0 int32)

//go:linkname Fn2528 github.com/goccy/spidermonkeywasm2go/p0.Fn2528
func Fn2528(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2529 github.com/goccy/spidermonkeywasm2go/p0.Fn2529
func Fn2529(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2530 github.com/goccy/spidermonkeywasm2go/p0.Fn2530
func Fn2530(m *base.Module, l0 int32) int32

//go:linkname Fn2536 github.com/goccy/spidermonkeywasm2go/p7.Fn2536
func Fn2536(m *base.Module, l0 int32)

//go:linkname Fn2545 github.com/goccy/spidermonkeywasm2go/p0.Fn2545
func Fn2545(m *base.Module, l0 int32) int32

//go:linkname Fn2546 github.com/goccy/spidermonkeywasm2go/p0.Fn2546
func Fn2546(m *base.Module, l0 int32)

//go:linkname Fn2551 github.com/goccy/spidermonkeywasm2go/p0.Fn2551
func Fn2551(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2552 github.com/goccy/spidermonkeywasm2go/p0.Fn2552
func Fn2552(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2560 github.com/goccy/spidermonkeywasm2go/p5.Fn2560
func Fn2560(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2566 github.com/goccy/spidermonkeywasm2go/p0.Fn2566
func Fn2566(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2567 github.com/goccy/spidermonkeywasm2go/p0.Fn2567
func Fn2567(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2571 github.com/goccy/spidermonkeywasm2go/p0.Fn2571
func Fn2571(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn2573 github.com/goccy/spidermonkeywasm2go/p0.Fn2573
func Fn2573(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2577 github.com/goccy/spidermonkeywasm2go/p0.Fn2577
func Fn2577(m *base.Module, l0 int32)

//go:linkname Fn2578 github.com/goccy/spidermonkeywasm2go/p0.Fn2578
func Fn2578(m *base.Module, l0 int32)

//go:linkname Fn2581 github.com/goccy/spidermonkeywasm2go/p0.Fn2581
func Fn2581(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2582 github.com/goccy/spidermonkeywasm2go/p0.Fn2582
func Fn2582(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn2583 github.com/goccy/spidermonkeywasm2go/p6.Fn2583
func Fn2583(m *base.Module, l0 int32) int32

//go:linkname Fn2584 github.com/goccy/spidermonkeywasm2go/p7.Fn2584
func Fn2584(m *base.Module, l0 int32)

//go:linkname Fn2589 github.com/goccy/spidermonkeywasm2go/p0.Fn2589
func Fn2589(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2592 github.com/goccy/spidermonkeywasm2go/p4.Fn2592
func Fn2592(m *base.Module, l0 int32)

//go:linkname Fn2595 github.com/goccy/spidermonkeywasm2go/p0.Fn2595
func Fn2595(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2600 github.com/goccy/spidermonkeywasm2go/p0.Fn2600
func Fn2600(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2609 github.com/goccy/spidermonkeywasm2go/p0.Fn2609
func Fn2609(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2611 github.com/goccy/spidermonkeywasm2go/p0.Fn2611
func Fn2611(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2620 github.com/goccy/spidermonkeywasm2go/p0.Fn2620
func Fn2620(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn2631 github.com/goccy/spidermonkeywasm2go/p6.Fn2631
func Fn2631(m *base.Module, l0 int32)

//go:linkname Fn2632 github.com/goccy/spidermonkeywasm2go/p0.Fn2632
func Fn2632(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2634 github.com/goccy/spidermonkeywasm2go/p0.Fn2634
func Fn2634(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2641 github.com/goccy/spidermonkeywasm2go/p0.Fn2641
func Fn2641(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2679 github.com/goccy/spidermonkeywasm2go/p6.Fn2679
func Fn2679(m *base.Module, l0 int32) int32

//go:linkname Fn2698 github.com/goccy/spidermonkeywasm2go/p0.Fn2698
func Fn2698(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2707 github.com/goccy/spidermonkeywasm2go/p7.Fn2707
func Fn2707(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2709 github.com/goccy/spidermonkeywasm2go/p0.Fn2709
func Fn2709(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2710 github.com/goccy/spidermonkeywasm2go/p7.Fn2710
func Fn2710(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2711 github.com/goccy/spidermonkeywasm2go/p7.Fn2711
func Fn2711(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2712 github.com/goccy/spidermonkeywasm2go/p4.Fn2712
func Fn2712(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2715 github.com/goccy/spidermonkeywasm2go/p0.Fn2715
func Fn2715(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2717 github.com/goccy/spidermonkeywasm2go/p7.Fn2717
func Fn2717(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2720 github.com/goccy/spidermonkeywasm2go/p0.Fn2720
func Fn2720(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2721 github.com/goccy/spidermonkeywasm2go/p0.Fn2721
func Fn2721(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2722 github.com/goccy/spidermonkeywasm2go/p0.Fn2722
func Fn2722(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2724 github.com/goccy/spidermonkeywasm2go/p0.Fn2724
func Fn2724(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2728 github.com/goccy/spidermonkeywasm2go/p0.Fn2728
func Fn2728(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2730 github.com/goccy/spidermonkeywasm2go/p0.Fn2730
func Fn2730(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2731 github.com/goccy/spidermonkeywasm2go/p4.Fn2731
func Fn2731(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2732 github.com/goccy/spidermonkeywasm2go/p4.Fn2732
func Fn2732(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2733 github.com/goccy/spidermonkeywasm2go/p6.Fn2733
func Fn2733(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2734 github.com/goccy/spidermonkeywasm2go/p7.Fn2734
func Fn2734(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2736 github.com/goccy/spidermonkeywasm2go/p0.Fn2736
func Fn2736(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2741 github.com/goccy/spidermonkeywasm2go/p0.Fn2741
func Fn2741(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2742 github.com/goccy/spidermonkeywasm2go/p0.Fn2742
func Fn2742(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2743 github.com/goccy/spidermonkeywasm2go/p0.Fn2743
func Fn2743(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2745 github.com/goccy/spidermonkeywasm2go/p0.Fn2745
func Fn2745(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2752 github.com/goccy/spidermonkeywasm2go/p5.Fn2752
func Fn2752(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2753 github.com/goccy/spidermonkeywasm2go/p4.Fn2753
func Fn2753(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2755 github.com/goccy/spidermonkeywasm2go/p0.Fn2755
func Fn2755(m *base.Module, l0 int32, l1 int32, l2 int32) int32

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

//go:linkname Fn2789 github.com/goccy/spidermonkeywasm2go/p6.Fn2789
func Fn2789(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2797 github.com/goccy/spidermonkeywasm2go/p4.Fn2797
func Fn2797(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2802 github.com/goccy/spidermonkeywasm2go/p6.Fn2802
func Fn2802(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn2812 github.com/goccy/spidermonkeywasm2go/p6.Fn2812
func Fn2812(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2828 github.com/goccy/spidermonkeywasm2go/p2.Fn2828
func Fn2828(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2829 github.com/goccy/spidermonkeywasm2go/p5.Fn2829
func Fn2829(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2831 github.com/goccy/spidermonkeywasm2go/p0.Fn2831
func Fn2831(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2832 github.com/goccy/spidermonkeywasm2go/p0.Fn2832
func Fn2832(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2833 github.com/goccy/spidermonkeywasm2go/p5.Fn2833
func Fn2833(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2834 github.com/goccy/spidermonkeywasm2go/p6.Fn2834
func Fn2834(m *base.Module, l0 int32) int32

//go:linkname Fn2840 github.com/goccy/spidermonkeywasm2go/p2.Fn2840
func Fn2840(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2843 github.com/goccy/spidermonkeywasm2go/p5.Fn2843
func Fn2843(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2844 github.com/goccy/spidermonkeywasm2go/p6.Fn2844
func Fn2844(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2847 github.com/goccy/spidermonkeywasm2go/p6.Fn2847
func Fn2847(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2848 github.com/goccy/spidermonkeywasm2go/p5.Fn2848
func Fn2848(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2851 github.com/goccy/spidermonkeywasm2go/p0.Fn2851
func Fn2851(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2852 github.com/goccy/spidermonkeywasm2go/p6.Fn2852
func Fn2852(m *base.Module, l0 int32)

//go:linkname Fn2855 github.com/goccy/spidermonkeywasm2go/p5.Fn2855
func Fn2855(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2856 github.com/goccy/spidermonkeywasm2go/p5.Fn2856
func Fn2856(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2870 github.com/goccy/spidermonkeywasm2go/p0.Fn2870
func Fn2870(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2875 github.com/goccy/spidermonkeywasm2go/p4.Fn2875
func Fn2875(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2876 github.com/goccy/spidermonkeywasm2go/p0.Fn2876
func Fn2876(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2880 github.com/goccy/spidermonkeywasm2go/p4.Fn2880
func Fn2880(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2882 github.com/goccy/spidermonkeywasm2go/p0.Fn2882
func Fn2882(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2886 github.com/goccy/spidermonkeywasm2go/p0.Fn2886
func Fn2886(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2889 github.com/goccy/spidermonkeywasm2go/p4.Fn2889
func Fn2889(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2894 github.com/goccy/spidermonkeywasm2go/p6.Fn2894
func Fn2894(m *base.Module, l0 int32) int32

//go:linkname Fn2895 github.com/goccy/spidermonkeywasm2go/p5.Fn2895
func Fn2895(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2900 github.com/goccy/spidermonkeywasm2go/p6.Fn2900
func Fn2900(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2911 github.com/goccy/spidermonkeywasm2go/p6.Fn2911
func Fn2911(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2913 github.com/goccy/spidermonkeywasm2go/p6.Fn2913
func Fn2913(m *base.Module, l0 int32) int64

//go:linkname Fn2918 github.com/goccy/spidermonkeywasm2go/p4.Fn2918
func Fn2918(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn2929 github.com/goccy/spidermonkeywasm2go/p0.Fn2929
func Fn2929(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2934 github.com/goccy/spidermonkeywasm2go/p4.Fn2934
func Fn2934(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2939 github.com/goccy/spidermonkeywasm2go/p7.Fn2939
func Fn2939(m *base.Module, l0 int32)

//go:linkname Fn2942 github.com/goccy/spidermonkeywasm2go/p7.Fn2942
func Fn2942(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2947 github.com/goccy/spidermonkeywasm2go/p0.Fn2947
func Fn2947(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2974 github.com/goccy/spidermonkeywasm2go/p6.Fn2974
func Fn2974(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2976 github.com/goccy/spidermonkeywasm2go/p6.Fn2976
func Fn2976(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2993 github.com/goccy/spidermonkeywasm2go/p6.Fn2993
func Fn2993(m *base.Module, l0 int32)

//go:linkname Fn2995 github.com/goccy/spidermonkeywasm2go/p5.Fn2995
func Fn2995(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2998 github.com/goccy/spidermonkeywasm2go/p6.Fn2998
func Fn2998(m *base.Module, l0 int32) int32

//go:linkname Fn2999 github.com/goccy/spidermonkeywasm2go/p6.Fn2999
func Fn2999(m *base.Module, l0 int32) int32

//go:linkname Fn3034 github.com/goccy/spidermonkeywasm2go/p0.Fn3034
func Fn3034(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3035 github.com/goccy/spidermonkeywasm2go/p0.Fn3035
func Fn3035(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3037 github.com/goccy/spidermonkeywasm2go/p6.Fn3037
func Fn3037(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3044 github.com/goccy/spidermonkeywasm2go/p6.Fn3044
func Fn3044(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3048 github.com/goccy/spidermonkeywasm2go/p6.Fn3048
func Fn3048(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3051 github.com/goccy/spidermonkeywasm2go/p6.Fn3051
func Fn3051(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3094 github.com/goccy/spidermonkeywasm2go/p6.Fn3094
func Fn3094(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3095 github.com/goccy/spidermonkeywasm2go/p4.Fn3095
func Fn3095(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3096 github.com/goccy/spidermonkeywasm2go/p5.Fn3096
func Fn3096(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3148 github.com/goccy/spidermonkeywasm2go/p6.Fn3148
func Fn3148(m *base.Module, l0 float64) float64

//go:linkname Fn3149 github.com/goccy/spidermonkeywasm2go/p2.Fn3149
func Fn3149(m *base.Module, l0 int32) int32

//go:linkname Fn3151 github.com/goccy/spidermonkeywasm2go/p2.Fn3151
func Fn3151(m *base.Module, l0 int32)

//go:linkname Fn3153 github.com/goccy/spidermonkeywasm2go/p0.Fn3153
func Fn3153(m *base.Module, l0 int32) int32

//go:linkname Fn3155 github.com/goccy/spidermonkeywasm2go/p6.Fn3155
func Fn3155(m *base.Module, l0 int32) int32

//go:linkname Fn3156 github.com/goccy/spidermonkeywasm2go/p7.Fn3156
func Fn3156(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3157 github.com/goccy/spidermonkeywasm2go/p7.Fn3157
func Fn3157(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3158 github.com/goccy/spidermonkeywasm2go/p6.Fn3158
func Fn3158(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3166 github.com/goccy/spidermonkeywasm2go/p0.Fn3166
func Fn3166(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3187 github.com/goccy/spidermonkeywasm2go/p0.Fn3187
func Fn3187(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3196 github.com/goccy/spidermonkeywasm2go/p6.Fn3196
func Fn3196(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3228 github.com/goccy/spidermonkeywasm2go/p0.Fn3228
func Fn3228(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3234 github.com/goccy/spidermonkeywasm2go/p6.Fn3234
func Fn3234(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3236 github.com/goccy/spidermonkeywasm2go/p7.Fn3236
func Fn3236(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3479 github.com/goccy/spidermonkeywasm2go/p6.Fn3479
func Fn3479(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn3480 github.com/goccy/spidermonkeywasm2go/p7.Fn3480
func Fn3480(m *base.Module, l0 int32)

//go:linkname Fn3492 github.com/goccy/spidermonkeywasm2go/p7.Fn3492
func Fn3492(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3499 github.com/goccy/spidermonkeywasm2go/p7.Fn3499
func Fn3499(m *base.Module, l0 int32)

//go:linkname Fn3522 github.com/goccy/spidermonkeywasm2go/p2.Fn3522
func Fn3522(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3523 github.com/goccy/spidermonkeywasm2go/p0.Fn3523
func Fn3523(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3525 github.com/goccy/spidermonkeywasm2go/p6.Fn3525
func Fn3525(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3528 github.com/goccy/spidermonkeywasm2go/p7.Fn3528
func Fn3528(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3532 github.com/goccy/spidermonkeywasm2go/p7.Fn3532
func Fn3532(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3535 github.com/goccy/spidermonkeywasm2go/p7.Fn3535
func Fn3535(m *base.Module, l0 int32) int32

//go:linkname Fn3536 github.com/goccy/spidermonkeywasm2go/p0.Fn3536
func Fn3536(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3538 github.com/goccy/spidermonkeywasm2go/p0.Fn3538
func Fn3538(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3539 github.com/goccy/spidermonkeywasm2go/p0.Fn3539
func Fn3539(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3540 github.com/goccy/spidermonkeywasm2go/p0.Fn3540
func Fn3540(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3546 github.com/goccy/spidermonkeywasm2go/p0.Fn3546
func Fn3546(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3549 github.com/goccy/spidermonkeywasm2go/p0.Fn3549
func Fn3549(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3550 github.com/goccy/spidermonkeywasm2go/p0.Fn3550
func Fn3550(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3556 github.com/goccy/spidermonkeywasm2go/p0.Fn3556
func Fn3556(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3557 github.com/goccy/spidermonkeywasm2go/p0.Fn3557
func Fn3557(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3559 github.com/goccy/spidermonkeywasm2go/p0.Fn3559
func Fn3559(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3560 github.com/goccy/spidermonkeywasm2go/p0.Fn3560
func Fn3560(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3562 github.com/goccy/spidermonkeywasm2go/p6.Fn3562
func Fn3562(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3564 github.com/goccy/spidermonkeywasm2go/p0.Fn3564
func Fn3564(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3566 github.com/goccy/spidermonkeywasm2go/p7.Fn3566
func Fn3566(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3567 github.com/goccy/spidermonkeywasm2go/p2.Fn3567
func Fn3567(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3575 github.com/goccy/spidermonkeywasm2go/p6.Fn3575
func Fn3575(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3577 github.com/goccy/spidermonkeywasm2go/p2.Fn3577
func Fn3577(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3584 github.com/goccy/spidermonkeywasm2go/p6.Fn3584
func Fn3584(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3585 github.com/goccy/spidermonkeywasm2go/p6.Fn3585
func Fn3585(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3589 github.com/goccy/spidermonkeywasm2go/p6.Fn3589
func Fn3589(m *base.Module, l0 int32)

//go:linkname Fn3604 github.com/goccy/spidermonkeywasm2go/p6.Fn3604
func Fn3604(m *base.Module, l0 int32)

//go:linkname Fn3621 github.com/goccy/spidermonkeywasm2go/p5.Fn3621
func Fn3621(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3628 github.com/goccy/spidermonkeywasm2go/p5.Fn3628
func Fn3628(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3633 github.com/goccy/spidermonkeywasm2go/p0.Fn3633
func Fn3633(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3636 github.com/goccy/spidermonkeywasm2go/p0.Fn3636
func Fn3636(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3663 github.com/goccy/spidermonkeywasm2go/p5.Fn3663
func Fn3663(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3709 github.com/goccy/spidermonkeywasm2go/p0.Fn3709
func Fn3709(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3721 github.com/goccy/spidermonkeywasm2go/p6.Fn3721
func Fn3721(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3722 github.com/goccy/spidermonkeywasm2go/p6.Fn3722
func Fn3722(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3735 github.com/goccy/spidermonkeywasm2go/p6.Fn3735
func Fn3735(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3736 github.com/goccy/spidermonkeywasm2go/p4.Fn3736
func Fn3736(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3755 github.com/goccy/spidermonkeywasm2go/p0.Fn3755
func Fn3755(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3811 github.com/goccy/spidermonkeywasm2go/p6.Fn3811
func Fn3811(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3812 github.com/goccy/spidermonkeywasm2go/p0.Fn3812
func Fn3812(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3813 github.com/goccy/spidermonkeywasm2go/p0.Fn3813
func Fn3813(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3814 github.com/goccy/spidermonkeywasm2go/p0.Fn3814
func Fn3814(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3815 github.com/goccy/spidermonkeywasm2go/p0.Fn3815
func Fn3815(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3816 github.com/goccy/spidermonkeywasm2go/p0.Fn3816
func Fn3816(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3817 github.com/goccy/spidermonkeywasm2go/p0.Fn3817
func Fn3817(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3818 github.com/goccy/spidermonkeywasm2go/p0.Fn3818
func Fn3818(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3819 github.com/goccy/spidermonkeywasm2go/p0.Fn3819
func Fn3819(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3820 github.com/goccy/spidermonkeywasm2go/p0.Fn3820
func Fn3820(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3821 github.com/goccy/spidermonkeywasm2go/p0.Fn3821
func Fn3821(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3822 github.com/goccy/spidermonkeywasm2go/p0.Fn3822
func Fn3822(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3823 github.com/goccy/spidermonkeywasm2go/p0.Fn3823
func Fn3823(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3865 github.com/goccy/spidermonkeywasm2go/p0.Fn3865
func Fn3865(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3867 github.com/goccy/spidermonkeywasm2go/p0.Fn3867
func Fn3867(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3868 github.com/goccy/spidermonkeywasm2go/p5.Fn3868
func Fn3868(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

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

//go:linkname Fn4024 github.com/goccy/spidermonkeywasm2go/p6.Fn4024
func Fn4024(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4027 github.com/goccy/spidermonkeywasm2go/p5.Fn4027
func Fn4027(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn4028 github.com/goccy/spidermonkeywasm2go/p5.Fn4028
func Fn4028(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn4029 github.com/goccy/spidermonkeywasm2go/p5.Fn4029
func Fn4029(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn4030 github.com/goccy/spidermonkeywasm2go/p5.Fn4030
func Fn4030(m *base.Module, l0 int32, l1 float32, l2 int32, l3 int32)

//go:linkname Fn4031 github.com/goccy/spidermonkeywasm2go/p5.Fn4031
func Fn4031(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32)

//go:linkname Fn4032 github.com/goccy/spidermonkeywasm2go/p5.Fn4032
func Fn4032(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn4033 github.com/goccy/spidermonkeywasm2go/p5.Fn4033
func Fn4033(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn4034 github.com/goccy/spidermonkeywasm2go/p6.Fn4034
func Fn4034(m *base.Module, l0 int64) int32

//go:linkname Fn4035 github.com/goccy/spidermonkeywasm2go/p5.Fn4035
func Fn4035(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4036 github.com/goccy/spidermonkeywasm2go/p6.Fn4036
func Fn4036(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn4037 github.com/goccy/spidermonkeywasm2go/p5.Fn4037
func Fn4037(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn4038 github.com/goccy/spidermonkeywasm2go/p5.Fn4038
func Fn4038(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn4039 github.com/goccy/spidermonkeywasm2go/p5.Fn4039
func Fn4039(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn4040 github.com/goccy/spidermonkeywasm2go/p5.Fn4040
func Fn4040(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn4041 github.com/goccy/spidermonkeywasm2go/p5.Fn4041
func Fn4041(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn4042 github.com/goccy/spidermonkeywasm2go/p5.Fn4042
func Fn4042(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn4043 github.com/goccy/spidermonkeywasm2go/p5.Fn4043
func Fn4043(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn4044 github.com/goccy/spidermonkeywasm2go/p6.Fn4044
func Fn4044(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn4045 github.com/goccy/spidermonkeywasm2go/p5.Fn4045
func Fn4045(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn4046 github.com/goccy/spidermonkeywasm2go/p6.Fn4046
func Fn4046(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn4054 github.com/goccy/spidermonkeywasm2go/p6.Fn4054
func Fn4054(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4055 github.com/goccy/spidermonkeywasm2go/p6.Fn4055
func Fn4055(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4056 github.com/goccy/spidermonkeywasm2go/p6.Fn4056
func Fn4056(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4057 github.com/goccy/spidermonkeywasm2go/p6.Fn4057
func Fn4057(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4086 github.com/goccy/spidermonkeywasm2go/p5.Fn4086
func Fn4086(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4087 github.com/goccy/spidermonkeywasm2go/p5.Fn4087
func Fn4087(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4090 github.com/goccy/spidermonkeywasm2go/p6.Fn4090
func Fn4090(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4093 github.com/goccy/spidermonkeywasm2go/p6.Fn4093
func Fn4093(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4094 github.com/goccy/spidermonkeywasm2go/p7.Fn4094
func Fn4094(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4115 github.com/goccy/spidermonkeywasm2go/p5.Fn4115
func Fn4115(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4143 github.com/goccy/spidermonkeywasm2go/p6.Fn4143
func Fn4143(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4176 github.com/goccy/spidermonkeywasm2go/p6.Fn4176
func Fn4176(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4184 github.com/goccy/spidermonkeywasm2go/p6.Fn4184
func Fn4184(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4185 github.com/goccy/spidermonkeywasm2go/p7.Fn4185
func Fn4185(m *base.Module, l0 int32)

//go:linkname Fn4186 github.com/goccy/spidermonkeywasm2go/p7.Fn4186
func Fn4186(m *base.Module, l0 int32)

//go:linkname Fn4192 github.com/goccy/spidermonkeywasm2go/p6.Fn4192
func Fn4192(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4193 github.com/goccy/spidermonkeywasm2go/p7.Fn4193
func Fn4193(m *base.Module, l0 int32)

//go:linkname Fn4200 github.com/goccy/spidermonkeywasm2go/p6.Fn4200
func Fn4200(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4201 github.com/goccy/spidermonkeywasm2go/p6.Fn4201
func Fn4201(m *base.Module, l0 int32) int32

//go:linkname Fn4212 github.com/goccy/spidermonkeywasm2go/p6.Fn4212
func Fn4212(m *base.Module, l0 int32) int32

//go:linkname Fn4216 github.com/goccy/spidermonkeywasm2go/p7.Fn4216
func Fn4216(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4222 github.com/goccy/spidermonkeywasm2go/p6.Fn4222
func Fn4222(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4229 github.com/goccy/spidermonkeywasm2go/p2.Fn4229
func Fn4229(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4230 github.com/goccy/spidermonkeywasm2go/p6.Fn4230
func Fn4230(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4244 github.com/goccy/spidermonkeywasm2go/p4.Fn4244
func Fn4244(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4266 github.com/goccy/spidermonkeywasm2go/p4.Fn4266
func Fn4266(m *base.Module, l0 int32) int32

//go:linkname Fn4288 github.com/goccy/spidermonkeywasm2go/p7.Fn4288
func Fn4288(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4307 github.com/goccy/spidermonkeywasm2go/p4.Fn4307
func Fn4307(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4317 github.com/goccy/spidermonkeywasm2go/p6.Fn4317
func Fn4317(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4318 github.com/goccy/spidermonkeywasm2go/p6.Fn4318
func Fn4318(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4319 github.com/goccy/spidermonkeywasm2go/p4.Fn4319
func Fn4319(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4330 github.com/goccy/spidermonkeywasm2go/p4.Fn4330
func Fn4330(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4332 github.com/goccy/spidermonkeywasm2go/p5.Fn4332
func Fn4332(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4335 github.com/goccy/spidermonkeywasm2go/p5.Fn4335
func Fn4335(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4336 github.com/goccy/spidermonkeywasm2go/p5.Fn4336
func Fn4336(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4337 github.com/goccy/spidermonkeywasm2go/p2.Fn4337
func Fn4337(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4340 github.com/goccy/spidermonkeywasm2go/p6.Fn4340
func Fn4340(m *base.Module, l0 int32, l1 int32, l2 float64, l3 float64)

//go:linkname Fn4346 github.com/goccy/spidermonkeywasm2go/p7.Fn4346
func Fn4346(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4396 github.com/goccy/spidermonkeywasm2go/p7.Fn4396
func Fn4396(m *base.Module, l0 int32)

//go:linkname Fn4402 github.com/goccy/spidermonkeywasm2go/p6.Fn4402
func Fn4402(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4406 github.com/goccy/spidermonkeywasm2go/p4.Fn4406
func Fn4406(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4409 github.com/goccy/spidermonkeywasm2go/p5.Fn4409
func Fn4409(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4410 github.com/goccy/spidermonkeywasm2go/p6.Fn4410
func Fn4410(m *base.Module, l0 int32) int32

//go:linkname Fn4411 github.com/goccy/spidermonkeywasm2go/p7.Fn4411
func Fn4411(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32

//go:linkname Fn4412 github.com/goccy/spidermonkeywasm2go/p7.Fn4412
func Fn4412(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4413 github.com/goccy/spidermonkeywasm2go/p6.Fn4413
func Fn4413(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4414 github.com/goccy/spidermonkeywasm2go/p6.Fn4414
func Fn4414(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4415 github.com/goccy/spidermonkeywasm2go/p6.Fn4415
func Fn4415(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4416 github.com/goccy/spidermonkeywasm2go/p6.Fn4416
func Fn4416(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4417 github.com/goccy/spidermonkeywasm2go/p6.Fn4417
func Fn4417(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4420 github.com/goccy/spidermonkeywasm2go/p6.Fn4420
func Fn4420(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4423 github.com/goccy/spidermonkeywasm2go/p5.Fn4423
func Fn4423(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4425 github.com/goccy/spidermonkeywasm2go/p5.Fn4425
func Fn4425(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4426 github.com/goccy/spidermonkeywasm2go/p4.Fn4426
func Fn4426(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4427 github.com/goccy/spidermonkeywasm2go/p6.Fn4427
func Fn4427(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4448 github.com/goccy/spidermonkeywasm2go/p6.Fn4448
func Fn4448(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4449 github.com/goccy/spidermonkeywasm2go/p7.Fn4449
func Fn4449(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4452 github.com/goccy/spidermonkeywasm2go/p5.Fn4452
func Fn4452(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4453 github.com/goccy/spidermonkeywasm2go/p6.Fn4453
func Fn4453(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4454 github.com/goccy/spidermonkeywasm2go/p6.Fn4454
func Fn4454(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4455 github.com/goccy/spidermonkeywasm2go/p5.Fn4455
func Fn4455(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4494 github.com/goccy/spidermonkeywasm2go/p2.Fn4494
func Fn4494(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4501 github.com/goccy/spidermonkeywasm2go/p7.Fn4501
func Fn4501(m *base.Module, l0 int32)

//go:linkname Fn4505 github.com/goccy/spidermonkeywasm2go/p5.Fn4505
func Fn4505(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4508 github.com/goccy/spidermonkeywasm2go/p7.Fn4508
func Fn4508(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4509 github.com/goccy/spidermonkeywasm2go/p7.Fn4509
func Fn4509(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4510 github.com/goccy/spidermonkeywasm2go/p6.Fn4510
func Fn4510(m *base.Module, l0 int32) int32

//go:linkname Fn4513 github.com/goccy/spidermonkeywasm2go/p6.Fn4513
func Fn4513(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4514 github.com/goccy/spidermonkeywasm2go/p4.Fn4514
func Fn4514(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 float64, l7 int32) int32

//go:linkname Fn4515 github.com/goccy/spidermonkeywasm2go/p6.Fn4515
func Fn4515(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4521 github.com/goccy/spidermonkeywasm2go/p6.Fn4521
func Fn4521(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4534 github.com/goccy/spidermonkeywasm2go/p6.Fn4534
func Fn4534(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4535 github.com/goccy/spidermonkeywasm2go/p6.Fn4535
func Fn4535(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4536 github.com/goccy/spidermonkeywasm2go/p6.Fn4536
func Fn4536(m *base.Module, l0 int32) int32

//go:linkname Fn4541 github.com/goccy/spidermonkeywasm2go/p5.Fn4541
func Fn4541(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4545 github.com/goccy/spidermonkeywasm2go/p6.Fn4545
func Fn4545(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4546 github.com/goccy/spidermonkeywasm2go/p6.Fn4546
func Fn4546(m *base.Module, l0 int32) int32

//go:linkname Fn4547 github.com/goccy/spidermonkeywasm2go/p2.Fn4547
func Fn4547(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4557 github.com/goccy/spidermonkeywasm2go/p4.Fn4557
func Fn4557(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4563 github.com/goccy/spidermonkeywasm2go/p7.Fn4563
func Fn4563(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4564 github.com/goccy/spidermonkeywasm2go/p6.Fn4564
func Fn4564(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4567 github.com/goccy/spidermonkeywasm2go/p4.Fn4567
func Fn4567(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4568 github.com/goccy/spidermonkeywasm2go/p7.Fn4568
func Fn4568(m *base.Module, l0 int32) int32

//go:linkname Fn4574 github.com/goccy/spidermonkeywasm2go/p6.Fn4574
func Fn4574(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4576 github.com/goccy/spidermonkeywasm2go/p4.Fn4576
func Fn4576(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn4598 github.com/goccy/spidermonkeywasm2go/p2.Fn4598
func Fn4598(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4599 github.com/goccy/spidermonkeywasm2go/p5.Fn4599
func Fn4599(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn4601 github.com/goccy/spidermonkeywasm2go/p6.Fn4601
func Fn4601(m *base.Module, l0 int32) int32

//go:linkname Fn4622 github.com/goccy/spidermonkeywasm2go/p7.Fn4622
func Fn4622(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4645 github.com/goccy/spidermonkeywasm2go/p6.Fn4645
func Fn4645(m *base.Module, l0 int32) int32

//go:linkname Fn4646 github.com/goccy/spidermonkeywasm2go/p6.Fn4646
func Fn4646(m *base.Module, l0 int32)

//go:linkname Fn4647 github.com/goccy/spidermonkeywasm2go/p6.Fn4647
func Fn4647(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4650 github.com/goccy/spidermonkeywasm2go/p5.Fn4650
func Fn4650(m *base.Module, l0 int32) int32

//go:linkname Fn4657 github.com/goccy/spidermonkeywasm2go/p6.Fn4657
func Fn4657(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4658 github.com/goccy/spidermonkeywasm2go/p6.Fn4658
func Fn4658(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4684 github.com/goccy/spidermonkeywasm2go/p4.Fn4684
func Fn4684(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4694 github.com/goccy/spidermonkeywasm2go/p5.Fn4694
func Fn4694(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4695 github.com/goccy/spidermonkeywasm2go/p0.Fn4695
func Fn4695(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4708 github.com/goccy/spidermonkeywasm2go/p7.Fn4708
func Fn4708(m *base.Module, l0 int32)

//go:linkname Fn4725 github.com/goccy/spidermonkeywasm2go/p7.Fn4725
func Fn4725(m *base.Module, l0 int32)

//go:linkname Fn4728 github.com/goccy/spidermonkeywasm2go/p6.Fn4728
func Fn4728(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4733 github.com/goccy/spidermonkeywasm2go/p5.Fn4733
func Fn4733(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4735 github.com/goccy/spidermonkeywasm2go/p5.Fn4735
func Fn4735(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4742 github.com/goccy/spidermonkeywasm2go/p5.Fn4742
func Fn4742(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5135 github.com/goccy/spidermonkeywasm2go/p6.Fn5135
func Fn5135(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5144 github.com/goccy/spidermonkeywasm2go/p0.Fn5144
func Fn5144(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5148 github.com/goccy/spidermonkeywasm2go/p0.Fn5148
func Fn5148(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn5149 github.com/goccy/spidermonkeywasm2go/p0.Fn5149
func Fn5149(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn5150 github.com/goccy/spidermonkeywasm2go/p6.Fn5150
func Fn5150(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5151 github.com/goccy/spidermonkeywasm2go/p2.Fn5151
func Fn5151(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5338 github.com/goccy/spidermonkeywasm2go/p5.Fn5338
func Fn5338(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5339 github.com/goccy/spidermonkeywasm2go/p4.Fn5339
func Fn5339(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5341 github.com/goccy/spidermonkeywasm2go/p5.Fn5341
func Fn5341(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5342 github.com/goccy/spidermonkeywasm2go/p5.Fn5342
func Fn5342(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5346 github.com/goccy/spidermonkeywasm2go/p4.Fn5346
func Fn5346(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5348 github.com/goccy/spidermonkeywasm2go/p6.Fn5348
func Fn5348(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5426 github.com/goccy/spidermonkeywasm2go/p7.Fn5426
func Fn5426(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5458 github.com/goccy/spidermonkeywasm2go/p7.Fn5458
func Fn5458(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5538 github.com/goccy/spidermonkeywasm2go/p6.Fn5538
func Fn5538(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5556 github.com/goccy/spidermonkeywasm2go/p6.Fn5556
func Fn5556(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5557 github.com/goccy/spidermonkeywasm2go/p7.Fn5557
func Fn5557(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5560 github.com/goccy/spidermonkeywasm2go/p7.Fn5560
func Fn5560(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5561 github.com/goccy/spidermonkeywasm2go/p7.Fn5561
func Fn5561(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5562 github.com/goccy/spidermonkeywasm2go/p4.Fn5562
func Fn5562(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5563 github.com/goccy/spidermonkeywasm2go/p0.Fn5563
func Fn5563(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5564 github.com/goccy/spidermonkeywasm2go/p0.Fn5564
func Fn5564(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5572 github.com/goccy/spidermonkeywasm2go/p4.Fn5572
func Fn5572(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5574 github.com/goccy/spidermonkeywasm2go/p7.Fn5574
func Fn5574(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5582 github.com/goccy/spidermonkeywasm2go/p6.Fn5582
func Fn5582(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5586 github.com/goccy/spidermonkeywasm2go/p5.Fn5586
func Fn5586(m *base.Module, l0 int32)

//go:linkname Fn5587 github.com/goccy/spidermonkeywasm2go/p6.Fn5587
func Fn5587(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5588 github.com/goccy/spidermonkeywasm2go/p6.Fn5588
func Fn5588(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5589 github.com/goccy/spidermonkeywasm2go/p6.Fn5589
func Fn5589(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5590 github.com/goccy/spidermonkeywasm2go/p6.Fn5590
func Fn5590(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5591 github.com/goccy/spidermonkeywasm2go/p4.Fn5591
func Fn5591(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5594 github.com/goccy/spidermonkeywasm2go/p5.Fn5594
func Fn5594(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5595 github.com/goccy/spidermonkeywasm2go/p6.Fn5595
func Fn5595(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5597 github.com/goccy/spidermonkeywasm2go/p6.Fn5597
func Fn5597(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5600 github.com/goccy/spidermonkeywasm2go/p6.Fn5600
func Fn5600(m *base.Module, l0 int32)

//go:linkname Fn5614 github.com/goccy/spidermonkeywasm2go/p2.Fn5614
func Fn5614(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5615 github.com/goccy/spidermonkeywasm2go/p4.Fn5615
func Fn5615(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5616 github.com/goccy/spidermonkeywasm2go/p4.Fn5616
func Fn5616(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5617 github.com/goccy/spidermonkeywasm2go/p6.Fn5617
func Fn5617(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn5618 github.com/goccy/spidermonkeywasm2go/p6.Fn5618
func Fn5618(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn5621 github.com/goccy/spidermonkeywasm2go/p6.Fn5621
func Fn5621(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5622 github.com/goccy/spidermonkeywasm2go/p6.Fn5622
func Fn5622(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5623 github.com/goccy/spidermonkeywasm2go/p4.Fn5623
func Fn5623(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5624 github.com/goccy/spidermonkeywasm2go/p5.Fn5624
func Fn5624(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5637 github.com/goccy/spidermonkeywasm2go/p7.Fn5637
func Fn5637(m *base.Module, l0 float64) float64

//go:linkname Fn5639 github.com/goccy/spidermonkeywasm2go/p6.Fn5639
func Fn5639(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5647 github.com/goccy/spidermonkeywasm2go/p5.Fn5647
func Fn5647(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5651 github.com/goccy/spidermonkeywasm2go/p6.Fn5651
func Fn5651(m *base.Module, l0 int32)

//go:linkname Fn5653 github.com/goccy/spidermonkeywasm2go/p6.Fn5653
func Fn5653(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5658 github.com/goccy/spidermonkeywasm2go/p6.Fn5658
func Fn5658(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5660 github.com/goccy/spidermonkeywasm2go/p7.Fn5660
func Fn5660(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5662 github.com/goccy/spidermonkeywasm2go/p0.Fn5662
func Fn5662(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5676 github.com/goccy/spidermonkeywasm2go/p5.Fn5676
func Fn5676(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5677 github.com/goccy/spidermonkeywasm2go/p5.Fn5677
func Fn5677(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5679 github.com/goccy/spidermonkeywasm2go/p6.Fn5679
func Fn5679(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5682 github.com/goccy/spidermonkeywasm2go/p5.Fn5682
func Fn5682(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5687 github.com/goccy/spidermonkeywasm2go/p5.Fn5687
func Fn5687(m *base.Module, l0 int32)

//go:linkname Fn5690 github.com/goccy/spidermonkeywasm2go/p6.Fn5690
func Fn5690(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5691 github.com/goccy/spidermonkeywasm2go/p7.Fn5691
func Fn5691(m *base.Module, l0 int32)

//go:linkname Fn5693 github.com/goccy/spidermonkeywasm2go/p6.Fn5693
func Fn5693(m *base.Module, l0 int32) int32

//go:linkname Fn5694 github.com/goccy/spidermonkeywasm2go/p7.Fn5694
func Fn5694(m *base.Module, l0 int32) int32

//go:linkname Fn5700 github.com/goccy/spidermonkeywasm2go/p4.Fn5700
func Fn5700(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5706 github.com/goccy/spidermonkeywasm2go/p5.Fn5706
func Fn5706(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5707 github.com/goccy/spidermonkeywasm2go/p6.Fn5707
func Fn5707(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5708 github.com/goccy/spidermonkeywasm2go/p5.Fn5708
func Fn5708(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5710 github.com/goccy/spidermonkeywasm2go/p7.Fn5710
func Fn5710(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5711 github.com/goccy/spidermonkeywasm2go/p5.Fn5711
func Fn5711(m *base.Module, l0 int32, l1 int32)

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

//go:linkname Fn5718 github.com/goccy/spidermonkeywasm2go/p5.Fn5718
func Fn5718(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5719 github.com/goccy/spidermonkeywasm2go/p6.Fn5719
func Fn5719(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5720 github.com/goccy/spidermonkeywasm2go/p7.Fn5720
func Fn5720(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5721 github.com/goccy/spidermonkeywasm2go/p6.Fn5721
func Fn5721(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5723 github.com/goccy/spidermonkeywasm2go/p6.Fn5723
func Fn5723(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5725 github.com/goccy/spidermonkeywasm2go/p7.Fn5725
func Fn5725(m *base.Module, l0 int32)

//go:linkname Fn5726 github.com/goccy/spidermonkeywasm2go/p7.Fn5726
func Fn5726(m *base.Module, l0 int32)

//go:linkname Fn5731 github.com/goccy/spidermonkeywasm2go/p6.Fn5731
func Fn5731(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5739 github.com/goccy/spidermonkeywasm2go/p6.Fn5739
func Fn5739(m *base.Module, l0 int32)

//go:linkname Fn5743 github.com/goccy/spidermonkeywasm2go/p6.Fn5743
func Fn5743(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5744 github.com/goccy/spidermonkeywasm2go/p7.Fn5744
func Fn5744(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5763 github.com/goccy/spidermonkeywasm2go/p5.Fn5763
func Fn5763(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5764 github.com/goccy/spidermonkeywasm2go/p6.Fn5764
func Fn5764(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5765 github.com/goccy/spidermonkeywasm2go/p5.Fn5765
func Fn5765(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5809 github.com/goccy/spidermonkeywasm2go/p6.Fn5809
func Fn5809(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5810 github.com/goccy/spidermonkeywasm2go/p6.Fn5810
func Fn5810(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5811 github.com/goccy/spidermonkeywasm2go/p6.Fn5811
func Fn5811(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5812 github.com/goccy/spidermonkeywasm2go/p6.Fn5812
func Fn5812(m *base.Module, l0 int32)

//go:linkname Fn5827 github.com/goccy/spidermonkeywasm2go/p6.Fn5827
func Fn5827(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5828 github.com/goccy/spidermonkeywasm2go/p6.Fn5828
func Fn5828(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5869 github.com/goccy/spidermonkeywasm2go/p6.Fn5869
func Fn5869(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5870 github.com/goccy/spidermonkeywasm2go/p6.Fn5870
func Fn5870(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5876 github.com/goccy/spidermonkeywasm2go/p6.Fn5876
func Fn5876(m *base.Module, l0 int32) int32

//go:linkname Fn5890 github.com/goccy/spidermonkeywasm2go/p4.Fn5890
func Fn5890(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5904 github.com/goccy/spidermonkeywasm2go/p7.Fn5904
func Fn5904(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5905 github.com/goccy/spidermonkeywasm2go/p7.Fn5905
func Fn5905(m *base.Module, l0 int32) int32

//go:linkname Fn5908 github.com/goccy/spidermonkeywasm2go/p7.Fn5908
func Fn5908(m *base.Module, l0 int32)

//go:linkname Fn5913 github.com/goccy/spidermonkeywasm2go/p5.Fn5913
func Fn5913(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5940 github.com/goccy/spidermonkeywasm2go/p5.Fn5940
func Fn5940(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5941 github.com/goccy/spidermonkeywasm2go/p7.Fn5941
func Fn5941(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5951 github.com/goccy/spidermonkeywasm2go/p4.Fn5951
func Fn5951(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5956 github.com/goccy/spidermonkeywasm2go/p6.Fn5956
func Fn5956(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6028 github.com/goccy/spidermonkeywasm2go/p4.Fn6028
func Fn6028(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6040 github.com/goccy/spidermonkeywasm2go/p6.Fn6040
func Fn6040(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6048 github.com/goccy/spidermonkeywasm2go/p7.Fn6048
func Fn6048(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6049 github.com/goccy/spidermonkeywasm2go/p7.Fn6049
func Fn6049(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6052 github.com/goccy/spidermonkeywasm2go/p6.Fn6052
func Fn6052(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6065 github.com/goccy/spidermonkeywasm2go/p6.Fn6065
func Fn6065(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6080 github.com/goccy/spidermonkeywasm2go/p7.Fn6080
func Fn6080(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6092 github.com/goccy/spidermonkeywasm2go/p7.Fn6092
func Fn6092(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6093 github.com/goccy/spidermonkeywasm2go/p4.Fn6093
func Fn6093(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)

//go:linkname Fn6094 github.com/goccy/spidermonkeywasm2go/p7.Fn6094
func Fn6094(m *base.Module)

//go:linkname Fn6113 github.com/goccy/spidermonkeywasm2go/p7.Fn6113
func Fn6113(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6174 github.com/goccy/spidermonkeywasm2go/p5.Fn6174
func Fn6174(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6175 github.com/goccy/spidermonkeywasm2go/p6.Fn6175
func Fn6175(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6177 github.com/goccy/spidermonkeywasm2go/p6.Fn6177
func Fn6177(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6178 github.com/goccy/spidermonkeywasm2go/p5.Fn6178
func Fn6178(m *base.Module, l0 int32)

//go:linkname Fn6180 github.com/goccy/spidermonkeywasm2go/p6.Fn6180
func Fn6180(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6198 github.com/goccy/spidermonkeywasm2go/p6.Fn6198
func Fn6198(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6221 github.com/goccy/spidermonkeywasm2go/p4.Fn6221
func Fn6221(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6234 github.com/goccy/spidermonkeywasm2go/p6.Fn6234
func Fn6234(m *base.Module, l0 int32) int32

//go:linkname Fn6248 github.com/goccy/spidermonkeywasm2go/p6.Fn6248
func Fn6248(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6249 github.com/goccy/spidermonkeywasm2go/p2.Fn6249
func Fn6249(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn6250 github.com/goccy/spidermonkeywasm2go/p6.Fn6250
func Fn6250(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6251 github.com/goccy/spidermonkeywasm2go/p7.Fn6251
func Fn6251(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6252 github.com/goccy/spidermonkeywasm2go/p6.Fn6252
func Fn6252(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6253 github.com/goccy/spidermonkeywasm2go/p4.Fn6253
func Fn6253(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6272 github.com/goccy/spidermonkeywasm2go/p6.Fn6272
func Fn6272(m *base.Module) int32

//go:linkname Fn6280 github.com/goccy/spidermonkeywasm2go/p5.Fn6280
func Fn6280(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6281 github.com/goccy/spidermonkeywasm2go/p7.Fn6281
func Fn6281(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6283 github.com/goccy/spidermonkeywasm2go/p5.Fn6283
func Fn6283(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn6334 github.com/goccy/spidermonkeywasm2go/p2.Fn6334
func Fn6334(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6346 github.com/goccy/spidermonkeywasm2go/p7.Fn6346
func Fn6346(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6643 github.com/goccy/spidermonkeywasm2go/p4.Fn6643
func Fn6643(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6652 github.com/goccy/spidermonkeywasm2go/p5.Fn6652
func Fn6652(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6653 github.com/goccy/spidermonkeywasm2go/p6.Fn6653
func Fn6653(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6654 github.com/goccy/spidermonkeywasm2go/p7.Fn6654
func Fn6654(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6655 github.com/goccy/spidermonkeywasm2go/p6.Fn6655
func Fn6655(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6656 github.com/goccy/spidermonkeywasm2go/p6.Fn6656
func Fn6656(m *base.Module, l0 int32) int32

//go:linkname Fn6669 github.com/goccy/spidermonkeywasm2go/p6.Fn6669
func Fn6669(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6676 github.com/goccy/spidermonkeywasm2go/p1.Fn6676
func Fn6676(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6678 github.com/goccy/spidermonkeywasm2go/p7.Fn6678
func Fn6678(m *base.Module, l0 int32)

//go:linkname Fn6681 github.com/goccy/spidermonkeywasm2go/p6.Fn6681
func Fn6681(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6683 github.com/goccy/spidermonkeywasm2go/p4.Fn6683
func Fn6683(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6684 github.com/goccy/spidermonkeywasm2go/p7.Fn6684
func Fn6684(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6687 github.com/goccy/spidermonkeywasm2go/p6.Fn6687
func Fn6687(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6688 github.com/goccy/spidermonkeywasm2go/p6.Fn6688
func Fn6688(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6689 github.com/goccy/spidermonkeywasm2go/p6.Fn6689
func Fn6689(m *base.Module, l0 int32, l1 int32) int32

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

//go:linkname Fn6702 github.com/goccy/spidermonkeywasm2go/p7.Fn6702
func Fn6702(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6703 github.com/goccy/spidermonkeywasm2go/p7.Fn6703
func Fn6703(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6705 github.com/goccy/spidermonkeywasm2go/p7.Fn6705
func Fn6705(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6710 github.com/goccy/spidermonkeywasm2go/p7.Fn6710
func Fn6710(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6711 github.com/goccy/spidermonkeywasm2go/p7.Fn6711
func Fn6711(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6712 github.com/goccy/spidermonkeywasm2go/p7.Fn6712
func Fn6712(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6715 github.com/goccy/spidermonkeywasm2go/p7.Fn6715
func Fn6715(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6720 github.com/goccy/spidermonkeywasm2go/p7.Fn6720
func Fn6720(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6721 github.com/goccy/spidermonkeywasm2go/p6.Fn6721
func Fn6721(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6748 github.com/goccy/spidermonkeywasm2go/p7.Fn6748
func Fn6748(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6751 github.com/goccy/spidermonkeywasm2go/p7.Fn6751
func Fn6751(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6753 github.com/goccy/spidermonkeywasm2go/p1.Fn6753
func Fn6753(m *base.Module, l0 int32) int32

//go:linkname Fn6759 github.com/goccy/spidermonkeywasm2go/p7.Fn6759
func Fn6759(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6767 github.com/goccy/spidermonkeywasm2go/p7.Fn6767
func Fn6767(m *base.Module, l0 int32) int32

//go:linkname Fn6771 github.com/goccy/spidermonkeywasm2go/p7.Fn6771
func Fn6771(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6772 github.com/goccy/spidermonkeywasm2go/p7.Fn6772
func Fn6772(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6773 github.com/goccy/spidermonkeywasm2go/p4.Fn6773
func Fn6773(m *base.Module, l0 int32) int32

//go:linkname Fn6774 github.com/goccy/spidermonkeywasm2go/p6.Fn6774
func Fn6774(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6778 github.com/goccy/spidermonkeywasm2go/p7.Fn6778
func Fn6778(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6783 github.com/goccy/spidermonkeywasm2go/p5.Fn6783
func Fn6783(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6784 github.com/goccy/spidermonkeywasm2go/p6.Fn6784
func Fn6784(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6823 github.com/goccy/spidermonkeywasm2go/p6.Fn6823
func Fn6823(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6824 github.com/goccy/spidermonkeywasm2go/p5.Fn6824
func Fn6824(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6827 github.com/goccy/spidermonkeywasm2go/p1.Fn6827
func Fn6827(m *base.Module, l0 int32) int32

//go:linkname Fn6852 github.com/goccy/spidermonkeywasm2go/p5.Fn6852
func Fn6852(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6886 github.com/goccy/spidermonkeywasm2go/p6.Fn6886
func Fn6886(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6889 github.com/goccy/spidermonkeywasm2go/p5.Fn6889
func Fn6889(m *base.Module, l0 int32, l1 int32, l2 int32) int64

//go:linkname Fn6893 github.com/goccy/spidermonkeywasm2go/p7.Fn6893
func Fn6893(m *base.Module, l0 int32) int32

//go:linkname Fn6899 github.com/goccy/spidermonkeywasm2go/p7.Fn6899
func Fn6899(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6912 github.com/goccy/spidermonkeywasm2go/p5.Fn6912
func Fn6912(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6915 github.com/goccy/spidermonkeywasm2go/p6.Fn6915
func Fn6915(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6935 github.com/goccy/spidermonkeywasm2go/p7.Fn6935
func Fn6935(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6937 github.com/goccy/spidermonkeywasm2go/p0.Fn6937
func Fn6937(m *base.Module, l0 int32) int32

//go:linkname Fn6940 github.com/goccy/spidermonkeywasm2go/p0.Fn6940
func Fn6940(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6941 github.com/goccy/spidermonkeywasm2go/p0.Fn6941
func Fn6941(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6955 github.com/goccy/spidermonkeywasm2go/p7.Fn6955
func Fn6955(m *base.Module, l0 int32)

//go:linkname Fn6956 github.com/goccy/spidermonkeywasm2go/p5.Fn6956
func Fn6956(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6958 github.com/goccy/spidermonkeywasm2go/p6.Fn6958
func Fn6958(m *base.Module, l0 int32) int32

//go:linkname Fn6959 github.com/goccy/spidermonkeywasm2go/p5.Fn6959
func Fn6959(m *base.Module, l0 int32) int32

//go:linkname Fn6966 github.com/goccy/spidermonkeywasm2go/p0.Fn6966
func Fn6966(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6967 github.com/goccy/spidermonkeywasm2go/p7.Fn6967
func Fn6967(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6969 github.com/goccy/spidermonkeywasm2go/p7.Fn6969
func Fn6969(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6972 github.com/goccy/spidermonkeywasm2go/p6.Fn6972
func Fn6972(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7002 github.com/goccy/spidermonkeywasm2go/p6.Fn7002
func Fn7002(m *base.Module, l0 int32) int32

//go:linkname Fn7069 github.com/goccy/spidermonkeywasm2go/p4.Fn7069
func Fn7069(m *base.Module, l0 int32) int32

//go:linkname Fn7103 github.com/goccy/spidermonkeywasm2go/p6.Fn7103
func Fn7103(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7104 github.com/goccy/spidermonkeywasm2go/p6.Fn7104
func Fn7104(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7105 github.com/goccy/spidermonkeywasm2go/p6.Fn7105
func Fn7105(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7106 github.com/goccy/spidermonkeywasm2go/p7.Fn7106
func Fn7106(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7109 github.com/goccy/spidermonkeywasm2go/p5.Fn7109
func Fn7109(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7113 github.com/goccy/spidermonkeywasm2go/p5.Fn7113
func Fn7113(m *base.Module, l0 int32) int32

//go:linkname Fn7116 github.com/goccy/spidermonkeywasm2go/p4.Fn7116
func Fn7116(m *base.Module, l0 int32)

//go:linkname Fn7121 github.com/goccy/spidermonkeywasm2go/p6.Fn7121
func Fn7121(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7124 github.com/goccy/spidermonkeywasm2go/p6.Fn7124
func Fn7124(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7132 github.com/goccy/spidermonkeywasm2go/p6.Fn7132
func Fn7132(m *base.Module, l0 int32) int32

//go:linkname Fn7133 github.com/goccy/spidermonkeywasm2go/p6.Fn7133
func Fn7133(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn7150 github.com/goccy/spidermonkeywasm2go/p6.Fn7150
func Fn7150(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn7152 github.com/goccy/spidermonkeywasm2go/p6.Fn7152
func Fn7152(m *base.Module, l0 int32)

//go:linkname Fn7154 github.com/goccy/spidermonkeywasm2go/p5.Fn7154
func Fn7154(m *base.Module, l0 int32) int32

//go:linkname Fn7155 github.com/goccy/spidermonkeywasm2go/p6.Fn7155
func Fn7155(m *base.Module, l0 int32) int32

//go:linkname Fn7157 github.com/goccy/spidermonkeywasm2go/p7.Fn7157
func Fn7157(m *base.Module, l0 int32) int32

//go:linkname Fn7161 github.com/goccy/spidermonkeywasm2go/p7.Fn7161
func Fn7161(m *base.Module, l0 int32) int32

//go:linkname Fn7163 github.com/goccy/spidermonkeywasm2go/p4.Fn7163
func Fn7163(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7164 github.com/goccy/spidermonkeywasm2go/p6.Fn7164
func Fn7164(m *base.Module, l0 int32, l1 int32)

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

//go:linkname Fn7184 github.com/goccy/spidermonkeywasm2go/p6.Fn7184
func Fn7184(m *base.Module, l0 int32)

//go:linkname Fn7185 github.com/goccy/spidermonkeywasm2go/p6.Fn7185
func Fn7185(m *base.Module, l0 int32) int32

//go:linkname Fn7187 github.com/goccy/spidermonkeywasm2go/p6.Fn7187
func Fn7187(m *base.Module, l0 int32) int32

//go:linkname Fn7215 github.com/goccy/spidermonkeywasm2go/p5.Fn7215
func Fn7215(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7230 github.com/goccy/spidermonkeywasm2go/p6.Fn7230
func Fn7230(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7241 github.com/goccy/spidermonkeywasm2go/p6.Fn7241
func Fn7241(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7347 github.com/goccy/spidermonkeywasm2go/p7.Fn7347
func Fn7347(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7348 github.com/goccy/spidermonkeywasm2go/p5.Fn7348
func Fn7348(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7362 github.com/goccy/spidermonkeywasm2go/p6.Fn7362
func Fn7362(m *base.Module, l0 int32)

//go:linkname Fn7369 github.com/goccy/spidermonkeywasm2go/p6.Fn7369
func Fn7369(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7377 github.com/goccy/spidermonkeywasm2go/p6.Fn7377
func Fn7377(m *base.Module, l0 int32)

//go:linkname Fn7414 github.com/goccy/spidermonkeywasm2go/p1.Fn7414
func Fn7414(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7423 github.com/goccy/spidermonkeywasm2go/p2.Fn7423
func Fn7423(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7425 github.com/goccy/spidermonkeywasm2go/p4.Fn7425
func Fn7425(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7429 github.com/goccy/spidermonkeywasm2go/p7.Fn7429
func Fn7429(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7432 github.com/goccy/spidermonkeywasm2go/p6.Fn7432
func Fn7432(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7433 github.com/goccy/spidermonkeywasm2go/p5.Fn7433
func Fn7433(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7436 github.com/goccy/spidermonkeywasm2go/p7.Fn7436
func Fn7436(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7437 github.com/goccy/spidermonkeywasm2go/p4.Fn7437
func Fn7437(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7439 github.com/goccy/spidermonkeywasm2go/p6.Fn7439
func Fn7439(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7440 github.com/goccy/spidermonkeywasm2go/p6.Fn7440
func Fn7440(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7442 github.com/goccy/spidermonkeywasm2go/p7.Fn7442
func Fn7442(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7445 github.com/goccy/spidermonkeywasm2go/p6.Fn7445
func Fn7445(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7456 github.com/goccy/spidermonkeywasm2go/p5.Fn7456
func Fn7456(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7457 github.com/goccy/spidermonkeywasm2go/p6.Fn7457
func Fn7457(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7459 github.com/goccy/spidermonkeywasm2go/p6.Fn7459
func Fn7459(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7460 github.com/goccy/spidermonkeywasm2go/p5.Fn7460
func Fn7460(m *base.Module, l0 int32, l1 float64, l2 float64, l3 float64) int32

//go:linkname Fn7461 github.com/goccy/spidermonkeywasm2go/p4.Fn7461
func Fn7461(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7462 github.com/goccy/spidermonkeywasm2go/p6.Fn7462
func Fn7462(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7463 github.com/goccy/spidermonkeywasm2go/p6.Fn7463
func Fn7463(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn7464 github.com/goccy/spidermonkeywasm2go/p7.Fn7464
func Fn7464(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7465 github.com/goccy/spidermonkeywasm2go/p6.Fn7465
func Fn7465(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7468 github.com/goccy/spidermonkeywasm2go/p6.Fn7468
func Fn7468(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7469 github.com/goccy/spidermonkeywasm2go/p7.Fn7469
func Fn7469(m *base.Module, l0 int32, l1 int32, l2 float64)

//go:linkname Fn7470 github.com/goccy/spidermonkeywasm2go/p6.Fn7470
func Fn7470(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn7471 github.com/goccy/spidermonkeywasm2go/p5.Fn7471
func Fn7471(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7473 github.com/goccy/spidermonkeywasm2go/p5.Fn7473
func Fn7473(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int64, l6 int32) int32

//go:linkname Fn7474 github.com/goccy/spidermonkeywasm2go/p5.Fn7474
func Fn7474(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn7476 github.com/goccy/spidermonkeywasm2go/p6.Fn7476
func Fn7476(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7478 github.com/goccy/spidermonkeywasm2go/p5.Fn7478
func Fn7478(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7479 github.com/goccy/spidermonkeywasm2go/p6.Fn7479
func Fn7479(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn7480 github.com/goccy/spidermonkeywasm2go/p7.Fn7480
func Fn7480(m *base.Module, l0 int32) int32

//go:linkname Fn7483 github.com/goccy/spidermonkeywasm2go/p6.Fn7483
func Fn7483(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7485 github.com/goccy/spidermonkeywasm2go/p4.Fn7485
func Fn7485(m *base.Module, l0 int32, l1 float64, l2 float64, l3 float64, l4 float64, l5 float64, l6 float64)

//go:linkname Fn7486 github.com/goccy/spidermonkeywasm2go/p7.Fn7486
func Fn7486(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7488 github.com/goccy/spidermonkeywasm2go/p6.Fn7488
func Fn7488(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7490 github.com/goccy/spidermonkeywasm2go/p6.Fn7490
func Fn7490(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32) int32

//go:linkname Fn7491 github.com/goccy/spidermonkeywasm2go/p7.Fn7491
func Fn7491(m *base.Module, l0 int32, l1 float64, l2 int32) int32

//go:linkname Fn7492 github.com/goccy/spidermonkeywasm2go/p4.Fn7492
func Fn7492(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7493 github.com/goccy/spidermonkeywasm2go/p2.Fn7493
func Fn7493(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7496 github.com/goccy/spidermonkeywasm2go/p7.Fn7496
func Fn7496(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7497 github.com/goccy/spidermonkeywasm2go/p6.Fn7497
func Fn7497(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn7499 github.com/goccy/spidermonkeywasm2go/p4.Fn7499
func Fn7499(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32

//go:linkname Fn7500 github.com/goccy/spidermonkeywasm2go/p6.Fn7500
func Fn7500(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7509 github.com/goccy/spidermonkeywasm2go/p6.Fn7509
func Fn7509(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7552 github.com/goccy/spidermonkeywasm2go/p5.Fn7552
func Fn7552(m *base.Module, l0 float64, l1 int32) int32

//go:linkname Fn7681 github.com/goccy/spidermonkeywasm2go/p7.Fn7681
func Fn7681(m *base.Module, l0 int32) int32

//go:linkname Fn7683 github.com/goccy/spidermonkeywasm2go/p6.Fn7683
func Fn7683(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7684 github.com/goccy/spidermonkeywasm2go/p5.Fn7684
func Fn7684(m *base.Module, l0 int32, l1 float64, l2 float64, l3 float64, l4 float64, l5 float64, l6 float64) int32

//go:linkname Fn7685 github.com/goccy/spidermonkeywasm2go/p5.Fn7685
func Fn7685(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn7686 github.com/goccy/spidermonkeywasm2go/p6.Fn7686
func Fn7686(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn7687 github.com/goccy/spidermonkeywasm2go/p6.Fn7687
func Fn7687(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7688 github.com/goccy/spidermonkeywasm2go/p6.Fn7688
func Fn7688(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7689 github.com/goccy/spidermonkeywasm2go/p4.Fn7689
func Fn7689(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn7695 github.com/goccy/spidermonkeywasm2go/p6.Fn7695
func Fn7695(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7696 github.com/goccy/spidermonkeywasm2go/p6.Fn7696
func Fn7696(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7701 github.com/goccy/spidermonkeywasm2go/p5.Fn7701
func Fn7701(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7702 github.com/goccy/spidermonkeywasm2go/p7.Fn7702
func Fn7702(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7703 github.com/goccy/spidermonkeywasm2go/p6.Fn7703
func Fn7703(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn7704 github.com/goccy/spidermonkeywasm2go/p7.Fn7704
func Fn7704(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7705 github.com/goccy/spidermonkeywasm2go/p6.Fn7705
func Fn7705(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7706 github.com/goccy/spidermonkeywasm2go/p4.Fn7706
func Fn7706(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7707 github.com/goccy/spidermonkeywasm2go/p7.Fn7707
func Fn7707(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7709 github.com/goccy/spidermonkeywasm2go/p5.Fn7709
func Fn7709(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7718 github.com/goccy/spidermonkeywasm2go/p5.Fn7718
func Fn7718(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn7905 github.com/goccy/spidermonkeywasm2go/p5.Fn7905
func Fn7905(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7906 github.com/goccy/spidermonkeywasm2go/p5.Fn7906
func Fn7906(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7908 github.com/goccy/spidermonkeywasm2go/p5.Fn7908
func Fn7908(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7909 github.com/goccy/spidermonkeywasm2go/p6.Fn7909
func Fn7909(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7910 github.com/goccy/spidermonkeywasm2go/p5.Fn7910
func Fn7910(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7911 github.com/goccy/spidermonkeywasm2go/p6.Fn7911
func Fn7911(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7912 github.com/goccy/spidermonkeywasm2go/p6.Fn7912
func Fn7912(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7913 github.com/goccy/spidermonkeywasm2go/p5.Fn7913
func Fn7913(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7914 github.com/goccy/spidermonkeywasm2go/p5.Fn7914
func Fn7914(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7915 github.com/goccy/spidermonkeywasm2go/p5.Fn7915
func Fn7915(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7916 github.com/goccy/spidermonkeywasm2go/p5.Fn7916
func Fn7916(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7917 github.com/goccy/spidermonkeywasm2go/p6.Fn7917
func Fn7917(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7918 github.com/goccy/spidermonkeywasm2go/p5.Fn7918
func Fn7918(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7919 github.com/goccy/spidermonkeywasm2go/p5.Fn7919
func Fn7919(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7920 github.com/goccy/spidermonkeywasm2go/p5.Fn7920
func Fn7920(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7921 github.com/goccy/spidermonkeywasm2go/p5.Fn7921
func Fn7921(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7923 github.com/goccy/spidermonkeywasm2go/p5.Fn7923
func Fn7923(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7924 github.com/goccy/spidermonkeywasm2go/p5.Fn7924
func Fn7924(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7931 github.com/goccy/spidermonkeywasm2go/p5.Fn7931
func Fn7931(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7932 github.com/goccy/spidermonkeywasm2go/p5.Fn7932
func Fn7932(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7933 github.com/goccy/spidermonkeywasm2go/p6.Fn7933
func Fn7933(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7934 github.com/goccy/spidermonkeywasm2go/p5.Fn7934
func Fn7934(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7935 github.com/goccy/spidermonkeywasm2go/p5.Fn7935
func Fn7935(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7936 github.com/goccy/spidermonkeywasm2go/p6.Fn7936
func Fn7936(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7937 github.com/goccy/spidermonkeywasm2go/p5.Fn7937
func Fn7937(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7940 github.com/goccy/spidermonkeywasm2go/p6.Fn7940
func Fn7940(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7941 github.com/goccy/spidermonkeywasm2go/p6.Fn7941
func Fn7941(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7942 github.com/goccy/spidermonkeywasm2go/p4.Fn7942
func Fn7942(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7943 github.com/goccy/spidermonkeywasm2go/p6.Fn7943
func Fn7943(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7946 github.com/goccy/spidermonkeywasm2go/p6.Fn7946
func Fn7946(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7947 github.com/goccy/spidermonkeywasm2go/p5.Fn7947
func Fn7947(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7948 github.com/goccy/spidermonkeywasm2go/p5.Fn7948
func Fn7948(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7949 github.com/goccy/spidermonkeywasm2go/p4.Fn7949
func Fn7949(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7950 github.com/goccy/spidermonkeywasm2go/p7.Fn7950
func Fn7950(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7965 github.com/goccy/spidermonkeywasm2go/p4.Fn7965
func Fn7965(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn7966 github.com/goccy/spidermonkeywasm2go/p7.Fn7966
func Fn7966(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7967 github.com/goccy/spidermonkeywasm2go/p5.Fn7967
func Fn7967(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7968 github.com/goccy/spidermonkeywasm2go/p6.Fn7968
func Fn7968(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7969 github.com/goccy/spidermonkeywasm2go/p4.Fn7969
func Fn7969(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn7988 github.com/goccy/spidermonkeywasm2go/p6.Fn7988
func Fn7988(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8085 github.com/goccy/spidermonkeywasm2go/p7.Fn8085
func Fn8085(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8091 github.com/goccy/spidermonkeywasm2go/p5.Fn8091
func Fn8091(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8094 github.com/goccy/spidermonkeywasm2go/p6.Fn8094
func Fn8094(m *base.Module, l0 int32)

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

//go:linkname Fn8108 github.com/goccy/spidermonkeywasm2go/p7.Fn8108
func Fn8108(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8109 github.com/goccy/spidermonkeywasm2go/p6.Fn8109
func Fn8109(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8110 github.com/goccy/spidermonkeywasm2go/p6.Fn8110
func Fn8110(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8131 github.com/goccy/spidermonkeywasm2go/p6.Fn8131
func Fn8131(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8132 github.com/goccy/spidermonkeywasm2go/p6.Fn8132
func Fn8132(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8133 github.com/goccy/spidermonkeywasm2go/p7.Fn8133
func Fn8133(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8136 github.com/goccy/spidermonkeywasm2go/p5.Fn8136
func Fn8136(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8137 github.com/goccy/spidermonkeywasm2go/p4.Fn8137
func Fn8137(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8142 github.com/goccy/spidermonkeywasm2go/p1.Fn8142
func Fn8142(m *base.Module, l0 int32) int32

//go:linkname Fn8145 github.com/goccy/spidermonkeywasm2go/p6.Fn8145
func Fn8145(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8147 github.com/goccy/spidermonkeywasm2go/p1.Fn8147
func Fn8147(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8148 github.com/goccy/spidermonkeywasm2go/p7.Fn8148
func Fn8148(m *base.Module, l0 int32) int32

//go:linkname Fn8150 github.com/goccy/spidermonkeywasm2go/p7.Fn8150
func Fn8150(m *base.Module, l0 int32) int32

//go:linkname Fn8152 github.com/goccy/spidermonkeywasm2go/p1.Fn8152
func Fn8152(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8153 github.com/goccy/spidermonkeywasm2go/p6.Fn8153
func Fn8153(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8154 github.com/goccy/spidermonkeywasm2go/p7.Fn8154
func Fn8154(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8156 github.com/goccy/spidermonkeywasm2go/p7.Fn8156
func Fn8156(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8157 github.com/goccy/spidermonkeywasm2go/p1.Fn8157
func Fn8157(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8158 github.com/goccy/spidermonkeywasm2go/p1.Fn8158
func Fn8158(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8159 github.com/goccy/spidermonkeywasm2go/p2.Fn8159
func Fn8159(m *base.Module) int32

//go:linkname Fn8163 github.com/goccy/spidermonkeywasm2go/p6.Fn8163
func Fn8163(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8165 github.com/goccy/spidermonkeywasm2go/p7.Fn8165
func Fn8165(m *base.Module, l0 int32) int32

//go:linkname Fn8177 github.com/goccy/spidermonkeywasm2go/p7.Fn8177
func Fn8177(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8182 github.com/goccy/spidermonkeywasm2go/p1.Fn8182
func Fn8182(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8183 github.com/goccy/spidermonkeywasm2go/p6.Fn8183
func Fn8183(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32)

//go:linkname Fn8185 github.com/goccy/spidermonkeywasm2go/p5.Fn8185
func Fn8185(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8187 github.com/goccy/spidermonkeywasm2go/p1.Fn8187
func Fn8187(m *base.Module, l0 int32)

//go:linkname Fn8192 github.com/goccy/spidermonkeywasm2go/p1.Fn8192
func Fn8192(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8193 github.com/goccy/spidermonkeywasm2go/p7.Fn8193
func Fn8193(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8202 github.com/goccy/spidermonkeywasm2go/p5.Fn8202
func Fn8202(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8204 github.com/goccy/spidermonkeywasm2go/p7.Fn8204
func Fn8204(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8208 github.com/goccy/spidermonkeywasm2go/p6.Fn8208
func Fn8208(m *base.Module, l0 int32) int32

//go:linkname Fn8210 github.com/goccy/spidermonkeywasm2go/p4.Fn8210
func Fn8210(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8211 github.com/goccy/spidermonkeywasm2go/p6.Fn8211
func Fn8211(m *base.Module, l0 int32) int32

//go:linkname Fn8254 github.com/goccy/spidermonkeywasm2go/p6.Fn8254
func Fn8254(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8255 github.com/goccy/spidermonkeywasm2go/p5.Fn8255
func Fn8255(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8257 github.com/goccy/spidermonkeywasm2go/p5.Fn8257
func Fn8257(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8258 github.com/goccy/spidermonkeywasm2go/p6.Fn8258
func Fn8258(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8262 github.com/goccy/spidermonkeywasm2go/p5.Fn8262
func Fn8262(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8269 github.com/goccy/spidermonkeywasm2go/p5.Fn8269
func Fn8269(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8271 github.com/goccy/spidermonkeywasm2go/p4.Fn8271
func Fn8271(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8274 github.com/goccy/spidermonkeywasm2go/p4.Fn8274
func Fn8274(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8278 github.com/goccy/spidermonkeywasm2go/p6.Fn8278
func Fn8278(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8282 github.com/goccy/spidermonkeywasm2go/p5.Fn8282
func Fn8282(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8285 github.com/goccy/spidermonkeywasm2go/p6.Fn8285
func Fn8285(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8286 github.com/goccy/spidermonkeywasm2go/p6.Fn8286
func Fn8286(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8287 github.com/goccy/spidermonkeywasm2go/p7.Fn8287
func Fn8287(m *base.Module) float64

//go:linkname Fn8289 github.com/goccy/spidermonkeywasm2go/p7.Fn8289
func Fn8289(m *base.Module, l0 int32) int32

//go:linkname Fn8290 github.com/goccy/spidermonkeywasm2go/p5.Fn8290
func Fn8290(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8291 github.com/goccy/spidermonkeywasm2go/p6.Fn8291
func Fn8291(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8293 github.com/goccy/spidermonkeywasm2go/p6.Fn8293
func Fn8293(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8296 github.com/goccy/spidermonkeywasm2go/p5.Fn8296
func Fn8296(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8299 github.com/goccy/spidermonkeywasm2go/p6.Fn8299
func Fn8299(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8300 github.com/goccy/spidermonkeywasm2go/p5.Fn8300
func Fn8300(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8301 github.com/goccy/spidermonkeywasm2go/p6.Fn8301
func Fn8301(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8308 github.com/goccy/spidermonkeywasm2go/p7.Fn8308
func Fn8308(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8309 github.com/goccy/spidermonkeywasm2go/p7.Fn8309
func Fn8309(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8310 github.com/goccy/spidermonkeywasm2go/p6.Fn8310
func Fn8310(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8328 github.com/goccy/spidermonkeywasm2go/p4.Fn8328
func Fn8328(m *base.Module, l0 int32) int32

//go:linkname Fn8330 github.com/goccy/spidermonkeywasm2go/p7.Fn8330
func Fn8330(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8334 github.com/goccy/spidermonkeywasm2go/p7.Fn8334
func Fn8334(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8337 github.com/goccy/spidermonkeywasm2go/p6.Fn8337
func Fn8337(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8341 github.com/goccy/spidermonkeywasm2go/p5.Fn8341
func Fn8341(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn8342 github.com/goccy/spidermonkeywasm2go/p4.Fn8342
func Fn8342(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8350 github.com/goccy/spidermonkeywasm2go/p6.Fn8350
func Fn8350(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8353 github.com/goccy/spidermonkeywasm2go/p6.Fn8353
func Fn8353(m *base.Module, l0 int32)

//go:linkname Fn8355 github.com/goccy/spidermonkeywasm2go/p6.Fn8355
func Fn8355(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8356 github.com/goccy/spidermonkeywasm2go/p5.Fn8356
func Fn8356(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8358 github.com/goccy/spidermonkeywasm2go/p6.Fn8358
func Fn8358(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8369 github.com/goccy/spidermonkeywasm2go/p5.Fn8369
func Fn8369(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn8378 github.com/goccy/spidermonkeywasm2go/p1.Fn8378
func Fn8378(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8379 github.com/goccy/spidermonkeywasm2go/p6.Fn8379
func Fn8379(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8384 github.com/goccy/spidermonkeywasm2go/p7.Fn8384
func Fn8384(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8385 github.com/goccy/spidermonkeywasm2go/p4.Fn8385
func Fn8385(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8390 github.com/goccy/spidermonkeywasm2go/p1.Fn8390
func Fn8390(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8403 github.com/goccy/spidermonkeywasm2go/p6.Fn8403
func Fn8403(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8422 github.com/goccy/spidermonkeywasm2go/p6.Fn8422
func Fn8422(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8425 github.com/goccy/spidermonkeywasm2go/p6.Fn8425
func Fn8425(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8426 github.com/goccy/spidermonkeywasm2go/p4.Fn8426
func Fn8426(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8430 github.com/goccy/spidermonkeywasm2go/p6.Fn8430
func Fn8430(m *base.Module, l0 int32) int32

//go:linkname Fn8444 github.com/goccy/spidermonkeywasm2go/p6.Fn8444
func Fn8444(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8449 github.com/goccy/spidermonkeywasm2go/p6.Fn8449
func Fn8449(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8451 github.com/goccy/spidermonkeywasm2go/p6.Fn8451
func Fn8451(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8452 github.com/goccy/spidermonkeywasm2go/p6.Fn8452
func Fn8452(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8453 github.com/goccy/spidermonkeywasm2go/p5.Fn8453
func Fn8453(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8462 github.com/goccy/spidermonkeywasm2go/p1.Fn8462
func Fn8462(m *base.Module, l0 int32) int32

//go:linkname Fn8478 github.com/goccy/spidermonkeywasm2go/p5.Fn8478
func Fn8478(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8485 github.com/goccy/spidermonkeywasm2go/p7.Fn8485
func Fn8485(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8486 github.com/goccy/spidermonkeywasm2go/p7.Fn8486
func Fn8486(m *base.Module, l0 int32)

//go:linkname Fn8487 github.com/goccy/spidermonkeywasm2go/p5.Fn8487
func Fn8487(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8489 github.com/goccy/spidermonkeywasm2go/p6.Fn8489
func Fn8489(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8509 github.com/goccy/spidermonkeywasm2go/p6.Fn8509
func Fn8509(m *base.Module, l0 int32, l1 int32, l2 int32)

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

//go:linkname Fn8520 github.com/goccy/spidermonkeywasm2go/p5.Fn8520
func Fn8520(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8522 github.com/goccy/spidermonkeywasm2go/p7.Fn8522
func Fn8522(m *base.Module, l0 int32) int32

//go:linkname Fn8524 github.com/goccy/spidermonkeywasm2go/p7.Fn8524
func Fn8524(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8527 github.com/goccy/spidermonkeywasm2go/p6.Fn8527
func Fn8527(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8530 github.com/goccy/spidermonkeywasm2go/p7.Fn8530
func Fn8530(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8531 github.com/goccy/spidermonkeywasm2go/p6.Fn8531
func Fn8531(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8536 github.com/goccy/spidermonkeywasm2go/p6.Fn8536
func Fn8536(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8538 github.com/goccy/spidermonkeywasm2go/p5.Fn8538
func Fn8538(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8539 github.com/goccy/spidermonkeywasm2go/p6.Fn8539
func Fn8539(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8540 github.com/goccy/spidermonkeywasm2go/p6.Fn8540
func Fn8540(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8541 github.com/goccy/spidermonkeywasm2go/p6.Fn8541
func Fn8541(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8545 github.com/goccy/spidermonkeywasm2go/p7.Fn8545
func Fn8545(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8549 github.com/goccy/spidermonkeywasm2go/p6.Fn8549
func Fn8549(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8558 github.com/goccy/spidermonkeywasm2go/p4.Fn8558
func Fn8558(m *base.Module, l0 int32) int32

//go:linkname Fn8561 github.com/goccy/spidermonkeywasm2go/p7.Fn8561
func Fn8561(m *base.Module, l0 int32) int32

//go:linkname Fn8613 github.com/goccy/spidermonkeywasm2go/p7.Fn8613
func Fn8613(m *base.Module, l0 int32)

//go:linkname Fn8614 github.com/goccy/spidermonkeywasm2go/p7.Fn8614
func Fn8614(m *base.Module, l0 int32)

//go:linkname Fn8617 github.com/goccy/spidermonkeywasm2go/p6.Fn8617
func Fn8617(m *base.Module, l0 int32, l1 int32, l2 int32) int32

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

//go:linkname Fn8628 github.com/goccy/spidermonkeywasm2go/p6.Fn8628
func Fn8628(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8630 github.com/goccy/spidermonkeywasm2go/p6.Fn8630
func Fn8630(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8631 github.com/goccy/spidermonkeywasm2go/p1.Fn8631
func Fn8631(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8633 github.com/goccy/spidermonkeywasm2go/p1.Fn8633
func Fn8633(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8642 github.com/goccy/spidermonkeywasm2go/p2.Fn8642
func Fn8642(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn8673 github.com/goccy/spidermonkeywasm2go/p6.Fn8673
func Fn8673(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn8679 github.com/goccy/spidermonkeywasm2go/p4.Fn8679
func Fn8679(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8680 github.com/goccy/spidermonkeywasm2go/p6.Fn8680
func Fn8680(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8681 github.com/goccy/spidermonkeywasm2go/p6.Fn8681
func Fn8681(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8693 github.com/goccy/spidermonkeywasm2go/p4.Fn8693
func Fn8693(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8698 github.com/goccy/spidermonkeywasm2go/p5.Fn8698
func Fn8698(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8699 github.com/goccy/spidermonkeywasm2go/p5.Fn8699
func Fn8699(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8704 github.com/goccy/spidermonkeywasm2go/p6.Fn8704
func Fn8704(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8711 github.com/goccy/spidermonkeywasm2go/p7.Fn8711
func Fn8711(m *base.Module, l0 int32)

//go:linkname Fn8714 github.com/goccy/spidermonkeywasm2go/p6.Fn8714
func Fn8714(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8717 github.com/goccy/spidermonkeywasm2go/p6.Fn8717
func Fn8717(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8742 github.com/goccy/spidermonkeywasm2go/p5.Fn8742
func Fn8742(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8743 github.com/goccy/spidermonkeywasm2go/p7.Fn8743
func Fn8743(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8746 github.com/goccy/spidermonkeywasm2go/p6.Fn8746
func Fn8746(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8749 github.com/goccy/spidermonkeywasm2go/p7.Fn8749
func Fn8749(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8821 github.com/goccy/spidermonkeywasm2go/p7.Fn8821
func Fn8821(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8856 github.com/goccy/spidermonkeywasm2go/p6.Fn8856
func Fn8856(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8902 github.com/goccy/spidermonkeywasm2go/p6.Fn8902
func Fn8902(m *base.Module, l0 int32) int32

//go:linkname Fn8982 github.com/goccy/spidermonkeywasm2go/p7.Fn8982
func Fn8982(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8998 github.com/goccy/spidermonkeywasm2go/p4.Fn8998
func Fn8998(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9013 github.com/goccy/spidermonkeywasm2go/p7.Fn9013
func Fn9013(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9014 github.com/goccy/spidermonkeywasm2go/p4.Fn9014
func Fn9014(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9016 github.com/goccy/spidermonkeywasm2go/p7.Fn9016
func Fn9016(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9042 github.com/goccy/spidermonkeywasm2go/p6.Fn9042
func Fn9042(m *base.Module, l0 int32)

//go:linkname Fn9048 github.com/goccy/spidermonkeywasm2go/p6.Fn9048
func Fn9048(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9056 github.com/goccy/spidermonkeywasm2go/p2.Fn9056
func Fn9056(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9063 github.com/goccy/spidermonkeywasm2go/p4.Fn9063
func Fn9063(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9064 github.com/goccy/spidermonkeywasm2go/p6.Fn9064
func Fn9064(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9066 github.com/goccy/spidermonkeywasm2go/p5.Fn9066
func Fn9066(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9068 github.com/goccy/spidermonkeywasm2go/p6.Fn9068
func Fn9068(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9072 github.com/goccy/spidermonkeywasm2go/p5.Fn9072
func Fn9072(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn9074 github.com/goccy/spidermonkeywasm2go/p4.Fn9074
func Fn9074(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn9086 github.com/goccy/spidermonkeywasm2go/p6.Fn9086
func Fn9086(m *base.Module, l0 int32) int32

//go:linkname Fn9088 github.com/goccy/spidermonkeywasm2go/p2.Fn9088
func Fn9088(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9091 github.com/goccy/spidermonkeywasm2go/p5.Fn9091
func Fn9091(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9113 github.com/goccy/spidermonkeywasm2go/p6.Fn9113
func Fn9113(m *base.Module, l0 int32)

//go:linkname Fn9114 github.com/goccy/spidermonkeywasm2go/p5.Fn9114
func Fn9114(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9117 github.com/goccy/spidermonkeywasm2go/p5.Fn9117
func Fn9117(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9121 github.com/goccy/spidermonkeywasm2go/p6.Fn9121
func Fn9121(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9123 github.com/goccy/spidermonkeywasm2go/p6.Fn9123
func Fn9123(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9128 github.com/goccy/spidermonkeywasm2go/p6.Fn9128
func Fn9128(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9164 github.com/goccy/spidermonkeywasm2go/p7.Fn9164
func Fn9164(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn9165 github.com/goccy/spidermonkeywasm2go/p6.Fn9165
func Fn9165(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn9261 github.com/goccy/spidermonkeywasm2go/p6.Fn9261
func Fn9261(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9278 github.com/goccy/spidermonkeywasm2go/p6.Fn9278
func Fn9278(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9279 github.com/goccy/spidermonkeywasm2go/p6.Fn9279
func Fn9279(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9280 github.com/goccy/spidermonkeywasm2go/p5.Fn9280
func Fn9280(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9281 github.com/goccy/spidermonkeywasm2go/p6.Fn9281
func Fn9281(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn9283 github.com/goccy/spidermonkeywasm2go/p6.Fn9283
func Fn9283(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

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

//go:linkname Fn9296 github.com/goccy/spidermonkeywasm2go/p6.Fn9296
func Fn9296(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9297 github.com/goccy/spidermonkeywasm2go/p6.Fn9297
func Fn9297(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9298 github.com/goccy/spidermonkeywasm2go/p5.Fn9298
func Fn9298(m *base.Module, l0 int32)

//go:linkname Fn9299 github.com/goccy/spidermonkeywasm2go/p5.Fn9299
func Fn9299(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn9300 github.com/goccy/spidermonkeywasm2go/p5.Fn9300
func Fn9300(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9304 github.com/goccy/spidermonkeywasm2go/p6.Fn9304
func Fn9304(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9306 github.com/goccy/spidermonkeywasm2go/p5.Fn9306
func Fn9306(m *base.Module, l0 int32) float64

//go:linkname Fn9311 github.com/goccy/spidermonkeywasm2go/p6.Fn9311
func Fn9311(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn9312 github.com/goccy/spidermonkeywasm2go/p6.Fn9312
func Fn9312(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9313 github.com/goccy/spidermonkeywasm2go/p6.Fn9313
func Fn9313(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9315 github.com/goccy/spidermonkeywasm2go/p7.Fn9315
func Fn9315(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9317 github.com/goccy/spidermonkeywasm2go/p6.Fn9317
func Fn9317(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9318 github.com/goccy/spidermonkeywasm2go/p6.Fn9318
func Fn9318(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9320 github.com/goccy/spidermonkeywasm2go/p5.Fn9320
func Fn9320(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9324 github.com/goccy/spidermonkeywasm2go/p6.Fn9324
func Fn9324(m *base.Module, l0 int32) int32

//go:linkname Fn9326 github.com/goccy/spidermonkeywasm2go/p5.Fn9326
func Fn9326(m *base.Module, l0 int32) int32

//go:linkname Fn9328 github.com/goccy/spidermonkeywasm2go/p5.Fn9328
func Fn9328(m *base.Module, l0 int32)

//go:linkname Fn9334 github.com/goccy/spidermonkeywasm2go/p1.Fn9334
func Fn9334(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9335 github.com/goccy/spidermonkeywasm2go/p4.Fn9335
func Fn9335(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9344 github.com/goccy/spidermonkeywasm2go/p4.Fn9344
func Fn9344(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9347 github.com/goccy/spidermonkeywasm2go/p7.Fn9347
func Fn9347(m *base.Module, l0 int32)

//go:linkname Fn9369 github.com/goccy/spidermonkeywasm2go/p6.Fn9369
func Fn9369(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9370 github.com/goccy/spidermonkeywasm2go/p6.Fn9370
func Fn9370(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9425 github.com/goccy/spidermonkeywasm2go/p5.Fn9425
func Fn9425(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9475 github.com/goccy/spidermonkeywasm2go/p4.Fn9475
func Fn9475(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9486 github.com/goccy/spidermonkeywasm2go/p5.Fn9486
func Fn9486(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9487 github.com/goccy/spidermonkeywasm2go/p5.Fn9487
func Fn9487(m *base.Module, l0 int32)

//go:linkname Fn9503 github.com/goccy/spidermonkeywasm2go/p6.Fn9503
func Fn9503(m *base.Module, l0 int32) int32

//go:linkname Fn9507 github.com/goccy/spidermonkeywasm2go/p5.Fn9507
func Fn9507(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9508 github.com/goccy/spidermonkeywasm2go/p7.Fn9508
func Fn9508(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9510 github.com/goccy/spidermonkeywasm2go/p6.Fn9510
func Fn9510(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9511 github.com/goccy/spidermonkeywasm2go/p6.Fn9511
func Fn9511(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9513 github.com/goccy/spidermonkeywasm2go/p6.Fn9513
func Fn9513(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9515 github.com/goccy/spidermonkeywasm2go/p5.Fn9515
func Fn9515(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn9521 github.com/goccy/spidermonkeywasm2go/p5.Fn9521
func Fn9521(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9524 github.com/goccy/spidermonkeywasm2go/p6.Fn9524
func Fn9524(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9540 github.com/goccy/spidermonkeywasm2go/p5.Fn9540
func Fn9540(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9559 github.com/goccy/spidermonkeywasm2go/p7.Fn9559
func Fn9559(m *base.Module, l0 int32) int32

//go:linkname Fn9563 github.com/goccy/spidermonkeywasm2go/p6.Fn9563
func Fn9563(m *base.Module, l0 int32) int32

//go:linkname Fn9565 github.com/goccy/spidermonkeywasm2go/p5.Fn9565
func Fn9565(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn9595 github.com/goccy/spidermonkeywasm2go/p6.Fn9595
func Fn9595(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9608 github.com/goccy/spidermonkeywasm2go/p6.Fn9608
func Fn9608(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9609 github.com/goccy/spidermonkeywasm2go/p6.Fn9609
func Fn9609(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9614 github.com/goccy/spidermonkeywasm2go/p5.Fn9614
func Fn9614(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn9615 github.com/goccy/spidermonkeywasm2go/p7.Fn9615
func Fn9615(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9617 github.com/goccy/spidermonkeywasm2go/p5.Fn9617
func Fn9617(m *base.Module, l0 int32)

//go:linkname Fn9623 github.com/goccy/spidermonkeywasm2go/p5.Fn9623
func Fn9623(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9634 github.com/goccy/spidermonkeywasm2go/p1.Fn9634
func Fn9634(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9642 github.com/goccy/spidermonkeywasm2go/p6.Fn9642
func Fn9642(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9661 github.com/goccy/spidermonkeywasm2go/p7.Fn9661
func Fn9661(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9674 github.com/goccy/spidermonkeywasm2go/p6.Fn9674
func Fn9674(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9679 github.com/goccy/spidermonkeywasm2go/p4.Fn9679
func Fn9679(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn9680 github.com/goccy/spidermonkeywasm2go/p4.Fn9680
func Fn9680(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9684 github.com/goccy/spidermonkeywasm2go/p6.Fn9684
func Fn9684(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9685 github.com/goccy/spidermonkeywasm2go/p5.Fn9685
func Fn9685(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn9687 github.com/goccy/spidermonkeywasm2go/p7.Fn9687
func Fn9687(m *base.Module, l0 int32)

//go:linkname Fn9689 github.com/goccy/spidermonkeywasm2go/p5.Fn9689
func Fn9689(m *base.Module, l0 int32, l1 float64, l2 float64, l3 int32, l4 int32)

//go:linkname Fn9690 github.com/goccy/spidermonkeywasm2go/p6.Fn9690
func Fn9690(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn9691 github.com/goccy/spidermonkeywasm2go/p7.Fn9691
func Fn9691(m *base.Module, l0 int32)

//go:linkname Fn9693 github.com/goccy/spidermonkeywasm2go/p7.Fn9693
func Fn9693(m *base.Module, l0 int32)

//go:linkname Fn9698 github.com/goccy/spidermonkeywasm2go/p4.Fn9698
func Fn9698(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9705 github.com/goccy/spidermonkeywasm2go/p6.Fn9705
func Fn9705(m *base.Module, l0 int32)

//go:linkname Fn9746 github.com/goccy/spidermonkeywasm2go/p6.Fn9746
func Fn9746(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9755 github.com/goccy/spidermonkeywasm2go/p6.Fn9755
func Fn9755(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9757 github.com/goccy/spidermonkeywasm2go/p7.Fn9757
func Fn9757(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9758 github.com/goccy/spidermonkeywasm2go/p7.Fn9758
func Fn9758(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9759 github.com/goccy/spidermonkeywasm2go/p6.Fn9759
func Fn9759(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9760 github.com/goccy/spidermonkeywasm2go/p5.Fn9760
func Fn9760(m *base.Module, l0 int32, l1 int32, l2 float64)

//go:linkname Fn9761 github.com/goccy/spidermonkeywasm2go/p7.Fn9761
func Fn9761(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9840 github.com/goccy/spidermonkeywasm2go/p6.Fn9840
func Fn9840(m *base.Module, l0 int32) int32

//go:linkname Fn9841 github.com/goccy/spidermonkeywasm2go/p2.Fn9841
func Fn9841(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9843 github.com/goccy/spidermonkeywasm2go/p7.Fn9843
func Fn9843(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9932 github.com/goccy/spidermonkeywasm2go/p7.Fn9932
func Fn9932(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9934 github.com/goccy/spidermonkeywasm2go/p7.Fn9934
func Fn9934(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9946 github.com/goccy/spidermonkeywasm2go/p7.Fn9946
func Fn9946(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9947 github.com/goccy/spidermonkeywasm2go/p7.Fn9947
func Fn9947(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9948 github.com/goccy/spidermonkeywasm2go/p6.Fn9948
func Fn9948(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9961 github.com/goccy/spidermonkeywasm2go/p5.Fn9961
func Fn9961(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9987 github.com/goccy/spidermonkeywasm2go/p2.Fn9987
func Fn9987(m *base.Module, l0 int32) int32

//go:linkname Fn10014 github.com/goccy/spidermonkeywasm2go/p6.Fn10014
func Fn10014(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10017 github.com/goccy/spidermonkeywasm2go/p5.Fn10017
func Fn10017(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10081 github.com/goccy/spidermonkeywasm2go/p2.Fn10081
func Fn10081(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn10097 github.com/goccy/spidermonkeywasm2go/p5.Fn10097
func Fn10097(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn10100 github.com/goccy/spidermonkeywasm2go/p6.Fn10100
func Fn10100(m *base.Module, l0 int32) int32

//go:linkname Fn10107 github.com/goccy/spidermonkeywasm2go/p6.Fn10107
func Fn10107(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10108 github.com/goccy/spidermonkeywasm2go/p7.Fn10108
func Fn10108(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn10119 github.com/goccy/spidermonkeywasm2go/p5.Fn10119
func Fn10119(m *base.Module, l0 int32) int32

//go:linkname Fn10124 github.com/goccy/spidermonkeywasm2go/p5.Fn10124
func Fn10124(m *base.Module, l0 int32) int32

//go:linkname Fn10128 github.com/goccy/spidermonkeywasm2go/p6.Fn10128
func Fn10128(m *base.Module, l0 int32)

//go:linkname Fn10150 github.com/goccy/spidermonkeywasm2go/p2.Fn10150
func Fn10150(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn10188 github.com/goccy/spidermonkeywasm2go/p5.Fn10188
func Fn10188(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10194 github.com/goccy/spidermonkeywasm2go/p7.Fn10194
func Fn10194(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10197 github.com/goccy/spidermonkeywasm2go/p6.Fn10197
func Fn10197(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn10199 github.com/goccy/spidermonkeywasm2go/p5.Fn10199
func Fn10199(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10200 github.com/goccy/spidermonkeywasm2go/p5.Fn10200
func Fn10200(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10207 github.com/goccy/spidermonkeywasm2go/p5.Fn10207
func Fn10207(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10212 github.com/goccy/spidermonkeywasm2go/p6.Fn10212
func Fn10212(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10213 github.com/goccy/spidermonkeywasm2go/p7.Fn10213
func Fn10213(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10214 github.com/goccy/spidermonkeywasm2go/p4.Fn10214
func Fn10214(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10226 github.com/goccy/spidermonkeywasm2go/p2.Fn10226
func Fn10226(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10227 github.com/goccy/spidermonkeywasm2go/p6.Fn10227
func Fn10227(m *base.Module, l0 int32)

//go:linkname Fn10232 github.com/goccy/spidermonkeywasm2go/p6.Fn10232
func Fn10232(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10237 github.com/goccy/spidermonkeywasm2go/p5.Fn10237
func Fn10237(m *base.Module, l0 int32) int32

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

//go:linkname Fn10328 github.com/goccy/spidermonkeywasm2go/p1.Fn10328
func Fn10328(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10330 github.com/goccy/spidermonkeywasm2go/p5.Fn10330
func Fn10330(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10338 github.com/goccy/spidermonkeywasm2go/p6.Fn10338
func Fn10338(m *base.Module, l0 int32) int32

//go:linkname Fn10344 github.com/goccy/spidermonkeywasm2go/p5.Fn10344
func Fn10344(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10389 github.com/goccy/spidermonkeywasm2go/p6.Fn10389
func Fn10389(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10408 github.com/goccy/spidermonkeywasm2go/p4.Fn10408
func Fn10408(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10419 github.com/goccy/spidermonkeywasm2go/p4.Fn10419
func Fn10419(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10428 github.com/goccy/spidermonkeywasm2go/p7.Fn10428
func Fn10428(m *base.Module, l0 int32)

//go:linkname Fn10450 github.com/goccy/spidermonkeywasm2go/p5.Fn10450
func Fn10450(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10459 github.com/goccy/spidermonkeywasm2go/p5.Fn10459
func Fn10459(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10491 github.com/goccy/spidermonkeywasm2go/p6.Fn10491
func Fn10491(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10508 github.com/goccy/spidermonkeywasm2go/p4.Fn10508
func Fn10508(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10509 github.com/goccy/spidermonkeywasm2go/p6.Fn10509
func Fn10509(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10548 github.com/goccy/spidermonkeywasm2go/p5.Fn10548
func Fn10548(m *base.Module, l0 int32) int32

//go:linkname Fn10566 github.com/goccy/spidermonkeywasm2go/p7.Fn10566
func Fn10566(m *base.Module, l0 int32) int32

//go:linkname Fn10571 github.com/goccy/spidermonkeywasm2go/p5.Fn10571
func Fn10571(m *base.Module) int32

//go:linkname Fn10572 github.com/goccy/spidermonkeywasm2go/p7.Fn10572
func Fn10572(m *base.Module) int32

//go:linkname Fn10573 github.com/goccy/spidermonkeywasm2go/p6.Fn10573
func Fn10573(m *base.Module, l0 int32) int32

//go:linkname Fn10578 github.com/goccy/spidermonkeywasm2go/p4.Fn10578
func Fn10578(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10581 github.com/goccy/spidermonkeywasm2go/p6.Fn10581
func Fn10581(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10621 github.com/goccy/spidermonkeywasm2go/p6.Fn10621
func Fn10621(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10628 github.com/goccy/spidermonkeywasm2go/p7.Fn10628
func Fn10628(m *base.Module, l0 int32)

//go:linkname Fn10629 github.com/goccy/spidermonkeywasm2go/p5.Fn10629
func Fn10629(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10633 github.com/goccy/spidermonkeywasm2go/p4.Fn10633
func Fn10633(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10635 github.com/goccy/spidermonkeywasm2go/p7.Fn10635
func Fn10635(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10645 github.com/goccy/spidermonkeywasm2go/p5.Fn10645
func Fn10645(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10746 github.com/goccy/spidermonkeywasm2go/p6.Fn10746
func Fn10746(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10765 github.com/goccy/spidermonkeywasm2go/p6.Fn10765
func Fn10765(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10777 github.com/goccy/spidermonkeywasm2go/p4.Fn10777
func Fn10777(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn10781 github.com/goccy/spidermonkeywasm2go/p5.Fn10781
func Fn10781(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn10812 github.com/goccy/spidermonkeywasm2go/p5.Fn10812
func Fn10812(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10814 github.com/goccy/spidermonkeywasm2go/p6.Fn10814
func Fn10814(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10815 github.com/goccy/spidermonkeywasm2go/p4.Fn10815
func Fn10815(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10818 github.com/goccy/spidermonkeywasm2go/p6.Fn10818
func Fn10818(m *base.Module, l0 int32, l1 float64) float64

//go:linkname Fn10819 github.com/goccy/spidermonkeywasm2go/p6.Fn10819
func Fn10819(m *base.Module, l0 int32, l1 float64) float64

//go:linkname Fn10870 github.com/goccy/spidermonkeywasm2go/p4.Fn10870
func Fn10870(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10872 github.com/goccy/spidermonkeywasm2go/p7.Fn10872
func Fn10872(m *base.Module, l0 int32) int32

//go:linkname Fn10905 github.com/goccy/spidermonkeywasm2go/p5.Fn10905
func Fn10905(m *base.Module, l0 int32)

//go:linkname Fn10909 github.com/goccy/spidermonkeywasm2go/p6.Fn10909
func Fn10909(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10916 github.com/goccy/spidermonkeywasm2go/p6.Fn10916
func Fn10916(m *base.Module, l0 int32)

//go:linkname Fn10917 github.com/goccy/spidermonkeywasm2go/p6.Fn10917
func Fn10917(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10921 github.com/goccy/spidermonkeywasm2go/p5.Fn10921
func Fn10921(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10922 github.com/goccy/spidermonkeywasm2go/p4.Fn10922
func Fn10922(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10924 github.com/goccy/spidermonkeywasm2go/p7.Fn10924
func Fn10924(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10926 github.com/goccy/spidermonkeywasm2go/p7.Fn10926
func Fn10926(m *base.Module)

//go:linkname Fn10928 github.com/goccy/spidermonkeywasm2go/p2.Fn10928
func Fn10928(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10929 github.com/goccy/spidermonkeywasm2go/p2.Fn10929
func Fn10929(m *base.Module, l0 float64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn10934 github.com/goccy/spidermonkeywasm2go/p4.Fn10934
func Fn10934(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10935 github.com/goccy/spidermonkeywasm2go/p5.Fn10935
func Fn10935(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10957 github.com/goccy/spidermonkeywasm2go/p2.Fn10957
func Fn10957(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10961 github.com/goccy/spidermonkeywasm2go/p7.Fn10961
func Fn10961(m *base.Module, l0 int32)

//go:linkname Fn10962 github.com/goccy/spidermonkeywasm2go/p7.Fn10962
func Fn10962(m *base.Module, l0 int32)

//go:linkname Fn10963 github.com/goccy/spidermonkeywasm2go/p7.Fn10963
func Fn10963(m *base.Module, l0 int32)
