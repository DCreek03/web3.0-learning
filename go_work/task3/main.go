package main

import (
	"fmt"
	"time"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 题目1：基本CRUD操作
// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、
//
//	grade （学生年级，字符串类型）。
//
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
type Student struct {
	Id    int
	Name  string
	Age   int
	Grade string
}

func q1(db *gorm.DB) {

	db.AutoMigrate(&Student{})

	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	db.Create(&student)
	// 批量插入一些学生数据
	students := []*Student{
		{Name: "李四", Age: 19, Grade: "二年级"},
		{Name: "林一", Age: 18, Grade: "二年级"},
		{Name: "罗二", Age: 20, Grade: "二年级"},
		{Name: "王五", Age: 10, Grade: "一年级"},
	}
	db.Create(students)

	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	var students2 []Student
	result := db.Where("Age > ?", 18).Find(&students2)
	if result.Error != nil {
		panic("查询失败！" + result.Error.Error())
	}
	for _, student := range students {
		fmt.Printf("姓名：%s\n", student.Name)
	}

	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	result2 := db.Model(&Student{}).Where("name = ?", "张三").Update("Grade", "四年级")
	if result2.Error != nil {
		panic("修改失败！" + result2.Error.Error())
	}

	// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	result3 := db.Where("age < ?", 15).Delete(&Student{})
	if result3.Error != nil {
		panic("删除失败！" + result3.Error.Error())
	}
}

// 题目2：事务语句
// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和
// transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
// 在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
// 向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
type Account struct {
	Id      int
	Balance float64
}

type Transaction struct {
	Id            int
	FromAccountId int
	ToAccountId   int
	Amount        float64
}

func q2(db *gorm.DB) {
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

	// 添加A、B两个账户
	// accounts := []*Account{
	// 	{Balance: 300.0},
	// 	{Balance: 99.0},
	// }
	// db.Create(accounts)

	// 此时1中300 2中99
	err := transferMoney(1, 2, db)
	if err != nil {
		fmt.Println("转账失败:", err)
	} else {
		fmt.Println("转账成功!")
	}

}

func transferMoney(fromId int, toId int, db *gorm.DB) error {
	// 开启事务
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	// 先查询转出账号中金额是否足够
	var account Account
	result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&account, fromId)
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("查询转出账户失败: %v", result.Error)
	}
	if account.Balance < 100 {
		tx.Rollback()
		return fmt.Errorf("账号: %d 中余额不足！", fromId)
	}

	// 查询转入账号
	var toAccount Account
	result = tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&toAccount, toId)
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("查询转入账户失败: %v", result.Error)
	}

	// 更新余额
	result = tx.Model(&Account{}).Where("Id = ?", fromId).Update("balance", gorm.Expr("balance - ?", 100.0))
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("更新转出账户失败: %v", result.Error)
	}
	result = tx.Model(&Account{}).Where("Id = ?", toId).Update("balance", gorm.Expr("balance + ?", 100.0))
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("更新转入账户失败: %v", result.Error)
	}
	// 写入表
	tr := Transaction{
		FromAccountId: fromId,
		ToAccountId:   toId,
		Amount:        100,
	}
	if result := tx.Create(&tr); result.Error != nil {
		return fmt.Errorf("创建交易记录失败: %v", result.Error)
	}

	return tx.Commit().Error
}

//	题目1：使用SQL扩展库进行查询
//
// 假设你已经使用Sql连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
type Employee struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func q3(db *sqlx.DB) {

	var findITDepart []Employee
	const sql = `select * from employees where department = ?`
	err := db.Select(&findITDepart, sql, "技术部")
	if err != nil {
		fmt.Println("查询部门为技术部的员工失败，失败原因是:", err)
	}
	for _, e := range findITDepart {
		fmt.Printf("员工：%s，部门：%s，工资：%f\n", e.Name, e.Department, e.Salary)
	}

	var maxSalaryEmployee Employee
	const sql1 = `select * from employees order by salary desc limit 1`
	err = db.Get(&maxSalaryEmployee, sql1)
	if err != nil {
		fmt.Println("查询工资最高的员工失败，失败原因是:", err)
	}
	fmt.Printf("员工：%s，部门：%s，工资：%f\n", maxSalaryEmployee.Name, maxSalaryEmployee.Department, maxSalaryEmployee.Salary)

}

