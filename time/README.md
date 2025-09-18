# golang time工具类包

+ GetCurrentDateTime 获取当前时间，格式为：2006-01-02 15:04:05
+ GetCurrentTime 获取当前时间，格式为：2006-01-02
+ GetCurrentTimestamp 获取当前时间戳，秒级 eg:1758192515
+ GetCurrentTimestampMilli 获取当前时间戳，毫秒级 eg:1758192515157
+ TimestampToDateTime 时间转换:时间戳 --> YYYY-MM-DD HH:mm:ss
+ TimestampToDate 时间转换：时间戳 --> YYYY-MM-DD
+ DateTimeToTimestamp 时间转换：YYYY-MM-DD HH:mm:ss --> 时间戳
+ FormatYYYYMMDD 时间格式转换:YYYYMMDD -> YYYY-MM-DD eg: 20060102 --> 2006-01-02
+ FormatYYYYMMDDReverse 时间格式转换:YYYY-MM-DD -> YYYYMMDD eg: 2006-01-02 --> 20060102
+ GetFirstAndLastDayOfMonth 获取指定YYYY-MM的第一天和最后一天 eg: 2025-09 --> 2025-09-01,2025-09-30
+ GetFirstAndLastDateTimeOfMonth 获取指定YYYY-MM的第一天0时0分0秒和最后一天23时59分59秒的YYYY-MM-DD HH:mm:ss
+ GetFirstAndLastTimestampOfMonth 获取指定YYYY-MM的第一天0时0分0秒和最后一天23时59分59秒的时间戳
