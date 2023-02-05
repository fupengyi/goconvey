package goconvey

# Profiles
## GoConvey Test Package Profiles
使用 Web UI 时，能够自定义在特定包中运行 go test 时使用的标志会很好*。以某种方式将包标记为永久忽略或禁用也很好，这样当 UI 在将来某个时间启
动时，您不必再次单击忽略该包。好吧，有一种方法可以通过在该包中创建一个满足以下正则表达式的文本文件来完成这两件事：
	.+\.goconvey

好名字的例子：
example.goconvey
hi.goconvey

坏名字的例子：
.goconvey
hi.txt

该文本文件的内容可能包括空白行、注释行（以# 或// 开头的行）、测试标志（-short 等）或单词ignore（大小写不重要）。唯一硬性规定是，如果你包含
单词 ignore，它必须出现在任何测试标志之前。

创建这样一个文件的结果是，如果发现 ignore 作为第一个非注释、非空行，GoConvey 服务器将忽略该包，直到该行被删除或注释。否则，所有非空白、非注
释行都将被视为在对包调用 go test 时使用的参数（例如保存文件时）。大多数这些参数可能是由 golang 测试命令定义的参数，但您可以在包中包含任意参
数。只有几个注意事项需要注意：
	1.-v 总是被使用，所以没有必要包含这个标志。
	2.-cover、-covermode 和 -coverprofile 由 GoConvey 服务器指定，因此包括这些标志将无效。 （让我知道您是否真的希望能够为 set 以外的
值指定 -covermode，这是 GoConvey 默认使用的值。）
	3.-tags 将被适当地传递给 go test、go list 等调用，所以如果你需要一个特殊的标签来构建你的包或你的测试，它应该被正确地拾取。
	4.goconvey 核心开发人员并没有真正使用很多分析/基准标记标志，因此我们不确定它们在应用于 goconvey 包配置文件时是否会像往常一样工作。

远不是一个详尽的清单，这里有一些预期的用例：
	1.使用 -short 标志可以让您切换长时间运行的集成测试的执行。
	2.使用 -run 标志可以让您专注于包中的特定测试功能。结合 FocusConvey，您可以将测试执行限制在单个测试函数中的单个测试用例——如果显示大量调
试信息，这很方便。
	3.使用 -timeout 标志可能有助于防止在错误无限循环的情况下手动重启 goconvey 服务器（来吧，承认它——你知道它发生在你身上！）。

请参阅示例包 (github.com/smartystreets/goconvey/examples) 中的 examples.goconvey 配置文件，作为示例。测试愉快！
配置文件仅适用于包含包（同一文件夹），不适用于任何嵌套包（子文件夹）。

// Uncomment the next line to disable the package when running the GoConvey UI:	在运行 GoConvey UI 时禁用包：
IGNORE

// Uncomment the next line to limit testing to the specified test function name pattern: 将测试限制为指定的测试函数名称模式：
-run=TestAssertionsAreAvailableFromConveyPackage

// Uncomment the next line to limit testing to those tests that don't bail when testing.Short() is true: 当 testing.Short() 为真时不退出的测试：
-short

// include any additional `go test` flags or application-specific flags below: // 在下面包含任何额外的 `go test` 标志或特定于应用程序的标志：

-timeout=1s
