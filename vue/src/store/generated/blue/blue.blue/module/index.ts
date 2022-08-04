// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreatePool } from "./types/blue/tx";
import { MsgVote } from "./types/blue/tx";
import { MsgAddQuestion } from "./types/blue/tx";
import { MsgCreatePrefPool } from "./types/blue/tx";


const types = [
  ["/blue.blue.MsgCreatePool", MsgCreatePool],
  ["/blue.blue.MsgVote", MsgVote],
  ["/blue.blue.MsgAddQuestion", MsgAddQuestion],
  ["/blue.blue.MsgCreatePrefPool", MsgCreatePrefPool],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCreatePool: (data: MsgCreatePool): EncodeObject => ({ typeUrl: "/blue.blue.MsgCreatePool", value: MsgCreatePool.fromPartial( data ) }),
    msgVote: (data: MsgVote): EncodeObject => ({ typeUrl: "/blue.blue.MsgVote", value: MsgVote.fromPartial( data ) }),
    msgAddQuestion: (data: MsgAddQuestion): EncodeObject => ({ typeUrl: "/blue.blue.MsgAddQuestion", value: MsgAddQuestion.fromPartial( data ) }),
    msgCreatePrefPool: (data: MsgCreatePrefPool): EncodeObject => ({ typeUrl: "/blue.blue.MsgCreatePrefPool", value: MsgCreatePrefPool.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
