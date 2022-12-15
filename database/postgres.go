package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Postgres struct {
	PostgresDB *gorm.DB
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Connection() error {
	dsn := "host=localhost user=root password=1234 dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	p.PostgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return err
	}

	return nil
}
