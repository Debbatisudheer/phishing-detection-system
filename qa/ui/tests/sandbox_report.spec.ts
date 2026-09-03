import { test } from "@playwright/test";
import { SandboxReportDetailsPage } from "../pages/SandboxReportDetailsPage";

test(
    "Sandbox Report Details Smoke Test",
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

        const report =
            new SandboxReportDetailsPage(
                page,
            );

        await report.goto();

        await report.verifyReportDetails();

    },
);