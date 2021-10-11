package db

import (
	"context"

	"github.com/bigchange/go-practice/testcase/model"
	"github.com/go-xorm/xorm"
)

var pdbInstance *primaryDB

type primaryDB struct {
	engine *xorm.Engine
}

func New(engine *xorm.Engine) DB {
	pdbInstance = &primaryDB{engine}
	return pdbInstance
}

func GetPrimaryDB()  DB {
	return pdbInstance
}

func (p *primaryDB) GetData(ctx context.Context, i int64) (*model.M1, error) {
	panic("implement me")
}
