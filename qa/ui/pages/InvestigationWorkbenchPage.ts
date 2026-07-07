import { Page, expect } from "@playwright/test";

export class InvestigationWorkbenchPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto(
            "/investigation",
        );

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Investigation Workbench",
                },
            ),
        ).toBeVisible();

    }

    async investigate(
        ioc: string,
    ) {

        await this.page
            .getByPlaceholder(
                "Enter IOC",
            )
            .fill(ioc);

        await this.page
            .getByRole(
                "button",
                {
                    name: "Investigate",
                },
            )
            .click();

    }

    async verifyResults() {

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "IOC Details",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                "Occurrences:",
            ),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                "Risk Level:",
            ),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                "Verdict:",
            ),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                "MITRE:",
            ),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                "Reputation:",
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Sources",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Files",
                },
            ),
        ).toBeVisible();

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Analyst Notes",
                },
            ),
        ).toBeVisible();

    }

}