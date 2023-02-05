package goconvey

import (
	"database/sql"
	"testing"
)

# Decorating tests to provide common logic
有时，一个测试套件有一些特定的资源或状态，必须为每个测试设置然后拆除。这是一种通过将测试包装在辅助函数中来简化该过程的方法。
这对于其他共享的公共状态很有用，比如数据库事务，在其中设置事务，让测试对其进行操作，检查它是否仍然存在，最后在下一次测试之前执行回滚。

## Basic principle
基本思想是这样的：
	* 提供执行以下操作的函数：
		i.接受一个闭包 f 作为参数，它将被初始化的资源或状态调用
		ii.创建并返回执行以下操作的 func()：
			a.初始化资源
			b.如果需要，为资源提供重置
			c.执行闭包 f
	* 然后在测试中使用这个函数：
		i.创建一个将初始化的资源或状态作为参数的闭包
		ii.将此闭包传递给创建的辅助函数
		iii.将辅助函数的结果作为块传递给 Convey 调用

注意：报告失败时，将通过遍历堆栈跟踪查找以 _test.go 结尾的文件中的第一个位置来报告失败。出于这个原因，对 So 的调用出现在这样的文件中很重要。
否则，错误将在堆栈中的不同（可能更高）级别报告，并且不清楚到底发生了什么故障。

## Example
package main

import (
	"database/sql"
	"testing"
	
	_ "github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
)

func WithTransaction(db *sql.DB, f func(tx *sql.Tx)) func() {
	return func() {
		tx, err := db.Begin()
		So(err, ShouldBeNil)
		
		Reset(func() {
			/* Verify that the transaction is alive by executing a command */			/* 通过执行命令验证事务是否存活 */
			_, err := tx.Exec("SELECT 1")
			So(err, ShouldBeNil)
			
			tx.Rollback()
		})
		
		/* Here we invoke the actual test-closure and provide the transaction */		/* 这里我们调用实际的测试闭包并提供交易 */
		f(tx)
	}
}

func TestUsers(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://localhost?sslmode=disable")
	if err != nil {
		panic(err)
	}
	
	Convey("Given a user in the database", t, WithTransaction(db, func(tx *sql.Tx) {
		_, err := tx.Exec(`INSERT INTO "Users" ("id", "name") VALUES (1, 'Test User')`)
		So(err, ShouldBeNil)
		
		Convey("Attempting to retrieve the user should return the user", func() {
			var name string
			
			data := tx.QueryRow(`SELECT "name" FROM "Users" WHERE "id" = 1`)
			err = data.Scan(&name)
			
			So(err, ShouldBeNil)
			So(name, ShouldEqual, "Test User")
		})
	}))
}

/* Required table to run the test:								运行测试所需的表：
CREATE TABLE "public"."Users" (
	"id" INTEGER NOT NULL UNIQUE,
	"name" CHARACTER VARYING( 2044 ) NOT NULL
);
*/

Output:
=== RUN TestUsers

Given a user in the database ✔✔
Attempting to retrieve the user should return the user ✔✔✔

--- PASS: TestUsers (0.01 seconds)
PASS
ok      test    0.022s
