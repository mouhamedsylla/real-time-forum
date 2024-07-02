package controllers

import (
	"net/http"
	"real-time-forum/services/chat/database"
	"real-time-forum/services/chat/models"
	"real-time-forum/utils"
	"strconv"
)

func (pm *GetPrivateMessage) HTTPServe() http.Handler {
	return http.HandlerFunc(pm.getPrivateMessage)
}

func (pm *GetPrivateMessage) EndPoint() string {
	return "/chat/message/private/:senderId/:receiverId"
}

func (pm *GetPrivateMessage) SetMethods() []string {
	return []string{http.MethodGet}
}

func (pm *GetPrivateMessage) getPrivateMessage(w http.ResponseWriter, r *http.Request) {
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)

	sendId, _ := strconv.Atoi(CustomRoute["senderId"])
	receiveId, _ := strconv.Atoi(CustomRoute["receiverId"])
	database.DbChat.Storage.Custom.
		Where("senderId", sendId).And("receiverId", receiveId).
		Or("senderId", receiveId).And("receiverId", sendId)

	result := database.DbChat.Storage.Scan(models.Message{}, "Id", "CreatedAt", "SenderId", "ReceiverId", "Content").([]models.Message)
	database.DbChat.Storage.Custom.Clear()

	params := r.URL.Query()
	limit := params.Get("limit")
	page := params.Get("page")
	if limit != "" && page != "" {
		limitInt, _ := strconv.Atoi(limit)
		pageInt, _ := strconv.Atoi(page)
		chunk := chunkMessage(limitInt, result)
		result = chunk[pageInt]
		utils.ResponseWithJSON(w, result, http.StatusOK)
		return
	}

	utils.ResponseWithJSON(w, result, http.StatusOK)
}

func chunkMessage(limit int, message []models.Message) map[int][]models.Message {
	// Reverse the input slice
	reversed := make([]models.Message, len(message))
	for i, v := range message {
		reversed[len(message)-1-i] = v
	}

	// Chunk the reversed slice
	chunk := make(map[int][]models.Message)
	chunkIndex := 1
	for _, v := range reversed {
		if len(chunk[chunkIndex]) < limit {
			chunk[chunkIndex] = append(chunk[chunkIndex], v)
		}
		if len(chunk[chunkIndex]) == limit {
			chunkIndex++
		}
	}
	return chunk
}
