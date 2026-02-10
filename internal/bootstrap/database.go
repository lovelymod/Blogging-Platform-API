package bootstrap

import (
	"blogging-platform-api/internal/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(env *ENV) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.SUPABASE_HOST,
		env.SUPABASE_USER,
		env.SUPABASE_PASSWORD,
		env.SUPABASE_DB,
		env.SUPABASE_PORT,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// สั่งสร้าง Table อัตโนมัติจาก Entity ที่เรานิยามไว้
	err = db.AutoMigrate(&entity.Blog{})
	if err != nil {
		log.Fatal("Migration Failed:", err)
	}

	return db
}
