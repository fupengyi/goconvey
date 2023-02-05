package goconvey

# Custom Assertions
有时，测试套件可能需要一个过于具体而无法包含在通用存储库中的断言。不用担心，只需使用以下签名实现一个函数（替换括号中的部分和字符串值）：
func should<do-something>(actual interface{}, expected ...interface{}) string {
	if <some-important-condition-is-met(actual, expected)> {
		return ""   // empty string means the assertion passed						// 空字符串表示断言通过
	}
	return "<some descriptive message detailing why the assertion failed...>"
}

假设我实现了以下断言：
func shouldScareGophersMoreThan(actual interface{}, expected ...interface{}) string {
	if actual == "BOO!" && expected[0] == "boo" {
		return ""
	}
	return "Ha! You'll have to get a lot friendlier with the capslock if you want to scare a gopher!"
}

然后我可以在测试中调用 So() 方法时使用断言函数：
Convey("All caps always makes text more meaningful", func() {
	So("BOO!", shouldScareGophersMoreThan, "boo")
})

## Next
如果您还不知道如何操作，那么是时候学习如何运行您的测试了。
