<!-- > go test -bench ".*"
go的基础测试用例 -->

[TOC]

## 单元测试

完整代码地址： <https://github1s.com/pengyang0317/GoLearning/blob/main/src/ch1/ch1_test.go>

### 1. 为什么要编写测试用例

提高代码质量：编写单元测试可以帮助发现代码中的潜在问题和错误，提高代码的质量和稳定性。

简化代码调试：当代码出现问题时，单元测试可以帮助快速定位问题，并且可以在修改代码后，快速验证修改是否解决了问题。

支持重构：当需要修改现有代码时，编写单元测试可以帮助确保修改不会破坏现有的代码功能。

提高开发效率：编写单元测试可以帮助开发人员更快地理解代码的功能和实现方式，减少开发和调试时间。

改善团队协作：编写单元测试可以帮助团队成员更好地理解代码实现、功能和接口，提高团队协作效率。

总之，编写单元测试可以提高代码质量、简化代码调试、支持重构、提高开发效率和改善团队协作等方面的好处，是现代软件开发过程中必不可少的一环。

### 2.单元测试编码

go test命令会自动查找当前目录下以_test.go结尾的文件，并运行其中的测试函数。测试函数的命名必须以Test开头，并且第一个字母必须大写。如果所有测试都通过，则输出PASS信息，否则输出错误信息。

```
func TestAdd(t *testing.T) {
 if got := Add(1, 2); got != 3 {
  t.Errorf("Add(1,2) = %d; want 3", got)
 }
}
```

### 3.跳过耗时单元测试

在 Go 语言中，可以使用 testing 包提供的 t.Skip() 函数来跳过单元测试。

当某个单元测试因为某些原因需要跳过时，可以在测试函数中调用 t.Skip() 函数，该函数会输出一个消息，告诉测试运行器该测试被跳过，而不是失败。

```
func TestSkip(t *testing.T) {
 if t.Short() {
  t.Skip("skipping test in short mode.")
 }
 if got := Add(1, 2); got != 3 {
  t.Errorf("Add(1,2) = %d; want 3", got)
 }
}

// go test -short

```

### 4.基于表格驱动测试

在 Go 语言中，可以使用表格驱动测试来编写简洁、易于维护的测试代码。表格驱动测试的基本思想是将测试数据和预期结果放在一个表格中，然后在测试函数中迭代这个表格，分别执行每组测试数据，并验证测试结果是否符合预期。

```
func TestTableAdd(t *testing.T) {
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
```

### 5.benchmark性能测试编码

Go语言的测试框架testing中也支持基准测试（benchmark testing），用于测试函数的性能和比较不同实现的性能差异。基准测试的代码与单元测试类似，但需要使用Benchmark函数来定义测试函数，以及使用b.N来表示测试的次数。
go test -bench=.命令会运行所有基准测试函数，并输出测试结果。
通常情况下，测试次数越多，测试结果越准确。

```
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
```
