package main

import (
	"fmt"
	"time"
)

var jokes []string = []string{
	"If you spell Chuck Norris in Scrabble, you win. Forever.",
	"Chuck Norris breathes air ... five times a day.",
	"Chuck Norris does not sleep. He waits.",
	"The chief export of Chuck Norris is pain.",
	"Chuck Norris can dribble a bowling ball.",
	"If you want a list of Chuck Norris' enemies, just check the extinct species list.",
	"Chuck Norris once shot an enemy plane down with his finger, by yelling,'Bang!'.",
	"Chuck Norris has a grizzly bear rug in his bedroom. It's not dead. It's just too scared to move.",
}

func main() {
	index := time.Now().Unix() % int64(len(jokes))
	fmt.Println(jokes[index])
}
