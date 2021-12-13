package db

import (
	"context"

	"github.com/bigchange/go-practice/testcase/model"
)




type DB interface {
	// mock funcs list here
	GetData(context.Context, int64) (*model.M1, error)
	// ...
}

