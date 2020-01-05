package dao

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Uid        string `gorm:"not null"` // 6位随机id
	Name       string
	Password   string `gorm:"not null"`
	Email      string
	Mobile     string
	InviteCode string // 邀请码
	Status     int    `gorm:"default 0;not null"` // 是否禁用 0 禁用 1 启用
}

func getDb() (*gorm.DB, error) {
	return gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?timeout=5s&readTimeout=3s&writeTimeout=3s&parseTime=true&loc=Local&charset=utf8mb4,utf8")
}

func insert(db *gorm.DB, tableName string, model interface{}) (interface{}, error) {
	if err := db.Table(tableName).Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func TestInsert(t *testing.T) {
	db, err := getDb()
	defer db.Close()
	if err != nil {
		panic(err)
	}

	// 创建表的时候名字不是复数
	db.SingularTable(true)
	// 启用Logger，显示详细日志
	db.LogMode(true)
	// Migrate the schema
	db.AutoMigrate(&User{})

	// 创建
	if user, err := insert(db, "user", &User{
		Uid:        "fdjkfd",
		Name:       "xiaoming",
		Password:   "1298978787",
		Email:      "12jfkdi897",
		Mobile:     "18238029999",
		InviteCode: "fjdki",
		Status:     0,
	}); err != nil || user == nil {
		panic(err)
	} else {
		t.Log(user.(*User).ID)
	}
}

func find(db *gorm.DB, tableName string, args interface{}) *gorm.DB {
	return db.Table(tableName).Where(args)
}

func TestFind(t *testing.T) {
	db, err := getDb()
	defer db.Close()
	if err != nil {
		panic(err)
	}

	// 创建表的时候名字不是复数
	db.SingularTable(true)
	// 启用Logger，显示详细日志
	db.LogMode(true)
	// Migrate the schema
	//db.AutoMigrate(&User{})

	user := User{
		Uid: "fdjkfd",
	}
	var users []User
	if err := find(db, "user", user).Limit(3).Offset((1 - 1) * 3).Find(&users).Error; err != nil {
		t.Error(err)
	}
	for _, v := range users {
		t.Log(v)
	}
}

func delete(db *gorm.DB, tableName string, args interface{}) (error, int64) {
	if argsBytes, err := json.Marshal(args); err != nil {
		return err, 0
	} else {
		var dat map[string]interface{}
		if err = json.Unmarshal(argsBytes, &dat); err != nil {
			return err, 0
		}
		if dat["ID"].(float64) == 0 {
			return fmt.Errorf("ID Is Nil"), 0
		}
	}
	d := db.Table(tableName).Where(args).Delete(args)
	return d.Error, d.RowsAffected
}

func TestDelete(t *testing.T) {
	db, err := getDb()
	defer db.Close()
	if err != nil {
		panic(err)
	}

	// 创建表的时候名字不是复数
	db.SingularTable(true)
	// 启用Logger，显示详细日志
	db.LogMode(true)
	// Migrate the schema
	//db.AutoMigrate(&User{})

	user := User{
		Model: gorm.Model{ID: 4},
		Uid:   "fdjkfz",
	}
	if err, i := delete(db, "user", user); err != nil {
		t.Error(err)
	} else {
		t.Log(i)
	}
}

func update(db *gorm.DB, tableName string, whereArgs interface{}, updateArgs interface{}) (error, int64) {
	u := db.Table(tableName).Where(whereArgs).Updates(updateArgs)
	return u.Error, u.RowsAffected
}

func TestUpdate(t *testing.T) {
	db, err := getDb()
	defer db.Close()
	if err != nil {
		panic(err)
	}

	// 创建表的时候名字不是复数
	db.SingularTable(true)
	// 启用Logger，显示详细日志
	db.LogMode(true)
	// Migrate the schema
	//db.AutoMigrate(&User{})

	user := User{
		Name: "xiaoming",
		Uid:  "fdjkfd",
	}
	if err, i := update(db, "user", User{
		Name: "xiaohong",
		Uid:  "fdjkfz",
	}, user); err != nil {
		t.Error(err)
	} else {
		t.Log(i)
	}
}
