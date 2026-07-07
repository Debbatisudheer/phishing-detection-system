import { Page, expect } from "@playwright/test";

export class RecentFindingsPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto("/recent-findings");

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Recent Findings",
                },
            ),
        ).toBeVisible();

    }

    async verifyFindingLoaded() {

        await expect(
            this.page.locator("h3").first(),
        ).toBeVisible();

    }

    async verifyRiskScore() {

        await expect(
            this.page.getByText(
                /Risk Score:/,
            ).first(),
        ).toBeVisible();

    }

    async verifyRiskLevel() {

        await expect(
            this.page.getByText(
                /Risk Level:/,
            ).first(),
        ).toBeVisible();

    }

    async verifyVerdict() {

        await expect(
            this.page.getByText(
                /Verdict:/,
            ).first(),
        ).toBeVisible();

    }

    async verifyFileLink() {

        await expect(
            this.page.locator("h3 a").first(),
        ).toBeVisible();

    }

}