package goconvey

import "testing"

# Home
欢迎使用 GoConvey，这是一个很好用的 gophers 测试工具。

## Main Features
	1.与 go test 集成
	2.可读的彩色控制台输出
	3.全自动网络用户界面
	4.大量的回归测试
	5.测试代码生成器
查看与其他 Go 测试工具相比所有功能的综合表。

## Get going in 25 seconds
1.在您的终端中：
	# make sure your GOPATH is set							# 确保你的 GOPATH 已设置
	
	$ cd <project path>
	$ go get github.com/smartystreets/goconvey
	$ $GOPATH/bin/goconvey

2.在您的浏览器中：
	http://localhost:8080
如果您有现有的 Go 测试，它们将自动运行并且结果将显示在您的浏览器中。

## Your first GoConvey test
打开任何 _test.go 文件并将其放入其中，自定义您的包声明：
package package_name

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestIntegerStuff(t *testing.T) {
	Convey("Given some integer with a starting value", t, func() {
		x := 1
		
		Convey("When the integer is incremented", func() {
			x++
			
			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}
保存文件，然后浏览浏览器窗口，您会看到新测试已经运行。
更改断言（带有 So() 的行）使测试失败，然后在浏览器中查看输出更改。
您也可以像往常一样从终端运行测试，使用 go test。如果您希望测试在终端中自动运行，请查看自动测试脚本。

## Required Reading
如果我能确保每个 GoConvey 用户只从这个存储库中读取一段代码，那将是隔离执行测试( isolated execution tests)。这些测试是 GoConvey 执行模
型的最佳文档。

## Full Documentation
有关断言、编写测试、执行等的详细信息，请参阅 Documentation 索引。
