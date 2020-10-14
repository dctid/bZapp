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

const addEventSubmission = `payload=%7B%22type%22%3A%22view_submission%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22trigger_id%22%3A%221296901702898.260884079521.664b4e09d83d7c1e1bb094385e9b49a3%22%2C%22view%22%3A%7B%22id%22%3A%22V018X2J9UA0%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22wrQ%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22r1Bj6%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22iMR%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%228HQ4g%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22T2I%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22Bdcei%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22nxzS%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22YtXZ%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22MJYT%22%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_title_input_block-1%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add+Event%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22action_id%22%3A%22add_event_title%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Title%22%2C%22emoji%22%3Atrue%7D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_day_input_block-1%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Day%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22action_id%22%3A%22add_event_day%22%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Today%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22today%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Tomorrow%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22tomorrow%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_hours_input_block-1%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Hour%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_hour%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+hour%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%229+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-9%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2210+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-10%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2211+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-11%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2212+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-12%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%221+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-1%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%222+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-2%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%223+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-3%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%224+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-4%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_mins_input_block-1%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_mins%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2200%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-0%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2215%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-15%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2230%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-30%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2245%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-45%22%7D%5D%7D%7D%5D%2C%22private_metadata%22%3A%22%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22add_event_title_input_block-1%22%3A%7B%22add_event_title%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22value%22%3A%22retrob%22%7D%7D%2C%22add_event_day_input_block-1%22%3A%7B%22add_event_day%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Tomorrow%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22tomorrow%22%7D%7D%7D%2C%22add_event_hours_input_block-1%22%3A%7B%22add_event_hour%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%224+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-4%22%7D%7D%7D%2C%22add_event_mins_input_block-1%22%3A%7B%22add_event_mins%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2245%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-45%22%7D%7D%7D%7D%7D%2C%22hash%22%3A%221597247878.AteqceCB%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp+-+Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Afalse%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Cancel%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3Anull%2C%22root_view_id%22%3A%22V018X2J9UA0%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22response_urls%22%3A%5B%5D%7D`

const removeEventAction = `payload=%7B%22type%22%3A%22block_actions%22%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22container%22%3A%7B%22type%22%3A%22view%22%2C%22view_id%22%3A%22V0190BP4VR9%22%7D%2C%22trigger_id%22%3A%221318817490673.260884079521.a444f1a8fbf298fa8e96049d3f552aa0%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22view%22%3A%7B%22id%22%3A%22V0190BP4VR9%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22fRK%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22laXE%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22me%3De3%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%2289Q4%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%2210%3A15+dsfg%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_event%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_today_0%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22uYw%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22eiYz%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22QCSi%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22Zcj%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22vH%5C%2F%22%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_title_input_block-9%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add+Event%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22action_id%22%3A%22add_event_title%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Title%22%2C%22emoji%22%3Atrue%7D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_day_input_block-9%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Day%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22action_id%22%3A%22add_event_day%22%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Today%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22today%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Tomorrow%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22tomorrow%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_hours_input_block-9%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Hour%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_hour%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+hour%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%229+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-9%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2210+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-10%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2211+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-11%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2212+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-12%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%221+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-1%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%222+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-2%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%223+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-3%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%224+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-4%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_mins_input_block-9%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_mins%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2200%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-0%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2215%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-15%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2230%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-30%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2245%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-45%22%7D%5D%7D%7D%5D%2C%22private_metadata%22%3A%22%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22add_event_title_input_block-0%22%3A%7B%22add_event_title%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22value%22%3A%22dsfg%22%7D%7D%2C%22add_event_day_input_block-0%22%3A%7B%22add_event_day%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Today%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22today%22%7D%7D%7D%2C%22add_event_hours_input_block-0%22%3A%7B%22add_event_hour%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2210+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-10%22%7D%7D%7D%2C%22add_event_mins_input_block-0%22%3A%7B%22add_event_mins%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2215%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-15%22%7D%7D%7D%7D%7D%2C%22hash%22%3A%221597700322.rf6rGLjZ%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp+-+Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Atrue%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Back%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3A%22V019CQ0MQSV%22%2C%22root_view_id%22%3A%22V019CQ0MQSV%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22actions%22%3A%5B%7B%22action_id%22%3A%22remove_event%22%2C%22block_id%22%3A%2289Q4%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_today_0%22%2C%22type%22%3A%22button%22%2C%22action_ts%22%3A%221597700329.377454%22%7D%5D%7D`

