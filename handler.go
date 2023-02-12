package main

import (
	"context"
	chat "dousheng/kitex_gen/chat"
	interact "dousheng/kitex_gen/interact"
	relation "dousheng/kitex_gen/relation"
	user "dousheng/kitex_gen/user"
	video "dousheng/kitex_gen/video"
)

// VideoSvrImpl implements the last service interface defined in the IDL.
type VideoSvrImpl struct{}

// Feed implements the VideoSvrImpl interface.
func (s *VideoSvrImpl) Feed(ctx context.Context, req *video.DouyinFeedRequest) (resp *video.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// Publish implements the VideoSvrImpl interface.
func (s *VideoSvrImpl) Publish(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the VideoSvrImpl interface.
func (s *VideoSvrImpl) PublishList(ctx context.Context, req *video.DouyinPublishListRequest) (resp *video.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// Register implements the UserSvrImpl interface.
func (s *UserSvrImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserSvrImpl interface.
func (s *UserSvrImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// GetById implements the UserSvrImpl interface.
func (s *UserSvrImpl) GetById(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}

// Fav implements the InteractSvrImpl interface.
func (s *InteractSvrImpl) Fav(ctx context.Context, req *interact.DouyinFavoriteActionRequest) (resp *interact.DouyinFavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FavList implements the InteractSvrImpl interface.
func (s *InteractSvrImpl) FavList(ctx context.Context, req *interact.DouyinFavoriteListRequest) (resp *interact.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// Comment implements the InteractSvrImpl interface.
func (s *InteractSvrImpl) Comment(ctx context.Context, req *interact.DouyinCommentActionRequest) (resp *interact.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the InteractSvrImpl interface.
func (s *InteractSvrImpl) CommentList(ctx context.Context, req *interact.DouyinCommentListRequest) (resp *interact.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	return
}

// Sub implements the RelationSvrImpl interface.
func (s *RelationSvrImpl) Sub(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// SubList implements the RelationSvrImpl interface.
func (s *RelationSvrImpl) SubList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// FansList implements the RelationSvrImpl interface.
func (s *RelationSvrImpl) FansList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// FriendList implements the RelationSvrImpl interface.
func (s *RelationSvrImpl) FriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (resp *relation.DouyinRelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetMsg implements the ChatSvrImpl interface.
func (s *ChatSvrImpl) GetMsg(ctx context.Context, req *chat.DouyinRelationActionRequest) (resp *chat.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PostMsg implements the ChatSvrImpl interface.
func (s *ChatSvrImpl) PostMsg(ctx context.Context, req *chat.DouyinMessageChatRequest) (resp *chat.DouyinMessageChatResponse, err error) {
	// TODO: Your code here...
	return
}
