package subscription

type SubscriptionRepository interface {
	SaveSubscription(teacherID, yearID, method string, qtty int) error
	UpdateSubscription(subID, method string, qtty int) error
	DeleteSubscription(subID string) error
}
