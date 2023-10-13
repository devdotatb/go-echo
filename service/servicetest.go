package service

type serviceTest struct {
	hellotext string
	userlist  []User
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewServiceTest(text string) *serviceTest {
	serviceTest := serviceTest{
		hellotext: text,
	}
	return &serviceTest
}

func (s *serviceTest) GetHello() string {
	return s.hellotext
}

func (s *serviceTest) UpdateHello(text string) {
	s.hellotext = text
}

func (s *serviceTest) GetAllUser() []User {

}
func (s *serviceTest) (u User){

}
