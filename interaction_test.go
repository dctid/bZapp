package bZapp

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/dctid/bZapp/format"
	"github.com/dctid/bZapp/mocks"
	"github.com/dctid/bZapp/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

const editEventsActionButton = `payload=%7B%22type%22%3A%22block_actions%22%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22container%22%3A%7B%22type%22%3A%22view%22%2C%22view_id%22%3A%22V018GCUV2GK%22%7D%2C%22trigger_id%22%3A%221288231154914.260884079521.ba1595ee20fab577e5ac042a518713fd%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22view%22%3A%7B%22id%22%3A%22V018GCUV2GK%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%225e2%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22R4Dd8%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22shm%3Dp%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22%3DIqU%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22%3DTCli%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22KtU%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%221%5C%2F6%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22fzN3%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22EBY%22%7D%2C%7B%22type%22%3A%22actions%22%2C%22block_id%22%3A%22actions_block%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22edit_events%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22edit_events%22%7D%5D%7D%5D%2C%22private_metadata%22%3A%22%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%7D%7D%2C%22hash%22%3A%221596810888.Sjj3E6JN%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Afalse%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Cancel%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Submit%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3Anull%2C%22root_view_id%22%3A%22V018GCUV2GK%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22actions%22%3A%5B%7B%22action_id%22%3A%22edit_events%22%2C%22block_id%22%3A%22actions_block%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22edit_events%22%2C%22type%22%3A%22button%22%2C%22action_ts%22%3A%221596810895.807186%22%7D%5D%7D`

const addEventSubmission = `payload=%7B%22type%22%3A%22view_submission%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22trigger_id%22%3A%221426592852789.260884079521.a2d8a66ff913636845d0efec5fa19beb%22%2C%22view%22%3A%7B%22id%22%3A%22V01CMKMUWUS%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22b7e%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22geM6%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%221nlPI%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22WyKVYV%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%2210%3A00+ads%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_event%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_today_WyKVYV%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22YX9%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22YQBW%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22Cyv%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22PTjSgI%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%2211%3A15+dfs%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_event%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_tomorrow_PTjSgI%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%2248vk%22%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_title_input_block-4%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add+Event%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22action_id%22%3A%22add_event_title%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Title%22%2C%22emoji%22%3Atrue%7D%2C%22dispatch_action_config%22%3A%7B%22trigger_actions_on%22%3A%5B%22on_enter_pressed%22%5D%7D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_day_input_block-4%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Day%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22action_id%22%3A%22add_event_day%22%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Today%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22today%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Tomorrow%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22tomorrow%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_hours_input_block-4%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Hour%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_hour%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+hour%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%229+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-9%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2210+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-10%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2211+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-11%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2212+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-12%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%221+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-1%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%222+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-2%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%223+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-3%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%224+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-4%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_mins_input_block-4%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_mins%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2200%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-0%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2215%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-15%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2230%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-30%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2245%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-45%22%7D%5D%7D%7D%5D%2C%22private_metadata%22%3A%22%7B%5C%22Index%5C%22%3A4%2C%5C%22Events%5C%22%3A%7B%5C%22TodaysEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22WyKVYV%5C%22%2C%5C%22Title%5C%22%3A%5C%22ads%5C%22%2C%5C%22Day%5C%22%3A%5C%22today%5C%22%2C%5C%22Hour%5C%22%3A10%2C%5C%22Min%5C%22%3A0%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%2C%5C%22TomorrowsEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22PTjSgI%5C%22%2C%5C%22Title%5C%22%3A%5C%22dfs%5C%22%2C%5C%22Day%5C%22%3A%5C%22tomorrow%5C%22%2C%5C%22Hour%5C%22%3A11%2C%5C%22Min%5C%22%3A15%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%7D%2C%5C%22Goals%5C%22%3Anull%7D%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22add_event_title_input_block-4%22%3A%7B%22add_event_title%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22value%22%3A%22sd%22%7D%7D%2C%22add_event_day_input_block-4%22%3A%7B%22add_event_day%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Today%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22today%22%7D%7D%7D%2C%22add_event_hours_input_block-4%22%3A%7B%22add_event_hour%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%221+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-1%22%7D%7D%7D%2C%22add_event_mins_input_block-4%22%3A%7B%22add_event_mins%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2215%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-15%22%7D%7D%7D%7D%7D%2C%22hash%22%3A%221602731098.6UhF390G%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp+-+Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Atrue%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Back%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3A%22V01DBEZ35GQ%22%2C%22root_view_id%22%3A%22V01DBEZ35GQ%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22response_urls%22%3A%5B%5D%7D`

const removeEventAction = `payload=%7B%22type%22%3A%22block_actions%22%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22container%22%3A%7B%22type%22%3A%22view%22%2C%22view_id%22%3A%22V01CMKMUWUS%22%7D%2C%22trigger_id%22%3A%221435839543684.260884079521.4c15d03291d5cce5b585433ce962f85f%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22view%22%3A%7B%22id%22%3A%22V01CMKMUWUS%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22u174F%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%223iU%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22rX4ME%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22WyKVYV%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%2210%3A00+ads%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_event%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_today_WyKVYV%22%7D%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22YUBFMb%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%2210%3A30+wer%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_event%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_today_YUBFMb%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22r3Mke%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22oISK%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22kGrf%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22PTjSgI%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%2211%3A15+dfs%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_event%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_tomorrow_PTjSgI%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22aoB1b%22%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_title_input_block-4%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add+Event%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22action_id%22%3A%22add_event_title%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Title%22%2C%22emoji%22%3Atrue%7D%2C%22dispatch_action_config%22%3A%7B%22trigger_actions_on%22%3A%5B%22on_enter_pressed%22%5D%7D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_day_input_block-4%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Day%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22action_id%22%3A%22add_event_day%22%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Today%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22today%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Tomorrow%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22tomorrow%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_hours_input_block-4%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Hour%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_hour%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+hour%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%229+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-9%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2210+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-10%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2211+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-11%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2212+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-12%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%221+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-1%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%222+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-2%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%223+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-3%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%224+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-4%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_mins_input_block-4%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_mins%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2200%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-0%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2215%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-15%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2230%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-30%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2245%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-45%22%7D%5D%7D%7D%5D%2C%22private_metadata%22%3A%22%7B%5C%22Index%5C%22%3A4%2C%5C%22Events%5C%22%3A%7B%5C%22TodaysEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22WyKVYV%5C%22%2C%5C%22Title%5C%22%3A%5C%22ads%5C%22%2C%5C%22Day%5C%22%3A%5C%22today%5C%22%2C%5C%22Hour%5C%22%3A10%2C%5C%22Min%5C%22%3A0%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%2C%7B%5C%22Id%5C%22%3A%5C%22YUBFMb%5C%22%2C%5C%22Title%5C%22%3A%5C%22wer%5C%22%2C%5C%22Day%5C%22%3A%5C%22today%5C%22%2C%5C%22Hour%5C%22%3A10%2C%5C%22Min%5C%22%3A30%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%2C%5C%22TomorrowsEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22PTjSgI%5C%22%2C%5C%22Title%5C%22%3A%5C%22dfs%5C%22%2C%5C%22Day%5C%22%3A%5C%22tomorrow%5C%22%2C%5C%22Hour%5C%22%3A11%2C%5C%22Min%5C%22%3A15%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%7D%2C%5C%22Goals%5C%22%3Anull%7D%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22add_event_title_input_block-4%22%3A%7B%22add_event_title%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22value%22%3Anull%7D%7D%2C%22add_event_day_input_block-4%22%3A%7B%22add_event_day%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22selected_option%22%3Anull%7D%7D%2C%22add_event_hours_input_block-4%22%3A%7B%22add_event_hour%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3Anull%7D%7D%2C%22add_event_mins_input_block-4%22%3A%7B%22add_event_mins%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3Anull%7D%7D%7D%7D%2C%22hash%22%3A%221602731085.3929ecke%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp+-+Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Atrue%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Back%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3A%22V01DBEZ35GQ%22%2C%22root_view_id%22%3A%22V01DBEZ35GQ%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22actions%22%3A%5B%7B%22action_id%22%3A%22remove_event%22%2C%22block_id%22%3A%22YUBFMb%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_today_YUBFMb%22%2C%22type%22%3A%22button%22%2C%22action_ts%22%3A%221602731096.674302%22%7D%5D%7D`

