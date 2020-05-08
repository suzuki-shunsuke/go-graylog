resource "graylog_event_notification" "legacy-http-alarm-callback" {
  title       = "legacy-http-alarm-callback"
  description = "Migrated legacy alarm callback"

  config = jsonencode({
    type          = "legacy-alarm-callback-notification-v1"
    callback_type = "org.graylog2.alarmcallbacks.HTTPAlarmCallback"
    configuration = {
      url = "https://example.com"
    }
  })
}

resource "graylog_event_notification" "legacy-email-alarm-callback" {
  title       = "legacy-email-alarm-callback"
  description = "Migrated legacy alarm callback"

  config = jsonencode({
    type          = "legacy-alarm-callback-notification-v1"
    callback_type = "org.graylog2.alarmcallbacks.EmailAlarmCallback"
    configuration = {
      sender  = "graylog@example.org"
      subject = "Graylog alert for stream "
      user_receivers = [
        "username"
      ]
      body = "hello world"
      email_receivers = [
        "graylog@example.com"
      ]
    }
  })
}

resource "graylog_event_notification" "legacy-slack-alarm-callback" {
  title       = "legacy-slack-alarm-callback"
  description = "Migrated legacy alarm callback"

  config = jsonencode({
    type          = "legacy-alarm-callback-notification-v1"
    callback_type = "org.graylog2.plugins.slack.callback.SlackAlarmCallback"
    configuration = {
      graylog2_url   = "https://graylog.example.com"
      notify_channel = false
      link_names     = true
      custom_message = "hello world"
      webhook_url    = "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
      color          = "#FF0000"
      backlog_items  = 5
      user_name      = "Graylog"
      channel        = "#general"
    }
  })
}

resource "graylog_event_notification" "http" {
  title       = "http"
  description = "http notification"

  config = jsonencode({
    type = "http-notification-v1"
    url  = "http://example.com"
  })
}

resource "graylog_event_notification" "email" {
  title       = "email"
  description = "email notification"

  config = jsonencode({
    type          = "email-notification-v1"
    sender        = "graylog@example.com"
    subject       = "Graylog event notification"
    body_template = "hello"
    email_recipients = [
      "graylog@example.com"
    ]
    user_recipients = [
      "admin"
    ]
  })
}
