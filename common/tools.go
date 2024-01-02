package common

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var (
	confPath string      = filepath.Join(GetCurrPath(), "config.toml")
	config   Config      = GetConfig()
	client   http.Client = InitClient()
	version  string      = "v0.0.1"
)

// DefaultConfig 默认配置
var DefaultConfig = Config{
	DownloadUrl:     "https://zinc.docking.org/scaffolds/DWYHDSLIWMUSOO-UHFFFAOYSA-N/",
	PerPage:         100,
	DownloadDirName: filepath.Join(GetCurrPath(), "download"),
}

// InitClient 初始化 http client 配置
func InitClient() http.Client {
	client := http.Client{Timeout: 300 * time.Second}
	return client
}

// GetVersion 获取版本号
func GetVersion() string {
	return version
}

// GetConfig 获取项目的配置信息
func GetConfig() Config {
	var config Config
	// 如果配置文件存在
	if _, err := os.Stat(confPath); err == nil {
		// 读取配置文件
		if _, err := toml.DecodeFile(confPath, &config); err != nil {
			log.Fatal(err)
		}
	} else {
		config = DefaultConfig
	}
	return config
}

// GetCurrPath 获取当前路径
func GetCurrPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return path
}

// GenerateFetchURL 生成请求的url
func GenerateFetchURL(page int) string {
	url := config.DownloadUrl + "?count=" + strconv.Itoa(config.PerPage) + "&page=" + strconv.Itoa(page)
	return url
}

// GetReqContent 请求 url 获取数据, 返回 Content []byte
func GetReqContent(url string) []byte {
	var content []byte
	req, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
		return []byte{}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err, "close body error")
		}
	}(req.Body)

	content, err = io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err, "read body error")
		return []byte{}
	}
	return content
}

// CreateDir 创建目录
func CreateDir(dirPath string) {
	// os.MkdirAll()如果目录已存在会跳过创建,不会报错。
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		log.Fatal("Create dir error: ", err)
	}
}

// ClearDir 清空目录
func ClearDir(dirPath string) {
	// 读取目录下的所有文件
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal("read dir error: ", err)
	}
	// 遍历所有文件
	for _, file := range files {
		// 拼接文件路径
		filePath := filepath.Join(dirPath, file.Name())
		// 删除文件
		err := os.Remove(filePath)
		if err != nil {
			log.Fatal("remove file error: ", err)
		}
	}
}

// RemoveDir 删除目录
func RemoveDir(dirPath string) {
	err := os.RemoveAll(dirPath)
	if err != nil {
		log.Fatal("remove dir error: ", err)
	}
}

// RenameDir 重命名目录
func RenameDir(oldDirPath string, newDirPath string) {
	err := os.Rename(oldDirPath, newDirPath)
	if err != nil {
		log.Fatal("rename dir error: ", err)
	}
}

// CreateFile 创建文件
func CreateFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("create file error: ", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("close file error: ", err)
		}
	}(file)
}

// RemoveFile 删除文件
func RemoveFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Fatal("remove file error: ", err)
	}
}

// IsDirExist 目录是否存在
func IsDirExist(dirPath string) bool {
	_, err := os.Stat(dirPath)
	return err == nil || os.IsExist(err)
}

// SaveContentToSDF 保存 content 到 sdf 文件中
func SaveContentToSDF(value string, content []byte) {
	fileName := fmt.Sprintf("%s.sdf", value)
	filePath := filepath.Join(config.DownloadDirName, config.DownloadSubDirName, fileName)
	// 判断文件是否存在
	if _, err := os.Stat(filePath); err == nil {
		// 文件存在，直接返回
		fmt.Println("file is exist, file name is:", fileName)
		return
	}
	// 创建文件
	jsonFile, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(jsonFile)
	_, err = jsonFile.Write(content)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("save body to json file success, file name is:", fileName)
}

// GetFileList 获取文件列表
func GetFileList(dirPath string) []string {
	var fileList []string
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("read dir error:", err)
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			fileList = append(fileList, file.Name())
		}
	}
	return fileList
}
