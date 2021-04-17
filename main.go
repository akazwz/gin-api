package main

import (
	"fmt"
	"github.com/akaedison/go-gin-demo/global"
	"github.com/akaedison/go-gin-demo/initialize"
)

// @title Golang Restful Api
// @version 1.0
// @description Golang Restful Api Demo
// @termsOfService https://akazwz.com

// @contact.name API Support
// @contact.url https://akazwz.com
// @contact.email akaedison@icloud.com

// @license.name MIT
// @license.url MIT

// @host 127.0.0.1:8000
// @BasePath /v1
func main() {
	//viper初始化配置
	global.VP = initialize.InitViper()
	if global.VP == nil {
		fmt.Println("配置初始化失败")
	}

	//gorm初始化数据库
	global.GDB = initialize.InitDB()
	if global.GDB != nil {
		initialize.CreateTables(global.GDB)
		db, _ := global.GDB.DB()
		defer db.Close()
	} else {
		fmt.Println("数据库连接失败")
		return
	}

	str := `
█████╗ ██╗  ██╗ █████╗ ███████╗██╗    ██╗███████╗                                      
██╔══██╗██║ ██╔╝██╔══██╗╚══███╔╝██║    ██║╚══███╔╝                                      
███████║█████╔╝ ███████║  ███╔╝ ██║ █╗ ██║  ███╔╝                                       
██╔══██║██╔═██╗ ██╔══██║ ███╔╝  ██║███╗██║ ███╔╝                                        
██║  ██║██║  ██╗██║  ██║███████╗╚███╔███╔╝███████╗                                      
╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝ ╚══╝╚══╝ ╚══════╝                                      
                                                                                        
 ██████╗  ██████╗        ██████╗ ██╗███╗   ██╗      ██████╗ ███████╗███╗   ███╗ ██████╗ 
██╔════╝ ██╔═══██╗      ██╔════╝ ██║████╗  ██║      ██╔══██╗██╔════╝████╗ ████║██╔═══██╗
██║  ███╗██║   ██║█████╗██║  ███╗██║██╔██╗ ██║█████╗██║  ██║█████╗  ██╔████╔██║██║   ██║
██║   ██║██║   ██║╚════╝██║   ██║██║██║╚██╗██║╚════╝██║  ██║██╔══╝  ██║╚██╔╝██║██║   ██║
╚██████╔╝╚██████╔╝      ╚██████╔╝██║██║ ╚████║      ██████╔╝███████╗██║ ╚═╝ ██║╚██████╔╝
 ╚═════╝  ╚═════╝        ╚═════╝ ╚═╝╚═╝  ╚═══╝      ╚═════╝ ╚══════╝╚═╝     ╚═╝ ╚═════╝ 
                                                                                        `
	fmt.Println(str)

	routers := initialize.Routers()
	addr := fmt.Sprintf(":%d", global.CFG.Server.Addr)

	_ = routers.Run(addr)

}
