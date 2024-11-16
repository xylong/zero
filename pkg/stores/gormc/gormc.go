package gormc

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

// NewEngine 实力化gorm
func NewEngine(c Config) *gorm.DB {
	if c.Separation == true { //启用读写分离
		return newRWDBEngine(c)
	}
	engine, err := gorm.Open(postgres.Open(c.DSN), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	if c.Debug {
		engine = engine.Debug()
	}

	sqlDB, err := engine.DB()

	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(c.MaxIdleConns)

	sqlDB.SetMaxOpenConns(c.MaxOpenConns)

	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	return engine
}

// 创建读写分离实例
func newRWDBEngine(c Config) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: c.Master,
	}))
	if err != nil {
		panic(err)
	}
	var replicas, sources []gorm.Dialector
	for _, v := range c.Replicas {
		replicas = append(replicas, postgres.New(postgres.Config{
			DSN: v,
		}))
	}
	for _, v := range c.Sources {
		sources = append(sources, postgres.New(postgres.Config{
			DSN: v,
		}))
	}
	err = db.Use(
		dbresolver.
			Register(dbresolver.Config{
				Sources:  sources,
				Replicas: replicas,
				Policy:   dbresolver.RandomPolicy{},
			}).
			SetMaxIdleConns(c.MaxIdleConns).
			SetConnMaxLifetime(time.Minute * 30).
			SetMaxOpenConns(c.MaxOpenConns),
	)
	if err != nil {
		panic(err)
	}
	return db
}
