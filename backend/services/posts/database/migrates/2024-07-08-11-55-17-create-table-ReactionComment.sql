CREATE TABLE IF NOT EXISTS ReactionComment (
	Value TEXT ,
	CommentId INTEGER ,
	UserId INTEGER ,
	FOREIGN KEY (CommentId) REFERENCES Comments (Id)
)