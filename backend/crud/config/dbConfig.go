package config

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func InitDB() {
	var err error
	connString := "server=DEVELOPER7\\MSSQLSERVER19;user id=sefa;password=sefa123;database=Encore;"
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("Veritabanı bağlantısı oluşturulamadı: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Veritabanına erişilemiyor: %v", err)
	}

	log.Println("Veritabanı bağlantısı başarıyla oluşturuldu")
}
