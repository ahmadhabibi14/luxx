package auth

var (
	register_success     string = "User created successfully!"
	login_success        string = "Login successful !"
	invalid_input        string = "The payload or input provided is invalid. Please check your request and try again."
	user_not_found       string = "User not found"
	invalid_password     string = "Password does not match the user's password."
	error_generate_token string = "Unable to generate session token"
	username_exist       string = "Username already exists, try another one"
	email_exist          string = "Email already exist, try another one"
	uncaught_error       string = "Uncaught error"
)
