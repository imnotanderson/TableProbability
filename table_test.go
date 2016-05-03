package TableProbability

import (
	"fmt"
	"testing"
)

func TestNewTable(t *testing.T) {
	table := NewTable()
	table.AddData(&MyData{
		fVal: 1,
		str:  "a",
	})
	table.AddData(&MyData{
		fVal: 1,
		str:  "b",
	})
	data := table.GetRandom().(*MyData)
	fmt.Println(data.str)
}

type MyData struct {
	fVal float32
	str  string
}

func (data *MyData) GetfVal() float32 {
	return data.fVal
}
