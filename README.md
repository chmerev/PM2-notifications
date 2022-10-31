# PM2 Notifications
This is a small program that will allow you to monitor the activity of processes in PM2 and if the process is not found in the "Online" status, the program will send a notification to Telegram.

**Compile the file**
`env GOOS=linux GOARCH=amd64 go build -o pm2-notification -v cmd/main.go`

**Add the file to the server and create a .sh file**
```bash
#!/bin/bash
/monitors/pm2-notification \
-names name1,name2,name3 \
-tg  your telegram bot key \
-chat telegram id chat, where you need to send notifications
```

**Add a job to CRON**
`*/5 * * * * root /monitors/pm2-notification.sh`

Now we can monitor PM2 processes and if the process becomes unavailable you will be notified