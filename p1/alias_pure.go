//go:build (!amd64 && !arm64) || purego

package p1

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

//go:linkname Fn28 github.com/goccy/spidermonkeywasm2go/p4.Fn28
func Fn28(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn29 github.com/goccy/spidermonkeywasm2go/p4.Fn29
func Fn29(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn53 github.com/goccy/spidermonkeywasm2go/p4.Fn53
func Fn53(m *base.Module, l0 int32)

//go:linkname Fn56 github.com/goccy/spidermonkeywasm2go/p2.Fn56
func Fn56(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn77 github.com/goccy/spidermonkeywasm2go/p4.Fn77
func Fn77(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn81 github.com/goccy/spidermonkeywasm2go/p2.Fn81
func Fn81(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn82 github.com/goccy/spidermonkeywasm2go/p4.Fn82
func Fn82(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn83 github.com/goccy/spidermonkeywasm2go/p4.Fn83
func Fn83(m *base.Module)

//go:linkname Fn113 github.com/goccy/spidermonkeywasm2go/p4.Fn113
func Fn113(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn116 github.com/goccy/spidermonkeywasm2go/p4.Fn116
func Fn116(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn119 github.com/goccy/spidermonkeywasm2go/p4.Fn119
func Fn119(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn122 github.com/goccy/spidermonkeywasm2go/p4.Fn122
func Fn122(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn125 github.com/goccy/spidermonkeywasm2go/p4.Fn125
func Fn125(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn128 github.com/goccy/spidermonkeywasm2go/p4.Fn128
func Fn128(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn131 github.com/goccy/spidermonkeywasm2go/p4.Fn131
func Fn131(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn134 github.com/goccy/spidermonkeywasm2go/p4.Fn134
func Fn134(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn159 github.com/goccy/spidermonkeywasm2go/p4.Fn159
func Fn159(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn214 github.com/goccy/spidermonkeywasm2go/p4.Fn214
func Fn214(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn215 github.com/goccy/spidermonkeywasm2go/p4.Fn215
func Fn215(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn325 github.com/goccy/spidermonkeywasm2go/p4.Fn325
func Fn325(m *base.Module, l0 int32)

//go:linkname Fn326 github.com/goccy/spidermonkeywasm2go/p4.Fn326
func Fn326(m *base.Module) int32

//go:linkname Fn329 github.com/goccy/spidermonkeywasm2go/p4.Fn329
func Fn329(m *base.Module, l0 int32)

//go:linkname Fn330 github.com/goccy/spidermonkeywasm2go/p4.Fn330
func Fn330(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn331 github.com/goccy/spidermonkeywasm2go/p4.Fn331
func Fn331(m *base.Module, l0 int32)

//go:linkname Fn332 github.com/goccy/spidermonkeywasm2go/p4.Fn332
func Fn332(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn432 github.com/goccy/spidermonkeywasm2go/p4.Fn432
func Fn432(m *base.Module, l0 int32) int32

//go:linkname Fn434 github.com/goccy/spidermonkeywasm2go/p4.Fn434
func Fn434(m *base.Module, l0 int32) int32

//go:linkname Fn442 github.com/goccy/spidermonkeywasm2go/p4.Fn442
func Fn442(m *base.Module, l0 int32) int32

//go:linkname Fn443 github.com/goccy/spidermonkeywasm2go/p4.Fn443
func Fn443(m *base.Module, l0 int32) int32

//go:linkname Fn444 github.com/goccy/spidermonkeywasm2go/p4.Fn444
func Fn444(m *base.Module, l0 int32)

//go:linkname Fn445 github.com/goccy/spidermonkeywasm2go/p4.Fn445
func Fn445(m *base.Module, l0 int32)

//go:linkname Fn453 github.com/goccy/spidermonkeywasm2go/p4.Fn453
func Fn453(m *base.Module, l0 int32)

//go:linkname Fn461 github.com/goccy/spidermonkeywasm2go/p3.Fn461
func Fn461(m *base.Module) int32

//go:linkname Fn479 github.com/goccy/spidermonkeywasm2go/p4.Fn479
func Fn479(m *base.Module, l0 int32) int32

//go:linkname Fn487 github.com/goccy/spidermonkeywasm2go/p3.Fn487
func Fn487(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn488 github.com/goccy/spidermonkeywasm2go/p4.Fn488
func Fn488(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn492 github.com/goccy/spidermonkeywasm2go/p4.Fn492
func Fn492(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn493 github.com/goccy/spidermonkeywasm2go/p4.Fn493
func Fn493(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn496 github.com/goccy/spidermonkeywasm2go/p4.Fn496
func Fn496(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn497 github.com/goccy/spidermonkeywasm2go/p4.Fn497
func Fn497(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn500 github.com/goccy/spidermonkeywasm2go/p4.Fn500
func Fn500(m *base.Module, l0 int32)

//go:linkname Fn501 github.com/goccy/spidermonkeywasm2go/p4.Fn501
func Fn501(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn511 github.com/goccy/spidermonkeywasm2go/p4.Fn511
func Fn511(m *base.Module, l0 int32) int32

//go:linkname Fn515 github.com/goccy/spidermonkeywasm2go/p4.Fn515
func Fn515(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn516 github.com/goccy/spidermonkeywasm2go/p4.Fn516
func Fn516(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn517 github.com/goccy/spidermonkeywasm2go/p4.Fn517
func Fn517(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn520 github.com/goccy/spidermonkeywasm2go/p4.Fn520
func Fn520(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn523 github.com/goccy/spidermonkeywasm2go/p4.Fn523
func Fn523(m *base.Module, l0 int32) int32

//go:linkname Fn524 github.com/goccy/spidermonkeywasm2go/p4.Fn524
func Fn524(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn526 github.com/goccy/spidermonkeywasm2go/p4.Fn526
func Fn526(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn527 github.com/goccy/spidermonkeywasm2go/p3.Fn527
func Fn527(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn528 github.com/goccy/spidermonkeywasm2go/p4.Fn528
func Fn528(m *base.Module, l0 int32) int32

//go:linkname Fn533 github.com/goccy/spidermonkeywasm2go/p4.Fn533
func Fn533(m *base.Module, l0 float64, l1 int32) float64

//go:linkname Fn536 github.com/goccy/spidermonkeywasm2go/p4.Fn536
func Fn536(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn537 github.com/goccy/spidermonkeywasm2go/p4.Fn537
func Fn537(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn543 github.com/goccy/spidermonkeywasm2go/p2.Fn543
func Fn543(m *base.Module, l0 int32) int32

//go:linkname Fn544 github.com/goccy/spidermonkeywasm2go/p3.Fn544
func Fn544(m *base.Module, l0 int32)

//go:linkname Fn545 github.com/goccy/spidermonkeywasm2go/p4.Fn545
func Fn545(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn546 github.com/goccy/spidermonkeywasm2go/p3.Fn546
func Fn546(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn550 github.com/goccy/spidermonkeywasm2go/p4.Fn550
func Fn550(m *base.Module, l0 int32)

//go:linkname Fn553 github.com/goccy/spidermonkeywasm2go/p3.Fn553
func Fn553(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn554 github.com/goccy/spidermonkeywasm2go/p3.Fn554
func Fn554(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn558 github.com/goccy/spidermonkeywasm2go/p0.Fn558
func Fn558(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn566 github.com/goccy/spidermonkeywasm2go/p0.Fn566
func Fn566(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn567 github.com/goccy/spidermonkeywasm2go/p0.Fn567
func Fn567(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn578 github.com/goccy/spidermonkeywasm2go/p4.Fn578
func Fn578(m *base.Module, l0 int32) int32

//go:linkname Fn580 github.com/goccy/spidermonkeywasm2go/p4.Fn580
func Fn580(m *base.Module, l0 int32) int32

//go:linkname Fn582 github.com/goccy/spidermonkeywasm2go/p3.Fn582
func Fn582(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn586 github.com/goccy/spidermonkeywasm2go/p4.Fn586
func Fn586(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn588 github.com/goccy/spidermonkeywasm2go/p3.Fn588
func Fn588(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn598 github.com/goccy/spidermonkeywasm2go/p4.Fn598
func Fn598(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn604 github.com/goccy/spidermonkeywasm2go/p4.Fn604
func Fn604(m *base.Module, l0 int32) int32

//go:linkname Fn607 github.com/goccy/spidermonkeywasm2go/p0.Fn607
func Fn607(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn608 github.com/goccy/spidermonkeywasm2go/p0.Fn608
func Fn608(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn609 github.com/goccy/spidermonkeywasm2go/p0.Fn609
func Fn609(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn610 github.com/goccy/spidermonkeywasm2go/p4.Fn610
func Fn610(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn611 github.com/goccy/spidermonkeywasm2go/p3.Fn611
func Fn611(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn612 github.com/goccy/spidermonkeywasm2go/p4.Fn612
func Fn612(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn622 github.com/goccy/spidermonkeywasm2go/p0.Fn622
func Fn622(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn630 github.com/goccy/spidermonkeywasm2go/p4.Fn630
func Fn630(m *base.Module, l0 int32) int32

//go:linkname Fn633 github.com/goccy/spidermonkeywasm2go/p4.Fn633
func Fn633(m *base.Module, l0 int32) int32

//go:linkname Fn634 github.com/goccy/spidermonkeywasm2go/p4.Fn634
func Fn634(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn646 github.com/goccy/spidermonkeywasm2go/p4.Fn646
func Fn646(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn647 github.com/goccy/spidermonkeywasm2go/p4.Fn647
func Fn647(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn649 github.com/goccy/spidermonkeywasm2go/p3.Fn649
func Fn649(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn650 github.com/goccy/spidermonkeywasm2go/p4.Fn650
func Fn650(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn651 github.com/goccy/spidermonkeywasm2go/p3.Fn651
func Fn651(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn652 github.com/goccy/spidermonkeywasm2go/p4.Fn652
func Fn652(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn658 github.com/goccy/spidermonkeywasm2go/p4.Fn658
func Fn658(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn659 github.com/goccy/spidermonkeywasm2go/p0.Fn659
func Fn659(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn660 github.com/goccy/spidermonkeywasm2go/p0.Fn660
func Fn660(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn663 github.com/goccy/spidermonkeywasm2go/p0.Fn663
func Fn663(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn667 github.com/goccy/spidermonkeywasm2go/p0.Fn667
func Fn667(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn678 github.com/goccy/spidermonkeywasm2go/p0.Fn678
func Fn678(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn679 github.com/goccy/spidermonkeywasm2go/p0.Fn679
func Fn679(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn688 github.com/goccy/spidermonkeywasm2go/p0.Fn688
func Fn688(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn699 github.com/goccy/spidermonkeywasm2go/p0.Fn699
func Fn699(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn700 github.com/goccy/spidermonkeywasm2go/p0.Fn700
func Fn700(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn711 github.com/goccy/spidermonkeywasm2go/p3.Fn711
func Fn711(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn714 github.com/goccy/spidermonkeywasm2go/p0.Fn714
func Fn714(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn717 github.com/goccy/spidermonkeywasm2go/p0.Fn717
func Fn717(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn720 github.com/goccy/spidermonkeywasm2go/p4.Fn720
func Fn720(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn721 github.com/goccy/spidermonkeywasm2go/p0.Fn721
func Fn721(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn722 github.com/goccy/spidermonkeywasm2go/p3.Fn722
func Fn722(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn723 github.com/goccy/spidermonkeywasm2go/p4.Fn723
func Fn723(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn734 github.com/goccy/spidermonkeywasm2go/p4.Fn734
func Fn734(m *base.Module, l0 int32, l1 int32) int64

//go:linkname Fn740 github.com/goccy/spidermonkeywasm2go/p4.Fn740
func Fn740(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn741 github.com/goccy/spidermonkeywasm2go/p4.Fn741
func Fn741(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn742 github.com/goccy/spidermonkeywasm2go/p4.Fn742
func Fn742(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn763 github.com/goccy/spidermonkeywasm2go/p4.Fn763
func Fn763(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn764 github.com/goccy/spidermonkeywasm2go/p3.Fn764
func Fn764(m *base.Module, l0 int32)

//go:linkname Fn795 github.com/goccy/spidermonkeywasm2go/p3.Fn795
func Fn795(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn796 github.com/goccy/spidermonkeywasm2go/p0.Fn796
func Fn796(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn816 github.com/goccy/spidermonkeywasm2go/p2.Fn816
func Fn816(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn817 github.com/goccy/spidermonkeywasm2go/p4.Fn817
func Fn817(m *base.Module, l0 int32)

//go:linkname Fn818 github.com/goccy/spidermonkeywasm2go/p4.Fn818
func Fn818(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn829 github.com/goccy/spidermonkeywasm2go/p4.Fn829
func Fn829(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn830 github.com/goccy/spidermonkeywasm2go/p4.Fn830
func Fn830(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn852 github.com/goccy/spidermonkeywasm2go/p0.Fn852
func Fn852(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn853 github.com/goccy/spidermonkeywasm2go/p0.Fn853
func Fn853(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn854 github.com/goccy/spidermonkeywasm2go/p0.Fn854
func Fn854(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn867 github.com/goccy/spidermonkeywasm2go/p0.Fn867
func Fn867(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn873 github.com/goccy/spidermonkeywasm2go/p0.Fn873
func Fn873(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn896 github.com/goccy/spidermonkeywasm2go/p0.Fn896
func Fn896(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn900 github.com/goccy/spidermonkeywasm2go/p4.Fn900
func Fn900(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn902 github.com/goccy/spidermonkeywasm2go/p3.Fn902
func Fn902(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn903 github.com/goccy/spidermonkeywasm2go/p3.Fn903
func Fn903(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn904 github.com/goccy/spidermonkeywasm2go/p3.Fn904
func Fn904(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn905 github.com/goccy/spidermonkeywasm2go/p3.Fn905
func Fn905(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn906 github.com/goccy/spidermonkeywasm2go/p3.Fn906
func Fn906(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn907 github.com/goccy/spidermonkeywasm2go/p4.Fn907
func Fn907(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn908 github.com/goccy/spidermonkeywasm2go/p4.Fn908
func Fn908(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn909 github.com/goccy/spidermonkeywasm2go/p4.Fn909
func Fn909(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn910 github.com/goccy/spidermonkeywasm2go/p4.Fn910
func Fn910(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn911 github.com/goccy/spidermonkeywasm2go/p3.Fn911
func Fn911(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn912 github.com/goccy/spidermonkeywasm2go/p3.Fn912
func Fn912(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn913 github.com/goccy/spidermonkeywasm2go/p3.Fn913
func Fn913(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn914 github.com/goccy/spidermonkeywasm2go/p3.Fn914
func Fn914(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn915 github.com/goccy/spidermonkeywasm2go/p3.Fn915
func Fn915(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn935 github.com/goccy/spidermonkeywasm2go/p0.Fn935
func Fn935(m *base.Module, l0 int32, l1 int32) int32

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

//go:linkname Fn1021 github.com/goccy/spidermonkeywasm2go/p3.Fn1021
func Fn1021(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1024 github.com/goccy/spidermonkeywasm2go/p0.Fn1024
func Fn1024(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1029 github.com/goccy/spidermonkeywasm2go/p0.Fn1029
func Fn1029(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1032 github.com/goccy/spidermonkeywasm2go/p4.Fn1032
func Fn1032(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn1038 github.com/goccy/spidermonkeywasm2go/p4.Fn1038
func Fn1038(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1039 github.com/goccy/spidermonkeywasm2go/p2.Fn1039
func Fn1039(m *base.Module, l0 int32) int32

//go:linkname Fn1040 github.com/goccy/spidermonkeywasm2go/p3.Fn1040
func Fn1040(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1041 github.com/goccy/spidermonkeywasm2go/p3.Fn1041
func Fn1041(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn1044 github.com/goccy/spidermonkeywasm2go/p0.Fn1044
func Fn1044(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1045 github.com/goccy/spidermonkeywasm2go/p4.Fn1045
func Fn1045(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1046 github.com/goccy/spidermonkeywasm2go/p4.Fn1046
func Fn1046(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1049 github.com/goccy/spidermonkeywasm2go/p4.Fn1049
func Fn1049(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1055 github.com/goccy/spidermonkeywasm2go/p0.Fn1055
func Fn1055(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1062 github.com/goccy/spidermonkeywasm2go/p0.Fn1062
func Fn1062(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1064 github.com/goccy/spidermonkeywasm2go/p4.Fn1064
func Fn1064(m *base.Module, l0 int32) int32

//go:linkname Fn1107 github.com/goccy/spidermonkeywasm2go/p4.Fn1107
func Fn1107(m *base.Module, l0 int32) int32

//go:linkname Fn1199 github.com/goccy/spidermonkeywasm2go/p3.Fn1199
func Fn1199(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1256 github.com/goccy/spidermonkeywasm2go/p4.Fn1256
func Fn1256(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1257 github.com/goccy/spidermonkeywasm2go/p0.Fn1257
func Fn1257(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1284 github.com/goccy/spidermonkeywasm2go/p4.Fn1284
func Fn1284(m *base.Module, l0 int32) int32

//go:linkname Fn1285 github.com/goccy/spidermonkeywasm2go/p4.Fn1285
func Fn1285(m *base.Module, l0 int32) int32

//go:linkname Fn1286 github.com/goccy/spidermonkeywasm2go/p4.Fn1286
func Fn1286(m *base.Module, l0 int32) int32

//go:linkname Fn1325 github.com/goccy/spidermonkeywasm2go/p3.Fn1325
func Fn1325(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1326 github.com/goccy/spidermonkeywasm2go/p4.Fn1326
func Fn1326(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1341 github.com/goccy/spidermonkeywasm2go/p0.Fn1341
func Fn1341(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1342 github.com/goccy/spidermonkeywasm2go/p3.Fn1342
func Fn1342(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1343 github.com/goccy/spidermonkeywasm2go/p3.Fn1343
func Fn1343(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1344 github.com/goccy/spidermonkeywasm2go/p3.Fn1344
func Fn1344(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1345 github.com/goccy/spidermonkeywasm2go/p2.Fn1345
func Fn1345(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1348 github.com/goccy/spidermonkeywasm2go/p0.Fn1348
func Fn1348(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1349 github.com/goccy/spidermonkeywasm2go/p2.Fn1349
func Fn1349(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1355 github.com/goccy/spidermonkeywasm2go/p4.Fn1355
func Fn1355(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1392 github.com/goccy/spidermonkeywasm2go/p2.Fn1392
func Fn1392(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1398 github.com/goccy/spidermonkeywasm2go/p4.Fn1398
func Fn1398(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1399 github.com/goccy/spidermonkeywasm2go/p4.Fn1399
func Fn1399(m *base.Module, l0 int32) int32

//go:linkname Fn1414 github.com/goccy/spidermonkeywasm2go/p2.Fn1414
func Fn1414(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1428 github.com/goccy/spidermonkeywasm2go/p4.Fn1428
func Fn1428(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1429 github.com/goccy/spidermonkeywasm2go/p4.Fn1429
func Fn1429(m *base.Module, l0 int32) int32

//go:linkname Fn1444 github.com/goccy/spidermonkeywasm2go/p0.Fn1444
func Fn1444(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1448 github.com/goccy/spidermonkeywasm2go/p4.Fn1448
func Fn1448(m *base.Module, l0 int32) int32

//go:linkname Fn1450 github.com/goccy/spidermonkeywasm2go/p3.Fn1450
func Fn1450(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1454 github.com/goccy/spidermonkeywasm2go/p0.Fn1454
func Fn1454(m *base.Module, l0 int32) int32

//go:linkname Fn1464 github.com/goccy/spidermonkeywasm2go/p4.Fn1464
func Fn1464(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1476 github.com/goccy/spidermonkeywasm2go/p0.Fn1476
func Fn1476(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1478 github.com/goccy/spidermonkeywasm2go/p4.Fn1478
func Fn1478(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1479 github.com/goccy/spidermonkeywasm2go/p4.Fn1479
func Fn1479(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1503 github.com/goccy/spidermonkeywasm2go/p4.Fn1503
func Fn1503(m *base.Module, l0 int32) int32

//go:linkname Fn1507 github.com/goccy/spidermonkeywasm2go/p3.Fn1507
func Fn1507(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn1508 github.com/goccy/spidermonkeywasm2go/p4.Fn1508
func Fn1508(m *base.Module, l0 int32)

//go:linkname Fn1516 github.com/goccy/spidermonkeywasm2go/p3.Fn1516
func Fn1516(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn1532 github.com/goccy/spidermonkeywasm2go/p3.Fn1532
func Fn1532(m *base.Module, l0 float64, l1 float64) float64

//go:linkname Fn1590 github.com/goccy/spidermonkeywasm2go/p4.Fn1590
func Fn1590(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1610 github.com/goccy/spidermonkeywasm2go/p2.Fn1610
func Fn1610(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1616 github.com/goccy/spidermonkeywasm2go/p4.Fn1616
func Fn1616(m *base.Module, l0 int32)

//go:linkname Fn1623 github.com/goccy/spidermonkeywasm2go/p4.Fn1623
func Fn1623(m *base.Module, l0 int32)

//go:linkname Fn1634 github.com/goccy/spidermonkeywasm2go/p0.Fn1634
func Fn1634(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1636 github.com/goccy/spidermonkeywasm2go/p4.Fn1636
func Fn1636(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1645 github.com/goccy/spidermonkeywasm2go/p0.Fn1645
func Fn1645(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32

//go:linkname Fn1686 github.com/goccy/spidermonkeywasm2go/p3.Fn1686
func Fn1686(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1692 github.com/goccy/spidermonkeywasm2go/p3.Fn1692
func Fn1692(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1694 github.com/goccy/spidermonkeywasm2go/p4.Fn1694
func Fn1694(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1695 github.com/goccy/spidermonkeywasm2go/p3.Fn1695
func Fn1695(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1696 github.com/goccy/spidermonkeywasm2go/p2.Fn1696
func Fn1696(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1697 github.com/goccy/spidermonkeywasm2go/p3.Fn1697
func Fn1697(m *base.Module, l0 int32)

//go:linkname Fn1699 github.com/goccy/spidermonkeywasm2go/p4.Fn1699
func Fn1699(m *base.Module, l0 int32)

//go:linkname Fn1700 github.com/goccy/spidermonkeywasm2go/p4.Fn1700
func Fn1700(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1705 github.com/goccy/spidermonkeywasm2go/p4.Fn1705
func Fn1705(m *base.Module, l0 int32) int32

//go:linkname Fn1707 github.com/goccy/spidermonkeywasm2go/p4.Fn1707
func Fn1707(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1708 github.com/goccy/spidermonkeywasm2go/p4.Fn1708
func Fn1708(m *base.Module, l0 int32) int32

//go:linkname Fn1712 github.com/goccy/spidermonkeywasm2go/p4.Fn1712
func Fn1712(m *base.Module, l0 int32)

//go:linkname Fn1715 github.com/goccy/spidermonkeywasm2go/p3.Fn1715
func Fn1715(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1720 github.com/goccy/spidermonkeywasm2go/p0.Fn1720
func Fn1720(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1723 github.com/goccy/spidermonkeywasm2go/p3.Fn1723
func Fn1723(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1738 github.com/goccy/spidermonkeywasm2go/p0.Fn1738
func Fn1738(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1741 github.com/goccy/spidermonkeywasm2go/p4.Fn1741
func Fn1741(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1742 github.com/goccy/spidermonkeywasm2go/p3.Fn1742
func Fn1742(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1746 github.com/goccy/spidermonkeywasm2go/p2.Fn1746
func Fn1746(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1748 github.com/goccy/spidermonkeywasm2go/p3.Fn1748
func Fn1748(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1758 github.com/goccy/spidermonkeywasm2go/p4.Fn1758
func Fn1758(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1770 github.com/goccy/spidermonkeywasm2go/p4.Fn1770
func Fn1770(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1775 github.com/goccy/spidermonkeywasm2go/p3.Fn1775
func Fn1775(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn1776 github.com/goccy/spidermonkeywasm2go/p4.Fn1776
func Fn1776(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1778 github.com/goccy/spidermonkeywasm2go/p4.Fn1778
func Fn1778(m *base.Module, l0 int32) int32

//go:linkname Fn1781 github.com/goccy/spidermonkeywasm2go/p0.Fn1781
func Fn1781(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1787 github.com/goccy/spidermonkeywasm2go/p0.Fn1787
func Fn1787(m *base.Module, l0 int32) int32

//go:linkname Fn1790 github.com/goccy/spidermonkeywasm2go/p0.Fn1790
func Fn1790(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1794 github.com/goccy/spidermonkeywasm2go/p3.Fn1794
func Fn1794(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1796 github.com/goccy/spidermonkeywasm2go/p0.Fn1796
func Fn1796(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1798 github.com/goccy/spidermonkeywasm2go/p0.Fn1798
func Fn1798(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1805 github.com/goccy/spidermonkeywasm2go/p0.Fn1805
func Fn1805(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1806 github.com/goccy/spidermonkeywasm2go/p0.Fn1806
func Fn1806(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1811 github.com/goccy/spidermonkeywasm2go/p2.Fn1811
func Fn1811(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1813 github.com/goccy/spidermonkeywasm2go/p3.Fn1813
func Fn1813(m *base.Module, l0 int32)

//go:linkname Fn1814 github.com/goccy/spidermonkeywasm2go/p3.Fn1814
func Fn1814(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1820 github.com/goccy/spidermonkeywasm2go/p0.Fn1820
func Fn1820(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1821 github.com/goccy/spidermonkeywasm2go/p0.Fn1821
func Fn1821(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1824 github.com/goccy/spidermonkeywasm2go/p0.Fn1824
func Fn1824(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1826 github.com/goccy/spidermonkeywasm2go/p0.Fn1826
func Fn1826(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1827 github.com/goccy/spidermonkeywasm2go/p0.Fn1827
func Fn1827(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1850 github.com/goccy/spidermonkeywasm2go/p4.Fn1850
func Fn1850(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1869 github.com/goccy/spidermonkeywasm2go/p0.Fn1869
func Fn1869(m *base.Module, l0 int32) int32

//go:linkname Fn1872 github.com/goccy/spidermonkeywasm2go/p0.Fn1872
func Fn1872(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1874 github.com/goccy/spidermonkeywasm2go/p0.Fn1874
func Fn1874(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1888 github.com/goccy/spidermonkeywasm2go/p4.Fn1888
func Fn1888(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1904 github.com/goccy/spidermonkeywasm2go/p4.Fn1904
func Fn1904(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1914 github.com/goccy/spidermonkeywasm2go/p3.Fn1914
func Fn1914(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1915 github.com/goccy/spidermonkeywasm2go/p3.Fn1915
func Fn1915(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1916 github.com/goccy/spidermonkeywasm2go/p4.Fn1916
func Fn1916(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1918 github.com/goccy/spidermonkeywasm2go/p0.Fn1918
func Fn1918(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn1920 github.com/goccy/spidermonkeywasm2go/p0.Fn1920
func Fn1920(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1923 github.com/goccy/spidermonkeywasm2go/p4.Fn1923
func Fn1923(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1932 github.com/goccy/spidermonkeywasm2go/p0.Fn1932
func Fn1932(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1934 github.com/goccy/spidermonkeywasm2go/p4.Fn1934
func Fn1934(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1937 github.com/goccy/spidermonkeywasm2go/p0.Fn1937
func Fn1937(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1938 github.com/goccy/spidermonkeywasm2go/p4.Fn1938
func Fn1938(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1939 github.com/goccy/spidermonkeywasm2go/p0.Fn1939
func Fn1939(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn1940 github.com/goccy/spidermonkeywasm2go/p0.Fn1940
func Fn1940(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn1943 github.com/goccy/spidermonkeywasm2go/p0.Fn1943
func Fn1943(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1946 github.com/goccy/spidermonkeywasm2go/p0.Fn1946
func Fn1946(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn1958 github.com/goccy/spidermonkeywasm2go/p4.Fn1958
func Fn1958(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn1961 github.com/goccy/spidermonkeywasm2go/p3.Fn1961
func Fn1961(m *base.Module, l0 int32)

//go:linkname Fn1962 github.com/goccy/spidermonkeywasm2go/p4.Fn1962
func Fn1962(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2013 github.com/goccy/spidermonkeywasm2go/p4.Fn2013
func Fn2013(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2014 github.com/goccy/spidermonkeywasm2go/p4.Fn2014
func Fn2014(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2018 github.com/goccy/spidermonkeywasm2go/p2.Fn2018
func Fn2018(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2038 github.com/goccy/spidermonkeywasm2go/p4.Fn2038
func Fn2038(m *base.Module, l0 int32)

//go:linkname Fn2040 github.com/goccy/spidermonkeywasm2go/p0.Fn2040
func Fn2040(m *base.Module, l0 int32)

//go:linkname Fn2041 github.com/goccy/spidermonkeywasm2go/p0.Fn2041
func Fn2041(m *base.Module, l0 int32)

//go:linkname Fn2042 github.com/goccy/spidermonkeywasm2go/p0.Fn2042
func Fn2042(m *base.Module, l0 int32)

//go:linkname Fn2043 github.com/goccy/spidermonkeywasm2go/p3.Fn2043
func Fn2043(m *base.Module, l0 int32, l1 int32, l2 int32)

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

//go:linkname Fn2060 github.com/goccy/spidermonkeywasm2go/p0.Fn2060
func Fn2060(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2063 github.com/goccy/spidermonkeywasm2go/p4.Fn2063
func Fn2063(m *base.Module, l0 int32)

//go:linkname Fn2065 github.com/goccy/spidermonkeywasm2go/p4.Fn2065
func Fn2065(m *base.Module, l0 int32) int32

//go:linkname Fn2066 github.com/goccy/spidermonkeywasm2go/p0.Fn2066
func Fn2066(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2073 github.com/goccy/spidermonkeywasm2go/p0.Fn2073
func Fn2073(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2078 github.com/goccy/spidermonkeywasm2go/p0.Fn2078
func Fn2078(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2080 github.com/goccy/spidermonkeywasm2go/p0.Fn2080
func Fn2080(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2081 github.com/goccy/spidermonkeywasm2go/p3.Fn2081
func Fn2081(m *base.Module, l0 int32)

//go:linkname Fn2085 github.com/goccy/spidermonkeywasm2go/p4.Fn2085
func Fn2085(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn2095 github.com/goccy/spidermonkeywasm2go/p0.Fn2095
func Fn2095(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32

//go:linkname Fn2107 github.com/goccy/spidermonkeywasm2go/p4.Fn2107
func Fn2107(m *base.Module, l0 int32)

//go:linkname Fn2108 github.com/goccy/spidermonkeywasm2go/p0.Fn2108
func Fn2108(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2109 github.com/goccy/spidermonkeywasm2go/p4.Fn2109
func Fn2109(m *base.Module, l0 int32)

//go:linkname Fn2110 github.com/goccy/spidermonkeywasm2go/p0.Fn2110
func Fn2110(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2111 github.com/goccy/spidermonkeywasm2go/p0.Fn2111
func Fn2111(m *base.Module, l0 int32) int32

//go:linkname Fn2112 github.com/goccy/spidermonkeywasm2go/p0.Fn2112
func Fn2112(m *base.Module, l0 int32) int32

//go:linkname Fn2115 github.com/goccy/spidermonkeywasm2go/p0.Fn2115
func Fn2115(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2116 github.com/goccy/spidermonkeywasm2go/p0.Fn2116
func Fn2116(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2117 github.com/goccy/spidermonkeywasm2go/p0.Fn2117
func Fn2117(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2120 github.com/goccy/spidermonkeywasm2go/p4.Fn2120
func Fn2120(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2121 github.com/goccy/spidermonkeywasm2go/p2.Fn2121
func Fn2121(m *base.Module, l0 int32) int32

//go:linkname Fn2123 github.com/goccy/spidermonkeywasm2go/p4.Fn2123
func Fn2123(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2124 github.com/goccy/spidermonkeywasm2go/p2.Fn2124
func Fn2124(m *base.Module, l0 int32) int32

//go:linkname Fn2132 github.com/goccy/spidermonkeywasm2go/p0.Fn2132
func Fn2132(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2135 github.com/goccy/spidermonkeywasm2go/p0.Fn2135
func Fn2135(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2136 github.com/goccy/spidermonkeywasm2go/p0.Fn2136
func Fn2136(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2137 github.com/goccy/spidermonkeywasm2go/p0.Fn2137
func Fn2137(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2145 github.com/goccy/spidermonkeywasm2go/p3.Fn2145
func Fn2145(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2146 github.com/goccy/spidermonkeywasm2go/p3.Fn2146
func Fn2146(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2155 github.com/goccy/spidermonkeywasm2go/p0.Fn2155
func Fn2155(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2157 github.com/goccy/spidermonkeywasm2go/p0.Fn2157
func Fn2157(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2165 github.com/goccy/spidermonkeywasm2go/p3.Fn2165
func Fn2165(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2166 github.com/goccy/spidermonkeywasm2go/p3.Fn2166
func Fn2166(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2167 github.com/goccy/spidermonkeywasm2go/p4.Fn2167
func Fn2167(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2168 github.com/goccy/spidermonkeywasm2go/p0.Fn2168
func Fn2168(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2169 github.com/goccy/spidermonkeywasm2go/p4.Fn2169
func Fn2169(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2170 github.com/goccy/spidermonkeywasm2go/p3.Fn2170
func Fn2170(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2174 github.com/goccy/spidermonkeywasm2go/p4.Fn2174
func Fn2174(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2175 github.com/goccy/spidermonkeywasm2go/p4.Fn2175
func Fn2175(m *base.Module, l0 int32)

//go:linkname Fn2199 github.com/goccy/spidermonkeywasm2go/p3.Fn2199
func Fn2199(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2200 github.com/goccy/spidermonkeywasm2go/p4.Fn2200
func Fn2200(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2201 github.com/goccy/spidermonkeywasm2go/p3.Fn2201
func Fn2201(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2262 github.com/goccy/spidermonkeywasm2go/p0.Fn2262
func Fn2262(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2264 github.com/goccy/spidermonkeywasm2go/p4.Fn2264
func Fn2264(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2265 github.com/goccy/spidermonkeywasm2go/p4.Fn2265
func Fn2265(m *base.Module, l0 int32, l1 float64) int32

//go:linkname Fn2266 github.com/goccy/spidermonkeywasm2go/p4.Fn2266
func Fn2266(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2267 github.com/goccy/spidermonkeywasm2go/p4.Fn2267
func Fn2267(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2268 github.com/goccy/spidermonkeywasm2go/p4.Fn2268
func Fn2268(m *base.Module, l0 int32, l1 float64) int32

//go:linkname Fn2269 github.com/goccy/spidermonkeywasm2go/p4.Fn2269
func Fn2269(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2289 github.com/goccy/spidermonkeywasm2go/p0.Fn2289
func Fn2289(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2293 github.com/goccy/spidermonkeywasm2go/p3.Fn2293
func Fn2293(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2294 github.com/goccy/spidermonkeywasm2go/p3.Fn2294
func Fn2294(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2297 github.com/goccy/spidermonkeywasm2go/p3.Fn2297
func Fn2297(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2308 github.com/goccy/spidermonkeywasm2go/p0.Fn2308
func Fn2308(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2314 github.com/goccy/spidermonkeywasm2go/p0.Fn2314
func Fn2314(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2318 github.com/goccy/spidermonkeywasm2go/p3.Fn2318
func Fn2318(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2341 github.com/goccy/spidermonkeywasm2go/p4.Fn2341
func Fn2341(m *base.Module, l0 int32) int32

//go:linkname Fn2347 github.com/goccy/spidermonkeywasm2go/p3.Fn2347
func Fn2347(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2351 github.com/goccy/spidermonkeywasm2go/p0.Fn2351
func Fn2351(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2353 github.com/goccy/spidermonkeywasm2go/p0.Fn2353
func Fn2353(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2354 github.com/goccy/spidermonkeywasm2go/p0.Fn2354
func Fn2354(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2357 github.com/goccy/spidermonkeywasm2go/p0.Fn2357
func Fn2357(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2361 github.com/goccy/spidermonkeywasm2go/p4.Fn2361
func Fn2361(m *base.Module, l0 int32)

//go:linkname Fn2364 github.com/goccy/spidermonkeywasm2go/p4.Fn2364
func Fn2364(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2369 github.com/goccy/spidermonkeywasm2go/p0.Fn2369
func Fn2369(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2372 github.com/goccy/spidermonkeywasm2go/p4.Fn2372
func Fn2372(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2401 github.com/goccy/spidermonkeywasm2go/p4.Fn2401
func Fn2401(m *base.Module, l0 int32)

//go:linkname Fn2404 github.com/goccy/spidermonkeywasm2go/p0.Fn2404
func Fn2404(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2407 github.com/goccy/spidermonkeywasm2go/p4.Fn2407
func Fn2407(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2408 github.com/goccy/spidermonkeywasm2go/p4.Fn2408
func Fn2408(m *base.Module, l0 int32)

//go:linkname Fn2409 github.com/goccy/spidermonkeywasm2go/p2.Fn2409
func Fn2409(m *base.Module, l0 int32) int32

//go:linkname Fn2410 github.com/goccy/spidermonkeywasm2go/p3.Fn2410
func Fn2410(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2456 github.com/goccy/spidermonkeywasm2go/p0.Fn2456
func Fn2456(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2457 github.com/goccy/spidermonkeywasm2go/p0.Fn2457
func Fn2457(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2458 github.com/goccy/spidermonkeywasm2go/p4.Fn2458
func Fn2458(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2462 github.com/goccy/spidermonkeywasm2go/p4.Fn2462
func Fn2462(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2466 github.com/goccy/spidermonkeywasm2go/p4.Fn2466
func Fn2466(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2468 github.com/goccy/spidermonkeywasm2go/p0.Fn2468
func Fn2468(m *base.Module, l0 int32) int32

//go:linkname Fn2474 github.com/goccy/spidermonkeywasm2go/p4.Fn2474
func Fn2474(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2477 github.com/goccy/spidermonkeywasm2go/p2.Fn2477
func Fn2477(m *base.Module, l0 int32) int32

//go:linkname Fn2483 github.com/goccy/spidermonkeywasm2go/p3.Fn2483
func Fn2483(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2493 github.com/goccy/spidermonkeywasm2go/p3.Fn2493
func Fn2493(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2502 github.com/goccy/spidermonkeywasm2go/p4.Fn2502
func Fn2502(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2503 github.com/goccy/spidermonkeywasm2go/p4.Fn2503
func Fn2503(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2504 github.com/goccy/spidermonkeywasm2go/p4.Fn2504
func Fn2504(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2505 github.com/goccy/spidermonkeywasm2go/p4.Fn2505
func Fn2505(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2506 github.com/goccy/spidermonkeywasm2go/p0.Fn2506
func Fn2506(m *base.Module, l0 int32) int32

//go:linkname Fn2517 github.com/goccy/spidermonkeywasm2go/p3.Fn2517
func Fn2517(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2621 github.com/goccy/spidermonkeywasm2go/p4.Fn2621
func Fn2621(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2622 github.com/goccy/spidermonkeywasm2go/p3.Fn2622
func Fn2622(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) int32

//go:linkname Fn2623 github.com/goccy/spidermonkeywasm2go/p3.Fn2623
func Fn2623(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) int32

//go:linkname Fn2634 github.com/goccy/spidermonkeywasm2go/p0.Fn2634
func Fn2634(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2641 github.com/goccy/spidermonkeywasm2go/p3.Fn2641
func Fn2641(m *base.Module, l0 int32)

//go:linkname Fn2644 github.com/goccy/spidermonkeywasm2go/p4.Fn2644
func Fn2644(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2645 github.com/goccy/spidermonkeywasm2go/p2.Fn2645
func Fn2645(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2660 github.com/goccy/spidermonkeywasm2go/p0.Fn2660
func Fn2660(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2663 github.com/goccy/spidermonkeywasm2go/p3.Fn2663
func Fn2663(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2664 github.com/goccy/spidermonkeywasm2go/p4.Fn2664
func Fn2664(m *base.Module, l0 int32)

//go:linkname Fn2671 github.com/goccy/spidermonkeywasm2go/p0.Fn2671
func Fn2671(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn2679 github.com/goccy/spidermonkeywasm2go/p0.Fn2679
func Fn2679(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2688 github.com/goccy/spidermonkeywasm2go/p3.Fn2688
func Fn2688(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2694 github.com/goccy/spidermonkeywasm2go/p4.Fn2694
func Fn2694(m *base.Module, l0 int32)

//go:linkname Fn2706 github.com/goccy/spidermonkeywasm2go/p4.Fn2706
func Fn2706(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2710 github.com/goccy/spidermonkeywasm2go/p4.Fn2710
func Fn2710(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2865 github.com/goccy/spidermonkeywasm2go/p3.Fn2865
func Fn2865(m *base.Module, l0 int32) int32

//go:linkname Fn2866 github.com/goccy/spidermonkeywasm2go/p4.Fn2866
func Fn2866(m *base.Module, l0 float64) int32

//go:linkname Fn2867 github.com/goccy/spidermonkeywasm2go/p4.Fn2867
func Fn2867(m *base.Module, l0 float64) int32

//go:linkname Fn2868 github.com/goccy/spidermonkeywasm2go/p4.Fn2868
func Fn2868(m *base.Module, l0 float64) int32

//go:linkname Fn2869 github.com/goccy/spidermonkeywasm2go/p4.Fn2869
func Fn2869(m *base.Module, l0 float64) int32

//go:linkname Fn2870 github.com/goccy/spidermonkeywasm2go/p4.Fn2870
func Fn2870(m *base.Module, l0 float64) int32

//go:linkname Fn2894 github.com/goccy/spidermonkeywasm2go/p4.Fn2894
func Fn2894(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn2900 github.com/goccy/spidermonkeywasm2go/p4.Fn2900
func Fn2900(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2903 github.com/goccy/spidermonkeywasm2go/p4.Fn2903
func Fn2903(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2904 github.com/goccy/spidermonkeywasm2go/p4.Fn2904
func Fn2904(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2905 github.com/goccy/spidermonkeywasm2go/p4.Fn2905
func Fn2905(m *base.Module, l0 int32, l1 int32) int32

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

//go:linkname Fn2928 github.com/goccy/spidermonkeywasm2go/p0.Fn2928
func Fn2928(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2938 github.com/goccy/spidermonkeywasm2go/p0.Fn2938
func Fn2938(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2943 github.com/goccy/spidermonkeywasm2go/p3.Fn2943
func Fn2943(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn2944 github.com/goccy/spidermonkeywasm2go/p4.Fn2944
func Fn2944(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2945 github.com/goccy/spidermonkeywasm2go/p3.Fn2945
func Fn2945(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2947 github.com/goccy/spidermonkeywasm2go/p4.Fn2947
func Fn2947(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2948 github.com/goccy/spidermonkeywasm2go/p3.Fn2948
func Fn2948(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2950 github.com/goccy/spidermonkeywasm2go/p3.Fn2950
func Fn2950(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2951 github.com/goccy/spidermonkeywasm2go/p4.Fn2951
func Fn2951(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2952 github.com/goccy/spidermonkeywasm2go/p4.Fn2952
func Fn2952(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2953 github.com/goccy/spidermonkeywasm2go/p3.Fn2953
func Fn2953(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2954 github.com/goccy/spidermonkeywasm2go/p4.Fn2954
func Fn2954(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2955 github.com/goccy/spidermonkeywasm2go/p4.Fn2955
func Fn2955(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2956 github.com/goccy/spidermonkeywasm2go/p4.Fn2956
func Fn2956(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2957 github.com/goccy/spidermonkeywasm2go/p4.Fn2957
func Fn2957(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn2958 github.com/goccy/spidermonkeywasm2go/p4.Fn2958
func Fn2958(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2959 github.com/goccy/spidermonkeywasm2go/p4.Fn2959
func Fn2959(m *base.Module, l0 int32, l1 float64) int32

//go:linkname Fn2961 github.com/goccy/spidermonkeywasm2go/p4.Fn2961
func Fn2961(m *base.Module, l0 int32)

//go:linkname Fn2963 github.com/goccy/spidermonkeywasm2go/p3.Fn2963
func Fn2963(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2964 github.com/goccy/spidermonkeywasm2go/p3.Fn2964
func Fn2964(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2965 github.com/goccy/spidermonkeywasm2go/p3.Fn2965
func Fn2965(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2966 github.com/goccy/spidermonkeywasm2go/p4.Fn2966
func Fn2966(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2967 github.com/goccy/spidermonkeywasm2go/p4.Fn2967
func Fn2967(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn2968 github.com/goccy/spidermonkeywasm2go/p4.Fn2968
func Fn2968(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2969 github.com/goccy/spidermonkeywasm2go/p2.Fn2969
func Fn2969(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2972 github.com/goccy/spidermonkeywasm2go/p3.Fn2972
func Fn2972(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn2975 github.com/goccy/spidermonkeywasm2go/p4.Fn2975
func Fn2975(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2976 github.com/goccy/spidermonkeywasm2go/p4.Fn2976
func Fn2976(m *base.Module, l0 int32) int32

//go:linkname Fn2977 github.com/goccy/spidermonkeywasm2go/p3.Fn2977
func Fn2977(m *base.Module, l0 int32) int32

//go:linkname Fn2978 github.com/goccy/spidermonkeywasm2go/p4.Fn2978
func Fn2978(m *base.Module, l0 int32)

//go:linkname Fn2979 github.com/goccy/spidermonkeywasm2go/p4.Fn2979
func Fn2979(m *base.Module, l0 int32) int32

//go:linkname Fn2981 github.com/goccy/spidermonkeywasm2go/p4.Fn2981
func Fn2981(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2982 github.com/goccy/spidermonkeywasm2go/p4.Fn2982
func Fn2982(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn2995 github.com/goccy/spidermonkeywasm2go/p3.Fn2995
func Fn2995(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn2999 github.com/goccy/spidermonkeywasm2go/p4.Fn2999
func Fn2999(m *base.Module, l0 float64) int32

//go:linkname Fn3013 github.com/goccy/spidermonkeywasm2go/p0.Fn3013
func Fn3013(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3014 github.com/goccy/spidermonkeywasm2go/p0.Fn3014
func Fn3014(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3015 github.com/goccy/spidermonkeywasm2go/p0.Fn3015
func Fn3015(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3016 github.com/goccy/spidermonkeywasm2go/p0.Fn3016
func Fn3016(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3017 github.com/goccy/spidermonkeywasm2go/p0.Fn3017
func Fn3017(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3018 github.com/goccy/spidermonkeywasm2go/p0.Fn3018
func Fn3018(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3019 github.com/goccy/spidermonkeywasm2go/p0.Fn3019
func Fn3019(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3020 github.com/goccy/spidermonkeywasm2go/p0.Fn3020
func Fn3020(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3021 github.com/goccy/spidermonkeywasm2go/p0.Fn3021
func Fn3021(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3022 github.com/goccy/spidermonkeywasm2go/p0.Fn3022
func Fn3022(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3023 github.com/goccy/spidermonkeywasm2go/p0.Fn3023
func Fn3023(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3024 github.com/goccy/spidermonkeywasm2go/p0.Fn3024
func Fn3024(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3025 github.com/goccy/spidermonkeywasm2go/p4.Fn3025
func Fn3025(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3026 github.com/goccy/spidermonkeywasm2go/p4.Fn3026
func Fn3026(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3027 github.com/goccy/spidermonkeywasm2go/p4.Fn3027
func Fn3027(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3028 github.com/goccy/spidermonkeywasm2go/p4.Fn3028
func Fn3028(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3029 github.com/goccy/spidermonkeywasm2go/p4.Fn3029
func Fn3029(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3030 github.com/goccy/spidermonkeywasm2go/p4.Fn3030
func Fn3030(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3031 github.com/goccy/spidermonkeywasm2go/p4.Fn3031
func Fn3031(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3032 github.com/goccy/spidermonkeywasm2go/p4.Fn3032
func Fn3032(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3033 github.com/goccy/spidermonkeywasm2go/p4.Fn3033
func Fn3033(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3034 github.com/goccy/spidermonkeywasm2go/p4.Fn3034
func Fn3034(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3035 github.com/goccy/spidermonkeywasm2go/p4.Fn3035
func Fn3035(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3036 github.com/goccy/spidermonkeywasm2go/p4.Fn3036
func Fn3036(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3037 github.com/goccy/spidermonkeywasm2go/p4.Fn3037
func Fn3037(m *base.Module, l0 int32) int32

//go:linkname Fn3038 github.com/goccy/spidermonkeywasm2go/p4.Fn3038
func Fn3038(m *base.Module, l0 int32) int32

//go:linkname Fn3039 github.com/goccy/spidermonkeywasm2go/p4.Fn3039
func Fn3039(m *base.Module, l0 int32) int32

//go:linkname Fn3040 github.com/goccy/spidermonkeywasm2go/p4.Fn3040
func Fn3040(m *base.Module, l0 int32) int32

//go:linkname Fn3041 github.com/goccy/spidermonkeywasm2go/p4.Fn3041
func Fn3041(m *base.Module, l0 int32) int32

//go:linkname Fn3042 github.com/goccy/spidermonkeywasm2go/p4.Fn3042
func Fn3042(m *base.Module, l0 int32) int32

//go:linkname Fn3043 github.com/goccy/spidermonkeywasm2go/p4.Fn3043
func Fn3043(m *base.Module, l0 int32) int32

//go:linkname Fn3044 github.com/goccy/spidermonkeywasm2go/p4.Fn3044
func Fn3044(m *base.Module, l0 int32) int32

//go:linkname Fn3045 github.com/goccy/spidermonkeywasm2go/p4.Fn3045
func Fn3045(m *base.Module, l0 int32) int32

//go:linkname Fn3046 github.com/goccy/spidermonkeywasm2go/p4.Fn3046
func Fn3046(m *base.Module, l0 int32) int32

//go:linkname Fn3047 github.com/goccy/spidermonkeywasm2go/p4.Fn3047
func Fn3047(m *base.Module, l0 int32) int32

//go:linkname Fn3048 github.com/goccy/spidermonkeywasm2go/p4.Fn3048
func Fn3048(m *base.Module, l0 int32) int32

//go:linkname Fn3049 github.com/goccy/spidermonkeywasm2go/p4.Fn3049
func Fn3049(m *base.Module, l0 int32) int32

//go:linkname Fn3050 github.com/goccy/spidermonkeywasm2go/p4.Fn3050
func Fn3050(m *base.Module, l0 int32) int32

//go:linkname Fn3051 github.com/goccy/spidermonkeywasm2go/p4.Fn3051
func Fn3051(m *base.Module, l0 int32) int32

//go:linkname Fn3052 github.com/goccy/spidermonkeywasm2go/p4.Fn3052
func Fn3052(m *base.Module, l0 int32) int32

//go:linkname Fn3053 github.com/goccy/spidermonkeywasm2go/p4.Fn3053
func Fn3053(m *base.Module, l0 int32) int32

//go:linkname Fn3054 github.com/goccy/spidermonkeywasm2go/p4.Fn3054
func Fn3054(m *base.Module, l0 int32) int32

//go:linkname Fn3055 github.com/goccy/spidermonkeywasm2go/p4.Fn3055
func Fn3055(m *base.Module, l0 int32) int32

//go:linkname Fn3056 github.com/goccy/spidermonkeywasm2go/p4.Fn3056
func Fn3056(m *base.Module, l0 int32) int32

//go:linkname Fn3057 github.com/goccy/spidermonkeywasm2go/p4.Fn3057
func Fn3057(m *base.Module, l0 int32) int32

//go:linkname Fn3058 github.com/goccy/spidermonkeywasm2go/p4.Fn3058
func Fn3058(m *base.Module, l0 int32) int32

//go:linkname Fn3059 github.com/goccy/spidermonkeywasm2go/p4.Fn3059
func Fn3059(m *base.Module, l0 int32) int32

//go:linkname Fn3060 github.com/goccy/spidermonkeywasm2go/p4.Fn3060
func Fn3060(m *base.Module, l0 int32) int32

//go:linkname Fn3102 github.com/goccy/spidermonkeywasm2go/p0.Fn3102
func Fn3102(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3111 github.com/goccy/spidermonkeywasm2go/p3.Fn3111
func Fn3111(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3125 github.com/goccy/spidermonkeywasm2go/p3.Fn3125
func Fn3125(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3126 github.com/goccy/spidermonkeywasm2go/p0.Fn3126
func Fn3126(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3127 github.com/goccy/spidermonkeywasm2go/p3.Fn3127
func Fn3127(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3129 github.com/goccy/spidermonkeywasm2go/p3.Fn3129
func Fn3129(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3131 github.com/goccy/spidermonkeywasm2go/p3.Fn3131
func Fn3131(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3133 github.com/goccy/spidermonkeywasm2go/p3.Fn3133
func Fn3133(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3135 github.com/goccy/spidermonkeywasm2go/p3.Fn3135
func Fn3135(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3137 github.com/goccy/spidermonkeywasm2go/p3.Fn3137
func Fn3137(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3139 github.com/goccy/spidermonkeywasm2go/p3.Fn3139
func Fn3139(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3141 github.com/goccy/spidermonkeywasm2go/p3.Fn3141
func Fn3141(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3143 github.com/goccy/spidermonkeywasm2go/p3.Fn3143
func Fn3143(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3145 github.com/goccy/spidermonkeywasm2go/p3.Fn3145
func Fn3145(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3147 github.com/goccy/spidermonkeywasm2go/p3.Fn3147
func Fn3147(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3155 github.com/goccy/spidermonkeywasm2go/p0.Fn3155
func Fn3155(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3157 github.com/goccy/spidermonkeywasm2go/p0.Fn3157
func Fn3157(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3159 github.com/goccy/spidermonkeywasm2go/p0.Fn3159
func Fn3159(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3161 github.com/goccy/spidermonkeywasm2go/p0.Fn3161
func Fn3161(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3166 github.com/goccy/spidermonkeywasm2go/p0.Fn3166
func Fn3166(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3168 github.com/goccy/spidermonkeywasm2go/p0.Fn3168
func Fn3168(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3170 github.com/goccy/spidermonkeywasm2go/p0.Fn3170
func Fn3170(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3175 github.com/goccy/spidermonkeywasm2go/p0.Fn3175
func Fn3175(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3195 github.com/goccy/spidermonkeywasm2go/p4.Fn3195
func Fn3195(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3198 github.com/goccy/spidermonkeywasm2go/p4.Fn3198
func Fn3198(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3201 github.com/goccy/spidermonkeywasm2go/p4.Fn3201
func Fn3201(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3204 github.com/goccy/spidermonkeywasm2go/p2.Fn3204
func Fn3204(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3205 github.com/goccy/spidermonkeywasm2go/p3.Fn3205
func Fn3205(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3206 github.com/goccy/spidermonkeywasm2go/p2.Fn3206
func Fn3206(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3207 github.com/goccy/spidermonkeywasm2go/p3.Fn3207
func Fn3207(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3208 github.com/goccy/spidermonkeywasm2go/p2.Fn3208
func Fn3208(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3209 github.com/goccy/spidermonkeywasm2go/p3.Fn3209
func Fn3209(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3210 github.com/goccy/spidermonkeywasm2go/p4.Fn3210
func Fn3210(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3213 github.com/goccy/spidermonkeywasm2go/p2.Fn3213
func Fn3213(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3214 github.com/goccy/spidermonkeywasm2go/p3.Fn3214
func Fn3214(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3284 github.com/goccy/spidermonkeywasm2go/p3.Fn3284
func Fn3284(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3285 github.com/goccy/spidermonkeywasm2go/p3.Fn3285
func Fn3285(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3306 github.com/goccy/spidermonkeywasm2go/p3.Fn3306
func Fn3306(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3363 github.com/goccy/spidermonkeywasm2go/p4.Fn3363
func Fn3363(m *base.Module, l0 int32) int32

//go:linkname Fn3364 github.com/goccy/spidermonkeywasm2go/p2.Fn3364
func Fn3364(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3365 github.com/goccy/spidermonkeywasm2go/p2.Fn3365
func Fn3365(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3367 github.com/goccy/spidermonkeywasm2go/p2.Fn3367
func Fn3367(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn3371 github.com/goccy/spidermonkeywasm2go/p3.Fn3371
func Fn3371(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3372 github.com/goccy/spidermonkeywasm2go/p3.Fn3372
func Fn3372(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn3373 github.com/goccy/spidermonkeywasm2go/p3.Fn3373
func Fn3373(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3392 github.com/goccy/spidermonkeywasm2go/p0.Fn3392
func Fn3392(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3404 github.com/goccy/spidermonkeywasm2go/p3.Fn3404
func Fn3404(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3417 github.com/goccy/spidermonkeywasm2go/p2.Fn3417
func Fn3417(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3424 github.com/goccy/spidermonkeywasm2go/p0.Fn3424
func Fn3424(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3426 github.com/goccy/spidermonkeywasm2go/p0.Fn3426
func Fn3426(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3427 github.com/goccy/spidermonkeywasm2go/p0.Fn3427
func Fn3427(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3492 github.com/goccy/spidermonkeywasm2go/p4.Fn3492
func Fn3492(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3537 github.com/goccy/spidermonkeywasm2go/p4.Fn3537
func Fn3537(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3540 github.com/goccy/spidermonkeywasm2go/p4.Fn3540
func Fn3540(m *base.Module, l0 int32) int32

//go:linkname Fn3576 github.com/goccy/spidermonkeywasm2go/p4.Fn3576
func Fn3576(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3577 github.com/goccy/spidermonkeywasm2go/p3.Fn3577
func Fn3577(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3629 github.com/goccy/spidermonkeywasm2go/p4.Fn3629
func Fn3629(m *base.Module, l0 int32)

//go:linkname Fn3669 github.com/goccy/spidermonkeywasm2go/p3.Fn3669
func Fn3669(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3671 github.com/goccy/spidermonkeywasm2go/p3.Fn3671
func Fn3671(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3672 github.com/goccy/spidermonkeywasm2go/p3.Fn3672
func Fn3672(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3674 github.com/goccy/spidermonkeywasm2go/p4.Fn3674
func Fn3674(m *base.Module, l0 int32) int32

//go:linkname Fn3675 github.com/goccy/spidermonkeywasm2go/p4.Fn3675
func Fn3675(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3676 github.com/goccy/spidermonkeywasm2go/p4.Fn3676
func Fn3676(m *base.Module, l0 int32, l1 int32, l2 int32) int32

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

//go:linkname Fn3738 github.com/goccy/spidermonkeywasm2go/p4.Fn3738
func Fn3738(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3742 github.com/goccy/spidermonkeywasm2go/p0.Fn3742
func Fn3742(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3743 github.com/goccy/spidermonkeywasm2go/p0.Fn3743
func Fn3743(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3744 github.com/goccy/spidermonkeywasm2go/p0.Fn3744
func Fn3744(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3745 github.com/goccy/spidermonkeywasm2go/p2.Fn3745
func Fn3745(m *base.Module, l0 int32)

//go:linkname Fn3746 github.com/goccy/spidermonkeywasm2go/p4.Fn3746
func Fn3746(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3747 github.com/goccy/spidermonkeywasm2go/p4.Fn3747
func Fn3747(m *base.Module, l0 int32)

//go:linkname Fn3754 github.com/goccy/spidermonkeywasm2go/p0.Fn3754
func Fn3754(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3758 github.com/goccy/spidermonkeywasm2go/p4.Fn3758
func Fn3758(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3759 github.com/goccy/spidermonkeywasm2go/p4.Fn3759
func Fn3759(m *base.Module, l0 int32)

//go:linkname Fn3761 github.com/goccy/spidermonkeywasm2go/p0.Fn3761
func Fn3761(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3763 github.com/goccy/spidermonkeywasm2go/p4.Fn3763
func Fn3763(m *base.Module, l0 int32)

//go:linkname Fn3764 github.com/goccy/spidermonkeywasm2go/p4.Fn3764
func Fn3764(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3772 github.com/goccy/spidermonkeywasm2go/p4.Fn3772
func Fn3772(m *base.Module, l0 int32)

//go:linkname Fn3773 github.com/goccy/spidermonkeywasm2go/p3.Fn3773
func Fn3773(m *base.Module, l0 float64, l1 float64, l2 float64) float64

//go:linkname Fn3774 github.com/goccy/spidermonkeywasm2go/p3.Fn3774
func Fn3774(m *base.Module, l0 int32, l1 float64)

//go:linkname Fn3779 github.com/goccy/spidermonkeywasm2go/p4.Fn3779
func Fn3779(m *base.Module, l0 float64) float64

//go:linkname Fn3780 github.com/goccy/spidermonkeywasm2go/p3.Fn3780
func Fn3780(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3781 github.com/goccy/spidermonkeywasm2go/p4.Fn3781
func Fn3781(m *base.Module, l0 int32) int32

//go:linkname Fn3785 github.com/goccy/spidermonkeywasm2go/p4.Fn3785
func Fn3785(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn3788 github.com/goccy/spidermonkeywasm2go/p3.Fn3788
func Fn3788(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3789 github.com/goccy/spidermonkeywasm2go/p4.Fn3789
func Fn3789(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3790 github.com/goccy/spidermonkeywasm2go/p4.Fn3790
func Fn3790(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3792 github.com/goccy/spidermonkeywasm2go/p3.Fn3792
func Fn3792(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn3793 github.com/goccy/spidermonkeywasm2go/p4.Fn3793
func Fn3793(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3795 github.com/goccy/spidermonkeywasm2go/p0.Fn3795
func Fn3795(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3805 github.com/goccy/spidermonkeywasm2go/p4.Fn3805
func Fn3805(m *base.Module, l0 int32, l1 float64, l2 int32) int32

//go:linkname Fn3812 github.com/goccy/spidermonkeywasm2go/p4.Fn3812
func Fn3812(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3813 github.com/goccy/spidermonkeywasm2go/p0.Fn3813
func Fn3813(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3814 github.com/goccy/spidermonkeywasm2go/p3.Fn3814
func Fn3814(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn3815 github.com/goccy/spidermonkeywasm2go/p4.Fn3815
func Fn3815(m *base.Module, l0 int32, l1 int32, l2 int32) float64

//go:linkname Fn3816 github.com/goccy/spidermonkeywasm2go/p3.Fn3816
func Fn3816(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn3817 github.com/goccy/spidermonkeywasm2go/p4.Fn3817
func Fn3817(m *base.Module, l0 int32, l1 int32, l2 int32) float64

//go:linkname Fn3818 github.com/goccy/spidermonkeywasm2go/p4.Fn3818
func Fn3818(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3820 github.com/goccy/spidermonkeywasm2go/p0.Fn3820
func Fn3820(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3823 github.com/goccy/spidermonkeywasm2go/p0.Fn3823
func Fn3823(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3835 github.com/goccy/spidermonkeywasm2go/p4.Fn3835
func Fn3835(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3836 github.com/goccy/spidermonkeywasm2go/p4.Fn3836
func Fn3836(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn3837 github.com/goccy/spidermonkeywasm2go/p4.Fn3837
func Fn3837(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3838 github.com/goccy/spidermonkeywasm2go/p4.Fn3838
func Fn3838(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3899 github.com/goccy/spidermonkeywasm2go/p3.Fn3899
func Fn3899(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3907 github.com/goccy/spidermonkeywasm2go/p3.Fn3907
func Fn3907(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3909 github.com/goccy/spidermonkeywasm2go/p4.Fn3909
func Fn3909(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3910 github.com/goccy/spidermonkeywasm2go/p2.Fn3910
func Fn3910(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn3919 github.com/goccy/spidermonkeywasm2go/p4.Fn3919
func Fn3919(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3922 github.com/goccy/spidermonkeywasm2go/p0.Fn3922
func Fn3922(m *base.Module, l0 int32)

//go:linkname Fn3927 github.com/goccy/spidermonkeywasm2go/p0.Fn3927
func Fn3927(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3938 github.com/goccy/spidermonkeywasm2go/p0.Fn3938
func Fn3938(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3939 github.com/goccy/spidermonkeywasm2go/p0.Fn3939
func Fn3939(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn3941 github.com/goccy/spidermonkeywasm2go/p0.Fn3941
func Fn3941(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3942 github.com/goccy/spidermonkeywasm2go/p4.Fn3942
func Fn3942(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn3957 github.com/goccy/spidermonkeywasm2go/p3.Fn3957
func Fn3957(m *base.Module, l0 int32)

//go:linkname Fn3959 github.com/goccy/spidermonkeywasm2go/p4.Fn3959
func Fn3959(m *base.Module, l0 int32)

//go:linkname Fn3960 github.com/goccy/spidermonkeywasm2go/p4.Fn3960
func Fn3960(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn3961 github.com/goccy/spidermonkeywasm2go/p4.Fn3961
func Fn3961(m *base.Module, l0 int32) int32

//go:linkname Fn3962 github.com/goccy/spidermonkeywasm2go/p4.Fn3962
func Fn3962(m *base.Module, l0 int32) int32

//go:linkname Fn3963 github.com/goccy/spidermonkeywasm2go/p4.Fn3963
func Fn3963(m *base.Module, l0 int32) int32

//go:linkname Fn3964 github.com/goccy/spidermonkeywasm2go/p4.Fn3964
func Fn3964(m *base.Module, l0 int32) int32

//go:linkname Fn3965 github.com/goccy/spidermonkeywasm2go/p4.Fn3965
func Fn3965(m *base.Module, l0 int32) int32

//go:linkname Fn3966 github.com/goccy/spidermonkeywasm2go/p4.Fn3966
func Fn3966(m *base.Module, l0 int32) int32

//go:linkname Fn3967 github.com/goccy/spidermonkeywasm2go/p4.Fn3967
func Fn3967(m *base.Module, l0 int32) int32

//go:linkname Fn4041 github.com/goccy/spidermonkeywasm2go/p3.Fn4041
func Fn4041(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4046 github.com/goccy/spidermonkeywasm2go/p0.Fn4046
func Fn4046(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4063 github.com/goccy/spidermonkeywasm2go/p4.Fn4063
func Fn4063(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4064 github.com/goccy/spidermonkeywasm2go/p4.Fn4064
func Fn4064(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4065 github.com/goccy/spidermonkeywasm2go/p0.Fn4065
func Fn4065(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4068 github.com/goccy/spidermonkeywasm2go/p4.Fn4068
func Fn4068(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4115 github.com/goccy/spidermonkeywasm2go/p0.Fn4115
func Fn4115(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4160 github.com/goccy/spidermonkeywasm2go/p4.Fn4160
func Fn4160(m *base.Module, l0 int32)

//go:linkname Fn4376 github.com/goccy/spidermonkeywasm2go/p3.Fn4376
func Fn4376(m *base.Module, l0 int32) int32

//go:linkname Fn4378 github.com/goccy/spidermonkeywasm2go/p0.Fn4378
func Fn4378(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4379 github.com/goccy/spidermonkeywasm2go/p0.Fn4379
func Fn4379(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4380 github.com/goccy/spidermonkeywasm2go/p4.Fn4380
func Fn4380(m *base.Module, l0 int32)

//go:linkname Fn4389 github.com/goccy/spidermonkeywasm2go/p4.Fn4389
func Fn4389(m *base.Module, l0 int32) int32

//go:linkname Fn4390 github.com/goccy/spidermonkeywasm2go/p3.Fn4390
func Fn4390(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4391 github.com/goccy/spidermonkeywasm2go/p0.Fn4391
func Fn4391(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4392 github.com/goccy/spidermonkeywasm2go/p2.Fn4392
func Fn4392(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4393 github.com/goccy/spidermonkeywasm2go/p0.Fn4393
func Fn4393(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4394 github.com/goccy/spidermonkeywasm2go/p0.Fn4394
func Fn4394(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4397 github.com/goccy/spidermonkeywasm2go/p3.Fn4397
func Fn4397(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4505 github.com/goccy/spidermonkeywasm2go/p4.Fn4505
func Fn4505(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4515 github.com/goccy/spidermonkeywasm2go/p3.Fn4515
func Fn4515(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4523 github.com/goccy/spidermonkeywasm2go/p0.Fn4523
func Fn4523(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4525 github.com/goccy/spidermonkeywasm2go/p3.Fn4525
func Fn4525(m *base.Module, l0 int32)

//go:linkname Fn4526 github.com/goccy/spidermonkeywasm2go/p0.Fn4526
func Fn4526(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4571 github.com/goccy/spidermonkeywasm2go/p4.Fn4571
func Fn4571(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4572 github.com/goccy/spidermonkeywasm2go/p4.Fn4572
func Fn4572(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4577 github.com/goccy/spidermonkeywasm2go/p4.Fn4577
func Fn4577(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn4578 github.com/goccy/spidermonkeywasm2go/p4.Fn4578
func Fn4578(m *base.Module, l0 int32) int32

//go:linkname Fn4587 github.com/goccy/spidermonkeywasm2go/p4.Fn4587
func Fn4587(m *base.Module, l0 int32) int32

//go:linkname Fn4591 github.com/goccy/spidermonkeywasm2go/p3.Fn4591
func Fn4591(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4592 github.com/goccy/spidermonkeywasm2go/p3.Fn4592
func Fn4592(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4594 github.com/goccy/spidermonkeywasm2go/p4.Fn4594
func Fn4594(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4595 github.com/goccy/spidermonkeywasm2go/p2.Fn4595
func Fn4595(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4598 github.com/goccy/spidermonkeywasm2go/p4.Fn4598
func Fn4598(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4603 github.com/goccy/spidermonkeywasm2go/p4.Fn4603
func Fn4603(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4606 github.com/goccy/spidermonkeywasm2go/p4.Fn4606
func Fn4606(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4607 github.com/goccy/spidermonkeywasm2go/p3.Fn4607
func Fn4607(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4626 github.com/goccy/spidermonkeywasm2go/p4.Fn4626
func Fn4626(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4641 github.com/goccy/spidermonkeywasm2go/p4.Fn4641
func Fn4641(m *base.Module, l0 int32) int32

//go:linkname Fn4642 github.com/goccy/spidermonkeywasm2go/p4.Fn4642
func Fn4642(m *base.Module, l0 int32) int32

//go:linkname Fn4643 github.com/goccy/spidermonkeywasm2go/p3.Fn4643
func Fn4643(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4644 github.com/goccy/spidermonkeywasm2go/p4.Fn4644
func Fn4644(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4645 github.com/goccy/spidermonkeywasm2go/p4.Fn4645
func Fn4645(m *base.Module, l0 int32) int32

//go:linkname Fn4646 github.com/goccy/spidermonkeywasm2go/p4.Fn4646
func Fn4646(m *base.Module, l0 int32) int32

//go:linkname Fn4647 github.com/goccy/spidermonkeywasm2go/p4.Fn4647
func Fn4647(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4648 github.com/goccy/spidermonkeywasm2go/p4.Fn4648
func Fn4648(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4649 github.com/goccy/spidermonkeywasm2go/p3.Fn4649
func Fn4649(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4650 github.com/goccy/spidermonkeywasm2go/p3.Fn4650
func Fn4650(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4652 github.com/goccy/spidermonkeywasm2go/p4.Fn4652
func Fn4652(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4655 github.com/goccy/spidermonkeywasm2go/p2.Fn4655
func Fn4655(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn4656 github.com/goccy/spidermonkeywasm2go/p4.Fn4656
func Fn4656(m *base.Module, l0 int32)

//go:linkname Fn4657 github.com/goccy/spidermonkeywasm2go/p4.Fn4657
func Fn4657(m *base.Module, l0 int32)

//go:linkname Fn4660 github.com/goccy/spidermonkeywasm2go/p3.Fn4660
func Fn4660(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4661 github.com/goccy/spidermonkeywasm2go/p4.Fn4661
func Fn4661(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4662 github.com/goccy/spidermonkeywasm2go/p0.Fn4662
func Fn4662(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4668 github.com/goccy/spidermonkeywasm2go/p4.Fn4668
func Fn4668(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4671 github.com/goccy/spidermonkeywasm2go/p4.Fn4671
func Fn4671(m *base.Module, l0 int32) int32

//go:linkname Fn4676 github.com/goccy/spidermonkeywasm2go/p3.Fn4676
func Fn4676(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4708 github.com/goccy/spidermonkeywasm2go/p4.Fn4708
func Fn4708(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4747 github.com/goccy/spidermonkeywasm2go/p0.Fn4747
func Fn4747(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4788 github.com/goccy/spidermonkeywasm2go/p4.Fn4788
func Fn4788(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4789 github.com/goccy/spidermonkeywasm2go/p4.Fn4789
func Fn4789(m *base.Module, l0 int32) int32

//go:linkname Fn4806 github.com/goccy/spidermonkeywasm2go/p3.Fn4806
func Fn4806(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4807 github.com/goccy/spidermonkeywasm2go/p4.Fn4807
func Fn4807(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4809 github.com/goccy/spidermonkeywasm2go/p3.Fn4809
func Fn4809(m *base.Module, l0 int32) int32

//go:linkname Fn4812 github.com/goccy/spidermonkeywasm2go/p2.Fn4812
func Fn4812(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn4813 github.com/goccy/spidermonkeywasm2go/p0.Fn4813
func Fn4813(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn4814 github.com/goccy/spidermonkeywasm2go/p0.Fn4814
func Fn4814(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4815 github.com/goccy/spidermonkeywasm2go/p4.Fn4815
func Fn4815(m *base.Module, l0 int32)

//go:linkname Fn4820 github.com/goccy/spidermonkeywasm2go/p3.Fn4820
func Fn4820(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4826 github.com/goccy/spidermonkeywasm2go/p3.Fn4826
func Fn4826(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4841 github.com/goccy/spidermonkeywasm2go/p4.Fn4841
func Fn4841(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4843 github.com/goccy/spidermonkeywasm2go/p3.Fn4843
func Fn4843(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4845 github.com/goccy/spidermonkeywasm2go/p2.Fn4845
func Fn4845(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn4858 github.com/goccy/spidermonkeywasm2go/p3.Fn4858
func Fn4858(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4859 github.com/goccy/spidermonkeywasm2go/p0.Fn4859
func Fn4859(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4864 github.com/goccy/spidermonkeywasm2go/p2.Fn4864
func Fn4864(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4865 github.com/goccy/spidermonkeywasm2go/p2.Fn4865
func Fn4865(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4886 github.com/goccy/spidermonkeywasm2go/p4.Fn4886
func Fn4886(m *base.Module, l0 float64) int32

//go:linkname Fn4889 github.com/goccy/spidermonkeywasm2go/p4.Fn4889
func Fn4889(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4891 github.com/goccy/spidermonkeywasm2go/p4.Fn4891
func Fn4891(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn4892 github.com/goccy/spidermonkeywasm2go/p4.Fn4892
func Fn4892(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4893 github.com/goccy/spidermonkeywasm2go/p4.Fn4893
func Fn4893(m *base.Module, l0 int32)

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

//go:linkname Fn4901 github.com/goccy/spidermonkeywasm2go/p4.Fn4901
func Fn4901(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4902 github.com/goccy/spidermonkeywasm2go/p4.Fn4902
func Fn4902(m *base.Module, l0 int32)

//go:linkname Fn4905 github.com/goccy/spidermonkeywasm2go/p3.Fn4905
func Fn4905(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn4909 github.com/goccy/spidermonkeywasm2go/p0.Fn4909
func Fn4909(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn4925 github.com/goccy/spidermonkeywasm2go/p4.Fn4925
func Fn4925(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4928 github.com/goccy/spidermonkeywasm2go/p3.Fn4928
func Fn4928(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn4933 github.com/goccy/spidermonkeywasm2go/p3.Fn4933
func Fn4933(m *base.Module, l0 int32)

//go:linkname Fn4936 github.com/goccy/spidermonkeywasm2go/p4.Fn4936
func Fn4936(m *base.Module, l0 int32)

//go:linkname Fn4938 github.com/goccy/spidermonkeywasm2go/p4.Fn4938
func Fn4938(m *base.Module, l0 int32) int32

//go:linkname Fn4950 github.com/goccy/spidermonkeywasm2go/p3.Fn4950
func Fn4950(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn4969 github.com/goccy/spidermonkeywasm2go/p4.Fn4969
func Fn4969(m *base.Module, l0 int32)

//go:linkname Fn4970 github.com/goccy/spidermonkeywasm2go/p4.Fn4970
func Fn4970(m *base.Module, l0 int32)

//go:linkname Fn4981 github.com/goccy/spidermonkeywasm2go/p4.Fn4981
func Fn4981(m *base.Module, l0 int32)

//go:linkname Fn4984 github.com/goccy/spidermonkeywasm2go/p4.Fn4984
func Fn4984(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn4985 github.com/goccy/spidermonkeywasm2go/p4.Fn4985
func Fn4985(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5033 github.com/goccy/spidermonkeywasm2go/p4.Fn5033
func Fn5033(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5034 github.com/goccy/spidermonkeywasm2go/p4.Fn5034
func Fn5034(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5035 github.com/goccy/spidermonkeywasm2go/p4.Fn5035
func Fn5035(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5036 github.com/goccy/spidermonkeywasm2go/p4.Fn5036
func Fn5036(m *base.Module, l0 int32)

//go:linkname Fn5049 github.com/goccy/spidermonkeywasm2go/p4.Fn5049
func Fn5049(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5050 github.com/goccy/spidermonkeywasm2go/p4.Fn5050
func Fn5050(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5051 github.com/goccy/spidermonkeywasm2go/p4.Fn5051
func Fn5051(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5058 github.com/goccy/spidermonkeywasm2go/p4.Fn5058
func Fn5058(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5059 github.com/goccy/spidermonkeywasm2go/p4.Fn5059
func Fn5059(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5061 github.com/goccy/spidermonkeywasm2go/p3.Fn5061
func Fn5061(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5062 github.com/goccy/spidermonkeywasm2go/p4.Fn5062
func Fn5062(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5063 github.com/goccy/spidermonkeywasm2go/p3.Fn5063
func Fn5063(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5064 github.com/goccy/spidermonkeywasm2go/p3.Fn5064
func Fn5064(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5065 github.com/goccy/spidermonkeywasm2go/p3.Fn5065
func Fn5065(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5067 github.com/goccy/spidermonkeywasm2go/p3.Fn5067
func Fn5067(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5068 github.com/goccy/spidermonkeywasm2go/p3.Fn5068
func Fn5068(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5069 github.com/goccy/spidermonkeywasm2go/p3.Fn5069
func Fn5069(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5070 github.com/goccy/spidermonkeywasm2go/p3.Fn5070
func Fn5070(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5071 github.com/goccy/spidermonkeywasm2go/p3.Fn5071
func Fn5071(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5072 github.com/goccy/spidermonkeywasm2go/p4.Fn5072
func Fn5072(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5073 github.com/goccy/spidermonkeywasm2go/p3.Fn5073
func Fn5073(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5074 github.com/goccy/spidermonkeywasm2go/p3.Fn5074
func Fn5074(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5075 github.com/goccy/spidermonkeywasm2go/p3.Fn5075
func Fn5075(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5078 github.com/goccy/spidermonkeywasm2go/p4.Fn5078
func Fn5078(m *base.Module, l0 int32)

//go:linkname Fn5087 github.com/goccy/spidermonkeywasm2go/p4.Fn5087
func Fn5087(m *base.Module, l0 int32)

//go:linkname Fn5094 github.com/goccy/spidermonkeywasm2go/p2.Fn5094
func Fn5094(m *base.Module, l0 int32)

//go:linkname Fn5097 github.com/goccy/spidermonkeywasm2go/p3.Fn5097
func Fn5097(m *base.Module, l0 int32) int32

//go:linkname Fn5100 github.com/goccy/spidermonkeywasm2go/p4.Fn5100
func Fn5100(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5101 github.com/goccy/spidermonkeywasm2go/p4.Fn5101
func Fn5101(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5154 github.com/goccy/spidermonkeywasm2go/p2.Fn5154
func Fn5154(m *base.Module, l0 int32)

//go:linkname Fn5155 github.com/goccy/spidermonkeywasm2go/p4.Fn5155
func Fn5155(m *base.Module, l0 int32)

//go:linkname Fn5164 github.com/goccy/spidermonkeywasm2go/p4.Fn5164
func Fn5164(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5165 github.com/goccy/spidermonkeywasm2go/p2.Fn5165
func Fn5165(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5166 github.com/goccy/spidermonkeywasm2go/p2.Fn5166
func Fn5166(m *base.Module, l0 int32) int32

//go:linkname Fn5167 github.com/goccy/spidermonkeywasm2go/p4.Fn5167
func Fn5167(m *base.Module, l0 int32) int32

//go:linkname Fn5170 github.com/goccy/spidermonkeywasm2go/p3.Fn5170
func Fn5170(m *base.Module, l0 int32)

//go:linkname Fn5171 github.com/goccy/spidermonkeywasm2go/p3.Fn5171
func Fn5171(m *base.Module, l0 int32)

//go:linkname Fn5180 github.com/goccy/spidermonkeywasm2go/p4.Fn5180
func Fn5180(m *base.Module, l0 int32)

//go:linkname Fn5183 github.com/goccy/spidermonkeywasm2go/p3.Fn5183
func Fn5183(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5189 github.com/goccy/spidermonkeywasm2go/p4.Fn5189
func Fn5189(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5191 github.com/goccy/spidermonkeywasm2go/p4.Fn5191
func Fn5191(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn5192 github.com/goccy/spidermonkeywasm2go/p4.Fn5192
func Fn5192(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5195 github.com/goccy/spidermonkeywasm2go/p4.Fn5195
func Fn5195(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5241 github.com/goccy/spidermonkeywasm2go/p4.Fn5241
func Fn5241(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5242 github.com/goccy/spidermonkeywasm2go/p4.Fn5242
func Fn5242(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5243 github.com/goccy/spidermonkeywasm2go/p4.Fn5243
func Fn5243(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5244 github.com/goccy/spidermonkeywasm2go/p4.Fn5244
func Fn5244(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5245 github.com/goccy/spidermonkeywasm2go/p4.Fn5245
func Fn5245(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5246 github.com/goccy/spidermonkeywasm2go/p4.Fn5246
func Fn5246(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5247 github.com/goccy/spidermonkeywasm2go/p4.Fn5247
func Fn5247(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5248 github.com/goccy/spidermonkeywasm2go/p4.Fn5248
func Fn5248(m *base.Module, l0 int32) int32

//go:linkname Fn5249 github.com/goccy/spidermonkeywasm2go/p3.Fn5249
func Fn5249(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5250 github.com/goccy/spidermonkeywasm2go/p3.Fn5250
func Fn5250(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5251 github.com/goccy/spidermonkeywasm2go/p3.Fn5251
func Fn5251(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5252 github.com/goccy/spidermonkeywasm2go/p4.Fn5252
func Fn5252(m *base.Module)

//go:linkname Fn5253 github.com/goccy/spidermonkeywasm2go/p4.Fn5253
func Fn5253(m *base.Module, l0 int32)

//go:linkname Fn5255 github.com/goccy/spidermonkeywasm2go/p4.Fn5255
func Fn5255(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5259 github.com/goccy/spidermonkeywasm2go/p2.Fn5259
func Fn5259(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5260 github.com/goccy/spidermonkeywasm2go/p4.Fn5260
func Fn5260(m *base.Module)

//go:linkname Fn5262 github.com/goccy/spidermonkeywasm2go/p3.Fn5262
func Fn5262(m *base.Module, l0 int32)

//go:linkname Fn5265 github.com/goccy/spidermonkeywasm2go/p4.Fn5265
func Fn5265(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5267 github.com/goccy/spidermonkeywasm2go/p4.Fn5267
func Fn5267(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5274 github.com/goccy/spidermonkeywasm2go/p4.Fn5274
func Fn5274(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5275 github.com/goccy/spidermonkeywasm2go/p4.Fn5275
func Fn5275(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5276 github.com/goccy/spidermonkeywasm2go/p4.Fn5276
func Fn5276(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5279 github.com/goccy/spidermonkeywasm2go/p4.Fn5279
func Fn5279(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5411 github.com/goccy/spidermonkeywasm2go/p3.Fn5411
func Fn5411(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5428 github.com/goccy/spidermonkeywasm2go/p4.Fn5428
func Fn5428(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5429 github.com/goccy/spidermonkeywasm2go/p4.Fn5429
func Fn5429(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5432 github.com/goccy/spidermonkeywasm2go/p4.Fn5432
func Fn5432(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5435 github.com/goccy/spidermonkeywasm2go/p4.Fn5435
func Fn5435(m *base.Module, l0 int32)

//go:linkname Fn5436 github.com/goccy/spidermonkeywasm2go/p4.Fn5436
func Fn5436(m *base.Module, l0 int32) int32

//go:linkname Fn5437 github.com/goccy/spidermonkeywasm2go/p3.Fn5437
func Fn5437(m *base.Module, l0 int32)

//go:linkname Fn5438 github.com/goccy/spidermonkeywasm2go/p3.Fn5438
func Fn5438(m *base.Module, l0 int32)

//go:linkname Fn5439 github.com/goccy/spidermonkeywasm2go/p4.Fn5439
func Fn5439(m *base.Module, l0 int32) int32

//go:linkname Fn5440 github.com/goccy/spidermonkeywasm2go/p4.Fn5440
func Fn5440(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5442 github.com/goccy/spidermonkeywasm2go/p2.Fn5442
func Fn5442(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5443 github.com/goccy/spidermonkeywasm2go/p3.Fn5443
func Fn5443(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5444 github.com/goccy/spidermonkeywasm2go/p4.Fn5444
func Fn5444(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5445 github.com/goccy/spidermonkeywasm2go/p3.Fn5445
func Fn5445(m *base.Module, l0 int32) int32

//go:linkname Fn5446 github.com/goccy/spidermonkeywasm2go/p4.Fn5446
func Fn5446(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5447 github.com/goccy/spidermonkeywasm2go/p4.Fn5447
func Fn5447(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5448 github.com/goccy/spidermonkeywasm2go/p2.Fn5448
func Fn5448(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5449 github.com/goccy/spidermonkeywasm2go/p4.Fn5449
func Fn5449(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5450 github.com/goccy/spidermonkeywasm2go/p4.Fn5450
func Fn5450(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5451 github.com/goccy/spidermonkeywasm2go/p3.Fn5451
func Fn5451(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5452 github.com/goccy/spidermonkeywasm2go/p3.Fn5452
func Fn5452(m *base.Module, l0 int32)

//go:linkname Fn5453 github.com/goccy/spidermonkeywasm2go/p4.Fn5453
func Fn5453(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5454 github.com/goccy/spidermonkeywasm2go/p4.Fn5454
func Fn5454(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5455 github.com/goccy/spidermonkeywasm2go/p4.Fn5455
func Fn5455(m *base.Module, l0 int32) int32

//go:linkname Fn5456 github.com/goccy/spidermonkeywasm2go/p4.Fn5456
func Fn5456(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn5457 github.com/goccy/spidermonkeywasm2go/p3.Fn5457
func Fn5457(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5460 github.com/goccy/spidermonkeywasm2go/p4.Fn5460
func Fn5460(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5461 github.com/goccy/spidermonkeywasm2go/p3.Fn5461
func Fn5461(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5463 github.com/goccy/spidermonkeywasm2go/p3.Fn5463
func Fn5463(m *base.Module, l0 int32, l1 int32, l2 int64)

//go:linkname Fn5464 github.com/goccy/spidermonkeywasm2go/p4.Fn5464
func Fn5464(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5465 github.com/goccy/spidermonkeywasm2go/p4.Fn5465
func Fn5465(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5466 github.com/goccy/spidermonkeywasm2go/p4.Fn5466
func Fn5466(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5467 github.com/goccy/spidermonkeywasm2go/p4.Fn5467
func Fn5467(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5468 github.com/goccy/spidermonkeywasm2go/p4.Fn5468
func Fn5468(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5474 github.com/goccy/spidermonkeywasm2go/p4.Fn5474
func Fn5474(m *base.Module, l0 int32) int32

//go:linkname Fn5475 github.com/goccy/spidermonkeywasm2go/p3.Fn5475
func Fn5475(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5476 github.com/goccy/spidermonkeywasm2go/p3.Fn5476
func Fn5476(m *base.Module, l0 int32) int32

//go:linkname Fn5478 github.com/goccy/spidermonkeywasm2go/p4.Fn5478
func Fn5478(m *base.Module, l0 int32)

//go:linkname Fn5479 github.com/goccy/spidermonkeywasm2go/p3.Fn5479
func Fn5479(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5480 github.com/goccy/spidermonkeywasm2go/p3.Fn5480
func Fn5480(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5481 github.com/goccy/spidermonkeywasm2go/p4.Fn5481
func Fn5481(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5482 github.com/goccy/spidermonkeywasm2go/p3.Fn5482
func Fn5482(m *base.Module, l0 int32) int32

//go:linkname Fn5483 github.com/goccy/spidermonkeywasm2go/p2.Fn5483
func Fn5483(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5485 github.com/goccy/spidermonkeywasm2go/p4.Fn5485
func Fn5485(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn5489 github.com/goccy/spidermonkeywasm2go/p2.Fn5489
func Fn5489(m *base.Module, l0 int32) int32

//go:linkname Fn5490 github.com/goccy/spidermonkeywasm2go/p4.Fn5490
func Fn5490(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5500 github.com/goccy/spidermonkeywasm2go/p3.Fn5500
func Fn5500(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5502 github.com/goccy/spidermonkeywasm2go/p2.Fn5502
func Fn5502(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5554 github.com/goccy/spidermonkeywasm2go/p2.Fn5554
func Fn5554(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5566 github.com/goccy/spidermonkeywasm2go/p2.Fn5566
func Fn5566(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5869 github.com/goccy/spidermonkeywasm2go/p3.Fn5869
func Fn5869(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5871 github.com/goccy/spidermonkeywasm2go/p4.Fn5871
func Fn5871(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5872 github.com/goccy/spidermonkeywasm2go/p4.Fn5872
func Fn5872(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5873 github.com/goccy/spidermonkeywasm2go/p4.Fn5873
func Fn5873(m *base.Module, l0 int32) int32

//go:linkname Fn5886 github.com/goccy/spidermonkeywasm2go/p4.Fn5886
func Fn5886(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5891 github.com/goccy/spidermonkeywasm2go/p4.Fn5891
func Fn5891(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5892 github.com/goccy/spidermonkeywasm2go/p4.Fn5892
func Fn5892(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5894 github.com/goccy/spidermonkeywasm2go/p4.Fn5894
func Fn5894(m *base.Module, l0 int32)

//go:linkname Fn5895 github.com/goccy/spidermonkeywasm2go/p4.Fn5895
func Fn5895(m *base.Module, l0 int32)

//go:linkname Fn5896 github.com/goccy/spidermonkeywasm2go/p4.Fn5896
func Fn5896(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn5898 github.com/goccy/spidermonkeywasm2go/p4.Fn5898
func Fn5898(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5899 github.com/goccy/spidermonkeywasm2go/p3.Fn5899
func Fn5899(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5900 github.com/goccy/spidermonkeywasm2go/p3.Fn5900
func Fn5900(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5901 github.com/goccy/spidermonkeywasm2go/p4.Fn5901
func Fn5901(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5902 github.com/goccy/spidermonkeywasm2go/p4.Fn5902
func Fn5902(m *base.Module, l0 int32)

//go:linkname Fn5903 github.com/goccy/spidermonkeywasm2go/p3.Fn5903
func Fn5903(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn5904 github.com/goccy/spidermonkeywasm2go/p4.Fn5904
func Fn5904(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn5905 github.com/goccy/spidermonkeywasm2go/p4.Fn5905
func Fn5905(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5906 github.com/goccy/spidermonkeywasm2go/p4.Fn5906
func Fn5906(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5907 github.com/goccy/spidermonkeywasm2go/p4.Fn5907
func Fn5907(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5908 github.com/goccy/spidermonkeywasm2go/p4.Fn5908
func Fn5908(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5909 github.com/goccy/spidermonkeywasm2go/p3.Fn5909
func Fn5909(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn5910 github.com/goccy/spidermonkeywasm2go/p4.Fn5910
func Fn5910(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

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

//go:linkname Fn5918 github.com/goccy/spidermonkeywasm2go/p4.Fn5918
func Fn5918(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5919 github.com/goccy/spidermonkeywasm2go/p4.Fn5919
func Fn5919(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5920 github.com/goccy/spidermonkeywasm2go/p4.Fn5920
func Fn5920(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5921 github.com/goccy/spidermonkeywasm2go/p4.Fn5921
func Fn5921(m *base.Module, l0 int32) int32

//go:linkname Fn5922 github.com/goccy/spidermonkeywasm2go/p4.Fn5922
func Fn5922(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5923 github.com/goccy/spidermonkeywasm2go/p4.Fn5923
func Fn5923(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5924 github.com/goccy/spidermonkeywasm2go/p4.Fn5924
func Fn5924(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5925 github.com/goccy/spidermonkeywasm2go/p4.Fn5925
func Fn5925(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5927 github.com/goccy/spidermonkeywasm2go/p4.Fn5927
func Fn5927(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5929 github.com/goccy/spidermonkeywasm2go/p4.Fn5929
func Fn5929(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5930 github.com/goccy/spidermonkeywasm2go/p4.Fn5930
func Fn5930(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5932 github.com/goccy/spidermonkeywasm2go/p4.Fn5932
func Fn5932(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5934 github.com/goccy/spidermonkeywasm2go/p4.Fn5934
func Fn5934(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5937 github.com/goccy/spidermonkeywasm2go/p4.Fn5937
func Fn5937(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5938 github.com/goccy/spidermonkeywasm2go/p4.Fn5938
func Fn5938(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn5939 github.com/goccy/spidermonkeywasm2go/p4.Fn5939
func Fn5939(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5940 github.com/goccy/spidermonkeywasm2go/p4.Fn5940
func Fn5940(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5941 github.com/goccy/spidermonkeywasm2go/p4.Fn5941
func Fn5941(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5943 github.com/goccy/spidermonkeywasm2go/p2.Fn5943
func Fn5943(m *base.Module, l0 int32) int32

//go:linkname Fn5944 github.com/goccy/spidermonkeywasm2go/p2.Fn5944
func Fn5944(m *base.Module, l0 int32) int32

//go:linkname Fn5945 github.com/goccy/spidermonkeywasm2go/p3.Fn5945
func Fn5945(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5946 github.com/goccy/spidermonkeywasm2go/p3.Fn5946
func Fn5946(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5947 github.com/goccy/spidermonkeywasm2go/p3.Fn5947
func Fn5947(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5948 github.com/goccy/spidermonkeywasm2go/p3.Fn5948
func Fn5948(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5949 github.com/goccy/spidermonkeywasm2go/p3.Fn5949
func Fn5949(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5950 github.com/goccy/spidermonkeywasm2go/p3.Fn5950
func Fn5950(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5951 github.com/goccy/spidermonkeywasm2go/p2.Fn5951
func Fn5951(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5952 github.com/goccy/spidermonkeywasm2go/p3.Fn5952
func Fn5952(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5953 github.com/goccy/spidermonkeywasm2go/p2.Fn5953
func Fn5953(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5954 github.com/goccy/spidermonkeywasm2go/p3.Fn5954
func Fn5954(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5955 github.com/goccy/spidermonkeywasm2go/p3.Fn5955
func Fn5955(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5956 github.com/goccy/spidermonkeywasm2go/p3.Fn5956
func Fn5956(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5957 github.com/goccy/spidermonkeywasm2go/p2.Fn5957
func Fn5957(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5958 github.com/goccy/spidermonkeywasm2go/p2.Fn5958
func Fn5958(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5959 github.com/goccy/spidermonkeywasm2go/p4.Fn5959
func Fn5959(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5960 github.com/goccy/spidermonkeywasm2go/p3.Fn5960
func Fn5960(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5961 github.com/goccy/spidermonkeywasm2go/p4.Fn5961
func Fn5961(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5962 github.com/goccy/spidermonkeywasm2go/p4.Fn5962
func Fn5962(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5963 github.com/goccy/spidermonkeywasm2go/p4.Fn5963
func Fn5963(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5964 github.com/goccy/spidermonkeywasm2go/p4.Fn5964
func Fn5964(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5966 github.com/goccy/spidermonkeywasm2go/p2.Fn5966
func Fn5966(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5967 github.com/goccy/spidermonkeywasm2go/p4.Fn5967
func Fn5967(m *base.Module, l0 int32) int32

//go:linkname Fn5968 github.com/goccy/spidermonkeywasm2go/p4.Fn5968
func Fn5968(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5969 github.com/goccy/spidermonkeywasm2go/p4.Fn5969
func Fn5969(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5971 github.com/goccy/spidermonkeywasm2go/p2.Fn5971
func Fn5971(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5972 github.com/goccy/spidermonkeywasm2go/p4.Fn5972
func Fn5972(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn5973 github.com/goccy/spidermonkeywasm2go/p3.Fn5973
func Fn5973(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn5974 github.com/goccy/spidermonkeywasm2go/p4.Fn5974
func Fn5974(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5975 github.com/goccy/spidermonkeywasm2go/p4.Fn5975
func Fn5975(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5976 github.com/goccy/spidermonkeywasm2go/p4.Fn5976
func Fn5976(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5977 github.com/goccy/spidermonkeywasm2go/p4.Fn5977
func Fn5977(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5978 github.com/goccy/spidermonkeywasm2go/p4.Fn5978
func Fn5978(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5979 github.com/goccy/spidermonkeywasm2go/p4.Fn5979
func Fn5979(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5983 github.com/goccy/spidermonkeywasm2go/p4.Fn5983
func Fn5983(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5984 github.com/goccy/spidermonkeywasm2go/p4.Fn5984
func Fn5984(m *base.Module, l0 int32) int32

//go:linkname Fn5985 github.com/goccy/spidermonkeywasm2go/p4.Fn5985
func Fn5985(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5986 github.com/goccy/spidermonkeywasm2go/p4.Fn5986
func Fn5986(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn5988 github.com/goccy/spidermonkeywasm2go/p4.Fn5988
func Fn5988(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5989 github.com/goccy/spidermonkeywasm2go/p4.Fn5989
func Fn5989(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn5990 github.com/goccy/spidermonkeywasm2go/p3.Fn5990
func Fn5990(m *base.Module, l0 int32) int32

//go:linkname Fn5991 github.com/goccy/spidermonkeywasm2go/p4.Fn5991
func Fn5991(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5992 github.com/goccy/spidermonkeywasm2go/p4.Fn5992
func Fn5992(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5993 github.com/goccy/spidermonkeywasm2go/p4.Fn5993
func Fn5993(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5994 github.com/goccy/spidermonkeywasm2go/p4.Fn5994
func Fn5994(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn5995 github.com/goccy/spidermonkeywasm2go/p4.Fn5995
func Fn5995(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn5996 github.com/goccy/spidermonkeywasm2go/p4.Fn5996
func Fn5996(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn5999 github.com/goccy/spidermonkeywasm2go/p3.Fn5999
func Fn5999(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6000 github.com/goccy/spidermonkeywasm2go/p3.Fn6000
func Fn6000(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6001 github.com/goccy/spidermonkeywasm2go/p4.Fn6001
func Fn6001(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6002 github.com/goccy/spidermonkeywasm2go/p4.Fn6002
func Fn6002(m *base.Module, l0 int32) int32

//go:linkname Fn6003 github.com/goccy/spidermonkeywasm2go/p4.Fn6003
func Fn6003(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6004 github.com/goccy/spidermonkeywasm2go/p4.Fn6004
func Fn6004(m *base.Module, l0 int32)

//go:linkname Fn6005 github.com/goccy/spidermonkeywasm2go/p4.Fn6005
func Fn6005(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6006 github.com/goccy/spidermonkeywasm2go/p4.Fn6006
func Fn6006(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6008 github.com/goccy/spidermonkeywasm2go/p4.Fn6008
func Fn6008(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6009 github.com/goccy/spidermonkeywasm2go/p4.Fn6009
func Fn6009(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6010 github.com/goccy/spidermonkeywasm2go/p4.Fn6010
func Fn6010(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6011 github.com/goccy/spidermonkeywasm2go/p4.Fn6011
func Fn6011(m *base.Module, l0 int32) int32

//go:linkname Fn6012 github.com/goccy/spidermonkeywasm2go/p4.Fn6012
func Fn6012(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6013 github.com/goccy/spidermonkeywasm2go/p4.Fn6013
func Fn6013(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6014 github.com/goccy/spidermonkeywasm2go/p4.Fn6014
func Fn6014(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6015 github.com/goccy/spidermonkeywasm2go/p3.Fn6015
func Fn6015(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6016 github.com/goccy/spidermonkeywasm2go/p4.Fn6016
func Fn6016(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6017 github.com/goccy/spidermonkeywasm2go/p4.Fn6017
func Fn6017(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6018 github.com/goccy/spidermonkeywasm2go/p4.Fn6018
func Fn6018(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6019 github.com/goccy/spidermonkeywasm2go/p4.Fn6019
func Fn6019(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6020 github.com/goccy/spidermonkeywasm2go/p4.Fn6020
func Fn6020(m *base.Module, l0 int32) int32

//go:linkname Fn6021 github.com/goccy/spidermonkeywasm2go/p4.Fn6021
func Fn6021(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6022 github.com/goccy/spidermonkeywasm2go/p4.Fn6022
func Fn6022(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6023 github.com/goccy/spidermonkeywasm2go/p4.Fn6023
func Fn6023(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6024 github.com/goccy/spidermonkeywasm2go/p4.Fn6024
func Fn6024(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6025 github.com/goccy/spidermonkeywasm2go/p4.Fn6025
func Fn6025(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6026 github.com/goccy/spidermonkeywasm2go/p4.Fn6026
func Fn6026(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6027 github.com/goccy/spidermonkeywasm2go/p4.Fn6027
func Fn6027(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6028 github.com/goccy/spidermonkeywasm2go/p4.Fn6028
func Fn6028(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6029 github.com/goccy/spidermonkeywasm2go/p4.Fn6029
func Fn6029(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6030 github.com/goccy/spidermonkeywasm2go/p4.Fn6030
func Fn6030(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6031 github.com/goccy/spidermonkeywasm2go/p4.Fn6031
func Fn6031(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6032 github.com/goccy/spidermonkeywasm2go/p4.Fn6032
func Fn6032(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6033 github.com/goccy/spidermonkeywasm2go/p4.Fn6033
func Fn6033(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6034 github.com/goccy/spidermonkeywasm2go/p4.Fn6034
func Fn6034(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6035 github.com/goccy/spidermonkeywasm2go/p4.Fn6035
func Fn6035(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6036 github.com/goccy/spidermonkeywasm2go/p4.Fn6036
func Fn6036(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6037 github.com/goccy/spidermonkeywasm2go/p4.Fn6037
func Fn6037(m *base.Module, l0 int32) int32

//go:linkname Fn6038 github.com/goccy/spidermonkeywasm2go/p4.Fn6038
func Fn6038(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6041 github.com/goccy/spidermonkeywasm2go/p3.Fn6041
func Fn6041(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6042 github.com/goccy/spidermonkeywasm2go/p2.Fn6042
func Fn6042(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6044 github.com/goccy/spidermonkeywasm2go/p4.Fn6044
func Fn6044(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6045 github.com/goccy/spidermonkeywasm2go/p3.Fn6045
func Fn6045(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6046 github.com/goccy/spidermonkeywasm2go/p4.Fn6046
func Fn6046(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6047 github.com/goccy/spidermonkeywasm2go/p4.Fn6047
func Fn6047(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6048 github.com/goccy/spidermonkeywasm2go/p3.Fn6048
func Fn6048(m *base.Module, l0 int32, l1 int64, l2 int32)

//go:linkname Fn6049 github.com/goccy/spidermonkeywasm2go/p4.Fn6049
func Fn6049(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6050 github.com/goccy/spidermonkeywasm2go/p4.Fn6050
func Fn6050(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6051 github.com/goccy/spidermonkeywasm2go/p4.Fn6051
func Fn6051(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6052 github.com/goccy/spidermonkeywasm2go/p3.Fn6052
func Fn6052(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6053 github.com/goccy/spidermonkeywasm2go/p4.Fn6053
func Fn6053(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6054 github.com/goccy/spidermonkeywasm2go/p4.Fn6054
func Fn6054(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6055 github.com/goccy/spidermonkeywasm2go/p4.Fn6055
func Fn6055(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6056 github.com/goccy/spidermonkeywasm2go/p4.Fn6056
func Fn6056(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6057 github.com/goccy/spidermonkeywasm2go/p4.Fn6057
func Fn6057(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6058 github.com/goccy/spidermonkeywasm2go/p4.Fn6058
func Fn6058(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6059 github.com/goccy/spidermonkeywasm2go/p4.Fn6059
func Fn6059(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6060 github.com/goccy/spidermonkeywasm2go/p4.Fn6060
func Fn6060(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6061 github.com/goccy/spidermonkeywasm2go/p4.Fn6061
func Fn6061(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6062 github.com/goccy/spidermonkeywasm2go/p4.Fn6062
func Fn6062(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6063 github.com/goccy/spidermonkeywasm2go/p4.Fn6063
func Fn6063(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6064 github.com/goccy/spidermonkeywasm2go/p4.Fn6064
func Fn6064(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6067 github.com/goccy/spidermonkeywasm2go/p4.Fn6067
func Fn6067(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6072 github.com/goccy/spidermonkeywasm2go/p4.Fn6072
func Fn6072(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6078 github.com/goccy/spidermonkeywasm2go/p3.Fn6078
func Fn6078(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6079 github.com/goccy/spidermonkeywasm2go/p3.Fn6079
func Fn6079(m *base.Module, l0 int32)

//go:linkname Fn6080 github.com/goccy/spidermonkeywasm2go/p3.Fn6080
func Fn6080(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6089 github.com/goccy/spidermonkeywasm2go/p4.Fn6089
func Fn6089(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6090 github.com/goccy/spidermonkeywasm2go/p4.Fn6090
func Fn6090(m *base.Module, l0 int32)

//go:linkname Fn6091 github.com/goccy/spidermonkeywasm2go/p4.Fn6091
func Fn6091(m *base.Module, l0 int32) int32

//go:linkname Fn6092 github.com/goccy/spidermonkeywasm2go/p4.Fn6092
func Fn6092(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6093 github.com/goccy/spidermonkeywasm2go/p4.Fn6093
func Fn6093(m *base.Module, l0 int32) int32

//go:linkname Fn6098 github.com/goccy/spidermonkeywasm2go/p4.Fn6098
func Fn6098(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6099 github.com/goccy/spidermonkeywasm2go/p3.Fn6099
func Fn6099(m *base.Module)

//go:linkname Fn6100 github.com/goccy/spidermonkeywasm2go/p4.Fn6100
func Fn6100(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6102 github.com/goccy/spidermonkeywasm2go/p2.Fn6102
func Fn6102(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6106 github.com/goccy/spidermonkeywasm2go/p3.Fn6106
func Fn6106(m *base.Module, l0 int32, l1 int32, l2 int32) int64

//go:linkname Fn6107 github.com/goccy/spidermonkeywasm2go/p2.Fn6107
func Fn6107(m *base.Module, l0 int32)

//go:linkname Fn6121 github.com/goccy/spidermonkeywasm2go/p4.Fn6121
func Fn6121(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6122 github.com/goccy/spidermonkeywasm2go/p4.Fn6122
func Fn6122(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6123 github.com/goccy/spidermonkeywasm2go/p4.Fn6123
func Fn6123(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6124 github.com/goccy/spidermonkeywasm2go/p4.Fn6124
func Fn6124(m *base.Module, l0 int32) int32

//go:linkname Fn6125 github.com/goccy/spidermonkeywasm2go/p4.Fn6125
func Fn6125(m *base.Module, l0 int32) int32

//go:linkname Fn6126 github.com/goccy/spidermonkeywasm2go/p4.Fn6126
func Fn6126(m *base.Module, l0 int32)

//go:linkname Fn6128 github.com/goccy/spidermonkeywasm2go/p3.Fn6128
func Fn6128(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6129 github.com/goccy/spidermonkeywasm2go/p3.Fn6129
func Fn6129(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6130 github.com/goccy/spidermonkeywasm2go/p4.Fn6130
func Fn6130(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6131 github.com/goccy/spidermonkeywasm2go/p4.Fn6131
func Fn6131(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6132 github.com/goccy/spidermonkeywasm2go/p4.Fn6132
func Fn6132(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6133 github.com/goccy/spidermonkeywasm2go/p3.Fn6133
func Fn6133(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6135 github.com/goccy/spidermonkeywasm2go/p4.Fn6135
func Fn6135(m *base.Module, l0 int32)

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

//go:linkname Fn6159 github.com/goccy/spidermonkeywasm2go/p0.Fn6159
func Fn6159(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6184 github.com/goccy/spidermonkeywasm2go/p4.Fn6184
func Fn6184(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6186 github.com/goccy/spidermonkeywasm2go/p4.Fn6186
func Fn6186(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6187 github.com/goccy/spidermonkeywasm2go/p4.Fn6187
func Fn6187(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6188 github.com/goccy/spidermonkeywasm2go/p4.Fn6188
func Fn6188(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6189 github.com/goccy/spidermonkeywasm2go/p2.Fn6189
func Fn6189(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6190 github.com/goccy/spidermonkeywasm2go/p2.Fn6190
func Fn6190(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6191 github.com/goccy/spidermonkeywasm2go/p4.Fn6191
func Fn6191(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6192 github.com/goccy/spidermonkeywasm2go/p4.Fn6192
func Fn6192(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6216 github.com/goccy/spidermonkeywasm2go/p4.Fn6216
func Fn6216(m *base.Module, l0 int32)

//go:linkname Fn6217 github.com/goccy/spidermonkeywasm2go/p4.Fn6217
func Fn6217(m *base.Module, l0 int32)

//go:linkname Fn6229 github.com/goccy/spidermonkeywasm2go/p3.Fn6229
func Fn6229(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6236 github.com/goccy/spidermonkeywasm2go/p4.Fn6236
func Fn6236(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6264 github.com/goccy/spidermonkeywasm2go/p4.Fn6264
func Fn6264(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6284 github.com/goccy/spidermonkeywasm2go/p4.Fn6284
func Fn6284(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6286 github.com/goccy/spidermonkeywasm2go/p2.Fn6286
func Fn6286(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6287 github.com/goccy/spidermonkeywasm2go/p2.Fn6287
func Fn6287(m *base.Module, l0 int32) int32

//go:linkname Fn6302 github.com/goccy/spidermonkeywasm2go/p3.Fn6302
func Fn6302(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6303 github.com/goccy/spidermonkeywasm2go/p3.Fn6303
func Fn6303(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6319 github.com/goccy/spidermonkeywasm2go/p4.Fn6319
func Fn6319(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6320 github.com/goccy/spidermonkeywasm2go/p4.Fn6320
func Fn6320(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6321 github.com/goccy/spidermonkeywasm2go/p4.Fn6321
func Fn6321(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6322 github.com/goccy/spidermonkeywasm2go/p4.Fn6322
func Fn6322(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6323 github.com/goccy/spidermonkeywasm2go/p4.Fn6323
func Fn6323(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6325 github.com/goccy/spidermonkeywasm2go/p3.Fn6325
func Fn6325(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6326 github.com/goccy/spidermonkeywasm2go/p3.Fn6326
func Fn6326(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32

//go:linkname Fn6332 github.com/goccy/spidermonkeywasm2go/p4.Fn6332
func Fn6332(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6333 github.com/goccy/spidermonkeywasm2go/p2.Fn6333
func Fn6333(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6342 github.com/goccy/spidermonkeywasm2go/p4.Fn6342
func Fn6342(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6344 github.com/goccy/spidermonkeywasm2go/p3.Fn6344
func Fn6344(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6346 github.com/goccy/spidermonkeywasm2go/p4.Fn6346
func Fn6346(m *base.Module, l0 int32) int32

//go:linkname Fn6347 github.com/goccy/spidermonkeywasm2go/p4.Fn6347
func Fn6347(m *base.Module, l0 int32, l1 int32, l2 int64) int32

//go:linkname Fn6351 github.com/goccy/spidermonkeywasm2go/p3.Fn6351
func Fn6351(m *base.Module, l0 int32, l1 int64, l2 int32) int32

//go:linkname Fn6352 github.com/goccy/spidermonkeywasm2go/p4.Fn6352
func Fn6352(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6378 github.com/goccy/spidermonkeywasm2go/p2.Fn6378
func Fn6378(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

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

//go:linkname Fn6389 github.com/goccy/spidermonkeywasm2go/p4.Fn6389
func Fn6389(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6391 github.com/goccy/spidermonkeywasm2go/p4.Fn6391
func Fn6391(m *base.Module, l0 int32) int32

//go:linkname Fn6392 github.com/goccy/spidermonkeywasm2go/p3.Fn6392
func Fn6392(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6393 github.com/goccy/spidermonkeywasm2go/p3.Fn6393
func Fn6393(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6394 github.com/goccy/spidermonkeywasm2go/p2.Fn6394
func Fn6394(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6395 github.com/goccy/spidermonkeywasm2go/p4.Fn6395
func Fn6395(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6396 github.com/goccy/spidermonkeywasm2go/p4.Fn6396
func Fn6396(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6400 github.com/goccy/spidermonkeywasm2go/p4.Fn6400
func Fn6400(m *base.Module, l0 int32) int32

//go:linkname Fn6402 github.com/goccy/spidermonkeywasm2go/p3.Fn6402
func Fn6402(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6404 github.com/goccy/spidermonkeywasm2go/p2.Fn6404
func Fn6404(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6418 github.com/goccy/spidermonkeywasm2go/p4.Fn6418
func Fn6418(m *base.Module, l0 int32) int32

//go:linkname Fn6419 github.com/goccy/spidermonkeywasm2go/p4.Fn6419
func Fn6419(m *base.Module, l0 int32)

//go:linkname Fn6420 github.com/goccy/spidermonkeywasm2go/p4.Fn6420
func Fn6420(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6427 github.com/goccy/spidermonkeywasm2go/p3.Fn6427
func Fn6427(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6428 github.com/goccy/spidermonkeywasm2go/p3.Fn6428
func Fn6428(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6430 github.com/goccy/spidermonkeywasm2go/p4.Fn6430
func Fn6430(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6435 github.com/goccy/spidermonkeywasm2go/p4.Fn6435
func Fn6435(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6453 github.com/goccy/spidermonkeywasm2go/p3.Fn6453
func Fn6453(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6468 github.com/goccy/spidermonkeywasm2go/p4.Fn6468
func Fn6468(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6477 github.com/goccy/spidermonkeywasm2go/p4.Fn6477
func Fn6477(m *base.Module, l0 int32)

//go:linkname Fn6548 github.com/goccy/spidermonkeywasm2go/p4.Fn6548
func Fn6548(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6565 github.com/goccy/spidermonkeywasm2go/p4.Fn6565
func Fn6565(m *base.Module, l0 int32) int32

//go:linkname Fn6570 github.com/goccy/spidermonkeywasm2go/p3.Fn6570
func Fn6570(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6574 github.com/goccy/spidermonkeywasm2go/p4.Fn6574
func Fn6574(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6575 github.com/goccy/spidermonkeywasm2go/p4.Fn6575
func Fn6575(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6578 github.com/goccy/spidermonkeywasm2go/p4.Fn6578
func Fn6578(m *base.Module, l0 int32) int32

//go:linkname Fn6580 github.com/goccy/spidermonkeywasm2go/p3.Fn6580
func Fn6580(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6583 github.com/goccy/spidermonkeywasm2go/p4.Fn6583
func Fn6583(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6584 github.com/goccy/spidermonkeywasm2go/p4.Fn6584
func Fn6584(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn6587 github.com/goccy/spidermonkeywasm2go/p4.Fn6587
func Fn6587(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6588 github.com/goccy/spidermonkeywasm2go/p4.Fn6588
func Fn6588(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6590 github.com/goccy/spidermonkeywasm2go/p4.Fn6590
func Fn6590(m *base.Module, l0 int32, l1 int32, l2 int32)

//go:linkname Fn6595 github.com/goccy/spidermonkeywasm2go/p3.Fn6595
func Fn6595(m *base.Module, l0 int32) int32

//go:linkname Fn6601 github.com/goccy/spidermonkeywasm2go/p4.Fn6601
func Fn6601(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64) int32

//go:linkname Fn6602 github.com/goccy/spidermonkeywasm2go/p3.Fn6602
func Fn6602(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32

//go:linkname Fn6603 github.com/goccy/spidermonkeywasm2go/p4.Fn6603
func Fn6603(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6604 github.com/goccy/spidermonkeywasm2go/p4.Fn6604
func Fn6604(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6605 github.com/goccy/spidermonkeywasm2go/p4.Fn6605
func Fn6605(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6606 github.com/goccy/spidermonkeywasm2go/p4.Fn6606
func Fn6606(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6607 github.com/goccy/spidermonkeywasm2go/p4.Fn6607
func Fn6607(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6608 github.com/goccy/spidermonkeywasm2go/p4.Fn6608
func Fn6608(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6609 github.com/goccy/spidermonkeywasm2go/p4.Fn6609
func Fn6609(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn6610 github.com/goccy/spidermonkeywasm2go/p4.Fn6610
func Fn6610(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6611 github.com/goccy/spidermonkeywasm2go/p4.Fn6611
func Fn6611(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6612 github.com/goccy/spidermonkeywasm2go/p4.Fn6612
func Fn6612(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6613 github.com/goccy/spidermonkeywasm2go/p3.Fn6613
func Fn6613(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6614 github.com/goccy/spidermonkeywasm2go/p4.Fn6614
func Fn6614(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6615 github.com/goccy/spidermonkeywasm2go/p3.Fn6615
func Fn6615(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6616 github.com/goccy/spidermonkeywasm2go/p3.Fn6616
func Fn6616(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6617 github.com/goccy/spidermonkeywasm2go/p3.Fn6617
func Fn6617(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6618 github.com/goccy/spidermonkeywasm2go/p4.Fn6618
func Fn6618(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6619 github.com/goccy/spidermonkeywasm2go/p4.Fn6619
func Fn6619(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6620 github.com/goccy/spidermonkeywasm2go/p3.Fn6620
func Fn6620(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6621 github.com/goccy/spidermonkeywasm2go/p4.Fn6621
func Fn6621(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6622 github.com/goccy/spidermonkeywasm2go/p3.Fn6622
func Fn6622(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn6623 github.com/goccy/spidermonkeywasm2go/p3.Fn6623
func Fn6623(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6624 github.com/goccy/spidermonkeywasm2go/p3.Fn6624
func Fn6624(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6627 github.com/goccy/spidermonkeywasm2go/p4.Fn6627
func Fn6627(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6628 github.com/goccy/spidermonkeywasm2go/p4.Fn6628
func Fn6628(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6629 github.com/goccy/spidermonkeywasm2go/p3.Fn6629
func Fn6629(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6630 github.com/goccy/spidermonkeywasm2go/p4.Fn6630
func Fn6630(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6631 github.com/goccy/spidermonkeywasm2go/p2.Fn6631
func Fn6631(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6632 github.com/goccy/spidermonkeywasm2go/p3.Fn6632
func Fn6632(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6633 github.com/goccy/spidermonkeywasm2go/p3.Fn6633
func Fn6633(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6634 github.com/goccy/spidermonkeywasm2go/p4.Fn6634
func Fn6634(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6673 github.com/goccy/spidermonkeywasm2go/p3.Fn6673
func Fn6673(m *base.Module, l0 int32)

//go:linkname Fn6674 github.com/goccy/spidermonkeywasm2go/p4.Fn6674
func Fn6674(m *base.Module, l0 int32)

//go:linkname Fn6675 github.com/goccy/spidermonkeywasm2go/p2.Fn6675
func Fn6675(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6679 github.com/goccy/spidermonkeywasm2go/p3.Fn6679
func Fn6679(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6680 github.com/goccy/spidermonkeywasm2go/p2.Fn6680
func Fn6680(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6681 github.com/goccy/spidermonkeywasm2go/p3.Fn6681
func Fn6681(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6683 github.com/goccy/spidermonkeywasm2go/p2.Fn6683
func Fn6683(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32

//go:linkname Fn6685 github.com/goccy/spidermonkeywasm2go/p4.Fn6685
func Fn6685(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6686 github.com/goccy/spidermonkeywasm2go/p4.Fn6686
func Fn6686(m *base.Module, l0 int32)

//go:linkname Fn6687 github.com/goccy/spidermonkeywasm2go/p2.Fn6687
func Fn6687(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32)

//go:linkname Fn6694 github.com/goccy/spidermonkeywasm2go/p2.Fn6694
func Fn6694(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6696 github.com/goccy/spidermonkeywasm2go/p2.Fn6696
func Fn6696(m *base.Module, l0 int32, l1 int32)

//go:linkname Fn6697 github.com/goccy/spidermonkeywasm2go/p4.Fn6697
func Fn6697(m *base.Module, l0 int32, l1 int32, l2 int32) int32

//go:linkname Fn6698 github.com/goccy/spidermonkeywasm2go/p3.Fn6698
func Fn6698(m *base.Module, l0 int32, l1 int32) int32

//go:linkname Fn6699 github.com/goccy/spidermonkeywasm2go/p4.Fn6699
func Fn6699(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32)

//go:linkname Fn6714 github.com/goccy/spidermonkeywasm2go/p4.Fn6714
func Fn6714(m *base.Module, l0 int32) int32

//go:linkname Fn6715 github.com/goccy/spidermonkeywasm2go/p4.Fn6715
func Fn6715(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32

//go:linkname Fn6716 github.com/goccy/spidermonkeywasm2go/p2.Fn6716
func Fn6716(m *base.Module, l0 int32, l1 int32) float64

//go:linkname Fn6717 github.com/goccy/spidermonkeywasm2go/p3.Fn6717
func Fn6717(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64, l6 int32) float64

//go:linkname Fn6719 github.com/goccy/spidermonkeywasm2go/p4.Fn6719
func Fn6719(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32
