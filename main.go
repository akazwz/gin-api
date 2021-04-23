package main

import (
	"database/sql"
	"fmt"
	"github.com/akazwz/go-gin-demo/global"
	"github.com/akazwz/go-gin-demo/initialize"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @title Golang Restful Api
// @version 1.0
// @description Golang Restful Api Demo
// @termsOfService https://akazwz.com

// @contact.name API Support
// @contact.url https://akazwz.com
// @contact.email akazwz@icloud.com

// @license.name MIT
// @license.url MIT

// @host localhost:8000
// @BasePath /v1
func main() {
	//viper初始化配置
	global.VP = initialize.InitViper()
	if global.VP == nil {
		fmt.Println("配置初始化失败")
	}

	gin.SetMode(global.CFG.Server.Mode)
	//gorm初始化数据库
	global.GDB = initialize.InitDB()
	if global.GDB != nil {
		initialize.CreateTables(global.GDB)
		db, _ := global.GDB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
	} else {
		fmt.Println("数据库连接失败")
		return
	}
	time.Sleep(10 * time.Microsecond)
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
	ReadTimeout := global.CFG.Server.ReadTimeout
	WriteTimeout := global.CFG.Server.WriteTimeout

	s := &http.Server{
		Addr:           addr,
		Handler:        routers,
		ReadTimeout:    ReadTimeout * time.Second,
		WriteTimeout:   WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()

}
