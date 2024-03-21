# TEXTUALITY
A very simple text editor written in go.
This is a learning project, intended to build confidence working with go following a project-based learning approach.

ChatGPT has been used as a search and learning tool but no code has been directly lifted. The goal is to understand and interpret the skeleton provided by the AI assistant, not use it as a crutch.

This project originally started as a code alone to this video tutorial and uses the general layout and approach defined by the creator: [Program a text editor in go - Part 1](https://www.youtube.com/watch?v=mVFXBZUBe2s&list=PLLfIBXQeu3aa0NI4RT5OuRQsLo6gtLwGN&index=1)
The github for the above tutorial can be found here:
[EGO repository](https://github.com/maksimKorzh/ego/tree/main)

## Current features
- Display buffer and file reading
- Status bar
  - Lines in file
  - Current file mode
  - file name and position
- Basic syntax highlighting
  - Character based highlighting
  - Non-varchar character highlighting
  - Mathematic operator highlighting
  - numeric highlighting

## Planned features
- Rich Syntax highlighting
  - Keyword awareness
  - comment highlighting
  - word based highlighting
- Copy andd paste (single step)
- Undo and Redo
- Insert and delete
- navigation


### Build from source
- You'll need go version 1.23.xx installed
```
cd Textuality
go build -o main.exe .
```

### Run locally
to run without declaring a file:

`.\main.exe`

this will create a file called out.txt and open it. A sample file is included in the project called test.txt. To open it simple run

`.\main.exe test.txt`