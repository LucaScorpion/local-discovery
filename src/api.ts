import { Request, Response } from 'express';
import { AgentInfo } from './AgentInfo';
import { getAgents, registerAgent, removeAgent } from './agentRegistry';
import { getRequestIp } from './getRequestIp';

export function apiGetAgents(req: Request, res: Response<AgentInfo[]>): void {
  res.json(getAgents(getRequestIp(req)));
}

export function apiRegisterAgent(
  req: Request,
  res: Response<AgentInfo[]>
): void {
  const info: AgentInfo = req.body;
  const requestIp = getRequestIp(req);

  // Remove the old agent, register the new one.
  removeAgent(requestIp, info.address);
  registerAgent(requestIp, info);

  apiGetAgents(req, res);
}
