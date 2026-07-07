import { Page, expect } from "@playwright/test";

export class LiveAlertsPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto("/live-alerts");

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Live Alerts",
                },
            ),
        ).toBeVisible();

    }

    async verifyPageLoaded() {

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Live Alerts",
                },
            ),
        ).toBeVisible();

    }

    async verifyAlertsSection() {

        await expect(
            this.page.locator("body"),
        ).toContainText(
            "Live Alerts",
        );

    }

}