package main

//
//import (
//	"fmt"
//	"net/http"
//	"os"
//
//	"github.com/arl/statsviz"
//	"github.com/urfave/cli"
//)
//
//// @title gin-guestbook Example API
//// @version 1.0
//// @description This is a sample server guestbook server.
//
//// @BasePath /api/v1
//func main() {
//	app := cli.NewApp()
//	app.Name = "gin-guestbook"
//	app.Usage = "gin-guestbook -c config/config.local.json"
//	printVersion := false
//	app.Flags = []cli.Flag{
//		cli.StringFlag{
//			Name:  "conf, c",
//			Value: "config/config.local.json",
//			Usage: "config/config.{local|dev|test|pre|prod}.json",
//		},
//		cli.BoolFlag{
//			Name:        "version, v",
//			Required:    false,
//			Usage:       "-v",
//			Destination: &printVersion,
//		},
//	}
//
//	app.Action = func(c *cli.Context) error {
//		if printVersion {
//			fmt.Printf("{%#v}", version.Get())
//			return nil
//		}
//
//		conf := c.String("conf")
//		config.Init(conf)
//		err := sentinelm.InitSentinelByCustom()
//		if err != nil {
//			return err
//		}
//		err = models.Database(config.Conf.MySQL)
//		if err != nil {
//			return err
//		}
//		go func() {
//			err = rpc.InitRPCService()
//			if err != nil {
//				panic(err)
//			}
//		}()
//		server := router.InitRouter()
//		go runMointer()
//		server.GinEngine.Run(":8000")
//		return nil
//	}
//	app.Run(os.Args)
//}
//
//func runMointer() {
//	// Register statsviz handlers on the default serve mux.
//	statsviz.RegisterDefault()
//	http.ListenAndServe(":8001", nil)
//}
