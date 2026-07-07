import { test } from "@playwright/test";
import { DashboardPage } from "../pages/DashboardPage";

test(
    "Dashboard Smoke Test",
    async ({ page }) => {

        page.on(
            "dialog",
            async dialog => {
                await dialog.accept();
            },
        );

        await page.goto("/login");

        await page
            .getByPlaceholder("Username")
            .fill("sudheer");

        await page
            .getByPlaceholder("Password")
            .fill("password123");

        await page
            .getByRole(
                "button",
                {
                    name: "Login",
                },
            )
            .click();

        const dashboard =
            new DashboardPage(page);

        await dashboard.verifyDashboardLoaded();

        await dashboard.verifyStatisticsCards();

        await dashboard.verifyRiskDistribution();

        await dashboard.verifySystemHealth();

    },
);