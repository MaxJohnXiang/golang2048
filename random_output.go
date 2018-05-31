package main

import "github.com/shiyanlou/termbox-go"
import "math/rand"
import "time"


func draw() {
	w, h := termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y :=0 ; y < h; y ++ {
		for x := 0; x <w; x++ {
			termbox.SetCell(x,y,' ',termbox.ColorDefault,termbox.Attribute(rand.Int()%8)+1)
		}
	}
	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close() 
	event_quene := make(chan termbox.Event)
	go func() {
		for {
			event_quene <- termbox.PollEvent()
		}
	}()

	draw()

	for {
		select {
		case ev := <-event_quene:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				return 
			}
		default:
			draw()
			time.Sleep(10 * time.Millisecond)
		}
	}
}
