package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

var mode int
var ROWS, COLS int
var offset_col, offset_row int
var current_col, current_row int
var source_file string
var text_buffer = [][]rune{}
var copy_buffer = [][]rune{}
var undo_buffer = [][]rune{}
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
	statusBar := new(StatusBar)
	if mode > 0 {
		statusBar.mode_status = "EDIT: "
	} else {
		statusBar.mode_status = "VIEW: "
	}
	statusBar.file_status = source_file[:min(8, len(source_file))] + " - " + strconv.Itoa(len(text_buffer)) + " lines"
	if modified {
		statusBar.file_status += " modified"
	} else {
		statusBar.file_status += " saved"
	}
	statusBar.cursor_status = " Row " + strconv.Itoa(current_row+1) + ", Col " + strconv.Itoa(current_col+1) + " "
	if len(copy_buffer) > 0 {
		statusBar.copy_status = " [copy] "
	}
	if len(undo_buffer) > 0 {
		statusBar.undo_status = " [undo] "
	}
	used_space := len(statusBar.mode_status) + len(statusBar.file_status) + len(statusBar.cursor_status) + len(statusBar.copy_status) + len(statusBar.undo_status)
	spaces := strings.Repeat(" ", used_space)
	message := statusBar.GetStatusString(spaces)
	print_message(0, ROWS, termbox.ColorBlack, termbox.ColorWhite, message)
	termbox.Flush()
}

func process_events() {
	event := termbox.PollEvent()
	switch event.Type {
	case termbox.EventKey:
		switch event.Key {
		case termbox.KeyEsc:
			write_file(source_file)
			termbox.Close()
		}
	}
}

func clear_and_display_buffer() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		fmt.Println(err)
	}
	tokenedBuffer := tokenize(text_buffer)
	var col, row int
	for row = 0; row < ROWS; row++ {
		text_bufferRow := row + offset_row
		for col = 0; col < COLS; col++ {
			text_bufferCol := col + offset_col
			if text_bufferRow >= 0 && text_bufferRow < len(tokenedBuffer) && text_bufferCol < len(text_buffer[text_bufferRow]) {
				for _, character := range tokenedBuffer[text_bufferRow] {
					characterStyle := highlight(character)
					for _, v := range character.Value {
						termbox.SetCell(col, row, v, characterStyle.foreground, characterStyle.background)
						col++
					}
				}
			} else if row+offset_row > len(text_buffer)-1 {
				termbox.SetCell(0, row, rune('*'), termbox.ColorBlue, termbox.ColorRed)
			}
		}
		termbox.SetChar(col, row, rune('\n'))
	}
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

func print_message(col, row int, fg, bg termbox.Attribute, msg string) {
	for _, ch := range msg {
		termbox.SetCell(col, row, ch, fg, bg)
		col += runewidth.RuneWidth(ch)
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

type StatusBar struct {
	mode_status   string
	file_status   string
	copy_status   string
	undo_status   string
	cursor_status string
}

type CharacterStyle struct {
	foreground termbox.Attribute
	background termbox.Attribute
}

func highlight(token Token) CharacterStyle {
	switch token.Type {
	case "TokenTypeDigit":
		return CharacterStyle{termbox.ColorRed, termbox.ColorDefault}
	case "TokenTypeCharacter":
		if is_keyword(token.Value) {
			return CharacterStyle{termbox.ColorMagenta, termbox.ColorBlack}
		} else if is_type_keyword(token.Value) {
			return CharacterStyle{termbox.ColorCyan, termbox.ColorBlack}
		} else {
			return CharacterStyle{termbox.ColorDefault, termbox.ColorBlack}
		}
	case "TokenTypeSymbol":
		return CharacterStyle{termbox.ColorWhite, termbox.ColorDefault}
	case "TokenTypeMath":
		return CharacterStyle{termbox.ColorGreen, termbox.ColorDefault}
	case "TokenTypePunct":
		return CharacterStyle{termbox.ColorYellow, termbox.ColorDefault}
	case "TokenTypeOther":
		return CharacterStyle{termbox.ColorDefault, termbox.ColorDefault}
	default:
		return CharacterStyle{termbox.ColorDefault, termbox.ColorDefault}
	}
}

func is_keyword(token string) bool {
	keywords := []string{
		"false", "NaN", "none", "break", "byte",
		"case", "catch", "class", "const", "continue", "def", "do",
		"elif", "else", "else:", "enum", "export", "extends", "extern",
		"finally", "float", "for", "from", "func", "function",
		"global", "if", "import", " in", "lambda", "try", "except",
		"nil", "not", "null", "pass", "print", "raise", "return",
		"self", "short", "signed", "sizeof", "static", "struct", "switch",
		"this", "throw", "throws", "true", "True", "typedef", "typeof",
		"undefined", "union", "unsigned", "until", "var", "void",
		"while", "with", "yield",
	}
	return slices.Contains(keywords, strings.ToLower(token))
}

func is_type_keyword(token string) bool {
	typeKeywords := []string{
		"int", "double", "decimal", "long",
		"string", "char", "rune",
	}
	return slices.Contains(typeKeywords, strings.ToLower(token))
}

func (m StatusBar) GetStatusString(spaces string) string {
	return m.mode_status + m.file_status + m.copy_status + m.undo_status + spaces + m.cursor_status
}
