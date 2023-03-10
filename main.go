package main

import (
	"berlinger/controller"
	"berlinger/rules"
	"fmt"
)

//! in this project you learn to call methods from different structs in some other struct USING INTERFACES
// (ie: call methords from controller and rules struct in handler struct using interfaces)
// later todo: how to call them in func main 

type Handler struct {
	HandlerInterface 
	Controller        controller.Interface	
	Rules             rules.Interface	
}

type HandlerInterface interface{
	Handlersfunc()
}

func (h *Handler) Handlersfunc(){
	h.Rules.Rulesfunc()
	h.Controller.Controllerfunc()
	fmt.Println("")
}

func NewHandler() HandlerInterface {
	return &Handler{
		Controller: controller.NewController(), //controller: bcoz pkg name is controller and rules
		Rules:      rules.NewRules(),			 
		// these newhandler, newrules, and newcontrollers are used to create a instance in different page
	}
}

func main(){
	handler := NewHandler()
	handler.Handlersfunc()
}