package utils

import (
	models "ecommerce-website/app/Models"
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
	"net/url"
)

type ErrorResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"message"`
}

// GetError : This is helper function to prepare error model.
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func GetError(err error, w http.ResponseWriter) {

	log.Printf(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		Success:      false,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(message)

}

func Validate(product models.Product) url.Values {
	errors := url.Values{}

	if product.Name == "" {
		errors.Add("name", "Please enter product name!")
	}

	if product.Description == "" {
		errors.Add("description", "Please enter product description!")
	}

	if product.Price == 0 {
		errors.Add("price", "Please enter product price!")
	}

	if product.Price > 99999999 {
		errors.Add("prices", "Product price can not exceed length 8!")
	}

	if product.Ratings <= 0 {
		errors.Add("ratings", "Product ratings can not be negative or empty!")
	}

	if len(product.Images) == 0 {
		errors.Add("images", "Product images can not be empty!")
	}

	if product.Category == "" {
		errors.Add("category", "Product category can not be empty!")
	}

	if product.Stock == 0 {
		errors.Add("stock", "Please enter product stock!")
	}

	return errors
}

func UserRegisterValidation(user models.User) url.Values {

	errors := url.Values{}

	if user.Name == "" {
		errors.Add("name", "Please enter user name!")
	}

	if len(user.Name) > 30 || len(user.Name) < 4 {
		errors.Add("name", "User name should lie between length 4 and 30!")
	}

	if len(user.Email) == 0 {
		errors.Add("email", "User email address is mandatory!")
	}

	if len(user.Password) < 8 {
		errors.Add("password", "User password should be atleast 8 characters long!")
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		errors.Add("email", "Invalid email address given!")
	}

	if len(user.Avatar) < 1 {
		errors.Add("avatar", "User should provide profile image!")
	}

	return errors
}

func UserLoginValidation(user models.User) url.Values {

	errors := url.Values{}

	if len(user.Email) == 0 {
		errors.Add("email", "User email address is mandatory!")
	}

	if len(user.Password) == 0 {
		errors.Add("password", "User password is required to login!")
	}

	return errors
}
