import { test } from "@playwright/test";
import { CorrelationPage } from "../pages/CorrelationPage";

test(
    "IOC Correlation Smoke Test",
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

        const correlation =
            new CorrelationPage(
                page,
            );

        await correlation.goto();

        await correlation.verifyTable();

    },
);