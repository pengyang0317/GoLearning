package ch1

import (
	"fmt"
	"strings"
	"testing"
)

// 表格驱动测试
func TestAdd(t *testing.T) {
	var tests = []struct {
		ia, ib, want int
	}{
		{ia: 1, ib: 2, want: 3},
		{ia: 4, ib: 5, want: 9},
		{ia: 2, ib: 5, want: 7},
	}

	for _, tt := range tests {
		if got := Add(tt.ia, tt.ib); got != tt.want {
			t.Errorf("Add(%d, %d) = %d; want %d", tt.ia, tt.ib, got, tt.want)
		}
	}

}

// 性能测试
func BenchmarkAdd(b *testing.B) {
	var ia, ib, ic int
	ia = 3
	ib = 4
	ic = 7
	for i := 0; i < b.N; i++ {
		if got := Add(ia, ib); got != ic {
			fmt.Printf("Add(%d,%d) = %d; want %d", ia, ib, ic, got)
		}

	}
}

const numbers = 10000

// Sprintf性能测试
func BenchmarkSprintf(b *testing.B) {
	// 重置计时器
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = fmt.Sprintf("%s%d", str, j)
		}
	}
	b.StopTimer()
}

// StringAdd性能测试
func BenchmarkStringAdd(b *testing.B) {
	// 重置计时器
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = str + string(rune(j))
		}

	}
	b.StopTimer()
}

// StringBuilder性能测试
func BenchmarkStringBuilder(b *testing.B) {
	// 重置计时器
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < numbers; j++ {
			builder.WriteString(string(rune(j)))
		}
	}
	b.StopTimer()
}
