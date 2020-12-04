package main

import "fmt"

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{code: code}
}

func (s *securityCode) checkCode(intcomingCode int) error {
	if s.code != intcomingCode {
		return fmt.Errorf("security code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}