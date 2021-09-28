package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

type DataRow struct{
	Id int
	Name string
	Lname string
	Age int
}

var (
	Db *sql.DB
	ctx context.Context
)

func init(){
	connString := "server=192.168.1.105;user id=sa;password=P@ssw0rd;port=1433;database=GmFitsRepoUAT_THB_OLD;encrypt=disable;"

    db , err := sql.Open("mssql",connString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err := db.Ping(); err != nil {
		fmt.Println(err.Error())
		return
	}
	// ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// defer cancel()

	// status := "up"
	// if err := db.PingContext(ctx); err != nil {
	// 	status = "down"
	// }
 
	fmt.Println("Success")

	Db = db
}


// one field
func Select01() *DataRow {
	var cnt int
	if err :=  Db.QueryRow("select count(1) from tb_demo").Scan(&cnt); err != nil {
		fmt.Println(err)
		return nil
	}
	//fmt.Println(cnt)
	return new(DataRow)
}

// one row
func Select01_1() *DataRow {
	var data DataRow
	if err :=  Db.QueryRow("select * from tb_demo").Scan(&data.Id, &data.Name, &data.Lname, &data.Age); err != nil {
		fmt.Println(err)
		return nil
	}
	//fmt.Println(data)
	return &data
}

// n row
func Select01_2() ([]DataRow, error) {

	var list []DataRow
	
	// if err :=  Db.QueryRow("select ID, Name, Lname, Age from tb_demo").Scan(&data.Id, &data.Name, &data.Lname, &data.Age); err != nil {
	// 	fmt.Println(err)
	// 	return nil
	// }
   
	rows , err := Db.Query("select ID, Name, Lname, Age from tb_demo")
    if err != nil {
		fmt.Println(err)
		return nil, err	
	}

	defer rows.Close()

	for (rows.Next()){
		var data DataRow
		if err := rows.Scan(&data.Id, &data.Name, &data.Lname, &data.Age); err != nil {
			fmt.Println(err)
     		return nil, err	
		}

		list = append(list,data)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
     	return nil, err	
	}

	return list, nil
}

// n result set
func Select01_3() ([]DataRow, error) {
	var list []DataRow

	rows , err := Db.Query("select ID, Name, Lname, Age from tb_demo;select ID, Name, Lname, Age from tb_demo demo2; ")
    if err != nil {
		fmt.Println(err)
		return nil, err	
	}

	defer rows.Close()

	for (rows.Next()){
		var data DataRow
		if err := rows.Scan(&data.Id, &data.Name, &data.Lname, &data.Age); err != nil {
			fmt.Println(err)
     		return nil, err	
		}

		list = append(list,data)
	}

	rows.NextResultSet()

	for (rows.Next()){
		var data DataRow
		if err := rows.Scan(&data.Id, &data.Name, &data.Lname, &data.Age); err != nil {
			fmt.Println(err)
     		return nil, err	
		}

		data.Name = "Set 2 " + data.Name
		list = append(list,data)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
     	return nil, err	
	}

	return list, nil
}


func Insert() (int64, error) {

	sql := "insert into tb_demo(Name, Lname, Age) values (?,?,?);"

	result, err := Db.Exec(sql, "insert A", "insert B", 55)
	if err != nil {
       fmt.Println(err)
	   return -1, err
	}

	// if  result.LastInsertId() is not support, should use select ID = convert(bigint, SCOPE_IDENTITY())") and scan value
	return result.RowsAffected()
}

func Update() (int64, error) {

	sql := "update tb_demo set Name = ? where Id = ?"
	
	result, err := Db.Exec(sql,"AAAAA" , 2)
	if err != nil {
		fmt.Println(err)
		return -1, err
	 }

	 return result.RowsAffected()
}

func Delete()  (int64, error) {

	sql := "delete from tb_demo where Id = ?"
	
	result, err := Db.Exec(sql , 2)
	if err != nil {
		fmt.Println(err)
		return -1, err
	 }
	 return result.RowsAffected()
}

func ExecTrans()(int64, error){

	ctx = context.Background()

    tx, err := Db.BeginTx(ctx,nil)
	if err != nil {
		return -1, err
	}

	//defer tx.Rollback()

	_ , err = tx.ExecContext(ctx, "insert into tb_demo(Name,Lname,age) values (?,?,?)" , "ZZZ","ZZZ",60)
	if err !=nil {
		tx.Rollback()
		return -1, err
	}


	if err = tx.Commit(); err != nil {
        return -1, err
    }

   return 1 , nil

}

func ExStore01() {

}

func ExStore02() (*DataRow,error) {

    rows , err := Db.Query("exec Store01 ?", 555)
	if err != nil{
		return nil , err
	}

	var data DataRow

	for rows.Next(){
		err = rows.Scan(&data.Id)
		if err != nil{
			return nil , err
		}
	}
	return &data , nil
}



func Close() error {
	return Db.Close()
}