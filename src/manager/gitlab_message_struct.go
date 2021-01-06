package manager

import "time"

// gitlab base 通知消息结构
type GitlabBaseMessage struct {
	ObjectKind 			string				`json:"object_kind" query:"object_kind" form:"object_kind"`
	Project				Project				`json:"project" query:"project" form:"project"`
}

type Project struct {
	Name	string	`json:"name" query:"name" form:"name"`
}

// gitlab pipeline 通知消息结构
type GitlabPipelineMessage struct {
	ObjectAttributes	ObjectAttributes	`json:"object_attributes" query:"object_attributes" form:"object_attributes"`
	User				User				`json:"user" query:"user" form:"user"`
	Commit				Commit				`json:"commit" query:"commit" form:"commit"`
}

type ObjectAttributes struct {
	Ref				string		`json:"ref" query:"ref" form:"ref"`
	Status			string		`json:"status" query:"status" form:"status"`
	CreatedAt		string		`json:"created_at" query:"created_at" form:"created_at"`
	Duration		int			`json:"duration" query:"duration" form:"duration"`
	SourceBranch	string		`json:"source_branch" query:"source_branch" form:"source_branch"`
	TargetBranch	string		`json:"target_branch" query:"target_branch" form:"target_branch"`
	Title			string		`json:"title" query:"title" form:"title"`
	Url				string		`json:"url" query:"url" form:"url"`
	State			string		`json:"state" query:"state" form:"state"`
}

type User struct {
	Name		string	`json:"name" query:"name" form:"name"`
	UserName	string	`json:"username" query:"username" form:"username"`
	Email		string	`json:"email" query:"email" form:"email"`
}

type Commit struct {
	Id			string		`json:"id" query:"id" form:"id"`
	Title		string		`json:"title" query:"title" form:"title"`
	Url			string		`json:"url" query:"url" form:"url"`
	Timestamp	time.Time	`json:"timestamp" query:"timestamp" form:"timestamp"`
}

// gitlab push 通知消息结构
type GitlabPushMessage struct {
	UserName	string			`json:"user_name" query:"user_name" form:"user_name"`
	Ref			string			`json:"ref" query:"ref" form:"ref"`
	Commits		[]CommitsData	`json:"commits" query:"commits" form:"commits"`
}

type CommitsData struct {
	Id		string	`json:"id" query:"id" form:"id"`
	Title	string	`json:"title" query:"title" form:"title"`
	Url		string	`json:"url" query:"url" form:"url"`
}

// gitlab merge 通知消息结构
type GitlabMergeMessage struct {
	ObjectAttributes	ObjectAttributes	`json:"object_attributes" query:"object_attributes" form:"object_attributes"`
	User				User				`json:"user" query:"user" form:"user"`
}
