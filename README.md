wearhacksline
========

1) Send SMS to specific channel via Twilio. The purpose was of this project was to allow the hackathon attendees to notify the organizing team of anything that was making them feel uncomfortable. 

2) Reply back to number from which the notifiication was received.

## Usage

1) To notify the organizing team on Slack via SMS, just send a message to the Twilio number provided

    
2) To reply back to the notifier:

    /sms <number_to_text> <message body>


## Installation

### Setup your own server

Make sure to change the **Slash Command** URL to whatever your URL is.

##### Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy?template=https://github.com/ueg1990/wearhacksline/tree/master)

And then:

```bash
$ heroku config:set SLACK_WEARHACKS_WEBHOOK_URL=<URL>
$ heroku config:set TWILIO_ACCOUNT_SID=<ACCOUNT_SID> 
$ heroku config:set TWILIO_AUTH_TOKEN=<AUTH_TOKEN>
$ heroku config:set TWILIO_NUMBER=<NUMBER>
$ heroku config:set PORT=<PORT>
```

### Setup Integration

- Go to your channel
- Click on **Configure Integrations**.
- Scroll all the way down to **DIY Integrations & Customizations** section.

#### Add a new slash command with the following settings:

- Click on **Add** next to **Slash Commands**.

  - Command: `/sms`
  - URL: `http://YOUR-URL.com/sms`
  - Method: `POST`

  ![](http://i.imgur.com/zLrHkf5.png)

All other settings can be set on your own discretion.

#### Set up a new incoming webhook

Click on **Add** next to **Incoming WebHooks**.

  - Choose a channel to integrate with (this doesn't matter -- it'll always respond to the channel you called it from)
  - Note the new Webhook URL.

  ![](http://i.imgur.com/tgiTLdj.png)
  
### Setup Twilio

- Create a Twilio Account
- Go to your [Twilio Account](https://www.twilio.com/user/account/settings) to retrieve the Twilio Account SID and Auth Token associated with your account
- Update the Messaging Request URL to your URL with route `/twiml`:
    
    ![](http://i.imgur.com/Mkf7HGa.png)

## Contributing

- Please use the [issue tracker]() to report any bugs or file feature requests.

- PRs to add new sources are welcome.
