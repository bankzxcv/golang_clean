package repository

// repository จัดการ data, ดึง data ยังไง ก็พอ ไม่ต้องสนใจเรื่อง business logic ใดๆ
// เช่นกำหนด data มี interface อะไรบ้าง

// adapter is struct
type Customer struct {
	CustomerID  int    `db:"customer_id"`
	Name        string `db:"name"`
	DataOfBirth string `db:"date_of_birth"`
	City        string `db:"city"`
	ZipCode     string `db:"zipcode"`
	Status      bool   `db:"status"`
}

// PORT is interface !!! -- ปลั้กนี้มีกี่ขา มี data อะไรบ้างมาเก็บใน repository นี้
type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error) // ใช้ pointer เพราะถ้าไม่เจอ จะ return nil ได้
}
