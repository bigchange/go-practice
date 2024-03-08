package leetcode

import "testing"

func TestStockSpanner_Next(t *testing.T) {
	type args struct {
		price []int
	}
	this := &StockSpanner{
		stack:   []Item901{{-1, 1e5 + 1}}  ,
		NextIndex: 0,
	}
	tests := []struct {
		name   string
		args   args
		want   []int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{100, 80, 60, 70, 60, 75, 85}}, []int{1, 1, 1, 2, 1, 4, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, price := range tt.args.price {
				if got := this.Next(price); got != tt.want[i] {
					t.Errorf("Next() = %v, want %v", got, tt.want[i])
				}
			}

		})
	}
}
