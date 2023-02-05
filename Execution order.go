package goconvey

# Execution order
作为 FAQ 的扩展，本文档回答了一些关于 GoConvey 执行模型的重要问题，例如：
	1.如何定义在每次测试之前运行的"Setup"方法？
	2.为什么我的嵌套测试不按顺序运行？

令人惊讶的是，这些问题是相关的。这是 GoConvey 的一位用户的雄辩解释，阐明了这些问题：
例如，考虑伪代码：
	Convey A
		So 1
		Convey B
			So 2
		Convey C
			So 3

我最初认为会按 A1B2C3 执行，换句话说，按顺序执行。大家知道，这其实是先执行A1B2，再执行A1C3。意识到这一点后，我才真正意识到 goconvey 的强大
功能，因为它允许您使用 log(n) 语句编写两个 n 测试。这种“基于树的”行为测试消除了如此多的重复设置代码，并且更容易阅读完整性（相对于单元测试页
面），同时仍然允许非常好的隔离测试（对于树中的每个分支）。

在上面的伪代码中，Convey A 充当 Convey B 和 Convey C 的"Setup"方法，并且分别运行。这是一个更复杂的例子：
Convey A
	So 1
	Convey B
		So 2
		Convey Q
			So 9
	Convey C
		So 3

你能猜出输出是什么吗？
A1B2Q9A1C3 是正确答案。
欢迎您仔细阅读 GoConvey 项目本身中记录此行为的测试。

## Gotchas
请记住，Go 中的每个 Convey() 调用都会创建一个新的范围。您应该使用 Foo = &Bar{} 以便为先前声明的变量分配新值。使用 foo := &Bar{} 在当前
范围内创建一个新变量。例子：
Convey("Setup", func() {
	foo := &Bar{}
	Convey("This creates a new variable foo in this scope", func() {
		foo := &Bar{}
	}
	Convey("This assigns a new value to the previous declared foo", func() {
		foo = &Bar{}
	}
}

如果您想知道如何实现“拆卸”功能，请参阅重置功能页面。
