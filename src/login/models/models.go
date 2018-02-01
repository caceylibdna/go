package models

import (
	"database/sql"
	"errors"
	"fmt"
        _ "github.com/lib/pq"
        "bytes"
        "github.com/astaxie/beego"
        "os/exec"
        "strings"
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

func dbconn(dbtpye string) *sql.DB{
        host := beego.AppConfig.String("postgresqlhost")
        port, err1 := beego.AppConfig.Int("postgresqlport")
        user := beego.AppConfig.String("postgresqluser")
        password := beego.AppConfig.String("postgresqlpass")
        dbname := beego.AppConfig.String("postgresqldb")
        if err1 != nil {
            panic(err1)
        }

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

        db := dbconn("postgresql")
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
		return errors.New("username or password is wrong!")
	}
	return nil
}

func AddCredentials(user User) error {

        db := dbconn("postgresql")
        defer db.Close()
        fmt.Println("testtest")
        var lastInsertId int
        err := db.QueryRow("insert into credentials(cre_type,cre_name,cre_passwd) values($1,$2,$3) returning cre_id;",user.CreTyp, user.CreUsername, user.CrePassword).Scan(&lastInsertId)

        if err != nil {
            fmt.Println("ERROR2")
            panic(err)
        }

        return nil
}

func AddNetworks(user User) error {

        db := dbconn("postgresql")
        defer db.Close()
        fmt.Println("testtest")
        var lastInsertId int
        err := db.QueryRow("insert into networks(network_name, network_network) values($1,$2) returning network_id;",user.NetName, user.Network).Scan(&lastInsertId)

        if err != nil {
            fmt.Println("ERROR2")
            panic(err)
        }
        
        return nil
}

func AddShifts(user User) string {

        db := dbconn("postgresql")
        defer db.Close()
        fmt.Println("testtest")
        var lastInsertId int
        err := db.QueryRow("insert into tmp_shifts(shift_name, shift_username, shift_password, shift_network) values($1,$2,$3,$4) returning shift_id;", "scantask", user.CreUsername, user.CrePassword, user.Network).Scan(&lastInsertId)

        if err != nil {
            fmt.Println("ERROR2")
            panic(err)
        }

        cmd := exec.Command("/root/demo/duster/src/login/scripts/getWin32Services.sh", user.CreUsername, user.CrePassword, user.Network)
        cmd.Stdin = strings.NewReader("some input")
        var out bytes.Buffer
        cmd.Stdout = &out
        cmd.Run()

        cmd1:= exec.Command("/root/demo/duster/src/login/scripts/getAppByService.sh", user.Network)
        cmd1.Stdin = strings.NewReader("some input")
        var out1 bytes.Buffer
        cmd1.Stdout = &out1
        cmd1.Run()
        fmt.Println("parsedata")
        fmt.Println(out1.String())

        return out1.String()
}

