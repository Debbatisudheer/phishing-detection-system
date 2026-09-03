import { test } from "@playwright/test";
import { UpdateCasePage } from "../pages/UpdateCasePage";

test(
    "Update Case Smoke Test",
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

        const updateCase =
            new UpdateCasePage(
                page,
            );

        await updateCase.goto();

        await updateCase.loadCaseNotes();

        await updateCase.verifyTimelineVisible();

        await updateCase.updateCase();

        await updateCase.addTimelineNote();

    },
);