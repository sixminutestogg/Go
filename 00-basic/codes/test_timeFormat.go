package main

import (
	"fmt"
	"time"
)

var (
	cst *time.Location
	uct *time.Location
	err error
)

const CSTLayout = "2006-01-02 15:04:05"

func main() {
	fmt.Println("在开发过程中，我们有时会遇到这样的问题，将 `2020-11-08T08:18:46+08:00` 转成 `2020-11-08 08:18:46`")

	const (
		ANSIC       = "Mon Jan _2 15:04:05 2006"
		UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
		RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
		RFC822      = "02 Jan 06 15:04 MST"
		RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
		RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
		RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
		RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
		RFC3339     = "2006-01-02T15:04:05Z07:00"
		RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
		Kitchen     = "3:04PM"
		// Handy time stamps.
		Stamp      = "Jan _2 15:04:05"
		StampMilli = "Jan _2 15:04:05.000"
		StampMicro = "Jan _2 15:04:05.000000"
		StampNano  = "Jan _2 15:04:05.000000000"
	)

	fmt.Println(ANSIC)

	fmt.Println(cst)
	fmt.Println(uct)
	fmt.Println(err)

	fmt.Println("CSTLayout China Standard Time Layout")

}

func init() {

	var err error
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	fmt.Println(location)
}

func RFC3339ToCSTLayout(value string) (string, error) {
	parse, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}
	return parse.In(cst).Format(CSTLayout), nil
}
