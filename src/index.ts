import { Logger } from '@luca_scorpion/tinylogger';
import { AddressInfo } from 'net';
import { PORT } from './constants';
import { app } from './app';

const log = new Logger('index');
log.debug('Starting discovery server...');

async function bootstrap(): Promise<void> {
  // Start the server.
  const server = app.listen(PORT, () => {
    const serverPort = (server.address() as AddressInfo).port;
    log.info(`Discovery server running on port ${serverPort}`);
  });
}

bootstrap().catch((err) => {
  log.error('An unexpected error occurred:');
  log.error(err.stack);
});
