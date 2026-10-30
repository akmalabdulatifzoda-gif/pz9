package repo

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Open подключается к PostgreSQL через GORM.
// TranslateError включает перевод ошибок драйвера в gorm.ErrDuplicatedKey и т.п.
func Open(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Silent), // не пишем SQL с хэшами в лог
	})
}
