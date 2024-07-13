import bcrypt from 'bcrypt';


async function hashPassword(password: string): Promise<string> {
    return await bcrypt.hash(password, 10);
}

async function comparePassword(password: string, hashedPassword: string): Promise<boolean> {
    return await bcrypt.compare(password, hashedPassword);
}

export { hashPassword, comparePassword };