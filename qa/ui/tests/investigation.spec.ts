import { test } from "@playwright/test";
import { InvestigationWorkbenchPage } from "../pages/InvestigationWorkbenchPage";

test(
    "Investigation Workbench Smoke Test",
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

        const workbench =
            new InvestigationWorkbenchPage(
                page,
            );

        await workbench.goto();

        //
        // Use any IOC that exists in your database.
        //

        await workbench.investigate(
            "google.com",
        );

        await workbench.verifyResults();

    },
);