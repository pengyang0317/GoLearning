package handler

import (
	"context"
	"crypto/sha512"
	"lgo/pz-shop-server/user-srv/global"
	"strings"
	"time"

	"lgo/pz-shop-server/user-srv/model"
	userpb "lgo/pz-shop-server/user-srv/proto"

	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
}

func Paginate(req *userpb.GetUserRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if req.Page == 0 {
			req.Page = 1
		}
		if req.Size == 0 {
			req.Size = 10
		}
		offset := (req.Page - 1) * req.Size
		return db.Offset(int(offset)).Limit(int(req.Size))
	}
}

func (s *UserServer) GetUserList(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	// 获取用户列表
	var users []*userpb.User

	//TODO: 这里没有找到取分页总数的方法。 临时用查询两次的方案解决
	// 查询总数
	tx := global.DB.Find(&users)
	if tx.Error != nil { // 查询失败
		return nil, tx.Error
	}
	userRespone := &userpb.GetUserResponse{}
	userRespone.Totol = int32(tx.RowsAffected)

	// 分页
	tx = global.DB.Scopes(Paginate(req)).Find(&users)

	if tx.Error != nil { // 查询失败
		return nil, tx.Error
	}
	userRespone.Data = append(userRespone.Data, users...)

	return userRespone, nil
}

func (s *UserServer) GetUserByMobile(ctx context.Context, req *userpb.GetUserByMobileRequest) (*userpb.GetUserByMobileResponse, error) {
	user := &userpb.User{}
	tx := global.DB.Where("mobile = ?", req.Mobile).First(user)

	if tx.Error != nil { // 查询失败
		return nil, tx.Error
	}

	return &userpb.GetUserByMobileResponse{User: user}, nil

}

func (s *UserServer) GetUserById(ctx context.Context, req *userpb.GetUserByIdRequest) (*userpb.GetUserByIdResponse, error) {

	user := &userpb.User{}
	tx := global.DB.Where("id = ?", req.Id).First(user)
	if tx.Error != nil { // 查询失败
		return nil, tx.Error
	}

	return &userpb.GetUserByIdResponse{User: user}, nil
}

func ModelToCreateUserResponse(user *model.User) *userpb.CreateUserResponse {
	req := &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:       user.ID,
			PassWord: user.PassWord,
			NickName: user.NickName,
			Gender:   user.Gender,
			Role:     int32(user.Role),
			Mobile:   user.Mobile,
		},
	}
	if user.BirthDay != nil {
		req.User.BirthDay = uint64(user.BirthDay.Unix())
	}
	return req

}

func (s *UserServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {

	//因为使用了手机做唯一索引，所以这里需要先查询是否存在
	user := &model.User{}
	tx := global.DB.Where("mobile = ?", req.Mobile).First(user)
	if tx.RowsAffected > 0 { // 查询失败
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	user.Mobile = req.Mobile
	user.NickName = req.NickName
	user.PassWord = req.PassWord

	tx = global.DB.Create(user)
	if tx.Error != nil { // 查询失败
		return nil, tx.Error
	}
	results := ModelToCreateUserResponse(user)
	return results, nil

}
func ModelToUpdateUserResponse(user *model.User) *userpb.UpdateUserResponse {
	req := &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:       user.ID,
			PassWord: user.PassWord,
			NickName: user.NickName,
			Gender:   user.Gender,
			Role:     int32(user.Role),
			Mobile:   user.Mobile,
		},
	}
	if user.BirthDay != nil {
		req.User.BirthDay = uint64(user.BirthDay.Unix())
	}
	return req

}
func (s *UserServer) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	user := &model.User{}

	tx := global.DB.Where("id = ?", req.Id).First(user)
	if tx.Error != nil { // 查询失败
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	birthDay := time.Unix(int64(req.BirthDay), 0)

	user.NickName = req.NickName
	user.Gender = req.Gender
	user.BirthDay = &birthDay

	tx = global.DB.Save(user)

	if tx.Error != nil { // 查询失败
		return nil, status.Errorf(codes.Internal, tx.Error.Error())
	}

	return ModelToUpdateUserResponse(user), nil
}

func (s *UserServer) CheckPassword(ctx context.Context, req *userpb.CheckPasswordRequest) (*userpb.CheckPasswordResponse, error) {
	//校验密码
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	passwordInfo := strings.Split(req.EncryptedPassWord, "$")
	check := password.Verify(req.PassWord, passwordInfo[2], passwordInfo[3], options)
	return &userpb.CheckPasswordResponse{Success: check}, nil

}
