package log

func Init(level, fileName string) {
	InitLogger(&Configuration{
		JSONFormat:      true,
		LogLevel:        "debug",
		StacktraceLevel: "fatal",
		Console:         &ConsoleConfiguration{},
		File: &FileConfiguration{
			Filename:   "./crm.log",
			MaxSize:    10,
			MaxAge:     14,
			MaxBackups: 10,
		},
	})
}
