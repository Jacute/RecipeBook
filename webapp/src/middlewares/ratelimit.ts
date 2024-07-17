import rateLimit from "express-rate-limit"
import * as config from "../config";


const limiter = rateLimit({
    windowMs: config.RATELIMIT_PERIOD,
    max: config.RATELIMIT_COUNT,
    statusCode: 400
})

export { limiter };