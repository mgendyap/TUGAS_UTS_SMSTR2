packapackage messageModel

import (
	"utsstrukdat/db"
)

func FindMessage(user string) *[]db.Field_Message {
	var data_Message *db.Message = &db.Data_Message
	temp := data_Message.Next

	var response []db.Field_Message
	for temp != nil {
		if temp.Data.To == user {

			response = append(response, temp.Data)
		}
		temp = temp.Next
	}
	return &response
}

