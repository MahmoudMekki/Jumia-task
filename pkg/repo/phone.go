package repo

import (
	model "jumia-task/pkg/models"
	"regexp"
	"strings"
)

type PhoneRepo interface {
	ValidateNumber(number,regex string) bool
	GetCountryKeyAndPhoneNumber(phone string)(key string,number string)
	GetCountryPhoneDetails(key string)(model.Phone,bool)
	GetAvailablePhones()(phones []model.Phone)
}

func NewPhoneRepo() PhoneRepo {
	return &phoneImp{}
}

type phoneImp struct {

}

func (p phoneImp)ValidateNumber(number,regex string) bool {
	re := regexp.MustCompile(regex)
	isValid := re.MatchString(number)
	return isValid
}
func (p phoneImp)GetCountryKeyAndPhoneNumber(phone string)(key string,number string){
	nums:=strings.Split(phone," ")
	return nums[0],nums[1]
}
func (p phoneImp)GetCountryPhoneDetails(key string)(model.Phone,bool){
	v,existed :=model.AvailableCodes[key]
	if !existed{
		return model.Phone{},false
	}
	return v,true
}

func(p phoneImp)GetAvailablePhones()(phones []model.Phone){
	for _,v:=range model.AvailableCodes{
		phone := model.Phone{Country: v.Country,Code: v.Code}
		phones = append(phones,phone)
	}
	return phones
}