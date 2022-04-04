package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ischenkx/vk-test-task/internal/transport/web/client"
	chatForms "github.com/ischenkx/vk-test-task/internal/transport/web/controllers/chats/forms"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/dto"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/users/forms"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

func ValidateInteger(input string) error {
	_, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return errors.New("Invalid number")
	}
	return nil
}

func promptString(message string) *promptui.Prompt {
	return &promptui.Prompt{
		Label: message,
	}
}

func promptInt(message string) *promptui.Prompt {
	return &promptui.Prompt{
		Label:    message,
		Validate: ValidateInteger,
	}
}

func promptOffsetCount() (int, int, error) {
	offsetStr, err := promptInt("offset").Run()
	if err != nil {
		return 0, 0, err
	}

	countStr, err := promptInt("count").Run()
	if err != nil {
		return 0, 0, err
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return 0, 0, err
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 0, 0, err
	}
	return offset, count, nil
}

func output(data interface{}, prefixTabs int) {
	prefix := strings.Repeat("\t", prefixTabs)
	switch obj := data.(type) {
	case dto.User:
		fmt.Printf(prefix+"username: '%s'\n", obj.Username)
		fmt.Printf(prefix+"id: '%s'\n", obj.ID)
	case dto.Chat:
		fmt.Printf(prefix+"name: '%s'\n", obj.Name)
		fmt.Printf(prefix+"description: '%s'\n", obj.Description)
		fmt.Printf(prefix+"owner id: '%s'\n", obj.OwnerID)
		fmt.Printf(prefix+"id: '%s'\n", obj.ID)
	case dto.Message:
		fmt.Printf(prefix+"user id: '%s'\n", obj.UserID)
		fmt.Printf(prefix+"chat id: '%s'\n", obj.ChatID)
		fmt.Printf(prefix+"payload: '%s'\n", obj.Payload)
		fmt.Printf(prefix+"time: '%s'\n", obj.TimeStamp)
		fmt.Printf(prefix+"last update: '%s'\n", obj.LastUpdate)
		fmt.Printf(prefix+"id: '%s'\n", obj.ID)

	case dto.FriendRequest:
		fmt.Printf(prefix+"from: '%s'\n", obj.FromID)
		fmt.Printf(prefix+"to: '%s'\n", obj.ToID)
		fmt.Printf(prefix+"time: '%s'\n", obj.TimeStamp)
		fmt.Printf(prefix+"id: '%s'\n", obj.ID)
	case error:
		fmt.Println(prefix+"error:", obj)
	default:
		fmt.Print(prefix)
		fmt.Println(obj)
	}
}

func outputBreakLine(tabs int) {
	output(strings.Repeat("-", 27), tabs)
}

