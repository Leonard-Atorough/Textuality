package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

var mode int
var ROWS, COLS int
var offset_col, offset_row int
var current_col, current_row int
var source_file string
var text_buffer = [][]rune{}
var copy_buffer = [][]rune{}
var modified bool

func main() {
	run_editor()
}

func run_editor() {
	init_editor()
	set_source_file()
	run_editor_loop()
}

func run_editor_loop() {
	for {
		update_terminal_size()
		clear_and_display_buffer()
		display_status_bar()
		process_events()
	}
}

func display_status_bar() {
	panic("unimplemented")
}

func process_events() {
	event := termbox.PollEvent()
	switch event.Type {
	case termbox.EventKey:
		switch event.Key {
		case termbox.KeyEsc:
			termbox.Close()
		}
	}
}

func clear_and_display_buffer() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		fmt.Println(err)
	}
	display_text_buffer()
	termbox.Flush()

}

func update_terminal_size() {
	COLS, ROWS = termbox.Size()
	ROWS--
	if COLS < 78 {
		COLS = 78
	}
}

func set_source_file() {
	if len(os.Args) > 1 {
		source_file = os.Args[1]
		read_file(source_file)
	} else {
		source_file = "out.txt"
		text_buffer = append(text_buffer, []rune{})
	}
}

func init_editor() {
	err := termbox.Init()
	if err != nil {
		fmt.Println("Failed to initialize termbox.", err)
		os.Exit(1)
	}
}

// // func print_message(col, row int, fg, bg termbox.Attribute, msg string) {
// // 	for _, ch := range msg {
// // 		termbox.SetCell(col, row, ch, fg, bg)
// // 		col += runewidth.RuneWidth(ch)
// // 	}
// // }

func display_text_buffer() {
	var col, row int
	for row = 0; row < ROWS; row++ {
		text_bufferRow := row + offset_row
		for col = 0; col < COLS; col++ {
			text_bufferCol := col + offset_col
			if text_bufferRow >= 0 && text_bufferRow < len(text_buffer) && text_bufferCol < len(text_buffer[text_bufferRow]) {
				if text_buffer[text_bufferRow][text_bufferCol] != '\t' {
					termbox.SetChar(col, row, text_buffer[text_bufferRow][text_bufferCol])
				} else {
					termbox.SetCell(col, row, rune(' '), termbox.ColorDefault, termbox.ColorGreen)
				}
			} else if row+offset_row > len(text_buffer)-1 {
				termbox.SetCell(0, row, rune('*'), termbox.ColorBlue, termbox.ColorRed)
			}
		}
		termbox.SetChar(col, row, rune('\n'))
	}
}

func read_file(filename string) {
	data, err := os.Open(filename)
	if err != nil {
		source_file = filename
		text_buffer = append(text_buffer, []rune{})
		return
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		text_buffer = append(text_buffer, []rune{})

		for i := 0; i < len(line); i++ {
			text_buffer[lineNumber] = append(text_buffer[lineNumber], rune(line[i]))
		}
		lineNumber++
	}
}

func write_file(filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, row := range text_buffer {
		for _, r := range row {
			if _, err := writer.WriteRune(r); err != nil {
				fmt.Println(err)
				return
			}
		}
		// Write a newline character after each row
		if _, err := writer.WriteString("\n"); err != nil {
			fmt.Println(err)
			return
		}
	}
	writer.Flush()
}
