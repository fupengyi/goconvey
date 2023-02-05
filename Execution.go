package goconvey
# Execution
要运行测试，请在终端中执行您最擅长的操作：
	go test

示例输出：（实际输出是彩色的）
.....

5 assertions and counting

....

9 assertions and counting

PASS
ok  	github.com/smartystreets/goconvey/examples	0.022s

========================================================================================================================
您还可以使用 -v 进行详细说明：
	go test -v

示例输出：（实际输出是彩色的）
=== RUN TestScoring

Subject: Bowling Game Scoring
Given a fresh score card
When all gutter balls are thrown
The score should be zero ✔
When all throws knock down only one pin
The score should be 20 ✔
When a spare is thrown
The score should include a spare bonus. ✔
When a strike is thrown
The score should include a strike bonus. ✔
When all strikes are thrown
The score should be 300. ✔

5 assertions and counting

--- PASS: TestScoring (0.00 seconds)
=== RUN TestSpec

Subject: Integer incrementation and decrementation
Given a starting integer value
When incremented
The value should be greater by one ✔
The value should NOT be what it used to be ✔
When decremented
The value should be lesser by one ✔
The value should NOT be what it used to be ✔

9 assertions and counting

--- PASS: TestSpec (0.00 seconds)
PASS
ok  	github.com/smartystreets/goconvey/examples	0.023s
========================================================================================================================

## Auto-test and web UI
如果您厌倦了按 ↑，一直按 Enter，请尝试自动运行测试。
如果您完全厌倦了终端，请查看在浏览器中优雅地显示测试输出的 Web UI。

## Next
最后，了解 Skip 以跳过/忽略范围和断言。
