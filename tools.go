package tools

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// OpenBrowser 打开浏览器
func OpenBrowser(url string) {
	cmd := exec.Command("main", "/c", "start", url)

	_, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	if err = cmd.Start(); err != nil { // 运行命令
		fmt.Println(err)
	}
}
func Input(prompt string) string {
	var input string
	for {
		fmt.Printf(prompt)
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("输入错误:", err)
			continue
		}
		if strings.TrimSpace(input) != "" {
			return input
		}
	}
}

func InputInteger(prompt string) int {
	for {
		if i, ok := strconv.Atoi(Input(prompt)); ok == nil {
			return i
		} else {
			fmt.Println("输入内容不是数字:", ok)
		}
	}
}

// InArray 判断字符串是否在数组中
func InArray(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func TestTime(year int, month int, day int) bool {
	localTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
	t1, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	t2, err := time.Parse("2006-01-02 15:04:05", localTime)
	return err == nil && t2.Before(t1)
}
