import { test } from "@playwright/test";

import { IOCGraphPage }
from "../pages/IOCGraphPage";

import { IOCTrendsPage }
from "../pages/IOCTrendsPage";

import { IOCNetworkGraphPage }
from "../pages/IOCNetworkGraphPage";

test(
    "IOC Visualization Smoke Test",
    async ({ page }) => {

        page.on(
            "dialog",
            async dialog => {

                await dialog.accept();

            },
        );

        await page.goto("/login");

        await page
            .getByPlaceholder(
                "Username",
            )
            .fill("sudheer");

        await page
            .getByPlaceholder(
                "Password",
            )
            .fill("password123");

        await Promise.all([

            page.waitForURL("**/"),

            page
                .getByRole("button", {
                    name: "Login",
                })
                .click(),

        ]);

        //
        // IOC Graph
        //

        const graph =
            new IOCGraphPage(page);

        await graph.goto();

        await graph.verifyTable();

        //
        // IOC Trends
        //

        const trends =
            new IOCTrendsPage(page);

        await trends.goto();

        await trends.verifyStatistics();

        await trends.verifyChart();

        //
        // IOC Network
        //

        const network =
            new IOCNetworkGraphPage(page);

        await network.goto();

        await network.verifyGraph();

    },
);