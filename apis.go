
	package main

	import(
		"github.com/SERV4BIZ/letsgo/global"
		"github.com/SERV4BIZ/letsgo/letsgoapp"

		networks "github.com/SERV4BIZ/letsgo/apis/networks"

	)

	func LetsGoAPI() {
		letsgoapp.RegisterAPIHandler = func(rep *global.Request) {
			rep.AddAPI("networks/Ping",networks.Ping)

		}
	}