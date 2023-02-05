package goconvey

import (
	"os"
	"testing"
)

# Overview ¶
包 convey 包含该项目的所有面向公众的入口点。这意味着永远不需要用户从该项目导入任何其他包，因为它们用于内部目的。

# Index ¶
Constants
Variables
func Convey(items ...interface{})
func FocusConvey(items ...interface{})
func Print(items ...interface{}) (written int, err error)
func PrintConsoleStatistics()
func Printf(format string, items ...interface{}) (written int, err error)
func Println(items ...interface{}) (written int, err error)
func Reset(action func())
func SetDefaultFailureMode(mode FailureMode)
func SetDefaultStackMode(mode StackMode)
func SkipConvey(items ...interface{})
func SkipSo(stuff ...interface{})
func So(actual interface{}, assert Assertion, expected ...interface{})
func SoMsg(msg string, actual interface{}, assert Assertion, expected ...interface{})
func SuppressConsoleStatistics()
type Assertion
type C
type FailureMode
type StackMode

# Constants ¶
const (
	
	// FailureContinues is a failure mode which prevents failing				// FailureContinues 是一种防止失败的失败模式
	// So()-assertions from halting Convey-block execution, instead				// So()-来自停止 Convey-block 执行的断言，而不是
	// allowing the test to continue past failing So()-assertions.				// 允许测试在失败的 So() 断言之后继续。
	FailureContinues FailureMode = "continue"
	
	// FailureHalts is the default setting for a top-level Convey()-block		// FailureHalts 是顶层 Convey() 块的默认设置
	// and will cause all failing So()-assertions to halt further execution		// 并且会导致所有失败的 So() 断言停止进一步执行
	// in that test-arm and continue on to the next arm.						// 在那个测试臂中并继续到下一个臂。
	FailureHalts FailureMode = "halt"
	
	// FailureInherits is the default setting for failure-mode, it will			// FailureInherits 是失败模式的默认设置，它将
	// default to the failure-mode of the parent block. You should never		// 默认为父块的失败模式。你永远不应该
	// need to specify this mode in your tests..								// 需要在你的测试中指定这个模式..
	FailureInherits FailureMode = "inherits"
	
	// StackError is a stack mode which tells Convey to print stack traces		// StackError 是一种堆栈模式，它告诉 Convey 打印堆栈跟踪
	// only for errors and not for test failures								// 仅针对errors而不针对测试失败
	StackError StackMode = "error"
	
	// StackFail is a stack mode which tells Convey to print stack traces		// StackFail 是一种堆栈模式，它告诉 Convey 打印堆栈跟踪
	// for both errors and test failures										// 对于errors和测试失败
	StackFail StackMode = "fail"
	
	// StackInherits is the default setting for stack-mode, it will				// StackInherits 是堆栈模式的默认设置，它将
	// default to the stack-mode of the parent block. You should never			// 默认为父块的堆栈模式。你永远不应该
	// need to specify this mode in your tests..								// 需要在你的测试中指定这个模式..
	StackInherits StackMode = "inherits"
)

