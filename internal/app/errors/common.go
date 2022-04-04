package errors

import "errors"

var ResourceInaccessible = errors.New("resource inaccessible")
var DoesNotExist = errors.New("does not exist")
var NotAuthorized = errors.New("not authorized")
var AlreadyAuthorized = errors.New("already authorized")
var RightsViolation = errors.New("not enough rights")
