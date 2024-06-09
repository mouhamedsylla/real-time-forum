package posts

// func (p *ReactionPost) HTTPServe() http.Handler {
// 	return http.HandlerFunc(p.ReactionPost)
// }

// func (p *ReactionPost) EndPoint() string {
// 	return "/posts/reaction"
// }

// func (p *ReactionPost) SetMethods() []string {
// 	return []string{"GET"}
// }

// func (p *ReactionPost) ReactionPost(w http.ResponseWriter, r *http.Request) {

// 	// storage.Custom.Where("Post_id", "1")

// 	storage.Custom.OrderBy("Id", 1).Limit(1)

// 	if result == nil {
// 		utils.ResponseWithJSON(w, "Service Posts.ReactionPost: 400 BadRequest", http.StatusBadRequest)
// 		return
// 	}

// 	storage.Custom.Clear()

// 	fmt.Println(result)

// 	utils.ResponseWithJSON(w, result, http.StatusOK)
// }
