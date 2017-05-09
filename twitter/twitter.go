package twitter

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/redigo/redis"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

func twitterApi() *anaconda.TwitterApi {
	viper.AddConfigPath("twitter")
	viper.SetConfigName("config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	consumer_key := viper.GetString("twitter.consumer_key")
	consumer_secret := viper.GetString("twitter.consumer_secret")
	access_token := viper.GetString("twitter.access_token")
	access_token_secret := viper.GetString("twitter.access_token_secret")

	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret)
	api := anaconda.NewTwitterApi(access_token, access_token_secret)
	return api
}

func InsertSearchRecord(resp anaconda.SearchResponse) {
	session, _ := mgo.Dial("mongodb://localhost/gabula_dev")
	defer session.Close()
	db := session.DB("gabula_dev")
	col := db.C("tweet")
	for _, tweet := range resp.Statuses {
		col.Insert(tweet)
	}
}

func GetSearchAll(api *anaconda.TwitterApi, resp anaconda.SearchResponse, count int) int {
	if count < 0 {
		return 0
	}
	count--
	searchResult, err := resp.GetNext(api)
	if err != nil {
		return -1
	}
	query := searchResult.Metadata.Query
	nextResults := searchResult.Metadata.NextResults
	if nextResults != "" {
		RedisStore(query, nextResults)
	}
	InsertSearchRecord(searchResult)
	return GetSearchAll(api, searchResult, count)
}

func GetSearch(query string) {
	api := twitterApi()
	nextResults, err := RedisGet(query)
	values := url.Values{}
	if nextResults != "" {
		nextUrl, _ := url.Parse(nextResults)
		values = nextUrl.Query()
	} else {
		values.Set("result_type", "recent")
		values.Set("count", "100") // default: 15, maximum: 100
	}
	sr, err := api.GetSearch(query, values)
	if err != nil {
		panic(err)
	}
	InsertSearchRecord(sr)
	GetSearchAll(api, sr, 10)
}

func RedisStore(key string, value string) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	c.Do("SET", key, value)
}

func RedisGet(key string) (string, error) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	msg, err := redis.String(c.Do("GET", key))
	if err != nil {
		return "", err
	}
	return msg, err
}
