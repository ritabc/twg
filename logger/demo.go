package logger

import (
	"errors"
	"log"
	"os"
	"sync"
)

// DemoV1 has a dependency on a logger, but creates it itself
func DemoV1() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	err := doTheThing()
	if err != nil {
		logger.Println("error in doTheThing():", err)
	}
}

// DemoV2 has the same dependency, but it doesn't construct the logger itself, it gets passed in
func DemoV2(logger *log.Logger) {
	err := doTheThing()
	if err != nil {
		logger.Println("error in doTheThing():", err)
	}
}

// Implementation Agnostic Ex1: Functions
// DemoV3 is a way to make the function accept something more general than the very specific *log.Logger
// Call this with:
// 	loggger := log.New(...)
//  DemoV3(log.Println)
// This code doesn't care about the log package
func DemoV3(logFn func(...interface{})) {
	err := doTheThing()
	if err != nil {
		logFn("error in doTheThing():", err)
	}
}

// Implementation Agnostic Ex2: with Interfaces
type Logger interface {
	Println(...interface{})
	Printf(string, ...interface{})
}

// Call this with:
//   logger := log.New(...)
//   DemoV4(logger)
func DemoV4(logger Logger) {
	err := doTheThing()
	if err != nil {
		logger.Println("error in doTheThing():", err)
		// logger.Printf("error: %s\n", err)
	}
}

// Implementation Agnostic Ex3: with methods, get around passing in logger to Demo funcs
type Thing struct {
	Logger interface {
		Println(...interface{})
		Printf(string, ...interface{})
	}
}

func (t Thing) DemoV5() {
	err := doTheThing()
	if err != nil {
		t.Logger.Println("error in doTheThing():", err)
		// t.Logger.Printf("error: %s\n", err)
	}
}

// DI and useful zero values
// DemoV6 copied from DemoV4
func DemoV6(logger Logger) {
	// One way to be able to pass in nil
	if logger == nil {
		logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	}
	err := doTheThing()
	if err != nil {
		logger.Println("error in doTheThing():", err)
		// logger.Printf("error: %s\n", err)
	}
}

type ThingV2 struct {
	Logger interface {
		Println(...interface{})
		Printf(string, ...interface{})
	}
	once sync.Once
}

func (t *ThingV2) logger() Logger {
	// run the passed in function only once, even if its called multiple times (for ie multiple goroutines)
	t.once.Do(func() {
		if t.Logger == nil {
			t.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
		}
	})
	return t.Logger
}

func (t *ThingV2) DemoV7() {
	err := doTheThing()
	if err != nil {
		t.Logger.Println("error in doTheThing():", err)
		// t.Logger.Printf("error: %s\n", err)
	}
}
func doTheThing() error {
	return errors.New("error opening file: abc.txt")
}
