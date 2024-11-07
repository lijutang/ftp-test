package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jlaffaye/ftp"
)

func main() {
	// FTP 服务器信息
	ftpServer := "192.168.31.243"
	ftpUser := "ftpuser"
	ftpPass := "ftpuser"
	// 在FTP服务器上的目标路径
	remoteFilePath := "/home/ftpuser/shared/file1.txt"
	localFilePath := "E:/workspace/go_dev/test/ftp/downloadedfile.txt"

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

	// 从FTP服务器下载文件
	resp, err := c.Retr(remoteFilePath)
	if err != nil {
		log.Fatalf("Failed to retrieve file from FTP server: %v", err)
	}
	defer resp.Close()

	// 创建本地文件
	out, err := os.Create(localFilePath)
	if err != nil {
		log.Fatalf("Failed to create local file: %v", err)
	}
	defer out.Close()

	// 将FTP响应的内容复制到本地文件
	if _, err := io.Copy(out, resp); err != nil {
		log.Fatalf("Failed to write file to local disk: %v", err)
	}

	fmt.Println("File downloaded successfully!")
}
