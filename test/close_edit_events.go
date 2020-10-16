package test

const CloseEditEvents = `payload=%7B%22type%22%3A%22view_closed%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22view%22%3A%7B%22id%22%3A%22V01CMKMUWUS%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22d4K%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22nnUo%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22COY%2B%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22WyKVYV%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%2210%3A00+ads%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_event%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_today_WyKVYV%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22cgY%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22vcYD%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22loQ%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22PTjSgI%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%2211%3A15+dfs%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_event%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_tomorrow_PTjSgI%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22ofgEq%22%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_title_input_block-5%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add+Event%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22action_id%22%3A%22add_event_title%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Title%22%2C%22emoji%22%3Atrue%7D%2C%22dispatch_action_config%22%3A%7B%22trigger_actions_on%22%3A%5B%22on_enter_pressed%22%5D%7D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_day_input_block-5%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Day%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22action_id%22%3A%22add_event_day%22%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Today%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22today%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Tomorrow%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22tomorrow%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_hours_input_block-5%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Hour%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_hour%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+hour%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%229+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-9%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2210+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-10%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2211+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-11%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2212+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-12%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%221+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-1%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%222+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-2%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%223+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-3%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%224+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-4%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_mins_input_block-5%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_mins%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2200%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-0%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2215%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-15%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2230%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-30%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2245%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-45%22%7D%5D%7D%7D%5D%2C%22private_metadata%22%3A%22%7B%5C%22Index%5C%22%3A5%2C%5C%22Events%5C%22%3A%7B%5C%22TodaysEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22WyKVYV%5C%22%2C%5C%22Title%5C%22%3A%5C%22ads%5C%22%2C%5C%22Day%5C%22%3A%5C%22today%5C%22%2C%5C%22Hour%5C%22%3A10%2C%5C%22Min%5C%22%3A0%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%2C%5C%22TomorrowsEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22PTjSgI%5C%22%2C%5C%22Title%5C%22%3A%5C%22dfs%5C%22%2C%5C%22Day%5C%22%3A%5C%22tomorrow%5C%22%2C%5C%22Hour%5C%22%3A11%2C%5C%22Min%5C%22%3A15%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%7D%2C%5C%22Goals%5C%22%3Anull%7D%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22add_event_title_input_block-4%22%3A%7B%22add_event_title%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22value%22%3A%22sd%22%7D%7D%2C%22add_event_day_input_block-4%22%3A%7B%22add_event_day%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Today%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22today%22%7D%7D%7D%2C%22add_event_hours_input_block-4%22%3A%7B%22add_event_hour%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%221+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-1%22%7D%7D%7D%2C%22add_event_mins_input_block-4%22%3A%7B%22add_event_mins%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2215%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-15%22%7D%7D%7D%7D%7D%2C%22hash%22%3A%221602731683.vC8dpkrM%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp+-+Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Atrue%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Back%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3A%22V01DBEZ35GQ%22%2C%22root_view_id%22%3A%22V01DBEZ35GQ%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22is_cleared%22%3Afalse%7D`

const SummaryModal = `{
	"view": {
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
				"type": "divider"
			},
			{
				"elements": [
					{
						"text": "*Today'sEvents*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
			},
			{
				"type": "divider"
			},
			{
				"block_id": "WyKVYV",
				"text": {
					"text": ":small_orange_diamond: 10:00ads",
					"type": "mrkdwn"
				},
				"type": "section"
			},
			{
				"type": "divider"
			},
			{
				"elements": [
					{
						"text": "*Tomorrow'sEvents*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
			},
			{
				"type": "divider"
			},
			{
				"block_id": "PTjSgI",
				"text": {
					"text": ":small_orange_diamond: 11:15dfs",
					"type": "mrkdwn"
				},
				"type": "section"
			},
			{
				"type": "divider"
			},
			{
				"text": {
					"text": "Goals",
					"type": "plain_text"
				},
				"type": "header"
			},
			{
				"type": "divider"
			},
			{
				"text": {
					"text": "_Nogoalsyet_",
					"type": "mrkdwn"
				},
				"type": "section"
			},
			{
				"type": "divider"
			},
			{
				"block_id": "actions_block",
				"elements": [
					{
						"action_id": "edit_events",
						"text": {
							"emoji": true,
							"text": "EditEvents",
							"type": "plain_text"
						},
						"type": "button",
						"value": "edit_events"
					},
					{
						"action_id": "edit_goals",
						"text": {
							"emoji": true,
							"text": "EditGoals",
							"type": "plain_text"
						},
						"type": "button",
						"value": "edit_goals"
					}
				],
				"type": "actions"
			},
			{
				"block_id": "convo_input_id",
				"element": {
					"action_id": "conversation_select_action_id",
					"default_to_current_conversation": true,
					"response_url_enabled": true,
					"type": "conversations_select"
				},
				"label": {
					"text": "Selectachanneltoposttheresulton",
					"type": "plain_text"
				},
				"type": "input"
			}
		],
		"close": {
			"emoji": true,
			"text": "Cancel",
			"type": "plain_text"
		},
		"private_metadata": "{\"Index\":5,\"Events\":{\"TodaysEvents\":[{\"Id\":\"WyKVYV\",\"Title\":\"ads\",\"Day\":\"today\",\"Hour\":10,\"Min\":0,\"AmPm\":\"AM\"}],\"TomorrowsEvents\":[{\"Id\":\"PTjSgI\",\"Title\":\"dfs\",\"Day\":\"tomorrow\",\"Hour\":11,\"Min\":15,\"AmPm\":\"AM\"}]},\"Goals\":null}",
		"submit": {
			"emoji": true,
			"text": "Submit",
			"type": "plain_text"
		},
		"title": {
			"emoji": true,
			"text": "bZapp",
			"type": "plain_text"
		},
		"type": "modal"
	},
	"view_id": "V01DBEZ35GQ"
}`
