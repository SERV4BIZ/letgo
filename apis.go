
	package main

	import(
		"github.com/SERV4BIZ/letgo/global"
		"github.com/SERV4BIZ/letgo/letgo"

		networks "./apis/networks"

	)

	func LetGoAPI() {
		letgo.RegisterAPIHandler = func(rep *global.Request) {
			rep.AddAPI("networks/Ping",networks.Ping)

		}
	}