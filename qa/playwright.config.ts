import { defineConfig, devices } from "@playwright/test";
import dotenv from "dotenv";

dotenv.config({
  path: "./configs/.env.dev",
});

export default defineConfig({

  testDir: "./ui/tests",

  timeout: 60000,

  expect: {
    timeout: 10000,
  },

  fullyParallel: false,

  forbidOnly: !!process.env.CI,

  retries: process.env.CI ? 2 : 0,

  workers: 1,

  reporter: [
    ["html"],
    ["list"]
  ],

  use: {

    baseURL: process.env.UI_BASE_URL,

    trace: "retain-on-failure",

    screenshot: "only-on-failure",

    video: "retain-on-failure",

    actionTimeout: 15000,

    navigationTimeout: 30000,

    ignoreHTTPSErrors: true,
  },

  projects: [

    {
      name: "chromium",

      use: {
        ...devices["Desktop Chrome"],
      },
    },

    {
      name: "firefox",

      use: {
        ...devices["Desktop Firefox"],
      },
    },

    {
      name: "webkit",

      use: {
        ...devices["Desktop Safari"],
      },
    },

  ],

});