package service

import (
	"gin-mall/conf"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadAvatarToLocalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId)) // 后面用于路径拼接
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg" //todo:把file后缀提取出来
	content, err := ioutil.ReadAll(file)       //读取上传文件的内容
	if err != nil {
		return "", err
	}
	//把读取到的内同传入avatarPath中，若文件不存在按照给定的权限创建文件
	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return
	}
	//返回的filePath是数据库存储的
	return "user" + bId + "/" + userName + ".jpg", nil
}

// DirExistOrNot 判断文件夹路径是否存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir() //返回s是否是一个目录
}

func UploadProductToLocalStatic(file multipart.File, userId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId)) // 后面用于路径拼接
	basePath := "." + conf.ProductPath + "boss" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := basePath + productName + ".jpg" //todo:把file后缀提取出来
	content, err := ioutil.ReadAll(file)           //读取上传文件的内容
	if err != nil {
		return "", err
	}
	//把读取到的内同传入avatarPath中，若文件不存在按照给定的权限创建文件
	err = ioutil.WriteFile(productPath, content, 0666)
	if err != nil {
		return
	}
	//返回的filePath是数据库存储的
	return "boss" + bId + "/" + productName + ".jpg", nil
}

// CreateDir 创建目录
func CreateDir(dirName string) bool {
	//创建目录
	err := os.MkdirAll(dirName, 755)
	if err != nil {
		return false
	}
	return true
}
