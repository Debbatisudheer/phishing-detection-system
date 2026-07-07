import { test } from "@playwright/test";
import { ThreatHuntingPage } from "../pages/ThreatHuntingPage";

test(
    "Threat Hunting Dashboard Smoke Test",
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

        const hunting =
            new ThreatHuntingPage(
                page,
            );

        await hunting.goto();

        await hunting.verifyStatistics();

        await hunting.verifyChart();

        await hunting.verifyMITRETable();

    },
);