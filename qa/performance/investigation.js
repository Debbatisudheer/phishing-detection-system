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

    // Investigation API

    const response = http.get(
        "http://localhost:8081/api/investigation-summary?ioc=google.com",
        {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        },
    );

    check(response, {
        "Investigation Loaded": (r) => r.status === 200,
    });

    sleep(1);
}

/*
Investigation Performance Results
Metric	Result	Status
Virtual Users	10	✅
Duration	30 sec	✅
HTTP Requests	600	✅
Failed Requests	0%	🟢
Successful Checks	100%	🟢
Average Response	8.56 ms	🟢
Median	2.01 ms	🟢
P95	5.15 ms	🟢 Excellent
Maximum	340.39 ms	🟢
*/