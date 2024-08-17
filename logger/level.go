package logger

type Level int

const (
	Level_DEBUG_INFO Level = -1
	Level_INFO       Level = 0
	Level_WARNING    Level = 1
	Level_ERROR      Level = 2
)

var (
	Level_name = map[int]string{
		-1: "DEBUG_INFO",
		0:  "INFO",
		1:  "WARNING",
		2:  "ERROR",
	}
	Level_value = map[string]int{
		"DEBUG_INFO": -1,
		"INFO":       0,
		"WARNING":    1,
		"ERROR":      2,
	}
)

func (x Level) String() string {
	return Level_name[int(x)]
}

func (x Level) Number() int {
	return int(x)
}
