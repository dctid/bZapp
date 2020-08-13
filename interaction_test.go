package bZapp

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/dctid/bZapp/mocks"
	"github.com/dctid/bZapp/test"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

var editEventsActionButton = `payload=%7B%22type%22%3A%22block_actions%22%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22container%22%3A%7B%22type%22%3A%22view%22%2C%22view_id%22%3A%22V018GCUV2GK%22%7D%2C%22trigger_id%22%3A%221288231154914.260884079521.ba1595ee20fab577e5ac042a518713fd%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22view%22%3A%7B%22id%22%3A%22V018GCUV2GK%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%225e2%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22R4Dd8%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22shm%3Dp%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22%3DIqU%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22%3DTCli%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22KtU%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%221%5C%2F6%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22fzN3%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22EBY%22%7D%2C%7B%22type%22%3A%22actions%22%2C%22block_id%22%3A%22actions_block%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22edit_events%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22edit_events%22%7D%5D%7D%5D%2C%22private_metadata%22%3A%22%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%7D%7D%2C%22hash%22%3A%221596810888.Sjj3E6JN%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Afalse%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Cancel%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Submit%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3Anull%2C%22root_view_id%22%3A%22V018GCUV2GK%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22actions%22%3A%5B%7B%22action_id%22%3A%22edit_events%22%2C%22block_id%22%3A%22actions_block%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22edit_events%22%2C%22type%22%3A%22button%22%2C%22action_ts%22%3A%221596810895.807186%22%7D%5D%7D`

var addEventSubmission = `payload=%7B%22type%22%3A%22view_submission%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22trigger_id%22%3A%221296901702898.260884079521.664b4e09d83d7c1e1bb094385e9b49a3%22%2C%22view%22%3A%7B%22id%22%3A%22V018X2J9UA0%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22wrQ%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22r1Bj6%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22iMR%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%228HQ4g%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22T2I%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22Bdcei%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22nxzS%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22YtXZ%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22MJYT%22%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_title_input_block%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add+Event%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22action_id%22%3A%22add_event_title%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Title%22%2C%22emoji%22%3Atrue%7D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_day_input_block%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Day%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22action_id%22%3A%22add_event_day%22%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Today%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22today%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Tomorrow%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22tomorrow%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_hours_input_block%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Hour%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_hour%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+hour%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%229+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-9%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2210+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-10%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2211+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-11%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2212+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-12%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%221+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-1%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%222+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-2%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%223+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-3%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%224+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-4%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_mins_input_block%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_mins%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2200%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-0%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2215%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-15%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2230%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-30%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2245%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-45%22%7D%5D%7D%7D%5D%2C%22private_metadata%22%3A%22%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22add_event_title_input_block%22%3A%7B%22add_event_title%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22value%22%3A%22retrob%22%7D%7D%2C%22add_event_day_input_block%22%3A%7B%22add_event_day%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Tomorrow%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22tomorrow%22%7D%7D%7D%2C%22add_event_hours_input_block%22%3A%7B%22add_event_hour%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%224+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-4%22%7D%7D%7D%2C%22add_event_mins_input_block%22%3A%7B%22add_event_mins%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2245%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-45%22%7D%7D%7D%7D%7D%2C%22hash%22%3A%221597247878.AteqceCB%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp+-+Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Afalse%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Cancel%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3Anull%2C%22root_view_id%22%3A%22V018X2J9UA0%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22response_urls%22%3A%5B%5D%7D`

