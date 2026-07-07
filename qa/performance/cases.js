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

    const response = http.get(
        "http://localhost:8081/api/cases",
        {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        },
    );

    check(response, {
        "Cases Loaded": (r) => r.status === 200,
    });

    sleep(1);
}

/*
Results Summary
Metric	Result	Status
Virtual Users	10	✅
Duration	30 sec	✅
Requests	600	✅
Failed Requests	0%	🟢
Successful Checks	100%	🟢
Average Response	9.36 ms	🟢
P95	17.44 ms	🟢
Maximum	310.78 ms	🟢

Everything is well below your threshold of 500 ms, so this test is a clean pass.
*/