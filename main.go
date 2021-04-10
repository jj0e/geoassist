package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jj0e/geoguessr-assist/guess"
	"github.com/jj0e/geoguessr-assist/utils"
)

func main() {
	fmt.Print(utils.GetTimestamp(), "Welcome to GeoAssist\n")
	fmt.Print(utils.GetTimestamp(), "Enter uuid: ")

	reader := bufio.NewReader(os.Stdin)
	uuid, _ := reader.ReadString('\n')
	uuid = strings.TrimSuffix(uuid, "\r\n")

	guess.Run(uuid)
}
