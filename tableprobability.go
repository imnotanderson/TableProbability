package TableProbability

import (
	"fmt"
	"math/rand"
)

type Table struct {
	dataList []TableItem
	isSort   bool
	fMax     float32
}

type TableItem struct {
	data TableData
	fVal float32
}

type TableData interface {
	GetfVal() float32
}

func NewTable() *Table {
	return &Table{
		dataList: []TableItem{},
		isSort:   false,
		fMax:     0,
	}
}

func (t *Table) AddData(data TableData) {
	t.isSort = false
	t.dataList = append(t.dataList, TableItem{
		data: data,
		fVal: 0,
	})
}

func (t *Table) Sort() {
	t.fMax = 0
	for i, v := range t.dataList {
		t.fMax += v.data.GetfVal()
		v.fVal = t.fMax
		t.dataList[i] = v
	}
	t.isSort = true
}

func (t *Table) GetRandom() TableData {
	if t.isSort == false {
		t.Sort()
	}
	fVal := rand.Float32() * t.fMax
	for _, v := range t.dataList {
		if v.fVal > fVal {
			return v.data
		}
	}
	fmt.Println("err :GetRandom nil")
	return nil
}
