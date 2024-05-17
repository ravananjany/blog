package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jany/blog/client/service"
	"github.com/jany/blog/constants"
	models "github.com/jany/blog/model"
	"github.com/jany/blog/util"
	"github.com/sirupsen/logrus"
)

var blogclient service.BlogPost
var logger *logrus.Logger

func createPost(c *gin.Context) {
	funcDesc := "CreatePost"
	logger.Info("enter  " + funcDesc)
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := blogclient.CreateBlogPost(c.Request.Context(), post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func readPost(c *gin.Context) {
	funcDesc := "ReadPost"
	logger.Info("enter  " + funcDesc)
	postId := c.Param("id")

	res, err := blogclient.ReadBlogPost(c.Request.Context(), postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		util.WithFields("client", "api").Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}
func updatePost(c *gin.Context) {
	funcDesc := "UpdatePost"
	logger.Info("enter  " + funcDesc)
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		util.WithFields("client", "api").Error(err)
		return
	}
	res, err := blogclient.UpdateBlogPost(c.Request.Context(), post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		util.WithFields("client", "api").Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func deletePost(c *gin.Context) {
	funcDesc := "DeletePost"
	logger.Info("enter  " + funcDesc)
	postId := c.Param("id")

	res, err := blogclient.DeleteBlogPost(c.Request.Context(), postId)
	if err != nil {
		if err == constants.ErrNotFound {
			c.JSON(http.StatusNotFound, err.Error())
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		util.WithFields("client", "api").Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": res})
}

func getall(c *gin.Context) {
	funcDesc := "getall"
	logger.Info("enter  " + funcDesc)
	res, err := blogclient.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"Blogs": res})
}
