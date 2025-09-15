// Original file: proto/session.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { SessionRequest as _session_SessionRequest, SessionRequest__Output as _session_SessionRequest__Output } from '../session/SessionRequest';
import type { SessionResponse as _session_SessionResponse, SessionResponse__Output as _session_SessionResponse__Output } from '../session/SessionResponse';
import type { deleteSessionRequest as _session_deleteSessionRequest, deleteSessionRequest__Output as _session_deleteSessionRequest__Output } from '../session/deleteSessionRequest';
import type { deleteSessionResponse as _session_deleteSessionResponse, deleteSessionResponse__Output as _session_deleteSessionResponse__Output } from '../session/deleteSessionResponse';
import type { getSessionRequest as _session_getSessionRequest, getSessionRequest__Output as _session_getSessionRequest__Output } from '../session/getSessionRequest';
import type { getSessionResponse as _session_getSessionResponse, getSessionResponse__Output as _session_getSessionResponse__Output } from '../session/getSessionResponse';
import type { updateSessionRequest as _session_updateSessionRequest, updateSessionRequest__Output as _session_updateSessionRequest__Output } from '../session/updateSessionRequest';
import type { updateSessionResponse as _session_updateSessionResponse, updateSessionResponse__Output as _session_updateSessionResponse__Output } from '../session/updateSessionResponse';

export interface SessionServiceClient extends grpc.Client {
  CreateSession(argument: _session_SessionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_session_SessionResponse__Output>): grpc.ClientUnaryCall;
  CreateSession(argument: _session_SessionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_session_SessionResponse__Output>): grpc.ClientUnaryCall;
  CreateSession(argument: _session_SessionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_session_SessionResponse__Output>): grpc.ClientUnaryCall;
  CreateSession(argument: _session_SessionRequest, callback: grpc.requestCallback<_session_SessionResponse__Output>): grpc.ClientUnaryCall;
  createSession(argument: _session_SessionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_session_SessionResponse__Output>): grpc.ClientUnaryCall;
  createSession(argument: _session_SessionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_session_SessionResponse__Output>): grpc.ClientUnaryCall;
  createSession(argument: _session_SessionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_session_SessionResponse__Output>): grpc.ClientUnaryCall;
  createSession(argument: _session_SessionRequest, callback: grpc.requestCallback<_session_SessionResponse__Output>): grpc.ClientUnaryCall;
  
  deleteSession(argument: _session_deleteSessionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_session_deleteSessionResponse__Output>): grpc.ClientUnaryCall;
  deleteSession(argument: _session_deleteSessionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_session_deleteSessionResponse__Output>): grpc.ClientUnaryCall;
  deleteSession(argument: _session_deleteSessionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_session_deleteSessionResponse__Output>): grpc.ClientUnaryCall;
  deleteSession(argument: _session_deleteSessionRequest, callback: grpc.requestCallback<_session_deleteSessionResponse__Output>): grpc.ClientUnaryCall;
  
  getSession(argument: _session_getSessionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_session_getSessionResponse__Output>): grpc.ClientUnaryCall;
  getSession(argument: _session_getSessionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_session_getSessionResponse__Output>): grpc.ClientUnaryCall;
  getSession(argument: _session_getSessionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_session_getSessionResponse__Output>): grpc.ClientUnaryCall;
  getSession(argument: _session_getSessionRequest, callback: grpc.requestCallback<_session_getSessionResponse__Output>): grpc.ClientUnaryCall;
  
  updateSession(argument: _session_updateSessionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_session_updateSessionResponse__Output>): grpc.ClientUnaryCall;
  updateSession(argument: _session_updateSessionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_session_updateSessionResponse__Output>): grpc.ClientUnaryCall;
  updateSession(argument: _session_updateSessionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_session_updateSessionResponse__Output>): grpc.ClientUnaryCall;
  updateSession(argument: _session_updateSessionRequest, callback: grpc.requestCallback<_session_updateSessionResponse__Output>): grpc.ClientUnaryCall;
  
}

export interface SessionServiceHandlers extends grpc.UntypedServiceImplementation {
  CreateSession: grpc.handleUnaryCall<_session_SessionRequest__Output, _session_SessionResponse>;
  
  deleteSession: grpc.handleUnaryCall<_session_deleteSessionRequest__Output, _session_deleteSessionResponse>;
  
  getSession: grpc.handleUnaryCall<_session_getSessionRequest__Output, _session_getSessionResponse>;
  
  updateSession: grpc.handleUnaryCall<_session_updateSessionRequest__Output, _session_updateSessionResponse>;
  
}

export interface SessionServiceDefinition extends grpc.ServiceDefinition {
  CreateSession: MethodDefinition<_session_SessionRequest, _session_SessionResponse, _session_SessionRequest__Output, _session_SessionResponse__Output>
  deleteSession: MethodDefinition<_session_deleteSessionRequest, _session_deleteSessionResponse, _session_deleteSessionRequest__Output, _session_deleteSessionResponse__Output>
  getSession: MethodDefinition<_session_getSessionRequest, _session_getSessionResponse, _session_getSessionRequest__Output, _session_getSessionResponse__Output>
  updateSession: MethodDefinition<_session_updateSessionRequest, _session_updateSessionResponse, _session_updateSessionRequest__Output, _session_updateSessionResponse__Output>
}
