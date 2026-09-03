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

    // Threat Intelligence

    const response = http.get(

        "http://localhost:8081/api/threat-intel",

        {

            headers: {

                Authorization: `Bearer ${token}`,

            },

        },

    );

    check(response, {

        "Threat Intel Loaded": (r) => r.status === 200,

    });

    sleep(1);

}

/*
Threat Intelligence Performance Results
Metric	Result	Status
Virtual Users	10	✅
Duration	30 seconds	✅
Requests	586	✅
Failed Requests	0%	🟢
Successful Checks	100%	🟢
Average Response	18.15 ms	🟢
Median	1.66 ms	🟢
P90	5.26 ms	🟢
P95	86.11 ms	🟢 Excellent
Max	691 ms	Acceptable (single outlier)
Verdict
✅ Login successful
✅ Threat Intelligence API successful
✅ Zero failures
✅ P95 well below the 500 ms threshold
✅ Performance test passed

*/