package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"plugin"

	"github.com/asaskevich/EventBus"
	"github.com/joho/godotenv"
	"insanitygaming.net/bans/src/gb/services/addons"
	"insanitygaming.net/bans/src/gb/services/database"
	"insanitygaming.net/bans/src/gb/services/logger"
)

func loadAddons(ctx context.Context, bus EventBus.Bus) {
	dirs, err := os.ReadDir("src/addons")
	if err != nil {
		logger.Logger().Fatal("Error reading addons directory")
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			logger.Logger().Info("Running addon: " + dir.Name())
			addon := fmt.Sprintf("src/addons/%s/addon.so", dir.Name())
			if _, err := os.Stat(addon); err == nil {
				logger.Logger().Info("Found addon: " + addon)
				p, err := plugin.Open(addon)
				if err != nil {
					logger.Logger().Errorf("Error loading addon: " + addon)
					continue
				}

				name, err := p.Lookup("Name")
				if err != nil {
					logger.Logger().Errorf("Error loading addon %s Failed to load metadata", dir.Name())
					continue
				}

				author, err := p.Lookup("Author")
				if err != nil {
					logger.Logger().Errorf("Error loading addon %s Failed to load metadata", name)
					continue
				}

				version, err := p.Lookup("Version")
				if err != nil {
					logger.Logger().Errorf("Error loading addon %s Failed to load metadata", name)
					continue
				}

				description, err := p.Lookup("Description")
				if err != nil {
					logger.Logger().Errorf("Error loading addon %s Failed to load metadata", name)
					continue
				}

				url, err := p.Lookup("Url")
				if err != nil {
					logger.Logger().Errorf("Error loading addon %s Failed to load metadata", name)
					continue
				}

				logger.Logger().Infof("Loading addon %s Version: %s Author: %s Description: %s Url: %s", name, version, author, description, url)

				f, err := p.Lookup("Setup")
				if err != nil {
					logger.Logger().Errorf("Error loading addon %s failed to find setup function", name)
					continue
				}
				if f != nil {
					//TODO: Implement preventing setup from running multiple times
					//		Create table to store addon id and version and compare to see if version > is greater than current version
					// 		If greater, run setup
					setupFunc, ok := f.(func(context.Context))
					if !ok {
						logger.Logger().Errorf("Error loading addon %s has incorrect setup signature looking for func Setup(context.Context)", name)
						continue
					}
					setupFunc(ctx)
				}

				f, err = p.Lookup("Run")
				if err != nil {
					logger.Logger().Errorf("Error loading addon %s failed to find run function", name)
					continue
				}

				runFunc, ok := f.(func(EventBus.Bus))
				if !ok {
					logger.Logger().Errorf("Error loading addon %s has incorrect run signature looking for func Run(EventBus.Bus)", name)
					continue
				}

				runFunc(bus)

				addons.Register(*name.(*string), &addons.Addon{
					Plugin:      p,
					Name:        *name.(*string),
					Version:     *version.(*string),
					Author:      *author.(*string),
					Description: *description.(*string),
					Url:         *url.(*string),
				})
			}
		}
	}
}

func main() {
	setup := flag.Bool("setup", false, "Install the database")
	flag.Parse()
	var log = logger.Logger()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.Connect()
	defer database.Close()

	if *setup {
		database.RunSetup()
		return
	}

	background := context.Background()

	bus := EventBus.New()

	loadAddons(background, bus)

	rows, err := database.Query(background, "SELECT * FROM gb_bans")
	fmt.Print(rows, err)
}
