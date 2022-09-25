package service

import (
	"fmt"
	"tdd_config/iofn/repository"
)

type UserService struct {
	repository.MemberRepository
}

func (service UserService) ReadDetailAll()  {
	memberName := service.FindNameById(int64(1))
	fmt.Println(memberName)
}