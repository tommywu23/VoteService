package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/tommywu23/VoteService/models"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DB map[string][]string `yaml:"db,omitempty"`
}

type App struct {
	db *sql.DB
}

type Req struct {
	Event  string `json:"event"`
	Params string `json:"params"`
}

var (
	Settings Config
)

func loadConfig() {
	filename, _ := filepath.Abs("./config.yaml")

	yamlFile, _ := ioutil.ReadFile(filename)

	_ = yaml.Unmarshal(yamlFile, &Settings)

}

func main() {
	loadConfig()
	fmt.Println(Settings.DB["url"][0])

	db, err := sql.Open("postgres", Settings.DB["url"][0])

	if err != nil {
		fmt.Println(err)
	}

	app := &App{db: db}

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/vote", app.voteCreate)
		v1.GET("/vote", app.voteGet)
	}

	r.Run(":6600")
}

func (app *App) voteCreate(c *gin.Context) {
	var v models.VoteBase
	c.Bind(&v)
	d, _ := json.Marshal(v)
	var req = Req{}
	req.Event = "vote"
	req.Params = string(d)
	p, _ := json.Marshal(req)
	http.Post("http://192.168.22.67:10678/notify", "application/json", bytes.NewBuffer(p))

	//result, err := app.db.Exec("insert into vote(data) value($1)", v)

	c.JSON(200, v)
}

func (app *App) voteGet(c *gin.Context) {
	c.JSON(200, "")
}
