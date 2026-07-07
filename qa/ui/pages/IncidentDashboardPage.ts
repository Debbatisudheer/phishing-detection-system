import { Page, expect } from "@playwright/test";

export class IncidentDashboardPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto(
            "/incidents",
        );

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Incident Dashboard",
                },
            ),
        ).toBeVisible();

    }

    async verifyStatistics() {

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Total Incidents",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Open Incidents",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Closed Incidents",
                },
            ),
        ).toBeVisible();

    }

    async verifyRecentIncidents() {

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Recent Incidents",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "columnheader",
                {
                    name: "ID",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "columnheader",
                {
                    name: "File",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "columnheader",
                {
                    name: "Analyst",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "columnheader",
                {
                    name: "Status",
                },
            ),
        ).toBeVisible();

    }

}