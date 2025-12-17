package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employees struct {
	Id         int
	Name       string
	Department string
	Salary     int
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		fmt.Println("err:", err)
	}

	defer db.Close()

	employees, err := getEmployees(db)

	if err != nil {
		fmt.Println("查询失败:", err)
	}

	for _, emp := range employees {
		fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 薪水: %d\n",
			emp.Id, emp.Name, emp.Department, emp.Salary)
	}

	maxEmp, err := getMaxSalary(db)

	if err != nil {
		fmt.Println("查询失败:", err)
	}

	fmt.Printf("工资最高的员工：\n")
	fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %d\n",
		maxEmp.Id, maxEmp.Name,
		maxEmp.Department, maxEmp.Salary)

}

func getEmployees(db *sqlx.DB) ([]Employees, error) {
	var employees []Employees

	err := db.Select(&employees, "select * from employees where department = ?", "技术部")

	if err != nil {
		return nil, fmt.Errorf("查询员工失败:%v", err)
	}
	return employees, err
}

func getMaxSalary(db *sqlx.DB) (*Employees, error) {
	var emp Employees

	err := db.Get(&emp, "select * from employees order by salary desc limit 1")

	if err != nil {
		return nil, fmt.Errorf("查询最高工资员工失败: %v", err)
	}
	return &emp, err
}
