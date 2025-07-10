package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type user struct {
	Id      int `gorm:"primaryKey"`
	Name    string
	PostNum int
	Posts   []post
}
type post struct {
	Id         int `gorm:"primaryKey"`
	Content    string
	UserId     int
	CommentNum int
	PostStatus string
	Comments   []comment
}
type comment struct {
	Id      int `gorm:"primaryKey"`
	Content string
	PostId  int
}

func (p *post) AfterCreate(tx *gorm.DB) error {
	// 更新用户帖子数
	userReturn := tx.Model(&user{}).
		Where("id = ?", p.UserId).
		Update("post_num", gorm.Expr("post_num + ?", 1))
	if userReturn.Error != nil {
		return userReturn.Error
	}
	//更新帖子状态 默认为无评论
	if err := tx.Model(&post{}).
		Where("id = ?", p.Id).
		Update("post_status", gorm.Expr("CASE WHEN comment_num = 0 THEN '无评论' ELSE '有评论' END")).Error; err != nil {
		return err
	}
	return nil
}

//func (p *post) AfterUpdate(tx *gorm.DB) error {
//	// 更新用户帖子数
//	return tx.Model(&user{}).
//		Where("id = ?", p.UserId).
//		Update("post_num", gorm.Expr("post_num + ?", 1)).Error
//}
//func (c *comment) AfterUpdate(tx *gorm.DB) error {
//	// 更新帖子评论数
//	if err := tx.Model(&post{}).
//		Where("id = ?", c.PostId).
//		Update("comment_num", gorm.Expr("comment_num + ?", 1)).Error; err != nil {
//		return err
//	}
//	//更新帖子状态 PostStatus  若没有评论则更新为无评论
//	var post post
//	if err := tx.First(&post, c.PostId).Error; err != nil {
//		return err
//	}
//	if post.CommentNum == 0 {
//		post.PostStatus = "无评论"
//	}
//	return tx.Save(&post).Error
//}

func (c *comment) AfterCreate(tx *gorm.DB) error {
	//post1 := post{}
	//if err := tx.First(&post1, c.PostId).Error; err != nil {
	//	return err
	//}
	//var postStatus string
	//if post1.CommentNum == 0 {
	//	postStatus = "无评论"
	//} else {
	//	postStatus = "有评论"
	//}
	//var commentnmu int = post1.CommentNum + 1
	//// 更新帖子评论数
	////更新帖子状态 PostStatus  若没有评论则更新为无评论
	//var post2 = post{}
	//if err := tx.Model(&post{}).
	//	Where("id = ?", c.PostId).
	//	Update("comment_num", gorm.Expr("comment_num + ?", 1)).
	//	Update("post_status", "有评论").Error; err != nil {
	//	return err
	//}
	if err := tx.Model(&post{}).
		Where("id = ?", c.PostId).
		UpdateColumns(map[string]interface{}{"comment_num": gorm.Expr("comment_num + ?", 1), "post_status": "有评论"}).
		Error; err != nil {
		return err
	}
	//UPDATE `posts` SET `comment_num`=comment_num + 1 WHERE id = 2
	//UPDATE `posts` SET `post_status`='有评论' WHERE id = 2
	//if err := tx.Model(&post2).
	//	Where("id = ?", c.PostId).
	//	Updates(&post{CommentNum: commentnmu, PostStatus: postStatus}).Error; err != nil {
	//	return err
	//}
	return nil
}

func (c *comment) AfterDelete(tx *gorm.DB) error {
	//删除评论时 帖子评论数减一 若评论数为0 则更新为无评论
	if err := tx.Model(&post{}).
		Where("id = ?", c.PostId).
		UpdateColumns(map[string]interface{}{"comment_num": gorm.Expr("comment_num - ?", 1), "post_status": gorm.Expr(`
                CASE 
                    WHEN comment_num <= 1 THEN '无评论' 
                    ELSE post_status 
                END`)}).
		Error; err != nil {
		return err
	}
	return nil
}

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/studyweb3?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//初始化
	db.AutoMigrate(&user{}, &post{}, &comment{})
	db.Debug().Create(&user{
		Name: "张三",
		Posts: []post{
			{Content: "张三的帖子",
				Comments: []comment{
					{Content: "张三的评论"},
					{Content: "李四的评论"},
				}},
			{Content: "张三的帖子1",
				Comments: []comment{
					{Content: "张三的评论1"},
					{Content: "李四的评论2"},
					{Content: "王五的评论3"},
				}},
			{Content: "张三的帖子2", Comments: []comment{{Content: "张三的评论4"}}},
		},
	})
	//删除
	coment := comment{Id: 6, PostId: 3}
	db.Delete(&coment)
	//user := user{Id: 1}
	//db.Preload("Posts").Preload("Posts.Comments").Find(&user)
	//fmt.Println(user)
	////查询评论条数最多的帖子   psots 关联 comments表 使用 group by 函数查询出条数最多的帖子
	//post := post{}
	//db.Debug().
	//	Model(post).
	//	Select("posts.*,count(comments.id) as Comments").
	//	Joins("left join comments on comments.post_id = posts.id").
	//	Group("posts.id").
	//	Order("Comments DESC").Limit(1).Find(&post)
	//fmt.Println(post)
}
