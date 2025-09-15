import grpc from "@grpc/grpc-js"
import protoLoader from "@grpc/proto-loader"

import type { ProtoGrpcType } from "../proto/session.ts"
import path from "path"
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const PROTO_PATH = path.join(__dirname, '..', 'proto', 'session.proto')
const packageDefinition = protoLoader.loadSync(PROTO_PATH)
const grpcObj =(grpc.loadPackageDefinition(packageDefinition)) as unknown as ProtoGrpcType
const sessionProto = grpcObj.session

export const grpcClient = new sessionProto.SessionService("127.0.0.1:5003",grpc.credentials.createInsecure())