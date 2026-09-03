import { Page, expect } from "@playwright/test";

export class UpdateCasePage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto("/update-case");

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "Update Case",
                },
            ),
        ).toBeVisible();

    }

    async loadCaseNotes() {

        await this.page
            .getByPlaceholder("Case ID")
            .fill("1");

        await this.page
            .getByRole("button", {
                name: "Load Notes",
            })
            .click();

        await expect(
            this.page.getByRole("heading", {
                name: "Analyst Timeline",
            }),
        ).toBeVisible();

    }

    async updateCase() {

        await this.page
            .getByPlaceholder("Status")
            .fill("IN_PROGRESS");

        await this.page
            .getByPlaceholder("Case Notes")
            .fill("Playwright automated update.");

        await this.page
            .getByRole("button", {
                name: "Update Case",
            })
            .click();

    }

    async addTimelineNote() {

        await this.page
            .getByPlaceholder("Add Investigation Note")
            .fill("Playwright timeline note.");

        await this.page
            .getByRole("button", {
                name: "Add Note",
            })
            .click();

    }

    async verifyTimelineVisible() {

        await expect(
            this.page.getByRole("heading", {
                name: "Analyst Timeline",
            }),
        ).toBeVisible();

    }

}