const submitPayload = `payload=%7B%22type%22%3A%22view_submission%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22trigger_id%22%3A%221453560111536.260884079521.e5de4788286c9ff634c1f636b4403f34%22%2C%22view%22%3A%7B%22id%22%3A%22V01CTRTFJ0L%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22U%3DZ7%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22ctI%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22QDL8I%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22coEbHc%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%229%3A00+asdf%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22ldv%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22eZ1bq%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%229WSq%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22DZosTr%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%221%3A15+qewr%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%2217a%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22wRY%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22Goals%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%220Fzv6%22%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22U%3DQX%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22NpxGZ%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ACustomer+Questions%3F%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%229ehY%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22rXH%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22n%2BNG%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22BGxR%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATeam+Needs%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22P4oIt%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22bSAnHN%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22sfd%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22DKQT%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%229Oz%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ALearnings%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22Hra%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%2234405%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22N5nhe%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22ajKF%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AQuestions%3F%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22e%2BZ%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22RrMdIA%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22afsasdf%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22WaX%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22%3DCO%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AOther%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22Z0HE%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22rDU3%5C%2F%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22%2BgVwJ%22%7D%2C%7B%22type%22%3A%22actions%22%2C%22block_id%22%3A%22actions_block%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22edit_events%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22edit_events%22%7D%2C%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22edit_goals%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Edit+Goals%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22edit_goals%22%7D%5D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22convo_input_id%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+a+channel+to+post+the+result+on%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22conversations_select%22%2C%22action_id%22%3A%22conversation_select_action_id%22%2C%22default_to_current_conversation%22%3Atrue%2C%22response_url_enabled%22%3Atrue%2C%22initial_conversation%22%3A%22C0133NX0GQN%22%7D%7D%5D%2C%22private_metadata%22%3A%22%7B%5C%22Index%5C%22%3A6%2C%5C%22Events%5C%22%3A%7B%5C%22TodaysEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22coEbHc%5C%22%2C%5C%22Title%5C%22%3A%5C%22asdf%5C%22%2C%5C%22Day%5C%22%3A%5C%22today%5C%22%2C%5C%22Hour%5C%22%3A9%2C%5C%22Min%5C%22%3A0%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%2C%5C%22TomorrowsEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22DZosTr%5C%22%2C%5C%22Title%5C%22%3A%5C%22qewr%5C%22%2C%5C%22Day%5C%22%3A%5C%22tomorrow%5C%22%2C%5C%22Hour%5C%22%3A1%2C%5C%22Min%5C%22%3A15%2C%5C%22AmPm%5C%22%3A%5C%22PM%5C%22%7D%5D%7D%2C%5C%22Goals%5C%22%3A%7B%5C%22Questions%3F%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22RrMdIA%5C%22%2C%5C%22Value%5C%22%3A%5C%22afsasdf%5C%22%7D%5D%2C%5C%22Team+Needs%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22bSAnHN%5C%22%2C%5C%22Value%5C%22%3A%5C%22sfd%5C%22%7D%5D%7D%7D%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22convo_input_id%22%3A%7B%22conversation_select_action_id%22%3A%7B%22type%22%3A%22conversations_select%22%2C%22selected_conversation%22%3A%22C0133NX0GQN%22%7D%7D%7D%7D%2C%22hash%22%3A%221602733478.zZkCQ5Bg%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Afalse%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Cancel%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Submit%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3Anull%2C%22root_view_id%22%3A%22V01CTRTFJ0L%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22response_urls%22%3A%5B%7B%22block_id%22%3A%22convo_input_id%22%2C%22action_id%22%3A%22conversation_select_action_id%22%2C%22channel_id%22%3A%22C0133NX0GQN%22%2C%22response_url%22%3A%22https%3A%5C%2F%5C%2Fhooks.slack.com%5C%2Fapp%5C%2FT7NS02BFB%5C%2F1422986473414%5C%2FG33REsfxvjW86raJ1X7wuiJC%22%7D%5D%7D`

const closeEditEvents = `payload=%7B%22type%22%3A%22view_closed%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22view%22%3A%7B%22id%22%3A%22V01CMKMUWUS%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22d4K%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22nnUo%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22COY%2B%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22WyKVYV%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%2210%3A00+ads%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_event%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_today_WyKVYV%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22cgY%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22vcYD%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22loQ%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22PTjSgI%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%2211%3A15+dfs%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_event%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_tomorrow_PTjSgI%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22ofgEq%22%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_title_input_block-5%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add+Event%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22action_id%22%3A%22add_event_title%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Title%22%2C%22emoji%22%3Atrue%7D%2C%22dispatch_action_config%22%3A%7B%22trigger_actions_on%22%3A%5B%22on_enter_pressed%22%5D%7D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_day_input_block-5%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Day%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22action_id%22%3A%22add_event_day%22%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Today%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22today%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Tomorrow%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22tomorrow%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_hours_input_block-5%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Hour%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_hour%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+hour%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%229+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-9%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2210+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-10%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2211+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-11%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2212+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-12%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%221+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-1%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%222+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-2%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%223+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-3%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%224+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-4%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_mins_input_block-5%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_mins%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2200%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-0%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2215%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-15%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2230%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-30%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2245%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-45%22%7D%5D%7D%7D%5D%2C%22private_metadata%22%3A%22%7B%5C%22Index%5C%22%3A5%2C%5C%22Events%5C%22%3A%7B%5C%22TodaysEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22WyKVYV%5C%22%2C%5C%22Title%5C%22%3A%5C%22ads%5C%22%2C%5C%22Day%5C%22%3A%5C%22today%5C%22%2C%5C%22Hour%5C%22%3A10%2C%5C%22Min%5C%22%3A0%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%2C%5C%22TomorrowsEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22PTjSgI%5C%22%2C%5C%22Title%5C%22%3A%5C%22dfs%5C%22%2C%5C%22Day%5C%22%3A%5C%22tomorrow%5C%22%2C%5C%22Hour%5C%22%3A11%2C%5C%22Min%5C%22%3A15%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%7D%2C%5C%22Goals%5C%22%3Anull%7D%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22add_event_title_input_block-4%22%3A%7B%22add_event_title%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22value%22%3A%22sd%22%7D%7D%2C%22add_event_day_input_block-4%22%3A%7B%22add_event_day%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Today%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22today%22%7D%7D%7D%2C%22add_event_hours_input_block-4%22%3A%7B%22add_event_hour%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%221+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-1%22%7D%7D%7D%2C%22add_event_mins_input_block-4%22%3A%7B%22add_event_mins%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2215%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-15%22%7D%7D%7D%7D%7D%2C%22hash%22%3A%221602731683.vC8dpkrM%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp+-+Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Atrue%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Back%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3A%22V01DBEZ35GQ%22%2C%22root_view_id%22%3A%22V01DBEZ35GQ%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22is_cleared%22%3Afalse%7D`

