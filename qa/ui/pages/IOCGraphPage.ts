import { Page, expect } from "@playwright/test";

export class IOCGraphPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto("/ioc-graph");

        await expect(
            this.page.getByRole("heading", {
                name: "IOC Relationship View",
            }),
        ).toBeVisible();

    }

    async verifyTable() {

        await expect(
            this.page.getByRole("columnheader", {
                name: "IOC",
            }),
        ).toBeVisible();

        await expect(
            this.page.getByRole("columnheader", {
                name: "Source",
            }),
        ).toBeVisible();

        await expect(
            this.page.getByRole("columnheader", {
                name: "File",
            }),
        ).toBeVisible();

    }

}