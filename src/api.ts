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
  // TODO: body validation
  const info: AgentInfo = req.body;
  const requestIp = getRequestIp(req);

  // Remove the old agent, register the new one.
  removeAgent(requestIp, info.address);
  registerAgent(requestIp, info);

  res.status(201);
  apiGetAgents(req, res);
}

export function apiDeleteAgent(req: Request, res: Response<AgentInfo[]>): void {
  removeAgent(getRequestIp(req), req.params.address);
  apiGetAgents(req, res);
}
