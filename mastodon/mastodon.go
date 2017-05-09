package mastodon

import (
	"context"
	"fmt"
	"log"

	mastodon "github.com/mattn/go-mastodon"
	"github.com/spf13/viper"
)

type Config struct {
	ClientID     string `yaml:client_id`
	ClientSecret string `yaml:client_secret`
	Email        string `yaml:email`
	Password     string `yaml:password`
}

func Mastodon() {
	viper.AddConfigPath("mastodon")
	viper.SetConfigName("config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	email := viper.GetString("mastodon.email")
	password := viper.GetString("mastodon.password")

	client_id, client_secret := GetConfidential()
	c := mastodon.NewClient(&mastodon.Config{
		Server:       "https://mstdn.jp",
		ClientID:     client_id,
		ClientSecret: client_secret,
	})

	err := client.Authenticate(context.Background(), email, password)
	if err != nil {
		log.Fatal(err)
	}
	timeline, err := client.GetTimelineHome(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	for i := len(timeline) - 1; i >= 0; i-- {
		fmt.Println(timeline[i])
	}

}

func GetConfidential() (string, string) {
	app, _ := mastodon.RegisterApp(context.Background(), &mastodon.AppConfig{
		Server:     "https://mstdn.jp",
		ClientName: "locona",
		Scopes:     "read write follow",
		Website:    "https://github.com/mattn/go-mastodon",
	})
	return app.ClientID, app.ClientSecret
}
