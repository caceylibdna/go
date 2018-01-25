package models

import (
	"database/sql"
	"errors"
	"fmt"
        _ "github.com/lib/pq"
)

type User struct {
	Id       int `PK`
	Username string
	Pwd      string
        LoadId   string
        CreTyp   string
        CrePassword  string
        CreUsername  string
        NetName      string
        Network      string
}

func pqconn(host string, port int, user string, password string, dbname string) *sql.DB{
        dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+
                  "password=%s dbname=%s sslmode=disable",
                  host, port, user, password, dbname)
        fmt.Println(dbinfo)
        db, err := sql.Open("postgres", dbinfo)
        if err != nil {
            panic(err)
        } else {
            return db
        }
}

func ValidateUser(user User) error {
        
        db := pqconn("192.168.11.16", 5432, "bdna", "bdna", "bdna")
        defer db.Close()

        var flag int
        flag = 0
        rows, err := db.Query("SELECT login_name, login_passwd FROM logins")
        if err != nil {
            panic(err)
        } 
        for rows.Next() {
            var login_name string
            var login_passwd string
            err = rows.Scan(&login_name, &login_passwd)
            if err != nil {
                panic(err)
            }
            fmt.Println(user.Username)
            fmt.Println(login_name)
            if user.Username == login_name && user.Pwd == login_passwd{
                flag = 1
            }

        }
        fmt.Println(flag)

	if flag == 0 {
		return errors.New("用户名或密码错误！")
	}
	return nil
}

func AddCreNetworkToTb(user User) error {

        db := pqconn("192.168.11.16", 5432, "bdna", "bdna", "bdna")
        defer db.Close()
        fmt.Println("testtest")
        var lastInsertId int
        err := db.QueryRow("insert into credentials(cre_type,cre_name,cre_passwd) values($1,$2,$3) returning cre_id;",user.CreTyp,"test","test").Scan(&lastInsertId)
        fmt.Println("NNNNOERROR")
        fmt.Println(err)
        if err != nil {
            fmt.Println("NNNNOERROR2")
            panic(err)            
        }
        
        return nil
}

