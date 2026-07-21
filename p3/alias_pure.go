//go:build !amd64 && !arm64

package p3

import (
	base "github.com/goccy/spidermonkeywasm2go/base"
	_ "unsafe"
)

//go:linkname Fn36 github.com/goccy/spidermonkeywasm2go/p7.Fn36
func Fn36(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn102 github.com/goccy/spidermonkeywasm2go/p7.Fn102
func Fn102(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn103 github.com/goccy/spidermonkeywasm2go/p7.Fn103
func Fn103(m *base.Module, l0 int64)

//go:linkname Fn111 github.com/goccy/spidermonkeywasm2go/p7.Fn111
func Fn111(m *base.Module, l0 int32)

//go:linkname Fn116 github.com/goccy/spidermonkeywasm2go/p7.Fn116
func Fn116(m *base.Module, l0 int32) int32

//go:linkname Fn117 github.com/goccy/spidermonkeywasm2go/p6.Fn117
func Fn117(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn118 github.com/goccy/spidermonkeywasm2go/p7.Fn118
func Fn118(m *base.Module, l0 int32)

//go:linkname Fn124 github.com/goccy/spidermonkeywasm2go/p7.Fn124
func Fn124(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn125 github.com/goccy/spidermonkeywasm2go/p7.Fn125
func Fn125(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn126 github.com/goccy/spidermonkeywasm2go/p7.Fn126
func Fn126(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn127 github.com/goccy/spidermonkeywasm2go/p7.Fn127
func Fn127(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn145 github.com/goccy/spidermonkeywasm2go/p7.Fn145
func Fn145(m *base.Module, l0 int32)

//go:linkname Fn146 github.com/goccy/spidermonkeywasm2go/p7.Fn146
func Fn146(m *base.Module)

//go:linkname Fn148 github.com/goccy/spidermonkeywasm2go/p0.Fn148
func Fn148(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn166 github.com/goccy/spidermonkeywasm2go/p7.Fn166
func Fn166(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn169 github.com/goccy/spidermonkeywasm2go/p6.Fn169
func Fn169(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn173 github.com/goccy/spidermonkeywasm2go/p4.Fn173
func Fn173(m *base.Module, l0 float64) float64

//go:linkname Fn175 github.com/goccy/spidermonkeywasm2go/p5.Fn175
func Fn175(m *base.Module, l0 float64) float64

//go:linkname Fn176 github.com/goccy/spidermonkeywasm2go/p5.Fn176
func Fn176(m *base.Module, l0 float64) float64

//go:linkname Fn177 github.com/goccy/spidermonkeywasm2go/p5.Fn177
func Fn177(m *base.Module, l0 float64) int32

//go:linkname Fn179 github.com/goccy/spidermonkeywasm2go/p4.Fn179
func Fn179(m *base.Module, l0 int32) float64

//go:linkname Fn180 github.com/goccy/spidermonkeywasm2go/p5.Fn180
func Fn180(m *base.Module, l0 int32) int64

//go:linkname Fn181 github.com/goccy/spidermonkeywasm2go/p5.Fn181
func Fn181(m *base.Module, l0 float64) float64

//go:linkname Fn183 github.com/goccy/spidermonkeywasm2go/p6.Fn183
func Fn183(m *base.Module, l0 float64, l1 float64, l2 float64) float64

//go:linkname Fn184 github.com/goccy/spidermonkeywasm2go/p6.Fn184
func Fn184(m *base.Module, l0 float64) float64

//go:linkname Fn185 github.com/goccy/spidermonkeywasm2go/p5.Fn185
func Fn185(m *base.Module, l0 float64) float64

//go:linkname Fn187 github.com/goccy/spidermonkeywasm2go/p6.Fn187
func Fn187(m *base.Module, l0 float64) float64

//go:linkname Fn188 github.com/goccy/spidermonkeywasm2go/p6.Fn188
func Fn188(m *base.Module, l0 float64) float64

//go:linkname Fn189 github.com/goccy/spidermonkeywasm2go/p6.Fn189
func Fn189(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn190 github.com/goccy/spidermonkeywasm2go/p6.Fn190
func Fn190(m *base.Module, l0 int32, l1 int32, l2 int32) int64

//go:linkname Fn192 github.com/goccy/spidermonkeywasm2go/p7.Fn192
func Fn192(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn205 github.com/goccy/spidermonkeywasm2go/p7.Fn205
func Fn205(m *base.Module, l0 int32)

//go:linkname Fn209 github.com/goccy/spidermonkeywasm2go/p7.Fn209
func Fn209(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn213 github.com/goccy/spidermonkeywasm2go/p6.Fn213
func Fn213(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn215 github.com/goccy/spidermonkeywasm2go/p6.Fn215
func Fn215(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn221 github.com/goccy/spidermonkeywasm2go/p6.Fn221
func Fn221(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn224 github.com/goccy/spidermonkeywasm2go/p7.Fn224
func Fn224(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn231 github.com/goccy/spidermonkeywasm2go/p4.Fn231
func Fn231(m *base.Module, l0 int64) int64

//go:linkname Fn235 github.com/goccy/spidermonkeywasm2go/p5.Fn235
func Fn235(m *base.Module, l0 int64) int64

//go:linkname Fn239 github.com/goccy/spidermonkeywasm2go/p5.Fn239
func Fn239(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn240 github.com/goccy/spidermonkeywasm2go/p6.Fn240
func Fn240(m *base.Module, l0 int32) int64

//go:linkname Fn241 github.com/goccy/spidermonkeywasm2go/p7.Fn241
func Fn241(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn242 github.com/goccy/spidermonkeywasm2go/p7.Fn242
func Fn242(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn243 github.com/goccy/spidermonkeywasm2go/p7.Fn243
func Fn243(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn249 github.com/goccy/spidermonkeywasm2go/p6.Fn249
func Fn249(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn252 github.com/goccy/spidermonkeywasm2go/p7.Fn252
func Fn252(m *base.Module, l0 int32) int32

//go:linkname Fn253 github.com/goccy/spidermonkeywasm2go/p7.Fn253
func Fn253(m *base.Module, l0 int32)

//go:linkname Fn259 github.com/goccy/spidermonkeywasm2go/p6.Fn259
func Fn259(m *base.Module, l0 int32) int32

//go:linkname Fn260 github.com/goccy/spidermonkeywasm2go/p7.Fn260
func Fn260(m *base.Module, l0 int32) int32

//go:linkname Fn261 github.com/goccy/spidermonkeywasm2go/p7.Fn261
func Fn261(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn265 github.com/goccy/spidermonkeywasm2go/p6.Fn265
func Fn265(m *base.Module, l0 int32) int32

//go:linkname Fn268 github.com/goccy/spidermonkeywasm2go/p7.Fn268
func Fn268(m *base.Module, l0 int32)

//go:linkname Fn273 github.com/goccy/spidermonkeywasm2go/p6.Fn273
func Fn273(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn277 github.com/goccy/spidermonkeywasm2go/p7.Fn277
func Fn277(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn278 github.com/goccy/spidermonkeywasm2go/p7.Fn278
func Fn278(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn280 github.com/goccy/spidermonkeywasm2go/p6.Fn280
func Fn280(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn285 github.com/goccy/spidermonkeywasm2go/p7.Fn285
func Fn285(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn286 github.com/goccy/spidermonkeywasm2go/p7.Fn286
func Fn286(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn287 github.com/goccy/spidermonkeywasm2go/p7.Fn287
func Fn287(m *base.Module, l0 int32) int32

//go:linkname Fn288 github.com/goccy/spidermonkeywasm2go/p7.Fn288
func Fn288(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn290 github.com/goccy/spidermonkeywasm2go/p7.Fn290
func Fn290(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn298 github.com/goccy/spidermonkeywasm2go/p6.Fn298
func Fn298(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn302 github.com/goccy/spidermonkeywasm2go/p5.Fn302
func Fn302(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn306 github.com/goccy/spidermonkeywasm2go/p6.Fn306
func Fn306(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float32

//go:linkname Fn320 github.com/goccy/spidermonkeywasm2go/p7.Fn320
func Fn320(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn322 github.com/goccy/spidermonkeywasm2go/p7.Fn322
func Fn322(m *base.Module)

//go:linkname Fn347 github.com/goccy/spidermonkeywasm2go/p7.Fn347
func Fn347(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn348 github.com/goccy/spidermonkeywasm2go/p7.Fn348
func Fn348(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn349 github.com/goccy/spidermonkeywasm2go/p6.Fn349
func Fn349(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn352 github.com/goccy/spidermonkeywasm2go/p5.Fn352
func Fn352(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn355 github.com/goccy/spidermonkeywasm2go/p5.Fn355
func Fn355(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn358 github.com/goccy/spidermonkeywasm2go/p5.Fn358
func Fn358(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn361 github.com/goccy/spidermonkeywasm2go/p5.Fn361
func Fn361(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn364 github.com/goccy/spidermonkeywasm2go/p5.Fn364
func Fn364(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn367 github.com/goccy/spidermonkeywasm2go/p5.Fn367
func Fn367(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn370 github.com/goccy/spidermonkeywasm2go/p5.Fn370
func Fn370(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn373 github.com/goccy/spidermonkeywasm2go/p5.Fn373
func Fn373(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn374 github.com/goccy/spidermonkeywasm2go/p7.Fn374
func Fn374(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn380 github.com/goccy/spidermonkeywasm2go/p7.Fn380
func Fn380(m *base.Module, l0 int32) int32

//go:linkname Fn382 github.com/goccy/spidermonkeywasm2go/p7.Fn382
func Fn382(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn391 github.com/goccy/spidermonkeywasm2go/p5.Fn391
func Fn391(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn392 github.com/goccy/spidermonkeywasm2go/p6.Fn392
func Fn392(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn396 github.com/goccy/spidermonkeywasm2go/p7.Fn396
func Fn396(m *base.Module, l0 int32)

//go:linkname Fn397 github.com/goccy/spidermonkeywasm2go/p7.Fn397
func Fn397(m *base.Module, l0 int32)

//go:linkname Fn399 github.com/goccy/spidermonkeywasm2go/p7.Fn399
func Fn399(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn434 github.com/goccy/spidermonkeywasm2go/p7.Fn434
func Fn434(m *base.Module)

//go:linkname Fn464 github.com/goccy/spidermonkeywasm2go/p5.Fn464
func Fn464(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn482 github.com/goccy/spidermonkeywasm2go/p5.Fn482
func Fn482(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn484 github.com/goccy/spidermonkeywasm2go/p7.Fn484
func Fn484(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn525 github.com/goccy/spidermonkeywasm2go/p5.Fn525
func Fn525(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn532 github.com/goccy/spidermonkeywasm2go/p5.Fn532
func Fn532(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn534 github.com/goccy/spidermonkeywasm2go/p5.Fn534
func Fn534(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn540 github.com/goccy/spidermonkeywasm2go/p5.Fn540
func Fn540(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn573 github.com/goccy/spidermonkeywasm2go/p7.Fn573
func Fn573(m *base.Module, l0 int32)

//go:linkname Fn662 github.com/goccy/spidermonkeywasm2go/p7.Fn662
func Fn662(m *base.Module, l0 int32) int32

//go:linkname Fn663 github.com/goccy/spidermonkeywasm2go/p7.Fn663
func Fn663(m *base.Module, l0 int32)

//go:linkname Fn667 github.com/goccy/spidermonkeywasm2go/p6.Fn667
func Fn667(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn677 github.com/goccy/spidermonkeywasm2go/p7.Fn677
func Fn677(m *base.Module, l0 int32) int32

//go:linkname Fn679 github.com/goccy/spidermonkeywasm2go/p7.Fn679
func Fn679(m *base.Module, l0 int32)

//go:linkname Fn680 github.com/goccy/spidermonkeywasm2go/p7.Fn680
func Fn680(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn690 github.com/goccy/spidermonkeywasm2go/p6.Fn690
func Fn690(m *base.Module, l0 int32) int32

//go:linkname Fn692 github.com/goccy/spidermonkeywasm2go/p7.Fn692
func Fn692(m *base.Module, l0 int32)

//go:linkname Fn721 github.com/goccy/spidermonkeywasm2go/p7.Fn721
func Fn721(m *base.Module, l0 float32) float32

//go:linkname Fn722 github.com/goccy/spidermonkeywasm2go/p7.Fn722
func Fn722(m *base.Module, l0 float64, l1 int32) float64

//go:linkname Fn727 github.com/goccy/spidermonkeywasm2go/p7.Fn727
func Fn727(m *base.Module, l0 float32) float32

//go:linkname Fn728 github.com/goccy/spidermonkeywasm2go/p6.Fn728
func Fn728(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn732 github.com/goccy/spidermonkeywasm2go/p7.Fn732
func Fn732(m *base.Module, l0 int32)

//go:linkname Fn733 github.com/goccy/spidermonkeywasm2go/p6.Fn733
func Fn733(m *base.Module, l0 int32) int32

//go:linkname Fn746 github.com/goccy/spidermonkeywasm2go/p2.Fn746
func Fn746(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn749 github.com/goccy/spidermonkeywasm2go/p7.Fn749
func Fn749(m *base.Module, l0 int32) int32

//go:linkname Fn759 github.com/goccy/spidermonkeywasm2go/p7.Fn759
func Fn759(m *base.Module, l0 int32)

//go:linkname Fn776 github.com/goccy/spidermonkeywasm2go/p7.Fn776
func Fn776(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn777 github.com/goccy/spidermonkeywasm2go/p7.Fn777
func Fn777(m *base.Module, l0 int32) int32

//go:linkname Fn778 github.com/goccy/spidermonkeywasm2go/p4.Fn778
func Fn778(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn780 github.com/goccy/spidermonkeywasm2go/p7.Fn780
func Fn780(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn781 github.com/goccy/spidermonkeywasm2go/p6.Fn781
func Fn781(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn782 github.com/goccy/spidermonkeywasm2go/p7.Fn782
func Fn782(m *base.Module, l0 int32, l1 int32, l2 int32) int32

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

//go:linkname Fn797 github.com/goccy/spidermonkeywasm2go/p7.Fn797
func Fn797(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn798 github.com/goccy/spidermonkeywasm2go/p7.Fn798
func Fn798(m *base.Module, l0 int32) int32

//go:linkname Fn799 github.com/goccy/spidermonkeywasm2go/p7.Fn799
func Fn799(m *base.Module, l0 int32)

//go:linkname Fn805 github.com/goccy/spidermonkeywasm2go/p7.Fn805
func Fn805(m *base.Module, l0 int32) int32

//go:linkname Fn811 github.com/goccy/spidermonkeywasm2go/p7.Fn811
func Fn811(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn814 github.com/goccy/spidermonkeywasm2go/p6.Fn814
func Fn814(m *base.Module, l0 int32) int32

//go:linkname Fn816 github.com/goccy/spidermonkeywasm2go/p6.Fn816
func Fn816(m *base.Module, l0 int32) int32

//go:linkname Fn826 github.com/goccy/spidermonkeywasm2go/p7.Fn826
func Fn826(m *base.Module)

//go:linkname Fn827 github.com/goccy/spidermonkeywasm2go/p2.Fn827
func Fn827(m *base.Module, l0 int32, l1 int32, l2 int32) float64

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

//go:linkname Fn841 github.com/goccy/spidermonkeywasm2go/p0.Fn841
func Fn841(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn842 github.com/goccy/spidermonkeywasm2go/p5.Fn842
func Fn842(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn843 github.com/goccy/spidermonkeywasm2go/p6.Fn843
func Fn843(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn847 github.com/goccy/spidermonkeywasm2go/p0.Fn847
func Fn847(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn850 github.com/goccy/spidermonkeywasm2go/p0.Fn850
func Fn850(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn855 github.com/goccy/spidermonkeywasm2go/p7.Fn855
func Fn855(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn858 github.com/goccy/spidermonkeywasm2go/p0.Fn858
func Fn858(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn859 github.com/goccy/spidermonkeywasm2go/p0.Fn859
func Fn859(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn860 github.com/goccy/spidermonkeywasm2go/p0.Fn860
func Fn860(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn863 github.com/goccy/spidermonkeywasm2go/p0.Fn863
func Fn863(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn864 github.com/goccy/spidermonkeywasm2go/p6.Fn864
func Fn864(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn866 github.com/goccy/spidermonkeywasm2go/p5.Fn866
func Fn866(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn867 github.com/goccy/spidermonkeywasm2go/p4.Fn867
func Fn867(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn872 github.com/goccy/spidermonkeywasm2go/p6.Fn872
func Fn872(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn876 github.com/goccy/spidermonkeywasm2go/p6.Fn876
func Fn876(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn877 github.com/goccy/spidermonkeywasm2go/p6.Fn877
func Fn877(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn879 github.com/goccy/spidermonkeywasm2go/p0.Fn879
func Fn879(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn880 github.com/goccy/spidermonkeywasm2go/p5.Fn880
func Fn880(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn881 github.com/goccy/spidermonkeywasm2go/p6.Fn881
func Fn881(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn883 github.com/goccy/spidermonkeywasm2go/p0.Fn883
func Fn883(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn884 github.com/goccy/spidermonkeywasm2go/p0.Fn884
func Fn884(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn885 github.com/goccy/spidermonkeywasm2go/p0.Fn885
func Fn885(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn893 github.com/goccy/spidermonkeywasm2go/p6.Fn893
func Fn893(m *base.Module, l0 int32) int32

//go:linkname Fn895 github.com/goccy/spidermonkeywasm2go/p0.Fn895
func Fn895(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn896 github.com/goccy/spidermonkeywasm2go/p6.Fn896
func Fn896(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn898 github.com/goccy/spidermonkeywasm2go/p7.Fn898
func Fn898(m *base.Module, l0 int32) int32

//go:linkname Fn900 github.com/goccy/spidermonkeywasm2go/p5.Fn900
func Fn900(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn912 github.com/goccy/spidermonkeywasm2go/p6.Fn912
func Fn912(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn924 github.com/goccy/spidermonkeywasm2go/p6.Fn924
func Fn924(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn940 github.com/goccy/spidermonkeywasm2go/p7.Fn940
func Fn940(m *base.Module, l0 int32) int32

//go:linkname Fn941 github.com/goccy/spidermonkeywasm2go/p5.Fn941
func Fn941(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1009 github.com/goccy/spidermonkeywasm2go/p7.Fn1009
func Fn1009(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1042 github.com/goccy/spidermonkeywasm2go/p6.Fn1042
func Fn1042(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64, l5 int32, l6 int32) int32

//go:linkname Fn1043 github.com/goccy/spidermonkeywasm2go/p5.Fn1043
func Fn1043(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1045 github.com/goccy/spidermonkeywasm2go/p6.Fn1045
func Fn1045(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1049 github.com/goccy/spidermonkeywasm2go/p6.Fn1049
func Fn1049(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1050 github.com/goccy/spidermonkeywasm2go/p6.Fn1050
func Fn1050(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1088 github.com/goccy/spidermonkeywasm2go/p0.Fn1088
func Fn1088(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1090 github.com/goccy/spidermonkeywasm2go/p7.Fn1090
func Fn1090(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1092 github.com/goccy/spidermonkeywasm2go/p0.Fn1092
func Fn1092(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1094 github.com/goccy/spidermonkeywasm2go/p0.Fn1094
func Fn1094(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1095 github.com/goccy/spidermonkeywasm2go/p0.Fn1095
func Fn1095(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1096 github.com/goccy/spidermonkeywasm2go/p0.Fn1096
func Fn1096(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1097 github.com/goccy/spidermonkeywasm2go/p0.Fn1097
func Fn1097(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1099 github.com/goccy/spidermonkeywasm2go/p0.Fn1099
func Fn1099(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1100 github.com/goccy/spidermonkeywasm2go/p6.Fn1100
func Fn1100(m *base.Module, l0 int32) int32

//go:linkname Fn1108 github.com/goccy/spidermonkeywasm2go/p0.Fn1108
func Fn1108(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1116 github.com/goccy/spidermonkeywasm2go/p0.Fn1116
func Fn1116(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1117 github.com/goccy/spidermonkeywasm2go/p0.Fn1117
func Fn1117(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1125 github.com/goccy/spidermonkeywasm2go/p0.Fn1125
func Fn1125(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1127 github.com/goccy/spidermonkeywasm2go/p0.Fn1127
func Fn1127(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1128 github.com/goccy/spidermonkeywasm2go/p0.Fn1128
func Fn1128(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1130 github.com/goccy/spidermonkeywasm2go/p0.Fn1130
func Fn1130(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1136 github.com/goccy/spidermonkeywasm2go/p5.Fn1136
func Fn1136(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1137 github.com/goccy/spidermonkeywasm2go/p6.Fn1137
func Fn1137(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1139 github.com/goccy/spidermonkeywasm2go/p0.Fn1139
func Fn1139(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1140 github.com/goccy/spidermonkeywasm2go/p0.Fn1140
func Fn1140(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1142 github.com/goccy/spidermonkeywasm2go/p0.Fn1142
func Fn1142(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1144 github.com/goccy/spidermonkeywasm2go/p6.Fn1144
func Fn1144(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1145 github.com/goccy/spidermonkeywasm2go/p0.Fn1145
func Fn1145(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1170 github.com/goccy/spidermonkeywasm2go/p6.Fn1170
func Fn1170(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1171 github.com/goccy/spidermonkeywasm2go/p5.Fn1171
func Fn1171(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1172 github.com/goccy/spidermonkeywasm2go/p6.Fn1172
func Fn1172(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1185 github.com/goccy/spidermonkeywasm2go/p4.Fn1185
func Fn1185(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn1278 github.com/goccy/spidermonkeywasm2go/p0.Fn1278
func Fn1278(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1281 github.com/goccy/spidermonkeywasm2go/p6.Fn1281
func Fn1281(m *base.Module, l0 int32) int32

//go:linkname Fn1282 github.com/goccy/spidermonkeywasm2go/p6.Fn1282
func Fn1282(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1289 github.com/goccy/spidermonkeywasm2go/p4.Fn1289
func Fn1289(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1290 github.com/goccy/spidermonkeywasm2go/p6.Fn1290
func Fn1290(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1305 github.com/goccy/spidermonkeywasm2go/p6.Fn1305
func Fn1305(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1306 github.com/goccy/spidermonkeywasm2go/p0.Fn1306
func Fn1306(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1320 github.com/goccy/spidermonkeywasm2go/p5.Fn1320
func Fn1320(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1322 github.com/goccy/spidermonkeywasm2go/p6.Fn1322
func Fn1322(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1323 github.com/goccy/spidermonkeywasm2go/p6.Fn1323
func Fn1323(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1361 github.com/goccy/spidermonkeywasm2go/p7.Fn1361
func Fn1361(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1362 github.com/goccy/spidermonkeywasm2go/p6.Fn1362
func Fn1362(m *base.Module, l0 int32) int32

//go:linkname Fn1363 github.com/goccy/spidermonkeywasm2go/p4.Fn1363
func Fn1363(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1364 github.com/goccy/spidermonkeywasm2go/p6.Fn1364
func Fn1364(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1412 github.com/goccy/spidermonkeywasm2go/p0.Fn1412
func Fn1412(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1415 github.com/goccy/spidermonkeywasm2go/p0.Fn1415
func Fn1415(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1416 github.com/goccy/spidermonkeywasm2go/p7.Fn1416
func Fn1416(m *base.Module, l0 int32) int32

//go:linkname Fn1442 github.com/goccy/spidermonkeywasm2go/p6.Fn1442
func Fn1442(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn1444 github.com/goccy/spidermonkeywasm2go/p0.Fn1444
func Fn1444(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1451 github.com/goccy/spidermonkeywasm2go/p0.Fn1451
func Fn1451(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1452 github.com/goccy/spidermonkeywasm2go/p6.Fn1452
func Fn1452(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1453 github.com/goccy/spidermonkeywasm2go/p5.Fn1453
func Fn1453(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1455 github.com/goccy/spidermonkeywasm2go/p7.Fn1455
func Fn1455(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1457 github.com/goccy/spidermonkeywasm2go/p0.Fn1457
func Fn1457(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1459 github.com/goccy/spidermonkeywasm2go/p5.Fn1459
func Fn1459(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1465 github.com/goccy/spidermonkeywasm2go/p6.Fn1465
func Fn1465(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn1468 github.com/goccy/spidermonkeywasm2go/p5.Fn1468
func Fn1468(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1469 github.com/goccy/spidermonkeywasm2go/p6.Fn1469
func Fn1469(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1470 github.com/goccy/spidermonkeywasm2go/p7.Fn1470
func Fn1470(m *base.Module, l0 int32)

//go:linkname Fn1471 github.com/goccy/spidermonkeywasm2go/p6.Fn1471
func Fn1471(m *base.Module, l0 int32)

//go:linkname Fn1472 github.com/goccy/spidermonkeywasm2go/p5.Fn1472
func Fn1472(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1473 github.com/goccy/spidermonkeywasm2go/p6.Fn1473
func Fn1473(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1478 github.com/goccy/spidermonkeywasm2go/p7.Fn1478
func Fn1478(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1480 github.com/goccy/spidermonkeywasm2go/p0.Fn1480
func Fn1480(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1481 github.com/goccy/spidermonkeywasm2go/p0.Fn1481
func Fn1481(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1482 github.com/goccy/spidermonkeywasm2go/p1.Fn1482
func Fn1482(m *base.Module, l0 int32, l1 int32) int32

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

//go:linkname Fn1501 github.com/goccy/spidermonkeywasm2go/p6.Fn1501
func Fn1501(m *base.Module, l0 int32) int32

//go:linkname Fn1503 github.com/goccy/spidermonkeywasm2go/p0.Fn1503
func Fn1503(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1505 github.com/goccy/spidermonkeywasm2go/p7.Fn1505
func Fn1505(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1506 github.com/goccy/spidermonkeywasm2go/p0.Fn1506
func Fn1506(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1510 github.com/goccy/spidermonkeywasm2go/p7.Fn1510
func Fn1510(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1512 github.com/goccy/spidermonkeywasm2go/p7.Fn1512
func Fn1512(m *base.Module, l0 int32)

//go:linkname Fn1516 github.com/goccy/spidermonkeywasm2go/p7.Fn1516
func Fn1516(m *base.Module, l0 int32)

//go:linkname Fn1517 github.com/goccy/spidermonkeywasm2go/p7.Fn1517
func Fn1517(m *base.Module, l0 int32) int32

//go:linkname Fn1518 github.com/goccy/spidermonkeywasm2go/p5.Fn1518
func Fn1518(m *base.Module, l0 float64, l1 float64, l2 float64) float64

//go:linkname Fn1524 github.com/goccy/spidermonkeywasm2go/p6.Fn1524
func Fn1524(m *base.Module, l0 int32, l1 float64) float64

//go:linkname Fn1530 github.com/goccy/spidermonkeywasm2go/p6.Fn1530
func Fn1530(m *base.Module, l0 int32, l1 int32, l2 float64) int32

//go:linkname Fn1531 github.com/goccy/spidermonkeywasm2go/p1.Fn1531
func Fn1531(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1532 github.com/goccy/spidermonkeywasm2go/p7.Fn1532
func Fn1532(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1533 github.com/goccy/spidermonkeywasm2go/p6.Fn1533
func Fn1533(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1590 github.com/goccy/spidermonkeywasm2go/p4.Fn1590
func Fn1590(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1603 github.com/goccy/spidermonkeywasm2go/p6.Fn1603
func Fn1603(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1605 github.com/goccy/spidermonkeywasm2go/p7.Fn1605
func Fn1605(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1627 github.com/goccy/spidermonkeywasm2go/p4.Fn1627
func Fn1627(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1628 github.com/goccy/spidermonkeywasm2go/p0.Fn1628
func Fn1628(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1629 github.com/goccy/spidermonkeywasm2go/p0.Fn1629
func Fn1629(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1630 github.com/goccy/spidermonkeywasm2go/p6.Fn1630
func Fn1630(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1631 github.com/goccy/spidermonkeywasm2go/p6.Fn1631
func Fn1631(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1635 github.com/goccy/spidermonkeywasm2go/p7.Fn1635
func Fn1635(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1640 github.com/goccy/spidermonkeywasm2go/p0.Fn1640
func Fn1640(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1650 github.com/goccy/spidermonkeywasm2go/p6.Fn1650
func Fn1650(m *base.Module, l0 int32, l1 float64, l2 int32) int32

//go:linkname Fn1651 github.com/goccy/spidermonkeywasm2go/p7.Fn1651
func Fn1651(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1668 github.com/goccy/spidermonkeywasm2go/p0.Fn1668
func Fn1668(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1670 github.com/goccy/spidermonkeywasm2go/p0.Fn1670
func Fn1670(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1757 github.com/goccy/spidermonkeywasm2go/p7.Fn1757
func Fn1757(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1768 github.com/goccy/spidermonkeywasm2go/p7.Fn1768
func Fn1768(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1770 github.com/goccy/spidermonkeywasm2go/p7.Fn1770
func Fn1770(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1779 github.com/goccy/spidermonkeywasm2go/p6.Fn1779
func Fn1779(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1780 github.com/goccy/spidermonkeywasm2go/p7.Fn1780
func Fn1780(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1786 github.com/goccy/spidermonkeywasm2go/p6.Fn1786
func Fn1786(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1789 github.com/goccy/spidermonkeywasm2go/p6.Fn1789
func Fn1789(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1790 github.com/goccy/spidermonkeywasm2go/p7.Fn1790
func Fn1790(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1814 github.com/goccy/spidermonkeywasm2go/p6.Fn1814
func Fn1814(m *base.Module, l0 int32) int32

//go:linkname Fn1815 github.com/goccy/spidermonkeywasm2go/p6.Fn1815
func Fn1815(m *base.Module, l0 int32) int32

//go:linkname Fn1816 github.com/goccy/spidermonkeywasm2go/p6.Fn1816
func Fn1816(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1823 github.com/goccy/spidermonkeywasm2go/p5.Fn1823
func Fn1823(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1826 github.com/goccy/spidermonkeywasm2go/p5.Fn1826
func Fn1826(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1827 github.com/goccy/spidermonkeywasm2go/p0.Fn1827
func Fn1827(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1854 github.com/goccy/spidermonkeywasm2go/p0.Fn1854
func Fn1854(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1855 github.com/goccy/spidermonkeywasm2go/p7.Fn1855
func Fn1855(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1856 github.com/goccy/spidermonkeywasm2go/p5.Fn1856
func Fn1856(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1857 github.com/goccy/spidermonkeywasm2go/p0.Fn1857
func Fn1857(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1862 github.com/goccy/spidermonkeywasm2go/p5.Fn1862
func Fn1862(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn1874 github.com/goccy/spidermonkeywasm2go/p6.Fn1874
func Fn1874(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1877 github.com/goccy/spidermonkeywasm2go/p6.Fn1877
func Fn1877(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1880 github.com/goccy/spidermonkeywasm2go/p7.Fn1880
func Fn1880(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1881 github.com/goccy/spidermonkeywasm2go/p7.Fn1881
func Fn1881(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1896 github.com/goccy/spidermonkeywasm2go/p7.Fn1896
func Fn1896(m *base.Module, l0 int32) int32

//go:linkname Fn1897 github.com/goccy/spidermonkeywasm2go/p6.Fn1897
func Fn1897(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1898 github.com/goccy/spidermonkeywasm2go/p7.Fn1898
func Fn1898(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1899 github.com/goccy/spidermonkeywasm2go/p7.Fn1899
func Fn1899(m *base.Module, l0 int32) int32

//go:linkname Fn1900 github.com/goccy/spidermonkeywasm2go/p7.Fn1900
func Fn1900(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1901 github.com/goccy/spidermonkeywasm2go/p5.Fn1901
func Fn1901(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1902 github.com/goccy/spidermonkeywasm2go/p6.Fn1902
func Fn1902(m *base.Module, l0 int32) int32

//go:linkname Fn1922 github.com/goccy/spidermonkeywasm2go/p4.Fn1922
func Fn1922(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1963 github.com/goccy/spidermonkeywasm2go/p0.Fn1963
func Fn1963(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn1978 github.com/goccy/spidermonkeywasm2go/p7.Fn1978
func Fn1978(m *base.Module, l0 int32) int64

//go:linkname Fn1979 github.com/goccy/spidermonkeywasm2go/p7.Fn1979
func Fn1979(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1985 github.com/goccy/spidermonkeywasm2go/p0.Fn1985
func Fn1985(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1987 github.com/goccy/spidermonkeywasm2go/p0.Fn1987
func Fn1987(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2003 github.com/goccy/spidermonkeywasm2go/p0.Fn2003
func Fn2003(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2008 github.com/goccy/spidermonkeywasm2go/p5.Fn2008
func Fn2008(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2034 github.com/goccy/spidermonkeywasm2go/p7.Fn2034
func Fn2034(m *base.Module, l0 int32) int32

//go:linkname Fn2041 github.com/goccy/spidermonkeywasm2go/p7.Fn2041
func Fn2041(m *base.Module, l0 int32) int32

//go:linkname Fn2057 github.com/goccy/spidermonkeywasm2go/p7.Fn2057
func Fn2057(m *base.Module, l0 int32) int32

//go:linkname Fn2058 github.com/goccy/spidermonkeywasm2go/p6.Fn2058
func Fn2058(m *base.Module, l0 int32)

//go:linkname Fn2059 github.com/goccy/spidermonkeywasm2go/p6.Fn2059
func Fn2059(m *base.Module, l0 int32)

//go:linkname Fn2062 github.com/goccy/spidermonkeywasm2go/p7.Fn2062
func Fn2062(m *base.Module, l0 int32) int32

//go:linkname Fn2063 github.com/goccy/spidermonkeywasm2go/p0.Fn2063
func Fn2063(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2064 github.com/goccy/spidermonkeywasm2go/p0.Fn2064
func Fn2064(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2065 github.com/goccy/spidermonkeywasm2go/p7.Fn2065
func Fn2065(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2066 github.com/goccy/spidermonkeywasm2go/p0.Fn2066
func Fn2066(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2068 github.com/goccy/spidermonkeywasm2go/p0.Fn2068
func Fn2068(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2075 github.com/goccy/spidermonkeywasm2go/p0.Fn2075
func Fn2075(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2083 github.com/goccy/spidermonkeywasm2go/p6.Fn2083
func Fn2083(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2092 github.com/goccy/spidermonkeywasm2go/p7.Fn2092
func Fn2092(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2094 github.com/goccy/spidermonkeywasm2go/p6.Fn2094
func Fn2094(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2099 github.com/goccy/spidermonkeywasm2go/p6.Fn2099
func Fn2099(m *base.Module, l0 int32)

//go:linkname Fn2101 github.com/goccy/spidermonkeywasm2go/p4.Fn2101
func Fn2101(m *base.Module, l0 int32)

//go:linkname Fn2102 github.com/goccy/spidermonkeywasm2go/p7.Fn2102
func Fn2102(m *base.Module, l0 int32) int32

//go:linkname Fn2105 github.com/goccy/spidermonkeywasm2go/p6.Fn2105
func Fn2105(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn2108 github.com/goccy/spidermonkeywasm2go/p7.Fn2108
func Fn2108(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn2109 github.com/goccy/spidermonkeywasm2go/p4.Fn2109
func Fn2109(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn2112 github.com/goccy/spidermonkeywasm2go/p7.Fn2112
func Fn2112(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2123 github.com/goccy/spidermonkeywasm2go/p5.Fn2123
func Fn2123(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2148 github.com/goccy/spidermonkeywasm2go/p6.Fn2148
func Fn2148(m *base.Module, l0 int32)

//go:linkname Fn2152 github.com/goccy/spidermonkeywasm2go/p0.Fn2152
func Fn2152(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2192 github.com/goccy/spidermonkeywasm2go/p6.Fn2192
func Fn2192(m *base.Module, l0 int32) int32

//go:linkname Fn2197 github.com/goccy/spidermonkeywasm2go/p6.Fn2197
func Fn2197(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2199 github.com/goccy/spidermonkeywasm2go/p6.Fn2199
func Fn2199(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2229 github.com/goccy/spidermonkeywasm2go/p6.Fn2229
func Fn2229(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2230 github.com/goccy/spidermonkeywasm2go/p0.Fn2230
func Fn2230(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2234 github.com/goccy/spidermonkeywasm2go/p6.Fn2234
func Fn2234(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2255 github.com/goccy/spidermonkeywasm2go/p0.Fn2255
func Fn2255(m *base.Module, l0 int32) int32

//go:linkname Fn2256 github.com/goccy/spidermonkeywasm2go/p6.Fn2256
func Fn2256(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2260 github.com/goccy/spidermonkeywasm2go/p0.Fn2260
func Fn2260(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2261 github.com/goccy/spidermonkeywasm2go/p0.Fn2261
func Fn2261(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2262 github.com/goccy/spidermonkeywasm2go/p0.Fn2262
func Fn2262(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2266 github.com/goccy/spidermonkeywasm2go/p6.Fn2266
func Fn2266(m *base.Module, l0 int32)

//go:linkname Fn2277 github.com/goccy/spidermonkeywasm2go/p6.Fn2277
func Fn2277(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2278 github.com/goccy/spidermonkeywasm2go/p5.Fn2278
func Fn2278(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2293 github.com/goccy/spidermonkeywasm2go/p0.Fn2293
func Fn2293(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2301 github.com/goccy/spidermonkeywasm2go/p0.Fn2301
func Fn2301(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2304 github.com/goccy/spidermonkeywasm2go/p6.Fn2304
func Fn2304(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2314 github.com/goccy/spidermonkeywasm2go/p7.Fn2314
func Fn2314(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn2318 github.com/goccy/spidermonkeywasm2go/p0.Fn2318
func Fn2318(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2337 github.com/goccy/spidermonkeywasm2go/p0.Fn2337
func Fn2337(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2338 github.com/goccy/spidermonkeywasm2go/p7.Fn2338
func Fn2338(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2340 github.com/goccy/spidermonkeywasm2go/p5.Fn2340
func Fn2340(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2343 github.com/goccy/spidermonkeywasm2go/p7.Fn2343
func Fn2343(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2344 github.com/goccy/spidermonkeywasm2go/p0.Fn2344
func Fn2344(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2413 github.com/goccy/spidermonkeywasm2go/p6.Fn2413
func Fn2413(m *base.Module, l0 int32)

//go:linkname Fn2437 github.com/goccy/spidermonkeywasm2go/p5.Fn2437
func Fn2437(m *base.Module, l0 int32)

//go:linkname Fn2439 github.com/goccy/spidermonkeywasm2go/p5.Fn2439
func Fn2439(m *base.Module, l0 int32) int32

//go:linkname Fn2442 github.com/goccy/spidermonkeywasm2go/p6.Fn2442
func Fn2442(m *base.Module, l0 int32) int32

//go:linkname Fn2447 github.com/goccy/spidermonkeywasm2go/p7.Fn2447
func Fn2447(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2448 github.com/goccy/spidermonkeywasm2go/p0.Fn2448
func Fn2448(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2449 github.com/goccy/spidermonkeywasm2go/p6.Fn2449
func Fn2449(m *base.Module, l0 int32) int32

//go:linkname Fn2451 github.com/goccy/spidermonkeywasm2go/p6.Fn2451
func Fn2451(m *base.Module, l0 int32)

//go:linkname Fn2457 github.com/goccy/spidermonkeywasm2go/p0.Fn2457
func Fn2457(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2461 github.com/goccy/spidermonkeywasm2go/p5.Fn2461
func Fn2461(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2488 github.com/goccy/spidermonkeywasm2go/p4.Fn2488
func Fn2488(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2494 github.com/goccy/spidermonkeywasm2go/p4.Fn2494
func Fn2494(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2496 github.com/goccy/spidermonkeywasm2go/p4.Fn2496
func Fn2496(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2507 github.com/goccy/spidermonkeywasm2go/p2.Fn2507
func Fn2507(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2513 github.com/goccy/spidermonkeywasm2go/p5.Fn2513
func Fn2513(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2514 github.com/goccy/spidermonkeywasm2go/p6.Fn2514
func Fn2514(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2533 github.com/goccy/spidermonkeywasm2go/p0.Fn2533
func Fn2533(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2537 github.com/goccy/spidermonkeywasm2go/p6.Fn2537
func Fn2537(m *base.Module, l0 int32)

//go:linkname Fn2538 github.com/goccy/spidermonkeywasm2go/p0.Fn2538
func Fn2538(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2539 github.com/goccy/spidermonkeywasm2go/p0.Fn2539
func Fn2539(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2540 github.com/goccy/spidermonkeywasm2go/p0.Fn2540
func Fn2540(m *base.Module, l0 int32) int32

//go:linkname Fn2546 github.com/goccy/spidermonkeywasm2go/p7.Fn2546
func Fn2546(m *base.Module, l0 int32)

//go:linkname Fn2555 github.com/goccy/spidermonkeywasm2go/p0.Fn2555
func Fn2555(m *base.Module, l0 int32) int32

//go:linkname Fn2556 github.com/goccy/spidermonkeywasm2go/p0.Fn2556
func Fn2556(m *base.Module, l0 int32)

//go:linkname Fn2561 github.com/goccy/spidermonkeywasm2go/p0.Fn2561
func Fn2561(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2562 github.com/goccy/spidermonkeywasm2go/p0.Fn2562
func Fn2562(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2570 github.com/goccy/spidermonkeywasm2go/p5.Fn2570
func Fn2570(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2576 github.com/goccy/spidermonkeywasm2go/p0.Fn2576
func Fn2576(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2577 github.com/goccy/spidermonkeywasm2go/p0.Fn2577
func Fn2577(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2581 github.com/goccy/spidermonkeywasm2go/p0.Fn2581
func Fn2581(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn2583 github.com/goccy/spidermonkeywasm2go/p0.Fn2583
func Fn2583(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2587 github.com/goccy/spidermonkeywasm2go/p0.Fn2587
func Fn2587(m *base.Module, l0 int32)

//go:linkname Fn2588 github.com/goccy/spidermonkeywasm2go/p0.Fn2588
func Fn2588(m *base.Module, l0 int32)

//go:linkname Fn2591 github.com/goccy/spidermonkeywasm2go/p0.Fn2591
func Fn2591(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2592 github.com/goccy/spidermonkeywasm2go/p0.Fn2592
func Fn2592(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn2593 github.com/goccy/spidermonkeywasm2go/p6.Fn2593
func Fn2593(m *base.Module, l0 int32) int32

//go:linkname Fn2594 github.com/goccy/spidermonkeywasm2go/p7.Fn2594
func Fn2594(m *base.Module, l0 int32)

//go:linkname Fn2599 github.com/goccy/spidermonkeywasm2go/p0.Fn2599
func Fn2599(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2602 github.com/goccy/spidermonkeywasm2go/p4.Fn2602
func Fn2602(m *base.Module, l0 int32)

//go:linkname Fn2605 github.com/goccy/spidermonkeywasm2go/p0.Fn2605
func Fn2605(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2610 github.com/goccy/spidermonkeywasm2go/p0.Fn2610
func Fn2610(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2619 github.com/goccy/spidermonkeywasm2go/p0.Fn2619
func Fn2619(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2621 github.com/goccy/spidermonkeywasm2go/p0.Fn2621
func Fn2621(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2630 github.com/goccy/spidermonkeywasm2go/p0.Fn2630
func Fn2630(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn2641 github.com/goccy/spidermonkeywasm2go/p6.Fn2641
func Fn2641(m *base.Module, l0 int32)

//go:linkname Fn2642 github.com/goccy/spidermonkeywasm2go/p0.Fn2642
func Fn2642(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2644 github.com/goccy/spidermonkeywasm2go/p0.Fn2644
func Fn2644(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2651 github.com/goccy/spidermonkeywasm2go/p0.Fn2651
func Fn2651(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2689 github.com/goccy/spidermonkeywasm2go/p6.Fn2689
func Fn2689(m *base.Module, l0 int32) int32

//go:linkname Fn2708 github.com/goccy/spidermonkeywasm2go/p0.Fn2708
func Fn2708(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2717 github.com/goccy/spidermonkeywasm2go/p7.Fn2717
func Fn2717(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2719 github.com/goccy/spidermonkeywasm2go/p0.Fn2719
func Fn2719(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2720 github.com/goccy/spidermonkeywasm2go/p7.Fn2720
func Fn2720(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2721 github.com/goccy/spidermonkeywasm2go/p7.Fn2721
func Fn2721(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2722 github.com/goccy/spidermonkeywasm2go/p4.Fn2722
func Fn2722(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2725 github.com/goccy/spidermonkeywasm2go/p0.Fn2725
func Fn2725(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2727 github.com/goccy/spidermonkeywasm2go/p7.Fn2727
func Fn2727(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2730 github.com/goccy/spidermonkeywasm2go/p0.Fn2730
func Fn2730(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2731 github.com/goccy/spidermonkeywasm2go/p0.Fn2731
func Fn2731(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2732 github.com/goccy/spidermonkeywasm2go/p0.Fn2732
func Fn2732(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2734 github.com/goccy/spidermonkeywasm2go/p0.Fn2734
func Fn2734(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2738 github.com/goccy/spidermonkeywasm2go/p0.Fn2738
func Fn2738(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2740 github.com/goccy/spidermonkeywasm2go/p0.Fn2740
func Fn2740(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2741 github.com/goccy/spidermonkeywasm2go/p4.Fn2741
func Fn2741(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2742 github.com/goccy/spidermonkeywasm2go/p4.Fn2742
func Fn2742(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2743 github.com/goccy/spidermonkeywasm2go/p6.Fn2743
func Fn2743(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2744 github.com/goccy/spidermonkeywasm2go/p7.Fn2744
func Fn2744(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2746 github.com/goccy/spidermonkeywasm2go/p0.Fn2746
func Fn2746(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2751 github.com/goccy/spidermonkeywasm2go/p0.Fn2751
func Fn2751(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2752 github.com/goccy/spidermonkeywasm2go/p0.Fn2752
func Fn2752(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2753 github.com/goccy/spidermonkeywasm2go/p0.Fn2753
func Fn2753(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2755 github.com/goccy/spidermonkeywasm2go/p0.Fn2755
func Fn2755(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2762 github.com/goccy/spidermonkeywasm2go/p5.Fn2762
func Fn2762(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2763 github.com/goccy/spidermonkeywasm2go/p4.Fn2763
func Fn2763(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2765 github.com/goccy/spidermonkeywasm2go/p0.Fn2765
func Fn2765(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2774 github.com/goccy/spidermonkeywasm2go/p7.Fn2774
func Fn2774(m *base.Module, l0 int32)

//go:linkname Fn2787 github.com/goccy/spidermonkeywasm2go/p6.Fn2787
func Fn2787(m *base.Module, l0 int32)

//go:linkname Fn2788 github.com/goccy/spidermonkeywasm2go/p7.Fn2788
func Fn2788(m *base.Module, l0 int32) int32

//go:linkname Fn2793 github.com/goccy/spidermonkeywasm2go/p6.Fn2793
func Fn2793(m *base.Module, l0 int32)

//go:linkname Fn2794 github.com/goccy/spidermonkeywasm2go/p7.Fn2794
func Fn2794(m *base.Module, l0 int32) int32

//go:linkname Fn2799 github.com/goccy/spidermonkeywasm2go/p6.Fn2799
func Fn2799(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2807 github.com/goccy/spidermonkeywasm2go/p4.Fn2807
func Fn2807(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2812 github.com/goccy/spidermonkeywasm2go/p6.Fn2812
func Fn2812(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn2822 github.com/goccy/spidermonkeywasm2go/p6.Fn2822
func Fn2822(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2838 github.com/goccy/spidermonkeywasm2go/p2.Fn2838
func Fn2838(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2839 github.com/goccy/spidermonkeywasm2go/p5.Fn2839
func Fn2839(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2841 github.com/goccy/spidermonkeywasm2go/p0.Fn2841
func Fn2841(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2842 github.com/goccy/spidermonkeywasm2go/p0.Fn2842
func Fn2842(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2843 github.com/goccy/spidermonkeywasm2go/p5.Fn2843
func Fn2843(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2844 github.com/goccy/spidermonkeywasm2go/p6.Fn2844
func Fn2844(m *base.Module, l0 int32) int32

//go:linkname Fn2850 github.com/goccy/spidermonkeywasm2go/p2.Fn2850
func Fn2850(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2853 github.com/goccy/spidermonkeywasm2go/p5.Fn2853
func Fn2853(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2854 github.com/goccy/spidermonkeywasm2go/p6.Fn2854
func Fn2854(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2857 github.com/goccy/spidermonkeywasm2go/p6.Fn2857
func Fn2857(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2858 github.com/goccy/spidermonkeywasm2go/p5.Fn2858
func Fn2858(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2861 github.com/goccy/spidermonkeywasm2go/p0.Fn2861
func Fn2861(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2862 github.com/goccy/spidermonkeywasm2go/p6.Fn2862
func Fn2862(m *base.Module, l0 int32)

//go:linkname Fn2865 github.com/goccy/spidermonkeywasm2go/p5.Fn2865
func Fn2865(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2866 github.com/goccy/spidermonkeywasm2go/p5.Fn2866
func Fn2866(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2880 github.com/goccy/spidermonkeywasm2go/p0.Fn2880
func Fn2880(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2885 github.com/goccy/spidermonkeywasm2go/p4.Fn2885
func Fn2885(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2886 github.com/goccy/spidermonkeywasm2go/p0.Fn2886
func Fn2886(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2890 github.com/goccy/spidermonkeywasm2go/p4.Fn2890
func Fn2890(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2892 github.com/goccy/spidermonkeywasm2go/p0.Fn2892
func Fn2892(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2896 github.com/goccy/spidermonkeywasm2go/p0.Fn2896
func Fn2896(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2899 github.com/goccy/spidermonkeywasm2go/p4.Fn2899
func Fn2899(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2904 github.com/goccy/spidermonkeywasm2go/p6.Fn2904
func Fn2904(m *base.Module, l0 int32) int32

//go:linkname Fn2905 github.com/goccy/spidermonkeywasm2go/p5.Fn2905
func Fn2905(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2910 github.com/goccy/spidermonkeywasm2go/p6.Fn2910
func Fn2910(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2921 github.com/goccy/spidermonkeywasm2go/p6.Fn2921
func Fn2921(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2923 github.com/goccy/spidermonkeywasm2go/p6.Fn2923
func Fn2923(m *base.Module, l0 int32) int64

//go:linkname Fn2928 github.com/goccy/spidermonkeywasm2go/p4.Fn2928
func Fn2928(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn2939 github.com/goccy/spidermonkeywasm2go/p0.Fn2939
func Fn2939(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2944 github.com/goccy/spidermonkeywasm2go/p4.Fn2944
func Fn2944(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2949 github.com/goccy/spidermonkeywasm2go/p7.Fn2949
func Fn2949(m *base.Module, l0 int32)

//go:linkname Fn2952 github.com/goccy/spidermonkeywasm2go/p7.Fn2952
func Fn2952(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2957 github.com/goccy/spidermonkeywasm2go/p0.Fn2957
func Fn2957(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2984 github.com/goccy/spidermonkeywasm2go/p6.Fn2984
func Fn2984(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2986 github.com/goccy/spidermonkeywasm2go/p6.Fn2986
func Fn2986(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3003 github.com/goccy/spidermonkeywasm2go/p6.Fn3003
func Fn3003(m *base.Module, l0 int32)

//go:linkname Fn3005 github.com/goccy/spidermonkeywasm2go/p5.Fn3005
func Fn3005(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3008 github.com/goccy/spidermonkeywasm2go/p6.Fn3008
func Fn3008(m *base.Module, l0 int32) int32

//go:linkname Fn3009 github.com/goccy/spidermonkeywasm2go/p6.Fn3009
func Fn3009(m *base.Module, l0 int32) int32

//go:linkname Fn3044 github.com/goccy/spidermonkeywasm2go/p0.Fn3044
func Fn3044(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3045 github.com/goccy/spidermonkeywasm2go/p0.Fn3045
func Fn3045(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3047 github.com/goccy/spidermonkeywasm2go/p6.Fn3047
func Fn3047(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3054 github.com/goccy/spidermonkeywasm2go/p6.Fn3054
func Fn3054(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3058 github.com/goccy/spidermonkeywasm2go/p6.Fn3058
func Fn3058(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3061 github.com/goccy/spidermonkeywasm2go/p6.Fn3061
func Fn3061(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3104 github.com/goccy/spidermonkeywasm2go/p6.Fn3104
func Fn3104(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3105 github.com/goccy/spidermonkeywasm2go/p4.Fn3105
func Fn3105(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3106 github.com/goccy/spidermonkeywasm2go/p5.Fn3106
func Fn3106(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3158 github.com/goccy/spidermonkeywasm2go/p6.Fn3158
func Fn3158(m *base.Module, l0 float64) float64

//go:linkname Fn3159 github.com/goccy/spidermonkeywasm2go/p2.Fn3159
func Fn3159(m *base.Module, l0 int32) int32

//go:linkname Fn3161 github.com/goccy/spidermonkeywasm2go/p2.Fn3161
func Fn3161(m *base.Module, l0 int32)

//go:linkname Fn3163 github.com/goccy/spidermonkeywasm2go/p0.Fn3163
func Fn3163(m *base.Module, l0 int32) int32

//go:linkname Fn3165 github.com/goccy/spidermonkeywasm2go/p6.Fn3165
func Fn3165(m *base.Module, l0 int32) int32

//go:linkname Fn3166 github.com/goccy/spidermonkeywasm2go/p7.Fn3166
func Fn3166(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3167 github.com/goccy/spidermonkeywasm2go/p7.Fn3167
func Fn3167(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3168 github.com/goccy/spidermonkeywasm2go/p6.Fn3168
func Fn3168(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3176 github.com/goccy/spidermonkeywasm2go/p0.Fn3176
func Fn3176(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3197 github.com/goccy/spidermonkeywasm2go/p0.Fn3197
func Fn3197(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3206 github.com/goccy/spidermonkeywasm2go/p6.Fn3206
func Fn3206(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3238 github.com/goccy/spidermonkeywasm2go/p0.Fn3238
func Fn3238(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3244 github.com/goccy/spidermonkeywasm2go/p6.Fn3244
func Fn3244(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3246 github.com/goccy/spidermonkeywasm2go/p7.Fn3246
func Fn3246(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3489 github.com/goccy/spidermonkeywasm2go/p6.Fn3489
func Fn3489(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn3490 github.com/goccy/spidermonkeywasm2go/p7.Fn3490
func Fn3490(m *base.Module, l0 int32)

//go:linkname Fn3502 github.com/goccy/spidermonkeywasm2go/p7.Fn3502
func Fn3502(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3510 github.com/goccy/spidermonkeywasm2go/p7.Fn3510
func Fn3510(m *base.Module, l0 int32)

//go:linkname Fn3533 github.com/goccy/spidermonkeywasm2go/p2.Fn3533
func Fn3533(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3534 github.com/goccy/spidermonkeywasm2go/p0.Fn3534
func Fn3534(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3536 github.com/goccy/spidermonkeywasm2go/p6.Fn3536
func Fn3536(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3539 github.com/goccy/spidermonkeywasm2go/p7.Fn3539
func Fn3539(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3543 github.com/goccy/spidermonkeywasm2go/p7.Fn3543
func Fn3543(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3546 github.com/goccy/spidermonkeywasm2go/p7.Fn3546
func Fn3546(m *base.Module, l0 int32) int32

//go:linkname Fn3547 github.com/goccy/spidermonkeywasm2go/p0.Fn3547
func Fn3547(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3549 github.com/goccy/spidermonkeywasm2go/p0.Fn3549
func Fn3549(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3550 github.com/goccy/spidermonkeywasm2go/p0.Fn3550
func Fn3550(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3551 github.com/goccy/spidermonkeywasm2go/p0.Fn3551
func Fn3551(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3557 github.com/goccy/spidermonkeywasm2go/p0.Fn3557
func Fn3557(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3560 github.com/goccy/spidermonkeywasm2go/p0.Fn3560
func Fn3560(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3561 github.com/goccy/spidermonkeywasm2go/p0.Fn3561
func Fn3561(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3567 github.com/goccy/spidermonkeywasm2go/p0.Fn3567
func Fn3567(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3568 github.com/goccy/spidermonkeywasm2go/p0.Fn3568
func Fn3568(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3570 github.com/goccy/spidermonkeywasm2go/p0.Fn3570
func Fn3570(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3571 github.com/goccy/spidermonkeywasm2go/p0.Fn3571
func Fn3571(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3573 github.com/goccy/spidermonkeywasm2go/p6.Fn3573
func Fn3573(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3575 github.com/goccy/spidermonkeywasm2go/p0.Fn3575
func Fn3575(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3577 github.com/goccy/spidermonkeywasm2go/p7.Fn3577
func Fn3577(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3578 github.com/goccy/spidermonkeywasm2go/p2.Fn3578
func Fn3578(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn3586 github.com/goccy/spidermonkeywasm2go/p6.Fn3586
func Fn3586(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3588 github.com/goccy/spidermonkeywasm2go/p2.Fn3588
func Fn3588(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3595 github.com/goccy/spidermonkeywasm2go/p6.Fn3595
func Fn3595(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3596 github.com/goccy/spidermonkeywasm2go/p6.Fn3596
func Fn3596(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3600 github.com/goccy/spidermonkeywasm2go/p6.Fn3600
func Fn3600(m *base.Module, l0 int32)

//go:linkname Fn3615 github.com/goccy/spidermonkeywasm2go/p6.Fn3615
func Fn3615(m *base.Module, l0 int32)

//go:linkname Fn3632 github.com/goccy/spidermonkeywasm2go/p5.Fn3632
func Fn3632(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3639 github.com/goccy/spidermonkeywasm2go/p5.Fn3639
func Fn3639(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3644 github.com/goccy/spidermonkeywasm2go/p0.Fn3644
func Fn3644(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3647 github.com/goccy/spidermonkeywasm2go/p0.Fn3647
func Fn3647(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3674 github.com/goccy/spidermonkeywasm2go/p5.Fn3674
func Fn3674(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3720 github.com/goccy/spidermonkeywasm2go/p0.Fn3720
func Fn3720(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3732 github.com/goccy/spidermonkeywasm2go/p6.Fn3732
func Fn3732(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3733 github.com/goccy/spidermonkeywasm2go/p6.Fn3733
func Fn3733(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3746 github.com/goccy/spidermonkeywasm2go/p6.Fn3746
func Fn3746(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3747 github.com/goccy/spidermonkeywasm2go/p4.Fn3747
func Fn3747(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3766 github.com/goccy/spidermonkeywasm2go/p0.Fn3766
func Fn3766(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3822 github.com/goccy/spidermonkeywasm2go/p6.Fn3822
func Fn3822(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3823 github.com/goccy/spidermonkeywasm2go/p0.Fn3823
func Fn3823(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3824 github.com/goccy/spidermonkeywasm2go/p0.Fn3824
func Fn3824(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3825 github.com/goccy/spidermonkeywasm2go/p0.Fn3825
func Fn3825(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3826 github.com/goccy/spidermonkeywasm2go/p0.Fn3826
func Fn3826(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3827 github.com/goccy/spidermonkeywasm2go/p0.Fn3827
func Fn3827(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3828 github.com/goccy/spidermonkeywasm2go/p0.Fn3828
func Fn3828(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3829 github.com/goccy/spidermonkeywasm2go/p0.Fn3829
func Fn3829(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3830 github.com/goccy/spidermonkeywasm2go/p0.Fn3830
func Fn3830(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3831 github.com/goccy/spidermonkeywasm2go/p0.Fn3831
func Fn3831(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3832 github.com/goccy/spidermonkeywasm2go/p0.Fn3832
func Fn3832(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3833 github.com/goccy/spidermonkeywasm2go/p0.Fn3833
func Fn3833(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3834 github.com/goccy/spidermonkeywasm2go/p0.Fn3834
func Fn3834(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3876 github.com/goccy/spidermonkeywasm2go/p0.Fn3876
func Fn3876(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3878 github.com/goccy/spidermonkeywasm2go/p0.Fn3878
func Fn3878(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3879 github.com/goccy/spidermonkeywasm2go/p5.Fn3879
func Fn3879(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3880 github.com/goccy/spidermonkeywasm2go/p0.Fn3880
func Fn3880(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3882 github.com/goccy/spidermonkeywasm2go/p0.Fn3882
func Fn3882(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3884 github.com/goccy/spidermonkeywasm2go/p0.Fn3884
func Fn3884(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3886 github.com/goccy/spidermonkeywasm2go/p0.Fn3886
func Fn3886(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3888 github.com/goccy/spidermonkeywasm2go/p0.Fn3888
func Fn3888(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3890 github.com/goccy/spidermonkeywasm2go/p0.Fn3890
func Fn3890(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3892 github.com/goccy/spidermonkeywasm2go/p0.Fn3892
func Fn3892(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3894 github.com/goccy/spidermonkeywasm2go/p0.Fn3894
func Fn3894(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3896 github.com/goccy/spidermonkeywasm2go/p0.Fn3896
func Fn3896(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3898 github.com/goccy/spidermonkeywasm2go/p0.Fn3898
func Fn3898(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4036 github.com/goccy/spidermonkeywasm2go/p6.Fn4036
func Fn4036(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4039 github.com/goccy/spidermonkeywasm2go/p5.Fn4039
func Fn4039(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn4040 github.com/goccy/spidermonkeywasm2go/p5.Fn4040
func Fn4040(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn4041 github.com/goccy/spidermonkeywasm2go/p5.Fn4041
func Fn4041(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn4042 github.com/goccy/spidermonkeywasm2go/p5.Fn4042
func Fn4042(m *base.Module, l0 int32, l1 float32, l2 int32, l3 int32)

//go:linkname Fn4043 github.com/goccy/spidermonkeywasm2go/p5.Fn4043
func Fn4043(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32)

//go:linkname Fn4044 github.com/goccy/spidermonkeywasm2go/p5.Fn4044
func Fn4044(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn4045 github.com/goccy/spidermonkeywasm2go/p5.Fn4045
func Fn4045(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn4046 github.com/goccy/spidermonkeywasm2go/p6.Fn4046
func Fn4046(m *base.Module, l0 int64) int32

//go:linkname Fn4047 github.com/goccy/spidermonkeywasm2go/p5.Fn4047
func Fn4047(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4048 github.com/goccy/spidermonkeywasm2go/p6.Fn4048
func Fn4048(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn4049 github.com/goccy/spidermonkeywasm2go/p5.Fn4049
func Fn4049(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn4050 github.com/goccy/spidermonkeywasm2go/p5.Fn4050
func Fn4050(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn4051 github.com/goccy/spidermonkeywasm2go/p5.Fn4051
func Fn4051(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn4052 github.com/goccy/spidermonkeywasm2go/p5.Fn4052
func Fn4052(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn4053 github.com/goccy/spidermonkeywasm2go/p5.Fn4053
func Fn4053(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn4054 github.com/goccy/spidermonkeywasm2go/p5.Fn4054
func Fn4054(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn4055 github.com/goccy/spidermonkeywasm2go/p5.Fn4055
func Fn4055(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn4056 github.com/goccy/spidermonkeywasm2go/p6.Fn4056
func Fn4056(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn4057 github.com/goccy/spidermonkeywasm2go/p5.Fn4057
func Fn4057(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn4058 github.com/goccy/spidermonkeywasm2go/p6.Fn4058
func Fn4058(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn4066 github.com/goccy/spidermonkeywasm2go/p6.Fn4066
func Fn4066(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4067 github.com/goccy/spidermonkeywasm2go/p6.Fn4067
func Fn4067(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4068 github.com/goccy/spidermonkeywasm2go/p6.Fn4068
func Fn4068(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4069 github.com/goccy/spidermonkeywasm2go/p6.Fn4069
func Fn4069(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4098 github.com/goccy/spidermonkeywasm2go/p5.Fn4098
func Fn4098(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4099 github.com/goccy/spidermonkeywasm2go/p5.Fn4099
func Fn4099(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4102 github.com/goccy/spidermonkeywasm2go/p6.Fn4102
func Fn4102(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4105 github.com/goccy/spidermonkeywasm2go/p6.Fn4105
func Fn4105(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4106 github.com/goccy/spidermonkeywasm2go/p7.Fn4106
func Fn4106(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4127 github.com/goccy/spidermonkeywasm2go/p5.Fn4127
func Fn4127(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4155 github.com/goccy/spidermonkeywasm2go/p6.Fn4155
func Fn4155(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4188 github.com/goccy/spidermonkeywasm2go/p6.Fn4188
func Fn4188(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4196 github.com/goccy/spidermonkeywasm2go/p6.Fn4196
func Fn4196(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4197 github.com/goccy/spidermonkeywasm2go/p7.Fn4197
func Fn4197(m *base.Module, l0 int32)

//go:linkname Fn4198 github.com/goccy/spidermonkeywasm2go/p7.Fn4198
func Fn4198(m *base.Module, l0 int32)

//go:linkname Fn4204 github.com/goccy/spidermonkeywasm2go/p6.Fn4204
func Fn4204(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4205 github.com/goccy/spidermonkeywasm2go/p7.Fn4205
func Fn4205(m *base.Module, l0 int32)

//go:linkname Fn4212 github.com/goccy/spidermonkeywasm2go/p6.Fn4212
func Fn4212(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4213 github.com/goccy/spidermonkeywasm2go/p6.Fn4213
func Fn4213(m *base.Module, l0 int32) int32

//go:linkname Fn4224 github.com/goccy/spidermonkeywasm2go/p6.Fn4224
func Fn4224(m *base.Module, l0 int32) int32

//go:linkname Fn4228 github.com/goccy/spidermonkeywasm2go/p7.Fn4228
func Fn4228(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4234 github.com/goccy/spidermonkeywasm2go/p6.Fn4234
func Fn4234(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4241 github.com/goccy/spidermonkeywasm2go/p2.Fn4241
func Fn4241(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4242 github.com/goccy/spidermonkeywasm2go/p6.Fn4242
func Fn4242(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4256 github.com/goccy/spidermonkeywasm2go/p4.Fn4256
func Fn4256(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4278 github.com/goccy/spidermonkeywasm2go/p4.Fn4278
func Fn4278(m *base.Module, l0 int32) int32

//go:linkname Fn4300 github.com/goccy/spidermonkeywasm2go/p7.Fn4300
func Fn4300(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4319 github.com/goccy/spidermonkeywasm2go/p4.Fn4319
func Fn4319(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4329 github.com/goccy/spidermonkeywasm2go/p6.Fn4329
func Fn4329(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4330 github.com/goccy/spidermonkeywasm2go/p6.Fn4330
func Fn4330(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4331 github.com/goccy/spidermonkeywasm2go/p4.Fn4331
func Fn4331(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4342 github.com/goccy/spidermonkeywasm2go/p4.Fn4342
func Fn4342(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4344 github.com/goccy/spidermonkeywasm2go/p5.Fn4344
func Fn4344(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4347 github.com/goccy/spidermonkeywasm2go/p5.Fn4347
func Fn4347(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4348 github.com/goccy/spidermonkeywasm2go/p6.Fn4348
func Fn4348(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4349 github.com/goccy/spidermonkeywasm2go/p2.Fn4349
func Fn4349(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4352 github.com/goccy/spidermonkeywasm2go/p6.Fn4352
func Fn4352(m *base.Module, l0 int32, l1 int32, l2 float64, l3 float64)

//go:linkname Fn4358 github.com/goccy/spidermonkeywasm2go/p7.Fn4358
func Fn4358(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4408 github.com/goccy/spidermonkeywasm2go/p7.Fn4408
func Fn4408(m *base.Module, l0 int32)

//go:linkname Fn4414 github.com/goccy/spidermonkeywasm2go/p6.Fn4414
func Fn4414(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4418 github.com/goccy/spidermonkeywasm2go/p4.Fn4418
func Fn4418(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4421 github.com/goccy/spidermonkeywasm2go/p5.Fn4421
func Fn4421(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4422 github.com/goccy/spidermonkeywasm2go/p6.Fn4422
func Fn4422(m *base.Module, l0 int32) int32

//go:linkname Fn4423 github.com/goccy/spidermonkeywasm2go/p7.Fn4423
func Fn4423(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32

//go:linkname Fn4424 github.com/goccy/spidermonkeywasm2go/p7.Fn4424
func Fn4424(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4425 github.com/goccy/spidermonkeywasm2go/p6.Fn4425
func Fn4425(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4426 github.com/goccy/spidermonkeywasm2go/p6.Fn4426
func Fn4426(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4427 github.com/goccy/spidermonkeywasm2go/p6.Fn4427
func Fn4427(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4428 github.com/goccy/spidermonkeywasm2go/p6.Fn4428
func Fn4428(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4429 github.com/goccy/spidermonkeywasm2go/p6.Fn4429
func Fn4429(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4432 github.com/goccy/spidermonkeywasm2go/p6.Fn4432
func Fn4432(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4435 github.com/goccy/spidermonkeywasm2go/p5.Fn4435
func Fn4435(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4437 github.com/goccy/spidermonkeywasm2go/p5.Fn4437
func Fn4437(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4438 github.com/goccy/spidermonkeywasm2go/p4.Fn4438
func Fn4438(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4439 github.com/goccy/spidermonkeywasm2go/p6.Fn4439
func Fn4439(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4460 github.com/goccy/spidermonkeywasm2go/p6.Fn4460
func Fn4460(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4461 github.com/goccy/spidermonkeywasm2go/p7.Fn4461
func Fn4461(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4464 github.com/goccy/spidermonkeywasm2go/p5.Fn4464
func Fn4464(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4465 github.com/goccy/spidermonkeywasm2go/p6.Fn4465
func Fn4465(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4466 github.com/goccy/spidermonkeywasm2go/p6.Fn4466
func Fn4466(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4467 github.com/goccy/spidermonkeywasm2go/p5.Fn4467
func Fn4467(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4506 github.com/goccy/spidermonkeywasm2go/p2.Fn4506
func Fn4506(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4513 github.com/goccy/spidermonkeywasm2go/p7.Fn4513
func Fn4513(m *base.Module, l0 int32)

//go:linkname Fn4517 github.com/goccy/spidermonkeywasm2go/p5.Fn4517
func Fn4517(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4520 github.com/goccy/spidermonkeywasm2go/p7.Fn4520
func Fn4520(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4521 github.com/goccy/spidermonkeywasm2go/p7.Fn4521
func Fn4521(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4522 github.com/goccy/spidermonkeywasm2go/p6.Fn4522
func Fn4522(m *base.Module, l0 int32) int32

//go:linkname Fn4525 github.com/goccy/spidermonkeywasm2go/p6.Fn4525
func Fn4525(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4526 github.com/goccy/spidermonkeywasm2go/p4.Fn4526
func Fn4526(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 float64, l7 int32) int32

//go:linkname Fn4527 github.com/goccy/spidermonkeywasm2go/p6.Fn4527
func Fn4527(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4533 github.com/goccy/spidermonkeywasm2go/p6.Fn4533
func Fn4533(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4546 github.com/goccy/spidermonkeywasm2go/p6.Fn4546
func Fn4546(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4547 github.com/goccy/spidermonkeywasm2go/p6.Fn4547
func Fn4547(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4548 github.com/goccy/spidermonkeywasm2go/p6.Fn4548
func Fn4548(m *base.Module, l0 int32) int32

//go:linkname Fn4553 github.com/goccy/spidermonkeywasm2go/p5.Fn4553
func Fn4553(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4557 github.com/goccy/spidermonkeywasm2go/p6.Fn4557
func Fn4557(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4558 github.com/goccy/spidermonkeywasm2go/p6.Fn4558
func Fn4558(m *base.Module, l0 int32) int32

//go:linkname Fn4559 github.com/goccy/spidermonkeywasm2go/p2.Fn4559
func Fn4559(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4569 github.com/goccy/spidermonkeywasm2go/p4.Fn4569
func Fn4569(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4575 github.com/goccy/spidermonkeywasm2go/p7.Fn4575
func Fn4575(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4576 github.com/goccy/spidermonkeywasm2go/p6.Fn4576
func Fn4576(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4579 github.com/goccy/spidermonkeywasm2go/p4.Fn4579
func Fn4579(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4580 github.com/goccy/spidermonkeywasm2go/p7.Fn4580
func Fn4580(m *base.Module, l0 int32) int32

//go:linkname Fn4586 github.com/goccy/spidermonkeywasm2go/p6.Fn4586
func Fn4586(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4588 github.com/goccy/spidermonkeywasm2go/p4.Fn4588
func Fn4588(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn4610 github.com/goccy/spidermonkeywasm2go/p2.Fn4610
func Fn4610(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4611 github.com/goccy/spidermonkeywasm2go/p5.Fn4611
func Fn4611(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn4613 github.com/goccy/spidermonkeywasm2go/p6.Fn4613
func Fn4613(m *base.Module, l0 int32) int32

//go:linkname Fn4634 github.com/goccy/spidermonkeywasm2go/p7.Fn4634
func Fn4634(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4657 github.com/goccy/spidermonkeywasm2go/p6.Fn4657
func Fn4657(m *base.Module, l0 int32) int32

//go:linkname Fn4658 github.com/goccy/spidermonkeywasm2go/p6.Fn4658
func Fn4658(m *base.Module, l0 int32)

//go:linkname Fn4659 github.com/goccy/spidermonkeywasm2go/p6.Fn4659
func Fn4659(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4662 github.com/goccy/spidermonkeywasm2go/p5.Fn4662
func Fn4662(m *base.Module, l0 int32) int32

//go:linkname Fn4669 github.com/goccy/spidermonkeywasm2go/p6.Fn4669
func Fn4669(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4670 github.com/goccy/spidermonkeywasm2go/p6.Fn4670
func Fn4670(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4696 github.com/goccy/spidermonkeywasm2go/p4.Fn4696
func Fn4696(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4706 github.com/goccy/spidermonkeywasm2go/p5.Fn4706
func Fn4706(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4707 github.com/goccy/spidermonkeywasm2go/p0.Fn4707
func Fn4707(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4720 github.com/goccy/spidermonkeywasm2go/p7.Fn4720
func Fn4720(m *base.Module, l0 int32)

//go:linkname Fn4737 github.com/goccy/spidermonkeywasm2go/p7.Fn4737
func Fn4737(m *base.Module, l0 int32)

//go:linkname Fn4740 github.com/goccy/spidermonkeywasm2go/p6.Fn4740
func Fn4740(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4745 github.com/goccy/spidermonkeywasm2go/p5.Fn4745
func Fn4745(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4747 github.com/goccy/spidermonkeywasm2go/p5.Fn4747
func Fn4747(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4754 github.com/goccy/spidermonkeywasm2go/p5.Fn4754
func Fn4754(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5147 github.com/goccy/spidermonkeywasm2go/p6.Fn5147
func Fn5147(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5156 github.com/goccy/spidermonkeywasm2go/p0.Fn5156
func Fn5156(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5160 github.com/goccy/spidermonkeywasm2go/p0.Fn5160
func Fn5160(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn5161 github.com/goccy/spidermonkeywasm2go/p0.Fn5161
func Fn5161(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn5162 github.com/goccy/spidermonkeywasm2go/p6.Fn5162
func Fn5162(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5163 github.com/goccy/spidermonkeywasm2go/p2.Fn5163
func Fn5163(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5350 github.com/goccy/spidermonkeywasm2go/p5.Fn5350
func Fn5350(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5351 github.com/goccy/spidermonkeywasm2go/p4.Fn5351
func Fn5351(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5353 github.com/goccy/spidermonkeywasm2go/p5.Fn5353
func Fn5353(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5354 github.com/goccy/spidermonkeywasm2go/p5.Fn5354
func Fn5354(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5358 github.com/goccy/spidermonkeywasm2go/p4.Fn5358
func Fn5358(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5360 github.com/goccy/spidermonkeywasm2go/p6.Fn5360
func Fn5360(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5438 github.com/goccy/spidermonkeywasm2go/p7.Fn5438
func Fn5438(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5470 github.com/goccy/spidermonkeywasm2go/p7.Fn5470
func Fn5470(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5550 github.com/goccy/spidermonkeywasm2go/p6.Fn5550
func Fn5550(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5568 github.com/goccy/spidermonkeywasm2go/p6.Fn5568
func Fn5568(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5569 github.com/goccy/spidermonkeywasm2go/p7.Fn5569
func Fn5569(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5572 github.com/goccy/spidermonkeywasm2go/p7.Fn5572
func Fn5572(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5573 github.com/goccy/spidermonkeywasm2go/p7.Fn5573
func Fn5573(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5574 github.com/goccy/spidermonkeywasm2go/p4.Fn5574
func Fn5574(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5575 github.com/goccy/spidermonkeywasm2go/p0.Fn5575
func Fn5575(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5576 github.com/goccy/spidermonkeywasm2go/p0.Fn5576
func Fn5576(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5584 github.com/goccy/spidermonkeywasm2go/p4.Fn5584
func Fn5584(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5586 github.com/goccy/spidermonkeywasm2go/p7.Fn5586
func Fn5586(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5594 github.com/goccy/spidermonkeywasm2go/p6.Fn5594
func Fn5594(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5598 github.com/goccy/spidermonkeywasm2go/p5.Fn5598
func Fn5598(m *base.Module, l0 int32)

//go:linkname Fn5599 github.com/goccy/spidermonkeywasm2go/p6.Fn5599
func Fn5599(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5600 github.com/goccy/spidermonkeywasm2go/p6.Fn5600
func Fn5600(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5601 github.com/goccy/spidermonkeywasm2go/p6.Fn5601
func Fn5601(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5602 github.com/goccy/spidermonkeywasm2go/p6.Fn5602
func Fn5602(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5603 github.com/goccy/spidermonkeywasm2go/p4.Fn5603
func Fn5603(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5606 github.com/goccy/spidermonkeywasm2go/p5.Fn5606
func Fn5606(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5607 github.com/goccy/spidermonkeywasm2go/p6.Fn5607
func Fn5607(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5609 github.com/goccy/spidermonkeywasm2go/p6.Fn5609
func Fn5609(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5612 github.com/goccy/spidermonkeywasm2go/p6.Fn5612
func Fn5612(m *base.Module, l0 int32)

//go:linkname Fn5626 github.com/goccy/spidermonkeywasm2go/p2.Fn5626
func Fn5626(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5627 github.com/goccy/spidermonkeywasm2go/p4.Fn5627
func Fn5627(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5628 github.com/goccy/spidermonkeywasm2go/p4.Fn5628
func Fn5628(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5629 github.com/goccy/spidermonkeywasm2go/p6.Fn5629
func Fn5629(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn5630 github.com/goccy/spidermonkeywasm2go/p6.Fn5630
func Fn5630(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn5633 github.com/goccy/spidermonkeywasm2go/p6.Fn5633
func Fn5633(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5634 github.com/goccy/spidermonkeywasm2go/p6.Fn5634
func Fn5634(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5635 github.com/goccy/spidermonkeywasm2go/p4.Fn5635
func Fn5635(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5636 github.com/goccy/spidermonkeywasm2go/p5.Fn5636
func Fn5636(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5649 github.com/goccy/spidermonkeywasm2go/p7.Fn5649
func Fn5649(m *base.Module, l0 float64) float64

//go:linkname Fn5651 github.com/goccy/spidermonkeywasm2go/p6.Fn5651
func Fn5651(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5659 github.com/goccy/spidermonkeywasm2go/p5.Fn5659
func Fn5659(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5663 github.com/goccy/spidermonkeywasm2go/p6.Fn5663
func Fn5663(m *base.Module, l0 int32)

//go:linkname Fn5665 github.com/goccy/spidermonkeywasm2go/p6.Fn5665
func Fn5665(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5670 github.com/goccy/spidermonkeywasm2go/p6.Fn5670
func Fn5670(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5672 github.com/goccy/spidermonkeywasm2go/p7.Fn5672
func Fn5672(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5674 github.com/goccy/spidermonkeywasm2go/p0.Fn5674
func Fn5674(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5688 github.com/goccy/spidermonkeywasm2go/p5.Fn5688
func Fn5688(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5689 github.com/goccy/spidermonkeywasm2go/p5.Fn5689
func Fn5689(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5691 github.com/goccy/spidermonkeywasm2go/p6.Fn5691
func Fn5691(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5694 github.com/goccy/spidermonkeywasm2go/p5.Fn5694
func Fn5694(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5699 github.com/goccy/spidermonkeywasm2go/p5.Fn5699
func Fn5699(m *base.Module, l0 int32)

//go:linkname Fn5702 github.com/goccy/spidermonkeywasm2go/p6.Fn5702
func Fn5702(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5703 github.com/goccy/spidermonkeywasm2go/p7.Fn5703
func Fn5703(m *base.Module, l0 int32)

//go:linkname Fn5705 github.com/goccy/spidermonkeywasm2go/p6.Fn5705
func Fn5705(m *base.Module, l0 int32) int32

//go:linkname Fn5706 github.com/goccy/spidermonkeywasm2go/p7.Fn5706
func Fn5706(m *base.Module, l0 int32) int32

//go:linkname Fn5712 github.com/goccy/spidermonkeywasm2go/p4.Fn5712
func Fn5712(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5718 github.com/goccy/spidermonkeywasm2go/p5.Fn5718
func Fn5718(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5719 github.com/goccy/spidermonkeywasm2go/p6.Fn5719
func Fn5719(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5720 github.com/goccy/spidermonkeywasm2go/p5.Fn5720
func Fn5720(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5722 github.com/goccy/spidermonkeywasm2go/p7.Fn5722
func Fn5722(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5723 github.com/goccy/spidermonkeywasm2go/p5.Fn5723
func Fn5723(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5724 github.com/goccy/spidermonkeywasm2go/p6.Fn5724
func Fn5724(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5725 github.com/goccy/spidermonkeywasm2go/p7.Fn5725
func Fn5725(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5726 github.com/goccy/spidermonkeywasm2go/p6.Fn5726
func Fn5726(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5727 github.com/goccy/spidermonkeywasm2go/p7.Fn5727
func Fn5727(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5728 github.com/goccy/spidermonkeywasm2go/p6.Fn5728
func Fn5728(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5729 github.com/goccy/spidermonkeywasm2go/p6.Fn5729
func Fn5729(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5730 github.com/goccy/spidermonkeywasm2go/p5.Fn5730
func Fn5730(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5731 github.com/goccy/spidermonkeywasm2go/p6.Fn5731
func Fn5731(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5732 github.com/goccy/spidermonkeywasm2go/p7.Fn5732
func Fn5732(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5733 github.com/goccy/spidermonkeywasm2go/p6.Fn5733
func Fn5733(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5735 github.com/goccy/spidermonkeywasm2go/p6.Fn5735
func Fn5735(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5737 github.com/goccy/spidermonkeywasm2go/p7.Fn5737
func Fn5737(m *base.Module, l0 int32)

//go:linkname Fn5738 github.com/goccy/spidermonkeywasm2go/p7.Fn5738
func Fn5738(m *base.Module, l0 int32)

//go:linkname Fn5743 github.com/goccy/spidermonkeywasm2go/p6.Fn5743
func Fn5743(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5751 github.com/goccy/spidermonkeywasm2go/p6.Fn5751
func Fn5751(m *base.Module, l0 int32)

//go:linkname Fn5755 github.com/goccy/spidermonkeywasm2go/p6.Fn5755
func Fn5755(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5756 github.com/goccy/spidermonkeywasm2go/p7.Fn5756
func Fn5756(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5775 github.com/goccy/spidermonkeywasm2go/p5.Fn5775
func Fn5775(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5776 github.com/goccy/spidermonkeywasm2go/p6.Fn5776
func Fn5776(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5777 github.com/goccy/spidermonkeywasm2go/p5.Fn5777
func Fn5777(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5821 github.com/goccy/spidermonkeywasm2go/p6.Fn5821
func Fn5821(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5822 github.com/goccy/spidermonkeywasm2go/p6.Fn5822
func Fn5822(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5823 github.com/goccy/spidermonkeywasm2go/p6.Fn5823
func Fn5823(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5824 github.com/goccy/spidermonkeywasm2go/p6.Fn5824
func Fn5824(m *base.Module, l0 int32)

//go:linkname Fn5839 github.com/goccy/spidermonkeywasm2go/p6.Fn5839
func Fn5839(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5840 github.com/goccy/spidermonkeywasm2go/p6.Fn5840
func Fn5840(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5881 github.com/goccy/spidermonkeywasm2go/p6.Fn5881
func Fn5881(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5882 github.com/goccy/spidermonkeywasm2go/p6.Fn5882
func Fn5882(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5888 github.com/goccy/spidermonkeywasm2go/p6.Fn5888
func Fn5888(m *base.Module, l0 int32) int32

//go:linkname Fn5902 github.com/goccy/spidermonkeywasm2go/p4.Fn5902
func Fn5902(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5916 github.com/goccy/spidermonkeywasm2go/p7.Fn5916
func Fn5916(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5917 github.com/goccy/spidermonkeywasm2go/p7.Fn5917
func Fn5917(m *base.Module, l0 int32) int32

//go:linkname Fn5920 github.com/goccy/spidermonkeywasm2go/p7.Fn5920
func Fn5920(m *base.Module, l0 int32)

//go:linkname Fn5925 github.com/goccy/spidermonkeywasm2go/p5.Fn5925
func Fn5925(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5952 github.com/goccy/spidermonkeywasm2go/p5.Fn5952
func Fn5952(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5953 github.com/goccy/spidermonkeywasm2go/p7.Fn5953
func Fn5953(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5963 github.com/goccy/spidermonkeywasm2go/p4.Fn5963
func Fn5963(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5968 github.com/goccy/spidermonkeywasm2go/p6.Fn5968
func Fn5968(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6040 github.com/goccy/spidermonkeywasm2go/p4.Fn6040
func Fn6040(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6052 github.com/goccy/spidermonkeywasm2go/p6.Fn6052
func Fn6052(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6060 github.com/goccy/spidermonkeywasm2go/p7.Fn6060
func Fn6060(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6061 github.com/goccy/spidermonkeywasm2go/p7.Fn6061
func Fn6061(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6064 github.com/goccy/spidermonkeywasm2go/p6.Fn6064
func Fn6064(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6077 github.com/goccy/spidermonkeywasm2go/p6.Fn6077
func Fn6077(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6092 github.com/goccy/spidermonkeywasm2go/p7.Fn6092
func Fn6092(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6104 github.com/goccy/spidermonkeywasm2go/p7.Fn6104
func Fn6104(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6105 github.com/goccy/spidermonkeywasm2go/p4.Fn6105
func Fn6105(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32)

//go:linkname Fn6106 github.com/goccy/spidermonkeywasm2go/p7.Fn6106
func Fn6106(m *base.Module)

//go:linkname Fn6125 github.com/goccy/spidermonkeywasm2go/p7.Fn6125
func Fn6125(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6186 github.com/goccy/spidermonkeywasm2go/p5.Fn6186
func Fn6186(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6187 github.com/goccy/spidermonkeywasm2go/p6.Fn6187
func Fn6187(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6189 github.com/goccy/spidermonkeywasm2go/p6.Fn6189
func Fn6189(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6190 github.com/goccy/spidermonkeywasm2go/p5.Fn6190
func Fn6190(m *base.Module, l0 int32)

//go:linkname Fn6192 github.com/goccy/spidermonkeywasm2go/p6.Fn6192
func Fn6192(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6210 github.com/goccy/spidermonkeywasm2go/p6.Fn6210
func Fn6210(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6233 github.com/goccy/spidermonkeywasm2go/p4.Fn6233
func Fn6233(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6246 github.com/goccy/spidermonkeywasm2go/p6.Fn6246
func Fn6246(m *base.Module, l0 int32) int32

//go:linkname Fn6260 github.com/goccy/spidermonkeywasm2go/p6.Fn6260
func Fn6260(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6261 github.com/goccy/spidermonkeywasm2go/p2.Fn6261
func Fn6261(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn6262 github.com/goccy/spidermonkeywasm2go/p6.Fn6262
func Fn6262(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6263 github.com/goccy/spidermonkeywasm2go/p7.Fn6263
func Fn6263(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6264 github.com/goccy/spidermonkeywasm2go/p6.Fn6264
func Fn6264(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6265 github.com/goccy/spidermonkeywasm2go/p4.Fn6265
func Fn6265(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6284 github.com/goccy/spidermonkeywasm2go/p6.Fn6284
func Fn6284(m *base.Module) int32

//go:linkname Fn6292 github.com/goccy/spidermonkeywasm2go/p5.Fn6292
func Fn6292(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6293 github.com/goccy/spidermonkeywasm2go/p7.Fn6293
func Fn6293(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6295 github.com/goccy/spidermonkeywasm2go/p5.Fn6295
func Fn6295(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn6346 github.com/goccy/spidermonkeywasm2go/p2.Fn6346
func Fn6346(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6358 github.com/goccy/spidermonkeywasm2go/p7.Fn6358
func Fn6358(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6655 github.com/goccy/spidermonkeywasm2go/p4.Fn6655
func Fn6655(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6664 github.com/goccy/spidermonkeywasm2go/p5.Fn6664
func Fn6664(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6665 github.com/goccy/spidermonkeywasm2go/p6.Fn6665
func Fn6665(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6666 github.com/goccy/spidermonkeywasm2go/p7.Fn6666
func Fn6666(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6667 github.com/goccy/spidermonkeywasm2go/p6.Fn6667
func Fn6667(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6668 github.com/goccy/spidermonkeywasm2go/p6.Fn6668
func Fn6668(m *base.Module, l0 int32) int32

//go:linkname Fn6681 github.com/goccy/spidermonkeywasm2go/p6.Fn6681
func Fn6681(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6688 github.com/goccy/spidermonkeywasm2go/p1.Fn6688
func Fn6688(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6690 github.com/goccy/spidermonkeywasm2go/p7.Fn6690
func Fn6690(m *base.Module, l0 int32)

//go:linkname Fn6693 github.com/goccy/spidermonkeywasm2go/p6.Fn6693
func Fn6693(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6695 github.com/goccy/spidermonkeywasm2go/p4.Fn6695
func Fn6695(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6696 github.com/goccy/spidermonkeywasm2go/p7.Fn6696
func Fn6696(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6699 github.com/goccy/spidermonkeywasm2go/p6.Fn6699
func Fn6699(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6700 github.com/goccy/spidermonkeywasm2go/p6.Fn6700
func Fn6700(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6701 github.com/goccy/spidermonkeywasm2go/p6.Fn6701
func Fn6701(m *base.Module, l0 int32, l1 int32) int32

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

//go:linkname Fn6714 github.com/goccy/spidermonkeywasm2go/p7.Fn6714
func Fn6714(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6715 github.com/goccy/spidermonkeywasm2go/p7.Fn6715
func Fn6715(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6717 github.com/goccy/spidermonkeywasm2go/p7.Fn6717
func Fn6717(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6722 github.com/goccy/spidermonkeywasm2go/p7.Fn6722
func Fn6722(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6723 github.com/goccy/spidermonkeywasm2go/p7.Fn6723
func Fn6723(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6724 github.com/goccy/spidermonkeywasm2go/p7.Fn6724
func Fn6724(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6727 github.com/goccy/spidermonkeywasm2go/p7.Fn6727
func Fn6727(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6732 github.com/goccy/spidermonkeywasm2go/p7.Fn6732
func Fn6732(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6733 github.com/goccy/spidermonkeywasm2go/p6.Fn6733
func Fn6733(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6760 github.com/goccy/spidermonkeywasm2go/p7.Fn6760
func Fn6760(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6763 github.com/goccy/spidermonkeywasm2go/p7.Fn6763
func Fn6763(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6765 github.com/goccy/spidermonkeywasm2go/p1.Fn6765
func Fn6765(m *base.Module, l0 int32) int32

//go:linkname Fn6771 github.com/goccy/spidermonkeywasm2go/p7.Fn6771
func Fn6771(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6779 github.com/goccy/spidermonkeywasm2go/p7.Fn6779
func Fn6779(m *base.Module, l0 int32) int32

//go:linkname Fn6783 github.com/goccy/spidermonkeywasm2go/p7.Fn6783
func Fn6783(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6784 github.com/goccy/spidermonkeywasm2go/p7.Fn6784
func Fn6784(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6785 github.com/goccy/spidermonkeywasm2go/p4.Fn6785
func Fn6785(m *base.Module, l0 int32) int32

//go:linkname Fn6786 github.com/goccy/spidermonkeywasm2go/p6.Fn6786
func Fn6786(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6790 github.com/goccy/spidermonkeywasm2go/p7.Fn6790
func Fn6790(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6795 github.com/goccy/spidermonkeywasm2go/p5.Fn6795
func Fn6795(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6796 github.com/goccy/spidermonkeywasm2go/p6.Fn6796
func Fn6796(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6835 github.com/goccy/spidermonkeywasm2go/p6.Fn6835
func Fn6835(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6836 github.com/goccy/spidermonkeywasm2go/p5.Fn6836
func Fn6836(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6839 github.com/goccy/spidermonkeywasm2go/p1.Fn6839
func Fn6839(m *base.Module, l0 int32) int32

//go:linkname Fn6864 github.com/goccy/spidermonkeywasm2go/p5.Fn6864
func Fn6864(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6898 github.com/goccy/spidermonkeywasm2go/p6.Fn6898
func Fn6898(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6901 github.com/goccy/spidermonkeywasm2go/p5.Fn6901
func Fn6901(m *base.Module, l0 int32, l1 int32, l2 int32) int64

//go:linkname Fn6905 github.com/goccy/spidermonkeywasm2go/p7.Fn6905
func Fn6905(m *base.Module, l0 int32) int32

//go:linkname Fn6911 github.com/goccy/spidermonkeywasm2go/p7.Fn6911
func Fn6911(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6924 github.com/goccy/spidermonkeywasm2go/p5.Fn6924
func Fn6924(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6927 github.com/goccy/spidermonkeywasm2go/p6.Fn6927
func Fn6927(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6947 github.com/goccy/spidermonkeywasm2go/p7.Fn6947
func Fn6947(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6949 github.com/goccy/spidermonkeywasm2go/p0.Fn6949
func Fn6949(m *base.Module, l0 int32) int32

//go:linkname Fn6952 github.com/goccy/spidermonkeywasm2go/p0.Fn6952
func Fn6952(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6953 github.com/goccy/spidermonkeywasm2go/p0.Fn6953
func Fn6953(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6967 github.com/goccy/spidermonkeywasm2go/p7.Fn6967
func Fn6967(m *base.Module, l0 int32)

//go:linkname Fn6968 github.com/goccy/spidermonkeywasm2go/p5.Fn6968
func Fn6968(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6971 github.com/goccy/spidermonkeywasm2go/p5.Fn6971
func Fn6971(m *base.Module, l0 int32) int32

//go:linkname Fn6978 github.com/goccy/spidermonkeywasm2go/p0.Fn6978
func Fn6978(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6979 github.com/goccy/spidermonkeywasm2go/p7.Fn6979
func Fn6979(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6981 github.com/goccy/spidermonkeywasm2go/p7.Fn6981
func Fn6981(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6984 github.com/goccy/spidermonkeywasm2go/p6.Fn6984
func Fn6984(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7081 github.com/goccy/spidermonkeywasm2go/p4.Fn7081
func Fn7081(m *base.Module, l0 int32) int32

//go:linkname Fn7115 github.com/goccy/spidermonkeywasm2go/p6.Fn7115
func Fn7115(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7116 github.com/goccy/spidermonkeywasm2go/p6.Fn7116
func Fn7116(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7117 github.com/goccy/spidermonkeywasm2go/p6.Fn7117
func Fn7117(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7118 github.com/goccy/spidermonkeywasm2go/p7.Fn7118
func Fn7118(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7121 github.com/goccy/spidermonkeywasm2go/p5.Fn7121
func Fn7121(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7125 github.com/goccy/spidermonkeywasm2go/p5.Fn7125
func Fn7125(m *base.Module, l0 int32) int32

//go:linkname Fn7127 github.com/goccy/spidermonkeywasm2go/p4.Fn7127
func Fn7127(m *base.Module, l0 int32) int32

//go:linkname Fn7133 github.com/goccy/spidermonkeywasm2go/p6.Fn7133
func Fn7133(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7136 github.com/goccy/spidermonkeywasm2go/p6.Fn7136
func Fn7136(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7144 github.com/goccy/spidermonkeywasm2go/p6.Fn7144
func Fn7144(m *base.Module, l0 int32) int32

//go:linkname Fn7145 github.com/goccy/spidermonkeywasm2go/p6.Fn7145
func Fn7145(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn7162 github.com/goccy/spidermonkeywasm2go/p6.Fn7162
func Fn7162(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn7164 github.com/goccy/spidermonkeywasm2go/p6.Fn7164
func Fn7164(m *base.Module, l0 int32)

//go:linkname Fn7166 github.com/goccy/spidermonkeywasm2go/p5.Fn7166
func Fn7166(m *base.Module, l0 int32) int32

//go:linkname Fn7167 github.com/goccy/spidermonkeywasm2go/p6.Fn7167
func Fn7167(m *base.Module, l0 int32) int32

//go:linkname Fn7169 github.com/goccy/spidermonkeywasm2go/p7.Fn7169
func Fn7169(m *base.Module, l0 int32) int32

//go:linkname Fn7173 github.com/goccy/spidermonkeywasm2go/p7.Fn7173
func Fn7173(m *base.Module, l0 int32) int32

//go:linkname Fn7175 github.com/goccy/spidermonkeywasm2go/p4.Fn7175
func Fn7175(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7176 github.com/goccy/spidermonkeywasm2go/p6.Fn7176
func Fn7176(m *base.Module, l0 int32, l1 int32)

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

//go:linkname Fn7196 github.com/goccy/spidermonkeywasm2go/p6.Fn7196
func Fn7196(m *base.Module, l0 int32)

//go:linkname Fn7197 github.com/goccy/spidermonkeywasm2go/p6.Fn7197
func Fn7197(m *base.Module, l0 int32) int32

//go:linkname Fn7199 github.com/goccy/spidermonkeywasm2go/p6.Fn7199
func Fn7199(m *base.Module, l0 int32) int32

//go:linkname Fn7227 github.com/goccy/spidermonkeywasm2go/p5.Fn7227
func Fn7227(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7242 github.com/goccy/spidermonkeywasm2go/p6.Fn7242
func Fn7242(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7253 github.com/goccy/spidermonkeywasm2go/p6.Fn7253
func Fn7253(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7359 github.com/goccy/spidermonkeywasm2go/p7.Fn7359
func Fn7359(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7360 github.com/goccy/spidermonkeywasm2go/p5.Fn7360
func Fn7360(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7374 github.com/goccy/spidermonkeywasm2go/p6.Fn7374
func Fn7374(m *base.Module, l0 int32)

//go:linkname Fn7381 github.com/goccy/spidermonkeywasm2go/p6.Fn7381
func Fn7381(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7389 github.com/goccy/spidermonkeywasm2go/p6.Fn7389
func Fn7389(m *base.Module, l0 int32)

//go:linkname Fn7426 github.com/goccy/spidermonkeywasm2go/p1.Fn7426
func Fn7426(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7435 github.com/goccy/spidermonkeywasm2go/p2.Fn7435
func Fn7435(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7437 github.com/goccy/spidermonkeywasm2go/p4.Fn7437
func Fn7437(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7441 github.com/goccy/spidermonkeywasm2go/p7.Fn7441
func Fn7441(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7444 github.com/goccy/spidermonkeywasm2go/p6.Fn7444
func Fn7444(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7445 github.com/goccy/spidermonkeywasm2go/p5.Fn7445
func Fn7445(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7448 github.com/goccy/spidermonkeywasm2go/p7.Fn7448
func Fn7448(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7449 github.com/goccy/spidermonkeywasm2go/p4.Fn7449
func Fn7449(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7451 github.com/goccy/spidermonkeywasm2go/p6.Fn7451
func Fn7451(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7452 github.com/goccy/spidermonkeywasm2go/p6.Fn7452
func Fn7452(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7454 github.com/goccy/spidermonkeywasm2go/p7.Fn7454
func Fn7454(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7457 github.com/goccy/spidermonkeywasm2go/p6.Fn7457
func Fn7457(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7468 github.com/goccy/spidermonkeywasm2go/p5.Fn7468
func Fn7468(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7469 github.com/goccy/spidermonkeywasm2go/p6.Fn7469
func Fn7469(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7471 github.com/goccy/spidermonkeywasm2go/p6.Fn7471
func Fn7471(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7472 github.com/goccy/spidermonkeywasm2go/p5.Fn7472
func Fn7472(m *base.Module, l0 int32, l1 float64, l2 float64, l3 float64) int32

//go:linkname Fn7473 github.com/goccy/spidermonkeywasm2go/p4.Fn7473
func Fn7473(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7474 github.com/goccy/spidermonkeywasm2go/p6.Fn7474
func Fn7474(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7475 github.com/goccy/spidermonkeywasm2go/p6.Fn7475
func Fn7475(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn7476 github.com/goccy/spidermonkeywasm2go/p7.Fn7476
func Fn7476(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7477 github.com/goccy/spidermonkeywasm2go/p6.Fn7477
func Fn7477(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7480 github.com/goccy/spidermonkeywasm2go/p6.Fn7480
func Fn7480(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7481 github.com/goccy/spidermonkeywasm2go/p7.Fn7481
func Fn7481(m *base.Module, l0 int32, l1 int32, l2 float64)

//go:linkname Fn7482 github.com/goccy/spidermonkeywasm2go/p6.Fn7482
func Fn7482(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn7483 github.com/goccy/spidermonkeywasm2go/p5.Fn7483
func Fn7483(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7485 github.com/goccy/spidermonkeywasm2go/p5.Fn7485
func Fn7485(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int64, l6 int32) int32

//go:linkname Fn7486 github.com/goccy/spidermonkeywasm2go/p5.Fn7486
func Fn7486(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn7488 github.com/goccy/spidermonkeywasm2go/p6.Fn7488
func Fn7488(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7490 github.com/goccy/spidermonkeywasm2go/p5.Fn7490
func Fn7490(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn7491 github.com/goccy/spidermonkeywasm2go/p6.Fn7491
func Fn7491(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn7492 github.com/goccy/spidermonkeywasm2go/p7.Fn7492
func Fn7492(m *base.Module, l0 int32) int32

//go:linkname Fn7495 github.com/goccy/spidermonkeywasm2go/p6.Fn7495
func Fn7495(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7497 github.com/goccy/spidermonkeywasm2go/p4.Fn7497
func Fn7497(m *base.Module, l0 int32, l1 float64, l2 float64, l3 float64, l4 float64, l5 float64, l6 float64)

//go:linkname Fn7498 github.com/goccy/spidermonkeywasm2go/p7.Fn7498
func Fn7498(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7500 github.com/goccy/spidermonkeywasm2go/p7.Fn7500
func Fn7500(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7502 github.com/goccy/spidermonkeywasm2go/p6.Fn7502
func Fn7502(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32) int32

//go:linkname Fn7503 github.com/goccy/spidermonkeywasm2go/p7.Fn7503
func Fn7503(m *base.Module, l0 int32, l1 float64, l2 int32) int32

//go:linkname Fn7504 github.com/goccy/spidermonkeywasm2go/p4.Fn7504
func Fn7504(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7505 github.com/goccy/spidermonkeywasm2go/p2.Fn7505
func Fn7505(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7508 github.com/goccy/spidermonkeywasm2go/p7.Fn7508
func Fn7508(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7509 github.com/goccy/spidermonkeywasm2go/p6.Fn7509
func Fn7509(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn7511 github.com/goccy/spidermonkeywasm2go/p4.Fn7511
func Fn7511(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32

//go:linkname Fn7512 github.com/goccy/spidermonkeywasm2go/p6.Fn7512
func Fn7512(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7521 github.com/goccy/spidermonkeywasm2go/p6.Fn7521
func Fn7521(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7564 github.com/goccy/spidermonkeywasm2go/p5.Fn7564
func Fn7564(m *base.Module, l0 float64, l1 int32) int32

//go:linkname Fn7693 github.com/goccy/spidermonkeywasm2go/p7.Fn7693
func Fn7693(m *base.Module, l0 int32) int32

//go:linkname Fn7695 github.com/goccy/spidermonkeywasm2go/p6.Fn7695
func Fn7695(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7696 github.com/goccy/spidermonkeywasm2go/p5.Fn7696
func Fn7696(m *base.Module, l0 int32, l1 float64, l2 float64, l3 float64, l4 float64, l5 float64, l6 float64) int32

//go:linkname Fn7697 github.com/goccy/spidermonkeywasm2go/p5.Fn7697
func Fn7697(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn7698 github.com/goccy/spidermonkeywasm2go/p6.Fn7698
func Fn7698(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn7699 github.com/goccy/spidermonkeywasm2go/p6.Fn7699
func Fn7699(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7700 github.com/goccy/spidermonkeywasm2go/p6.Fn7700
func Fn7700(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7701 github.com/goccy/spidermonkeywasm2go/p4.Fn7701
func Fn7701(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn7707 github.com/goccy/spidermonkeywasm2go/p6.Fn7707
func Fn7707(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7708 github.com/goccy/spidermonkeywasm2go/p6.Fn7708
func Fn7708(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7713 github.com/goccy/spidermonkeywasm2go/p5.Fn7713
func Fn7713(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7714 github.com/goccy/spidermonkeywasm2go/p7.Fn7714
func Fn7714(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7715 github.com/goccy/spidermonkeywasm2go/p6.Fn7715
func Fn7715(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn7716 github.com/goccy/spidermonkeywasm2go/p7.Fn7716
func Fn7716(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7717 github.com/goccy/spidermonkeywasm2go/p6.Fn7717
func Fn7717(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7718 github.com/goccy/spidermonkeywasm2go/p4.Fn7718
func Fn7718(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7719 github.com/goccy/spidermonkeywasm2go/p7.Fn7719
func Fn7719(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7721 github.com/goccy/spidermonkeywasm2go/p5.Fn7721
func Fn7721(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7730 github.com/goccy/spidermonkeywasm2go/p5.Fn7730
func Fn7730(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn7917 github.com/goccy/spidermonkeywasm2go/p5.Fn7917
func Fn7917(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7918 github.com/goccy/spidermonkeywasm2go/p5.Fn7918
func Fn7918(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7920 github.com/goccy/spidermonkeywasm2go/p5.Fn7920
func Fn7920(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7921 github.com/goccy/spidermonkeywasm2go/p6.Fn7921
func Fn7921(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7922 github.com/goccy/spidermonkeywasm2go/p5.Fn7922
func Fn7922(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7923 github.com/goccy/spidermonkeywasm2go/p6.Fn7923
func Fn7923(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7924 github.com/goccy/spidermonkeywasm2go/p6.Fn7924
func Fn7924(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7925 github.com/goccy/spidermonkeywasm2go/p5.Fn7925
func Fn7925(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7926 github.com/goccy/spidermonkeywasm2go/p5.Fn7926
func Fn7926(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7927 github.com/goccy/spidermonkeywasm2go/p5.Fn7927
func Fn7927(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7928 github.com/goccy/spidermonkeywasm2go/p5.Fn7928
func Fn7928(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7929 github.com/goccy/spidermonkeywasm2go/p6.Fn7929
func Fn7929(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7930 github.com/goccy/spidermonkeywasm2go/p5.Fn7930
func Fn7930(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7931 github.com/goccy/spidermonkeywasm2go/p5.Fn7931
func Fn7931(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7932 github.com/goccy/spidermonkeywasm2go/p5.Fn7932
func Fn7932(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7933 github.com/goccy/spidermonkeywasm2go/p5.Fn7933
func Fn7933(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7935 github.com/goccy/spidermonkeywasm2go/p5.Fn7935
func Fn7935(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7936 github.com/goccy/spidermonkeywasm2go/p5.Fn7936
func Fn7936(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7943 github.com/goccy/spidermonkeywasm2go/p5.Fn7943
func Fn7943(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7944 github.com/goccy/spidermonkeywasm2go/p5.Fn7944
func Fn7944(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7945 github.com/goccy/spidermonkeywasm2go/p6.Fn7945
func Fn7945(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7946 github.com/goccy/spidermonkeywasm2go/p5.Fn7946
func Fn7946(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7947 github.com/goccy/spidermonkeywasm2go/p5.Fn7947
func Fn7947(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn7948 github.com/goccy/spidermonkeywasm2go/p6.Fn7948
func Fn7948(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7949 github.com/goccy/spidermonkeywasm2go/p5.Fn7949
func Fn7949(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn7952 github.com/goccy/spidermonkeywasm2go/p6.Fn7952
func Fn7952(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7953 github.com/goccy/spidermonkeywasm2go/p6.Fn7953
func Fn7953(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7954 github.com/goccy/spidermonkeywasm2go/p4.Fn7954
func Fn7954(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7955 github.com/goccy/spidermonkeywasm2go/p6.Fn7955
func Fn7955(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7958 github.com/goccy/spidermonkeywasm2go/p6.Fn7958
func Fn7958(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn7959 github.com/goccy/spidermonkeywasm2go/p5.Fn7959
func Fn7959(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn7960 github.com/goccy/spidermonkeywasm2go/p5.Fn7960
func Fn7960(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7961 github.com/goccy/spidermonkeywasm2go/p4.Fn7961
func Fn7961(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn7962 github.com/goccy/spidermonkeywasm2go/p7.Fn7962
func Fn7962(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7977 github.com/goccy/spidermonkeywasm2go/p4.Fn7977
func Fn7977(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn7978 github.com/goccy/spidermonkeywasm2go/p7.Fn7978
func Fn7978(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn7979 github.com/goccy/spidermonkeywasm2go/p5.Fn7979
func Fn7979(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7980 github.com/goccy/spidermonkeywasm2go/p6.Fn7980
func Fn7980(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn7981 github.com/goccy/spidermonkeywasm2go/p4.Fn7981
func Fn7981(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8000 github.com/goccy/spidermonkeywasm2go/p6.Fn8000
func Fn8000(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8097 github.com/goccy/spidermonkeywasm2go/p7.Fn8097
func Fn8097(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8103 github.com/goccy/spidermonkeywasm2go/p5.Fn8103
func Fn8103(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8106 github.com/goccy/spidermonkeywasm2go/p6.Fn8106
func Fn8106(m *base.Module, l0 int32)

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

//go:linkname Fn8120 github.com/goccy/spidermonkeywasm2go/p7.Fn8120
func Fn8120(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8121 github.com/goccy/spidermonkeywasm2go/p6.Fn8121
func Fn8121(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8122 github.com/goccy/spidermonkeywasm2go/p6.Fn8122
func Fn8122(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8143 github.com/goccy/spidermonkeywasm2go/p6.Fn8143
func Fn8143(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8144 github.com/goccy/spidermonkeywasm2go/p6.Fn8144
func Fn8144(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8145 github.com/goccy/spidermonkeywasm2go/p7.Fn8145
func Fn8145(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8148 github.com/goccy/spidermonkeywasm2go/p5.Fn8148
func Fn8148(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8149 github.com/goccy/spidermonkeywasm2go/p4.Fn8149
func Fn8149(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8154 github.com/goccy/spidermonkeywasm2go/p1.Fn8154
func Fn8154(m *base.Module, l0 int32) int32

//go:linkname Fn8157 github.com/goccy/spidermonkeywasm2go/p6.Fn8157
func Fn8157(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8159 github.com/goccy/spidermonkeywasm2go/p1.Fn8159
func Fn8159(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8160 github.com/goccy/spidermonkeywasm2go/p7.Fn8160
func Fn8160(m *base.Module, l0 int32) int32

//go:linkname Fn8162 github.com/goccy/spidermonkeywasm2go/p7.Fn8162
func Fn8162(m *base.Module, l0 int32) int32

//go:linkname Fn8164 github.com/goccy/spidermonkeywasm2go/p1.Fn8164
func Fn8164(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8165 github.com/goccy/spidermonkeywasm2go/p6.Fn8165
func Fn8165(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8166 github.com/goccy/spidermonkeywasm2go/p7.Fn8166
func Fn8166(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8168 github.com/goccy/spidermonkeywasm2go/p7.Fn8168
func Fn8168(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8169 github.com/goccy/spidermonkeywasm2go/p1.Fn8169
func Fn8169(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8170 github.com/goccy/spidermonkeywasm2go/p1.Fn8170
func Fn8170(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8171 github.com/goccy/spidermonkeywasm2go/p2.Fn8171
func Fn8171(m *base.Module) int32

//go:linkname Fn8175 github.com/goccy/spidermonkeywasm2go/p6.Fn8175
func Fn8175(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8177 github.com/goccy/spidermonkeywasm2go/p7.Fn8177
func Fn8177(m *base.Module, l0 int32) int32

//go:linkname Fn8189 github.com/goccy/spidermonkeywasm2go/p7.Fn8189
func Fn8189(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8194 github.com/goccy/spidermonkeywasm2go/p1.Fn8194
func Fn8194(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8195 github.com/goccy/spidermonkeywasm2go/p6.Fn8195
func Fn8195(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32)

//go:linkname Fn8197 github.com/goccy/spidermonkeywasm2go/p5.Fn8197
func Fn8197(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8199 github.com/goccy/spidermonkeywasm2go/p1.Fn8199
func Fn8199(m *base.Module, l0 int32)

//go:linkname Fn8204 github.com/goccy/spidermonkeywasm2go/p1.Fn8204
func Fn8204(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8205 github.com/goccy/spidermonkeywasm2go/p7.Fn8205
func Fn8205(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8214 github.com/goccy/spidermonkeywasm2go/p5.Fn8214
func Fn8214(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8216 github.com/goccy/spidermonkeywasm2go/p7.Fn8216
func Fn8216(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8220 github.com/goccy/spidermonkeywasm2go/p6.Fn8220
func Fn8220(m *base.Module, l0 int32) int32

//go:linkname Fn8222 github.com/goccy/spidermonkeywasm2go/p4.Fn8222
func Fn8222(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8223 github.com/goccy/spidermonkeywasm2go/p6.Fn8223
func Fn8223(m *base.Module, l0 int32) int32

//go:linkname Fn8266 github.com/goccy/spidermonkeywasm2go/p6.Fn8266
func Fn8266(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8267 github.com/goccy/spidermonkeywasm2go/p5.Fn8267
func Fn8267(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8269 github.com/goccy/spidermonkeywasm2go/p5.Fn8269
func Fn8269(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8270 github.com/goccy/spidermonkeywasm2go/p6.Fn8270
func Fn8270(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8274 github.com/goccy/spidermonkeywasm2go/p5.Fn8274
func Fn8274(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8281 github.com/goccy/spidermonkeywasm2go/p5.Fn8281
func Fn8281(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8283 github.com/goccy/spidermonkeywasm2go/p4.Fn8283
func Fn8283(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8286 github.com/goccy/spidermonkeywasm2go/p4.Fn8286
func Fn8286(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8290 github.com/goccy/spidermonkeywasm2go/p6.Fn8290
func Fn8290(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8294 github.com/goccy/spidermonkeywasm2go/p5.Fn8294
func Fn8294(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8297 github.com/goccy/spidermonkeywasm2go/p6.Fn8297
func Fn8297(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8298 github.com/goccy/spidermonkeywasm2go/p6.Fn8298
func Fn8298(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8299 github.com/goccy/spidermonkeywasm2go/p7.Fn8299
func Fn8299(m *base.Module) float64

//go:linkname Fn8301 github.com/goccy/spidermonkeywasm2go/p7.Fn8301
func Fn8301(m *base.Module, l0 int32) int32

//go:linkname Fn8302 github.com/goccy/spidermonkeywasm2go/p5.Fn8302
func Fn8302(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8303 github.com/goccy/spidermonkeywasm2go/p6.Fn8303
func Fn8303(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8305 github.com/goccy/spidermonkeywasm2go/p6.Fn8305
func Fn8305(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn8308 github.com/goccy/spidermonkeywasm2go/p5.Fn8308
func Fn8308(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8311 github.com/goccy/spidermonkeywasm2go/p6.Fn8311
func Fn8311(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8312 github.com/goccy/spidermonkeywasm2go/p5.Fn8312
func Fn8312(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8313 github.com/goccy/spidermonkeywasm2go/p6.Fn8313
func Fn8313(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8317 github.com/goccy/spidermonkeywasm2go/p6.Fn8317
func Fn8317(m *base.Module, l0 int32) int32

//go:linkname Fn8320 github.com/goccy/spidermonkeywasm2go/p7.Fn8320
func Fn8320(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8321 github.com/goccy/spidermonkeywasm2go/p7.Fn8321
func Fn8321(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8322 github.com/goccy/spidermonkeywasm2go/p6.Fn8322
func Fn8322(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8340 github.com/goccy/spidermonkeywasm2go/p4.Fn8340
func Fn8340(m *base.Module, l0 int32) int32

//go:linkname Fn8342 github.com/goccy/spidermonkeywasm2go/p7.Fn8342
func Fn8342(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8344 github.com/goccy/spidermonkeywasm2go/p7.Fn8344
func Fn8344(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8346 github.com/goccy/spidermonkeywasm2go/p7.Fn8346
func Fn8346(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8349 github.com/goccy/spidermonkeywasm2go/p6.Fn8349
func Fn8349(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8353 github.com/goccy/spidermonkeywasm2go/p5.Fn8353
func Fn8353(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32

//go:linkname Fn8354 github.com/goccy/spidermonkeywasm2go/p4.Fn8354
func Fn8354(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8362 github.com/goccy/spidermonkeywasm2go/p6.Fn8362
func Fn8362(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8365 github.com/goccy/spidermonkeywasm2go/p6.Fn8365
func Fn8365(m *base.Module, l0 int32)

//go:linkname Fn8367 github.com/goccy/spidermonkeywasm2go/p6.Fn8367
func Fn8367(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8368 github.com/goccy/spidermonkeywasm2go/p5.Fn8368
func Fn8368(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8370 github.com/goccy/spidermonkeywasm2go/p6.Fn8370
func Fn8370(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8381 github.com/goccy/spidermonkeywasm2go/p5.Fn8381
func Fn8381(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn8390 github.com/goccy/spidermonkeywasm2go/p1.Fn8390
func Fn8390(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8391 github.com/goccy/spidermonkeywasm2go/p6.Fn8391
func Fn8391(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8396 github.com/goccy/spidermonkeywasm2go/p7.Fn8396
func Fn8396(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8397 github.com/goccy/spidermonkeywasm2go/p4.Fn8397
func Fn8397(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8402 github.com/goccy/spidermonkeywasm2go/p1.Fn8402
func Fn8402(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8415 github.com/goccy/spidermonkeywasm2go/p6.Fn8415
func Fn8415(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8434 github.com/goccy/spidermonkeywasm2go/p6.Fn8434
func Fn8434(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8435 github.com/goccy/spidermonkeywasm2go/p6.Fn8435
func Fn8435(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8437 github.com/goccy/spidermonkeywasm2go/p6.Fn8437
func Fn8437(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8438 github.com/goccy/spidermonkeywasm2go/p4.Fn8438
func Fn8438(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8439 github.com/goccy/spidermonkeywasm2go/p5.Fn8439
func Fn8439(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8442 github.com/goccy/spidermonkeywasm2go/p6.Fn8442
func Fn8442(m *base.Module, l0 int32) int32

//go:linkname Fn8448 github.com/goccy/spidermonkeywasm2go/p6.Fn8448
func Fn8448(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8456 github.com/goccy/spidermonkeywasm2go/p6.Fn8456
func Fn8456(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8461 github.com/goccy/spidermonkeywasm2go/p6.Fn8461
func Fn8461(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8463 github.com/goccy/spidermonkeywasm2go/p6.Fn8463
func Fn8463(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn8464 github.com/goccy/spidermonkeywasm2go/p6.Fn8464
func Fn8464(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8465 github.com/goccy/spidermonkeywasm2go/p5.Fn8465
func Fn8465(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8474 github.com/goccy/spidermonkeywasm2go/p1.Fn8474
func Fn8474(m *base.Module, l0 int32) int32

//go:linkname Fn8490 github.com/goccy/spidermonkeywasm2go/p5.Fn8490
func Fn8490(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8497 github.com/goccy/spidermonkeywasm2go/p7.Fn8497
func Fn8497(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8498 github.com/goccy/spidermonkeywasm2go/p7.Fn8498
func Fn8498(m *base.Module, l0 int32)

//go:linkname Fn8499 github.com/goccy/spidermonkeywasm2go/p5.Fn8499
func Fn8499(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8501 github.com/goccy/spidermonkeywasm2go/p6.Fn8501
func Fn8501(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8521 github.com/goccy/spidermonkeywasm2go/p6.Fn8521
func Fn8521(m *base.Module, l0 int32, l1 int32, l2 int32)

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

//go:linkname Fn8532 github.com/goccy/spidermonkeywasm2go/p5.Fn8532
func Fn8532(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8534 github.com/goccy/spidermonkeywasm2go/p7.Fn8534
func Fn8534(m *base.Module, l0 int32) int32

//go:linkname Fn8536 github.com/goccy/spidermonkeywasm2go/p7.Fn8536
func Fn8536(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8539 github.com/goccy/spidermonkeywasm2go/p6.Fn8539
func Fn8539(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8542 github.com/goccy/spidermonkeywasm2go/p7.Fn8542
func Fn8542(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8543 github.com/goccy/spidermonkeywasm2go/p6.Fn8543
func Fn8543(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8548 github.com/goccy/spidermonkeywasm2go/p6.Fn8548
func Fn8548(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8550 github.com/goccy/spidermonkeywasm2go/p5.Fn8550
func Fn8550(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8551 github.com/goccy/spidermonkeywasm2go/p6.Fn8551
func Fn8551(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8552 github.com/goccy/spidermonkeywasm2go/p6.Fn8552
func Fn8552(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8553 github.com/goccy/spidermonkeywasm2go/p6.Fn8553
func Fn8553(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn8557 github.com/goccy/spidermonkeywasm2go/p7.Fn8557
func Fn8557(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8561 github.com/goccy/spidermonkeywasm2go/p6.Fn8561
func Fn8561(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8570 github.com/goccy/spidermonkeywasm2go/p4.Fn8570
func Fn8570(m *base.Module, l0 int32) int32

//go:linkname Fn8573 github.com/goccy/spidermonkeywasm2go/p7.Fn8573
func Fn8573(m *base.Module, l0 int32) int32

//go:linkname Fn8625 github.com/goccy/spidermonkeywasm2go/p7.Fn8625
func Fn8625(m *base.Module, l0 int32)

//go:linkname Fn8626 github.com/goccy/spidermonkeywasm2go/p7.Fn8626
func Fn8626(m *base.Module, l0 int32)

//go:linkname Fn8629 github.com/goccy/spidermonkeywasm2go/p6.Fn8629
func Fn8629(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8632 github.com/goccy/spidermonkeywasm2go/p6.Fn8632
func Fn8632(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8633 github.com/goccy/spidermonkeywasm2go/p1.Fn8633
func Fn8633(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8635 github.com/goccy/spidermonkeywasm2go/p6.Fn8635
func Fn8635(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8636 github.com/goccy/spidermonkeywasm2go/p1.Fn8636
func Fn8636(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8637 github.com/goccy/spidermonkeywasm2go/p6.Fn8637
func Fn8637(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8638 github.com/goccy/spidermonkeywasm2go/p1.Fn8638
func Fn8638(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8640 github.com/goccy/spidermonkeywasm2go/p6.Fn8640
func Fn8640(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8642 github.com/goccy/spidermonkeywasm2go/p6.Fn8642
func Fn8642(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8643 github.com/goccy/spidermonkeywasm2go/p1.Fn8643
func Fn8643(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8645 github.com/goccy/spidermonkeywasm2go/p1.Fn8645
func Fn8645(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8654 github.com/goccy/spidermonkeywasm2go/p2.Fn8654
func Fn8654(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn8685 github.com/goccy/spidermonkeywasm2go/p6.Fn8685
func Fn8685(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn8691 github.com/goccy/spidermonkeywasm2go/p4.Fn8691
func Fn8691(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn8692 github.com/goccy/spidermonkeywasm2go/p6.Fn8692
func Fn8692(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8693 github.com/goccy/spidermonkeywasm2go/p6.Fn8693
func Fn8693(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn8705 github.com/goccy/spidermonkeywasm2go/p4.Fn8705
func Fn8705(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8710 github.com/goccy/spidermonkeywasm2go/p5.Fn8710
func Fn8710(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn8711 github.com/goccy/spidermonkeywasm2go/p5.Fn8711
func Fn8711(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8716 github.com/goccy/spidermonkeywasm2go/p6.Fn8716
func Fn8716(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8723 github.com/goccy/spidermonkeywasm2go/p7.Fn8723
func Fn8723(m *base.Module, l0 int32)

//go:linkname Fn8726 github.com/goccy/spidermonkeywasm2go/p6.Fn8726
func Fn8726(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8729 github.com/goccy/spidermonkeywasm2go/p6.Fn8729
func Fn8729(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn8754 github.com/goccy/spidermonkeywasm2go/p5.Fn8754
func Fn8754(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8755 github.com/goccy/spidermonkeywasm2go/p7.Fn8755
func Fn8755(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn8758 github.com/goccy/spidermonkeywasm2go/p6.Fn8758
func Fn8758(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8761 github.com/goccy/spidermonkeywasm2go/p7.Fn8761
func Fn8761(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn8833 github.com/goccy/spidermonkeywasm2go/p7.Fn8833
func Fn8833(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn8868 github.com/goccy/spidermonkeywasm2go/p6.Fn8868
func Fn8868(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn8914 github.com/goccy/spidermonkeywasm2go/p6.Fn8914
func Fn8914(m *base.Module, l0 int32) int32

//go:linkname Fn8994 github.com/goccy/spidermonkeywasm2go/p7.Fn8994
func Fn8994(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9010 github.com/goccy/spidermonkeywasm2go/p4.Fn9010
func Fn9010(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9025 github.com/goccy/spidermonkeywasm2go/p7.Fn9025
func Fn9025(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9026 github.com/goccy/spidermonkeywasm2go/p4.Fn9026
func Fn9026(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9028 github.com/goccy/spidermonkeywasm2go/p7.Fn9028
func Fn9028(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9054 github.com/goccy/spidermonkeywasm2go/p6.Fn9054
func Fn9054(m *base.Module, l0 int32)

//go:linkname Fn9060 github.com/goccy/spidermonkeywasm2go/p6.Fn9060
func Fn9060(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9068 github.com/goccy/spidermonkeywasm2go/p2.Fn9068
func Fn9068(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9075 github.com/goccy/spidermonkeywasm2go/p4.Fn9075
func Fn9075(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9076 github.com/goccy/spidermonkeywasm2go/p6.Fn9076
func Fn9076(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9084 github.com/goccy/spidermonkeywasm2go/p5.Fn9084
func Fn9084(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn9086 github.com/goccy/spidermonkeywasm2go/p4.Fn9086
func Fn9086(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn9098 github.com/goccy/spidermonkeywasm2go/p6.Fn9098
func Fn9098(m *base.Module, l0 int32) int32

//go:linkname Fn9100 github.com/goccy/spidermonkeywasm2go/p2.Fn9100
func Fn9100(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9103 github.com/goccy/spidermonkeywasm2go/p5.Fn9103
func Fn9103(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9125 github.com/goccy/spidermonkeywasm2go/p6.Fn9125
func Fn9125(m *base.Module, l0 int32)

//go:linkname Fn9126 github.com/goccy/spidermonkeywasm2go/p5.Fn9126
func Fn9126(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9129 github.com/goccy/spidermonkeywasm2go/p5.Fn9129
func Fn9129(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9133 github.com/goccy/spidermonkeywasm2go/p6.Fn9133
func Fn9133(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9135 github.com/goccy/spidermonkeywasm2go/p6.Fn9135
func Fn9135(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9140 github.com/goccy/spidermonkeywasm2go/p6.Fn9140
func Fn9140(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9176 github.com/goccy/spidermonkeywasm2go/p7.Fn9176
func Fn9176(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn9177 github.com/goccy/spidermonkeywasm2go/p6.Fn9177
func Fn9177(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn9273 github.com/goccy/spidermonkeywasm2go/p6.Fn9273
func Fn9273(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9290 github.com/goccy/spidermonkeywasm2go/p6.Fn9290
func Fn9290(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9291 github.com/goccy/spidermonkeywasm2go/p6.Fn9291
func Fn9291(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9292 github.com/goccy/spidermonkeywasm2go/p5.Fn9292
func Fn9292(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9293 github.com/goccy/spidermonkeywasm2go/p6.Fn9293
func Fn9293(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn9295 github.com/goccy/spidermonkeywasm2go/p6.Fn9295
func Fn9295(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

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

//go:linkname Fn9308 github.com/goccy/spidermonkeywasm2go/p6.Fn9308
func Fn9308(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9309 github.com/goccy/spidermonkeywasm2go/p6.Fn9309
func Fn9309(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9310 github.com/goccy/spidermonkeywasm2go/p5.Fn9310
func Fn9310(m *base.Module, l0 int32)

//go:linkname Fn9311 github.com/goccy/spidermonkeywasm2go/p6.Fn9311
func Fn9311(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn9312 github.com/goccy/spidermonkeywasm2go/p5.Fn9312
func Fn9312(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9316 github.com/goccy/spidermonkeywasm2go/p6.Fn9316
func Fn9316(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9318 github.com/goccy/spidermonkeywasm2go/p5.Fn9318
func Fn9318(m *base.Module, l0 int32) float64

//go:linkname Fn9323 github.com/goccy/spidermonkeywasm2go/p6.Fn9323
func Fn9323(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn9324 github.com/goccy/spidermonkeywasm2go/p6.Fn9324
func Fn9324(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9325 github.com/goccy/spidermonkeywasm2go/p6.Fn9325
func Fn9325(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9327 github.com/goccy/spidermonkeywasm2go/p7.Fn9327
func Fn9327(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9329 github.com/goccy/spidermonkeywasm2go/p6.Fn9329
func Fn9329(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9330 github.com/goccy/spidermonkeywasm2go/p6.Fn9330
func Fn9330(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9332 github.com/goccy/spidermonkeywasm2go/p5.Fn9332
func Fn9332(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9336 github.com/goccy/spidermonkeywasm2go/p6.Fn9336
func Fn9336(m *base.Module, l0 int32) int32

//go:linkname Fn9338 github.com/goccy/spidermonkeywasm2go/p5.Fn9338
func Fn9338(m *base.Module, l0 int32) int32

//go:linkname Fn9340 github.com/goccy/spidermonkeywasm2go/p5.Fn9340
func Fn9340(m *base.Module, l0 int32)

//go:linkname Fn9346 github.com/goccy/spidermonkeywasm2go/p1.Fn9346
func Fn9346(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9347 github.com/goccy/spidermonkeywasm2go/p4.Fn9347
func Fn9347(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn9356 github.com/goccy/spidermonkeywasm2go/p4.Fn9356
func Fn9356(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9359 github.com/goccy/spidermonkeywasm2go/p7.Fn9359
func Fn9359(m *base.Module, l0 int32)

//go:linkname Fn9381 github.com/goccy/spidermonkeywasm2go/p6.Fn9381
func Fn9381(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9382 github.com/goccy/spidermonkeywasm2go/p6.Fn9382
func Fn9382(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9437 github.com/goccy/spidermonkeywasm2go/p5.Fn9437
func Fn9437(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9487 github.com/goccy/spidermonkeywasm2go/p4.Fn9487
func Fn9487(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9498 github.com/goccy/spidermonkeywasm2go/p5.Fn9498
func Fn9498(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9499 github.com/goccy/spidermonkeywasm2go/p5.Fn9499
func Fn9499(m *base.Module, l0 int32)

//go:linkname Fn9515 github.com/goccy/spidermonkeywasm2go/p6.Fn9515
func Fn9515(m *base.Module, l0 int32) int32

//go:linkname Fn9519 github.com/goccy/spidermonkeywasm2go/p5.Fn9519
func Fn9519(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9520 github.com/goccy/spidermonkeywasm2go/p7.Fn9520
func Fn9520(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9522 github.com/goccy/spidermonkeywasm2go/p6.Fn9522
func Fn9522(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9523 github.com/goccy/spidermonkeywasm2go/p6.Fn9523
func Fn9523(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9525 github.com/goccy/spidermonkeywasm2go/p6.Fn9525
func Fn9525(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9527 github.com/goccy/spidermonkeywasm2go/p5.Fn9527
func Fn9527(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn9533 github.com/goccy/spidermonkeywasm2go/p5.Fn9533
func Fn9533(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9536 github.com/goccy/spidermonkeywasm2go/p6.Fn9536
func Fn9536(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9552 github.com/goccy/spidermonkeywasm2go/p5.Fn9552
func Fn9552(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9571 github.com/goccy/spidermonkeywasm2go/p7.Fn9571
func Fn9571(m *base.Module, l0 int32) int32

//go:linkname Fn9575 github.com/goccy/spidermonkeywasm2go/p6.Fn9575
func Fn9575(m *base.Module, l0 int32) int32

//go:linkname Fn9577 github.com/goccy/spidermonkeywasm2go/p5.Fn9577
func Fn9577(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn9607 github.com/goccy/spidermonkeywasm2go/p6.Fn9607
func Fn9607(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9620 github.com/goccy/spidermonkeywasm2go/p6.Fn9620
func Fn9620(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9621 github.com/goccy/spidermonkeywasm2go/p6.Fn9621
func Fn9621(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9626 github.com/goccy/spidermonkeywasm2go/p5.Fn9626
func Fn9626(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn9627 github.com/goccy/spidermonkeywasm2go/p7.Fn9627
func Fn9627(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9629 github.com/goccy/spidermonkeywasm2go/p5.Fn9629
func Fn9629(m *base.Module, l0 int32)

//go:linkname Fn9635 github.com/goccy/spidermonkeywasm2go/p5.Fn9635
func Fn9635(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9646 github.com/goccy/spidermonkeywasm2go/p1.Fn9646
func Fn9646(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9654 github.com/goccy/spidermonkeywasm2go/p6.Fn9654
func Fn9654(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9673 github.com/goccy/spidermonkeywasm2go/p7.Fn9673
func Fn9673(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9686 github.com/goccy/spidermonkeywasm2go/p6.Fn9686
func Fn9686(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9691 github.com/goccy/spidermonkeywasm2go/p4.Fn9691
func Fn9691(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn9692 github.com/goccy/spidermonkeywasm2go/p4.Fn9692
func Fn9692(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9696 github.com/goccy/spidermonkeywasm2go/p6.Fn9696
func Fn9696(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn9697 github.com/goccy/spidermonkeywasm2go/p5.Fn9697
func Fn9697(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn9699 github.com/goccy/spidermonkeywasm2go/p7.Fn9699
func Fn9699(m *base.Module, l0 int32)

//go:linkname Fn9701 github.com/goccy/spidermonkeywasm2go/p5.Fn9701
func Fn9701(m *base.Module, l0 int32, l1 float64, l2 float64, l3 int32, l4 int32)

//go:linkname Fn9702 github.com/goccy/spidermonkeywasm2go/p6.Fn9702
func Fn9702(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn9703 github.com/goccy/spidermonkeywasm2go/p7.Fn9703
func Fn9703(m *base.Module, l0 int32)

//go:linkname Fn9705 github.com/goccy/spidermonkeywasm2go/p7.Fn9705
func Fn9705(m *base.Module, l0 int32)

//go:linkname Fn9710 github.com/goccy/spidermonkeywasm2go/p4.Fn9710
func Fn9710(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9717 github.com/goccy/spidermonkeywasm2go/p6.Fn9717
func Fn9717(m *base.Module, l0 int32)

//go:linkname Fn9758 github.com/goccy/spidermonkeywasm2go/p6.Fn9758
func Fn9758(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9767 github.com/goccy/spidermonkeywasm2go/p6.Fn9767
func Fn9767(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9769 github.com/goccy/spidermonkeywasm2go/p7.Fn9769
func Fn9769(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9770 github.com/goccy/spidermonkeywasm2go/p7.Fn9770
func Fn9770(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn9771 github.com/goccy/spidermonkeywasm2go/p6.Fn9771
func Fn9771(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9772 github.com/goccy/spidermonkeywasm2go/p5.Fn9772
func Fn9772(m *base.Module, l0 int32, l1 int32, l2 float64)

//go:linkname Fn9773 github.com/goccy/spidermonkeywasm2go/p7.Fn9773
func Fn9773(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9852 github.com/goccy/spidermonkeywasm2go/p6.Fn9852
func Fn9852(m *base.Module, l0 int32) int32

//go:linkname Fn9853 github.com/goccy/spidermonkeywasm2go/p2.Fn9853
func Fn9853(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn9855 github.com/goccy/spidermonkeywasm2go/p7.Fn9855
func Fn9855(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9944 github.com/goccy/spidermonkeywasm2go/p7.Fn9944
func Fn9944(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9946 github.com/goccy/spidermonkeywasm2go/p7.Fn9946
func Fn9946(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9958 github.com/goccy/spidermonkeywasm2go/p7.Fn9958
func Fn9958(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn9959 github.com/goccy/spidermonkeywasm2go/p7.Fn9959
func Fn9959(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn9960 github.com/goccy/spidermonkeywasm2go/p6.Fn9960
func Fn9960(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn9973 github.com/goccy/spidermonkeywasm2go/p5.Fn9973
func Fn9973(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn9999 github.com/goccy/spidermonkeywasm2go/p2.Fn9999
func Fn9999(m *base.Module, l0 int32) int32

//go:linkname Fn10026 github.com/goccy/spidermonkeywasm2go/p6.Fn10026
func Fn10026(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10029 github.com/goccy/spidermonkeywasm2go/p5.Fn10029
func Fn10029(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10093 github.com/goccy/spidermonkeywasm2go/p2.Fn10093
func Fn10093(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn10109 github.com/goccy/spidermonkeywasm2go/p5.Fn10109
func Fn10109(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn10112 github.com/goccy/spidermonkeywasm2go/p6.Fn10112
func Fn10112(m *base.Module, l0 int32) int32

//go:linkname Fn10119 github.com/goccy/spidermonkeywasm2go/p6.Fn10119
func Fn10119(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10120 github.com/goccy/spidermonkeywasm2go/p7.Fn10120
func Fn10120(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn10128 github.com/goccy/spidermonkeywasm2go/p7.Fn10128
func Fn10128(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10131 github.com/goccy/spidermonkeywasm2go/p5.Fn10131
func Fn10131(m *base.Module, l0 int32) int32

//go:linkname Fn10136 github.com/goccy/spidermonkeywasm2go/p5.Fn10136
func Fn10136(m *base.Module, l0 int32) int32

//go:linkname Fn10140 github.com/goccy/spidermonkeywasm2go/p6.Fn10140
func Fn10140(m *base.Module, l0 int32)

//go:linkname Fn10144 github.com/goccy/spidermonkeywasm2go/p7.Fn10144
func Fn10144(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10145 github.com/goccy/spidermonkeywasm2go/p6.Fn10145
func Fn10145(m *base.Module, l0 int32) int32

//go:linkname Fn10149 github.com/goccy/spidermonkeywasm2go/p6.Fn10149
func Fn10149(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10155 github.com/goccy/spidermonkeywasm2go/p5.Fn10155
func Fn10155(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10159 github.com/goccy/spidermonkeywasm2go/p4.Fn10159
func Fn10159(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10162 github.com/goccy/spidermonkeywasm2go/p2.Fn10162
func Fn10162(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn10200 github.com/goccy/spidermonkeywasm2go/p5.Fn10200
func Fn10200(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10206 github.com/goccy/spidermonkeywasm2go/p7.Fn10206
func Fn10206(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10209 github.com/goccy/spidermonkeywasm2go/p6.Fn10209
func Fn10209(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn10211 github.com/goccy/spidermonkeywasm2go/p5.Fn10211
func Fn10211(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10212 github.com/goccy/spidermonkeywasm2go/p5.Fn10212
func Fn10212(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10219 github.com/goccy/spidermonkeywasm2go/p5.Fn10219
func Fn10219(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10224 github.com/goccy/spidermonkeywasm2go/p6.Fn10224
func Fn10224(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10225 github.com/goccy/spidermonkeywasm2go/p7.Fn10225
func Fn10225(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10226 github.com/goccy/spidermonkeywasm2go/p4.Fn10226
func Fn10226(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10238 github.com/goccy/spidermonkeywasm2go/p2.Fn10238
func Fn10238(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10239 github.com/goccy/spidermonkeywasm2go/p6.Fn10239
func Fn10239(m *base.Module, l0 int32)

//go:linkname Fn10244 github.com/goccy/spidermonkeywasm2go/p6.Fn10244
func Fn10244(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10249 github.com/goccy/spidermonkeywasm2go/p5.Fn10249
func Fn10249(m *base.Module, l0 int32) int32

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

//go:linkname Fn10340 github.com/goccy/spidermonkeywasm2go/p1.Fn10340
func Fn10340(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn10342 github.com/goccy/spidermonkeywasm2go/p5.Fn10342
func Fn10342(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10350 github.com/goccy/spidermonkeywasm2go/p6.Fn10350
func Fn10350(m *base.Module, l0 int32) int32

//go:linkname Fn10356 github.com/goccy/spidermonkeywasm2go/p5.Fn10356
func Fn10356(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10401 github.com/goccy/spidermonkeywasm2go/p6.Fn10401
func Fn10401(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10420 github.com/goccy/spidermonkeywasm2go/p4.Fn10420
func Fn10420(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10431 github.com/goccy/spidermonkeywasm2go/p4.Fn10431
func Fn10431(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10440 github.com/goccy/spidermonkeywasm2go/p7.Fn10440
func Fn10440(m *base.Module, l0 int32)

//go:linkname Fn10462 github.com/goccy/spidermonkeywasm2go/p5.Fn10462
func Fn10462(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10471 github.com/goccy/spidermonkeywasm2go/p5.Fn10471
func Fn10471(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10503 github.com/goccy/spidermonkeywasm2go/p6.Fn10503
func Fn10503(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10520 github.com/goccy/spidermonkeywasm2go/p4.Fn10520
func Fn10520(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10521 github.com/goccy/spidermonkeywasm2go/p6.Fn10521
func Fn10521(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10560 github.com/goccy/spidermonkeywasm2go/p5.Fn10560
func Fn10560(m *base.Module, l0 int32) int32

//go:linkname Fn10578 github.com/goccy/spidermonkeywasm2go/p7.Fn10578
func Fn10578(m *base.Module, l0 int32) int32

//go:linkname Fn10583 github.com/goccy/spidermonkeywasm2go/p5.Fn10583
func Fn10583(m *base.Module) int32

//go:linkname Fn10584 github.com/goccy/spidermonkeywasm2go/p7.Fn10584
func Fn10584(m *base.Module) int32

//go:linkname Fn10585 github.com/goccy/spidermonkeywasm2go/p6.Fn10585
func Fn10585(m *base.Module, l0 int32) int32

//go:linkname Fn10590 github.com/goccy/spidermonkeywasm2go/p4.Fn10590
func Fn10590(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10593 github.com/goccy/spidermonkeywasm2go/p6.Fn10593
func Fn10593(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10633 github.com/goccy/spidermonkeywasm2go/p6.Fn10633
func Fn10633(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10640 github.com/goccy/spidermonkeywasm2go/p7.Fn10640
func Fn10640(m *base.Module, l0 int32)

//go:linkname Fn10641 github.com/goccy/spidermonkeywasm2go/p5.Fn10641
func Fn10641(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10645 github.com/goccy/spidermonkeywasm2go/p4.Fn10645
func Fn10645(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn10647 github.com/goccy/spidermonkeywasm2go/p7.Fn10647
func Fn10647(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10657 github.com/goccy/spidermonkeywasm2go/p5.Fn10657
func Fn10657(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10758 github.com/goccy/spidermonkeywasm2go/p6.Fn10758
func Fn10758(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10777 github.com/goccy/spidermonkeywasm2go/p6.Fn10777
func Fn10777(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10789 github.com/goccy/spidermonkeywasm2go/p4.Fn10789
func Fn10789(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn10793 github.com/goccy/spidermonkeywasm2go/p5.Fn10793
func Fn10793(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn10824 github.com/goccy/spidermonkeywasm2go/p5.Fn10824
func Fn10824(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn10826 github.com/goccy/spidermonkeywasm2go/p6.Fn10826
func Fn10826(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10827 github.com/goccy/spidermonkeywasm2go/p4.Fn10827
func Fn10827(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10830 github.com/goccy/spidermonkeywasm2go/p6.Fn10830
func Fn10830(m *base.Module, l0 int32, l1 float64) float64

//go:linkname Fn10831 github.com/goccy/spidermonkeywasm2go/p6.Fn10831
func Fn10831(m *base.Module, l0 int32, l1 float64) float64

//go:linkname Fn10882 github.com/goccy/spidermonkeywasm2go/p4.Fn10882
func Fn10882(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10884 github.com/goccy/spidermonkeywasm2go/p7.Fn10884
func Fn10884(m *base.Module, l0 int32) int32

//go:linkname Fn10917 github.com/goccy/spidermonkeywasm2go/p5.Fn10917
func Fn10917(m *base.Module, l0 int32)

//go:linkname Fn10921 github.com/goccy/spidermonkeywasm2go/p6.Fn10921
func Fn10921(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10928 github.com/goccy/spidermonkeywasm2go/p6.Fn10928
func Fn10928(m *base.Module, l0 int32)

//go:linkname Fn10929 github.com/goccy/spidermonkeywasm2go/p6.Fn10929
func Fn10929(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn10933 github.com/goccy/spidermonkeywasm2go/p5.Fn10933
func Fn10933(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10934 github.com/goccy/spidermonkeywasm2go/p4.Fn10934
func Fn10934(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn10936 github.com/goccy/spidermonkeywasm2go/p7.Fn10936
func Fn10936(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn10938 github.com/goccy/spidermonkeywasm2go/p7.Fn10938
func Fn10938(m *base.Module)

//go:linkname Fn10940 github.com/goccy/spidermonkeywasm2go/p2.Fn10940
func Fn10940(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10941 github.com/goccy/spidermonkeywasm2go/p2.Fn10941
func Fn10941(m *base.Module, l0 float64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn10946 github.com/goccy/spidermonkeywasm2go/p4.Fn10946
func Fn10946(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10947 github.com/goccy/spidermonkeywasm2go/p5.Fn10947
func Fn10947(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn10969 github.com/goccy/spidermonkeywasm2go/p2.Fn10969
func Fn10969(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn10973 github.com/goccy/spidermonkeywasm2go/p7.Fn10973
func Fn10973(m *base.Module, l0 int32)

//go:linkname Fn10974 github.com/goccy/spidermonkeywasm2go/p7.Fn10974
func Fn10974(m *base.Module, l0 int32)

//go:linkname Fn10975 github.com/goccy/spidermonkeywasm2go/p7.Fn10975
func Fn10975(m *base.Module, l0 int32)
