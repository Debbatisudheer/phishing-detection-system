import { Page, expect } from "@playwright/test";

export class SearchPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto("/search");

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "IOC Search",
                },
            ),
        ).toBeVisible();

    }

    async searchIOC(
        query: string,
    ) {

        await this.page
            .getByPlaceholder(
                "Search IOC",
            )
            .fill(query);

        await this.page
            .getByRole(
                "button",
                {
                    name: "Search",
                },
            )
            .click();

    }

    async verifySearchCompleted() {

        await expect(

            this.page.locator(
                "body",
            ),

        ).toContainText(

            /Risk:|No results found/,

        );

    }

    async verifyResultCard() {

        const body =
            await this.page
                .locator("body")
                .textContent();

        if (
            body?.includes(
                "No results found",
            )
        ) {
            return;
        }

        await expect(
            this.page.locator("h3").nth(1),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                /Risk:/,
            ).first(),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                /Verdict:/,
            ).first(),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                /SHA256:/,
            ).first(),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                /MITRE:/,
            ).first(),
        ).toBeVisible();

    }

}