package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ApiConf struct {
	ConsumerKey       string `yaml:"consumer_key"`
	ConsumerSecret    string `yaml:"consumer_secret"`
	AccessToken       string `yaml:"access_token"`
	AccessTokenSecret string `yaml:"access_token_secret"`
}

func main() {
	buf, _ := ioutil.ReadFile("config.yaml")
	var apiConf ApiConf
	if err := yaml.Unmarshal(buf, &apiConf); err != nil {
		panic(err)
	}

	anaconda.SetConsumerKey(apiConf.ConsumerKey)
	anaconda.SetConsumerSecret(apiConf.ConsumerSecret)
	api := anaconda.NewTwitterApi(apiConf.AccessToken, apiConf.AccessTokenSecret)

	searchResult, _ := api.GetSearch("kabu", nil)
	for _, tweet := range searchResult.Statuses {
		fmt.Println("TweetId: ", tweet.Id, "Tweet:", tweet.Text)
	}

	fmt.Println("=----------------")

	searchResultNext, _ := searchResult.GetNext(api)
	for _, tweet := range searchResultNext.Statuses {
		fmt.Println("TweetId: ", tweet.Id, "Tweet:", tweet.Text)
	}
}
