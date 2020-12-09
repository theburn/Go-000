package main

import (
	sql "database/sql" //just show
	"fmt"

	"github.com/pkg/errors"
)

type DAO struct{}

type Biz struct{}

type Service struct{}

func (dao *DAO) Query(sqlString string) (interface{}, error) {
	// just show
	tx := sql.Tx{}
	return tx.Exec(sqlString)
}

func (srv *Service) Do(dao *DAO) error {
	rs, err := dao.Query("select uid from user where uid = 'G20190282010006';")

	switch {
	case err == sql.ErrNoRows:
		return errors.Wrap(err, "The user can not found!")
	case err != nil:
		return errors.Wrap(err, "Unknown error!")
	default:
		return work(rs)
	}

}

func (biz *Biz) Do(dao *DAO) error {
	srv := Service{}
	err := srv.Do(dao)

	if errors.Is(err, sql.ErrNoRows) { // biz handler, not found user and only print.
		fmt.Println(err.Error())
		return nil
	}

	return err

}

func work(i interface{}) error {
	// just show
	fmt.Println(i)
	return nil
}

func main() {
	dao := DAO{}
	biz := Biz{}

	if err := biz.Do(&dao); err != nil {
		fmt.Println("Call Failed!")
	}

}
