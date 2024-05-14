package service

import (
	"context"
	"fmt"
	"time"

	"github.com/jany/blog/constants"
	models "github.com/jany/blog/model"
	"github.com/jany/blog/protos/blog"
	"github.com/jany/blog/util"
	"github.com/sirupsen/logrus"
)

type BlogPost interface {
	CreateBlogPost(ctx context.Context, input models.Post) (models.Post, error)
	ReadBlogPost(ctx context.Context, id string) (models.Post, error)
	UpdateBlogPost(ctx context.Context, input models.Post) (models.Post, error)
	DeleteBlogPost(ctx context.Context, id string) (string, error)
}

type blogPost struct {
	client blog.BlogServiceClient
	logger *logrus.Logger
}

func NewBlog(c blog.BlogServiceClient, l *logrus.Logger) BlogPost {
	return &blogPost{
		client: c,
		logger: l,
	}
}

func (b *blogPost) CreateBlogPost(ctx context.Context, input models.Post) (models.Post, error) {
	funcDesc := "CreateBlogPost"
	b.logger.Info("enter  " + funcDesc)

	res, err := b.client.CreateBlogPost(ctx, &blog.Post{
		Title:           input.Title,
		Content:         input.Content,
		Author:          input.Author,
		PublicationData: time.Now().String(),
		Tags:            input.Tags,
	})

	if err != nil {
		util.WithFields("client", "blog").Error(err)
		return models.Post{}, err
	}

	input.ID = res.PostId
	b.logger.Info(fmt.Sprintf("response from the server: %v\n", res))
	b.logger.Info("exit  " + funcDesc)
	return util.ConvertToResource(res), nil
}

func (b *blogPost) ReadBlogPost(ctx context.Context, id string) (models.Post, error) {
	funcDesc := "ReadBlogPost"
	b.logger.Info("enter  " + funcDesc)

	res, err := b.client.ReadBlogPost(ctx, &blog.Id{PostId: id})

	if err != nil {
		util.WithFields("client", "blog").Error(err)
		return models.Post{}, err
	}
	b.logger.Info(fmt.Sprintf("response from the server: %v\n", res))
	b.logger.Info("exit  " + funcDesc)
	return models.Post{
		ID:              res.PostId,
		Title:           res.Title,
		Content:         res.Content,
		Author:          res.Author,
		PublicationData: res.PublicationData,
		Tags:            res.Tags,
	}, nil
}

func (b *blogPost) UpdateBlogPost(ctx context.Context, input models.Post) (models.Post, error) {
	funcDesc := "UpdateBlogPost"
	b.logger.Info("enter  " + funcDesc)

	if input.PublicationData != "" {
		util.WithFields("client", "blog").Error(constants.ErrPubNotPermitted)
		return models.Post{}, constants.ErrPubNotPermitted
	}

	res, err := b.client.UpdateBlogPost(ctx, &blog.Post{
		PostId:  input.ID,
		Title:   input.Title,
		Content: input.Content,
		Author:  input.Author,
		Tags:    input.Tags,
	})

	if err != nil {
		util.WithFields("client", "blog").Error(err)
		return models.Post{}, err
	}

	b.logger.Info(fmt.Sprintf("response from the server: %v\n", res))
	b.logger.Info("exit  " + funcDesc)
	return models.Post{
		ID:              res.PostId,
		Title:           res.Title,
		Content:         res.Content,
		Author:          res.Author,
		PublicationData: res.PublicationData,
		Tags:            res.Tags,
	}, nil
}

func (b *blogPost) DeleteBlogPost(ctx context.Context, id string) (string, error) {
	funcDesc := "DeleteBlogPost"
	b.logger.Info("enter  " + funcDesc)

	res, err := b.client.DeleteBlogPost(ctx, &blog.Id{PostId: id})

	if err != nil {
		util.WithFields("client", "blog").Error(err)
		return res.Message, err
	}
	b.logger.Info(fmt.Sprintf("response from the server: %v\n", res))
	return res.Message, nil
}
