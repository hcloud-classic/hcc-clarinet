package model

import errors "innogrid.com/hcloud-classic/hcc_errors"

// Group : Contain infos of the group
type Group struct {
	ID     int64             `json:"group_id"`
	Name   string            `json:"group_name"`
	Errors []errors.HccError `json:"errors"`
}

// GroupList : Contain list of groups
type GroupList struct {
	Groups   []Group           `json:"group_list"`
	TotalNum int               `json:"total_num"`
	Errors   []errors.HccError `json:"errors"`
}
