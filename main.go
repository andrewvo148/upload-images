package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"gin-website-starter/models"
	"github.com/astaxie/beego/orm"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)
var ORM orm.Ormer
var token string
func init() {
	models.ConnectToDb()
	ORM = models.GetOrmObject()

}
func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/", indexHandler)
	r.POST("/upload", uploadHandler);
	return r
}

func main() {

	token = os.Getenv("token")
	if token == "" {
		token = generateTokenRandom()
	}
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"auth": token,
	})
}

func uploadHandler(c *gin.Context) {
	tk := c.PostForm("auth")
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	if tk != token || !strings.HasPrefix(file.Header.Get("Content-type"), "image/") || file.Size > 8388608 {
		c.String(http.StatusBadRequest, fmt.Sprint("file upload bad request"))
		return
	}
	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, "tmp/" + filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	img := models.Images{
		FileName: filename,
		Size: file.Size,
		ContentType: file.Header.Get("Content-type"),
	}
	_, err = ORM.Insert(&img)
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprint("Failed to uploaded the image"))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully", file.Filename))
}

func generateTokenRandom() string {
	rb := make([]byte, 32)
	_, err := rand.Read(rb)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(rb)
}





