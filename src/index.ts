import express from 'express';
import { Logger } from '@luca_scorpion/tinylogger';
import { AddressInfo } from 'net';
import { PORT } from './constants';
import { apiGetAgents, apiRegisterAgent } from './api';
import { cors } from './cors';

const log = new Logger('index');
log.debug('Starting discovery server...');

async function bootstrap(): Promise<void> {
  // Start express.
  const app = express();
  app.use(express.json());
  app.use(cors);

  // API routes.
  app.get('/api/agents', apiGetAgents);
  app.post('/api/agents', apiRegisterAgent);

  // Done!
  const server = app.listen(PORT, () => {
    const serverPort = (server.address() as AddressInfo).port;
    log.info(`Discovery server running on port ${serverPort}`);
  });
}

bootstrap().catch((err) => {
  log.error('An unexpected error occurred:');
  log.error(err.stack);
});
