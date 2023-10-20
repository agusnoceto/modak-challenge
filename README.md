# Modak challenge
This repository contains the proposal solution for Modak's Code Challenge described below.

# Rate-Limited Notification Service
## The Task

We have a Notification system that sends out email notifications of various types (supdatesupdate, daily news, project invitations, etc). We need to protect recipients from getting too many emails, either due to system errors or due to abuse, so let's limit the number of emails sent to them by implementing a rate-limited version of NotificationService.

The system must reject requests that are over the limit.

Some sample notification types and rate limit rules, e.g.:

Status: not more than 2 per minute for each recipient

News: not more than 1 per day for each recipient

Marketing: not more than 3 per hour for each recipient

Etc. these are just samples, the system might have several rate limit rules!

# How to Run it

Please run `make help` for information on how to run this app.

## The Rules

The rules are configurable, and presented in a json file in the main directory of the app, named `rules.json`. If you want to add/modify/remove any  rule, please follow the format in said json, taking into account that the acceptable interval values are the following:
- "ns": Nanosecond
- "us": Microsecond
- "µs": Microsecond
- "μs": Microsecond
- "ms": Millisecond
- "s":  Second
- "m":  Minute
- "h":  Hour

