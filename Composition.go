package goconvey

import "testing"

# Composition
使用 GoConvey 编写自记录测试非常容易。

## Examples
首先，查看示例文件夹以获得基本概念。我们建议查看 isolated_execution_test.go 以更透彻地了解如何编写测试用例。

## Functions
有关导出的函数和断言，请参阅 GoDoc。您会对 convey 包最感兴趣。

## Quick tutorial
在您的测试文件中，导入所需的包：
import(
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)
（为方便起见，请注意传送包的点符号。）

由于 GoConvey 使用 go test，所以设置一个 Test 函数：
func TestSomething(t *testing.T) {

}

为了设置测试用例，我们使用 Convey() 来定义范围/上下文/行为/想法，并使用 So() 进行断言。例如：
Convey("1 should equal 1", t, func() {
	So(1, ShouldEqual, 1)
})

有一个有效的 GoConvey 测试。请注意，我们传入了 *testing.T 对象。只有对 Convey() 的顶级调用需要这样做。对于嵌套调用，您必须省略它。例如：
Convey("Comparing two variables", t, func() {
	myVar := "Hello, world!"
	
	Convey(`"Asdf" should NOT equal "qwerty"`, func() {
		So("Asdf", ShouldNotEqual, "qwerty")
	})
	
	Convey("myVar should not be nil", func() {
		So(myVar, ShouldNotBeNil)
	})
})

如果您尚未实施测试或范围，只需将其功能设置为 nil 即可跳过它：
Convey("This isn't yet implemented", nil)

## Next
接下来，您应该了解标准断言。您也可以跳到执行测试或跳到 Skip 以使测试更方便。
