package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/huynhtrongtien/dove/entities"
	"github.com/huynhtrongtien/dove/models"
	"github.com/huynhtrongtien/dove/pkg/crypto"
)

type IUser interface {
	Create(ctx context.Context, data *entities.User) (string, error)
	Read(ctx context.Context, id int64) (*entities.User, error)
	ReadByUsername(ctx context.Context, username string) (*entities.User, error)
	Update(ctx context.Context, data *entities.User) error
	Authenticate(ctx context.Context, username, password string) (int64, string, error)
}

type User struct {
	Model models.IUser
}

func NewUser() IUser {
	return &User{
		Model: models.User{},
	}
}

func (p *User) Create(ctx context.Context, data *entities.User) (string, error) {
	data.UUID = uuid.NewString()
	userId, err := p.Model.Create(ctx, data)
	if err != nil {
		return "", err
	}

	return JWTMaker.CreateToken(userId)
}

func (p *User) Read(ctx context.Context, id int64) (*entities.User, error) {
	return p.Model.First(ctx, map[string]any{"id": id})
}

func (p *User) ReadByUsername(ctx context.Context, username string) (*entities.User, error) {
	return p.Model.First(ctx, map[string]any{"username": username})
}

func (p *User) Update(ctx context.Context, data *entities.User) error {
	return p.Model.Update(ctx, data)
}

func (p *User) Authenticate(ctx context.Context, username, password string) (int64, string, error) {
	data, err := p.Model.First(ctx, map[string]any{"username": username})
	if err != nil {
		return 0, "", err
	}

	if crypto.CheckPasswordHash(password, data.PasswordHash) {
		return 0, "", fmt.Errorf("password is incorrect")
	}

	jwtToken, err := JWTMaker.CreateToken(data.ID)
	if err != nil {
		return 0, "", err
	}
	return data.ID, jwtToken, nil
}
