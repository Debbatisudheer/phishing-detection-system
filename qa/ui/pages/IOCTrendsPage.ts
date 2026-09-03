import { Page, expect } from "@playwright/test";

export class IOCTrendsPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto("/ioc-trends");

        await expect(
            this.page.getByRole("heading", {
                name: "IOC Trends Dashboard",
            }),
        ).toBeVisible();

    }

    async verifyStatistics() {

        await expect(
            this.page.getByText(
                "Total IOC Events",
            ),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                "Peak Day",
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

}