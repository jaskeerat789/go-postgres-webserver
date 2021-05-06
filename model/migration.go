package model

func (db *DB) Migrate() {
	db.DB.AutoMigrate(&Person{})
	db.DB.AutoMigrate(&Book{})
	db.Log.Info("Completed auto migrations")
}
