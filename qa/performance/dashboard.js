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

    // Login first
    const loginPayload = JSON.stringify({

        username: "sudheer",

        password: "password123",

    });

    const loginParams = {

        headers: {

            "Content-Type": "application/json",

        },

    };

    const loginResponse = http.post(

        "http://localhost:8081/api/login",

        loginPayload,

        loginParams,

    );

    check(loginResponse, {

        "Login Success": (r) => r.status === 200,

    });

    const token =
        JSON.parse(
            loginResponse.body,
        ).token;

    // Dashboard request
    const dashboardResponse = http.get(

        "http://localhost:8081/api/dashboard",

        {

            headers: {

                Authorization:
                    `Bearer ${token}`,

            },

        },

    );

    check(dashboardResponse, {

        "Dashboard Loaded": (r) =>
            r.status === 200,

    });

    sleep(1);

}

/*
Dashboard Performance Summary
Metric	Result	Assessment
Virtual Users	10	✅
Duration	30 seconds	✅
Total HTTP Requests	600	✅
Failed Requests	0	🟢 Excellent
Successful Checks	100%	🟢 Excellent
Average Response	12.18 ms	🟢 Excellent
Median	1.70 ms	🟢 Excellent
P95	4.93 ms	🟢 Outstanding
Max Response	711.02 ms	🟡 Acceptable (single outlier)

The P95 latency of 4.93 ms is especially impressive. It means almost all authenticated dashboard requests completed in under 5 ms during this test.

*/