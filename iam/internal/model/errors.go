package model

import "errors"

var SessionNotFound = errors.New("session not found")

var AlreadyRegistered = errors.New("user is already registered")

var AuthErr = errors.New("invalid username or password")