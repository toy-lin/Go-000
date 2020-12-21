package biz

// DO
type User struct {
    Username string
    Password string
}

type UserRepo interface {
    CheckExists(*User) (bool, error)
}

func NewUserUseCase(repo UserRepo) *UserUseCase {
    return &UserUseCase{repo: repo}
}

type UserUseCase struct {
    repo UserRepo
}

func (u UserUseCase) Exists(user *User) (bool, error) {
    return u.repo.CheckExists(user)
}
