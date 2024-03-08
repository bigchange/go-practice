/*
 * @Author: Jerry You
 * @CreatedDate: 2023/4/17
 * @Last Modified by: lolaliva
 * @Last Modified time: 2023/4/17 15:48
 */
package designpatterns

import (
	"fmt"
	"testing"
)

var info = Info{}

// initialize
var loadFile = func(info *Info, err error) error {
	fmt.Println("loadFile file initialized")
	info.Name = "Hao Chen"
	info.Namespace = "MegaEase"
	info.OtherThings = "We are running as remote team."
	return nil
}

var NameVisitorFunc = func(info *Info, err error) error {
	if err != nil {
		return err
	}
	fmt.Println("NameVisitor() before call function")
	// 操作对象
	info.Name = "new name modified by nameVisitor"
	fmt.Println("NameVisitor() after call function")
	return err
}

// 操作日志： 一般放到最后，查看修改结果
var LogVisitorFunc = func(info *Info, err error) error {
	if err != nil {
		return err
	}
	fmt.Println("LogVisitor() before call function")
	// log
	fmt.Printf("==> Name=%s, NameSpace=%s, other=%s\n", info.Name, info.Namespace, info.OtherThings)
	fmt.Println("LogVisitor() after call function")
	return err
}

var OtherThingsVisitorFunc = func(info *Info, err error) error {
	if err != nil {
		return err
	}
	fmt.Println("OtherThingsVisitor() before call function")
	// 操作对象其他字段
	info.Namespace = "new namespace modified by OtherThingsVisitor"
	info.OtherThings = "new other things modified by OtherThingsVisitor"
	fmt.Println("OtherThingsVisitor() after call function")
	return err
}

func TestDecoratedVisitor_Visit(t *testing.T) {
	type fields struct {
		visitor    Visitor
		decorators []VisitorFunc
	}
	type args struct {
		fn VisitorFunc
	}
	var v Visitor = &info
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "visits-decorator", fields: fields{visitor: v, decorators: []VisitorFunc{NameVisitorFunc, OtherThingsVisitorFunc, LogVisitorFunc}}, args: args{fn: loadFile}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewDecoratedVisitor(tt.fields.visitor, tt.fields.decorators...)
			if err := v.Visit(tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("Visit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Visitor 嵌套使用
// 层次： v = OtherV(NameV(LogV(Info自身)))
//  v.Visit(f)
func TestNormalVisitor(t *testing.T) {
	var v Visitor = &info
	v = LogVisitor{v}
	v = NameVisitor{v}
	v = OtherThingsVisitor{v}
	v.Visit(loadFile)
}
