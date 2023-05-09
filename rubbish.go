package oslib

import (
	"log"
	"os"
	"time"
)

func AddLog(ip string) {
	// 写入 ip_log.txt 文件
	f, err := os.OpenFile("ip_log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// 获取当前时间
	now := time.Now()
	// 写入文件
	if _, err := f.WriteString(now.Format("2006-01-02 15:04:05") + " " + ip + "\n"); err != nil {
		log.Fatal(err)
	}
}
