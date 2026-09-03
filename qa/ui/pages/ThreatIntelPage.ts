import { Page, expect } from "@playwright/test";

export class ThreatIntelPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto("/api/threat-intel");

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Threat Intelligence Dashboard",
                },
            ),
        ).toBeVisible();

    }

    async verifyStatistics() {

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Total IOCs",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Total Alerts",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Critical Files",
                },
            ),
        ).toBeVisible();

    }

    async verifyTopFiles() {

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Top Risk Files",
                },
            ),
        ).toBeVisible();

    }

    async verifyTopIOCs() {

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Top IOCs",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "table",
            ),
        ).toBeVisible();

    }

    async lookupIOC(
        value: string,
    ) {

        await this.page
            .getByPlaceholder(
                "Enter IOC",
            )
            .fill(value);

        await this.page
            .getByRole(
                "button",
                {
                    name: "Lookup",
                },
            )
            .click();

    }

    async verifyLookupCompleted() {

        await expect(
            this.page.locator(
                "body",
            ),
        ).toContainText(
            /Reputation:|Source:/,
        );

    }

}