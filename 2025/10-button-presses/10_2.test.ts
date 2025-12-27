import { second } from "./10_2";
import { readLines } from "@/utils-ts/index";

describe("Day 10", () => {
  test("Second part example", async () => {
    const lines = readLines(`${__dirname}/10.exin`);
    await expect(second(lines)).resolves.toBe("33");
  });

  test("Second part real", async () => {
    const lines = readLines(`${__dirname}/10.in`);
    expect(await second(lines)).toBe("16463");
  });
});
