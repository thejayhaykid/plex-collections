package models

// MapUserDAOtoUserDTO for HTTP response
func MapUserDAOtoUserDTO(user User) UserDTO {
	return UserDTO{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Role:   user.Role,
		Active: user.Active,
	}
}

// MapUserDAOToUserAuthenticationResponse maps a user dao to an authentication response
func MapUserDAOToUserAuthenticationResponse(user User, token string) UserAuthenticationResponse {
	return UserAuthenticationResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Role:   user.Role,
		Active: user.Active,
		Token:  token,
	}
}
