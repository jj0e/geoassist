package main

import (
	"os"
	"fmt"
	"bufio"
	"github.com/jj0e/geoguessr-assist/utils"
	"github.com/jj0e/geoguessr-assist/guess"
)

func main() {
	fmt.Print(utils.GetTimestamp(), "Welcome to GeoAssist\n")
	fmt.Print(utils.GetTimestamp(), "Enter uuid: ")

	reader := bufio.NewReader(os.Stdin)
	uuid, _ := reader.ReadString('\n')

	guess.Run(uuid)
}
