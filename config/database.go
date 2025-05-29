package config

import (
	"log"
	"os"

	"github.com/V-enekoder/backend-housing-app/src/schema"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dbName := os.Getenv("DB_NAME")

	if dbName == "" {
		dbName = "default_app.db"
		log.Printf("Advertencia: La variable de entorno DB_NAME no está definida. Usando '%s' por defecto.\n", dbName)
	}

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	var gormErr error
	DB, gormErr = gorm.Open(sqlite.Open(dbName), gormConfig)

	if gormErr != nil {
		log.Fatalf("Error al conectar con la base de datos SQLite (%s): %v", dbName, gormErr)
	}

	log.Printf("Conectado exitosamente a la base de datos SQLite: %s\n", dbName)
}

func SyncDB() {
	if DB == nil {
		log.Fatal("Error: La conexión a la base de datos no ha sido inicializada. Llama a ConnectDB() primero.")
	}

	log.Println("Sincronizando modelos con la base de datos SQLite...")

	modelsToMigrate := []interface{}{
		&schema.User{},
		&schema.Category{},
		&schema.Project{},
		&schema.Model{},
	}

	err := DB.AutoMigrate(modelsToMigrate...)
	if err != nil {
		log.Fatalf("Error durante la AutoMigración de GORM: %v", err)
	}
	log.Println("Modelos sincronizados exitosamente con SQLite.")
}
