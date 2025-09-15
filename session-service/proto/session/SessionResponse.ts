// Original file: proto/session.proto

import type { Timestamp as _google_protobuf_Timestamp, Timestamp__Output as _google_protobuf_Timestamp__Output } from '../google/protobuf/Timestamp';

export interface SessionResponse {
  'sessionId'?: (string);
  'expiresAt'?: (_google_protobuf_Timestamp | null);
}

export interface SessionResponse__Output {
  'sessionId'?: (string);
  'expiresAt'?: (_google_protobuf_Timestamp__Output);
}
