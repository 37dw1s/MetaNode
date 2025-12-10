package q1

import (
	"context"
	"fmt"
	"log"
	"time"
)

/*
*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var emps []Employee
	baseSQL := `SELECT * FROM employees WHERE department = ?`
	query := db.Rebind(baseSQL)
	if err := db.SelectContext(ctx, &emps, query, "技术部"); err != nil {
		log.Fatal(err)
	}

	var top Employee
	baseTop := `SELECT *
	            FROM employees
	            ORDER BY salary DESC
	            LIMIT 1`
	queryTop := db.Rebind(baseTop)
	if err := db.GetContext(ctx, &top, queryTop); err != nil {
		log.Fatal(err)
	}
}

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}
