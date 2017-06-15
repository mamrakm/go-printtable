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

//users struct contains data about users from /etc/passwd
type users struct {
	uid     int
	gid     int
	name    string
	homedir string
}

func getUsersList() []users {
	usr := make([]users, 32)
	pwd, err := os.Open("/etc/passwd")
	defer pwd.Close()
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
		usr = append(usr, users{uid: uid, gid: gid, name: data[0], homedir: data[5]})
	}
	return usr
}

func main() {
	res := getUsersList()
	fmt.Println(printUsers(res))
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

func printUsers(s []users) string {
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
		printLineSeparator(ret, 4*l+3)
		ret.WriteString("|")
		for _, line := range u {
			pad := strings.Repeat(" ", l-len(line))
			ret.WriteString(line + pad + "|")
		}
		ret.WriteString("\n")
		printLineSeparator(ret, 4*l+3)
	}
	return ret.String()
}

func printLineSeparator(buff *bytes.Buffer, i int) string {
	buff.WriteString("+")
	for l := 0; l < i; l++ {
		buff.WriteString("-")
	}
	buff.WriteString("+")
	buff.WriteString("\n")
	return buff.String()
}