const submitPayload = `payload=%7B%22type%22%3A%22view_submission%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22trigger_id%22%3A%221321168290897.260884079521.8654a183f68d6c418fc4bae6d12229a7%22%2C%22view%22%3A%7B%22id%22%3A%22V0192F54PSN%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22jLlGw%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22zhX%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22dH3qZ%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22FpV6%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%2210%3A45+gsfd%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%229GK1%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22Tf49%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22hRmAB%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22xWAF3%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%224%3A30+kljh%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22eiyfo%22%7D%2C%7B%22type%22%3A%22actions%22%2C%22block_id%22%3A%22actions_block%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22edit_events%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22edit_events%22%7D%5D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22convo_input_id%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+a+channel+to+post+the+result+on%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22conversations_select%22%2C%22action_id%22%3A%22conversation_select_action_id%22%2C%22default_to_current_conversation%22%3Atrue%2C%22response_url_enabled%22%3Atrue%2C%22initial_conversation%22%3A%22C0133NX0GQN%22%7D%7D%5D%2C%22private_metadata%22%3A%22%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22convo_input_id%22%3A%7B%22conversation_select_action_id%22%3A%7B%22type%22%3A%22conversations_select%22%2C%22selected_conversation%22%3A%22C0133NX0GQN%22%7D%7D%7D%7D%2C%22hash%22%3A%221597780066.zmloiCmP%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Afalse%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Cancel%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Submit%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3Anull%2C%22root_view_id%22%3A%22V0192F54PSN%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22response_urls%22%3A%5B%7B%22block_id%22%3A%22convo_input_id%22%2C%22action_id%22%3A%22conversation_select_action_id%22%2C%22channel_id%22%3A%22C0133NX0GQN%22%2C%22response_url%22%3A%22https%3A%5C%2F%5C%2Fhooks.slack.com%5C%2Fapp%5C%2FT7NS02BFB%5C%2F1308748543923%5C%2FIQu4PNxJQeofD8m8RucVb5d3%22%7D%5D%7D`

const closeEditEvents = `payload=%7B%22type%22%3A%22view_closed%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22view%22%3A%7B%22id%22%3A%22V0198LMG1JQ%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22WF9v6%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22rgGj%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22nRj8%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22qKG%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%2210%3A45+gsfd%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_event%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_today_0%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22mKgM%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22I%5C%2FRXh%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22Vtku%2B%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22F2vL%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22KVeN%22%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_title_input_block-2%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add+Event%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22action_id%22%3A%22add_event_title%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Title%22%2C%22emoji%22%3Atrue%7D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_day_input_block-2%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Day%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22action_id%22%3A%22add_event_day%22%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Today%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22today%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Tomorrow%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22tomorrow%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_hours_input_block-2%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Hour%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_hour%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+hour%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%229+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-9%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2210+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-10%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2211+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-11%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2212+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-12%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%221+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-1%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%222+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-2%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%223+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-3%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%224+PM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-4%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_event_mins_input_block-2%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_event_mins%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+Minutes%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2200%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-0%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2215%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-15%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2230%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-30%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2245%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-45%22%7D%5D%7D%7D%5D%2C%22private_metadata%22%3A%22%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22add_event_title_input_block-1%22%3A%7B%22add_event_title%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22value%22%3A%22gsfd%22%7D%7D%2C%22add_event_day_input_block-1%22%3A%7B%22add_event_day%22%3A%7B%22type%22%3A%22radio_buttons%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Today%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22today%22%7D%7D%7D%2C%22add_event_hours_input_block-1%22%3A%7B%22add_event_hour%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2210+AM%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22hour-10%22%7D%7D%7D%2C%22add_event_mins_input_block-1%22%3A%7B%22add_event_mins%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%2245%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22min-45%22%7D%7D%7D%7D%7D%2C%22hash%22%3A%221597779374.8S8jehyM%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp+-+Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Atrue%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Back%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3A%22V0192F54PSN%22%2C%22root_view_id%22%3A%22V0192F54PSN%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22is_cleared%22%3Afalse%7D`

