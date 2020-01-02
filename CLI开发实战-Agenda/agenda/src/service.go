package src

import (
	"fmt"
	"os"
	"log"
	"strings"
	"github.com/sysu-five/agenda/entity"
)

var logflag bool
var log_file *os.File
var host_name, host_password string

func IsLogin() bool {
	return logflag
}

func Init() {
	entity.Init()
	temp,err := os.OpenFile("data/agenda.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0600)
	log_file = temp
	if err != nil {
		log.Fatalln("Open file error")
	}
	host := entity.ReadHost()
	if len(host) == 0 {
		logflag = false
	} else {
		logflag = true
		host_name = strings.Replace(host[0],"\n","",-1)
	}
}

func RegisterUser(name string, password string, email string, telephone string) {
	debug_log := log.New(log_file,"[Operation]",log.LstdFlags)
	i := entity.RegisterUser(name,password,email,telephone)
	if i {
		debug_log.Println(name, "register successfully")
	} else {
		debug_log.Println(name, "register failed")
	}
	defer log_file.Close()
}

func Login(name string, password string) {
	debug_log := log.New(log_file,"[Operation]",log.LstdFlags)
	if entity.CheckUserExist(name) {
		host_name = name
		host_password = password
		temp_user := entity.FindUser(name)
		if temp_user.Password != password {
			debug_log.Println(name, "Login Failed! Wrong Password!")
			fmt.Println("Wrong Password!")
		} else {
			debug_log.Println(name, "Login Succeeded!")
			fmt.Println("Login Successfully! Welcome to Agenda!")
			entity.WriteHost(name)
		}
	} else {
		debug_log.Println(name, "Login Failed! No such user!")
		fmt.Println("The user does not exist!")
	}
	defer log_file.Close()
}

func LogOut() {
	debug_log := log.New(log_file,"[Operation]",log.LstdFlags)
	debug_log.Println(host_name, "Logout Successfully!")
	entity.LogOut()
	fmt.Println("Log out successfully!")
	defer log_file.Close()
}

func QueryUser() {
	debug_log := log.New(log_file,"[Operation]",log.LstdFlags)
	debug_log.Println(host_name, "Query user successfully!")
	userlist := entity.QueryUser()
	for i,u := range userlist {
		fmt.Println(i+1, u.Name, u.Email, u.Telephone)
	}
	defer log_file.Close() 	
}