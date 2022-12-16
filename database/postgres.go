package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type Postgres struct {
	PostgresDB *gorm.DB
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Connection() error {
	dsn := "host=localhost user=root password=1234 dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn2 := "host=localhost user=root password=1234 dbname=gorm2 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//dsn3 := "host=localhost user=root password=1234 dbname=gorm3 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	p.PostgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	//dbresolver.

	p.PostgresDB.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{
			postgres.Open(dsn2),
		},
		Policy: dbresolver.RandomPolicy{},
	}, "read"))
	//or 이 방식으로 하면 repo 수정 안해도 find 시 Replicas 사용
	//Register에 Sources로 등록시 write 역활을
	/*
		p.PostgresDB.Use(dbresolver.Register(dbresolver.Config{
			Replicas: []gorm.Dialector{
				postgres.Open(dsn2),
			},
			Policy: dbresolver.RandomPolicy{},
		}).Register(dbresolver.Config{
			Sources: []gorm.Dialector{
				postgres.Open(dsn3),
			},
		}))
	*/
	if err != nil {
		return err
	}

	return nil
}
