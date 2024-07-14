import jwt, { Algorithm } from 'jsonwebtoken';
import * as config from '../config';


interface VerifyTokenResult {
    tokenExp: boolean;
    error: any;
    decoded: any;
}


function encodeToken(username: string): string {
    const payload = {
        username: username
    }
    const options = {
        algorithm: "RS256" as Algorithm,
        expiresIn: config.TOKEN_LIFETIME
    };
    const token = jwt.sign(
        payload,
        config.JWT_PRIVATE_KEY,
        options
    );
    return token;
}

function verifyToken(token: string): VerifyTokenResult {
    const options = {
        algorithms: ["RS256"] as Algorithm[],
        maxAge: config.TOKEN_LIFETIME
    }
    try {
        const decoded = jwt.verify(token, config.JWT_PUBLIC_KEY, options);
        return { tokenExp: false, error: null, decoded: decoded };
    } catch (err) {
        return { tokenExp: true, error: err, decoded: null };
    }
}

export { encodeToken, verifyToken };