package template_4_your_project_name

import "errors"

// Domain-specific errors for the template4YourProjectName service
var (
	ErrNotFound                             = errors.New("template_4_your_project_name not found")
	ErrAlreadyExists                        = errors.New("template_4_your_project_name already exists")
	ErrTypetemplate4YourProjectNameNotFound = errors.New("type template_4_your_project_name not found")
	ErrUnauthorized                         = errors.New("unauthorized")
	ErrInvalidInput                         = errors.New("invalid input")
	ErrNotOwner                             = errors.New("user is not the owner")
	ErrAdminRequired                        = errors.New("admin privileges required")
)