const editGoalsActionButton = `payload=%7B%22type%22%3A%22block_actions%22%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22container%22%3A%7B%22type%22%3A%22view%22%2C%22view_id%22%3A%22V01CQ79HN5S%22%7D%2C%22trigger_id%22%3A%221411346195543.260884079521.14fdd4f0ec90fe20a07ea8dc9429d891%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22view%22%3A%7B%22id%22%3A%22V01CQ79HN5S%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22uv7%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22aJO%5C%2Fu%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AToday%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22TNF%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%220x%2BuJ%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22%5C%2F%3DGFQ%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22rn%5C%2Fv%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATomorrow%27s+Events%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22v%5C%2Fk9X%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22IT%5C%2F%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+events+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22Y4N9%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22287%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22Goals%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22tjv5u%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%222dHPr%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22Yntk%22%7D%2C%7B%22type%22%3A%22actions%22%2C%22block_id%22%3A%22actions_block%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22edit_events%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Edit+Events%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22edit_events%22%7D%2C%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22edit_goals%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Edit+Goals%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22edit_goals%22%7D%5D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22convo_input_id%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Select+a+channel+to+post+the+result+on%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22conversations_select%22%2C%22action_id%22%3A%22conversation_select_action_id%22%2C%22default_to_current_conversation%22%3Atrue%2C%22response_url_enabled%22%3Atrue%2C%22initial_conversation%22%3A%22C0133NX0GQN%22%7D%7D%5D%2C%22private_metadata%22%3A%22%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22convo_input_id%22%3A%7B%22conversation_select_action_id%22%3A%7B%22type%22%3A%22conversations_select%22%2C%22selected_conversation%22%3A%22C0133NX0GQN%22%7D%7D%7D%7D%2C%22hash%22%3A%221602618823.sl68Li82%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Afalse%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Cancel%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Submit%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3Anull%2C%22root_view_id%22%3A%22V01CQ79HN5S%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22actions%22%3A%5B%7B%22action_id%22%3A%22edit_goals%22%2C%22block_id%22%3A%22actions_block%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Edit+Goals%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22edit_goals%22%2C%22type%22%3A%22button%22%2C%22action_ts%22%3A%221602618829.770185%22%7D%5D%7D`

const addGoalSubmission = `payload=%7B%22type%22%3A%22view_submission%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22trigger_id%22%3A%221416625528774.260884079521.682e20059f237a92f0391dcba6a064cb%22%2C%22view%22%3A%7B%22id%22%3A%22V01CFG53H99%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22gzF%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22A6T2X%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ACustomer+Questions%3F%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%224WtFP%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22FzT%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22wCg%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%229eI%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATeam+Needs%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22zzoeZ%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22wvDnw%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%229qw%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22V8z%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ALearnings%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22USX9G%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22OmVn8%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%220cKt%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22h3X4f%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AQuestions%3F%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22Y48F%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%222%3DS47%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%226FB%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22wCU%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AOther%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%224Eo%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22Wn4%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22wrz6%22%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_goal_category_input_block-4%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal+to+Add+to%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_goal_category%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Choose+Goal%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Customer+Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Customer+Questions%3F%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Team+Needs%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Team+Needs%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Learnings%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Learnings%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Questions%3F%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Other%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Other%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_goal_input_block-4%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal+to+Add%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22action_id%22%3A%22add_goal%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal%22%2C%22emoji%22%3Atrue%7D%2C%22dispatch_action_config%22%3A%7B%22trigger_actions_on%22%3A%5B%22on_enter_pressed%22%5D%7D%7D%7D%5D%2C%22private_metadata%22%3A%22%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22add_goal_category_input_block-4%22%3A%7B%22add_goal_category%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Customer+Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Customer+Questions%3F%22%7D%7D%7D%2C%22add_goal_input_block-4%22%3A%7B%22add_goal%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22value%22%3A%22adsfads%22%7D%7D%7D%7D%2C%22hash%22%3A%221602533827.7T3dy93R%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp+-+Edit+Goals%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Atrue%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Back%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3A%22V01C0HPLSEB%22%2C%22root_view_id%22%3A%22V01C0HPLSEB%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22response_urls%22%3A%5B%5D%7D`

