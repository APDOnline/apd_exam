package main

import (
	//log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/urfave/cli"

	"os"

	"github.com/Sirupsen/logrus"
	"github.com/apd_exam/data"
	"github.com/apd_exam/route"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

func main() {

	app := cli.NewApp()
	app.Name = "childrenlab"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			EnvVar: "DEBUG",
			Name:   "debug",
			Usage:  "Debug",
			Value:  "false",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_USER",
			Name:   "database_user",
			Usage:  "Database user name",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_PASSWORD",
			Name:   "database_password",
			Usage:  "Database password",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_IP",
			Name:   "database_IP",
			Usage:  "Database IP address with port number",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_NAME",
			Name:   "database_name",
			Usage:  "Database name",
			Value:  "swing_test_record",
		},
	}

	app.Action = func(c *cli.Context) error {
		data.DatabaseInfo = data.DatabaseStruct{
			Name:     c.String("database_name"),
			User:     c.String("database_user"),
			Password: c.String("database_password"),
			IP:       c.String("database_IP"),
		}

		if c.Bool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}

		db := data.NewDB()
		defer db.Close()

		if err := data.CreateSchema(db); err != nil {
			errors.Cause(err)
			log.Fatal(err)
		}

		if err := data.DefaultData(db); err != nil {
			errors.Cause(err)
		}

		e := echo.New()
		e.Logger.SetLevel(log.DEBUG)
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:80"},
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		}))

		route.InitRoute(e)

		e.Logger.Fatal(e.Start(":8112"))

		/*		if c.Bool("debug") {
					return r.Run(":8111")
				} else {
					return r.RunTLS(":8111", ".ssh/childrenlab.chained.crt", ".ssh/childrenlab.com.key")
				}*/

		return nil

	}

	app.Run(os.Args)

}
