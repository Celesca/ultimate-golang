package main

import "fmt"

type (
	PostgresDb struct{}
	MockDb     struct{}
)

func InsertPlayerItemPostgres(db *PostgresDb) {
	db.Insert()
}

func (p *PostgresDb) Insert() {
	fmt.Println("Postgres Insert")
}

func InsertPlayerItemMock(db *MockDb) {
	db.Insert()
}

// เพราะเราไม่รู้ว่า function block มันจะเป็น human error ไม่ใช่หลักการที่ดี
func (m *MockDb) Insert() {
	fmt.Println("Mock Insert")
}

func main() {
	postgresDb := &PostgresDb{}
	InsertPlayerItemPostgres(postgresDb)

	mockDb := &MockDb{}
	InsertPlayerItemMock(mockDb)

}
