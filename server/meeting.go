package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mattermost/mattermost-plugin-webex/server/webex"

	"github.com/mattermost/mattermost-server/v5/model"
)

type meetingDetails struct {
	startedByUserID     string
	meetingRoomOfUserID string
	channelID           string
	meetingStatus       string
	roomURL             string
}

type meetingPosts struct {
	createdJoinPost  *model.Post
	createdStartPost *model.Post
}

// startMeeting starts a meeting using details.meetingRoomOfUserId's room
// returns the joinPost, startPost, http status code and a descriptive error
func (p *Plugin) startMeeting(details meetingDetails) (*meetingPosts, int, error) {
	roomURL, err := p.getRoomURLFromMMId(details.meetingRoomOfUserID)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	details.roomURL = roomURL
	return p.startMeetingFromRoomURL(details)
}

// startMeetingFromroomURL starts a meeting using details.roomURL, ignoring details.meetingRoomOfUserId
func (p *Plugin) startMeetingFromRoomURL(details meetingDetails) (*meetingPosts, int, error) {
	webexJoinURL := p.makeJoinURL(details.roomURL)
	webexStartURL := p.makeStartURL(details.roomURL)

	joinPost := &model.Post{
		UserId:    details.startedByUserID,
		ChannelId: details.channelID,
		Message:   fmt.Sprintf("Meeting started at %s.", webexJoinURL),
		Type:      "custom_webex",
		Props: map[string]interface{}{
			"meeting_link":     webexJoinURL,
			"meeting_status":   details.meetingStatus,
			"meeting_topic":    "Webex Meeting",
			"starting_user_id": details.startedByUserID,
		},
	}

	createdJoinPost, appErr := p.API.CreatePost(joinPost)
	if appErr != nil {
		return nil, appErr.StatusCode, appErr
	}

	startPost := &model.Post{
		UserId:    p.botUserID,
		ChannelId: details.channelID,
		Message:   fmt.Sprintf("To start the meeting, click here: %s.", webexStartURL),
	}

	var createdStartPost *model.Post
	if details.meetingStatus == webex.StatusStarted {
		createdStartPost = p.API.SendEphemeralPost(details.startedByUserID, startPost)
	}

	return &meetingPosts{createdJoinPost, createdStartPost}, http.StatusOK, nil
}

func (p *Plugin) makeJoinURL(meetingURL string) string {
	if p.getConfiguration().URLConversion {
		meetingURL = strings.Replace(meetingURL, "webex.com/meet/", "webex.com/join/", 1)
	}

	return meetingURL
}

func (p *Plugin) makeStartURL(meetingURL string) string {
	if p.getConfiguration().URLConversion {
		meetingURL = strings.Replace(meetingURL, "webex.com/meet/", "webex.com/start/", 1)
	}

	return meetingURL
}

func (p *Plugin) getURLFromRoomID(roomID string) (string, error) {
	roomURL, cerr := p.webexClient.GetPersonalMeetingRoomURL(roomID, "", "")
	if cerr != nil {
		return "", cerr
	}

	return roomURL, nil
}

func (p *Plugin) getURLFromNameOrEmail(userName, email string) (string, error) {
	roomURL, err := p.webexClient.GetPersonalMeetingRoomURL("", userName, email)
	if err != nil {
		return "", err
	}

	return roomURL, nil
}

// getroomURLFromMMId will find the correct url for mattermostUserId, or return a message explaining why it couldn't.
func (p *Plugin) getRoomURLFromMMId(mattermostUserID string) (string, error) {
	var roomURL string
	if roomID, err := p.getRoom(mattermostUserID); err == nil && roomID != "" {
		// Look for their url using roomId
		roomURL, err = p.getURLFromRoomID(roomID)
		if err != nil {
			return "", fmt.Errorf("no Personal Room link found at `%s` for the room: `%s`", p.getConfiguration().SiteHost, roomID)
		}
	} else {
		// Look for their url using userName or email
		email, userName, err := p.getEmailAndUserName(mattermostUserID)
		if err != nil {
			return "", fmt.Errorf("error getting email and Username: %v", err)
		}
		roomURL, err = p.getURLFromNameOrEmail(userName, email)
		if err != nil {
			return "", fmt.Errorf("no Personal Room link found at `%s` for your Username: `%s`, or your email: `%s`. Try setting a room manually with `/webex room <room id>`", p.getConfiguration().SiteHost, userName, email)
		}
	}

	return roomURL, nil
}
