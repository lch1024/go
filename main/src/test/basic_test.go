package test

import "testing"

// 命令行执行
// 1. go test . 测试当前目录下文件
// 2. go test -coverprofile=c.out  测试代码覆盖率输出到c.out文件
// 3. go tool cover -html c.out 将c.out文件已html方式 输出出来
// 4. go test -bench . 对当前文件夹下 文件进行性能测试
// 5. go test -bench . -cpuprofile cpu.out cpu使用率文件
// 6. go tool pprof cpu.out 查看cpu.out二进制文件 进入bin 使用web方式查看

// 测试
// Test开头 接下来字母要大写 使用模板测试
// Benchmark开头 接下来字母大写  测试性能

// 测试Test开头 接下来字母要大写
func TestTriangle(t *testing.T) {
	tests := []struct {a, b, c int} {
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	for _, test := range tests {
		if actual := calcTriangle(test.a, test.b); actual != test.c {
			t.Errorf("calcTriangle(%d, %d); got %d; expected %d", test.a, test.b, actual, test.c)
		}
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct{
		s string
		ans int
	} {
		{"abcabcbb", 3},
		{"bbbbbb", 1},
		{"pwwkew", 3},

		{"", 0},
		{"我爱敲代码!", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, test := range tests {
		if actual := lengthOfLongestSubstring(test.s); actual != test.ans {
			t.Errorf("got %d for input %s expect %d", actual, test.s, test.ans)
		}
	}
}

// 性能测试Benchmark开头 接下来字母要大写
func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8
	for i := 0; i < 13; i++ {
		s += s
	}

	// 准备的时间 重置 只关注for循环内
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actual := lengthOfLongestSubstring(s)
		if actual != ans {
			b.Errorf("got %d for input %s expect %d", actual, s, ans)
		}
	}
}
