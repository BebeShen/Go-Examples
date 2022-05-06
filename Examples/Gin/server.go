package main
 
import (
   "fmt"
   "net/http"

   "github.com/gin-gonic/gin"
)
 
func main() {
   gin.SetMode(gin.DebugMode)
   apiServer := gin.Default()
   apiServer.GET("/ping", Ping)
   apiServer.GET("/ready", Ready)
   err := apiServer.Run(":8787" )
   if err != nil {
       fmt.Println("health check:", err)
   }
}
 
//Ping health check
func Ready(c *gin.Context) {
   c.JSON(http.StatusOK, gin.H{
       "status": "UP",
   })
}
func Ping(c *gin.Context) {
   c.JSON(http.StatusOK, gin.H{
       "message": "PONG test",
   })
}