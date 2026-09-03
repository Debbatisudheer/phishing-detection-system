import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
    vus: 10,
    duration: "30s",

    thresholds: {
        http_req_duration: ["p(95)<500"],
        http_req_failed: ["rate<0.01"],
    },
};

export default function () {

    // Login

    const login = http.post(
        "http://localhost:8081/api/login",
        JSON.stringify({
            username: "sudheer",
            password: "password123",
        }),
        {
            headers: {
                "Content-Type": "application/json",
            },
        },
    );

    check(login, {
        "Login Success": (r) => r.status === 200,
    });

    const token = JSON.parse(login.body).token;

    const headers = {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    };

    // Dashboard
    const dashboard = http.get(
        "http://localhost:8081/api/dashboard",
        headers,
    );

    check(dashboard, {
        "Dashboard Loaded": (r) => r.status === 200,
    });

    // Search
    const search = http.get(
        "http://localhost:8081/api/search?q=google",
        headers,
    );

    check(search, {
        "Search Loaded": (r) => r.status === 200,
    });

    // Threat Intelligence
    const intel = http.get(
        "http://localhost:8081/api/threat-intel",
        headers,
    );

    check(intel, {
        "Threat Intel Loaded": (r) => r.status === 200,
    });

    // Cases
    const cases = http.get(
        "http://localhost:8081/api/cases",
        headers,
    );

    check(cases, {
        "Cases Loaded": (r) => r.status === 200,
    });

    // Investigation
    const investigation = http.get(
        "http://localhost:8081/api/investigation-summary?ioc=google.com",
        headers,
    );

    check(investigation, {
        "Investigation Loaded": (r) => r.status === 200,
    });

    // Sandbox
    const sandbox = http.get(
        "http://localhost:8081/api/sandbox/reports",
        headers,
    );

    check(sandbox, {
        "Sandbox Loaded": (r) => r.status === 200,
    });

    sleep(1);
}

/*
Performance Testing Summary
Test	Status
Login	✅
Dashboard	✅
Search	✅
Playground	✅
Threat Intelligence	✅
Cases	✅
Investigation	✅
Sandbox	✅
Full User Workflow	✅

Your final workflow test is especially strong:

✅ 2002 HTTP Requests
✅ 0 Failed Requests
✅ 100% Checks Passed
✅ P95 = 10.26 ms
✅ Average = 9.83 ms

That is an excellent result for a local development environment.
*/