func main() {

	appClient := client.New(&http.Client{})

	if len(os.Args) >= 2 {
		appClient.SetBaseUrl(os.Args[1])
	}

	ctx := context.Background()
	for {
		res, err := promptString("command").Run()
		if err != nil {
			panic(err)
		}

		switch res {
		case "seturl":
			url, err := promptString("url").Run()
			if err != nil {
				output(err, 1)
				continue
			}
			appClient.SetBaseUrl(url)

		case "info":
			info, err := appClient.Info(ctx)

			if err != nil {
				output(err, 1)
			} else {
				output(info, 1)
			}

		case "login":
			username, err := promptString("username").Run()
			if err != nil {
				output(err, 1)
				continue
			}
			password, err := promptString("password").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			user, token, err := appClient.Login(forms.Login{
				Username: username,
				Password: password,
			})

			if err != nil {
				output(err, 1)
				continue
			}

			appClient.SetToken(token)
			output(user, 1)

		case "logout":
			err := appClient.Logout(ctx)

			if err != nil {
				output(err, 1)
				continue
			}

			appClient.SetToken("")

		case "register":
			username, err := promptString("username").Run()
			if err != nil {
				output(err, 1)
				continue
			}
			password, err := promptString("password").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			user, token, err := appClient.Register(forms.Register{
				Username: username,
				Password: password,
			})

			if err != nil {
				output(err, 1)
				continue
			}

			appClient.SetToken(token)
			output(user, 1)

		case "chats":
			offset, count, err := promptOffsetCount()
			if err != nil {
				output(err, 1)
				continue
			}

			chats, err := appClient.Chats(forms.GetChats{
				Offset: offset,
				Count:  count,
			})

			if err != nil {
				output(err, 1)
				continue
			}

			for _, chat := range chats {
				output(chat, 1)
				outputBreakLine(1)
			}

		case "friends":
			offset, count, err := promptOffsetCount()
			if err != nil {
				output(err, 1)
				continue
			}

			friends, err := appClient.Friends(forms.GetFriends{
				Offset: offset,
				Count:  count,
			})

			if err != nil {
				output(err, 1)
				continue
			}

			for _, friend := range friends {
				output(friend, 1)
				outputBreakLine(1)
			}

		case "incoming-fr":
			offset, count, err := promptOffsetCount()
			if err != nil {
				output(err, 1)
				continue
			}

			requests, err := appClient.IncomingFriendRequests(forms.GetIncomingFriendRequests{
				Offset: offset,
				Count:  count,
			})

			if err != nil {
				output(err, 1)
				continue
			}

			for _, req := range requests {
				output(req, 1)
				outputBreakLine(1)
			}

		case "outgoing-fr":
			offset, count, err := promptOffsetCount()
			if err != nil {
				output(err, 1)
				continue
			}

			requests, err := appClient.OutgoingFriendRequests(forms.GetOutgoingFriendRequests{
				Offset: offset,
				Count:  count,
			})

			if err != nil {
				output(err, 1)
				continue
			}

			for _, req := range requests {
				output(req, 1)
				outputBreakLine(1)
			}

		case "send-fr":
			to, err := promptString("to").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			req, err := appClient.SendFriendRequest(forms.SendFriendRequest{
				To: to,
			})

			if err != nil {
				output(err, 1)
				continue
			}
			output(req, 1)

		case "accept-fr":
			id, err := promptString("id").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			err = appClient.AcceptFriendRequest(forms.AcceptFriendRequest{
				ID: id,
			})

			if err != nil {
				output(err, 1)
				continue
			}

		case "decline-fr":
			id, err := promptString("id").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			err = appClient.DeclineFriendRequest(forms.DeclineFriendRequest{
				ID: id,
			})

			if err != nil {
				output(err, 1)
				continue
			}

		case "chat":
			id, err := promptString("id").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			res, err := appClient.Chat(chatForms.GetChat{ID: id})

			if err != nil {
				output(err, 1)
				continue
			}

			output(res, 1)

		case "delete-chat":
			chatId, err := promptString("chat id").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			err = appClient.DeleteChat(chatForms.DeleteChat{
				ID: chatId,
			})

			if err != nil {
				output(err, 1)
				continue
			}

		case "create-chat":
			name, err := promptString("name").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			desc, err := promptString("description").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			chat, err := appClient.CreateChat(chatForms.CreateChat{
				Name:        name,
				Description: desc,
			})

			if err != nil {
				output(err, 1)
				continue
			}

			output(chat, 1)

		case "create-chat-member":
			chatID, err := promptString("chat id").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			userID, err := promptString("user id").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			err = appClient.CreateChatMember(chatForms.CreateChatMember{
				ChatID: chatID,
				UserID: userID,
			})

			if err != nil {
				output(err, 1)
				continue
			}

		case "delete-chat-member":
			chatID, err := promptString("chat id").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			userID, err := promptString("user id").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			err = appClient.DeleteChatMember(chatForms.DeleteChatMember{
				ChatID: chatID,
				UserID: userID,
			})

			if err != nil {
				output(err, 1)
				continue
			}

		case "get-chat-members":
			chatID, err := promptString("chat id").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			offset, count, err := promptOffsetCount()
			if err != nil {
				output(err, 1)
				continue
			}

			members, err := appClient.GetChatMembers(chatForms.GetChatMembers{
				ChatID: chatID,
				Offset: offset,
				Count:  count,
			})

			if err != nil {
				output(err, 1)
				continue
			}

			for _, mem := range members {
				output(mem, 1)
				outputBreakLine(1)
			}

		case "send-message":
			chatId, err := promptString("chat-id").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			payload, err := promptString("payload").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			res, err := appClient.SendMessage(chatForms.SendMessage{
				ChatID:  chatId,
				Payload: payload,
			})

			if err != nil {
				output(err, 1)
				continue
			}

			output(res, 1)

		case "update-message":
			id, err := promptString("id").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			payload, err := promptString("payload").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			res, err := appClient.UpdateMessage(chatForms.UpdateMessage{
				ID:      id,
				Payload: payload,
			})

			if err != nil {
				output(err, 1)
				continue
			}

			output(res, 1)

		case "delete-message":
			id, err := promptString("id").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			err = appClient.DeleteMessage(chatForms.DeleteMessage{
				ID: id,
			})

			if err != nil {
				output(err, 1)
				continue
			}

		case "messages":

			offset, count, err := promptOffsetCount()
			if err != nil {
				output(err, 1)
				continue
			}
			chatID, err := promptString("chat id").Run()
			if err != nil {
				output(err, 1)
				continue
			}

			messages, err := appClient.GetMessages(chatForms.GetMessages{
				ChatID: chatID,
				Offset: offset,
				Count:  count,
			})

			if err != nil {
				output(err, 1)
				continue
			}

			for _, mes := range messages {
				output(mes, 1)
				outputBreakLine(1)
			}
		case "kill":
			return
		}
	}
}
