package goconvey

import "testing"

SMARTY 免责声明：根据相关许可协议的条款，该软件可免费供您使用。这个软件是免费的，就像小狗一样，是一份礼物。享受你的新责任。这意味着虽然我们可
能会考虑增强请求，但我们可能会或可能不会自行决定接受请求。

GoConvey 是很棒的 Go 测试
欢迎使用 GoConvey，这是一个很好用的 gophers Go 测试工具。与 go test 一起工作。根据您的观看乐趣在终端或浏览器中使用它。查看全功能导览。

GoConvey 支持 Go 的当前版本（请参阅官方 Go 发布策略）。目前这意味着支持 Go 1.16 和 Go 1.17。

Features:
	1.直接与 go test 集成
	2.全自动 Web UI（也适用于本机 Go 测试）
	3.大量的回归测试
	4.显示测试覆盖率
	5.可读的彩色控制台输出（无论是否是 IT 人员，任何经理都可以理解）
	6.测试代码生成器
	7.桌面通知（可选）
	8.立即在 Sublime Text 中打开问题行（需要一些组装）
您可以在 StackOverflow 上询问有关如何使用 GoConvey 的问题。使用标签 go 和 goconvey。

Menu:
	1.Installation
	2.Quick start
	3.Documentation
	4.Screenshots
	5.Contributors

# Installation
	$ go get github.com/smartystreets/goconvey

# Quick start
做个测试，例如：
package package_name

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	
	// Only pass t into top-level Convey calls								// 只将 t 传递给顶级 Convey 调用
	Convey("Given some integer with a starting value", t, func() {			// 给定一些具有起始值的整数
		x := 1
		
		Convey("When the integer is incremented", func() {					// 当整数递增时
			x++
			
			Convey("The value should be greater by one", func() {			// 该值应大一
				So(x, ShouldEqual, 2)
			})
		})
	})
}

## In the browser
在您的项目路径中启动 GoConvey Web 服务器：
	$ $GOPATH/bin/goconvey
	
然后在您的浏览器中观看测试结果显示：
	http://localhost:8080

如果浏览器没有自动打开，请点击http://localhost:8080 手动打开。

There you have it.
	(web 画面)

只要 GoConvey 正在运行，测试结果就会在您的浏览器窗口中自动更新
	(web 画面)
该设计是响应式的，因此如果您需要将浏览器放在代码旁边，您可以将其压缩得非常紧。
Web UI 支持传统的 Go 测试，因此即使您不使用 GoConvey 测试也可以使用它。

## In the terminal
只做你最擅长的事：
	$ go test

或者，如果您希望输出包含故事：
	$ go test -v

# Documentation
查看:
	1.GoConvey wiki,
	2.GoConvey reference
	3.和分散在整个项目中的 *_test.go 文件。

# Screenshots
对于 Web UI 和终端屏幕截图，请查看完整功能导览。

# Contributors
GoConvey 由 SmartyStreets 和几位贡献者提供给您（谢谢！）。
