// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var fighter_pb = require('./fighter_pb.js');

function serialize_fighter_UnitByIDRequest(arg) {
  if (!(arg instanceof fighter_pb.UnitByIDRequest)) {
    throw new Error('Expected argument of type fighter.UnitByIDRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_fighter_UnitByIDRequest(buffer_arg) {
  return fighter_pb.UnitByIDRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_fighter_UnitByIDResponse(arg) {
  if (!(arg instanceof fighter_pb.UnitByIDResponse)) {
    throw new Error('Expected argument of type fighter.UnitByIDResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_fighter_UnitByIDResponse(buffer_arg) {
  return fighter_pb.UnitByIDResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var FighterServiceService = exports.FighterServiceService = {
  unitByID: {
    path: '/fighter.FighterService/UnitByID',
    requestStream: false,
    responseStream: false,
    requestType: fighter_pb.UnitByIDRequest,
    responseType: fighter_pb.UnitByIDResponse,
    requestSerialize: serialize_fighter_UnitByIDRequest,
    requestDeserialize: deserialize_fighter_UnitByIDRequest,
    responseSerialize: serialize_fighter_UnitByIDResponse,
    responseDeserialize: deserialize_fighter_UnitByIDResponse,
  },
};

exports.FighterServiceClient = grpc.makeGenericClientConstructor(FighterServiceService);
