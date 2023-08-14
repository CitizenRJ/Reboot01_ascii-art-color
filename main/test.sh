#!/bin/bash
echo "ascii-art:"
echo go run . "hello" 
echo go run . "HELLO" 
echo go run . "HeLlo HuMaN" 
echo go run . "1Hello 2There" 
echo go run . "Hello\nThere" 
echo go run . "Hello\n\nThere" 
echo go run . "{Hello & There #}" 
echo go run . 'hello There 1 to 2!' 
echo go run . "MaD3IrA&LiSboN" 
echo go run . "1a\"#FdwHywR&/()=" 
echo go run . "{|}~" 
echo go run . "[\]^_ 'a" 
echo go run . "RGB" 
echo go run . ":;<=>?@" 
echo go run . '\!" #$%&'"'"'()*+,-./' 
echo go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" 
echo go run . "abcdefghijklmnopqrstuvwxyz" 
echo "ascii-art-fs:"
echo go run . "hello" standard 
echo go run . "hello world" shadow 
echo go run . "nice 2 meet you" thinkertoy 
echo go run . "you & me" standard 
echo go run . "123" shadow 
echo go run . "/(\")" thinkertoy 
echo go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" shadow 
echo go run . "\"#$%&/()*+,-./" thinkertoy 
echo go run . "It's Working" thinkertoy
echo "ascii-art-output: "
echo go run . --output=banner.txt "hello" standard
echo go run . --output=banner.txt "Hello There!" shadow
