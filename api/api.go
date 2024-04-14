package api

// subscribeToUser godoc
//
//	@Summary	Subscribe to user
//	@Tags		Users
//	@Produce	json
//	@Param		id	path	int	true	"ID of user to subscribe to"	minimum(1)
//	@Success	200
//	@Failure	404	{object}	RestError
//	@Failure	500	{object}	RestError
//	@Router		/v1/users/{id}/subscription [post]
func (s Server) subscribeToUser() {}

// unsubscribeFromUser godoc
//
//	@Summary	Unsubscribe from user
//	@Tags		Users
//	@Produce	json
//	@Param		id	path	int	true	"ID of user to unsubscribe from"	minimum(1)
//	@Success	200
//	@Failure	404	{object}	RestError
//	@Failure	500	{object}	RestError
//	@Router		/v1/users/{id}/subscription [delete]
func (s Server) unsubscribeFromUser() {}

// uploadImage godoc
//
//	@Summary	Upload image
//	@Tags		Images
//	@Accept		jpeg
//	@Produce	json
//	@Param		request	body		string	true	"JPEG image"
//	@Success	201		{object}	ImageInfo
//	@Failure	422		{object}	RestError
//	@Failure	500		{object}	RestError
//	@Router		/v1/images [post]
func (s Server) uploadImage() {}

// createPost godoc
//
//	@Summary	Create post
//	@Tags		Posts
//	@Accept		json
//	@Produce	json
//	@Param		body	body		Post	true	"Post object"
//	@Success	201		{object}	Post
//	@Failure	422		{object}	RestError
//	@Failure	500		{object}	RestError
//	@Router		/v1/posts [post]
func (s Server) createPost() {}

// likePost godoc
//
//	@Summary	Like post
//	@Tags		Posts
//	@Produce	json
//	@Param		id	path		int	true	"ID of post to like"	minimum(1)
//	@Success	200	{object}	Post
//	@Failure	404	{object}	RestError
//	@Failure	500	{object}	RestError
//	@Router		/v1/posts/{id}/like [post]
func (s Server) likePost() {}

// commentPost godoc
//
//	@Summary	Comment post
//	@Tags		Posts
//	@Accept		json
//	@Produce	json
//	@Param		id		path	int		true	"ID of post to comment"	minimum(1)
//	@Param		body	body	Comment	true	"Comment object"
//	@Success	200
//	@Failure	404	{object}	RestError
//	@Failure	422	{object}	RestError
//	@Failure	500	{object}	RestError
//	@Router		/v1/posts/{id}/comment [post]
func (s Server) commentPost() {}

// getPostsByUser godoc
//
//	@Summary	Get posts by user
//	@Tags		Posts
//	@Produce	json
//	@Param		id		path		int	true	"ID of user to get posts"	minimum(1)
//	@Param		limit	query		int	false	"limit"						minimum(1)
//	@Param		offset	query		int	false	"offset"					minimum(0)
//	@Success	200		{object}	PaginatedPosts
//	@Failure	404		{object}	RestError
//	@Failure	500		{object}	RestError
//	@Router		/v1/posts/users/{id} [get]
func (s Server) getPostsByUser() {}

// postsTop godoc
//
//	@Summary	Get posts top
//	@Tags		Posts
//	@Accept		json
//	@Produce	json
//	@Param		limit	query		int			false	"limit"		minimum(1)
//	@Param		offset	query		int			false	"offset"	minimum(0)
//	@Param		body	body		PostsTop	true	"PostsTop object"
//	@Success	200		{object}	PaginatedPosts
//	@Failure	422		{object}	RestError
//	@Failure	500		{object}	RestError
//	@Router		/v1/posts/top [get]
func (s Server) postsTop() {}

// getFeedByUser godoc
//
//	@Summary	Get feed by user
//	@Tags		Feed
//	@Produce	json
//	@Param		id		path		int	true	"ID of user to get feed"	minimum(1)
//	@Param		limit	query		int	false	"limit"						minimum(1)
//	@Param		offset	query		int	false	"offset"					minimum(0)
//	@Success	200		{object}	PaginatedPosts
//	@Failure	404		{object}	RestError
//	@Failure	500		{object}	RestError
//	@Router		/v1/feed/users/{id} [get]
func (s Server) getFeedByUser() {}

// createChat godoc
//
//	@Summary	Create chat
//	@Tags		Chats
//	@Accept		json
//	@Produce	json
//	@Param		body	body		Chat	true	"Chat object"
//	@Success	201		{object}	Chat
//	@Failure	422		{object}	RestError
//	@Failure	500		{object}	RestError
//	@Router		/v1/chats [post]
func (s Server) createChat() {}

// updateChat godoc
//
//	@Summary	Update chat by ID
//	@Tags		Chats
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int		true	"ID of chat to be updated"	minimum(1)
//	@Param		body	body		Chat	true	"Chat object"
//	@Success	200		{object}	Chat
//	@Failure	404		{object}	RestError
//	@Failure	422		{object}	RestError
//	@Failure	500		{object}	RestError
//	@Router		/v1/chats/{id} [patch]
func (s Server) updateChat() {}

// deleteChat godoc
//
//	@Summary	Delete chat by ID
//	@Tags		Chats
//	@Produce	json
//	@Param		id	path	int	true	"ID of chat to be deleted"	minimum(1)
//	@Success	204
//	@Failure	404	{object}	RestError
//	@Failure	500	{object}	RestError
//	@Router		/v1/chats/{id} [delete]
func (s Server) deleteChat() {}

// getChat godoc
//
//	@Summary	Get chat messages
//	@Tags		Chats
//	@Produce	json
//	@Param		id		path		int	true	"ID of chat to be retrieved"	minimum(1)
//	@Param		limit	query		int	false	"limit"							minimum(1)
//	@Param		offset	query		int	false	"offset"						minimum(0)
//	@Success	201		{object}	PaginatedMessages
//	@Failure	404		{object}	RestError
//	@Failure	500		{object}	RestError
//	@Router		/v1/chats/{id} [get]
func (s Server) getChatMessages() {}

// getChatsByUser godoc
//
//	@Summary	Get chats by user
//	@Tags		Chats
//	@Produce	json
//	@Param		id		path		int	true	"ID of user to get chats"	minimum(1)
//	@Param		limit	query		int	false	"limit"						minimum(1)
//	@Param		offset	query		int	false	"offset"					minimum(0)
//	@Success	200		{object}	PaginatedChats
//	@Failure	404		{object}	RestError
//	@Failure	500		{object}	RestError
//	@Router		/v1/chats/users/{id} [get]
func (s Server) getChatsByUser() {}

// sendMessage godoc
//
//	@Summary	Send message to chat
//	@Tags		Chats
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int		true	"ID of chat to send message"	minimum(1)
//	@Param		body	body		Message	true	"Message object"
//	@Success	201		{object}	PaginatedMessages
//	@Failure	422		{object}	RestError
//	@Failure	500		{object}	RestError
//	@Router		/v1/chats/{id}/message [post]
func (s Server) sendMessage() {}
