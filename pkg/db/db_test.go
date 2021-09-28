package db_test


import (
	"golearn/pkg/db"
	"fmt"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
 
)


func TestOuputExample(t *testing.T) {

	t.Run("Test ExampleTest success", func(t *testing.T) {

    	DB, mock, err := sqlmock.New()
    
		if err != nil {
			fmt.Println(err)
		}

		defer DB.Close()
		mock.ExpectExec("insert").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

		db.Db = DB
		_, err = db.Insert()
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		 
	})
	 
	t.Run("Test ExampleTest02 success", func(t *testing.T) {

    	DB, mock, err := sqlmock.New()
    
		if err != nil {
			fmt.Println(err)
		}

		defer DB.Close()

		rows := sqlmock.NewRows([]string{"Id", "Name", "Lname","Age"}).AddRow(1, "post 1", "hello",10)

		mock.ExpectQuery("select (.+) from tb_demo").WillReturnRows(rows)

		db.Db = DB
		_ = db.Select01_1()
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		 
	})
}