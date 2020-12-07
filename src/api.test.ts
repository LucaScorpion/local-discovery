// A default import does not work with supertest.
// eslint-disable-next-line import/no-named-default
import { default as request } from 'supertest';
import { app } from './app';
import { registerAgent, removeAgent } from './agentRegistry';
import { AgentInfo } from './AgentInfo';

function randomAgent(): AgentInfo {
  return {
    name: Math.random().toString().substring(2),
    version: Math.round(Math.random() * 10).toString(),
    address: Math.random().toString().substring(2),
    platform: 'test',
    hostname: Math.random().toString().substring(2),
  };
}

describe('api', (): void => {
  describe('GET /api/agents', (): void => {
    it('returns an empty list if no agents are registered', async (): Promise<void> => {
      await request(app)
        .get('/api/agents')
        .expect(200)
        .expect('Content-Type', 'application/json; charset=utf-8')
        .expect([]);
    });

    it('returns all registered agents for the request ip', async (): Promise<void> => {
      const agentOne = randomAgent();
      registerAgent('::1', agentOne);
      const agentTwo = randomAgent();
      registerAgent('::1', agentTwo);
      const agentThree = randomAgent();
      registerAgent('6.7.8.9', agentThree);

      try {
        await request(app)
          .get('/api/agents')
          .expect(200)
          .expect('Content-Type', 'application/json; charset=utf-8')
          .expect([agentTwo, agentOne]);
      } finally {
        removeAgent('::1', agentOne.address);
        removeAgent('::1', agentTwo.address);
        removeAgent('6.7.8.9', agentThree.address);
      }
    });
  });

  describe('POST /api/agents', () => {
    it('registers a new agent and returns all agents', async (): Promise<void> => {
      const agent = randomAgent();

      try {
        await request(app)
          .post('/api/agents')
          .send(agent)
          .expect(201)
          .expect('Content-Type', 'application/json; charset=utf-8')
          .expect([agent]);
      } finally {
        removeAgent('::1', agent.address);
      }
    });

    it('overwrites an existing agent with the same address', async (): Promise<void> => {
      const oldAgent = randomAgent();
      registerAgent('::1', oldAgent);
      const newAgent = randomAgent();
      newAgent.address = oldAgent.address;
      const otherAgent = randomAgent();
      registerAgent('::1', otherAgent);

      try {
        await request(app)
          .post('/api/agents')
          .send(newAgent)
          .expect(201)
          .expect('Content-Type', 'application/json; charset=utf-8')
          .expect([newAgent, otherAgent]);
      } finally {
        removeAgent('::1', newAgent.address);
        removeAgent('::1', otherAgent.address);
      }
    });
  });

  describe('DELETE /api/agents/:address', () => {
    it('removes the agent', async (): Promise<void> => {
      const agent = randomAgent();
      registerAgent('::1', agent);
      const remove = randomAgent();
      registerAgent('::1', remove);

      try {
        await request(app).get('/api/agents').expect([remove, agent]);

        await request(app)
          .del(`/api/agents/${remove.address}`)
          .expect(200)
          .expect('Content-Type', 'application/json; charset=utf-8')
          .expect([agent]);
      } finally {
        removeAgent('::1', agent.address);
      }
    });

    it('does nothing if no agent with that address exists', async (): Promise<void> => {
      await request(app)
        .del('/api/agents/does.not.exist')
        .expect(200)
        .expect('Content-Type', 'application/json; charset=utf-8')
        .expect([]);
    });
  });
});