const add2ndGoalSubmission = `payload=%7B%22type%22%3A%22view_submission%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22trigger_id%22%3A%221410528021047.260884079521.2dd5693d02f4b233d7c5b345ff489f5c%22%2C%22view%22%3A%7B%22id%22%3A%22V01CE58KMTP%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22%3DIPPq%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22axP3%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ACustomer+Questions%3F%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22QqhF%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22CfOUbJ%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22adsf%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_goal%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_Customer+Questions%3F_CfOUbJ%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22A6Kk%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22x0oY9%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATeam+Needs%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22GrQL%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22wABJ%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%226AAm%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22GVuWC%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ALearnings%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22Jap%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22w3eYj%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22IXor%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22f9j%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AQuestions%3F%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22qkaE6%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22zCw%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22twZ%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22%3DrUq%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AOther%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%220oBE2%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22v8SlY%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22Rn%5C%2Fp%22%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_goal_category_input_block-2%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal+to+Add+to%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_goal_category%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Choose+Goal%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Customer+Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Customer+Questions%3F%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Team+Needs%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Team+Needs%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Learnings%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Learnings%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Questions%3F%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Other%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Other%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_goal_input_block-2%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal+to+Add%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22action_id%22%3A%22add_goal%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal%22%2C%22emoji%22%3Atrue%7D%2C%22dispatch_action_config%22%3A%7B%22trigger_actions_on%22%3A%5B%22on_enter_pressed%22%5D%7D%7D%7D%5D%2C%22private_metadata%22%3A%22%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22add_goal_category_input_block-2%22%3A%7B%22add_goal_category%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3A%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Customer+Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Customer+Questions%3F%22%7D%7D%7D%2C%22add_goal_input_block-2%22%3A%7B%22add_goal%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22value%22%3A%22hlkj%22%7D%7D%7D%7D%2C%22hash%22%3A%221602601527.0KQGWgbM%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp+-+Edit+Goals%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Atrue%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Back%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3A%22V01CE58EDEH%22%2C%22root_view_id%22%3A%22V01CE58EDEH%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22response_urls%22%3A%5B%5D%7D`

