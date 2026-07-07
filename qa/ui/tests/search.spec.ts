import { test } from "@playwright/test";
import { SearchPage } from "../pages/SearchPage";

test(
    "IOC Search Test",
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

        await Promise.all([
            page.waitForURL("**/"),
            page.getByRole(
                "button",
                {
                    name: "Login",
                },
            ).click(),
        ]);

        const search =
            new SearchPage(page);

        await search.goto();

        await search.searchIOC(
            "invoice",
        );

        await search.verifySearchCompleted();

        await search.verifyResultCard();

    },
);