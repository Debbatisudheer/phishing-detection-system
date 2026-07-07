import { Page, expect } from "@playwright/test";

export class CasesPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto("/cases");

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Cases",
                },
            ),
        ).toBeVisible();

    }

    async verifyCasesLoaded() {

        await expect(
            this.page.locator("h3").first(),
        ).toBeVisible();

    }

    async verifyAnalystVisible() {

        await expect(
            this.page.getByText(
                /Analyst:/,
            ).first(),
        ).toBeVisible();

    }

    async verifyStatusVisible() {

        await expect(
            this.page.getByText(
                /Status:/,
            ).first(),
        ).toBeVisible();

    }

    async verifyNotesVisible() {

        await expect(
            this.page.getByText(
                /Notes:/,
            ).first(),
        ).toBeVisible();

    }

}