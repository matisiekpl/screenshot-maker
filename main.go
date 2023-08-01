package main

import (
	"embed"
	_ "embed"
	"image/png"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

//go:embed font.ttf
var fontBytes []byte

//go:embed devices
var devices embed.FS

func main() {
	app := &cli.App{
		Name:  "screenshot-maker",
		Usage: "render screenshot",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "device",
				Value: "iphone-11",
				Usage: "device used as a mockup",
			},
			&cli.StringFlag{
				Name:     "input",
				Usage:    "screenshot filename",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "text",
				Usage:    "text above mockup",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "text-color",
				Usage: "text above mockup",
				Value: "#ffffff",
			},
			&cli.StringFlag{
				Name:  "background-color",
				Usage: "text above mockup",
				Value: "#2D7FE5",
			},
			&cli.StringFlag{
				Name:  "output",
				Usage: "output filename",
				Value: "screenshot.png",
			},
		},
		Action: func(c *cli.Context) error {
			var screenshot Screenshot
			for _, s := range screenshots {
				if s.Name == c.String("device") {
					screenshot = s
				}
			}
			if screenshot.Name == "" {
				logrus.Panicf("cannot find device: %s", c.String("device"))
			}
			backgroundColor, err := ParseHexColor(c.String("background-color"))
			if err != nil {
				logrus.Panic("cannot parse background color: %s", c.String("background-color"))
			}
			textColor, err := ParseHexColor(c.String("text-color"))
			if err != nil {
				logrus.Panic("cannot parse text color: %s", c.String("text-color"))
			}
			output := screenshot.render(c.String("input"), c.String("text"), backgroundColor, textColor)

			f, err := os.Create(c.String("output"))
			if err != nil {
				logrus.Panicf("cannot write output file: %s", err)
			}
			err = png.Encode(f, output)
			if err != nil {
				logrus.Panicf("cannot encode output file: %s", err)
			}
			logrus.Infof("wrote screenshot to: %s", c.String("output"))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
