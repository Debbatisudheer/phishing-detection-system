import { Page, expect } from "@playwright/test";

export class SandboxDashboardPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto(
            "/sandbox",
        );

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Sandbox Dashboard",
                },
            ),
        ).toBeVisible();

    }

    async verifyRefreshButton() {

        await expect(
            this.page.getByRole(
                "button",
                {
                    name: "Refresh",
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
                    name: "File",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "columnheader",
                {
                    name: "Risk Score",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "columnheader",
                {
                    name: "Risk Level",
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

        await expect(
            this.page.getByRole(
                "columnheader",
                {
                    name: "Action",
                },
            ),
        ).toBeVisible();

    }

}