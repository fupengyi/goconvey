package goconvey

# Assertions
GoConvey 附带了许多标准断言，您可以将其与 So() 一起使用。

## General Equality
So(thing1, ShouldEqual, thing2)			// 应该等于
So(thing1, ShouldNotEqual, thing2)		// 应该not等于
So(thing1, ShouldResemble, thing2)		// 应该相似		// a deep equals for arrays, slices, maps, and structs
So(thing1, ShouldNotResemble, thing2)	// 应该not相似	// 数组、切片、映射和结构的深度相等
So(thing1, ShouldPointTo, thing2)		// 应该指向
So(thing1, ShouldNotPointTo, thing2)	// 应该not指向
So(thing1, ShouldBeNil)					// 应该为零
So(thing1, ShouldNotBeNil)				// 应该not为零
So(thing1, ShouldBeTrue)				// 应该是真的
So(thing1, ShouldBeFalse)				// 应该not是真的
So(thing1, ShouldBeZeroValue)			// 应该是零值

## Numeric Quantity comparison
So(1, ShouldBeGreaterThan, 0)					// 应该大于
So(1, ShouldBeGreaterThanOrEqualTo, 0)			// 应该大于或等于
So(1, ShouldBeLessThan, 2)						// 应该小于
So(1, ShouldBeLessThanOrEqualTo, 2)				// 应该小于或等于
So(1.1, ShouldBeBetween, .8, 1.2)				// 应该介于
So(1.1, ShouldNotBeBetween, 2, 3)				// 不应该介于两者之间
So(1.1, ShouldBeBetweenOrEqual, .9, 1.1)		// 应该介于或等于
So(1.1, ShouldNotBeBetweenOrEqual, 1000, 2000)	// 不应该介于或等于
So(1.0, ShouldAlmostEqual, 0.99999999, .0001)   // 应该几乎相等		// tolerance is optional; default 0.0000000001
So(1.0, ShouldNotAlmostEqual, 0.9, .0001)		// 不应该几乎相等		// 公差是可选的；默认 0.0000000001

## Collections
So([]int{2, 4, 6}, ShouldContain, 4)						// 应该包含
So([]int{2, 4, 6}, ShouldNotContain, 5)						// 应该not包含
So(4, ShouldBeIn, ...[]int{2, 4, 6})						// 应该在
So(4, ShouldNotBeIn, ...[]int{1, 3, 5})						// 应该not在
So([]int{}, ShouldBeEmpty)									// 应该为空(切片)
So([]int{1}, ShouldNotBeEmpty)								// 应该not为空(切片)
So(map[string]string{"a": "b"}, ShouldContainKey, "a")		// 应该包含Key
So(map[string]string{"a": "b"}, ShouldNotContainKey, "b")	// 应该not包含Key
So(map[string]string{"a": "b"}, ShouldNotBeEmpty)			// 应该not为空(映射)
So(map[string]string{}, ShouldBeEmpty)						// 应该为空(映射)
So(map[string]string{"a": "b"}, ShouldHaveLength, 1) 		// 应该有长度 supports map, slice, chan, and string

## Strings
So("asdf", ShouldStartWith, "as")				// 应该开始于
So("asdf", ShouldNotStartWith, "df")			// 应该not开始于
So("asdf", ShouldEndWith, "df")					// 应该结束于
So("asdf", ShouldNotEndWith, "df")				// 应该not结束于
So("asdf", ShouldContainSubstring, "sd")		// 应该包含子串 // 可选的“预期发生”参数？optional 'expected occurences' arguments?
So("asdf", ShouldNotContainSubstring, "er")		// 应该not包含子串
So("adsf", ShouldBeBlank)						// 应该是空白
So("asdf", ShouldNotBeBlank)					// 应该not是空白

## panic
So(func(), ShouldPanic)				// 应该恐慌
So(func(), ShouldNotPanic)			// 应该not恐慌
So(func(), ShouldPanicWith, "")		// or errors.New("something")
So(func(), ShouldNotPanicWith, "")	// or errors.New("something")

## Type checking
So(1, ShouldHaveSameTypeAs, 0)				// 应该具有相同的类型作
So(1, ShouldNotHaveSameTypeAs, "asdf")		// 应该not具有相同的类型

## time.Time (and time.Duration)
So(time.Now(), ShouldHappenBefore, time.Now())							// 应该发生在之前
So(time.Now(), ShouldHappenOnOrBefore, time.Now())						// 应该发生在或之前
So(time.Now(), ShouldHappenAfter, time.Now())							// 应该发生在之后
So(time.Now(), ShouldHappenOnOrAfter, time.Now())						// 应该发生在或之后
So(time.Now(), ShouldHappenBetween, time.Now(), time.Now())				// 应该发生在...中间
So(time.Now(), ShouldHappenOnOrBetween, time.Now(), time.Now())			// 应该发生在或在...中间
So(time.Now(), ShouldNotHappenOnOrBetween, time.Now(), time.Now())		// 应该not发生在或在...中间
So(time.Now(), ShouldHappenWithin, duration, time.Now())				// 应该发生在...时间段
So(time.Now(), ShouldNotHappenWithin, duration, time.Now())				// 应该not发生在...时间段

感谢 github.com/jacobsa 提供了出色的 oglematchers 库，这些方法中的许多方法都利用它来完成它们的工作。

## Next
接下来，了解如何构建您自己的断言(building your own assertions)。
