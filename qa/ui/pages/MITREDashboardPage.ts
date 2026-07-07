import { Page, expect } from "@playwright/test";

export class MITREDashboardPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto(
            "/mitre",
        );

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "MITRE Dashboard",
                },
            ),
        ).toBeVisible();

    }

    async verifyDashboardLoaded() {

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "MITRE Dashboard",
                },
            ),
        ).toBeVisible();

    }

    async verifyTechniqueCards() {

        await expect(
            this.page.getByText(
                "Techniques",
            ).first(),
        ).toBeVisible();

    }

}