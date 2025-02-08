// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.5

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUsersIdPrefix       = "cache:users:id:"
	cacheUsersEmailPrefix    = "cache:users:email:"
	cacheUsersNicknamePrefix = "cache:users:nickname:"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Users, error)
		FindOneByEmail(ctx context.Context, email string) (*Users, error)
		FindOneByNickname(ctx context.Context, nickname sql.NullString) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		Id        int64          `db:"id"`         // snowflake ID
		Email     string         `db:"email"`      // user email
		Password  string         `db:"password"`   // argon2 hash
		Status    int64          `db:"status"`     // 0:unvarified, 1:varified
		Nickname  sql.NullString `db:"nickname"`   // user nickname
		Sex       int64          `db:"sex"`        // 0:unknown, 1:male, 2:female
		AvatarUrl sql.NullString `db:"avatar_url"` // avatar url
		Info      sql.NullString `db:"info"`       // user info
		CreatedAt time.Time      `db:"created_at"`
		UpdatedAt time.Time      `db:"updated_at"`
	}
)

func newUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`users`",
	}
}

func (m *defaultUsersModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	usersEmailKey := fmt.Sprintf("%s%v", cacheUsersEmailPrefix, data.Email)
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)
	usersNicknameKey := fmt.Sprintf("%s%v", cacheUsersNicknamePrefix, data.Nickname)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, usersEmailKey, usersIdKey, usersNicknameKey)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, usersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByEmail(ctx context.Context, email string) (*Users, error) {
	usersEmailKey := fmt.Sprintf("%s%v", cacheUsersEmailPrefix, email)
	var resp Users
	err := m.QueryRowIndexCtx(ctx, &resp, usersEmailKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", usersRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, email); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByNickname(ctx context.Context, nickname sql.NullString) (*Users, error) {
	usersNicknameKey := fmt.Sprintf("%s%v", cacheUsersNicknamePrefix, nickname)
	var resp Users
	err := m.QueryRowIndexCtx(ctx, &resp, usersNicknameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `nickname` = ? limit 1", usersRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, nickname); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (sql.Result, error) {
	usersEmailKey := fmt.Sprintf("%s%v", cacheUsersEmailPrefix, data.Email)
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersNicknameKey := fmt.Sprintf("%s%v", cacheUsersNicknamePrefix, data.Nickname)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, usersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Email, data.Password, data.Status, data.Nickname, data.Sex, data.AvatarUrl, data.Info)
	}, usersEmailKey, usersIdKey, usersNicknameKey)
	return ret, err
}

func (m *defaultUsersModel) Update(ctx context.Context, newData *Users) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	usersEmailKey := fmt.Sprintf("%s%v", cacheUsersEmailPrefix, data.Email)
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersNicknameKey := fmt.Sprintf("%s%v", cacheUsersNicknamePrefix, data.Nickname)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Email, newData.Password, newData.Status, newData.Nickname, newData.Sex, newData.AvatarUrl, newData.Info, newData.Id)
	}, usersEmailKey, usersIdKey, usersNicknameKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUsersIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}
