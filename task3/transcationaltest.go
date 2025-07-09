package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:root@tcp(127.0.0.1:3306)/studyweb3?charset=utf8mb4&parseTime=True&loc=Local"

type account struct {
	Id      int `gorm:"primaryKey"`
	Name    string
	Balance int
}

type transaction struct {
	Id            int `gorm:"primaryKey"`
	FromAccountId int
	ToAccountId   int
	Amount        int
}

func transfer(db *gorm.DB, fromId, toId, amount int) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var fromAccount account
		if err := tx.First(&fromAccount, fromId).Error; err != nil {
			return err
		}
		if fromAccount.Balance < amount {
			return fmt.Errorf("余额不足")
		}
		fromAccount.Balance -= amount
		if err := tx.Save(&fromAccount).Error; err != nil {
			return err
		}
		var toAccount account
		if err := tx.First(&toAccount, toId).Error; err != nil {
			return err
		}
		toAccount.Balance += amount
		if err := tx.Save(&toAccount).Error; err != nil {
			return err
		}
		if err := tx.Create(&transaction{FromAccountId: fromId, ToAccountId: toId, Amount: amount}).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	})
}
func teansfer1(db *gorm.DB, fromId, toId, amount int) error {
	tx := db.Begin()
	var fromAccount account
	if err := tx.First(&fromAccount, fromId).Error; err != nil {
		return err
	}
	if fromAccount.Balance < amount {
		return fmt.Errorf("余额不足")
	}
	fromAccount.Balance -= amount
	if err := tx.Save(&fromAccount).Error; err != nil {
		tx.Rollback()
		return err
	}
	var toAccount account
	if err := tx.First(&toAccount, toId).Error; err != nil {
		tx.Rollback()
		return err
	}
	toAccount.Balance += amount
	if err := tx.Save(&toAccount).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&transaction{FromAccountId: fromId, ToAccountId: toId, Amount: amount}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func main() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//初始化
	//db.AutoMigrate(&account{}, &transaction{})
	//accounts := []account{
	//	{Name: "张三", Balance: 150},
	//	{Name: "李四", Balance: 50},
	//}
	//db.Create(&accounts)
	err2 := transfer(db, 1, 2, 100)
	if err2 != nil {
		fmt.Println(err2)
	}
	err1 := teansfer1(db, 1, 2, 100)
	if err1 != nil {
		fmt.Println(err1)
	}
}
