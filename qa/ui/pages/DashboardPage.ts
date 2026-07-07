import { Page, expect } from "@playwright/test";

export class DashboardPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto("/");

    }

    async verifyDashboardLoaded() {

        await expect(

            this.page.getByRole(
                "heading",
                {
                    name: "SOC Dashboard",
                },
            ),

        ).toBeVisible();

    }

    async verifyStatisticsCards() {

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Total Analyzed",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Allow",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Suspicious",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Quarantine",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Critical",
                },
            ),
        ).toBeVisible();

    }

    async verifyRiskDistribution() {

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Risk Distribution",
                },
            ),
        ).toBeVisible();

    }

    async verifySystemHealth() {

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "System Health",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                "Database Status",
            ),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                "Auto Cleanup",
            ),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                "Retention",
            ),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                "Last Cleanup",
            ),
        ).toBeVisible();

    }

}