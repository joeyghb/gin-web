package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"os"
)

type Configuration struct {
	User string
	Pwd  string
}

type HomeSpent struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Date  string        "bson:`date`"
	Item  string        `bson:"item"`
	Cost  int           `bson:"cost"`
	User  string        `bson:"user"`
	Type  string        `bson:"type"`
	Store string        `bson:"store"`
	Memo  string        `bson:"memo"`
}

func GetConfig(c *gin.Context) {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"User": configuration.User,
		"Pwd":  configuration.Pwd,
	})
}

//测试查询
func TestFindHomeSpent(c *gin.Context) {

	//User := "Joey"
	//id := "5c91fc95eb5d0ee41f6e28c7"
	//objectId := bson.ObjectIdHex(id)

	session, err := mgo.Dial("128.199.143.113:27017")
	if err != nil {
		panic(err)
	} else {
		println("Connection Success")
	}
	collections := session.DB("home").C("homespent")
	homespent := HomeSpent{}

	//err = collections.Find(bson.M{"_id": objectId}).One(&homespent)
	err = collections.Find(bson.M{"User": "Joey"}).One(&homespent)

	if err != nil {
		println("find data error")
	}

	strArticleID := bson.ObjectId(homespent.ID).Hex()

	fmt.Printf("ID_Hex : %s\n", homespent.ID)
	fmt.Printf("ID : %s\n", strArticleID)

	fmt.Printf("Date : %s\n", homespent.Date)
	fmt.Printf("Item : %s\n", homespent.Item)
	fmt.Printf("Cost : %s\n", homespent.Cost)
	fmt.Printf("User : %s\n", homespent.User)
	fmt.Printf("Type : %s\n", homespent.Type)
	fmt.Printf("Store : %s\n", homespent.Store)
	fmt.Printf("Memo : %s\n", homespent.Memo)

	c.String(http.StatusOK, fmt.Sprintf("find homespent : %s", homespent.Memo))

	/*
		c.JSON(http.StatusOK, gin.H{
			"User": homespent.User,
			"Cost": homespent.Cost,
		})*/
}

func TestInsertHomeSpent(c *gin.Context) {
	/*
		sn, _ := c.GetQuery("sn")
		name, _ := c.GetQuery("name")
		res := "no insert"
		if sn != "" && name != "" {
			session, err := mgo.Dial("128.199.143.113")
			if err != nil {
				panic(err)
			} else {
				println("连接成功")
			}
			collections := session.DB("home").C("homespent")

			homespent := &HomeSpent{
				User: name,
				Sn:   sn,
			}
			err = collections.Insert(homespent)
			if err != nil {
				res = "no insert"
			}
			res = "insert!"
		}
		c.String(http.StatusOK, res)*/
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello docker !")
	})

	r.GET("/query", func(c *gin.Context) {
		TestFindHomeSpent(c)
	})

	r.GET("/getconfig", func(c *gin.Context) {
		GetConfig(c)
	})

	r.GET("/insert", func(c *gin.Context) {
		TestInsertHomeSpent(c)
	})
}

func main() {
	r := gin.New()
	RegisterRoutes(r)
	r.Run(":18080")
	select {}
}
