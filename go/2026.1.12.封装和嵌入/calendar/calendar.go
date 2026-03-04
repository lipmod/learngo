package calendar

import (
	"errors"
	"fmt"
)

type Date struct {
	year  int
	month int
	day   int
	time  //time 是一个嵌入到Data内的类型
}

type time int //自定义类型可以是struct也可以是基础的这种类型,几乎所有的类型都可以

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) Year() int {
	return d.year
}

// 这个方法会被提升到Data,可以使用Data直接调用
func (e time) Settime() {
	fmt.Println("ok")
}

// 该方法不会被提升，外部无法使用
func (e time) settime() {
	fmt.Println("ok")
}
