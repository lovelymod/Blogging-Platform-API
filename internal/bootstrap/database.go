package bootstrap

import (
	"blogging-platform-api/internal/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(config *entity.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.SUPABASE_HOST,
		config.SUPABASE_USER,
		config.SUPABASE_PASSWORD,
		config.SUPABASE_DB,
		config.SUPABASE_PORT,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// สั่งสร้าง Table อัตโนมัติจาก Entity ที่เรานิยามไว้
	err = db.AutoMigrate(&entity.Blog{}, &entity.Tag{}, &entity.User{}, &entity.RefreshToken{})
	if err != nil {
		log.Fatal("Migration Failed:", err)
	}

	return db
}
