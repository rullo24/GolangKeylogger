package main

type keyModifierPair struct {
	BaseKey     rune
	ModifiedKey rune
}

// Modifier Key Codes {virtualKeyCode: KeyPair} --> Based on US Keyboard Layout
var keyModifierMap = map[int]keyModifierPair {
	8: {8, 8}, // virtual key code for delete key
	9: {'\t', '\t'}, // virtual key code for tab
	13: {'\n', '\n'}, // virtual key code for return (next line)
	32:  {' ', ' '}, // virtual key code for spacebar
	49:  {'1', '!'}, // example virtual key code for 1
	50:  {'2', '@'}, // example virtual key code for 2
	51:  {'3', '#'}, // example virtual key code for 3
	52:  {'4', '$'}, // example virtual key code for 4
	53:  {'5', '%'}, // example virtual key code for 5
	54:  {'6', '^'}, // example virtual key code for 6
	55:  {'7', '&'}, // example virtual key code for 7
	56:  {'8', '*'}, // example virtual key code for 8
	57:  {'9', '('}, // example virtual key code for 9
	48:  {'0', ')'}, // example virtual key code for 0
	65:  {'a', 'A'}, // example virtual key code for a
	66:  {'b', 'B'}, // example virtual key code for b
	67:  {'c', 'C'}, // example virtual key code for c
	68:  {'d', 'D'}, // example virtual key code for d
	69:  {'e', 'E'}, // example virtual key code for e
	70:  {'f', 'F'}, // example virtual key code for f
	71:  {'g', 'G'}, // example virtual key code for g
	72:  {'h', 'H'}, // example virtual key code for h
	73:  {'i', 'I'}, // example virtual key code for i
	74:  {'j', 'J'}, // example virtual key code for j
	75:  {'k', 'K'}, // example virtual key code for k
	76:  {'l', 'L'}, // example virtual key code for l
	77:  {'m', 'M'}, // example virtual key code for m
	78:  {'n', 'N'}, // example virtual key code for n
	79:  {'o', 'O'}, // example virtual key code for o
	80:  {'p', 'P'}, // example virtual key code for p
	81:  {'q', 'Q'}, // example virtual key code for q
	82:  {'r', 'R'}, // example virtual key code for r
	83:  {'s', 'S'}, // example virtual key code for s
	84:  {'t', 'T'}, // example virtual key code for t
	85:  {'u', 'U'}, // example virtual key code for u
	86:  {'v', 'V'}, // example virtual key code for v
	87:  {'w', 'W'}, // example virtual key code for w
	88:  {'x', 'X'}, // example virtual key code for x
	89:  {'y', 'Y'}, // example virtual key code for y
	90:  {'z', 'Z'}, // example virtual key code for z
	186: {';', ':'},  // example virtual key code for ;
	187: {'=', '+'}, // example virtual key code for =
	188: {',', '<'},  // example virtual key code for ,
	189: {'-', '_'}, // example virtual key code for -
	190: {'.', '>'},  // example virtual key code for .
	191: {'/', '?'},  // example virtual key code for /
	192: {'`', '~'}, // example virtual key code for `
	219: {'[', '{'}, // example virtual key code for [
	220: {'\\', '|'}, // example virtual key code for \
	221: {']', '}'}, // example virtual key code for ]
	222: {'\'', '"'}, // example virtual key code for '
}	