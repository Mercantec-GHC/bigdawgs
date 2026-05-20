# Big Dogs: Rise of the Pack

Big Dogs: Rise of the Pack is a browser-based multiplayer strategy game inspired by games like Travian and Grepolis.

This repository contains the Ruby on Rails frontend application responsible for:
- User authentication
- Gameplay interface
- Buildings and upgrades
- Marketplace system
- Map system
- Communication with backend services

The project was created as a school project.

---

# Features

- Login and account creation
- Dynamic resource bar using Hotwire
- Building system
- Marketplace with dynamic prices
- Multiplayer map
- JWT communication between services
- Background jobs using Solid Queue
- Hotwire + Stimulus frontend interactions

---

# Technologies

- Ruby on Rails 8
- PostgreSQL
- Hotwire
  - Turbo
  - Stimulus
- Slim
- Bootstrap 5
- Faraday
- JWT
- Solid Queue
- Docker

---

# Project Architecture

The Rails application acts as both:
- Frontend/UI
- Gateway to backend services

The project communicates with multiple external backends:
- Game Engine
- Marketplace Backend
- Account/Auth Systems

Communication between services happens through API calls using Faraday and JWT tokens.

---

# Installation

## Requirements

- Ruby 3.x
- PostgreSQL
- Node.js
- Yarn
- Docker (optional)

---

# Setup

Clone the repository:

```bash
git clone https://github.com/Mercantec-GHC/h5-msp-bigdawgs-1
cd h5-msp-bigdawgs-1
