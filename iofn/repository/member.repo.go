package repository

type MemberRepository struct {}

func (u MemberRepository) FindNameById(id int64) string {
	member := mockMember()
	return member.name
}

func mockMember() member {
	return member{
		id:       1,
		name:     "teddy",
		password: "password",
	}
}
