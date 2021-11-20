package error_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

//应该wrap,因为根据业务来说,一般对数据源的操作错误都会导致业务中断,无法向下进行,所以只能抛出到最上层
//在dao层对异常进行wrap,然后返回最上层,在最上层对异常进行判断,打印出堆栈信息的日志,同时处理异常

func TestError(t *testing.T) {
	t.Log("hello")
	_, err := daoException()
	//errors.Is()
	if errors.Cause(err) == sql.ErrNoRows {
		//直接打印发现没有行号信息
		fmt.Printf("err value is :%T--- %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
	}

}

//模拟在dao层抛出error的函数
func daoException() (interface{}, error) {
	//在该值的基础上,可以带入错误的自定义信息,但是导致了上层对该error的类型无法做出准确的判断
	//return nil, fmt.Errorf("authenticate: %v", sql.ErrNoRows)
	//要对异常只做一次处理,可以避免额外的噪音,因为日志在日志中心中是断开的,无法进行连续的查询
	//fmt.Println(sql.ErrNoRows)
	return nil, errors.Wrap(sql.ErrNoRows, "sql error")
}

//在一个applications中可以选择 wrap error策略,具有最高可用性的包只能返回根错误值,例如Go标准库中的如sql.ErrNoRows。
