package main

import (
	"fmt"
	"time"
)

func main() {

	for _, v := range `	独立寒秋，湘江北去，橘子洲头。
	看万山红遍，层林尽染；
	漫江碧透，百舸争流。
	鹰击长空，鱼翔浅底，
	万类霜天竞自由。
	怅寥廓，问苍茫大地，谁主沉浮？

	携来百侣曾游，
	忆往昔峥嵘岁月稠。
	恰同学少年，风华正茂；
	书生意气，挥斥方遒。
	指点江山，激扬文字，
	粪土当年万户侯。
	曾记否，到中流击水，浪遏飞舟！` {
		if v == '，' {
			time.Sleep(time.Millisecond * 335)
		} else if v == '。' {
			time.Sleep(time.Millisecond * 345)
		}
		time.Sleep(time.Millisecond * 231)
		fmt.Printf("%c", v)
	}
}
