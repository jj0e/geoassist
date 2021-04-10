package guess

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jj0e/geoguessr-assist/files"
	"github.com/jj0e/geoguessr-assist/utils"
)

func Run(uuid string) {
	fmt.Print("\033[H\033[2J")
	fm := files.New()
	conf := fm.LoadConfig()
	if conf.Email == "" || conf.Password == "" {
		fmt.Print(utils.GetTimestamp(), "Enter login email: ")
		reader := bufio.NewReader(os.Stdin)
		email, _ := reader.ReadString('\n')
		fmt.Print(utils.GetTimestamp(), "Enter login password: ")
		reader = bufio.NewReader(os.Stdin)
		password, _ := reader.ReadString('\n')
		conf.Email = email
		conf.Password = password
		fm.UpdateConfig(&conf)
		fmt.Print("\033[H\033[2J")
	}

	instance := New(conf.Email, conf.Password)
	instance.Join(uuid)
	instance.Watch()
}
