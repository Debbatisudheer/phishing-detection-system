import { Page, expect } from "@playwright/test";

export class CampaignPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto("/campaigns");

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Campaign Dashboard",
                },
            ),
        ).toBeVisible();

    }

    async verifyCampaignTable() {

        await expect(
            this.page.getByRole(
                "table",
            ).first(),
        ).toBeVisible();

    }

    async viewSources() {

        const button =
            this.page.getByRole(
                "button",
                {
                    name: "View Sources",
                },
            ).first();

        if (
            await button.isVisible()
        ) {

            await button.click();

        }

    }

    async verifySourcesTable() {

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "IOC Sources",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "table",
            ).nth(1),
        ).toBeVisible();

    }

}