package store

import (
	authModel "finback/internal/auth/model"
	contentModel "finback/internal/content/model"
	"sync"
	"time"
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
		UsersByID: map[string]*authModel.User{
			// Seed a demo current_user so frontend can request progress for 'current_user'
			"current_user": &authModel.User{
				ID:               "current_user",
				Name:             "Demo User",
				Email:            "demo@example.com",
				ExperienceLevel:  "beginner",
				Goals:            []string{},
				PreferredTopics:  []string{},
				Assessment:       map[string]float64{},
				CompletedLessons: []string{"lesson-budget-1"},
				Progress:         map[string]float64{"budgeting": 85},
				TotalPoints:      85,
				CreatedAt:        time.Now().UTC(),
			},
		},
		IDByEmail:  map[string]string{"demo@example.com": "current_user"},
		NextUserID: 2,
		Topics: map[string]contentModel.Topic{
			"budgeting":       {ID: "budgeting", Title: "Budgeting Basics", Description: "Plan income and expenses using simple monthly budgeting rules.", Difficulty: "beginner", Tags: []string{"cash flow", "planning"}},
			"saving":          {ID: "saving", Title: "Saving Habits", Description: "Build a steady saving routine and automate progress.", Difficulty: "beginner", Tags: []string{"saving", "goals"}},
			"fraud-security":  {ID: "fraud-security", Title: "Fraud and Security", Description: "Recognize scams, phishing, and safe digital behavior.", Difficulty: "beginner", Tags: []string{"fraud", "security"}},
			"investing":       {ID: "investing", Title: "Investing Basics", Description: "Understand risk, return, and diversification.", Difficulty: "intermediate", Tags: []string{"investing", "risk"}},
			"debt-management": {ID: "debt-management", Title: "Debt Management", Description: "Learn healthy borrowing and debt payoff strategies.", Difficulty: "intermediate", Tags: []string{"debt", "interest"}},
		},
		Lessons: map[string]contentModel.Lesson{
			"lesson-budget-1": {
				ID: "lesson-budget-1", TopicID: "budgeting", Title: "The 50/30/20 Idea",
				Summary:          "A simple framework to split income into needs, wants, and savings.",
				EstimatedMinutes: 8,
				Content:          []string{"List monthly income.", "Separate fixed and variable expenses.", "Set savings before wants."},
				Quiz:             []contentModel.QuizQuestion{{ID: "q1", Prompt: "Which category is usually planned before wants?", Options: []string{"Savings", "Entertainment", "Luxury shopping"}, Answer: "Savings"}},
				Steps: []contentModel.LessonStep{
					{ID: "b1", Order: 1, Title: "What Is a Budget?", Body: "A budget is a plan for your money. It shows how much comes in (income) and how much goes out (expenses) each month.", Example: "Ana earns $2,500/month. Her rent is $900, food $400, transport $150.", Tip: "Even a rough budget beats no budget at all."},
					{ID: "b2", Order: 2, Title: "The 50/30/20 Rule", Body: "50% needs, 30% wants, 20% savings and debt repayment.", Example: "$2,500 income → $1,250 for needs, $750 for wants, $500 for savings.", Tip: "This rule is a starting point, not a strict law."},
				},
				Outcomes: []contentModel.LessonOutcome{
					{Text: "Understand the basic structure of budgeting"},
					{Text: "Know the 50/30/20 rule"},
					{Text: "Be able to create a simple monthly budget"},
				},
			},
			"lesson-saving-1": {
				ID: "lesson-saving-1", TopicID: "saving", Title: "Automate Saving",
				Summary:          "Use transfers and small goals to save consistently.",
				EstimatedMinutes: 7,
				Content:          []string{"Create a target.", "Automate a small amount after payday.", "Track progress weekly."},
				Quiz:             []contentModel.QuizQuestion{{ID: "q1", Prompt: "What helps make saving consistent?", Options: []string{"Automation", "Guesswork", "Ignoring goals"}, Answer: "Automation"}},
				Steps: []contentModel.LessonStep{
					{ID: "s1", Order: 1, Title: "Why Saving Matters", Body: "Saving money isn't about being cheap — it's about buying freedom.", Example: "Maria saves $200/month. After 2 years she has $4,800.", Tip: "Savings = options. More savings = more choices in life."},
					{ID: "s2", Order: 2, Title: "Pay Yourself First", Body: "Automate a transfer to savings the same day your paycheck arrives.", Example: "Set up automatic transfer of $300 on every payday.", Tip: "Automation removes the need for discipline."},
				},
				Outcomes: []contentModel.LessonOutcome{
					{Text: "Understand the importance of saving"},
					{Text: "Learn automation strategies"},
					{Text: "Create a personal savings plan"},
				},
			},
			"lesson-fraud-1": {
				ID: "lesson-fraud-1", TopicID: "fraud-security", Title: "Spot a Phishing Message",
				Summary:          "Recognize suspicious urgency and fake links.",
				EstimatedMinutes: 7,
				Content:          []string{"Do not trust urgency alone.", "Check sender details.", "Never share one-time codes."},
				Quiz:             []contentModel.QuizQuestion{{ID: "q1", Prompt: "What is a common phishing sign?", Options: []string{"Urgent pressure", "Clear official process", "Verified office visit"}, Answer: "Urgent pressure"}},
				Steps: []contentModel.LessonStep{
					{ID: "f1", Order: 1, Title: "What Is Phishing?", Body: "Phishing is a scam where criminals impersonate trusted entities.", Example: "You receive an email claiming to be from your bank asking to verify account.", Tip: "Banks never ask for passwords via email."},
					{ID: "f2", Order: 2, Title: "Protect Yourself", Body: "Verify sender, check URLs, never share sensitive info.", Example: "Hover over links to see actual URL before clicking.", Tip: "When in doubt, call the official number."},
				},
				Outcomes: []contentModel.LessonOutcome{
					{Text: "Recognize phishing attempts"},
					{Text: "Understand security best practices"},
					{Text: "Know how to report suspicious emails"},
				},
			},
			"lesson-invest-1": {
				ID: "lesson-invest-1", TopicID: "investing", Title: "Risk and Diversification",
				Summary:          "Different assets behave differently, which is why diversification matters.",
				EstimatedMinutes: 10,
				Content:          []string{"Higher return usually means higher risk.", "Diversification spreads exposure.", "Long-term thinking matters."},
				Quiz:             []contentModel.QuizQuestion{{ID: "q1", Prompt: "Why diversify?", Options: []string{"Spread risk", "Guarantee huge returns", "Avoid learning"}, Answer: "Spread risk"}},
				Steps: []contentModel.LessonStep{
					{ID: "i1", Order: 1, Title: "Understanding Risk", Body: "Risk and return are related. Higher potential returns come with higher risk.", Example: "Stocks: high risk/high reward. Bonds: low risk/low reward.", Tip: "Your risk tolerance depends on your age and goals."},
					{ID: "i2", Order: 2, Title: "Diversification Strategy", Body: "Don't put all eggs in one basket. Spread investments across asset types.", Example: "Portfolio: 60% stocks, 30% bonds, 10% cash.", Tip: "Diversification reduces overall portfolio volatility."},
				},
				Outcomes: []contentModel.LessonOutcome{
					{Text: "Understand risk-return relationship"},
					{Text: "Learn diversification principles"},
					{Text: "Create a basic investment strategy"},
				},
			},
			"lesson-debt-1": {
				ID: "lesson-debt-1", TopicID: "debt-management", Title: "Interest and Repayment",
				Summary:          "Learn why paying high-interest debt earlier matters.",
				EstimatedMinutes: 9,
				Content:          []string{"List debts by interest rate.", "Avoid adding new debt during payoff.", "Use avalanche or snowball strategy."},
				Quiz:             []contentModel.QuizQuestion{{ID: "q1", Prompt: "Which debt is often prioritized in the avalanche method?", Options: []string{"Highest interest", "Lowest balance", "Newest debt"}, Answer: "Highest interest"}},
				Steps: []contentModel.LessonStep{
					{ID: "d1", Order: 1, Title: "Understanding Debt", Body: "Debt can be a tool or a trap. Interest rates determine the cost.", Example: "$10,000 at 5% vs 20% APR means $500 vs $2,000 yearly.", Tip: "High-interest debt is your biggest wealth enemy."},
					{ID: "d2", Order: 2, Title: "Repayment Strategies", Body: "Snowball: smallest balance first. Avalanche: highest interest first.", Example: "Avalanche saves more interest. Snowball gives quick wins.", Tip: "Choose what keeps you motivated."},
				},
				Outcomes: []contentModel.LessonOutcome{
					{Text: "Calculate debt interest"},
					{Text: "Choose appropriate repayment strategy"},
					{Text: "Create a debt payoff plan"},
				},
			},
		},
	}
}
