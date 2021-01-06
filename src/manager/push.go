package manager

import (
	"feishu-notice/src/tool"
	"github.com/labstack/echo/v4"
	"log"
)

const PushObjectKind = "push"

type Push struct {
	GitlabBaseMessage
	GitlabPushMessage
	WebHook	string
}

func (p *Push) PushNotice(c echo.Context) bool {
	p.WebHook = c.Param("web_hook")
	if err := c.Bind(p); err != nil {
		log.Println("push 参数绑定错误")
		return false
	}
	if !p.checkObjectKind() {
		log.Println("just for push notice")
		return false
	}
	templateFile := "push_card.json"
	d := new(pushCardData)
	d.formatData(p)
	return tool.SendFeishuMessage(p.WebHook, tool.Render(d, templateFile))
}

// 判断消息类型是否为 push
func (p *Push) checkObjectKind() bool {
	return p.ObjectKind == PushObjectKind
}

// 飞书卡片模板需要的参数
type pushCardData struct {
	ProjectName			string
	Ref					string
	UserName			string
	TitleColor			string
	Commits				[]CommitsData
	CommitsLastIndex	int
}

// 根据 gitlab 参数初始化飞书卡片参数
func (d *pushCardData) formatData(p *Push) {
	d.ProjectName = p.Project.Name
	d.Ref = p.Ref
	d.UserName = p.UserName
	d.TitleColor = "Turquoise"
	for _, commit := range p.Commits {
		commit.Id = commit.Id[0:8]
		d.Commits = append(d.Commits, commit)
	}
	d.CommitsLastIndex = len(p.Commits) - 1
}