const editGoalsActionButton = `payload=%7B%22type%22%3A%22block_actions%22%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22container%22%3A%7B%22type%22%3A%22view%22%2C%22view_id%22%3A%22V01CQ79HN5S%22%7D%2C%22trigger_id%22%3A%221411346195543.260884079521.14fdd4f0ec90fe20a07ea8dc9429d891%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22view%22%3A%7B%22id%22%3A%22V01CQ79HN5S%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22uv7%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22aJO%5C%2Fu%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22TNF%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%220x%2BuJ%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22%5C%2F%3DGFQ%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22rn%5C%2Fv%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22v%5C%2Fk9X%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22IT%5C%2F%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22Y4N9%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22287%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22Goals%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22tjv5u%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%222dHPr%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22Yntk%22%7D%2C%7B%22type%22%3A%22actions%22%2C%22block_id%22%3A%22actions_block%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22edit_events%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22edit_events%22%7D%2C%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22edit_goals%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Edit+Goals%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22edit_goals%22%7D%5D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22convo_input_id%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+a+channel+to+post+the+result+on%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22conversations_select%22%2C%22action_id%22%3A%22conversation_select_action_id%22%2C%22default_to_current_conversation%22%3Atrue%2C%22response_url_enabled%22%3Atrue%2C%22initial_conversation%22%3A%22C0133NX0GQN%22%7D%7D%5D%2C%22private_metadata%22%3A%22%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22convo_input_id%22%3A%7B%22conversation_select_action_id%22%3A%7B%22type%22%3A%22conversations_select%22%2C%22selected_conversation%22%3A%22C0133NX0GQN%22%7D%7D%7D%7D%2C%22hash%22%3A%221602618823.sl68Li82%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Afalse%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Cancel%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Submit%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3Anull%2C%22root_view_id%22%3A%22V01CQ79HN5S%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22actions%22%3A%5B%7B%22action_id%22%3A%22edit_goals%22%2C%22block_id%22%3A%22actions_block%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Edit+Goals%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22edit_goals%22%2C%22type%22%3A%22button%22%2C%22action_ts%22%3A%221602618829.770185%22%7D%5D%7D`

const addGoalSubmission = `payload=%7B%22type%22%3A%22view_submission%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22trigger_id%22%3A%221429700110306.260884079521.fb856e1ee2ca59f7d5ecad40e9bb5251%22%2C%22view%22%3A%7B%22id%22%3A%22V01DBFTR588%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%2263mm%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22f%3DE%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ACustomer+Questions%3F%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22%3D3ELE%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22IEv%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22SW1ga%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22LQdD%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATeam+Needs%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22xmBH%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22Zu8L%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22Po84%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22wQKE0%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ALearnings%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22%2BTTl7%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22VyH%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22mHIwq%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22ZdV9%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AQuestions%3F%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22G1e4P%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22n5HP9%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22y0s1c%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22Aw2Kt%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AOther%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%224zp%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%228MMKp%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%222iTP%22%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_goal_category_input_block-6%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal+to+Add+to%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_goal_category%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Choose+Goal%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Customer+Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Customer+Questions%3F%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Team+Needs%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Team+Needs%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Learnings%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Learnings%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Questions%3F%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Other%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Other%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_goal_input_block-6%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal+to+Add%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22action_id%22%3A%22add_goal%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal%22%2C%22emoji%22%3Atrue%7D%2C%22dispatch_action_config%22%3A%7B%22trigger_actions_on%22%3A%5B%22on_enter_pressed%22%5D%7D%7D%7D%5D%2C%22private_metadata%22%3A%22%7B%5C%22Index%5C%22%3A6%2C%5C%22Events%5C%22%3A%7B%5C%22TodaysEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22WyKVYV%5C%22%2C%5C%22Title%5C%22%3A%5C%22ads%5C%22%2C%5C%22Day%5C%22%3A%5C%22today%5C%22%2C%5C%22Hour%5C%22%3A10%2C%5C%22Min%5C%22%3A0%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%2C%5C%22TomorrowsEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22PTjSgI%5C%22%2C%5C%22Title%5C%22%3A%5C%22dfs%5C%22%2C%5C%22Day%5C%22%3A%5C%22tomorrow%5C%22%2C%5C%22Hour%5C%22%3A11%2C%5C%22Min%5C%22%3A15%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%7D%2C%5C%22Goals%5C%22%3Anull%7D%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22add_goal_category_input_block-6%22%3A%7B%22add_goal_category%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Team+Needs%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Team+Needs%22%7D%7D%7D%2C%22add_goal_input_block-6%22%3A%7B%22add_goal%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22value%22%3A%22lskfd%22%7D%7D%7D%7D%2C%22hash%22%3A%221602732145.Q2P95KmT%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp+-+Edit+Goals%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Atrue%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Back%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3A%22V01DBEZ35GQ%22%2C%22root_view_id%22%3A%22V01DBEZ35GQ%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22response_urls%22%3A%5B%5D%7D`

const add2ndGoalSubmission = `payload=%7B%22type%22%3A%22view_submission%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22trigger_id%22%3A%221442332863425.260884079521.02e4f5e5fd1ee60ab5596f23d9225917%22%2C%22view%22%3A%7B%22id%22%3A%22V01DBFTR588%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22KOEN%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22ZZrUc%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ACustomer+Questions%3F%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22Lp0a%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22HqqD%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22puSL%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22h6H%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATeam+Needs%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22GC6Of%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22YbiWhf%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22lskfd%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_goal%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_Team+Needs_YbiWhf%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22gXt%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%223cfUK%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ALearnings%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22vrJ%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22BxM%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%225%2BqV%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22%2BZ95%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AQuestions%3F%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22R5%2Bd8%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22i2BxM%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22aPy%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%228F%5C%2FZI%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AOther%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22ed5%5C%2Fm%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22ekQD9%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22avbE%22%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_goal_category_input_block-7%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal+to+Add+to%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_goal_category%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Choose+Goal%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Customer+Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Customer+Questions%3F%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Team+Needs%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Team+Needs%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Learnings%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Learnings%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Questions%3F%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Other%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Other%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_goal_input_block-7%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal+to+Add%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22action_id%22%3A%22add_goal%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal%22%2C%22emoji%22%3Atrue%7D%2C%22dispatch_action_config%22%3A%7B%22trigger_actions_on%22%3A%5B%22on_enter_pressed%22%5D%7D%7D%7D%5D%2C%22private_metadata%22%3A%22%7B%5C%22Index%5C%22%3A7%2C%5C%22Events%5C%22%3A%7B%5C%22TodaysEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22WyKVYV%5C%22%2C%5C%22Title%5C%22%3A%5C%22ads%5C%22%2C%5C%22Day%5C%22%3A%5C%22today%5C%22%2C%5C%22Hour%5C%22%3A10%2C%5C%22Min%5C%22%3A0%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%2C%5C%22TomorrowsEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22PTjSgI%5C%22%2C%5C%22Title%5C%22%3A%5C%22dfs%5C%22%2C%5C%22Day%5C%22%3A%5C%22tomorrow%5C%22%2C%5C%22Hour%5C%22%3A11%2C%5C%22Min%5C%22%3A15%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%7D%2C%5C%22Goals%5C%22%3A%7B%5C%22Team+Needs%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22YbiWhf%5C%22%2C%5C%22Value%5C%22%3A%5C%22lskfd%5C%22%7D%5D%7D%7D%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22add_goal_category_input_block-7%22%3A%7B%22add_goal_category%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Questions%3F%22%7D%7D%7D%2C%22add_goal_input_block-7%22%3A%7B%22add_goal%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22value%22%3A%22adsf%22%7D%7D%7D%7D%2C%22hash%22%3A%221602732470.LikZgL6J%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp+-+Edit+Goals%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Atrue%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Back%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3A%22V01DBEZ35GQ%22%2C%22root_view_id%22%3A%22V01DBEZ35GQ%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22response_urls%22%3A%5B%5D%7D`

