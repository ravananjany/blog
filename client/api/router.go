package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jany/blog/client/service"
	"github.com/jany/blog/config"
	"github.com/jany/blog/protos/blog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartRest(port string) {
	router := gin.Default()
	router.POST("/blog", createPost)
	router.DELETE("/blog/:id", deletePost)
	router.PUT("/blog", updatePost)
	router.GET("/blog/:id", readPost)
	router.Run(fmt.Sprintf("localhost:%s", port))
}
func StartClient() {
	log, conf := config.LoadConfig()
	logger = log
	log.Info("connecting client")
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", conf.Grpcport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connecting to server", err)
	}
	defer conn.Close()
	blogclient = service.NewBlog(blog.NewBlogServiceClient(conn), log)
	log.Info("starting rest api")
	StartRest(conf.Apiport)
}

func Start() {
	StartClient()
}
