package request

type VoteGroupRequest struct {
	Name string `json:"name"`
}

type UserRequest struct {
	Name        string `json:"name"`
	Dni         string `json:"dni"`
	HouseID     uint   `json:"house_id"`
	VoteGroupID uint   `json:"vote_group_id"`
}
type HouseRequest struct {
	Name        string `json:"name"`
	VoteGroupID uint   `json:"vote_group_id"`
}
