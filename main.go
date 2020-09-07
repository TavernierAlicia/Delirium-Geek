package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//define variables
var (
	log        *zap.Logger
	identified bool
)

func main() {

	//zap stuff
	log, _ = zap.NewProduction()
	defer log.Sync()

	//Define router
	router := gin.Default()

	//to include html
	router.LoadHTMLFiles("templates/index.html", "templates/indexVerified.html", "templates/tchat.html")

	//to include js
	router.Static("/js", "./js")

	//to include css
	router.Static("/css", "./css")

	//to include images
	router.Static("/media", "./media")

	//GET requests
	router.NoRoute(errorPage)

	// Routes
	router.GET("/", indexPage)
	router.GET("/index", indexPage)
	router.GET("/indexVerified", indexVerified)

	// define the route to ws connection
	router.GET("/tchat", func(c *gin.Context) {
		log.Info("Socket connection request")
		// tchatFunc(c.Writer, c.Request, c)
		c.JSON(200, gin.H{
			"event": "waitingidentification",
		})
	})

	router.Run("127.0.0.1:3000")
}

func indexPage(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func indexVerified(c *gin.Context) {
	c.HTML(200, "indexVerified.html", nil)
}

//handle custom error 404
func errorPage(c *gin.Context) {
	var cause string
	//get infos about the path
	host := "127.0.0.1"
	fullpath := host + c.Request.URL.Path

	cause = fullpath + " cette page n'est pas accessible pour le moment ou n'existe pas"

	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": cause})
}

func receptForm(c *gin.Context) {
	c.Request.ParseForm()

	pseudo := strings.Join(c.Request.PostForm["pseudo"], " ")
	picture := strings.Join(c.Request.PostForm["picture"], " ")

	fmt.Println("PSEUDO: ", pseudo, "PICTURE PATH: ", picture)

	if pseudo != "" {
		c.Redirect(http.StatusMovedPermanently, "/indexVerified")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/index")
	}

	// insert db funct call
}

// func tchatFunc(w http.ResponseWriter, r *http.Request, c *gin.Context) {
// 	// set identification to false
// 	identified = false

// 	var wsupgrader = websocket.Upgrader{
// 		ReadBufferSize:  1024,
// 		WriteBufferSize: 1024,
// 	}

// 	wsupgrader.CheckOrigin = func(r *http.Request) bool { return true }
// 	conn, err := wsupgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		fmt.Println("Failed to set websocket upgrade: ", err)
// 		return
// 	}

// 	// Loop on listening
// 	for {
// 		for identified == false {

// 		}
// 	}
// }
