import { Page, expect } from "@playwright/test";

export class MITREHeatmapPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto(
            "/mitre-heatmap",
        );

        await expect(
            this.page.getByRole(
                "heading",
                {
                    name: "MITRE ATT&CK Heatmap",
                },
            ),
        ).toBeVisible();

    }

    async verifyStatistics() {

        await expect(
            this.page.getByText(
                "Total Techniques",
            ),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                "Total Detections",
            ),
        ).toBeVisible();

        await expect(
            this.page.getByText(
                "Top Technique",
            ),
        ).toBeVisible();

    }

    async verifyChart() {

        await expect(
            this.page.locator(
                ".recharts-wrapper",
            ),
        ).toBeVisible();

    }

}