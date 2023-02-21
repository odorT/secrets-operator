package notification

import (
	"fmt"
	"github.com/slack-go/slack"
	"go.uber.org/zap"
	"secrets-operator/config"
	"secrets-operator/internal/core/domain"
)

type slackNotifier struct {
	cfg    *config.Config
	l      *zap.SugaredLogger
	client *slack.Client
}

func NewSlackNotifier(cfg *config.Config, l *zap.SugaredLogger) *slackNotifier {

	return &slackNotifier{
		cfg:    cfg,
		l:      l,
		client: slack.New(cfg.SlackAuthToken, slack.OptionDebug(cfg.SlackDebugEnabled)),
	}
}

func (sl slackNotifier) SendMessage(message domain.FindingsReport) error {

	attachment := slack.Attachment{
		Title: fmt.Sprintf("Found new hard coded secrets in %s 😐", message.RepoName),
		Color: "#FF0000",
		Text:  fmt.Sprintf("%s's commit included secret(ish) information. Please check", message.CommitAuthor),
		Fields: []slack.AttachmentField{
			{
				Title: "Repository",
				Value: message.RepoURL,
			},
			{
				Title: "Pipeline",
				Value: fmt.Sprintf("%s", message.BuildPipelineURL()),
			},
			{
				Title: "Commit",
				Value: fmt.Sprintf("%s", message.BuildCommitURL()),
			},
			{
				Title: "How many findings found?",
				Value: fmt.Sprintf("%d", len(message.Findings)),
			},
			{
				Title: "Date",
				Value: message.Timestamp.String(),
			},
		},
	}

	_, _, err := sl.client.PostMessage(
		sl.cfg.SlackChannelId,
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		return err
	}

	return nil
}