// 题目2：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
type Book struct {
	Id     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func q4(db *sqlx.DB) {

	var findBooks []*Book
	const sql = `select * from books where price > ？`
	err := db.Select(&findBooks, sql, 50.0)
	if err != nil {
		fmt.Println("查询查询价格大于 50 元的书籍，失败原因是:", err)
	}
	for _, b := range findBooks {
		fmt.Printf("书名：%s，作者：%s，价格：%f\n", b.Title, b.Author, b.Price)
	}
}

// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
// Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

// ====================== 实体定义 ======================
type User struct {
	ID         uint   `gorm:"primaryKey"`
	Username   string `gorm:"uniqueIndex;size:50;not null"`
	Email      string `gorm:"uniqueIndex;size:100;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Posts      []Post         `gorm:"foreignKey:AuthorID"` // 一对多关系
	PostsCount int            `gorm:"default:0"`           // 用户文章数量统计字段

}

type Post struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"size:200;not null"`
	Content       string `gorm:"type:text;not null"`
	AuthorID      uint   `gorm:"index;not null"` // 外键
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Comments      []Comment      `gorm:"foreignKey:PostID"`                    // 一对多关系
	CommentStatus string         `gorm:"type:enum('无评论','有评论');default:'无评论'"` // 评论状态
}

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	AuthorID  uint   `gorm:"index;not null"` // 外键
	PostID    uint   `gorm:"index;not null"` // 外键
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func GetSampleUsers() []User {
	return []User{
		{
			ID:       1,
			Username: "tech_guru",
			Email:    "guru@tech.com",
		},
		{
			ID:       2,
			Username: "code_master",
			Email:    "master@code.dev",
		},
		{
			ID:       3,
			Username: "dev_learner",
			Email:    "learner@dev.net",
		},
	}
}

func GetSamplePosts() []Post {
	return []Post{
		{
			ID:       1,
			Title:    "深入理解Golang并发模型",
			Content:  "Go语言的并发模型是其最强大的特性之一...",
			AuthorID: 1,
		},
		{
			ID:       2,
			Title:    "GORM高级技巧大全",
			Content:  "本文将介绍GORM的各种高级用法和最佳实践...",
			AuthorID: 1,
		},
		{
			ID:       3,
			Title:    "从零构建RESTful API",
			Content:  "使用Go和Gin框架构建高性能API服务...",
			AuthorID: 2,
		},
		{
			ID:       4,
			Title:    "数据库优化实战",
			Content:  "如何优化SQL查询提升应用性能...",
			AuthorID: 2,
		},
		{
			ID:       5,
			Title:    "微服务架构设计模式",
			Content:  "微服务架构的常见模式和反模式...",
			AuthorID: 3,
		},
	}
}

func GetSampleComments() []Comment {
	return []Comment{
		{
			ID:       1,
			Content:  "非常有深度的文章！",
			AuthorID: 2,
			PostID:   1,
		},
		{
			ID:       2,
			Content:  "期待更多关于channel的内容",
			AuthorID: 3,
			PostID:   1,
		},
		{
			ID:       3,
			Content:  "GORM的关联查询确实很方便",
			AuthorID: 1,
			PostID:   2,
		},
		{
			ID:       4,
			Content:  "解决了我的实际问题",
			AuthorID: 3,
			PostID:   3,
		},
		{
			ID:       5,
			Content:  "优化后性能提升明显",
			AuthorID: 1,
			PostID:   4,
		},
		{
			ID:       6,
			Content:  "实例代码能否分享一下？",
			AuthorID: 3,
			PostID:   4,
		},
		{
			ID:       7,
			Content:  "架构设计思路很清晰",
			AuthorID: 2,
			PostID:   5,
		},
	}
}

func q5(db *gorm.DB) {
	models := []interface{}{
		&User{},
		&Post{},
		&Comment{},
	}

	db.AutoMigrate(models...)
	db.Create(GetSampleUsers())
	db.Create(GetSamplePosts())
	db.Create(GetSampleComments())
}

// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。

func q6(db *gorm.DB) {
	// 查询 tech_guru 发表的文章
	var user User
	result := db.Where("username = ?", "tech_guru").First(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	var post []*Post
	result = db.Where("author_id = ?", user.ID).Find(&post)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	for _, p := range post {
		fmt.Println(p)
	}

	subQuery := db.Model(&Comment{}).
		Select("post_id, COUNT(*) AS comment_count").
		Group("post_id").
		Order("comment_count DESC").
		Limit(1)
	var maxCommentPost Post
	result = db.Joins("JOIN (?) AS max_post ON posts.id = max_post.post_id", subQuery).First(&maxCommentPost)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	fmt.Println(maxCommentPost)

}

// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (p *Post) AfterCreate(db *gorm.DB) (err error) {
	result := db.Model(&User{}).Where("id = ?", p.AuthorID).Update("posts_count", gorm.Expr("posts_count + ?", 1))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *Comment) AfterDelete(db *gorm.DB) (err error) {
	var commentCount int64
	result := db.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount)
	if result.Error != nil {
		return result.Error
	}
	status := "有评论"
	if commentCount == 0 {
		status = "无评论"
		result = db.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", status)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil

}
func q7(db *gorm.DB) {

	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	// 创建用户
	user := User{Username: "张三122UUU", Email: "zhangsan@example.com111UUU"}
	db.Create(&user)

	// 创建文章 (触发AfterCreate钩子)
	post := Post{
		Title:    "我的第一篇博客",
		Content:  "欢迎来到我的博客...",
		AuthorID: user.ID,
	}
	db.Create(&post) // 此时用户的PostsCount自动+1

	// 创建评论
	comment := Comment{
		Content:  "很好的文章!",
		PostID:   post.ID,
		AuthorID: user.ID,
	}
	db.Create(&comment)

	// 删除评论 (触发AfterDelete钩子)
	db.Delete(&comment) // 删除后检查文章状态
}

func main() {
	// gorm的连接
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:!!!@tcp(127.0.0.1:3306)/study_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// sqlx 的连接
	// db, err := sqlx.Connect("mysql", "@tcp(127.0.0.1:3306)/study_go?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	// q1(db)

	// q2(db)

	// q3(db)

	// q4(db)

	// q5(db)
	// q6(db)
	q7(db)

}
