CREATE TABLE IF NOT EXISTS ReactionComment (
	Value TEXT ,
	CommentId INTEGER ,
	FOREIGN KEY (CommentId) REFERENCES Comments (Id)
)