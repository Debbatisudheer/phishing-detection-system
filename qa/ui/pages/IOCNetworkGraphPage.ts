import { Page, expect } from "@playwright/test";

export class IOCNetworkGraphPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto("/ioc-network");

        await expect(
            this.page.getByRole("heading", {
                name: "IOC Relationship Network",
            }),
        ).toBeVisible();

    }

    async verifyGraph() {

        await expect(
            this.page.locator(
                ".react-flow",
            ),
        ).toBeVisible();

    }

}