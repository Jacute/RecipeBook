import fs from 'fs';
import path from 'path';


const JWT_PRIVATE_KEY = fs.readFileSync(path.join(__dirname, '../jwt_keys/private.pem'));
const JWT_PUBLIC_KEY = fs.readFileSync(path.join(__dirname, '../jwt_keys/public.pem'));
const TOKEN_LIFETIME = "60 days";
const COOKIE_LIFETIME = 1000 * 60 * 60 * 24 * 60;

export { JWT_PRIVATE_KEY, JWT_PUBLIC_KEY, TOKEN_LIFETIME, COOKIE_LIFETIME };