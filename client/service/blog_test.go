package service

import (
	"context"
	"testing"

	"github.com/jany/blog/constants"
	models "github.com/jany/blog/model"
	"github.com/jany/blog/protos/blog"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

type mockServiceClient struct {
	blog.BlogServiceClient
	failure bool
}

func (m *mockServiceClient) CreateBlogPost(ctx context.Context, in *blog.Post, opts ...grpc.CallOption) (*blog.Post, error) {
	return &blog.Post{}, nil
}

func (m *mockServiceClient) ReadBlogPost(ctx context.Context, in *blog.Id, opts ...grpc.CallOption) (*blog.Post, error) {
	if m.failure {
		return nil, constants.ErrNotFound
	}
	return &blog.Post{}, nil
}

func (m *mockServiceClient) UpdateBlogPost(ctx context.Context, in *blog.Post, opts ...grpc.CallOption) (*blog.Post, error) {
	if m.failure {
		return nil, constants.ErrNotFound
	}
	return &blog.Post{}, nil
}

func (m *mockServiceClient) DeleteBlogPost(context.Context, *blog.Id, ...grpc.CallOption) (*blog.DeleteResponse, error) {
	if m.failure {
		return &blog.DeleteResponse{Message: "failure"}, constants.ErrNotFound
	}
	return &blog.DeleteResponse{Message: "success"}, nil
}

func TestCreatePost(t *testing.T) {
	type feilds struct {
		ctx  context.Context
		post models.Post
	}

	tests := []struct {
		name string
		args feilds
	}{
		{name: "create success", args: feilds{context.Background(), models.Post{}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := &blogPost{client: &mockServiceClient{}, logger: logrus.New()}
			obj.CreateBlogPost(tt.args.ctx, tt.args.post)
		})
	}

}

func TestReadPost(t *testing.T) {
	type feilds struct {
		ctx    context.Context
		postID *blog.Id
		client *mockServiceClient
	}

	tests := []struct {
		name string
		args feilds
	}{
		{name: "read success", args: feilds{context.Background(), &blog.Id{}, &mockServiceClient{}}},
		{name: "read failure", args: feilds{context.Background(), &blog.Id{}, &mockServiceClient{failure: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := &blogPost{client: tt.args.client, logger: logrus.New()}
			_, err := obj.ReadBlogPost(tt.args.ctx, tt.args.postID.PostId)
			if tt.name == "read failure" {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

		})
	}

}

func TestUpdate(t *testing.T) {
	type feilds struct {
		ctx    context.Context
		postID models.Post
		client *mockServiceClient
	}

	tests := []struct {
		name string
		args feilds
	}{
		{name: "update success", args: feilds{context.Background(), models.Post{}, &mockServiceClient{}}},
		{name: "update failure", args: feilds{context.Background(), models.Post{}, &mockServiceClient{failure: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := &blogPost{client: tt.args.client, logger: logrus.New()}
			_, err := obj.UpdateBlogPost(tt.args.ctx, tt.args.postID)
			if tt.name == "update failure" {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

		})
	}

}
func TestDeletePost(t *testing.T) {
	type feilds struct {
		ctx    context.Context
		postID *blog.Id
		client *mockServiceClient
	}

	tests := []struct {
		name string
		args feilds
	}{
		{name: "delete success", args: feilds{context.Background(), &blog.Id{}, &mockServiceClient{}}},
		{name: "delete failure", args: feilds{context.Background(), &blog.Id{}, &mockServiceClient{failure: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := &blogPost{client: tt.args.client, logger: logrus.New()}
			_, err := obj.DeleteBlogPost(tt.args.ctx, tt.args.postID.PostId)
			if tt.name == "delete failure" {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

		})
	}

}
