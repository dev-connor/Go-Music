package dblayer

import (
	"backend/src/src/models"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con)
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

func (db *DBORM) GetPromos() (products models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

func (db *DBORM) GetCustomerByName(firstname, lastname string) (customer models.Customer, err error) {
	return customer, db.Where(&models.Customer{FirstName: firstname, LastName: lastname}).Find(&customer).Error
}

func (db *DBORM) GetCustomerByID(id int) (customer models.Customer, err error) {
	return customer, db.First(&customer, id).Error
}

func (db *DBORM) GetProduct(id int) (product models.Product, err error) {
	return product, db.First(&product, id).Error
}

func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {
	hashPassword(&customer.Pass)
	customer.LoggedIn = true
	err := db.Create(&customer).Error
	customer.Pass = ""
	return customer, err
}

func hashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}
	// bcrypt 패키지에서 사용할 수 있게 패스워드 문자열을 바이트 슬라이스로 변환한다.
	sBytes := []byte(*s)
	// GenerateFromPassword() 메서드는 패스워드 해시를 반환한다.
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// 패스워드 문자열을 해시 값으로 바꾼다.
	*s = string(hashedBytes[:])
	return nil
}

func (db *DBORM) SignInUser(email, pass string) (customer models.Customer, err error) {
	// 사용자 행을 나타내는 *gorm.DB 타입 할당
	result := db.Table("Customers").Where(&models.Customer{Email: email})
	// 입력된 이메일로 사용자 정보 조회
	err = result.First(&customer).Error
	if err != nil {
		return customer, err
	}
	// 패스워드 문자열과 해시 값 비교
	if !checkPassword(customer.Pass, pass) {
		// 불일치 시 에러 반환
		return customer, ErrINVALIDPASSWORD
	}
	// 공유되지 않도록 패스워드 문자열은 지운다.
	customer.Pass = ""
	// loggedin 필드 업데이트
	err = result.Update("loggedin", 1).Error
	if err != nil {
		return customer, err
	}
	// 사용자 행 반환
	return customer, result.Find(&customer).Error
}

func checkPassword(existingHash, incomingPass string) bool {
	// 해시와 패스워드 문자열이 일치하지 않으면
	// 아래 메서드는 에러를 반환한다.
	return bcrypt.CompareHashAndPassword([]byte(existingHash),
		[]byte(incomingPass)) == nil
}

func (db *DBORM) SignOutUserById(id int) error {
	// ID 에 해당하는 사용자 구조체 생성
	customer := models.Customer{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	// 사용자의 상태를 로그아웃 상태로 업데이트한다.
	return db.Table("Customers").Where(&customer).Update("loggedin", 0).Error
}

func (db *DBORM) GetCustomerOrdersByID(id int) (orders []models.Order, err error) {
	return orders, db.Table("orders").Select("*").Joins("join customers on customers.id = customer_id").Joins("join products on products.id = product_id").Where("customer_id=?", id).Scan(&orders).Error
}