const removeGoalAction = `payload=%7B%22type%22%3A%22block_actions%22%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22container%22%3A%7B%22type%22%3A%22view%22%2C%22view_id%22%3A%22V01DBFTR588%22%7D%2C%22trigger_id%22%3A%221435868726004.260884079521.a652912744dfe368ad98783b6348eaaf%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22view%22%3A%7B%22id%22%3A%22V01DBFTR588%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%227R8Zf%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22qXz%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ACustomer+Questions%3F%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22wlaWs%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22rORVV%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22WBnt%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22%3DSab%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATeam+Needs%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%2249bPU%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22YbiWhf%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22lskfd%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_goal%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_Team+Needs_YbiWhf%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22crnz%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22KRJ%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ALearnings%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%222nFK%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22mopNVQ%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22sdfg%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_goal%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_Learnings_mopNVQ%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22ZI8%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%220nc90%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AQuestions%3F%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22%5C%2FWp2%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22GqrOzx%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22adsf%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_goal%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_Questions%3F_GqrOzx%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%2238vc%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22cLb%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AOther%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22esr%5C%2FC%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22qTiZR%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22NR2%22%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_goal_category_input_block-9%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal+to+Add+to%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_goal_category%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Choose+Goal%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Customer+Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Customer+Questions%3F%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Team+Needs%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Team+Needs%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Learnings%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Learnings%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Questions%3F%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Other%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Other%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_goal_input_block-9%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal+to+Add%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22action_id%22%3A%22add_goal%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal%22%2C%22emoji%22%3Atrue%7D%2C%22dispatch_action_config%22%3A%7B%22trigger_actions_on%22%3A%5B%22on_enter_pressed%22%5D%7D%7D%7D%5D%2C%22private_metadata%22%3A%22%7B%5C%22Index%5C%22%3A9%2C%5C%22Events%5C%22%3A%7B%5C%22TodaysEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22WyKVYV%5C%22%2C%5C%22Title%5C%22%3A%5C%22ads%5C%22%2C%5C%22Day%5C%22%3A%5C%22today%5C%22%2C%5C%22Hour%5C%22%3A10%2C%5C%22Min%5C%22%3A0%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%2C%5C%22TomorrowsEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22PTjSgI%5C%22%2C%5C%22Title%5C%22%3A%5C%22dfs%5C%22%2C%5C%22Day%5C%22%3A%5C%22tomorrow%5C%22%2C%5C%22Hour%5C%22%3A11%2C%5C%22Min%5C%22%3A15%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%7D%2C%5C%22Goals%5C%22%3A%7B%5C%22Learnings%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22mopNVQ%5C%22%2C%5C%22Value%5C%22%3A%5C%22sdfg%5C%22%7D%5D%2C%5C%22Questions%3F%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22GqrOzx%5C%22%2C%5C%22Value%5C%22%3A%5C%22adsf%5C%22%7D%5D%2C%5C%22Team+Needs%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22YbiWhf%5C%22%2C%5C%22Value%5C%22%3A%5C%22lskfd%5C%22%7D%5D%7D%7D%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22add_goal_category_input_block-9%22%3A%7B%22add_goal_category%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3Anull%7D%7D%2C%22add_goal_input_block-9%22%3A%7B%22add_goal%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22value%22%3Anull%7D%7D%7D%7D%2C%22hash%22%3A%221602732853.6v4zg0RN%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp+-+Edit+Goals%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Atrue%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Back%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3A%22V01DBEZ35GQ%22%2C%22root_view_id%22%3A%22V01DBEZ35GQ%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22actions%22%3A%5B%7B%22action_id%22%3A%22remove_goal%22%2C%22block_id%22%3A%22GqrOzx%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_Questions%3F_GqrOzx%22%2C%22type%22%3A%22button%22%2C%22action_ts%22%3A%221602732863.821337%22%7D%5D%7D`

const closeEditGoals = `payload=%7B%22type%22%3A%22view_closed%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22view%22%3A%7B%22id%22%3A%22V01DBFTR588%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22vdJ%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22q%5C%2F2%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ACustomer+Questions%3F%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22ZVqzz%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22J35s%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22HRqaD%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%227YN%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATeam+Needs%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22JU0O%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22YbiWhf%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22lskfd%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_goal%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_Team+Needs_YbiWhf%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22sMPV%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22ZlUW%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ALearnings%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%229U%3D%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22mopNVQ%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22sdfg%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_goal%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_Learnings_mopNVQ%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22J7Jx%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%224Vy0%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AQuestions%3F%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22oXQdi%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22Ko9L%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22oakYf%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22QcBbg%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AOther%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22pi5Z6%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%228ud4%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22IFX%22%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_goal_category_input_block-9%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal+to+Add+to%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_goal_category%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Choose+Goal%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Customer+Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Customer+Questions%3F%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Team+Needs%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Team+Needs%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Learnings%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Learnings%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Questions%3F%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Other%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Other%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_goal_input_block-9%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal+to+Add%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22action_id%22%3A%22add_goal%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal%22%2C%22emoji%22%3Atrue%7D%2C%22dispatch_action_config%22%3A%7B%22trigger_actions_on%22%3A%5B%22on_enter_pressed%22%5D%7D%7D%7D%5D%2C%22private_metadata%22%3A%22%7B%5C%22Index%5C%22%3A9%2C%5C%22Events%5C%22%3A%7B%5C%22TodaysEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22WyKVYV%5C%22%2C%5C%22Title%5C%22%3A%5C%22ads%5C%22%2C%5C%22Day%5C%22%3A%5C%22today%5C%22%2C%5C%22Hour%5C%22%3A10%2C%5C%22Min%5C%22%3A0%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%2C%5C%22TomorrowsEvents%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22PTjSgI%5C%22%2C%5C%22Title%5C%22%3A%5C%22dfs%5C%22%2C%5C%22Day%5C%22%3A%5C%22tomorrow%5C%22%2C%5C%22Hour%5C%22%3A11%2C%5C%22Min%5C%22%3A15%2C%5C%22AmPm%5C%22%3A%5C%22AM%5C%22%7D%5D%7D%2C%5C%22Goals%5C%22%3A%7B%5C%22Learnings%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22mopNVQ%5C%22%2C%5C%22Value%5C%22%3A%5C%22sdfg%5C%22%7D%5D%2C%5C%22Questions%3F%5C%22%3A%5B%5D%2C%5C%22Team+Needs%5C%22%3A%5B%7B%5C%22Id%5C%22%3A%5C%22YbiWhf%5C%22%2C%5C%22Value%5C%22%3A%5C%22lskfd%5C%22%7D%5D%7D%7D%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22add_goal_category_input_block-8%22%3A%7B%22add_goal_category%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Learnings%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Learnings%22%7D%7D%7D%2C%22add_goal_input_block-8%22%3A%7B%22add_goal%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22value%22%3A%22sdfg%22%7D%7D%7D%7D%2C%22hash%22%3A%221602732865.hKuDijG7%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp+-+Edit+Goals%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Atrue%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Back%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3A%22V01DBEZ35GQ%22%2C%22root_view_id%22%3A%22V01DBEZ35GQ%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22is_cleared%22%3Afalse%7D`

