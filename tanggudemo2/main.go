package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "tanggudemo2/routers"
)

func main() {

	beego.Run()
}
