import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { Timestamp as _google_protobuf_Timestamp, Timestamp__Output as _google_protobuf_Timestamp__Output } from './google/protobuf/Timestamp.ts';
import type { SessionRequest as _session_SessionRequest, SessionRequest__Output as _session_SessionRequest__Output } from './session/SessionRequest.ts';
import type { SessionResponse as _session_SessionResponse, SessionResponse__Output as _session_SessionResponse__Output } from './session/SessionResponse.ts';
import type { SessionServiceClient as _session_SessionServiceClient, SessionServiceDefinition as _session_SessionServiceDefinition } from './session/SessionService.ts';
import type { deleteSessionRequest as _session_deleteSessionRequest, deleteSessionRequest__Output as _session_deleteSessionRequest__Output } from './session/deleteSessionRequest.ts';
import type { deleteSessionResponse as _session_deleteSessionResponse, deleteSessionResponse__Output as _session_deleteSessionResponse__Output } from './session/deleteSessionResponse.ts';
import type { getSessionRequest as _session_getSessionRequest, getSessionRequest__Output as _session_getSessionRequest__Output } from './session/getSessionRequest.ts';
import type { getSessionResponse as _session_getSessionResponse, getSessionResponse__Output as _session_getSessionResponse__Output } from './session/getSessionResponse.ts';
import type { updateSessionRequest as _session_updateSessionRequest, updateSessionRequest__Output as _session_updateSessionRequest__Output } from './session/updateSessionRequest.ts';
import type { updateSessionResponse as _session_updateSessionResponse, updateSessionResponse__Output as _session_updateSessionResponse__Output } from './session/updateSessionResponse.ts';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  google: {
    protobuf: {
      Timestamp: MessageTypeDefinition<_google_protobuf_Timestamp, _google_protobuf_Timestamp__Output>
    }
  }
  session: {
    SessionRequest: MessageTypeDefinition<_session_SessionRequest, _session_SessionRequest__Output>
    SessionResponse: MessageTypeDefinition<_session_SessionResponse, _session_SessionResponse__Output>
    SessionService: SubtypeConstructor<typeof grpc.Client, _session_SessionServiceClient> & { service: _session_SessionServiceDefinition }
    deleteSessionRequest: MessageTypeDefinition<_session_deleteSessionRequest, _session_deleteSessionRequest__Output>
    deleteSessionResponse: MessageTypeDefinition<_session_deleteSessionResponse, _session_deleteSessionResponse__Output>
    getSessionRequest: MessageTypeDefinition<_session_getSessionRequest, _session_getSessionRequest__Output>
    getSessionResponse: MessageTypeDefinition<_session_getSessionResponse, _session_getSessionResponse__Output>
    updateSessionRequest: MessageTypeDefinition<_session_updateSessionRequest, _session_updateSessionRequest__Output>
    updateSessionResponse: MessageTypeDefinition<_session_updateSessionResponse, _session_updateSessionResponse__Output>
  }
}

