import { test } from "@playwright/test";

import { CampaignPage }
from "../pages/CampaignPage";

import { CampaignTimelinePage }
from "../pages/CampaignTimelinePage";

test(
    "Campaign Smoke Test",
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

            page.waitForURL(
                "**/",
            ),

            page
                .getByRole(
                    "button",
                    {
                        name: "Login",
                    },
                )
                .click(),

        ]);

        const campaign =
            new CampaignPage(page);

        await campaign.goto();

        await campaign.verifyCampaignTable();

        await campaign.viewSources();

        await campaign.verifySourcesTable();

        const timeline =
            new CampaignTimelinePage(page);

        await timeline.goto();

        await timeline.verifyTimelineTable();

    },
);