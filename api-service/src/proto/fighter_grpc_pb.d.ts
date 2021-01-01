// package: fighter
// file: fighter.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as fighter_pb from "./fighter_pb";

interface IFighterServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    unitByID: IFighterServiceService_IUnitByID;
}

interface IFighterServiceService_IUnitByID extends grpc.MethodDefinition<fighter_pb.UnitByIDRequest, fighter_pb.UnitByIDResponse> {
    path: "/fighter.FighterService/UnitByID";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<fighter_pb.UnitByIDRequest>;
    requestDeserialize: grpc.deserialize<fighter_pb.UnitByIDRequest>;
    responseSerialize: grpc.serialize<fighter_pb.UnitByIDResponse>;
    responseDeserialize: grpc.deserialize<fighter_pb.UnitByIDResponse>;
}

export const FighterServiceService: IFighterServiceService;

export interface IFighterServiceServer {
    unitByID: grpc.handleUnaryCall<fighter_pb.UnitByIDRequest, fighter_pb.UnitByIDResponse>;
}

export interface IFighterServiceClient {
    unitByID(request: fighter_pb.UnitByIDRequest, callback: (error: grpc.ServiceError | null, response: fighter_pb.UnitByIDResponse) => void): grpc.ClientUnaryCall;
    unitByID(request: fighter_pb.UnitByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: fighter_pb.UnitByIDResponse) => void): grpc.ClientUnaryCall;
    unitByID(request: fighter_pb.UnitByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: fighter_pb.UnitByIDResponse) => void): grpc.ClientUnaryCall;
}

export class FighterServiceClient extends grpc.Client implements IFighterServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public unitByID(request: fighter_pb.UnitByIDRequest, callback: (error: grpc.ServiceError | null, response: fighter_pb.UnitByIDResponse) => void): grpc.ClientUnaryCall;
    public unitByID(request: fighter_pb.UnitByIDRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: fighter_pb.UnitByIDResponse) => void): grpc.ClientUnaryCall;
    public unitByID(request: fighter_pb.UnitByIDRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: fighter_pb.UnitByIDResponse) => void): grpc.ClientUnaryCall;
}
