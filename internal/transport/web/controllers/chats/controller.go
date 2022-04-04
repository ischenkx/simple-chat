package chats

import (
	"encoding/json"
	"github.com/ischenkx/vk-test-task/internal/app"
	appForms "github.com/ischenkx/vk-test-task/internal/app/forms"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/chats/forms"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/common"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/common/result"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/dto"
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

func (c *Controller) GetChat(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.GetChat
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	chat, err := c.app.Chats().Get(ctx, form.ID)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	var chatDto dto.Chat

	if err := chatDto.Load(ctx, chat); err != nil {
		result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
		return
	}

	result.WriteSilent(w, result.Ok(chatDto))
}

func (c *Controller) DeleteChat(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.DeleteChat
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	chat, err := c.app.Chats().Get(ctx, form.ID)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	if err := chat.Delete(ctx); err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	result.WriteSilent(w, result.Ok(nil))
}

func (c *Controller) CreateChat(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.CreateChat
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	chat, err := c.app.Chats().Create(ctx, appForms.ChatCreationForm{
		Name:        form.Name,
		Description: form.Description,
	})

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	var chatDto dto.Chat

	if err := chatDto.Load(ctx, chat); err != nil {
		result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
		return
	}

	result.WriteSilent(w, result.Ok(chatDto))
}

func (c *Controller) GetChatMembers(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.GetChatMembers
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	chat, err := c.app.Chats().Get(ctx, form.ChatID)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	members, err := chat.Members(ctx, form.Offset, form.Count)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	var memberDtos []dto.User

	for _, mem := range members {
		user, err := mem.User(ctx)
		if err != nil {
			result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
			return
		}
		var memberDto dto.User
		if err := memberDto.Load(ctx, user); err != nil {
			result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
			return
		}
		memberDtos = append(memberDtos, memberDto)
	}

	result.WriteSilent(w, result.Ok(memberDtos))
}

func (c *Controller) DeleteChatMember(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.DeleteChatMember
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	chat, err := c.app.Chats().Get(ctx, form.ChatID)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	member, err := chat.Member(ctx, form.UserID)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	if err := member.Delete(ctx); err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	result.WriteSilent(w, result.Ok(nil))
}

func (c *Controller) CreateChatMember(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.CreateChatMember
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	chat, err := c.app.Chats().Get(ctx, form.ChatID)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	_, err = chat.Add(ctx, form.UserID, 0)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	result.WriteSilent(w, result.Ok(nil))
}

func (c *Controller) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.UpdateMessage
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	message, err := c.app.Chats().GetMessage(ctx, form.ID)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	if err := message.Update(ctx, appForms.MessageUpdate{Payload: form.Payload}); err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	var messageDto dto.Message

	if err := messageDto.Load(ctx, message); err != nil {
		result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
		return
	}

	result.WriteSilent(w, result.Ok(messageDto))
}

func (c *Controller) SendMessage(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.SendMessage
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	chat, err := c.app.Chats().Get(ctx, form.ChatID)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	member, err := chat.Member(ctx, ctx.User().ID())

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	mes, err := member.SendMessage(ctx, appForms.SendMessage{Payload: form.Payload})

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	var messageDto dto.Message

	if err := messageDto.Load(ctx, mes); err != nil {
		result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
		return
	}

	result.WriteSilent(w, result.Ok(messageDto))
}

func (c *Controller) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.DeleteMessage
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	message, err := c.app.Chats().GetMessage(ctx, form.ID)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	if err := message.Delete(ctx); err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	result.WriteSilent(w, result.Ok(nil))
}

func (c *Controller) GetMessages(w http.ResponseWriter, r *http.Request) {
	ctx, ok := util.AppContext(r.Context())
	if !ok {
		result.WriteSilent(w, result.New(nil, common.InternalServerErr))
		return
	}

	var form forms.GetMessages
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		result.WriteSilent(w, result.New(nil, common.IncorrectInputErr))
		return
	}

	chat, err := c.app.Chats().Get(ctx, form.ChatID)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	messages, err := chat.Messages(ctx, form.Offset, form.Count)

	if err != nil {
		result.WriteSilent(w, result.Err(common.CustomErrorCode, err.Error()))
		return
	}

	var messageDtos []dto.Message

	for _, req := range messages {
		var messageDto dto.Message
		if err := messageDto.Load(ctx, req); err != nil {
			result.WriteSilent(w, result.New(nil, common.FailedToLoadErr))
			return
		}
		messageDtos = append(messageDtos, messageDto)
	}

	result.WriteSilent(w, result.Ok(messageDtos))
}

func (c *Controller) init() {
	c.mux.HandleFunc("/createChatMember", c.CreateChatMember)
	c.mux.HandleFunc("/deleteChatMember", c.DeleteChatMember)
	c.mux.HandleFunc("/getChatMembers", c.GetChatMembers)
	c.mux.HandleFunc("/getChat", c.GetChat)
	c.mux.HandleFunc("/deleteChat", c.DeleteChat)
	c.mux.HandleFunc("/createChat", c.CreateChat)
	c.mux.HandleFunc("/sendMessage", c.SendMessage)
	c.mux.HandleFunc("/updateMessage", c.UpdateMessage)
	c.mux.HandleFunc("/deleteMessage", c.DeleteMessage)
	c.mux.HandleFunc("/getMessages", c.GetMessages)
}

func NewController(app *app.App) *Controller {
	controller := &Controller{
		app: app,
		mux: http.NewServeMux(),
	}

	controller.init()
	return controller
}
