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

    // Sandbox Dashboard

    const response = http.get(
        "http://localhost:8081/api/sandbox/reports",
        {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        },
    );

    check(response, {
        "Sandbox Reports Loaded": (r) => r.status === 200,
    });

    sleep(1);
}

/*
Sandbox Performance Results
Metric	Result	Status
Virtual Users	10	✅
Duration	30 sec	✅
Requests	598	✅
Failed Requests	0%	🟢
Successful Checks	100%	🟢
Average Response	8.54 ms	🟢
P95	6.73 ms	🟢
Maximum	363.16 ms	🟢

Verdict: ✅ Passed.
*/