const removeGoalAction = `payload=%7B%22type%22%3A%22block_actions%22%2C%22user%22%3A%7B%22id%22%3A%22U7QNBA36K%22%2C%22username%22%3A%22cdorman1%22%2C%22name%22%3A%22cdorman1%22%2C%22team_id%22%3A%22T7NS02BFB%22%7D%2C%22api_app_id%22%3A%22A0131JT7VPF%22%2C%22token%22%3A%228KTh0sVRkeZozlTxrBRqk1NO%22%2C%22container%22%3A%7B%22type%22%3A%22view%22%2C%22view_id%22%3A%22V01CS7KGBE0%22%7D%2C%22trigger_id%22%3A%221421365539206.260884079521.b3c71db5dfca1d3404eb774a919f7e35%22%2C%22team%22%3A%7B%22id%22%3A%22T7NS02BFB%22%2C%22domain%22%3A%22ford-community%22%7D%2C%22view%22%3A%7B%22id%22%3A%22V01CS7KGBE0%22%2C%22team_id%22%3A%22T7NS02BFB%22%2C%22type%22%3A%22modal%22%2C%22blocks%22%3A%5B%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22kiJ5q%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22IqbI5%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ACustomer+Questions%3F%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22GlD%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22oiFbDZ%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22asdf%22%2C%22verbatim%22%3Afalse%7D%2C%22accessory%22%3A%7B%22type%22%3A%22button%22%2C%22action_id%22%3A%22remove_goal%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_Customer+Questions%3F_oiFbDZ%22%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22YoMqB%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22UnyN%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ATeam+Needs%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22gjB%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%228nfie%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22pp9%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22Acx%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2ALearnings%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22s%5C%2Fea%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%228xlW%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22F6J%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22fmM2%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AQuestions%3F%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22XwplK%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22d5%2Bd1%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22OOD%22%7D%2C%7B%22type%22%3A%22context%22%2C%22block_id%22%3A%22RRbT%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22%2AOther%2A%22%2C%22verbatim%22%3Afalse%7D%5D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%221dhF%22%7D%2C%7B%22type%22%3A%22section%22%2C%22block_id%22%3A%22UuI%3DY%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22_No+goals+yet_%22%2C%22verbatim%22%3Afalse%7D%7D%2C%7B%22type%22%3A%22divider%22%2C%22block_id%22%3A%22vHIBr%22%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_goal_category_input_block-2%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal+to+Add+to%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22static_select%22%2C%22action_id%22%3A%22add_goal_category%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Choose+Goal%22%2C%22emoji%22%3Atrue%7D%2C%22options%22%3A%5B%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Customer+Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Customer+Questions%3F%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Team+Needs%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Team+Needs%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Learnings%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Learnings%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Questions%3F%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Questions%3F%22%7D%2C%7B%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Other%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22goal-Other%22%7D%5D%7D%7D%2C%7B%22type%22%3A%22input%22%2C%22block_id%22%3A%22add_goal_input_block-2%22%2C%22label%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal+to+Add%22%2C%22emoji%22%3Atrue%7D%2C%22optional%22%3Afalse%2C%22dispatch_action%22%3Afalse%2C%22element%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22action_id%22%3A%22add_goal%22%2C%22placeholder%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Goal%22%2C%22emoji%22%3Atrue%7D%2C%22dispatch_action_config%22%3A%7B%22trigger_actions_on%22%3A%5B%22on_enter_pressed%22%5D%7D%7D%7D%5D%2C%22private_metadata%22%3A%22%22%2C%22callback_id%22%3A%22%22%2C%22state%22%3A%7B%22values%22%3A%7B%22add_goal_category_input_block-2%22%3A%7B%22add_goal_category%22%3A%7B%22type%22%3A%22static_select%22%2C%22selected_option%22%3Anull%7D%7D%2C%22add_goal_input_block-2%22%3A%7B%22add_goal%22%3A%7B%22type%22%3A%22plain_text_input%22%2C%22value%22%3Anull%7D%7D%7D%7D%2C%22hash%22%3A%221602686861.waneDoPu%22%2C%22title%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22bZapp+-+Edit+Goals%22%2C%22emoji%22%3Atrue%7D%2C%22clear_on_close%22%3Afalse%2C%22notify_on_close%22%3Atrue%2C%22close%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Back%22%2C%22emoji%22%3Atrue%7D%2C%22submit%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Add%22%2C%22emoji%22%3Atrue%7D%2C%22previous_view_id%22%3A%22V01CS7KAJQL%22%2C%22root_view_id%22%3A%22V01CS7KAJQL%22%2C%22app_id%22%3A%22A0131JT7VPF%22%2C%22external_id%22%3A%22%22%2C%22app_installed_team_id%22%3A%22T7NS02BFB%22%2C%22bot_id%22%3A%22B0133F8RE11%22%7D%2C%22actions%22%3A%5B%7B%22action_id%22%3A%22remove_goal%22%2C%22block_id%22%3A%22oiFbDZ%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22text%22%3A%22Remove%22%2C%22emoji%22%3Atrue%7D%2C%22value%22%3A%22remove_Customer+Questions%3F_oiFbDZ%22%2C%22type%22%3A%22button%22%2C%22action_ts%22%3A%221602686863.223793%22%7D%5D%7D`

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
				url: getUrl("https://slack.com/api/views.push"),
				body: format.PrettyJsonNoError(fmt.Sprintf(
					`{
								"trigger_id": "1288231154914.260884079521.ba1595ee20fab577e5ac042a518713fd",
    							"view": %s
							}`, editEventsModal)),
			},
		},
		{
			name: "remove actions",
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
								"view_id": "V0190BP4VR9",
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
				url:  getUrl("https://hooks.slack.com/app/T7NS02BFB/1308748543923/IQu4PNxJQeofD8m8RucVb5d3"),
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
								"view_id": "V01CS7KGBE0",
    							"view": %s
							}`, removeGoalsModal)),
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
	"private_metadata": "{\"Index\":1,\"Events\":{\"TodaysEvents\":[],\"TomorrowsEvents\":[]},\"Goals\":{\"CustomerQuestions?\":[],\"Learnings\":[],\"Other\":[],\"Questions?\":[],\"TeamNeeds\":[]}}",
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
        "block_id": "add_event_title_input_block-9",
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
        "block_id": "add_event_day_input_block-9",
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
        "block_id": "add_event_hours_input_block-9",
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
        "block_id": "add_event_mins_input_block-9",
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
	"private_metadata": "{\"Index\":9,\"Events\":{\"TodaysEvents\":[],\"TomorrowsEvents\":[]},\"Goals\":{\"CustomerQuestions?\":[],\"Learnings\":[],\"Other\":[],\"Questions?\":[],\"TeamNeeds\":[]}}",
    "notify_on_close": true
  }
`

