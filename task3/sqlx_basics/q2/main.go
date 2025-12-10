package q2

import (
	"context"
	"log"
	"time"
)

/*
*
### 题目2：实现类型安全映射
- 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
  - 要求 ：
  - 定义一个 Book 结构体，包含与 books 表对应的字段。
  - 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/
func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var books []Book

	query := `
        SELECT *
        FROM books
        WHERE price > ?
        ORDER BY price DESC
    `
	query = db.Rebind(query)

	if err := db.SelectContext(ctx, &books, query, 50); err != nil {
		log.Fatal(err)
	}
}

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}
