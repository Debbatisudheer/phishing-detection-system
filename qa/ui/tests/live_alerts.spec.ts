import { test } from "@playwright/test";
import { LiveAlertsPage } from "../pages/LiveAlertsPage";

test(
    "Live Alerts Smoke Test",
    async ({ page }) => {

        page.on(
            "dialog",
            async dialog => {
                await dialog.accept();
            },
        );

        await page.goto("/login");

        await page
            .getByPlaceholder("Username")
            .fill("sudheer");

        await page
            .getByPlaceholder("Password")
            .fill("password123");

        await Promise.all([

            page.waitForURL("**/"),

            page.getByRole(
                "button",
                {
                    name: "Login",
                },
            ).click(),

        ]);

        const liveAlerts =
            new LiveAlertsPage(page);

        await liveAlerts.goto();

        await liveAlerts.verifyPageLoaded();

        await liveAlerts.verifyAlertsSection();

    },
);