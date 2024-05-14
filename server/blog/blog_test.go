package blog

import (
	"context"
	"testing"

	"github.com/jany/blog/protos/blog"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateBlogPost(t *testing.T) {
	svr := NewBlogServer(logrus.New())

	type feilds struct {
		post *blog.Post
		ctx  context.Context
	}

	tests := []struct {
		name string
		args feilds
	}{
		{name: "success in create", args: feilds{post: &blog.Post{}, ctx: context.Background()}},
	}

	for _, tt := range tests {
		_, err := svr.CreateBlogPost(tt.args.ctx, tt.args.post)
		assert.Nil(t, err)
	}
}

func TestUpdatePost(t *testing.T) {
	svr := NewBlogServer(logrus.New())

	type feilds struct {
		post *blog.Post
		ctx  context.Context
	}

	tests := []struct {
		name string
		args feilds
	}{
		{name: "success in update", args: feilds{post: &blog.Post{}, ctx: context.Background()}},
		{name: "failure in update", args: feilds{post: &blog.Post{}, ctx: context.Background()}},
	}

	for _, tt := range tests {
		if tt.name == "success in update" {
			post, _ := svr.CreateBlogPost(tt.args.ctx, &blog.Post{Title: "random", Author: "jany"})
			tt.args.post = post
		}
		_, err := svr.UpdateBlogPost(tt.args.ctx, tt.args.post)

		if tt.name == "failure in update" {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

func TestReadPost(t *testing.T) {
	svr := NewBlogServer(logrus.New())

	type feilds struct {
		post *blog.Post
		ctx  context.Context
	}

	tests := []struct {
		name string
		args feilds
	}{
		{name: "success in read", args: feilds{post: &blog.Post{}, ctx: context.Background()}},
		{name: "failure in read", args: feilds{post: &blog.Post{}, ctx: context.Background()}},
	}

	for _, tt := range tests {
		if tt.name == "success in read" {
			post, _ := svr.CreateBlogPost(tt.args.ctx, &blog.Post{Title: "random", Author: "jany"})
			tt.args.post = post
		}
		id := tt.args.post.PostId
		_, err := svr.ReadBlogPost(tt.args.ctx, &blog.Id{PostId: id})

		if tt.name == "failure in read" {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

func TestDeletePost(t *testing.T) {
	svr := NewBlogServer(logrus.New())

	type feilds struct {
		post *blog.Post
		ctx  context.Context
	}

	tests := []struct {
		name string
		args feilds
	}{
		{name: "success in delete", args: feilds{post: &blog.Post{}, ctx: context.Background()}},
		{name: "failure in delete", args: feilds{post: &blog.Post{}, ctx: context.Background()}},
	}
	for _, tt := range tests {
		if tt.name == "success in delete" {
			post, _ := svr.CreateBlogPost(tt.args.ctx, &blog.Post{Title: "random", Author: "jany"})
			tt.args.post = post
		}
		id := tt.args.post.PostId
		_, err := svr.DeleteBlogPost(tt.args.ctx, &blog.Id{PostId: id})

		if tt.name == "failure in delete" {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}