var removeAction = `payload=%7B%22type%22%3A%22block_actions%22%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22container%22%3A%7B%22type%22%3A%22view%22%2C%22view_id%22%3A%22V018M8X511B%22%7D%2C%22trigger_id%22%3A%221294678404786.260884079521.830a105661b8cf7f5f1d799002760772%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22view%22%3A%7B%22id%22%3A%22V018M8X511B%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22d2%5C%2Fe6%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22VAG%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22ZMPU1%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%224MWhf%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%2211%3A15+asdf%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_event%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_today_0%22%7D%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22ivQ%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%2211%3A30+dfgh%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_event%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_today_1%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22CUj%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%227PD%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%2296L%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22bpxH%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%2210%3A15+asdf%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_event%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_tomorrow_0%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22yPup%22%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_title_input_block%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add+Event%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22action_id%22%3A%22add_event_title%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Title%22%2C%22emoji%22%3Atrue%7D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_day_input_block%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Day%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22action_id%22%3A%22add_event_day%22%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Today%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22today%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Tomorrow%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22tomorrow%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_hours_input_block%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Hour%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_hour%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+hour%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%229+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-9%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2210+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-10%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2211+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-11%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2212+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-12%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%221+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-1%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%222+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-2%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%223+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-3%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%224+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-4%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_mins_input_block%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_mins%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2200%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-0%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2215%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-15%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2230%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-30%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2245%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-45%22%7D%5D%7D%7D%5D%2C%22private_metadata%22%3A%22test+metadata%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22add_event_title_input_block%22%3A%7B%22add_event_title%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22value%22%3A%22dfgh%22%7D%7D%2C%22add_event_day_input_block%22%3A%7B%22add_event_day%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Today%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22today%22%7D%7D%7D%2C%22add_event_hours_input_block%22%3A%7B%22add_event_hour%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2211+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-11%22%7D%7D%7D%2C%22add_event_mins_input_block%22%3A%7B%22add_event_mins%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2230%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-30%22%7D%7D%7D%7D%7D%2C%22hash%22%3A%221597168475.XN1EGhyB%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp+-+Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Afalse%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Cancel%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3Anull%2C%22root_view_id%22%3A%22V018M8X511B%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22actions%22%3A%5B%7B%22action_id%22%3A%22remove_event%22%2C%22block_id%22%3A%22ivQ%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_today_1%22%2C%22type%22%3A%22button%22%2C%22action_ts%22%3A%221597168476.485233%22%7D%5D%7D`

var submitPayload = `payload=%7B%22type%22%3A%22view_submission%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22trigger_id%22%3A%221303586025956.260884079521.a0500980fb7c927ba17ee21467d6b19e%22%2C%22view%22%3A%7B%22id%22%3A%22V01940P5D6D%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22da7yw%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22TWCbB%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22ysw%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22FIU%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%2210%3A15+asd%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22NE9%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22Lqr%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22A9H%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22Hlp%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22QEW%22%7D%2C%7B%22type%22%3A%22actions%22%2C%22block_id%22%3A%22actions_block%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22edit_events%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22edit_events%22%7D%5D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22convo_input_id%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+a+channel+to+post+the+result+on%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22conversations_select%22%2C%22action_id%22%3A%22conversation_select_action_id%22%2C%22default_to_current_conversation%22%3Atrue%2C%22response_url_enabled%22%3Atrue%2C%22initial_conversation%22%3A%22C0133NX0GQN%22%7D%7D%5D%2C%22private_metadata%22%3A%22%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22convo_input_id%22%3A%7B%22conversation_select_action_id%22%3A%7B%22type%22%3A%22conversations_select%22%2C%22selected_conversation%22%3A%22C0133NX0GQN%22%7D%7D%7D%7D%2C%22hash%22%3A%221597259257.JYHyTTUK%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Afalse%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Cancel%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Submit%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3Anull%2C%22root_view_id%22%3A%22V01940P5D6D%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22response_urls%22%3A%5B%7B%22block_id%22%3A%22convo_input_id%22%2C%22action_id%22%3A%22conversation_select_action_id%22%2C%22channel_id%22%3A%22C0133NX0GQN%22%2C%22response_url%22%3A%22https%3A%5C%2F%5C%2Fhooks.slack.com%5C%2Fapp%5C%2FT7NS02BFB%5C%2F1296028017141%5C%2FE6IdypsIaJ3iyrpWMBxX7pNt%22%7D%5D%7D`

