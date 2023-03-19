package service

import (
	"clean/dto"
	"clean/entity"
	"clean/repository"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

//disini merupakan lapisan yang mengurus binis logic
//service dependant ke repository
type userService struct{
	userRepo repository.UserRepository
}
type UserService interface{
	Register(ctx context.Context, userDTO dto.RegisterUser) (entity.User, error)
	Login(ctx context.Context, userDTO dto.LoginUser)(entity.User,error)
	GetAllUser(ctx context.Context)([]entity.User,error)
	GetUserInfo(ctx context.Context, id uint64)(entity.User,error)
	UpdateNama(ctx context.Context,namaLama string, nama string)(entity.User,error)
	DeleteAkun(ctx context.Context, ID uint64)(entity.User,error)
}
func NewUserService(ur repository.UserRepository) UserService{
	return &userService{
		userRepo: ur,
	}
}
func (s *userService) Register(ctx context.Context, userDTO dto.RegisterUser) (entity.User,error){
	var user entity.User
	user.ID=userDTO.ID
	user.Nama=userDTO.Nama
	user.Email=userDTO.Email
	user.Password=userDTO.Password	
	user.Role = userDTO.Role

	userWithEmail, err :=s.userRepo.GetUserByEmail(ctx, nil, userDTO.Email)

	if err != nil {
		return entity.User{},err
	}
	if !(userWithEmail.Email == "") {
		return entity.User{}, errors.New("duplicate user")
	}
	
	hashPassword,e := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if e != nil {
		return entity.User{},err
	}
	user.Password=string(hashPassword)
	createdUser, err :=s.userRepo.BuatUser(ctx,nil,user)
	if err != nil {
		return entity.User{},err
	}
	
	return createdUser,nil
}
func (s *userService) Login(ctx context.Context, userDTO dto.LoginUser) (entity.User,error) {
	userWithEmail, err :=s.userRepo.GetUserByEmail(ctx, nil, userDTO.Email)

	if err != nil {
		return entity.User{},err
	}
	if userWithEmail.Email == "" {
		return entity.User{}, errors.New("user tidak ditemukan")
	}

	er := bcrypt.CompareHashAndPassword([]byte(userWithEmail.Password),[]byte(userDTO.Password))
	if er != nil{
		return entity.User{},er
	}
	return userWithEmail,nil
}
func (s *userService) GetUserInfo(ctx context.Context, id uint64)(entity.User,error){
	user,err := s.userRepo.GetUserByID(ctx,nil,id)
	if err !=nil {
		return entity.User{},err
	}
	return user,nil
}
func (s *userService) UpdateNama(ctx context.Context,namaLama string, namaBaru string)(entity.User, error){
	user,err := s.userRepo.UpdateNama(ctx,nil,namaLama,namaBaru)
	if err !=nil {
		return entity.User{},err
	}
	return user,nil
}
func (s *userService) DeleteAkun(ctx context.Context,ID uint64)(entity.User, error){
	user,err := s.userRepo.DeleteUser(ctx,nil,ID)
	if err !=nil {
		return entity.User{},err
	}
	return user,nil
}
func (s *userService) GetAllUser(ctx context.Context)([]entity.User,error){
	users,err :=s.userRepo.GetAllUser(ctx,nil)
	if err != nil{
		return []entity.User{},err
	}
	return users,nil
}
