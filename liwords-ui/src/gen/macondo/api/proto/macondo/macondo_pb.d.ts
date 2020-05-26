// package: macondo
// file: macondo/api/proto/macondo/macondo.proto

import * as jspb from "google-protobuf";

export class GameHistory extends jspb.Message {
  clearTurnsList(): void;
  getTurnsList(): Array<GameTurn>;
  setTurnsList(value: Array<GameTurn>): void;
  addTurns(value?: GameTurn, index?: number): GameTurn;

  clearPlayersList(): void;
  getPlayersList(): Array<PlayerInfo>;
  setPlayersList(value: Array<PlayerInfo>): void;
  addPlayers(value?: PlayerInfo, index?: number): PlayerInfo;

  getVersion(): number;
  setVersion(value: number): void;

  getOriginalGcg(): string;
  setOriginalGcg(value: string): void;

  getLexicon(): string;
  setLexicon(value: string): void;

  getIdAuth(): string;
  setIdAuth(value: string): void;

  getUid(): string;
  setUid(value: string): void;

  getTitle(): string;
  setTitle(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  clearLastKnownRacksList(): void;
  getLastKnownRacksList(): Array<string>;
  setLastKnownRacksList(value: Array<string>): void;
  addLastKnownRacks(value: string, index?: number): string;

  getFlipPlayers(): boolean;
  setFlipPlayers(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GameHistory.AsObject;
  static toObject(includeInstance: boolean, msg: GameHistory): GameHistory.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GameHistory, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GameHistory;
  static deserializeBinaryFromReader(message: GameHistory, reader: jspb.BinaryReader): GameHistory;
}

export namespace GameHistory {
  export type AsObject = {
    turnsList: Array<GameTurn.AsObject>,
    playersList: Array<PlayerInfo.AsObject>,
    version: number,
    originalGcg: string,
    lexicon: string,
    idAuth: string,
    uid: string,
    title: string,
    description: string,
    lastKnownRacksList: Array<string>,
    flipPlayers: boolean,
  }
}

export class GameEvent extends jspb.Message {
  getNickname(): string;
  setNickname(value: string): void;

  getNote(): string;
  setNote(value: string): void;

  getRack(): string;
  setRack(value: string): void;

  getType(): GameEvent.TypeMap[keyof GameEvent.TypeMap];
  setType(value: GameEvent.TypeMap[keyof GameEvent.TypeMap]): void;

  getCumulative(): number;
  setCumulative(value: number): void;

  getRow(): number;
  setRow(value: number): void;

  getColumn(): number;
  setColumn(value: number): void;

  getDirection(): GameEvent.DirectionMap[keyof GameEvent.DirectionMap];
  setDirection(value: GameEvent.DirectionMap[keyof GameEvent.DirectionMap]): void;

  getPosition(): string;
  setPosition(value: string): void;

  getPlayedTiles(): string;
  setPlayedTiles(value: string): void;

  getExchanged(): string;
  setExchanged(value: string): void;

  getScore(): number;
  setScore(value: number): void;

  getBonus(): number;
  setBonus(value: number): void;

  getEndRackPoints(): number;
  setEndRackPoints(value: number): void;

  getLostScore(): number;
  setLostScore(value: number): void;

  getIsBingo(): boolean;
  setIsBingo(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GameEvent.AsObject;
  static toObject(includeInstance: boolean, msg: GameEvent): GameEvent.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GameEvent, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GameEvent;
  static deserializeBinaryFromReader(message: GameEvent, reader: jspb.BinaryReader): GameEvent;
}

export namespace GameEvent {
  export type AsObject = {
    nickname: string,
    note: string,
    rack: string,
    type: GameEvent.TypeMap[keyof GameEvent.TypeMap],
    cumulative: number,
    row: number,
    column: number,
    direction: GameEvent.DirectionMap[keyof GameEvent.DirectionMap],
    position: string,
    playedTiles: string,
    exchanged: string,
    score: number,
    bonus: number,
    endRackPoints: number,
    lostScore: number,
    isBingo: boolean,
  }

  export interface TypeMap {
    TILE_PLACEMENT_MOVE: 0;
    PHONY_TILES_RETURNED: 1;
    PASS: 2;
    CHALLENGE_BONUS: 3;
    EXCHANGE: 4;
    END_RACK_PTS: 5;
    TIME_PENALTY: 6;
    END_RACK_PENALTY: 7;
    UNSUCCESSFUL_CHALLENGE_TURN_LOSS: 8;
  }

  export const Type: TypeMap;

  export interface DirectionMap {
    HORIZONTAL: 0;
    VERTICAL: 1;
  }

  export const Direction: DirectionMap;
}

export class GameTurn extends jspb.Message {
  clearEventsList(): void;
  getEventsList(): Array<GameEvent>;
  setEventsList(value: Array<GameEvent>): void;
  addEvents(value?: GameEvent, index?: number): GameEvent;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GameTurn.AsObject;
  static toObject(includeInstance: boolean, msg: GameTurn): GameTurn.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GameTurn, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GameTurn;
  static deserializeBinaryFromReader(message: GameTurn, reader: jspb.BinaryReader): GameTurn;
}

export namespace GameTurn {
  export type AsObject = {
    eventsList: Array<GameEvent.AsObject>,
  }
}

export class PlayerInfo extends jspb.Message {
  getNickname(): string;
  setNickname(value: string): void;

  getRealName(): string;
  setRealName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlayerInfo.AsObject;
  static toObject(includeInstance: boolean, msg: PlayerInfo): PlayerInfo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PlayerInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlayerInfo;
  static deserializeBinaryFromReader(message: PlayerInfo, reader: jspb.BinaryReader): PlayerInfo;
}

export namespace PlayerInfo {
  export type AsObject = {
    nickname: string,
    realName: string,
  }
}
