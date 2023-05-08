package logger

import "log"

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC)
}

type Logger struct {
	name string
}

func NewLogger(name string) *Logger {
	return &Logger{
		name: name,
	}
}

func (l Logger) Print(args ...any) {
	log.Print(append([]any{l.name + ":"}, args...)...)
}

func (l Logger) Println(args ...any) {
	log.Println(append([]any{l.name + ":"}, args...)...)
}

func (l Logger) Printf(args ...any) {
	if format, ok := args[0].(string); ok {
		log.Printf("%s: "+format, append([]any{l.name + ":"}, args[1:]...)...)
	} else {
		panic("First argument Logger.Printf call should be the format string")
	}
}

func (l Logger) Fatal(args ...any) {
	log.Fatal(append([]any{l.name + ":"}, args...)...)
}

func (l Logger) Fatalln(args ...any) {
	log.Fatalln(append([]any{l.name + ":"}, args...)...)
}

func (l Logger) Fatalf(args ...any) {
	if format, ok := args[0].(string); ok {
		log.Fatalf("%s: "+format, append([]any{l.name + ":"}, args[1:]...)...)
	} else {
		panic("First argument Logger.Fatalf call should be the format string")
	}
}
