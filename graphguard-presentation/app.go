package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx               context.Context
	helloWorldChannel chan struct{}
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (app *App) Fibonacci(n uint64) uint64 {
	// if n < 2 {
	// 	return n
	// }
	// var a, b uint64 = 0, 1
	// for n--; n > 0; n-- {
	// 	a += b
	// 	a, b = b, a
	// }
	// return b
	i := uint64(0)
	for ; i < n; i++ {
	}
	return i
}

func (app *App) Events_HelloWorld_FireOnce() {
	runtime.EventsEmit(app.ctx, "HelloWorld", "Hello World!")
}

func (app *App) Events_HelloWorld_Start() {
	if app.helloWorldChannel != nil {
		return
	}

	c := make(chan struct{})
	app.helloWorldChannel = c
	rand.Seed(time.Now().UTC().UnixNano())
	go func(c <-chan struct{}) {
		for {
			duration := rand.Uint32() / 1_000_000
			time.Sleep(time.Millisecond * time.Duration(duration))
			runtime.EventsEmit(app.ctx, "HelloWorld", "Hello World!")
			select {
			case <-c:
				return
			default:
			}
		}
	}(c)
}

func (app *App) Events_HelloWorld_Stop() {
	close(app.helloWorldChannel)
	app.helloWorldChannel = nil
}
