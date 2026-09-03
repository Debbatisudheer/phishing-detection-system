import { test } from "@playwright/test";
import { RecentFindingsPage } from "../pages/RecentFindingsPage";

test(
    "Recent Findings Smoke Test",
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

        const findings =
            new RecentFindingsPage(page);

        await findings.goto();

        await findings.verifyFindingLoaded();

        await findings.verifyRiskScore();

        await findings.verifyRiskLevel();

        await findings.verifyVerdict();

        await findings.verifyFileLink();

    },
);