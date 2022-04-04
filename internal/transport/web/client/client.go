package client

import (
	"bytes"
	"context"
	"encoding/json"
	chatForms "github.com/ischenkx/vk-test-task/internal/transport/web/controllers/chats/forms"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/common/result"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/dto"
	userForms "github.com/ischenkx/vk-test-task/internal/transport/web/controllers/users/forms"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	http    *http.Client
	cookies map[string]*http.Cookie
	baseUrl string
}

func (c *Client) url(suffix string) string {
	if !strings.HasPrefix(suffix, "/") {
		suffix += "/"
	}
	base := strings.TrimSuffix(c.baseUrl, "/")
	return base + suffix
}

func (c *Client) encode(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (c *Client) decode(bts []byte, data interface{}) error {
	var res result.Result
	res.Data = data

	err := json.Unmarshal(bts, &res)
	if err != nil {
		return err
	}

	if res.Error == nil {
		return nil
	}

	return res.Error
}

func (c *Client) prepareRawRequest(req *http.Request) {
	for _, cookie := range c.cookies {
		req.AddCookie(cookie)
	}
}

func (c *Client) rawRequest(method, path string, data interface{}) ([]byte, error) {
	url := c.url(path)
	body, err := c.encode(data)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	c.prepareRawRequest(req)

	res, err := c.http.Do(req)

	if err != nil {
		return nil, err
	}

	for _, cookie := range res.Cookies() {
		c.cookies[cookie.Name] = cookie
	}

	return io.ReadAll(res.Body)
}

func (c *Client) request(path, method string, input, output interface{}) error {
	bts, err := c.rawRequest(method, path, input)

	if err != nil {
		return err
	}

	return c.decode(bts, output)
}

func (c *Client) post(path string, input, output interface{}) error {
	return c.request(path, "POST", input, output)
}

func (c *Client) get(path string, input, output interface{}) error {
	return c.request(path, "GET", input, output)
}

func (c *Client) Register(form userForms.Register) (dto.User, string, error) {
	var res dto.User

	if err := c.post("/users/register", form, &res); err != nil {
		return res, "", err
	}

	token := ""

	if cookie, ok := c.cookies["auth_token"]; ok {
		token = cookie.Value
	}

	return res, token, nil
}

func (c *Client) Login(form userForms.Login) (dto.User, string, error) {
	var res dto.User

	if err := c.post("/users/login", form, &res); err != nil {
		return res, "", err
	}

	token := ""

	if cookie, ok := c.cookies["auth_token"]; ok {
		token = cookie.Value
	}

	return res, token, nil
}

func (c *Client) Logout(ctx context.Context) error {
	if err := c.post("/users/logout", nil, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) Info(ctx context.Context) (dto.User, error) {
	var res dto.User

	if err := c.get("/users/getInfo", nil, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (c *Client) Friends(form userForms.GetFriends) ([]dto.User, error) {
	var res []dto.User

	if err := c.post("/users/getFriends", form, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (c *Client) Chats(form userForms.GetChats) ([]dto.Chat, error) {
	var res []dto.Chat

	if err := c.post("/users/getChats", form, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (c *Client) IncomingFriendRequests(form userForms.GetIncomingFriendRequests) ([]dto.FriendRequest, error) {
	var res []dto.FriendRequest

	if err := c.post("/users/getIncomingFriendRequests", form, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (c *Client) OutgoingFriendRequests(form userForms.GetOutgoingFriendRequests) ([]dto.FriendRequest, error) {
	var res []dto.FriendRequest

	if err := c.post("/users/getOutgoingFriendRequests", form, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (c *Client) SendFriendRequest(form userForms.SendFriendRequest) (dto.FriendRequest, error) {
	var res dto.FriendRequest

	if err := c.post("/users/sendFriendRequest", form, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (c *Client) DeclineFriendRequest(form userForms.DeclineFriendRequest) error {
	if err := c.post("/users/declineFriendRequest", form, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) AcceptFriendRequest(form userForms.AcceptFriendRequest) error {
	if err := c.post("/users/acceptFriendRequest", form, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) Chat(form chatForms.GetChat) (dto.Chat, error) {
	var res dto.Chat
	if err := c.post("/chats/getChat", form, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (c *Client) CreateChatMember(form chatForms.CreateChatMember) error {
	if err := c.post("/chats/createChatMember", form, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteChatMember(form chatForms.DeleteChatMember) error {
	if err := c.post("/chats/deleteChatMembers", form, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteChat(form chatForms.DeleteChat) error {
	if err := c.post("/chats/deleteChat", form, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateChat(form chatForms.CreateChat) (dto.Chat, error) {
	var res dto.Chat
	if err := c.post("/chats/createChat", form, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (c *Client) SendMessage(form chatForms.SendMessage) (dto.Message, error) {
	var res dto.Message
	if err := c.post("/chats/sendMessage", form, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (c *Client) UpdateMessage(form chatForms.UpdateMessage) (dto.Message, error) {
	var res dto.Message
	if err := c.post("/chats/updateMessage", form, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (c *Client) DeleteMessage(form chatForms.DeleteMessage) error {
	if err := c.post("/chats/deleteMessage", form, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetMessages(form chatForms.GetMessages) ([]dto.Message, error) {
	var res []dto.Message
	if err := c.post("/chats/getMessages", form, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (c *Client) GetChatMembers(form chatForms.GetChatMembers) ([]dto.User, error) {
	var res []dto.User
	if err := c.post("/chats/getChatMembers", form, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (c *Client) SetToken(token string) {
	c.cookies["auth_token"] = &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	}
}

func (c *Client) SetBaseUrl(url string) {
	c.baseUrl = url
}

func New(client *http.Client) *Client {
	return &Client{
		http:    client,
		cookies: map[string]*http.Cookie{},
		baseUrl: "",
	}
}
