package main

import (
    "database/sql"
    "fmt"
    _"github.com/ziutek/mymysql/godrv"
)

const(
    DB_NAME = "p2p"
    DB_USER = "p2p_dev"
    DB_PASS = "p2p_dev"
    DB_HOST = "127.0.0.1"
    DB_PORT = "3306"
)

type LoanInfo struct{
    ID uint64 `json:"id"`
}

func OpenDB() *sql.DB{
    db, err := sql.Open("mymysql",fmt.Sprintf("tcp:%s:%s*%s/%s/%s",DB_HOST, DB_PORT,DB_NAME, DB_USER, DB_PASS))
//    db, err := sql.Open("mymysql", "tcp:127.0.0.1:3306*p2p/p2p_dev/p2p_dev")

//      db, err := sql.Open("mymysql", "tcp:203.195.192.116:36000*p2p/p2p_dev/p2p_dev")

    if err != nil{
        panic(err)
    }
    return db
}

func main(){
    db := OpenDB()
    rows,err := db.Query("select FuiUid,FstrName from t_user_info")
    if err != nil{
        panic(err)
    }
    defer rows.Close()
    num,err := rows.Columns()
    fmt.Printf("Columns name:%v \n", num);

    for rows.Next(){
        var uid uint64
        var name string
        err := rows.Scan(&uid, &name)
        if err != nil{
            panic(err)
        }
        fmt.Printf("uid:%v name:%v\n", uid, name)
    }

    b := false;
    fmt.Println("b:", b);
}
