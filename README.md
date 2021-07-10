# Mattermost Webex Cloud Plugin 
[![Build Status](https://img.shields.io/circleci/project/github/mattermost/mattermost-plugin-webex/master)](https://circleci.com/gh/mattermost/mattermost-plugin-webex)
[![Code Coverage](https://img.shields.io/codecov/c/github/mattermost/mattermost-plugin-webex/master)](https://codecov.io/gh/mattermost/mattermost-plugin-webex)
[![Release](https://img.shields.io/github/v/release/mattermost/mattermost-plugin-webex)](https://github.com/mattermost/mattermost-plugin-webex/releases/latest)
[![HW](https://img.shields.io/github/issues/mattermost/mattermost-plugin-webex/Up%20For%20Grabs?color=dark%20green&label=Help%20Wanted)](https://github.com/mattermost/mattermost-plugin-webex/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3A%22Up+For+Grabs%22+label%3A%22Help+Wanted%22)

**Maintainer:** [@cpoile](https://github.com/cpoile)
**Co-Maintainer:** [@mickmister](https://github.com/mickmister)

Start and join voice calls, video calls and use screen sharing with your team members via Cisco Webex Meetings. (We do not support Cisco Webex Teams)

Once enabled, clicking a meeting icon in a Mattermost channel invites team members to join a Webex meeting, hosted using the credentials of the user who initiated the meeting.

![image](https://user-images.githubusercontent.com/915956/65968532-2bff2400-e418-11e9-8479-3e43d5890862.gif)


## Installation

System Administrators need to install the Webex plugin on the Mattermost server to make this functionality available to Mattermost end-users.

### Plugin Installation via Marketplace
In Mattermost version 5.16 and above, the easiest way to install a Mattermost plugin is by clicking on the "Settings" menu button above the channel list. Select "Plugin Marketplace", search for "Webex" and click "Install"

### Manual Plugin Installation via Upload
Alternatively, download one of the binary releases from the GitHub page for the webex plugin

1. Go to Settings --> Plugins --> Upload Plugin. Select the file you downloaded, upload it to the server. In server 5.14+, plugins will automatically be distributed across an Enterprise cluster of Mattermost servers, prior to v5.14 you will need to deploy the plugin on each server manually.

2. Go to settings --> Plugin Management and Enable the Webex Meeting Plugin

## Configuration
Go to Settings --> Scroll down to the Plugins section, and click on Webex Plugin

Insert the Webex Meetings URL for your organization. It is often in the format of `<companyname>.my.webex.com` or `<companyname>.webex.com`.

Depending on your situation, you will want to disable the URL conversion (known unsupported on some case with Linux clients).

## Usage
Easily start and join Webex meetings directly from Mattermost

### Starting a Meeting
There are two methods to initiate a new Webex Meeting from within Mattermost:

1. Clicking the Webex Meeting Button at the top right of the channel 
2. By typing `/webex start` and pressing 'enter' in a chat window


### Joining a Meeting from a channel
If you are the meeting organizer and want to start the meeting for other participants, click on the link that is shown below the "Join Meeting" button. This link brings you directly to the meeting and will ask you to login to Webex if you haven't already.

If you are joining a meeting as a participant, you will only see the "Join Meeting" button in your channel. Simply click it to be brought to the Webex meeting.

After initiating a meeting, if you are the organizer - you will see a second link to start the meeting.

### Advanced Options - Sharing Meetings on behalf of others

`/webex <room id>` - Shares a Join Meeting link for the Webex Personal Room meeting that is associated with the specified Personal Room ID, whether it’s your Personal Meeting Room ID or someone else’s.

`/webex <@mattermost_username>` - Shares a Join Meeting link for the Webex Personal Room meeting that is associated with that Mattermost team member.

If you type `/webex help` in any channel conversation you will be presented with your available options.


### If your email address for Webex login is different than your Mattermost login email address
In some cases, you may need to configure your Webex username manually.
By default, the Webex plugin will use the email address associated with your Mattermost account to setup new meetings. Sometimes, users will use a different email address to login to Webex than they use to login to Mattermost. In this case, you will need to configure the Webex plugin to use the email address associated with your Webex account to setup meetings.

If your email address or username is different between your Webex and Mattermost accounts - you may encounter this error:

`No Personal Room link found at <mycompany>.my.webex.com for your userName: bob, or your email: bob@bob.com. Try setting a room manually with /webex room <room id>.`

This error indicates you need to configure your personal room ID.
To use a specific Webex account instead of using your email address from Mattermost - type `/webex room <room id>` - where `<room id>` is your Webex room ID. Meetings you start will use this ID. This Webex Room ID can be found by:

1. Logging in to your Webex account
2. On the home screen you will see a URL with a username within it ('camille' highlighted in red here as an example).  That username is what you will enter as <room id>.

This setting is required only if your Webex account email address is different from your Mattermost account email address, or if the username of your email address does not match your Personal Meeting Room ID or User name on your Webex site.

If anything changes in the future, and your email address with Mattermost and Webex get changed, you can reset the room ID using `/webex room-reset`

To display your current username settings, simply type `/webex info`
  

## Development

This plugin contains both a server and web app portion.

Use `make dist` to build distributions of the plugin that you can upload to a Mattermost server for testing.

Use `make check-style` to check the style for the whole plugin.

### Server

Inside the `/server` directory, you will find the Go files that make up the server-side of the plugin. Within there, build the plugin like you would any other Go application.

### Web App

Inside the `/webapp` directory, you will find the JS and React files that make up the client-side of the plugin. Within there, modify files and components as necessary. Test your syntax by running `npm run build`.
