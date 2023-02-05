package goconvey
# Auto test
有两种方法可以使用 GoConvey 自动运行测试：在终端或浏览器中。

1. 在终端中：auto-run.py
要在终端中查看测试报告，请使用类似 auto-run.py 的文件，它曾与 GoConvey 捆绑在一起。首先下载脚本并将其放在路径中的某个位置。
	cd <folder_with_tests_or_packages>
	auto-run.py -v
从项目的顶级文件夹运行它，对当前目录下的 .go 文件的任何更改都会触发测试运行。 -v 选项启用“详细”模式，并且完全是可选的。

2.在浏览器中：goconvey server
当您使用 Web UI 查看测试结果时，测试已经自动运行。

## Performance
方法 #2（在浏览器中）的优点之一是测试跨多个 goroutine 运行以尽快返回结果。
