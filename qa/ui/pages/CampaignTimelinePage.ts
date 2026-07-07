import { Page, expect } from "@playwright/test";

export class CampaignTimelinePage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto(
            "/campaign-timeline",
        );

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Campaign Timeline",
                },
            ),
        ).toBeVisible();

    }

    async verifyTimelineTable() {

        await expect(
            this.page.getByRole(
                "table",
            ),
        ).toBeVisible();

    }

}