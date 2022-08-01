package main     
import (
    "fmt"
	"github.com/Trace-yz/HelloGo/model"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
    var i int = 10
	i = 30
	i = 40
	fmt.Println("i=",i)
	
	fmt.Println(model.HeroName)

	r := gin.Default()
	r.GET("/ping", ping)
	r.Run()
}