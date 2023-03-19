package repository

//berhubungan langsung dengan database
import (
	"clean/entity"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	BuatUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
	GetAllUser(ctx context.Context,tx *gorm.DB) ([]entity.User, error)
	UpdateNama(ctx context.Context, tx *gorm.DB, namaLama string, namaBaru string) (entity.User, error)
	DeleteUser(ctx context.Context, tx *gorm.DB, ID uint64) (entity.User, error)
	GetUserByID(ctx context.Context, tx *gorm.DB, userID uint64) (entity.User, error)
	GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, error)
	GetUserByNama(ctx context.Context, tx *gorm.DB, nama string) (entity.User, error)
}


func NewUserRepository(dbv *gorm.DB) UserRepository {
	return &userRepository{
		db: dbv,
	}
}
func (r *userRepository) BuatUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error) {
	var err error
	if tx == nil {
		tx = r.db.WithContext(ctx).Create(&user)
		err = tx.Error
	} else {
		err = r.db.WithContext(ctx).Create(&user).Error
	}
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetAllUser(ctx context.Context, tx *gorm.DB) ([]entity.User, error) {
	var daftarUser []entity.User
	var err error
	if tx == nil {
		tx = r.db.WithContext(ctx).Preload("Blogs").Preload("Likes").Preload("Comments").Find(&daftarUser)
		err = tx.Error
	} else {
		err = r.db.WithContext(ctx).Preload("Blogs").Preload("Likes").Preload("Comments").Find(&daftarUser).Error
	}
	if err != nil {
		return []entity.User{}, err
	}
	return daftarUser, nil
}
func (r *userRepository) UpdateNama(ctx context.Context, tx *gorm.DB, namaLama string, namaBaru string) (entity.User, error) {
	var err error

	if tx == nil {
		tx = r.db.WithContext(ctx).Model(&entity.User{}).Where("nama = ?", namaLama).Update("nama", namaBaru)
	} else {
		err = r.db.WithContext(ctx).Model(&entity.User{}).Where("nama = ?", namaLama).Update("nama", namaBaru).Error
	}
	if err != nil {
		return entity.User{}, err
	}
	var user entity.User
	t := r.db.WithContext(ctx).Debug().Where("nama = ?", namaBaru).Take(&user)
	if t.Error != nil {
		return entity.User{}, t.Error
	}

	return user, nil
}
func (r *userRepository) DeleteUser(ctx context.Context, tx *gorm.DB, ID uint64) (entity.User, error) {
	var err error
	var user entity.User
	if tx == nil {
		tx = r.db.WithContext(ctx).Debug().Where("id = ?", ID).Preload("Blogs").Preload("Likes").Preload("Comments").Take(&user)
	} else {
		err = r.db.WithContext(ctx).Debug().Where("id = ?", ID).Preload("Blogs").Preload("Likes").Preload("Comments").Take(&user).Error
	}
	if err != nil {
		return entity.User{}, nil
	}
	tx = r.db.WithContext(ctx).Debug().Where("id = ?", ID).Delete(&entity.User{})
	if tx.Error != nil {
		return entity.User{}, nil
	}
	return user, nil
}
func (r *userRepository) GetUserByID(ctx context.Context, tx *gorm.DB, userID uint64) (entity.User, error) {
	var user entity.User
	var err error
	if tx == nil {
		tx = r.db.WithContext(ctx).Debug().Where("id = ?", userID).Preload("Blogs").Preload("Likes").Preload("Comments").Take(&user)
		err = tx.Error
	} else {
		err = r.db.WithContext(ctx).Debug().Where("id = ?", userID).Preload("Blogs").Preload("Likes").Preload("Comments").Take(&user).Error
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, fmt.Errorf("user with ID %d not found", userID)
		}
		return entity.User{}, err
	}
	return user, nil
}
func (r *userRepository) GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, error) {
	var user entity.User
	var err error
	if tx == nil {
		tx = r.db.WithContext(ctx).Debug().Where("email = ?", email).Preload("Blogs").Preload("Likes").Preload("Comments").Take(&user)
	} else {
		err = r.db.WithContext(ctx).Debug().Where("email = ?", email).Preload("Blogs").Preload("Likes").Preload("Comments").Take(&user).Error
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, fmt.Errorf("user with email %s not found", email)
		}
		return entity.User{}, err
	}
	return user, nil
}
func (r *userRepository) GetUserByNama(ctx context.Context, tx *gorm.DB, nama string) (entity.User, error) {
	var user entity.User
	var err error
	if tx == nil {
		tx = r.db.WithContext(ctx).Debug().Where("nama = ?", nama).Preload("Blogs").Preload("Likes").Preload("Comments").Take(&user)
		err = tx.Error
	} else {
		err = r.db.WithContext(ctx).Debug().Where("nama = ?", nama).Preload("Blogs").Preload("Likes").Preload("Comments").Take(&user).Error
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, fmt.Errorf("user with email %s not found", nama)
		}
		return entity.User{}, err
	}
	return user, nil
}
