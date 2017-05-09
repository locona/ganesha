package mastodon

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	mastodon "github.com/mattn/go-mastodon"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	ClientID     string `yaml:client_id`
	ClientSecret string `yaml:client_secret`
	Email        string `yaml:email`
	Password     string `yaml:password`
}

func Mastodon() {
	buf, _ := ioutil.ReadFile("mastodon/config.yaml")
	client_id, client_secret := GetConfidential()
	var config Config
	if err := yaml.Unmarshal(buf, &config); err != nil {
		panic(err)
	}
	client := mastodon.NewClient(&mastodon.Config{
		Server:       "https://mstdn.jp",
		ClientID:     client_id,
		ClientSecret: client_secret,
	})

	fmt.Println("email: ", config.Email)
	fmt.Println("password: ", config.Password)

	err := client.Authenticate(context.Background(), config.Email, config.Password)
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
