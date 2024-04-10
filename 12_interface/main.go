package main

import "fmt"

type (
	Database interface { // 2 methods
		Insert() error // return error type
		Update() error
	}

	PostgresDb struct{} // Real database

	MockDb struct{}
)

// ในกรณีใช้ร่วมกับ interface Db โดยอัตโนมัติ ถ้าเรากลัวเขียนไม่ครบ

// การ Implement Constructor ขึ้นมาเพื่อเช็คว่าเราทำครบรึยัง
func NewPostgresDb() Database {
	return &PostgresDb{}
}

func (p *PostgresDb) Insert() error {
	fmt.Println("Real Insert!!")
	return nil
}

func (p *PostgresDb) Update() error {
	fmt.Println("Real Update !!")
	return nil
}

func (m *MockDb) Insert() error {
	fmt.Println("Mock Insert!!")
	return nil
}

func (m *MockDb) Update() error {
	fmt.Println("Mock Update !!")
	return nil
}

// เราไม่สนใจว่าจะเป็น Db ไหน แต่เรารู้ว่ามันสามารถ Insert ได้
func InsertPlayerItem(db Database) error {
	return db.Insert()
}

func UpdatePlayerItem(db Database) error {
	return db.Update()
}

// ทำให้มันเป็น Interface เสมอ

func main() {

	postgres := &PostgresDb{}
	mock := &MockDb{}

	InsertPlayerItem(postgres)
	UpdatePlayerItem(postgres)

	InsertPlayerItem(mock)
	UpdatePlayerItem(mock)

}