# Variables ¶
var (
	ShouldEqual          = assertions.ShouldEqual
	ShouldNotEqual       = assertions.ShouldNotEqual
	ShouldAlmostEqual    = assertions.ShouldAlmostEqual
	ShouldNotAlmostEqual = assertions.ShouldNotAlmostEqual
	ShouldEqualJSON      = assertions.ShouldEqualJSON
	ShouldResemble       = assertions.ShouldResemble
	ShouldNotResemble    = assertions.ShouldNotResemble
	ShouldPointTo        = assertions.ShouldPointTo
	ShouldNotPointTo     = assertions.ShouldNotPointTo
	ShouldBeNil          = assertions.ShouldBeNil
	ShouldNotBeNil       = assertions.ShouldNotBeNil
	ShouldBeTrue         = assertions.ShouldBeTrue
	ShouldBeFalse        = assertions.ShouldBeFalse
	ShouldBeZeroValue    = assertions.ShouldBeZeroValue
	ShouldNotBeZeroValue = assertions.ShouldNotBeZeroValue
	
	ShouldBeGreaterThan          = assertions.ShouldBeGreaterThan
	ShouldBeGreaterThanOrEqualTo = assertions.ShouldBeGreaterThanOrEqualTo
	ShouldBeLessThan             = assertions.ShouldBeLessThan
	ShouldBeLessThanOrEqualTo    = assertions.ShouldBeLessThanOrEqualTo
	ShouldBeBetween              = assertions.ShouldBeBetween
	ShouldNotBeBetween           = assertions.ShouldNotBeBetween
	ShouldBeBetweenOrEqual       = assertions.ShouldBeBetweenOrEqual
	ShouldNotBeBetweenOrEqual    = assertions.ShouldNotBeBetweenOrEqual
	
	ShouldContain       = assertions.ShouldContain
	ShouldNotContain    = assertions.ShouldNotContain
	ShouldContainKey    = assertions.ShouldContainKey
	ShouldNotContainKey = assertions.ShouldNotContainKey
	ShouldBeIn          = assertions.ShouldBeIn
	ShouldNotBeIn       = assertions.ShouldNotBeIn
	ShouldBeEmpty       = assertions.ShouldBeEmpty
	ShouldNotBeEmpty    = assertions.ShouldNotBeEmpty
	ShouldHaveLength    = assertions.ShouldHaveLength
	
	ShouldStartWith           = assertions.ShouldStartWith
	ShouldNotStartWith        = assertions.ShouldNotStartWith
	ShouldEndWith             = assertions.ShouldEndWith
	ShouldNotEndWith          = assertions.ShouldNotEndWith
	ShouldBeBlank             = assertions.ShouldBeBlank
	ShouldNotBeBlank          = assertions.ShouldNotBeBlank
	ShouldContainSubstring    = assertions.ShouldContainSubstring
	ShouldNotContainSubstring = assertions.ShouldNotContainSubstring
	
	ShouldPanic        = assertions.ShouldPanic
	ShouldNotPanic     = assertions.ShouldNotPanic
	ShouldPanicWith    = assertions.ShouldPanicWith
	ShouldNotPanicWith = assertions.ShouldNotPanicWith
	
	ShouldHaveSameTypeAs    = assertions.ShouldHaveSameTypeAs
	ShouldNotHaveSameTypeAs = assertions.ShouldNotHaveSameTypeAs
	ShouldImplement         = assertions.ShouldImplement
	ShouldNotImplement      = assertions.ShouldNotImplement
	
	ShouldHappenBefore         = assertions.ShouldHappenBefore
	ShouldHappenOnOrBefore     = assertions.ShouldHappenOnOrBefore
	ShouldHappenAfter          = assertions.ShouldHappenAfter
	ShouldHappenOnOrAfter      = assertions.ShouldHappenOnOrAfter
	ShouldHappenBetween        = assertions.ShouldHappenBetween
	ShouldHappenOnOrBetween    = assertions.ShouldHappenOnOrBetween
	ShouldNotHappenOnOrBetween = assertions.ShouldNotHappenOnOrBetween
	ShouldHappenWithin         = assertions.ShouldHappenWithin
	ShouldNotHappenWithin      = assertions.ShouldNotHappenWithin
	ShouldBeChronological      = assertions.ShouldBeChronological
	
	ShouldBeError = assertions.ShouldBeError
	ShouldWrap    = assertions.ShouldWrap
)

# Functions ¶
func Convey(items ...interface{})
Convey 是声明规范范围时使用的方法。每个作用域都有一个描述和一个 func()，它可能包含对 Convey()、Reset() 或 Should 式断言的其他调用。只要
您认为合适，Convey 调用就可以嵌套。

重要说明：测试方法中的顶级 Convey() 必须符合以下签名：
Convey(description string, t *testing.T, action func())

所有其他调用应如下所示（无需传入 *testing.T）：
Convey(description string, action func())
别担心，如果你弄错了，goconvey 会 panic，所以你可以修复它。

此外，您可以通过执行以下操作显式获得对 Convey 上下文的访问权限：
Convey(description string, action func(c C))
如果您想将上下文传递给 goroutine，或者将处理程序中的上下文关闭到在 goroutine 中调用您的处理程序的库（想到 httptest），您可能需要这样做。

所有 Convey()-blocks 还接受 FailureMode 的可选参数，该参数设置 goconvey 应如何处理块和嵌套块中 So()-assertions 的失败。有关可用选项
，请参阅此文件中的常量。
默认情况下，它将继承自其父块，并且顶级块默认为 FailureHalts 设置。

此参数插入块本身之前：
Convey(description string, t *testing.T, mode FailureMode, action func())
Convey(description string, mode FailureMode, action func())
请参阅示例包以获取示例。

func FocusConvey(items ...interface{})
FocusConvey 与 SkipConvey 具有相反的效果。如果顶级 Convey 更改为“FocusConvey”，则只会运行使用 FocusConvey 定义的嵌套范围。其余的将被
完全忽略。这在调试重复运行异常函数的大型套件时非常方便，因为您可以禁用除一个函数之外的所有函数，而无需大量的 `SkipConvey` 调用，只需对 FocusConvey
的目标调用链。

