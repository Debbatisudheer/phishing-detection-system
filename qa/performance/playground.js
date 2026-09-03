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

    const token =
        JSON.parse(login.body).token;

    // Playground Email Analysis

    const analysis = http.post(
        "http://localhost:8081/api/analyze-email",

        JSON.stringify({

            subject: "Urgent Password Reset",

            body: "Click https://google.com to reset your password immediately.",

        }),

        {

            headers: {

                Authorization: `Bearer ${token}`,

                "Content-Type": "application/json",

            },

        },

    );

    check(analysis, {

        "Playground Loaded": (r) => r.status === 200,

    });

    sleep(1);

}

/*  
Playground Performance Results
Metric	Result	Status
Virtual Users	10	✅
Duration	30 sec	✅
HTTP Requests	600	✅
Failed Requests	0%	🟢
Successful Checks	100%	🟢
Average Response	7.66 ms	🟢
Median	1.99 ms	🟢
P95	10.56 ms	🟢 Excellent
Max Response	360.64 ms	🟢

Verdict: ✅ Passed. The Playground API handled the load comfortably with zero failures.
*/