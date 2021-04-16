package main

import (
	"fmt"
	"github.com/akaedison/go-gin-demo/global"
	"github.com/akaedison/go-gin-demo/initialize"
)

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
