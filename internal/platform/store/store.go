package store

import (
	authModel "finback/internal/auth/model"
	contentModel "finback/internal/content/model"
	"sync"
)

type Store struct {
	Mu         sync.RWMutex
	UsersByID  map[string]*authModel.User
	IDByEmail  map[string]string
	NextUserID int
	Topics     map[string]contentModel.Topic
	Lessons    map[string]contentModel.Lesson
}

func New() *Store {
	return &Store{
		UsersByID:  map[string]*authModel.User{},
		IDByEmail:  map[string]string{},
		NextUserID: 1,
		Topics: map[string]contentModel.Topic{
			"budgeting":       {ID: "budgeting", Title: "Budgeting Basics", Description: "Plan income and expenses using simple monthly budgeting rules.", Difficulty: "beginner", Tags: []string{"cash flow", "planning"}},
			"saving":          {ID: "saving", Title: "Saving Habits", Description: "Build a steady saving routine and automate progress.", Difficulty: "beginner", Tags: []string{"saving", "goals"}},
			"fraud-security":  {ID: "fraud-security", Title: "Fraud and Security", Description: "Recognize scams, phishing, and safe digital behavior.", Difficulty: "beginner", Tags: []string{"fraud", "security"}},
			"investing":       {ID: "investing", Title: "Investing Basics", Description: "Understand risk, return, and diversification.", Difficulty: "intermediate", Tags: []string{"investing", "risk"}},
			"debt-management": {ID: "debt-management", Title: "Debt Management", Description: "Learn healthy borrowing and debt payoff strategies.", Difficulty: "intermediate", Tags: []string{"debt", "interest"}},
		},
		Lessons: map[string]contentModel.Lesson{
			"lesson-budget-1": {ID: "lesson-budget-1", TopicID: "budgeting", Title: "The 50/30/20 Idea", Summary: "A simple framework to split income into needs, wants, and savings.", EstimatedMinutes: 8, Content: []string{"List monthly income.", "Separate fixed and variable expenses.", "Set savings before wants."}, Quiz: []contentModel.QuizQuestion{{ID: "q1", Prompt: "Which category is usually planned before wants?", Options: []string{"Savings", "Entertainment", "Luxury shopping"}, Answer: "Savings"}}},
			"lesson-saving-1": {ID: "lesson-saving-1", TopicID: "saving", Title: "Automate Saving", Summary: "Use transfers and small goals to save consistently.", EstimatedMinutes: 7, Content: []string{"Create a target.", "Automate a small amount after payday.", "Track progress weekly."}, Quiz: []contentModel.QuizQuestion{{ID: "q1", Prompt: "What helps make saving consistent?", Options: []string{"Automation", "Guesswork", "Ignoring goals"}, Answer: "Automation"}}},
			"lesson-fraud-1":  {ID: "lesson-fraud-1", TopicID: "fraud-security", Title: "Spot a Phishing Message", Summary: "Recognize suspicious urgency and fake links.", EstimatedMinutes: 7, Content: []string{"Do not trust urgency alone.", "Check sender details.", "Never share one-time codes."}, Quiz: []contentModel.QuizQuestion{{ID: "q1", Prompt: "What is a common phishing sign?", Options: []string{"Urgent pressure", "Clear official process", "Verified office visit"}, Answer: "Urgent pressure"}}},
			"lesson-invest-1": {ID: "lesson-invest-1", TopicID: "investing", Title: "Risk and Diversification", Summary: "Different assets behave differently, which is why diversification matters.", EstimatedMinutes: 10, Content: []string{"Higher return usually means higher risk.", "Diversification spreads exposure.", "Long-term thinking matters."}, Quiz: []contentModel.QuizQuestion{{ID: "q1", Prompt: "Why diversify?", Options: []string{"Spread risk", "Guarantee huge returns", "Avoid learning"}, Answer: "Spread risk"}}},
			"lesson-debt-1":   {ID: "lesson-debt-1", TopicID: "debt-management", Title: "Interest and Repayment", Summary: "Learn why paying high-interest debt earlier matters.", EstimatedMinutes: 9, Content: []string{"List debts by interest rate.", "Avoid adding new debt during payoff.", "Use avalanche or snowball strategy."}, Quiz: []contentModel.QuizQuestion{{ID: "q1", Prompt: "Which debt is often prioritized in the avalanche method?", Options: []string{"Highest interest", "Lowest balance", "Newest debt"}, Answer: "Highest interest"}}},
		},
	}
}
