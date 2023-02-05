package goconvey、

FAQ
# What is GoConvey?
基本上，GoConvey 是内置 Go 测试工具的扩展。它促进了 Go 中的行为驱动开发 (BDD)，尽管这不是使用它的唯一方法。许多人继续编写传统的 Go 测试，
但更喜欢 GoConvey 的 Web UI 来报告测试结果。
GoConvey 有两个主要部分：
	1.全面的 BDD 框架
	2.服务器和网络用户界面
这两个部分都是可选的，可以根据您的工作流程独立使用。

# 我可以将 GoConvey 与 go test 一起使用吗？
是的，这就是重点！

# 为什么我的嵌套 Convey 块以奇怪的顺序执行？
请阅读隔离执行测试。它们是 GoConvey 执行模型的最佳文档形式，它非常强大但一开始并不明显。

# What is the web UI?
这是您浏览器中的测试结果。当监视目录中的文件被保存或更改时，它会自动更新。有关详细信息，请参阅 Web UI wiki 页面。

# 如何让测试自动运行？
如果您使用带有 goconvey（服务器）的 Web UI 在浏览器中观看测试，那么当 .go 文件更改时，测试已经自动运行。
要在终端中运行测试，请查看如何使用 auto-run.py 自动测试脚本。 （这真的很容易！）

# Web UI 是否适用于传统的 Go 测试？
是的！如果您还没有将所有测试移植到 GoConvey，您的传统 Go 测试用例仍将运行并且结果将在浏览器中报告。

# 我怎样才能让调试输出出现在我的断言旁边，而不是出现在函数级别？
使用 convey.Print 或 convey.Printf 或 convey.Println 就像在 fmt 包中一样。这将导致输出显示在您的断言中，而不是显示在更高级别。

# 我怎样才能强制测试在失败后继续执行？
默认情况下，测试失败或恐慌会导致该范围内的未来测试停止。要让测试继续运行，您可以在 Convey() 调用中传入 FailureMode：
Convey("A", t, FailureContinues, func() {
	// ...
})
所有嵌套的 Conveys 都将继承该设置。您还可以使用 SetDefaultFailureMode() 在 init() 函数中设置默认的 FailureMode，如下所示：
func init() {
	SetDefaultFailureMode(FailureContinues)
}

# 为什么我无法打开测试文件的覆盖率报告（404 未找到）？
你必须确保你正在测试的包存在于你的 $GOPATH 中。

# 是否支持 GoConvey？
不是商业意义上的，不。尽管它是由 SmartyStreets“赞助”的（因为一些开发人员有公司时间来处理它），它还是按原样出现，所以“使用风险自负”——尽管老
实说你您可能会非常喜欢它。:)
但是，是的，在开源这个词的意义上，它是受支持的，意思是：它不是一个废弃的项目。它非常活跃。随时提交带有贡献的请求请求！

# GoConvey 从哪里开始？
当 SmartyStreets 决定放弃其 .NET 代码（几乎是所有代码）以支持 Go 时，所有开发人员都开始学习和实践 Go。虽然 Go 的内置 go test 命令和标准
测试包让我们兴奋，但实际上它让我们缺乏。

在我们用 Go 重写第一个项目后，使用标准库测试，我们发现我们的测试没有清楚地记录我们的代码在做什么，以及它应该如何表现。断言不明确，因为它们是落后的。

最重要的是（由其他更轻量级的库解决），我们认为带有可视化显示测试结果的浏览器选项卡真的很酷，而且听起来很有趣。所以最初，GoConvey 是为 SmartyStreets
内部使用而编写的。然后我们决定把它作为我们送给 Golang 社区的礼物。

# 你会添加这个或那个功能吗？
也许吧，但 GoConvey 对我们在 SmartyStreets 的现状来说已经足够好了。欢迎您这样做。

# 我如何贡献？
请参阅我们的贡献者页面。

# 如果我是测试新手，您推荐阅读哪些内容？
Unit Testing								单元测试
Test-Driven Development (TDD)				测试驱动开发 (TDD)
Behavior-Driven Development (BDD)			行为驱动开发 (BDD)
BDD vs. TDD
Laws of TDD									TDD 法则
Integration Testing							集成测试

# 我们为什么要创造这个？ （go test 还不够好吗？）
我们对内置的 GoLang 测试工具不满意。不，实际上我们很高兴语言带有一些内置的东西。我们非常喜欢 go test，所以我们决定直接与 go test 集成，而
不是从头开始创建一些东西。我们只是习惯了更具描述性的东西，这有助于使用大量样板代码测试大型系统。

# 为什么叫 GoConvey？
我们之前使用过几种不同的 BDD 工具，每种工具都有自己的语言，您应该使用它来指定被测系统的行为。 BDD 中的“Given, When, Then” 与 BDD 中的“
Establish, Because, It” 和 TDD 中的“Arrange, Act, Assert” 是几个例子。最后，您使用测试工具来指定或“传达(convey)”系统应该做什么。因
此，您将用于编写规范的主要函数名为 Convey。用于指定系统的语言由您决定，尽管我们通常使用“Given, When, Then”风格。

# 我应该小心不要在我的代码中产生某些输出吗？
好吧，如果我听说过的话，这是一个晦涩的问题，但我很高兴你问了。永远不要输出以下任何一种模式：
	>->->OPEN-JSON->->->
	<-<-<-CLOSE-JSON<-<-<
这些模式用于分隔 JSON 块，以便 Web 服务器可以正确解析测试输出。
