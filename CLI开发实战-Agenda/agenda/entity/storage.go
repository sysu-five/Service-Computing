package entity

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"encoding/json"
)

var users []User

func UserJsonDecode(js []byte) User {
	var temp User
	err := json.Unmarshal(js, &temp)
	if err != nil {
		fmt.Println("Decode error")
	}
	return temp
}

func UserJsonEncode(user User) []byte {
	js,err := json.Marshal(user)
	if err != nil {
		fmt.Println("Encode error")
		os.Exit(1)
	}
	return js
}

func WriteUserToFile(user User) {
	file,err := os.OpenFile("data/User.txt",os.O_WRONLY|os.O_APPEND|os.O_CREATE,0600)
	if err != nil {
		fmt.Println("Open file error", err.Error())
		os.Exit(1)
	}
	file.WriteString(string(UserJsonEncode(user)))
	file.WriteString("\n")
	file.Close()
}

func ReadUserFromFile() []User {
	var temp []User
	file,err := os.OpenFile("data/User.txt",os.O_CREATE|os.O_RDONLY|os.O_APPEND,0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	for {
		line,err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		temp = append(temp,UserJsonDecode([]byte(line)))
	}
	return temp
}

func WriteHost(name string) {
	file,err := os.Create("data/Host.txt")
	if err != nil {
		fmt.Println("Open file error", err.Error())
		os.Exit(1)
	}
	file.WriteString(name)
	file.WriteString("\n")
	file.Close()
}

func ReadHost() []string {
	var temp []string
	file,err := os.OpenFile("data/Host.txt",os.O_CREATE|os.O_RDONLY|os.O_APPEND,0600)
	if err != nil {
		panic(err)
	}
	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		temp = append(temp,line)
	}
	file.Close()
	return temp
}

func Init() {
	temp_user := ReadUserFromFile()
	for i :=0; i < len(temp_user); i ++ {
		users = append(users, temp_user[i])
	}
}

func CheckUserExist(name string) bool {
	for i := 0; i < len(users); i ++ {
		if users[i].Name == name {
			return true
		}
	}
	return false
}

func RegisterUser(name string, password string, email string, telephone string) bool {
	var temp_user User
	err := false
	if CheckUserExist(name) {
		fmt.Println("This username has been used, please try another one!")
		err = true
	}
	if err {
		return false
	}
	temp_user.Name = name
	temp_user.Password = password
	temp_user.Email = email
	temp_user.Telephone = telephone
	users = append(users,temp_user)
	WriteUserToFile(temp_user)
	fmt.Println("Register successfully!")
	return true
}

func FindUser(name string) User {
	for i := 0; i < len(users); i ++ {
		if users[i].Name == name {
			return users[i]
		}
	}
	return User{"null","null","null","null"}
}

func LogOut() {
	os.Truncate("data/Host.txt", 0)
}

func QueryUser() []User {
	return users
}