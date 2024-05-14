package util

import (
	"strconv"

	models "github.com/jany/blog/model"
	"github.com/jany/blog/protos/blog"
	"github.com/sirupsen/logrus"
)

func WithFields(pac string, file string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"package": pac,
		"go":      file,
	})
}

func StringTOI(s string) (int, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func ConvertToResource(p *blog.Post) models.Post {
	post := models.Post{
		ID:              p.PostId,
		Title:           p.Title,
		Author:          p.Author,
		Content:         p.Content,
		PublicationData: p.PublicationData,
		Tags:            p.Tags,
	}
	return post
}
