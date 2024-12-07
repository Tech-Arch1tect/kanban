# go_react_web_app_template

This is a starting point for other projects. It's basic, but useful for quick proof-of-concepts.

The interaction between client <--> server is all based around an  OpenAPI spec. The OpenAPI spec is built from annotations within the Go server controllers. We then use openapi generator to create a basic API client for the frontend to utilise.

## Server
* Authentication (login, register, forgot-password, 2fa (totp))
* Authorisation (roles)
* Main technologies: Go, Gin framework, GORM

#### Improvements to make

* ~~Better database module - generic interface to create/edit/update/delete which each model repository can implement or extend~~ <https://github.com/Tech-Arch1tect/go_react_web_app_template/commit/716160a06ec414233113c8097088582fccea5175>
* Cleaner code
* ~~Security enhancements~~ (~~brute force protection~~, ~~CSRF~~, etc?)

## Client
* Authentication (login, register, forgot-password, 2fa (totp))
* Admin / client menu ready to extend
* Basic user managment (e.g. change users roles etc)
* Main technologies: React (+tanstack router), Typescript, Vite

#### Improvements to make
* Move to tanstack query
* Improve component organisation
* Better error handling