func TestInteraction(t *testing.T) {
	Client = &mocks.MockClient{}

	type args struct {
		ctx   context.Context
		event events.APIGatewayProxyRequest
	}
	type do struct {
		url  *url.URL
		body string
	}
	var gotDo do

	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
		wantDo  do
	}{
		{
			name: "open edit events",
			args: args{event: events.APIGatewayProxyRequest{Body: editEventsActionButton}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url: getUrl("https://slack.com/api/views.push"),
				body: format.PrettyJsonNoError(fmt.Sprintf(
					`{
								"trigger_id": "1288231154914.260884079521.ba1595ee20fab577e5ac042a518713fd",
    							"view": %s
							}`, editEventsModal)),
			},
		},
		{
			name: "remove event",
			args: args{event: events.APIGatewayProxyRequest{Body: removeEventAction}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url: getUrl("https://slack.com/api/views.update"),
				body: format.PrettyJsonNoError(fmt.Sprintf(
					`{
								"view_id": "V01CMKMUWUS",
    							"view": %s
							}`, removeEventsModal)),
			},
		},
		{
			name: "add event submission",
			args: args{event: events.APIGatewayProxyRequest{Body: addEventSubmission}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
				Body:       format.PrettyJsonNoError(addEventSubmissionResponse),
			},
			wantErr: false,
			wantDo:  do{},
		},
		{
			name: "submit and send message to channel",
			args: args{event: events.APIGatewayProxyRequest{Body: submitPayload}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
			},
			wantErr: false,
			wantDo: do{
				url:  getUrl("https://hooks.slack.com/app/T7NS02BFB/1422986473414/G33REsfxvjW86raJ1X7wuiJC"),
				body: format.PrettyJsonNoError(submissionJson),
			},
		},
		{
			name: "close edit events",
			args: args{event: events.APIGatewayProxyRequest{Body: closeEditEvents}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url:  getUrl("https://slack.com/api/views.update"),
				body: format.PrettyJsonNoError(summaryModal),
			},
		},
		{
			name: "open edit goals actions",
			args: args{event: events.APIGatewayProxyRequest{Body: editGoalsActionButton}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url: getUrl("https://slack.com/api/views.push"),
				body: format.PrettyJsonNoError(fmt.Sprintf(
					`{
								"trigger_id": "1411346195543.260884079521.14fdd4f0ec90fe20a07ea8dc9429d891",
    							"view": %s
							}`, editGoalsModal)),
			},
		},
		{
			name: "add goal submission",
			args: args{event: events.APIGatewayProxyRequest{Body: addGoalSubmission}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
				Body:       format.PrettyJsonNoError(addGoalSubmissionResponse),
			},
			wantErr: false,
			wantDo:  do{},
		},
		{
			name: "add goal 2nd submission",
			args: args{event: events.APIGatewayProxyRequest{Body: add2ndGoalSubmission}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
				Body:       format.PrettyJsonNoError(add2ndGoalSubmissionResponse),
			},
			wantErr: false,
			wantDo:  do{},
		},
		{
			name: "remove goal actions",
			args: args{event: events.APIGatewayProxyRequest{Body: removeGoalAction}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url: getUrl("https://slack.com/api/views.update"),
				body: format.PrettyJsonNoError(fmt.Sprintf(
					`{
								"view_id": "V01DBFTR588",
    							"view": %s
							}`, removeGoalsModal)),
			},
		},
		{
			name: "close edit goals",
			args: args{event: events.APIGatewayProxyRequest{Body: closeEditGoals}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url:  getUrl("https://slack.com/api/views.update"),
				body: format.PrettyJsonNoError(summaryModalWithGoals),
			},
		},
	}
	for _, tt := range tests {
		model.Hash = func() string {
			return "Fake hash"
		}
		gotDo = do{}
		mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
			log.Printf("url %s ", req.URL)
			body, _ := ioutil.ReadAll(req.Body)
			gotDo = do{
				url:  req.URL,
				body: format.PrettyJsonNoError(string(body)),
			}

			return &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(response))),
				StatusCode: 200,
			}, nil
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := Interaction(tt.args.ctx, tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("Interaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				if !assert.EqualValues(t, tt.want.Body, format.PrettyJsonNoError(got.Body)) {
					t.Errorf("Interaction() got = %v, want %v", got, tt.want)
				}
			}
			if tt.wantDo != (do{}) {
				assert.EqualValues(t, tt.wantDo.url, gotDo.url)
				assert.EqualValues(t, tt.wantDo.body, gotDo.body)
			} else {
				assert.EqualValues(t, do{}, gotDo)
			}
		})
	}
}

func getUrl(urlString string) *url.URL {
	result, _ := url.Parse(urlString)
	return result
}

