package auth

type AuthRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginResponse struct {
	Message string `json:"message"`
}

type AuthChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=8"`
}

type AuthChangePasswordResponse struct {
	Message string `json:"message"`
}

type AuthChangeUsernameRequest struct {
	Username string `json:"username" binding:"required"`
}

type AuthChangeUsernameResponse struct {
	Message string `json:"message"`
}

type AuthPasswordResetRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type AuthPasswordResetResponse struct {
	Message string `json:"message"`
}

type AuthResetPasswordRequest struct {
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type AuthResetPasswordResponse struct {
	Message string `json:"message"`
}

type AuthGetCSRFTokenResponse struct {
	CSRFToken string `json:"csrf_token"`
}

type AuthGenerateTOTPResponse struct {
	Secret string `json:"secret"`
}

type AuthEnableTOTPRequest struct {
	Code string `json:"code" binding:"required"`
}

type AuthEnableTOTPResponse struct {
	Message string `json:"message"`
}

type AuthDisableTOTPRequest struct {
	Code string `json:"code" binding:"required"`
}

type AuthDisableTOTPResponse struct {
	Message string `json:"message"`
}

type AuthConfirmTOTPRequest struct {
	Code string `json:"code" binding:"required"`
}

type AuthConfirmTOTPResponse struct {
	Message string `json:"message"`
}
