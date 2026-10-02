package myTypes

import "errors"

var ErrDatabseOffline error = errors.New("Database Offline")
var ErrBlogNotFound error = errors.New("Blog post with this name not found")
var ErrUnknownDatabaseError error = errors.New("Encountered an unknown databse error")
