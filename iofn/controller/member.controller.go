package controller

// request mapping url + method 는 겹칠 수 없으니
// controller 는 같은 repo 에서 사용한다.
// openapi && api test

// GET
// 회원정보 상세보기 : admin
// 회원정보 상세보기 : user
// 회원 리스트 조회 : admin
// 권한 리스트 조회 : any

// POST
// 회원 등록 : admin
// 회원 가입 신청 : any

// PUT
// 회원 권한 변경 : admin
// 회원 비밀번호 변경 : user
// 회원 개인정보 변경 : user
// 회원 권한 변경 : admin
// 회원 정지 : admin
// 회원 삭제 : admin

// GET / AUTH
// 회원 로그인 : any
// 회원 로그아웃 : any
