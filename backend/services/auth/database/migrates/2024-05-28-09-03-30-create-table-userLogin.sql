CREATE TABLE IF NOT EXISTS userLogin (
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	CreatedAt TEXT DEFAULT CURRENT_TIMESTAMP,
	email TEXT NOT NULL,
	password TEXT NOT NULL
)