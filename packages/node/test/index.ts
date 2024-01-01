import { Hono } from 'hono';
import { serve } from '@hono/node-server';
import { s } from '../src';
const app = new Hono();

app.get('/', (c) => c.text('Test Server'));

s(() => {
    serve({ fetch: app.fetch, port: Number(process.env.port!) });
}, { env: 'test' });
