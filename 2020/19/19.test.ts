
import { first, second } from './19';
import { readLines } from '@/utils-ts/index';

describe('Day 19', () => {
    
    test('First part example', () => {
        const lines = readLines('2020/19/19.exin');
        expect(first(lines)).toBe('2');
    });
    
    test('First part real', () => {
        const lines = readLines('2020/19/19.in');
        expect(first(lines)).toBe('230');
    });
});
