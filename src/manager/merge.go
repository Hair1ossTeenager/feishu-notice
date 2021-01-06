package manager

import (
	"feishu-notice/src/tool"
	"github.com/labstack/echo/v4"
	"log"
)

const MergeObjectKind = "merge_request"

type Merge struct {
	GitlabBaseMessage
	GitlabMergeMessage
	WebHook	string
}

func (m *Merge) MergeNotice(c echo.Context) bool {
	m.WebHook = c.Param("web_hook")
	if err := c.Bind(m); err != nil {
		log.Println("merge 参数绑定错误")
		return false
	}
	if !m.checkObjectKind() {
		log.Println("just for merge notice")
		return false
	}
	templateFile := "merge_card.json"
	d := new(mergeCardData)
	d.formatData(m)
	return tool.SendFeishuMessage(m.WebHook, tool.Render(d, templateFile))
}

// 判断消息类型是否为 merge
func (m *Merge) checkObjectKind() bool {
	return m.ObjectKind == MergeObjectKind
}

// 飞书卡片模板需要的参数
type mergeCardData struct {
	ProjectName			string
	UserName			string
	TitleColor			string
	SourceBranch		string
	TargetBranch		string
	Title				string
	Status				string
	Url					string
}

// 根据 gitlab 参数初始化飞书卡片参数
func (d *mergeCardData) formatData(m *Merge) {
	d.ProjectName = m.Project.Name
	d.UserName = m.User.UserName
	d.TitleColor = "Turquoise"
	d.SourceBranch = m.ObjectAttributes.SourceBranch
	d.TargetBranch = m.ObjectAttributes.TargetBranch
	d.Title = m.ObjectAttributes.Title
	d.Status = m.ObjectAttributes.State
	d.Url = m.ObjectAttributes.Url
}

