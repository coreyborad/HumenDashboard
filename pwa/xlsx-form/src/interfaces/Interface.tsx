export interface User {
    email: string;
    password: string;
};

export interface Accounting {
    date: Date | null;
    payer: string;
    category: string;
    subCategory: string;
    note: string;
    cost: number;
}
