import { test, expect } from "@playwright/test";
import { PlaygroundPage } from "../pages/PlaygroundPage";

test(
    "Playground Email Analysis",
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

        await playground.analyzeEmail(
            "Urgent Password Reset",
            "Click https://google.com to reset your password.",
        );

        await playground.verifyAnalysisCompleted();

        await playground.verifyRiskScore();

        await playground.verifyRiskLevel();

        await playground.verifyFindings();

    },
);