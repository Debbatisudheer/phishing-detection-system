import { test } from "@playwright/test";
import { SandboxDashboardPage } from "../pages/SandboxDashboardPage";

test(
    "Sandbox Dashboard Smoke Test",
    async ({ page }) => {

        page.on(
            "dialog",
            async dialog => {

                await dialog.accept();

            },
        );

        await page.goto("/login");

        await page
            .getByPlaceholder(
                "Username",
            )
            .fill("sudheer");

        await page
            .getByPlaceholder(
                "Password",
            )
            .fill("password123");

        await Promise.all([

            page.waitForURL("**/"),

            page
                .getByRole(
                    "button",
                    {
                        name: "Login",
                    },
                )
                .click(),

        ]);

        const sandbox =
            new SandboxDashboardPage(
                page,
            );

        await sandbox.goto();

        await sandbox.verifyRefreshButton();

        await sandbox.verifyTable();

    },
);