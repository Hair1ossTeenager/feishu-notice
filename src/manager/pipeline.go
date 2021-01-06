package manager

import (
	"feishu-notice/src/tool"
	"github.com/labstack/echo/v4"
	"log"
)

const PipelineObjectKind = "pipeline"

type Pipeline struct {
	GitlabBaseMessage
	GitlabPipelineMessage
	WebHook	string
}

func (p *Pipeline) PipelineNotice(c echo.Context) bool {
	p.WebHook = c.Param("web_hook")
	if err := c.Bind(p); err != nil {
		log.Println("pipeline 参数绑定错误")
		return false
	}
	if !p.checkObjectKind() {
		log.Println("just for pipeline kind")
		return false
	}
	if !p.checkPipelineStatus() {
		return false
	}
	templateFile := "pipeline_card.json"
	d := new(pipelineCardData)
	d.formatData(p)
	return tool.SendFeishuMessage(p.WebHook, tool.Render(d, templateFile))
}

// 判断消息类型是否为 pipeline
func (p *Pipeline) checkObjectKind() bool {
	return p.ObjectKind == PipelineObjectKind
}

// 判断 pipeline 状态，只处理 success 或 failed 状态
func (p *Pipeline) checkPipelineStatus() bool {
	return p.ObjectAttributes.Status == "success" || p.ObjectAttributes.Status == "failed"
}

// 判断 pipeline 状态是否为 success
func (p *Pipeline) checkPipelineSuccess() bool {
	return p.ObjectAttributes.Status == "success"
}

// 飞书卡片模板需要的参数
type pipelineCardData struct {
	BuildStatus		string
	ProjectName		string
	TitleColor		string
	CommitMessage	string
	UserName		string
	CommitId		string
	CommitRef		string
	CommitTime		string
	Success			bool
	Duration		string
	CommitUrl		string
}

// 根据 gitlab 参数初始化飞书卡片参数
func (d *pipelineCardData) formatData(p *Pipeline) {
	d.ProjectName = p.Project.Name
	if p.ObjectAttributes.Status == "success" {
		d.BuildStatus = "构建成功"
		d.TitleColor = "Green"
	} else {
		d.BuildStatus = "构建失败"
		d.TitleColor = "Red"
	}
	d.CommitMessage = p.Commit.Title
	d.UserName = p.User.Name
	d.CommitId = p.Commit.Id[0:8]
	d.CommitRef = p.ObjectAttributes.Ref
	var layout = "2006-01-02 15:04:05"
	d.CommitTime = p.Commit.Timestamp.Format(layout)
	d.Success = p.checkPipelineSuccess()
	d.Duration = tool.SecondsFormat(p.ObjectAttributes.Duration)
	d.CommitUrl = p.Commit.Url
}
