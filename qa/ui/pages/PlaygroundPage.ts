import { Page, expect } from "@playwright/test";

export class PlaygroundPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto("/playground");

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Demo Library",
                },
            ),
        ).toBeVisible();

    }

    async analyzeEmail(
        subject: string,
        body: string,
    ) {

        await this.page
            .locator("input")
            .nth(3)
            .fill(subject);

        await this.page
            .locator("textarea")
            .fill(body);

        await this.page
            .getByRole(
                "button",
                {
                    name: "Analyze Email",
                },
            )
            .click();

    }

    async analyzeDemoFile() {

        await this.page
            .getByRole(
                "button",
                {
                    name: /DOCM/,
                },
            )
            .click();

        await this.page
            .getByRole(
                "button",
                {
                    name: "Analyze Email",
                },
            )
            .click();

    }

    async verifyAnalysisCompleted() {

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Analysis Summary",
                },
            ),
        ).toBeVisible();

    }

    async verifyRiskScore() {

        await expect(
            this.page.locator(
                ".risk-card h1",
            ),
        ).toBeVisible();

    }

    async verifyRiskLevel() {

        await expect(
            this.page.locator(
                ".risk-card span",
            ),
        ).toBeVisible();

    }

    async verifySandboxStatus() {

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Sandbox Status",
                },
            ),
        ).toBeVisible();

    }

    async verifyFindings() {

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Detection Findings",
                },
            ),
        ).toBeVisible();

    }

}