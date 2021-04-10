package guess

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jj0e/geoguessr-assist/files"
	"github.com/jj0e/geoguessr-assist/utils"
)

func Run(uuid string) {
	fm := files.New()
	conf := fm.LoadConfig()
	if conf.Email == "" || conf.Password == "" {
		fmt.Print(utils.GetTimestamp(), "Enter login email: ")
		reader := bufio.NewReader(os.Stdin)
		email, _ := reader.ReadString('\n')
		email = strings.TrimSuffix(email, "\r\n")
		fmt.Print(utils.GetTimestamp(), "Enter login password: ")
		reader = bufio.NewReader(os.Stdin)
		password, _ := reader.ReadString('\n')
		password = strings.TrimSuffix(password, "\r\n")
		conf.Email = email
		conf.Password = password
		fm.UpdateConfig(&conf)
	}

	instance := New(conf.Email, conf.Password)
	instance.Join(uuid)
	instance.Watch()
}
