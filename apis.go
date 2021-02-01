
	package main

	import(
		"github.com/SERV4BIZ/letgo/global"
		"github.com/SERV4BIZ/letgo/letgoapp"

		networks "./apis/networks"

	)

	func LetGoAPI() {
		letgoapp.RegisterAPIHandler = func(rep *global.Request) {
			rep.AddAPI("networks/Ping",networks.Ping)

		}
	}