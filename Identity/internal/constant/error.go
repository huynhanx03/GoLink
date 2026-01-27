package constant

// Specific Error Messages
const (
	MsgInvalidAuth       = "invalid username or password"
	MsgInvalidCredData   = "invalid credential data"
	MsgUserNoPass        = "user has no password set"
	MsgPassIncorrect     = "current password incorrect"
	MsgUsernameExists    = "username already exists"
	MsgUserNotMember     = "user is not a member of this tenant"
	MsgUnauthorized      = "unauthorized"
	MsgResourceNotFound  = "resource not found"
	MsgRebuildTreeFailed = "failed to rebuild role tree"
	MsgInvalidParentID   = "cannot set parent to self"
)
