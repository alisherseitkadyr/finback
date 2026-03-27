package repository

import (
	"errors"
	"finback/internal/auth/model"
	"finback/internal/platform/store"
	"fmt"
	"strings"
	"time"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository struct {
	store *store.Store
}

func NewUserRepository(store *store.Store) *UserRepository {
	return &UserRepository{store: store}
}

func (r *UserRepository) Register(name, email, level string, goals, preferred []string) (*model.User, error) {
	r.store.Mu.Lock()
	defer r.store.Mu.Unlock()

	email = strings.ToLower(strings.TrimSpace(email))
	if id, ok := r.store.IDByEmail[email]; ok {
		return r.store.UsersByID[id], nil
	}

	user := &model.User{
		ID:               fmt.Sprintf("user-%03d", r.store.NextUserID),
		Name:             strings.TrimSpace(name),
		Email:            email,
		ExperienceLevel:  level,
		Goals:            cloneStrings(goals),
		PreferredTopics:  cloneStrings(preferred),
		Assessment:       map[string]float64{},
		CompletedLessons: []string{},
		Progress:         map[string]float64{},
		CreatedAt:        time.Now().UTC(),
	}
	r.store.NextUserID++
	r.store.UsersByID[user.ID] = user
	r.store.IDByEmail[email] = user.ID
	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	r.store.Mu.RLock()
	defer r.store.Mu.RUnlock()
	id, ok := r.store.IDByEmail[strings.ToLower(strings.TrimSpace(email))]
	if !ok {
		return nil, ErrUserNotFound
	}
	return r.store.UsersByID[id], nil
}

func (r *UserRepository) GetByID(userID string) (*model.User, error) {
	r.store.Mu.RLock()
	defer r.store.Mu.RUnlock()
	user, ok := r.store.UsersByID[userID]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (r *UserRepository) Save(user *model.User) error {
	r.store.Mu.Lock()
	defer r.store.Mu.Unlock()
	r.store.UsersByID[user.ID] = user
	return nil
}

func cloneStrings(v []string) []string {
	out := make([]string, len(v))
	copy(out, v)
	return out
}
