/**
 * @Author: Anpw
 * @Description:
 * @File:  file
 * @Version: 1.0.0
 * @Date: 2021/6/15 19:53
 */

package upload

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"selfblog/global"
	"selfblog/pkg/util"
	"strings"
)

type FileType int

const TypeImage FileType = iota + 1

//GetFileName
/**
 * @Author: Anpw
 * @Description: 将原始文件名加密处理
 * @param name
 * @return string
 */
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

//GetFileExt
/**
 * @Author: Anpw
 * @Description: 获取文件后缀
 * @param name
 * @return string
 */
func GetFileExt(name string) string {
	return path.Ext(name)
}

//GetSavePath
/**
 * @Author: Anpw
 * @Description: 获取文件保存地址
 * @return string
 */
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
				return true
			}
		}
	}
	return false
}

func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
