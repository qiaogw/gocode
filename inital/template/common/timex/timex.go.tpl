package timex

import "time"

const (
datetimeTemplate = "2006-01-02 15:04:05" //常规类型
dateTemplate     = "2006-01-02"          //其他类型
timeTemplate     = "15:04:05"            //其他类型
)

//TimeToDatetimeStr time 转换为 "2006-01-02 15:04:05"
func TimeToDatetimeStr(time time.Time) string {
return time.Format(datetimeTemplate)
}

//TimeToDateStr time 转换为 "2006-01-02"
func TimeToDateStr(time time.Time) string {
return time.Format(dateTemplate)
}

//TimeToTimeStr time 转换为 "15:04:05"
func TimeToTimeStr(time time.Time) string {
return time.Format(timeTemplate)
}

//DatetimeStrToTime "2006-01-02 15:04:05" 转换为 time
func DatetimeStrToTime(t string) time.Time {
stamp, _ := time.ParseInLocation(datetimeTemplate, t, time.Local)
return stamp
}

//DateStrToTime "2006-01-02 " 转换为 time
func DateStrToTime(t string) time.Time {
stamp, _ := time.ParseInLocation(dateTemplate, t, time.Local)
return stamp
}


