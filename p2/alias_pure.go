//go:build (!amd64 && !arm64) || purego

package p2

import (
	base "github.com/goccy/spidermonkeywasm2go/base"
	_ "unsafe"
)

//go:linkname Fn20 github.com/goccy/spidermonkeywasm2go/p4.Fn20
func Fn20(m *base.Module, l0 int32) int32

//go:linkname Fn21 github.com/goccy/spidermonkeywasm2go/p4.Fn21
func Fn21(m *base.Module, l0 int32) int64

//go:linkname Fn22 github.com/goccy/spidermonkeywasm2go/p4.Fn22
func Fn22(m *base.Module, l0 int32)

//go:linkname Fn23 github.com/goccy/spidermonkeywasm2go/p4.Fn23
func Fn23(m *base.Module, l0 int32) int64

//go:linkname Fn24 github.com/goccy/spidermonkeywasm2go/p4.Fn24
func Fn24(m *base.Module, l0 int32)

//go:linkname Fn27 github.com/goccy/spidermonkeywasm2go/p4.Fn27
func Fn27(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn28 github.com/goccy/spidermonkeywasm2go/p4.Fn28
func Fn28(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn29 github.com/goccy/spidermonkeywasm2go/p4.Fn29
func Fn29(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn30 github.com/goccy/spidermonkeywasm2go/p4.Fn30
func Fn30(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn47 github.com/goccy/spidermonkeywasm2go/p4.Fn47
func Fn47(m *base.Module, l0 int32)

//go:linkname Fn48 github.com/goccy/spidermonkeywasm2go/p4.Fn48
func Fn48(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn49 github.com/goccy/spidermonkeywasm2go/p4.Fn49
func Fn49(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn52 github.com/goccy/spidermonkeywasm2go/p4.Fn52
func Fn52(m *base.Module, l0 int32)

//go:linkname Fn57 github.com/goccy/spidermonkeywasm2go/p4.Fn57
func Fn57(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn67 github.com/goccy/spidermonkeywasm2go/p4.Fn67
func Fn67(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn70 github.com/goccy/spidermonkeywasm2go/p3.Fn70
func Fn70(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn71 github.com/goccy/spidermonkeywasm2go/p4.Fn71
func Fn71(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn72 github.com/goccy/spidermonkeywasm2go/p4.Fn72
func Fn72(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn73 github.com/goccy/spidermonkeywasm2go/p4.Fn73
func Fn73(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn77 github.com/goccy/spidermonkeywasm2go/p4.Fn77
func Fn77(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn79 github.com/goccy/spidermonkeywasm2go/p4.Fn79
func Fn79(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn80 github.com/goccy/spidermonkeywasm2go/p4.Fn80
func Fn80(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn82 github.com/goccy/spidermonkeywasm2go/p4.Fn82
func Fn82(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn83 github.com/goccy/spidermonkeywasm2go/p4.Fn83
func Fn83(m *base.Module)

//go:linkname Fn84 github.com/goccy/spidermonkeywasm2go/p4.Fn84
func Fn84(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn115 github.com/goccy/spidermonkeywasm2go/p3.Fn115
func Fn115(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn118 github.com/goccy/spidermonkeywasm2go/p3.Fn118
func Fn118(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn121 github.com/goccy/spidermonkeywasm2go/p3.Fn121
func Fn121(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn124 github.com/goccy/spidermonkeywasm2go/p3.Fn124
func Fn124(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn127 github.com/goccy/spidermonkeywasm2go/p3.Fn127
func Fn127(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn130 github.com/goccy/spidermonkeywasm2go/p3.Fn130
func Fn130(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn133 github.com/goccy/spidermonkeywasm2go/p3.Fn133
func Fn133(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn136 github.com/goccy/spidermonkeywasm2go/p3.Fn136
func Fn136(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn139 github.com/goccy/spidermonkeywasm2go/p4.Fn139
func Fn139(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn140 github.com/goccy/spidermonkeywasm2go/p4.Fn140
func Fn140(m *base.Module, l0 int32)

//go:linkname Fn142 github.com/goccy/spidermonkeywasm2go/p4.Fn142
func Fn142(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn144 github.com/goccy/spidermonkeywasm2go/p4.Fn144
func Fn144(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn146 github.com/goccy/spidermonkeywasm2go/p4.Fn146
func Fn146(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn147 github.com/goccy/spidermonkeywasm2go/p4.Fn147
func Fn147(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn152 github.com/goccy/spidermonkeywasm2go/p4.Fn152
func Fn152(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn154 github.com/goccy/spidermonkeywasm2go/p4.Fn154
func Fn154(m *base.Module)

//go:linkname Fn155 github.com/goccy/spidermonkeywasm2go/p4.Fn155
func Fn155(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn191 github.com/goccy/spidermonkeywasm2go/p4.Fn191
func Fn191(m *base.Module)

//go:linkname Fn214 github.com/goccy/spidermonkeywasm2go/p4.Fn214
func Fn214(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn215 github.com/goccy/spidermonkeywasm2go/p4.Fn215
func Fn215(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn224 github.com/goccy/spidermonkeywasm2go/p3.Fn224
func Fn224(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn239 github.com/goccy/spidermonkeywasm2go/p3.Fn239
func Fn239(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn241 github.com/goccy/spidermonkeywasm2go/p4.Fn241
func Fn241(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn282 github.com/goccy/spidermonkeywasm2go/p3.Fn282
func Fn282(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn289 github.com/goccy/spidermonkeywasm2go/p3.Fn289
func Fn289(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn291 github.com/goccy/spidermonkeywasm2go/p3.Fn291
func Fn291(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn297 github.com/goccy/spidermonkeywasm2go/p3.Fn297
func Fn297(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn331 github.com/goccy/spidermonkeywasm2go/p4.Fn331
func Fn331(m *base.Module, l0 int32)

//go:linkname Fn433 github.com/goccy/spidermonkeywasm2go/p4.Fn433
func Fn433(m *base.Module, l0 int32)

//go:linkname Fn434 github.com/goccy/spidermonkeywasm2go/p4.Fn434
func Fn434(m *base.Module, l0 int32) int32

//go:linkname Fn442 github.com/goccy/spidermonkeywasm2go/p4.Fn442
func Fn442(m *base.Module, l0 int32) int32

//go:linkname Fn444 github.com/goccy/spidermonkeywasm2go/p4.Fn444
func Fn444(m *base.Module, l0 int32)

//go:linkname Fn449 github.com/goccy/spidermonkeywasm2go/p4.Fn449
func Fn449(m *base.Module, l0 int32) int64

//go:linkname Fn451 github.com/goccy/spidermonkeywasm2go/p4.Fn451
func Fn451(m *base.Module, l0 int32) int32

//go:linkname Fn459 github.com/goccy/spidermonkeywasm2go/p4.Fn459
func Fn459(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn470 github.com/goccy/spidermonkeywasm2go/p4.Fn470
func Fn470(m *base.Module, l0 float64, l1 int32) float64

//go:linkname Fn474 github.com/goccy/spidermonkeywasm2go/p4.Fn474
func Fn474(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn478 github.com/goccy/spidermonkeywasm2go/p4.Fn478
func Fn478(m *base.Module, l0 int32)

//go:linkname Fn479 github.com/goccy/spidermonkeywasm2go/p4.Fn479
func Fn479(m *base.Module, l0 int32) int32

//go:linkname Fn492 github.com/goccy/spidermonkeywasm2go/p4.Fn492
func Fn492(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn494 github.com/goccy/spidermonkeywasm2go/p4.Fn494
func Fn494(m *base.Module, l0 int32) int32

//go:linkname Fn502 github.com/goccy/spidermonkeywasm2go/p4.Fn502
func Fn502(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn512 github.com/goccy/spidermonkeywasm2go/p4.Fn512
func Fn512(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn513 github.com/goccy/spidermonkeywasm2go/p4.Fn513
func Fn513(m *base.Module, l0 int32) int32

//go:linkname Fn514 github.com/goccy/spidermonkeywasm2go/p3.Fn514
func Fn514(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn516 github.com/goccy/spidermonkeywasm2go/p4.Fn516
func Fn516(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn517 github.com/goccy/spidermonkeywasm2go/p4.Fn517
func Fn517(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn518 github.com/goccy/spidermonkeywasm2go/p4.Fn518
func Fn518(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn523 github.com/goccy/spidermonkeywasm2go/p4.Fn523
func Fn523(m *base.Module, l0 int32) int32

//go:linkname Fn530 github.com/goccy/spidermonkeywasm2go/p4.Fn530
func Fn530(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn538 github.com/goccy/spidermonkeywasm2go/p4.Fn538
func Fn538(m *base.Module)

//go:linkname Fn539 github.com/goccy/spidermonkeywasm2go/p4.Fn539
func Fn539(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn541 github.com/goccy/spidermonkeywasm2go/p4.Fn541
func Fn541(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn542 github.com/goccy/spidermonkeywasm2go/p4.Fn542
func Fn542(m *base.Module, l0 int32) int32

//go:linkname Fn544 github.com/goccy/spidermonkeywasm2go/p3.Fn544
func Fn544(m *base.Module, l0 int32)

//go:linkname Fn545 github.com/goccy/spidermonkeywasm2go/p4.Fn545
func Fn545(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn546 github.com/goccy/spidermonkeywasm2go/p3.Fn546
func Fn546(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn548 github.com/goccy/spidermonkeywasm2go/p4.Fn548
func Fn548(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn549 github.com/goccy/spidermonkeywasm2go/p4.Fn549
func Fn549(m *base.Module, l0 int32) int32

//go:linkname Fn550 github.com/goccy/spidermonkeywasm2go/p4.Fn550
func Fn550(m *base.Module, l0 int32)

//go:linkname Fn551 github.com/goccy/spidermonkeywasm2go/p4.Fn551
func Fn551(m *base.Module, l0 int32)

//go:linkname Fn552 github.com/goccy/spidermonkeywasm2go/p4.Fn552
func Fn552(m *base.Module, l0 int32)

//go:linkname Fn556 github.com/goccy/spidermonkeywasm2go/p4.Fn556
func Fn556(m *base.Module, l0 int32) int32

//go:linkname Fn557 github.com/goccy/spidermonkeywasm2go/p4.Fn557
func Fn557(m *base.Module, l0 int32) int32

//go:linkname Fn559 github.com/goccy/spidermonkeywasm2go/p0.Fn559
func Fn559(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn560 github.com/goccy/spidermonkeywasm2go/p4.Fn560
func Fn560(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn561 github.com/goccy/spidermonkeywasm2go/p0.Fn561
func Fn561(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn565 github.com/goccy/spidermonkeywasm2go/p4.Fn565
func Fn565(m *base.Module, l0 int32)

//go:linkname Fn566 github.com/goccy/spidermonkeywasm2go/p0.Fn566
func Fn566(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn567 github.com/goccy/spidermonkeywasm2go/p0.Fn567
func Fn567(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn569 github.com/goccy/spidermonkeywasm2go/p3.Fn569
func Fn569(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn580 github.com/goccy/spidermonkeywasm2go/p4.Fn580
func Fn580(m *base.Module, l0 int32) int32

//go:linkname Fn581 github.com/goccy/spidermonkeywasm2go/p0.Fn581
func Fn581(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn582 github.com/goccy/spidermonkeywasm2go/p3.Fn582
func Fn582(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn586 github.com/goccy/spidermonkeywasm2go/p4.Fn586
func Fn586(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn588 github.com/goccy/spidermonkeywasm2go/p3.Fn588
func Fn588(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn593 github.com/goccy/spidermonkeywasm2go/p4.Fn593
func Fn593(m *base.Module, l0 int32)

//go:linkname Fn598 github.com/goccy/spidermonkeywasm2go/p4.Fn598
func Fn598(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn607 github.com/goccy/spidermonkeywasm2go/p0.Fn607
func Fn607(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn608 github.com/goccy/spidermonkeywasm2go/p0.Fn608
func Fn608(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn609 github.com/goccy/spidermonkeywasm2go/p0.Fn609
func Fn609(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn610 github.com/goccy/spidermonkeywasm2go/p4.Fn610
func Fn610(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn612 github.com/goccy/spidermonkeywasm2go/p4.Fn612
func Fn612(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn615 github.com/goccy/spidermonkeywasm2go/p4.Fn615
func Fn615(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn630 github.com/goccy/spidermonkeywasm2go/p4.Fn630
func Fn630(m *base.Module, l0 int32) int32

//go:linkname Fn631 github.com/goccy/spidermonkeywasm2go/p4.Fn631
func Fn631(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn632 github.com/goccy/spidermonkeywasm2go/p4.Fn632
func Fn632(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn633 github.com/goccy/spidermonkeywasm2go/p4.Fn633
func Fn633(m *base.Module, l0 int32) int32

//go:linkname Fn634 github.com/goccy/spidermonkeywasm2go/p4.Fn634
func Fn634(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn635 github.com/goccy/spidermonkeywasm2go/p3.Fn635
func Fn635(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn636 github.com/goccy/spidermonkeywasm2go/p4.Fn636
func Fn636(m *base.Module, l0 int32) int32

//go:linkname Fn637 github.com/goccy/spidermonkeywasm2go/p4.Fn637
func Fn637(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn638 github.com/goccy/spidermonkeywasm2go/p4.Fn638
func Fn638(m *base.Module, l0 int32) int32

//go:linkname Fn639 github.com/goccy/spidermonkeywasm2go/p3.Fn639
func Fn639(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn646 github.com/goccy/spidermonkeywasm2go/p4.Fn646
func Fn646(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn647 github.com/goccy/spidermonkeywasm2go/p4.Fn647
func Fn647(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn658 github.com/goccy/spidermonkeywasm2go/p4.Fn658
func Fn658(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn659 github.com/goccy/spidermonkeywasm2go/p0.Fn659
func Fn659(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn680 github.com/goccy/spidermonkeywasm2go/p0.Fn680
func Fn680(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn695 github.com/goccy/spidermonkeywasm2go/p4.Fn695
func Fn695(m *base.Module, l0 int32) int64

//go:linkname Fn701 github.com/goccy/spidermonkeywasm2go/p0.Fn701
func Fn701(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn711 github.com/goccy/spidermonkeywasm2go/p3.Fn711
func Fn711(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn714 github.com/goccy/spidermonkeywasm2go/p0.Fn714
func Fn714(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn721 github.com/goccy/spidermonkeywasm2go/p0.Fn721
func Fn721(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn734 github.com/goccy/spidermonkeywasm2go/p4.Fn734
func Fn734(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn740 github.com/goccy/spidermonkeywasm2go/p4.Fn740
func Fn740(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn741 github.com/goccy/spidermonkeywasm2go/p4.Fn741
func Fn741(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn744 github.com/goccy/spidermonkeywasm2go/p3.Fn744
func Fn744(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn757 github.com/goccy/spidermonkeywasm2go/p3.Fn757
func Fn757(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn758 github.com/goccy/spidermonkeywasm2go/p4.Fn758
func Fn758(m *base.Module, l0 int32) int32

//go:linkname Fn761 github.com/goccy/spidermonkeywasm2go/p3.Fn761
func Fn761(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn763 github.com/goccy/spidermonkeywasm2go/p4.Fn763
func Fn763(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn764 github.com/goccy/spidermonkeywasm2go/p3.Fn764
func Fn764(m *base.Module, l0 int32)

//go:linkname Fn765 github.com/goccy/spidermonkeywasm2go/p3.Fn765
func Fn765(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn766 github.com/goccy/spidermonkeywasm2go/p0.Fn766
func Fn766(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn768 github.com/goccy/spidermonkeywasm2go/p0.Fn768
func Fn768(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn793 github.com/goccy/spidermonkeywasm2go/p0.Fn793
func Fn793(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn794 github.com/goccy/spidermonkeywasm2go/p4.Fn794
func Fn794(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn795 github.com/goccy/spidermonkeywasm2go/p3.Fn795
func Fn795(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn796 github.com/goccy/spidermonkeywasm2go/p0.Fn796
func Fn796(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn801 github.com/goccy/spidermonkeywasm2go/p3.Fn801
func Fn801(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn812 github.com/goccy/spidermonkeywasm2go/p4.Fn812
func Fn812(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn815 github.com/goccy/spidermonkeywasm2go/p4.Fn815
func Fn815(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn817 github.com/goccy/spidermonkeywasm2go/p4.Fn817
func Fn817(m *base.Module, l0 int32)

//go:linkname Fn819 github.com/goccy/spidermonkeywasm2go/p4.Fn819
func Fn819(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn820 github.com/goccy/spidermonkeywasm2go/p4.Fn820
func Fn820(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn825 github.com/goccy/spidermonkeywasm2go/p3.Fn825
func Fn825(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn826 github.com/goccy/spidermonkeywasm2go/p4.Fn826
func Fn826(m *base.Module, l0 int32) int32

//go:linkname Fn832 github.com/goccy/spidermonkeywasm2go/p3.Fn832
func Fn832(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn844 github.com/goccy/spidermonkeywasm2go/p0.Fn844
func Fn844(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn846 github.com/goccy/spidermonkeywasm2go/p0.Fn846
func Fn846(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn848 github.com/goccy/spidermonkeywasm2go/p4.Fn848
func Fn848(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn850 github.com/goccy/spidermonkeywasm2go/p0.Fn850
func Fn850(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn852 github.com/goccy/spidermonkeywasm2go/p0.Fn852
func Fn852(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn853 github.com/goccy/spidermonkeywasm2go/p0.Fn853
func Fn853(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn855 github.com/goccy/spidermonkeywasm2go/p0.Fn855
func Fn855(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn856 github.com/goccy/spidermonkeywasm2go/p0.Fn856
func Fn856(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn857 github.com/goccy/spidermonkeywasm2go/p4.Fn857
func Fn857(m *base.Module, l0 int32) int32

//go:linkname Fn865 github.com/goccy/spidermonkeywasm2go/p0.Fn865
func Fn865(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn871 github.com/goccy/spidermonkeywasm2go/p0.Fn871
func Fn871(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn873 github.com/goccy/spidermonkeywasm2go/p0.Fn873
func Fn873(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn874 github.com/goccy/spidermonkeywasm2go/p0.Fn874
func Fn874(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn881 github.com/goccy/spidermonkeywasm2go/p0.Fn881
func Fn881(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn882 github.com/goccy/spidermonkeywasm2go/p0.Fn882
func Fn882(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn884 github.com/goccy/spidermonkeywasm2go/p0.Fn884
func Fn884(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn885 github.com/goccy/spidermonkeywasm2go/p0.Fn885
func Fn885(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn886 github.com/goccy/spidermonkeywasm2go/p0.Fn886
func Fn886(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn887 github.com/goccy/spidermonkeywasm2go/p0.Fn887
func Fn887(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn888 github.com/goccy/spidermonkeywasm2go/p0.Fn888
func Fn888(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn889 github.com/goccy/spidermonkeywasm2go/p0.Fn889
func Fn889(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn890 github.com/goccy/spidermonkeywasm2go/p0.Fn890
func Fn890(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn893 github.com/goccy/spidermonkeywasm2go/p3.Fn893
func Fn893(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn895 github.com/goccy/spidermonkeywasm2go/p0.Fn895
func Fn895(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn896 github.com/goccy/spidermonkeywasm2go/p0.Fn896
func Fn896(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn898 github.com/goccy/spidermonkeywasm2go/p0.Fn898
func Fn898(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn900 github.com/goccy/spidermonkeywasm2go/p4.Fn900
func Fn900(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn901 github.com/goccy/spidermonkeywasm2go/p0.Fn901
func Fn901(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn935 github.com/goccy/spidermonkeywasm2go/p0.Fn935
func Fn935(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn965 github.com/goccy/spidermonkeywasm2go/p0.Fn965
func Fn965(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn999 github.com/goccy/spidermonkeywasm2go/p4.Fn999
func Fn999(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1001 github.com/goccy/spidermonkeywasm2go/p4.Fn1001
func Fn1001(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1017 github.com/goccy/spidermonkeywasm2go/p4.Fn1017
func Fn1017(m *base.Module, l0 int32) int32

//go:linkname Fn1018 github.com/goccy/spidermonkeywasm2go/p4.Fn1018
func Fn1018(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1019 github.com/goccy/spidermonkeywasm2go/p0.Fn1019
func Fn1019(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1020 github.com/goccy/spidermonkeywasm2go/p0.Fn1020
func Fn1020(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn1021 github.com/goccy/spidermonkeywasm2go/p3.Fn1021
func Fn1021(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1022 github.com/goccy/spidermonkeywasm2go/p4.Fn1022
func Fn1022(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1026 github.com/goccy/spidermonkeywasm2go/p0.Fn1026
func Fn1026(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1032 github.com/goccy/spidermonkeywasm2go/p4.Fn1032
func Fn1032(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn1037 github.com/goccy/spidermonkeywasm2go/p0.Fn1037
func Fn1037(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1038 github.com/goccy/spidermonkeywasm2go/p4.Fn1038
func Fn1038(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1040 github.com/goccy/spidermonkeywasm2go/p3.Fn1040
func Fn1040(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1041 github.com/goccy/spidermonkeywasm2go/p3.Fn1041
func Fn1041(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn1046 github.com/goccy/spidermonkeywasm2go/p4.Fn1046
func Fn1046(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1049 github.com/goccy/spidermonkeywasm2go/p4.Fn1049
func Fn1049(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1050 github.com/goccy/spidermonkeywasm2go/p4.Fn1050
func Fn1050(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1052 github.com/goccy/spidermonkeywasm2go/p0.Fn1052
func Fn1052(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn1053 github.com/goccy/spidermonkeywasm2go/p4.Fn1053
func Fn1053(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1055 github.com/goccy/spidermonkeywasm2go/p0.Fn1055
func Fn1055(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1056 github.com/goccy/spidermonkeywasm2go/p0.Fn1056
func Fn1056(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1057 github.com/goccy/spidermonkeywasm2go/p0.Fn1057
func Fn1057(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1064 github.com/goccy/spidermonkeywasm2go/p4.Fn1064
func Fn1064(m *base.Module, l0 int32) int32

//go:linkname Fn1065 github.com/goccy/spidermonkeywasm2go/p0.Fn1065
func Fn1065(m *base.Module, l0 int32) int32

//go:linkname Fn1066 github.com/goccy/spidermonkeywasm2go/p0.Fn1066
func Fn1066(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1067 github.com/goccy/spidermonkeywasm2go/p4.Fn1067
func Fn1067(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1084 github.com/goccy/spidermonkeywasm2go/p4.Fn1084
func Fn1084(m *base.Module, l0 int32)

//go:linkname Fn1088 github.com/goccy/spidermonkeywasm2go/p4.Fn1088
func Fn1088(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1089 github.com/goccy/spidermonkeywasm2go/p3.Fn1089
func Fn1089(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1090 github.com/goccy/spidermonkeywasm2go/p0.Fn1090
func Fn1090(m *base.Module, l0 int32) int32

//go:linkname Fn1107 github.com/goccy/spidermonkeywasm2go/p4.Fn1107
func Fn1107(m *base.Module, l0 int32) int32

//go:linkname Fn1108 github.com/goccy/spidermonkeywasm2go/p3.Fn1108
func Fn1108(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1176 github.com/goccy/spidermonkeywasm2go/p4.Fn1176
func Fn1176(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1196 github.com/goccy/spidermonkeywasm2go/p4.Fn1196
func Fn1196(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1205 github.com/goccy/spidermonkeywasm2go/p4.Fn1205
func Fn1205(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64, l5 int32, l6 int32) int32

//go:linkname Fn1206 github.com/goccy/spidermonkeywasm2go/p3.Fn1206
func Fn1206(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1208 github.com/goccy/spidermonkeywasm2go/p4.Fn1208
func Fn1208(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1212 github.com/goccy/spidermonkeywasm2go/p4.Fn1212
func Fn1212(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1213 github.com/goccy/spidermonkeywasm2go/p4.Fn1213
func Fn1213(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1246 github.com/goccy/spidermonkeywasm2go/p4.Fn1246
func Fn1246(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1247 github.com/goccy/spidermonkeywasm2go/p4.Fn1247
func Fn1247(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1253 github.com/goccy/spidermonkeywasm2go/p4.Fn1253
func Fn1253(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1256 github.com/goccy/spidermonkeywasm2go/p4.Fn1256
func Fn1256(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1257 github.com/goccy/spidermonkeywasm2go/p0.Fn1257
func Fn1257(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1258 github.com/goccy/spidermonkeywasm2go/p4.Fn1258
func Fn1258(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1259 github.com/goccy/spidermonkeywasm2go/p4.Fn1259
func Fn1259(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1260 github.com/goccy/spidermonkeywasm2go/p4.Fn1260
func Fn1260(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1285 github.com/goccy/spidermonkeywasm2go/p4.Fn1285
func Fn1285(m *base.Module, l0 int32) int32

//go:linkname Fn1286 github.com/goccy/spidermonkeywasm2go/p4.Fn1286
func Fn1286(m *base.Module, l0 int32) int32

//go:linkname Fn1287 github.com/goccy/spidermonkeywasm2go/p4.Fn1287
func Fn1287(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1314 github.com/goccy/spidermonkeywasm2go/p0.Fn1314
func Fn1314(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1317 github.com/goccy/spidermonkeywasm2go/p4.Fn1317
func Fn1317(m *base.Module, l0 int32) int32

//go:linkname Fn1318 github.com/goccy/spidermonkeywasm2go/p4.Fn1318
func Fn1318(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1324 github.com/goccy/spidermonkeywasm2go/p3.Fn1324
func Fn1324(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1325 github.com/goccy/spidermonkeywasm2go/p3.Fn1325
func Fn1325(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1326 github.com/goccy/spidermonkeywasm2go/p4.Fn1326
func Fn1326(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1327 github.com/goccy/spidermonkeywasm2go/p0.Fn1327
func Fn1327(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1328 github.com/goccy/spidermonkeywasm2go/p4.Fn1328
func Fn1328(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1335 github.com/goccy/spidermonkeywasm2go/p4.Fn1335
func Fn1335(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1338 github.com/goccy/spidermonkeywasm2go/p4.Fn1338
func Fn1338(m *base.Module, l0 int32)

//go:linkname Fn1347 github.com/goccy/spidermonkeywasm2go/p4.Fn1347
func Fn1347(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1353 github.com/goccy/spidermonkeywasm2go/p3.Fn1353
func Fn1353(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1355 github.com/goccy/spidermonkeywasm2go/p4.Fn1355
func Fn1355(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1393 github.com/goccy/spidermonkeywasm2go/p4.Fn1393
func Fn1393(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1394 github.com/goccy/spidermonkeywasm2go/p4.Fn1394
func Fn1394(m *base.Module, l0 int32) int32

//go:linkname Fn1396 github.com/goccy/spidermonkeywasm2go/p4.Fn1396
func Fn1396(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1398 github.com/goccy/spidermonkeywasm2go/p4.Fn1398
func Fn1398(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1426 github.com/goccy/spidermonkeywasm2go/p4.Fn1426
func Fn1426(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1428 github.com/goccy/spidermonkeywasm2go/p4.Fn1428
func Fn1428(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1444 github.com/goccy/spidermonkeywasm2go/p0.Fn1444
func Fn1444(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1448 github.com/goccy/spidermonkeywasm2go/p4.Fn1448
func Fn1448(m *base.Module, l0 int32) int32

//go:linkname Fn1469 github.com/goccy/spidermonkeywasm2go/p4.Fn1469
func Fn1469(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1471 github.com/goccy/spidermonkeywasm2go/p0.Fn1471
func Fn1471(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1473 github.com/goccy/spidermonkeywasm2go/p3.Fn1473
func Fn1473(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1476 github.com/goccy/spidermonkeywasm2go/p0.Fn1476
func Fn1476(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1478 github.com/goccy/spidermonkeywasm2go/p4.Fn1478
func Fn1478(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1479 github.com/goccy/spidermonkeywasm2go/p4.Fn1479
func Fn1479(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1483 github.com/goccy/spidermonkeywasm2go/p3.Fn1483
func Fn1483(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1484 github.com/goccy/spidermonkeywasm2go/p4.Fn1484
func Fn1484(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1489 github.com/goccy/spidermonkeywasm2go/p3.Fn1489
func Fn1489(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1490 github.com/goccy/spidermonkeywasm2go/p1.Fn1490
func Fn1490(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1586 github.com/goccy/spidermonkeywasm2go/p4.Fn1586
func Fn1586(m *base.Module, l0 int32)

//go:linkname Fn1607 github.com/goccy/spidermonkeywasm2go/p4.Fn1607
func Fn1607(m *base.Module, l0 int32) int32

//go:linkname Fn1608 github.com/goccy/spidermonkeywasm2go/p4.Fn1608
func Fn1608(m *base.Module, l0 int32) int32

//go:linkname Fn1609 github.com/goccy/spidermonkeywasm2go/p0.Fn1609
func Fn1609(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1611 github.com/goccy/spidermonkeywasm2go/p4.Fn1611
func Fn1611(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1612 github.com/goccy/spidermonkeywasm2go/p4.Fn1612
func Fn1612(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1613 github.com/goccy/spidermonkeywasm2go/p3.Fn1613
func Fn1613(m *base.Module, l0 int32)

//go:linkname Fn1614 github.com/goccy/spidermonkeywasm2go/p4.Fn1614
func Fn1614(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1615 github.com/goccy/spidermonkeywasm2go/p3.Fn1615
func Fn1615(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1617 github.com/goccy/spidermonkeywasm2go/p4.Fn1617
func Fn1617(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1623 github.com/goccy/spidermonkeywasm2go/p4.Fn1623
func Fn1623(m *base.Module, l0 int32)

//go:linkname Fn1626 github.com/goccy/spidermonkeywasm2go/p4.Fn1626
func Fn1626(m *base.Module, l0 int32) int32

//go:linkname Fn1634 github.com/goccy/spidermonkeywasm2go/p0.Fn1634
func Fn1634(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1636 github.com/goccy/spidermonkeywasm2go/p4.Fn1636
func Fn1636(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1638 github.com/goccy/spidermonkeywasm2go/p3.Fn1638
func Fn1638(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1645 github.com/goccy/spidermonkeywasm2go/p0.Fn1645
func Fn1645(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32

//go:linkname Fn1651 github.com/goccy/spidermonkeywasm2go/p3.Fn1651
func Fn1651(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1658 github.com/goccy/spidermonkeywasm2go/p0.Fn1658
func Fn1658(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1683 github.com/goccy/spidermonkeywasm2go/p3.Fn1683
func Fn1683(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1686 github.com/goccy/spidermonkeywasm2go/p3.Fn1686
func Fn1686(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1692 github.com/goccy/spidermonkeywasm2go/p3.Fn1692
func Fn1692(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1693 github.com/goccy/spidermonkeywasm2go/p3.Fn1693
func Fn1693(m *base.Module, l0 int32) int32

//go:linkname Fn1694 github.com/goccy/spidermonkeywasm2go/p4.Fn1694
func Fn1694(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1695 github.com/goccy/spidermonkeywasm2go/p3.Fn1695
func Fn1695(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1697 github.com/goccy/spidermonkeywasm2go/p3.Fn1697
func Fn1697(m *base.Module, l0 int32)

//go:linkname Fn1698 github.com/goccy/spidermonkeywasm2go/p4.Fn1698
func Fn1698(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1699 github.com/goccy/spidermonkeywasm2go/p4.Fn1699
func Fn1699(m *base.Module, l0 int32)

//go:linkname Fn1701 github.com/goccy/spidermonkeywasm2go/p3.Fn1701
func Fn1701(m *base.Module, l0 int32)

//go:linkname Fn1704 github.com/goccy/spidermonkeywasm2go/p4.Fn1704
func Fn1704(m *base.Module, l0 int32) int32

//go:linkname Fn1705 github.com/goccy/spidermonkeywasm2go/p4.Fn1705
func Fn1705(m *base.Module, l0 int32) int32

//go:linkname Fn1707 github.com/goccy/spidermonkeywasm2go/p4.Fn1707
func Fn1707(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1709 github.com/goccy/spidermonkeywasm2go/p4.Fn1709
func Fn1709(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1710 github.com/goccy/spidermonkeywasm2go/p0.Fn1710
func Fn1710(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1711 github.com/goccy/spidermonkeywasm2go/p4.Fn1711
func Fn1711(m *base.Module, l0 int32) int32

//go:linkname Fn1712 github.com/goccy/spidermonkeywasm2go/p4.Fn1712
func Fn1712(m *base.Module, l0 int32)

//go:linkname Fn1713 github.com/goccy/spidermonkeywasm2go/p4.Fn1713
func Fn1713(m *base.Module, l0 int32)

//go:linkname Fn1715 github.com/goccy/spidermonkeywasm2go/p3.Fn1715
func Fn1715(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1716 github.com/goccy/spidermonkeywasm2go/p4.Fn1716
func Fn1716(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1720 github.com/goccy/spidermonkeywasm2go/p0.Fn1720
func Fn1720(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1724 github.com/goccy/spidermonkeywasm2go/p3.Fn1724
func Fn1724(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1742 github.com/goccy/spidermonkeywasm2go/p3.Fn1742
func Fn1742(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1747 github.com/goccy/spidermonkeywasm2go/p3.Fn1747
func Fn1747(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1748 github.com/goccy/spidermonkeywasm2go/p3.Fn1748
func Fn1748(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1756 github.com/goccy/spidermonkeywasm2go/p4.Fn1756
func Fn1756(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1770 github.com/goccy/spidermonkeywasm2go/p4.Fn1770
func Fn1770(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1775 github.com/goccy/spidermonkeywasm2go/p3.Fn1775
func Fn1775(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1776 github.com/goccy/spidermonkeywasm2go/p4.Fn1776
func Fn1776(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1777 github.com/goccy/spidermonkeywasm2go/p4.Fn1777
func Fn1777(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1778 github.com/goccy/spidermonkeywasm2go/p4.Fn1778
func Fn1778(m *base.Module, l0 int32) int32

//go:linkname Fn1781 github.com/goccy/spidermonkeywasm2go/p0.Fn1781
func Fn1781(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1785 github.com/goccy/spidermonkeywasm2go/p0.Fn1785
func Fn1785(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1786 github.com/goccy/spidermonkeywasm2go/p0.Fn1786
func Fn1786(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1787 github.com/goccy/spidermonkeywasm2go/p0.Fn1787
func Fn1787(m *base.Module, l0 int32) int32

//go:linkname Fn1790 github.com/goccy/spidermonkeywasm2go/p0.Fn1790
func Fn1790(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1800 github.com/goccy/spidermonkeywasm2go/p0.Fn1800
func Fn1800(m *base.Module, l0 int32) int32

//go:linkname Fn1802 github.com/goccy/spidermonkeywasm2go/p0.Fn1802
func Fn1802(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1805 github.com/goccy/spidermonkeywasm2go/p0.Fn1805
func Fn1805(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1806 github.com/goccy/spidermonkeywasm2go/p0.Fn1806
func Fn1806(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1813 github.com/goccy/spidermonkeywasm2go/p3.Fn1813
func Fn1813(m *base.Module, l0 int32)

//go:linkname Fn1814 github.com/goccy/spidermonkeywasm2go/p3.Fn1814
func Fn1814(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1819 github.com/goccy/spidermonkeywasm2go/p3.Fn1819
func Fn1819(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1820 github.com/goccy/spidermonkeywasm2go/p0.Fn1820
func Fn1820(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1821 github.com/goccy/spidermonkeywasm2go/p0.Fn1821
func Fn1821(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1824 github.com/goccy/spidermonkeywasm2go/p0.Fn1824
func Fn1824(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1825 github.com/goccy/spidermonkeywasm2go/p0.Fn1825
func Fn1825(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn1827 github.com/goccy/spidermonkeywasm2go/p0.Fn1827
func Fn1827(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1848 github.com/goccy/spidermonkeywasm2go/p0.Fn1848
func Fn1848(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1850 github.com/goccy/spidermonkeywasm2go/p4.Fn1850
func Fn1850(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1861 github.com/goccy/spidermonkeywasm2go/p4.Fn1861
func Fn1861(m *base.Module, l0 int32)

//go:linkname Fn1867 github.com/goccy/spidermonkeywasm2go/p4.Fn1867
func Fn1867(m *base.Module, l0 int32) int32

//go:linkname Fn1868 github.com/goccy/spidermonkeywasm2go/p3.Fn1868
func Fn1868(m *base.Module, l0 int32) int32

//go:linkname Fn1869 github.com/goccy/spidermonkeywasm2go/p0.Fn1869
func Fn1869(m *base.Module, l0 int32) int32

//go:linkname Fn1872 github.com/goccy/spidermonkeywasm2go/p0.Fn1872
func Fn1872(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1874 github.com/goccy/spidermonkeywasm2go/p0.Fn1874
func Fn1874(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1877 github.com/goccy/spidermonkeywasm2go/p4.Fn1877
func Fn1877(m *base.Module, l0 int32)

//go:linkname Fn1888 github.com/goccy/spidermonkeywasm2go/p4.Fn1888
func Fn1888(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1889 github.com/goccy/spidermonkeywasm2go/p3.Fn1889
func Fn1889(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1896 github.com/goccy/spidermonkeywasm2go/p0.Fn1896
func Fn1896(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1903 github.com/goccy/spidermonkeywasm2go/p0.Fn1903
func Fn1903(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1910 github.com/goccy/spidermonkeywasm2go/p0.Fn1910
func Fn1910(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1911 github.com/goccy/spidermonkeywasm2go/p0.Fn1911
func Fn1911(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1914 github.com/goccy/spidermonkeywasm2go/p3.Fn1914
func Fn1914(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1917 github.com/goccy/spidermonkeywasm2go/p4.Fn1917
func Fn1917(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1918 github.com/goccy/spidermonkeywasm2go/p0.Fn1918
func Fn1918(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1927 github.com/goccy/spidermonkeywasm2go/p0.Fn1927
func Fn1927(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1939 github.com/goccy/spidermonkeywasm2go/p0.Fn1939
func Fn1939(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1940 github.com/goccy/spidermonkeywasm2go/p0.Fn1940
func Fn1940(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1943 github.com/goccy/spidermonkeywasm2go/p0.Fn1943
func Fn1943(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1944 github.com/goccy/spidermonkeywasm2go/p0.Fn1944
func Fn1944(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1946 github.com/goccy/spidermonkeywasm2go/p0.Fn1946
func Fn1946(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1947 github.com/goccy/spidermonkeywasm2go/p4.Fn1947
func Fn1947(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1949 github.com/goccy/spidermonkeywasm2go/p3.Fn1949
func Fn1949(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1952 github.com/goccy/spidermonkeywasm2go/p4.Fn1952
func Fn1952(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1953 github.com/goccy/spidermonkeywasm2go/p0.Fn1953
func Fn1953(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1959 github.com/goccy/spidermonkeywasm2go/p3.Fn1959
func Fn1959(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1965 github.com/goccy/spidermonkeywasm2go/p4.Fn1965
func Fn1965(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1968 github.com/goccy/spidermonkeywasm2go/p4.Fn1968
func Fn1968(m *base.Module, l0 int32) int32

//go:linkname Fn1969 github.com/goccy/spidermonkeywasm2go/p4.Fn1969
func Fn1969(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1979 github.com/goccy/spidermonkeywasm2go/p4.Fn1979
func Fn1979(m *base.Module, l0 int32) int32

//go:linkname Fn1981 github.com/goccy/spidermonkeywasm2go/p4.Fn1981
func Fn1981(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2038 github.com/goccy/spidermonkeywasm2go/p4.Fn2038
func Fn2038(m *base.Module, l0 int32)

//go:linkname Fn2040 github.com/goccy/spidermonkeywasm2go/p0.Fn2040
func Fn2040(m *base.Module, l0 int32)

//go:linkname Fn2041 github.com/goccy/spidermonkeywasm2go/p0.Fn2041
func Fn2041(m *base.Module, l0 int32)

//go:linkname Fn2045 github.com/goccy/spidermonkeywasm2go/p0.Fn2045
func Fn2045(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2047 github.com/goccy/spidermonkeywasm2go/p4.Fn2047
func Fn2047(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2048 github.com/goccy/spidermonkeywasm2go/p0.Fn2048
func Fn2048(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2049 github.com/goccy/spidermonkeywasm2go/p4.Fn2049
func Fn2049(m *base.Module, l0 int32)

//go:linkname Fn2054 github.com/goccy/spidermonkeywasm2go/p0.Fn2054
func Fn2054(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2057 github.com/goccy/spidermonkeywasm2go/p3.Fn2057
func Fn2057(m *base.Module, l0 int32)

//go:linkname Fn2062 github.com/goccy/spidermonkeywasm2go/p3.Fn2062
func Fn2062(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2063 github.com/goccy/spidermonkeywasm2go/p4.Fn2063
func Fn2063(m *base.Module, l0 int32)

//go:linkname Fn2066 github.com/goccy/spidermonkeywasm2go/p0.Fn2066
func Fn2066(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2073 github.com/goccy/spidermonkeywasm2go/p0.Fn2073
func Fn2073(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2078 github.com/goccy/spidermonkeywasm2go/p0.Fn2078
func Fn2078(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2079 github.com/goccy/spidermonkeywasm2go/p0.Fn2079
func Fn2079(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2080 github.com/goccy/spidermonkeywasm2go/p0.Fn2080
func Fn2080(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2085 github.com/goccy/spidermonkeywasm2go/p4.Fn2085
func Fn2085(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn2086 github.com/goccy/spidermonkeywasm2go/p0.Fn2086
func Fn2086(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2087 github.com/goccy/spidermonkeywasm2go/p0.Fn2087
func Fn2087(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2088 github.com/goccy/spidermonkeywasm2go/p0.Fn2088
func Fn2088(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2095 github.com/goccy/spidermonkeywasm2go/p0.Fn2095
func Fn2095(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn2107 github.com/goccy/spidermonkeywasm2go/p4.Fn2107
func Fn2107(m *base.Module, l0 int32)

//go:linkname Fn2108 github.com/goccy/spidermonkeywasm2go/p0.Fn2108
func Fn2108(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2110 github.com/goccy/spidermonkeywasm2go/p0.Fn2110
func Fn2110(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2117 github.com/goccy/spidermonkeywasm2go/p0.Fn2117
func Fn2117(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2127 github.com/goccy/spidermonkeywasm2go/p4.Fn2127
func Fn2127(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2129 github.com/goccy/spidermonkeywasm2go/p4.Fn2129
func Fn2129(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2130 github.com/goccy/spidermonkeywasm2go/p4.Fn2130
func Fn2130(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2131 github.com/goccy/spidermonkeywasm2go/p0.Fn2131
func Fn2131(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2133 github.com/goccy/spidermonkeywasm2go/p4.Fn2133
func Fn2133(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2134 github.com/goccy/spidermonkeywasm2go/p0.Fn2134
func Fn2134(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2135 github.com/goccy/spidermonkeywasm2go/p0.Fn2135
func Fn2135(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2136 github.com/goccy/spidermonkeywasm2go/p0.Fn2136
func Fn2136(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2137 github.com/goccy/spidermonkeywasm2go/p0.Fn2137
func Fn2137(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2142 github.com/goccy/spidermonkeywasm2go/p0.Fn2142
func Fn2142(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2144 github.com/goccy/spidermonkeywasm2go/p0.Fn2144
func Fn2144(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2145 github.com/goccy/spidermonkeywasm2go/p3.Fn2145
func Fn2145(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2146 github.com/goccy/spidermonkeywasm2go/p3.Fn2146
func Fn2146(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2147 github.com/goccy/spidermonkeywasm2go/p4.Fn2147
func Fn2147(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2148 github.com/goccy/spidermonkeywasm2go/p4.Fn2148
func Fn2148(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2155 github.com/goccy/spidermonkeywasm2go/p0.Fn2155
func Fn2155(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2156 github.com/goccy/spidermonkeywasm2go/p0.Fn2156
func Fn2156(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2158 github.com/goccy/spidermonkeywasm2go/p0.Fn2158
func Fn2158(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2162 github.com/goccy/spidermonkeywasm2go/p3.Fn2162
func Fn2162(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2165 github.com/goccy/spidermonkeywasm2go/p3.Fn2165
func Fn2165(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2166 github.com/goccy/spidermonkeywasm2go/p3.Fn2166
func Fn2166(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2168 github.com/goccy/spidermonkeywasm2go/p0.Fn2168
func Fn2168(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2170 github.com/goccy/spidermonkeywasm2go/p3.Fn2170
func Fn2170(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2175 github.com/goccy/spidermonkeywasm2go/p4.Fn2175
func Fn2175(m *base.Module, l0 int32)

//go:linkname Fn2187 github.com/goccy/spidermonkeywasm2go/p4.Fn2187
func Fn2187(m *base.Module, l0 int32)

//go:linkname Fn2188 github.com/goccy/spidermonkeywasm2go/p4.Fn2188
func Fn2188(m *base.Module, l0 int32) int32

//go:linkname Fn2194 github.com/goccy/spidermonkeywasm2go/p4.Fn2194
func Fn2194(m *base.Module, l0 int32)

//go:linkname Fn2195 github.com/goccy/spidermonkeywasm2go/p4.Fn2195
func Fn2195(m *base.Module, l0 int32) int32

//go:linkname Fn2199 github.com/goccy/spidermonkeywasm2go/p3.Fn2199
func Fn2199(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2200 github.com/goccy/spidermonkeywasm2go/p4.Fn2200
func Fn2200(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2201 github.com/goccy/spidermonkeywasm2go/p3.Fn2201
func Fn2201(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2209 github.com/goccy/spidermonkeywasm2go/p3.Fn2209
func Fn2209(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2242 github.com/goccy/spidermonkeywasm2go/p4.Fn2242
func Fn2242(m *base.Module, l0 int32) int32

//go:linkname Fn2262 github.com/goccy/spidermonkeywasm2go/p0.Fn2262
func Fn2262(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2275 github.com/goccy/spidermonkeywasm2go/p4.Fn2275
func Fn2275(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2276 github.com/goccy/spidermonkeywasm2go/p4.Fn2276
func Fn2276(m *base.Module, l0 int32)

//go:linkname Fn2281 github.com/goccy/spidermonkeywasm2go/p4.Fn2281
func Fn2281(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2283 github.com/goccy/spidermonkeywasm2go/p0.Fn2283
func Fn2283(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2284 github.com/goccy/spidermonkeywasm2go/p0.Fn2284
func Fn2284(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2285 github.com/goccy/spidermonkeywasm2go/p3.Fn2285
func Fn2285(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2286 github.com/goccy/spidermonkeywasm2go/p3.Fn2286
func Fn2286(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2289 github.com/goccy/spidermonkeywasm2go/p0.Fn2289
func Fn2289(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2290 github.com/goccy/spidermonkeywasm2go/p4.Fn2290
func Fn2290(m *base.Module, l0 int32)

//go:linkname Fn2293 github.com/goccy/spidermonkeywasm2go/p3.Fn2293
func Fn2293(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2294 github.com/goccy/spidermonkeywasm2go/p3.Fn2294
func Fn2294(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2295 github.com/goccy/spidermonkeywasm2go/p4.Fn2295
func Fn2295(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2297 github.com/goccy/spidermonkeywasm2go/p3.Fn2297
func Fn2297(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2300 github.com/goccy/spidermonkeywasm2go/p4.Fn2300
func Fn2300(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2308 github.com/goccy/spidermonkeywasm2go/p0.Fn2308
func Fn2308(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2311 github.com/goccy/spidermonkeywasm2go/p0.Fn2311
func Fn2311(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2312 github.com/goccy/spidermonkeywasm2go/p0.Fn2312
func Fn2312(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2314 github.com/goccy/spidermonkeywasm2go/p0.Fn2314
func Fn2314(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2318 github.com/goccy/spidermonkeywasm2go/p3.Fn2318
func Fn2318(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2320 github.com/goccy/spidermonkeywasm2go/p0.Fn2320
func Fn2320(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2324 github.com/goccy/spidermonkeywasm2go/p0.Fn2324
func Fn2324(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2328 github.com/goccy/spidermonkeywasm2go/p3.Fn2328
func Fn2328(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2330 github.com/goccy/spidermonkeywasm2go/p4.Fn2330
func Fn2330(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2333 github.com/goccy/spidermonkeywasm2go/p4.Fn2333
func Fn2333(m *base.Module, l0 int32) int32

//go:linkname Fn2334 github.com/goccy/spidermonkeywasm2go/p3.Fn2334
func Fn2334(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2338 github.com/goccy/spidermonkeywasm2go/p4.Fn2338
func Fn2338(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2348 github.com/goccy/spidermonkeywasm2go/p4.Fn2348
func Fn2348(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2351 github.com/goccy/spidermonkeywasm2go/p0.Fn2351
func Fn2351(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2353 github.com/goccy/spidermonkeywasm2go/p0.Fn2353
func Fn2353(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2356 github.com/goccy/spidermonkeywasm2go/p3.Fn2356
func Fn2356(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2358 github.com/goccy/spidermonkeywasm2go/p4.Fn2358
func Fn2358(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2361 github.com/goccy/spidermonkeywasm2go/p4.Fn2361
func Fn2361(m *base.Module, l0 int32)

//go:linkname Fn2364 github.com/goccy/spidermonkeywasm2go/p4.Fn2364
func Fn2364(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2369 github.com/goccy/spidermonkeywasm2go/p0.Fn2369
func Fn2369(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2400 github.com/goccy/spidermonkeywasm2go/p3.Fn2400
func Fn2400(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2405 github.com/goccy/spidermonkeywasm2go/p4.Fn2405
func Fn2405(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2408 github.com/goccy/spidermonkeywasm2go/p4.Fn2408
func Fn2408(m *base.Module, l0 int32)

//go:linkname Fn2410 github.com/goccy/spidermonkeywasm2go/p3.Fn2410
func Fn2410(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2413 github.com/goccy/spidermonkeywasm2go/p0.Fn2413
func Fn2413(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2422 github.com/goccy/spidermonkeywasm2go/p4.Fn2422
func Fn2422(m *base.Module, l0 float64) float64

//go:linkname Fn2456 github.com/goccy/spidermonkeywasm2go/p0.Fn2456
func Fn2456(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2457 github.com/goccy/spidermonkeywasm2go/p0.Fn2457
func Fn2457(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2459 github.com/goccy/spidermonkeywasm2go/p4.Fn2459
func Fn2459(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2461 github.com/goccy/spidermonkeywasm2go/p3.Fn2461
func Fn2461(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2466 github.com/goccy/spidermonkeywasm2go/p4.Fn2466
func Fn2466(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2468 github.com/goccy/spidermonkeywasm2go/p0.Fn2468
func Fn2468(m *base.Module, l0 int32) int32

//go:linkname Fn2471 github.com/goccy/spidermonkeywasm2go/p3.Fn2471
func Fn2471(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2474 github.com/goccy/spidermonkeywasm2go/p4.Fn2474
func Fn2474(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2478 github.com/goccy/spidermonkeywasm2go/p4.Fn2478
func Fn2478(m *base.Module, l0 int32) int32

//go:linkname Fn2479 github.com/goccy/spidermonkeywasm2go/p1.Fn2479
func Fn2479(m *base.Module, l0 int32)

//go:linkname Fn2480 github.com/goccy/spidermonkeywasm2go/p3.Fn2480
func Fn2480(m *base.Module, l0 int32) int32

//go:linkname Fn2481 github.com/goccy/spidermonkeywasm2go/p4.Fn2481
func Fn2481(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2482 github.com/goccy/spidermonkeywasm2go/p4.Fn2482
func Fn2482(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2491 github.com/goccy/spidermonkeywasm2go/p0.Fn2491
func Fn2491(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2510 github.com/goccy/spidermonkeywasm2go/p0.Fn2510
func Fn2510(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2519 github.com/goccy/spidermonkeywasm2go/p4.Fn2519
func Fn2519(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2581 github.com/goccy/spidermonkeywasm2go/p4.Fn2581
func Fn2581(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2583 github.com/goccy/spidermonkeywasm2go/p3.Fn2583
func Fn2583(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2634 github.com/goccy/spidermonkeywasm2go/p0.Fn2634
func Fn2634(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2639 github.com/goccy/spidermonkeywasm2go/p0.Fn2639
func Fn2639(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn2641 github.com/goccy/spidermonkeywasm2go/p3.Fn2641
func Fn2641(m *base.Module, l0 int32)

//go:linkname Fn2644 github.com/goccy/spidermonkeywasm2go/p4.Fn2644
func Fn2644(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2660 github.com/goccy/spidermonkeywasm2go/p0.Fn2660
func Fn2660(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2665 github.com/goccy/spidermonkeywasm2go/p4.Fn2665
func Fn2665(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2666 github.com/goccy/spidermonkeywasm2go/p4.Fn2666
func Fn2666(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2668 github.com/goccy/spidermonkeywasm2go/p4.Fn2668
func Fn2668(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2679 github.com/goccy/spidermonkeywasm2go/p0.Fn2679
func Fn2679(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2693 github.com/goccy/spidermonkeywasm2go/p4.Fn2693
func Fn2693(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2694 github.com/goccy/spidermonkeywasm2go/p4.Fn2694
func Fn2694(m *base.Module, l0 int32)

//go:linkname Fn2706 github.com/goccy/spidermonkeywasm2go/p4.Fn2706
func Fn2706(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2710 github.com/goccy/spidermonkeywasm2go/p4.Fn2710
func Fn2710(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2895 github.com/goccy/spidermonkeywasm2go/p4.Fn2895
func Fn2895(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2896 github.com/goccy/spidermonkeywasm2go/p4.Fn2896
func Fn2896(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2897 github.com/goccy/spidermonkeywasm2go/p1.Fn2897
func Fn2897(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2898 github.com/goccy/spidermonkeywasm2go/p0.Fn2898
func Fn2898(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2900 github.com/goccy/spidermonkeywasm2go/p4.Fn2900
func Fn2900(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2902 github.com/goccy/spidermonkeywasm2go/p4.Fn2902
func Fn2902(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2907 github.com/goccy/spidermonkeywasm2go/p4.Fn2907
func Fn2907(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2908 github.com/goccy/spidermonkeywasm2go/p4.Fn2908
func Fn2908(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2909 github.com/goccy/spidermonkeywasm2go/p4.Fn2909
func Fn2909(m *base.Module, l0 int32) int32

//go:linkname Fn2910 github.com/goccy/spidermonkeywasm2go/p0.Fn2910
func Fn2910(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2913 github.com/goccy/spidermonkeywasm2go/p0.Fn2913
func Fn2913(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2923 github.com/goccy/spidermonkeywasm2go/p0.Fn2923
func Fn2923(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2924 github.com/goccy/spidermonkeywasm2go/p0.Fn2924
func Fn2924(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2926 github.com/goccy/spidermonkeywasm2go/p3.Fn2926
func Fn2926(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2928 github.com/goccy/spidermonkeywasm2go/p0.Fn2928
func Fn2928(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2930 github.com/goccy/spidermonkeywasm2go/p0.Fn2930
func Fn2930(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2931 github.com/goccy/spidermonkeywasm2go/p0.Fn2931
func Fn2931(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2933 github.com/goccy/spidermonkeywasm2go/p0.Fn2933
func Fn2933(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2934 github.com/goccy/spidermonkeywasm2go/p0.Fn2934
func Fn2934(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2937 github.com/goccy/spidermonkeywasm2go/p0.Fn2937
func Fn2937(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2938 github.com/goccy/spidermonkeywasm2go/p0.Fn2938
func Fn2938(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2945 github.com/goccy/spidermonkeywasm2go/p3.Fn2945
func Fn2945(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2951 github.com/goccy/spidermonkeywasm2go/p4.Fn2951
func Fn2951(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2956 github.com/goccy/spidermonkeywasm2go/p4.Fn2956
func Fn2956(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2957 github.com/goccy/spidermonkeywasm2go/p4.Fn2957
func Fn2957(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2984 github.com/goccy/spidermonkeywasm2go/p0.Fn2984
func Fn2984(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2994 github.com/goccy/spidermonkeywasm2go/p4.Fn2994
func Fn2994(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2995 github.com/goccy/spidermonkeywasm2go/p3.Fn2995
func Fn2995(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3014 github.com/goccy/spidermonkeywasm2go/p0.Fn3014
func Fn3014(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3070 github.com/goccy/spidermonkeywasm2go/p4.Fn3070
func Fn3070(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3071 github.com/goccy/spidermonkeywasm2go/p0.Fn3071
func Fn3071(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3072 github.com/goccy/spidermonkeywasm2go/p0.Fn3072
func Fn3072(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3073 github.com/goccy/spidermonkeywasm2go/p0.Fn3073
func Fn3073(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3074 github.com/goccy/spidermonkeywasm2go/p0.Fn3074
func Fn3074(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3075 github.com/goccy/spidermonkeywasm2go/p0.Fn3075
func Fn3075(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3076 github.com/goccy/spidermonkeywasm2go/p0.Fn3076
func Fn3076(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3077 github.com/goccy/spidermonkeywasm2go/p0.Fn3077
func Fn3077(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3078 github.com/goccy/spidermonkeywasm2go/p0.Fn3078
func Fn3078(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3079 github.com/goccy/spidermonkeywasm2go/p0.Fn3079
func Fn3079(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3080 github.com/goccy/spidermonkeywasm2go/p0.Fn3080
func Fn3080(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3081 github.com/goccy/spidermonkeywasm2go/p0.Fn3081
func Fn3081(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3082 github.com/goccy/spidermonkeywasm2go/p0.Fn3082
func Fn3082(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3110 github.com/goccy/spidermonkeywasm2go/p0.Fn3110
func Fn3110(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3124 github.com/goccy/spidermonkeywasm2go/p0.Fn3124
func Fn3124(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3126 github.com/goccy/spidermonkeywasm2go/p0.Fn3126
func Fn3126(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3127 github.com/goccy/spidermonkeywasm2go/p3.Fn3127
func Fn3127(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3128 github.com/goccy/spidermonkeywasm2go/p0.Fn3128
func Fn3128(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3130 github.com/goccy/spidermonkeywasm2go/p0.Fn3130
func Fn3130(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3132 github.com/goccy/spidermonkeywasm2go/p0.Fn3132
func Fn3132(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3134 github.com/goccy/spidermonkeywasm2go/p0.Fn3134
func Fn3134(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3136 github.com/goccy/spidermonkeywasm2go/p0.Fn3136
func Fn3136(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3138 github.com/goccy/spidermonkeywasm2go/p0.Fn3138
func Fn3138(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3140 github.com/goccy/spidermonkeywasm2go/p0.Fn3140
func Fn3140(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3142 github.com/goccy/spidermonkeywasm2go/p0.Fn3142
func Fn3142(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3144 github.com/goccy/spidermonkeywasm2go/p0.Fn3144
func Fn3144(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3146 github.com/goccy/spidermonkeywasm2go/p0.Fn3146
func Fn3146(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3148 github.com/goccy/spidermonkeywasm2go/p3.Fn3148
func Fn3148(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3149 github.com/goccy/spidermonkeywasm2go/p3.Fn3149
func Fn3149(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3151 github.com/goccy/spidermonkeywasm2go/p4.Fn3151
func Fn3151(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3191 github.com/goccy/spidermonkeywasm2go/p4.Fn3191
func Fn3191(m *base.Module, l0 int32)

//go:linkname Fn3283 github.com/goccy/spidermonkeywasm2go/p4.Fn3283
func Fn3283(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3286 github.com/goccy/spidermonkeywasm2go/p3.Fn3286
func Fn3286(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn3287 github.com/goccy/spidermonkeywasm2go/p3.Fn3287
func Fn3287(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn3288 github.com/goccy/spidermonkeywasm2go/p3.Fn3288
func Fn3288(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn3289 github.com/goccy/spidermonkeywasm2go/p3.Fn3289
func Fn3289(m *base.Module, l0 int32, l1 float32, l2 int32, l3 int32)

//go:linkname Fn3290 github.com/goccy/spidermonkeywasm2go/p3.Fn3290
func Fn3290(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32)

//go:linkname Fn3291 github.com/goccy/spidermonkeywasm2go/p3.Fn3291
func Fn3291(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn3292 github.com/goccy/spidermonkeywasm2go/p3.Fn3292
func Fn3292(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32)

//go:linkname Fn3293 github.com/goccy/spidermonkeywasm2go/p4.Fn3293
func Fn3293(m *base.Module, l0 int64) int32

//go:linkname Fn3294 github.com/goccy/spidermonkeywasm2go/p3.Fn3294
func Fn3294(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3295 github.com/goccy/spidermonkeywasm2go/p4.Fn3295
func Fn3295(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn3296 github.com/goccy/spidermonkeywasm2go/p3.Fn3296
func Fn3296(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn3297 github.com/goccy/spidermonkeywasm2go/p3.Fn3297
func Fn3297(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn3298 github.com/goccy/spidermonkeywasm2go/p3.Fn3298
func Fn3298(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn3299 github.com/goccy/spidermonkeywasm2go/p3.Fn3299
func Fn3299(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn3300 github.com/goccy/spidermonkeywasm2go/p3.Fn3300
func Fn3300(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn3301 github.com/goccy/spidermonkeywasm2go/p3.Fn3301
func Fn3301(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn3302 github.com/goccy/spidermonkeywasm2go/p3.Fn3302
func Fn3302(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn3303 github.com/goccy/spidermonkeywasm2go/p3.Fn3303
func Fn3303(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn3304 github.com/goccy/spidermonkeywasm2go/p3.Fn3304
func Fn3304(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn3305 github.com/goccy/spidermonkeywasm2go/p4.Fn3305
func Fn3305(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64

//go:linkname Fn3313 github.com/goccy/spidermonkeywasm2go/p4.Fn3313
func Fn3313(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3314 github.com/goccy/spidermonkeywasm2go/p4.Fn3314
func Fn3314(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3315 github.com/goccy/spidermonkeywasm2go/p4.Fn3315
func Fn3315(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3316 github.com/goccy/spidermonkeywasm2go/p4.Fn3316
func Fn3316(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3345 github.com/goccy/spidermonkeywasm2go/p3.Fn3345
func Fn3345(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3346 github.com/goccy/spidermonkeywasm2go/p3.Fn3346
func Fn3346(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3349 github.com/goccy/spidermonkeywasm2go/p4.Fn3349
func Fn3349(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3352 github.com/goccy/spidermonkeywasm2go/p4.Fn3352
func Fn3352(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3353 github.com/goccy/spidermonkeywasm2go/p4.Fn3353
func Fn3353(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3368 github.com/goccy/spidermonkeywasm2go/p4.Fn3368
func Fn3368(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3369 github.com/goccy/spidermonkeywasm2go/p3.Fn3369
func Fn3369(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn3370 github.com/goccy/spidermonkeywasm2go/p3.Fn3370
func Fn3370(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3374 github.com/goccy/spidermonkeywasm2go/p3.Fn3374
func Fn3374(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3375 github.com/goccy/spidermonkeywasm2go/p4.Fn3375
func Fn3375(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3376 github.com/goccy/spidermonkeywasm2go/p3.Fn3376
func Fn3376(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn3377 github.com/goccy/spidermonkeywasm2go/p3.Fn3377
func Fn3377(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3391 github.com/goccy/spidermonkeywasm2go/p4.Fn3391
func Fn3391(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3392 github.com/goccy/spidermonkeywasm2go/p0.Fn3392
func Fn3392(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3404 github.com/goccy/spidermonkeywasm2go/p3.Fn3404
func Fn3404(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3438 github.com/goccy/spidermonkeywasm2go/p3.Fn3438
func Fn3438(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3491 github.com/goccy/spidermonkeywasm2go/p4.Fn3491
func Fn3491(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3492 github.com/goccy/spidermonkeywasm2go/p4.Fn3492
func Fn3492(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3502 github.com/goccy/spidermonkeywasm2go/p4.Fn3502
func Fn3502(m *base.Module, l0 int32)

//go:linkname Fn3504 github.com/goccy/spidermonkeywasm2go/p4.Fn3504
func Fn3504(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3516 github.com/goccy/spidermonkeywasm2go/p4.Fn3516
func Fn3516(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3539 github.com/goccy/spidermonkeywasm2go/p3.Fn3539
func Fn3539(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3540 github.com/goccy/spidermonkeywasm2go/p4.Fn3540
func Fn3540(m *base.Module, l0 int32) int32

//go:linkname Fn3541 github.com/goccy/spidermonkeywasm2go/p4.Fn3541
func Fn3541(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32

//go:linkname Fn3542 github.com/goccy/spidermonkeywasm2go/p4.Fn3542
func Fn3542(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3546 github.com/goccy/spidermonkeywasm2go/p4.Fn3546
func Fn3546(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3547 github.com/goccy/spidermonkeywasm2go/p4.Fn3547
func Fn3547(m *base.Module, l0 int32)

//go:linkname Fn3554 github.com/goccy/spidermonkeywasm2go/p4.Fn3554
func Fn3554(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3555 github.com/goccy/spidermonkeywasm2go/p4.Fn3555
func Fn3555(m *base.Module, l0 int32) int32

//go:linkname Fn3559 github.com/goccy/spidermonkeywasm2go/p3.Fn3559
func Fn3559(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3565 github.com/goccy/spidermonkeywasm2go/p4.Fn3565
func Fn3565(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3566 github.com/goccy/spidermonkeywasm2go/p4.Fn3566
func Fn3566(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3567 github.com/goccy/spidermonkeywasm2go/p4.Fn3567
func Fn3567(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3568 github.com/goccy/spidermonkeywasm2go/p4.Fn3568
func Fn3568(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3569 github.com/goccy/spidermonkeywasm2go/p4.Fn3569
func Fn3569(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3572 github.com/goccy/spidermonkeywasm2go/p4.Fn3572
func Fn3572(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3575 github.com/goccy/spidermonkeywasm2go/p3.Fn3575
func Fn3575(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3576 github.com/goccy/spidermonkeywasm2go/p4.Fn3576
func Fn3576(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3577 github.com/goccy/spidermonkeywasm2go/p3.Fn3577
func Fn3577(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3578 github.com/goccy/spidermonkeywasm2go/p3.Fn3578
func Fn3578(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3579 github.com/goccy/spidermonkeywasm2go/p4.Fn3579
func Fn3579(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3587 github.com/goccy/spidermonkeywasm2go/p4.Fn3587
func Fn3587(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3588 github.com/goccy/spidermonkeywasm2go/p4.Fn3588
func Fn3588(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3589 github.com/goccy/spidermonkeywasm2go/p4.Fn3589
func Fn3589(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3592 github.com/goccy/spidermonkeywasm2go/p3.Fn3592
func Fn3592(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3593 github.com/goccy/spidermonkeywasm2go/p4.Fn3593
func Fn3593(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3594 github.com/goccy/spidermonkeywasm2go/p4.Fn3594
func Fn3594(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3595 github.com/goccy/spidermonkeywasm2go/p3.Fn3595
func Fn3595(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3596 github.com/goccy/spidermonkeywasm2go/p3.Fn3596
func Fn3596(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3597 github.com/goccy/spidermonkeywasm2go/p3.Fn3597
func Fn3597(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3601 github.com/goccy/spidermonkeywasm2go/p4.Fn3601
func Fn3601(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3602 github.com/goccy/spidermonkeywasm2go/p4.Fn3602
func Fn3602(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3619 github.com/goccy/spidermonkeywasm2go/p3.Fn3619
func Fn3619(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 float64, l7 int32) int32

//go:linkname Fn3628 github.com/goccy/spidermonkeywasm2go/p1.Fn3628
func Fn3628(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3629 github.com/goccy/spidermonkeywasm2go/p4.Fn3629
func Fn3629(m *base.Module, l0 int32)

//go:linkname Fn3644 github.com/goccy/spidermonkeywasm2go/p4.Fn3644
func Fn3644(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3653 github.com/goccy/spidermonkeywasm2go/p4.Fn3653
func Fn3653(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3669 github.com/goccy/spidermonkeywasm2go/p3.Fn3669
func Fn3669(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3671 github.com/goccy/spidermonkeywasm2go/p3.Fn3671
func Fn3671(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3672 github.com/goccy/spidermonkeywasm2go/p3.Fn3672
func Fn3672(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3673 github.com/goccy/spidermonkeywasm2go/p4.Fn3673
func Fn3673(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3674 github.com/goccy/spidermonkeywasm2go/p4.Fn3674
func Fn3674(m *base.Module, l0 int32) int32

//go:linkname Fn3675 github.com/goccy/spidermonkeywasm2go/p4.Fn3675
func Fn3675(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3676 github.com/goccy/spidermonkeywasm2go/p4.Fn3676
func Fn3676(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3677 github.com/goccy/spidermonkeywasm2go/p4.Fn3677
func Fn3677(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3701 github.com/goccy/spidermonkeywasm2go/p3.Fn3701
func Fn3701(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3702 github.com/goccy/spidermonkeywasm2go/p4.Fn3702
func Fn3702(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3707 github.com/goccy/spidermonkeywasm2go/p4.Fn3707
func Fn3707(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3712 github.com/goccy/spidermonkeywasm2go/p3.Fn3712
func Fn3712(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3713 github.com/goccy/spidermonkeywasm2go/p3.Fn3713
func Fn3713(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3715 github.com/goccy/spidermonkeywasm2go/p4.Fn3715
func Fn3715(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3718 github.com/goccy/spidermonkeywasm2go/p3.Fn3718
func Fn3718(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3723 github.com/goccy/spidermonkeywasm2go/p3.Fn3723
func Fn3723(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3732 github.com/goccy/spidermonkeywasm2go/p4.Fn3732
func Fn3732(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3733 github.com/goccy/spidermonkeywasm2go/p4.Fn3733
func Fn3733(m *base.Module, l0 int32)

//go:linkname Fn3734 github.com/goccy/spidermonkeywasm2go/p3.Fn3734
func Fn3734(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3735 github.com/goccy/spidermonkeywasm2go/p4.Fn3735
func Fn3735(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3736 github.com/goccy/spidermonkeywasm2go/p4.Fn3736
func Fn3736(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3737 github.com/goccy/spidermonkeywasm2go/p3.Fn3737
func Fn3737(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3740 github.com/goccy/spidermonkeywasm2go/p4.Fn3740
func Fn3740(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3742 github.com/goccy/spidermonkeywasm2go/p0.Fn3742
func Fn3742(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3743 github.com/goccy/spidermonkeywasm2go/p0.Fn3743
func Fn3743(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3744 github.com/goccy/spidermonkeywasm2go/p0.Fn3744
func Fn3744(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3746 github.com/goccy/spidermonkeywasm2go/p4.Fn3746
func Fn3746(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3747 github.com/goccy/spidermonkeywasm2go/p4.Fn3747
func Fn3747(m *base.Module, l0 int32)

//go:linkname Fn3748 github.com/goccy/spidermonkeywasm2go/p0.Fn3748
func Fn3748(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3754 github.com/goccy/spidermonkeywasm2go/p0.Fn3754
func Fn3754(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3758 github.com/goccy/spidermonkeywasm2go/p4.Fn3758
func Fn3758(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3760 github.com/goccy/spidermonkeywasm2go/p4.Fn3760
func Fn3760(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3761 github.com/goccy/spidermonkeywasm2go/p0.Fn3761
func Fn3761(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3762 github.com/goccy/spidermonkeywasm2go/p0.Fn3762
func Fn3762(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3763 github.com/goccy/spidermonkeywasm2go/p4.Fn3763
func Fn3763(m *base.Module, l0 int32)

//go:linkname Fn3764 github.com/goccy/spidermonkeywasm2go/p4.Fn3764
func Fn3764(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3765 github.com/goccy/spidermonkeywasm2go/p0.Fn3765
func Fn3765(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3767 github.com/goccy/spidermonkeywasm2go/p4.Fn3767
func Fn3767(m *base.Module, l0 int32) int32

//go:linkname Fn3772 github.com/goccy/spidermonkeywasm2go/p4.Fn3772
func Fn3772(m *base.Module, l0 int32)

//go:linkname Fn3773 github.com/goccy/spidermonkeywasm2go/p3.Fn3773
func Fn3773(m *base.Module, l0 float64, l1 float64, l2 float64) float64

//go:linkname Fn3776 github.com/goccy/spidermonkeywasm2go/p3.Fn3776
func Fn3776(m *base.Module, l0 int64) int32

//go:linkname Fn3779 github.com/goccy/spidermonkeywasm2go/p4.Fn3779
func Fn3779(m *base.Module, l0 float64) float64

//go:linkname Fn3782 github.com/goccy/spidermonkeywasm2go/p0.Fn3782
func Fn3782(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3786 github.com/goccy/spidermonkeywasm2go/p4.Fn3786
func Fn3786(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3790 github.com/goccy/spidermonkeywasm2go/p4.Fn3790
func Fn3790(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3795 github.com/goccy/spidermonkeywasm2go/p0.Fn3795
func Fn3795(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3805 github.com/goccy/spidermonkeywasm2go/p4.Fn3805
func Fn3805(m *base.Module, l0 int32, l1 float64, l2 int32) int32

//go:linkname Fn3810 github.com/goccy/spidermonkeywasm2go/p0.Fn3810
func Fn3810(m *base.Module, l0 int32, l1 float64) int32

//go:linkname Fn3820 github.com/goccy/spidermonkeywasm2go/p0.Fn3820
func Fn3820(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3822 github.com/goccy/spidermonkeywasm2go/p0.Fn3822
func Fn3822(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3825 github.com/goccy/spidermonkeywasm2go/p4.Fn3825
func Fn3825(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3831 github.com/goccy/spidermonkeywasm2go/p4.Fn3831
func Fn3831(m *base.Module, l0 int32, l1 int32, l2 float64) int32

//go:linkname Fn3832 github.com/goccy/spidermonkeywasm2go/p1.Fn3832
func Fn3832(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3833 github.com/goccy/spidermonkeywasm2go/p3.Fn3833
func Fn3833(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64) int32

//go:linkname Fn3834 github.com/goccy/spidermonkeywasm2go/p4.Fn3834
func Fn3834(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3896 github.com/goccy/spidermonkeywasm2go/p4.Fn3896
func Fn3896(m *base.Module, l0 int32) int32

//go:linkname Fn3897 github.com/goccy/spidermonkeywasm2go/p4.Fn3897
func Fn3897(m *base.Module, l0 int32)

//go:linkname Fn3898 github.com/goccy/spidermonkeywasm2go/p4.Fn3898
func Fn3898(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3899 github.com/goccy/spidermonkeywasm2go/p3.Fn3899
func Fn3899(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3901 github.com/goccy/spidermonkeywasm2go/p3.Fn3901
func Fn3901(m *base.Module, l0 int32) int32

//go:linkname Fn3908 github.com/goccy/spidermonkeywasm2go/p4.Fn3908
func Fn3908(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3909 github.com/goccy/spidermonkeywasm2go/p4.Fn3909
func Fn3909(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3935 github.com/goccy/spidermonkeywasm2go/p3.Fn3935
func Fn3935(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3940 github.com/goccy/spidermonkeywasm2go/p3.Fn3940
func Fn3940(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3941 github.com/goccy/spidermonkeywasm2go/p0.Fn3941
func Fn3941(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3945 github.com/goccy/spidermonkeywasm2go/p3.Fn3945
func Fn3945(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3946 github.com/goccy/spidermonkeywasm2go/p0.Fn3946
func Fn3946(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3948 github.com/goccy/spidermonkeywasm2go/p4.Fn3948
func Fn3948(m *base.Module, l0 int32)

//go:linkname Fn3959 github.com/goccy/spidermonkeywasm2go/p4.Fn3959
func Fn3959(m *base.Module, l0 int32)

//go:linkname Fn3976 github.com/goccy/spidermonkeywasm2go/p4.Fn3976
func Fn3976(m *base.Module, l0 int32)

//go:linkname Fn3979 github.com/goccy/spidermonkeywasm2go/p3.Fn3979
func Fn3979(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3984 github.com/goccy/spidermonkeywasm2go/p3.Fn3984
func Fn3984(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3986 github.com/goccy/spidermonkeywasm2go/p3.Fn3986
func Fn3986(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3993 github.com/goccy/spidermonkeywasm2go/p3.Fn3993
func Fn3993(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4005 github.com/goccy/spidermonkeywasm2go/p4.Fn4005
func Fn4005(m *base.Module, l0 int32)

//go:linkname Fn4008 github.com/goccy/spidermonkeywasm2go/p3.Fn4008
func Fn4008(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4013 github.com/goccy/spidermonkeywasm2go/p3.Fn4013
func Fn4013(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4044 github.com/goccy/spidermonkeywasm2go/p4.Fn4044
func Fn4044(m *base.Module, l0 int32)

//go:linkname Fn4046 github.com/goccy/spidermonkeywasm2go/p0.Fn4046
func Fn4046(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4048 github.com/goccy/spidermonkeywasm2go/p3.Fn4048
func Fn4048(m *base.Module, l0 int32) int32

//go:linkname Fn4049 github.com/goccy/spidermonkeywasm2go/p4.Fn4049
func Fn4049(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4055 github.com/goccy/spidermonkeywasm2go/p4.Fn4055
func Fn4055(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn4057 github.com/goccy/spidermonkeywasm2go/p3.Fn4057
func Fn4057(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4068 github.com/goccy/spidermonkeywasm2go/p4.Fn4068
func Fn4068(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4306 github.com/goccy/spidermonkeywasm2go/p0.Fn4306
func Fn4306(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn4376 github.com/goccy/spidermonkeywasm2go/p3.Fn4376
func Fn4376(m *base.Module, l0 int32) int32

//go:linkname Fn4387 github.com/goccy/spidermonkeywasm2go/p4.Fn4387
func Fn4387(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4388 github.com/goccy/spidermonkeywasm2go/p1.Fn4388
func Fn4388(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn4389 github.com/goccy/spidermonkeywasm2go/p4.Fn4389
func Fn4389(m *base.Module, l0 int32) int32

//go:linkname Fn4390 github.com/goccy/spidermonkeywasm2go/p3.Fn4390
func Fn4390(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4396 github.com/goccy/spidermonkeywasm2go/p4.Fn4396
func Fn4396(m *base.Module, l0 int32)

//go:linkname Fn4398 github.com/goccy/spidermonkeywasm2go/p0.Fn4398
func Fn4398(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn4399 github.com/goccy/spidermonkeywasm2go/p0.Fn4399
func Fn4399(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn4400 github.com/goccy/spidermonkeywasm2go/p3.Fn4400
func Fn4400(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4402 github.com/goccy/spidermonkeywasm2go/p0.Fn4402
func Fn4402(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4523 github.com/goccy/spidermonkeywasm2go/p0.Fn4523
func Fn4523(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4525 github.com/goccy/spidermonkeywasm2go/p3.Fn4525
func Fn4525(m *base.Module, l0 int32)

//go:linkname Fn4526 github.com/goccy/spidermonkeywasm2go/p0.Fn4526
func Fn4526(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4577 github.com/goccy/spidermonkeywasm2go/p4.Fn4577
func Fn4577(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn4578 github.com/goccy/spidermonkeywasm2go/p4.Fn4578
func Fn4578(m *base.Module, l0 int32) int32

//go:linkname Fn4585 github.com/goccy/spidermonkeywasm2go/p3.Fn4585
func Fn4585(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn4586 github.com/goccy/spidermonkeywasm2go/p3.Fn4586
func Fn4586(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4588 github.com/goccy/spidermonkeywasm2go/p3.Fn4588
func Fn4588(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4589 github.com/goccy/spidermonkeywasm2go/p3.Fn4589
func Fn4589(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4591 github.com/goccy/spidermonkeywasm2go/p3.Fn4591
func Fn4591(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4592 github.com/goccy/spidermonkeywasm2go/p3.Fn4592
func Fn4592(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4596 github.com/goccy/spidermonkeywasm2go/p3.Fn4596
func Fn4596(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4598 github.com/goccy/spidermonkeywasm2go/p4.Fn4598
func Fn4598(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4614 github.com/goccy/spidermonkeywasm2go/p4.Fn4614
func Fn4614(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4640 github.com/goccy/spidermonkeywasm2go/p1.Fn4640
func Fn4640(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn4652 github.com/goccy/spidermonkeywasm2go/p4.Fn4652
func Fn4652(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4660 github.com/goccy/spidermonkeywasm2go/p3.Fn4660
func Fn4660(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4662 github.com/goccy/spidermonkeywasm2go/p0.Fn4662
func Fn4662(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4668 github.com/goccy/spidermonkeywasm2go/p4.Fn4668
func Fn4668(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4675 github.com/goccy/spidermonkeywasm2go/p4.Fn4675
func Fn4675(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4677 github.com/goccy/spidermonkeywasm2go/p4.Fn4677
func Fn4677(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4708 github.com/goccy/spidermonkeywasm2go/p4.Fn4708
func Fn4708(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4774 github.com/goccy/spidermonkeywasm2go/p0.Fn4774
func Fn4774(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4787 github.com/goccy/spidermonkeywasm2go/p4.Fn4787
func Fn4787(m *base.Module, l0 int32) int32

//go:linkname Fn4788 github.com/goccy/spidermonkeywasm2go/p4.Fn4788
func Fn4788(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4789 github.com/goccy/spidermonkeywasm2go/p4.Fn4789
func Fn4789(m *base.Module, l0 int32) int32

//go:linkname Fn4806 github.com/goccy/spidermonkeywasm2go/p3.Fn4806
func Fn4806(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4807 github.com/goccy/spidermonkeywasm2go/p4.Fn4807
func Fn4807(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4810 github.com/goccy/spidermonkeywasm2go/p4.Fn4810
func Fn4810(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4811 github.com/goccy/spidermonkeywasm2go/p4.Fn4811
func Fn4811(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4813 github.com/goccy/spidermonkeywasm2go/p0.Fn4813
func Fn4813(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4814 github.com/goccy/spidermonkeywasm2go/p0.Fn4814
func Fn4814(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4815 github.com/goccy/spidermonkeywasm2go/p4.Fn4815
func Fn4815(m *base.Module, l0 int32)

//go:linkname Fn4816 github.com/goccy/spidermonkeywasm2go/p4.Fn4816
func Fn4816(m *base.Module, l0 int32)

//go:linkname Fn4817 github.com/goccy/spidermonkeywasm2go/p3.Fn4817
func Fn4817(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4818 github.com/goccy/spidermonkeywasm2go/p4.Fn4818
func Fn4818(m *base.Module, l0 int32)

//go:linkname Fn4821 github.com/goccy/spidermonkeywasm2go/p4.Fn4821
func Fn4821(m *base.Module, l0 int32)

//go:linkname Fn4824 github.com/goccy/spidermonkeywasm2go/p4.Fn4824
func Fn4824(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4826 github.com/goccy/spidermonkeywasm2go/p3.Fn4826
func Fn4826(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4831 github.com/goccy/spidermonkeywasm2go/p4.Fn4831
func Fn4831(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4835 github.com/goccy/spidermonkeywasm2go/p3.Fn4835
func Fn4835(m *base.Module, l0 int32)

//go:linkname Fn4836 github.com/goccy/spidermonkeywasm2go/p4.Fn4836
func Fn4836(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4837 github.com/goccy/spidermonkeywasm2go/p4.Fn4837
func Fn4837(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4838 github.com/goccy/spidermonkeywasm2go/p4.Fn4838
func Fn4838(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4839 github.com/goccy/spidermonkeywasm2go/p4.Fn4839
func Fn4839(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4842 github.com/goccy/spidermonkeywasm2go/p4.Fn4842
func Fn4842(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4843 github.com/goccy/spidermonkeywasm2go/p3.Fn4843
func Fn4843(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4844 github.com/goccy/spidermonkeywasm2go/p4.Fn4844
func Fn4844(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4846 github.com/goccy/spidermonkeywasm2go/p4.Fn4846
func Fn4846(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn4849 github.com/goccy/spidermonkeywasm2go/p4.Fn4849
func Fn4849(m *base.Module, l0 int32)

//go:linkname Fn4863 github.com/goccy/spidermonkeywasm2go/p1.Fn4863
func Fn4863(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4866 github.com/goccy/spidermonkeywasm2go/p4.Fn4866
func Fn4866(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn4867 github.com/goccy/spidermonkeywasm2go/p4.Fn4867
func Fn4867(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn4869 github.com/goccy/spidermonkeywasm2go/p4.Fn4869
func Fn4869(m *base.Module, l0 int32)

//go:linkname Fn4870 github.com/goccy/spidermonkeywasm2go/p4.Fn4870
func Fn4870(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4871 github.com/goccy/spidermonkeywasm2go/p4.Fn4871
func Fn4871(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4872 github.com/goccy/spidermonkeywasm2go/p3.Fn4872
func Fn4872(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4873 github.com/goccy/spidermonkeywasm2go/p3.Fn4873
func Fn4873(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4875 github.com/goccy/spidermonkeywasm2go/p4.Fn4875
func Fn4875(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4876 github.com/goccy/spidermonkeywasm2go/p3.Fn4876
func Fn4876(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4887 github.com/goccy/spidermonkeywasm2go/p4.Fn4887
func Fn4887(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4894 github.com/goccy/spidermonkeywasm2go/p3.Fn4894
func Fn4894(m *base.Module, l0 int32)

//go:linkname Fn4895 github.com/goccy/spidermonkeywasm2go/p4.Fn4895
func Fn4895(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4896 github.com/goccy/spidermonkeywasm2go/p4.Fn4896
func Fn4896(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4897 github.com/goccy/spidermonkeywasm2go/p4.Fn4897
func Fn4897(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4899 github.com/goccy/spidermonkeywasm2go/p4.Fn4899
func Fn4899(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4905 github.com/goccy/spidermonkeywasm2go/p3.Fn4905
func Fn4905(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn4907 github.com/goccy/spidermonkeywasm2go/p4.Fn4907
func Fn4907(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4909 github.com/goccy/spidermonkeywasm2go/p0.Fn4909
func Fn4909(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4911 github.com/goccy/spidermonkeywasm2go/p4.Fn4911
func Fn4911(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4912 github.com/goccy/spidermonkeywasm2go/p4.Fn4912
func Fn4912(m *base.Module, l0 int32)

//go:linkname Fn4919 github.com/goccy/spidermonkeywasm2go/p4.Fn4919
func Fn4919(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4920 github.com/goccy/spidermonkeywasm2go/p4.Fn4920
func Fn4920(m *base.Module, l0 int32)

//go:linkname Fn4922 github.com/goccy/spidermonkeywasm2go/p3.Fn4922
func Fn4922(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4923 github.com/goccy/spidermonkeywasm2go/p3.Fn4923
func Fn4923(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4925 github.com/goccy/spidermonkeywasm2go/p4.Fn4925
func Fn4925(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4928 github.com/goccy/spidermonkeywasm2go/p3.Fn4928
func Fn4928(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4933 github.com/goccy/spidermonkeywasm2go/p3.Fn4933
func Fn4933(m *base.Module, l0 int32)

//go:linkname Fn4936 github.com/goccy/spidermonkeywasm2go/p4.Fn4936
func Fn4936(m *base.Module, l0 int32)

//go:linkname Fn4937 github.com/goccy/spidermonkeywasm2go/p4.Fn4937
func Fn4937(m *base.Module, l0 int32) int32

//go:linkname Fn4938 github.com/goccy/spidermonkeywasm2go/p4.Fn4938
func Fn4938(m *base.Module, l0 int32) int32

//go:linkname Fn4950 github.com/goccy/spidermonkeywasm2go/p3.Fn4950
func Fn4950(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4951 github.com/goccy/spidermonkeywasm2go/p4.Fn4951
func Fn4951(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4952 github.com/goccy/spidermonkeywasm2go/p3.Fn4952
func Fn4952(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4954 github.com/goccy/spidermonkeywasm2go/p4.Fn4954
func Fn4954(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4955 github.com/goccy/spidermonkeywasm2go/p3.Fn4955
func Fn4955(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4956 github.com/goccy/spidermonkeywasm2go/p4.Fn4956
func Fn4956(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4957 github.com/goccy/spidermonkeywasm2go/p4.Fn4957
func Fn4957(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4958 github.com/goccy/spidermonkeywasm2go/p4.Fn4958
func Fn4958(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4959 github.com/goccy/spidermonkeywasm2go/p4.Fn4959
func Fn4959(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4960 github.com/goccy/spidermonkeywasm2go/p4.Fn4960
func Fn4960(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4961 github.com/goccy/spidermonkeywasm2go/p4.Fn4961
func Fn4961(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4962 github.com/goccy/spidermonkeywasm2go/p3.Fn4962
func Fn4962(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4963 github.com/goccy/spidermonkeywasm2go/p4.Fn4963
func Fn4963(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4964 github.com/goccy/spidermonkeywasm2go/p4.Fn4964
func Fn4964(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4965 github.com/goccy/spidermonkeywasm2go/p4.Fn4965
func Fn4965(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4966 github.com/goccy/spidermonkeywasm2go/p4.Fn4966
func Fn4966(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4967 github.com/goccy/spidermonkeywasm2go/p4.Fn4967
func Fn4967(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4969 github.com/goccy/spidermonkeywasm2go/p4.Fn4969
func Fn4969(m *base.Module, l0 int32)

//go:linkname Fn4970 github.com/goccy/spidermonkeywasm2go/p4.Fn4970
func Fn4970(m *base.Module, l0 int32)

//go:linkname Fn4971 github.com/goccy/spidermonkeywasm2go/p4.Fn4971
func Fn4971(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4975 github.com/goccy/spidermonkeywasm2go/p4.Fn4975
func Fn4975(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4976 github.com/goccy/spidermonkeywasm2go/p4.Fn4976
func Fn4976(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4981 github.com/goccy/spidermonkeywasm2go/p4.Fn4981
func Fn4981(m *base.Module, l0 int32)

//go:linkname Fn4985 github.com/goccy/spidermonkeywasm2go/p4.Fn4985
func Fn4985(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4986 github.com/goccy/spidermonkeywasm2go/p4.Fn4986
func Fn4986(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4991 github.com/goccy/spidermonkeywasm2go/p4.Fn4991
func Fn4991(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4992 github.com/goccy/spidermonkeywasm2go/p4.Fn4992
func Fn4992(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5005 github.com/goccy/spidermonkeywasm2go/p3.Fn5005
func Fn5005(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5006 github.com/goccy/spidermonkeywasm2go/p3.Fn5006
func Fn5006(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5033 github.com/goccy/spidermonkeywasm2go/p4.Fn5033
func Fn5033(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5034 github.com/goccy/spidermonkeywasm2go/p4.Fn5034
func Fn5034(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5035 github.com/goccy/spidermonkeywasm2go/p4.Fn5035
func Fn5035(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5036 github.com/goccy/spidermonkeywasm2go/p4.Fn5036
func Fn5036(m *base.Module, l0 int32)

//go:linkname Fn5050 github.com/goccy/spidermonkeywasm2go/p4.Fn5050
func Fn5050(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5051 github.com/goccy/spidermonkeywasm2go/p4.Fn5051
func Fn5051(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5062 github.com/goccy/spidermonkeywasm2go/p4.Fn5062
func Fn5062(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5092 github.com/goccy/spidermonkeywasm2go/p4.Fn5092
func Fn5092(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5093 github.com/goccy/spidermonkeywasm2go/p4.Fn5093
func Fn5093(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5099 github.com/goccy/spidermonkeywasm2go/p4.Fn5099
func Fn5099(m *base.Module, l0 int32) int32

//go:linkname Fn5125 github.com/goccy/spidermonkeywasm2go/p4.Fn5125
func Fn5125(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5127 github.com/goccy/spidermonkeywasm2go/p4.Fn5127
func Fn5127(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5128 github.com/goccy/spidermonkeywasm2go/p4.Fn5128
func Fn5128(m *base.Module, l0 int32) int32

//go:linkname Fn5131 github.com/goccy/spidermonkeywasm2go/p4.Fn5131
func Fn5131(m *base.Module, l0 int32)

//go:linkname Fn5136 github.com/goccy/spidermonkeywasm2go/p3.Fn5136
func Fn5136(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5155 github.com/goccy/spidermonkeywasm2go/p4.Fn5155
func Fn5155(m *base.Module, l0 int32)

//go:linkname Fn5163 github.com/goccy/spidermonkeywasm2go/p3.Fn5163
func Fn5163(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5164 github.com/goccy/spidermonkeywasm2go/p4.Fn5164
func Fn5164(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5173 github.com/goccy/spidermonkeywasm2go/p4.Fn5173
func Fn5173(m *base.Module, l0 int32)

//go:linkname Fn5174 github.com/goccy/spidermonkeywasm2go/p3.Fn5174
func Fn5174(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5179 github.com/goccy/spidermonkeywasm2go/p4.Fn5179
func Fn5179(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5183 github.com/goccy/spidermonkeywasm2go/p3.Fn5183
func Fn5183(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5188 github.com/goccy/spidermonkeywasm2go/p3.Fn5188
func Fn5188(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5194 github.com/goccy/spidermonkeywasm2go/p4.Fn5194
func Fn5194(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5251 github.com/goccy/spidermonkeywasm2go/p3.Fn5251
func Fn5251(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5252 github.com/goccy/spidermonkeywasm2go/p4.Fn5252
func Fn5252(m *base.Module)

//go:linkname Fn5255 github.com/goccy/spidermonkeywasm2go/p4.Fn5255
func Fn5255(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5257 github.com/goccy/spidermonkeywasm2go/p3.Fn5257
func Fn5257(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5258 github.com/goccy/spidermonkeywasm2go/p4.Fn5258
func Fn5258(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5260 github.com/goccy/spidermonkeywasm2go/p4.Fn5260
func Fn5260(m *base.Module)

//go:linkname Fn5262 github.com/goccy/spidermonkeywasm2go/p3.Fn5262
func Fn5262(m *base.Module, l0 int32)

//go:linkname Fn5263 github.com/goccy/spidermonkeywasm2go/p4.Fn5263
func Fn5263(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5264 github.com/goccy/spidermonkeywasm2go/p4.Fn5264
func Fn5264(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5265 github.com/goccy/spidermonkeywasm2go/p4.Fn5265
func Fn5265(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5266 github.com/goccy/spidermonkeywasm2go/p4.Fn5266
func Fn5266(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5267 github.com/goccy/spidermonkeywasm2go/p4.Fn5267
func Fn5267(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5268 github.com/goccy/spidermonkeywasm2go/p4.Fn5268
func Fn5268(m *base.Module, l0 int32) int32

//go:linkname Fn5269 github.com/goccy/spidermonkeywasm2go/p4.Fn5269
func Fn5269(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5270 github.com/goccy/spidermonkeywasm2go/p3.Fn5270
func Fn5270(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5271 github.com/goccy/spidermonkeywasm2go/p4.Fn5271
func Fn5271(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5272 github.com/goccy/spidermonkeywasm2go/p4.Fn5272
func Fn5272(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5274 github.com/goccy/spidermonkeywasm2go/p4.Fn5274
func Fn5274(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5275 github.com/goccy/spidermonkeywasm2go/p4.Fn5275
func Fn5275(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5276 github.com/goccy/spidermonkeywasm2go/p4.Fn5276
func Fn5276(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5279 github.com/goccy/spidermonkeywasm2go/p4.Fn5279
func Fn5279(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5284 github.com/goccy/spidermonkeywasm2go/p3.Fn5284
func Fn5284(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5287 github.com/goccy/spidermonkeywasm2go/p4.Fn5287
func Fn5287(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5296 github.com/goccy/spidermonkeywasm2go/p3.Fn5296
func Fn5296(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5297 github.com/goccy/spidermonkeywasm2go/p3.Fn5297
func Fn5297(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5298 github.com/goccy/spidermonkeywasm2go/p3.Fn5298
func Fn5298(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5299 github.com/goccy/spidermonkeywasm2go/p3.Fn5299
func Fn5299(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5313 github.com/goccy/spidermonkeywasm2go/p4.Fn5313
func Fn5313(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5316 github.com/goccy/spidermonkeywasm2go/p4.Fn5316
func Fn5316(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5318 github.com/goccy/spidermonkeywasm2go/p4.Fn5318
func Fn5318(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5319 github.com/goccy/spidermonkeywasm2go/p4.Fn5319
func Fn5319(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5329 github.com/goccy/spidermonkeywasm2go/p4.Fn5329
func Fn5329(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5338 github.com/goccy/spidermonkeywasm2go/p4.Fn5338
func Fn5338(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5352 github.com/goccy/spidermonkeywasm2go/p3.Fn5352
func Fn5352(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5355 github.com/goccy/spidermonkeywasm2go/p4.Fn5355
func Fn5355(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5362 github.com/goccy/spidermonkeywasm2go/p4.Fn5362
func Fn5362(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5363 github.com/goccy/spidermonkeywasm2go/p4.Fn5363
func Fn5363(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5364 github.com/goccy/spidermonkeywasm2go/p3.Fn5364
func Fn5364(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn5365 github.com/goccy/spidermonkeywasm2go/p4.Fn5365
func Fn5365(m *base.Module)

//go:linkname Fn5455 github.com/goccy/spidermonkeywasm2go/p4.Fn5455
func Fn5455(m *base.Module, l0 int32) int32

//go:linkname Fn5470 github.com/goccy/spidermonkeywasm2go/p4.Fn5470
func Fn5470(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5471 github.com/goccy/spidermonkeywasm2go/p3.Fn5471
func Fn5471(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5478 github.com/goccy/spidermonkeywasm2go/p4.Fn5478
func Fn5478(m *base.Module, l0 int32)

//go:linkname Fn5480 github.com/goccy/spidermonkeywasm2go/p3.Fn5480
func Fn5480(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5488 github.com/goccy/spidermonkeywasm2go/p3.Fn5488
func Fn5488(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5492 github.com/goccy/spidermonkeywasm2go/p4.Fn5492
func Fn5492(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5494 github.com/goccy/spidermonkeywasm2go/p4.Fn5494
func Fn5494(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5500 github.com/goccy/spidermonkeywasm2go/p3.Fn5500
func Fn5500(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5501 github.com/goccy/spidermonkeywasm2go/p4.Fn5501
func Fn5501(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5503 github.com/goccy/spidermonkeywasm2go/p3.Fn5503
func Fn5503(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn5555 github.com/goccy/spidermonkeywasm2go/p3.Fn5555
func Fn5555(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5564 github.com/goccy/spidermonkeywasm2go/p4.Fn5564
func Fn5564(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5869 github.com/goccy/spidermonkeywasm2go/p3.Fn5869
func Fn5869(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5870 github.com/goccy/spidermonkeywasm2go/p4.Fn5870
func Fn5870(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5871 github.com/goccy/spidermonkeywasm2go/p4.Fn5871
func Fn5871(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5872 github.com/goccy/spidermonkeywasm2go/p4.Fn5872
func Fn5872(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5873 github.com/goccy/spidermonkeywasm2go/p4.Fn5873
func Fn5873(m *base.Module, l0 int32) int32

//go:linkname Fn5886 github.com/goccy/spidermonkeywasm2go/p4.Fn5886
func Fn5886(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5893 github.com/goccy/spidermonkeywasm2go/p1.Fn5893
func Fn5893(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5895 github.com/goccy/spidermonkeywasm2go/p4.Fn5895
func Fn5895(m *base.Module, l0 int32)

//go:linkname Fn5896 github.com/goccy/spidermonkeywasm2go/p4.Fn5896
func Fn5896(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn5898 github.com/goccy/spidermonkeywasm2go/p4.Fn5898
func Fn5898(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5900 github.com/goccy/spidermonkeywasm2go/p3.Fn5900
func Fn5900(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5901 github.com/goccy/spidermonkeywasm2go/p4.Fn5901
func Fn5901(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5904 github.com/goccy/spidermonkeywasm2go/p4.Fn5904
func Fn5904(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5905 github.com/goccy/spidermonkeywasm2go/p4.Fn5905
func Fn5905(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5906 github.com/goccy/spidermonkeywasm2go/p4.Fn5906
func Fn5906(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5911 github.com/goccy/spidermonkeywasm2go/p3.Fn5911
func Fn5911(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5912 github.com/goccy/spidermonkeywasm2go/p4.Fn5912
func Fn5912(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn5913 github.com/goccy/spidermonkeywasm2go/p4.Fn5913
func Fn5913(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5914 github.com/goccy/spidermonkeywasm2go/p4.Fn5914
func Fn5914(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5915 github.com/goccy/spidermonkeywasm2go/p4.Fn5915
func Fn5915(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5916 github.com/goccy/spidermonkeywasm2go/p4.Fn5916
func Fn5916(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5917 github.com/goccy/spidermonkeywasm2go/p3.Fn5917
func Fn5917(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5919 github.com/goccy/spidermonkeywasm2go/p4.Fn5919
func Fn5919(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5920 github.com/goccy/spidermonkeywasm2go/p4.Fn5920
func Fn5920(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5922 github.com/goccy/spidermonkeywasm2go/p4.Fn5922
func Fn5922(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5926 github.com/goccy/spidermonkeywasm2go/p3.Fn5926
func Fn5926(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5927 github.com/goccy/spidermonkeywasm2go/p4.Fn5927
func Fn5927(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5928 github.com/goccy/spidermonkeywasm2go/p4.Fn5928
func Fn5928(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5929 github.com/goccy/spidermonkeywasm2go/p4.Fn5929
func Fn5929(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5931 github.com/goccy/spidermonkeywasm2go/p3.Fn5931
func Fn5931(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn5932 github.com/goccy/spidermonkeywasm2go/p4.Fn5932
func Fn5932(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5935 github.com/goccy/spidermonkeywasm2go/p3.Fn5935
func Fn5935(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5937 github.com/goccy/spidermonkeywasm2go/p4.Fn5937
func Fn5937(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5938 github.com/goccy/spidermonkeywasm2go/p4.Fn5938
func Fn5938(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5941 github.com/goccy/spidermonkeywasm2go/p4.Fn5941
func Fn5941(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5942 github.com/goccy/spidermonkeywasm2go/p4.Fn5942
func Fn5942(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5959 github.com/goccy/spidermonkeywasm2go/p4.Fn5959
func Fn5959(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5965 github.com/goccy/spidermonkeywasm2go/p4.Fn5965
func Fn5965(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5968 github.com/goccy/spidermonkeywasm2go/p4.Fn5968
func Fn5968(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5969 github.com/goccy/spidermonkeywasm2go/p4.Fn5969
func Fn5969(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5970 github.com/goccy/spidermonkeywasm2go/p1.Fn5970
func Fn5970(m *base.Module, l0 int32) int32

//go:linkname Fn5976 github.com/goccy/spidermonkeywasm2go/p4.Fn5976
func Fn5976(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5977 github.com/goccy/spidermonkeywasm2go/p4.Fn5977
func Fn5977(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5978 github.com/goccy/spidermonkeywasm2go/p4.Fn5978
func Fn5978(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5979 github.com/goccy/spidermonkeywasm2go/p4.Fn5979
func Fn5979(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5980 github.com/goccy/spidermonkeywasm2go/p4.Fn5980
func Fn5980(m *base.Module, l0 int32, l1 float64) int32

//go:linkname Fn5981 github.com/goccy/spidermonkeywasm2go/p4.Fn5981
func Fn5981(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5982 github.com/goccy/spidermonkeywasm2go/p4.Fn5982
func Fn5982(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5983 github.com/goccy/spidermonkeywasm2go/p4.Fn5983
func Fn5983(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5984 github.com/goccy/spidermonkeywasm2go/p4.Fn5984
func Fn5984(m *base.Module, l0 int32) int32

//go:linkname Fn5988 github.com/goccy/spidermonkeywasm2go/p4.Fn5988
func Fn5988(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5989 github.com/goccy/spidermonkeywasm2go/p4.Fn5989
func Fn5989(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5990 github.com/goccy/spidermonkeywasm2go/p3.Fn5990
func Fn5990(m *base.Module, l0 int32) int32

//go:linkname Fn5991 github.com/goccy/spidermonkeywasm2go/p4.Fn5991
func Fn5991(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5995 github.com/goccy/spidermonkeywasm2go/p4.Fn5995
func Fn5995(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6000 github.com/goccy/spidermonkeywasm2go/p3.Fn6000
func Fn6000(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6001 github.com/goccy/spidermonkeywasm2go/p4.Fn6001
func Fn6001(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6013 github.com/goccy/spidermonkeywasm2go/p4.Fn6013
func Fn6013(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6039 github.com/goccy/spidermonkeywasm2go/p4.Fn6039
func Fn6039(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6040 github.com/goccy/spidermonkeywasm2go/p3.Fn6040
func Fn6040(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6043 github.com/goccy/spidermonkeywasm2go/p1.Fn6043
func Fn6043(m *base.Module, l0 int32) int32

//go:linkname Fn6065 github.com/goccy/spidermonkeywasm2go/p4.Fn6065
func Fn6065(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6066 github.com/goccy/spidermonkeywasm2go/p4.Fn6066
func Fn6066(m *base.Module, l0 int32, l1 int32, l2 int64)

//go:linkname Fn6068 github.com/goccy/spidermonkeywasm2go/p3.Fn6068
func Fn6068(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6069 github.com/goccy/spidermonkeywasm2go/p4.Fn6069
func Fn6069(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6075 github.com/goccy/spidermonkeywasm2go/p3.Fn6075
func Fn6075(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6077 github.com/goccy/spidermonkeywasm2go/p3.Fn6077
func Fn6077(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6079 github.com/goccy/spidermonkeywasm2go/p3.Fn6079
func Fn6079(m *base.Module, l0 int32)

//go:linkname Fn6080 github.com/goccy/spidermonkeywasm2go/p3.Fn6080
func Fn6080(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6086 github.com/goccy/spidermonkeywasm2go/p3.Fn6086
func Fn6086(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6103 github.com/goccy/spidermonkeywasm2go/p4.Fn6103
func Fn6103(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6104 github.com/goccy/spidermonkeywasm2go/p4.Fn6104
func Fn6104(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6106 github.com/goccy/spidermonkeywasm2go/p3.Fn6106
func Fn6106(m *base.Module, l0 int32, l1 int32, l2 int32) int64

//go:linkname Fn6110 github.com/goccy/spidermonkeywasm2go/p4.Fn6110
func Fn6110(m *base.Module, l0 int32) int32

//go:linkname Fn6116 github.com/goccy/spidermonkeywasm2go/p4.Fn6116
func Fn6116(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6123 github.com/goccy/spidermonkeywasm2go/p4.Fn6123
func Fn6123(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6129 github.com/goccy/spidermonkeywasm2go/p3.Fn6129
func Fn6129(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6132 github.com/goccy/spidermonkeywasm2go/p4.Fn6132
func Fn6132(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6136 github.com/goccy/spidermonkeywasm2go/p3.Fn6136
func Fn6136(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6143 github.com/goccy/spidermonkeywasm2go/p4.Fn6143
func Fn6143(m *base.Module, l0 int32)

//go:linkname Fn6145 github.com/goccy/spidermonkeywasm2go/p4.Fn6145
func Fn6145(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6152 github.com/goccy/spidermonkeywasm2go/p4.Fn6152
func Fn6152(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6154 github.com/goccy/spidermonkeywasm2go/p0.Fn6154
func Fn6154(m *base.Module, l0 int32) int32

//go:linkname Fn6157 github.com/goccy/spidermonkeywasm2go/p0.Fn6157
func Fn6157(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6158 github.com/goccy/spidermonkeywasm2go/p0.Fn6158
func Fn6158(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6172 github.com/goccy/spidermonkeywasm2go/p4.Fn6172
func Fn6172(m *base.Module, l0 int32)

//go:linkname Fn6173 github.com/goccy/spidermonkeywasm2go/p3.Fn6173
func Fn6173(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6175 github.com/goccy/spidermonkeywasm2go/p3.Fn6175
func Fn6175(m *base.Module, l0 int32) int32

//go:linkname Fn6176 github.com/goccy/spidermonkeywasm2go/p3.Fn6176
func Fn6176(m *base.Module, l0 int32) int32

//go:linkname Fn6179 github.com/goccy/spidermonkeywasm2go/p4.Fn6179
func Fn6179(m *base.Module, l0 int32)

//go:linkname Fn6183 github.com/goccy/spidermonkeywasm2go/p0.Fn6183
func Fn6183(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6184 github.com/goccy/spidermonkeywasm2go/p4.Fn6184
func Fn6184(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6186 github.com/goccy/spidermonkeywasm2go/p4.Fn6186
func Fn6186(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6188 github.com/goccy/spidermonkeywasm2go/p4.Fn6188
func Fn6188(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6216 github.com/goccy/spidermonkeywasm2go/p4.Fn6216
func Fn6216(m *base.Module, l0 int32)

//go:linkname Fn6218 github.com/goccy/spidermonkeywasm2go/p4.Fn6218
func Fn6218(m *base.Module, l0 int32) int32

//go:linkname Fn6236 github.com/goccy/spidermonkeywasm2go/p4.Fn6236
func Fn6236(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6285 github.com/goccy/spidermonkeywasm2go/p3.Fn6285
func Fn6285(m *base.Module, l0 int32) int32

//go:linkname Fn6288 github.com/goccy/spidermonkeywasm2go/p3.Fn6288
func Fn6288(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6289 github.com/goccy/spidermonkeywasm2go/p4.Fn6289
func Fn6289(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6302 github.com/goccy/spidermonkeywasm2go/p3.Fn6302
func Fn6302(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6319 github.com/goccy/spidermonkeywasm2go/p4.Fn6319
func Fn6319(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6320 github.com/goccy/spidermonkeywasm2go/p4.Fn6320
func Fn6320(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6321 github.com/goccy/spidermonkeywasm2go/p4.Fn6321
func Fn6321(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6322 github.com/goccy/spidermonkeywasm2go/p4.Fn6322
func Fn6322(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6325 github.com/goccy/spidermonkeywasm2go/p3.Fn6325
func Fn6325(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6328 github.com/goccy/spidermonkeywasm2go/p3.Fn6328
func Fn6328(m *base.Module, l0 int32) int32

//go:linkname Fn6331 github.com/goccy/spidermonkeywasm2go/p3.Fn6331
func Fn6331(m *base.Module, l0 int32)

//go:linkname Fn6332 github.com/goccy/spidermonkeywasm2go/p4.Fn6332
func Fn6332(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6335 github.com/goccy/spidermonkeywasm2go/p4.Fn6335
func Fn6335(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6336 github.com/goccy/spidermonkeywasm2go/p4.Fn6336
func Fn6336(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6337 github.com/goccy/spidermonkeywasm2go/p3.Fn6337
func Fn6337(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6342 github.com/goccy/spidermonkeywasm2go/p4.Fn6342
func Fn6342(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6346 github.com/goccy/spidermonkeywasm2go/p4.Fn6346
func Fn6346(m *base.Module, l0 int32) int32

//go:linkname Fn6347 github.com/goccy/spidermonkeywasm2go/p4.Fn6347
func Fn6347(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6364 github.com/goccy/spidermonkeywasm2go/p4.Fn6364
func Fn6364(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn6366 github.com/goccy/spidermonkeywasm2go/p4.Fn6366
func Fn6366(m *base.Module, l0 int32)

//go:linkname Fn6368 github.com/goccy/spidermonkeywasm2go/p3.Fn6368
func Fn6368(m *base.Module, l0 int32) int32

//go:linkname Fn6369 github.com/goccy/spidermonkeywasm2go/p4.Fn6369
func Fn6369(m *base.Module, l0 int32) int32

//go:linkname Fn6371 github.com/goccy/spidermonkeywasm2go/p4.Fn6371
func Fn6371(m *base.Module, l0 int32) int32

//go:linkname Fn6375 github.com/goccy/spidermonkeywasm2go/p4.Fn6375
func Fn6375(m *base.Module, l0 int32) int32

//go:linkname Fn6377 github.com/goccy/spidermonkeywasm2go/p4.Fn6377
func Fn6377(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6379 github.com/goccy/spidermonkeywasm2go/p4.Fn6379
func Fn6379(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6380 github.com/goccy/spidermonkeywasm2go/p3.Fn6380
func Fn6380(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6381 github.com/goccy/spidermonkeywasm2go/p3.Fn6381
func Fn6381(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6382 github.com/goccy/spidermonkeywasm2go/p4.Fn6382
func Fn6382(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6383 github.com/goccy/spidermonkeywasm2go/p4.Fn6383
func Fn6383(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6384 github.com/goccy/spidermonkeywasm2go/p3.Fn6384
func Fn6384(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6385 github.com/goccy/spidermonkeywasm2go/p3.Fn6385
func Fn6385(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6386 github.com/goccy/spidermonkeywasm2go/p3.Fn6386
func Fn6386(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6387 github.com/goccy/spidermonkeywasm2go/p3.Fn6387
func Fn6387(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6388 github.com/goccy/spidermonkeywasm2go/p3.Fn6388
func Fn6388(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6397 github.com/goccy/spidermonkeywasm2go/p4.Fn6397
func Fn6397(m *base.Module, l0 int32)

//go:linkname Fn6398 github.com/goccy/spidermonkeywasm2go/p4.Fn6398
func Fn6398(m *base.Module, l0 int32) int32

//go:linkname Fn6400 github.com/goccy/spidermonkeywasm2go/p4.Fn6400
func Fn6400(m *base.Module, l0 int32) int32

//go:linkname Fn6401 github.com/goccy/spidermonkeywasm2go/p0.Fn6401
func Fn6401(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6407 github.com/goccy/spidermonkeywasm2go/p3.Fn6407
func Fn6407(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6408 github.com/goccy/spidermonkeywasm2go/p4.Fn6408
func Fn6408(m *base.Module, l0 int32)

//go:linkname Fn6416 github.com/goccy/spidermonkeywasm2go/p4.Fn6416
func Fn6416(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6418 github.com/goccy/spidermonkeywasm2go/p4.Fn6418
func Fn6418(m *base.Module, l0 int32) int32

//go:linkname Fn6419 github.com/goccy/spidermonkeywasm2go/p4.Fn6419
func Fn6419(m *base.Module, l0 int32)

//go:linkname Fn6420 github.com/goccy/spidermonkeywasm2go/p4.Fn6420
func Fn6420(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6428 github.com/goccy/spidermonkeywasm2go/p3.Fn6428
func Fn6428(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6437 github.com/goccy/spidermonkeywasm2go/p3.Fn6437
func Fn6437(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn6443 github.com/goccy/spidermonkeywasm2go/p4.Fn6443
func Fn6443(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6453 github.com/goccy/spidermonkeywasm2go/p3.Fn6453
func Fn6453(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6454 github.com/goccy/spidermonkeywasm2go/p4.Fn6454
func Fn6454(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6504 github.com/goccy/spidermonkeywasm2go/p4.Fn6504
func Fn6504(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6506 github.com/goccy/spidermonkeywasm2go/p4.Fn6506
func Fn6506(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6508 github.com/goccy/spidermonkeywasm2go/p3.Fn6508
func Fn6508(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6560 github.com/goccy/spidermonkeywasm2go/p4.Fn6560
func Fn6560(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6561 github.com/goccy/spidermonkeywasm2go/p3.Fn6561
func Fn6561(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6562 github.com/goccy/spidermonkeywasm2go/p3.Fn6562
func Fn6562(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6563 github.com/goccy/spidermonkeywasm2go/p4.Fn6563
func Fn6563(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6569 github.com/goccy/spidermonkeywasm2go/p4.Fn6569
func Fn6569(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6576 github.com/goccy/spidermonkeywasm2go/p4.Fn6576
func Fn6576(m *base.Module, l0 int32)

//go:linkname Fn6577 github.com/goccy/spidermonkeywasm2go/p4.Fn6577
func Fn6577(m *base.Module, l0 int32)

//go:linkname Fn6578 github.com/goccy/spidermonkeywasm2go/p4.Fn6578
func Fn6578(m *base.Module, l0 int32) int32

//go:linkname Fn6581 github.com/goccy/spidermonkeywasm2go/p4.Fn6581
func Fn6581(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6582 github.com/goccy/spidermonkeywasm2go/p4.Fn6582
func Fn6582(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6589 github.com/goccy/spidermonkeywasm2go/p4.Fn6589
func Fn6589(m *base.Module, l0 int32)

//go:linkname Fn6592 github.com/goccy/spidermonkeywasm2go/p3.Fn6592
func Fn6592(m *base.Module, l0 int32) int32

//go:linkname Fn6593 github.com/goccy/spidermonkeywasm2go/p3.Fn6593
func Fn6593(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6594 github.com/goccy/spidermonkeywasm2go/p3.Fn6594
func Fn6594(m *base.Module, l0 int32) int32

//go:linkname Fn6595 github.com/goccy/spidermonkeywasm2go/p3.Fn6595
func Fn6595(m *base.Module, l0 int32) int32

//go:linkname Fn6597 github.com/goccy/spidermonkeywasm2go/p3.Fn6597
func Fn6597(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6626 github.com/goccy/spidermonkeywasm2go/p1.Fn6626
func Fn6626(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6635 github.com/goccy/spidermonkeywasm2go/p1.Fn6635
func Fn6635(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6636 github.com/goccy/spidermonkeywasm2go/p1.Fn6636
func Fn6636(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6637 github.com/goccy/spidermonkeywasm2go/p3.Fn6637
func Fn6637(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6641 github.com/goccy/spidermonkeywasm2go/p4.Fn6641
func Fn6641(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6642 github.com/goccy/spidermonkeywasm2go/p3.Fn6642
func Fn6642(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6644 github.com/goccy/spidermonkeywasm2go/p4.Fn6644
func Fn6644(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6645 github.com/goccy/spidermonkeywasm2go/p3.Fn6645
func Fn6645(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6646 github.com/goccy/spidermonkeywasm2go/p3.Fn6646
func Fn6646(m *base.Module, l0 int32) int32

//go:linkname Fn6647 github.com/goccy/spidermonkeywasm2go/p4.Fn6647
func Fn6647(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn6673 github.com/goccy/spidermonkeywasm2go/p3.Fn6673
func Fn6673(m *base.Module, l0 int32)

//go:linkname Fn6677 github.com/goccy/spidermonkeywasm2go/p4.Fn6677
func Fn6677(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6679 github.com/goccy/spidermonkeywasm2go/p3.Fn6679
func Fn6679(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6684 github.com/goccy/spidermonkeywasm2go/p4.Fn6684
func Fn6684(m *base.Module, l0 int32)

//go:linkname Fn6685 github.com/goccy/spidermonkeywasm2go/p4.Fn6685
func Fn6685(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6689 github.com/goccy/spidermonkeywasm2go/p3.Fn6689
func Fn6689(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6690 github.com/goccy/spidermonkeywasm2go/p3.Fn6690
func Fn6690(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6695 github.com/goccy/spidermonkeywasm2go/p1.Fn6695
func Fn6695(m *base.Module, l0 float64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6700 github.com/goccy/spidermonkeywasm2go/p3.Fn6700
func Fn6700(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6701 github.com/goccy/spidermonkeywasm2go/p3.Fn6701
func Fn6701(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6705 github.com/goccy/spidermonkeywasm2go/p3.Fn6705
func Fn6705(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6709 github.com/goccy/spidermonkeywasm2go/p3.Fn6709
func Fn6709(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32) int32

//go:linkname Fn6710 github.com/goccy/spidermonkeywasm2go/p3.Fn6710
func Fn6710(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32) int32

//go:linkname Fn6711 github.com/goccy/spidermonkeywasm2go/p3.Fn6711
func Fn6711(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn6712 github.com/goccy/spidermonkeywasm2go/p3.Fn6712
func Fn6712(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6720 github.com/goccy/spidermonkeywasm2go/p4.Fn6720
func Fn6720(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32
