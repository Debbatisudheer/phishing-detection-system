import { test } from "@playwright/test";

import { MITREDashboardPage }
from "../pages/MITREDashboardPage";

import { MITREHeatmapPage }
from "../pages/MITREHeatmapPage";

test(
    "MITRE Smoke Test",
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
                .getByRole(
                    "button",
                    {
                        name: "Login",
                    },
                )
                .click(),

        ]);

        const dashboard =
            new MITREDashboardPage(page);

        await dashboard.goto();

        await dashboard.verifyDashboardLoaded();

        await dashboard.verifyTechniqueCards();

        const heatmap =
            new MITREHeatmapPage(page);

        await heatmap.goto();

        await heatmap.verifyStatistics();

        await heatmap.verifyChart();

    },
);