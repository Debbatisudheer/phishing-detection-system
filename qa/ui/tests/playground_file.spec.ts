import { test, expect } from "@playwright/test";
import { PlaygroundPage } from "../pages/PlaygroundPage";

test(
    "Playground Demo File Analysis",
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
            page.getByRole("button", {
                name: "Login",
            }).click(),
        ]);

        const playground =
            new PlaygroundPage(page);

        await playground.goto();

        await playground.analyzeDemoFile();

        await playground.verifyAnalysisCompleted();

        await playground.verifyRiskScore();

        await playground.verifyRiskLevel();

        await playground.verifySandboxStatus();

        await playground.verifyFindings();

    },
);