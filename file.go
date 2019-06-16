package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

type myfile struct {
	a int
	//os.File
}

var testfile myfile
func main() {
/*
	err := testfile.mkdir("testfile/ss", 0777)
	if err != nil {
		fmt.Println(err)
	}
*/
/*
	err := testfile.rmdirrf("testfile")
	if err != nil {
		fmt.Println(err)
	}
*/
/*
	newFile, err := testfile.create("./src/github/newfile.txt")

	if err != nil {
		log.Fatal(err)
		//fmt.Println(err)
	}
	newFile.Close()
	*/

//获取文件信息
	/*
	fileInfo, err := os.Stat("./src/github/newfile.txt")
	if err != nil && os.IsNotExist(err) {
		log.Fatal("文件不存在！")
	}
	fmt.Println(fileInfo.Name())
	*/

//重命名
	/*
	err := os.Rename("./src/github/newfile.txt", "./src/github/newfile.log")

	if err != nil {
		log.Fatal(err)
	}
	*/

//打开和关闭
	/*
	filePath := "./src/github/newfile.log"
	file,err := os.Open(filePath)

	if err != nil {
		 fmt.Println(err)
	}
	buf := make([]byte, 1024)
	for  {
		n,_ := file.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
	file.Close()
	*/

//截断文件
	/*
	filePath := "./src/github/newfile.log"
	err := os.Truncate(filePath, 1024)  //传0怎么表示删除文件
	if err != nil {
		fmt.Println(err)
	}
	*/

//复制文件
	/*
	filePath := "./src/github/newfile.log"
	originalFile,err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	newFile,_ := os.Create(filePath + "_copy")

	bytes,err := io.Copy(newFile, originalFile)

	fmt.Println(bytes)
	newFile.Sync()
	*/

	/*
	filePath := "./src/github/newfile.log"
	newFile,err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND,0777)
	defer newFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	ret,err := newFile.WriteString("xyz")
	fmt.Println(ret)
	newFile.Close()

	newFile,err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND,0777)


	fmt.Println(ioutil.ReadAll(newFile))
	os.Exit(0)

	buf := make([]byte,1024)
	for{
		n, _ := newFile.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])

	}
	*/


	logrus.WithFields(logrus.Fields{
		"message": "error message",
	}).Info("First message")


	os.Exit(0)


}



//方法
func (f *myfile) mkdir (dir string, perm os.FileMode) (err error) {
	err = os.MkdirAll(dir, perm)
	return
}

func (f *myfile) rmdir (dir string) (err error) {
	err = os.Remove(dir)
	return
}

func (f *myfile) rmdirrf (dir string) (err error) {
	err = os.RemoveAll(dir)
	return
}

func (f *myfile) create (filePath string) (newFile *os.File, err error) {
	newFile, err = os.Create(filePath)
	return
}






