package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicacao de linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{Name: "host", Value: "devbook.com.br"},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) error {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
	return nil
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servidores, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
