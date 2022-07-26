/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../blue/params";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Pool } from "../blue/pool";
import { PrefPool } from "../blue/pref_pool";

export const protobufPackage = "blue.blue";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryPoolsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryPoolsResponse {
  Pool: Pool[];
  pagination: PageResponse | undefined;
}

export interface QueryShowpoolRequest {
  id: number;
}

export interface QueryShowpoolResponse {
  Pool: Pool | undefined;
}

export interface QueryGetPrefPoolRequest {
  id: number;
}

export interface QueryGetPrefPoolResponse {
  PrefPool: PrefPool | undefined;
}

export interface QueryAllPrefPoolRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllPrefPoolResponse {
  PrefPool: PrefPool[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryPoolsRequest: object = {};

export const QueryPoolsRequest = {
  encode(message: QueryPoolsRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryPoolsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryPoolsRequest } as QueryPoolsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryPoolsRequest {
    const message = { ...baseQueryPoolsRequest } as QueryPoolsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryPoolsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryPoolsRequest>): QueryPoolsRequest {
    const message = { ...baseQueryPoolsRequest } as QueryPoolsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryPoolsResponse: object = {};

export const QueryPoolsResponse = {
  encode(
    message: QueryPoolsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Pool) {
      Pool.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryPoolsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryPoolsResponse } as QueryPoolsResponse;
    message.Pool = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Pool.push(Pool.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryPoolsResponse {
    const message = { ...baseQueryPoolsResponse } as QueryPoolsResponse;
    message.Pool = [];
    if (object.Pool !== undefined && object.Pool !== null) {
      for (const e of object.Pool) {
        message.Pool.push(Pool.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryPoolsResponse): unknown {
    const obj: any = {};
    if (message.Pool) {
      obj.Pool = message.Pool.map((e) => (e ? Pool.toJSON(e) : undefined));
    } else {
      obj.Pool = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryPoolsResponse>): QueryPoolsResponse {
    const message = { ...baseQueryPoolsResponse } as QueryPoolsResponse;
    message.Pool = [];
    if (object.Pool !== undefined && object.Pool !== null) {
      for (const e of object.Pool) {
        message.Pool.push(Pool.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryShowpoolRequest: object = { id: 0 };

export const QueryShowpoolRequest = {
  encode(
    message: QueryShowpoolRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryShowpoolRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryShowpoolRequest } as QueryShowpoolRequest;
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

  fromJSON(object: any): QueryShowpoolRequest {
    const message = { ...baseQueryShowpoolRequest } as QueryShowpoolRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryShowpoolRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryShowpoolRequest>): QueryShowpoolRequest {
    const message = { ...baseQueryShowpoolRequest } as QueryShowpoolRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryShowpoolResponse: object = {};

export const QueryShowpoolResponse = {
  encode(
    message: QueryShowpoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Pool !== undefined) {
      Pool.encode(message.Pool, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryShowpoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryShowpoolResponse } as QueryShowpoolResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Pool = Pool.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryShowpoolResponse {
    const message = { ...baseQueryShowpoolResponse } as QueryShowpoolResponse;
    if (object.Pool !== undefined && object.Pool !== null) {
      message.Pool = Pool.fromJSON(object.Pool);
    } else {
      message.Pool = undefined;
    }
    return message;
  },

  toJSON(message: QueryShowpoolResponse): unknown {
    const obj: any = {};
    message.Pool !== undefined &&
      (obj.Pool = message.Pool ? Pool.toJSON(message.Pool) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryShowpoolResponse>
  ): QueryShowpoolResponse {
    const message = { ...baseQueryShowpoolResponse } as QueryShowpoolResponse;
    if (object.Pool !== undefined && object.Pool !== null) {
      message.Pool = Pool.fromPartial(object.Pool);
    } else {
      message.Pool = undefined;
    }
    return message;
  },
};

const baseQueryGetPrefPoolRequest: object = { id: 0 };

export const QueryGetPrefPoolRequest = {
  encode(
    message: QueryGetPrefPoolRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPrefPoolRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPrefPoolRequest,
    } as QueryGetPrefPoolRequest;
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

  fromJSON(object: any): QueryGetPrefPoolRequest {
    const message = {
      ...baseQueryGetPrefPoolRequest,
    } as QueryGetPrefPoolRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetPrefPoolRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPrefPoolRequest>
  ): QueryGetPrefPoolRequest {
    const message = {
      ...baseQueryGetPrefPoolRequest,
    } as QueryGetPrefPoolRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetPrefPoolResponse: object = {};

export const QueryGetPrefPoolResponse = {
  encode(
    message: QueryGetPrefPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.PrefPool !== undefined) {
      PrefPool.encode(message.PrefPool, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetPrefPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPrefPoolResponse,
    } as QueryGetPrefPoolResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.PrefPool = PrefPool.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPrefPoolResponse {
    const message = {
      ...baseQueryGetPrefPoolResponse,
    } as QueryGetPrefPoolResponse;
    if (object.PrefPool !== undefined && object.PrefPool !== null) {
      message.PrefPool = PrefPool.fromJSON(object.PrefPool);
    } else {
      message.PrefPool = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetPrefPoolResponse): unknown {
    const obj: any = {};
    message.PrefPool !== undefined &&
      (obj.PrefPool = message.PrefPool
        ? PrefPool.toJSON(message.PrefPool)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPrefPoolResponse>
  ): QueryGetPrefPoolResponse {
    const message = {
      ...baseQueryGetPrefPoolResponse,
    } as QueryGetPrefPoolResponse;
    if (object.PrefPool !== undefined && object.PrefPool !== null) {
      message.PrefPool = PrefPool.fromPartial(object.PrefPool);
    } else {
      message.PrefPool = undefined;
    }
    return message;
  },
};

const baseQueryAllPrefPoolRequest: object = {};

export const QueryAllPrefPoolRequest = {
  encode(
    message: QueryAllPrefPoolRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPrefPoolRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllPrefPoolRequest,
    } as QueryAllPrefPoolRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPrefPoolRequest {
    const message = {
      ...baseQueryAllPrefPoolRequest,
    } as QueryAllPrefPoolRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPrefPoolRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPrefPoolRequest>
  ): QueryAllPrefPoolRequest {
    const message = {
      ...baseQueryAllPrefPoolRequest,
    } as QueryAllPrefPoolRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllPrefPoolResponse: object = {};

export const QueryAllPrefPoolResponse = {
  encode(
    message: QueryAllPrefPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.PrefPool) {
      PrefPool.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllPrefPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllPrefPoolResponse,
    } as QueryAllPrefPoolResponse;
    message.PrefPool = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.PrefPool.push(PrefPool.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPrefPoolResponse {
    const message = {
      ...baseQueryAllPrefPoolResponse,
    } as QueryAllPrefPoolResponse;
    message.PrefPool = [];
    if (object.PrefPool !== undefined && object.PrefPool !== null) {
      for (const e of object.PrefPool) {
        message.PrefPool.push(PrefPool.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPrefPoolResponse): unknown {
    const obj: any = {};
    if (message.PrefPool) {
      obj.PrefPool = message.PrefPool.map((e) =>
        e ? PrefPool.toJSON(e) : undefined
      );
    } else {
      obj.PrefPool = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPrefPoolResponse>
  ): QueryAllPrefPoolResponse {
    const message = {
      ...baseQueryAllPrefPoolResponse,
    } as QueryAllPrefPoolResponse;
    message.PrefPool = [];
    if (object.PrefPool !== undefined && object.PrefPool !== null) {
      for (const e of object.PrefPool) {
        message.PrefPool.push(PrefPool.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Pools items. */
  Pools(request: QueryPoolsRequest): Promise<QueryPoolsResponse>;
  /** Queries a list of Showpool items. */
  Showpool(request: QueryShowpoolRequest): Promise<QueryShowpoolResponse>;
  /** Queries a PrefPool by id. */
  PrefPool(request: QueryGetPrefPoolRequest): Promise<QueryGetPrefPoolResponse>;
  /** Queries a list of PrefPool items. */
  PrefPoolAll(
    request: QueryAllPrefPoolRequest
  ): Promise<QueryAllPrefPoolResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("blue.blue.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Pools(request: QueryPoolsRequest): Promise<QueryPoolsResponse> {
    const data = QueryPoolsRequest.encode(request).finish();
    const promise = this.rpc.request("blue.blue.Query", "Pools", data);
    return promise.then((data) => QueryPoolsResponse.decode(new Reader(data)));
  }

  Showpool(request: QueryShowpoolRequest): Promise<QueryShowpoolResponse> {
    const data = QueryShowpoolRequest.encode(request).finish();
    const promise = this.rpc.request("blue.blue.Query", "Showpool", data);
    return promise.then((data) =>
      QueryShowpoolResponse.decode(new Reader(data))
    );
  }

  PrefPool(
    request: QueryGetPrefPoolRequest
  ): Promise<QueryGetPrefPoolResponse> {
    const data = QueryGetPrefPoolRequest.encode(request).finish();
    const promise = this.rpc.request("blue.blue.Query", "PrefPool", data);
    return promise.then((data) =>
      QueryGetPrefPoolResponse.decode(new Reader(data))
    );
  }

  PrefPoolAll(
    request: QueryAllPrefPoolRequest
  ): Promise<QueryAllPrefPoolResponse> {
    const data = QueryAllPrefPoolRequest.encode(request).finish();
    const promise = this.rpc.request("blue.blue.Query", "PrefPoolAll", data);
    return promise.then((data) =>
      QueryAllPrefPoolResponse.decode(new Reader(data))
    );
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
