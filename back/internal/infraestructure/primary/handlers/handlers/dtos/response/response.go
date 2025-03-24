package response

import "time"

type BaseResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message,omitempty"`
}

type UserResponse struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Dni           string `json:"dni"`
	HouseID       uint   `json:"house_id"`
	HouseName     string `json:"house_name"`
	VoteGroupID   uint   `json:"vote_group_id"`
	VoteGroupName string `json:"vote_group_name"`
}

type HouseResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	VoteGroupID uint      `json:"vote_group_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type VoteGroupResponse struct {
	ID        uint            `json:"id"`
	Name      string          `json:"name"`
	CreatedAt time.Time       `json:"created_at"`
	Houses    []HouseResponse `json:"houses"`
}
