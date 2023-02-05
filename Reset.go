package goconvey
# Reset
当您的 Conveys 涉及一些设置时，您可能需要在测试之后或测试之间拆除。在这些情况下使用 Reset() 进行清理。 Convey 的 Reset() 在同一范围内的
每个 Convey() 的末尾运行。

例如：
Convey("Top-level", t, func() {
	
	// setup (run before each `Convey` at this scope):						// 设置（在此范围内的每个 `Convey` 之前运行）：
	db.Open()
	db.Initialize()
	
	Convey("Test a query", t, func() {
		db.Query()
		// TODO: assertions here
	})
	
	Convey("Test inserts", t, func() {
		db.Insert()
		// TODO: assertions here
	})
	
	Reset(func() {
		// This reset is run after each `Convey` at the same scope.			// 此 reset 在同一范围内的每个 `Convey` 之后运行。
		db.Close()
	})
	
})























