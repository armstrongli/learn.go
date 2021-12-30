package main

import (
	"math"
	"sort"
)

type RandItem struct {
	Name    string
	FatRate float64
}

type FatRateRank struct {
	items []RandItem
}

func (r *FatRateRank) inputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		if minFatRate > item {
			minFatRate = item
		}
	}

	found := false
	for i, item := range r.items {
		if item.Name == name {
			if item.FatRate >= minFatRate {
				item.FatRate = minFatRate
			}
			r.items[i] = item
			found = true
			break
		}
	}
	if !found {
		r.items = append(r.items, RandItem{
			Name:    name,
			FatRate: minFatRate,
		})
	}
}

func (r *FatRateRank) getRank(name string) (rank int, fatRate float64) {
	sort.Slice(r.items, func(i, j int) bool {
		return r.items[i].FatRate < r.items[j].FatRate
	})
	frs := map[float64]struct{}{}
	for _, item := range r.items {
		frs[item.FatRate] = struct{}{}
		if item.Name == name {
			fatRate = item.FatRate
		}
	}
	rankArr := make([]float64, 0, len(frs))
	for k := range frs {
		rankArr = append(rankArr, k)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		if frItem == fatRate {
			rank = i + 1
			break
		}
	}

	// for nameItem, frItem := range personFatRate {
	// 	fatRate2PersonMap[frItem] = append(fatRate2PersonMap[frItem], nameItem)
	// 	rankArr = append(rankArr, frItem)
	// }
	// sort.Float64s(rankArr)
	// for i, frItem := range rankArr {
	// 	_names := fatRate2PersonMap[frItem]
	// 	for _, _name := range _names {
	// 		if _name == name {
	// 			rank = i + 1
	// 			fatRate = frItem
	// 			return
	// 		}
	// 	}
	// }
	return
}
