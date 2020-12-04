// A default import does not work with supertest.
// eslint-disable-next-line import/no-named-default
import { default as request } from 'supertest';
import { app } from './app';
import { registerAgent, removeAgent } from './agentRegistry';
import { AgentInfo } from './AgentInfo';

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
      const agentOne: AgentInfo = {
        name: 'one',
        version: '1.1.1',
        address: '192.168.1.1',
        platform: 'linux',
        hostname: 'test',
      };
      registerAgent('::1', agentOne);
      const agentTwo: AgentInfo = {
        name: 'two',
        version: '1.2.0',
        address: '192.168.1.2',
        platform: 'windows',
        hostname: 'another',
      };
      registerAgent('::1', agentTwo);
      const agentThree: AgentInfo = {
        name: 'three',
        version: '2.0.2',
        address: '10.0.0.10',
        platform: 'mac',
        hostname: 'nope',
      };
      registerAgent('6.7.8.9', agentThree);

      try {
        await request(app)
          .get('/api/agents')
          .expect(200)
          .expect('Content-Type', 'application/json; charset=utf-8')
          .expect([agentOne, agentTwo]);
      } finally {
        removeAgent('::1', '192.168.1.1');
        removeAgent('::1', '192.168.1.2');
        removeAgent('6.7.8.9', '10.0.0.10');
      }
    });
  });
});
