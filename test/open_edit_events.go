package test

const EditEventsActionButton = `payload=%7B%22type%22%3A%22block_actions%22%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22container%22%3A%7B%22type%22%3A%22view%22%2C%22view_id%22%3A%22V018GCUV2GK%22%7D%2C%22trigger_id%22%3A%221288231154914.260884079521.ba1595ee20fab577e5ac042a518713fd%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22view%22%3A%7B%22id%22%3A%22V018GCUV2GK%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%225e2%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22R4Dd8%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22shm%3Dp%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22%3DIqU%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22%3DTCli%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22KtU%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%221%5C%2F6%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22fzN3%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22EBY%22%7D%2C%7B%22type%22%3A%22actions%22%2C%22block_id%22%3A%22actions_block%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22edit_events%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22edit_events%22%7D%5D%7D%5D%2C%22private_metadata%22%3A%22%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%7D%7D%2C%22hash%22%3A%221596810888.Sjj3E6JN%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Afalse%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Cancel%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Submit%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3Anull%2C%22root_view_id%22%3A%22V018GCUV2GK%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22actions%22%3A%5B%7B%22action_id%22%3A%22edit_events%22%2C%22block_id%22%3A%22actions_block%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22edit_events%22%2C%22type%22%3A%22button%22%2C%22action_ts%22%3A%221596810895.807186%22%7D%5D%7D`

const EditEventsModal = `{
	"title": {
		"type": "plain_text",
		"text": "bZapp - Edit Events",
		"emoji": true
	},
	"notify_on_close": true,
	"private_metadata": "{\"Index\":1,\"Events\":{\"TodaysEvents\":null,\"TomorrowsEvents\":null},\"Goals\":null}",
	"submit": {
		"type": "plain_text",
		"text": "Add",
		"emoji": true
	},
	"type": "modal",
	"close": {
		"type": "plain_text",
		"text": "Back",
		"emoji": true
	},
	"blocks": [
		{
			"type": "divider"
		},
		{
			"text": {
				"text": "Events",
				"type": "plain_text"
			},
			"type": "header"
		},
		{
			"type": "divider"
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "_No events yet_"
			}
		},
		{
			"type": "divider"
		},
		{
			"type": "input",
			"block_id": "add_event_title_input_block-1",
			"element": {
				"action_id": "add_event_title",
				"type": "plain_text_input",
				"placeholder": {
					"type": "plain_text",
					"text": "Title"
				}
			},
			"label": {
				"type": "plain_text",
				"text": "Add Event"
			}
		},
		{
			"type": "input",
			"block_id": "add_event_day_input_block-1",
			"element": {
				"type": "radio_buttons",
				"action_id": "add_event_day",
				"options": [
					{
						"text": {
							"type": "plain_text",
							"text": "Today",
							"emoji": true
						},
						"value": "today"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "Tomorrow",
							"emoji": true
						},
						"value": "tomorrow"
					}
				]
			},
			"label": {
				"type": "plain_text",
				"text": "Day",
				"emoji": true
			}
		},
		{
			"type": "input",
			"block_id": "add_event_hours_input_block-1",
			"element": {
				"type": "static_select",
				"placeholder": {
					"type": "plain_text",
					"text": "Select hour",
					"emoji": true
				},
				"action_id": "add_event_hour",
				"options": [
					{
						"text": {
							"type": "plain_text",
							"text": "9 AM",
							"emoji": true
						},
						"value": "hour-9"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "10 AM",
							"emoji": true
						},
						"value": "hour-10"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "11 AM",
							"emoji": true
						},
						"value": "hour-11"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "12 PM",
							"emoji": true
						},
						"value": "hour-12"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "1 PM",
							"emoji": true
						},
						"value": "hour-1"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "2 PM",
							"emoji": true
						},
						"value": "hour-2"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "3 PM",
							"emoji": true
						},
						"value": "hour-3"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "4 PM",
							"emoji": true
						},
						"value": "hour-4"
					}
				]
			},
			"label": {
				"type": "plain_text",
				"text": "Hour",
				"emoji": true
			}
		},
		{
			"type": "input",
			"block_id": "add_event_mins_input_block-1",
			"element": {
				"type": "static_select",
				"placeholder": {
					"type": "plain_text",
					"text": "Select Minutes",
					"emoji": true
				},
				"action_id": "add_event_mins",
				"options": [
					{
						"text": {
							"type": "plain_text",
							"text": "00",
							"emoji": true
						},
						"value": "min-0"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "15",
							"emoji": true
						},
						"value": "min-15"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "30",
							"emoji": true
						},
						"value": "min-30"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "45",
							"emoji": true
						},
						"value": "min-45"
					}
				]
			},
			"label": {
				"type": "plain_text",
				"text": "Minutes",
				"emoji": true
			}
		}
	]
}`

