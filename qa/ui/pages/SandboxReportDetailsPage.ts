import { Page, expect } from "@playwright/test";

export class SandboxReportDetailsPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto(
            "/sandbox-report/1",
        );

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Sandbox Report",
                },
            ),
        ).toBeVisible();

    }

    async verifyReportDetails() {

        await expect(
            this.page.getByText(
                "ID:",
            ),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                "File:",
            ),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                "Risk Score:",
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
            this.page.getByRole(
                "heading",
                {
                    name: "Findings",
                },
            ),
        ).toBeVisible();

    }

}