package notification

import (
	"server/database/repository"
	"server/models"
	"server/services/eventBus"
	"strings"
)

func (ns *NotificationSubscriber) HandleTaskDescriptionMentionEvent(change eventBus.Change[models.Task], u models.User) {
	user, found := ns.StringContainsMention(change.New.Description)
	if !found {
		return
	}

	ns.me.Publish("mentioned", eventBus.TaskOrComment{Task: &change.Old, MentionedUser: &user}, eventBus.TaskOrComment{Task: &change.New, MentionedUser: &u}, u)
}

func (ns *NotificationSubscriber) HandleCommentTextMentionEvent(change eventBus.Change[models.Comment], u models.User) {
	user, found := ns.StringContainsMention(change.New.Text)
	if !found {
		return
	}

	ns.me.Publish("mentioned", eventBus.TaskOrComment{Comment: &change.Old, MentionedUser: &user}, eventBus.TaskOrComment{Comment: &change.New, MentionedUser: &u}, u)
}

func (ns *NotificationSubscriber) StringContainsMention(text string) (models.User, bool) {
	atIndex := strings.Index(text, "@")
	if atIndex == -1 {
		return models.User{}, false
	}

	mentionPortion := text[atIndex+1:]

	spaceIndex := strings.IndexAny(mentionPortion, " \t\n")
	var foundMention string
	if spaceIndex == -1 {
		foundMention = mentionPortion
	} else {
		foundMention = mentionPortion[:spaceIndex]
	}

	foundMention = strings.TrimSpace(foundMention)
	if foundMention == "" {
		return models.User{}, false
	}

	user, err := ns.db.UserRepository.GetFirst(repository.WithWhere("username = ?", foundMention))
	if err != nil {
		return models.User{}, false
	}

	return user, true
}
