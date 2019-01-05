#!/bin/bash

TIME="10"
URL="https://api.telegram.org/bot $TELEGRAM_TOKEN / sendMessage"
TEXT="Deploy the status: $1 %0A% 0AProject: + $CI_PROJECT_NAME % 0AURL: + $ CI_PROJECT_URL / Pipelines / $CI_PIPELINE_ID /% 0ABranch: + $CI_COMMIT_REF_SLUG "

curl -s --max-time $TIME -d "chat_id = $TELEGRAM_CHATID & disable_web_page_preview = 1 & text = $TEXT "  $ URL > /dev/null