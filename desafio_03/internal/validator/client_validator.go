package validator

import (
	"errors"
	"regexp"
	"time"
)

var (
	ErrNameIsEmpty       = errors.New("nome não pode ser vazio")
	ErrBirthdayIsFuture = errors.New("data de nasicmento não pode ser futura")
	ErrIncomeIsInvalid   = errors.New("salário não pode ser negativo")
	ErrCpfIsInvalid      = errors.New("CPF deve conter 11 dígitos numéricos")
	ErrChildrenIsInvalid = errors.New("quantidade de filhos não pode ser negativa")
)

func ValidateClientName(name string) error {
	if len(name) == 0 {
		return ErrNameIsEmpty
	}
	return nil
}

func ValidateClientBirthDate(birthDate time.Time) error {
	if birthDate.After(time.Now()) {
		return ErrBirthdayIsFuture
	}
	return nil
}

func ValidateIncome(income float64) error {
	if income < 0 {
		return ErrIncomeIsInvalid
	}
	return nil
}

func ValidateCPF(cpf string) error {
	match, _ := regexp.MatchString(`^\d{11}$`, cpf)
	if !match {
		return ErrCpfIsInvalid
	}
	return nil
}

func ValidateChildren(children int) error {
	if children < 0 {
		return ErrChildrenIsInvalid
	}
	return nil
}
