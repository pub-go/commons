# code.gopub.tech/commons

```bash
go get code.gopub.tech/commons@latest
```

```go
// code.gopub.tech/commons/arg
var xxx = ...
t.Logf("xxx = %v", arg.JSON(xxx))

// code.gopub.tech/commons/assert
assert.Equal(t, 1, 1)
assert.DeepEqual(t, []int{1}, []int{1})
assert.True(t, values.IsNotNil(xxx))
assert.False(t, values.IsNil(xxx))
assert.Nil(t, xxx)
assert.NotNil(t, xxx)
assert.ShouldNotPanic(t, func(){})
assert.ShouldPanic(t, func(){panic("xxx")})

// code.gopub.tech/commons/conv
conv.Bytes2String([]byte(`abc`)) == "abc"
conv.String2ReadOnlyBytes("abc") // []byte(`abc`)

// code.gopub.tech/commons/choose
s := choose.If(true, "T", "F")
assert.Equal(t, s, "T")
choose.IfLazy(bool, onTrue, onFalse funcs.Supplier[T])
choose.IfLazyT(bool, onTrue funcs.Supplier[T], onFalse T)
choose.IfLazyF(bool, onTrue T, onFalse funcs.Supplier[T])

// code.gopub.tech/commons/funcs
type (
	// Supplier 产生一个元素
	Supplier[T any] func() T
	// Consumer 消费一个元素
	Consumer[T any] func(T)
	// Function 将一个类型转为另一个类型
	Function[T, R any] func(T) R
	// Predicate 断言是否满足指定条件
	Predicate[T any] Function[T, bool]
	// UnaryOperator 对输入进行一元运算返回相同类型的结果
	UnaryOperator[T any] Function[T, T]
	// BiFunction 将两个类型转为第三个类型
	BiFunction[X, Y, R any] func(X, Y) R
	// BinaryOperator 输入两个相同类型的参数，对其做二元运算，返回相同类型的结果
	BinaryOperator[T any] BiFunction[T, T, T]
)

func TestIdentidy(t *testing.T) {
	assert.Equal(t, funcs.Identidy(1), 1)
}

func TestNot(t *testing.T) {
	assert.True(t, values.IsZero(0))
	assert.True(t, funcs.Not(values.IsZero[int])(1))
}

func TestPartial(t *testing.T) {
	format10to := funcs.Partial2(strconv.FormatInt, 10)
	assert.Equal(t, format10to(10), "10")
	assert.Equal(t, format10to(2), "1010")
}

// code.gopub.tech/commons/iters
func TestSeq(t *testing.T) {
	iters.Range(0, 10).
		Peek(func(i int) {
			t.Logf("from source: %d", i)
		}).
		Filter(func(i int) bool { return i%2 == 0 }).
		Peek(func(i int) {
			t.Logf("after filter: %d", i)
		}).
		Map(func(i int) int { return 2 * i }).
		Peek(func(i int) {
			t.Logf("map to *2: %d", i)
		}).
		FlatMap(func(i int) iter.Seq[int] {
			return iters.Repeat(i).Limit(2).Seq()
		}).
		Peek(func(i int) {
			t.Logf("after flatten: %d", i)
		}).
		Distinct(funcs.Identidy).
		Sorted(order.Reversed[int]).
		Skip(1).
		ForEach(func(i int) {
			t.Logf("got: %v", i)
		})
}

// code.gopub.tech/commons/jsons
t.Logf("xxx=%s", jsons.ToJSON(xxx))
// 区别于 arg.JSON(xxx): 格式化时(调用到String函数时)才会 to json
logs.Debug(ctx, "xxx=%v", arg.JSON(xxx))
// 即使 logs level 更高, 不打印 Debug 日志, 也会执行 to json
logs.Debug(ctx, "xxx=%v", jsons.JSON(xxx))

// code.gopub.tech/commons/nums
type(
    Signed,
    Unsigned,
    Int,
    Float,
    Complex,
    Number
)
nums.To[float64](MyInt(10)) == float64(10)

// code.gopub.tech/commons/order
type: order.Comparator
order.Natural
order.Reversed
order.Reverse

// code.gopub.tech/commons/values
tuple := values.Make2("a", 2)
tuple.Val1 == "a"
tuple.Val2 == 2
values.IsNotZero(tuple.Val1)

// code.gopub.tech/commons/values
values.IsNil(xx)
values.IsNotNil(xx)
values.Zero[string]() == ""
values.IsZero("") == true

```
