# The Mi programming language

<img align="right" src="images/mi.png" title="Mi logo" height=150>

Mi is a small percussion music programming language made for fun.
I'm a big fan of music (percussion more precisely) and programming, why not merging both?
I don't want programming to replace music software because these kind of software are much more suitable and offer much more tools for music manipulation; but I find it fun to play music from some code.

- [The Mi programming language](#the-mi-programming-language)
- [Installation](#installation-and-usage)
- [Example](#example)
- [Syntax](#syntax)
    - [Durations of notes](#durations-of-notes)
- [Roadmap](#roadmap)

## Dependencies

- Golang (any version)

## Installation and Usage

Installing Mi is a pretty simple 2-step procedure.

First of all, you have to run go get to pull the repository to your GO path.<br>
```go get github.com/eagle453/mi```<br><br>
Then you need to install `aurora`<br>
```go get -u github.com/logrusorgru/aurora```<br><br>
Then just simply run your `mi` file<br>
```go run mi.go your-mi-file.mi```

```bash
got get github.com/eagle453/mi
# Open the mi folder
go run main.go your-mi-file.mi
```

> Mi will generate a midi file from your `.mi` file.

## Example

This is a simple syntax example for drums:

```ruby
bpm:120

HC,K. HC.   HC,S. HC.
HC,K. HC,K. HC,S. HC.
HC,K. HC.   HC,S. HC.
HC,K. HC,K. HC,S. HC.
```

`bpm:x` for x BPM

`K` = kick, `S` = snare, `HC` = hi-hat closed

### Syntax

The syntax of Mi is meant to be simple enough to be understood easily.

## Durations of notes

In Mi, we must specify the duration of each note by appending one of these symbols:

- `1/4`it's the default duration
- `~~` 1 bar
- `~` 1/2 bar
- `.` 1/8 bar
- `..` 1/16 bar
- `...` 1/32 bar
- `....` 1/64 bar
- `.....` 1/128 bar

(you can find the duration of the notes in music in the folder `images` and the file `notes-durations.png`)

### Roadmap

- [x] Grammar
- [x] Parser
    - [ ] Optimize code and Mi syntax
