import http from "k6/http";
import { check, sleep } from "k6";

export const options = {

    stages: [

        { duration: "30s", target: 50 },

        { duration: "1m", target: 50 },

        { duration: "30s", target: 0 },

    ],

    thresholds: {

        http_req_duration: ["p(95)<1000"],

        http_req_failed: ["rate<0.05"],

    },

};

export default function () {

    const payload = JSON.stringify({

        username: "sudheer",

        password: "password123",

    });

    const response = http.post(

        "http://localhost:8081/api/login",

        payload,

        {

            headers: {

                "Content-Type": "application/json",

            },

        },

    );

    check(response, {

        "Login Success": (r) => r.status === 200,

    });

    sleep(1);

}

/*
50 Virtual Users Load Test Results
Metric	Result	Status
Max Virtual Users	50	✅
Duration	2 minutes	✅
Total Requests	4528	✅
Failed Requests	0%	🟢
Successful Checks	100%	🟢
Average Response Time	1.82 ms	🟢
P95 Response Time	3.99 ms	🟢
Max Response Time	177.48 ms	🟢
Verdict

Your backend handled:

✅ 50 concurrent users
✅ 4528 login requests
✅ Zero failures
✅ Very low latency
✅ Stable throughout the test

This is a strong result.
*/