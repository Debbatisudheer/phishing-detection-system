import { Page, expect } from "@playwright/test";

export class AlertHistoryPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto(
            "/alert-history",
        );

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Alert History",
                },
            ),
        ).toBeVisible();

    }

    async verifyTable() {

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
                    name: "Time",
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
                    name: "Risk",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "columnheader",
                {
                    name: "Verdict",
                },
            ),
        ).toBeVisible();

    }

}