const addEventSubmissionResponse = `{
	"response_action": "update",
	"view": {
		"type": "modal",
		"title": {
			"type": "plain_text",
			"text": "bZapp-EditEvents",
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
				"block_id": "Fakehash",
				"text": {
					"type": "mrkdwn",
					"text": "4:45 retrob"
				},
				"accessory": {
					"action_id": "remove_event",
					"text": {
						"emoji": true,
						"text": "Remove",
						"type": "plain_text"
					},
					"type": "button",
					"value": "remove_tomorrow_Fakehash"
				}
			},
			{
				"type": "divider"
			},
      {
        "type": "input",
        "block_id": "add_event_title_input_block-2",
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
        "block_id": "add_event_day_input_block-2",
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
        "block_id": "add_event_hours_input_block-2",
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
        "block_id": "add_event_mins_input_block-2",
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
		"notify_on_close": true,
		"private_metadata": "{\"Index\":2,\"Events\":{\"TodaysEvents\":[],\"TomorrowsEvents\":[{\"Id\":\"Fakehash\",\"Title\":\"retrob\",\"Day\":\"tomorrow\",\"Hour\":4,\"Min\":45,\"AmPm\":\"PM\"}]},\"Goals\":{\"CustomerQuestions?\":[],\"Learnings\":[],\"Other\":[],\"Questions?\":[],\"TeamNeeds\":[]}}",
		"submit": {
			"type": "plain_text",
			"text": "Add",
			"emoji": true
		}
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
			"block_id": "FpV6",
			"text": {
				"text": "10:45gsfd",
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
			"block_id": "xWAF3",
			"text": {
				"text": "4:30kljh",
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
  "view_id": "V0192F54PSN",
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
        "block_id": "qKG",
        "text": {
          "text": "10:45gsfd",
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
        "text": {
          "text": "_Noeventsyet_",
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
    "private_metadata": "{\"Index\":0,\"Events\":{\"TodaysEvents\":[{\"Id\":\"qKG\",\"Title\":\"gsfd\",\"Day\":\"today\",\"Hour\":10,\"Min\":45,\"AmPm\":\"AM\"}],\"TomorrowsEvents\":[]},\"Goals\":{\"CustomerQuestions?\":[],\"Learnings\":[],\"Other\":[],\"Questions?\":[],\"TeamNeeds\":[]}}",
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
  }
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
		"private_metadata": "{\"Index\":1,\"Events\":{\"TodaysEvents\":[],\"TomorrowsEvents\":[]},\"Goals\":{\"CustomerQuestions?\":[],\"Learnings\":[],\"Other\":[],\"Questions?\":[],\"TeamNeeds\":[]}}",
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
          "text": "adsfads"
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
          "value": "remove_Customer Questions?_Fakehash"
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
        "block_id": "add_goal_category_input_block-5",
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
        "block_id": "add_goal_input_block-5",
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
    "notify_on_close": true,
	"private_metadata": "{\"Index\":5,\"Events\":{\"TodaysEvents\":[],\"TomorrowsEvents\":[]},\"Goals\":{\"CustomerQuestions?\":[{\"Id\":\"Fakehash\",\"Value\":\"adsfads\"}],\"Learnings\":[],\"Other\":[],\"Questions?\":[],\"TeamNeeds\":[]}}"
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
          "text": "adsf"
        },
        "block_id": "CfOUbJ",
        "accessory": {
          "type": "button",
          "text": {
            "type": "plain_text",
            "text": "Remove",
            "emoji": true
          },
          "action_id": "remove_goal",
          "value": "remove_CustomerQuestions?_CfOUbJ"
        }
      },
      {
        "type": "section",
        "text": {
          "type": "mrkdwn",
          "text": "hlkj"
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
          "value": "remove_CustomerQuestions?_Fakehash"
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
        "block_id": "add_goal_category_input_block-3",
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
        "block_id": "add_goal_input_block-3",
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
    "notify_on_close": true,
	"private_metadata": "{\"Index\":3,\"Events\":{\"TodaysEvents\":[],\"TomorrowsEvents\":[]},\"Goals\":{\"CustomerQuestions?\":[{\"Id\":\"CfOUbJ\",\"Value\":\"adsf\"},{\"Id\":\"Fakehash\",\"Value\":\"hlkj\"}],\"Learnings\":[],\"Other\":[],\"Questions?\":[],\"TeamNeeds\":[]}}"
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
			"block_id": "add_goal_category_input_block-2",
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
			"block_id": "add_goal_input_block-2",
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
	"notify_on_close": true,
	"private_metadata": "{\"Index\":2,\"Events\":{\"TodaysEvents\":[],\"TomorrowsEvents\":[]},\"Goals\":{\"CustomerQuestions?\":[],\"Learnings\":[],\"Other\":[],\"Questions?\":[],\"TeamNeeds\":[]}}"
}`
