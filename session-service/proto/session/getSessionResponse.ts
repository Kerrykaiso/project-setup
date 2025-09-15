// Original file: proto/session.proto

import type { Timestamp as _google_protobuf_Timestamp, Timestamp__Output as _google_protobuf_Timestamp__Output } from '../google/protobuf/Timestamp';

export interface getSessionResponse {
  'sessionId'?: (string);
  'expiresAt'?: (_google_protobuf_Timestamp | null);
  'userId'?: (string);
}

export interface getSessionResponse__Output {
  'sessionId'?: (string);
  'expiresAt'?: (_google_protobuf_Timestamp__Output);
  'userId'?: (string);
}
