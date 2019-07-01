# The Mi programming language

<img align="right" src="images/m2.png" title="Mi logo" height=200>

Mi is a small percussion music programming language for fun.
I'm a big fan of music (percussion more precisely) and programming, why not merge the two ?
I don't want programming to replace music software because these software is much more suitable and offers much more tools for music but I find it fun to play music by coding.

- [Installation](#installation)
- [Compiling](#compiling)
- [Example](#example)
- [Syntax](#syntax)
    - [Durations of notes](#durations-of-notes)
    - [Octaves](#octaves)

## Installation

To install Mi:

`gem install mi`

And verifying that it works:

`mi --version`

## Compiling

To create a MIDI file:

`mi -o --midi sample.mi`

To create PDF file:

`mi -o --pdf sample.mi`

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

<img align="left" src="images/notes-durations.png" title="Durations-of-notes" height=150>

In Mi we must specify the duration of the note after this one with `.`:

- `1/4`it's the default duration
- `~~` 1 bar
- `~` 1/2 bar
- `.` 1/8 bar
- `..` 1/16 bar
- `...` 1/32 bar
- `....` 1/64 bar
- `.....` 1/128 bar

<br><br>

### Roadmap

- [x] Grammar
- [ ] Parser
    - [ ] Lexer
    - [ ] Parser