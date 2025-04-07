package dto

type (
	// UserDetails ...
	UserDetails struct {
		ID    uint32  `json:"id"`
		Name  string  `json:"name"`
		Email string  `json:"email"`
		Phone *string `json:"phone"`
	}

	// UserStats ...
	UserStats struct {
		QuestionCount uint32 `json:"question_count"`
		PendingCount  uint32 `json:"pending_count"`
		InReviewCount uint32 `json:"in_review_count"`
		ApprovedCount uint32 `json:"approved_count"`
		RejectedCount uint32 `json:"rejected_count"`
		PaidCount     uint32 `json:"paid_count"`
		UnpaidCount   uint32 `json:"unpaid_count"`
		TotalPayment  uint32 `json:"total_payment"`
	}
)
