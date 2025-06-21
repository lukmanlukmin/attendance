// Package bootstrap ...
package bootstrap

import (
	"attendance/constant"
	"attendance/model/db"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/lukmanlukmin/go-lib/database"
	"github.com/lukmanlukmin/go-lib/log"
	"github.com/lukmanlukmin/go-lib/util"
)

// LoadDefaultData ...
func LoadDefaultData(bs *Bootstrap) {
	log.Info("making sure system has default data")
	LoadDefaultUser(context.TODO(), bs)
	log.Info("done making sure has default data")
}

// DefaultUser ...
type DefaultUser struct {
	Username string
	Password string
	FullName string
	Role     string
}

// LoadDefaultUser ...
func LoadDefaultUser(ctx context.Context, bs *Bootstrap) {
	// Role
	requiredRole := []string{constant.RoleAdmin, constant.RoleEmployee}
	roles, err := bs.Repository.DB.Role.GetAll(ctx)
	if err != nil {
		log.WithError(err).Fatal("failed to prepare default data - 0")
	}
	if len(roles) < len(requiredRole) {
		log.Info("creating default role")
		err = database.BeginTransaction(ctx, bs.Repository.Store.GetMaster(), func(ctx context.Context) error {
			for _, role := range requiredRole {
				if err := bs.Repository.DB.Role.Create(ctx, &db.Role{
					Name: role,
				}); err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			log.WithError(err).Fatal("failed to prepare default data - 1")
		}
		roles, err = bs.Repository.DB.Role.GetAll(ctx)
		if err != nil {
			log.WithError(err).Fatal("failed to prepare default data - 2")
		}
	}

	users := []DefaultUser{}
	for i := 0; i <= 100; i++ {
		if i == 0 {
			hashPasswd, _ := util.HashPassword(constant.RoleAdmin)
			users = append(users, DefaultUser{
				Username: constant.RoleAdmin,
				Password: hashPasswd,
				Role:     constant.RoleAdmin,
				FullName: constant.RoleAdmin,
			})
			continue
		}
		hashPasswd, _ := util.HashPassword(fmt.Sprintf("password%d", i))
		users = append(users, DefaultUser{
			Username: fmt.Sprintf("%s%d", constant.RoleEmployee, i),
			Password: hashPasswd,
			Role:     constant.RoleEmployee,
			FullName: fmt.Sprintf("%s %d", constant.RoleEmployee, i),
		})
	}

	// check is already have user data
	userSample, err := bs.Repository.DB.User.GetByUsername(ctx, users[0].Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.WithError(err).Fatal("failed to prepare default data - 3")
	}
	if userSample != nil {
		return
	}

	employeeRole := db.Role{}
	for _, role := range roles {
		if role.Name == constant.RoleEmployee {
			employeeRole = role
			break
		}
	}
	for _, user := range users {
		dataUser, err := bs.Repository.DB.User.GetByUsername(ctx, user.Username)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			log.WithError(err).Fatal("failed to prepare default data - 4", user.Username)
		}
		if dataUser != nil {
			continue
		}
		// admin
		if user.Role == constant.RoleAdmin {
			log.Info("creating default admin")
			err = database.BeginTransaction(ctx, bs.Repository.Store.GetMaster(), func(ctx context.Context) error {
				userAdmin := &db.User{
					Username: user.Username,
					Password: user.Password,
				}
				err = bs.Repository.DB.User.Create(ctx, userAdmin)
				if err != nil {
					return err
				}
				for _, role := range roles {
					err = bs.Repository.DB.UserRole.Create(ctx, &db.UserRole{
						UserID: userAdmin.ID,
						RoleID: role.ID,
					})
					if err != nil {
						return err
					}
					if role.Name == constant.RoleEmployee {
						err = bs.Repository.DB.Employee.Create(ctx, &db.Employee{
							UserID:   userAdmin.ID,
							Salary:   0,
							FullName: user.FullName,
						})
						if err != nil {
							return err
						}
					}
				}
				return nil
			})
			if err != nil {
				log.WithError(err).Fatal("failed to prepare default data - 5")
			}
			continue
		}

		// employee
		if user.Role == constant.RoleEmployee {
			log.Info("creating default employee")
			err = database.BeginTransaction(ctx, bs.Repository.Store.GetMaster(), func(ctx context.Context) error {
				userEmployee := &db.User{
					Username: user.Username,
					Password: user.Password,
				}
				err = bs.Repository.DB.User.Create(ctx, userEmployee)
				if err != nil {
					return err
				}
				err = bs.Repository.DB.UserRole.Create(ctx, &db.UserRole{
					UserID: userEmployee.ID,
					RoleID: employeeRole.ID,
				})
				if err != nil {
					return err
				}

				rand.Seed(time.Now().UnixNano())
				randomNumber := rand.Intn(9) + 1
				err = bs.Repository.DB.Employee.Create(ctx, &db.Employee{
					UserID:   userEmployee.ID,
					Salary:   randomNumber * 1000000,
					FullName: user.FullName,
				})
				if err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				log.WithError(err).Fatal("failed to prepare default data - 6")
			}
		}

	}
}
