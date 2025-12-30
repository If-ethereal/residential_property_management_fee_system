package db_conn

import (
	"fmt"
	"strconv"

	"math/rand"

	//"fmt"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"os"
	"errors"
	"time"
)

type Crud interface {
	Add()
	Update()
	Delete()
	Find()
}
type User struct {
	gorm.Model
	Account  uint `gorm:"primaryKey;unique"`
	Password string
	Name     string
	// 如果需要反向引用，添加以下字段
	Courses []Course `gorm:"many2many:user_courses;foreignKey:Account;joinForeignKey:UserAccount;references:Code;joinReferences:CourseCode"`
}

type Course struct {
	gorm.Model
	Code    int `gorm:"primaryKey;unique"`
	Name    string
	Credits int
	Users   []User `gorm:"many2many:user_courses;foreignKey:Code;joinForeignKey:CourseCode;references:Account;joinReferences:UserAccount"`
}

type UserCourse struct {
	UserAccount uint `gorm:"primaryKey"`
	CourseCode  int  `gorm:"primaryKey"`
}

func (co *Course) BeforeCreate(tx *gorm.DB) (err error) {
	if co.Code == 0 {
		co.Code = rand.Intn(900000) + 100000
	}
	return err
}

func (co Course) Add(db *gorm.DB) (bool, error) {
	course := co
	result := db.Create(&course)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (co Course) Find(db *gorm.DB) ([]Course, error) {

	courses := []Course{}
	result := db.Where(Course{Code: co.Code, Name: co.Name, Credits: co.Credits}).Find(&courses)
	return courses, result.Error

}
func (co Course) ConnectStudent(db *gorm.DB, user User) (bool, error) {
	err := db.Model(&user).Association("Courses").Append([]Course{co})
	if err != nil {
		return false, err
	}
	return true, err
}

func (co Course) DeleteConnect(db *gorm.DB, user User) (bool, error) {
	err := db.Model(&user).Association("Courses").Delete([]Course{co})
	if err != nil {
		return false, err
	}
	return true, err
}

func (user *User) FindNoConnect(db *gorm.DB) ([]Course, error) {

	allCourses, err := Course{}.Find(db)
	if err != nil {
		return nil, fmt.Errorf("查询所有课程失败: %v", err)
	}

	if err := db.Preload("Courses").Find(user).Error; err != nil {
		return nil, fmt.Errorf("查询用户课程失败: %v", err)
	}

	selected := make(map[uint]bool)
	for _, c := range user.Courses {
		selected[c.ID] = true
	}

	var availableCourses []Course
	for _, course := range allCourses {
		if !selected[course.ID] {
			availableCourses = append(availableCourses, course)
		}
	}

	return availableCourses, nil
}
func (user *User) FindOwn(db *gorm.DB) (bool, error) {
	result := db.Preload("Courses").Find(user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func Add(db *gorm.DB, account uint, password string, name string) (bool, error) {
	var user User
	user = User{Account: account, Password: password, Name: name}
	result := db.Create(&user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func Find(db *gorm.DB, account uint, password string, name string) ([]User, error) {
	user := []User{}
	var result *gorm.DB
	if account != 0 {
		result = db.Where(&User{Account: account, Password: password}).Find(&user)
	} else {
		result = db.Where("name LIKE ?", "%"+name+"%").Find(&user)
	}
	fmt.Println(result)
	if result.RowsAffected == 0 {
		return user, errors.New("该账号不存在")
	}
	return user, result.Error
}
func ChangeStrtouint(account []string) []uint {
	var uintSlice []uint
	fmt.Println("111")
	for _, s := range account {
		sa := s[1 : len(s)-1]
		num, _ := strconv.Atoi(sa)
		fmt.Println(sa)
		uintSlice = append(uintSlice, uint(num))
	}
	return uintSlice
}
func Findmore(db *gorm.DB, account []string) ([]User, error) {
	user := []User{}
	var result *gorm.DB
	uintSlice := ChangeStrtouint(account)
	fmt.Println("111")
	result = db.Where("account in ?", uintSlice).Find(&user)
	return user, result.Error
}
func Update(db *gorm.DB, account uint, password string, name string) error {
	user, err := Find(db, account, "", "")
	fmt.Println(err)
	user[0].Name = name
	user[0].Password = password
	result := db.Save(&user[0])
	return result.Error

}
func Delete(db *gorm.DB, account []string) error {
	uintSlice := ChangeStrtouint(account)
	result := db.Where("account in ?", uintSlice).Delete(&User{})
	return result.Error
}

func CheckPassword(db *gorm.DB, account uint, password string) bool {
	_, err := Find(db, account, password, "")
	if account == 0 {
		fmt.Println("账号不能为空")
		return false
	}
	if password == "" {
		fmt.Println("密码不能为空")
		return false
	}
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}

//func main() {
//	//Xing := []string{"张", "王", "李", "曹", "曾"}
//	//Ming := []string{"一", "二", "三", "四", "五"}
//	dsn := "host=192.168.15.10 user=postgres password=postgres dbname=postgres port=25433 sslmode=disable TimeZone=Asia/Shanghai"
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		fmt.Println("failed to connect database", err)
//	}
//	var account uint
//	fmt.Scanf("%d", &account)
//	fmt.Println(Find(db, account, ""))
//	//db.AutoMigrate(&User{})
//	//user := []User{}
//	//
//	//for i := 0; i < 10; i++ {
//	//	a := uint(rand.Intn(900000) + 100000)
//	//	Xrand := rand.Intn(5)
//	//	Mrand := rand.Intn(5)
//	//	name := fmt.Sprint(Xing[Xrand], Ming[Mrand])
//	//	user = append(user, User{Account: a, Password: "123456", Name: name})
//	//}
//	//db.Create(&user)
//
//}
