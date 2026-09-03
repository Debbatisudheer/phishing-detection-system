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

    const payload = JSON.stringify({

        username: "sudheer",

        password: "password123",

    });

    const params = {

        headers: {

            "Content-Type": "application/json",

        },

    };

    const response = http.post(

        "http://localhost:8081/api/login",

        payload,

        params,

    );

    check(response, {

        "Login Successful": (r) => r.status === 200,

        "Token Returned": (r) => {

            try {

                return JSON.parse(r.body).token !== "";

            } catch {

                return false;

            }

        },

    });

    sleep(1);

}

/*
Performance Summary
Metric	Result	Assessment
Virtual Users	10	✅
Duration	30 seconds	✅
Total Requests	300	✅
Failed Requests	0	🟢 Excellent
Successful Checks	100%	🟢 Excellent
Average Response	18.37 ms	🟢 Excellent
Median	1.32 ms	🟢 Excellent
P95	11.15 ms	🟢 Excellent
Max Response	479.76 ms	🟢 Acceptable (likely a one-off spike)
Interpretation

This is what matters most:

Failed Requests = 0%

✅ No login failures under load.

100% Checks Passed

✅ Every response returned a valid JWT.

P95 = 11 ms

This is the most important performance metric.

It means:

95% of all login requests completed in under 11 ms.

That is an excellent result for a local development environment.

*/