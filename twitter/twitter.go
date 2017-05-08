package twitter

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gopkg.in/mgo.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/url"
)

type Config struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// func main() {
// api := twitterApi(c)
// fmt.Println("api", api)
//
// session, _ := mgo.Dial("mongodb://localhost/gabula_dev")
// defer session.Close()
// db := session.DB("gabula_dev")
//
// // db, err := gorm.Open("mysql", "root@gabula_dev?charset=utf8&parseTime=True&loc=Local")
// col := db.C("tweet")
// values := url.Values{}
// values.Set("result_type", "recent")
// values.Set("count", "100") // default: 15, maximum: 100
//
// searchResult, _ := api.GetSearch("from:...", values)
// count := 0
// for _, tweet := range searchResult.Statuses {
// col.Insert(tweet)
// count += 1
// }
// fmt.Println("count", count)
//
// searchResultNext, _ := searchResult.GetNext(api)
// for _, tweet := range searchResultNext.Statuses {
// fmt.Println(tweet.Text)
// }
// }

func twitterApi() *anaconda.TwitterApi {
	buf, _ := ioutil.ReadFile("config/twitter.yaml")
	var c Config
	if err := yaml.Unmarshal(buf, &c); err != nil {
		panic(err)
	}

	anaconda.SetConsumerKey(c.ConsumerKey)
	anaconda.SetConsumerSecret(c.ConsumerSecret)
	api := anaconda.NewTwitterApi(c.AccessToken, c.AccessTokenSecret)
	return api
}

func GetSearch(query string) {
	api := twitterApi()
	values := url.Values{}
	values.Set("result_type", "recent")
	values.Set("count", "100") // default: 15, maximum: 100
	searchResult, _ := api.GetSearch(query, values)
	count := 0
	for _, tweet := range searchResult.Statuses {
		// col.Insert(tweet)
		fmt.Println(tweet)

		count += 1
	}
}
