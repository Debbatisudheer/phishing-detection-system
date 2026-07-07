import { Page, expect } from "@playwright/test";

export class LoginPage {

    constructor(
        private page: Page,
    ) {}

    async goto() {

        await this.page.goto("/");

    }

    async login(
        username: string,
        password: string,
    ) {

        await this.page.getByPlaceholder(
            "Username",
        ).fill(
            username,
        );

        await this.page.getByPlaceholder(
            "Password",
        ).fill(
            password,
        );

        await this.page.getByRole(
            "button",
            {
                name: "Login",
            },
        ).click();

    }

    async verifyLoginPage() {

        await expect(

            this.page.getByRole(
                "heading",
                {
                    name: "Login",
                },
            ),

        ).toBeVisible();

    }

}