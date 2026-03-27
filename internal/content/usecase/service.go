package usecase

import (
	authModel "finback/internal/auth/model"
	authRepo "finback/internal/auth/repository"
	contentModel "finback/internal/content/model"
	contentRepo "finback/internal/content/repository"
	"fmt"
	"sort"
	"strings"
)

type TopicItem struct {
	Topic   contentModel.Topic    `json:"topic"`
	Lessons []contentModel.Lesson `json:"lessons"`
}

type AssessmentAnswer struct {
	TopicID string  `json:"topic_id"`
	Score   float64 `json:"score"`
}

type SubmitAssessmentRequest struct {
	UserID  string             `json:"user_id"`
	Answers []AssessmentAnswer `json:"answers"`
}

type CompleteLessonRequest struct {
	UserID           string  `json:"user_id"`
	LessonID         string  `json:"lesson_id"`
	QuizScore        float64 `json:"quiz_score"`
	TimeSpentMinutes int     `json:"time_spent_minutes"`
}

type Recommendation struct {
	TopicID  string  `json:"topic_id"`
	Title    string  `json:"title"`
	Reason   string  `json:"reason"`
	Score    float64 `json:"score"`
	Priority string  `json:"priority"`
}

type Service struct {
	contentRepo *contentRepo.Repository
	userRepo    *authRepo.UserRepository
}

func NewService(contentRepo *contentRepo.Repository, userRepo *authRepo.UserRepository) *Service {
	return &Service{contentRepo: contentRepo, userRepo: userRepo}
}

func (s *Service) ListTopics() []TopicItem {
	topics := s.contentRepo.ListTopics()
	items := make([]TopicItem, 0, len(topics))
	for _, topic := range topics {
		items = append(items, TopicItem{
			Topic:   topic,
			Lessons: s.contentRepo.LessonsByTopic(topic.ID),
		})
	}
	return items
}

func (s *Service) GetTopic(topicID string) (*TopicItem, error) {
	topic, err := s.contentRepo.GetTopic(topicID)
	if err != nil {
		return nil, err
	}
	return &TopicItem{Topic: *topic, Lessons: s.contentRepo.LessonsByTopic(topicID)}, nil
}

func (s *Service) GetLesson(lessonID string) (*contentModel.Lesson, error) {
	return s.contentRepo.GetLesson(lessonID)
}

func (s *Service) SubmitAssessment(req SubmitAssessmentRequest) (map[string]any, error) {
	user, err := s.userRepo.GetByID(req.UserID)
	if err != nil {
		return nil, err
	}
	for _, a := range req.Answers {
		if _, err := s.contentRepo.GetTopic(a.TopicID); err != nil {
			return nil, fmt.Errorf("unknown topic_id: %s", a.TopicID)
		}
		if a.Score < 0 {
			a.Score = 0
		}
		if a.Score > 100 {
			a.Score = 100
		}
		user.Assessment[a.TopicID] = a.Score
	}
	if err := s.userRepo.Save(user); err != nil {
		return nil, err
	}
	return map[string]any{
		"assessment":      user.Assessment,
		"recommendations": buildRecommendations(s.contentRepo.ListTopics(), user),
	}, nil
}

func (s *Service) CompleteLesson(req CompleteLessonRequest) (map[string]any, error) {
	user, err := s.userRepo.GetByID(req.UserID)
	if err != nil {
		return nil, err
	}
	lesson, err := s.contentRepo.GetLesson(req.LessonID)
	if err != nil {
		return nil, err
	}
	if !contains(user.CompletedLessons, lesson.ID) {
		user.CompletedLessons = append(user.CompletedLessons, lesson.ID)
	}
	if req.QuizScore < 0 {
		req.QuizScore = 0
	}
	if req.QuizScore > 100 {
		req.QuizScore = 100
	}
	current := user.Progress[lesson.TopicID]
	updated := current + (req.QuizScore-current)*0.6
	if updated == 0 {
		updated = req.QuizScore
	}
	user.Progress[lesson.TopicID] = round1(updated)
	user.TotalPoints += int(req.QuizScore)
	if err := s.userRepo.Save(user); err != nil {
		return nil, err
	}
	return map[string]any{
		"completed_lesson":   lesson,
		"topic_progress":     user.Progress[lesson.TopicID],
		"total_points":       user.TotalPoints,
		"time_spent_minutes": req.TimeSpentMinutes,
	}, nil
}

func (s *Service) GetProgress(userID string) (map[string]any, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"user_id":           user.ID,
		"completed_lessons": user.CompletedLessons,
		"topic_progress":    user.Progress,
		"assessment":        user.Assessment,
		"total_points":      user.TotalPoints,
	}, nil
}

func (s *Service) GetRecommendations(userID string) (map[string]any, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return map[string]any{"items": buildRecommendations(s.contentRepo.ListTopics(), user)}, nil
}

func buildRecommendations(topics []contentModel.Topic, user *authModel.User) []Recommendation {
	items := make([]Recommendation, 0, len(topics))
	for _, topic := range topics {
		score := 40.0
		reasons := []string{}

		if contains(user.PreferredTopics, topic.ID) {
			score += 35
			reasons = append(reasons, "matches preferred topic from onboarding")
		}
		if a, ok := user.Assessment[topic.ID]; ok {
			score += (100 - a) * 0.5
			if a < 60 {
				reasons = append(reasons, fmt.Sprintf("assessment score is low (%.0f%%)", a))
			}
		}
		if p, ok := user.Progress[topic.ID]; ok {
			score += (100 - p) * 0.2
			if p >= 80 {
				score -= 20
				reasons = append(reasons, "already strong progress in this topic")
			}
		}
		for _, goal := range user.Goals {
			g := strings.ToLower(goal)
			switch {
			case strings.Contains(g, "save") && (topic.ID == "saving" || topic.ID == "budgeting"):
				score += 12
			case strings.Contains(g, "debt") && topic.ID == "debt-management":
				score += 14
			case strings.Contains(g, "invest") && topic.ID == "investing":
				score += 14
			case (strings.Contains(g, "security") || strings.Contains(g, "fraud")) && topic.ID == "fraud-security":
				score += 14
			}
		}
		priority := "medium"
		if score >= 95 {
			priority = "high"
		} else if score < 65 {
			priority = "low"
		}
		reason := "general foundation topic"
		if len(reasons) > 0 {
			reason = strings.Join(uniqueStrings(reasons), "; ")
		}
		items = append(items, Recommendation{
			TopicID:  topic.ID,
			Title:    topic.Title,
			Reason:   reason,
			Score:    round1(score),
			Priority: priority,
		})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Score == items[j].Score {
			return items[i].Title < items[j].Title
		}
		return items[i].Score > items[j].Score
	})
	if len(items) > 5 {
		items = items[:5]
	}
	return items
}

func contains(items []string, target string) bool {
	for _, item := range items {
		if strings.EqualFold(item, target) {
			return true
		}
	}
	return false
}

func uniqueStrings(items []string) []string {
	seen := map[string]struct{}{}
	out := []string{}
	for _, item := range items {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		out = append(out, item)
	}
	return out
}

func round1(v float64) float64 {
	return float64(int(v*10+0.5)) / 10
}
