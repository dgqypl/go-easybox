package repository

import "context"

type IdType interface {
	int64 | string
}

type Field struct {
	Key   string
	Value any
}

// IRepository 定义了业务repository一些通用方法的接口
type IRepository[T any, I IdType] interface {
	Insert(ctx context.Context, t T) (I, error)
	UpdateById(ctx context.Context, id I, t T) error
	UpdateFieldsById(ctx context.Context, id I, fields []*Field) error
	GetById(ctx context.Context, id I) (T, error)
	GetByIds(ctx context.Context, ids []I) ([]T, error)
	DeleteById(ctx context.Context, id I) error
	Count(ctx context.Context) (int64, error)
}

// DefaultRepository 作为 IRepository 的默认实现，业务repository将其作为一个匿名属性，并覆写需要的方法。代码样例如下：
//
// 	type User struct {
// 		Id   string
// 		Name string
// 	}
//
// 	type UserRepository interface {
// 		IRepository[*User, string]
// 	}
//
// 	type userRepository struct {
// 		DefaultRepository[*User, string]
// 	}
//
// 	func NewUserRepository() UserRepository {
// 		return &userRepository{}
// 	}
//
// 	func (repo *userRepository) Insert(ctx context.Context, u *User) (string, error) {
// 		panic("insert user")
// 	}
type DefaultRepository[T any, I IdType] struct {
	IRepository[T, I]
}

func (repo *DefaultRepository[T, I]) Insert(ctx context.Context, t T) (I, error) {
	panic("Insert")
}

func (repo *DefaultRepository[T, I]) UpdateById(ctx context.Context, id I, t T) error {
	panic("UpdateById")
}

func (repo *DefaultRepository[T, I]) UpdateFieldsById(ctx context.Context, id I, fields []*Field) error {
	panic("UpdateFieldsById")
}
func (repo *DefaultRepository[T, I]) GetById(ctx context.Context, id I) (T, error) {
	panic("GetById")
}

func (repo *DefaultRepository[T, I]) GetByIds(ctx context.Context, ids []I) ([]T, error) {
	panic("GetByIds")
}

func (repo *DefaultRepository[T, I]) DeleteById(ctx context.Context, id I) error {
	panic("DeleteById")
}

func (repo *DefaultRepository[T, I]) Count(ctx context.Context) (int64, error) {
	panic("Count")
}

type Entity[I IdType] interface {
	GetEntityId() I
}

// SortEntitiesByIds 对业务repository的GetByIds方法提供入参 []I 与返回值 []T 元素顺序保持一致的功能
func SortEntitiesByIds[E Entity[I], I IdType](ids []I, entities []E) []E {
	entityMap := make(map[I]E)
	for _, entity := range entities {
		entityMap[entity.GetEntityId()] = entity
	}

	result := make([]E, len(entities))
	for i, v := range ids {
		result[i] = entityMap[v]
	}

	return result
}
