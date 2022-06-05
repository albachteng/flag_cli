package main

import (
    "fmt"
    "os"
    "gopkg.in/urfave/cli.v1"
)

func main() {
    app := cli.NewApp()
    app.Usage = "Count up or down."
    app.Commands = []cli.Command{
        {
        Name: "up",
        ShortName: "u",
        Usage: "Count up",
        Flags: []cli.Flag{
            cli.IntFlag{
                Name: "stop, s",
                Usage: "Value to count up to", 
                Value: 10,
            },
        },
        Action: func(c *cli.Context) error {
            start := c.Int("stop")
            if start <= 0 {
                fmt.Println("stop cannot be negative.")
            }
            for i := 1; i <= start; i++ {
                fmt.Println(i)
            }
            return nil
        },
        },
        {
        Name: "down",
        ShortName: "d",
        Usage: "Count down",
        Flags: []cli.Flag{
            cli.IntFlag{
                Name: "stop, s",
                Usage: "start counting down from", 
                Value: 10,
            },
        },
        Action: func(c *cli.Context) error {
            start := c.Int("start")
            if start < 0 {
                fmt.Println("start cannot be negative")
            }
            for i := start; i >= 0; i-- {
                fmt.Println(i)
            }
            return nil
        },

    },
}
app.Run(os.Args)
}
