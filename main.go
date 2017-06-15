package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var arr = []string{
	"George", "Michael", "Robert", "Dennis", "Ian",
}

//Users struct contains data about users from /etc/passwd
type Users struct {
	uid     int
	gid     int
	name    string
	homedir string
}

func getUsersList() []Users {
	users := make([]Users, 32)
	pwd, err := os.Open("/etc/passwd")
	if err != nil {
		log.Println("Unable to open /etc/passwd: ", err)
	}
	buff := bufio.NewReader(pwd)
	for {
		line, err := buff.ReadString('\n')
		if err != nil {
			break
		}
		data := strings.Split(line, ":")
		uid, err := strconv.Atoi(data[2])
		if err != nil {
			log.Fatal("Unable to read UID: ", err)
		}
		gid, err := strconv.Atoi(data[3])
		if err != nil {
			log.Fatal("Unable to read GID: ", err)
		}
		users = append(users, Users{uid: uid, gid: gid, name: data[0], homedir: data[5]})
	}
	return users
}

func main() {
	res := getUsersList()
	fmt.Println(printUsers(res))
}

func printArr(in []string) string {
	max := maxArrayLen(in)
	var ret string
	buf := bytes.NewBufferString(ret)
	str := strconv.Itoa(max)
	_, err := buf.WriteString(str)
	if err != nil {
		log.Println("Unable to write string: ", err)
	}
	return buf.String()
}

func maxArrayLen(in []string) (max int) {
	max = 0
	for _, v := range in {
		if len(v) > max {
			max = len(v)
		}
	}
	return max
}

func printUsers(s []Users) string {
	var strbuff string
	ret := bytes.NewBufferString(strbuff)
	for _, user := range s {
		if user.name == "" {
			continue
		}
		u := []string{
			strconv.Itoa(user.uid),
			strconv.Itoa(user.gid),
			user.name,
			user.homedir,
		}
		l := maxArrayLen(u)
		printBox(ret, 4*l+3)
		ret.WriteString("|")
		for _, line := range u {
			pad := strings.Repeat(" ", l-len(line))
			ret.WriteString(line + pad + "|")
		}
		ret.WriteString("\n")
		printBox(ret, 4*l+3)
	}
	return ret.String()
}

func printBox(buff *bytes.Buffer, i int) string {
	buff.WriteString("+")
	for l := 0; l < i; l++ {
		buff.WriteString("-")
	}
	buff.WriteString("+")
	buff.WriteString("\n")
	return buff.String()
}
