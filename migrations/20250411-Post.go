package main

import (
	"log"
	"sv_be_posts/config"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	db := config.DB

	sql := `
	CREATE TABLE IF NOT EXISTS posts (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(200),
		content TEXT,
		category VARCHAR(100),
		status VARCHAR(100) CHECK (status IN ('Publish', 'Draft', 'Thrash')),
		created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
	`

	if err := db.Exec(sql).Error; err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migration completed successfully!")
}
