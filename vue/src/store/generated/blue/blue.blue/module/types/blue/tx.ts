/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Pool } from "../blue/pool";

export const protobufPackage = "blue.blue";

export interface MsgCreatePool {
  creator: string;
  title: string;
  question: string;
  answerRange: number;
}

export interface MsgCreatePoolResponse {
  id: number;
}

export interface MsgVote {
  creator: string;
  poolId: string;
  answer: string;
}

export interface MsgVoteResponse {
  pool: Pool | undefined;
}

const baseMsgCreatePool: object = {
  creator: "",
  title: "",
  question: "",
  answerRange: 0,
};

export const MsgCreatePool = {
  encode(message: MsgCreatePool, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.question !== "") {
      writer.uint32(26).string(message.question);
    }
    if (message.answerRange !== 0) {
      writer.uint32(32).uint64(message.answerRange);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreatePool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreatePool } as MsgCreatePool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.title = reader.string();
          break;
        case 3:
          message.question = reader.string();
          break;
        case 4:
          message.answerRange = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreatePool {
    const message = { ...baseMsgCreatePool } as MsgCreatePool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.question !== undefined && object.question !== null) {
      message.question = String(object.question);
    } else {
      message.question = "";
    }
    if (object.answerRange !== undefined && object.answerRange !== null) {
      message.answerRange = Number(object.answerRange);
    } else {
      message.answerRange = 0;
    }
    return message;
  },

  toJSON(message: MsgCreatePool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.title !== undefined && (obj.title = message.title);
    message.question !== undefined && (obj.question = message.question);
    message.answerRange !== undefined &&
      (obj.answerRange = message.answerRange);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreatePool>): MsgCreatePool {
    const message = { ...baseMsgCreatePool } as MsgCreatePool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.question !== undefined && object.question !== null) {
      message.question = object.question;
    } else {
      message.question = "";
    }
    if (object.answerRange !== undefined && object.answerRange !== null) {
      message.answerRange = object.answerRange;
    } else {
      message.answerRange = 0;
    }
    return message;
  },
};

const baseMsgCreatePoolResponse: object = { id: 0 };

export const MsgCreatePoolResponse = {
  encode(
    message: MsgCreatePoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreatePoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreatePoolResponse } as MsgCreatePoolResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreatePoolResponse {
    const message = { ...baseMsgCreatePoolResponse } as MsgCreatePoolResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreatePoolResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreatePoolResponse>
  ): MsgCreatePoolResponse {
    const message = { ...baseMsgCreatePoolResponse } as MsgCreatePoolResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgVote: object = { creator: "", poolId: "", answer: "" };

export const MsgVote = {
  encode(message: MsgVote, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.poolId !== "") {
      writer.uint32(18).string(message.poolId);
    }
    if (message.answer !== "") {
      writer.uint32(26).string(message.answer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgVote {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgVote } as MsgVote;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.poolId = reader.string();
          break;
        case 3:
          message.answer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgVote {
    const message = { ...baseMsgVote } as MsgVote;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.poolId !== undefined && object.poolId !== null) {
      message.poolId = String(object.poolId);
    } else {
      message.poolId = "";
    }
    if (object.answer !== undefined && object.answer !== null) {
      message.answer = String(object.answer);
    } else {
      message.answer = "";
    }
    return message;
  },

  toJSON(message: MsgVote): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.poolId !== undefined && (obj.poolId = message.poolId);
    message.answer !== undefined && (obj.answer = message.answer);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgVote>): MsgVote {
    const message = { ...baseMsgVote } as MsgVote;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.poolId !== undefined && object.poolId !== null) {
      message.poolId = object.poolId;
    } else {
      message.poolId = "";
    }
    if (object.answer !== undefined && object.answer !== null) {
      message.answer = object.answer;
    } else {
      message.answer = "";
    }
    return message;
  },
};

const baseMsgVoteResponse: object = {};

export const MsgVoteResponse = {
  encode(message: MsgVoteResponse, writer: Writer = Writer.create()): Writer {
    if (message.pool !== undefined) {
      Pool.encode(message.pool, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgVoteResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgVoteResponse } as MsgVoteResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pool = Pool.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgVoteResponse {
    const message = { ...baseMsgVoteResponse } as MsgVoteResponse;
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = Pool.fromJSON(object.pool);
    } else {
      message.pool = undefined;
    }
    return message;
  },

  toJSON(message: MsgVoteResponse): unknown {
    const obj: any = {};
    message.pool !== undefined &&
      (obj.pool = message.pool ? Pool.toJSON(message.pool) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgVoteResponse>): MsgVoteResponse {
    const message = { ...baseMsgVoteResponse } as MsgVoteResponse;
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = Pool.fromPartial(object.pool);
    } else {
      message.pool = undefined;
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreatePool(request: MsgCreatePool): Promise<MsgCreatePoolResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Vote(request: MsgVote): Promise<MsgVoteResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreatePool(request: MsgCreatePool): Promise<MsgCreatePoolResponse> {
    const data = MsgCreatePool.encode(request).finish();
    const promise = this.rpc.request("blue.blue.Msg", "CreatePool", data);
    return promise.then((data) =>
      MsgCreatePoolResponse.decode(new Reader(data))
    );
  }

  Vote(request: MsgVote): Promise<MsgVoteResponse> {
    const data = MsgVote.encode(request).finish();
    const promise = this.rpc.request("blue.blue.Msg", "Vote", data);
    return promise.then((data) => MsgVoteResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
