package db

import (
	"context"

	"github.com/bigchange/go-practice/testcase/model"
	"github.com/go-xorm/xorm"
)

var db DB

type primaryDB struct {
	engine *xorm.Engine
}

func New(engine *xorm.Engine) DB {
	db = &primaryDB{engine}
	return db
}

func GetDB()  DB {
	return db
}

func SetDB(p DB) {
	db = p
}
func (p *primaryDB) GetData(ctx context.Context, i int64) (*model.M1, error) {
	panic("implement me")
}
