package data

type Service interface {
	FindAll() ([]Nama, error)
	FindById(ID int) (Nama, error)
	Create(data User) (Nama, error)
	Update(ID int, data User) (Nama, error)
	Delete(ID int) (Nama, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Nama, error) {
	namas, err := s.repository.FindAll()
	return namas, err
}

func (s *service) FindById(ID int) (Nama, error) {
	namas, err := s.repository.FindById(ID)
	return namas, err
}

func (s *service) Create(inputRequest User) (Nama, error) {

	age, _ := inputRequest.Age.Int64()
	rating, _ := inputRequest.Rating.Int64()

	user := Nama{
		Name:        inputRequest.Name,
		Email:       inputRequest.Email,
		Age:         int(age),
		Rating:      int(rating),
		Description: inputRequest.Description,
	}

	newRequest, err := s.repository.Create(user)
	return newRequest, err

}

func (s *service) Update(ID int, inputRequest User) (Nama, error) {
	nama, err := s.repository.FindById(ID)

	age, _ := inputRequest.Age.Int64()
	rating, _ := inputRequest.Rating.Int64()

	nama.Name = inputRequest.Name
	nama.Email = inputRequest.Email
	nama.Age = int(age)
	nama.Rating = int(rating)
	nama.Description = inputRequest.Description

	newRequest, err := s.repository.Update(nama)
	return newRequest, err

}

func (s *service) Delete(ID int) (Nama, error) {
	hapusnama, err := s.repository.FindById(ID)
	newRequest, err := s.repository.Delete(hapusnama)
	return newRequest, err

}
