package repository

import (
	"errors"
	contentModel "finback/internal/content/model"
	"finback/internal/platform/store"
	"sort"
)

var (
	ErrTopicNotFound  = errors.New("topic not found")
	ErrLessonNotFound = errors.New("lesson not found")
)

type Repository struct {
	store *store.Store
}

func New(store *store.Store) *Repository {
	return &Repository{store: store}
}

func (r *Repository) ListTopics() []contentModel.Topic {
	r.store.Mu.RLock()
	defer r.store.Mu.RUnlock()
	items := make([]contentModel.Topic, 0, len(r.store.Topics))
	for _, t := range r.store.Topics {
		items = append(items, t)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].Title < items[j].Title })
	return items
}

func (r *Repository) GetTopic(topicID string) (*contentModel.Topic, error) {
	r.store.Mu.RLock()
	defer r.store.Mu.RUnlock()
	t, ok := r.store.Topics[topicID]
	if !ok {
		return nil, ErrTopicNotFound
	}
	return &t, nil
}

func (r *Repository) GetLesson(lessonID string) (*contentModel.Lesson, error) {
	r.store.Mu.RLock()
	defer r.store.Mu.RUnlock()
	l, ok := r.store.Lessons[lessonID]
	if !ok {
		return nil, ErrLessonNotFound
	}
	return &l, nil
}

func (r *Repository) LessonsByTopic(topicID string) []contentModel.Lesson {
	r.store.Mu.RLock()
	defer r.store.Mu.RUnlock()
	items := []contentModel.Lesson{}
	for _, l := range r.store.Lessons {
		if l.TopicID == topicID {
			items = append(items, l)
		}
	}
	sort.Slice(items, func(i, j int) bool { return items[i].Title < items[j].Title })
	return items
}
