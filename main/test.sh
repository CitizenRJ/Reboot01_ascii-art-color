#!/bin/bash
#!/bin/bash

echo "go run . "" | cat -e"
go run . "" | cat -e 
echo
echo "---------------------------------------------------------"

echo "run . "\n" | cat -e"
go run . "\n" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "Hello\n" | cat -e"
go run . "Hello\n"  | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "hello" | cat -e"
go run . "hello" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "HeLlO" | cat -e"
go run . "HeLlO" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "Hello There" | cat -e"
go run . "Hello There" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "1Hello 2There" | cat -e"
go run . "1Hello 2There" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "{Hello There}" | cat -e"
go run . "{Hello There}" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "Hello\nThere" | cat -e"
go run . "Hello\nThere" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "Hello\n\nThere" | cat -e"
go run . "Hello\n\nThere" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "hello" | cat -e"
go run . "hello" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "HELLO" | cat -e"
go run . "HELLO" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "HeLlo HuMaN" | cat -e"
go run . "HeLlo HuMaN" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "1Hello 2There" | cat -e"
go run . "1Hello 2There" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "Hello\nThere" | cat -e"
go run . "Hello\nThere" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "Hello\n\nThere" | cat -e"
go run . "Hello\n\nThere" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . 'hello There 1 to 2!' | cat -e"
go run . 'hello There 1 to 2!' | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "MaD3IrA\&LiSboN" | cat -e"
go run . "MaD3IrA&LiSboN" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "1a\"#FdwHywR\&/\(\)=" | cat -e"
go run . "1a\"#FdwHywR&/()=" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "{\|}~" | cat -e"
go run . "{|}~" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "[\\]^_ \'a" | cat -e"
go run . "[\]^_ 'a" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "RGB" | cat -e"
go run . "RGB" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . ":\;\<=\>?\@" | cat -e"
go run . ":;<=>?@" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . '\\!\" #$%&\'\"'\"'()*+,-./' | cat -e"
go run . '\!" #$%&'"'"'()*+,-./' | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" | cat -e"
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "abcdefghijklmnopqrstuvwxyz" | cat -e"
go run . "abcdefghijklmnopqrstuvwxyz" | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "banana" standard abc"
go run . "banana" standard abc | cat -e 
echo
echo "---------------------------------------------------------"

echo "run . "hello" standard | cat -e"
go run . "hello" standard | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "hello world" shadow | cat -e"
go run . "hello world" shadow | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "nice 2 meet you" thinkertoy | cat -e"
go run . "nice 2 meet you" thinkertoy | cat -e 
echo
echo "---------------------------------------------------------"
echo "go run . "you \& me" standard | cat -e"
go run . "you & me" standard | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "123" shadow | cat -e"
go run . "123" shadow | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "/\(\"\)" thinkertoy | cat -e"
go run . "/(\")" thinkertoy | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" shadow | cat -e"
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" shadow | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . '\\!\" #$%&\'\"'\"'()*+,-./' thinkertoy | cat -e"
go run . '\!" #$%&'"'"'()*+,-./' thinkertoy | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . "It\'s Working" thinkertoy | cat -e"
go run . "It's Working" thinkertoy | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --output=banner01.txt \"hello\" standard"
go run . --output=banner01.txt "hello" standard | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --output=banner02.txt \"Hello There!\" shadow"
go run . --output=banner02.txt "Hello There!" shadow | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --output=test00.txt \"First\nTest\" shadow | cat -e"
go run . --output=test00.txt "First\nTest" shadow | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --output=test01.txt \"hello\" standard | cat -e"
go run . --output=test01.txt "hello" standard | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --output=test02.txt \"123 -> \#$%\" standard | cat -e"
go run . --output=test02.txt "123 -> #$%" standard | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --output=test03.txt \"432 -> \#$%\&@\" shadow | cat -e"
go run . --output=test03.txt "432 -> #$%&@" shadow | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --output=test04.txt \"There\" shadow | cat -e"
go run . --output=test04.txt "There" shadow | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --output=test05.txt \"123 -> \"#$%@\" thinkertoy | cat -e"
go run . --output=test05.txt "123 -> \"#$%@" thinkertoy | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . -output=test06.txt \"2 you\" thinkertoy | cat -e"
go run . --output=test06.txt "2 you" thinkertoy | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --output=test07.txt 'Testing long output!' standard | cat -e"
go run . --output=test07.txt 'Testing long output!' standard | cat -e 
echo
echo "---------------------------------------------------------"

echo "go run . --color red "banana""
go run . --color red "banana" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=red \"hello world\""
go run . --color=red "hello world" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=green \"1 + 1 = 2\""
go run . --color=green "1 + 1 = 2"  
echo
echo "---------------------------------------------------------"

echo "go run . --color=yellow \"(%&) ??\""
go run . --color=yellow "(%&) ??" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=yellow  ello \"Hello\""
go run . --color=yellow  ello "Hello" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=yellow  e \"Hello\""
go run .  --color=yellow  e \"Hello\" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=yellow  el \"Hello\""
go run . --color=yellow  el \"Hello\" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=orange GuYs \"HeY GuYs\""
go run . --color=orange GuYs "HeY GuYs"  
echo
echo "---------------------------------------------------------"

echo "go run . --color=blue B 'RGB()'"
go run . --color=blue B 'RGB()'  
echo
echo "---------------------------------------------------------"

echo "go run . --color=cyan B \"ReBoOt\""
go run . --color=cyan B "ReBoOt"  
echo
echo "---------------------------------------------------------"

echo "go run . --color=purple ! \"Reboot 01!\""
go run . --color=purple ! "Reboot 01!" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=purple R \"#Reboot 01!!!\""
go run . --color=purple R "#Reboot 01!!!" 
echo
echo "---------------------------------------------------------"

echo "go run . --color=purple R \"#R eb oot &&&&01!!!\""
go run . --color=purple R "#R eb oot &&&&01!!!"  
echo
echo "---------------------------------------------------------"

echo "go run . --color=purple R --color=cyan ! \"#R eb oot &&&&01!!!\""
go run . --color=purple R --color=cyan ! "#R eb oot &&&&01!!!"  
echo
echo "---------------------------------------------------------"