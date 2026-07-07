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
    const loginPayload = JSON.stringify({

        username: "sudheer",

        password: "password123",

    });

    const loginResponse = http.post(

        "http://localhost:8081/api/login",

        loginPayload,

        {

            headers: {

                "Content-Type": "application/json",

            },

        },

    );

    check(loginResponse, {

        "Login Success": (r) => r.status === 200,

    });

    const token =
        JSON.parse(loginResponse.body).token;

    // Search API
    const response = http.get(

        "http://localhost:8081/api/search?q=google.com",

        {

            headers: {

                Authorization: `Bearer ${token}`,

            },

        },

    );

    check(response, {

        "Search Success": (r) => r.status === 200,

    });

    sleep(1);

}

/*
Search Performance Results
Metric	Result	Status
Virtual Users	10	✅
Duration	30 sec	✅
HTTP Requests	600	✅
Failed Requests	0%	🟢
Successful Checks	100%	🟢
Average Response	10.46 ms	🟢
Median	2.32 ms	🟢
P95	9.13 ms	🟢 Excellent
Max Response	470.9 ms	🟢

Again...

✅ Login works
✅ Search works
✅ No failures
✅ Under 500ms threshold
✅ Excellent latency

*/