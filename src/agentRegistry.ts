import { AgentInfo } from './AgentInfo';
import { KEEP_AGENT_TIME } from './constants';

interface AgentWithTimeout extends AgentInfo {
  removalTimeout: NodeJS.Timeout;
}

/**
 * The agent registry.
 * This contains all known agents by public IP.
 */
const agents: { [publicIp: string]: AgentWithTimeout[] } = {};

/**
 * Get all known agents for a public IP.
 *
 * @param publicIp The public IP.
 */
export function getAgents(publicIp: string): AgentInfo[] {
  return (agents[publicIp] || []).map(
    // Remove the removalTimeout from the agent object.
    ({ name, version, address, platform, hostname }) => ({
      name,
      version,
      address,
      platform,
      hostname,
    })
  );
}

/**
 * Remove an agent from the registry and clear its removal timeout.
 *
 * @param publicIp The public IP.
 * @param address The local address of the agent to remove.
 */
export function removeAgent(publicIp: string, address: string): void {
  const agentList = agents[publicIp] || [];

  for (let i = 0; i < agentList.length; i++) {
    if (agentList[i].address === address) {
      const removed = agentList.splice(i, 1)[0];
      agents[publicIp] = agentList;
      clearTimeout(removed.removalTimeout);
    }
  }
}

/**
 * Register a new agent.
 *
 * @param publicIp The public IP.
 * @param agent The agent to register.
 */
export function registerAgent(publicIp: string, agent: AgentInfo): void {
  const agentWithTimeout = {
    ...agent,
    removalTimeout: setTimeout(
      () => removeAgent(publicIp, agent.address),
      KEEP_AGENT_TIME * 1000
    ),
  };

  // Add the new agent to the list, store it.
  const currentAgents = agents[publicIp] || [];
  agents[publicIp] = [agentWithTimeout, ...currentAgents];
}