var addEventSubmittionJson = `{
  "type": "view_submission",
  "team": {
    "id": "T7NS02BFB",
    "domain": "ford-community"
  },
  "user": {
    "id": "U7QNBA36K",
    "username": "cdorman1",
    "name": "cdorman1",
    "team_id": "T7NS02BFB"
  },
  "api_app_id": "A0131JT7VPF",
  "token": "8KTh0sVRkeZozlTxrBRqk1NO",
  "trigger_id": "1294422417044.260884079521.581460426a7ec54dcf700ed47d087735",
  "view": {
    "id": "V018GCUV2GK",
    "team_id": "T7NS02BFB",
    "type": "modal",
    "blocks": [
      {
        "type": "divider",
        "block_id": "ChVGD"
      },
      {
        "type": "context",
        "block_id": "hln",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Today's Events*",
            "verbatim": false
          }
        ]
      },
      {
        "type": "divider",
        "block_id": "NPGWv"
      },
      {
        "type": "section",
        "block_id": "ADL6a",
        "text": {
          "type": "mrkdwn",
          "text": "9:15 Standup",
          "verbatim": false
        },
        "accessory": {
          "type": "button",
          "text": {
            "type": "plain_text",
            "text": "Remove",
            "emoji": true
          },
          "value": "remove_today_1",
          "action_id": "=F5"
        }
      },
      {
        "type": "divider",
        "block_id": "LzFxD"
      },
      {
        "type": "context",
        "block_id": "mXLy",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Tomorrow's Events*",
            "verbatim": false
          }
        ]
      },
      {
        "type": "divider",
        "block_id": "bZK"
      },
      {
        "type": "section",
        "block_id": "5ybS",
        "text": {
          "type": "mrkdwn",
          "text": "_No events yet_",
          "verbatim": false
        }
      },
      {
        "type": "divider",
        "block_id": "Q3Pm"
      },
      {
        "type": "input",
        "block_id": "mOC",
        "label": {
          "type": "plain_text",
          "text": "Add Event",
          "emoji": true
        },
        "optional": false,
        "element": {
          "type": "plain_text_input",
          "action_id": "add_event",
          "placeholder": {
            "type": "plain_text",
            "text": "Title",
            "emoji": true
          }
        }
      },
      {
        "type": "input",
        "block_id": "pzhZ",
        "label": {
          "type": "plain_text",
          "text": "Day",
          "emoji": true
        },
        "optional": false,
        "element": {
          "type": "radio_buttons",
          "action_id": "today_or_tomorrow",
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
        "block_id": "89kBN",
        "label": {
          "type": "plain_text",
          "text": "Hour",
          "emoji": true
        },
        "optional": false,
        "element": {
          "type": "static_select",
          "action_id": "hours_select",
          "placeholder": {
            "type": "plain_text",
            "text": "Select hour",
            "emoji": true
          },
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
        "block_id": "zJnq",
        "label": {
          "type": "plain_text",
          "text": "Minutes",
          "emoji": true
        },
        "optional": false,
        "element": {
          "type": "static_select",
          "action_id": "mins_select",
          "placeholder": {
            "type": "plain_text",
            "text": "Select Minutes",
            "emoji": true
          },
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
    "private_metadata": "test metadata",
    "callback_id": "",
    "state": {
      "values": {
        "mOC": {
          "add_event": {
            "type": "plain_text_input",
            "value": "new test event"
          }
        },
        "pzhZ": {
          "today_or_tomorrow": {
            "type": "radio_buttons",
            "selected_option": {
              "text": {
                "type": "plain_text",
                "text": "Today",
                "emoji": true
              },
              "value": "today"
            }
          }
        },
        "89kBN": {
          "hours_select": {
            "type": "static_select",
            "selected_option": {
              "text": {
                "type": "plain_text",
                "text": "10 AM",
                "emoji": true
              },
              "value": "hour-10"
            }
          }
        },
        "zJnq": {
          "mins_select": {
            "type": "static_select",
            "selected_option": {
              "text": {
                "type": "plain_text",
                "text": "15",
                "emoji": true
              },
              "value": "min-15"
            }
          }
        }
      }
    },
    "hash": "1596810897.DFwzhtyn",
    "title": {
      "type": "plain_text",
      "text": "bZapp - Edit Events",
      "emoji": true
    },
    "clear_on_close": false,
    "notify_on_close": false,
    "close": {
      "type": "plain_text",
      "text": "Cancel",
      "emoji": true
    },
    "submit": {
      "type": "plain_text",
      "text": "Add",
      "emoji": true
    },
    "previous_view_id": null,
    "root_view_id": "V018GCUV2GK",
    "app_id": "A0131JT7VPF",
    "external_id": "",
    "app_installed_team_id": "T7NS02BFB",
    "bot_id": "B0133F8RE11"
  },
  "response_urls": []
}`

