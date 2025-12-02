package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

/**
 * error类型: 数据库连接失败，文件读写错误
 * defer语句: 资源清理(关闭文件，释放锁，数据库连接)
 * panic/recover: 模拟程序崩溃，捕获并处理panic
 */
func main() {
	//userName, err := queryDatabase(100)
	//
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(userName)
	//}

	err := readFile("C:\\Users\\PC\\Desktop\\workspace\\golang-workspace\\golang-practice\\internal\\base\\1-7-error-handler\\test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("文件读取成功")
	}
}

type BusinessError struct {
	Code    int
	Message string
	Time    time.Time
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Time: %s", e.Code, e.Message, e.Time.Format(time.RFC3339))
}

func queryDatabase(userID int) (string, error) {
	if userID <= 0 {
		return "", &BusinessError{
			Code:    1001,
			Message: "用户不存在",
			Time:    time.Now(),
		}
	}

	if userID == 999 {
		return "", errors.New("数据库连接超时")
	}

	return "用户数据", nil
}

func readFile(fileName string) error {
	// open file
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("打开文件失败: %w", err)
	}
	// close os resource anywhere
	defer file.Close()

	// read file
	buf := make([]byte, 100)
	_, readErr := file.Read(buf)
	if readErr != nil {
		return fmt.Errorf("读取文件失败: %w", err)
	}

	return nil
}
