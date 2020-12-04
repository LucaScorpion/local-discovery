import { Request } from 'express';

/**
 * Get the request IP from a request.
 * This mainly ensures that the IP address is always the same when coming from localhost,
 * since this can be either an IPv4 or IPv6 address depending on the client.
 *
 * @param req The request to get the IP from.
 * @return The request IP.
 */
export function getRequestIp(req: Request): string {
  // Make sure the IP address is always the same when coming from localhost.
  return req.ip === '::1' ? '::ffff:127.0.0.1' : req.ip;
}
