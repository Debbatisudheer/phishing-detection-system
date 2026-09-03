import { test } from "@playwright/test";
import { ThreatIntelPage } from "../pages/ThreatIntelPage";

test(
    "Threat Intelligence Smoke Test",
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

            page
                .getByRole(
                    "button",
                    {
                        name: "Login",
                    },
                )
                .click(),

        ]);

        const intel =
            new ThreatIntelPage(page);

        await intel.goto();

        await intel.verifyStatistics();

        await intel.verifyTopFiles();

        await intel.verifyTopIOCs();

        await intel.lookupIOC(
            "google.com",
        );

        await intel.verifyLookupCompleted();

    },
);