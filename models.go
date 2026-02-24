package models

// BulkMessage represents the structure for bulk messaging.
type BulkMessage struct {
    ID          string   `json:"id"` // Unique identifier for the bulk message
    SenderID    string   `json:"sender_id"` // ID of the sender
    RecipientIDs []string `json:"recipient_ids"` // List of recipient IDs
    Message     string   `json:"message"` // The message to be sent
    SentAt      string   `json:"sent_at"` // Timestamp of when the message was sent
}

// RateLimiting holds the data for rate limiting mechanisms.
type RateLimiting struct {
    Limit      int    `json:"limit"`       // Maximum number of messages allowed
    Interval   int    `json:"interval"`    // Time interval (in seconds) for the limit
    Exceeded   bool   `json:"exceeded"`    // Flag to indicate if the limit is exceeded
}

// BanPrevention stores information for preventing users from being banned.
type BanPrevention struct {
    UserID   string  `json:"user_id"`          // ID of the user
    Banned    bool    `json:"banned"`           // Whether the user is banned
    BanReason string   `json:"ban_reason"`      // Reason for the ban
    BanUntil  string   `json:"ban_until"`       // Timestamp until when the user is banned
    Appeal     bool    `json:"appeal"`         // Flag to indicate if the user can appeal the ban
}