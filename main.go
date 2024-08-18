package main

import (
	"fmt"
	"strings"
)



func main()  {
	str := make([]string,0)
	str = append(str,"camp_migration_status:155949-33087")
	for _,v := range str {
		appEnvInfo := strings.Split(v,":")
		appEnv := strings.Split(appEnvInfo[1],"-")
		appId,envId := appEnv[0],appEnv[1]
		fmt.Println(appId,envId)
	}
}