import http from "k6/http";
import { check, sleep } from "k6";

export const options = {

    stages: [

        { duration: "30s", target: 100 },

        { duration: "2m", target: 100 },

        { duration: "30s", target: 0 },

    ],

    thresholds: {

        http_req_duration: ["p(95)<1000"],

        http_req_failed: ["rate<0.05"],

    },

};

export default function () {

    const response = http.post(

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

    check(response, {

        "Login Success": (r) => r.status === 200,

    });

    sleep(1);

}

/*

100 Virtual Users Load Test Results
Metric	Result	Status
Max Virtual Users	100	✅
Duration	3 minutes	✅
Total Requests	14,924	✅
Failed Requests	0%	🟢
Successful Checks	100%	🟢
Average Response Time	7.83 ms	🟢
Median	1.23 ms	🟢
P95	4.39 ms	🟢
Maximum	1.3 s	Acceptable (rare outlier)
Verdict
✅ 100 concurrent users handled successfully
✅ Nearly 15,000 requests
✅ Zero failures
✅ P95 well below the 1-second threshold
✅ System remained stable

At this point, your backend has demonstrated solid performance under moderate concurrent load.
*/