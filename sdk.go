package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type SdkDB struct {
	*gorm.DB
}

func New(db *gorm.DB) *SdkDB {
	return &SdkDB{DB: db}
}

func (db *SdkDB) GenTable(table string) {
	g := gen.NewGenerator(gen.Config{
		OutPath: "sdk",
		Mode:    gen.WithDefaultQuery,
	})
	g.UseDB(db.DB)
	g.GenerateModel(table)
	g.Execute()
}

func NewGorm(host, user, password, database, timeZone, schema string, port int) (*gorm.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s search_path=%s client_encoding=UTF8",
		host, port, user, password, database, timeZone, schema)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get *sql.DB from GORM: %v", err)
	}
	err = sqlDB.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}
	fmt.Println("Connected to Gorm is set up.")
	return db, nil
}
