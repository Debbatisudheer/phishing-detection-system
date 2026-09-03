import http from "k6/http";
import { check, sleep } from "k6";

export const options = {

    stages: [

        { duration: "1m", target: 250 },

        { duration: "3m", target: 250 },

        { duration: "1m", target: 0 },

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
This is actually a very useful result. The test did not fail because your application crashed—it failed because it crossed your performance threshold.

What happened?

You configured this threshold:

thresholds: {
    http_req_duration: ["p(95)<1000"],
}

This means:

95% of all requests must finish in under 1 second (1000 ms).

Your actual result was:

p(95)=8.08s

So 95% of the requests completed within 8.08 seconds, which is much slower than your target.

The good news

Look at these metrics:

Metric	Result	Status
250 Virtual Users	✅	Excellent
Total Requests	30,011	Excellent
Failed Requests	0%	🟢
Login Success	100%	🟢
Application Crash	No	🟢
Server Errors	0	🟢

Your backend did not crash.

It simply became slow under 250 concurrent users.

That's exactly what load testing is designed to discover.

Interpretation

From your tests:

10 Users

P95 ≈ 4 ms

✅ Excellent

50 Users

P95 ≈ 4 ms

✅ Excellent

100 Users

P95 ≈ 4 ms

✅ Excellent

250 Users

P95 = 8.08 seconds

⚠ Performance degradation begins here.

This is actually valuable

Real load testing is supposed to find the point where the system slows down.

Now you know:

Your current backend comfortably supports around 100 concurrent users.

Around 250 concurrent users, response times increase significantly, although the system remains functional.

project report? Something like:

Load testing was performed using Grafana k6. The application successfully handled up to 100 concurrent users while maintaining a P95 response time below 5 ms. Under a load of 250 concurrent users, no request failures occurred; however, the P95 response time increased to approximately 8 seconds, indicating that the application's throughput limit was reached and additional optimization or scaling would be required for higher concurrency.

That is a professional conclusion.

Should we continue to 500 users?

I would not recommend it right now.

Why?

At 250 users, you've already identified the capacity limit. Running 500 users will almost certainly produce even higher response times and won't add much new information unless your goal is to perform a stress test.
*/