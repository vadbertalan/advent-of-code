import type { Config } from 'jest';

const config: Config = {
  preset: 'ts-jest',
  testEnvironment: 'node',
  moduleNameMapper: {
    '^@/(.*)\\.js$': '<rootDir>/$1.ts',
    '^@/(.*)$': '<rootDir>/$1',
  },
  transform: {
    '^.+\\.tsx?$': ['ts-jest', {
        useESM: true,
    }],
  },
  extensionsToTreatAsEsm: ['.ts'],
  // Since we use ESM in node (type: module might not be set in package.json but we use .js import extensions), 
  // let's ensure ts-jest handles it.
};

export default config;
