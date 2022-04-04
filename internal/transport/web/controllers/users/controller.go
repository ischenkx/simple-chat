package users

import (
	"encoding/json"
	"github.com/ischenkx/vk-test-task/internal/app"
	appForms "github.com/ischenkx/vk-test-task/internal/app/forms"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/common"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/common/auth"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/common/result"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/dto"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/users/forms"
	"github.com/ischenkx/vk-test-task/internal/transport/web/util"
	"net/http"
)

type Controller struct {
	app *app.App
	mux *http.ServeMux
}

func (c *Controller) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c.mux.ServeHTTP(writer, request)
}

func (c *Controller) SendFriendRequest(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.SendFriendRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	if ctx.User() == nil {
		result.WriteSilent(w, result.New(nil, common.UnauthorizedErr))
		return
	}

	friendRequest, err := ctx.User().SendFriendRequest(ctx, form.To)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	var requestDto dto.FriendRequest
	if err := requestDto.Load(ctx, friendRequest); err != nil {
		result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
		return
	}

	result.WriteSilent(w, result.Ok(requestDto))
}

func (c *Controller) DeclineFriendRequest(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.DeclineFriendRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	if ctx.User() == nil {
		result.WriteSilent(w, result.New(nil, common.UnauthorizedErr))
		return
	}

	friendRequest, err := ctx.User().IncomingFriendRequest(ctx, form.ID)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	if err := friendRequest.Decline(ctx); err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	result.WriteSilent(w, result.Ok(nil))
}

func (c *Controller) AcceptFriendRequest(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.AcceptFriendRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	if ctx.User() == nil {
		result.WriteSilent(w, result.New(nil, common.UnauthorizedErr))
		return
	}

	friendRequest, err := ctx.User().IncomingFriendRequest(ctx, form.ID)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	if err := friendRequest.Accept(ctx); err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	result.WriteSilent(w, result.Ok(nil))
}

func (c *Controller) GetInfo(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	if ctx.User() == nil {
		result.WriteSilent(w, result.New(nil, common.UnauthorizedErr))
		return
	}

	var user dto.User
	if err := user.Load(ctx, ctx.User()); err != nil {
		result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
		return
	}

	result.WriteSilent(w, result.Ok(user))
}

func (c *Controller) GetOutgoingFriendRequests(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.GetOutgoingFriendRequests
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	if ctx.User() == nil {
		result.WriteSilent(w, result.New(nil, common.UnauthorizedErr))
		return
	}

	friendRequests, err := ctx.User().OutgoingFriendRequests(ctx, form.Offset, form.Count)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	var requestDtos []dto.FriendRequest

	for _, req := range friendRequests {
		var dtoRequest dto.FriendRequest
		if err := dtoRequest.Load(ctx, req); err != nil {
			result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
			return
		}
		requestDtos = append(requestDtos, dtoRequest)
	}

	result.WriteSilent(w, result.Ok(requestDtos))
}

func (c *Controller) GetIncomingFriendRequests(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.GetIncomingFriendRequests
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	if ctx.User() == nil {
		result.WriteSilent(w, result.New(nil, common.UnauthorizedErr))
		return
	}

	friendRequests, err := ctx.User().IncomingFriendRequests(ctx, form.Offset, form.Count)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	var requests []dto.FriendRequest

	for _, req := range friendRequests {
		var dtoRequest dto.FriendRequest
		if err := dtoRequest.Load(ctx, req); err != nil {
			result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
			return
		}
		requests = append(requests, dtoRequest)
	}

	result.WriteSilent(w, result.Ok(requests))
}

func (c *Controller) GetFriends(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.GetFriends
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	if ctx.User() == nil {
		result.WriteSilent(w, result.New(nil, common.UnauthorizedErr))
		return
	}

	friends, err := ctx.User().Friends(ctx, form.Offset, form.Count)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	var friendsDto []dto.User

	for _, friend := range friends {
		var friendDto dto.User
		user, err := friend.Friend(ctx)
		if err != nil {
			result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
			return
		}
		if err := friendDto.Load(ctx, user); err != nil {
			result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
			return
		}
		friendsDto = append(friendsDto, friendDto)
	}

	result.WriteSilent(w, result.Ok(friendsDto))
}

func (c *Controller) GetChats(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.GetChats
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	if ctx.User() == nil {
		result.WriteSilent(w, result.New(nil, common.UnauthorizedErr))
		return
	}

	members, err := ctx.User().Chats(ctx, form.Offset, form.Count)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	var chatsDto []dto.Chat

	for _, member := range members {
		var chatDto dto.Chat
		chat, err := member.Chat(ctx)
		if err != nil {
			result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
			return
		}
		if err := chatDto.Load(ctx, chat); err != nil {
			result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
			return
		}
		chatsDto = append(chatsDto, chatDto)
	}

	result.WriteSilent(w, result.Ok(chatsDto))
}

func (c *Controller) Logout(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	if ctx.User() == nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, "not authorized"))
		return
	}

	auth.DeleteVerificationToken(w, r)
	ctx.SetUser(nil)

	result.WriteSilent(w, result.Ok(nil))
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.Login
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	if ctx.User() != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, "already authorized"))
		return
	}

	user, err := c.app.Users().Login(ctx, appForms.UserLogin{
		Username: form.Username,
		Password: form.Password,
	})

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	ctx.SetUser(user)
	token, err := c.app.Auth().GenerateToken(ctx, user.ID())

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, "failed to generate a jwt token"))
		return
	}

	auth.StoreVerificationToken(w, r, token)

	var userDto dto.User
	if err := userDto.Load(ctx, user); err != nil {
		result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
		return
	}

	result.WriteSilent(w, result.Ok(userDto))
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.Register
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	if ctx.User() != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, "already authorized"))
		return
	}

	user, err := c.app.Users().Register(ctx, appForms.UserRegistration{
		Username: form.Username,
		Password: form.Password,
	})

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	ctx.SetUser(user)
	token, err := c.app.Auth().GenerateToken(ctx, user.ID())

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, "failed to generate a jwt token"))
		return
	}

	auth.StoreVerificationToken(w, r, token)

	var userDto dto.User
	if err := userDto.Load(ctx, user); err != nil {
		result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
		return
	}

	result.WriteSilent(w, result.Ok(userDto))
}

func (c *Controller) init() {
	c.mux.HandleFunc("/register", c.Register)
	c.mux.HandleFunc("/login", c.Login)
	c.mux.HandleFunc("/logout", c.Logout)
	c.mux.HandleFunc("/getInfo", c.GetInfo)
	c.mux.HandleFunc("/getFriends", c.GetFriends)
	c.mux.HandleFunc("/getChats", c.GetChats)
	c.mux.HandleFunc("/getIncomingFriendRequests", c.GetIncomingFriendRequests)
	c.mux.HandleFunc("/getOutgoingFriendRequests", c.GetOutgoingFriendRequests)
	c.mux.HandleFunc("/sendFriendRequest", c.SendFriendRequest)
	c.mux.HandleFunc("/declineFriendRequest", c.DeclineFriendRequest)
	c.mux.HandleFunc("/acceptFriendRequest", c.AcceptFriendRequest)
}

func NewController(app *app.App) *Controller {
	controller := &Controller{
		app: app,
		mux: http.NewServeMux(),
	}

	controller.init()
	return controller
}
