import { test } from "@playwright/test";
import { AlertHistoryPage } from "../pages/AlertHistoryPage";

test(
    "Alert History Smoke Test",
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

        const alertHistory =
            new AlertHistoryPage(
                page,
            );

        await alertHistory.goto();

        await alertHistory.verifyTable();

    },
);