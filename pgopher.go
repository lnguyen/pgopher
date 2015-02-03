package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	_ "github.com/lib/pq"
)

var VERSION = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "pgopher"
	app.Version = VERSION
	app.Author = "Long Nguyen"
	app.Email = "long.nguyen11288@gmail.com"
	app.Usage = "Running PostgresSQL Queries"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "",
			Usage: "hostname",
		},
		cli.StringFlag{
			Name:  "port",
			Value: "",
			Usage: "port",
		},
		cli.StringFlag{
			Name:  "username, u",
			Value: "",
			Usage: "username",
		},
		cli.StringFlag{
			Name:  "password, p",
			Value: "",
			Usage: "password",
		},
		cli.StringFlag{
			Name:  "uri",
			Value: "",
			Usage: "uri",
		},
		cli.StringFlag{
			Name:  "query, q",
			Value: "",
			Usage: "query to be ran",
		},
		cli.StringFlag{
			Name:  "database, d",
			Value: "",
			Usage: "database",
		},
	}
	app.Action = func(c *cli.Context) {
		uri, err := getUri(c)
		if err != nil {
			log.Fatal(err)
		}
		query := c.String("query")
		db, err := sql.Open("postgres", uri+"?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Query: " + query + " completed")
	}

	app.Run(os.Args)
}

func getUri(c *cli.Context) (string, error) {
	uri := ""
	username := c.String("username")
	password := c.String("password")
	host := c.String("host")
	port := c.String("port")
	database := c.String("database")
	if username != "" && password != "" &&
		host != "" && port != "" {
		uri = "postgres://" + username + ":" + password +
			"@" + host + ":" + port + "/" + database
	} else if c.String("uri") != "" {
		uri = c.String("uri")
	} else {
		return "", errors.New("Error creating uri, please enter username, password, host, port, database or uri")
	}
	return uri, nil
}
