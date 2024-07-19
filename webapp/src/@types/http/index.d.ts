import * as http from 'http';

declare module 'http' {
  interface ServerResponse {
    locals: {
      nonce?: string;
    };
  }
}