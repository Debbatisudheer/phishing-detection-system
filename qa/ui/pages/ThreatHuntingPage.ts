import { Page, expect } from "@playwright/test";

export class ThreatHuntingPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto(
            "/threat-hunting",
        );

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Threat Hunting Dashboard",
                },
            ),
        ).toBeVisible();

    }

    async verifyStatistics() {

        await expect(
            this.page.getByText(
                "Critical Files",
            ),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                "Quarantine Files",
            ),
        ).toBeVisible();

    }

    async verifyChart() {

        await expect(
            this.page.locator(
                ".recharts-wrapper",
            ),
        ).toBeVisible();

    }

    async verifyMITRETable() {

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "MITRE Techniques",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "table",
            ),
        ).toBeVisible();

    }

}