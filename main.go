package main

import (
	"gotodo/conf"
	"gotodo/server"
	//"os/exec"
)

//shutdown.exe -s -t 00 -f 强制关机

// 项目根目录的main.go文件，用于启动项目
func main() {
	//cmdExec("calc","")
	conf.Init() // 初始化项目配置。数据库 etc
	router := server.NewRouter()
	router.Run(":3000")
}


/*func cmdExec(command string,param string,){
	c := exec.Command("cmd", "/C", command, param)
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
*/
