import { Page, expect } from "@playwright/test";

export class CorrelationPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto(
            "/correlation",
        );

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "IOC Correlation",
                },
            ),
        ).toBeVisible();

    }

    async verifyTable() {

        await expect(
            this.page.getByRole(
                "columnheader",
                {
                    name: "IOC",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "columnheader",
                {
                    name: "Occurrences",
                },
            ),
        ).toBeVisible();

    }

}