var parsedRemoveAction = `{
  "type": "block_actions",
  "user": {
    "id": "U7QNBA36K",
    "username": "cdorman1",
    "name": "cdorman1",
    "team_id": "T7NS02BFB"
  },
  "api_app_id": "A0131JT7VPF",
  "token": "8KTh0sVRkeZozlTxrBRqk1NO",
  "container": {
    "type": "view",
    "view_id": "V018Z0RV60H"
  },
  "trigger_id": "1316246398256.260884079521.dd90960e72b227fd104c50827a859bd2",
  "team": {
    "id": "T7NS02BFB",
    "domain": "ford-community"
  },
  "view": {
    "id": "V018Z0RV60H",
    "team_id": "T7NS02BFB",
    "type": "modal",
    "blocks": [
      {
        "type": "divider",
        "block_id": "S8fmI"
      },
      {
        "type": "context",
        "block_id": "QO\/H6",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Today's Events*",
            "verbatim": false
          }
        ]
      },
      {
        "type": "divider",
        "block_id": "rVpPn"
      },
      {
        "type": "section",
        "block_id": "bSKL",
        "text": {
          "type": "mrkdwn",
          "text": "10:30 asdf",
          "verbatim": false
        },
        "accessory": {
          "type": "button",
          "action_id": "remove_today_1",
          "text": {
            "type": "plain_text",
            "text": "Remove",
            "emoji": true
          },
          "value": "remove_today_1"
        }
      },
      {
        "type": "divider",
        "block_id": "=Z2C"
      },
      {
        "type": "context",
        "block_id": "T7p",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Tomorrow's Events*",
            "verbatim": false
          }
        ]
      },
      {
        "type": "divider",
        "block_id": "h+egt"
      },
      {
        "type": "section",
        "block_id": "HAp",
        "text": {
          "type": "mrkdwn",
          "text": "_No events yet_",
          "verbatim": false
        }
      },
      {
        "type": "divider",
        "block_id": "AIA"
      },
      {
        "type": "actions",
        "block_id": "actions_block",
        "elements": [
          {
            "type": "button",
            "action_id": "edit_events",
            "text": {
              "type": "plain_text",
              "text": "Edit Events",
              "emoji": true
            },
            "value": "edit_events"
          }
        ]
      }
    ],
    "private_metadata": "test metadata",
    "callback_id": "",
    "state": {
      "values": {
        "add_event_title_input_block": {
          "add_event_title": {
            "type": "plain_text_input",
            "value": "asdf"
          }
        },
        "add_event_day_input_block": {
          "add_event_day": {
            "type": "radio_buttons",
            "selected_option": {
              "text": {
                "type": "plain_text",
                "text": "Today",
                "emoji": true
              },
              "value": "today"
            }
          }
        },
        "add_event_hours_input_block": {
          "add_event_hour": {
            "type": "static_select",
            "selected_option": {
              "text": {
                "type": "plain_text",
                "text": "10 AM",
                "emoji": true
              },
              "value": "hour-10"
            }
          }
        },
        "add_event_mins_input_block": {
          "add_event_mins": {
            "type": "static_select",
            "selected_option": {
              "text": {
                "type": "plain_text",
                "text": "30",
                "emoji": true
              },
              "value": "min-30"
            }
          }
        }
      }
    },
    "hash": "1597087652.cTolAgdn",
    "title": {
      "type": "plain_text",
      "text": "bZapp",
      "emoji": true
    },
    "clear_on_close": false,
    "notify_on_close": false,
    "close": {
      "type": "plain_text",
      "text": "Cancel",
      "emoji": true
    },
    "submit": {
      "type": "plain_text",
      "text": "Submit",
      "emoji": true
    },
    "previous_view_id": null,
    "root_view_id": "V018Z0RV60H",
    "app_id": "A0131JT7VPF",
    "external_id": "",
    "app_installed_team_id": "T7NS02BFB",
    "bot_id": "B0133F8RE11"
  },
  "actions": [
    {
      "action_id": "remove_today_1",
      "block_id": "bSKL",
      "text": {
        "type": "plain_text",
        "text": "Remove",
        "emoji": true
      },
      "value": "remove_today_1",
      "type": "button",
      "action_ts": "1597087744.664898"
    }
  ]
}`

