package data

type UID int64

type Captain struct {
	UID  UID    `json:"uid"`
	Name string `json:"name"`
}
