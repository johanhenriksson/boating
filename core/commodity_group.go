package core

var PRECIOUS_METALS = NewCommodityGroup(1001, "Precious Metals", GOLD, SILVER)
var METALS = NewCommodityGroup(1000, "Metals", PRECIOUS_METALS, GOLD, SILVER)

type CommodityGroup struct {
    Name    string
    Members map[ComType]Tradable
    ctype   ComType
}

func NewCommodityGroup(ctype ComType, name string, includes ...Tradable) *CommodityGroup {
    cg := &CommodityGroup {
        ctype: ctype,
        Name: name,
        Members: make(map[ComType]Tradable),
    }
    for _, com := range includes {
        if com.Is(cg.Type()) {
            continue
        }
        cg.Include(com)
    }
    return cg
}

func (cg *CommodityGroup) Include(com Tradable) *CommodityGroup {
    cg.Members[com.Type()] = com
    return cg
}

func (cg *CommodityGroup) Type() ComType {
    return cg.ctype
}

func (cg *CommodityGroup) Is(ctype ComType) bool {
    for _, com := range cg.Members {
        if com.Type() == ctype {
            return true
        }
    }
    return false
}

