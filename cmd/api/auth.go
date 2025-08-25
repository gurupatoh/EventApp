package main

type registerRequest struct{
	Email string  `json :"email" binding:"required,email`
}