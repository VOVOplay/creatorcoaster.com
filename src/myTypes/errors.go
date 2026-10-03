package myTypes

import "errors"

var ErrDatabseOffline error = errors.New("Database Offline")
var ErrArticleNotFound error = errors.New("Article not found")
var ErrUnknownDatabaseError error = errors.New("Encountered an unknown databse error")