var editEvent = `{
  "type": "block_actions",
  "user": {
    "id": "U7QNBA36K",
    "username": "cdorman1",
    "name": "cdorman1",
    "team_id": "T7NS02BFB"
  },
  "api_app_id": "A0131JT7VPF",
  "token": "8KTh0sVRkeZozlTxrBRqk1NO",
  "container": {
    "type": "view",
    "view_id": "V018Z0RV60H"
  },
  "trigger_id": "1292606539763.260884079521.3246a76652e607ba32b4f5a8b6d73bf4",
  "team": {
    "id": "T7NS02BFB",
    "domain": "ford-community"
  },
  "view": {
    "id": "V018Z0RV60H",
    "team_id": "T7NS02BFB",
    "type": "modal",
    "blocks": [
      {
        "type": "divider",
        "block_id": "ZDQ97"
      },
      {
        "type": "context",
        "block_id": "7fu",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Today's Events*",
            "verbatim": false
          }
        ]
      },
      {
        "type": "divider",
        "block_id": "NJt"
      },
      {
        "type": "section",
        "block_id": "phu",
        "text": {
          "type": "mrkdwn",
          "text": "_No events yet_",
          "verbatim": false
        }
      },
      {
        "type": "divider",
        "block_id": "Ts3"
      },
      {
        "type": "context",
        "block_id": "Oat",
        "elements": [
          {
            "type": "mrkdwn",
            "text": "*Tomorrow's Events*",
            "verbatim": false
          }
        ]
      },
      {
        "type": "divider",
        "block_id": "qQQph"
      },
      {
        "type": "section",
        "block_id": "xwem2",
        "text": {
          "type": "mrkdwn",
          "text": "_No events yet_",
          "verbatim": false
        }
      },
      {
        "type": "divider",
        "block_id": "p2e\/I"
      },
      {
        "type": "actions",
        "block_id": "actions_block",
        "elements": [
          {
            "type": "button",
            "action_id": "edit_events",
            "text": {
              "type": "plain_text",
              "text": "Edit Events",
              "emoji": true
            },
            "value": "edit_events"
          }
        ]
      }
    ],
    "private_metadata": "",
    "callback_id": "",
    "state": {
      "values": {}
    },
    "hash": "1597087576.E3v2G5oS",
    "title": {
      "type": "plain_text",
      "text": "bZapp",
      "emoji": true
    },
    "clear_on_close": false,
    "notify_on_close": false,
    "close": {
      "type": "plain_text",
      "text": "Cancel",
      "emoji": true
    },
    "submit": {
      "type": "plain_text",
      "text": "Submit",
      "emoji": true
    },
    "previous_view_id": null,
    "root_view_id": "V018Z0RV60H",
    "app_id": "A0131JT7VPF",
    "external_id": "",
    "app_installed_team_id": "T7NS02BFB",
    "bot_id": "B0133F8RE11"
  },
  "actions": [
    {
      "action_id": "edit_events",
      "block_id": "actions_block",
      "text": {
        "type": "plain_text",
        "text": "Edit Events",
        "emoji": true
      },
      "value": "edit_events",
      "type": "button",
      "action_ts": "1597087578.755152"
    }
  ]
}`

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
			name: "open edit actions",
			args: args{event: events.APIGatewayProxyRequest{Body: editEventsActionButton}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url: getUrl("https://slack.com/api/views.update"),
				body: test.PrettyJsonNoError(fmt.Sprintf(
					`{
								"view_id": "V018GCUV2GK",
    							"view": %s
							}`, editEventsModal)),
			},
		},
		{
			name: "remove actions",
			args: args{event: events.APIGatewayProxyRequest{Body: removeAction}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url: getUrl("https://slack.com/api/views.update"),
				body: test.PrettyJsonNoError(fmt.Sprintf(
					`{
								"view_id": "V018M8X511B",
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
				Body: test.PrettyJsonNoError(addEventSubmissionResponse),
			},
			wantErr: false,
			wantDo: do{},
		},
	}
	for _, tt := range tests {
		gotDo = do{}
		mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
			log.Printf("url %s ", req.URL)
			body, _ := ioutil.ReadAll(req.Body)
			gotDo = do{
				url:  req.URL,
				body: test.PrettyJsonNoError(string(body)),
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
				if !assert.EqualValues(t, tt.want.Body, test.PrettyJsonNoError(got.Body)) {
					t.Errorf("Interaction() got = %v, want %v", got, tt.want)
				}
			}
			if tt.wantDo != (do{}) {
				assert.EqualValues(t, tt.wantDo.url, gotDo.url)
				assert.EqualValues(t, tt.wantDo.body, gotDo.body)
			} else {
				assert.EqualValues(t, do{}, gotDo)
			}
			//if !reflect.DeepEqual(gotDo, tt.wantDo) {
			//	t.Errorf("Interaction() gotDo = %v, wantDo %v", gotDo, tt.wantDo)
			//}
		})
	}
}

func getUrl(urlString string) *url.URL {
	result, _ := url.Parse(urlString)
	return result
}

var editEventsModal = `{
	"title": {
		"type": "plain_text",
		"text": "bZapp - Edit Events",
		"emoji": true
	},
	"submit": {
		"type": "plain_text",
		"text": "Add",
		"emoji": true
	},
	"type": "modal",
	"close": {
		"type": "plain_text",
		"text": "Cancel",
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
			"block_id": "add_event_title_input_block",
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
			"block_id": "add_event_day_input_block",
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
			"block_id": "add_event_hours_input_block",
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
			"block_id": "add_event_mins_input_block",
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

var removeEventsModal = `{
	"title": {
		"type": "plain_text",
		"text": "bZapp - Edit Events",
		"emoji": true
	},
	"submit": {
		"type": "plain_text",
		"text": "Add",
		"emoji": true
	},
	"type": "modal",
	"close": {
		"type": "plain_text",
		"text": "Cancel",
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
"block_id": "4MWhf",
			"text": {
				"type": "mrkdwn",
				"text": "11:15asdf"
			},
			"accessory": {
				"action_id": "remove_event",
				"type": "button",
				"text": {
					"type": "plain_text",
					"text": "Remove",
					"emoji": true
				},
				"value": "remove_today_0"
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
				"text": "10:15asdf"
			},
			"accessory": {
				"action_id": "remove_event",
				"type": "button",
				"value": "remove_tomorrow_0",
				"text": {
					"type": "plain_text",
					"text": "Remove",
					"emoji": true
				}
			},
			"block_id": "bpxH"
		},
		{
			"type": "divider"
		},
		{
			"type": "input",
			"block_id": "add_event_title_input_block",
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
			"block_id": "add_event_day_input_block",
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
			"block_id": "add_event_hours_input_block",
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
			"block_id": "add_event_mins_input_block",
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

const addEventSubmissionResponse = `{
	"response_action": "update",
	"view": {
		"type": "modal",
		"title": {
			"type": "plain_text",
			"text": "bZapp",
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
					"text": "4:45 retrob"
				}
			},
			{
				"type": "divider"
			},
			{
				"type": "actions",
				"block_id": "actions_block",
				"elements": [
					{
						"type": "button",
						"text": {
							"type": "plain_text",
							"text": "Edit Events",
							"emoji": true
						},
						"action_id": "edit_events",
						"value": "edit_events"
					}
				]
			}
		],
		"close": {
			"type": "plain_text",
			"text": "Cancel",
			"emoji": true
		},
		"submit": {
			"type": "plain_text",
			"text": "Submit",
			"emoji": true
		}
	}
}`