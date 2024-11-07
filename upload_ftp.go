package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jlaffaye/ftp"
)

func main() {
	// FTP 服务器信息
	ftpServer := "192.168.31.243"
	ftpUser := "ftpuser"
	ftpPass := "ftpuser"
	localFilePath := "E:/workspace/go_dev/test/file1.txt"
	// 在FTP服务器上的目标路径
	remoteFilePath := "/home/ftpuser/shared/file1.txt"

	// 连接到FTP服务器
	c, err := ftp.Dial(fmt.Sprintf("%s:%d", ftpServer, 21))
	if err != nil {
		log.Fatalf("Failed to connect to FTP server: %v", err)
	}
	defer c.Quit()

	// 登录到FTP服务器
	if err := c.Login(ftpUser, ftpPass); err != nil {
		log.Fatalf("Failed to login to FTP server: %v", err)
	}

	// 打开本地文件
	file, err := os.Open(localFilePath)
	if err != nil {
		log.Fatalf("Failed to open local file: %v", err)
	}
	defer file.Close()

	// 创建FTP存储器，用于存储文件
	// 创建FTP存储器，用于存储文件
	err = c.Stor(remoteFilePath, file)
	if err != nil {
		log.Fatalf("Failed to create FTP storer: %v", err)
	}

	fmt.Println("File uploaded successfully!")
}
