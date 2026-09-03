import { test } from "@playwright/test";

import { IncidentDashboardPage }
from "../pages/IncidentDashboardPage";

test(
    "Incident Dashboard Smoke Test",
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

        const incident =
            new IncidentDashboardPage(
                page,
            );

        await incident.goto();

        await incident.verifyStatistics();

        await incident.verifyRecentIncidents();

    },
);