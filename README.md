# The Mi programming language

<img align="right" src="images/mi.png" title="Mi logo" height=150>

Mi is a small percussion music programming language for fun.
I'm a big fan of music (percussion more precisely) and programming, why not merge the two ?
I don't want programming to replace music software because these software is much more suitable and offers much more tools for music but I find it fun to play music by coding.

- [The Mi programming language](#the-mi-programming-langugae)
- [Installation](#installation)
- [Example](#example)
- [Syntax](#syntax)
    - [Durations of notes](#durations-of-notes)

## Installation

To install Mi:

```
git clone https://github.com/eagle453/Mi.git
cd Mi
go run main.go your-mi-file.mi
```

Mi will generate a midi file from your .mi file.

## Example

This is a simple syntax example for drums:

```
bpm:120

HC,K. HC.   HC,S. HC.
HC,K. HC,K. HC,S. HC.
HC,K. HC.   HC,S. HC.
HC,K. HC,K. HC,S. HC.
```

`bpm:x` for x BPM

`K` = kick, `S` = snare, `HC` = hi-hat closed

### Syntax

The Mi syntax is meant to be simple enough to understand.

## Durations of notes

In Mi we must specify the duration of the note after this one with `.`:

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
