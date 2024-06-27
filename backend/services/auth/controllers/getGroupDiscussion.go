package controllers

import (
	"log"
	"net/http"
	"os"
	"real-time-forum/services/auth/database"
	"real-time-forum/services/auth/models"
	"real-time-forum/utils"
)

func (gtuser *GetGroupUserDiscussion) HTTPServe() http.Handler {
	return http.HandlerFunc(gtuser.GetGroupUserDiscussion)
}

func (gtuser *GetGroupUserDiscussion) EndPoint() string {
	return "/auth/getGroupUser/:userId"
}

func (gtuser *GetGroupUserDiscussion) SetMethods() []string {
	return []string{"GET"}
}

func (gtuser *GetGroupUserDiscussion) GetGroupUserDiscussion(w http.ResponseWriter, r *http.Request) {
	CustomRouter := r.Context().Value("CustomRoute").(map[string]string)
	userId, ok := CustomRouter["userId"]
	if !ok {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}

	AuthClient.SetMethod("GET")
	err := utils.LoadEnv("../../.env")
	if err != nil {
		log.Println(err)
	}
	baseUrl := os.Getenv("CHAT_SERVICE")
	AuthClient.SetBaseURL(baseUrl[1 : len(baseUrl)-1])

	var users_discussionsId models.UserContact
	err = AuthClient.Call("chat", "message/private/users/"+userId, nil, &users_discussionsId)
	if err != nil {
		log.Println(err)
	}

	database.Db.Storage.Custom.WhereIn("Id", IntSliceToInterfaceSlice(users_discussionsId.UsersId))
	result, ok := database.Db.Storage.Scan(models.UserRegister{}, "Id", "Nickname", "FirstName", "LastName").([]models.UserRegister)
	database.Db.Storage.Custom.Clear()

	if !ok {
		utils.ResponseWithJSON(w, "Service Auth.GetGroupUserDiscussion: Bad Request", http.StatusBadRequest)
		return
	}

	utils.ResponseWithJSON(w, result, http.StatusOK)
}

func IntSliceToInterfaceSlice(ints []int) []interface{} {
	interfaces := make([]interface{}, len(ints))
	for i, v := range ints {
		interfaces[i] = v
	}
	return interfaces
}