const editEventsModal = `{
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
			"type": "context",
			"elements": [
				{
					"type": "mrkdwn",
					"text": "*Today's Events*"
				}
			]
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
			"type": "context",
			"elements": [
				{
					"type": "mrkdwn",
					"text": "*Tomorrow's Events*"
				}
			]
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

const removeEventsModal = `{
	"type": "modal",
	"title": {
		"type": "plain_text",
		"text": "bZapp - Edit Events",
		"emoji": true
	},
	"blocks": [
		{
			"type": "divider"
		},
		{
			"type": "context",
			"elements": [
				{
					"type": "mrkdwn",
					"text": "*Today's Events*"
				}
			]
		},
		{
			"type": "divider"
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "10:00 ads"
			},
			"block_id": "WyKVYV",
			"accessory": {
				"type": "button",
				"text": {
					"type": "plain_text",
					"text": "Remove",
					"emoji": true
				},
				"action_id": "remove_event",
				"value": "remove_today_WyKVYV"
			}
		},
		{
			"type": "divider"
		},
		{
			"type": "context",
			"elements": [
				{
					"type": "mrkdwn",
					"text": "*Tomorrow's Events*"
				}
			]
		},
		{
			"type": "divider"
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "11:15 dfs"
			},
			"block_id": "PTjSgI",
			"accessory": {
				"type": "button",
				"text": {
					"type": "plain_text",
					"text": "Remove",
					"emoji": true
				},
				"action_id": "remove_event",
				"value": "remove_tomorrow_PTjSgI"
			}
		},
		{
			"type": "divider"
		},
		{
			"type": "input",
			"block_id": "add_event_title_input_block-4",
			"label": {
				"type": "plain_text",
				"text": "Add Event"
			},
			"element": {
				"type": "plain_text_input",
				"action_id": "add_event_title",
				"placeholder": {
					"type": "plain_text",
					"text": "Title"
				}
			}
		},
		{
			"type": "input",
			"block_id": "add_event_day_input_block-4",
			"label": {
				"type": "plain_text",
				"text": "Day",
				"emoji": true
			},
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
			}
		},
		{
			"type": "input",
			"block_id": "add_event_hours_input_block-4",
			"label": {
				"type": "plain_text",
				"text": "Hour",
				"emoji": true
			},
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
			}
		},
		{
			"type": "input",
			"block_id": "add_event_mins_input_block-4",
			"label": {
				"type": "plain_text",
				"text": "Minutes",
				"emoji": true
			},
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
			}
		}
	],
	"close": {
		"type": "plain_text",
		"text": "Back",
		"emoji": true
	},
	"submit": {
		"type": "plain_text",
		"text": "Add",
		"emoji": true
	},
	"private_metadata": "{\"Index\":4,\"Events\":{\"TodaysEvents\":[{\"Id\":\"WyKVYV\",\"Title\":\"ads\",\"Day\":\"today\",\"Hour\":10,\"Min\":0,\"AmPm\":\"AM\"}],\"TomorrowsEvents\":[{\"Id\":\"PTjSgI\",\"Title\":\"dfs\",\"Day\":\"tomorrow\",\"Hour\":11,\"Min\":15,\"AmPm\":\"AM\"}]},\"Goals\":null}",
	"notify_on_close": true
}
`

const addEventSubmissionResponse = `{
  "response_action": "update",
  "view": {
    "type": "modal",
    "title": {
      "type": "plain_text",
      "text": "bZapp - Edit Events",
      "emoji": true
    },
    "blocks": [
      {
        "type": "divider"
      },
      {
        "type": "context",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Today's Events*"
          }
        ]
      },
      {
        "type": "divider"
      },
      {
        "type": "section",
        "text": {
          "type": "mrkdwn",
          "text": "10:00 ads"
        },
        "block_id": "WyKVYV",
        "accessory": {
          "type": "button",
          "text": {
            "type": "plain_text",
            "text": "Remove",
            "emoji": true
          },
          "action_id": "remove_event",
          "value": "remove_today_WyKVYV"
        }
      },
      {
        "type": "section",
        "text": {
          "type": "mrkdwn",
          "text": "1:15 sd"
        },
        "block_id": "Fakehash",
        "accessory": {
          "type": "button",
          "text": {
            "type": "plain_text",
            "text": "Remove",
            "emoji": true
          },
          "action_id": "remove_event",
          "value": "remove_today_Fakehash"
        }
      },
      {
        "type": "divider"
      },
      {
        "type": "context",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Tomorrow's Events*"
          }
        ]
      },
      {
        "type": "divider"
      },
      {
        "type": "section",
        "text": {
          "type": "mrkdwn",
          "text": "11:15 dfs"
        },
        "block_id": "PTjSgI",
        "accessory": {
          "type": "button",
          "text": {
            "type": "plain_text",
            "text": "Remove",
            "emoji": true
          },
          "action_id": "remove_event",
          "value": "remove_tomorrow_PTjSgI"
        }
      },
      {
        "type": "divider"
      },
      {
        "type": "input",
        "block_id": "add_event_title_input_block-5",
        "label": {
          "type": "plain_text",
          "text": "Add Event"
        },
        "element": {
          "type": "plain_text_input",
          "action_id": "add_event_title",
          "placeholder": {
            "type": "plain_text",
            "text": "Title"
          }
        }
      },
      {
        "type": "input",
        "block_id": "add_event_day_input_block-5",
        "label": {
          "type": "plain_text",
          "text": "Day",
          "emoji": true
        },
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
        }
      },
      {
        "type": "input",
        "block_id": "add_event_hours_input_block-5",
        "label": {
          "type": "plain_text",
          "text": "Hour",
          "emoji": true
        },
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
        }
      },
      {
        "type": "input",
        "block_id": "add_event_mins_input_block-5",
        "label": {
          "type": "plain_text",
          "text": "Minutes",
          "emoji": true
        },
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
        }
      }
    ],
    "close": {
      "type": "plain_text",
      "text": "Back",
      "emoji": true
    },
    "submit": {
      "type": "plain_text",
      "text": "Add",
      "emoji": true
    },
    "private_metadata": "{\"Index\":5,\"Events\":{\"TodaysEvents\":[{\"Id\":\"WyKVYV\",\"Title\":\"ads\",\"Day\":\"today\",\"Hour\":10,\"Min\":0,\"AmPm\":\"AM\"},{\"Id\":\"Fakehash\",\"Title\":\"sd\",\"Day\":\"today\",\"Hour\":1,\"Min\":15,\"AmPm\":\"PM\"}],\"TomorrowsEvents\":[{\"Id\":\"PTjSgI\",\"Title\":\"dfs\",\"Day\":\"tomorrow\",\"Hour\":11,\"Min\":15,\"AmPm\":\"AM\"}]},\"Goals\":null}",
    "notify_on_close": true
  }
}`

const submissionJson = `{
	"blocks": [
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
			"block_id": "coEbHc",
			"text": {
				"text": "9:00asdf",
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
			"block_id": "DZosTr",
			"text": {
				"text": "1:15qewr",
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
					"text": "Goals",
					"type": "mrkdwn"
				}
			],
			"type": "context"
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
					"text": "*CustomerQuestions?*",
					"type": "mrkdwn"
				}
			],
			"type": "context"
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
			"elements": [
				{
					"text": "*TeamNeeds*",
					"type": "mrkdwn"
				}
			],
			"type": "context"
		},
		{
			"type": "divider"
		},
		{
			"block_id": "bSAnHN",
			"text": {
				"text": "sfd",
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
					"text": "*Learnings*",
					"type": "mrkdwn"
				}
			],
			"type": "context"
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
			"elements": [
				{
					"text": "*Questions?*",
					"type": "mrkdwn"
				}
			],
			"type": "context"
		},
		{
			"type": "divider"
		},
		{
			"block_id": "RrMdIA",
			"text": {
				"text": "afsasdf",
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
					"text": "*Other*",
					"type": "mrkdwn"
				}
			],
			"type": "context"
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
		}
	],
	"delete_original": false,
	"replace_original": false,
	"response_type": "in_channel",
	"text": "bZapp-Today'sStandupSummary"
}`

const summaryModal = `{
	"view": {
		"blocks": [
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
					"text": "10:00ads",
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
					"text": "11:15dfs",
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
						"text": "Goals",
						"type": "mrkdwn"
					}
				],
				"type": "context"
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
						"text": "*CustomerQuestions?*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
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
				"elements": [
					{
						"text": "*TeamNeeds*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
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
				"elements": [
					{
						"text": "*Learnings*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
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
				"elements": [
					{
						"text": "*Questions?*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
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
				"elements": [
					{
						"text": "*Other*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
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

const editGoalsModal = `{
		"blocks": [
			{
				"type": "divider"
			},
			{
				"elements": [
					{
						"text": "*CustomerQuestions?*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
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
				"elements": [
					{
						"text": "*TeamNeeds*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
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
				"elements": [
					{
						"text": "*Learnings*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
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
				"elements": [
					{
						"text": "*Questions?*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
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
				"elements": [
					{
						"text": "*Other*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
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
				"block_id": "add_goal_category_input_block-1",
				"element": {
					"action_id": "add_goal_category",
					"options": [
						{
							"text": {
								"emoji": true,
								"text": "CustomerQuestions?",
								"type": "plain_text"
							},
							"value": "goal-CustomerQuestions?"
						},
						{
							"text": {
								"emoji": true,
								"text": "TeamNeeds",
								"type": "plain_text"
							},
							"value": "goal-TeamNeeds"
						},
						{
							"text": {
								"emoji": true,
								"text": "Learnings",
								"type": "plain_text"
							},
							"value": "goal-Learnings"
						},
						{
							"text": {
								"emoji": true,
								"text": "Questions?",
								"type": "plain_text"
							},
							"value": "goal-Questions?"
						},
						{
							"text": {
								"emoji": true,
								"text": "Other",
								"type": "plain_text"
							},
							"value": "goal-Other"
						}
					],
					"placeholder": {
						"emoji": true,
						"text": "ChooseGoal",
						"type": "plain_text"
					},
					"type": "static_select"
				},
				"label": {
					"emoji": true,
					"text": "GoaltoAddto",
					"type": "plain_text"
				},
				"type": "input"
			},
			{
				"block_id": "add_goal_input_block-1",
				"element": {
					"action_id": "add_goal",
					"placeholder": {
						"text": "Goal",
						"type": "plain_text"
					},
					"type": "plain_text_input"
				},
				"label": {
					"emoji": true,
					"text": "GoaltoAdd",
					"type": "plain_text"
				},
				"type": "input"
			}
		],
		"close": {
			"emoji": true,
			"text": "Back",
			"type": "plain_text"
		},
		"notify_on_close": true,
		"private_metadata": "{\"Index\":1,\"Events\":{\"TodaysEvents\":null,\"TomorrowsEvents\":null},\"Goals\":null}",
		"submit": {
			"emoji": true,
			"text": "Add",
			"type": "plain_text"
		},
		"title": {
			"emoji": true,
			"text": "bZapp-EditGoals",
			"type": "plain_text"
		},
		"type": "modal"
	}`

const addGoalSubmissionResponse = `{
  "response_action": "update",
  "view": {
    "type": "modal",
    "title": {
      "type": "plain_text",
      "text": "bZapp - Edit Goals",
      "emoji": true
    },
    "blocks": [
      {
        "type": "divider"
      },
      {
        "type": "context",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Customer Questions?*"
          }
        ]
      },
      {
        "type": "divider"
      },
      {
        "type": "section",
        "text": {
          "type": "mrkdwn",
          "text": "_No goals yet_"
        }
      },
      {
        "type": "divider"
      },
      {
        "type": "context",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Team Needs*"
          }
        ]
      },
      {
        "type": "divider"
      },
      {
        "type": "section",
        "text": {
          "type": "mrkdwn",
          "text": "lskfd"
        },
        "block_id": "Fakehash",
        "accessory": {
          "type": "button",
          "text": {
            "type": "plain_text",
            "text": "Remove",
            "emoji": true
          },
          "action_id": "remove_goal",
          "value": "remove_Team Needs_Fakehash"
        }
      },
      {
        "type": "divider"
      },
      {
        "type": "context",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Learnings*"
          }
        ]
      },
      {
        "type": "divider"
      },
      {
        "type": "section",
        "text": {
          "type": "mrkdwn",
          "text": "_No goals yet_"
        }
      },
      {
        "type": "divider"
      },
      {
        "type": "context",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Questions?*"
          }
        ]
      },
      {
        "type": "divider"
      },
      {
        "type": "section",
        "text": {
          "type": "mrkdwn",
          "text": "_No goals yet_"
        }
      },
      {
        "type": "divider"
      },
      {
        "type": "context",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Other*"
          }
        ]
      },
      {
        "type": "divider"
      },
      {
        "type": "section",
        "text": {
          "type": "mrkdwn",
          "text": "_No goals yet_"
        }
      },
      {
        "type": "divider"
      },
      {
        "type": "input",
        "block_id": "add_goal_category_input_block-7",
        "label": {
          "type": "plain_text",
          "text": "Goal to Add to",
          "emoji": true
        },
        "element": {
          "type": "static_select",
          "placeholder": {
            "type": "plain_text",
            "text": "Choose Goal",
            "emoji": true
          },
          "action_id": "add_goal_category",
          "options": [
            {
              "text": {
                "type": "plain_text",
                "text": "Customer Questions?",
                "emoji": true
              },
              "value": "goal-Customer Questions?"
            },
            {
              "text": {
                "type": "plain_text",
                "text": "Team Needs",
                "emoji": true
              },
              "value": "goal-Team Needs"
            },
            {
              "text": {
                "type": "plain_text",
                "text": "Learnings",
                "emoji": true
              },
              "value": "goal-Learnings"
            },
            {
              "text": {
                "type": "plain_text",
                "text": "Questions?",
                "emoji": true
              },
              "value": "goal-Questions?"
            },
            {
              "text": {
                "type": "plain_text",
                "text": "Other",
                "emoji": true
              },
              "value": "goal-Other"
            }
          ]
        }
      },
      {
        "type": "input",
        "block_id": "add_goal_input_block-7",
        "label": {
          "type": "plain_text",
          "text": "Goal to Add",
          "emoji": true
        },
        "element": {
          "type": "plain_text_input",
          "action_id": "add_goal",
          "placeholder": {
            "type": "plain_text",
            "text": "Goal"
          }
        }
      }
    ],
    "close": {
      "type": "plain_text",
      "text": "Back",
      "emoji": true
    },
    "submit": {
      "type": "plain_text",
      "text": "Add",
      "emoji": true
    },
    "private_metadata": "{\"Index\":7,\"Events\":{\"TodaysEvents\":[{\"Id\":\"WyKVYV\",\"Title\":\"ads\",\"Day\":\"today\",\"Hour\":10,\"Min\":0,\"AmPm\":\"AM\"}],\"TomorrowsEvents\":[{\"Id\":\"PTjSgI\",\"Title\":\"dfs\",\"Day\":\"tomorrow\",\"Hour\":11,\"Min\":15,\"AmPm\":\"AM\"}]},\"Goals\":{\"Team Needs\":[{\"Id\":\"Fakehash\",\"Value\":\"lskfd\"}]}}",
    "notify_on_close": true
  }
}`

const add2ndGoalSubmissionResponse = `{
  "response_action": "update",
  "view": {
    "type": "modal",
    "title": {
      "type": "plain_text",
      "text": "bZapp - Edit Goals",
      "emoji": true
    },
    "blocks": [
      {
        "type": "divider"
      },
      {
        "type": "context",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Customer Questions?*"
          }
        ]
      },
      {
        "type": "divider"
      },
      {
        "type": "section",
        "text": {
          "type": "mrkdwn",
          "text": "_No goals yet_"
        }
      },
      {
        "type": "divider"
      },
      {
        "type": "context",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Team Needs*"
          }
        ]
      },
      {
        "type": "divider"
      },
      {
        "type": "section",
        "text": {
          "type": "mrkdwn",
          "text": "lskfd"
        },
        "block_id": "YbiWhf",
        "accessory": {
          "type": "button",
          "text": {
            "type": "plain_text",
            "text": "Remove",
            "emoji": true
          },
          "action_id": "remove_goal",
          "value": "remove_Team Needs_YbiWhf"
        }
      },
      {
        "type": "divider"
      },
      {
        "type": "context",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Learnings*"
          }
        ]
      },
      {
        "type": "divider"
      },
      {
        "type": "section",
        "text": {
          "type": "mrkdwn",
          "text": "_No goals yet_"
        }
      },
      {
        "type": "divider"
      },
      {
        "type": "context",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Questions?*"
          }
        ]
      },
      {
        "type": "divider"
      },
      {
        "type": "section",
        "text": {
          "type": "mrkdwn",
          "text": "adsf"
        },
        "block_id": "Fakehash",
        "accessory": {
          "type": "button",
          "text": {
            "type": "plain_text",
            "text": "Remove",
            "emoji": true
          },
          "action_id": "remove_goal",
          "value": "remove_Questions?_Fakehash"
        }
      },
      {
        "type": "divider"
      },
      {
        "type": "context",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Other*"
          }
        ]
      },
      {
        "type": "divider"
      },
      {
        "type": "section",
        "text": {
          "type": "mrkdwn",
          "text": "_No goals yet_"
        }
      },
      {
        "type": "divider"
      },
      {
        "type": "input",
        "block_id": "add_goal_category_input_block-8",
        "label": {
          "type": "plain_text",
          "text": "Goal to Add to",
          "emoji": true
        },
        "element": {
          "type": "static_select",
          "placeholder": {
            "type": "plain_text",
            "text": "Choose Goal",
            "emoji": true
          },
          "action_id": "add_goal_category",
          "options": [
            {
              "text": {
                "type": "plain_text",
                "text": "Customer Questions?",
                "emoji": true
              },
              "value": "goal-Customer Questions?"
            },
            {
              "text": {
                "type": "plain_text",
                "text": "Team Needs",
                "emoji": true
              },
              "value": "goal-Team Needs"
            },
            {
              "text": {
                "type": "plain_text",
                "text": "Learnings",
                "emoji": true
              },
              "value": "goal-Learnings"
            },
            {
              "text": {
                "type": "plain_text",
                "text": "Questions?",
                "emoji": true
              },
              "value": "goal-Questions?"
            },
            {
              "text": {
                "type": "plain_text",
                "text": "Other",
                "emoji": true
              },
              "value": "goal-Other"
            }
          ]
        }
      },
      {
        "type": "input",
        "block_id": "add_goal_input_block-8",
        "label": {
          "type": "plain_text",
          "text": "Goal to Add",
          "emoji": true
        },
        "element": {
          "type": "plain_text_input",
          "action_id": "add_goal",
          "placeholder": {
            "type": "plain_text",
            "text": "Goal"
          }
        }
      }
    ],
    "close": {
      "type": "plain_text",
      "text": "Back",
      "emoji": true
    },
    "submit": {
      "type": "plain_text",
      "text": "Add",
      "emoji": true
    },
    "private_metadata": "{\"Index\":8,\"Events\":{\"TodaysEvents\":[{\"Id\":\"WyKVYV\",\"Title\":\"ads\",\"Day\":\"today\",\"Hour\":10,\"Min\":0,\"AmPm\":\"AM\"}],\"TomorrowsEvents\":[{\"Id\":\"PTjSgI\",\"Title\":\"dfs\",\"Day\":\"tomorrow\",\"Hour\":11,\"Min\":15,\"AmPm\":\"AM\"}]},\"Goals\":{\"Questions?\":[{\"Id\":\"Fakehash\",\"Value\":\"adsf\"}],\"Team Needs\":[{\"Id\":\"YbiWhf\",\"Value\":\"lskfd\"}]}}",
    "notify_on_close": true
  }
}`

const removeGoalsModal = `{
	"type": "modal",
	"title": {
		"type": "plain_text",
		"text": "bZapp - Edit Goals",
		"emoji": true
	},
	"blocks": [
		{
			"type": "divider"
		},
		{
			"type": "context",
			"elements": [
				{
					"type": "mrkdwn",
					"text": "*Customer Questions?*"
				}
			]
		},
		{
			"type": "divider"
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "_No goals yet_"
			}
		},
		{
			"type": "divider"
		},
		{
			"type": "context",
			"elements": [
				{
					"type": "mrkdwn",
					"text": "*Team Needs*"
				}
			]
		},
		{
			"type": "divider"
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "lskfd"
			},
			"block_id": "YbiWhf",
			"accessory": {
				"type": "button",
				"text": {
					"type": "plain_text",
					"text": "Remove",
					"emoji": true
				},
				"action_id": "remove_goal",
				"value": "remove_Team Needs_YbiWhf"
			}
		},
		{
			"type": "divider"
		},
		{
			"type": "context",
			"elements": [
				{
					"type": "mrkdwn",
					"text": "*Learnings*"
				}
			]
		},
		{
			"type": "divider"
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "sdfg"
			},
			"block_id": "mopNVQ",
			"accessory": {
				"type": "button",
				"text": {
					"type": "plain_text",
					"text": "Remove",
					"emoji": true
				},
				"action_id": "remove_goal",
				"value": "remove_Learnings_mopNVQ"
			}
		},
		{
			"type": "divider"
		},
		{
			"type": "context",
			"elements": [
				{
					"type": "mrkdwn",
					"text": "*Questions?*"
				}
			]
		},
		{
			"type": "divider"
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "_No goals yet_"
			}
		},
		{
			"type": "divider"
		},
		{
			"type": "context",
			"elements": [
				{
					"type": "mrkdwn",
					"text": "*Other*"
				}
			]
		},
		{
			"type": "divider"
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "_No goals yet_"
			}
		},
		{
			"type": "divider"
		},
		{
			"type": "input",
			"block_id": "add_goal_category_input_block-9",
			"label": {
				"type": "plain_text",
				"text": "Goal to Add to",
				"emoji": true
			},
			"element": {
				"type": "static_select",
				"placeholder": {
					"type": "plain_text",
					"text": "Choose Goal",
					"emoji": true
				},
				"action_id": "add_goal_category",
				"options": [
					{
						"text": {
							"type": "plain_text",
							"text": "Customer Questions?",
							"emoji": true
						},
						"value": "goal-Customer Questions?"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "Team Needs",
							"emoji": true
						},
						"value": "goal-Team Needs"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "Learnings",
							"emoji": true
						},
						"value": "goal-Learnings"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "Questions?",
							"emoji": true
						},
						"value": "goal-Questions?"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "Other",
							"emoji": true
						},
						"value": "goal-Other"
					}
				]
			}
		},
		{
			"type": "input",
			"block_id": "add_goal_input_block-9",
			"label": {
				"type": "plain_text",
				"text": "Goal to Add",
				"emoji": true
			},
			"element": {
				"type": "plain_text_input",
				"action_id": "add_goal",
				"placeholder": {
					"type": "plain_text",
					"text": "Goal"
				}
			}
		}
	],
	"close": {
		"type": "plain_text",
		"text": "Back",
		"emoji": true
	},
	"submit": {
		"type": "plain_text",
		"text": "Add",
		"emoji": true
	},
	"private_metadata": "{\"Index\":9,\"Events\":{\"TodaysEvents\":[{\"Id\":\"WyKVYV\",\"Title\":\"ads\",\"Day\":\"today\",\"Hour\":10,\"Min\":0,\"AmPm\":\"AM\"}],\"TomorrowsEvents\":[{\"Id\":\"PTjSgI\",\"Title\":\"dfs\",\"Day\":\"tomorrow\",\"Hour\":11,\"Min\":15,\"AmPm\":\"AM\"}]},\"Goals\":{\"Learnings\":[{\"Id\":\"mopNVQ\",\"Value\":\"sdfg\"}],\"Questions?\":[],\"Team Needs\":[{\"Id\":\"YbiWhf\",\"Value\":\"lskfd\"}]}}",
	"notify_on_close": true
}`

const summaryModalWithGoals = `{
	"view": {
		"blocks": [
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
					"text": "10:00ads",
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
					"text": "11:15dfs",
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
						"text": "Goals",
						"type": "mrkdwn"
					}
				],
				"type": "context"
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
						"text": "*CustomerQuestions?*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
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
				"elements": [
					{
						"text": "*TeamNeeds*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
			},
			{
				"type": "divider"
			},
			{
				"block_id": "YbiWhf",
				"text": {
					"text": "lskfd",
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
						"text": "*Learnings*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
			},
			{
				"type": "divider"
			},
			{
				"block_id": "mopNVQ",
				"text": {
					"text": "sdfg",
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
						"text": "*Questions?*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
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
				"elements": [
					{
						"text": "*Other*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
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
		"private_metadata": "{\"Index\":9,\"Events\":{\"TodaysEvents\":[{\"Id\":\"WyKVYV\",\"Title\":\"ads\",\"Day\":\"today\",\"Hour\":10,\"Min\":0,\"AmPm\":\"AM\"}],\"TomorrowsEvents\":[{\"Id\":\"PTjSgI\",\"Title\":\"dfs\",\"Day\":\"tomorrow\",\"Hour\":11,\"Min\":15,\"AmPm\":\"AM\"}]},\"Goals\":{\"Learnings\":[{\"Id\":\"mopNVQ\",\"Value\":\"sdfg\"}],\"Questions?\":[],\"TeamNeeds\":[{\"Id\":\"YbiWhf\",\"Value\":\"lskfd\"}]}}",
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