package utils

import (
	"fmt"

	"github.com/cheggaaa/pb/v3"
)

type Bar struct {
	pb *pb.ProgressBar
}

func NewBar(count int, MyStrStart, MyStrEnd string) *Bar {
	pb.RegisterElement("mochaGrey", pb.ElementFunc(func(state *pb.State, args ...string) string {
		for _, arg := range args {
			return fmt.Sprintf("%s%s%s", CatppuccinMochaGrey, arg, ColorReset)
		}
		return ""
	}), false)
	pb.RegisterElement("mochaGreen", pb.ElementFunc(func(state *pb.State, args ...string) string {
		for _, arg := range args {
			return fmt.Sprintf("%s%s%s", CatppuccinMochaGreen, arg, ColorReset)
		}
		return ""
	}), false)
	pb.RegisterElement("mochaYellow", pb.ElementFunc(func(state *pb.State, args ...string) string {
		for _, arg := range args {
			return fmt.Sprintf("%s%s%s", CatppuccinMochaYellow, arg, ColorReset)
		}
		return ""
	}), false)

	tmpl := fmt.Sprintf(`{{counters . | mochaYellow . }} {{ bar . "[" "-" (cycle . "↖" "↗" "↘" "↙" | mochaGreen . ) "_" "]" | mochaGrey . }} %s%s%s {{string . "MyStr" | mochaGreen .}} %s `, CatppuccinMochaYellow, MyStrStart, ColorReset, MyStrEnd)
	bar := pb.ProgressBarTemplate(tmpl).Start(count)
	return &Bar{pb: bar}
}

func (b *Bar) Grow(num int, MyStrVal string) {
	b.pb.Set("MyStr", MyStrVal).Add(num)
}

func (b *Bar) Done() {
	b.pb.Finish()
}