func Print(items ...interface{}) (written int, err error)
Print 类似于 fmt.Print（它甚至调用 fmt.Print）。它确保输出与 Web UI 中的相应范围对齐。

func PrintConsoleStatistics()
可以随时调用 PrintConsoleStatistics 来打印断言统计信息。通常，执行此操作的最佳位置是在运行所有测试之后的 TestMain 函数中。像这样：
func TestMain(m *testing.M) {
	convey.SuppressConsoleStatistics()
	result := m.Run()
	convey.PrintConsoleStatistics()
	os.Exit(result)
}

func Printf(format string, items ...interface{}) (written int, err error)
Print 类似于 fmt.Printf（它甚至调用 fmt.Printf）。它确保输出与 Web UI 中的相应范围对齐。

func Println(items ...interface{}) (written int, err error)
Print 类似于 fmt.Println（它甚至调用 fmt.Println）。它确保输出与 Web UI 中的相应范围对齐。

func Reset(action func())
Reset 注册一个清理函数，在同一范围内的每个 Convey() 之后运行。有关简单用例，请参阅示例包。

func SetDefaultFailureMode(mode FailureMode)
SetDefaultFailureMode 允许您为所有 Convey 模块指定默认故障模式。它旨在用于 init 函数，以允许在整个 packgae 的所有测试中更改默认模式，
但它可以在任何地方使用。

func SetDefaultStackMode(mode StackMode)
SetDefaultStackMode 允许您为所有 Convey 模块指定默认堆栈模式。它旨在用于 init 函数，以允许在整个 packgae 的所有测试中更改默认模式，但它
可以在任何地方使用。

func SkipConvey(items ...interface{})
SkipConvey 与 Convey 类似，只是作用域不执行（这意味着在此作用域内定义的子作用域也不会运行）。将通知记者跳过此步骤。

func SkipSo(stuff ...interface{})
SkipSo 与 So 类似，不同之处在于本应传递给 So 的断言不会执行，并且会通知报告者断言已被跳过。

func So(actual interface{}, assert Assertion, expected ...interface{})
对被测系统进行断言的方法也是如此。 assertions 包中的大多数导出名称都以单词“Should”开头，并描述了第一个参数（实际）应该如何与任何最终（预期
）参数进行比较。接受多少最终参数取决于作为断言参数传入的特定断言。有关用例的示例包和有关特定断言方法的文档的断言包。失败的断言将导致调用 t.Fail()
——您永远不应在测试代码中调用此方法（或其他导致失败的方法）。将其留给 GoConvey。

func SoMsg(msg string, actual interface{}, assert Assertion, expected ...interface{})
SoMsg 是 So 的扩展，允许您指定一条消息来报告错误。

func SuppressConsoleStatistics()
SuppressConsoleStatistics 阻止自动打印控制台统计信息。显式调用 PrintConsoleStatistics 将强制打印统计信息。

# Types ¶
type Assertion func(actual interface{}, expected ...interface{}) string
断言是具有 convey.So() 方法可以处理的签名的函数的别名。任何未来的或自定义的断言都应该符合这个方法签名。如果断言通过，返回值应该是一个空字符
串，否则返回一个格式正确的失败消息。

type C
type C interface {
	Convey(items ...interface{})
	SkipConvey(items ...interface{})
	FocusConvey(items ...interface{})
	
	So(actual interface{}, assert Assertion, expected ...interface{})
	SoMsg(msg string, actual interface{}, assert Assertion, expected ...interface{})
	SkipSo(stuff ...interface{})
	
	Reset(action func())
	
	Println(items ...interface{}) (int, error)
	Print(items ...interface{}) (int, error)
	Printf(format string, items ...interface{}) (int, error)
}

C 是 Convey 上下文，您可以通过调用 Convey 来选择在您的操作中获取它，例如：
Convey(..., func(c C) {
	...
})
有关更多详细信息，请参阅有关 Convey 的文档。
此上下文中的所有方法的行为与此包中同名的全局函数相同。

type FailureMode string
FailureMode 是一种类型，它确定 So() 块在其断言失败时应如何失败。请参阅下面的常量以获得可接受的值

type StackMode string
StackMode 是一种类型，它确定 So() 块是否应报告其断言失败的堆栈跟踪。请参阅下面的常量以获得可接受的值
