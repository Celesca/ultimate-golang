package main

import "fmt"

type (
	Database interface {
		Insert()
	}

	PostgresDb struct {
		Database
	}
	MockDb struct {
		Database
	}
)

// Constructor return เป็น Interface
func NewPostgresDatabase() Database {
	return &PostgresDb{}
}

func (p *PostgresDb) Insert() {
	fmt.Println("Postgres Insert")
}

func NewMockDatabase() Database {
	return &MockDb{}
}

func (m *MockDb) Insert() {
	fmt.Println("Mock Insert")
}

// ไม่สนว่าคุณใช้ Database อะไร แค่เราใช้ Interface ขอแค่ Insert มีไหม
func InsertPlayerItem(db Database) {
	db.Insert()
}

func main() {
	postgresDb := NewPostgresDatabase()
	InsertPlayerItem(postgresDb)

	mockDb := NewMockDatabase()
	InsertPlayerItem(mockDb)

}
