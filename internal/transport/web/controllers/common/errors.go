package common

import "github.com/ischenkx/vk-test-task/internal/transport/web/controllers/common/result"

const CustomErrorCode = 42

var InternalServerErr = result.NewError(1, "server failure")
var IncorrectInputErr = result.NewError(2, "incorrect data")
var UnauthorizedErr = result.NewError(3, "unauthorized")
var FailedToLoadErr = result.NewError(4, "failed to load data")
