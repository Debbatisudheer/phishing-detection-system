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