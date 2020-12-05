import express from 'express';
import { cors } from './cors';
import { apiDeleteAgent, apiGetAgents, apiRegisterAgent } from './api';

// Set up express.
export const app = express();
app.use(express.json());
app.use(cors);

// API routes.
app.get('/api/agents', apiGetAgents);
app.post('/api/agents', apiRegisterAgent);
app.delete('/api/agents/:address', apiDeleteAgent);
