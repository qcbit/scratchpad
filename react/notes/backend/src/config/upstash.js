import { Ratelimit } from "@upstash/ratelimit";
import { Redis } from "@upstash/redis";

import dotenv from 'dotenv';

dotenv.config();

// Configure Upstash Redis and Rate Limit
const ratelimit = new Ratelimit({
  redis: Redis.fromEnv(),
  limiter: Ratelimit.slidingWindow(100, '60 s'), // 100 requests per 1 minute
});

export default ratelimit;