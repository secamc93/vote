package response

type BaseResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message,omitempty"`
}

type UserResponse struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Dni           string `json:"dni"`
	HouseID       uint   `json:"houseId"`
	HouseName     string `json:"houseName"`
	VoteGroupID   uint   `json:"voteGroupId"`
	VoteGroupName string `json:"voteGroupName"`
}
