<div align="center">
  
<img src="./.github/assets/images/logo.svg" width="200px" alt="Timeful logo" />

</div>
<br />
<div align="center">

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL_v3-orange.svg)](https://www.gnu.org/licenses/agpl-3.0)
[![Donate](https://img.shields.io/badge/-Donate%20with%20Paypal-blue?logo=paypal)](https://www.paypal.com/donate/?hosted_button_id=KWCH6LGJCP6E6)
[![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/timeful_app?label=%40timeful_app&labelColor=white)](https://x.com/timeful_app)
[![Discord](https://img.shields.io/badge/-Join%20Discord-7289DA?logo=discord&logoColor=white)](https://discord.gg/v6raNqYxx3)
[![Subreddit subscribers](https://img.shields.io/reddit/subreddit-subscribers/schej?label=join%20r%2Fschej)](https://www.reddit.com/r/schej/)

</div>

<img src="./.github/assets/images/hero.jpg" alt="Timeful hero" />

Timeful is a scheduling platform helps you find the best time for a group to meet. It is a free availability poll that is easy to use and integrates with your calendar.

Hosted version of the site: https://timeful.app

Built with [Vue 2](https://github.com/vuejs/vue), [MongoDB](https://github.com/mongodb/mongo), [Go](https://github.com/golang/go), and [TailwindCSS](https://github.com/tailwindlabs/tailwindcss)

## Demo

[![demo video](http://markdown-videos-api.jorgenkh.no/youtube/vFkBC8BrkOk)](https://www.youtube.com/watch?v=vFkBC8BrkOk)

## Features

- See when everybody's availability overlaps
- Easily specify date + time ranges to meet between
- Google calendar, Outlook, Apple calendar integration
- "Available" vs. "If needed" times
- Determine when a subset of people are available
- Schedule across different time zones
- Email notifications + reminders
- Duplicating polls
- Availability groups - stay up to date with people's real-time calendar availability
- Export availability as CSV
- Only show responses to event creator

## Self-hosting

1. **Prerequisites**
   - Node.js 16+ and npm
   - Go 1.20+
   - MongoDB running and accessible
2. **Environment**
   - Copy `.env.example` to `.env` in `/server` and fill values (Mongo URI, JWT secret, email/stripe keys as needed).
   - For the frontend, create `frontend/.env.local` with any `VUE_APP_*` overrides you need (e.g., API base URL).
3. **Install**
   - Frontend: `cd frontend && npm install`
   - Backend: `cd server && go mod download`
4. **Run locally**
   - Backend: `cd server && go run main.go`
   - Frontend: `cd frontend && npm run serve` (defaults to http://localhost:8080)
5. **Build for production**
   - Frontend: `npm run build` (outputs to `frontend/dist`)
   - Serve built assets via your preferred web server; point API calls to the running Go server.
6. **Updating**
   - Pull the latest changes, then re-run `npm install` (frontend) and `go mod download` (server) if dependencies changed.
   - Rebuild the frontend (`npm run build`) and restart the Go server after updates.

Tip: When deploying, ensure environment vars match your production Mongo/SMTP/Stripe config and that the frontend `VUE_APP_API_URL` points to your API host.
