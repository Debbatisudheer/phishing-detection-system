# 08 — Frontend Application

## Stack

React 19, Vite, React Router, Axios, Recharts, React Flow and React Toastify.

## Flow

`index.html` → `src/main.jsx` → `App.jsx` → `BrowserRouter` → application pages.

## Main areas

Login, dashboard, cases, search, recent findings, live alerts, threat hunting, playground, file details, threat intelligence, MITRE dashboard/heatmap, incidents, campaigns/timeline, IOC graphs/trends/network, investigation workbench and sandbox pages.

## API client

`src/services/api.js` uses `VITE_API_URL` or `http://localhost:8081`. Axios adds the JWT from local storage. A 401 removes the token and redirects to `/login`.

## Visualization

Recharts is used for charts and React Flow for graph/network investigation views.

## Deployment identity

This is a React + Vite SPA, not Next.js. The Go backend is a separate service.
