package blog

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/jany/blog/constants"
	"github.com/jany/blog/protos/blog"
	"github.com/jany/blog/util"
	"github.com/sirupsen/logrus"
)

type server struct {
	BlogPost map[string]*blog.Post
	mu       sync.Mutex
	blog.BlogServiceServer
	logger *logrus.Logger
}

func NewBlogServer(l *logrus.Logger) *server {
	return &server{
		BlogPost: make(map[string]*blog.Post),
		logger:   l,
	}
}

func (s *server) CreateBlogPost(ctx context.Context, post *blog.Post) (*blog.Post, error) {
	funcDesc := "CreateBlogPost"
	s.logger.Info("enter rest " + funcDesc)

	s.mu.Lock()
	defer s.mu.Unlock()

	post.PostId = uuid.New().String()
	s.BlogPost[post.PostId] = post

	return post, nil
}

func (s *server) ReadBlogPost(ctx context.Context, id *blog.Id) (*blog.Post, error) {
	funcDesc := "ReadBlogPost"
	s.logger.Info("enter rest " + funcDesc)

	s.mu.Lock()
	defer s.mu.Unlock()

	var (
		post *blog.Post
		ok   bool
	)

	if post, ok = s.BlogPost[id.PostId]; !ok {
		util.WithFields("server", "blog").Error(constants.ErrNotFound.Error())
		return nil, constants.ErrNotFound
	}
	return post, nil
}

func (s *server) UpdateBlogPost(ctx context.Context, post *blog.Post) (*blog.Post, error) {
	funcDesc := "UpdateBlogPost"
	s.logger.Info("enter rest " + funcDesc)

	s.mu.Lock()
	defer s.mu.Unlock()

	var (
		data *blog.Post
		ok   bool
	)

	if data, ok = s.BlogPost[post.PostId]; !ok {
		util.WithFields("server", "blog").Error(constants.ErrNotFound.Error())
		return nil, constants.ErrNotFound
	}

	post.PublicationData = data.PublicationData
	s.BlogPost[post.PostId] = post

	return post, nil
}

func (s *server) DeleteBlogPost(ctx context.Context, id *blog.Id) (*blog.DeleteResponse, error) {
	funcDesc := "DeleteBlogPost"
	s.logger.Info("enter rest " + funcDesc)

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.BlogPost[id.PostId]; !ok {
		util.WithFields("server", "blog").Error(constants.ErrNotFound.Error())
		return &blog.DeleteResponse{Message: "Failure"}, constants.ErrNotFound
	}
	delete(s.BlogPost, id.PostId)
	return &blog.DeleteResponse{Message: "Success"}, nil
}
