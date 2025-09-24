package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id string
	private bool
	createdAt time.Time
	name string
}
func NewBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{
		id: id,
		private: private,
		createdAt: createdAt,
		name: name,
	}
}

type BinList struct {
	List []*Bin
}
func (list *BinList) addBin(bin *Bin) *BinList {
	list.List = append(list.List, bin)
	return list
}

func main() {
	var temp = NewBin("123", true, time.Now(), "abraham")
	fmt.Println("1. BIN: ", temp)
	var gage = BinList{
		List: []*Bin{temp},
	}

	fmt.Println("2. BINLIST: ", gage)
	gage.addBin(temp)
		fmt.Println("3. BINLIST: ", gage)
}