package date

import "time"

const (
	DEFAULT_FORMAT_A string = "2006-01-02 15:04:05"
	DEFAULT_FORMAT_B string = "2006-01-02"
	DEFAULT_FORMAT_C string = "20060102150405"
)

/**
  日期格式化处理B
*/

/**
获取当前时间
*/
func GetNowTime() time.Time {
	now := time.Now()
	return now
}

func GetNowTimeString() string {
	now := time.Now()
	nowStr := now.Format(DEFAULT_FORMAT_A)
	return nowStr
}

func GetNowDateString() string {
	now := time.Now()
	nowStr := now.Format(DEFAULT_FORMAT_B)
	return nowStr
}

func GetNowAString() string {
	now := time.Now()
	nowStr := now.Format(DEFAULT_FORMAT_C)
	return nowStr
}
