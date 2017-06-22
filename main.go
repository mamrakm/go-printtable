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
	maxLengths := [4]int{0, 0, 0, 0}
	u := make([][4]string, len(s))
	for i, user := range s {
		u[i] = [4]string{
			strconv.Itoa(user.uid),
			strconv.Itoa(user.gid),
			user.name,
			user.homedir,
		}
		for j := range [...]int{0, 1, 2, 3} {
			if len(u[i][j]) > maxLengths[j] {
				maxLengths[j] = len(u[i][j])
			}
		}
	}
	sumLengths := maxLengths[0] + maxLengths[1] + maxLengths[2] + maxLengths[3]
	printLineSeparator(ret, sumLengths+3)
	ret.WriteString("|")
	for j, line := range [...]string{"UID", "GID", "User", "Homedir"} {
		pad := strings.Repeat(" ", maxLengths[j]-len(line))
		ret.WriteString(line + pad + "|")
	}
	ret.WriteString("\n")
	printLineSeparator(ret, sumLengths+3)
	for i := range u {
		if u[i][2] == "" {
			continue
		}
		ret.WriteString("|")
		for j, line := range u[i] {
			pad := strings.Repeat(" ", maxLengths[j]-len(line))
			ret.WriteString(line + pad + "|")
		}
		ret.WriteString("\n")
	}
	printLineSeparator(ret, sumLengths+3)
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
