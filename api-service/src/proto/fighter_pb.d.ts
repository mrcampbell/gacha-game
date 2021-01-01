// package: fighter
// file: fighter.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class Unit extends jspb.Message { 
    getId(): string;
    setId(value: string): Unit;

    getName(): string;
    setName(value: string): Unit;

    getElement(): Element;
    setElement(value: Element): Unit;

    getType(): UnitType;
    setType(value: UnitType): Unit;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Unit.AsObject;
    static toObject(includeInstance: boolean, msg: Unit): Unit.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Unit, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Unit;
    static deserializeBinaryFromReader(message: Unit, reader: jspb.BinaryReader): Unit;
}

export namespace Unit {
    export type AsObject = {
        id: string,
        name: string,
        element: Element,
        type: UnitType,
    }
}

export class UnitByIDRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): UnitByIDRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UnitByIDRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UnitByIDRequest): UnitByIDRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UnitByIDRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UnitByIDRequest;
    static deserializeBinaryFromReader(message: UnitByIDRequest, reader: jspb.BinaryReader): UnitByIDRequest;
}

export namespace UnitByIDRequest {
    export type AsObject = {
        id: string,
    }
}

export class UnitByIDResponse extends jspb.Message { 

    hasUnit(): boolean;
    clearUnit(): void;
    getUnit(): Unit | undefined;
    setUnit(value?: Unit): UnitByIDResponse;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UnitByIDResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UnitByIDResponse): UnitByIDResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UnitByIDResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UnitByIDResponse;
    static deserializeBinaryFromReader(message: UnitByIDResponse, reader: jspb.BinaryReader): UnitByIDResponse;
}

export namespace UnitByIDResponse {
    export type AsObject = {
        unit?: Unit.AsObject,
    }
}

export enum UnitType {
    UNIT_TYPE_UNKNOWN = 0,
    UNIT_TYPE_TANK = 1,
    UNIT_TYPE_HEALER = 2,
    UNIT_TYPE_SUPPORT = 3,
    UNIT_TYPE_ATTACKER = 4,
}

export enum Element {
    ELEMENT_UNKNOWN = 0,
    ELEMENT_FIRE = 1,
    ELEMENT_WATER = 2,
    ELEMENT_GRASS = 3,
}
