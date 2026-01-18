package response

var (

	// ========================
	// SUCCESS (1xxx)
	// ========================

	GetUsersSuccess = Code{
		Code:    1000,
		Message: "Get users successfully",
	}

	GetUserSuccess = Code{
		Code:    1001,
		Message: "Get user successfully",
	}

	CreateUserSuccess = Code{
		Code:    1002,
		Message: "Create user successfully",
	}

	UpdateUserSuccess = Code{
		Code:    1003,
		Message: "Update user successfully",
	}

	DeleteUserSuccess = Code{
		Code:    1004,
		Message: "Delete user successfully",
	}

	// ========================
	// VALIDATION (2xxx)
	// ========================

	ValidationError = Code{
		Code:      2001,
		ErrorCode: "VALIDATION_ERROR",
		Message:   "invalid data",
	}

	InvalidID = Code{
		Code:      2002,
		ErrorCode: "INVALID_ID",
		Message:   "invalid id",
	}

	MissingRequiredField = Code{
		Code:      2003,
		ErrorCode: "MISSING_REQUIRED_FIELD",
		Message:   "required field is missing",
	}

	// ========================
	// USER (3xxx)
	// ========================

	UserNotFound = Code{
		Code:      3001,
		ErrorCode: "USER_NOT_FOUND",
		Message:   "user not found",
	}

	UserDuplicate = Code{
		Code:      3002,
		ErrorCode: "USER_DUPLICATE",
		Message:   "user already exists",
	}

	UserUpdateFailed = Code{
		Code:      3003,
		ErrorCode: "USER_UPDATE_FAILED",
		Message:   "update user failed",
	}

	// ========================
	// AUTH (11xx)
	// ========================

	AuthInvalidToken = Code{
		Code:      1101,
		ErrorCode: "AUTH_INVALID_TOKEN",
		Message:   "invalid token",
	}

	AuthTokenExpired = Code{
		Code:      1102,
		ErrorCode: "AUTH_TOKEN_EXPIRED",
		Message:   "token expired",
	}

	AuthPermissionDenied = Code{
		Code:      1103,
		ErrorCode: "AUTH_PERMISSION_DENIED",
		Message:   "permission denied",
	}

	// ========================
	// SERVER (9xxx)
	// ========================

	InternalServerError = Code{
		Code:      9000,
		ErrorCode: "INTERNAL_SERVER_ERROR",
		Message:   "internal server error",
	}
)
