package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type students struct {
	Id     int `gorm:"primaryKey"`
	Name   string
	Age    int
	Gender string
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/studyweb3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//创建
	//db.AutoMigrate(&students{})
	//新增一条记录
	//student1 := students{Name: "张三", Age: 20, Gender: "三年及"}
	//db.Create(&student1)
	//查询并修改
	//student2 := students{}
	//db.Where("Age > ?", 18).Find(&student2)
	//fmt.Println(student2)
	//student2.Gender = "四年级"
	//db.Save(&student2)
	//新增两条记录
	//student3 := []*students{
	//	&students{Name: "张一", Age: 10, Gender: "一年级"},
	//	&students{Name: "张二", Age: 14, Gender: "二年级"},
	//}
	//db.Create(student3)
	//删除年龄小于15的
	student2 := students{}
	db.Debug().Where("Age < ?", 15).Delete(&student2)
}
