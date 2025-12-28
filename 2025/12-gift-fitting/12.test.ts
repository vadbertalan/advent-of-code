
import { first } from './12';
import { readLines } from '@/utils-ts/index';

describe('Day 12', () => {
    test('First part real', () => {
        const lines = readLines('2025/12-gift-fitting/12.in');
        expect(first(lines)).toBe('512');
    });
});
